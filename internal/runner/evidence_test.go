// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/gates"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
)

// ---- WP6: the declaration is written by a command

// report writes a file standing in for what a pipeline published and returns its path. It
// lies outside the repository, because that is where one arrives from: a download, not a
// file somebody put in the phase directory by hand.
func (f *fixture) report(name, content string) string {
	f.t.Helper()
	p := filepath.Join(f.t.TempDir(), name)
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		f.t.Fatal(err)
	}
	return p
}

// verification is the state every one of these tests needs: the four phases before it
// judged, and P4 started with its artifact written and no verdict over it yet. A phase is
// where a declaration belongs, and the sequence is what lets one begin.
func (f *fixture) verification() {
	f.t.Helper()
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	f.started("04-verification")
}

func (f *fixture) declare(phase, file string, e model.EvidenceItem) (*model.EvidenceItem, error) {
	f.t.Helper()
	return f.r.DeclareEvidence(key, phase, file, e)
}

func (f *fixture) declared(phase string) []model.EvidenceItem {
	f.t.Helper()
	var o model.Output
	f.readFront(phase, &o)
	return o.Evidence
}

func TestAnItemThatExistsIsDeclaredFromItsFileAndTheRunnerHashesIt(t *testing.T) {
	f := newFixture(t)
	f.verification()
	content := "<testsuite failures=\"0\"/>\n"
	e, err := f.declare("04-verification", f.report("junit.xml", content),
		model.EvidenceItem{Kind: "test-report", Job: "unit", Result: "pass", Format: "junit"})
	f.must(err)
	if e.Path != "evidence/junit.xml" {
		t.Fatalf("the report was not declared where it was copied: %q", e.Path)
	}
	if e.SHA256 != hashing.Hex(hashing.Normalise([]byte(content))) {
		t.Fatalf("the declared hash is not the hash of the content: %q", e.SHA256)
	}
	// The copy is beside the artifact, which is what makes the phase self contained and
	// what G-Evidence resolves the item through.
	b, err := os.ReadFile(filepath.Join(f.root, model.PhaseDir(key, "04-verification"), e.Path))
	f.must(err)
	if string(b) != content {
		t.Fatalf("the report was not copied into evidence/: %q", b)
	}
	if items := f.declared("04-verification"); len(items) != 1 || items[0].SHA256 != e.SHA256 {
		t.Fatalf("the declaration did not reach the frontmatter: %+v", items)
	}
}

// The item a declaration makes is the one G-Evidence judges, and it judges it by resolving
// the path and comparing the hash. Both directions are asserted here rather than in the
// trail, because a trail that demonstrated them would be a repository with a red verdict.
func TestADeclaredItemIsJudgedAndFailsWhenItsContentChanges(t *testing.T) {
	f := newFixture(t)
	f.verification()
	_, err := f.declare("04-verification", f.report("junit.xml", "<testsuite failures=\"0\"/>\n"),
		model.EvidenceItem{Kind: "test-report", Job: "unit", Result: "pass"})
	f.must(err)
	if g := f.finish("04-verification"); g.Status != "green" || check(g, "G-Evidence").Result != "pass" {
		t.Fatalf("a resolved declaration did not pass: %s", g.Status)
	}
	rel := model.PhaseDir(key, "04-verification") + "/evidence/junit.xml"
	f.write(rel, "<testsuite failures=\"3\"/>\n")
	g, err := f.r.compute(key, "04-verification")
	f.must(err)
	if c := check(g, "G-Evidence"); c.Result != "fail" ||
		!strings.Contains(c.Findings[0].Cause, "does not match its hash") {
		t.Fatalf("an edited report was not caught: %+v", c)
	}
	f.must(os.Remove(filepath.Join(f.root, rel)))
	g, err = f.r.compute(key, "04-verification")
	f.must(err)
	if c := check(g, "G-Evidence"); c.Result != "fail" ||
		!strings.Contains(c.Findings[0].Cause, "does not resolve") {
		t.Fatalf("a removed report was not caught: %+v", c)
	}
}

func TestAnItemInAStoreIsDeclaredWithTheHashThePipelinePublished(t *testing.T) {
	f := newFixture(t)
	f.verification()
	e, err := f.declare("04-verification", "", model.EvidenceItem{
		Kind: "scan", Job: "trivy", Result: "pass",
		URI: "https://ci.example/runs/4711/trivy.json", SHA256: hex64('a'),
	})
	f.must(err)
	if e.Path != "" || e.URI == "" || e.SHA256 != hex64('a') {
		t.Fatalf("a uri item was not declared as given: %+v", e)
	}
	if _, err := os.Stat(filepath.Join(f.root, model.PhaseDir(key, "04-verification"), "evidence")); err == nil {
		t.Fatal("a uri item created an evidence directory it has nothing to put in")
	}
}

