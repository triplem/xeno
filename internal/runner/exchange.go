// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"go.yaml.in/yaml/v3"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/gates"
	"github.com/triplem/xeno/internal/model"
)

// The exchange section 8 defines: a question asked with options, and the decision that
// settles it. Both are frontmatter blocks of the phase that holds them, sealed with it,
// which is why neither lives at intent level the way an assumption does — only an
// assumption has a state that changes.
//
// Both are written by amending one frontmatter field and nothing else. Going through
// SectionSet would re-render the body from the template, which is the template's business
// and not a question's, and the body of a phase somebody is in the middle of writing is
// not something these two commands have anything to say about.
//
// Neither looks at gate.yaml. A write after a verdict changes the artifact and so changes
// its hash, and the next phase finish says the phase changed after its verdict was
// written, which is the state #216 built and the same answer section set and learning
// record already give.

// RecordQuestion appends one question to the phase's open_questions block.
//
// The entry arrives as the YAML section 5 defines, because up to five options with a
// consequence each do not fit flags without inventing a separator, and because the shape
// read is then the shape written: the gate's own checks are the only authority on what a
// question is, and nothing can be well formed on the way in and malformed in the file.
// Unknown fields are refused rather than dropped, which is the standing rule about
// invented fields applied where a typo would otherwise pass as silence.
func (r *Runner) RecordQuestion(key, phase string, entry []byte) (*model.Question, error) {
	dec := yaml.NewDecoder(bytes.NewReader(entry))
	dec.KnownFields(true)
	var q model.Question
	if err := dec.Decode(&q); err != nil {
		return nil, refuse("the question could not be read as an open_questions entry: %v", err)
	}
	if q.Key != "" {
		return nil, refuse("a question's key is assigned and not given: leave key out and the next Q-n of the intent is written")
	}
	if strings.TrimSpace(q.Text) == "" {
		return nil, refuse("a question needs text: what is being asked, in the words it is being asked in")
	}
	front, o, body, err := r.artifact(key, phase)
	if err != nil {
		return nil, err
	}
	raised, _ := r.exchange(key)
	used := make([]string, 0, len(raised))
	for _, e := range raised {
		used = append(used, e.Key)
	}
	q.Key = nextExchangeID("Q", used)
	// QuestionAsked and not QuestionShape: the writer requires the recommendation section 8
	// asks for, and the gate does not, because a check added to the gate's half is applied to
	// every artifact in the trail and one sealed question fails it (#229). The asymmetry is
	// deliberate and QuestionShape's comment carries the measurement.
	if err := shapeRefusal(gates.QuestionAsked(
		model.PhaseDir(key, phase)+"/output.md", model.Output{OpenQuestions: []model.Question{q}})); err != nil {
		return nil, err
	}
	if err := r.amendFront(key, phase, front, "open_questions", append(o.OpenQuestions, q), body); err != nil {
		return nil, err
	}
	return &q, nil
}

// RecordDecision appends one decision to the phase's decisions block.
//
// Flags rather than a document, because the entry is six scalars; `assumption record` takes
// its flat record the same way. The person is never defaulted, to the git user or to
// anything else: A35's rule is that a plausible value in a field nobody produced is worse
// than an absent one, and for decided_by the stakes are the whole register, since who
// decided is the question it exists to answer.
func (r *Runner) RecordDecision(key, phase string, d model.Decision) (*model.Decision, error) {
	if d.ID != "" {
		return nil, refuse("a decision's id is assigned and not given: the next D-n of the intent is written")
	}
	if d.DecidedBy == "" {
		return nil, refuse("--by is required: taking a decision is a statement by a person")
	}
	if strings.TrimSpace(d.Rationale) == "" {
		return nil, refuse("--reason is required: a decision carries why it went that way, and a withdrawal carries why it was dropped")
	}
	switch {
	case d.Withdrawn && d.Chosen != "":
		return nil, refuse("a withdrawal has no chosen option: --withdraw says the question is dropped and --reason says why")
	case !d.Withdrawn && d.Chosen == "":
		return nil, refuse("--chosen is required: the option taken, in the words the person deciding used")
	}
	front, o, body, err := r.artifact(key, phase)
	if err != nil {
		return nil, err
	}
	raised, taken := r.exchange(key)
	if d.Resolves != "" && !raisedIn(raised, d.Resolves) {
		return nil, refuse("%s has raised no question %s; --resolves names a question of this intent, in any phase", key, d.Resolves)
	}
	used := make([]string, 0, len(taken))
	for _, e := range taken {
		used = append(used, e.ID)
	}
	d.ID = nextExchangeID("D", used)
	if err := shapeRefusal(gates.DecisionShape(
		model.PhaseDir(key, phase)+"/output.md", model.Output{Decisions: []model.Decision{d}})); err != nil {
		return nil, err
	}
	if err := r.amendFront(key, phase, front, "decisions", append(o.Decisions, d), body); err != nil {
		return nil, err
	}
	return &d, nil
}

