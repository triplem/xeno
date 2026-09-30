// SPDX-License-Identifier: Apache-2.0

package gitlab

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/enforcement"
)

func declared() enforcement.Declared {
	var d enforcement.Declared
	d.RequiredPipeline, d.AllowBypass = true, false
	d.Approvals.Required, d.Approvals.NotByAuthor = 1, true
	return d
}

func stateOf(qs []enforcement.Requirement, name string) (enforcement.State, string, string) {
	for _, q := range qs {
		if q.Name == name {
			return q.State, q.Actual, q.Note
		}
	}
	return "", "", ""
}

// host answers each endpoint from a map of path suffix to handler, so a test says what this
// host said and nothing about how the adapter asks.
type host struct {
	branch    any
	project   any
	approvals any
	rules     any
	status    map[string]int
}

func (h host) serve(t *testing.T) *httptest.Server {
	t.Helper()
	write := func(w http.ResponseWriter, key string, body any) {
		if code, ok := h.status[key]; ok {
			w.WriteHeader(code)
			return
		}
		if body == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(body)
	}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/protected_branches/"):
			write(w, "branch", h.branch)
		case strings.HasSuffix(r.URL.Path, "/approval_rules"):
			write(w, "rules", h.rules)
		case strings.HasSuffix(r.URL.Path, "/approvals"):
			write(w, "approvals", h.approvals)
		default:
			write(w, "project", h.project)
		}
	}))
}

func (h host) ask(t *testing.T) []enforcement.Requirement {
	t.Helper()
	srv := h.serve(t)
	defer srv.Close()
	qs, err := New(srv.Client(), srv.URL).Requirements("group/proj", "main", "t", declared())
	if err != nil {
		t.Fatal(err)
	}
	return qs
}

func grant(level int) map[string]any { return map[string]any{"access_level": level} }

// A65's deciding case. This host has no enforce_admins: it has three lists of grants and a
// force push flag, and which of them was found has to reach the report, because
// "administrators are included" is the other host's sentence and would be false here.
func TestAllowBypassNamesWhichBypassItFound(t *testing.T) {
	base := map[string]any{"name": "main"}

	for _, tc := range []struct {
		name   string
		branch map[string]any
		state  enforcement.State
		says   string
	}{
		{
			name:   "an unprotect grant is the largest and is reported first",
			branch: map[string]any{"name": "main", "unprotect_access_levels": []any{grant(40)}, "allow_force_push": true, "push_access_levels": []any{grant(40)}},
			state:  enforcement.Unmet, says: "unprotect",
		},
		{
			name:   "force push is reported above a push grant",
			branch: map[string]any{"name": "main", "allow_force_push": true, "push_access_levels": []any{grant(40)}},
			state:  enforcement.Unmet, says: "force push",
		},
		{
			name:   "a push grant is a bypass of the merge request",
			branch: map[string]any{"name": "main", "push_access_levels": []any{grant(40), grant(30)}},
			state:  enforcement.Unmet, says: "push without a merge request",
		},
		{
			name:   "none of the three is met",
			branch: base,
			state:  enforcement.Met, says: "no push, force push or unprotect grant",
		},
		{
			// Level 0 is "no one" in this host's permission model, so an explicit entry at
			// 0 expresses that nobody holds the permission. Counting it would report a
			// bypass held by nobody.
			name:   "a grant at no access is not a bypass",
			branch: map[string]any{"name": "main", "push_access_levels": []any{grant(0)}},
			state:  enforcement.Met, says: "no push, force push or unprotect grant",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			qs := host{branch: tc.branch, project: map[string]any{}, approvals: map[string]any{}, rules: []any{}}.ask(t)
			s, actual, _ := stateOf(qs, "allow_bypass")
			if s != tc.state {
				t.Errorf("state is %q, want %q", s, tc.state)
			}
			if !strings.Contains(actual, tc.says) {
				t.Errorf("it says %q, want it to name %q", actual, tc.says)
			}
			if strings.Contains(actual, "administrator") {
				t.Errorf("it used the other host's words: %q", actual)
			}
		})
	}
}

