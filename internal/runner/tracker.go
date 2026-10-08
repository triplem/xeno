// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/triplem/xeno/internal/enforcement"
	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/host"
	"github.com/triplem/xeno/internal/model"
)

// The tracker half of section 12: starting a phase reads the issue, and a job that has
// judged a phase writes the verdict back where people who do not work from a command line
// can read it. Both are on the side of the CLI that may use the network, and neither is
// reachable from a gate — `xeno gate ...` never touches the network, which is the property
// every verdict rests on, so the write-back is a step after the verdict and never part of
// one.

// IntakeProblem is the section of the intake the issue's content is written to. The intake
// has five sections and this is the one that holds the problem as stated; the issue is the
// one place it was stated by the person who raised it, so it belongs there and nowhere
// else.
//
// The id is the shipped template's. A project that overrides the intake template with a
// structure that has no such section gets nothing written and a sentence saying so, which
// is the same degradation an absent tracker block gets: what is lost is a convenience and
// never a part of the trail.
const IntakeProblem = "problem"

// Started is what `phase start` has to say about the tracker, beside having started the
// phase. Section is what was written, Note is why nothing was, and exactly one of them
// carries anything.
//
// It is a return value and not a refusal, because neither a project without a tracker nor
// a host that declines a read is a reason not to start a phase: Appendix A says an absent
// block means "a phase is started with the issue content supplied by hand", and that is
// the state this reports rather than fails on.
type Started struct {
	Section string
	Note    string
}

// trackerClient is the transport both halves use. The timeout is the one the branch rules
// port already uses, so that a host which has stopped answering costs one command the same
// wait wherever it is asked from.
func trackerClient() *http.Client { return &http.Client{Timeout: 20 * time.Second} }

// Start begins a phase, discarding what the tracker had to say about it. StartPhase below
// is the same act for a caller that reports it.
func (r *Runner) Start(key, phase string) error {
	_, err := r.StartPhase(key, phase)
	return err
}

// readIssue asks the host for the issue the intent belongs to. It is issueOf for a caller
// that holds the key of an intent that exists; `intent start` holds the id of one that does
// not yet, and calls issueOf itself.
func (r *Runner) readIssue(key string) (model.Issue, string, error) {
	var p struct {
		Tracker model.Tracker `yaml:"tracker"`
	}
	if err := fm.ReadYAML(r.abs(projectConfig), &p); err != nil && !os.IsNotExist(err) {
		return model.Issue{}, "", refuse("%s cannot be read: %v", projectConfig, err)
	}
	if p.Tracker == (model.Tracker{}) {
		return model.Issue{}, "no tracker is configured, so no issue was read; the block " +
			"is optional and the intake is then written from what somebody knows", nil
	}
	qualified, err := r.qualified(key)
	if err != nil {
		return model.Issue{}, "", err
	}
	return r.issueOf(p.Tracker, qualified)
}

// issueOf asks the host for the issue a qualified id names, through the tracker block
// given, which the caller has already found to be there.
//
// The three answers are an issue, a sentence saying why there is none, and an error. The
// middle one is the common case and is deliberately not an error here: a host that refuses
// the read, an issue that is not there, a machine with no token. What a caller does with
// it is the caller's — `phase start` prints it and starts the phase, as Appendix A says,
// and `intent start` refuses on it, as section 12 says, because a start that could not
// read the approval has nothing to stand on. The last is a credential that does not work
// and a host that answered something nobody can read, which the plan's credential table
// puts at exit code 2 — a token nobody renewed is a different failure from a tracker
// nobody configured, and the staircase is what separates them.
func (r *Runner) issueOf(t model.Tracker, qualified string) (model.Issue, string, error) {
	// An intent whose issue lies on a host this project is not configured for is an absence
	// and not a misconfiguration: `--for` takes a whole qualified id, so an intent may
	// legitimately belong to an issue the configured adapter cannot reach, and the one
	// honest answer about it is that nothing was read. It is resolved before the credential
	// is, so that such an intent costs no token and no call.
	project, issueKey, err := t.Locate(qualified)
	if err != nil {
		return model.Issue{}, fmt.Sprintf("no issue was read: %v", err), nil
	}
	// The whole block is there and some of it is missing: Appendix A's rule for that case
	// is that `xeno phase start` refuses rather than guessing, and the refusal names the
	// field. host.For makes it for the adapter and the address.
	h, err := host.For(t, trackerClient())
	if err != nil {
		return model.Issue{}, "", refuse("%v", err)
	}
	token, err := host.Credentials(t.Auth)
	// A machine with no token reads no issue and starts the phase all the same. The issue's
	// content is what the intake is saved from writing by hand, not what makes a phase
	// legitimate, and a clone without a token is the ordinary case for anybody reading the
	// trail or running the tests. A block that names no variable, or names a scheme neither
	// adapter implements, is the other case and still refuses: that is a configuration
	// somebody has to finish rather than a credential somebody has not got.
	if errors.Is(err, host.ErrNoToken) {
		return model.Issue{}, fmt.Sprintf("no issue was read: %v", err), nil
	}
	if err != nil {
		return model.Issue{}, "", err
	}
	issue, err := h.Issue(project, issueKey, token)
	if err != nil {
		return model.Issue{}, "", err
	}
	return issue, "", nil
}

