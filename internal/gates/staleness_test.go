// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
)

// Section 5 states two limits on the staleness half and gives each its reason. Both were
// missing (#236), and XENO-0245 found out by being blocked: its own specification commit moved
// two documents its scope named, the check read the gated phase's own lock, and the phase went
// red for doing its job.

// stalenessFixture writes a lock per phase, each recording the files named with their real
// hashes, in a repository with a commit range. Changing a file afterwards is then the only
// difference between what a lock says it was given and what is there.
func stalenessFixture(t *testing.T, perPhase map[string][]string, bodies map[string]string) (Ctx, string) {
	t.Helper()
	root := t.TempDir()
	gitRepo(t, root)
	for rel, body := range bodies {
		p := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	base := gitCommit(t, root, "the tree the phases were given")
	for phase, files := range perPhase {
		dir := filepath.Join(root, model.PhaseDir("PROJ-1", phase))
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
		lock := "intent: " + fixtureIntent + "\nphase: " + phase + "\nfiles:\n"
		for _, rel := range files {
			h, err := hashing.FileHash(filepath.Join(root, rel))
			if err != nil {
				t.Fatal(err)
			}
			lock += "  - path: " + rel + "\n    sha256: " + h + "\n"
		}
		if err := os.WriteFile(filepath.Join(dir, "context.lock.yaml"), []byte(lock), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return Ctx{Root: root, Key: "PROJ-1", QualifiedID: fixtureIntent, Base: base}, root
}

// "A phase that changes the files it read is not stale, it is working: that is what P3 does,
// and without this limit every implementation phase would report itself out of date the moment
// it did its job."
func TestTheLockOfThePhaseBeingGatedIsNotRead(t *testing.T) {
	c, root := stalenessFixture(t,
		map[string][]string{model.Phases[3]: {"internal/gates/gates.go"}},
		map[string]string{"internal/gates/gates.go": "package gates\n"})

	if err := os.WriteFile(filepath.Join(root, "internal/gates/gates.go"),
		[]byte("package gates // the phase did its job\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	c.Phase = model.Phases[3]
	c.Head = gitCommit(t, root, "the phase changes what it read")

	if fs := staleReads(c, 3); fs != nil {
		t.Fatalf("the gated phase's own lock was read: %v", findingText(fs))
	}
}

// The other direction, which is what the check is for: a phase whose ground moved after it
// finished. P1's lock is read when P3 is gated, and the finding names the phase and the file.
func TestAPrecedingPhasesMovedInputIsAFinding(t *testing.T) {
	c, root := stalenessFixture(t,
		map[string][]string{model.Phases[1]: {"internal/gates/gates.go"}},
		map[string]string{"internal/gates/gates.go": "package gates\n"})

	if err := os.WriteFile(filepath.Join(root, "internal/gates/gates.go"),
		[]byte("package gates // moved under P1\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	c.Phase = model.Phases[3]
	c.Head = gitCommit(t, root, "the ground moves under an earlier phase")

	fs := staleReads(c, 3)
	if len(fs) != 1 {
		t.Fatalf("a preceding phase's moved input gave %d findings: %v", len(fs), findingText(fs))
	}
	if !strings.Contains(findingText(fs), model.Phases[1]) ||
		!strings.Contains(findingText(fs), "internal/gates/gates.go") {
		t.Fatalf("the finding names neither the phase nor the file: %s", findingText(fs))
	}
}

// "It looks at the files the change under review touched, not at everything the scope names…
// in an active repository every rebase onto a moved default branch would set it off, and a
// check that fires constantly is one people learn to ignore."
func TestOnlyWhatTheChangeUnderReviewTouchedIsReported(t *testing.T) {
	c, root := stalenessFixture(t,
		map[string][]string{model.Phases[1]: {"internal/gates/gates.go", "docs/process-definition.md"}},
		map[string]string{
			"internal/gates/gates.go":    "package gates\n",
			"docs/process-definition.md": "# the process\n",
		})

	// The document moves first and is committed, as a rebase onto a moved default branch
	// would move it: it no longer matches the hash P1's lock recorded, and it is behind the
	// base of the range under review.
	if err := os.WriteFile(filepath.Join(root, "docs/process-definition.md"),
		[]byte("# the process, moved elsewhere\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	c.Base = gitCommit(t, root, "the document moves outside the change")

	// The source file moves inside the range, which is the one the check is asked about.
	if err := os.WriteFile(filepath.Join(root, "internal/gates/gates.go"),
		[]byte("package gates // in the change\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	c.Phase = model.Phases[3]
	c.Head = gitCommit(t, root, "the change under review")

	fs := staleReads(c, 3)
	if len(fs) != 1 {
		t.Fatalf("want one finding for the touched file, got %d: %v", len(fs), findingText(fs))
	}
	if !strings.Contains(findingText(fs), "internal/gates/gates.go") {
		t.Fatalf("the finding is about the untouched file: %s", findingText(fs))
	}
}

// Section 12 forbids working the range out here, so without one there is no change under
// review to narrow to and the half reports nothing. That is a silent pass of the kind #217 was
// opened about and it cannot be reported as anything else until #235; the test states the
// behaviour so that a reader finds it deliberate rather than missing.
func TestWithNoCommitRangeTheStalenessHalfReportsNothing(t *testing.T) {
	c, root := stalenessFixture(t,
		map[string][]string{model.Phases[1]: {"internal/gates/gates.go"}},
		map[string]string{"internal/gates/gates.go": "package gates\n"})

	if err := os.WriteFile(filepath.Join(root, "internal/gates/gates.go"),
		[]byte("package gates // moved\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	c.Phase = model.Phases[3]
	c.Base, c.Head = "", ""

	if fs := staleReads(c, 3); fs != nil {
		t.Fatalf("a run with no range reported %v", findingText(fs))
	}
}

// findingText is the findings as text, as causes is for a whole check: these tests call
// staleReads directly, because the two limits are its and a finding routed through the gate
// would answer for freshness's first half as well.
func findingText(fs []model.Finding) string {
	var b strings.Builder
	for _, f := range fs {
		b.WriteString(f.File + ": " + f.Cause + "\n")
	}
	return b.String()
}
