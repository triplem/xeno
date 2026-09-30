// SPDX-License-Identifier: Apache-2.0

// Package gitlab answers what a GitLab branch is configured to require.
//
// It is the second adapter behind host.BranchRules and the reason A65 decided the
// vocabulary the way it did: this host answers allow_bypass with three lists of grants
// where the other answers with one boolean, and reducing grants to "somebody may bypass"
// is a judgement about this permission model.
//
// It shares no code with the other adapter on purpose. The difference is structural and
// not cosmetic: that one answers every requirement from a single request, so one status
// code decides the whole report, while this one needs three and a tier limit on the
// approvals endpoints must not stop the other requirements being answered.
package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/triplem/xeno/internal/enforcement"
)

// GitLab's access levels, from its permissions reference. Only the boundary matters here:
// noAccess is "no one", and any grant above it is somebody who holds the permission.
const (
	noAccess = 0
)

// Adapter reads branch and project settings from one GitLab deployment. The base URL is a
// parameter because a self managed deployment has no canonical address.
type Adapter struct {
	client  *http.Client
	baseURL string
}

// New returns the adapter. A nil client gets one with the timeout the runner uses.
func New(c *http.Client, baseURL string) *Adapter {
	if c == nil {
		c = &http.Client{Timeout: 20 * time.Second}
	}
	return &Adapter{client: c, baseURL: baseURL}
}

// accessLevel is one grant on a protected branch: a level, and whoever holds it.
type accessLevel struct {
	AccessLevel int `json:"access_level"`
}

// protectedBranch is this host's answer about one branch.
type protectedBranch struct {
	Name                  string        `json:"name"`
	PushAccessLevels      []accessLevel `json:"push_access_levels"`
	MergeAccessLevels     []accessLevel `json:"merge_access_levels"`
	UnprotectAccessLevels []accessLevel `json:"unprotect_access_levels"`
	AllowForcePush        bool          `json:"allow_force_push"`
}

// project is the part of the project payload that answers a declared requirement. The
// pipeline setting is a project setting on this host and a branch setting on the other,
// which is one of the places the two hosts stop resembling each other.
//
// The setting is a pointer because this host omits it rather than sending false where the
// caller may not read the project's settings. Decoded into a bool, an omission becomes off,
// and a setting nobody could read is reported as a setting nobody made — which is the one
// collapse the report exists to prevent. Observed against gitlab.com, where an
// unauthenticated read of a public project carries no such field.
type project struct {
	OnlyAllowMergeIfPipelineSucceeds *bool `json:"only_allow_merge_if_pipeline_succeeds"`
}

// approvalSettings is the project wide approval configuration. Premium, so the request for
// it is the one most likely to be refused, and a refusal is the tier rather than a fault.
type approvalSettings struct {
	MergeRequestsAuthorApproval bool `json:"merge_requests_author_approval"`
}

// approvalRule is one rule of the project's approval configuration. Premium.
type approvalRule struct {
	ApprovalsRequired int `json:"approvals_required"`
}

// Requirements implements host.BranchRules.
//
// Each requirement is answered from the request that can answer it, and a request that
// could not be read degrades its own requirements to not-available with the reason. That is
// why this is not one fetch and a switch: a project on a tier without approval rules still
// has a protected branch worth reporting on.
func (a *Adapter) Requirements(repo, branch, token string, d enforcement.Declared) (
	[]enforcement.Requirement, error) {
	var out []enforcement.Requirement

	pb, pbWhy, err := a.branch(repo, branch, token)
	if err != nil {
		return nil, err
	}
	pr, prWhy, err := a.project(repo, token)
	if err != nil {
		return nil, err
	}

	switch {
	case prWhy != "":
		out = append(out, notAvailable(enforcement.NameRequiredPipeline, prWhy))
	case pr.OnlyAllowMergeIfPipelineSucceeds == nil:
		out = append(out, notAvailable(enforcement.NameRequiredPipeline,
			"the project payload carries no only_allow_merge_if_pipeline_succeeds, "+
				"so the token cannot read the project's merge settings"))
	default:
		out = append(out, requiredPipeline(pr))
	}

	switch {
	case pbWhy != "":
		out = append(out, notAvailable(enforcement.NameAllowBypass, pbWhy))
	default:
		out = append(out, allowBypass(pb))
	}

	approvals, err := a.approvals(repo, token, d)
	if err != nil {
		return nil, err
	}
	return append(out, approvals...), nil
}

