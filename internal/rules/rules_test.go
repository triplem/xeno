// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// tree writes a rule file at a path relative to a fresh repository root. Every test builds the
// tree it is about, because the thing under test is what a layout means and a shared fixture
// would hide which part of it each case depends on.
func tree(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for rel, body := range files {
		p := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

// rule is a well formed rule file, so that a test changing one field changes only that field.
func rule(id, scope, kind string, extra ...string) string {
	b := &strings.Builder{}
	b.WriteString("id: " + id + "\nversion: 1\nscope: " + scope + "\nkind: " + kind + "\n")
	b.WriteString("applies_to: [02-design]\nstatement: >\n  Something has to hold.\n")
	if kind == Checked {
		b.WriteString("check:\n  type: section-implies-section\n  when: { section: interfaces, non_empty: true }\n  then: { section: migration-notes, non_empty: true }\n")
	}
	for _, e := range extra {
		b.WriteString(e + "\n")
	}
	return b.String()
}

func TestNoRuleTreeIsNoRulesAndNoProblems(t *testing.T) {
	read, problems := Load(t.TempDir())
	if len(read) != 0 || len(problems) != 0 {
		t.Fatalf("read %d rules and %d problems, want none", len(read), len(problems))
	}
}

// An empty effective set still has a hash, because section 5 requires the field. What it says
// is that a tree was resolved and came out empty, which is not the same claim as a tree
// nobody looked at (A66).
func TestAnEmptySetHashesTheEmptyRendering(t *testing.T) {
	if got := Render(nil); got != "" {
		t.Fatalf("rendering of nothing is %q, want empty", got)
	}
	if got := Hash(nil); len(got) != 64 {
		t.Fatalf("hash of the empty rendering is %q, want sixty four hex characters", got)
	}
}

func TestPrecedenceIsSpecificityThenGivenOverLearned(t *testing.T) {
	root := tree(t, map[string]string{
		PluginDir + "/given/builtin/r.yaml":  rule("r", "builtin", Review),
		ConfigDir + "/given/provider/r.yaml": rule("r", "provider", Review),
		ConfigDir + "/given/org/r.yaml":      rule("r", "org", Review),
		ConfigDir + "/learned/org/r.yaml":    rule("r", "org", Review, "abstract: a transferable summary"),
	})
	read, problems := Load(root)
	if len(problems) != 0 {
		t.Fatalf("problems: %v", problems)
	}
	eff, collisions := Effective(read)
	if len(collisions) != 0 {
		t.Fatalf("collisions: %v", collisions)
	}
	if len(eff) != 1 {
		t.Fatalf("effective set has %d rules, want 1", len(eff))
	}
	if eff[0].Level != "org" || eff[0].Origin != Given {
		t.Fatalf("resolved to %s/%s, want given/org", eff[0].Origin, eff[0].Level)
	}
}

// Binding beats specificity, which is the whole point of it: a supplier keeps the few
// statements it cannot give up regardless of customer.
func TestBindingBeatsAMoreSpecificRule(t *testing.T) {
	root := tree(t, map[string]string{
		ConfigDir + "/given/provider/r.yaml": rule("r", "provider", Review, "binding: true"),
		ConfigDir + "/given/project/r.yaml":  rule("r", "project", Review),
	})
	read, _ := Load(root)
	eff, collisions := Effective(read)
	if len(collisions) != 0 {
		t.Fatalf("collisions: %v", collisions)
	}
	if len(eff) != 1 || eff[0].Level != "provider" {
		t.Fatalf("resolved to %+v, want the binding provider rule", eff)
	}
}

// The refusal section 9 writes for this case: red, and stays red. The resolver must not pick
// the more specific one, which is what it would do if nobody had read that sentence.
func TestTwoBindingRulesOnDifferentLevelsStayUnresolved(t *testing.T) {
	root := tree(t, map[string]string{
		ConfigDir + "/given/provider/r.yaml": rule("r", "provider", Review, "binding: true"),
		ConfigDir + "/given/org/r.yaml":      rule("r", "org", Review, "binding: true"),
	})
	read, _ := Load(root)
	eff, collisions := Effective(read)
	if len(eff) != 0 {
		t.Fatalf("effective set has %d rules, want the id left out", len(eff))
	}
	if len(collisions) != 1 {
		t.Fatalf("collisions: %v, want one", collisions)
	}
	if !strings.Contains(collisions[0].Cause, "given/provider/r.yaml") ||
		!strings.Contains(collisions[0].Cause, "given/org/r.yaml") {
		t.Fatalf("the collision names %q, want both files", collisions[0].Cause)
	}
}

func TestTheConfigurationErrorsSectionNineNames(t *testing.T) {
	cases := []struct {
		name, path, body, cause string
	}{
		{"scope against the path", ConfigDir + "/given/org/r.yaml",
			rule("r", "project", Review), "claims scope project where its path gives it org"},
		{"binding under learned", ConfigDir + "/learned/org/r.yaml",
			rule("r", "org", Review, "binding: true", "abstract: a summary"), "is binding under learned/"},
		{"binding under given/project", ConfigDir + "/given/project/r.yaml",
			rule("r", "project", Review, "binding: true"), "is binding under given/project/"},
		{"checked without a check", ConfigDir + "/given/org/r.yaml",
			strings.Replace(rule("r", "org", Checked), "check:\n  type: section-implies-section\n  when: { section: interfaces, non_empty: true }\n  then: { section: migration-notes, non_empty: true }\n", "", 1),
			"is kind checked and carries no check"},
		{"review with a check", ConfigDir + "/given/org/r.yaml",
			strings.Replace(rule("r", "org", Checked), "kind: checked", "kind: review", 1),
			"is kind review and carries a check"},
		{"a learned rule above project without an abstract", ConfigDir + "/learned/provider/r.yaml",
			rule("r", "provider", Review), "carries no abstract"},
		{"no id", ConfigDir + "/given/org/r.yaml",
			strings.Replace(rule("r", "org", Review), "id: r\n", "", 1), "has no id"},
		{"no statement", ConfigDir + "/given/org/r.yaml",
			strings.Replace(rule("r", "org", Review), "statement: >\n  Something has to hold.\n", "", 1), "has no statement"},
		{"an unknown kind", ConfigDir + "/given/org/r.yaml",
			strings.Replace(rule("r", "org", Review), "kind: review", "kind: advisory", 1), "has kind advisory"},
		{"not a rule file at all", ConfigDir + "/given/org/r.yaml",
			"- this is a list\n", "is not a readable rule file"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			read, problems := Load(tree(t, map[string]string{c.path: c.body}))
			if len(read) != 0 {
				t.Fatalf("read %d rules, want the file left out of the set", len(read))
			}
			var causes []string
			for _, p := range problems {
				causes = append(causes, p.Cause)
				if p.Next == "" {
					t.Errorf("problem %q has no next step", p.Cause)
				}
			}
			if !strings.Contains(strings.Join(causes, "\n"), c.cause) {
				t.Fatalf("problems %v, want one containing %q", causes, c.cause)
			}
		})
	}
}

// The finding is on the directory, because its presence is the error rather than anything in
// the files it holds.
func TestLearnedBuiltinIsRejectedOnTheDirectory(t *testing.T) {
	root := tree(t, map[string]string{
		ConfigDir + "/learned/builtin/r.yaml": rule("r", "builtin", Review),
	})
	read, problems := Load(root)
	if len(read) != 0 {
		t.Fatalf("read %d rules from learned/builtin, want none", len(read))
	}
	if len(problems) != 1 || !strings.HasSuffix(problems[0].Path, "learned/builtin") {
		t.Fatalf("problems %v, want one on the directory", problems)
	}
}

func TestADuplicateIdAtOneLevelIsAProblem(t *testing.T) {
	root := tree(t, map[string]string{
		ConfigDir + "/given/org/a.yaml": rule("r", "org", Review),
		ConfigDir + "/given/org/b.yaml": rule("r", "org", Review),
	})
	_, problems := Load(root)
	if len(problems) != 1 || !strings.Contains(problems[0].Cause, "already claims at the same level") {
		t.Fatalf("problems %v, want one duplicate id", problems)
	}
}

// The hash covers the set's content, so these three leave it alone and the fourth changes it.
func TestTheHashCoversTheContentAndNotTheFiles(t *testing.T) {
	base := map[string]string{
		ConfigDir + "/given/org/r.yaml":     rule("r", "org", Review),
		ConfigDir + "/given/project/s.yaml": rule("s", "project", Checked),
	}
	want := hashOf(t, tree(t, base))

	withComment := map[string]string{}
	for k, v := range base {
		withComment[k] = "# a comment nobody hashes\n" + v
	}
	if got := hashOf(t, tree(t, withComment)); got != want {
		t.Error("a comment changed rules_hash")
	}

	renamed := map[string]string{
		ConfigDir + "/given/org/zzz.yaml":    base[ConfigDir+"/given/org/r.yaml"],
		ConfigDir + "/given/project/aa.yaml": base[ConfigDir+"/given/project/s.yaml"],
	}
	if got := hashOf(t, tree(t, renamed)); got != want {
		t.Error("a file name changed rules_hash")
	}

	rewrapped := map[string]string{}
	for k, v := range base {
		rewrapped[k] = strings.Replace(v, "  Something has to hold.", "  Something\n  has to hold.", 1)
	}
	if got := hashOf(t, tree(t, rewrapped)); got != want {
		t.Error("a line break in a folded statement changed rules_hash")
	}

	edited := map[string]string{}
	for k, v := range base {
		edited[k] = strings.Replace(v, "Something has to hold.", "Something else has to hold.", 1)
	}
	if got := hashOf(t, tree(t, edited)); got == want {
		t.Error("an edited statement left rules_hash alone")
	}
}

// A rule that lost its precedence contest never applied, so it is not in the hash.
func TestOnlyTheEffectiveSetReachesTheHash(t *testing.T) {
	one := hashOf(t, tree(t, map[string]string{
		ConfigDir + "/given/project/r.yaml": rule("r", "project", Review),
	}))
	two := hashOf(t, tree(t, map[string]string{
		ConfigDir + "/given/project/r.yaml": rule("r", "project", Review),
		ConfigDir + "/given/org/r.yaml":     rule("r", "org", Review),
	}))
	if one != two {
		t.Error("a rule that lost its precedence contest reached rules_hash")
	}
}

// The definition is checked against a sha256sum pipeline rather than against itself, which is
// what Appendix B's byte exactness means and what the hashing tests in this repository already
// do for the hashes the appendix fixes.
func TestTheHashIsTheSha256OfTheRendering(t *testing.T) {
	root := tree(t, map[string]string{
		ConfigDir + "/given/org/r.yaml":     rule("r", "org", Review),
		ConfigDir + "/given/project/s.yaml": rule("s", "project", Checked),
	})
	read, _ := Load(root)
	eff, _ := Effective(read)
	rendering := Render(eff)

	cmd := exec.Command("sha256sum")
	cmd.Stdin = strings.NewReader(rendering)
	out, err := cmd.Output()
	if err != nil {
		t.Skipf("sha256sum is not available: %v", err)
	}
	want := strings.Fields(string(out))[0]
	if got := Hash(eff); got != want {
		t.Fatalf("Hash is %s, sha256sum of the rendering is %s", got, want)
	}
}

// The rendering is one line per rule, sorted by id, tab separated, so that a person can read
// it and recompute the hash by hand.
func TestTheRenderingIsOneSortedLinePerRule(t *testing.T) {
	root := tree(t, map[string]string{
		ConfigDir + "/given/project/s.yaml": rule("s", "project", Review),
		ConfigDir + "/given/org/a.yaml":     rule("a", "org", Checked),
	})
	read, _ := Load(root)
	eff, _ := Effective(read)
	lines := strings.Split(strings.TrimRight(Render(eff), "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("rendering has %d lines, want 2:\n%s", len(lines), Render(eff))
	}
	if !strings.HasPrefix(lines[0], "a\t") || !strings.HasPrefix(lines[1], "s\t") {
		t.Fatalf("lines are not sorted by id:\n%s", Render(eff))
	}
	if n := strings.Count(lines[0], "\t"); n != 8 {
		t.Fatalf("a line carries %d tabs, want 8", n)
	}
	if !strings.Contains(lines[0], "type=section-implies-section") {
		t.Fatalf("a checked rule's line does not carry its check type:\n%s", lines[0])
	}
}

// The parameters are flattened by this package rather than marshalled, so two trees whose
// checks differ only in key order agree.
func TestKeyOrderInACheckDoesNotReachTheHash(t *testing.T) {
	a := rule("r", "org", Checked)
	b := strings.Replace(a,
		"  when: { section: interfaces, non_empty: true }\n  then: { section: migration-notes, non_empty: true }",
		"  then: { non_empty: true, section: migration-notes }\n  when: { non_empty: true, section: interfaces }", 1)
	one := hashOf(t, tree(t, map[string]string{ConfigDir + "/given/org/r.yaml": a}))
	two := hashOf(t, tree(t, map[string]string{ConfigDir + "/given/org/r.yaml": b}))
	if one != two {
		t.Error("the order of a check's keys changed rules_hash")
	}
}

func hashOf(t *testing.T, root string) string {
	t.Helper()
	read, _ := Load(root)
	eff, _ := Effective(read)
	return Hash(eff)
}