// On this host the pipeline requirement is a project setting and not a property of the
// branch, which is one of the places the two hosts stop resembling each other.
func TestRequiredPipelineReadsTheProjectSetting(t *testing.T) {
	on := host{branch: map[string]any{"name": "main"},
		project:   map[string]any{"only_allow_merge_if_pipeline_succeeds": true},
		approvals: map[string]any{}, rules: []any{}}.ask(t)
	if s, _, _ := stateOf(on, "required_pipeline"); s != enforcement.Met {
		t.Errorf("with the setting on the state is %q, want %q", s, enforcement.Met)
	}

	off := host{branch: map[string]any{"name": "main"},
		project:   map[string]any{"only_allow_merge_if_pipeline_succeeds": false},
		approvals: map[string]any{}, rules: []any{}}.ask(t)
	if s, _, note := stateOf(off, "required_pipeline"); s != enforcement.Unmet || note == "" {
		t.Errorf("with the setting off the state is %q with note %q, want unmet and a reason", s, note)
	}
}

// The tier, which is the distinction the whole report rests on. A setting a tier does not
// have is nobody's oversight, so it is not available and never unmet — otherwise every Free
// project's report is red forever, which is the standing complaint the waive logic exists to
// prevent.
func TestATierWithoutApprovalRulesIsNotAvailableRatherThanUnmet(t *testing.T) {
	for _, status := range []int{http.StatusForbidden, http.StatusNotFound, http.StatusPaymentRequired} {
		qs := host{branch: map[string]any{"name": "main"}, project: map[string]any{},
			status: map[string]int{"approvals": status, "rules": status}}.ask(t)

		for _, name := range []string{"approvals.required", "approvals.not_by_author"} {
			s, _, note := stateOf(qs, name)
			if strings.Contains(note, "on a tier") {
				t.Errorf("%d: %s blames the tier for a status that is also a permission: %q", status, name, note)
			}
			if s != enforcement.NotAvailable {
				t.Errorf("%d: %s is %q, want %q", status, name, s, enforcement.NotAvailable)
			}
			if note == "" {
				t.Errorf("%d: %s says nothing about why", status, name)
			}
		}
		// The other three are answered anyway. A tier limit on one endpoint must not cost
		// the requirements another endpoint could answer, which is why this adapter makes
		// three requests rather than one.
		if s, _, _ := stateOf(qs, "allow_bypass"); s == "" || s == enforcement.NotAvailable {
			t.Errorf("%d: a tier limit on approvals lost allow_bypass (%q)", status, s)
		}
	}
}

// On a tier that has them, the two rows are compared. This host makes not_by_author a
// setting, so the requirement has a reachable unmet state that the other host's does not,
// which is A65's second deciding case.
func TestApprovalsAreComparedWhereTheTierHasThem(t *testing.T) {
	met := host{branch: map[string]any{"name": "main"}, project: map[string]any{},
		approvals: map[string]any{"merge_requests_author_approval": false},
		rules:     []any{map[string]any{"approvals_required": 2}}}.ask(t)
	if s, _, _ := stateOf(met, "approvals.required"); s != enforcement.Met {
		t.Errorf("two required against one declared is %q, want met", s)
	}
	if s, _, _ := stateOf(met, "approvals.not_by_author"); s != enforcement.Met {
		t.Errorf("author approval off is %q, want met", s)
	}

	unmet := host{branch: map[string]any{"name": "main"}, project: map[string]any{},
		approvals: map[string]any{"merge_requests_author_approval": true},
		rules:     []any{}}.ask(t)
	if s, _, _ := stateOf(unmet, "approvals.required"); s != enforcement.Unmet {
		t.Errorf("no rules against one declared is %q, want unmet", s)
	}
	if s, _, note := stateOf(unmet, "approvals.not_by_author"); s != enforcement.Unmet || note == "" {
		t.Errorf("author approval on is %q with note %q, want unmet and a reason", s, note)
	}
}

