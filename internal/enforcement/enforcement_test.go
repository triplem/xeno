// SPDX-License-Identifier: Apache-2.0

package enforcement

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func declared() Declared {
	var d Declared
	d.RequiredPipeline, d.AllowBypass = true, false
	d.Approvals.Required, d.Approvals.NotByAuthor = 1, true
	return d
}

func stateOf(r Report, name string) (State, string) {
	for _, q := range r.Requirements {
		if q.Name == name {
			return q.State, q.Note
		}
	}
	return "", ""
}

// The distinction the whole report rests on. A setting nobody made is somebody's
// oversight; a setting the host does not have is nobody's, and reporting it as one
// sends them looking for a checkbox that is not there.
func TestNotAvailableIsNotTheSameAsUnmet(t *testing.T) {
	absent := Compare("o/r", "main", declared(),
		Protection{Available: false, Reason: "not on this plan"}, time.Now())
	if s, _ := stateOf(absent, "required_pipeline"); s != NotAvailable {
		t.Fatalf("an unavailable setting is %q, want %q", s, NotAvailable)
	}

	unset := Compare("o/r", "main", declared(),
		Protection{Available: true, Protected: true, ReviewsExpressible: true}, time.Now())
	if s, _ := stateOf(unset, "required_pipeline"); s != Unmet {
		t.Fatalf("a setting nobody made is %q, want %q", s, Unmet)
	}
}

// WP10: a repository whose pipeline is not required is reported as such rather than
// passing quietly.
func TestAnUnprotectedBranchDoesNotPassQuietly(t *testing.T) {
	r := Compare("o/r", "main", declared(),
		Protection{Available: true, Protected: false, Reason: "the branch is not protected"}, time.Now())
	if s, _ := stateOf(r, "required_pipeline"); s != Unmet {
		t.Fatalf("state is %q, want %q", s, Unmet)
	}
	if r.Unmet() == 0 {
		t.Fatal("it passed quietly")
	}
}

// WP10: a requirement the host cannot express is reported once as waived rather than
// daily as unmet. Without this the report says the same thing forever and is ignored
// within a fortnight.
func TestWaivedTurnsAStandingComplaintIntoADecision(t *testing.T) {
	d := declared()
	// Four requirements, because the declaration carries four: the pipeline, the bypass,
	// the count of approvals and whether one may be the author's.
	before := Compare("o/r", "main", d, Protection{Available: false, Reason: "not on this plan"}, time.Now())
	if before.Unmet() != 4 {
		t.Fatalf("%d unmet before waiving, want 4", before.Unmet())
	}

	d.Waived = "not expressible on this plan, recorded 2026-09-25"
	after := Compare("o/r", "main", d, Protection{Available: false, Reason: "not on this plan"}, time.Now())
	if after.Unmet() != 0 {
		t.Fatalf("%d still unmet after waiving", after.Unmet())
	}
	s, note := stateOf(after, "required_pipeline")
	if s != Waived {
		t.Fatalf("state is %q, want %q", s, Waived)
	}
	if note != d.Waived {
		t.Fatalf("the note is %q, want the reason that was recorded", note)
	}
}

func TestAMetRequirementIsMet(t *testing.T) {
	r := Compare("o/r", "main", declared(), Protection{
		Available: true, Protected: true, RequiredStatusChecks: true,
		EnforceAdmins: true, ReviewsExpressible: true, RequiredApprovals: 2,
	}, time.Now())
	if r.Unmet() != 0 {
		t.Fatalf("%d unmet against a host that requires everything", r.Unmet())
	}
	for _, q := range r.Requirements {
		if q.State != Met {
			t.Errorf("%s is %q", q.Name, q.State)
		}
	}
}

// A 403 is the answer a private repository on the free plan gets, and it is not an
// error. Treating it as one would make the command fail where it has something to say.
func TestForbiddenMeansNotAvailableRatherThanAnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()
	p, err := Fetch(srv.Client(), srv.URL, "o/r", "main", "t")
	if err != nil {
		t.Fatalf("a 403 was treated as an error: %v", err)
	}
	if p.Available {
		t.Fatal("a 403 was read as the host offering protection")
	}
	if p.Reason == "" {
		t.Fatal("nothing says why it is unavailable")
	}
}

func TestNotFoundMeansTheBranchIsNotProtected(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()
	p, err := Fetch(srv.Client(), srv.URL, "o/r", "main", "t")
	if err != nil {
		t.Fatal(err)
	}
	if !p.Available || p.Protected {
		t.Fatalf("a 404 read as available=%v protected=%v", p.Available, p.Protected)
	}
}

func TestARejectedTokenIsAnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer srv.Close()
	if _, err := Fetch(srv.Client(), srv.URL, "o/r", "main", "t"); err == nil {
		t.Fatal("a rejected token was read as an answer")
	}
}

// Section 13 says the approvals block covers whether the author may give one, so the
// second half is compared rather than read and ignored. What satisfies it is the host's
// own rule and not a setting: GitHub refuses an approval from the pull request's author.
func TestNotByAuthorIsCompared(t *testing.T) {
	d := declared()
	protected := func(count int, lastPush bool) Protection {
		return Protection{Available: true, Protected: true, RequiredStatusChecks: true,
			EnforceAdmins: true, ReviewsExpressible: true, RequiredApprovals: count,
			LastPushApproval: lastPush}
	}

	r := Compare("o/r", "main", d, protected(1, false), time.Now())
	if s, note := stateOf(r, "approvals.not_by_author"); s != Met || note == "" {
		t.Fatalf("with one approval required the state is %q with note %q, want met and the stronger form named", s, note)
	}

	r = Compare("o/r", "main", d, protected(1, true), time.Now())
	if s, note := stateOf(r, "approvals.not_by_author"); s != Met || note != "" {
		t.Fatalf("with the last push covered the state is %q with note %q, want met and nothing left to add", s, note)
	}

	// An approval that does not exist is nobody's, so the absence does not satisfy a
	// requirement about who gives one.
	r = Compare("o/r", "main", d, protected(0, false), time.Now())
	if s, _ := stateOf(r, "approvals.not_by_author"); s != Unmet {
		t.Fatalf("with no approval required the state is %q, want %q", s, Unmet)
	}

	// Not declared is not compared: a project that says nothing about it gets no line,
	// which is what stateOf reports as the empty state.
	d.Approvals.NotByAuthor = false
	r = Compare("o/r", "main", d, protected(1, false), time.Now())
	if s, _ := stateOf(r, "approvals.not_by_author"); s != "" {
		t.Fatalf("an undeclared requirement produced a line with state %q", s)
	}
}

// merge_method is declared and not compared, which section 13 says outright. Reporting it
// as unchecked is what keeps a project from believing otherwise.
func TestMergeMethodIsReportedAsUnchecked(t *testing.T) {
	d := declared()
	d.MergeMethod = "no-squash"
	r := Compare("o/r", "main", d, Protection{Available: true, Protected: true}, time.Now())
	s, note := stateOf(r, "merge_method")
	if s != Unknown {
		t.Fatalf("state is %q, want %q", s, Unknown)
	}
	if note == "" {
		t.Fatal("nothing says why it is not compared")
	}
	// Unknown is neither met nor unmet, so it does not fail a run on its own.
	if before := Compare("o/r", "main", declared(), Protection{Available: true, Protected: true}, time.Now()); r.Unmet() != before.Unmet() {
		t.Fatalf("declaring merge_method changed the unmet count from %d to %d", before.Unmet(), r.Unmet())
	}
}
