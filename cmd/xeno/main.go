// SPDX-License-Identifier: Apache-2.0

// Command xeno is the runner. Exit codes follow one staircase throughout: 0 where the
// command did what was asked and the verdict is not red, 1 on a red verdict or a
// refusal with a reason, 2 where it could not run at all.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/triplem/xeno/internal/cost"
	"github.com/triplem/xeno/internal/enforcement"
	"github.com/triplem/xeno/internal/evidence"
	"github.com/triplem/xeno/internal/gates"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/runner"
)

const usage = `usage:
  xeno init           [--vendor] [--project OWNER/REPO] [--model ID] [--language TAG]
  xeno phase start    --intent KEY --phase NN [--evidence-from DIR] [--export]
  xeno phase finish   --intent KEY --phase NN [--summary PATH|-]   writes digest.md
  xeno gate run       --intent KEY --phase NN [--base REF --head REF] [--evidence-from DIR]
  xeno gate approve   FINDING --intent KEY --phase NN --by WHO --reason TEXT
  xeno gate override  FINDING --intent KEY --phase NN --by WHO --reason TEXT
  xeno obligation close FINDING --intent KEY --phase NN
  xeno assumption record --intent KEY --phase NN --text TEXT --origin WHERE --confidence HOW [--resolves KEY]
  xeno assumption confirm ID --intent KEY --by WHO
  xeno assumption reject  ID --intent KEY --by WHO
  xeno gate verify    [--intent KEY]            recompute and compare, write nothing (CI)
  xeno enforcement check [--branch NAME]        ask the host what it enforces (needs the network)
  xeno evidence attach --intent KEY --phase NN --from DIR
  xeno intent status  [--intent KEY]            without one, every intent by creation
  xeno intent close   --intent KEY --reason TEXT
  xeno section set    SECTION --intent KEY --phase NN [--file PATH]   reads stdin without --file
  xeno check commit-message [--pattern NAME] [--file PATH]   reads stdin without --file
  xeno cost turn                                reads a hook's JSON on stdin
  xeno version
common: --root DIR (default .), --no-next to leave out the next step`

func main() { os.Exit(run(os.Args[1:])) }

// opts is everything a command was given: the flags, the finding id that stands before
// them, the resolved phase and the runner they act on. One struct rather than twenty
// arguments, so that a command reads what it needs and a new flag reaches every command
// without touching any of them.
type opts struct {
	r       *runner.Runner
	cmd     string
	finding string
	phase   string

	root, key, phaseArg            string
	from, src, by, pattern, file   string
	summary                        string
	project, mdl, language         string
	pluginFrom, host, branch       string
	base, head, reason             string
	text, origin, confidence       string
	resolves                       string
	vendor, noNext, export, isJSON bool
}

// command is one entry of the surface. needsKey and needsPhase were two switches that
// listed their exceptions, which is a list a new command can be forgotten from twice;
// here they are a field, and the only place they are read is the dispatch below.
type command struct {
	needsKey   bool
	needsPhase bool
	run        func(*opts) int
}

var commands = map[string]command{
	"init":                 {run: cmdInit},
	"enforcement check":    {run: cmdEnforcementCheck},
	"check commit-message": {run: cmdCheckMessage},
	"gate verify":          {run: cmdGateVerify},
	"cost turn":            {run: cmdCostTurn},
	"intent status":        {run: cmdIntentStatus},
	"intent close":         {needsKey: true, run: cmdIntentClose},
	"assumption confirm":   {needsKey: true, run: cmdAssumptionDecide},
	"assumption reject":    {needsKey: true, run: cmdAssumptionDecide},
	"section set":          {needsKey: true, needsPhase: true, run: cmdSectionSet},
	"phase start":          {needsKey: true, needsPhase: true, run: cmdPhaseStart},
	"phase finish":         {needsKey: true, needsPhase: true, run: cmdPhaseFinish},
	"gate run":             {needsKey: true, needsPhase: true, run: cmdGateRun},
	"gate approve":         {needsKey: true, needsPhase: true, run: cmdGateApprove},
	"gate override":        {needsKey: true, needsPhase: true, run: cmdGateOverride},
	"assumption record":    {needsKey: true, needsPhase: true, run: cmdAssumptionRecord},
	"obligation close":     {needsKey: true, needsPhase: true, run: cmdObligationClose},
	"evidence attach":      {needsKey: true, needsPhase: true, run: cmdEvidenceAttach},
}

