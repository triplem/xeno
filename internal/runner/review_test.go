// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
)

// artifactText reads a phase's output.md whole, for the one assertion that is about the bytes
// rather than about the parse.
func (f *fixture) artifactText(phase string) string {
	f.t.Helper()
	b, err := os.ReadFile(filepath.Join(f.root, model.PhaseDir(key, phase), "output.md"))
	f.must(err)
	return string(b)
}

// The checklist section 9 defines was the last of the four frontmatter lists with no writer,
// and the only field a gate refuses the phase without that a person had to produce by editing
// a sealed artifact (#242). These cover the writer and the three refusals it moves from the
// verdict to the moment of the mistake.

// review is the phase the checklist belongs to. Named once here because the writer resolves it
// rather than taking it, which is the behaviour TestTheChecklistIsTheReviewPhases asserts.
//
// Its artifact is written with f.output rather than f.started, because Start enforces the phase
// sequence and P5 cannot be reached without a verdict on the four before it. What these tests
// need is the artifact the writer amends, not the sequence that leads to it.
var review = model.Phases[len(model.Phases)-1]

// reviewRules lays down a given/builtin tree holding two review rules and one checked rule.
// The checked one is there to be left out: G-Policy counts review rules and so does the
// writer, and a checked rule in the answered set would be a rule answered twice over.
func (f *fixture) reviewRules() {
	f.t.Helper()
	rule := func(id, kind string) {
		f.write(".xeno/plugin/rules/given/builtin/"+id+".yaml",
			"id: "+id+"\nversion: 1\nscope: builtin\nbinding: false\nkind: "+kind+
				"\napplies_to: [05-review]\nstatement: >\n  "+id+" says so.\n")
	}
	rule("deviations-are-traceable", "review")
	rule("interface-change-needs-a-migration-note", "review")
	rule("release-notes-are-filled", "checked")
}

func (f *fixture) answer(e model.ChecklistEntry) *ReviewAnswered {
	f.t.Helper()
	a, err := f.r.ReviewAnswer(key, e)
	f.must(err)
	return a
}

// TestAnAnswerIsWrittenAndWhatIsLeftIsReported is the writer's point. The report is why the
// command exists rather than a bare write: a checklist finished by reading phase finish's
// findings is the loop #242 was filed to close.
func TestAnAnswerIsWrittenAndWhatIsLeftIsReported(t *testing.T) {
	f := newFixture(t)
	f.reviewRules()
	f.output(review, "")

	a := f.answer(model.ChecklistEntry{Rule: "deviations-are-traceable", Result: "met"})
	if a.Replaced {
		t.Fatal("the first answer to a rule replaced something")
	}
	if len(a.Unanswered) != 1 || a.Unanswered[0] != "interface-change-needs-a-migration-note" {
		t.Fatalf("what is left was not reported: %+v", a.Unanswered)
	}

	var o model.Output
	f.readFront(review, &o)
	if len(o.ReviewChecklist) != 1 {
		t.Fatalf("the entry was not written: %+v", o.ReviewChecklist)
	}
	if e := o.ReviewChecklist[0]; e.Rule != "deviations-are-traceable" || e.Result != "met" {
		t.Fatalf("the entry was not written as given: %+v", e)
	}

	// And the checked rule is not among what is owed, because a checked rule is not answered
	// here at all.
	a = f.answer(model.ChecklistEntry{Rule: "interface-change-needs-a-migration-note",
		Result: "not-applicable", Note: "no interface changes"})
	if len(a.Unanswered) != 0 {
		t.Fatalf("a checked rule was counted as owing an answer: %+v", a.Unanswered)
	}
}

// Replacing rather than appending is the maintainer's decision on the one question #242 left
// open, and keeping the index is what makes a corrected answer a one-line diff. Two entries for
// one rule would be an ambiguity and not a history: nothing records which was written first.
func TestAnsweringARuleTwiceReplacesTheAnswerInPlace(t *testing.T) {
	f := newFixture(t)
	f.reviewRules()
	f.output(review, "")

	f.answer(model.ChecklistEntry{Rule: "deviations-are-traceable", Result: "met"})
	f.answer(model.ChecklistEntry{Rule: "interface-change-needs-a-migration-note",
		Result: "not-applicable", Note: "none"})
	a := f.answer(model.ChecklistEntry{Rule: "deviations-are-traceable",
		Result: "deviation", Note: "one, against criterion 3"})
	if !a.Replaced {
		t.Fatal("re-answering a rule did not report a replacement")
	}

	var o model.Output
	f.readFront(review, &o)
	if len(o.ReviewChecklist) != 2 {
		t.Fatalf("re-answering appended instead of replacing: %+v", o.ReviewChecklist)
	}
	// The position is the one it was first answered in, so the block does not reorder.
	if e := o.ReviewChecklist[0]; e.Rule != "deviations-are-traceable" ||
		e.Result != "deviation" || e.Note != "one, against criterion 3" {
		t.Fatalf("the replacement did not keep its place or its content: %+v", o.ReviewChecklist)
	}
}

