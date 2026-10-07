// SPDX-License-Identifier: Apache-2.0

package learning

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/rules"
)

// trail writes a repository holding the records given, keyed by the path each one takes under
// the intents directory, and returns its root. The records are written as text rather than
// through the runner, because what is under test is what the route makes of a record and not
// what wrote it; a trail from the runner would also have to carry a project.yaml and a plugin
// to get one written.
func trail(t *testing.T, records map[string]string, files ...string) string {
	t.Helper()
	root := t.TempDir()
	for rel, body := range records {
		p := filepath.Join(root, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	for _, rel := range files {
		p := filepath.Join(root, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte("a file the project reads\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

// record is one learning.yaml with the header a phase writes and the entries given.
func record(entries ...string) string {
	s := "intent: github.com/triplem/xeno#1\ncreated: \"2026-10-07T00:00:00Z\"\n" +
		"schema_version: \"1.0\"\nrunner_version: dev\nplugin_version: 0.0.0-dev\nlearnings:\n"
	return s + strings.Join(entries, "")
}

func entry(category, target string) string {
	return entryWith(category, target, "The proposal, which is a sentence somebody wrote about their own phase.")
}

func entryWith(category, target, proposal string) string {
	return "  - category: " + category + "\n    observation: what was noticed\n" +
		"    proposal: " + proposal + "\n    target: " + target + "\n"
}

// Every proposal read takes exactly one disposition, and the three counts add up to the number
// of entries. That is the whole of #289's complaint turned into an assertion: a route that
// carried the resolvable proposals and dropped the rest would leave nobody able to say which of
// them had been read, so an item in none of the three is a dropped record.
func TestEveryProposalTakesExactlyOneDisposition(t *testing.T) {
	root := trail(t, map[string]string{
		".xeno/intents/XENO-0200/phases/00-intake/learning.yaml": record(
			entry("project-convention", ProjectRules+"/a-rule.yaml"),
			entry("project-convention", "CLAUDE.md"),
			entry("context-rule", "docs/process-definition.md section 4, the verification phase"),
			entry("template", ".xeno/plugin/templates"),
		),
		".xeno/intents/XENO-0201/phases/00-intake/learning.yaml": "no_finding: true\n",
	}, "CLAUDE.md", ".xeno/plugin/templates/requirements/output.md")

	b := Plan(root, []string{"XENO-0200", "XENO-0201"})
	if b.Records != 2 || b.NoFinding != 1 {
		t.Fatalf("read %d records, %d of them empty; want 2 and 1", b.Records, b.NoFinding)
	}
	if n := len(b.Items); n != 4 {
		t.Fatalf("read %d proposals, want 4", n)
	}
	sum := b.Count(Carried) + b.Count(ForWording) + b.Count(Unresolved)
	if sum != len(b.Items) {
		t.Errorf("%d proposals and %d dispositions, so %d were dropped", len(b.Items), sum, len(b.Items)-sum)
	}
	want := map[Disposition]int{Carried: 1, ForWording: 1, Unresolved: 2}
	for d, n := range want {
		if got := b.Count(d); got != n {
			t.Errorf("%s: %d, want %d", d, got, n)
		}
	}
	// Every item is named in the description, by the record it came from and its position in
	// it, which is what a reader needs to find the entry again.
	for _, it := range b.Items {
		if !strings.Contains(b.Description, it.Path) {
			t.Errorf("the description does not name %s", it.Path)
		}
	}
}

// A target under the project level of the learned tree becomes a rule file, and the rules
// package reads the generated file back without a problem. The route's claim is that a reviewer
// is handed something the rule set will have, and the place to find out otherwise is here.
func TestACarriedProposalBecomesARuleTheRuleSetAccepts(t *testing.T) {
	root := trail(t, map[string]string{
		".xeno/intents/XENO-0200/phases/02-design/learning.yaml": record(
			entry("project-convention", ProjectRules+"/a-figure-names-its-population.yaml")),
	})
	b := Plan(root, []string{"XENO-0200"})
	out := filepath.Join(t.TempDir(), "bundle")
	problems, err := b.Write(out)
	if err != nil {
		t.Fatal(err)
	}
	if len(problems) > 0 {
		t.Fatalf("the generated tree does not resolve: %+v", problems)
	}
	loaded, _ := rules.Load(out)
	if len(loaded) != 1 {
		t.Fatalf("the generated tree holds %d rules, want 1", len(loaded))
	}
	r := loaded[0]
	if r.ID != "a-figure-names-its-population" || r.Scope != "project" || r.Kind != rules.Review {
		t.Errorf("generated %s, scope %s, kind %s", r.ID, r.Scope, r.Kind)
	}
	if r.Check != nil {
		t.Error("a proposal in prose produced a predicate, which nobody proposed")
	}
	if len(r.AppliesTo) != 1 || r.AppliesTo[0] != "02-design" {
		t.Errorf("applies_to is %v, want the phase the record was written in", r.AppliesTo)
	}
	if eff, probs := rules.Effective(loaded); len(probs) > 0 || len(eff) != 1 {
		t.Errorf("the effective set is %d rules with %d problems", len(eff), len(probs))
	}
	if _, err := os.Stat(filepath.Join(out, DescriptionFile)); err != nil {
		t.Errorf("no description beside the rules: %v", err)
	}
	if !strings.Contains(b.Patch, "new file mode 100644") || !strings.Contains(b.Patch, "+id: a-figure-names-its-population") {
		t.Errorf("the patch does not create the rule file:\n%s", b.Patch)
	}
}

// A target that reads as prose is listed with the reason rather than dropped, and nothing is
// generated for it. Section 10 says target is the file the merge request would change, so a
// sentence naming a section of a document is a finding about the record.
func TestAProseTargetIsListedWithItsReason(t *testing.T) {
	root := trail(t, map[string]string{
		".xeno/intents/XENO-0200/phases/00-intake/learning.yaml": record(
			entry("context-rule", "docs/process-definition.md section 4, the verification phase")),
	})
	b := Plan(root, []string{"XENO-0200"})
	if len(b.Files) != 0 {
		t.Fatalf("generated %d files from a target no file follows from", len(b.Files))
	}
	it := b.Items[0]
	if it.Disposition != Unresolved || !strings.Contains(it.Reason, "prose") {
		t.Fatalf("disposition %s, reason %q", it.Disposition, it.Reason)
	}
	if !strings.Contains(b.Description, it.Reason) {
		t.Error("the description does not carry the reason, so a reader cannot tell why it was not carried")
	}
}

// Section 10's exception is for project-convention and for nothing else. The same target under
// another category is unresolved, because a latitude for all four would make target mean "where
// this belongs" again, which is what the prose targets of #289 already were.
func TestOnlyAProjectConventionMayNameAFileOutsideTheRuleTree(t *testing.T) {
	for _, c := range []string{"project-convention", "context-rule", "template", "prompt"} {
		root := trail(t, map[string]string{
			".xeno/intents/XENO-0200/phases/00-intake/learning.yaml": record(entry(c, "CLAUDE.md")),
		}, "CLAUDE.md")
		b := Plan(root, []string{"XENO-0200"})
		got := b.Items[0].Disposition
		want := Unresolved
		if c == "project-convention" {
			want = ForWording
		}
		if got != want {
			t.Errorf("%s naming CLAUDE.md is %s, want %s (%s)", c, got, want, b.Items[0].Reason)
		}
	}
}

// Two proposals naming one rule file are carried into it together. A convention noticed twice is
// the ordinary case, and the second record overwriting the first would be a proposal reported as
// carried and then lost.
func TestTwoProposalsForOneRuleFileAreCarriedIntoIt(t *testing.T) {
	target := ProjectRules + "/a-count-names-its-population.yaml"
	root := trail(t, map[string]string{
		".xeno/intents/XENO-0200/phases/00-intake/learning.yaml": record(
			entryWith("project-convention", target, "The first proposal.")),
		".xeno/intents/XENO-0201/phases/04-verification/learning.yaml": record(
			entryWith("context-rule", target, "The second proposal.")),
	})
	b := Plan(root, []string{"XENO-0200", "XENO-0201"})
	if len(b.Files) != 1 {
		t.Fatalf("generated %d files for one target", len(b.Files))
	}
	c := b.Files[0].Content
	for _, want := range []string{
		"applies_to: [00-intake, 04-verification]",
		"XENO-0200/phases/00-intake/learning.yaml entry 1",
		"XENO-0201/phases/04-verification/learning.yaml entry 1",
	} {
		if !strings.Contains(c, want) {
			t.Errorf("the generated file does not carry %q:\n%s", want, c)
		}
	}
	if !strings.Contains(c, "The first proposal. The second proposal.") {
		t.Errorf("the statement does not carry both proposals:\n%s", c)
	}
}

// The record an intent writes when it closes belongs to no phase, so the rule is proposed for
// the whole sequence and a reviewer narrows it. Section 9 requires applies_to and will not take
// it empty, and guessing a phase for a record that names none would be worse than proposing it
// everywhere: a rule applying too widely is answered in the checklist, where somebody says so.
func TestARecordWithNoPhaseProposesTheRuleForEveryPhase(t *testing.T) {
	root := trail(t, map[string]string{
		".xeno/intents/XENO-0200/learning.yaml": record(
			entry("project-convention", ProjectRules+"/an-abandoned-intent-says-why.yaml")),
	})
	b := Plan(root, []string{"XENO-0200"})
	if len(b.Files) != 1 {
		t.Fatalf("generated %d files", len(b.Files))
	}
	if !strings.Contains(b.Files[0].Content,
		"applies_to: [00-intake, 01-requirements, 02-design, 03-implementation, 04-verification, 05-review]") {
		t.Errorf("applies_to narrows a record that named no phase:\n%s", b.Files[0].Content)
	}
}

// A rule already in the tree is proposed at one version higher. The counter counts changes to
// the file, so a proposal against an existing rule is a change to it and not a first version of
// it, and the tree is read for the number and never written.
func TestARuleAlreadyInTheTreeIsProposedAtOneVersionHigher(t *testing.T) {
	target := ProjectRules + "/already-here.yaml"
	root := trail(t, map[string]string{
		".xeno/intents/XENO-0200/phases/00-intake/learning.yaml": record(entry("project-convention", target)),
		target: "id: already-here\nversion: 4\nscope: project\nkind: review\n" +
			"applies_to: [00-intake]\nstatement: >\n  Something already agreed.\n",
	})
	b := Plan(root, []string{"XENO-0200"})
	if !strings.Contains(b.Files[0].Content, "version: 5") {
		t.Errorf("the proposal does not count the change:\n%s", b.Files[0].Content)
	}
	on, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(target)))
	if err != nil || !strings.Contains(string(on), "version: 4") {
		t.Errorf("the rule in the tree was touched: %v %s", err, on)
	}
}

// The route refuses a destination inside the trail it read. A learning takes effect after
// review and never on being noticed, and an output directory pointed at the trail would be
// exactly the edit section 10 forbids, with a flag in front of it.
func TestTheRouteRefusesToWriteIntoTheTrailItRead(t *testing.T) {
	root := trail(t, map[string]string{
		".xeno/intents/XENO-0200/phases/00-intake/learning.yaml": record(
			entry("project-convention", ProjectRules+"/a-rule.yaml")),
	})
	b := Plan(root, []string{"XENO-0200"})
	for _, out := range []string{
		filepath.Join(root, ".xeno"),
		filepath.Join(root, ProjectRules),
		filepath.Join(root, ".xeno/intents/XENO-0200"),
		"",
	} {
		if _, err := b.Write(out); err == nil {
			t.Errorf("writing to %q was allowed", out)
		}
	}
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(ProjectRules))); !os.IsNotExist(err) {
		t.Errorf("the rule tree of the repository was created after all: %v", err)
	}
}
