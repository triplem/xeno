// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"encoding/json"
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
}

func (h *fakeHost) serve(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.asked = append(h.asked, r.Method+" "+r.URL.Path)
		if strings.HasSuffix(r.URL.Path, "/comments") {
			var in map[string]string
			_ = json.NewDecoder(r.Body).Decode(&in)
			h.comments = append(h.comments, in["body"])
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"html_url": "https://example.invalid/comment/1"})
			return
		}
		if h.status != 0 {
			w.WriteHeader(h.status)
			return
		}
		_ = json.NewEncoder(w).Encode(h.issue)
	}))
	t.Cleanup(srv.Close)
	return srv
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
	if len(h.asked) != 1 || h.asked[0] != "GET /repos/triplem/xeno/issues/288" {
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
	if len(h.asked) != 1 {
		t.Fatalf("the intake asked the host %d times: %v", len(h.asked), h.asked)
	}
	// What the later start does is beside the point and is not asserted: this fixture has no
	// verdict on P0, so it is refused for that. What is asserted is that it refuses without
	// having asked the host, because the read is the intake's and not every phase's.
	_, _ = f.r.StartPhase(key, "01-requirements")
	if len(h.asked) != 1 {
		t.Fatalf("starting a later phase asked the host again: %v", h.asked)
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
