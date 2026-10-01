// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/rules"
)

func writeRule(t *testing.T, root, rel, body string) {
	t.Helper()
	p := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

const reviewRule = `id: migration-note
version: 1
scope: org
kind: review
applies_to: [02-design]
statement: >
  A change to a published interface comes with a migration note.
`

// Every repository is in this state until a rule set exists, this one included, so it is the
// first case rather than the edge.
func TestNoRuleTreePasses(t *testing.T) {
	root, _ := corpus(t)
	c := rulesGate(ctxFor(root))
	if c.Result != "pass" {
		t.Fatalf("G-Rules is %s over a repository with no rule tree, want pass: %v", c.Result, c.Findings)
	}
}

func TestAWellFormedTreePasses(t *testing.T) {
	root, _ := corpus(t)
	writeRule(t, root, rules.ConfigDir+"/given/org/migration-note.yaml", reviewRule)
	if c := rulesGate(ctxFor(root)); c.Result != "pass" {
		t.Fatalf("G-Rules is %s, want pass: %v", c.Result, c.Findings)
	}
}

// The gate words what internal/rules returns, so what is checked here is that a problem
// reaches a verdict with the three things section 7 requires of a red one.
func TestAMisplacedScopeIsRedAndNamesFileCauseAndNextStep(t *testing.T) {
	root, _ := corpus(t)
	writeRule(t, root, rules.ConfigDir+"/given/project/migration-note.yaml", reviewRule)
	c := rulesGate(ctxFor(root))
	if c.Result != "fail" {
		t.Fatalf("G-Rules is %s, want fail", c.Result)
	}
	if len(c.Findings) != 1 {
		t.Fatalf("%d findings, want 1: %v", len(c.Findings), c.Findings)
	}
	f := c.Findings[0]
	if !strings.HasSuffix(f.File, "migration-note.yaml") || f.Cause == "" || f.Next == "" {
		t.Fatalf("finding %+v does not name a file, a cause and a next step", f)
	}
}

// A collision is reported by Effective rather than by Load, so this is the one case that
// proves the gate reads both.
func TestABindingCollisionReachesTheVerdict(t *testing.T) {
	root, _ := corpus(t)
	binding := strings.Replace(reviewRule, "kind: review", "binding: true\nkind: review", 1)
	writeRule(t, root, rules.ConfigDir+"/given/provider/m.yaml",
		strings.Replace(binding, "scope: org", "scope: provider", 1))
	writeRule(t, root, rules.ConfigDir+"/given/org/m.yaml", binding)
	c := rulesGate(ctxFor(root))
	if c.Result != "fail" {
		t.Fatalf("G-Rules is %s, want fail", c.Result)
	}
	if len(c.Findings) != 1 || !strings.Contains(c.Findings[0].Cause, "collides") {
		t.Fatalf("findings %v, want one collision", c.Findings)
	}
}

// The gate was `not-implemented` since the table was written, and the table is the one place
// that decides what a phase reports.
func TestGRulesIsNoLongerReportedAsNotImplemented(t *testing.T) {
	root, _ := corpus(t)
	for _, s := range table {
		if s.id != "G-Rules" {
			continue
		}
		if got := s.fn(ctxFor(root)).Result; got == "not-implemented" {
			t.Fatal("G-Rules still reports not-implemented")
		}
		return
	}
	t.Fatal("G-Rules is not in the table")
}
