// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
)

// isRefusal is the assertion these tests make most: a command declining with a reason,
// rather than a failure or a write that went through.
func isRefusal(err error) bool {
	var ref *Refusal
	return errors.As(err, &ref)
}

// The question a person can act on: two options with their consequence, a recommendation,
// and the free entry section 8 asks for always. Without a key, because the key is assigned.
const askable = `text: which error behaviour?
options:
  - text: fail fast
    consequence: the caller retries
    recommended: true
  - text: retry internally
    consequence: the caller never sees it
  - text: something else, in the words of whoever decides
    free: true
`

// started writes a phase's artifact without finishing it, which is the state both commands
// are used in: the sections are being written and the verdict does not exist yet.
func (f *fixture) started(phase string) {
	f.t.Helper()
	f.must(f.r.Start(key, phase))
	f.output(phase, "")
}

func (f *fixture) question(phase, entry string) *model.Question {
	f.t.Helper()
	q, err := f.r.RecordQuestion(key, phase, []byte(entry))
	f.must(err)
	return q
}

// readFront reads the frontmatter of a phase's artifact as the gates read it.
func (f *fixture) readFront(phase string, o *model.Output) {
	f.t.Helper()
	_, err := fm.ReadFront(filepath.Join(f.root, model.PhaseDir(key, phase), "output.md"), o)
	f.must(err)
}

func (f *fixture) body(phase string) string {
	f.t.Helper()
	b, err := os.ReadFile(filepath.Join(f.root, model.PhaseDir(key, phase), "output.md"))
	f.must(err)
	_, body, err := fm.Split(b)
	f.must(err)
	return string(body)
}

// ---- WP5: the exchange is written by a command

func TestAQuestionIsWrittenWithTheKeyTheRunnerAssigns(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	if q := f.question("00-intake", askable); q.Key != "Q-1" {
		t.Fatalf("the first question of an intent is Q-1, not %s", q.Key)
	}
	var o model.Output
	f.readFront("00-intake", &o)
	if len(o.OpenQuestions) != 1 || o.OpenQuestions[0].Options[2].Free != true {
		t.Fatalf("the entry did not reach the frontmatter as it was given: %+v", o.OpenQuestions)
	}
}

// The key is unique over the intent and not over the file, because resolves is read from
// any later phase and a reference to one of two Q-1s would be ambiguous.
func TestKeysContinueAcrossThePhasesOfOneIntent(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	f.question("00-intake", askable)
	f.finish("00-intake")
	f.started("01-requirements")
	if q := f.question("01-requirements", askable); q.Key != "Q-2" {
		t.Fatalf("the second question of an intent is Q-2, not %s", q.Key)
	}
}

func TestAQuestionWithoutEnoughOptionsIsRefusedBeforeTheWrite(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	before := f.body("00-intake")
	_, err := f.r.RecordQuestion(key, "00-intake", []byte("text: bare\n"))
	if !isRefusal(err) || !strings.Contains(err.Error(), "two to four options") {
		t.Fatalf("a question with no options was accepted: %v", err)
	}
	var o model.Output
	f.readFront("00-intake", &o)
	if len(o.OpenQuestions) != 0 || f.body("00-intake") != before {
		t.Fatal("a refused question changed the artifact")
	}
}

// The exception section 8 names: an agent that found no options says so rather than
// inventing two.
func TestAnHonestNoOptionsIsAccepted(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	f.question("00-intake", "text: who owns the schedule?\nno_options: true\n")
}

func TestAKeyOnTheWayInIsRefused(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	_, err := f.r.RecordQuestion(key, "00-intake", []byte("key: Q-7\n"+askable))
	if !isRefusal(err) || !strings.Contains(err.Error(), "assigned and not given") {
		t.Fatalf("a given key was accepted: %v", err)
	}
}

// A field nothing enumerates is refused rather than dropped. Dropped, a mistyped
// consequence would leave a question that reads complete in the session and carries half of
// itself in the artifact.
func TestAnUnknownFieldInAQuestionIsRefused(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	_, err := f.r.RecordQuestion(key, "00-intake", []byte("text: t\nurgency: high\nno_options: true\n"))
	if !isRefusal(err) || !strings.Contains(err.Error(), "urgency") {
		t.Fatalf("an invented field was accepted: %v", err)
	}
}

