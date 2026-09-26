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

// setField rewrites one flat frontmatter line, the counterpart of dropField.
func setField(t *testing.T, root, file, field, value string) {
	t.Helper()
	p := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]), file)
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	var kept []string
	for _, line := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(line, field+":") {
			line = field + ": " + value
		}
		kept = append(kept, line)
	}
	if err := os.WriteFile(p, []byte(strings.Join(kept, "\n")), 0o644); err != nil {
		t.Fatal(err)
	}
}

// pluginTemplate writes a template the loader can resolve, so the artifact's `template`
// field can be compared against the version the repository carries.
func pluginTemplate(t *testing.T, root, version string) {
	t.Helper()
	dir := filepath.Join(root, ".xeno/plugin/templates/intake")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	write := func(name, content string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("template.yaml", "id: intake\nversion: "+version+"\ntitle_key: title\nsections:\n  - id: problem\n    required: true\n")
	write("strings.en.yaml", "title: Intake\nheadings:\n  problem: Problem\n")
}

// Appendix B: a hash field carries a sha256 or the placeholder, and the placeholder is
// honest only where nothing could have written the value.
func TestHashFieldShape(t *testing.T) {
	for _, tc := range []struct {
		name, field, value, tool, pluginVersion string
		wantFinding                             bool
	}{
		{name: "a sha256 passes", field: "context_hash", value: strings.Repeat("a", 64), tool: "claude-code"},
		{name: "a truncated hash is a finding", field: "context_hash", value: "abc123", tool: "claude-code", wantFinding: true},
		{name: "an invented value is a finding", field: "context_hash", value: "todo", tool: "claude-code", wantFinding: true},
		{name: "an uppercase hash is a finding", field: "context_hash", value: strings.Repeat("A", 64), tool: "claude-code", wantFinding: true},
		{name: "by-hand is a finding where a session wrote", field: "context_hash", value: "by-hand", tool: "claude-code", wantFinding: true},
		{name: "by-hand passes a manual artifact", field: "context_hash", value: "by-hand", tool: "manual"},
		{name: "by-hand passes where no writer exists", field: "secrets_hash", value: "by-hand", tool: "claude-code"},
		{name: "by-hand passes where the bundle is gone", field: "strings_hash", value: "by-hand", tool: "claude-code", pluginVersion: "2.0.0"},
		{name: "by-hand is a finding where the bundle is here", field: "strings_hash", value: "by-hand", tool: "claude-code", pluginVersion: "1.0.0", wantFinding: true},
	} {
		root, _ := corpus(t)
		if tc.pluginVersion != "" {
			pluginTemplate(t, root, tc.pluginVersion)
			setField(t, root, "output.md", "template", "intake@1.0.0")
		}
		for _, file := range []string{"output.md", "digest.md"} {
			setField(t, root, file, "tool", tc.tool)
			setField(t, root, file, tc.field, tc.value)
		}
		causes := causes(schema(ctxFor(root)))
		if strings.Contains(causes, tc.field) != tc.wantFinding {
			t.Errorf("%s: wanted finding=%v; causes: %s", tc.name, tc.wantFinding, causes)
		}
	}
}
