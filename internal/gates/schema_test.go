// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
)

const fixtureIntent = "git.example/group/proj#1"

// The corpus is read from testdata rather than built here, and the fields it asserts
// come from a manifest written out of section 5 rather than from the constants in this
// package. Written the other way round the test would only prove that the code agrees
// with itself.
func corpus(t *testing.T) (root string, required map[string][]string) {
	t.Helper()
	root = t.TempDir()
	dir := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir("testdata/phase")
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		b, err := os.ReadFile(filepath.Join("testdata/phase", e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, e.Name()), b, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if err := fm.ReadYAML("testdata/required-fields.yaml", &required); err != nil {
		t.Fatal(err)
	}
	return root, required
}

func ctxFor(root string) Ctx {
	return Ctx{Root: root, Key: "PROJ-1", Phase: model.Phases[0], QualifiedID: fixtureIntent}
}

// dropField removes the line that carries one key, which is what a forgotten field
// looks like. The fixtures are flat on purpose so that this stays a line and not a
// YAML rewrite that could change something else on the way.
func dropField(t *testing.T, root, file, field string) {
	t.Helper()
	p := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]), file)
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	var kept []string
	dropped := false
	for _, line := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(line, field+":") {
			dropped = true
			continue
		}
		kept = append(kept, line)
	}
	if !dropped {
		t.Fatalf("%s carries no %s to drop; the fixture and the manifest disagree", file, field)
	}
	if err := os.WriteFile(p, []byte(strings.Join(kept, "\n")), 0o644); err != nil {
		t.Fatal(err)
	}
}

func causes(c model.Check) string {
	var b strings.Builder
	for _, f := range c.Findings {
		b.WriteString(f.File + ": " + f.Cause + "\n")
	}
	return b.String()
}

func TestWellFormedCorpusPasses(t *testing.T) {
	root, _ := corpus(t)
	if got := schema(ctxFor(root)); len(got.Findings) != 0 {
		t.Fatalf("a well formed phase was rejected:\n%s", causes(got))
	}
}

// One malformed variant per required field, per artifact kind, from the manifest.
func TestEveryRequiredFieldIsReported(t *testing.T) {
	_, required := corpus(t)
	for file, fields := range required {
		for _, field := range fields {
			t.Run(file+"/"+field, func(t *testing.T) {
				root, _ := corpus(t)
				dropField(t, root, file, field)
				got := schema(ctxFor(root))
				want := "required field missing: " + field
				if !strings.Contains(causes(got), want) {
					t.Fatalf("dropping %s from %s produced no finding naming it:\n%s", field, file, causes(got))
				}
			})
		}
	}
}

// Absent is the schema that predates the field and stays readable; present has to be
// well formed. Both halves are the process definition's, section 5.
func TestSchemaVersionAbsentIsReadableAndMalformedIsNot(t *testing.T) {
	for _, file := range []string{"output.md", "digest.md", "context.lock.yaml", "learning.yaml"} {
		root, _ := corpus(t)
		dropField(t, root, file, "schema_version")
		if got := schema(ctxFor(root)); len(got.Findings) != 0 {
			t.Fatalf("%s without schema_version was rejected:\n%s", file, causes(got))
		}
	}

	root, _ := corpus(t)
	p := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]), "output.md")
	b, _ := os.ReadFile(p)
	edited := strings.Replace(string(b), `schema_version: "1.0"`, `schema_version: "one"`, 1)
	if edited == string(b) {
		t.Fatal("the fixture no longer carries the value this test edits")
	}
	if err := os.WriteFile(p, []byte(edited), 0o644); err != nil {
		t.Fatal(err)
	}
	got := schema(ctxFor(root))
	if !strings.Contains(causes(got), "is not major.minor") {
		t.Fatalf("a malformed schema_version passed:\n%s", causes(got))
	}
}
