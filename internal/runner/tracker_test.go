// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/template"
)

// fakeHost is a code host that answers from a map of path to payload and records what it
// was asked. Nothing here reaches a network: the project's `tracker.base_url` is a setting
// rather than a constant, which is what lets a test point the adapter at a local server and
// is the same property a self managed deployment relies on.
type fakeHost struct {
	issue    map[string]string
	status   int
	asked    []string
	comments []string
	// What the host says around the issue, in this host's own JSON: the labels and the
	// milestone on the issue, the comments under it, and the repository's open milestones.
	// Nil is an issue with none of them, which is what every test before #330 served.
	labels     []string
	milestone  map[string]any
	discussion []map[string]any
	milestones []map[string]any
}

// approved is the fake host an intent can be started against: the label and a comment
// whose first line is `/xeno approved`, which is the whole of what section 12 asks for.
func approved(issue map[string]string) *fakeHost {
	return &fakeHost{issue: issue, labels: []string{"xeno-approved"},
		discussion: []map[string]any{{
			"user":       map[string]string{"login": "maintainer"},
			"created_at": "2026-10-08T15:10:28Z",
			"body":       "/xeno approved\n\nbecause the shape was decided on the issue",
		}}}
}

func (h *fakeHost) serve(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.asked = append(h.asked, r.Method+" "+r.URL.Path)
		switch {
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/comments"):
			var in map[string]string
			_ = json.NewDecoder(r.Body).Decode(&in)
			h.comments = append(h.comments, in["body"])
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"html_url": "https://example.invalid/comment/1"})
			return
		case h.status != 0:
			w.WriteHeader(h.status)
			return
		case strings.HasSuffix(r.URL.Path, "/comments"):
			_ = json.NewEncoder(w).Encode(append([]map[string]any{}, h.discussion...))
			return
		case strings.HasSuffix(r.URL.Path, "/milestones"):
			_ = json.NewEncoder(w).Encode(append([]map[string]any{}, h.milestones...))
			return
		}
		out := map[string]any{"labels": []map[string]string{}, "milestone": h.milestone}
		for k, v := range h.issue {
			out[k] = v
		}
		for _, l := range h.labels {
			out["labels"] = append(out["labels"].([]map[string]string), map[string]string{"name": l})
		}
		_ = json.NewEncoder(w).Encode(out)
	}))
	t.Cleanup(srv.Close)
	return srv
}

// hostOf is the tracker host a qualified id carries for a fake server's address, which
// is the address without its scheme: the fixture's block points `base_url` there and A78
// derives the host from it.
func hostOf(baseURL string) string {
	return strings.TrimPrefix(strings.TrimPrefix(baseURL, "http://"), "https://")
}

// tracker writes a complete tracker block pointing at the given address, and sets the
// variable it names. Credentials come from the environment and never from the file, so the
// fixture has to do both halves.
func (f *fixture) tracker(baseURL string) {
	f.t.Helper()
	f.t.Setenv("XENO_TRACKER_TOKEN", "a-token")
	f.project("tracker:\n  adapter: github\n  project: triplem/xeno\n" +
		"  base_url: " + baseURL + "\n" +
		"  auth: { scheme: token, secret_env: XENO_TRACKER_TOKEN }\n")
}

// intentOn rewrites the fixture's intent.yaml so that its qualified id names the host the
// given address serves. An intent whose issue lies elsewhere is a case of its own below.
func (f *fixture) intentOn(baseURL, issue string) {
	f.t.Helper()
	host := strings.TrimPrefix(strings.TrimPrefix(baseURL, "http://"), "https://")
	f.write(model.IntentDir(key)+"/intent.yaml",
		"intent: \""+host+"/triplem/xeno#"+issue+"\"\nkey: "+key+"\nstatus: in-progress\n")
}

