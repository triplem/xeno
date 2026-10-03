// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
)

// phaseWithProfile writes a P0 profile, a lock recording a base, and the files that base names, so
// that the budget has two files to compare and a tree to measure.
func phaseWithProfile(t *testing.T, profile string, files map[string]string) Ctx {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, model.ContextProfile), []byte(profile), 0o644); err != nil {
		t.Fatal(err)
	}
	lock := "intent: " + fixtureIntent + "\nphase: " + model.Phases[0] + "\nfiles:\n"
	for rel, body := range files {
		p := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
		lock += "  - path: " + rel + "\n    sha256: " + strings.Repeat("0", 64) + "\n"
	}
	if err := os.WriteFile(filepath.Join(dir, "context.lock.yaml"), []byte(lock), 0o644); err != nil {
		t.Fatal(err)
	}
	return Ctx{Root: root, Key: "PROJ-1", Phase: model.Phases[0], QualifiedID: fixtureIntent}
}

func TestNoProfileAndNoBudgetAreNoFinding(t *testing.T) {
	if fs := budget(Ctx{Root: t.TempDir(), Key: "PROJ-1", Phase: model.Phases[0]}); fs != nil {
		t.Fatalf("a repository with no profile produced %v", fs)
	}
	c := phaseWithProfile(t, "include:\n  - src/**\n", map[string]string{"src/a.go": "package a\n"})
	if fs := budget(c); fs != nil {
		t.Fatalf("a profile with no budget produced %v", fs)
	}
}

// Section 5: a finding, with both numbers, and the phase decidable rather than blocked.
func TestAContextOverItsFileBudgetIsAFinding(t *testing.T) {
	c := phaseWithProfile(t, "include:\n  - src/**\nbudget:\n  files: 2\n", map[string]string{
		"src/a.go": "package a\n", "src/b.go": "package b\n", "src/c.go": "package c\n",
	})
	fs := budget(c)
	if len(fs) != 1 {
		t.Fatalf("%d findings, want 1: %v", len(fs), fs)
	}
	if !strings.Contains(fs[0].Cause, "3 files") || !strings.Contains(fs[0].Cause, "budget is 2") {
		t.Errorf("the cause is %q, want both numbers", fs[0].Cause)
	}
	if !strings.Contains(fs[0].Next, "raise the budget") {
		t.Errorf("the next step is %q", fs[0].Next)
	}
	// Within the budget, nothing.
	ok := phaseWithProfile(t, "include:\n  - src/**\nbudget:\n  files: 3\n", map[string]string{
		"src/a.go": "package a\n", "src/b.go": "package b\n", "src/c.go": "package c\n",
	})
	if fs := budget(ok); fs != nil {
		t.Fatalf("a context inside its budget produced %v", fs)
	}
}

func TestAContextOverItsByteBudgetIsAFinding(t *testing.T) {
	c := phaseWithProfile(t, "include:\n  - src/**\nbudget:\n  bytes: 10\n", map[string]string{
		"src/a.go": strings.Repeat("x", 40),
	})
	fs := budget(c)
	if len(fs) != 1 || !strings.Contains(fs[0].Cause, "40 bytes") {
		t.Fatalf("findings %v, want one naming the size", fs)
	}
}

// A file the lock names and the tree has lost is skipped: its absence is G-Freshness's finding, and
// one cause reported by two gates is what #160 avoided.
func TestAMissingFileIsNotCountedAndNotReportedTwice(t *testing.T) {
	c := phaseWithProfile(t, "include:\n  - src/**\nbudget:\n  bytes: 30\n", map[string]string{
		"src/a.go": strings.Repeat("x", 20),
	})
	if err := os.Remove(filepath.Join(c.Root, "src/a.go")); err != nil {
		t.Fatal(err)
	}
	if fs := budget(c); fs != nil {
		t.Fatalf("a lost file produced %v from the budget", fs)
	}
}