// The three refusals are G-Policy's own checks moved to the moment of writing. The gate stays
// the authority on a sealed phase; what this removes is finding out from a red verdict.
func TestAnAnswerThatTheGateWouldRejectIsRefused(t *testing.T) {
	cases := []struct {
		name  string
		entry model.ChecklistEntry
		says  string
	}{
		{"a rule outside the effective set", model.ChecklistEntry{
			Rule: "invented-rule", Result: "met"}, "no review rule of the effective set"},
		{"a checked rule, which is answered by its check and not here", model.ChecklistEntry{
			Rule: "release-notes-are-filled", Result: "met"}, "no review rule of the effective set"},
		{"no rule at all", model.ChecklistEntry{Result: "met"}, "names the rule it answers"},
		{"no result", model.ChecklistEntry{
			Rule: "deviations-are-traceable"}, "--result is required"},
		{"a result outside section 9's three", model.ChecklistEntry{
			Rule: "deviations-are-traceable", Result: "partly"}, "no checklist result"},
		{"a deviation with no note", model.ChecklistEntry{
			Rule: "deviations-are-traceable", Result: "deviation"}, "--note is required"},
		{"not-applicable with a blank note", model.ChecklistEntry{
			Rule: "deviations-are-traceable", Result: "not-applicable", Note: "  "}, "--note is required"},
		{"a source, which belongs to a lens entry", model.ChecklistEntry{
			Rule: "deviations-are-traceable", Result: "met", Source: "lens"}, "carries no source"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := newFixture(t)
			f.reviewRules()
			f.output(review, "")

			_, err := f.r.ReviewAnswer(key, c.entry)
			if !isRefusal(err) {
				t.Fatalf("want a refusal, got %v", err)
			}
			if !strings.Contains(err.Error(), c.says) {
				t.Fatalf("the refusal does not say why: %v", err)
			}
			// Nothing was written, which is why the cheap checks run before the file is opened.
			var o model.Output
			f.readFront(review, &o)
			if len(o.ReviewChecklist) != 0 {
				t.Fatalf("a refused answer was written anyway: %+v", o.ReviewChecklist)
			}
		})
	}
}

// met is the only result that needs no note, which is model.ChecklistNeedsNote's split and the
// one case where an absent note is the right answer rather than an omission.
func TestMetNeedsNoNote(t *testing.T) {
	f := newFixture(t)
	f.reviewRules()
	f.output(review, "")

	f.answer(model.ChecklistEntry{Rule: "deviations-are-traceable", Result: "met"})
	var o model.Output
	f.readFront(review, &o)
	if o.ReviewChecklist[0].Note != "" {
		t.Fatalf("a note was invented: %+v", o.ReviewChecklist[0])
	}
}

// The writer amends one field, so the body it did not write comes back byte for byte. Going
// through SectionSet would re-render it from the template, which is what exchange.go gives as
// the reason these frontmatter writers exist at all.
func TestTheBodySurvivesAnAnswer(t *testing.T) {
	f := newFixture(t)
	f.reviewRules()
	f.output(review, "")
	before := f.artifactText(review)
	_, body, ok := strings.Cut(before, "---\n\n")
	if !ok {
		t.Fatalf("the fixture's artifact has no body to keep: %q", before)
	}

	f.answer(model.ChecklistEntry{Rule: "deviations-are-traceable", Result: "met"})

	after := f.artifactText(review)
	if !strings.HasSuffix(after, body) {
		t.Fatalf("the body did not survive: want suffix %q, got %q", body, after)
	}
}

// The checklist is the review phase's artifact and the writer resolves it, which is why the
// command takes no phase: G-Policy judges the checklist only on the last phase, so a flag would
// offer a choice of one. An answer given while another phase is open still lands in P5.
func TestTheChecklistIsTheReviewPhases(t *testing.T) {
	f := newFixture(t)
	f.reviewRules()
	f.started(model.Phases[0])
	f.output(review, "")

	f.answer(model.ChecklistEntry{Rule: "deviations-are-traceable", Result: "met"})

	var intake model.Output
	f.readFront(model.Phases[0], &intake)
	if len(intake.ReviewChecklist) != 0 {
		t.Fatalf("the answer landed in %s: %+v", model.Phases[0], intake.ReviewChecklist)
	}
	var rev model.Output
	f.readFront(review, &rev)
	if len(rev.ReviewChecklist) != 1 {
		t.Fatalf("the answer did not land in %s: %+v", review, rev.ReviewChecklist)
	}
}

