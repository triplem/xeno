// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/triplem/xeno/internal/cost"
	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/template"
)

// The three figures the plan says will decide proportionality are cost per intent,
// artifacts per intent and reopens per intent, and until now all three were counted by hand
// into comments on #117. This file derives them from the trail instead, so that the same
// question asked at any commit gets the same answer out of the same tree.
//
// It reads and writes nothing, which is the property everything else here rests on. WP15
// says the same thing about the symbol index and for the same reason: a measurement a gate
// could read stops being a measurement, because then the cheapest way to move the figure is
// to change what is measured. Nothing in this file is reachable from the gate path, and the
// command over it has no writer.
//
// Everything is counted rather than judged. There is no threshold here and no verdict: the
// report says what the trail holds and the decision WP3 is waiting for is a person's.

// commonSections are the two sections common to all six templates, which WP3 names and
// keeps outside its budget: the budget is three to four *phase specific* required sections
// per template, so a required `decisions` in P2 is a common section promoted rather than a
// fourth of the phase's own.
var commonSections = map[string]bool{"open-questions": true, "decisions": true}

// recordedFiles are the files whose words are counted, which is the set #117 counted by
// hand: what a phase writes as prose. gate.yaml, context.lock.yaml and cost.yaml are the
// runner's own writing and their length says nothing about what a phase asked of anybody.
var recordedFiles = []string{"output.md", "digest.md", "learning.yaml"}

// IntentFigures is one intent's row. Nothing here is stored anywhere; every field is counted
// off the files the intent already has, which is why this is a reader and not a new artifact.
type IntentFigures struct {
	Key     string
	Created string // the date, as the listing prints it, or empty where it cannot be read
	State   string // as Summarise computes it: abandoned, complete, or where the work got to

	// Phases is how many of the six carry a verdict, and SixPhase says all of them do. An
	// intent that stopped at the intake is not the same measurement as one that went through,
	// so the population is stated rather than averaged over.
	Phases   int
	SixPhase bool

	// Sections is what was written, counted from the anchors in output.md rather than from
	// the templates: a sealed phase records the template it rendered from, and counting a
	// phase from today's template would report a number the phase was never asked for.
	Sections int
	Files    int // every file under the intent directory, which is what a reader of it finds
	Words    int // across recordedFiles, over every phase

	// Cost is what the ledger attributed, and CostPhases how many phases attributed anything.
	// The two travel together or not at all: section 11 calls the figure self-reported, and
	// over this trail it covers a handful of phases out of hundreds.
	Cost       cost.Totals
	CostPhases int

	// The three things the trail can honestly answer about a reopen. See Reopens.
	Released    int // findings a person had to approve or override for the phase to pass
	Obligations int // overrides whose obligation is still open
	Moved       int // phases whose files have changed since their verdict was written

	Problem string // why this row could not be read, where that happened
}

// Reopens is the reopen figure, and it is a proxy rather than the thing itself.
//
// A reopen in the plain sense is a phase judged twice, and the trail cannot answer that: a
// phase keeps one gate.yaml, section 11's rule that what is sealed is never rewritten means
// the verdict in it is the last one, and the runs before it left nothing behind. So what is
// counted is a phase that did not pass on its own merits — a finding released by a person's
// approval or override — which is the event a reopen would have been about. The two counts
// beside it are reported separately rather than added in: an open obligation is work the
// release deferred, and a phase whose files moved after its verdict is one a re-run would
// judge again.
func (f IntentFigures) Reopens() int { return f.Released }

// TemplateFigures is one template against WP3's budget. The budget is the lever WP3 names,
// and the reason it is in this report is that both halves of the argument have to be visible
// at once: every template is inside its budget and seventeen required sections still come out
// of the six, which is a fact about the total and not about any one template.
type TemplateFigures struct {
	Phase    string
	Ref      string
	Source   template.Source
	Required int // every required section
	Specific int // those of them that are not one of the two common sections
	Optional int
	Problem  string
}

// Figures is the whole report: the population it counted, the rows, and the totals over them.
type Figures struct {
	// Selected is every intent the report was asked about, in the order they were created.
	Selected []IntentFigures
	// Counted and SixPhase are the population statement. A count over the trail includes the
	// intent doing the counting, and #201 is this project's learning about a figure stated
	// without saying what was behind it, so the report says both numbers rather than one.
	Counted  int
	SixPhase int
	// Totals are summed over the six-phase intents alone, and PhasesCounted is six times
	// SixPhase: the denominator the cost coverage is a fraction of.
	Totals        IntentFigures
	PhasesCounted int
	Templates     []TemplateFigures
}

// Mean divides a total by the six-phase population, and reports whether there was one to
// divide by. Separated from the printing so that the arithmetic can be asserted without
// capturing output.
func (f *Figures) Mean(total int) (float64, bool) {
	if f.SixPhase == 0 {
		return 0, false
	}
	return float64(total) / float64(f.SixPhase), true
}

// CostCovered reports whether every phase in the population attributed something, which is
// the one condition under which a cost per intent is a figure rather than a guess. Over this
// repository it is false by three orders of magnitude, and the report prints the coverage in
// place of the mean when it is.
func (f *Figures) CostCovered() bool {
	return f.PhasesCounted > 0 && f.Totals.CostPhases == f.PhasesCounted
}