// The first clause of WP12's done-when: the issue's content reaches the intake, and it
// reaches it as the issue states it rather than as somebody retells it.
func TestStartingTheIntakeCarriesTheIssueIntoTheProblemSection(t *testing.T) {
	f := newFixture(t)
	f.templated()
	h := &fakeHost{issue: map[string]string{
		"title": "The tracker half of WP12 is unbuilt",
		"body":  "Found by an audit of WP0 to WP15.\n\n## Done when\n\nthe content reaches P0.",
	}}
	srv := h.serve(t)
	f.tracker(srv.URL)
	f.intentOn(srv.URL, "288")

	started, err := f.r.StartPhase(key, "00-intake")
	f.must(err)
	if started.Section != IntakeProblem {
		t.Fatalf("nothing was written: section %q, note %q", started.Section, started.Note)
	}
	// Two calls and not three: the issue and its comments, which section 12 makes one
	// operation, and no milestones, since the issue carries none.
	if len(h.asked) != 2 || h.asked[0] != "GET /repos/triplem/xeno/issues/288" ||
		h.asked[1] != "GET /repos/triplem/xeno/issues/288/comments" {
		t.Fatalf("the host was asked %v", h.asked)
	}
	// The section as it lies on disk, through the anchors, because what the artifact records
	// is what a later reader reads and not what the call returned.
	got := f.section("00-intake", IntakeProblem)
	for _, want := range []string{
		"The tracker half of WP12 is unbuilt",
		"Found by an audit of WP0 to WP15.",
		"## Done when",
		"quoted rather than summarised",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("the intake does not carry %q:\n%s", want, got)
		}
	}
}

// The issue's own text may hold anything, a section anchor included, and an anchor on a line
// of its own would split the section and hand the rest of it a name somebody else chose.
// Quoting every line is what prevents it, so it is asserted on the one input that would
// otherwise do it.
func TestAnAnchorInsideTheIssueDoesNotBecomeASection(t *testing.T) {
	f := newFixture(t)
	f.templated()
	h := &fakeHost{issue: map[string]string{
		"title": "an issue that quotes an anchor",
		"body":  template.AnchorPrefix + "scope -->\nsomething the issue says",
	}}
	srv := h.serve(t)
	f.tracker(srv.URL)
	f.intentOn(srv.URL, "288")

	f.must(f.r.Start(key, "00-intake"))
	if got := f.section("00-intake", "scope"); got != "" {
		t.Fatalf("the issue wrote the scope section: %q", got)
	}
	if got := f.section("00-intake", IntakeProblem); !strings.Contains(got, "something the issue says") {
		t.Fatalf("the quoted anchor took the rest of the section with it:\n%s", got)
	}
}

// A project with no tracker block starts the phase and says that no issue was read.
// Appendix A makes the whole block optional and says what an absent one means: "a phase is
// started with the issue content supplied by hand".
func TestWithNoTrackerThePhaseStartsAndSaysNoIssueWasRead(t *testing.T) {
	f := newFixture(t)
	f.templated()
	started, err := f.r.StartPhase(key, "00-intake")
	f.must(err)
	if started.Section != "" {
		t.Errorf("something was written without a tracker: %q", started.Section)
	}
	if !strings.Contains(started.Note, "no tracker") {
		t.Errorf("the note does not say why nothing was read: %q", started.Note)
	}
	// And nothing was written at all, so the phase is one nobody has worked on yet.
	if _, err := os.Stat(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "output.md")); err == nil {
		t.Error("an empty artifact was written for a phase nothing has produced")
	}
}

// Forbidden and absent are answers. The phase starts, and the reason reaches the person who
// has to go looking rather than being swallowed.
func TestAHostThatDeclinesTheReadStillStartsThePhase(t *testing.T) {
	for _, status := range []int{http.StatusForbidden, http.StatusNotFound} {
		f := newFixture(t)
		f.templated()
		h := &fakeHost{status: status}
		srv := h.serve(t)
		f.tracker(srv.URL)
		f.intentOn(srv.URL, "288")

		started, err := f.r.StartPhase(key, "00-intake")
		if err != nil {
			t.Fatalf("%d stopped the phase from starting: %v", status, err)
		}
		if started.Section != "" || started.Note == "" {
			t.Fatalf("%d gave section %q note %q", status, started.Section, started.Note)
		}
	}
}

