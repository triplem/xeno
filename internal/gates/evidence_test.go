// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
)

// declaration is the closed sets of section 4, read from testdata rather than from the
// constants the code uses. Written the other way round these tests would prove only that
// the package agrees with itself, which is how G-Build came to read a kind no document
// defines while every test passed.
type declaration struct {
	Kinds            []string `yaml:"kinds"`
	Results          []string `yaml:"results"`
	ResultRequiredOn []string `yaml:"result_required_on"`
	BuildKind        string   `yaml:"build_kind"`
}

func documented(t *testing.T) declaration {
	t.Helper()
	var d declaration
	if err := fm.ReadYAML("testdata/evidence-declaration.yaml", &d); err != nil {
		t.Fatal(err)
	}
	return d
}

// The sets the code carries are the sets section 4 fixes. This is the check that was
// missing: with it, changing a constant without changing the document fails here.
func TestTheClosedSetsAreTheDocumentsOwn(t *testing.T) {
	d := documented(t)
	for _, c := range []struct {
		name       string
		want, have []string
	}{
		{"kinds", d.Kinds, model.EvidenceKinds},
		{"results", d.Results, model.EvidenceResults},
		{"result_required_on", d.ResultRequiredOn, model.ResultRequiredKinds},
	} {
		if strings.Join(c.want, ",") != strings.Join(c.have, ",") {
			t.Errorf("%s: section 4 fixes %v, the code carries %v", c.name, c.want, c.have)
		}
	}
	if d.BuildKind != BuildKind {
		t.Errorf("G-Build reads %q, section 4 spells it %q", BuildKind, d.BuildKind)
	}
	if !model.OneOf(BuildKind, d.Kinds) {
		t.Errorf("G-Build reads %q, which is not in the set at all", BuildKind)
	}
}

// phaseWith writes a phase carrying one evidence declaration, as frontmatter of output.md.
// Only the fields these tests read are written; G-Schema's other findings are another
// test's subject and are filtered out below.
func phaseWith(t *testing.T, phase, item string) Ctx {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, model.PhaseDir("PROJ-1", phase))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	out := "---\nintent: " + fixtureIntent + "\nphase: " + phase + "\nevidence:\n" + item + "---\n\nbody\n"
	if err := os.WriteFile(filepath.Join(dir, "output.md"), []byte(out), 0o644); err != nil {
		t.Fatal(err)
	}
	return Ctx{Root: root, Key: "PROJ-1", Phase: phase, QualifiedID: fixtureIntent}
}

// declarationFindings is what EvidenceShape said about the one item, without the findings
// the deliberately sparse fixture provokes elsewhere in G-Schema.
func declarationFindings(t *testing.T, item string) string {
	t.Helper()
	c := phaseWith(t, model.Phases[0], item)
	var o model.Output
	if _, err := fm.ReadFront(c.abs(c.phaseRel(c.Phase)+"/output.md"), &o); err != nil {
		t.Fatal(err)
	}
	if len(o.Evidence) != 1 {
		t.Fatalf("the fixture declared %d items, not one", len(o.Evidence))
	}
	return causes(result(EvidenceShape("output.md", o)))
}

const sealed = "    sha256: " + "0000000000000000000000000000000000000000000000000000000000000000" +
	"\n    uri: https://example.invalid/report\n"

// Every kind the document defines is accepted, and one it does not is a finding. The loop
// is over the document's set, so a kind added there and not here fails.
func TestEveryDocumentedKindIsAcceptedAndOthersAreNot(t *testing.T) {
	d := documented(t)
	for _, kind := range d.Kinds {
		item := "  - kind: " + kind + "\n    job: j\n" + sealed
		if model.OneOf(kind, d.ResultRequiredOn) {
			item += "    result: pass\n"
		}
		if got := declarationFindings(t, item); got != "" {
			t.Errorf("%s is a kind section 4 defines and was rejected:\n%s", kind, got)
		}
	}
	got := declarationFindings(t, "  - kind: test\n    job: j\n    result: pass\n"+sealed)
	if !strings.Contains(got, "kind section 4 does not define") {
		t.Errorf("kind: test is not in the set and passed:\n%s", got)
	}
}

