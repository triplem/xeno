// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
)

// nexts is the next steps of a verdict's findings. causes() in schema_test.go prints the file
// and the cause; several criteria here are about what a finding tells somebody to do.
func nexts(c model.Check) string {
	var b strings.Builder
	for _, f := range c.Findings {
		b.WriteString(f.Next + "\n")
	}
	return b.String()
}

// shippedTemplate gives a fixture the design template, because section-implies-section checks a
// rule's section names against the template the phase renders from, and a repository without a
// vendored plugin cannot answer that at all.
func shippedTemplate(t *testing.T, root string) {
	t.Helper()
	dir := filepath.Join(root, ".xeno/plugin/templates/design")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		"template.yaml": "id: design\nversion: 1.0.0\nphase: 02-design\nsections:\n" +
			"  - { id: decisions, required: true }\n  - { id: alternatives, required: true }\n" +
			"  - { id: impact, required: true }\n",
		"strings.en.yaml": "language: en\ntitle: Design\nheadings:\n  decisions: Decisions\n" +
			"  alternatives: Alternatives\n  impact: Impact\n",
	}
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

// ---- section-implies-section, the one type that reads an artifact

// designPhase writes a 02-design artifact with the sections given, which is the phase whose
// template carries the sections section 9's example names.
func designPhase(t *testing.T, sections map[string]string, tree map[string]string) Ctx {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, model.PhaseDir("PROJ-1", "02-design"))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	shippedTemplate(t, root)
	body := "---\nintent: " + fixtureIntent + "\nphase: 02-design\n---\n"
	for id, text := range sections {
		body += "\n<!-- xeno:section:" + id + " -->\n## " + id + "\n\n" + text + "\n"
	}
	if err := os.WriteFile(filepath.Join(dir, "output.md"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	for rel, b := range tree {
		writeRule(t, root, rel, b)
	}
	return Ctx{Root: root, Key: "PROJ-1", Phase: "02-design", QualifiedID: fixtureIntent}
}

// impliesRule is section 9's own example, which is the rule this type exists for.
func impliesRule(when, then string) string {
	return "id: interfaces-need-a-note\nversion: 1\nscope: org\nkind: checked\napplies_to: [02-design]\n" +
		"statement: >\n  A change to a published interface must come with a migration note.\n" +
		"check:\n  type: section-implies-section\n" +
		"  when: { section: " + when + ", non_empty: true }\n" +
		"  then: { section: " + then + ", non_empty: true }\n"
}

func TestSectionImpliesSection(t *testing.T) {
	for _, tc := range []struct {
		name     string
		sections map[string]string
		want     string
	}{
		{"antecedent filled, consequent filled", map[string]string{
			"decisions": "the interface changed", "alternatives": "a migration note"}, ""},
		{"antecedent filled, consequent empty", map[string]string{
			"decisions": "the interface changed"}, "decisions is filled and alternatives is empty"},
		{"antecedent empty, consequent empty", map[string]string{"impact": "none"}, ""},
		{"antecedent whitespace only, consequent empty", map[string]string{
			"decisions": "   \n  ", "impact": "none"}, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := policy(designPhase(t, tc.sections, map[string]string{
				".xeno/config/rules/given/org/i.yaml": impliesRule("decisions", "alternatives"),
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

// A section the phase's template does not have is a configuration error in the rule, not an
// empty section: an empty one would make the rule vacuously green for ever.
func TestASectionTheTemplateDoesNotHaveIsAConfigurationError(t *testing.T) {
	c := policy(designPhase(t, map[string]string{"decisions": "something"}, map[string]string{
		".xeno/config/rules/given/org/i.yaml": impliesRule("decisions", "a-section-nobody-has"),
	}))
	if !strings.Contains(causes(c), `names section "a-section-nobody-has"`) {
		t.Fatalf("findings %q, want the unknown section", causes(c))
	}
}

func TestACheckWithNoSectionsIsAConfigurationError(t *testing.T) {
	bare := "id: bare\nversion: 1\nscope: org\nkind: checked\napplies_to: [02-design]\n" +
		"statement: >\n  Something.\ncheck:\n  type: section-implies-section\n"
	c := policy(designPhase(t, nil, map[string]string{".xeno/config/rules/given/org/b.yaml": bare}))
	if !strings.Contains(causes(c), "names no section in when or then") {
		t.Fatalf("findings %q, want the missing sections", causes(c))
	}
}

// ---- the four that read a commit range

func gitRepo(t *testing.T, root string) {
	t.Helper()
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not on the path")
	}
	for _, args := range [][]string{
		{"init", "--initial-branch=main"},
		{"config", "user.name", "A Committer"},
		{"config", "user.email", "committer@example.test"},
		{"config", "commit.gpgsign", "false"},
	} {
		cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %s: %v\n%s", strings.Join(args, " "), err, out)
		}
	}
}

func gitCommit(t *testing.T, root, message string) string {
	t.Helper()
	safe := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			return r
		}
		return '-'
	}, strings.ToLower(message))
	if err := os.WriteFile(filepath.Join(root, safe+".txt"), []byte(message), 0o644); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{{"add", "."}, {"commit", "-m", message}} {
		cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %s: %v\n%s", strings.Join(args, " "), err, out)
		}
	}
	cmd := exec.Command("git", "-C", root, "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(string(out))
}

// commitPhase is a repository that is also an intent: a rule tree, a phase, and a history.
func commitPhase(t *testing.T, tree map[string]string, messages ...string) (Ctx, string) {
	t.Helper()
	root := t.TempDir()
	gitRepo(t, root)
	base := gitCommit(t, root, "chore: the base")
	for _, m := range messages {
		gitCommit(t, root, m)
	}
	dir := filepath.Join(root, model.PhaseDir("PROJ-1", "05-review"))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	out := "---\nintent: " + fixtureIntent + "\nphase: 05-review\n---\n\nbody\n"
	if err := os.WriteFile(filepath.Join(dir, "output.md"), []byte(out), 0o644); err != nil {
		t.Fatal(err)
	}
	for rel, b := range tree {
		writeRule(t, root, rel, b)
	}
	return Ctx{Root: root, Key: "PROJ-1", Phase: "05-review", QualifiedID: fixtureIntent,
		Base: base, Head: "HEAD"}, base
}

func commitRule(id, typ, params string) string {
	return "id: " + id + "\nversion: 1\nscope: org\nkind: checked\napplies_to: [05-review]\n" +
		"statement: >\n  Something about the commits.\ncheck:\n  type: " + typ + "\n" +
		"  range: change-under-review\n" + params
}

const messageRuleFile = ".xeno/config/rules/given/org/subjects.yaml"

func TestCommitMessageJudgesEverySubject(t *testing.T) {
	tree := map[string]string{messageRuleFile: commitRule("subjects", "commit-message",
		"  pattern: conventional-commits\n")}

	c, _ := commitPhase(t, tree, "feat: one", "fix(parser): two")
	if got := policy(c); got.Result != "pass" {
		t.Fatalf("conforming subjects are %s: %s", got.Result, causes(got))
	}

	c, _ = commitPhase(t, tree, "feat: one", "made some changes")
	got := policy(c)
	if got.Result != "fail" {
		t.Fatalf("a non-conforming subject is %s, want fail", got.Result)
	}
	if !strings.Contains(causes(got), `does not match pattern "conventional-commits"`) {
		t.Fatalf("findings %q, want the pattern named", causes(got))
	}
	if !strings.Contains(causes(got), "made some changes") {
		t.Fatalf("findings %q, want the offending subject named", causes(got))
	}
}

func TestTheSecondShippedPatternWantsAReference(t *testing.T) {
	tree := map[string]string{messageRuleFile: commitRule("subjects", "commit-message",
		"  pattern: conventional-commits-with-issue\n")}

	c, _ := commitPhase(t, tree, "feat: the thing (#162)", "fix(rules): the other (!41)")
	if got := policy(c); got.Result != "pass" {
		t.Fatalf("subjects carrying a reference are %s: %s", got.Result, causes(got))
	}

	c, _ = commitPhase(t, tree, "feat: the thing")
	if got := policy(c); got.Result != "fail" {
		t.Fatalf("a subject with no reference is %s, want fail", got.Result)
	}
}

func TestAnUnknownPatternNameSaysWhatIsShipped(t *testing.T) {
	c, _ := commitPhase(t, map[string]string{messageRuleFile: commitRule("subjects", "commit-message",
		"  pattern: house-style\n")}, "feat: one")
	got := causes(policy(c))
	if !strings.Contains(got, `names pattern "house-style", which is not shipped`) {
		t.Fatalf("findings %q, want the unknown pattern", got)
	}
	if !strings.Contains(nexts(policy(c)), "conventional-commits-with-issue") {
		t.Fatalf("next steps %q, want the shipped names listed", nexts(policy(c)))
	}
}

// Without the exemption a merge subject is judged like any other, which is what section 9's
// example makes explicit by carrying the field.
func TestMergeCommitsAreExemptWhereTheRuleSaysSo(t *testing.T) {
	build := func(t *testing.T, params string) model.Check {
		t.Helper()
		root := t.TempDir()
		gitRepo(t, root)
		base := gitCommit(t, root, "chore: the base")
		for _, args := range [][]string{{"checkout", "-b", "side"}} {
			exec.Command("git", append([]string{"-C", root}, args...)...).Run()
		}
		gitCommit(t, root, "feat: on the side")
		exec.Command("git", "-C", root, "checkout", "main").Run()
		gitCommit(t, root, "feat: on main")
		if out, err := exec.Command("git", "-C", root, "merge", "--no-ff", "-m",
			"Merge branch 'side'", "side").CombinedOutput(); err != nil {
			t.Fatalf("merge: %v\n%s", err, out)
		}
		dir := filepath.Join(root, model.PhaseDir("PROJ-1", "05-review"))
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "output.md"),
			[]byte("---\nintent: "+fixtureIntent+"\nphase: 05-review\n---\n\nbody\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		writeRule(t, root, messageRuleFile, commitRule("subjects", "commit-message",
			"  pattern: conventional-commits\n"+params))
		return policy(Ctx{Root: root, Key: "PROJ-1", Phase: "05-review", QualifiedID: fixtureIntent,
			Base: base, Head: "HEAD"})
	}

	if got := build(t, "  exempt: [merge-commits]\n"); got.Result != "pass" {
		t.Fatalf("with the exemption the verdict is %s: %s", got.Result, causes(got))
	}
	got := build(t, "")
	if got.Result != "fail" || !strings.Contains(causes(got), "Merge branch") {
		t.Fatalf("without the exemption the merge subject was not judged: %s %s", got.Result, causes(got))
	}
}