// An intent whose issue lies on a host the project is not configured for is an absence and
// not a misconfiguration: `--for` takes a whole qualified id, so such an intent is
// legitimate, and the one honest answer about its issue is that nothing was read.
func TestAnIssueOnAnotherHostIsNotReadAndIsNotAnError(t *testing.T) {
	f := newFixture(t)
	f.templated()
	h := &fakeHost{issue: map[string]string{"title": "t", "body": "b"}}
	srv := h.serve(t)
	f.tracker(srv.URL)
	// The fixture's own intent.yaml names git.example, which is not what the block configures.
	started, err := f.r.StartPhase(key, "00-intake")
	f.must(err)
	if started.Section != "" {
		t.Errorf("an issue on another host was read anyway")
	}
	if !strings.Contains(started.Note, "cannot reach") {
		t.Errorf("the note does not say why: %q", started.Note)
	}
	if len(h.asked) != 0 {
		t.Errorf("a call went out for an issue this adapter cannot reach: %v", h.asked)
	}
}

// Only P0. Section 12 asks a phase start to read the issue, and the intake is where the
// problem as stated belongs; every later phase works from the intake rather than from the
// issue again, so a call per phase would be five calls for nothing.
func TestOnlyTheIntakeReadsTheIssue(t *testing.T) {
	f := newFixture(t)
	h := &fakeHost{issue: map[string]string{"title": "t", "body": "b"}}
	srv := h.serve(t)
	f.tracker(srv.URL)
	f.intentOn(srv.URL, "288")

	f.must(f.r.Start(key, "00-intake"))
	// Two calls, the issue and its comments, which is the one read section 12 names.
	if len(h.asked) != 2 {
		t.Fatalf("the intake asked the host %d times: %v", len(h.asked), h.asked)
	}
	// What the later start does is beside the point and is not asserted: this fixture has no
	// verdict on P0, so it is refused for that. What is asserted is that it refuses without
	// having asked the host, because the read is the intake's and not every phase's.
	_, _ = f.r.StartPhase(key, "01-requirements")
	if len(h.asked) != 2 {
		t.Fatalf("starting a later phase asked the host again: %v", h.asked)
	}
}

// ---- section 12's "Starting an intent": an issue nobody approved does not become one

// The label alone and the comment alone are each refused, and the refusal names the half
// that is missing: a label carries no reason, and a comment can be written by anybody.
func TestAnIssueNobodyApprovedDoesNotBecomeAnIntent(t *testing.T) {
	cases := []struct {
		name string
		host *fakeHost
		want string
	}{
		{"neither", &fakeHost{issue: map[string]string{"title": "t", "body": "b"}},
			"the label xeno-approved and a comment whose first line is /xeno approved"},
		{"the label alone", &fakeHost{issue: map[string]string{"title": "t"},
			labels: []string{"xeno-approved"}}, "missing a comment whose first line is /xeno approved"},
		{"the comment alone", &fakeHost{issue: map[string]string{"title": "t"},
			discussion: approved(nil).discussion}, "missing the label xeno-approved"},
		{"a comment that only contains the word", &fakeHost{issue: map[string]string{"title": "t"},
			labels: []string{"xeno-approved"}, discussion: []map[string]any{{
				"user": map[string]string{"login": "m"}, "body": "this could be approved later"}}},
			"missing a comment whose first line is /xeno approved"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := newFixture(t)
			f.tracker(c.host.serve(t).URL)
			_, err := f.r.IntentStart("NEW-1", "330", false)
			var ref *Refusal
			if !errors.As(err, &ref) || !strings.Contains(ref.Reason, c.want) {
				t.Fatalf("got %v, want a refusal saying %q", err, c.want)
			}
			if fm.Exists(filepath.Join(f.root, model.IntentDir("NEW-1"))) {
				t.Error("a refused start left a directory behind")
			}
		})
	}
}

// Both halves present, the intent starts exactly as it did before the clause.
func TestAnApprovedIssueBecomesAnIntent(t *testing.T) {
	f := newFixture(t)
	h := approved(map[string]string{"title": "t", "body": "b"})
	srv := h.serve(t)
	f.tracker(srv.URL)
	in, err := f.r.IntentStart("NEW-1", "330", false)
	f.must(err)
	if in.Intent != hostOf(srv.URL)+"/triplem/xeno#330" || in.Status != "in-progress" {
		t.Errorf("wrote %+v", in)
	}
}