// intake is the issue's content as the phase records it: by what the intent was
// authorised, then the qualified id, the title and the body, quoted rather than retold.
// Section 12 asks a phase start to read the issue, and what makes that worth more than the
// agent's summary of it is precisely that it is not a summary — the words the problem was
// raised in are the ones a reviewer later compares the requirements against.
//
// The sentence above the quote is the one section 12's "Starting an intent" has the intake
// record: who approved, when and why, and that the intent was started ahead of its
// milestone where that is so. It says what the issue carries at the moment P0 starts and
// not what it carried when the intent was started, because intent.yaml has no field for
// either and section 5 enumerates its fields; the two moments are usually minutes apart,
// and where they differ the intake says what is true when it is written. Where no
// approval is found, the sentence says so rather than saying nothing, since an intake
// silent on the point reads as one written before the clause.
//
// Every line of the issue is quoted with a Markdown blockquote marker, which is also what
// keeps a section anchor inside an issue body from becoming one. An anchor is recognised
// on a line of its own, so an issue whose text contained one would otherwise split this
// section in two and hand the rest of it a name somebody else chose; the quote marker puts
// a character in front of it and the line is text again.
func intake(qualified string, issue model.Issue, at string) string {
	var b strings.Builder
	if a, missing := issue.Approval(); len(missing) == 0 {
		if a.Reason == "" {
			fmt.Fprintf(&b, "Approved by @%s on %s, with no reason beside the word.\n", a.By, a.At)
		} else {
			fmt.Fprintf(&b, "Approved by @%s on %s: %s\n", a.By, a.At,
				strings.Join(strings.Fields(a.Reason), " "))
		}
	} else {
		fmt.Fprintf(&b, "No approval was found on the issue when this phase started: it is "+
			"missing %s.\n", strings.Join(missing, " and "))
	}
	if m := issue.Ahead(); m != nil {
		fmt.Fprintf(&b, "Started ahead of its milestone, %q, while %q is open.\n",
			issue.Milestone.Title, m.Title)
	}
	fmt.Fprintf(&b, "\n> **%s** — %s\n>\n", qualified, strings.TrimSpace(issue.Title))
	for _, line := range strings.Split(strings.ReplaceAll(issue.Body, "\r\n", "\n"), "\n") {
		if line = strings.TrimRight(line, " \t"); line == "" {
			b.WriteString(">\n")
			continue
		}
		fmt.Fprintf(&b, "> %s\n", line)
	}
	fmt.Fprintf(&b, "\nThe issue as it stood at %s, read by `xeno phase start` and quoted "+
		"rather than summarised. What this phase concludes about it belongs below.\n", at)
	return b.String()
}

// WriteBack is one write-back: the comment that was composed, and where it landed.
type WriteBack struct {
	Intent  string
	Phase   string
	Body    string
	Written model.Written
	// Posted is false for a dry run and for a host that declined the write. Written.Reason
	// says which.
	Posted bool
}

