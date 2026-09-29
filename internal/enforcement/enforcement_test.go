// SPDX-License-Identifier: Apache-2.0

package enforcement

import (
	"testing"
	"time"
)

func declared() Declared {
	var d Declared
	d.RequiredPipeline, d.AllowBypass = true, false
	d.Approvals.Required, d.Approvals.NotByAuthor = 1, true
	return d
}

func stateOf(r Report, name string) (State, string) {
	for _, q := range r.Requirements {
		if q.Name == name {
			return q.State, q.Note
		}
	}
	return "", ""
}

// answered stands in for an adapter. Since A65 the states are the host's and this package
// judges none of them, so a test of Compare says what the host answered rather than what it
// was configured to do: the second of those is internal/host/github's subject.
func answered(s State, actual, note string, names ...string) []Requirement {
	if len(names) == 0 {
		names = []string{NameRequiredPipeline, NameAllowBypass,
			NameApprovalsRequired, NameApprovalsNotByAuthor}
	}
	out := make([]Requirement, 0, len(names))
	for _, n := range names {
		out = append(out, Requirement{Name: n, Actual: actual, State: s, Note: note})
	}
	return out
}

// The distinction the whole report rests on. A setting nobody made is somebody's
// oversight; a setting the host does not have is nobody's, and reporting it as one
// sends them looking for a checkbox that is not there.
func TestNotAvailableIsNotTheSameAsUnmet(t *testing.T) {
	absent := Compare("o/r", "main", declared(),
		answered(NotAvailable, "not available", "not on this plan"), time.Now())
	if s, _ := stateOf(absent, "required_pipeline"); s != NotAvailable {
		t.Fatalf("an unavailable setting is %q, want %q", s, NotAvailable)
	}

	unset := Compare("o/r", "main", declared(),
		answered(Unmet, "no status check is required", ""), time.Now())
	if s, _ := stateOf(unset, "required_pipeline"); s != Unmet {
		t.Fatalf("a setting nobody made is %q, want %q", s, Unmet)
	}
}

// A declared requirement the adapter said nothing about is not available rather than met.
// An absence is the one answer a host cannot be asked to give explicitly, so the domain has
// to read it, and reading it as met would turn a host that went silent into a pass.
func TestARequirementTheHostDidNotAnswerForIsNotAvailable(t *testing.T) {
	r := Compare("o/r", "main", declared(),
		answered(Met, "required", "", NameRequiredPipeline), time.Now())
	if s, _ := stateOf(r, "allow_bypass"); s != NotAvailable {
		t.Fatalf("an unanswered requirement is %q, want %q", s, NotAvailable)
	}
}

// WP10: a repository whose pipeline is not required is reported as such rather than
// passing quietly.
func TestAnUnprotectedBranchDoesNotPassQuietly(t *testing.T) {
	r := Compare("o/r", "main", declared(),
		answered(Unmet, "the branch is not protected", ""), time.Now())
	if s, _ := stateOf(r, "required_pipeline"); s != Unmet {
		t.Fatalf("state is %q, want %q", s, Unmet)
	}
	if r.Unmet() == 0 {
		t.Fatal("it passed quietly")
	}
}

// WP10: a requirement the host cannot express is reported once as waived rather than
// daily as unmet. Without this the report says the same thing forever and is ignored
// within a fortnight.
func TestWaivedTurnsAStandingComplaintIntoADecision(t *testing.T) {
	d := declared()
	// Four requirements, because the declaration carries four: the pipeline, the bypass,
	// the count of approvals and whether one may be the author's.
	before := Compare("o/r", "main", d,
		answered(NotAvailable, "not available", "not on this plan"), time.Now())
	if before.Unmet() != 4 {
		t.Fatalf("%d unmet before waiving, want 4", before.Unmet())
	}

	d.Waived = "not expressible on this plan, recorded 2026-09-25"
	after := Compare("o/r", "main", d,
		answered(NotAvailable, "not available", "not on this plan"), time.Now())
	if after.Unmet() != 0 {
		t.Fatalf("%d still unmet after waiving", after.Unmet())
	}
	s, note := stateOf(after, "required_pipeline")
	if s != Waived {
		t.Fatalf("state is %q, want %q", s, Waived)
	}
	if note != d.Waived {
		t.Fatalf("the note is %q, want the reason that was recorded", note)
	}
}