// A milestone holds the issue while an earlier one is open, by due date with undated ones
// last; --now starts it anyway; the earliest open milestone and a closed one hold nothing.
func TestAMilestoneHoldsTheIssueUntilItsTurn(t *testing.T) {
	open := []map[string]any{
		{"title": "later", "number": 3, "state": "open"},
		{"title": "1.1", "due_on": "2026-12-31T00:00:00Z", "number": 2, "state": "open"},
		{"title": "1.0", "due_on": "2026-11-30T00:00:00Z", "number": 1, "state": "open"},
	}
	on := func(title string, number int, state string) map[string]any {
		return map[string]any{"title": title, "number": number, "state": state}
	}
	cases := []struct {
		name      string
		milestone map[string]any
		now       bool
		held      string
	}{
		{"the second milestone", on("1.1", 2, "open"), false, `"1.0" is open ahead of it`},
		{"the undated one", on("later", 3, "open"), false, `"1.0" is open ahead of it`},
		{"the second, with --now", on("1.1", 2, "open"), true, ""},
		{"the earliest", on("1.0", 1, "open"), false, ""},
		{"a closed one", on("0.9", 0, "closed"), false, ""},
		{"none", nil, false, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := newFixture(t)
			h := approved(map[string]string{"title": "t"})
			h.milestone, h.milestones = c.milestone, open
			f.tracker(h.serve(t).URL)
			_, err := f.r.IntentStart("NEW-1", "330", c.now)
			var ref *Refusal
			switch {
			case c.held == "" && err != nil:
				t.Fatalf("held: %v", err)
			case c.held != "" && (!errors.As(err, &ref) || !strings.Contains(ref.Reason, c.held)):
				t.Fatalf("got %v, want a refusal saying %s", err, c.held)
			}
			// The milestones are listed only where the issue carries one: a read nobody
			// consults is a call for nothing.
			listed := false
			for _, a := range h.asked {
				listed = listed || strings.HasSuffix(a, "/milestones")
			}
			if listed != (c.milestone != nil) {
				t.Errorf("the milestones were listed: %v, for an issue with milestone %v", listed, c.milestone)
			}
		})
	}
}

// Every answer but an issue is a refusal here, where phase start carries on: a start that
// could not read the approval has nothing to stand on. The missing token is the one that
// differs most from phase start, which has always started without one.
func TestAReadThatCannotHappenDoesNotStartTheIntent(t *testing.T) {
	t.Run("no token", func(t *testing.T) {
		f := newFixture(t)
		h := approved(map[string]string{"title": "t"})
		f.tracker(h.serve(t).URL)
		t.Setenv("XENO_TRACKER_TOKEN", "")
		_, err := f.r.IntentStart("NEW-1", "330", false)
		var ref *Refusal
		if !errors.As(err, &ref) || !strings.Contains(ref.Reason, "XENO_TRACKER_TOKEN") {
			t.Fatalf("got %v, want a refusal naming the variable", err)
		}
		if len(h.asked) != 0 {
			t.Errorf("a call went out without a token: %v", h.asked)
		}
	})
	for _, status := range []int{http.StatusNotFound, http.StatusForbidden} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			f := newFixture(t)
			f.tracker((&fakeHost{status: status}).serve(t).URL)
			_, err := f.r.IntentStart("NEW-1", "330", false)
			var ref *Refusal
			if !errors.As(err, &ref) || !strings.Contains(ref.Reason, "not started") {
				t.Fatalf("%d: got %v, want a refusal", status, err)
			}
		})
	}
	t.Run("another host", func(t *testing.T) {
		f := newFixture(t)
		h := approved(map[string]string{"title": "t"})
		f.tracker(h.serve(t).URL)
		_, err := f.r.IntentStart("NEW-1", "git.example/group/proj#4", false)
		var ref *Refusal
		if !errors.As(err, &ref) || !strings.Contains(ref.Reason, "cannot reach") {
			t.Fatalf("got %v, want a refusal naming the host", err)
		}
		if len(h.asked) != 0 {
			t.Errorf("a call went out for an issue this adapter cannot reach: %v", h.asked)
		}
	})
	t.Run("a rejected token is an error and not a refusal", func(t *testing.T) {
		f := newFixture(t)
		f.tracker((&fakeHost{status: http.StatusUnauthorized}).serve(t).URL)
		_, err := f.r.IntentStart("NEW-1", "330", false)
		var ref *Refusal
		if err == nil || errors.As(err, &ref) {
			t.Fatalf("got %v, want an error: a token the host rejects is exit 2, not 1", err)
		}
	})
}