// The pending form is the one the whole pull path rests on: what it writes is the pair the
// attach looks an entry up by, so this asserts the declaration and the attachment together.
func TestAPendingItemDeclaresItsKindAndJobAndIsWhatTheAttachBindsTo(t *testing.T) {
	f := newFixture(t)
	f.verification()
	e, err := f.declare("04-verification", "", model.EvidenceItem{
		Kind: "test-report", Job: "unit", Format: "go-test-json",
		ProducedBy: "go test -json ./...",
	})
	f.must(err)
	if !e.Pending() || e.Path != "" || e.Result != "" {
		t.Fatalf("a pending item carries more than its kind and its job: %+v", e)
	}
	if g := f.finish("04-verification"); g.Status != "provisional" {
		t.Fatalf("a phase with a pending declaration is not provisional: %s", g.Status)
	}
	f.r.EvidenceFrom = f.pipeline("pass")
	f.must(f.r.Start(key, "05-review"))
	g, err := f.r.readGate(key, "04-verification")
	f.must(err)
	if g.Status != "green" {
		t.Fatalf("the declaration was not bound by the attach: %s", g.Status)
	}
}

// Provenance is written where it is given and absent where it is not. Nothing defaults it:
// A35's rule is that a plausible value in a field nobody produced is worse than an absent
// one, and A-002 says what the field means on an item a pipeline has yet to produce.
func TestProvenanceIsWrittenWhereGivenAndAbsentWhereNot(t *testing.T) {
	f := newFixture(t)
	f.verification()
	_, err := f.declare("04-verification", "", model.EvidenceItem{
		Kind: "scan", Job: "lint", Format: "other", ProducedBy: "golangci-lint run",
	})
	f.must(err)
	_, err = f.declare("04-verification", "", model.EvidenceItem{Kind: "sbom", Job: "sbom"})
	f.must(err)
	items := f.declared("04-verification")
	if len(items) != 2 || items[0].ProducedBy == "" || items[0].Format != "other" {
		t.Fatalf("provenance was not written as given: %+v", items)
	}
	if items[1].ProducedBy != "" || items[1].Format != "" {
		t.Fatalf("provenance was defaulted where nobody gave it: %+v", items[1])
	}
}

// Every refusal in one table, because each is one sentence of section 4 and a table is
// where a missing one shows. None of them may write: the assertion after the loop is that
// the phase holds no declaration and no evidence directory at all.
func TestADeclarationIsRefusedBeforeItIsWritten(t *testing.T) {
	f := newFixture(t)
	f.verification()
	file := f.report("junit.xml", "<testsuite/>\n")
	for _, c := range []struct {
		why  string
		file string
		item model.EvidenceItem
	}{
		{"no kind", file, model.EvidenceItem{Job: "unit", Result: "pass"}},
		{"a kind section 4 does not define", file, model.EvidenceItem{Kind: "tests", Job: "unit", Result: "pass"}},
		{"a result section 4 does not define", file, model.EvidenceItem{Kind: "test-report", Job: "unit", Result: "success"}},
		{"a format section 4 does not define", file, model.EvidenceItem{Kind: "test-report", Job: "unit", Result: "pass", Format: "go-test"}},
		{"a test report that exists and reports nothing", file, model.EvidenceItem{Kind: "test-report", Job: "unit"}},
		{"both places at once", file, model.EvidenceItem{Kind: "scan", Job: "trivy", URI: "https://ci.example/x", SHA256: hex64('a')}},
		{"a hash beside the bytes", file, model.EvidenceItem{Kind: "scan", Job: "trivy", SHA256: hex64('a')}},
		{"a uri with nothing binding it", "", model.EvidenceItem{Kind: "scan", Job: "trivy", URI: "https://ci.example/x"}},
		{"a hash with nothing it belongs to", "", model.EvidenceItem{Kind: "scan", Job: "trivy", SHA256: hex64('a')}},
		{"a pending item with no job", "", model.EvidenceItem{Kind: "scan"}},
		{"a pending item reporting a run that has not happened", "", model.EvidenceItem{Kind: "scan", Job: "trivy", Result: "pass"}},
		{"a report that cannot be read", filepath.Join(f.t.TempDir(), "absent.xml"), model.EvidenceItem{Kind: "scan", Job: "trivy"}},
	} {
		if _, err := f.declare("04-verification", c.file, c.item); !isRefusal(err) {
			t.Fatalf("%s was not refused: %v", c.why, err)
		}
	}
	if items := f.declared("04-verification"); len(items) != 0 {
		t.Fatalf("a refusal wrote a declaration: %+v", items)
	}
	if _, err := os.Stat(filepath.Join(f.root, model.PhaseDir(key, "04-verification"), "evidence")); err == nil {
		t.Fatal("a refusal copied a report into evidence/")
	}
}

// The command's refusals and the gate's findings are one judgement, so the shape refusals
// above are the gate's own and not a second opinion. Asserted by calling the gate directly
// on an item the command declined.
func TestTheRefusalIsTheGatesOwnJudgement(t *testing.T) {
	item := model.EvidenceItem{Kind: "test-report", Job: "unit", Result: "success", SHA256: hex64('a'), Path: "evidence/x"}
	fs := gates.EvidenceShape("output.md", model.Output{Evidence: []model.EvidenceItem{item}})
	if len(fs) == 0 {
		t.Fatal("the gate accepts what the command refuses, so there are two opinions about shape")
	}
}