// ReportVerdict writes a phase's verdict to the issue the intent came from: the phase, the
// verdict, the findings with their decisions, and the enforcement report beside them.
//
// It is a command of its own and never part of `xeno gate`, for the reason
// EnforcementCheck gives: the gate path's freedom from the network is what makes a verdict
// reproducible, and a gate that called out would depend on whether a service answered. So
// this runs after the verdict, in the job that already has one, and reads `gate.yaml` from
// the repository rather than judging anything itself.
//
// One call and no read before it. Section 12 costs the write-back at "one call from a job
// that already runs", and finding a previous comment to edit instead would be two; a
// comment per run is the consequence, which is why the workflow that calls this runs it on
// a push to the default branch and not on every pull request event.
//
// An empty phase is the last one holding a verdict, as `review answer` resolves the phase
// it can only mean. A job reporting after a run has judged whatever it judged, and naming
// the phase in the pipeline would mean the pipeline knowing the trail.
func (r *Runner) ReportVerdict(key, phase string, dry bool) (*WriteBack, error) {
	if phase == "" {
		p, err := r.lastJudged(key)
		if err != nil {
			return nil, err
		}
		phase = p
	}
	if model.PhaseIndex(phase) < 0 {
		return nil, fmt.Errorf("unknown phase %q", phase)
	}
	g, err := r.readGate(key, phase)
	if err != nil {
		return nil, refuse("%s of %s has no verdict to report; run xeno phase finish first",
			phase, key)
	}
	var p struct {
		Tracker model.Tracker `yaml:"tracker"`
	}
	if err := fm.ReadYAML(r.abs(projectConfig), &p); err != nil && !os.IsNotExist(err) {
		return nil, refuse("%s cannot be read: %v", projectConfig, err)
	}
	w := &WriteBack{Intent: g.Intent, Phase: phase,
		Body: r.comment(key, phase, g, r.enforcementReport())}
	if dry {
		return w, nil
	}
	if p.Tracker == (model.Tracker{}) {
		w.Written.Reason = "no tracker is configured, so no comment was written; the block " +
			"is optional and Appendix A says a project without one writes none"
		return w, nil
	}
	project, issueKey, err := p.Tracker.Locate(g.Intent)
	if err != nil {
		// Same reading as on the read side: an issue on a host this project is not
		// configured for is somewhere nothing here can write, which the comment in the
		// pipeline log says rather than failing the step over.
		w.Written.Reason = err.Error()
		return w, nil
	}
	h, err := host.For(p.Tracker, trackerClient())
	if err != nil {
		return nil, refuse("%v", err)
	}
	token, err := host.Credentials(p.Tracker.Auth)
	if err != nil {
		return nil, err
	}
	if w.Written, err = h.Comment(project, issueKey, token, w.Body); err != nil {
		return nil, err
	}
	w.Posted = w.Written.Reason == ""
	return w, nil
}

// lastJudged is the furthest phase of the intent that holds a verdict. The furthest and
// not the newest by time: the phases are ordered and a later one is judged on an earlier
// one, so the last in the order is the one a run has just reached.
func (r *Runner) lastJudged(key string) (string, error) {
	for i := len(model.Phases) - 1; i >= 0; i-- {
		if fm.Exists(r.abs(model.PhaseDir(key, model.Phases[i]) + "/gate.yaml")) {
			return model.Phases[i], nil
		}
	}
	return "", refuse("no phase of %s holds a verdict, so there is nothing to report", key)
}

// enforcementReport is the report `xeno enforcement check` left in the local data
// location, or nil where that command has not run. Nil is ordinary: the check needs the
// network and a token of its own, and a job that could not ask the host still has a
// verdict worth reporting.
func (r *Runner) enforcementReport() *enforcement.Report {
	var rep enforcement.Report
	if err := fm.ReadYAML(ReportPath(r.Root), &rep); err != nil {
		return nil
	}
	return &rep
}

// statusSays is section 5's five derived statuses, each in the words that section defines
// it with. The status itself is in the comment as well, because it is what `gate.yaml` says
// and what somebody asking about it will quote; the sentence is there because the audience
// of this comment is everybody who does not work from a command line, and "overridden" is
// not self explanatory to a reader meeting it for the first time.
var statusSays = map[string]string{
	"green":       "every gate that applies to this phase passed.",
	"red":         "a gate failed and nobody has decided what to do about it yet.",
	"provisional": "a check is waiting for evidence a pipeline has yet to produce.",
	"approved":    "a person assessed every failing finding before the merge and accepted it. The findings stay failed; the approval is the record.",
	"overridden":  "the merge was taken first and the artifacts are still owed. That is what the obligation below is.",
}

// comment is the body of the write-back. It is Markdown, addressed to somebody reading an
// issue in a browser: the verdict first and in a sentence, then the gates as a table, then
// every finding with what to do about it and whatever decision a person has taken on it,
// then what the host is configured to enforce.
//
// It states where it came from and that it decides nothing. A comment that looks like the
// record would become a second place a verdict lives, which is what section 5 forbids of
// the schema and is no better in a comment; the trail in the repository is the record, and
// this is a copy of it for people who cannot read that.
func (r *Runner) comment(key, phase string, g *model.Gate, rep *enforcement.Report) string {
	var b strings.Builder
	fmt.Fprintf(&b, "## Xeno: %s is %s\n\n", phase, g.Status)
	if says := statusSays[g.Status]; says != "" {
		fmt.Fprintf(&b, "%s\n\n", says)
	}
	fmt.Fprintf(&b, "Intent `%s`, judged at %s.\n\n", key, g.RunAt)

	b.WriteString("| Gate | Result |\n| --- | --- |\n")
	for _, c := range g.Checks {
		fmt.Fprintf(&b, "| %s | %s |\n", c.Gate, c.Result)
	}
	b.WriteString("\n")
	b.WriteString(findingsOf(g))
	b.WriteString(enforcementOf(rep))
	fmt.Fprintf(&b, "\nWritten by `xeno report verdict` from `%s/gate.yaml`%s. It decides "+
		"nothing and records nothing: the trail in the repository is the record, and a "+
		"failing finding is decided with `xeno gate approve` or `xeno gate override`.\n",
		model.PhaseDir(key, phase), asOf(r.headCommit()))
	return b.String()
}

