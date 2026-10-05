// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"sort"
	"strings"

	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/rules"
)

// The review checklist section 9 defines, and the command that writes it. One entry answers
// one review rule of the effective set, and G-Policy counts from the rules to the entries,
// so an unanswered rule is a finding and an empty checklist is a red phase rather than a
// pass.
//
// It is kept out of exchange.go, which is section 8's pair of a question and the decision
// that settles it. This answers a rule, not a question. The two helpers are shared, because
// amending one frontmatter field and leaving the body alone is the same act here: going
// through SectionSet would re-render the body from the template, which is the template's
// business and not an answer's.
//
// Like those two, this does not look at gate.yaml. A write after a verdict changes the
// artifact and so changes its hash, and the next phase finish says the phase changed after
// its verdict was written, which is #216's mechanism. #242 asked for a refusal here and gave
// the decision commands as the precedent; D-6 is the decision it was reading, and that one
// is about gate approve and gate override, which carry `against`, the artifacts_hash the
// decision was taken on. A checklist entry carries no such field, so it asserts nothing a
// recompute could contradict.

// ReviewAnswered is what one answer did, which the command reports so that a checklist can be
// finished without reading the gate's findings to find out what is missing. It is recorded
// nowhere: the artifact holds the answers, and the set outstanding at one moment describes
// neither the phase nor its verdict.
type ReviewAnswered struct {
	Entry      model.ChecklistEntry
	Replaced   bool     // an answer for this rule was already there and was written over
	Unanswered []string // review rules of the effective set with no entry yet, sorted
}

// ReviewAnswer writes one checklist entry into the review phase's artifact.
//
// Flags rather than a document, because an entry is three scalars; decision record and
// assumption record take their flat records the same way, and the nested shapes are what
// question record and scope set read from stdin for.
//
// There is no phase argument. G-Policy judges the checklist only on the last phase, so the
// field belongs to one phase and a flag would offer a choice the format does not have. This
// is ScopeSet's reasoning in the other direction: the scope is P0's artifact and the
// checklist is P5's.
//
// An answer for a rule already answered replaces that entry where it stands, rather than
// appending a second. Nothing records when an entry was written, so two entries for one rule
// are an ambiguity and not a history, and keeping the index means a corrected answer is a
// one-line diff rather than a reordering.
func (r *Runner) ReviewAnswer(key string, e model.ChecklistEntry) (*ReviewAnswered, error) {
	// The checks that need neither the rule tree nor the artifact come first, so a malformed
	// call cannot leave a half-amended file behind.
	if strings.TrimSpace(e.Rule) == "" {
		return nil, refuse("a review answer names the rule it answers: the id as the effective set spells it")
	}
	if e.Source != "" {
		return nil, refuse("a review answer carries no source: section 12 gives source to a lens entry, " +
			"which answers no rule and is written by the lens")
	}
	switch {
	case e.Result == "":
		return nil, refuse("--result is required: %s", strings.Join(model.ChecklistResults, ", "))
	case !model.OneOf(e.Result, model.ChecklistResults):
		return nil, refuse("%q is no checklist result; section 9 fixes the three: %s",
			e.Result, strings.Join(model.ChecklistResults, ", "))
	case model.OneOf(e.Result, model.ChecklistNeedsNote) && strings.TrimSpace(e.Note) == "":
		return nil, refuse("--note is required for %s: write why the rule was passed over; "+
			"met is the only result that needs none", e.Result)
	}
	// The set the gate counts against. Shared with it rather than restated, so that a refusal
	// on the way in and a finding after the fact cannot disagree about what a rule is.
	//
	// A74 is the limit on what this can promise: G-Policy judges a phase against the set the
	// artifact recorded in rules_hash, and this is the set that resolves now. They agree
	// unless the rule tree changed since section set wrote the hash, and where they diverge
	// the gate is right and this refusal was advice.
	read, _ := rules.Load(r.Root)
	effective, _ := rules.Effective(read)
	var review []string
	for _, rl := range effective {
		if rl.Kind == rules.Review {
			review = append(review, rl.ID)
		}
	}
	if !model.OneOf(e.Rule, review) {
		if len(review) == 0 {
			return nil, refuse("the effective rule set holds no review rule, so there is nothing to answer; "+
				"%q cannot be one of them", e.Rule)
		}
		sort.Strings(review)
		return nil, refuse("%q is no review rule of the effective set; it holds %s",
			e.Rule, strings.Join(review, ", "))
	}
	phase := model.Phases[len(model.Phases)-1]
	front, o, body, err := r.artifact(key, phase)
	if err != nil {
		return nil, err
	}
	list, replaced := upsertChecklist(o.ReviewChecklist, e)
	if err := r.amendFront(key, phase, front, "review_checklist", list, body); err != nil {
		return nil, err
	}
	return &ReviewAnswered{Entry: e, Replaced: replaced, Unanswered: unanswered(review, list)}, nil
}

// upsertChecklist replaces the entry answering the same rule, or appends. The index is kept
// so that re-answering a rule does not reorder the block.
func upsertChecklist(list []model.ChecklistEntry, e model.ChecklistEntry) ([]model.ChecklistEntry, bool) {
	out := make([]model.ChecklistEntry, len(list))
	copy(out, list)
	for i, x := range out {
		if x.Rule == e.Rule {
			out[i] = e
			return out, true
		}
	}
	return append(out, e), false
}

// unanswered counts from the rules to the entries, as G-Policy does. Walking the entries
// instead would report an empty checklist as complete, which is the reading the gate's own
// comment warns against.
func unanswered(review []string, list []model.ChecklistEntry) []string {
	answered := map[string]bool{}
	for _, e := range list {
		if e.Rule != "" {
			answered[e.Rule] = true
		}
	}
	var out []string
	for _, id := range review {
		if !answered[id] {
			out = append(out, id)
		}
	}
	sort.Strings(out)
	return out
}
