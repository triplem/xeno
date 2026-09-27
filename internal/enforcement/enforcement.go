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
	"io"
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
	// LastPushApproval is the host's stronger form of "not by the author": the most
	// recent push has to be approved by somebody who did not make it. GitHub already
	// refuses an approval from the pull request's own author, so this is the part of the
	// requirement a setting can add rather than the whole of it.
	LastPushApproval bool
	Reason           string
}

// Fetch asks the host what the branch is configured to require.
//
// A 403 is the answer this repository gets, and it is not an error: branch protection,
// rulesets and required reviewers exist on GitHub and none of them is available for a
// private repository on the free plan. Reporting that as a missing setting would send
// somebody looking for a checkbox that is not there.
func Fetch(client *http.Client, baseURL, repo, branch, token string) (Protection, error) {
	body, p, err := fetchProtection(client, baseURL, repo, branch, token)
	if err != nil || body == nil {
		return p, err
	}
	return decodeProtection(body)
}

// fetchProtection is the transport half: one request, and the status codes turned into the
// three answers a host can give. A nil body with no error means the host answered without
// one, which is every case except 200.
//
// The URL, the header and the meaning of each status are this host's. When a second host
// arrives they move into its adapter and this function goes with them, which is #97; the
// split is worth having either way, because it is what lets the decode be tested from a
// fixture and the status mapping from a server that returns nothing else.
func fetchProtection(client *http.Client, baseURL, repo, branch, token string) ([]byte, Protection, error) {
	var p Protection
	url := fmt.Sprintf("%s/repos/%s/branches/%s/protection", strings.TrimRight(baseURL, "/"), repo, branch)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, p, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, p, fmt.Errorf("asking %s: %w", url, err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusForbidden:
		p.Reason = "the host does not offer branch protection for this repository on its current plan"
		return nil, p, nil
	case http.StatusNotFound:
		p.Available, p.Reason = true, "the branch is not protected"
		return nil, p, nil
	case http.StatusUnauthorized:
		return nil, p, fmt.Errorf("the token was rejected by %s", url)
	default:
		return nil, p, fmt.Errorf("%s answered %s", url, resp.Status)
	}
	raw, err := io.ReadAll(resp.Body)
	return raw, p, err
}

// decodeProtection is the domain half: what this host's answer says about the requirements
// a project declares. The field names are the host's and the states are not.
func decodeProtection(raw []byte) (Protection, error) {
	var p Protection
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
	p.Available, p.Protected = true, true
	p.RequiredStatusChecks = body.RequiredStatusChecks != nil
	p.EnforceAdmins = body.EnforceAdmins != nil && body.EnforceAdmins.Enabled
	if body.Reviews != nil {
		p.ReviewsExpressible, p.RequiredApprovals = true, body.Reviews.Count
		p.LastPushApproval = body.Reviews.LastPushApproval
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

// requirement is one line of the report before it is evaluated: whether the declaration
// asks for it, and what the host's answer says about it. A table rather than four switches,
// because every one of them had the same three way shape, and the states differ only in
// their words.
type requirement struct {
	name     string
	declared func(Declared) (string, bool)
	// evaluate answers for a host that has protection to report. Unavailable is handled
	// once, above, since the reason is the same for every requirement.
	evaluate func(Declared, Protection) (actual string, met bool, note string)
	// waived is the reason this requirement may carry of its own, beyond the block's.
	waived func(Declared) string
	// expressible says whether the host reports on this requirement at all. Absent means
	// availability alone decides.
	expressible func(Protection) bool
}

var requirements = []requirement{
	{
		name:     "required_pipeline",
		declared: func(d Declared) (string, bool) { return "true", d.RequiredPipeline },
		evaluate: func(_ Declared, p Protection) (string, bool, string) {
			if !p.Protected {
				return "the branch is not protected", false, "a red gate then produces a report that a merge walks past"
			}
			if p.RequiredStatusChecks {
				return "required", true, ""
			}
			return "no status check is required", false, "a red gate then produces a report that a merge walks past"
		},
	},
	{
		name:     "allow_bypass",
		declared: func(d Declared) (string, bool) { return "false", !d.AllowBypass },
		evaluate: func(_ Declared, p Protection) (string, bool, string) {
			if p.EnforceAdmins {
				return "administrators are included", true, ""
			}
			return "administrators may bypass", false, ""
		},
	},
	{
		name: "approvals.required",
		declared: func(d Declared) (string, bool) {
			return fmt.Sprintf("%d", d.Approvals.Required), d.Approvals.Required > 0
		},
		evaluate: func(d Declared, p Protection) (string, bool, string) {
			return fmt.Sprintf("%d", p.RequiredApprovals), p.RequiredApprovals >= d.Approvals.Required, ""
		},
		waived:      func(d Declared) string { return d.Approvals.Waived },
		expressible: func(p Protection) bool { return p.ReviewsExpressible },
	},
	{
		// Section 13 says the approvals block covers whether the author may give one. What
		// satisfies it is not a setting: GitHub refuses an approval from a pull request's
		// own author, so an approval that exists is somebody else's. With none required
		// there is nothing to be somebody else's, which is unmet rather than met.
		name:     "approvals.not_by_author",
		declared: func(d Declared) (string, bool) { return "true", d.Approvals.NotByAuthor },
		evaluate: func(_ Declared, p Protection) (string, bool, string) {
			switch {
			case p.RequiredApprovals < 1:
				return "no approval is required", false, "an approval that does not exist is nobody's"
			case p.LastPushApproval:
				return "the author cannot approve, and the last push needs another approval", true, ""
			default:
				return "the host refuses an approval from the author", true,
					"require_last_push_approval would extend it to whoever pushed last"
			}
		},
		waived:      func(d Declared) string { return d.Approvals.Waived },
		expressible: func(p Protection) bool { return p.ReviewsExpressible },
	},
}

// Compare turns a declaration and an answer into the report.
func Compare(repo, branch string, d Declared, p Protection, now time.Time) Report {
	rep := Report{Repository: repo, Branch: branch, CheckedAt: now.UTC().Format(time.RFC3339)}
	for _, q := range requirements {
		declared, asked := q.declared(d)
		if !asked {
			continue
		}
		w := d.Waived
		if q.waived != nil && q.waived(d) != "" {
			w = q.waived(d)
		}
		expressible := p.Available
		if expressible && q.expressible != nil {
			expressible = q.expressible(p)
		}
		if !expressible {
			rep.add(q.name, declared, "not available", waiveOr(w, NotAvailable), noteOf(w, p.Reason))
			continue
		}
		actual, met, note := q.evaluate(d, p)
		if met {
			rep.add(q.name, declared, actual, Met, note)
			continue
		}
		rep.add(q.name, declared, actual, waiveOr(w, Unmet), noteOf(w, note))
	}

	// merge_method is declared and deliberately not compared. Section 13 says "not
	// checked" for it, and the plan says it is meaningful once a commit predicate is
	// active, which no gate in this runner has yet: a squash replaces the commits a
	// predicate judged, so the requirement exists for a consumer that does not. Comparing
	// it would be a check the specification says is not made, which is a spec change
	// first. The field is reported as unchecked rather than left silent, so that a project
	// declaring it learns that nothing reads it.
	if d.MergeMethod != "" {
		rep.add("merge_method", d.MergeMethod, "not compared", Unknown,
			"section 13 leaves it unchecked; it is meaningful once a commit predicate is active")
	}
	return rep
}

func (r *Report) add(name, declared, actual string, s State, note string) {
	r.Requirements = append(r.Requirements,
		Requirement{Name: name, Declared: declared, Actual: actual, State: s, Note: note})
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