// requiredPipeline reads the project setting. On this host it is not a property of the
// branch at all, which is a closer fit to what the declaration means than the other host's
// required status check: it says no merge request merges on a failed pipeline, rather than
// naming a check by name.
func requiredPipeline(pr project) enforcement.Requirement {
	q := enforcement.Requirement{Name: enforcement.NameRequiredPipeline}
	if *pr.OnlyAllowMergeIfPipelineSucceeds {
		q.Actual, q.State = "only_allow_merge_if_pipeline_succeeds is on", enforcement.Met
		return q
	}
	q.Actual, q.State = "only_allow_merge_if_pipeline_succeeds is off", enforcement.Unmet
	q.Note = "a red gate then produces a report that a merge walks past"
	return q
}

// allowBypass is the row #104's table calls the one where this host stops resembling the
// other, and A65's deciding case. There is no enforce_admins here. What there is, is who
// may push to the branch without a merge request, whether the judged history may be
// replaced, and who may remove the protection altogether.
//
// The largest finding is reported rather than all of them or a count. Unprotecting the
// branch subsumes the other two, and a count would imply that somebody who can unprotect
// and somebody who can push are two of the same thing.
func allowBypass(pb protectedBranch) enforcement.Requirement {
	q := enforcement.Requirement{Name: enforcement.NameAllowBypass, State: enforcement.Unmet}
	switch {
	case grants(pb.UnprotectAccessLevels) > 0:
		q.Actual = fmt.Sprintf("%d grant(s) may unprotect the branch",
			grants(pb.UnprotectAccessLevels))
		q.Note = "unprotecting removes every other requirement on this branch with it"
	case pb.AllowForcePush:
		q.Actual = "the branch allows force push"
		q.Note = "the history a gate judged can be replaced after it was judged"
	case grants(pb.PushAccessLevels) > 0:
		q.Actual = fmt.Sprintf("%d grant(s) may push without a merge request",
			grants(pb.PushAccessLevels))
		q.Note = "a push that is not a merge request is not gated"
	default:
		q.Actual, q.State = "no push, force push or unprotect grant", enforcement.Met
	}
	return q
}

// grants counts the grants above "no one". A protected branch with an explicit level 0
// entry is expressing that nobody holds the permission, which is not a bypass.
func grants(ls []accessLevel) int {
	n := 0
	for _, l := range ls {
		if l.AccessLevel > noAccess {
			n++
		}
	}
	return n
}

// approvals answers the two rows that need Premium. Both come from requests that a tier
// without approval rules refuses, and a refusal is reported as not-available carrying the
// reason: a setting the tier does not have is nobody's oversight.
func (a *Adapter) approvals(repo, token string, d enforcement.Declared) (
	[]enforcement.Requirement, error) {
	set, setWhy, err := a.approvalSettings(repo, token)
	if err != nil {
		return nil, err
	}
	rules, rulesWhy, err := a.approvalRules(repo, token)
	if err != nil {
		return nil, err
	}

	var out []enforcement.Requirement
	if rulesWhy != "" {
		out = append(out, notAvailable(enforcement.NameApprovalsRequired, rulesWhy))
	} else {
		// The rules are separate requirements with their own eligible approvers, so the
		// count a merge request needs is their sum. Observed on a real project: two rules,
		// one requiring none and one requiring one, which is one approval.
		required := 0
		for _, r := range rules {
			required += r.ApprovalsRequired
		}
		q := enforcement.Requirement{
			Name:   enforcement.NameApprovalsRequired,
			Actual: fmt.Sprintf("%d", required),
			State:  enforcement.Unmet,
		}
		if required >= d.Approvals.Required {
			q.State = enforcement.Met
		}
		out = append(out, q)
	}

	if setWhy != "" {
		return append(out, notAvailable(enforcement.NameApprovalsNotByAuthor, setWhy)), nil
	}
	// Unlike the other host, this one does not refuse an author's own approval by itself:
	// merge_requests_author_approval is a setting, so the requirement has a reachable unmet
	// state here that it does not have there. A65 names this as the second case that decided
	// the vocabulary, and it is the reason the words differ rather than only the values.
	q := enforcement.Requirement{Name: enforcement.NameApprovalsNotByAuthor}
	if set.MergeRequestsAuthorApproval {
		q.Actual, q.State = "merge_requests_author_approval is on", enforcement.Unmet
		q.Note = "the author of a merge request may approve it"
		return append(out, q), nil
	}
	q.Actual, q.State = "merge_requests_author_approval is off", enforcement.Met
	return append(out, q), nil
}

