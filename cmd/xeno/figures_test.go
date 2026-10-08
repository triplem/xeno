// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
)

// ---- WP20: the proportionality figures, read off the trail rather than counted by hand

// twoIntents is a trail of exactly the two cases the figures have to keep apart: PROJ-1, which
// repo() leaves at the intake with no verdict anywhere, and PROJ-2, carried through all six
// phases. One of each is the whole point of the fixture: an intent that stopped at the intake
// is a different measurement from one that went through, and a report that averaged them
// together would describe a process nobody ran.
//
// PROJ-2's phases are written as files rather than driven through the commands, because what
// is under test is a reader. Driving six phases to a green verdict would make the fixture a
// test of the gates, and the gates have their own.
func twoIntents(t *testing.T) string {
	t.Helper()
	root := repo(t)
	writeUnder(t, root, model.IntentDir("PROJ-2")+"/intent.yaml",
		"intent: \"git.example/group/proj#2\"\nkey: PROJ-2\nstatus: in-progress\n"+
			"created: \"2026-09-21T10:00:00Z\"\n")

	for i, phase := range model.Phases {
		dir := model.PhaseDir("PROJ-2", phase)
		// Two sections with content and one with none. The empty one is the case worth having:
		// an absent section and a present empty one are the same silence, so the count is of
		// what was written and not of what was anchored.
		writeUnder(t, root, dir+"/output.md",
			"---\nintent: \"git.example/group/proj#2\"\nphase: "+phase+"\n---\n\n"+
				"# A phase\n\n<!-- xeno:section:problem -->\n## Problem\n\n"+
				"the first section of this phase\n\n"+
				"<!-- xeno:section:scope -->\n## Scope\n\nthe second section of this phase\n\n"+
				"<!-- xeno:section:open-questions -->\n## Open questions\n")
		writeUnder(t, root, dir+"/digest.md", "a digest of "+phase+"\n")
		writeUnder(t, root, dir+"/learning.yaml", "no_finding: true\n")
		// One phase attributed a cost and the other five did not, which is the state the whole
		// coverage rule exists for: the report may print what was attributed and may not turn
		// it into a cost per intent.
		if i == 0 {
			writeUnder(t, root, dir+"/cost.yaml", "phase: "+phase+"\nevidence: self-reported\n"+
				"tokens_in: 11\ntokens_out: 22\ntokens_cached: 33\n")
		}
		// The verdict last, and with the hash the directory actually has: a gate.yaml carrying
		// any other value would make Status report the phase as changed after its verdict, and
		// the fixture would then be measuring drift instead of a finished phase.
		h, err := hashing.DirHash(root, dir, hashing.PhaseExcluded)
		if err != nil {
			t.Fatal(err)
		}
		verdict := "status: green\n"
		// One released finding in one phase, so that the reopen figure has something to count.
		// An approval is what a person writes where a phase could not pass on its own.
		if i == len(model.Phases)-1 {
			verdict += "checks:\n  - gate: G-Schema\n    result: pass\n    provenance: internal\n" +
				"    findings:\n      - id: F-abc123\n        file: output.md\n" +
				"        cause: a section was thin\n        next: write it again\n" +
				"        decision:\n          type: approved\n          by: a.person\n" +
				"          at: \"2026-09-21T11:00:00Z\"\n          against: " + h + "\n" +
				"          reason: assessed and accepted\n"
		}
		writeUnder(t, root, dir+"/gate.yaml",
			"intent: \"git.example/group/proj#2\"\nphase: "+phase+"\n"+
				"run_at: \"2026-09-21T10:30:00Z\"\nartifacts_hash: "+h+"\n"+verdict)
	}
	return root
}

// TestTheFiguresSayWhatPopulationTheyCounted is the acceptance of #316's honesty clause. A
// count over the trail includes the intent doing the counting, and an intent that stopped at
// the intake is not the same measurement as one that went through, so the report states both
// numbers before it states anything derived from them.
func TestTheFiguresSayWhatPopulationTheyCounted(t *testing.T) {
	root := twoIntents(t)
	code, out, errw := invoke(t, "report", "figures", "--root", root, "--all")
	if code != 0 {
		t.Fatalf("report figures exited %d:\n%s%s", code, out, errw)
	}
	want := "2 intents counted, 1 with all six phases, 1 that stopped earlier."
	if !strings.Contains(out, want) {
		t.Errorf("the report does not say what it counted.\nwant a line: %s\ngot:\n%s", want, out)
	}
	if !strings.Contains(out, "PER INTENT, over the 1 with all six phases") {
		t.Errorf("the aggregate does not name its population:\n%s", out)
	}
	// The intent that stopped at the intake is a row and not an omission: a record missing
	// from a listing is worse than one that looks wrong in it, which is the listing's rule.
	if !strings.Contains(out, "PROJ-1") || !strings.Contains(out, "PROJ-2") {
		t.Errorf("one of the two intents is missing from the rows:\n%s", out)
	}
	for _, row := range strings.Split(out, "\n") {
		if strings.Contains(row, "PROJ-1") && !strings.Contains(row, "0/6") {
			t.Errorf("PROJ-1 reached no verdict and the row does not say so: %q", row)
		}
		if strings.Contains(row, "PROJ-2") && !strings.Contains(row, "6/6") {
			t.Errorf("PROJ-2 has all six phases and the row does not say so: %q", row)
		}
	}
}

