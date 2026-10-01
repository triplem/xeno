// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// repoRoot is this repository, which carries the shipped set. The set is read where it lives
// rather than copied into a fixture: what is under test is the files a project receives.
const repoRoot = "../.."

func shipped(t *testing.T) []Rule {
	t.Helper()
	read, problems := Load(repoRoot)
	if len(problems) != 0 {
		t.Fatalf("the shipped set does not load cleanly: %v", problems)
	}
	effective, collisions := Effective(read)
	if len(collisions) != 0 {
		t.Fatalf("the shipped set collides with itself: %v", collisions)
	}
	return effective
}

// Every file in given/builtin reaches the effective set, which is the gate's own reading of
// whether the set is well formed. Written the other way round the test would only prove that the
// files parse.
func TestTheShippedSetResolves(t *testing.T) {
	dir := filepath.Join(repoRoot, PluginDir, "given/builtin")
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("there is no shipped set at %s: %v", dir, err)
	}
	var files int
	for _, e := range entries {
		if !e.IsDir() && isRuleFile(e.Name()) {
			files++
		}
	}
	if files == 0 {
		t.Fatal("the shipped set is empty")
	}
	if got := len(shipped(t)); got != files {
		t.Fatalf("%d rule files and %d rules in force", files, got)
	}
}

// Section 9: the set is deliberately minimal and generic, and anything specific to a supplier, a
// customer or a project belongs in that party's own layer. These are the words that would show
// it had drifted.
func TestTheShippedSetNamesNothingSpecific(t *testing.T) {
	forbidden := []string{"go", "golang", "gofmt", "npm", "python", "java", "xeno", "triplem",
		"github", "gitlab", "docker", "kubernetes"}
	for _, r := range shipped(t) {
		words := strings.FieldsFunc(strings.ToLower(r.Statement), func(c rune) bool {
			return !(c >= 'a' && c <= 'z')
		})
		for _, w := range words {
			for _, f := range forbidden {
				if w == f {
					t.Errorf("%s names %q in its statement, which is specific to a language, a tool or a project", r.ID, f)
				}
			}
		}
	}
}

// Every rule in the set carries what a reader and a gate need: the level it is at, a statement
// to answer, and the phases it applies to.
func TestEveryShippedRuleIsWellFormedForAReader(t *testing.T) {
	for _, r := range shipped(t) {
		if r.Level != "builtin" || r.Scope != "builtin" {
			t.Errorf("%s is at %s and claims %s", r.ID, r.Level, r.Scope)
		}
		if r.Binding {
			t.Errorf("%s is binding; the shipped set binds nothing, so that a project can override it", r.ID)
		}
		if len(strings.Fields(r.Statement)) < 8 {
			t.Errorf("%s has a statement of %d words, which is too short to answer against",
				r.ID, len(strings.Fields(r.Statement)))
		}
		if !strings.HasSuffix(strings.TrimSpace(r.Statement), ".") {
			t.Errorf("%s's statement is not a sentence: %q", r.ID, r.Statement)
		}
		if len(r.AppliesTo) == 0 {
			t.Errorf("%s names no phase", r.ID)
		}
	}
}

// A checked rule in the set has to name a type the runner implements, which is the failure a
// project would otherwise meet on its first phase. The names are the registry's, duplicated here
// rather than imported, because internal/gates imports this package.
func TestAShippedCheckedRuleNamesAnImplementedType(t *testing.T) {
	implemented := map[string]bool{
		"section-implies-section": true, "section-non-empty": true,
		"commit-message": true, "commit-trailer": true,
		"commit-signature": true, "approver-not-author": true,
	}
	var checked int
	for _, r := range shipped(t) {
		if r.Kind != Checked {
			continue
		}
		checked++
		if r.Check == nil || !implemented[r.Check.Type] {
			t.Errorf("%s is checked and names type %q, which nothing implements", r.ID, r.Check.Type)
		}
	}
	if checked == 0 {
		t.Error("nothing in the shipped set is checked, so the set demonstrates no predicate at all")
	}
}

// The examples are adopted by copying, so what matters is that they resolve where a reader puts
// them. Nothing reads examples/ itself.
func TestTheExamplesResolveWhereAReaderPutsThem(t *testing.T) {
	src := filepath.Join(repoRoot, "examples/rules")
	entries, err := os.ReadDir(src)
	if err != nil {
		t.Fatalf("there are no examples at %s: %v", src, err)
	}
	root := t.TempDir()
	dir := filepath.Join(root, ConfigDir, "given/org")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	var copied int
	for _, e := range entries {
		if e.IsDir() || !isRuleFile(e.Name()) {
			continue
		}
		b, err := os.ReadFile(filepath.Join(src, e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, e.Name()), b, 0o644); err != nil {
			t.Fatal(err)
		}
		copied++
	}
	if copied == 0 {
		t.Fatal("there are no example rules")
	}
	read, problems := Load(root)
	if len(problems) != 0 {
		t.Fatalf("an example does not resolve when copied into given/org: %v", problems)
	}
	if len(read) != copied {
		t.Fatalf("%d examples copied and %d read", copied, len(read))
	}
	// Each one declares the level it is written for, so that copying it is the whole adoption.
	for _, r := range read {
		if r.Scope != "org" {
			t.Errorf("%s claims scope %s; the examples are written for given/org", r.ID, r.Scope)
		}
	}
}

// An example copied to the wrong level is a finding that names both the claim and the path,
// which is how a reader learns the two axes without a document.
func TestAnExampleAtTheWrongLevelSaysSo(t *testing.T) {
	b, err := os.ReadFile(filepath.Join(repoRoot, "examples/rules/conventional-commits.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	dir := filepath.Join(root, ConfigDir, "given/project")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "conventional-commits.yaml"), b, 0o644); err != nil {
		t.Fatal(err)
	}
	_, problems := Load(root)
	if len(problems) != 1 {
		t.Fatalf("problems %v, want one about the scope", problems)
	}
	if !strings.Contains(problems[0].Cause, "scope org") || !strings.Contains(problems[0].Cause, "project") {
		t.Fatalf("the finding is %q, want both the claim and the path", problems[0].Cause)
	}
}