// exchange is what the intent has raised and what it has settled, over every phase.
//
// The gate asks both of one file, because a duplicate key within a file is the duplicate it
// can see. A command can ask it of the intent and has to: `resolves` is read from any later
// phase, so two phases carrying a Q-1 each would make a reference to one of them ambiguous.
// It is also where the figure in the listing comes from, which is the other half of #188.
func (r *Runner) exchange(key string) (raised []model.Question, taken []model.Decision) {
	for _, p := range model.Phases {
		var o model.Output
		if _, err := fm.ReadFront(r.abs(model.PhaseDir(key, p)+"/output.md"), &o); err != nil {
			continue // a phase nothing has written yet, which is ordinary
		}
		raised = append(raised, o.OpenQuestions...)
		taken = append(taken, o.Decisions...)
	}
	return raised, taken
}

func raisedIn(raised []model.Question, key string) bool {
	for _, q := range raised {
		if q.Key == key {
			return true
		}
	}
	return false
}

// nextExchangeID continues the intent's own numbering rather than counting entries, so that
// a question dropped from a draft does not hand its key to the next one. It is
// nextAssumptionID's reason, and the ids are per intent as the register's are; the D- table
// of docs/assumptions.md is a separate sequence in a separate file, exactly as its A80 is
// separate from a register's A-001.
//
// Unpadded, because the four questions the trail already holds are Q-1 and the decisions
// recorded by hand are D-1 to D-7, and a key that sorts differently from the ones a reader
// has seen is a second scheme rather than a continuation.
func nextExchangeID(prefix string, used []string) string {
	high := 0
	for _, u := range used {
		var n int
		if _, err := fmt.Sscanf(u, prefix+"-%d", &n); err == nil && n > high {
			high = n
		}
	}
	return fmt.Sprintf("%s-%d", prefix, high+1)
}

// shapeRefusal turns what the gate would have reported into a refusal before the write.
// The first finding is the whole answer: both checks report one finding per entry and
// exactly one entry is being judged.
func shapeRefusal(fs []model.Finding) error {
	if len(fs) == 0 {
		return nil
	}
	return refuse("%s; %s", fs[0].Cause, fs[0].Next)
}

// artifact reads a phase's output.md three ways: the map SectionSet carries its frontmatter
// over in, the lists these two commands append to, and the body they write back untouched.
//
// It refuses where there is no file. The artifact is created by its first section write, and
// a question belongs to a phase that has one; creating it here would write a frontmatter
// carrying whatever the runner could work out on its own, which is what SectionSet's comment
// about leaving fields out rather than filling them is about.
func (r *Runner) artifact(key, phase string) (map[string]any, model.Output, []byte, error) {
	var o model.Output
	rel := model.PhaseDir(key, phase) + "/output.md"
	b, err := os.ReadFile(r.contentSource(key, phase, "output.md"))
	if err != nil {
		return nil, o, nil, refuse("%s has no %s yet; it is written by xeno section set, which creates the artifact", phase, rel)
	}
	f, body, err := fm.Split(b)
	if err != nil {
		return nil, o, nil, refuse("%s has no frontmatter to amend", rel)
	}
	front := map[string]any{}
	if err := yaml.Unmarshal(f, &front); err != nil {
		return nil, o, nil, refuse("the frontmatter of %s cannot be read: %v", rel, err)
	}
	if err := yaml.Unmarshal(f, &o); err != nil {
		return nil, o, nil, refuse("the frontmatter of %s cannot be read: %v", rel, err)
	}
	return front, o, body, nil
}

// amendFront replaces one frontmatter field and writes the file again, body byte for byte.
// The order is frontmatterOrder's, which is where section 5's field order lives and which
// already places both of these keys.
func (r *Runner) amendFront(key, phase string, front map[string]any, field string, value any, body []byte) error {
	front[field] = value
	out := "---\n" + frontmatter(front) + "---\n" + string(body)
	return os.WriteFile(r.contentPath(key, phase, "output.md"), []byte(out), 0o644)
}
