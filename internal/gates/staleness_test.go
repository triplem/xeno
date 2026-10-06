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

// Section 7: "A stale phase is not deleted, it is re-run or explicitly approved as still
// valid." The remedy named neither and named the one act Start refuses, so a person following
// it was refused, then offered section set and phase finish, which re-render the artifact and
// leave the finding where it was: the lock is written only by phase start (#237).
//
// These assert what the remedy names rather than its wording. A test matching the literal
// would break on every rewording and pass on a route silently dropped, which is the wrong way
// round for a string whose fault was that nothing read it.
func TestAStalePhasesRemedyNamesBothRoutesSectionSevenNames(t *testing.T) {
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
		t.Fatalf("want one finding, got %d: %v", len(fs), findingText(fs))
	}
	assertSealedLockRemedy(t, fs[0].Next, model.Phases[1], model.Phases[3])
}

// The same two routes for a file that has gone, which is the other shape of the ground moving
// under a sealed phase. The old remedy for this one offered "record why the file left", which
// is not an act either: nothing in the runner records that against a lock.
func TestAVanishedInputGetsTheSameTwoRoutes(t *testing.T) {
	c, root := stalenessFixture(t,
		map[string][]string{model.Phases[1]: {"internal/gates/gates.go"}},
		map[string]string{"internal/gates/gates.go": "package gates\n"})

	if err := os.Remove(filepath.Join(root, "internal/gates/gates.go")); err != nil {
		t.Fatal(err)
	}
	c.Phase = model.Phases[3]
	c.Head = gitCommit(t, root, "the file leaves")

	fs := staleReads(c, 3)
	if len(fs) != 1 {
		t.Fatalf("want one finding, got %d: %v", len(fs), findingText(fs))
	}
	if !strings.Contains(fs[0].Cause, "is gone") {
		t.Fatalf("the finding is not the vanished one: %s", fs[0].Cause)
	}
	assertSealedLockRemedy(t, fs[0].Next, model.Phases[1], model.Phases[3])
}

// The third case is not like the other two and its remedy already worked. An unreadable file is
// a condition of this machine — a permission, a broken link, a filesystem — so the next gate run
// clears the finding with nobody approving anything and nothing started over. The test exists so
// that a later pass does not make all three alike for symmetry.
func TestAnUnreadableInputIsTheTreesFaultAndNeedsNoRelease(t *testing.T) {
	c, root := stalenessFixture(t,
		map[string][]string{model.Phases[1]: {"internal/gates/gates.go"}},
		map[string]string{"internal/gates/gates.go": "package gates\n"})

	p := filepath.Join(root, "internal/gates/gates.go")
	if err := os.WriteFile(p, []byte("package gates // in the change\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	c.Phase = model.Phases[3]
	c.Head = gitCommit(t, root, "the change under review")
	// After the commit, so that the range still names the path: what is unreadable is the
	// working tree's copy, which is what the check hashes.
	if err := os.Chmod(p, 0o000); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(p, 0o644) })

	fs := staleReads(c, 3)
	if len(fs) != 1 {
		t.Fatalf("want one finding, got %d: %v", len(fs), findingText(fs))
	}
	if !strings.Contains(fs[0].Cause, "cannot be read") {
		t.Fatalf("the finding is not the unreadable one: %s", fs[0].Cause)
	}
	if !strings.Contains(fs[0].Next, "make it readable") {
		t.Fatalf("the remedy does not name the act that works: %s", fs[0].Next)
	}
	for _, unwanted := range []string{"approves", "start " + model.Phases[1] + " over"} {
		if strings.Contains(fs[0].Next, unwanted) {
			t.Fatalf("the remedy asks for %q, which this fault does not need: %s", unwanted, fs[0].Next)
		}
	}
}

// assertSealedLockRemedy holds the two routes and the reason to one place, because the two cases
// that share a remedy should fail together if a route is dropped.
func assertSealedLockRemedy(t *testing.T, next, stale, gated string) {
	t.Helper()
	for what, want := range map[string]string{
		"the re-run route":                 "start " + stale + " over",
		"what the re-run discards":         "every phase after it",
		"the release route":                "approves",
		"the phase the release is made on": gated,
		"why the obvious route does not":   "section set and phase finish will not clear it",
	} {
		if !strings.Contains(next, want) {
			t.Fatalf("the remedy does not name %s (%q): %s", what, want, next)
		}
	}
}
