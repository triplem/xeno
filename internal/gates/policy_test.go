// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/rules"
)

const p5 = "05-review"

// reviewPhase writes a P5 artifact carrying a checklist, and a rule tree beside it. Both are
// the subject here, so both are built per test rather than shared.
func reviewPhase(t *testing.T, checklist string, tree map[string]string) Ctx {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, model.PhaseDir("PROJ-1", p5))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	out := "---\nintent: " + fixtureIntent + "\nphase: " + p5 + "\n" + checklist + "---\n\nbody\n"
	if err := os.WriteFile(filepath.Join(dir, "output.md"), []byte(out), 0o644); err != nil {
		t.Fatal(err)
	}
	for rel, body := range tree {
		writeRule(t, root, rel, body)
	}
	return Ctx{Root: root, Key: "PROJ-1", Phase: p5, QualifiedID: fixtureIntent}
}

func aReviewRule(id string) string {
	return "id: " + id + "\nversion: 1\nscope: org\nkind: review\napplies_to: [02-design]\n" +
		"statement: >\n  Something has to hold.\n"
}

func checkedRule(id, typ string) string {
	return "id: " + id + "\nversion: 1\nscope: org\nkind: checked\napplies_to: [" + p5 + "]\n" +
		"statement: >\n  Something has to hold.\ncheck:\n  type: " + typ + "\n"
}

// Every repository is in this state until a rule set exists, this one included, and the gate is
// quiet at every phase rather than at the four that carry no checklist.
func TestNoRuleTreeIsGreenForPolicy(t *testing.T) {
	for _, phase := range model.Phases {
		root := t.TempDir()
		c := policy(Ctx{Root: root, Key: "PROJ-1", Phase: phase, QualifiedID: fixtureIntent})
		if c.Result != "pass" {
			t.Fatalf("G-Policy is %s at %s with no rule tree, want pass: %s", c.Result, phase, causes(c))
		}
	}
}

// The criterion the gate rests on: without it, a checklist with nothing in it passes.
func TestAReviewRuleWithNoEntryIsRed(t *testing.T) {
	c := policy(reviewPhase(t, "", map[string]string{
		rules.ConfigDir + "/given/org/migration-note.yaml": aReviewRule("migration-note"),
	}))
	if c.Result != "fail" {
		t.Fatalf("G-Policy is %s, want fail", c.Result)
	}
	if !strings.Contains(causes(c), "review rule migration-note has no checklist entry") {
		t.Fatalf("findings %q, want the missing entry", causes(c))
	}
}

func TestAnAnsweredReviewRulePasses(t *testing.T) {
	c := policy(reviewPhase(t, "review_checklist:\n  - rule: migration-note\n    result: met\n",
		map[string]string{rules.ConfigDir + "/given/org/migration-note.yaml": aReviewRule("migration-note")}))
	if c.Result != "pass" {
		t.Fatalf("G-Policy is %s, want pass: %s", c.Result, causes(c))
	}
}

func TestWhatAnEntryOwes(t *testing.T) {
	for _, tc := range []struct {
		name, entry, want string
	}{
		{"no result", "  - rule: migration-note\n", "carries no result"},
		{"a result outside the three", "  - rule: migration-note\n    result: done\n", `has result "done"`},
		{"a deviation with no note", "  - rule: migration-note\n    result: deviation\n", "is deviation and carries no note"},
		{"a not-applicable with no note", "  - rule: migration-note\n    result: not-applicable\n", "is not-applicable and carries no note"},
		{"a deviation with a note", "  - rule: migration-note\n    result: deviation\n    note: the interface is internal\n", ""},
		{"a met needs no note", "  - rule: migration-note\n    result: met\n", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := policy(reviewPhase(t, "review_checklist:\n"+tc.entry, map[string]string{
				rules.ConfigDir + "/given/org/migration-note.yaml": aReviewRule("migration-note"),
			}))
			got := causes(c)
			if tc.want == "" {
				if c.Result != "pass" {
					t.Fatalf("G-Policy is %s, want pass: %s", c.Result, got)
				}
				return
			}
			if !strings.Contains(got, tc.want) {
				t.Fatalf("findings %q, want one containing %q", got, tc.want)
			}
		})
	}
}

