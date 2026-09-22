// SPDX-License-Identifier: Apache-2.0

package gates

import (
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
