// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
)

// Section 5, "Acceptance criteria are identifiable": the numbered list from
// requirements@1.1.0, the mapping naming each number from verification@1.1.0, and the
// requirement carried by the template version so that an artifact declaring 1.0.0 is judged as
// it always was. Both halves had no reader until this (#258), and the anchor is what keeps the
// 130 sealed P1 and P4 artifacts of this trail out of scope.

// criteriaFixture writes one phase's artifact with a declared template ref and one section.
func criteriaFixture(t *testing.T, root, phase, ref, section, body string) {
	t.Helper()
	dir := filepath.Join(root, model.PhaseDir("PROJ-1", phase))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	art := "---\nintent: " + fixtureIntent + "\nphase: " + phase + "\nlanguage: en\ntemplate: " + ref +
		"\n---\n\n# Phase\n\n<!-- xeno:section:" + section + " -->\n## Heading\n\n" + body + "\n"
	if err := os.WriteFile(filepath.Join(dir, "output.md"), []byte(art), 0o644); err != nil {
		t.Fatal(err)
	}
}

func criteriaCtx(root, phase string) Ctx {
	return Ctx{Root: root, Key: "PROJ-1", Phase: phase, QualifiedID: fixtureIntent}
}

// A string compare puts 1.10.0 before 1.9.0, so the comparison is on the three numbers. A check
// keyed the other way would stop applying at the tenth minor version with nothing to say so.
func TestTemplateAtLeastComparesNumbersAndNotStrings(t *testing.T) {
	for _, c := range []struct {
		ref  string
		want bool
	}{
		{"requirements@1.1.0", true},
		{"requirements@1.0.0", false},
		{"requirements@1.0.9", false},
		{"requirements@1.2.0", true},
		{"requirements@1.10.0", true},
		{"requirements@2.0.0", true},
		{"verification@1.1.0", false}, // another id is not this clause's subject
		{"requirements@1", false},     // no minor to compare
		{"requirements@one.two", false},
		{"requirements", false},
		{"", false},
	} {
		if got := templateAtLeast(c.ref, "requirements", 1, 1); got != c.want {
			t.Errorf("templateAtLeast(%q) = %v, want %v", c.ref, got, c.want)
		}
	}
}

// The claim the numeric comparison exists for, asserted rather than left to the table above: a
// string compare puts "1.10.0" below "1.9.0" and the check would stop applying at the tenth minor.
func TestTenthMinorVersionIsLaterThanTheNinth(t *testing.T) {
	if !templateAtLeast("requirements@1.10.0", "requirements", 1, 9) {
		t.Fatal("1.10.0 did not compare as later than 1.9")
	}
	if templateAtLeast("requirements@1.9.0", "requirements", 1, 10) {
		t.Fatal("1.9.0 compared as later than 1.10")
	}
	if "1.10.0" >= "1.9.0" {
		t.Fatal("the premise is wrong: a string compare would have ordered these correctly")
	}
}

// "From requirements@1.1.0 the acceptance-criteria section is a numbered list."
func TestAtTheNewTemplateUnnumberedCriteriaAreReported(t *testing.T) {
	root := t.TempDir()
	criteriaFixture(t, root, model.Phases[1], "requirements@1.1.0", "acceptance-criteria",
		"The runner refuses a second start.\nThe lock is written once.")

	fs := numberedCriteria(criteriaCtx(root, model.Phases[1]))
	if len(fs) != 1 {
		t.Fatalf("want one finding, got %d: %v", len(fs), findingText(fs))
	}
	if !strings.Contains(fs[0].Cause, "requirements@1.1.0") {
		t.Fatalf("the finding does not name the version that requires it: %s", fs[0].Cause)
	}
}

func TestAtTheNewTemplateNumberedCriteriaPass(t *testing.T) {
	root := t.TempDir()
	criteriaFixture(t, root, model.Phases[1], "requirements@1.1.0", "acceptance-criteria",
		"1. The runner refuses a second start.\n2. The lock is written once.")

	if fs := numberedCriteria(criteriaCtx(root, model.Phases[1])); fs != nil {
		t.Fatalf("numbered criteria were reported: %v", findingText(fs))
	}
}

// The anchor, and the reason nothing in the trail is re-judged: "An artifact declaring
// requirements@1.0.0 is judged as it always was."
func TestAtTheOldTemplateUnnumberedCriteriaAreNotReported(t *testing.T) {
	root := t.TempDir()
	criteriaFixture(t, root, model.Phases[1], "requirements@1.0.0", "acceptance-criteria",
		"The runner refuses a second start.")

	if fs := numberedCriteria(criteriaCtx(root, model.Phases[1])); fs != nil {
		t.Fatalf("an artifact at the old template was judged by the new clause: %v", findingText(fs))
	}
}