// asOf names the tree the verdict was read from, where there is one. A comment is read
// weeks later and a verdict is about content rather than about a commit, so the commit is
// the one thing that tells a reader which state of the repository this was.
func asOf(commit string) string {
	if commit == "" {
		return ""
	}
	return ", as the tree stood at " + commit
}

// findingsOf is every finding of every check, with its remedy and its decision. Advisory
// ones are named as such, because section 5 makes them reported rather than held against
// the phase and a reader counting failures would otherwise count one that did not fail.
func findingsOf(g *model.Gate) string {
	var b strings.Builder
	var total, undecided int
	for _, c := range g.Checks {
		for _, f := range c.Findings {
			total++
			if f.Decision == nil && !f.Advisory {
				undecided++
			}
		}
	}
	if total == 0 {
		return "No gate reported a finding.\n\n"
	}
	fmt.Fprintf(&b, "### Findings\n\n%s, %s.\n\n", count(total, "finding", "findings"),
		decidedness(undecided))
	for _, c := range g.Checks {
		for _, f := range c.Findings {
			fmt.Fprintf(&b, "- **%s**", c.Gate)
			if f.File != "" {
				fmt.Fprintf(&b, " on `%s`", f.File)
			}
			if f.Advisory {
				b.WriteString(" (advisory, so it does not fail the check)")
			}
			fmt.Fprintf(&b, " — %s\n", oneLine(f.Cause))
			if f.Next != "" {
				fmt.Fprintf(&b, "  - what to do: %s\n", oneLine(f.Next))
			}
			b.WriteString(decisionOf(f))
			fmt.Fprintf(&b, "  - finding `%s`\n", f.ID)
		}
	}
	b.WriteString("\n")
	return b.String()
}

// decisionOf is the one line of a finding a person wrote, in the words section 5 gives the
// two kinds. An obligation that is still open is said in the sentence rather than left to
// the field name, because what it means — the merge was taken and the artifacts are owed —
// is the thing a reader of an issue has to know.
func decisionOf(f model.Finding) string {
	d := f.Decision
	if d == nil {
		if f.Advisory {
			return ""
		}
		return "  - nobody has decided this yet\n"
	}
	line := fmt.Sprintf("  - %s by %s on %s: %s", d.Type, d.By, d.At, oneLine(d.Reason))
	if d.Obligation == "open" {
		line += " — the artifacts are still owed"
	}
	return line + "\n"
}

// enforcementOf is what the host is configured to enforce, from the report the same job
// produced. Absent is said rather than left out: a comment that silently omitted the
// section would read as a host with nothing to report, and the whole reason the check
// exists is that an unchecked host is indistinguishable from a compliant one.
func enforcementOf(rep *enforcement.Report) string {
	if rep == nil {
		return "No enforcement report lay beside this verdict, so what the host is " +
			"configured to enforce goes unreported here. `xeno enforcement check` " +
			"produces it and needs the network and a token of its own.\n"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "### What the host enforces\n\n`%s`, branch `%s`, asked at %s.\n\n",
		rep.Repository, rep.Branch, rep.CheckedAt)
	b.WriteString("| Requirement | Declared | The host | |\n| --- | --- | --- | --- |\n")
	for _, q := range rep.Requirements {
		note := ""
		if q.Note != "" {
			note = " — " + oneLine(q.Note)
		}
		fmt.Fprintf(&b, "| `%s` | %s | %s | %s%s |\n",
			q.Name, q.Declared, oneLine(q.Actual), q.State, note)
	}
	if n := rep.Unmet(); n > 0 {
		fmt.Fprintf(&b, "\n%s neither met nor waived.\n",
			count(n, "requirement is", "requirements are"))
	}
	return b.String()
}

// oneLine keeps a cause, a remedy or a reason from breaking the table or the list it sits
// in. The text is written for a terminal and may carry newlines and pipes; a reader of the
// issue gets the same words on one line rather than a table that has come apart.
func oneLine(s string) string {
	return strings.Join(strings.Fields(strings.ReplaceAll(s, "|", "\\|")), " ")
}

func count(n int, one, many string) string {
	if n == 1 {
		return "One " + one
	}
	return fmt.Sprintf("%d %s", n, many)
}

func decidedness(undecided int) string {
	switch undecided {
	case 0:
		return "each of them either decided or advisory"
	case 1:
		return "one of them still undecided"
	default:
		return fmt.Sprintf("%d of them still undecided", undecided)
	}
}