func TestCommitTrailerRequiresItOnEveryCommit(t *testing.T) {
	tree := map[string]string{".xeno/config/rules/given/org/trailer.yaml": commitRule("trailer",
		"commit-trailer", "  trailer: Xeno-Intent\n")}

	c, _ := commitPhase(t, tree, "feat: one\n\nXeno-Intent: XENO-0219")
	if got := policy(c); got.Result != "pass" {
		t.Fatalf("a commit carrying the trailer is %s: %s", got.Result, causes(got))
	}

	c, _ = commitPhase(t, tree, "feat: one\n\nXeno-Intent: XENO-0219", "fix: two")
	got := policy(c)
	if got.Result != "fail" || !strings.Contains(causes(got), "carries no Xeno-Intent trailer") {
		t.Fatalf("a commit without the trailer is %s: %s", got.Result, causes(got))
	}
}

func TestCommitSignatureRequiresAVerifiedSignature(t *testing.T) {
	c, _ := commitPhase(t, map[string]string{".xeno/config/rules/given/org/sig.yaml": commitRule("sig",
		"commit-signature", "")}, "feat: unsigned")
	got := policy(c)
	if got.Result != "fail" || !strings.Contains(causes(got), "carries no signature that verifies") {
		t.Fatalf("an unsigned commit is %s: %s", got.Result, causes(got))
	}
}