// An artifact that does not exist yet is the refusal the other frontmatter writers give, and for
// the same reason: section set creates it, and an answer has nothing to amend before then.
func TestAnAnswerBeforeTheArtifactIsRefused(t *testing.T) {
	f := newFixture(t)
	f.reviewRules()

	_, err := f.r.ReviewAnswer(key, model.ChecklistEntry{
		Rule: "deviations-are-traceable", Result: "met"})
	if !isRefusal(err) {
		t.Fatalf("want a refusal, got %v", err)
	}
	if !strings.Contains(err.Error(), "xeno section set") {
		t.Fatalf("the refusal does not say what writes it: %v", err)
	}
}

// A tree with no review rule at all is its own refusal, because the message that lists what the
// set holds would otherwise list nothing and read as though the rule were misspelled.
func TestAnAnswerAgainstARuleSetWithNoReviewRuleIsRefused(t *testing.T) {
	f := newFixture(t)
	f.output(review, "")

	_, err := f.r.ReviewAnswer(key, model.ChecklistEntry{
		Rule: "deviations-are-traceable", Result: "met"})
	if !isRefusal(err) {
		t.Fatalf("want a refusal, got %v", err)
	}
	if !strings.Contains(err.Error(), "holds no review rule") {
		t.Fatalf("the refusal does not say the set is empty: %v", err)
	}
}

// ---- the lens entry section 12 allows, and the command that writes it

// The entry the four lens skills describe. Section 12 fixes what it carries: source: lens and
// no rule id, which is what keeps it out of the set G-Policy counts for completeness.
func TestALensEntryIsWrittenWithItsSourceAndNoRule(t *testing.T) {
	f := newFixture(t)
	f.reviewRules()
	f.output(review, "")

	n, err := f.r.ReviewLens(key, "deviation",
		"the security lens: the error path reports more than its caller may know")
	f.must(err)
	if n.Lens != 1 {
		t.Fatalf("the count of lens entries is %d after the first: %+v", n.Lens, n)
	}

	var o model.Output
	f.readFront(review, &o)
	if len(o.ReviewChecklist) != 1 {
		t.Fatalf("the entry was not written: %+v", o.ReviewChecklist)
	}
	e := o.ReviewChecklist[0]
	if e.Rule != "" {
		t.Errorf("the entry answers a rule, which is what keeps it out of the counted set: %+v", e)
	}
	if e.Source != model.ChecklistSourceLens {
		t.Errorf("the entry does not carry section 12's source: %+v", e)
	}
	if e.Result != "deviation" || !strings.Contains(e.Note, "security lens") {
		t.Errorf("the entry was not written as given: %+v", e)
	}
}

// An answer replaces the entry for its rule and a lens entry cannot: a rule id is what makes
// two answers to one thing recognisable as such. So four lenses leave four entries, and neither
// writer disturbs the other's.
func TestLensEntriesSitBesideEachOtherAndBesideTheAnswers(t *testing.T) {
	f := newFixture(t)
	f.reviewRules()
	f.output(review, "")

	f.answer(model.ChecklistEntry{Rule: "deviations-are-traceable", Result: "met"})
	for _, lens := range []string{"security", "privacy", "operations", "architecture"} {
		n, err := f.r.ReviewLens(key, "deviation", "the "+lens+" lens had something to say")
		f.must(err)
		if n.Entry.Note != "the "+lens+" lens had something to say" {
			t.Fatalf("the entry reported is not the one given: %+v", n.Entry)
		}
	}
	a := f.answer(model.ChecklistEntry{Rule: "interface-change-needs-a-migration-note",
		Result: "not-applicable", Note: "no interface changed"})
	if len(a.Unanswered) != 0 {
		t.Fatalf("the lens entries were counted as answers or as debts: %+v", a.Unanswered)
	}

	var o model.Output
	f.readFront(review, &o)
	if len(o.ReviewChecklist) != 6 {
		t.Fatalf("four lenses and two rules are six entries: %+v", o.ReviewChecklist)
	}
	// The two answers keep the places they were written in, which is what says the lens
	// entries were appended rather than written over one of them.
	if o.ReviewChecklist[0].Rule != "deviations-are-traceable" ||
		o.ReviewChecklist[5].Rule != "interface-change-needs-a-migration-note" {
		t.Fatalf("the answers moved: %+v", o.ReviewChecklist)
	}
}