func TestADecisionNeedsItsPersonItsReasonAndItsOption(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	for _, c := range []struct {
		why string
		d   model.Decision
		say string
	}{
		{"no person", model.Decision{Chosen: "fail fast", Rationale: "the caller retries"}, "--by"},
		{"no reason", model.Decision{Chosen: "fail fast", DecidedBy: "m.example"}, "--reason"},
		{"no option", model.Decision{Rationale: "r", DecidedBy: "m.example"}, "--chosen"},
		{"a withdrawal with an option", model.Decision{
			Withdrawn: true, Chosen: "fail fast", Rationale: "r", DecidedBy: "m.example"}, "--withdraw"},
		{"a given id", model.Decision{
			ID: "D-9", Chosen: "c", Rationale: "r", DecidedBy: "m.example"}, "assigned and not given"},
	} {
		_, err := f.r.RecordDecision(key, "00-intake", c.d)
		if !isRefusal(err) || !strings.Contains(err.Error(), c.say) {
			t.Fatalf("%s was accepted: %v", c.why, err)
		}
	}
	var o model.Output
	f.readFront("00-intake", &o)
	if len(o.Decisions) != 0 {
		t.Fatal("a refused decision reached the artifact")
	}
}

// A typo in --resolves would otherwise leave the question unresolved and the trail would
// say so three phases later, at P5, with a decision in it that looks like the answer.
func TestADecisionResolvingAQuestionNobodyRaisedIsRefused(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	_, err := f.r.RecordDecision(key, "00-intake", model.Decision{
		Resolves: "Q-4", Chosen: "fail fast", Rationale: "r", DecidedBy: "m.example"})
	if !isRefusal(err) || !strings.Contains(err.Error(), "no question Q-4") {
		t.Fatalf("a decision resolving nothing was accepted: %v", err)
	}
}

// The loop end to end, which is what #188 says has never once run: raised in one phase,
// settled in a later one by a person, and seen as resolved by the gate at P5.
func TestAQuestionRecordedAndDecidedIsResolvedForTheGate(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	q := f.question("00-intake", askable)
	f.finish("00-intake")
	f.started("01-requirements")
	d, err := f.r.RecordDecision(key, "01-requirements", model.Decision{
		Resolves: q.Key, Chosen: "fail fast", Rationale: "the caller retries",
		DecidedBy: "m.example", ProposedBy: "the agent"})
	f.must(err)
	if d.ID != "D-1" {
		t.Fatalf("the first decision of an intent is D-1, not %s", d.ID)
	}
	f.finish("01-requirements")
	for _, p := range model.Phases[2:5] {
		f.run(p, "")
	}
	g := f.run("05-review", "")
	if c := check(g, "G-Questions"); c.Result != "pass" {
		t.Fatalf("a decided question was not resolved: %+v", c)
	}
	if s := f.r.Summarise(key); s.Questions != 1 || s.Decisions != 1 || s.AskedNothing() {
		t.Fatalf("the figure disagrees with the trail: %+v", s)
	}
}

// Section 8's third exit, which has to exist so that nobody invents an assumption to get a
// gate green.
func TestAWithdrawalResolvesTheQuestionAsWell(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	q := f.question("00-intake", askable)
	_, err := f.r.RecordDecision(key, "00-intake", model.Decision{
		Resolves: q.Key, Withdrawn: true, Rationale: "the feature went away", DecidedBy: "m.example"})
	f.must(err)
	f.finish("00-intake")
	for _, p := range model.Phases[1:5] {
		f.run(p, "")
	}
	if g := f.run("05-review", ""); check(g, "G-Questions").Result != "pass" {
		t.Fatal("a withdrawal did not resolve the question")
	}
}

// The command amends one field. Everything else about the file is somebody else's work: the
// body is the template's and the other fields are the section writer's.
func TestRecordingLeavesTheBodyAndTheOtherFieldsAlone(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	before, template := f.body("00-intake"), f.frontField("00-intake", "template")
	f.question("00-intake", askable)
	if f.body("00-intake") != before {
		t.Fatal("the body changed")
	}
	if got := f.frontField("00-intake", "template"); got != template {
		t.Fatalf("another field changed: %q became %q", template, got)
	}
}

func TestRecordingIntoAPhaseWithNoArtifactIsRefused(t *testing.T) {
	f := newFixture(t)
	_, err := f.r.RecordQuestion(key, "00-intake", []byte(askable))
	if !isRefusal(err) || !strings.Contains(err.Error(), "section set") {
		t.Fatalf("a question was written into a phase nothing has written: %v", err)
	}
}

// ---- WP5: the figure

func TestTheFigureNamesAnIntentThatReachedReviewAskingNothing(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:5] {
		f.run(p, "")
		if s := f.r.Summarise(key); s.AskedNothing() {
			t.Fatalf("the figure appeared at %s, where an open question is still normal", p)
		}
	}
	f.run("05-review", "")
	s := f.r.Summarise(key)
	if s.State != "complete" || !s.AskedNothing() {
		t.Fatalf("an intent that asked nothing reached the merge unmarked: %+v", s)
	}
}