// The intake says by what it was authorised, above the quote: who, when and why, read at
// the moment P0 starts, and that the intent is ahead of its milestone where it still is.
func TestTheIntakeSaysByWhatItWasAuthorised(t *testing.T) {
	f := newFixture(t)
	f.templated()
	h := approved(map[string]string{"title": "t", "body": "b"})
	h.milestone = map[string]any{"title": "1.1", "number": 2, "state": "open"}
	h.milestones = []map[string]any{
		{"title": "1.0", "due_on": "2026-11-30T00:00:00Z", "number": 1, "state": "open"},
		{"title": "1.1", "due_on": "2026-12-31T00:00:00Z", "number": 2, "state": "open"},
	}
	srv := h.serve(t)
	f.tracker(srv.URL)
	f.intentOn(srv.URL, "330")

	f.must(f.r.Start(key, "00-intake"))
	got := f.section("00-intake", IntakeProblem)
	for _, want := range []string{
		"Approved by @maintainer on 2026-10-08T15:10:28Z: because the shape was decided on the issue",
		`Started ahead of its milestone, "1.1", while "1.0" is open.`,
		"> **" + hostOf(srv.URL) + "/triplem/xeno#330** — t",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("the intake does not say %q:\n%s", want, got)
		}
	}
	if !strings.Contains(got, "Approved by") || strings.Index(got, "Approved by") > strings.Index(got, "> **") {
		t.Errorf("the sentence is not above the quote:\n%s", got)
	}
}

// And where it finds none, it says so rather than nothing: an intake silent on the point
// reads as one written before the clause.
func TestTheIntakeSaysWhereNoApprovalWasFound(t *testing.T) {
	f := newFixture(t)
	f.templated()
	h := &fakeHost{issue: map[string]string{"title": "t", "body": "b"}}
	srv := h.serve(t)
	f.tracker(srv.URL)
	f.intentOn(srv.URL, "330")

	f.must(f.r.Start(key, "00-intake"))
	got := f.section("00-intake", IntakeProblem)
	if !strings.Contains(got, "No approval was found on the issue when this phase started: it is missing the label xeno-approved and a comment whose first line is /xeno approved.") {
		t.Errorf("the intake does not say what was missing:\n%s", got)
	}
}

// ---- the second clause: the verdict reaches the issue

// The write-back is one call, carrying the phase, the verdict and the findings.
func TestTheVerdictIsWrittenBackInOneCall(t *testing.T) {
	f := newFixture(t)
	f.templated()
	h := &fakeHost{issue: map[string]string{"title": "t", "body": "b"}}
	srv := h.serve(t)
	f.tracker(srv.URL)
	f.intentOn(srv.URL, "288")
	g := f.redByRendering()

	w, err := f.r.ReportVerdict(key, "", false)
	f.must(err)
	if !w.Posted {
		t.Fatalf("nothing was posted: %q", w.Written.Reason)
	}
	if w.Phase != "00-intake" {
		t.Errorf("reported on %q, and the only judged phase is 00-intake", w.Phase)
	}
	if len(h.comments) != 1 {
		t.Fatalf("the write-back made %d calls, and section 12 costs it one", len(h.comments))
	}
	body := h.comments[0]
	for _, want := range []string{"00-intake is red", g.Checks[0].Gate, "Findings",
		"undecided", "`" + key + "`"} {
		if !strings.Contains(body, want) {
			t.Errorf("the comment does not carry %q:\n%s", want, body)
		}
	}
	// A red verdict is reported and not suppressed, and reporting it is not itself a
	// failure: the gate run has already said the phase is red.
	if w.Written.URL == "" {
		t.Error("the address the host answered with did not reach the caller")
	}
}