func notAvailable(name, reason string) enforcement.Requirement {
	return enforcement.Requirement{
		Name: name, Actual: "not available", State: enforcement.NotAvailable, Note: reason,
	}
}

// branch asks about one protected branch. A 404 here is the branch not being protected,
// which is an answer about the branch and not a failure to read it.
func (a *Adapter) branch(repo, branch, token string) (protectedBranch, string, error) {
	var pb protectedBranch
	raw, status, err := a.get(
		fmt.Sprintf("projects/%s/protected_branches/%s", ident(repo), url.PathEscape(branch)),
		token)
	switch {
	case err != nil:
		return pb, "", err
	case status == http.StatusNotFound:
		return pb, "the branch is not protected", nil
	case status != http.StatusOK:
		return pb, unreadable("the protected branch", status), nil
	}
	return pb, "", json.Unmarshal(raw, &pb)
}

// project asks about the project itself. A 404 here is not an answer: a project that cannot
// be found cannot be reported on, and saying one requirement is unavailable would leave
// somebody looking for a setting in a project that is not there or not theirs to read.
func (a *Adapter) project(repo, token string) (project, string, error) {
	var pr project
	raw, status, err := a.get("projects/"+ident(repo), token)
	switch {
	case err != nil:
		return pr, "", err
	case status == http.StatusNotFound:
		return pr, "", fmt.Errorf(
			"no project %q at %s, or the token cannot read it", repo, a.baseURL)
	case status != http.StatusOK:
		return pr, unreadable("the project", status), nil
	}
	return pr, "", json.Unmarshal(raw, &pr)
}

func (a *Adapter) approvalSettings(repo, token string) (approvalSettings, string, error) {
	var s approvalSettings
	raw, status, err := a.get("projects/"+ident(repo)+"/approvals", token)
	if err != nil {
		return s, "", err
	}
	if status != http.StatusOK {
		return s, unreadable("the approval settings", status), nil
	}
	return s, "", json.Unmarshal(raw, &s)
}

func (a *Adapter) approvalRules(repo, token string) ([]approvalRule, string, error) {
	raw, status, err := a.get("projects/"+ident(repo)+"/approval_rules", token)
	if err != nil {
		return nil, "", err
	}
	if status != http.StatusOK {
		return nil, unreadable("the approval rules", status), nil
	}
	var rs []approvalRule
	return rs, "", json.Unmarshal(raw, &rs)
}

// unreadable is the reason a requirement goes unanswered. It names the status and stops
// there, because what a status means here was observed to be ambiguous: gitlab.com answers
// 403 on the approval settings both where the tier has no approval rules and where the token
// simply is not a member with enough access, and those are different problems with the same
// code. An earlier version of this sentence asserted the tier, which is wrong in the second
// case and sends somebody to a billing page over a permission.
//
// What the sentence can say truthfully is the one thing the report needs: this is not a
// setting nobody made. Which of the three it is belongs to whoever reads the status.
func unreadable(what string, status int) string {
	return fmt.Sprintf("%s answered %d; that is the tier, the token's access or the project, "+
		"and not a setting nobody made", what, status)
}

// get is one request. It returns the body and the status, and an error only for the cases
// where nothing about this project can be read.
//
// A rejected token is such a case: a token the host will not accept cannot answer anything,
// and reporting five requirements as not available would hide a misconfiguration behind the
// tier.
func (a *Adapter) get(path, token string) ([]byte, int, error) {
	u := fmt.Sprintf("%s/%s", strings.TrimRight(a.baseURL, "/"), path)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Accept", "application/json")
	if token != "" {
		req.Header.Set("PRIVATE-TOKEN", token)
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("asking %s: %w", u, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, fmt.Errorf("the token was rejected by %s", u)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, nil
	}
	raw, err := io.ReadAll(resp.Body)
	return raw, resp.StatusCode, err
}

// ident is this host's project identity: a namespaced path, URL encoded, so that the slash
// in it does not read as a path separator.
func ident(repo string) string { return url.PathEscape(repo) }