func TestAMetRequirementIsMet(t *testing.T) {
	r := Compare("o/r", "main", declared(), answered(Met, "required", ""), time.Now())
	if r.Unmet() != 0 {
		t.Fatalf("%d unmet against a host that requires everything", r.Unmet())
	}
	for _, q := range r.Requirements {
		if q.State != Met {
			t.Errorf("%s is %q", q.Name, q.State)
		}
	}
}

// The adapter's words reach the report unchanged. They are the only account of what the
// host actually said, so a domain that rephrased them would be inventing the evidence.
func TestTheHostsOwnWordsAreCarried(t *testing.T) {
	r := Compare("o/r", "main", declared(),
		answered(Met, "administrators are included", "and cannot be excluded"), time.Now())
	for _, q := range r.Requirements {
		if q.Actual != "administrators are included" {
			t.Errorf("%s reports actual %q, want the adapter's words", q.Name, q.Actual)
		}
		if q.Note != "and cannot be excluded" {
			t.Errorf("%s reports note %q, want the adapter's words", q.Name, q.Note)
		}
	}
}

// Not declared is not compared: a project that says nothing about it gets no line, which
// is what stateOf reports as the empty state. The filter is the domain's, which is why it
// is tested here and not against a host.
func TestAnUndeclaredRequirementGetsNoLine(t *testing.T) {
	d := declared()
	d.Approvals.NotByAuthor = false
	r := Compare("o/r", "main", d, answered(Met, "required", ""), time.Now())
	if s, _ := stateOf(r, "approvals.not_by_author"); s != "" {
		t.Fatalf("an undeclared requirement produced a line with state %q", s)
	}
}

// A name no section of the specification has is dropped rather than reported. There is
// nowhere in Report to say so that would not itself be a field the specification does not
// have, so the budget is held here and by a test over each adapter's names.
func TestANameTheDomainDoesNotKnowIsDropped(t *testing.T) {
	invented := []Requirement{{Name: "branch_naming", Actual: "enforced", State: Met}}
	r := Compare("o/r", "main", declared(), invented, time.Now())
	for _, q := range r.Requirements {
		if q.Name == "branch_naming" {
			t.Fatal("an adapter's invented requirement reached the report")
		}
	}
}

// merge_method is declared and not compared, which section 13 says outright. Reporting it
// as unchecked is what keeps a project from believing otherwise.
func TestMergeMethodIsReportedAsUnchecked(t *testing.T) {
	d := declared()
	d.MergeMethod = "no-squash"
	r := Compare("o/r", "main", d, answered(Met, "required", ""), time.Now())
	s, note := stateOf(r, "merge_method")
	if s != Unknown {
		t.Fatalf("state is %q, want %q", s, Unknown)
	}
	if note == "" {
		t.Fatal("nothing says why it is not compared")
	}
	// Unknown is neither met nor unmet, so it does not fail a run on its own.
	if before := Compare("o/r", "main", declared(), answered(Met, "required", ""), time.Now()); r.Unmet() != before.Unmet() {
		t.Fatalf("declaring merge_method changed the unmet count from %d to %d", before.Unmet(), r.Unmet())
	}
}

// Names is what an adapter builds its answers from, so it has to hold every name the
// report can carry. A name in the table and not in Names would be unreachable for an
// adapter and would read as a requirement no host can express.
func TestNamesHoldsEveryRequirementTheDomainKnows(t *testing.T) {
	in := make(map[string]bool, len(Names()))
	for _, n := range Names() {
		in[n] = true
	}
	for _, q := range requirements {
		if !in[q.name] {
			t.Errorf("%q is in the table and not in Names", q.name)
		}
	}
	if !in[NameMergeMethod] {
		t.Error("merge_method is reported and not in Names")
	}
}