func run(args []string) int {
	if len(args) >= 1 && args[0] == "version" {
		fmt.Println("xeno", model.RunnerVersion)
		return 0
	}
	if len(args) < 2 && (len(args) == 0 || args[0] != "init") {
		fmt.Fprintln(os.Stderr, usage)
		return 2
	}
	// Commands are two words except init, which is one. The split is where the flags
	// begin, not a property of the name.
	name, rest := args[0]+" "+args[1], args[2:]
	if args[0] == "init" {
		name, rest = "init", args[1:]
	}
	c, ok := commands[name]
	if !ok {
		fmt.Fprintln(os.Stderr, usage)
		return 2
	}

	o, code := parse(name, rest)
	if code != 0 {
		return code
	}
	if c.needsKey && o.key == "" {
		fmt.Fprintln(os.Stderr, "--intent is required")
		return 2
	}
	if c.needsPhase {
		phase, err := model.ResolvePhase(o.phaseArg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		o.phase = phase
	}
	return c.run(o)
}

// parse takes the finding id off the front and reads the flags. The id stands before them
// as the process definition writes these commands, and Go's flag package stops at the
// first argument that is not a flag, so it cannot be read back out afterwards.
func parse(name string, args []string) (*opts, int) {
	o := &opts{cmd: name}
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		o.finding, args = args[0], args[1:]
	}

	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.StringVar(&o.root, "root", ".", "repository root")
	fs.StringVar(&o.key, "intent", "", "intent key")
	fs.StringVar(&o.phaseArg, "phase", "", "phase id or number")
	fs.StringVar(&o.from, "evidence-from", "", "directory standing in for the pipeline artifact store")
	fs.StringVar(&o.src, "from", "", "directory standing in for the pipeline artifact store")
	fs.StringVar(&o.by, "by", "", "the person deciding")
	fs.StringVar(&o.pattern, "pattern", "conventional-commits", "a shipped pattern name")
	fs.StringVar(&o.file, "file", "", "the message to read, or stdin when absent")
	fs.StringVar(&o.summary, "summary", "", "the phase summary the digest is written from, - for stdin")
	fs.BoolVar(&o.vendor, "vendor", false, "copy the plugin in and pin it")
	fs.StringVar(&o.project, "project", "", "the tracker project the intents belong to")
	fs.StringVar(&o.mdl, "model", "", "the default model a phase uses")
	fs.StringVar(&o.language, "language", "en", "the language artifacts are written in")
	fs.StringVar(&o.pluginFrom, "plugin-from", ".xeno/plugin", "where to vendor the plugin from")
	fs.StringVar(&o.host, "host", "github", "which CI wrapper to generate")
	fs.StringVar(&o.branch, "branch", "", "the branch whose protection to read, main by default")
	// Both ends of the commit range. They are an input of the run and are deliberately
	// not recorded: after a squash a recorded range would point at commits that no
	// longer exist, and a field that is sometimes wrong is worse than no field. No gate
	// reads them until WP4 brings the commit predicates.
	fs.StringVar(&o.base, "base", "", "the base of the commit range under review")
	fs.StringVar(&o.head, "head", "", "the head of the commit range under review")
	fs.StringVar(&o.reason, "reason", "", "why")
	fs.StringVar(&o.text, "text", "", "the statement being assumed")
	fs.StringVar(&o.origin, "origin", "", "where the assumption came from")
	fs.StringVar(&o.confidence, "confidence", "", "how much weight it carries")
	fs.StringVar(&o.resolves, "resolves", "", "the open question this assumption answers")
	// The suggestion is off by default nowhere and on by default nowhere either: the
	// commands that change state say it, the ones a pipeline or a hook runs do not, and
	// this turns it off for the scripts that are neither.
	fs.BoolVar(&o.noNext, "no-next", false, "do not say what the next step is")
	// Print what a shell can eval instead of saying what the next step is. Both on one
	// stream would make the eval swallow a sentence meant for a person.
	fs.BoolVar(&o.export, "export", false, "print the phase's environment for a shell to eval")
	if err := fs.Parse(args); err != nil {
		return nil, 2
	}

	o.r = runner.New(o.root)
	o.r.EvidenceFrom = o.from
	return o, 0
}

