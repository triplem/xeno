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

// declarationFindings is what evidenceShape said about the one item, without the findings
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
	return causes(result(evidenceShape("output.md", o)))
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