// The asymmetry of section 4: required on two kinds, and absent elsewhere without
// complaint, because there the producer had no threshold to report against.
func TestResultIsRequiredOnTwoKindsAndOptionalOnTheRest(t *testing.T) {
	d := documented(t)
	for _, kind := range d.Kinds {
		got := declarationFindings(t, "  - kind: "+kind+"\n    job: j\n"+sealed)
		named := strings.Contains(got, "carries no result")
		if want := model.OneOf(kind, d.ResultRequiredOn); named != want {
			t.Errorf("%s without a result: finding=%v, section 4 wants %v:\n%s", kind, named, want, got)
		}
	}
}

// A pending item declares only its kind and its job, so it has no result to carry. Without
// this exemption every declaration of evidence a pipeline owes would be a finding, which is
// the state P4 is designed to be in between its finish and its pipeline.
func TestAPendingItemOwesNoResult(t *testing.T) {
	for _, kind := range documented(t).ResultRequiredOn {
		if got := declarationFindings(t, "  - kind: "+kind+"\n    job: j\n"); got != "" {
			t.Errorf("a pending %s was asked for a result the pipeline has not produced:\n%s", kind, got)
		}
	}
}

func TestAResultOutsideTheSetIsAFinding(t *testing.T) {
	got := declarationFindings(t, "  - kind: build-log\n    job: compile\n    result: green\n"+sealed)
	if !strings.Contains(got, "has result green") {
		t.Errorf("result: green is outside the set and passed:\n%s", got)
	}
}

// The regression this issue is about. G-Build read `build`, which section 4 does not
// define, so a declared build carrying result: fail passed. The kind is taken from the
// document, so a rename there fails here rather than silently making the gate inert again.
func TestGBuildJudgesTheKindTheDocumentDefines(t *testing.T) {
	kind := documented(t).BuildKind
	for _, c := range []struct {
		result string
		want   string
	}{{"fail", "fail"}, {"pass", "pass"}} {
		ctx := phaseWith(t, "03-implementation",
			"  - kind: "+kind+"\n    job: compile\n    result: "+c.result+"\n"+sealed)
		if got := build(ctx); got.Result != c.want {
			t.Errorf("a %s build declared as %s: G-Build said %s, want %s\n%s",
				c.result, kind, got.Result, c.want, causes(got))
		}
	}
}

// A build the pipeline still owes is pending, not failing: there is no result yet to be
// unsuccessful. Kept beside the test above because the two states used to be one.
func TestAPendingBuildIsPendingAndNotAFailure(t *testing.T) {
	ctx := phaseWith(t, "03-implementation", "  - kind: "+documented(t).BuildKind+"\n    job: compile\n")
	if got := build(ctx); got.Result != "pending" {
		t.Errorf("a build the pipeline owes read %s, want pending\n%s", got.Result, causes(got))
	}
}

// ---- G-Evidence on an attachment that binds nothing

// attachedPhase writes a phase with one pending declaration and one attached record, which
// is the state only a hand edit can now produce: the attach declines to write it.
func attachedPhase(t *testing.T, record string) Ctx {
	t.Helper()
	c := phaseWith(t, "04-verification", "  - kind: scan\n    job: trivy\n")
	evDir := c.abs(c.phaseRel(c.Phase) + "/evidence")
	if err := os.MkdirAll(evDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(evDir, "attached.yaml"), []byte(record), 0o644); err != nil {
		t.Fatal(err)
	}
	return c
}

