// SPDX-License-Identifier: Apache-2.0

// Package enforcement compares what a project declares it requires of its host against
// what the host is configured to do.
//
// It is the only package here that uses the network, and it is deliberately not
// reachable from the gate path: a verdict has to be reproducible from the repository
// alone, and a gate that called out would make it depend on whether a service answered.
// Nothing under internal/gates imports this.
package enforcement

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// State is what could be learned about one requirement. The three are different and
// collapsing them is what makes a report worthless: a setting nobody made is somebody's
// oversight, a setting the host does not have is nobody's, and a waived one is a
// decision already taken.
type State string

const (
	Met          State = "met"
	Unmet        State = "unmet"
	NotAvailable State = "not-available"
	Waived       State = "waived"
	Unknown      State = "unknown"
)

// Requirement is one line of the report.
type Requirement struct {
	Name     string `yaml:"name" json:"name"`
	Declared string `yaml:"declared" json:"declared"`
	Actual   string `yaml:"actual" json:"actual"`
	State    State  `yaml:"state" json:"state"`
	Note     string `yaml:"note,omitempty" json:"note,omitempty"`
}

// Report is written as a pipeline artifact and committed by nothing. CI verifies
// artifacts that exist and produces none, and a second place carrying a verdict beside
// gate.yaml is what the schema forbids.
type Report struct {
	Repository   string        `yaml:"repository"`
	Branch       string        `yaml:"branch"`
	CheckedAt    string        `yaml:"checked_at"`
	Requirements []Requirement `yaml:"requirements"`
}

// Unmet counts the requirements that are neither met nor excused. A waived one is
// excused by a decision and a not-available one that is not waived is not: that is the
// case a project has to look at once and then record.
func (r Report) Unmet() int {
	n := 0
	for _, q := range r.Requirements {
		if q.State == Unmet || q.State == NotAvailable {
			n++
		}
	}
	return n
}

// Protection is the part of the host's answer this compares against.
type Protection struct {
	Available            bool
	Protected            bool
	RequiredStatusChecks bool
	EnforceAdmins        bool
	RequiredApprovals    int
	// ReviewsExpressible says whether the host reports a review requirement at all.
	// Not the same as none being configured.
	ReviewsExpressible bool
	Reason             string
}

// Fetch asks the host what the branch is configured to require.
//
// A 403 is the answer this repository gets, and it is not an error: branch protection,
// rulesets and required reviewers exist on GitHub and none of them is available for a
// private repository on the free plan. Reporting that as a missing setting would send
// somebody looking for a checkbox that is not there.
func Fetch(client *http.Client, baseURL, repo, branch, token string) (Protection, error) {
	var p Protection
	url := fmt.Sprintf("%s/repos/%s/branches/%s/protection", strings.TrimRight(baseURL, "/"), repo, branch)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return p, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return p, fmt.Errorf("asking %s: %w", url, err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusForbidden:
		p.Reason = "the host does not offer branch protection for this repository on its current plan"
		return p, nil
	case http.StatusNotFound:
		p.Available, p.Reason = true, "the branch is not protected"
		return p, nil
	case http.StatusUnauthorized:
		return p, fmt.Errorf("the token was rejected by %s", url)
	default:
		return p, fmt.Errorf("%s answered %s", url, resp.Status)
	}

	var body struct {
		RequiredStatusChecks *struct {
			Contexts []string `json:"contexts"`
		} `json:"required_status_checks"`
		EnforceAdmins *struct {
			Enabled bool `json:"enabled"`
		} `json:"enforce_admins"`
		Reviews *struct {
			Count int `json:"required_approving_review_count"`
		} `json:"required_pull_request_reviews"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return p, err
	}
	p.Available, p.Protected = true, true
	p.RequiredStatusChecks = body.RequiredStatusChecks != nil
	p.EnforceAdmins = body.EnforceAdmins != nil && body.EnforceAdmins.Enabled
	if body.Reviews != nil {
		p.ReviewsExpressible, p.RequiredApprovals = true, body.Reviews.Count
	}
	return p, nil
}

// Declared is the enforcement block of project.yaml.
type Declared struct {
	RequiredPipeline bool   `yaml:"required_pipeline"`
	AllowBypass      bool   `yaml:"allow_bypass"`
	MergeMethod      string `yaml:"merge_method,omitempty"`
	Approvals        struct {
		Required    int    `yaml:"required"`
		NotByAuthor bool   `yaml:"not_by_author"`
		Waived      string `yaml:"waived,omitempty"`
	} `yaml:"approvals"`
	// Waived excuses the block as a whole where a host cannot express any of it.
	Waived string `yaml:"waived,omitempty"`
}

// Compare turns a declaration and an answer into the report.
func Compare(repo, branch string, d Declared, p Protection, now time.Time) Report {
	rep := Report{Repository: repo, Branch: branch, CheckedAt: now.UTC().Format(time.RFC3339)}
	add := func(name, declared, actual string, s State, note string) {
		rep.Requirements = append(rep.Requirements,
			Requirement{Name: name, Declared: declared, Actual: actual, State: s, Note: note})
	}

	if d.RequiredPipeline {
		switch {
		case !p.Available:
			add("required_pipeline", "true", "not available", waiveOr(d.Waived, NotAvailable), noteOf(d.Waived, p.Reason))
		case !p.Protected:
			add("required_pipeline", "true", "the branch is not protected",
				waiveOr(d.Waived, Unmet), noteOf(d.Waived, "a red gate then produces a report that a merge walks past"))
		case p.RequiredStatusChecks:
			add("required_pipeline", "true", "required", Met, "")
		default:
			add("required_pipeline", "true", "no status check is required",
				waiveOr(d.Waived, Unmet), noteOf(d.Waived, "a red gate then produces a report that a merge walks past"))
		}
	}

	if !d.AllowBypass {
		switch {
		case !p.Available:
			add("allow_bypass", "false", "not available", waiveOr(d.Waived, NotAvailable), noteOf(d.Waived, p.Reason))
		case p.EnforceAdmins:
			add("allow_bypass", "false", "administrators are included", Met, "")
		default:
			add("allow_bypass", "false", "administrators may bypass", waiveOr(d.Waived, Unmet), d.Waived)
		}
	}

	if d.Approvals.Required > 0 {
		declared := fmt.Sprintf("%d", d.Approvals.Required)
		w := d.Approvals.Waived
		if w == "" {
			w = d.Waived
		}
		switch {
		case !p.Available || !p.ReviewsExpressible:
			add("approvals.required", declared, "not available", waiveOr(w, NotAvailable), noteOf(w, p.Reason))
		case p.RequiredApprovals >= d.Approvals.Required:
			add("approvals.required", declared, fmt.Sprintf("%d", p.RequiredApprovals), Met, "")
		default:
			add("approvals.required", declared, fmt.Sprintf("%d", p.RequiredApprovals), waiveOr(w, Unmet), w)
		}
	}
	return rep
}

// waiveOr is where waived does its work. Without it a requirement the host cannot
// express is reported as unmet on every run forever, and a report that always says the
// same thing is ignored within a fortnight.
func waiveOr(waived string, otherwise State) State {
	if strings.TrimSpace(waived) != "" {
		return Waived
	}
	return otherwise
}

func noteOf(waived, reason string) string {
	if strings.TrimSpace(waived) != "" {
		return waived
	}
	return reason
}

// Token reads the enforcement token from the environment. It is never in the
// repository, which is the same rule the tracker credential follows.
func Token() string {
	for _, k := range []string{"XENO_ENFORCEMENT_TOKEN", "GITHUB_TOKEN"} {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return ""
}
