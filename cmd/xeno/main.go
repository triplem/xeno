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

	"github.com/triplem/xeno/internal/evidence"
	"github.com/triplem/xeno/internal/gates"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/runner"
)

const usage = `usage:
  xeno phase start    --intent KEY --phase NN [--evidence-from DIR]
  xeno phase finish   --intent KEY --phase NN
  xeno gate run       --intent KEY --phase NN [--evidence-from DIR]
  xeno gate approve   FINDING --intent KEY --phase NN --by WHO --reason TEXT
  xeno gate override  FINDING --intent KEY --phase NN --by WHO --reason TEXT
  xeno obligation close FINDING --intent KEY --phase NN
  xeno gate verify    [--intent KEY]            recompute and compare, write nothing (CI)
  xeno evidence attach --intent KEY --phase NN --from DIR
  xeno intent status  --intent KEY
  xeno intent close   --intent KEY --reason TEXT
  xeno check commit-message [--pattern NAME] [--file PATH]   reads stdin without --file
  xeno version
common: --root DIR (default .)`

func main() { os.Exit(run(os.Args[1:])) }

func run(args []string) int {
	if len(args) >= 1 && args[0] == "version" {
		fmt.Println("xeno", model.RunnerVersion)
		return 0
	}
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		return 2
	}
	cmd := args[0] + " " + args[1]

	// A finding id stands before the flags, as the process definition writes these
	// commands. Go's flag package stops at the first argument that is not a flag, so it
	// is taken off the front rather than read back out afterwards.
	finding := ""
	rest := args[2:]
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
	reason := fs.String("reason", "", "why")
	if err := fs.Parse(rest); err != nil {
		return 2
	}
	if *key == "" && cmd != "gate verify" && cmd != "check commit-message" {
		fmt.Fprintln(os.Stderr, "--intent is required")
		return 2
	}
	r := runner.New(*root)
	r.EvidenceFrom = *from

	phase := ""
	switch cmd {
	case "intent status", "intent close", "gate verify", "check commit-message":
	default:
		p, err := model.ResolvePhase(*phaseArg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		phase = p
	}

	switch cmd {
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
	case "intent close":
		return report(r.IntentClose(*key, *reason))
	case "phase start":
		return report(nil, r.Start(*key, phase))
	case "phase finish":
		return report(r.Finish(*key, phase))
	case "gate run":
		return report(r.GateRun(*key, phase))
	case "gate approve":
		return report(r.Decide(*key, phase, finding, "approved", *by, *reason))
	case "gate override":
		return report(r.Decide(*key, phase, finding, "overridden", *by, *reason))
	case "obligation close":
		return report(r.CloseObligation(*key, phase, finding))
	case "evidence attach":
		a, p, err := evidence.Attach(*root, *key, phase, *src)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		fmt.Printf("attached %d, pending %d; run xeno gate run to carry the verdict forward\n", a, p)
		return 0
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
		return 0
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
