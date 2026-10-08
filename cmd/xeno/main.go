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
	"github.com/triplem/xeno/internal/learning"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/runner"
)

const usage = `usage:
  xeno init           --host HOST [--vendor] [--project OWNER/REPO] [--model ID] [--language TAG]
  xeno intent start   --for ISSUE [--intent KEY]        writes intent.yaml
  xeno phase start    --intent KEY --phase NN [--evidence-from DIR] [--export]
  xeno phase finish   --intent KEY --phase NN [--summary PATH|-]   writes digest.md
  xeno gate run       --intent KEY --phase NN [--base REF --head REF] [--evidence-from DIR]
  xeno gate approve   FINDING --intent KEY --phase NN --by WHO --reason TEXT
  xeno gate override  FINDING --intent KEY --phase NN --by WHO --reason TEXT
  xeno obligation close FINDING --intent KEY --phase NN
  xeno learning record --intent KEY [--phase NN] --category C --observation T --proposal T --target P
  xeno learning record --intent KEY [--phase NN] --no-finding
  xeno learning propose --out DIR [--intent KEY]   generates the merge request against the rule set
  xeno question record --intent KEY --phase NN [--file PATH]   reads the entry on stdin
  xeno decision record --intent KEY --phase NN --chosen TEXT --reason TEXT --by WHO [--resolves KEY] [--proposed-by WHO]
  xeno decision record --intent KEY --phase NN --withdraw --resolves KEY --reason TEXT --by WHO
  xeno review answer  RULE --intent KEY --result R [--note TEXT]   answers one review rule
  xeno review lens    --intent KEY --result R --note TEXT   a lens's entry, which answers no rule
  xeno assumption record --intent KEY --phase NN --text TEXT --origin WHERE --confidence HOW [--resolves KEY]
  xeno assumption confirm ID --intent KEY --by WHO
  xeno assumption reject  ID --intent KEY --by WHO
  xeno gate verify    [--intent KEY]            recompute and compare, write nothing (CI)
  xeno intent verify  --base REF --head REF     every intent the range touches is finished or closed (CI)
  xeno enforcement check [--branch NAME]        ask the host what it enforces (needs the network)
  xeno report verdict --intent KEY [--phase NN] [--dry-run]   the verdict onto the issue (CI, needs the network)
  xeno report figures [--intent KEY] [--all]    cost, artifacts and reopens per intent, read from the trail
  xeno evidence declare --intent KEY --phase NN --kind K [--job J] [--result R]
                        [--file PATH | --uri URL --sha256 HEX] [--produced-by CMD] [--format F]
  xeno evidence attach --intent KEY --phase NN --from DIR
  xeno intent status  [--intent KEY] [--all]    without one, the last ten by creation
  xeno intent close   --intent KEY --reason TEXT
  xeno section set    SECTION --intent KEY --phase NN [--file PATH]   reads stdin without --file
  xeno scope set      --intent KEY [--file PATH]   P0's context scope, read from stdin
  xeno template show  --phase NN                the sections a phase owes, in order
  xeno symbol show    NAME                      where a name is defined, from the project's index
  xeno check commit-message [--pattern NAME] [--file PATH]   reads stdin without --file
  xeno cost turn                                reads a hook's JSON on stdin
  xeno mcp            [--root DIR]              the six process operations over stdio
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
	dest                           string
	text, origin, confidence       string
	resolves, chosen, proposedBy   string
	kind, job, result              string
	note                           string
	uri, sha256                    string
	producedBy, format             string
	withdraw                       bool
	dry                            bool
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
	"init":              {run: cmdInit},
	"enforcement check": {run: cmdEnforcementCheck},
	// The phase is optional, alone among the commands that take an intent and a phase: a job
	// reporting after a run has judged whatever it judged, and the trail knows which phase
	// that was, so the pipeline does not have to.
	"report verdict": {needsKey: true, run: cmdReportVerdict},
	// No key at all, and for `learning propose`'s reason: the figures are a reading of the
	// whole trail and one intent is a narrowing of it. The command writes nothing and no gate
	// reads it, which is WP15's rule about the symbol index applied to a measurement.
	"report figures":       {run: cmdReportFigures},
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
	"learning record": {needsKey: true, run: cmdLearningRecord},
	// No key either, and for the opposite reason: the route reads what a whole trail
	// proposed, and one intent is a narrowing of that rather than the case.
	"learning propose":   {run: cmdLearningPropose},
	"assumption confirm": {needsKey: true, run: cmdAssumptionDecide},
	"assumption reject":  {needsKey: true, run: cmdAssumptionDecide},
	"section set":        {needsKey: true, needsPhase: true, run: cmdSectionSet},
	"phase start":        {needsKey: true, needsPhase: true, run: cmdPhaseStart},
	"phase finish":       {needsKey: true, needsPhase: true, run: cmdPhaseFinish},
	"gate run":           {needsKey: true, needsPhase: true, run: cmdGateRun},
	"gate approve":       {needsKey: true, needsPhase: true, run: cmdGateApprove},
	"gate override":      {needsKey: true, needsPhase: true, run: cmdGateOverride},
	"assumption record":  {needsKey: true, needsPhase: true, run: cmdAssumptionRecord},
	"scope set":          {needsKey: true, run: cmdScopeSet},
	"question record":    {needsKey: true, needsPhase: true, run: cmdQuestionRecord},
	"decision record":    {needsKey: true, needsPhase: true, run: cmdDecisionRecord},
	// No phase: G-Policy judges the checklist only on the last phase, so there is one
	// phase it can mean and the writer resolves it, as scope set does for P0.
	"review answer": {needsKey: true, run: cmdReviewAnswer},
	// No phase either, and no positional: a lens entry answers no rule, so there is no
	// argument for one and cmdReviewLens refuses the stray the parser would otherwise drop.
	"review lens":      {needsKey: true, run: cmdReviewLens},
	"obligation close": {needsKey: true, needsPhase: true, run: cmdObligationClose},
	"evidence attach":  {needsKey: true, needsPhase: true, run: cmdEvidenceAttach},
	"evidence declare": {needsKey: true, needsPhase: true, run: cmdEvidenceDeclare},
	// The two reads behind fetch_template and query_symbol_index. A phase and no intent is
	// this pair alone: a template is resolved from the plugin or the project and the index
	// from .xeno/local/, so neither answer differs by which intent is open, and asking for a
	// key would be asking for one that decides nothing.
	"template show": {needsPhase: true, run: cmdTemplateShow},
	"symbol show":   {run: cmdSymbolShow},
}

func run(args []string, out, errw io.Writer) int {
	if len(args) >= 1 && args[0] == "version" {
		fmt.Fprintln(out, "xeno", model.RunnerVersion)
		return 0
	}
	if len(args) >= 1 && args[0] == "mcp" {
		return cmdMCP(args[1:], out, errw)
	}
	// Commands are two words except init, which is one. The split is where the flags
	// begin, not a property of the name, and the one word case is decided here rather
	// than corrected afterwards: the guard used to admit `init` with a single argument
	// and the line that split the name then indexed args[1], so the first command
	// anybody runs in an empty repository answered with a stack trace (#303).
	var name string
	var rest []string
	switch {
	case len(args) >= 1 && args[0] == "init":
		name, rest = "init", args[1:]
	case len(args) >= 2:
		name, rest = args[0]+" "+args[1], args[2:]
	default:
		fmt.Fprintln(errw, usage)
		return 2
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
	// The directory a job left the pipeline's results in. The runner reads it and does not
	// fetch: a gate never touches the network, so the fetch is a step of the job.
	fs.StringVar(&o.from, "evidence-from", "", "directory a job left the pipeline's results in")
	fs.StringVar(&o.src, "from", "", "directory a job left the pipeline's results in")
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
	// Section 4's declaration. --file is the report itself here and not a message to read,
	// which is the one place that flag means something else; --from was the alternative and
	// is the attach's word for a directory, and one flag naming a file here and a directory
	// there is how a reader learns to read it twice. There is no flag for the hash of a
	// file: where the bytes are present the runner computes it.
	fs.StringVar(&o.kind, "kind", "", "which kind of evidence, from section 4's set")
	fs.StringVar(&o.job, "job", "", "the job expected to produce it")
	fs.StringVar(&o.result, "result", "", "what the run reported against its own threshold, or how a review rule was answered")
	fs.StringVar(&o.note, "note", "", "why a review rule was passed over")
	fs.StringVar(&o.uri, "uri", "", "where a large or binary result lies, instead of --file")
	fs.StringVar(&o.sha256, "sha256", "", "the hash the pipeline published, with --uri")
	fs.StringVar(&o.producedBy, "produced-by", "", "the command as run")
	fs.StringVar(&o.format, "format", "", "the shape of the report, from section 4's set")
	// Section 10's four keys, and the statement that there were none. A record carrying
	// neither is what G-Learning calls empty, so one of the two has to be said.
	fs.StringVar(&o.category, "category", "", "which kind of learning, from section 10's four")
	fs.StringVar(&o.observation, "observation", "", "what was observed")
	fs.StringVar(&o.proposal, "proposal", "", "what should change because of it")
	fs.StringVar(&o.target, "target", "", "what the proposal applies to")
	// Where the merge request is written. A destination is the one thing the route cannot
	// derive: the trail it reads is the one place it may not write, so there is no default
	// that would be right.
	fs.StringVar(&o.dest, "out", "", "where the generated merge request is written, outside the trail it was read from")
	fs.BoolVar(&o.noFinding, "no-finding", false, "there was nothing to record, said rather than left out")
	// The suggestion is off by default nowhere and on by default nowhere either: the
	// commands that change state say it, the ones a pipeline or a hook runs do not, and
	// this turns it off for the scripts that are neither.
	fs.BoolVar(&o.noNext, "no-next", false, "do not say what the next step is")
	// Print what a shell can eval instead of saying what the next step is. Both on one
	// stream would make the eval swallow a sentence meant for a person.
	fs.BoolVar(&o.export, "export", false, "print the phase's environment for a shell to eval")
	fs.BoolVar(&o.all, "all", false, "list every intent, not the last ten")
	// Compose the comment and print it, writing nothing. It is how a change to the wording is
	// read before anybody receives it, and how a job logs what it would have said on a run
	// that holds no credential to say it with.
	fs.BoolVar(&o.dry, "dry-run", false, "print the comment the write-back would post, and post nothing")
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
		// A missing --host is where init could not run at all rather than refused to act:
		// it is the one flag the command cannot default, so it belongs on the same step of
		// the staircase as the --intent the dispatch asks for above. Init's own message
		// names the two hosts it knows, so the reason is written in one place.
		if strings.TrimSpace(o.host) == "" {
			return 2
		}
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
	printEnforcement(o.out, o.root, rep)
	if rep.Unmet() > 0 {
		return 1
	}
	return 0
}

// cmdReportVerdict is section 12's write-back: the phase, the verdict and the findings with
// their decisions, on the issue the intent came from, with the enforcement report beside
// them. It is the step a pipeline runs after the verdict and never a gate, because the gate
// path reaches no network and every verdict rests on that.
//
// The comment is printed as well as posted. A pipeline log that holds what was said is what
// makes the step readable on a run where the host declined the write, and the body is the
// same string either way.
func cmdReportVerdict(o *opts) int {
	// The phase is resolved here rather than by the dispatch, because the dispatch resolves
	// it for every command that declares one and this command's is optional: an empty value
	// means the last phase holding a verdict, and ResolvePhase would refuse it.
	phase := ""
	if o.phaseArg != "" {
		p, err := model.ResolvePhase(o.phaseArg)
		if err != nil {
			fmt.Fprintln(o.errw, err)
			return 2
		}
		phase = p
	}
	w, err := o.r.ReportVerdict(o.key, phase, o.dry)
	if code := o.report(nil, err); code != 0 {
		return code
	}
	fmt.Fprintln(o.out, w.Body)
	switch {
	case o.dry:
		fmt.Fprintf(o.out, "nothing was posted, because --dry-run. It would go to %s\n", w.Intent)
	case w.Posted && w.Written.URL != "":
		fmt.Fprintf(o.out, "posted to %s at %s\n", w.Intent, w.Written.URL)
	case w.Posted:
		fmt.Fprintf(o.out, "posted to %s\n", w.Intent)
	default:
		// Not an error and deliberately not a failing exit code either. A host that refuses
		// the write is the ordinary case on a run without the permission for it, and a job
		// that failed here would report a red step over a verdict it had already carried.
		fmt.Fprintln(o.errw, "nothing was posted:", w.Written.Reason)
	}
	return 0
}

// cmdReportFigures prints the three figures the plan names as deciding proportionality —
// cost per intent, artifacts per intent and reopens per intent — derived from the trail,
// together with the section count per template that WP3 names as the lever.
//
// It is the second half of `report`, and the opposite half: `report verdict` takes a verdict
// to the host, this takes nothing anywhere. It writes no file, posts nothing and no gate
// reads it, for the reason WP15 gives about the symbol index: a measurement a verdict could
// turn on stops being a measurement, because the cheapest way to move it is then to change
// what is counted.
//
// The figures were being counted by hand into comments on #117, which is the thing #117
// objected to: a number written down once, against a tree that has since moved. Printed by a
// command they are recomputable at any commit, and the comments become a record of what was
// true rather than the only place the numbers live.
func cmdReportFigures(o *opts) int {
	var keys []string
	if o.key != "" {
		keys = []string{o.key}
	}
	f, err := o.r.Figures(keys)
	if code := o.report(nil, err); code != 0 {
		return code
	}
	printPopulation(o.out, f)
	printFigureRows(o.out, f, o.all)
	printPerIntent(o.out, f)
	printCost(o.out, f)
	printTemplateBudget(o.out, f)
	return 0
}

// figuresRow is one intent's row. The first two columns are the listing's, so that a reader
// who has seen `xeno intent status` recognises the left edge; the rest are counts, and they
// are left aligned as every other table in this tool sets its cells.
const figuresRow = "%-10s  %-12s %-7s %-9s %-6s %-7s %-5s %s"

// templateBudgetRow is one template against WP3's budget: where it renders from, how many
// sections it requires, and how many of those are the phase's own rather than one of the two
// common ones the budget excludes.
const templateBudgetRow = "%-18s %-22s %-8s %-9s %-9s %s"

// budgetSpecific is the upper end of WP3's budget: three to four phase specific required
// sections per template, where adding a fifth means arguing another one away. A row over it
// is marked and nothing refuses: the template set is the plan's to change, and a report that
// failed on it would be a gate reading a measurement.
const budgetSpecific = 4

// printPopulation says what was counted before it says anything about it. A count over the
// trail includes the intent doing the counting, and the intents that stopped at the intake
// are a different measurement from the ones that went through, so both numbers are named and
// neither is folded into the other.
func printPopulation(out io.Writer, f *runner.Figures) {
	fmt.Fprintln(out, "POPULATION")
	fmt.Fprintf(out, "%d intents counted, %d with all six phases, %d that stopped earlier.\n",
		f.Counted, f.SixPhase, f.Counted-f.SixPhase)
	fmt.Fprintln(out, "Every figure per intent below is over the six-phase ones; the others are")
	fmt.Fprintln(out, "counted and not averaged in. The count includes the intent this was run for.")
}

// printFigureRows prints the rows oldest first, truncated the way the listing truncates: ten
// is a screen, and a hundred rows answer what has ever happened rather than what is
// happening. --all is there for the other question.
func printFigureRows(out io.Writer, f *runner.Figures, all bool) {
	rows := f.Selected
	shown, hidden := tail(len(rows), all)
	rows = rows[len(rows)-shown:]
	if len(rows) == 0 {
		return
	}
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, sprintRow(figuresRow,
		"CREATED", "INTENT", "PHASES", "SECTIONS", "FILES", "WORDS", "COST", "REOPENS"))
	for _, r := range rows {
		created := r.Created
		if created == "" {
			created = "?"
		}
		// The cost cell is a coverage and not a token count. A column of token sums over the
		// few phases that attributed anything would read as the cost of the intent, which is
		// the one thing #316 says this report may not print.
		line := sprintRow(figuresRow, created, r.Key,
			fmt.Sprintf("%d/%d", r.Phases, len(model.Phases)),
			fmt.Sprint(r.Sections), fmt.Sprint(r.Files), fmt.Sprint(r.Words),
			fmt.Sprintf("%d/%d", r.CostPhases, len(model.Phases)), fmt.Sprint(r.Reopens()))
		if r.Obligations > 0 {
			line += fmt.Sprintf("  (%d still owed)", r.Obligations)
		}
		if r.Moved > 0 {
			line += fmt.Sprintf("  (%d phase(s) moved since their verdict)", r.Moved)
		}
		if r.Problem != "" {
			line += "  (" + r.Problem + ")"
		}
		fmt.Fprintln(out, strings.TrimRight(line, " "))
	}
	if hidden > 0 {
		fmt.Fprintf(out, "\n%d older, --all to see them\n", hidden)
	}
}

// printPerIntent is the aggregate: the artifacts figure and the reopens figure, over the
// six-phase intents alone. The two counts beside reopens are printed rather than added into
// it, because they are different events and the figure the plan names is one number.
func printPerIntent(out io.Writer, f *runner.Figures) {
	fmt.Fprintf(out, "\nPER INTENT, over the %d with all six phases\n", f.SixPhase)
	if f.SixPhase == 0 {
		fmt.Fprintln(out, "nothing to average: no intent in the selection has all six phases.")
		return
	}
	mean := func(total int) float64 { m, _ := f.Mean(total); return m }
	fmt.Fprintf(out, "sections written   %8.1f\n", mean(f.Totals.Sections))
	fmt.Fprintf(out, "files              %8.1f\n", mean(f.Totals.Files))
	fmt.Fprintf(out, "words              %8.1f   across output.md, digest.md and learning.yaml\n",
		mean(f.Totals.Words))
	fmt.Fprintf(out, "reopens            %8.2f   released findings, approvals and overrides together\n",
		mean(f.Totals.Released))
	fmt.Fprintf(out, "%d obligation(s) still owed, %d phase(s) moved since their verdict.\n",
		f.Totals.Obligations, f.Totals.Moved)
	fmt.Fprintln(out, "The reopen figure is a proxy: a phase keeps one verdict, so the trail cannot")
	fmt.Fprintln(out, "say how often one was judged again. What it can say is where a person had to")
	fmt.Fprintln(out, "release a finding for a phase to pass.")
}

// printCost prints the cost figure with its coverage, or the coverage alone.
//
// The rule is the one #316 states: the figure travels with its coverage or it does not
// travel. A mean over the phases that recorded nothing would be a cost per intent orders of
// magnitude from the truth, so no mean is printed until every phase in the population has a
// record, and what is printed instead is what was attributed and how much of the population
// attributed it.
func printCost(out io.Writer, f *runner.Figures) {
	fmt.Fprintln(out, "\nCOST  self-reported, which section 11 requires of any cost figure shown")
	fmt.Fprintf(out, "%d of %d phases attributed anything.\n", f.Totals.CostPhases, f.PhasesCounted)
	t := f.Totals.Cost
	fmt.Fprintf(out, "tokens_in %d  tokens_out %d  tokens_cached %d, over those phases.\n",
		t.In, t.Out, t.Cached)
	if f.CostCovered() {
		in, _ := f.Mean(t.In)
		spent, _ := f.Mean(t.Out)
		cached, _ := f.Mean(t.Cached)
		fmt.Fprintf(out, "per intent: tokens_in %.0f  tokens_out %.0f  tokens_cached %.0f\n",
			in, spent, cached)
		return
	}
	fmt.Fprintln(out, "No cost per intent is printed, because not every phase counted has a record")
	fmt.Fprintln(out, "and a mean over the ones that do would be read as the cost of an intent.")
	fmt.Fprintln(out, "Within a phase that did attribute something the ledger holds the turns after")
	fmt.Fprintln(out, "`phase start` and no others, which internal/cost measured at about seven per")
	fmt.Fprintln(out, "cent of a session, because the work is done before that command is called.")
}

// printTemplateBudget is the other half of the argument WP3 names, and both halves have to be
// visible at once to be readable: every template is inside its budget of three to four phase
// specific required sections, and seventeen required sections still come out of the six.
func printTemplateBudget(out io.Writer, f *runner.Figures) {
	fmt.Fprintln(out, "\nTEMPLATES  WP3 budgets three to four phase-specific required sections each")
	fmt.Fprintln(out, sprintRow(templateBudgetRow,
		"PHASE", "TEMPLATE", "SOURCE", "REQUIRED", "SPECIFIC", "OPTIONAL"))
	required, over := 0, 0
	for _, t := range f.Templates {
		if t.Problem != "" {
			fmt.Fprintln(out, strings.TrimRight(sprintRow(templateBudgetRow,
				t.Phase, "?", "?", "?", "?", "?"), " ")+"  ("+t.Problem+")")
			continue
		}
		required += t.Required
		line := sprintRow(templateBudgetRow, t.Phase, t.Ref, string(t.Source),
			fmt.Sprint(t.Required), fmt.Sprint(t.Specific), fmt.Sprint(t.Optional))
		if t.Specific > budgetSpecific {
			over++
			line += fmt.Sprintf("  (over the budget of %d)", budgetSpecific)
		}
		fmt.Fprintln(out, strings.TrimRight(line, " "))
	}
	fmt.Fprintf(out, "%d required sections across the six, which is what one intent owes.\n", required)
	if over == 0 {
		fmt.Fprintln(out, "Every template is inside the budget, which is the argument: the budget is met")
		fmt.Fprintln(out, "per phase and the total is what a one-line fix pays.")
	}
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

// cmdScopeSet writes the intent's context scope and reports what its patterns resolve to.
//
// The entry is read whole rather than assembled from flags, which is cmdQuestionRecord's
// argument: include patterns, exclude patterns, a link per component and two budget numbers are
// a nested structure however they are spelled. There is no --phase, because the scope is P0's
// artifact and is read from P0 by every phase of the intent.
//
// The figure is printed because a budget set from a guess is what this intent's own P0 did, and
// it had to be corrected before the phase was judged. The sentence about that moment is here
// rather than in a comment: the scope lies inside P0's artifacts_hash, so a budget revised after
// the verdict makes the phase report as changed.
func cmdScopeSet(o *opts) int {
	entry, err := readMessage(o.file)
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 2
	}
	res, err := o.r.ScopeSet(o.key, []byte(entry))
	if code := o.report(nil, err); code != 0 {
		return code
	}
	fmt.Fprintf(o.out, "%s %s resolves %d files and %d bytes\n",
		o.key, model.ContextScope, res.Files, res.Bytes)
	fmt.Fprintln(o.out, "set the budget from these figures: it cannot be changed once "+
		"00-intake is judged, because the scope lies inside its artifacts_hash")
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
	started, err := o.r.StartPhase(o.key, o.phase)
	if code := o.report(nil, err); code != 0 {
		return code
	}
	// What the tracker had to say. The section written goes to stdout because it changed the
	// artifact; the reason nothing was written goes to stderr, because a phase that was
	// started with no issue content is working as Appendix A describes and the sentence is
	// something to know rather than something that failed. Not under --export, which prints
	// what a shell evals and must carry nothing else.
	if !o.export {
		if started.Section != "" {
			fmt.Fprintf(o.out, "%s carries the issue's content in the %s section\n",
				o.phase, started.Section)
		} else if started.Note != "" {
			fmt.Fprintln(o.errw, started.Note)
		}
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
	// Which lenses this phase works under, from project.yaml and the vendored plugin. Printed
	// for the same reason the changed set is: section 5's field list has no entry for it, and a
	// lens is a cost decision rather than part of the trail. Silent where none applies, which
	// is every project that enabled none — Appendix A's default. A name nobody answers to is
	// said out loud, because a misspelled one is otherwise indistinguishable from an empty list.
	applying, unknown := o.r.Lenses(o.phase)
	for _, l := range applying {
		fmt.Fprintf(o.out, "lens %s applies to %s: %s\n", l.ID, o.phase, l.Skill)
	}
	for _, name := range unknown {
		fmt.Fprintf(o.out, "lenses.enabled names %s, and the vendored plugin has no such lens\n", name)
	}
	// What a repeated phase has to read again, from the predecessor's lock and the tree. Printed
	// rather than recorded: section 5's field list has no entry for it, and the lock states what
	// was declared rather than what was read. Silent where nothing moved, which is also what a
	// repository with no context scope gets, since it declared no base to move.
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

// cmdLearningPropose carries what the records proposed to the rule set. It prints the figures
// and where the bundle is, and claims to have opened nothing, because it has not: what follows
// is a person reading the description and pushing the branch.
//
// A generated file the rules package will not have is a red exit. The route's claim is that a
// reviewer is handed something the rule set accepts, so a bundle that fails to load as a rule
// tree is a defect in the route rather than a finding about the records.
func cmdLearningPropose(o *opts) int {
	b, problems, err := o.r.ProposeLearning(o.key, o.dest)
	if code := o.report(nil, err); code != 0 {
		return code
	}
	fmt.Fprintf(o.out, "%d learning records read, %d stating no finding, %d proposals between the rest\n",
		b.Records, b.NoFinding, len(b.Items))
	fmt.Fprintf(o.out, "  carried     %4d into %d rule file(s)\n", b.Count(learning.Carried), len(b.Files))
	fmt.Fprintf(o.out, "  for wording %4d project-convention(s) naming a file the project reads\n",
		b.Count(learning.ForWording))
	fmt.Fprintf(o.out, "  unresolved  %4d naming no file this route can change\n", b.Count(learning.Unresolved))
	fmt.Fprintf(o.out, "%s holds the rule files, %s and %s; every proposal is listed in the description\n",
		o.dest, learning.PatchFile, learning.DescriptionFile)
	for _, p := range problems {
		fmt.Fprintf(o.errw, "the generated %s %s\n  next: %s\n", p.Path, p.Cause, p.Next)
	}
	if len(problems) > 0 {
		return 1
	}
	return 0
}

// cmdReviewAnswer answers one review rule of the effective set. The rule is the positional
// argument, as it is for gate approve and obligation close, and o.finding is where parse puts
// one.
//
// The report names what is still unanswered because the writer has resolved the set anyway,
// and a checklist finished by reading phase finish's findings is the loop this command exists
// to close.
func cmdReviewAnswer(o *opts) int {
	a, err := o.r.ReviewAnswer(o.key, model.ChecklistEntry{
		Rule: o.finding, Result: o.result, Note: o.note,
	})
	if code := o.report(nil, err); code != 0 {
		return code
	}
	wrote := "answers"
	if a.Replaced {
		wrote = "re-answers"
	}
	fmt.Fprintf(o.out, "%s %s %s %s: %s\n", o.key, model.Phases[len(model.Phases)-1], wrote,
		a.Entry.Rule, a.Entry.Result)
	if len(a.Unanswered) == 0 {
		fmt.Fprintln(o.out, "every review rule of the effective set is answered")
	} else {
		fmt.Fprintf(o.out, "still unanswered: %s\n", strings.Join(a.Unanswered, ", "))
	}
	return o.next(0)
}

// cmdReviewLens writes a lens's checklist entry. It is a command rather than a --source on
// review answer because the property section 12 rests on is that a lens entry answers no rule,
// and a flag can be given beside a rule id where an argument list with no place for one cannot.
//
// The stray positional is refused rather than ignored. parse takes a leading non-flag argument
// off the front for every command, so `xeno review lens migration-note --result met` would
// otherwise write an entry that does not answer the rule the person named, which is the mistake
// this command's shape exists to prevent.
//
// What it prints is the count rather than what is still unanswered: a lens owes none of the
// effective set, and naming what is outstanding would read as though it did.
func cmdReviewLens(o *opts) int {
	if o.finding != "" {
		fmt.Fprintf(o.errw, "review lens takes no rule and was given %q: a lens entry answers "+
			"none, and an entry for a rule of the effective set is written by xeno review answer\n",
			o.finding)
		return 2
	}
	n, err := o.r.ReviewLens(o.key, o.result, o.note)
	if code := o.report(nil, err); code != 0 {
		return code
	}
	fmt.Fprintf(o.out, "%s %s notes a lens entry: %s\n", o.key, model.Phases[len(model.Phases)-1],
		n.Entry.Result)
	fmt.Fprintf(o.out, "the checklist carries %d from lenses; none of them answers a rule, so "+
		"what G-Policy counts is unchanged\n", n.Lens)
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

// cmdEvidenceDeclare writes one declaration. The line it prints says which of section 4's
// three states the item is in, because that is what decides what happens next: a bound item
// is judged by G-Evidence now, and a pending one waits for a pipeline and for the attach.
func cmdEvidenceDeclare(o *opts) int {
	e, err := o.r.DeclareEvidence(o.key, o.phase, o.file, model.EvidenceItem{
		Kind: o.kind, Job: o.job, Result: o.result, URI: o.uri, SHA256: o.sha256,
		ProducedBy: o.producedBy, Format: o.format,
	})
	if code := o.report(nil, err); code != 0 {
		return code
	}
	where := "pending, until the pipeline publishes it"
	switch {
	case e.Path != "":
		where = "bound to " + e.Path + " by " + e.SHA256
	case e.URI != "":
		where = "bound to " + e.URI + " by " + e.SHA256
	}
	fmt.Fprintf(o.out, "%s %s declares %s: %s\n", o.key, o.phase, declared(e), where)
	return o.next(0)
}

// declared names the item the way the gate's findings do, kind and job, and leaves the job
// out where there is none rather than printing a trailing separator.
func declared(e *model.EvidenceItem) string {
	if e.Job == "" {
		return e.Kind
	}
	return e.Kind + "/" + e.Job
}

func cmdEvidenceAttach(o *opts) int {
	res, err := evidence.Attach(o.root, o.key, o.phase, o.src)
	if err != nil {
		fmt.Fprintln(o.errw, err)
		return 2
	}
	fmt.Fprintf(o.out, "attached %d, pending %d; run xeno gate run to carry the verdict forward\n",
		res.Attached, res.Pending)
	// Named on stderr and counted as pending: an entry that was declined was not attached,
	// and a run that only printed the counts would report it as evidence still to come
	// from a job that has already produced it.
	//
	// No line of advice is added after them. Each sentence says what to republish, because
	// the reasons are fixed in different places, and a single hint underneath was wrong for
	// two of the three the moment there were three (#221).
	for _, u := range res.Declined {
		fmt.Fprintln(o.errw, "  not attached:", u)
	}
	if len(res.Declined) > 0 {
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
		// A redo that nobody finished leaves nothing in the repository, because staged
		// content is not committed, so this line is the only place it is visible at all.
		if len(s.Staged) > 0 {
			line += fmt.Sprintf("  (%s staged, run xeno phase finish to apply)",
				strings.Join(s.Staged, ", "))
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

// templateRow is one section of a template: the id `section set` takes, whether the phase owes
// it, and the heading it renders under. The id is first because it is the cell somebody types.
const templateRow = "%-22s %-8s %s"

// symbolRow is one definition: the name, what the indexer called it, the file and line, and
// what encloses it. The container is last and often empty, which is a value and not a gap: a
// symbol at the top level has nothing enclosing it.
const symbolRow = "%-24s %-10s %-44s %s"

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
	intent, phase := cost.LivePhase(model.LocalPath(o.root, runner.PhaseEnvFile))
	if phase == "" {
		phase = cost.NoPhase
	}
	_ = cost.Append(o.root, cost.Turn{
		At: time.Now().UTC().Format(time.RFC3339), Session: hook.Session,
		Intent: intent, Phase: phase, Totals: totals,
	})
	return 0
}

// The two commands behind fetch_template and query_symbol_index, which WP11 asks for: every
// MCP operation has a command behind it, because the server is one way in and never the only
// one. An operation with no command is something an agent driven through MCP can do that a
// person at a terminal cannot, and section 13's arrangement is MCP instead of free text and
// never only MCP. Both call the runner method the server's tool calls, so there is one path
// to the answer and no second one to reconcile; what differs is the rendering, columns for a
// person here and JSON for a model there.

// cmdTemplateShow prints the sections a phase owes, in the order they are rendered in.
//
// The template ref and where it was resolved from are on the first line because they are what
// makes the list falsifiable: a project that replaced one template shows `project` here and the
// same sections everywhere else, and a list with nothing saying which template it came from
// cannot be held against the file that was rendered.
func cmdTemplateShow(o *opts) int {
	t, err := o.r.Template(o.phase)
	if code := o.report(nil, err); code != 0 {
		return code
	}
	fmt.Fprintf(o.out, "%s renders from %s (%s) in %s\n", o.phase, t.Ref(), t.Source, t.Bundle.Language)
	if t.Bundle.Title != "" {
		fmt.Fprintln(o.out, t.Bundle.Title)
	}
	fmt.Fprintln(o.out, sprintRow(templateRow, "SECTION", "REQUIRED", "HEADING"))
	for _, s := range t.Template.Sections {
		required := "no"
		if s.Required {
			required = "yes"
		}
		fmt.Fprintln(o.out, strings.TrimRight(
			sprintRow(templateRow, s.ID, required, t.Bundle.Headings[s.ID]), " "))
	}
	return 0
}

// cmdSymbolShow prints where a name is defined, from the index the project produced.
//
// It exits zero whether anything was found and whether there was an index to look in at all,
// for Runner.Symbols' reason: absent, unreadable, malformed and stale are four causes with one
// outcome, no gate reads the index, and a repository that has not produced one is the state
// every repository starts in. Exit 1 is a refusal with a reason, and the ordinary state of the
// thing is not one.
//
// What it does instead is say why there was nothing, because no index at all and a name the
// index does not hold would otherwise print the same silence, and only one of the two is a
// reason to go and look at the configuration.
//
// The provenance is printed beside the locations rather than asked for. An index is allowed to
// be stale or wrong, so a location with nothing saying what produced it and how old that is
// cannot be weighed against the source the reader also has.
func cmdSymbolShow(o *opts) int {
	if o.finding == "" {
		fmt.Fprintln(o.errw, "a name is required: xeno symbol show NAME")
		return 2
	}
	a := o.r.Symbols(o.finding)
	if a.Why != "" {
		fmt.Fprintf(o.out, "no index was read: %s\n", a.Why)
		return 0
	}
	fmt.Fprintf(o.out, "%s %s, %s old\n", a.Tool, a.ToolVersion, a.Age.Round(time.Minute))
	if len(a.Found) == 0 {
		fmt.Fprintf(o.out, "%s is not in it, and an index is allowed to be incomplete\n", a.Name)
		return 0
	}
	fmt.Fprintln(o.out, sprintRow(symbolRow, "NAME", "KIND", "WHERE", "IN"))
	for _, s := range a.Found {
		fmt.Fprintln(o.out, strings.TrimRight(sprintRow(symbolRow, s.Name, s.Kind,
			fmt.Sprintf("%s:%d", s.File, s.Line), s.Container), " "))
	}
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
		// What the gate did not look at, where it has such a bound to state. Section 5
		// enumerates the fields of a check, so this is not something gate.yaml can carry, and
		// the run that prints the result is where a reader meets it. A pass says that what was
		// read was clean and nothing about what was read, and for G-Secret that difference is
		// the half of the check the digest filter does instead (#284).
		if note := gates.Coverage(c.Gate); note != "" {
			fmt.Fprintf(o.out, "      covers %s\n", note)
		}
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
func printEnforcement(out io.Writer, root string, rep *enforcement.Report) {
	fmt.Fprintf(out, "%s, branch %s\n", rep.Repository, rep.Branch)
	for _, q := range rep.Requirements {
		fmt.Fprintf(out, "  %-13s %-20s declared %-6s actual %s\n", q.State, q.Name, q.Declared, q.Actual)
		if q.Note != "" {
			fmt.Fprintf(out, "                  %s\n", q.Note)
		}
	}
	fmt.Fprintf(out, "\nreport written to %s\n", runner.ReportPath(root))
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