// Observed against gitlab.com: this host omits the merge settings rather than sending false
// where the caller may not read them. Decoded into a bool an omission becomes off, and the
// requirement would be reported unmet — a setting nobody could read reported as a setting
// nobody made, which is the one collapse the report exists to prevent.
func TestAnOmittedSettingIsNotAvailableRatherThanOff(t *testing.T) {
	qs := host{branch: map[string]any{"name": "main"},
		project:   map[string]any{"id": 1},
		approvals: map[string]any{}, rules: []any{}}.ask(t)
	s, _, note := stateOf(qs, "required_pipeline")
	if s != enforcement.NotAvailable {
		t.Errorf("an omitted setting is %q, want %q", s, enforcement.NotAvailable)
	}
	if !strings.Contains(note, "only_allow_merge_if_pipeline_succeeds") {
		t.Errorf("the reason does not name the field: %q", note)
	}
}

func TestAnUnprotectedBranchIsAnAnswerAndNotAFailure(t *testing.T) {
	qs := host{project: map[string]any{}, approvals: map[string]any{}, rules: []any{},
		status: map[string]int{"branch": http.StatusNotFound}}.ask(t)
	s, _, note := stateOf(qs, "allow_bypass")
	if s != enforcement.NotAvailable || !strings.Contains(note, "not protected") {
		t.Errorf("an unprotected branch is %q with note %q", s, note)
	}
}

// A token the host will not accept cannot answer anything, so it is an error rather than
// five requirements reported as unavailable, which would hide a misconfiguration behind the
// tier.
func TestARejectedTokenIsAnError(t *testing.T) {
	h := host{status: map[string]int{"branch": http.StatusUnauthorized}}
	srv := h.serve(t)
	defer srv.Close()
	if _, err := New(srv.Client(), srv.URL).Requirements("g/p", "main", "t", declared()); err == nil {
		t.Fatal("a rejected token was read as an answer")
	}
}

// A project that cannot be found is not a requirement that is unavailable. Reporting it as
// one would leave somebody looking for a setting in a project that is not there.
func TestAMissingProjectIsAnError(t *testing.T) {
	h := host{branch: map[string]any{"name": "main"},
		status: map[string]int{"project": http.StatusNotFound}}
	srv := h.serve(t)
	defer srv.Close()
	_, err := New(srv.Client(), srv.URL).Requirements("g/p", "main", "t", declared())
	if err == nil {
		t.Fatal("a missing project was reported as a requirement")
	}
	if !strings.Contains(err.Error(), "g/p") {
		t.Errorf("the error does not name the project: %v", err)
	}
}

// The second standing rule, held where an adapter could break it.
func TestTheAdapterInventsNoRequirementName(t *testing.T) {
	known := make(map[string]bool)
	for _, n := range enforcement.Names() {
		known[n] = true
	}
	qs := host{branch: map[string]any{"name": "main"}, project: map[string]any{},
		approvals: map[string]any{}, rules: []any{}}.ask(t)
	if len(qs) != 4 {
		t.Errorf("%d requirements answered, want 4", len(qs))
	}
	for _, q := range qs {
		if !known[q.Name] {
			t.Errorf("the adapter answered with %q, which the domain does not know", q.Name)
		}
	}
}

// The payload is a real one, fetched from gitlab.com unauthenticated, which is the evidence
// A65 turned on. A hand written fixture would assert the field names this adapter already
// believes; this one asserts them against what the host actually sends.
func TestARealPayloadDecodes(t *testing.T) {
	raw, err := os.ReadFile("testdata/protected-branch.json")
	if err != nil {
		t.Fatal(err)
	}
	var pb protectedBranch
	if err := json.Unmarshal(raw, &pb); err != nil {
		t.Fatal(err)
	}
	if pb.Name == "" {
		t.Error("the branch name did not decode")
	}
	if grants(pb.PushAccessLevels) == 0 {
		t.Error("the push grants did not decode, and this project has two")
	}
	// The project the fixture came from permits pushes to its protected branch, so the
	// requirement is unmet and says so in this host's words.
	q := allowBypass(pb)
	if q.State != enforcement.Unmet {
		t.Errorf("state is %q, want %q", q.State, enforcement.Unmet)
	}
	if !strings.Contains(q.Actual, "push without a merge request") {
		t.Errorf("it says %q", q.Actual)
	}
}
