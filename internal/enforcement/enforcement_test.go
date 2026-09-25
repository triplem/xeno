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
	before := Compare("o/r", "main", d, Protection{Available: false, Reason: "not on this plan"}, time.Now())
	if before.Unmet() != 3 {
		t.Fatalf("%d unmet before waiving, want 3", before.Unmet())
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