// Figures reads the trail. With no keys it reads every intent, which is the case: a selection
// is a narrowing of the whole and not the other way round, as `learning propose` treats one.
//
// The population is always the whole selection and the totals are always over the six-phase
// intents within it. Those are two different numbers and both are printed, because an intent
// that stopped at the intake has a real row — three files and no verdict — and averaging it
// together with a finished one would report a process nobody ran.
func (r *Runner) Figures(keys []string) (*Figures, error) {
	out := &Figures{}
	if len(keys) == 0 {
		all, err := r.Intents()
		if err != nil {
			return nil, err
		}
		for _, s := range all {
			out.Selected = append(out.Selected, r.intentFigures(s))
		}
	} else {
		for _, key := range keys {
			if !fm.Exists(r.abs(model.IntentDir(key))) {
				return nil, refuse("no intent %s", key)
			}
			out.Selected = append(out.Selected, r.intentFigures(r.Summarise(key)))
		}
	}

	out.Counted = len(out.Selected)
	for _, f := range out.Selected {
		if !f.SixPhase {
			continue
		}
		out.SixPhase++
		out.Totals.Phases += f.Phases
		out.Totals.Sections += f.Sections
		out.Totals.Files += f.Files
		out.Totals.Words += f.Words
		out.Totals.CostPhases += f.CostPhases
		out.Totals.Cost.In += f.Cost.In
		out.Totals.Cost.Out += f.Cost.Out
		out.Totals.Cost.Cached += f.Cost.Cached
		out.Totals.Released += f.Released
		out.Totals.Obligations += f.Obligations
		out.Totals.Moved += f.Moved
	}
	out.PhasesCounted = out.SixPhase * len(model.Phases)
	out.Templates = r.templateFigures()
	return out, nil
}

// intentFigures counts one intent. It reports rather than fails, as Summarise does: a figure
// missing from a report is worse than one that says why it is not there.
func (r *Runner) intentFigures(s IntentSummary) IntentFigures {
	f := IntentFigures{Key: s.Key, Created: s.Created, State: s.State, Problem: s.Problem}
	if len(f.Created) > 10 {
		f.Created = f.Created[:10]
	}
	f.Files = countFiles(r.abs(model.IntentDir(s.Key)))

	states, err := r.Status(s.Key)
	if err != nil {
		if f.Problem == "" {
			f.Problem = "its phases cannot be read"
		}
		return f
	}
	for _, st := range states {
		if st.Status == "" {
			continue
		}
		f.Phases++
		if st.State == "changed-after-verdict" {
			f.Moved++
		}
	}
	f.SixPhase = f.Phases == len(model.Phases)

	for _, p := range model.Phases {
		dir := r.abs(model.PhaseDir(s.Key, p))
		f.Sections += countSections(filepath.Join(dir, "output.md"))
		for _, name := range recordedFiles {
			f.Words += countWords(filepath.Join(dir, name))
		}
		var c model.Cost
		if err := fm.ReadYAML(filepath.Join(dir, "cost.yaml"), &c); err == nil {
			f.CostPhases++
			f.Cost.In += c.TokensIn
			f.Cost.Out += c.TokensOut
			f.Cost.Cached += c.TokensCached
		}
		released, open := countReleased(filepath.Join(dir, "gate.yaml"))
		f.Released += released
		f.Obligations += open
	}
	return f
}

// templateFigures resolves the six templates the way a phase does, through Runner.Template,
// so that a project which replaced one is counted on the one it actually renders.
func (r *Runner) templateFigures() []TemplateFigures {
	out := make([]TemplateFigures, 0, len(model.Phases))
	for _, p := range model.Phases {
		t, err := r.Template(p)
		if err != nil {
			out = append(out, TemplateFigures{Phase: p, Problem: err.Error()})
			continue
		}
		fig := TemplateFigures{Phase: p, Ref: t.Ref(), Source: t.Source}
		for _, s := range t.Template.Sections {
			switch {
			case !s.Required:
				fig.Optional++
			default:
				fig.Required++
				if !commonSections[s.ID] {
					fig.Specific++
				}
			}
		}
		out = append(out, fig)
	}
	return out
}

// countFiles is every file under an intent directory, which is the figure #117 counted: what
// a person who opens the directory finds. A directory that is not there counts zero rather
// than failing, because an intent whose files cannot be read is a row and not an error.
func countFiles(dir string) int {
	n := 0
	_ = filepath.WalkDir(dir, func(_ string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			n++
		}
		return nil
	})
	return n
}

// countSections counts the sections a phase actually wrote, from the anchors rather than
// from a template: the anchors are in the file whose hash the verdict carries, so this
// counts what was sealed. An empty section is not counted, since an absent section and a
// present empty one are the same silence and G-Schema already refuses the second.
func countSections(path string) int {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0
	}
	_, body, err := fm.Split(b)
	if err != nil {
		return 0
	}
	n := 0
	for _, content := range template.Parse(string(body)) {
		if strings.TrimSpace(content) != "" {
			n++
		}
	}
	return n
}

// countWords counts words the way `wc -w` does, which is the tool #117's figures were
// counted with and therefore the one this has to agree with to be a recomputation of them.
func countWords(path string) int {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0
	}
	return len(strings.Fields(string(b)))
}

// countReleased counts the findings a person released and the obligations still open among
// them. An approval and an override are counted together, because what the figure is about
// is a phase that could not pass on its own; which of the two it was is a question about the
// finding and the verdict carries it.
func countReleased(path string) (released, open int) {
	var g model.Gate
	if err := fm.ReadYAML(path, &g); err != nil {
		return 0, 0
	}
	for _, c := range g.Checks {
		for _, f := range c.Findings {
			if f.Decision == nil {
				continue
			}
			released++
			if f.Decision.Type == "overridden" && f.Decision.Obligation != "closed" {
				open++
			}
		}
	}
	return released, open
}
