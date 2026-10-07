// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/template"
)

// Suggestion is the next step of the working sequence, read off the state the runner
// already computes. Section 6 fixes who does what with which command and in which
// order; this reads that table against one intent and adds nothing to it. No field, no
// gate and no rule.
//
// It is never an action. Nothing happens because it was suggested, and where the
// sequence says a person decides, the suggestion names the decision and leaves it
// unmade. Where the next step is not the runner's, Command is empty: offering a command
// that does not exist is worse than offering none, because it is tried.
type Suggestion struct {
	Phase   string
	Text    string
	Command string   // empty where the step is nobody's subcommand
	Owed    []string // obligations from an override, which are due whenever somebody gets to them
}

// Next answers for one intent. It returns a suggestion in every case, including the
// case where it does not know: a state the working sequence does not cover is said to
// be one, because a confident wrong suggestion is followed.
func (r *Runner) Next(key string) *Suggestion {
	if !fm.Exists(r.abs(model.IntentDir(key))) {
		// The command is named in the sentence rather than in Command, because the issue is
		// the one thing nothing here knows, and a command offered with a placeholder in it
		// would be typed as it stands and create an intent for an issue called ISSUE.
		return &Suggestion{Text: "there is no intent " + key + " yet. Create it with " +
			"xeno intent start --intent " + key + ", naming the issue it belongs to with " +
			"--for. Its assumptions.yaml is written by hand; no command creates it"}
	}
	states, err := r.Status(key)
	if err != nil {
		return &Suggestion{Text: "the state of " + key + " cannot be read, so there is nothing " +
			"to suggest: " + err.Error()}
	}
	owed := r.owed(key, states)
	for i, s := range states {
		if sug := r.next(key, i, s); sug != nil {
			sug.Owed = owed
			if s.Stale {
				sug.Text += ". Its predecessor changed after this phase started, so read it again first"
			}
			return sug
		}
	}
	return &Suggestion{Text: "every phase is finished and decided", Owed: owed}
}

// next answers for one phase, or nil where that phase is done and the walk goes on.
func (r *Runner) next(key string, i int, s PhaseState) *Suggestion {
	at := " --intent " + key + " --phase " + phaseNumber(s.Phase)
	switch s.State {
	case "not-started":
		// The fresh session is the one part of the context economy nothing checks. Section 6 has
		// a phase work from its predecessor's digest and section 5 has context.lock.yaml state
		// what it was given; a session carrying four earlier phases gives this one a context
		// neither describes. No harness is named: section 7 records XENO_HARNESS and never
		// branches on it, and a person working with commands alone has no such command at all.
		return &Suggestion{Phase: s.Phase, Text: "start " + s.Phase + " in a fresh session, so " +
			"that its context is what context.lock.yaml says it was given",
			Command: "xeno phase start" + at}
	case "running":
		return r.running(key, s, at)
	case "changed-after-verdict":
		return &Suggestion{Phase: s.Phase, Text: s.Phase + " changed after its verdict was " +
			"written, so the verdict is about something else. Judge it again",
			Command: "xeno phase finish" + at}
	case "finished":
		switch s.Status {
		case "red":
			return r.red(key, s)
		case "provisional":
			return &Suggestion{Phase: s.Phase, Text: s.Phase + " is provisional: the pipeline " +
				"owes evidence it declared. Nothing to do here, and the next phase's start " +
				"attaches whatever has arrived"}
		case "green", "approved", "overridden":
			if i == len(model.Phases)-1 {
				// And the session that is ending has nothing left to carry: the next intent
				// starts from the tree and from the scope its own intake declares, so none of
				// this one's context is read again. A hint rather than a step, which is what
				// this suggestion already is for the merge.
				return &Suggestion{Phase: s.Phase, Text: "P5 is decided, so the merge is next. " +
					"That is not a xeno command: commit, push, and let the review and the " +
					"pipeline run. Nothing of this intent's context is read again, so the next " +
					"one begins best in a fresh session"}
			}
			return nil
		default:
			return &Suggestion{Phase: s.Phase, Text: "the verdict of " + s.Phase + " reads " +
				strconv.Quote(s.Status) + ", which the working sequence does not cover. " +
				"No suggestion for it"}
		}
	}
	return &Suggestion{Phase: s.Phase, Text: "the state of " + s.Phase + " reads " +
		strconv.Quote(s.State) + ", which the working sequence does not cover. No suggestion for it"}
}

