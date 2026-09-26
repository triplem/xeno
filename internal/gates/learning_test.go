// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
)

// writeLearning replaces the record in the fixture, which is the file G-Learning judges.
func writeLearning(t *testing.T, root, body string) {
	t.Helper()
	header := "intent: " + fixtureIntent + "\nphase: 00-intake\ncreated: 2026-09-20T10:00:00Z\n" +
		"schema_version: \"1.0\"\nrunner_version: 0.1.0-dev\nplugin_version: 0.1.0-dev\n"
	p := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]), "learning.yaml")
	if err := os.WriteFile(p, []byte(header+body), 0o644); err != nil {
		t.Fatal(err)
	}
}

const wellFormed = `learnings:
  - category: project-convention
    observation: the register was read row by row
    proposal: reread it when a row's subject changes
    target: CLAUDE.md
`

// Section 10 defines a shape, and counting keys was not it: a record could invent a key
// and a category of its own and pass, which is A7 and how the gap was found.
func TestLearningShape(t *testing.T) {
	for _, tc := range []struct {
		name, body, want string
	}{
		{"no_finding is the honest empty case", "no_finding: true\n", ""},
		{"a well formed entry passes", wellFormed, ""},
		{"an empty record is a finding", "", "empty"},
		{"an invented key is a finding",
			"observations:\n  - it happened\n", "unknown key observations"},
		{"a category outside the set is a finding",
			strings.Replace(wellFormed, "project-convention", "process", 1), "category process"},
		{"a missing proposal is a finding",
			"learnings:\n  - category: template\n    observation: it happened\n    target: CLAUDE.md\n",
			"has no proposal"},
		{"a missing target is a finding",
			"learnings:\n  - category: template\n    observation: it happened\n    proposal: change it\n",
			"has no target"},
		{"an unknown key inside an entry is a finding",
			wellFormed + "    rule_id: R-1\n", "unknown key rule_id"},
		{"a bare list is a finding",
			"learnings:\n  - it happened\n", "not a mapping"},
		{"learnings that is not a list is a finding",
			"learnings: it happened\n", "not a list"},
	} {
		root, _ := corpus(t)
		writeLearning(t, root, tc.body)
		got := causes(learning(ctxFor(root)))
		switch {
		case tc.want == "" && got != "":
			t.Errorf("%s: wanted no finding, got:\n%s", tc.name, got)
		case tc.want != "" && !strings.Contains(got, tc.want):
			t.Errorf("%s: wanted a finding naming %q, got:\n%s", tc.name, tc.want, got)
		}
	}
}

// Every category the specification lists is accepted, so the closed set in the code is
// the one in section 10 rather than a subset somebody found convenient.
func TestEveryCategoryInTheSetIsAccepted(t *testing.T) {
	for _, category := range []string{"template", "prompt", "context-rule", "project-convention"} {
		root, _ := corpus(t)
		writeLearning(t, root, strings.Replace(wellFormed, "project-convention", category, 1))
		if got := causes(learning(ctxFor(root))); got != "" {
			t.Errorf("category %s was rejected:\n%s", category, got)
		}
	}
}