// Section 7's G-Test row: "mapping of acceptance criteria complete".
func TestAMappingThatMissesACriterionIsReported(t *testing.T) {
	root := t.TempDir()
	criteriaFixture(t, root, model.Phases[1], "requirements@1.1.0", "acceptance-criteria",
		"1. refuses a second start\n2. writes the lock once\n3. names the phase")
	criteriaFixture(t, root, model.Phases[4], "verification@1.1.0", "test-mapping",
		"| # | how |\n|---|---|\n| 1 | a test |\n| 3 | a test |")

	fs := mappingComplete(criteriaCtx(root, model.Phases[4]))
	if len(fs) != 1 {
		t.Fatalf("want one finding, got %d: %v", len(fs), findingText(fs))
	}
	if !strings.Contains(fs[0].Cause, "criteria 2") || strings.Contains(fs[0].Cause, "1, 2") {
		t.Fatalf("the finding does not name the uncovered criterion alone: %s", fs[0].Cause)
	}
}

// A table cell, a list item and a sentence all "name each criterion by its number", which is as
// much as section 5 says. The check reads the number as a token and not a format.
func TestAMappingNamesItsCriteriaInAnyForm(t *testing.T) {
	root := t.TempDir()
	criteriaFixture(t, root, model.Phases[1], "requirements@1.1.0", "acceptance-criteria",
		"1. refuses a second start\n2. writes the lock once\n3. names the phase")
	criteriaFixture(t, root, model.Phases[4], "verification@1.1.0", "test-mapping",
		"| 1 | a test |\n\n2. another test\n\nCriterion 3 is read by a person.")

	if fs := mappingComplete(criteriaCtx(root, model.Phases[4])); fs != nil {
		t.Fatalf("a mapping naming all three was reported: %v", findingText(fs))
	}
}

// A P4 at 1.1.0 behind a P1 at 1.0.0 is every intent already under way when the bump lands.
// Reading it as a failure would make the bump re-judge what the anchor exists to protect.
func TestAMappingIsNotJudgedAgainstAnOlderRequirementsPhase(t *testing.T) {
	root := t.TempDir()
	criteriaFixture(t, root, model.Phases[1], "requirements@1.0.0", "acceptance-criteria",
		"1. refuses a second start\n2. writes the lock once")
	criteriaFixture(t, root, model.Phases[4], "verification@1.1.0", "test-mapping", "nothing numbered here")

	if fs := mappingComplete(criteriaCtx(root, model.Phases[4])); fs != nil {
		t.Fatalf("a mapping was judged against a sealed requirements phase: %v", findingText(fs))
	}
}

// With no requirements phase there is nothing to map against, which is a state and not a fault.
func TestAMappingWithNoRequirementsPhaseReportsNothing(t *testing.T) {
	root := t.TempDir()
	criteriaFixture(t, root, model.Phases[4], "verification@1.1.0", "test-mapping", "1. a test")

	if fs := mappingComplete(criteriaCtx(root, model.Phases[4])); fs != nil {
		t.Fatalf("a mapping with no predecessor was reported: %v", findingText(fs))
	}
}

// 1 must not match inside 10, or a mapping covering only criterion 10 would satisfy criterion 1.
func TestANumberIsNamedAsATokenAndNotAsASubstring(t *testing.T) {
	if namesNumber("| 10 | a test |", "1") {
		t.Fatal("1 matched inside 10")
	}
	if !namesNumber("| 10 | a test |", "10") {
		t.Fatal("10 did not match itself")
	}
	if !namesNumber("1. a test", "1") {
		t.Fatal("a leading number was not found")
	}
}

// Each check belongs to one phase, so neither answers for an artifact it is not about.
func TestEachCheckReadsOnlyItsOwnPhase(t *testing.T) {
	root := t.TempDir()
	criteriaFixture(t, root, model.Phases[2], "design@1.1.0", "decisions", "nothing numbered")

	if fs := numberedCriteria(criteriaCtx(root, model.Phases[2])); fs != nil {
		t.Fatalf("the criteria check judged a design phase: %v", findingText(fs))
	}
	if fs := mappingComplete(criteriaCtx(root, model.Phases[2])); fs != nil {
		t.Fatalf("the mapping check judged a design phase: %v", findingText(fs))
	}
}