// running reads the template's required sections against what the artifact holds. The
// suggestion names the section rather than the idea of a section, because the command
// needs one and looking it up is the part somebody would otherwise do by hand.
func (r *Runner) running(key string, s PhaseState, at string) *Suggestion {
	t, err := template.Load(r.Root, model.TemplateID(s.Phase), r.language())
	if err != nil {
		return &Suggestion{Phase: s.Phase, Text: s.Phase + " is running, and its template does " +
			"not resolve, so which sections it wants cannot be said: " + err.Error()}
	}
	content := map[string]string{}
	if b, err := os.ReadFile(r.abs(model.PhaseDir(key, s.Phase) + "/output.md")); err == nil {
		if _, body, ferr := fm.Split(b); ferr == nil {
			content = template.Parse(string(body))
		}
	}
	missing := t.Missing(content)
	if len(missing) == 0 {
		return &Suggestion{Phase: s.Phase, Text: "the required sections of " + s.Phase +
			" are written. Judge it", Command: "xeno phase finish" + at}
	}
	return &Suggestion{Phase: s.Phase,
		Text:    s.Phase + " is running and still wants " + strings.Join(missing, ", "),
		Command: "xeno section set " + missing[0] + at}
}

// red names the finding and both ways out, and offers neither as a command. Releasing a
// finding is a statement by a second person, so a runner that put `gate approve` in the
// reader's hand would be nudging towards the decision it exists to record.
func (r *Runner) red(key string, s PhaseState) *Suggestion {
	g, err := r.readGate(key, s.Phase)
	if err != nil {
		return &Suggestion{Phase: s.Phase, Text: s.Phase + " is red and its verdict cannot be " +
			"read: " + err.Error()}
	}
	for _, c := range g.Checks {
		for _, f := range c.Findings {
			if f.Decision != nil {
				continue
			}
			return &Suggestion{Phase: s.Phase, Text: fmt.Sprintf(
				"%s is red on %s, %s: %s. Then judge it again. Or a second person releases "+
					"it, with gate approve or gate override, naming %s and a reason",
				s.Phase, f.ID, f.File, f.Next, f.ID)}
		}
	}
	return &Suggestion{Phase: s.Phase, Text: s.Phase + " is red with every finding decided, " +
		"which the working sequence does not cover. No suggestion for it"}
}

// owed collects the obligations an override left open. They are not the next step: an
// override exists so that the work can go on, and section 6 puts them at "later,
// whoever owes it". They are listed so that later is not never.
func (r *Runner) owed(key string, states []PhaseState) []string {
	var out []string
	for _, s := range states {
		if s.State != "finished" && s.State != "changed-after-verdict" {
			continue
		}
		g, err := r.readGate(key, s.Phase)
		if err != nil {
			continue
		}
		for _, c := range g.Checks {
			for _, f := range c.Findings {
				if f.Decision != nil && f.Decision.Type == "overridden" && f.Decision.Obligation != "closed" {
					out = append(out, "xeno obligation close "+f.ID+" --intent "+key+
						" --phase "+phaseNumber(s.Phase))
				}
			}
		}
	}
	return out
}

// phaseNumber is the ordering prefix of a phase id, which is what --phase takes in its
// short form. The prefix orders the phases and says nothing else, which is why A33 keeps
// the template id on the other half of the name.
func phaseNumber(phase string) string {
	return strings.SplitN(phase, "-", 2)[0]
}
