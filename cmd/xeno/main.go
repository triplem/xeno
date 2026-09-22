// SPDX-License-Identifier: Apache-2.0

// Command xeno is the runner. Exit codes follow one staircase throughout: 0 where the
// command did what was asked and the verdict is not red, 1 on a red verdict or a
// refusal with a reason, 2 where it could not run at all.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/triplem/xeno/internal/evidence"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/runner"
)

const usage = `usage:
  xeno phase start    --intent KEY --phase NN [--evidence-from DIR]
  xeno phase finish   --intent KEY --phase NN
  xeno gate run       --intent KEY --phase NN [--evidence-from DIR]
  xeno gate verify    [--intent KEY]            recompute and compare, write nothing (CI)
  xeno evidence attach --intent KEY --phase NN --from DIR
  xeno intent status  --intent KEY
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
	fs := flag.NewFlagSet(cmd, flag.ContinueOnError)
	root := fs.String("root", ".", "repository root")
	key := fs.String("intent", "", "intent key")
	phaseArg := fs.String("phase", "", "phase id or number")
	from := fs.String("evidence-from", "", "directory standing in for the pipeline artifact store")
	src := fs.String("from", "", "directory standing in for the pipeline artifact store")
	if err := fs.Parse(args[2:]); err != nil {
		return 2
	}
	if *key == "" && cmd != "gate verify" {
		fmt.Fprintln(os.Stderr, "--intent is required")
		return 2
	}
	r := runner.New(*root)
	r.EvidenceFrom = *from

	phase := ""
	if cmd != "intent status" && cmd != "gate verify" {
		p, err := model.ResolvePhase(*phaseArg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		phase = p
	}

	switch cmd {
	case "phase start":
		return report(nil, r.Start(*key, phase))
	case "phase finish":
		return report(r.Finish(*key, phase))
	case "gate run":
		return report(r.GateRun(*key, phase))
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