// TestTheThreeFiguresAreCountedOffTheTrail holds the artifacts and reopens figures to the
// files the fixture wrote. Twelve sections across six phases, two per phase, with the empty
// third counted nowhere; one released finding, which is the proxy the reopen figure is.
func TestTheThreeFiguresAreCountedOffTheTrail(t *testing.T) {
	root := twoIntents(t)
	code, out, errw := invoke(t, "report", "figures", "--root", root, "--intent", "PROJ-2")
	if code != 0 {
		t.Fatalf("report figures exited %d:\n%s%s", code, out, errw)
	}
	for _, want := range []string{
		"1 intents counted, 1 with all six phases, 0 that stopped earlier.",
		"sections written       12.0",
		"files                  26.0",
		"reopens                1.00",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("the figures do not carry %q:\n%s", want, out)
		}
	}
}

// TestTheCostFigureTravelsWithItsCoverage is the clause #316 is most explicit about: cost.yaml
// exists for a handful of phases out of hundreds, and a mean over the ones that have it,
// printed as the cost of an intent, would be wrong by orders of magnitude. So the coverage is
// printed, the attributed totals are printed, and no per-intent figure is.
func TestTheCostFigureTravelsWithItsCoverage(t *testing.T) {
	root := twoIntents(t)
	code, out, errw := invoke(t, "report", "figures", "--root", root, "--intent", "PROJ-2")
	if code != 0 {
		t.Fatalf("report figures exited %d:\n%s%s", code, out, errw)
	}
	for _, want := range []string{
		"self-reported",
		"1 of 6 phases attributed anything.",
		"tokens_in 11  tokens_out 22  tokens_cached 33, over those phases.",
		"No cost per intent is printed",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("the cost block does not carry %q:\n%s", want, out)
		}
	}
	if strings.Contains(out, "per intent: tokens_in") {
		t.Errorf("a cost per intent was printed over a coverage of one phase in six:\n%s", out)
	}
}

// TestTheSectionBudgetIsReportedPerTemplate is the other half of WP3's argument, which only
// reads as an argument when both halves are in one place: every template is inside the budget
// of three to four phase-specific required sections, and seventeen required sections still
// come out of the six.
func TestTheSectionBudgetIsReportedPerTemplate(t *testing.T) {
	root := twoIntents(t)
	code, out, errw := invoke(t, "report", "figures", "--root", root, "--intent", "PROJ-2")
	if code != 0 {
		t.Fatalf("report figures exited %d:\n%s%s", code, out, errw)
	}
	for _, want := range []string{
		"intake@1.0.0",
		"17 required sections across the six",
		"Every template is inside the budget",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("the template budget does not carry %q:\n%s", want, out)
		}
	}
}

// TestTheFiguresWriteNothing is the property the whole report rests on, and it is asserted
// over the trail rather than argued: a measurement a verdict could turn on stops being a
// measurement, because the cheapest way to move it is then to change what is counted. The
// sealed walk holds every command to leaving a judged phase alone; this holds this one to
// leaving the whole of .xeno alone, which is the half that is particular to it.
func TestTheFiguresWriteNothing(t *testing.T) {
	root := twoIntents(t)
	before := snapshot(t, root)
	code, out, errw := invoke(t, "report", "figures", "--root", root, "--all")
	if code != 0 {
		t.Fatalf("report figures exited %d:\n%s%s", code, out, errw)
	}
	if after := snapshot(t, root); after != before {
		t.Errorf("report figures changed something under .xeno:\n%s", diffLines(before, after))
	}
}

// snapshot is every file under .xeno with the hash of its content, as one string. Content and
// not modification times: what matters is whether anything was written, and a tool that
// rewrote a file with the same bytes would be as wrong as one that changed them.
func snapshot(t *testing.T, root string) string {
	t.Helper()
	var b strings.Builder
	base := filepath.Join(root, ".xeno")
	err := filepath.WalkDir(base, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		sum := sha256.Sum256(content)
		rel, err := filepath.Rel(base, path)
		if err != nil {
			return err
		}
		fmt.Fprintf(&b, "%s %s\n", hex.EncodeToString(sum[:]), filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return b.String()
}

// diffLines names the lines the two snapshots do not share, so that a failure says which file
// moved rather than printing two walls of hashes.
func diffLines(before, after string) string {
	had := map[string]bool{}
	for _, l := range strings.Split(before, "\n") {
		had[l] = true
	}
	var out []string
	for _, l := range strings.Split(after, "\n") {
		if l != "" && !had[l] {
			out = append(out, l)
		}
	}
	return strings.Join(out, "\n")
}
