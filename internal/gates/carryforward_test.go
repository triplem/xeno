// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
)

// Two checks that differ in nothing but their provenance. Same gate, same file, same
// cause, and therefore the same finding id: the decision is offered to both, and only
// one of them may keep it.
func TestDecisionSurvivesForXenoAndNeverForAnExternalFinding(t *testing.T) {
	const gate, file, cause = "G-Example", "phases/02-design/output.md", "a cause"
	id := hashing.FindingID(gate, "", file, cause)
	decided := map[string]*model.DecisionOnFinding{
		id: {Type: "approved", By: "a person", At: "2026-09-22T12:00:00Z", Reason: "assessed"},
	}

	check := func(provenance string) model.Check {
		return model.Check{
			Gate:       gate,
			Result:     "fail",
			Provenance: provenance,
			Findings:   []model.Finding{{File: file, Cause: cause}},
		}
	}

	own := check("xeno")
	carryForward(&own, decided)
	if own.Findings[0].ID != id {
		t.Fatalf("finding id %s, want %s", own.Findings[0].ID, id)
	}
	if own.Findings[0].Decision == nil {
		t.Fatal("a decision on an unchanged finding of this runner did not survive the re-run")
	}

	foreign := check(ExternalProvenance)
	carryForward(&foreign, decided)
	if foreign.Findings[0].ID != id {
		t.Fatalf("an external finding needs an id of its own for this run: got %q", foreign.Findings[0].ID)
	}
	if foreign.Findings[0].Decision != nil {
		t.Fatal("a decision on an external finding survived the re-run; its cause comes from foreign code and is stable only by that code's promise")
	}
}

// A25 said the rule that an external finding keeps no decision rests on every check going
// through carryForward, and that nothing enforced the routing. These are the properties a
// verdict has to have whichever code produced it, which is what makes a second path into
// it safe to add.
func TestInvariantsHoldForWhatRunProduces(t *testing.T) {
	root := t.TempDir()
	if err := Invariants([]model.Check{}); err != nil {
		t.Fatalf("an empty verdict failed: %v", err)
	}
	// A real run of a phase that has findings: the fixture root has no artifacts at all,
	// so every gate that can report does.
	checks := Run(Ctx{Root: root, Key: "PROJ-1", Phase: model.Phases[0]}, nil)
	found := 0
	for _, ch := range checks {
		found += len(ch.Findings)
	}
	if found == 0 {
		t.Fatal("the fixture produced no findings, so the invariants were not exercised")
	}
	if err := Invariants(checks); err != nil {
		t.Fatalf("what Run produced failed its own invariants: %v", err)
	}
}

func TestInvariantsRejectWhatCarryForwardWouldNeverProduce(t *testing.T) {
	one := func(f model.Finding, provenance string) []model.Check {
		return []model.Check{{Gate: "G-Schema", Result: "fail", Provenance: provenance, Findings: []model.Finding{f}}}
	}
	id := hashing.FindingID("G-Schema", "", "a.md", "because")

	for _, tc := range []struct {
		name   string
		checks []model.Check
		want   string
	}{
		{"a finding with no id", one(model.Finding{File: "a.md", Cause: "because"}, "xeno"), "no id"},
		{"an id that is not derived", one(model.Finding{ID: "F-000000", File: "a.md", Cause: "because"}, "xeno"),
			"does not carry the id"},
		{"a decision on foreign code", one(model.Finding{ID: id, File: "a.md", Cause: "because",
			Decision: &model.DecisionOnFinding{Type: "approved", By: "m.example"}}, ExternalProvenance),
			"external and carries a decision"},
	} {
		err := Invariants(tc.checks)
		if err == nil {
			t.Errorf("%s was accepted", tc.name)
			continue
		}
		if !strings.Contains(err.Error(), tc.want) {
			t.Errorf("%s: error is %q, wanted it to name %q", tc.name, err, tc.want)
		}
	}

	// The same decision on a check of this runner is the case that has to keep working.
	withDecision := one(model.Finding{ID: id, File: "a.md", Cause: "because",
		Decision: &model.DecisionOnFinding{Type: "approved", By: "m.example"}}, "xeno")
	if err := Invariants(withDecision); err != nil {
		t.Fatalf("a decision on a finding of this runner was rejected: %v", err)
	}
}
