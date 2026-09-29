// SPDX-License-Identifier: Apache-2.0

// Package github answers what a GitHub branch is configured to require.
//
// Everything GitHub's is here and nowhere else: the URL shape, the Accept header, the JSON
// field names, and the meaning of a 403. #97 enumerated those four and named the last as
// the one that makes this more than a URL swap — a 403 meaning "not available on this plan"
// is a fact about one host's pricing, and a second host answers the same question with a
// different status or with none.
package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/triplem/xeno/internal/enforcement"
)

// Adapter reads branch protection from one GitHub deployment. The base URL is a parameter
// rather than a constant because a self managed deployment has no canonical address, and
// because a tool that knows its host's address by heart is a tool that only ever had one.
type Adapter struct {
	client  *http.Client
	baseURL string
}

// New returns the adapter. A nil client gets one with the timeout the runner used, so that
// a caller which has no opinion about transports does not have to invent one.
func New(c *http.Client, baseURL string) *Adapter {
	if c == nil {
		c = &http.Client{Timeout: 20 * time.Second}
	}
	return &Adapter{client: c, baseURL: baseURL}
}

// protection is what this host's answer says, before it is turned into requirements. It is
// this package's own shape and deliberately not shared: it was enforcement.Protection until
// A65, and what that row rejected was exactly the idea that a second host could fill it.
type protection struct {
	available            bool
	protected            bool
	requiredStatusChecks bool
	enforceAdmins        bool
	requiredApprovals    int
	// reviewsExpressible says whether the host reports a review requirement at all. Not
	// the same as none being configured.
	reviewsExpressible bool
	// lastPushApproval is the host's stronger form of "not by the author": the most recent
	// push has to be approved by somebody who did not make it. GitHub already refuses an
	// approval from the pull request's own author, so this is the part of the requirement a
	// setting can add rather than the whole of it.
	lastPushApproval bool
	reason           string
}

// Requirements implements host.BranchRules.
//
// Every requirement this host can speak to is answered, whether or not the declaration asks
// for it: which ones were asked for is the domain's filter, and an adapter that guessed
// would be applying the declaration twice.
//
// Where the host cannot express a requirement the answer is not-available carrying the
// reason, rather than an omission. Both read as not-available in the report, and the
// difference is that a reason reaches the person who has to go looking for a setting — which
// is the whole point of telling a setting nobody made apart from one the host does not have.
func (a *Adapter) Requirements(repo, branch, token string, d enforcement.Declared) (
	[]enforcement.Requirement, error) {
	p, err := a.fetch(repo, branch, token)
	if err != nil {
		return nil, err
	}
	if !p.available {
		return unavailable(p.reason), nil
	}
	out := []enforcement.Requirement{
		a.requiredPipeline(p),
		a.allowBypass(p),
	}
	if !p.reviewsExpressible {
		// The host reports no review requirement at all, which is not the same as none
		// being configured. On an unprotected branch this is why: there is no protection
		// to carry a review rule.
		return append(out, notAvailable(enforcement.NameApprovalsRequired, p.reason),
			notAvailable(enforcement.NameApprovalsNotByAuthor, p.reason)), nil
	}
	return append(out, a.approvalsRequired(p, d), a.approvalsNotByAuthor(p)), nil
}

// unavailable is the answer where the host offers nothing about this branch at all. Each
// requirement is named, so the report says which ones went unanswered and why, instead of
// leaving the domain to infer an absence.
func unavailable(reason string) []enforcement.Requirement {
	names := []string{
		enforcement.NameRequiredPipeline, enforcement.NameAllowBypass,
		enforcement.NameApprovalsRequired, enforcement.NameApprovalsNotByAuthor,
	}
	out := make([]enforcement.Requirement, 0, len(names))
	for _, n := range names {
		out = append(out, notAvailable(n, reason))
	}
	return out
}

func notAvailable(name, reason string) enforcement.Requirement {
	return enforcement.Requirement{
		Name: name, Actual: "not available", State: enforcement.NotAvailable, Note: reason,
	}
}

func (a *Adapter) requiredPipeline(p protection) enforcement.Requirement {
	q := enforcement.Requirement{Name: enforcement.NameRequiredPipeline}
	switch {
	case !p.protected:
		q.Actual, q.State = "the branch is not protected", enforcement.Unmet
		q.Note = "a red gate then produces a report that a merge walks past"
	case p.requiredStatusChecks:
		q.Actual, q.State = "required", enforcement.Met
	default:
		q.Actual, q.State = "no status check is required", enforcement.Unmet
		q.Note = "a red gate then produces a report that a merge walks past"
	}
	return q
}