// `evidence/attached.yaml` lies outside the artifacts_hash by design, so it is the one file
// in a judged phase that a person can edit without making any verdict stale. This check is
// what the missing seal is replaced by, and it is the reason the guard exists twice.
func TestAnAttachmentWithAUriAndNoHashIsAFinding(t *testing.T) {
	c := attachedPhase(t, "- kind: scan\n  job: trivy\n  state: attached\n  result: pass\n"+
		"  uri: https://ci.example/a/7\n  sha256: \"\"\n")
	got := evidence(c)
	if got.Result != "fail" {
		t.Fatalf("an attachment nothing binds read %s, want fail\n%s", got.Result, causes(got))
	}
	if !strings.Contains(causes(got), "nothing binds it") {
		t.Errorf("the finding does not say what is wrong:\n%s", causes(got))
	}
	// It belongs to the file somebody had to edit, not to the declaration, which is correct.
	if !strings.Contains(causes(got), "evidence/attached.yaml") {
		t.Errorf("the finding does not name attached.yaml:\n%s", causes(got))
	}
}

func TestAnAttachmentWithAUriAndAHashPasses(t *testing.T) {
	c := attachedPhase(t, "- kind: scan\n  job: trivy\n  state: attached\n  result: pass\n"+
		"  uri: https://ci.example/a/7\n  sha256: "+strings.Repeat("3", 64)+"\n")
	if got := evidence(c); got.Result != "pass" {
		t.Fatalf("a bound attachment read %s, want pass\n%s", got.Result, causes(got))
	}
}

// Nothing attached at all is still pending rather than failing: the pipeline owes it and
// has not answered. The two states were never confused and this keeps them apart.
func TestNothingAttachedIsStillPending(t *testing.T) {
	c := phaseWith(t, "04-verification", "  - kind: scan\n    job: trivy\n")
	if got := evidence(c); got.Result != "pending" {
		t.Fatalf("an unanswered declaration read %s, want pending\n%s", got.Result, causes(got))
	}
}

// ---- #221: G-Evidence judges the result an attachment carries

// The other half of what #208 left. The declaration is sealed inside artifacts_hash and
// judged by G-Schema; this file is the one in a judged phase that can be edited without
// making any verdict stale, so what it says is checked here or nowhere. Fifteen records in
// this repository read `success` until the commit that added this check corrected them.
func TestAnAttachedResultOutsideTheSetIsAFinding(t *testing.T) {
	c := attachedPhase(t, "- kind: scan\n  job: trivy\n  state: attached\n  result: success\n"+
		"  uri: https://ci.example/a/7\n  sha256: "+strings.Repeat("3", 64)+"\n")
	got := evidence(c)
	if got.Result != "fail" {
		t.Fatalf("an attachment carrying result: success read %s, want fail\n%s", got.Result, causes(got))
	}
	if !strings.Contains(causes(got), "attached with result success") {
		t.Errorf("the finding does not say what the record carries:\n%s", causes(got))
	}
	// On attached.yaml, which is the file somebody had to edit for the record to say this.
	if !strings.Contains(causes(got), "evidence/attached.yaml") {
		t.Errorf("the finding does not name attached.yaml:\n%s", causes(got))
	}
}

// The set is the document's, read from testdata, so a value added there and not here fails.
func TestEveryDocumentedResultIsAcceptedOnAnAttachment(t *testing.T) {
	for _, res := range documented(t).Results {
		c := attachedPhase(t, "- kind: scan\n  job: trivy\n  state: attached\n  result: "+res+"\n"+
			"  uri: https://ci.example/a/7\n  sha256: "+strings.Repeat("3", 64)+"\n")
		if got := evidence(c); got.Result != "pass" {
			t.Errorf("an attachment reporting %s read %s, want pass\n%s", res, got.Result, causes(got))
		}
	}
}

// A value is judged and an absence is not, which keeps the entry this repository's own
// vulnerability scan publishes beside its report — a database metadata file with no result —
// out of the findings. Which kinds must carry one is the declaration's question.
func TestAnAttachmentWithNoResultIsNotAFinding(t *testing.T) {
	c := attachedPhase(t, "- kind: scan\n  job: trivy\n  state: attached\n"+
		"  uri: https://ci.example/a/7\n  sha256: "+strings.Repeat("3", 64)+"\n")
	if got := evidence(c); got.Result != "pass" {
		t.Fatalf("an attachment whose producer reports nothing read %s, want pass\n%s", got.Result, causes(got))
	}
}