// Section 12: a lens entry carries source: lens and no rule id, and cannot add to or subtract
// from the set the gate counts. So four of them leave a missing answer missing.
func TestLensEntriesDoNotAnswerARule(t *testing.T) {
	checklist := "review_checklist:\n"
	for _, lens := range []string{"security", "privacy", "operations", "architecture"} {
		checklist += "  - result: deviation\n    note: the " + lens + " lens had something to say\n    source: lens\n"
	}
	c := policy(reviewPhase(t, checklist, map[string]string{
		rules.ConfigDir + "/given/org/migration-note.yaml": aReviewRule("migration-note"),
	}))
	if !strings.Contains(causes(c), "review rule migration-note has no checklist entry") {
		t.Fatalf("findings %q, want the rule still unanswered", causes(c))
	}
}

// The gate keys on the missing rule id rather than on source, so a lens that omitted its own
// label still cannot answer a rule, and it still owes a result.
func TestALensEntryStillOwesAResult(t *testing.T) {
	c := policy(reviewPhase(t, "review_checklist:\n  - source: lens\n    note: something\n", nil))
	if !strings.Contains(causes(c), "carries no result") {
		t.Fatalf("findings %q, want the missing result", causes(c))
	}
	if !strings.Contains(causes(c), "from lens") {
		t.Fatalf("findings %q, want the entry named by its source", causes(c))
	}
}

func TestAnEntryForARuleOutsideTheSetIsRed(t *testing.T) {
	c := policy(reviewPhase(t, "review_checklist:\n  - rule: a-rule-nobody-has\n    result: met\n", nil))
	if !strings.Contains(causes(c), "which is no review rule of the effective set") {
		t.Fatalf("findings %q, want the unknown rule", causes(c))
	}
}

// A rule in force that nothing evaluates is the silently green verdict section 16 catalogues,
// so a type nothing implements is reported instead. Since #162 the five of section 9 are
// implemented, which leaves this as the case that keeps the registry a budget: a project naming
// a type of its own gets the finding, because project-defined predicates are outside v1.
func TestACheckedRuleWithNoImplementationIsRed(t *testing.T) {
	c := policy(reviewPhase(t, "", map[string]string{
		rules.ConfigDir + "/given/org/house.yaml": checkedRule("house-style", "house-linter-says-so"),
	}))
	if c.Result != "fail" {
		t.Fatalf("G-Policy is %s, want fail", c.Result)
	}
	if !strings.Contains(causes(c), `no implementation exists for predicate type "house-linter-says-so"`) {
		t.Fatalf("findings %q, want the unimplemented type", causes(c))
	}
}

// The five names of section 9, each one evaluated rather than reported.
func TestTheFiveTypesOfSectionNineAreRegistered(t *testing.T) {
	for _, name := range []string{"section-implies-section", "commit-message", "commit-trailer",
		"commit-signature", "approver-not-author"} {
		if _, ok := predicates[name]; !ok {
			t.Errorf("predicate type %s is not registered", name)
		}
	}
	if len(predicates) != 5 {
		t.Errorf("the registry holds %d types, want the five section 9 names", len(predicates))
	}
}

// Only where it applies. The same rule at a phase it does not name is nothing to say.
func TestACheckedRuleIsReadAtThePhasesItNames(t *testing.T) {
	root := t.TempDir()
	writeRule(t, root, rules.ConfigDir+"/given/org/interfaces.yaml", checkedRule("interfaces", "section-implies-section"))
	c := policy(Ctx{Root: root, Key: "PROJ-1", Phase: model.Phases[0], QualifiedID: fixtureIntent})
	if c.Result != "pass" {
		t.Fatalf("G-Policy is %s at 00-intake for a rule applying to 05-review, want pass: %s", c.Result, causes(c))
	}
}

// A phase that is quiet because nothing applied reports a pass. not-implemented is a statement
// about the runner, and this gate is implemented.
func TestGPolicyIsNoLongerReportedAsNotImplemented(t *testing.T) {
	root, _ := corpus(t)
	for _, s := range table {
		if s.id != "G-Policy" {
			continue
		}
		if got := s.fn(ctxFor(root)).Result; got == "not-implemented" {
			t.Fatal("G-Policy still reports not-implemented")
		}
		return
	}
	t.Fatal("G-Policy is not in the table")
}

// A tree that does not resolve is G-Rules's finding; repeating it here would report one broken
// file twice in one verdict.
func TestAnUnresolvableTreeIsNotReportedTwice(t *testing.T) {
	c := policy(reviewPhase(t, "", map[string]string{
		// scope against the path, which G-Rules reports
		rules.ConfigDir + "/given/org/misplaced.yaml": strings.Replace(aReviewRule("misplaced"), "scope: org", "scope: project", 1),
	}))
	if c.Result != "pass" {
		t.Fatalf("G-Policy is %s over a tree G-Rules rejects, want pass: %s", c.Result, causes(c))
	}
}