func TestAPairIsDeclaredOnce(t *testing.T) {
	f := newFixture(t)
	f.verification()
	_, err := f.declare("04-verification", "", model.EvidenceItem{Kind: "scan", Job: "trivy"})
	f.must(err)
	_, err = f.declare("04-verification", "", model.EvidenceItem{Kind: "scan", Job: "trivy"})
	if !isRefusal(err) || !strings.Contains(err.Error(), "scan/trivy") {
		t.Fatalf("the pair was declared twice: %v", err)
	}
	// A different job of the same kind is a different item, which is what the three scans
	// of this repository's own pipeline are.
	_, err = f.declare("04-verification", "", model.EvidenceItem{Kind: "scan", Job: "lint"})
	f.must(err)
	if items := f.declared("04-verification"); len(items) != 2 {
		t.Fatalf("two jobs of one kind are two items: %+v", items)
	}
}

func TestDeclaringOverAnotherReportIsRefused(t *testing.T) {
	f := newFixture(t)
	f.verification()
	_, err := f.declare("04-verification", f.report("report.json", "{\"first\":true}\n"),
		model.EvidenceItem{Kind: "scan", Job: "lint", Result: "pass"})
	f.must(err)
	_, err = f.declare("04-verification", f.report("report.json", "{\"second\":true}\n"),
		model.EvidenceItem{Kind: "scan", Job: "trivy", Result: "pass"})
	if !isRefusal(err) {
		t.Fatalf("a second report took the first one's name: %v", err)
	}
	b, err := os.ReadFile(filepath.Join(f.root, model.PhaseDir(key, "04-verification"), "evidence/report.json"))
	f.must(err)
	if string(b) != "{\"first\":true}\n" {
		t.Fatalf("the first report was overwritten: %q", b)
	}
}

// The artifact is created by its first section write, as it is for a question. A command
// that created it here would write a frontmatter of whatever the runner could work out on
// its own, which is the state G-Schema reports as a phase nothing has produced.
func TestDeclaringIntoAPhaseWithNoArtifactIsRefused(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	f.must(f.r.Start(key, "04-verification"))
	_, err := f.declare("04-verification", "", model.EvidenceItem{Kind: "scan", Job: "trivy"})
	if !isRefusal(err) || !strings.Contains(err.Error(), "section set") {
		t.Fatalf("a declaration into a phase with no artifact was not refused: %v", err)
	}
}

// A write after a verdict is staged, which is #304's answer to the state #216 described: the
// declaration is not lost, it does not rewrite the verdict, and it leaves the judged phase
// exactly as the verdict found it. The next finish applies it and judges once, so the stale
// window #216 was about never opens.
func TestADeclarationAfterAVerdictIsStagedAndLeavesTheSealAlone(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	g := f.run("04-verification", "")
	sealed := g.ArtifactsHash
	_, err := f.declare("04-verification", "", model.EvidenceItem{Kind: "scan", Job: "trivy"})
	f.must(err)
	after, err := f.r.readGate(key, "04-verification")
	f.must(err)
	if after.ArtifactsHash != sealed {
		t.Fatal("the declaration rewrote the verdict")
	}
	if f.hash("04-verification") != sealed {
		t.Fatal("the declaration changed the artifact the verdict covers")
	}
	if staged := f.r.staged(key, "04-verification"); len(staged) != 1 || staged[0] != "output.md" {
		t.Fatalf("the declaration was not staged: %v", staged)
	}
	// And the next finish is what moves both together. Finish directly, not through the
	// fixture's run: a second phase start is refused on a judged phase, which is #215, and the
	// refusal names these two commands as the way to redo one.
	g2, err := f.r.Finish(key, "04-verification", "")
	f.must(err)
	if g2.ArtifactsHash == sealed {
		t.Fatal("the finish did not apply the staged declaration")
	}
	if f.hash("04-verification") != g2.ArtifactsHash {
		t.Fatal("the finish sealed something other than what it judged")
	}
	if staged := f.r.staged(key, "04-verification"); staged != nil {
		t.Fatalf("the staged copy outlived the finish that applied it: %v", staged)
	}
}

// The body and the rest of the frontmatter are untouched, which is what makes this safe to
// run in the middle of writing a phase.
func TestADeclarationTouchesNothingElse(t *testing.T) {
	f := newFixture(t)
	f.verification()
	before, hash := f.body("04-verification"), f.frontField("04-verification", "context_hash")
	_, err := f.declare("04-verification", "", model.EvidenceItem{Kind: "scan", Job: "trivy"})
	f.must(err)
	if f.body("04-verification") != before {
		t.Fatal("the body changed")
	}
	if f.frontField("04-verification", "context_hash") != hash {
		t.Fatal("a frontmatter field other than evidence changed")
	}
}