// The gate the section puts it in, so that it is decidable like any other finding.
func TestTheBudgetFindingComesFromGSchema(t *testing.T) {
	root, _ := corpus(t)
	dir := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]))
	if err := os.WriteFile(filepath.Join(root, "src.go"), []byte(strings.Repeat("x", 40)), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, model.ContextProfile),
		[]byte("include:\n  - src.go\nbudget:\n  bytes: 1\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// The corpus's own lock carries no files, so the base is appended to it: the budget compares
	// the profile against the lock, and a lock with nothing in it is a phase that was given
	// nothing.
	lock, err := os.ReadFile(filepath.Join(dir, "context.lock.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	lock = append(lock, []byte("files:\n  - path: src.go\n    sha256: "+strings.Repeat("0", 64)+"\n")...)
	if err := os.WriteFile(filepath.Join(dir, "context.lock.yaml"), lock, 0o644); err != nil {
		t.Fatal(err)
	}
	got := schema(ctxFor(root))
	if !strings.Contains(causes(got), "the budget is 1") {
		t.Fatalf("G-Schema did not report the budget: %s", causes(got))
	}
}

// ---- a declared link whose document is not there (#172)

// #171's criterion, met by #172: "a link naming a document that does not exist is a finding against
// the profile rather than a silently missing file".
func TestALinkNamingADocumentThatIsNotThereIsAFinding(t *testing.T) {
	c := phaseWithProfile(t,
		"include:\n  - src/**\nlinks:\n  - component: src/payment\n    docs: docs/adr/0012-payments.md\n",
		map[string]string{"src/a.go": "package a\n"})
	fs := links(c)
	if len(fs) != 1 {
		t.Fatalf("%d findings, want 1: %v", len(fs), fs)
	}
	if !strings.Contains(fs[0].Cause, "src/payment") || !strings.Contains(fs[0].Cause, "docs/adr/0012-payments.md") {
		t.Errorf("the cause is %q, want the component and the path", fs[0].Cause)
	}
	if !strings.Contains(fs[0].File, model.ContextProfile) {
		t.Errorf("the finding names %q, want the profile: the profile is the claim", fs[0].File)
	}
	if fs[0].Next == "" {
		t.Error("the finding has no next step")
	}
}

// A link whose document is there is not a finding, and nothing is reported for a profile with no
// links or a repository with no profile.
func TestALinkWhoseDocumentExistsIsNoFinding(t *testing.T) {
	c := phaseWithProfile(t,
		"include:\n  - src/**\nlinks:\n  - component: src/payment\n    docs: docs/adr/0012.md\n",
		map[string]string{"src/a.go": "package a\n", "docs/adr/0012.md": "# payments\n"})
	if fs := links(c); fs != nil {
		t.Fatalf("a link whose document exists produced %v", fs)
	}
	bare := phaseWithProfile(t, "include:\n  - src/**\n", map[string]string{"src/a.go": "package a\n"})
	if fs := links(bare); fs != nil {
		t.Fatalf("a profile with no links produced %v", fs)
	}
	if fs := links(Ctx{Root: t.TempDir(), Key: "PROJ-1", Phase: model.Phases[0]}); fs != nil {
		t.Fatalf("a repository with no profile produced %v", fs)
	}
}

// The gate the finding belongs to, so that it is decidable like any other and reported at every
// phase rather than once at P0.
func TestTheLinkFindingComesFromGSchemaAtAnyPhase(t *testing.T) {
	root, _ := corpus(t)
	dir := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]))
	if err := os.WriteFile(filepath.Join(dir, model.ContextProfile),
		[]byte("include:\n  - src/**\nlinks:\n  - component: src\n    docs: docs/nobody.md\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	got := schema(ctxFor(root))
	if !strings.Contains(causes(got), "which is not in the tree") {
		t.Fatalf("G-Schema did not report the link: %s", causes(got))
	}
}
