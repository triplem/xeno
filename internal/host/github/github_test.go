// SPDX-License-Identifier: Apache-2.0

package github

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/triplem/xeno/internal/enforcement"
)

func declared() enforcement.Declared {
	var d enforcement.Declared
	d.RequiredPipeline, d.AllowBypass = true, false
	d.Approvals.Required, d.Approvals.NotByAuthor = 1, true
	return d
}

func stateOf(qs []enforcement.Requirement, name string) (enforcement.State, string) {
	for _, q := range qs {
		if q.Name == name {
			return q.State, q.Note
		}
	}
	return "", ""
}

// A 403 is the answer a private repository on the free plan gets, and it is not an
// error. Treating it as one would make the command fail where it has something to say.
func TestForbiddenMeansNotAvailableRatherThanAnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()
	p, err := New(srv.Client(), srv.URL).fetch("o/r", "main", "t")
	if err != nil {
		t.Fatalf("a 403 was treated as an error: %v", err)
	}
	if p.available {
		t.Fatal("a 403 was read as the host offering protection")
	}
	if p.reason == "" {
		t.Fatal("nothing says why it is unavailable")
	}
}

func TestNotFoundMeansTheBranchIsNotProtected(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()
	p, err := New(srv.Client(), srv.URL).fetch("o/r", "main", "t")
	if err != nil {
		t.Fatal(err)
	}
	if !p.available || p.protected {
		t.Fatalf("a 404 read as available=%v protected=%v", p.available, p.protected)
	}
}

func TestARejectedTokenIsAnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer srv.Close()
	if _, err := New(srv.Client(), srv.URL).fetch("o/r", "main", "t"); err == nil {
		t.Fatal("a rejected token was read as an answer")
	}
}

// Section 13 says the approvals block covers whether the author may give one, so the
// second half is compared rather than read and ignored. What satisfies it is the host's
// own rule and not a setting: GitHub refuses an approval from the pull request's author.
//
// It is this host's rule, which is why the case moved here with A65: on GitLab the same
// requirement is a setting and has a reachable unmet that this one does not.
func TestNotByAuthorIsCompared(t *testing.T) {
	a := New(nil, "https://example.invalid")
	protected := func(count int, lastPush bool) protection {
		return protection{available: true, protected: true, requiredStatusChecks: true,
			enforceAdmins: true, reviewsExpressible: true, requiredApprovals: count,
			lastPushApproval: lastPush}
	}
	answers := func(p protection) []enforcement.Requirement {
		return []enforcement.Requirement{a.approvalsNotByAuthor(p)}
	}

	r := answers(protected(1, false))
	if s, note := stateOf(r, "approvals.not_by_author"); s != enforcement.Met || note == "" {
		t.Fatalf("with one approval required the state is %q with note %q, want met and the stronger form named", s, note)
	}

	r = answers(protected(1, true))
	if s, note := stateOf(r, "approvals.not_by_author"); s != enforcement.Met || note != "" {
		t.Fatalf("with the last push covered the state is %q with note %q, want met and nothing left to add", s, note)
	}

	// An approval that does not exist is nobody's, so the absence does not satisfy a
	// requirement about who gives one.
	r = answers(protected(0, false))
	if s, _ := stateOf(r, "approvals.not_by_author"); s != enforcement.Unmet {
		t.Fatalf("with no approval required the state is %q, want %q", s, enforcement.Unmet)
	}
}

// A host that offers nothing still names every requirement, carrying the reason. An
// omission would read the same in the report and would lose the sentence that tells
// somebody there is no checkbox to go and find.
func TestAnUnavailableHostStillSaysWhyForEveryRequirement(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()
	qs, err := New(srv.Client(), srv.URL).Requirements("o/r", "main", "t", declared())
	if err != nil {
		t.Fatal(err)
	}
	if len(qs) != 4 {
		t.Fatalf("%d requirements answered, want 4", len(qs))
	}
	for _, q := range qs {
		if q.State != enforcement.NotAvailable {
			t.Errorf("%s is %q, want %q", q.Name, q.State, enforcement.NotAvailable)
		}
		if q.Note == "" {
			t.Errorf("%s says nothing about why", q.Name)
		}
	}
}