// next is the suggestion every command that changes state prints, and the reason each of
// these functions ends with it rather than returning a code directly.
func (o *opts) next(code int) int { return suggest(o.r, o.key, o.noNext, code) }

func cmdInit(o *opts) int {
	o.r.PluginSource = o.pluginFrom
	res, err := o.r.Init(runner.InitOptions{
		TrackerKey: o.project, Model: o.mdl, Language: o.language, Vendor: o.vendor, Host: o.host,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	printInit(res)
	return 0
}

func cmdEnforcementCheck(o *opts) int {
	rep, err := o.r.EnforcementCheck(o.branch, enforcement.Token())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	printEnforcement(rep)
	if rep.Unmet() > 0 {
		return 1
	}
	return 0
}

func cmdCheckMessage(o *opts) int {
	message, err := readMessage(o.file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	if err := gates.CheckMessage(o.pattern, message); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func cmdSectionSet(o *opts) int {
	content, err := readMessage(o.file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	t, err := o.r.SectionSet(o.key, o.phase, o.finding, content)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("%s rendered from %s (%s)\n", o.phase, t.Ref(), t.Source)
	return o.next(0)
}

func cmdIntentClose(o *opts) int { return o.next(report(o.r.IntentClose(o.key, o.reason))) }

func cmdPhaseStart(o *opts) int {
	if code := report(nil, o.r.Start(o.key, o.phase)); code != 0 {
		return code
	}
	if o.export {
		env, err := o.r.PhaseEnv(o.key, o.phase)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		fmt.Print(env)
		return 0
	}
	return o.next(0)
}

// cmdPhaseFinish passes the summary the digest is written from. A path, or "-" for stdin:
// absence cannot mean stdin here as it does for section set, because a summary is optional and
// a phase finished without one is judged exactly as before (section 5, #120).
func cmdPhaseFinish(o *opts) int {
	summary := ""
	if o.summary != "" {
		var err error
		if summary, err = readSummary(o.summary); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		if strings.TrimSpace(summary) == "" {
			fmt.Fprintln(os.Stderr, "--summary is empty: a digest with no summary is worse than none, since G-Schema would pass it")
			return 2
		}
	}
	return o.next(report(o.r.Finish(o.key, o.phase, summary)))
}

// readSummary reads a path, or stdin for "-". readMessage cannot serve: it reads stdin for an
// empty path, which is what "no summary" has to mean here.
func readSummary(path string) (string, error) {
	if path == "-" {
		b, err := io.ReadAll(os.Stdin)
		return string(b), err
	}
	b, err := os.ReadFile(path)
	return string(b), err
}

func cmdGateRun(o *opts) int {
	o.r.Base, o.r.Head = o.base, o.head
	return o.next(report(o.r.GateRun(o.key, o.phase)))
}

func cmdGateApprove(o *opts) int {
	return o.next(report(o.r.Decide(o.key, o.phase, o.finding, "approved", o.by, o.reason)))
}

func cmdGateOverride(o *opts) int {
	return o.next(report(o.r.Decide(o.key, o.phase, o.finding, "overridden", o.by, o.reason)))
}

func cmdAssumptionRecord(o *opts) int {
	a, err := o.r.RecordAssumption(o.key, o.phase, o.text, o.origin, o.confidence, o.resolves)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("%s recorded, open, from %s with %s confidence\n", a.ID, a.Origin, a.Confidence)
	return o.next(0)
}

func cmdAssumptionDecide(o *opts) int {
	status := "confirmed"
	if o.cmd == "assumption reject" {
		status = "rejected"
	}
	a, err := o.r.DecideAssumption(o.key, o.finding, status, o.by)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("%s %s by %s\n", a.ID, a.Status, o.by)
	return o.next(0)
}

func cmdObligationClose(o *opts) int {
	return o.next(report(o.r.CloseObligation(o.key, o.phase, o.finding)))
}

func cmdEvidenceAttach(o *opts) int {
	res, err := evidence.Attach(o.root, o.key, o.phase, o.src)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	fmt.Printf("attached %d, pending %d; run xeno gate run to carry the verdict forward\n",
		res.Attached, res.Pending)
	// Named on stderr and counted as pending: an entry nothing can bind was not attached,
	// and a run that only printed the counts would report it as evidence still to come
	// from a job that has already produced it.
	for _, u := range res.Unbindable {
		fmt.Fprintln(os.Stderr, "  not attached:", u)
	}
	if len(res.Unbindable) > 0 {
		fmt.Fprintln(os.Stderr, "  republish with a sha256; a uri is bound by its hash alone.")
		return o.next(1)
	}
	return o.next(0)
}

func cmdGateVerify(o *opts) int {
	res, err := o.r.Verify(o.key)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return 2
	}
	fmt.Printf("verified %d verdicts\n", res.Checked)
	for _, d := range res.Divergences {
		fmt.Printf("  DIVERGENT   %s %s: %s\n", d.Key, d.Phase, d.What)
	}
	for _, l := range res.Red {
		fmt.Printf("  RED         %s\n", l)
	}
	for _, l := range res.Provisional {
		fmt.Printf("  PROVISIONAL %s: evidence outstanding, binding only at the merge request\n", l)
	}
	if len(res.Divergences) > 0 || len(res.Red) > 0 {
		return 1
	}
	return 0
}

// cmdIntentStatus answers one of two questions. Without a key it lists every intent in the
// order it was created, which the key cannot express because it carried the issue number and
// issues are filed in a different order than work happens (#118). With one it reports the
// phases of that intent, unchanged.
//
// needsKey is false for this command alone in its family, so the key is checked here. gate
// verify takes --intent optionally in the same way.
func cmdIntentStatus(o *opts) int {
	if o.key == "" {
		return cmdIntentList(o)
	}
	states, err := o.r.Status(o.key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	for _, s := range states {
		line := fmt.Sprintf("%-18s %-22s %s", s.Phase, s.State, s.Status)
		if s.Stale {
			line += "  (predecessor changed since this phase started)"
		}
		fmt.Println(line)
	}
	return o.next(0)
}

// cmdIntentList prints the intents oldest first. A row whose intent.yaml could not be read
// says so in place of its date rather than going missing.
func cmdIntentList(o *opts) int {
	intents, err := o.r.Intents()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	for _, s := range intents {
		verdict := s.Verdict
		if s.Phase != "" {
			verdict = s.Phase + "  " + s.Verdict
		}
		// The date, not the timestamp: the order is what matters and a column of identical
		// times reads worse than a column of dates. The runner sorted on the whole value.
		date := s.Created
		if len(date) > 10 {
			date = date[:10]
		}
		if date == "" {
			date = "?"
		}
		line := fmt.Sprintf("%-10s  %-12s %-17s %s", date, s.Key, s.State, verdict)
		if s.Problem != "" {
			line += "  (" + s.Problem + ")"
		}
		fmt.Println(strings.TrimRight(line, " "))
	}
	return 0
}

// cmdCostTurn records what a turn cost, for the hook the plugin ships. Section 11 puts the
// figure in cost.yaml and says the local session logs are evaluated; a hook is the only place a
// token count is visible at all, because hook input carries transcript_path and no counts.
//
// It exits zero on every failure, without exception. This runs on every turn of every session in
// this repository, so a bug in it is not a missing figure but an unusable harness, and section 11
// says a phase without a cost record is complete. Nothing it cannot do is worth a turn.
//
// It reads counts and identifiers. A transcript is a whole conversation and none of it is copied.
func cmdCostTurn(o *opts) int {
	var hook struct {
		Transcript string `json:"transcript_path"`
		Session    string `json:"session_id"`
	}
	b, err := io.ReadAll(os.Stdin)
	if err != nil || json.Unmarshal(b, &hook) != nil || hook.Transcript == "" {
		return 0
	}
	totals, err := cost.TranscriptTotals(hook.Transcript)
	if err != nil || totals.Zero() {
		return 0
	}
	intent, phase := cost.LivePhase(o.root, ".xeno/local/phase.env")
	if phase == "" {
		phase = cost.NoPhase
	}
	_ = cost.Append(o.root, cost.Turn{
		At: time.Now().UTC().Format(time.RFC3339), Session: hook.Session,
		Intent: intent, Phase: phase, Totals: totals,
	})
	return 0
}

func suggest(r *runner.Runner, key string, off bool, code int) int {
	if off || code == 2 {
		return code
	}
	s := r.Next(key)
	fmt.Printf("\nnext: %s\n", s.Text)
	if s.Command != "" {
		fmt.Printf("  %s\n", s.Command)
	}
	for _, o := range s.Owed {
		fmt.Printf("  owed, whenever somebody gets to it:\n  %s\n", o)
	}
	return code
}

func report(g *model.Gate, err error) int {
	var ref *runner.Refusal
	switch {
	case errors.As(err, &ref):
		fmt.Fprintln(os.Stderr, "refused:", ref.Reason)
		return 1
	case err != nil:
		fmt.Fprintln(os.Stderr, "error:", err)
		return 2
	case g == nil:
		return 0
	}
	fmt.Printf("%s %s: %s\n", g.Intent, g.Phase, g.Status)
	for _, c := range g.Checks {
		fmt.Printf("  %-14s %s\n", c.Gate, c.Result)
		for _, f := range c.Findings {
			d := ""
			if f.Decision != nil {
				d = " [" + f.Decision.Type + "]"
			}
			fmt.Printf("      %s %s: %s%s\n        next: %s\n", f.ID, f.File, f.Cause, d, f.Next)
		}
	}
	if g.Status == "red" {
		return 1
	}
	return 0
}

// printInit says what was done, what was left alone, and what a person still has to do.
// The last list is the point: a first contact that leaves the project believing the gate
// is binding when it is not is worse than no first contact.
func printInit(res *runner.InitResult) {
	for _, p := range res.Created {
		fmt.Println("  created  ", p)
	}
	for _, p := range res.Kept {
		fmt.Println("  kept     ", p)
	}
	if len(res.Outstand) > 0 {
		fmt.Println("\nNot determined:")
		for _, s := range res.Outstand {
			fmt.Println("  -", s)
		}
	}
	fmt.Println("\nXeno cannot make these settings. Somebody with repository administration has to:")
	for _, s := range res.Manual {
		fmt.Println("  -", s)
	}
}

// printEnforcement prints the report. not-available is its own line rather than folded
// into unmet, because a setting the host does not have is nobody's oversight and
// reporting it as one sends somebody looking for a checkbox that is not there.
func printEnforcement(rep *enforcement.Report) {
	fmt.Printf("%s, branch %s\n", rep.Repository, rep.Branch)
	for _, q := range rep.Requirements {
		fmt.Printf("  %-13s %-20s declared %-6s actual %s\n", q.State, q.Name, q.Declared, q.Actual)
		if q.Note != "" {
			fmt.Printf("                  %s\n", q.Note)
		}
	}
	fmt.Printf("\nreport written to %s\n", runner.ReportPath)
	if rep.Unmet() > 0 {
		fmt.Printf("\n%d requirement(s) are neither met nor waived. Where the host cannot express one,\n"+
			"record it as waived in project.yaml with a reason and a date: an unmeetable requirement\n"+
			"becomes a decision in the repository rather than a complaint on every run.\n", rep.Unmet())
	}
}

// readMessage takes the message from a file, which is what a commit-msg hook has, or
// from stdin, which is what a pipe has.
func readMessage(path string) (string, error) {
	if path != "" {
		b, err := os.ReadFile(path)
		return string(b), err
	}
	b, err := io.ReadAll(os.Stdin)
	return string(b), err
}