// Section 9's result is required of a lens entry as it is of an answer, because G-Policy's one
// check on the list is that every entry carries one, lens entries included. The note is required
// of every result and not only of the two: the entry has no rule id, so without a note nothing
// in it says which lens wrote it or what it found.
func TestALensEntryThatSaysNothingIsRefused(t *testing.T) {
	cases := []struct {
		name         string
		result, note string
		says         string
	}{
		{"no result, which is the one thing G-Policy checks", "", "the lens found something",
			"--result is required"},
		{"a result outside section 9's three", "partly", "the lens found something",
			"no checklist result"},
		{"a deviation with no note", "deviation", "", "--note is required"},
		{"met with no note, which an answer may give and this may not", "met", "",
			"--note is required for a lens entry"},
		{"met with a blank note", "met", "  ", "--note is required for a lens entry"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := newFixture(t)
			f.reviewRules()
			f.output(review, "")

			_, err := f.r.ReviewLens(key, c.result, c.note)
			if !isRefusal(err) {
				t.Fatalf("want a refusal, got %v", err)
			}
			if !strings.Contains(err.Error(), c.says) {
				t.Fatalf("the refusal does not say why: %v", err)
			}
			var o model.Output
			f.readFront(review, &o)
			if len(o.ReviewChecklist) != 0 {
				t.Fatalf("a refused entry was written anyway: %+v", o.ReviewChecklist)
			}
		})
	}
}

// A lens entry answers no rule, so it is writable where the rule tree holds no review rule at
// all. That is the other half of its being outside the counted set, and it is the case an answer
// is refused for.
func TestALensEntryNeedsNoReviewRuleInTheSet(t *testing.T) {
	f := newFixture(t)
	f.output(review, "")

	_, err := f.r.ReviewLens(key, "deviation", "the operations lens: nothing sets the new variable")
	f.must(err)

	var o model.Output
	f.readFront(review, &o)
	if len(o.ReviewChecklist) != 1 {
		t.Fatalf("the entry was not written: %+v", o.ReviewChecklist)
	}
}

// The refusal the other frontmatter writers give, for the same reason: section set creates the
// artifact and there is nothing to amend before then. Held to the standard #242 set for the
// answer, since both writers amend the same field of the same file.
func TestALensEntryBeforeTheArtifactIsRefused(t *testing.T) {
	f := newFixture(t)

	_, err := f.r.ReviewLens(key, "deviation", "the privacy lens: the log carries an address")
	if !isRefusal(err) {
		t.Fatalf("want a refusal, got %v", err)
	}
	if !strings.Contains(err.Error(), "xeno section set") {
		t.Fatalf("the refusal does not say what writes it: %v", err)
	}
}

// The body is not this writer's business either, which is TestTheBodySurvivesAnAnswer's reason:
// the prose of the review-checklist section is what section set wrote, and a structured entry
// amends one frontmatter field beside it.
func TestTheBodySurvivesALensEntry(t *testing.T) {
	f := newFixture(t)
	f.reviewRules()
	f.output(review, "")
	before := f.artifactText(review)
	_, body, ok := strings.Cut(before, "---\n\n")
	if !ok {
		t.Fatalf("the fixture's artifact has no body to keep: %q", before)
	}

	_, err := f.r.ReviewLens(key, "met", "the architecture lens read the structure and it held")
	f.must(err)

	if after := f.artifactText(review); !strings.HasSuffix(after, body) {
		t.Fatalf("the body did not survive: want suffix %q, got %q", body, after)
	}
}

// The refusal of a source on an answer now names the command that writes one. It named "the
// lens" before, which resolved to nothing: there was no writer, and the four lens skills wrote
// their findings as prose because of it.
func TestTheAnswersRefusalOfASourceNamesTheWriterThatExists(t *testing.T) {
	f := newFixture(t)
	f.reviewRules()
	f.output(review, "")

	_, err := f.r.ReviewAnswer(key, model.ChecklistEntry{
		Rule: "deviations-are-traceable", Result: "met", Source: model.ChecklistSourceLens})
	if !isRefusal(err) {
		t.Fatalf("want a refusal, got %v", err)
	}
	if !strings.Contains(err.Error(), "xeno review lens") {
		t.Fatalf("the refusal does not name the command that writes one: %v", err)
	}
}