// For the reason G-Questions leaves an abandoned intent alone: abandoning means giving up,
// and a figure that reproached it would only produce an invented question.
func TestTheFigureLeavesAnAbandonedIntentAlone(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:6] {
		f.run(p, "")
	}
	f.write(model.IntentDir(key)+"/intent.yaml",
		"intent: \"git.example/group/proj#1\"\nkey: PROJ-1\nstatus: abandoned\n")
	if s := f.r.Summarise(key); s.AskedNothing() {
		t.Fatalf("the figure marked an abandoned intent: %+v", s)
	}
}

// A decision and no question is a different state from neither, so one entry of either kind
// is enough to answer the figure.
func TestOneEntryOfEitherKindAnswersTheFigure(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	f.run("01-requirements", "decisions:\n  - id: D-1\n    chosen: fail fast\n"+
		"    rationale: the caller retries\n    decided_by: m.example\n")
	for _, p := range model.Phases[2:6] {
		f.run(p, "")
	}
	if s := f.r.Summarise(key); s.AskedNothing() {
		t.Fatalf("an intent with a decision in it was marked as asking nothing: %+v", s)
	}
}

// ---- #229: the three things section 8 asks of a question's options

// bare is askable with the consequences taken off, which is the shape section 8 was written to
// forbid and which was well formed as far as anything could tell until #229.
const bare = `text: which error behaviour?
options:
  - text: fail fast
    recommended: true
  - text: retry internally
  - text: something else, in the words of whoever decides
    free: true
`

// unrecommended is askable with nothing marked, which is XENO-3's Q-2's shape: every
// consequence present and the decision handed back whole.
const unrecommended = `text: which error behaviour?
options:
  - text: fail fast
    consequence: the caller retries
  - text: retry internally
    consequence: the caller never sees it
  - text: something else, in the words of whoever decides
    free: true
`

func TestAnOptionWithNoConsequenceIsRefusedBeforeTheWrite(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	before := f.body("00-intake")

	_, err := f.r.RecordQuestion(key, "00-intake", []byte(bare))
	if !isRefusal(err) || !strings.Contains(err.Error(), "2 of 2 options with no consequence") {
		t.Fatalf("a question with bare options was accepted: %v", err)
	}
	var o model.Output
	f.readFront("00-intake", &o)
	if len(o.OpenQuestions) != 0 || f.body("00-intake") != before {
		t.Fatal("a refused question changed the artifact")
	}
}

// The free entry carries no consequence and cannot: it stands for an answer nobody has written
// yet. askable is the project's own well-formed question and its free entry is bare, so a check
// that required one would refuse the fixture rather than catch anything.
func TestTheFreeEntryNeedsNoConsequence(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	f.question("00-intake", askable)
}

// The recommendation is the writer's check and not the gate's, for the reason QuestionShape's
// comment carries: the gate judges every artifact in the trail and XENO-3's Q-2 fails this.
func TestAQuestionThatRecommendsNothingIsRefusedBeforeTheWrite(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")

	_, err := f.r.RecordQuestion(key, "00-intake", []byte(unrecommended))
	if !isRefusal(err) || !strings.Contains(err.Error(), "recommends 0 of its options") {
		t.Fatalf("a question recommending nothing was accepted: %v", err)
	}
}

// Three recommendations is no recommendation. Section 8 says "the agent's recommendation",
// singular, and a list with everything starred has made no choice.
func TestAQuestionThatRecommendsMoreThanOneIsRefused(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	two := strings.Replace(unrecommended,
		"    consequence: the caller never sees it\n",
		"    consequence: the caller never sees it\n    recommended: true\n", 1)
	two = strings.Replace(two,
		"    consequence: the caller retries\n",
		"    consequence: the caller retries\n    recommended: true\n", 1)

	_, err := f.r.RecordQuestion(key, "00-intake", []byte(two))
	if !isRefusal(err) || !strings.Contains(err.Error(), "recommends 2 of its options") {
		t.Fatalf("a question recommending two options was accepted: %v", err)
	}
}

// no_options is the exception section 8 names, and a question with no options to offer cannot
// recommend one. Exempt from both halves, here as in the gate.
func TestNoOptionsIsExemptFromBothChecks(t *testing.T) {
	f := newFixture(t)
	f.started("00-intake")
	f.question("00-intake", "text: who owns the schedule?\nno_options: true\n")
}