// Both halves are already recorded, which is what makes the separation of duties cheap to
// check: the decision carries by, the range carries its authors.
func TestApproverNotAuthor(t *testing.T) {
	rule := map[string]string{".xeno/config/rules/given/org/approver.yaml": commitRule("approver",
		"approver-not-author", "")}

	withDecision := func(t *testing.T, by string) model.Check {
		t.Helper()
		c, _ := commitPhase(t, rule, "feat: one")
		gate := "intent: " + fixtureIntent + "\nphase: 05-review\ncreated: 2026-10-01T00:00:00Z\n" +
			"schema_version: \"1.0\"\nrunner_version: 0.1.0-dev\nplugin_version: 0.1.0-dev\n" +
			"status: approved\nrun_at: 2026-10-01T00:00:00Z\nartifacts_hash: deadbeef\nchecks:\n" +
			"  - gate: G-Schema\n    result: fail\n    provenance: xeno\n    findings:\n" +
			"      - id: F-abc123\n        file: output.md\n        cause: something\n        next: fix it\n" +
			"        decision:\n          type: approved\n          by: " + by + "\n" +
			"          at: 2026-10-01T00:00:00Z\n          against: deadbeef\n          reason: looked at it\n"
		p := filepath.Join(c.Root, model.PhaseDir("PROJ-1", "05-review"), "gate.yaml")
		if err := os.WriteFile(p, []byte(gate), 0o644); err != nil {
			t.Fatal(err)
		}
		return policy(c)
	}

	got := withDecision(t, "committer@example.test")
	if got.Result != "fail" {
		t.Fatalf("an approval by an author of the range is %s, want fail", got.Result)
	}
	if !strings.Contains(causes(got), "F-abc123 was approved by") {
		t.Fatalf("findings %q, want the finding and the approver named", causes(got))
	}

	if got := withDecision(t, "somebody.else@example.test"); got.Result != "pass" {
		t.Fatalf("an approval by somebody else is %s: %s", got.Result, causes(got))
	}

	// The name, not only the email, because a decision carries free text.
	if got := withDecision(t, "A Committer"); got.Result != "fail" {
		t.Fatalf("an approval by the author's name is %s, want fail", got.Result)
	}
}