// A decision on a finding is part of the verdict, and the comment is the one place a reader
// without a clone sees that somebody has taken one.
func TestTheCommentCarriesTheDecisionOnAFinding(t *testing.T) {
	f := newFixture(t)
	f.templated()
	g := f.redByRendering()
	found := firstFinding(g)
	if found == nil {
		t.Fatal("this fixture is meant to have a finding to decide")
	}
	id := found.ID
	_, err := f.r.Decide(key, "00-intake", id, "approved", "A Person", "the field is absent on purpose")
	f.must(err)

	w, err := f.r.ReportVerdict(key, "00-intake", true)
	f.must(err)
	for _, want := range []string{"approved by A Person", "the field is absent on purpose", id} {
		if !strings.Contains(w.Body, want) {
			t.Errorf("the comment does not carry %q:\n%s", want, w.Body)
		}
	}
}

// The enforcement report lies beside the verdict in the comment, which is the second half
// of WP12's done-when. It is a pipeline artifact under the local data location and never
// part of the trail (A39), so the comment is where the two meet.
func TestTheEnforcementReportLiesBesideTheVerdict(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.redByRendering()
	f.write(".xeno/local/"+ReportFile, "repository: triplem/xeno\nbranch: main\n"+
		"checked_at: 2026-10-07T09:00:00Z\nrequirements:\n"+
		"  - name: required_pipeline\n    declared: \"true\"\n    actual: required\n    state: met\n"+
		"  - name: allow_bypass\n    declared: \"false\"\n    actual: administrators may bypass\n    state: unmet\n")

	w, err := f.r.ReportVerdict(key, "00-intake", true)
	f.must(err)
	for _, want := range []string{"What the host enforces", "required_pipeline", "met",
		"administrators may bypass", "neither met nor waived"} {
		if !strings.Contains(w.Body, want) {
			t.Errorf("the comment does not carry %q:\n%s", want, w.Body)
		}
	}
}

// Absent is said rather than left out. A comment that silently omitted the section would
// read as a host with nothing to report, and the whole reason the check exists is that an
// unchecked host is indistinguishable from a compliant one.
func TestAnAbsentEnforcementReportIsSaidAndNotOmitted(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.redByRendering()
	w, err := f.r.ReportVerdict(key, "00-intake", true)
	f.must(err)
	if !strings.Contains(w.Body, "No enforcement report") {
		t.Errorf("the comment says nothing about the missing report:\n%s", w.Body)
	}
}

// A dry run composes the comment and writes nothing, which is what lets a wording be read
// before anybody receives it and what lets a job without a credential log what it would
// have said.
func TestADryRunComposesTheCommentAndPostsNothing(t *testing.T) {
	f := newFixture(t)
	f.templated()
	h := &fakeHost{issue: map[string]string{"title": "t", "body": "b"}}
	srv := h.serve(t)
	f.tracker(srv.URL)
	f.intentOn(srv.URL, "288")
	f.redByRendering()

	w, err := f.r.ReportVerdict(key, "00-intake", true)
	f.must(err)
	if w.Posted || len(h.comments) != 0 {
		t.Fatalf("a dry run wrote something: posted=%v comments=%v", w.Posted, h.comments)
	}
	if w.Body == "" {
		t.Fatal("a dry run composed nothing")
	}
}

// A phase nobody has judged has no verdict to report, which is a refusal naming what to do
// rather than an empty comment.
func TestReportingAPhaseWithNoVerdictIsRefused(t *testing.T) {
	f := newFixture(t)
	if _, err := f.r.ReportVerdict(key, "", true); err == nil {
		t.Fatal("an intent with no verdict at all was reported on")
	}
	if _, err := f.r.ReportVerdict(key, "02-design", true); err == nil {
		t.Fatal("a phase with no verdict was reported on")
	}
}

// A cause or a note written for a terminal may hold newlines and pipes, and both break the
// list or the table it lands in. The reader of the issue gets the same words on one line.
func TestTextWrittenForATerminalDoesNotBreakTheTable(t *testing.T) {
	got := oneLine("one\ntwo | three\n  four")
	if got != "one two \\| three four" {
		t.Errorf("read %q", got)
	}
}

// section reads one section of a phase's output.md through the anchors, which is what a
// later reader of the artifact sees.
func (f *fixture) section(phase, id string) string {
	f.t.Helper()
	b, err := os.ReadFile(filepath.Join(f.root, model.PhaseDir(key, phase), "output.md"))
	if err != nil {
		return ""
	}
	_, body, err := fm.Split(b)
	f.must(err)
	return template.Parse(string(body))[id]
}
