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

// Section 5: "An advisory finding does not fail its check. advisory marks a finding that is
// reported rather than held against the phase: the check carrying it is pass, the phase is green,
// and the finding is in gate.yaml with its id, its cause and its remedy like any other."
//
// The clause was in the document and the code contradicted it, with budget called from schema and
// result failing a check on any finding at all (#235). Nothing asserted the contradiction either:
// every budget test called budget directly and read its findings, so the one thing section 5 is
// explicit about — what the check result should be — had no test.

func TestACheckOfOnlyAdvisoryFindingsPasses(t *testing.T) {
	ch := result([]model.Finding{
		advisory(finding("a.md", "over the budget", "narrow the scope")),
		advisory(finding("a.md", "over it again", "narrow the scope")),
	})
	if ch.Result != "pass" {
		t.Fatalf("a check of advisory findings is %q, want pass", ch.Result)
	}
	if len(ch.Findings) != 2 {
		t.Fatalf("the findings were dropped: %d of 2", len(ch.Findings))
	}
}

// The case a count of findings cannot express, and the whole of the distinction: the ordinary
// finding fails the check and the advisory one rides along in the verdict.
func TestACheckWithOneOrdinaryFindingStillFails(t *testing.T) {
	ch := result([]model.Finding{
		advisory(finding("a.md", "over the budget", "narrow the scope")),
		finding("a.md", "required field missing", "add it"),
	})
	if ch.Result != "fail" {
		t.Fatalf("a check carrying an ordinary finding is %q, want fail", ch.Result)
	}
	if len(ch.Findings) != 2 {
		t.Fatalf("the advisory finding was dropped: %d of 2", len(ch.Findings))
	}
}

func TestACheckOfOrdinaryFindingsStillFailsAndAnEmptyOnePasses(t *testing.T) {
	if r := result([]model.Finding{finding("a.md", "cause", "next")}).Result; r != "fail" {
		t.Fatalf("an ordinary finding gives %q, want fail", r)
	}
	if r := result(nil).Result; r != "pass" {
		t.Fatalf("no findings gives %q, want pass", r)
	}
}

// "The finding is advisory, so G-Schema stays pass and the phase stays green." Read through the
// gate rather than through budget, because the fault was the call site and not the check.
func TestABudgetOverrunLeavesGSchemaPassing(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	write := func(name, body string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("context-scope.yaml", "intent: "+fixtureIntent+"\nphase: "+model.Phases[0]+
		"\ninclude:\n  - a.txt\nbudget:\n  files: 1\n  bytes: 10\n")
	write("context.lock.yaml", "intent: "+fixtureIntent+"\nphase: "+model.Phases[0]+
		"\nfiles:\n  - path: a.txt\n    sha256: "+strings.Repeat("a", 64)+"\n    bytes: 99\n")

	c := Ctx{Root: root, Key: "PROJ-1", Phase: model.Phases[0], QualifiedID: fixtureIntent}
	fs := budget(c)
	if len(fs) != 1 {
		t.Fatalf("want one budget finding, got %d: %v", len(fs), findingText(fs))
	}
	if !fs[0].Advisory {
		t.Fatal("the budget finding is not advisory")
	}
	if ch := result(fs); ch.Result != "pass" {
		t.Fatalf("a check of the budget finding is %q, want pass", ch.Result)
	}
}

// The half the check result does not cover, and the one a passing check hides: Status counts
// undecided findings, and an advisory finding is undecided for ever because nothing asks anybody
// to decide it. Without this the phase is red through a check that passed — which is what the
// end-to-end run found after result alone had been changed.
func TestAPhaseWhoseOnlyFindingIsAdvisoryIsGreen(t *testing.T) {
	st, err := Status([]model.Check{{
		Gate: "G-Schema", Result: "pass",
		Findings: []model.Finding{{ID: "F-1", File: "a.md", Cause: "over the budget", Advisory: true}},
	}})
	if err != nil {
		t.Fatal(err)
	}
	if st != "green" {
		t.Fatalf("a phase whose only finding is advisory is %q, want green", st)
	}
}

// And it is still a finding: an ordinary one beside it makes the phase red, and deciding the
// advisory one is allowed rather than ignored.
func TestAnAdvisoryFindingNeitherHidesNorBlocksTheOthers(t *testing.T) {
	st, err := Status([]model.Check{{
		Gate: "G-Schema", Result: "fail",
		Findings: []model.Finding{
			{ID: "F-1", File: "a.md", Cause: "over the budget", Advisory: true},
			{ID: "F-2", File: "a.md", Cause: "required field missing"},
		},
	}})
	if err != nil {
		t.Fatal(err)
	}
	if st != "red" {
		t.Fatalf("a phase with an ordinary finding is %q, want red", st)
	}

	// Section 5: it "can be decided like any other finding".
	st, err = Status([]model.Check{{
		Gate: "G-Schema", Result: "pass",
		Findings: []model.Finding{{ID: "F-1", File: "a.md", Cause: "over the budget", Advisory: true,
			Decision: &model.DecisionOnFinding{Type: "approved", By: "a person", Reason: "the budget was guessed"}}},
	}})
	if err != nil {
		t.Fatal(err)
	}
	if st != "approved" {
		t.Fatalf("an approved advisory finding gives %q, want approved", st)
	}
}

// "It is the exception and stays one… nothing else writes the field." No code can tell a clause
// that legitimately asked for this from a check somebody found inconvenient, so the bound is the
// sentence and this count. It fails when a second check reaches for the field, which is the moment
// worth catching.
func TestOnlyTheBudgetClauseWritesTheAdvisoryField(t *testing.T) {
	b, err := os.ReadFile("gates.go")
	if err != nil {
		t.Fatal(err)
	}
	if n := strings.Count(string(b), "advisory(finding("); n != 2 {
		t.Fatalf("%d findings are marked advisory, want the budget clause's 2; section 5 bounds "+
			"the field to the clause that asks for it", n)
	}
	if n := strings.Count(string(b), "Advisory = true"); n != 1 {
		t.Fatalf("the field is set in %d places, want only the advisory helper", n)
	}
}

// Section 5: a finding's id does not depend on the run. The flag must stay out of the hash, or
// every budget finding's id moved the day it became advisory.
func TestTheAdvisoryFlagIsNotInTheFindingID(t *testing.T) {
	plain := finding("a.md", "over the budget", "narrow the scope")
	marked := advisory(plain)
	if got, want := hashing.FindingID("G-Schema", "", marked.File, marked.Cause),
		hashing.FindingID("G-Schema", "", plain.File, plain.Cause); got != want {
		t.Fatalf("the id changed with the flag: %s against %s", got, want)
	}
}