func (a *Adapter) allowBypass(p protection) enforcement.Requirement {
	q := enforcement.Requirement{Name: enforcement.NameAllowBypass}
	if p.enforceAdmins {
		q.Actual, q.State = "administrators are included", enforcement.Met
		return q
	}
	q.Actual, q.State = "administrators may bypass", enforcement.Unmet
	return q
}

func (a *Adapter) approvalsRequired(p protection, d enforcement.Declared) enforcement.Requirement {
	q := enforcement.Requirement{
		Name:   enforcement.NameApprovalsRequired,
		Actual: fmt.Sprintf("%d", p.requiredApprovals),
		State:  enforcement.Unmet,
	}
	if p.requiredApprovals >= d.Approvals.Required {
		q.State = enforcement.Met
	}
	return q
}

// approvalsNotByAuthor is the row #104's table calls the interesting one. Here the
// requirement is met by the host's own rule: GitHub refuses an approval from a pull
// request's own author, so an approval that exists is somebody else's. With none required
// there is nothing to be somebody else's, which is unmet rather than met. Section 13 says
// the approvals block covers whether the author may give one.
func (a *Adapter) approvalsNotByAuthor(p protection) enforcement.Requirement {
	q := enforcement.Requirement{Name: enforcement.NameApprovalsNotByAuthor}
	switch {
	case p.requiredApprovals < 1:
		q.Actual, q.State = "no approval is required", enforcement.Unmet
		q.Note = "an approval that does not exist is nobody's"
	case p.lastPushApproval:
		q.Actual = "the author cannot approve, and the last push needs another approval"
		q.State = enforcement.Met
	default:
		q.Actual = "the host refuses an approval from the author"
		q.State = enforcement.Met
		q.Note = "require_last_push_approval would extend it to whoever pushed last"
	}
	return q
}

// fetch asks the host what the branch is configured to require.
//
// A 403 is the answer this repository gets on a free private plan, and it is not an error:
// branch protection, rulesets and required reviewers exist on GitHub and none of them is
// available there. Reporting that as a missing setting would send somebody looking for a
// checkbox that is not there.
func (a *Adapter) fetch(repo, branch, token string) (protection, error) {
	body, p, err := a.transport(repo, branch, token)
	if err != nil || body == nil {
		return p, err
	}
	return decode(body)
}

// transport is one request, and the status codes turned into the three answers a host can
// give. A nil body with no error means the host answered without one, which is every case
// except 200. The URL, the header and the meaning of each status are this host's, which is
// why they are in this package and not in the domain.
func (a *Adapter) transport(repo, branch, token string) ([]byte, protection, error) {
	var p protection
	url := fmt.Sprintf("%s/repos/%s/branches/%s/protection",
		strings.TrimRight(a.baseURL, "/"), repo, branch)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, p, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, p, fmt.Errorf("asking %s: %w", url, err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusForbidden:
		p.reason = "the host does not offer branch protection for this repository on its current plan"
		return nil, p, nil
	case http.StatusNotFound:
		p.available, p.reason = true, "the branch is not protected"
		return nil, p, nil
	case http.StatusUnauthorized:
		return nil, p, fmt.Errorf("the token was rejected by %s", url)
	default:
		return nil, p, fmt.Errorf("%s answered %s", url, resp.Status)
	}
	raw, err := io.ReadAll(resp.Body)
	return raw, p, err
}

// decode reads this host's field names. The names are the host's and the states are not.
func decode(raw []byte) (protection, error) {
	var p protection
	var body struct {
		RequiredStatusChecks *struct {
			Contexts []string `json:"contexts"`
		} `json:"required_status_checks"`
		EnforceAdmins *struct {
			Enabled bool `json:"enabled"`
		} `json:"enforce_admins"`
		Reviews *struct {
			Count            int  `json:"required_approving_review_count"`
			LastPushApproval bool `json:"require_last_push_approval"`
		} `json:"required_pull_request_reviews"`
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return p, err
	}
	p.available, p.protected = true, true
	p.requiredStatusChecks = body.RequiredStatusChecks != nil
	p.enforceAdmins = body.EnforceAdmins != nil && body.EnforceAdmins.Enabled
	if body.Reviews != nil {
		p.reviewsExpressible, p.requiredApprovals = true, body.Reviews.Count
		p.lastPushApproval = body.Reviews.LastPushApproval
	}
	return p, nil
}