// An unprotected branch has no protection to carry a review rule, so the two approvals
// rows are not available rather than unmet. Reporting them as unmet would ask somebody to
// fix a setting that cannot exist until the branch is protected.
func TestAnUnprotectedBranchCannotExpressApprovals(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()
	qs, err := New(srv.Client(), srv.URL).Requirements("o/r", "main", "t", declared())
	if err != nil {
		t.Fatal(err)
	}
	if s, _ := stateOf(qs, "required_pipeline"); s != enforcement.Unmet {
		t.Errorf("required_pipeline is %q, want %q", s, enforcement.Unmet)
	}
	if s, _ := stateOf(qs, "approvals.required"); s != enforcement.NotAvailable {
		t.Errorf("approvals.required is %q, want %q", s, enforcement.NotAvailable)
	}
}

// The second standing rule, held where an adapter could break it: every name this adapter
// answers with is one the domain knows. A host may choose its own words, never its own
// name for what was asked.
func TestTheAdapterInventsNoRequirementName(t *testing.T) {
	known := make(map[string]bool)
	for _, n := range enforcement.Names() {
		known[n] = true
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"required_status_checks":{"contexts":["verify"]},
		  "enforce_admins":{"enabled":true},
		  "required_pull_request_reviews":{"required_approving_review_count":2}}`))
	}))
	defer srv.Close()
	qs, err := New(srv.Client(), srv.URL).Requirements("o/r", "main", "t", declared())
	if err != nil {
		t.Fatal(err)
	}
	for _, q := range qs {
		if !known[q.Name] {
			t.Errorf("the adapter answered with %q, which the domain does not know", q.Name)
		}
	}
}

// The split of fetch into transport and decode is what makes these two possible: the status
// mapping without a body, and the body without a server.
func TestFetchMapsTheAnswersAHostCanGive(t *testing.T) {
	for _, tc := range []struct {
		status    int
		available bool
		protected bool
		wantErr   bool
	}{
		{status: 200, available: true, protected: true},
		{status: 403},                  // not on this plan, and not an error
		{status: 404, available: true}, // the branch is not protected
		{status: 401, wantErr: true},   // the token was rejected
		{status: 500, wantErr: true},
	} {
		body := `{"enforce_admins":{"enabled":true}}`
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if got := r.Header.Get("Authorization"); got != "Bearer t" {
				t.Errorf("the token was not sent: %q", got)
			}
			if got := r.Header.Get("Accept"); got != "application/vnd.github+json" {
				t.Errorf("this host's Accept header was not sent: %q", got)
			}
			w.WriteHeader(tc.status)
			if tc.status == 200 {
				_, _ = w.Write([]byte(body))
			}
		}))
		p, err := New(srv.Client(), srv.URL).fetch("o/r", "main", "t")
		srv.Close()
		if (err != nil) != tc.wantErr {
			t.Fatalf("%d gave err=%v, wanted error=%v", tc.status, err, tc.wantErr)
		}
		if err != nil {
			continue
		}
		if p.available != tc.available || p.protected != tc.protected {
			t.Errorf("%d gave available=%v protected=%v, wanted %v and %v",
				tc.status, p.available, p.protected, tc.available, tc.protected)
		}
		if !tc.available && p.reason == "" {
			t.Errorf("%d says nothing about why", tc.status)
		}
	}
}

func TestDecodeProtectionReadsTheHostsFieldNames(t *testing.T) {
	raw := []byte(`{
	  "required_status_checks": {"contexts": ["verify"]},
	  "enforce_admins": {"enabled": true},
	  "required_pull_request_reviews": {"required_approving_review_count": 2,
	                                    "require_last_push_approval": true}
	}`)
	p, err := decode(raw)
	if err != nil {
		t.Fatal(err)
	}
	if !p.requiredStatusChecks || !p.enforceAdmins || p.requiredApprovals != 2 ||
		!p.reviewsExpressible || !p.lastPushApproval {
		t.Fatalf("decoded %+v", p)
	}

	// A branch with no review requirement at all: expressible is false, which is not the
	// same as zero approvals configured.
	p, err = decode([]byte(`{"enforce_admins":{"enabled":false}}`))
	if err != nil {
		t.Fatal(err)
	}
	if p.reviewsExpressible || p.requiredApprovals != 0 || p.enforceAdmins {
		t.Fatalf("decoded %+v", p)
	}
}
