// SPDX-License-Identifier: Apache-2.0

// Command xeno is the runner. Exit codes follow one staircase throughout: 0 where the
// command did what was asked and the verdict is not red, 1 on a red verdict or a
// refusal with a reason, 2 where it could not run at all.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/triplem/xeno/internal/enforcement"
	"github.com/triplem/xeno/internal/evidence"
	"github.com/triplem/xeno/internal/gates"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/runner"
)

const usage = `usage:
  xeno init           [--vendor] [--project OWNER/REPO] [--model ID] [--language TAG]
  xeno phase start    --intent KEY --phase NN [--evidence-from DIR]
  xeno phase finish   --intent KEY --phase NN
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
  xeno intent status  --intent KEY
  xeno intent close   --intent KEY --reason TEXT
  xeno section set    SECTION --intent KEY --phase NN [--file PATH]   reads stdin without --file
  xeno check commit-message [--pattern NAME] [--file PATH]   reads stdin without --file
  xeno version
common: --root DIR (default .), --no-next to leave out the next step`

func main() { os.Exit(run(os.Args[1:])) }

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
	cmd, rest0 := args[0]+" "+args[1], args[2:]
	if args[0] == "init" {
		cmd, rest0 = "init", args[1:]
	}

	// A finding id stands before the flags, as the process definition writes these
	// commands. Go's flag package stops at the first argument that is not a flag, so it
	// is taken off the front rather than read back out afterwards.
	finding := ""
	rest := rest0
	if len(rest) > 0 && !strings.HasPrefix(rest[0], "-") {
		finding, rest = rest[0], rest[1:]
	}

	fs := flag.NewFlagSet(cmd, flag.ContinueOnError)
	root := fs.String("root", ".", "repository root")
	key := fs.String("intent", "", "intent key")
	phaseArg := fs.String("phase", "", "phase id or number")
	from := fs.String("evidence-from", "", "directory standing in for the pipeline artifact store")
	src := fs.String("from", "", "directory standing in for the pipeline artifact store")
	by := fs.String("by", "", "the person deciding")
	pattern := fs.String("pattern", "conventional-commits", "a shipped pattern name")
	file := fs.String("file", "", "the message to read, or stdin when absent")
	vendor := fs.Bool("vendor", false, "copy the plugin in and pin it")
	project := fs.String("project", "", "the tracker project the intents belong to")
	mdl := fs.String("model", "", "the default model a phase uses")
	language := fs.String("language", "en", "the language artifacts are written in")
	pluginFrom := fs.String("plugin-from", ".xeno/plugin", "where to vendor the plugin from")
	host := fs.String("host", "github", "which CI wrapper to generate")
	branch := fs.String("branch", "", "the branch whose protection to read, main by default")
	// Both ends of the commit range. They are an input of the run and are deliberately
	// not recorded: after a squash a recorded range would point at commits that no
	// longer exist, and a field that is sometimes wrong is worse than no field. No gate
	// reads them until WP4 brings the commit predicates.
	base := fs.String("base", "", "the base of the commit range under review")
	head := fs.String("head", "", "the head of the commit range under review")
	reason := fs.String("reason", "", "why")
	text := fs.String("text", "", "the statement being assumed")
	origin := fs.String("origin", "", "where the assumption came from")
	confidence := fs.String("confidence", "", "how much weight it carries")
	resolves := fs.String("resolves", "", "the open question this assumption answers")
	// The suggestion is off by default nowhere and on by default nowhere either: the
	// commands that change state say it, the ones a pipeline or a hook runs do not, and
	// this turns it off for the scripts that are neither.
	noNext := fs.Bool("no-next", false, "do not say what the next step is")
	if err := fs.Parse(rest); err != nil {
		return 2
	}
	if *key == "" && cmd != "gate verify" && cmd != "check commit-message" && cmd != "init" &&
		cmd != "enforcement check" {
		fmt.Fprintln(os.Stderr, "--intent is required")
		return 2
	}
	r := runner.New(*root)
	r.EvidenceFrom = *from

	phase := ""
	switch cmd {
	case "intent status", "intent close", "gate verify", "check commit-message", "init",
		"enforcement check", "assumption confirm", "assumption reject":
	default:
		p, err := model.ResolvePhase(*phaseArg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		phase = p
	}

	switch cmd {
	case "init":
		r.PluginSource = *pluginFrom
		res, err := r.Init(runner.InitOptions{
			TrackerKey: *project, Model: *mdl, Language: *language, Vendor: *vendor, Host: *host,
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		printInit(res)
		return 0
	case "enforcement check":
		rep, err := r.EnforcementCheck(*branch, enforcement.Token())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		printEnforcement(rep)
		if rep.Unmet() > 0 {
			return 1
		}
		return 0
	case "check commit-message":
		message, err := readMessage(*file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		if err := gates.CheckMessage(*pattern, message); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	case "section set":
		content, err := readMessage(*file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		t, err := r.SectionSet(*key, phase, finding, content)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		fmt.Printf("%s rendered from %s (%s)\n", phase, t.Ref(), t.Source)
		return suggest(r, *key, *noNext, 0)
	case "intent close":
		return suggest(r, *key, *noNext, report(r.IntentClose(*key, *reason)))
	case "phase start":
		return suggest(r, *key, *noNext, report(nil, r.Start(*key, phase)))
	case "phase finish":
		return suggest(r, *key, *noNext, report(r.Finish(*key, phase)))
	case "gate run":
		r.Base, r.Head = *base, *head
		return suggest(r, *key, *noNext, report(r.GateRun(*key, phase)))
	case "gate approve":
		return suggest(r, *key, *noNext,
			report(r.Decide(*key, phase, finding, "approved", *by, *reason)))
	case "gate override":
		return suggest(r, *key, *noNext,
			report(r.Decide(*key, phase, finding, "overridden", *by, *reason)))
	case "assumption record":
		a, err := r.RecordAssumption(*key, phase, *text, *origin, *confidence, *resolves)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		fmt.Printf("%s recorded, open, from %s with %s confidence\n", a.ID, a.Origin, a.Confidence)
		return suggest(r, *key, *noNext, 0)
	case "assumption confirm", "assumption reject":
		status := "confirmed"
		if cmd == "assumption reject" {
			status = "rejected"
		}
		a, err := r.DecideAssumption(*key, finding, status, *by)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		fmt.Printf("%s %s by %s\n", a.ID, a.Status, *by)
		return suggest(r, *key, *noNext, 0)
	case "obligation close":
		return suggest(r, *key, *noNext, report(r.CloseObligation(*key, phase, finding)))
	case "evidence attach":
		a, p, err := evidence.Attach(*root, *key, phase, *src)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		fmt.Printf("attached %d, pending %d; run xeno gate run to carry the verdict forward\n", a, p)
		return suggest(r, *key, *noNext, 0)
	case "gate verify":
		res, err := r.Verify(*key)
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
	case "intent status":
		states, err := r.Status(*key)
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
		return suggest(r, *key, *noNext, 0)
	}
	fmt.Fprintln(os.Stderr, usage)
	return 2
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

// printInit says what was done, what was left alone, and what a person still has to do.
// The last list is the point: a first contact that leaves the project believing the gate
// is binding when it is not is worse than no first contact.
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

// suggest prints the next step of the working sequence and passes the exit code through.
// It runs after the command, so the state it reads is the state the command left behind
// and not the one it found. An internal error says nothing: a suggestion derived from a
// state the runner could not read would be a guess.
//
// A refusal does get one, and it is worth the most there: a phase start refused because
// its predecessor is red is exactly the moment somebody wants to be told what to do.
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