// No decision is green: a rule about who may release a finding is not a rule requiring one.
func TestApproverNotAuthorWithNoDecisionsPasses(t *testing.T) {
	c, _ := commitPhase(t, map[string]string{".xeno/config/rules/given/org/approver.yaml": commitRule("approver",
		"approver-not-author", "")}, "feat: one")
	if got := policy(c); got.Result != "pass" {
		t.Fatalf("no decisions is %s: %s", got.Result, causes(got))
	}
}

// The case WP4's done-when names as its own criterion: the range is an input of the run, so a
// missing one is a finding and never an empty range.
func TestARuleThatNeedsARangeAndDidNotGetOneIsRed(t *testing.T) {
	for _, typ := range []string{"commit-message", "commit-trailer", "commit-signature", "approver-not-author"} {
		t.Run(typ, func(t *testing.T) {
			root := t.TempDir()
			dir := filepath.Join(root, model.PhaseDir("PROJ-1", "05-review"))
			if err := os.MkdirAll(dir, 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(dir, "output.md"),
				[]byte("---\nintent: "+fixtureIntent+"\nphase: 05-review\n---\n\nbody\n"), 0o644); err != nil {
				t.Fatal(err)
			}
			writeRule(t, root, ".xeno/config/rules/given/org/r.yaml",
				commitRule("r", typ, "  pattern: conventional-commits\n  trailer: Xeno-Intent\n"))
			c := policy(Ctx{Root: root, Key: "PROJ-1", Phase: "05-review", QualifiedID: fixtureIntent})
			if c.Result != "fail" {
				t.Fatalf("%s with no range is %s, want fail", typ, c.Result)
			}
			if !strings.Contains(causes(c), "the commit range and it could not be read") {
				t.Fatalf("findings %q, want the range named", causes(c))
			}
			if !strings.Contains(nexts(c), "never inferred") {
				t.Fatalf("next steps %q, want them to say the range is passed in", nexts(c))
			}
		})
	}
}

// Five rules naming a commit type read one git log, which is the shape #160's gap asked for.
func TestTheRangeIsReadOncePerGateRun(t *testing.T) {
	tree := map[string]string{}
	for _, n := range []string{"a", "b", "c", "d", "e"} {
		tree[".xeno/config/rules/given/org/"+n+".yaml"] = commitRule("subjects-"+n, "commit-message",
			"  pattern: conventional-commits\n")
	}
	c, _ := commitPhase(t, tree, "made some changes")
	got := policy(c)
	if n := len(got.Findings); n != 5 {
		t.Fatalf("%d findings over five rules and one bad subject, want 5: %s", n, causes(got))
	}
}
