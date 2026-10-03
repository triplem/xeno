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
  xeno intent start   --for ISSUE [--intent KEY]        writes intent.yaml
  xeno phase start    --intent KEY --phase NN [--evidence-from DIR] [--export]
  xeno phase finish   --intent KEY --phase NN [--summary PATH|-]   writes digest.md
  xeno gate run       --intent KEY --phase NN [--base REF --head REF] [--evidence-from DIR]
  xeno gate approve   FINDING --intent KEY --phase NN --by WHO --reason TEXT
  xeno gate override  FINDING --intent KEY --phase NN --by WHO --reason TEXT
  xeno obligation close FINDING --intent KEY --phase NN
  xeno learning record --intent KEY [--phase NN] --category C --observation T --proposal T --target P
  xeno learning record --intent KEY [--phase NN] --no-finding
  xeno question record --intent KEY --phase NN [--file PATH]   reads the entry on stdin
  xeno decision record --intent KEY --phase NN --chosen TEXT --reason TEXT --by WHO [--resolves KEY] [--proposed-by WHO]
  xeno decision record --intent KEY --phase NN --withdraw --resolves KEY --reason TEXT --by WHO
  xeno assumption record --intent KEY --phase NN --text TEXT --origin WHERE --confidence HOW [--resolves KEY]
  xeno assumption confirm ID --intent KEY --by WHO
  xeno assumption reject  ID --intent KEY --by WHO
  xeno gate verify    [--intent KEY]            recompute and compare, write nothing (CI)
  xeno intent verify  --base REF --head REF     every intent the range touches is finished or closed (CI)
  xeno enforcement check [--branch NAME]        ask the host what it enforces (needs the network)
  xeno evidence attach --intent KEY --phase NN --from DIR
  xeno intent status  [--intent KEY] [--all]    without one, the last ten by creation
  xeno intent close   --intent KEY --reason TEXT
  xeno section set    SECTION --intent KEY --phase NN [--file PATH]   reads stdin without --file
  xeno check commit-message [--pattern NAME] [--file PATH]   reads stdin without --file
  xeno cost turn                                reads a hook's JSON on stdin
  xeno version
