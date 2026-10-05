// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
)

// Section 8 asks three things of a question's options: a consequence on each, the agent's
// recommendation with a reason, and always a free entry. The count and the free entry were
// read; the consequence is read from #229; the recommendation is read by QuestionAsked, which
// only the writer calls, and the reason has no field to be read from.
//
// The split is the subject of TestTheGateIsSilentOnTheRecommendation, which exists so that
// moving the check fails a test rather than failing `xeno gate verify` over the trail.

func option(text, consequence string, recommended bool) model.Option {
	return model.Option{Text: text, Consequence: consequence, Recommended: recommended}
}

// asked wraps one question the way both shape functions take it.
func asked(q model.Question) model.Output {
	return model.Output{OpenQuestions: []model.Question{q}}
}

// well is a question in the shape section 8 asks for, which both functions accept.
func well() model.Question {
	return model.Question{
		Key:  "Q-1",
		Text: "which error behaviour?",
		Options: []model.Option{
			option("fail fast", "the caller retries", true),
			option("retry internally", "the caller never sees it", false),
			{Text: "something else, in the words of whoever decides", Free: true},
		},
	}
}

func TestAnOptionWithoutAConsequenceIsAFinding(t *testing.T) {
	q := well()
	q.Options[1].Consequence = "   "

	fs := QuestionShape("output.md", asked(q))
	if len(fs) != 1 {
		t.Fatalf("want one finding, got %d: %+v", len(fs), fs)
	}
	if !strings.Contains(fs[0].Cause, "1 of 2 options with no consequence") {
		t.Fatalf("the finding does not say how many: %s", fs[0].Cause)
	}
}

// The free entry stands for an answer nobody has written yet, so it carries no consequence and
// is outside the set section 8 asks one of.
func TestTheFreeEntryIsNotAskedForAConsequence(t *testing.T) {
	if fs := QuestionShape("output.md", asked(well())); len(fs) != 0 {
		t.Fatalf("a well-formed question was a finding: %+v", fs)
	}
}

// **The asymmetry, pinned.** A question recommending nothing is not the gate's finding, because
// this function is reached through phaseResult and so through G-Schema, which runs from P0:
// XENO-3's Q-2 recommends no option, and a check here makes `xeno gate verify` report
// "DIVERGENT XENO-3 00-intake: committed status green, recomputed red" and exit 1. What is
// sealed is never rewritten, so #229 put the check in the writer.
//
// If this test fails because the gate now reports it, the trail is what to check before the
// test: either a sealed verdict has been released, or `gate verify` is about to break.
func TestTheGateIsSilentOnTheRecommendation(t *testing.T) {
	q := well()
	q.Options[0].Recommended = false

	if fs := QuestionShape("output.md", asked(q)); len(fs) != 0 {
		t.Fatalf("the gate reported the recommendation, which re-judges the trail: %+v", fs)
	}
	fs := QuestionAsked("output.md", asked(q))
	if len(fs) != 1 || !strings.Contains(fs[0].Cause, "recommends 0 of its options") {
		t.Fatalf("the writer did not require a recommendation: %+v", fs)
	}
}

// Exactly one: section 8 says "the agent's recommendation", and a list with everything starred
// has made no choice.
func TestTheWriterWantsExactlyOneRecommendation(t *testing.T) {
	q := well()
	q.Options[1].Recommended = true

	fs := QuestionAsked("output.md", asked(q))
	if len(fs) != 1 || !strings.Contains(fs[0].Cause, "recommends 2 of its options") {
		t.Fatalf("two recommendations were accepted: %+v", fs)
	}
}

// QuestionAsked is QuestionShape plus one check, so the half they share cannot drift. A question
// that is wrong in both ways is reported for both, by one function.
func TestTheWriterAddsToTheGatesChecksRatherThanReplacingThem(t *testing.T) {
	q := well()
	q.Options[0].Recommended = false
	q.Options[1].Consequence = ""

	fs := QuestionAsked("output.md", asked(q))
	if len(fs) != 2 {
		t.Fatalf("want the consequence and the recommendation, got %d: %+v", len(fs), fs)
	}
	joined := fs[0].Cause + " " + fs[1].Cause
	for _, want := range []string{"no consequence", "recommends 0"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("the findings do not cover %q: %s", want, joined)
		}
	}
}

// The exception section 8 names. A question with no options to offer has none to carry a
// consequence and none to recommend, and both functions leave it alone.
func TestNoOptionsIsExemptFromBoth(t *testing.T) {
	q := model.Question{Key: "Q-1", Text: "who owns the schedule?", NoOptions: true}

	if fs := QuestionShape("output.md", asked(q)); len(fs) != 0 {
		t.Fatalf("no_options was a finding for the gate: %+v", fs)
	}
	if fs := QuestionAsked("output.md", asked(q)); len(fs) != 0 {
		t.Fatalf("no_options was a finding for the writer: %+v", fs)
	}
}