common: --root DIR (default .), --no-next to leave out the next step
        --tool-version V on section set and phase finish, which record what wrote a phase;
        XENO_HARNESS_VERSION says the same thing for a whole session`

// main is the only place the real files appear. Everything below writes through what it is
// given, so that the exit code staircase this package promises can be asserted in process
// rather than by running the binary (#110).
func main() { os.Exit(run(os.Args[1:], os.Stdout, os.Stderr)) }

// opts is everything a command was given: the flags, the finding id that stands before
// them, the resolved phase and the runner they act on. One struct rather than twenty
// arguments, so that a command reads what it needs and a new flag reaches every command
// without touching any of them.
type opts struct {
	r       *runner.Runner
	cmd     string
	finding string
	phase   string

	root, key, phaseArg          string
	from, src, by, pattern, file string
	summary                      string
	all                          bool
	// Where this command writes. Supplied rather than global, so a test can read it.
	out, errw                      io.Writer
	project, mdl, language         string
	pluginFrom, host, branch       string
	base, head, reason             string
	issue, toolVersion             string
	category, observation          string
	proposal, target               string
	noFinding                      bool
	text, origin, confidence       string
	resolves, chosen, proposedBy   string
	withdraw                       bool
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
	"intent verify":        {run: cmdIntentVerify},
	"cost turn":            {run: cmdCostTurn},
	"intent start":         {run: cmdIntentStart},
	"intent status":        {run: cmdIntentStatus},
	"intent close":         {needsKey: true, run: cmdIntentClose},
	// The phase is optional alone among the commands that take one: section 10 owes a
	// record at the end of every phase and once more when an intent closes, and the
	// intent level record is the one G-Complete reads from `intent close`.
	"learning record":    {needsKey: true, run: cmdLearningRecord},
	"assumption confirm": {needsKey: true, run: cmdAssumptionDecide},
	"assumption reject":  {needsKey: true, run: cmdAssumptionDecide},
	"section set":        {needsKey: true, needsPhase: true, run: cmdSectionSet},
	"phase start":        {needsKey: true, needsPhase: true, run: cmdPhaseStart},
	"phase finish":       {needsKey: true, needsPhase: true, run: cmdPhaseFinish},
	"gate run":           {needsKey: true, needsPhase: true, run: cmdGateRun},
	"gate approve":       {needsKey: true, needsPhase: true, run: cmdGateApprove},
	"gate override":      {needsKey: true, needsPhase: true, run: cmdGateOverride},
	"assumption record":  {needsKey: true, needsPhase: true, run: cmdAssumptionRecord},
	"question record":    {needsKey: true, needsPhase: true, run: cmdQuestionRecord},
	"decision record":    {needsKey: true, needsPhase: true, run: cmdDecisionRecord},
	"obligation close":   {needsKey: true, needsPhase: true, run: cmdObligationClose},
	"evidence attach":    {needsKey: true, needsPhase: true, run: cmdEvidenceAttach},
}

func run(args []string, out, errw io.Writer) int {
	if len(args) >= 1 && args[0] == "version" {
		fmt.Fprintln(out, "xeno", model.RunnerVersion)
		return 0
	}
	if len(args) < 2 && (len(args) == 0 || args[0] != "init") {
		fmt.Fprintln(errw, usage)
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
		fmt.Fprintln(errw, usage)
		return 2
	}

	o, code := parse(name, rest, out, errw)
	if code != 0 {
		return code
	}
	if c.needsKey && o.key == "" {
		fmt.Fprintln(errw, "--intent is required")
		return 2
	}
	if c.needsPhase {
		phase, err := model.ResolvePhase(o.phaseArg)
		if err != nil {
			fmt.Fprintln(errw, err)
			return 2
		}
		o.phase = phase
	}
	return c.run(o)
}

// parse takes the finding id off the front and reads the flags. The id stands before them
// as the process definition writes these commands, and Go's flag package stops at the
// first argument that is not a flag, so it cannot be read back out afterwards.
func parse(name string, args []string, out, errw io.Writer) (*opts, int) {
	o := &opts{cmd: name, out: out, errw: errw}
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
	fs.StringVar(&o.host, "host", "", "which host to generate for: the CI wrapper and the tracker block")
	fs.StringVar(&o.branch, "branch", "", "the branch whose protection to read, main by default")
	// Both ends of the commit range. They are an input of the run and are deliberately
	// not recorded: after a squash a recorded range would point at commits that no
	// longer exist, and a field that is sometimes wrong is worse than no field. No gate
	// reads them until WP4 brings the commit predicates.
	fs.StringVar(&o.base, "base", "", "the base of the commit range under review")
	fs.StringVar(&o.head, "head", "", "the head of the commit range under review")
	fs.StringVar(&o.reason, "reason", "", "why")
	fs.StringVar(&o.issue, "for", "", "the issue an intent belongs to, as a key or as the whole qualified id")
	// The one field of section 5 the runner cannot know, so the harness says it. Absent
	// where it is not given, which is what every artifact written before this carries.
	fs.StringVar(&o.toolVersion, "tool-version", "",
		"the version of the harness writing this phase, overriding "+runner.HarnessVersionEnv)
	fs.StringVar(&o.text, "text", "", "the statement being assumed")
	fs.StringVar(&o.origin, "origin", "", "where the assumption came from")
	fs.StringVar(&o.confidence, "confidence", "", "how much weight it carries")
	fs.StringVar(&o.resolves, "resolves", "", "the open question this answers")
	// Section 8's exchange. --chosen is the option taken in the words of whoever took it,
	// which for the free entry is their own words and not an aside beside them. --reason
	// carries the rationale, because why is already this flag's word on gate approve and
	// intent close and section 8 calls a decision's rationale and a withdrawal's reason the
	// same thing.
	fs.StringVar(&o.chosen, "chosen", "", "the option taken, in the words of whoever took it")
	fs.StringVar(&o.proposedBy, "proposed-by", "", "who put the options, where that was not the person deciding")
	fs.BoolVar(&o.withdraw, "withdraw", false, "section 8's third exit: the question is dropped, with a reason")
	// Section 10's four keys, and the statement that there were none. A record carrying
	// neither is what G-Learning calls empty, so one of the two has to be said.
	fs.StringVar(&o.category, "category", "", "which kind of learning, from section 10's four")
	fs.StringVar(&o.observation, "observation", "", "what was observed")
	fs.StringVar(&o.proposal, "proposal", "", "what should change because of it")
	fs.StringVar(&o.target, "target", "", "what the proposal applies to")
	fs.BoolVar(&o.noFinding, "no-finding", false, "there was nothing to record, said rather than left out")
	// The suggestion is off by default nowhere and on by default nowhere either: the
	// commands that change state say it, the ones a pipeline or a hook runs do not, and
	// this turns it off for the scripts that are neither.
	fs.BoolVar(&o.noNext, "no-next", false, "do not say what the next step is")
	// Print what a shell can eval instead of saying what the next step is. Both on one
	// stream would make the eval swallow a sentence meant for a person.
	fs.BoolVar(&o.export, "export", false, "print the phase's environment for a shell to eval")
	fs.BoolVar(&o.all, "all", false, "list every intent, not the last ten")
	if err := fs.Parse(args); err != nil {
		return nil, 2
	}

	o.r = runner.New(o.root)
	o.r.EvidenceFrom = o.from
	// The flag beats the variable, and silence leaves what New read from the environment.
	// An argument is somebody saying it about this run; the variable is whatever set it,
	// which for a session is usually the entry point and not a person.
	if o.toolVersion != "" {
		o.r.ToolVersion = o.toolVersion
	}
	return o, 0
}

// next is the suggestion every command that changes state prints, and the reason each of
// these functions ends with it rather than returning a code directly.
func (o *opts) next(code int) int { return o.suggest(o.r, o.key, o.noNext, code) }

func cmdInit(o *opts) int {
	o.r.PluginSource = o.pluginFrom
	res, err := o.r.Init(runner.InitOptions{
		TrackerKey: o.project, Model: o.mdl, Language: o.language, Vendor: o.vendor, Host: o.host,
	})
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 1
	}
	printInit(o.out, res)
	return 0
}

func cmdEnforcementCheck(o *opts) int {
	rep, err := o.r.EnforcementCheck(o.branch, enforcement.Token())
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 2
	}
	printEnforcement(o.out, rep)
	if rep.Unmet() > 0 {
		return 1
	}
	return 0
}

func cmdCheckMessage(o *opts) int {
	message, err := readMessage(o.file)
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 2
	}
	if err := gates.CheckMessage(o.pattern, message); err != nil {
		fmt.Fprintln(o.errw, err)
		return 1
	}
	return 0
}

func cmdSectionSet(o *opts) int {
	content, err := readMessage(o.file)
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 2
	}
	t, err := o.r.SectionSet(o.key, o.phase, o.finding, content)
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 1
	}
	fmt.Fprintf(o.out, "%s rendered from %s (%s)\n", o.phase, t.Ref(), t.Source)
	return o.next(0)
}

// cmdQuestionRecord writes one open_questions entry. The entry is read whole rather than
// assembled from flags: two to four options with a consequence each, a recommendation and the
// free entry are a nested structure however they are spelled, and section set already reads
// stdin for the same reason.
func cmdQuestionRecord(o *opts) int {
	entry, err := readMessage(o.file)
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 2
	}
	q, err := o.r.RecordQuestion(o.key, o.phase, []byte(entry))
	if code := o.report(nil, err); code != 0 {
		return code
	}
	fmt.Fprintf(o.out, "%s %s raises %s\n", o.key, o.phase, q.Key)
	return o.next(0)
}

// cmdDecisionRecord writes one decisions entry, which is the half of section 8's exchange
// that needs a person. Nothing here defaults --by.
func cmdDecisionRecord(o *opts) int {
	d, err := o.r.RecordDecision(o.key, o.phase, model.Decision{
		Resolves: o.resolves, Chosen: o.chosen, Rationale: o.reason,
		DecidedBy: o.by, ProposedBy: o.proposedBy, Withdrawn: o.withdraw,
	})
	if code := o.report(nil, err); code != 0 {
		return code
	}
	line := fmt.Sprintf("%s %s records %s, decided by %s", o.key, o.phase, d.ID, d.DecidedBy)
	if d.Resolves != "" {
		verb := "resolving"
		if d.Withdrawn {
			verb = "withdrawing"
		}
		line += ", " + verb + " " + d.Resolves
	}
	fmt.Fprintln(o.out, line)
	return o.next(0)
}

// cmdIntentStart takes --intent optionally, alone among the commands that name an intent.
// Everywhere else the key says which intent is meant and there is nothing to derive it
// from; here the intent does not exist yet, and the key of a new one is the next number of
// the sequence the repository already holds. Given, it is used as given, which is what the
// first intent of a repository needs and what the older naming scheme is kept readable by.
func cmdIntentStart(o *opts) int {
	in, err := o.r.IntentStart(o.key, o.issue)
	if code := o.report(nil, err); code != 0 {
		return code
	}
	fmt.Fprintf(o.out, "%s created for %s, %s\n", in.Key, in.Intent, in.Status)
	o.key = in.Key
	return o.next(0)
}

func cmdIntentClose(o *opts) int { return o.next(o.report(o.r.IntentClose(o.key, o.reason))) }

func cmdPhaseStart(o *opts) int {
	if code := o.report(nil, o.r.Start(o.key, o.phase)); code != 0 {
		return code
	}
	if o.export {
		env, err := o.r.PhaseEnv(o.key, o.phase)
		if err != nil {
			fmt.Fprintln(o.errw, err)
			return 1
		}
		fmt.Fprint(o.out, env)
		return 0
	}
	// What a repeated phase has to read again, from the predecessor's lock and the tree. Printed
	// rather than recorded: section 5's field list has no entry for it, and the lock states what
	// was declared rather than what was read. Silent where nothing moved, which is also what a
	// repository with no context profile gets, since it declared no base to move.
	if changed := o.r.ChangedSince(o.key, o.phase); len(changed) > 0 {
		fmt.Fprintf(o.out, "\nchanged since %s, and nothing else needs rereading:\n", previousPhase(o.phase))
		for _, p := range changed {
			fmt.Fprintln(o.out, "  "+p)
		}
	}
	return o.next(0)
}

// previousPhase names the phase a changed set is measured against, for the one line that reports
// it. The caller has already established that the phase is not the first.
func previousPhase(phase string) string {
	return model.Phases[model.PhaseIndex(phase)-1]
}

// cmdPhaseFinish passes the summary the digest is written from. A path, or "-" for stdin:
// absence cannot mean stdin here as it does for section set, because a summary is optional and
// a phase finished without one is judged exactly as before (section 5, #120).
func cmdPhaseFinish(o *opts) int {
	summary := ""
	if o.summary != "" {
		var err error
		if summary, err = readSummary(o.summary); err != nil {
			fmt.Fprintln(o.errw, err)
			return 2
		}
		if strings.TrimSpace(summary) == "" {
			fmt.Fprintln(o.errw, "--summary is empty: a digest with no summary is worse than none, since G-Schema would pass it")
			return 2
		}
	}
	return o.next(o.report(o.r.Finish(o.key, o.phase, summary)))
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
	return o.next(o.report(o.r.GateRun(o.key, o.phase)))
}

func cmdGateApprove(o *opts) int {
	return o.next(o.report(o.r.Decide(o.key, o.phase, o.finding, "approved", o.by, o.reason)))
}

func cmdGateOverride(o *opts) int {
	return o.next(o.report(o.r.Decide(o.key, o.phase, o.finding, "overridden", o.by, o.reason)))
}

func cmdAssumptionRecord(o *opts) int {
	a, err := o.r.RecordAssumption(o.key, o.phase, o.text, o.origin, o.confidence, o.resolves)
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 1
	}
	fmt.Fprintf(o.out, "%s recorded, open, from %s with %s confidence\n", a.ID, a.Origin, a.Confidence)
	return o.next(0)
}

func cmdLearningRecord(o *opts) int {
	phase := ""
	if o.phaseArg != "" {
		var err error
		if phase, err = model.ResolvePhase(o.phaseArg); err != nil {
			fmt.Fprintln(o.errw, err)
			return 2
		}
	}
	rec, err := o.r.RecordLearning(o.key, phase, o.noFinding, model.LearningEntry{
		Category: o.category, Observation: o.observation,
		Proposal: o.proposal, Target: o.target,
	})
	if code := o.report(nil, err); code != 0 {
		return code
	}
	where := o.key
	if phase != "" {
		where += " " + phase
	}
	if rec.NoFinding {
		fmt.Fprintf(o.out, "%s records no finding\n", where)
	} else {
		fmt.Fprintf(o.out, "%s records %d learning(s), the last of category %s\n",
			where, len(rec.Learnings), rec.Learnings[len(rec.Learnings)-1].Category)
	}
	return o.next(0)
}

func cmdAssumptionDecide(o *opts) int {
	status := "confirmed"
	if o.cmd == "assumption reject" {
		status = "rejected"
	}
	a, err := o.r.DecideAssumption(o.key, o.finding, status, o.by)
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 1
	}
	fmt.Fprintf(o.out, "%s %s by %s\n", a.ID, a.Status, o.by)
	return o.next(0)
}

func cmdObligationClose(o *opts) int {
	return o.next(o.report(o.r.CloseObligation(o.key, o.phase, o.finding)))
}

func cmdEvidenceAttach(o *opts) int {
	res, err := evidence.Attach(o.root, o.key, o.phase, o.src)
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 2
	}
	fmt.Fprintf(o.out, "attached %d, pending %d; run xeno gate run to carry the verdict forward\n",
		res.Attached, res.Pending)
	// Named on stderr and counted as pending: an entry nothing can bind was not attached,
	// and a run that only printed the counts would report it as evidence still to come
	// from a job that has already produced it.
	for _, u := range res.Unbindable {
		fmt.Fprintln(o.errw, "  not attached:", u)
	}
	if len(res.Unbindable) > 0 {
		fmt.Fprintln(o.errw, "  republish with a sha256; a uri is bound by its hash alone.")
		return o.next(1)
	}
	return o.next(0)
}

func cmdGateVerify(o *opts) int {
	res, err := o.r.Verify(o.key)
	if err != nil {
		fmt.Fprintln(o.errw, "error:", err)
		return 2
	}
	fmt.Fprintf(o.out, "verified %d verdicts\n", res.Checked)
	for _, d := range res.Divergences {
		fmt.Fprintf(o.out, "  DIVERGENT   %s %s: %s\n", d.Key, d.Phase, d.What)
	}
	for _, l := range res.Red {
		fmt.Fprintf(o.out, "  RED         %s\n", l)
	}
	for _, l := range res.Provisional {
		fmt.Fprintf(o.out, "  PROVISIONAL %s: evidence outstanding, binding only at the merge request\n", l)
	}
	return verifyCode(res)
}

// verifyCode is the staircase for a verification, separated from the printing so that the one
// branch a CI wrapper depends on can be asserted without building a phase to produce it.
//
// A provisional verdict exits 0. Section 6 rests the evidence arrangement on that: the wrapper
// reports provisional, N declarations open rather than failing, because a P4 that failed
// verification for waiting on a pipeline is what section 6 names as the outcome that would make
// people stop taking verification seriously. A divergence or a red phase is 1 (#110).
func verifyCode(res *runner.VerifyResult) int {
	if len(res.Divergences) > 0 || len(res.Red) > 0 {
		return 1
	}
	return 0
}

// cmdIntentVerify is the merge check for section 8's two endings. An intent the change
// touches has either reached a decided P5 or been closed, and anything else is a trail that
// stops in the middle with every gate green on the phases it did write (#206).
//
// The range is required and never inferred, as section 9 says and as Commits already refuses:
// a base that does not resolve exits 2, "could not run", because a comparison made on no
// evidence reports the same thing as a clean one.
//
// A range that touches no intent says so in those words rather than printing nothing. That
// every change belongs to an intent is a rule of this project and an unenforced one, and
// enforcing it from inside a check about something else would bury it (#120).
func cmdIntentVerify(o *opts) int {
	res, err := o.r.Completeness(o.base, o.head)
	if err != nil {
		fmt.Fprintln(o.errw, "error:", err)
		return 2
	}
	if len(res.Touched) == 0 {
		fmt.Fprintln(o.out, "no intent is touched by this change")
		return 0
	}
	intents := "intents"
	if len(res.Touched) == 1 {
		intents = "intent"
	}
	fmt.Fprintf(o.out, "checked %d %s the change touches\n", len(res.Touched), intents)
	for _, s := range res.Unfinished {
		what := s.State
		if what == "" {
			what = s.Problem
		}
		fmt.Fprintf(o.out, "  UNFINISHED  %s: %s\n", s.Key, what)
	}
	if len(res.Unfinished) > 0 {
		fmt.Fprintln(o.out, "An intent ends at a decided P5 or at xeno intent close.")
		fmt.Fprintln(o.out, "Finish the phases, or close it with a reason: both are recorded,")
		fmt.Fprintln(o.out, "and a trail that stops in the middle is neither.")
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
		fmt.Fprintln(o.errw, err)
		return 2
	}
	if len(states) > 0 {
		fmt.Fprintln(o.out, sprintRow(phaseRow, "PHASE", "STATE", "VERDICT"))
	}
	for _, s := range states {
		line := fmt.Sprintf(phaseRow, s.Phase, s.State, s.Status)
		if s.Stale {
			line += "  (predecessor changed since this phase started)"
		}
		fmt.Fprintln(o.out, line)
	}
	// After the table, where the truncation note goes in the other form: it is a figure
	// about the intent and not a row of it. The sentence names both halves, because a trail
	// with a decision and no question is a different state from a trail with neither.
	if s := o.r.Summarise(o.key); s.AskedNothing() {
		fmt.Fprintln(o.out, "\nno question and no decision in any phase, so G-Questions has nothing to verify")
	}
	return o.next(0)
}

// cmdIntentList prints the intents oldest first. A row whose intent.yaml could not be read
// says so in place of its date rather than going missing.
// listRow is the format the heading and every row share, which is what keeps a heading from
// drifting from the column it labels. The state column is seventeen wide because the longest one
// is 03-implementation.
//
// The headings are upper case, as ps prints PID TTY TIME CMD and everything descended from it
// does. That is the convention for a label on tabular output, and a reader scanning a terminal
// recognises a row of capitals as the line that is not data. This project's rule that a heading
// names its section in words is about prose, in files and in issues; XENO-0207 cited it for a
// column label and set these in lower case, which #136 reversed.
const listRow = "%-10s  %-12s %-17s %s"

// listDefault is how many intents the listing shows without --all. A listing answers what is
// happening; sixty rows, fifty of which stopped at the intake, answer what has ever happened, and
// --all is there for that. Ten is a screen.
const listDefault = 10

// phaseRow is the same arrangement for the one-intent form, where state is a position and
// verdict is a judgement, which is the pair a heading is worth most for.
const phaseRow = "%-18s %-22s %s"

// tail says how many of n intents a listing shows and how many it leaves out. Separated from the
// printing so that the arithmetic can be asserted without capturing output.
func tail(n int, all bool) (shown, hidden int) {
	if all || n <= listDefault {
		return n, 0
	}
	return listDefault, n - listDefault
}

// sprintRow renders one row of a table, and exists so that a test can compare a heading against a
// row built from the same format string.
func sprintRow(format string, cells ...any) string {
	return fmt.Sprintf(format, cells...)
}

func cmdIntentList(o *opts) int {
	intents, err := o.r.Intents()
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 2
	}
	// The tail, taken here rather than in the runner: Intents() returns everything in order and
	// knows nothing about presentation, so this cannot change what is shown (#134).
	shown, hidden := tail(len(intents), o.all)
	intents = intents[len(intents)-shown:]
	// With the first row rather than before the loop, so that a repository holding no intents
	// prints nothing at all instead of a label for an absence.
	if len(intents) > 0 {
		fmt.Fprintln(o.out, sprintRow(listRow, "CREATED", "INTENT", "STATE", "PHASE"))
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
		// The figure, after the columns where the unreadable-record note already goes. A
		// mark and not a column: two more columns of numbers on every row would carry a
		// fact about a handful of intents, and #136 settled the widths.
		if s.AskedNothing() {
			line += "  (no question and no decision)"
		}
		fmt.Fprintln(o.out, strings.TrimRight(line, " "))
	}
	// After the table, so it is the last thing read and cannot be taken for a row. A truncated
	// listing that said nothing would read as a repository holding ten intents, which is a wrong
	// fact rather than a missing one.
	if hidden > 0 {
		fmt.Fprintf(o.out, "\n%d older, --all to see them\n", hidden)
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

func (o *opts) suggest(r *runner.Runner, key string, off bool, code int) int {
	if off || code == 2 {
		return code
	}
	s := r.Next(key)
	fmt.Fprintf(o.out, "\nnext: %s\n", s.Text)
	if s.Command != "" {
		fmt.Fprintf(o.out, "  %s\n", s.Command)
	}
	for _, owed := range s.Owed {
		fmt.Fprintf(o.out, "  owed, whenever somebody gets to it:\n  %s\n", owed)
	}
	return code
}

func (o *opts) report(g *model.Gate, err error) int {
	var ref *runner.Refusal
	switch {
	case errors.As(err, &ref):
		fmt.Fprintln(o.errw, "refused:", ref.Reason)
		return 1
	case err != nil:
		fmt.Fprintln(o.errw, "error:", err)
		return 2
	case g == nil:
		return 0
	}
	fmt.Fprintf(o.out, "%s %s: %s\n", g.Intent, g.Phase, g.Status)
	for _, c := range g.Checks {
		fmt.Fprintf(o.out, "  %-14s %s\n", c.Gate, c.Result)
		for _, f := range c.Findings {
			d := ""
			if f.Decision != nil {
				d = " [" + f.Decision.Type + "]"
			}
			fmt.Fprintf(o.out, "      %s %s: %s%s\n        next: %s\n", f.ID, f.File, f.Cause, d, f.Next)
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
func printInit(out io.Writer, res *runner.InitResult) {
	if res.PluginFrom != "" {
		fmt.Fprintln(out, "  plugin from", res.PluginFrom)
	}
	for _, p := range res.Created {
		fmt.Fprintln(out, "  created  ", p)
	}
	for _, p := range res.Kept {
		fmt.Fprintln(out, "  kept     ", p)
	}
	if len(res.Outstand) > 0 {
		fmt.Fprintln(out, "\nNot determined:")
		for _, s := range res.Outstand {
			fmt.Fprintln(out, "  -", s)
		}
	}
	fmt.Fprintln(out, "\nXeno cannot make these settings. Somebody with repository administration has to:")
	for _, s := range res.Manual {
		fmt.Fprintln(out, "  -", s)
	}
}

// printEnforcement prints the report. not-available is its own line rather than folded
// into unmet, because a setting the host does not have is nobody's oversight and
// reporting it as one sends somebody looking for a checkbox that is not there.
func printEnforcement(out io.Writer, rep *enforcement.Report) {
	fmt.Fprintf(out, "%s, branch %s\n", rep.Repository, rep.Branch)
	for _, q := range rep.Requirements {
		fmt.Fprintf(out, "  %-13s %-20s declared %-6s actual %s\n", q.State, q.Name, q.Declared, q.Actual)
		if q.Note != "" {
			fmt.Fprintf(out, "                  %s\n", q.Note)
		}
	}
	fmt.Fprintf(out, "\nreport written to %s\n", runner.ReportPath)
	if rep.Unmet() > 0 {
		fmt.Fprintf(out, "\n%d requirement(s) are neither met nor waived. Where the host cannot express one,\n"+
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
