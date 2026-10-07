// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/hashing"
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
		// Nothing about schema_version, rather than nothing at all: dropping a line from
		// context.lock.yaml changes the file, so the artifacts' context_hash stops
		// matching it, which is a finding of its own and the subject of another test.
		if got := causes(schema(ctxFor(root))); strings.Contains(got, "schema_version") {
			t.Fatalf("%s without schema_version was rejected:\n%s", file, got)
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
		{name: "the lock's own hash passes", field: "context_hash", value: lockHash, tool: "claude-code"},
		{name: "a truncated hash is a finding", field: "context_hash", value: "abc123", tool: "claude-code", wantFinding: true},
		{name: "an invented value is a finding", field: "context_hash", value: "todo", tool: "claude-code", wantFinding: true},
		{name: "an uppercase hash is a finding", field: "context_hash", value: strings.Repeat("A", 64), tool: "claude-code", wantFinding: true},
		{name: "by-hand is a finding where a session wrote", field: "context_hash", value: "by-hand", tool: "claude-code", wantFinding: true},
		// The tool field is self-reported, so it is not evidence that nothing could have written
		// the value, and #277 took it out of the exemption. What a manual artifact declares about
		// itself no longer decides a hash check.
		{name: "by-hand is a finding even where the artifact says it was manual", field: "context_hash", value: "by-hand", tool: "manual", wantFinding: true},
		{name: "by-hand passes where no writer exists", field: "secrets_hash", value: "by-hand", tool: "claude-code"},
		{name: "by-hand passes where the bundle is gone", field: "strings_hash", value: "by-hand", tool: "claude-code", pluginVersion: "2.0.0"},
		{name: "by-hand is a finding where the bundle is here", field: "strings_hash", value: "by-hand", tool: "claude-code", pluginVersion: "1.0.0", wantFinding: true},
	} {
		root, _ := corpus(t)
		if tc.value == lockHash {
			tc.value = fixtureLockHash(t, root)
		}
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

// lockHash is a placeholder in the table above, replaced per case with the hash of that
// root's own lock file: the table cannot know it, because corpus copies the fixture into a
// fresh directory for every case.
const lockHash = "<the lock's hash>"

func fixtureLockHash(t *testing.T, root string) string {
	t.Helper()
	h, err := hashing.FileHash(filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]), "context.lock.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	return h
}

// A hash of the right shape over the wrong content is the case a shape check cannot catch,
// and the one a trail rests on: the value says what the phase was produced from.
func TestHashValueIsRecomputed(t *testing.T) {
	root, _ := corpus(t)
	if got := causes(schema(ctxFor(root))); got != "" {
		t.Fatalf("the fixture does not match its own lock file:\n%s", got)
	}

	root, _ = corpus(t)
	setField(t, root, "output.md", "context_hash", strings.Repeat("d", 64))
	if got := causes(schema(ctxFor(root))); !strings.Contains(got, "context_hash does not match") {
		t.Fatalf("a well shaped hash over the wrong content passed:\n%s", got)
	}

	// The same for the bundle: the artifact names a version the repository carries, so the
	// value is recomputable and has to be right.
	root, _ = corpus(t)
	pluginTemplate(t, root, "1.0.0")
	setField(t, root, "output.md", "template", "intake@1.0.0")
	setField(t, root, "output.md", "strings_hash", strings.Repeat("e", 64))
	if got := causes(schema(ctxFor(root))); !strings.Contains(got, "strings_hash does not match") {
		t.Fatalf("a well shaped strings_hash over the wrong bundle passed:\n%s", got)
	}
}

// ---- undeclared evidence, keyed on the path relative to evidence/

// evidenceFile writes one file under evidence/ of the corpus phase, at a path relative to
// that directory, so a test can put one inside a subdirectory.
func evidenceFile(t *testing.T, root, rel, content string) {
	t.Helper()
	p := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]), "evidence", rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

// declare appends an evidence block to the corpus output.md, one item per path given. A
// path of "" is the declaration that carries none, which is a pending item.
func declare(t *testing.T, root string, paths ...string) {
	t.Helper()
	p := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]), "output.md")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	var block strings.Builder
	block.WriteString("evidence:\n")
	for i, path := range paths {
		block.WriteString("  - kind: other\n    job: j" + string(rune('a'+i)) + "\n")
		if path != "" {
			block.WriteString("    path: " + path + "\n")
			block.WriteString("    sha256: " + strings.Repeat("4", 64) + "\n")
		}
	}
	edited := strings.Replace(string(b), "---\n\n# Intake", block.String()+"---\n\n# Intake", 1)
	if edited == string(b) {
		t.Fatal("the corpus output.md no longer ends its frontmatter where this expects")
	}
	if err := os.WriteFile(p, []byte(edited), 0o644); err != nil {
		t.Fatal(err)
	}
}

// undeclared returns the causes of the findings this check produces, and nothing else the
// schema gate reports, so that a declaration added by hand cannot pass the test by
// producing a different finding.
func undeclared(t *testing.T, root string) string {
	t.Helper()
	var b strings.Builder
	for _, f := range schema(ctxFor(root)).Findings {
		if strings.Contains(f.Cause, "undeclared file in evidence/") {
			b.WriteString(f.File + "\n")
		}
	}
	return b.String()
}

// A pending item carries a kind and a job and no path, and declares no file. This passed
// before the change as well: filepath.Base of an empty string is ".", which matched no entry,
// so the set carried something inert rather than something wrong. The test is here because
// the entry is gone now and nothing else would notice if a later key brought it back.
func TestAPendingDeclarationDeclaresNoFile(t *testing.T) {
	root, _ := corpus(t)
	evidenceFile(t, root, "junit.xml", "<testsuite/>\n")
	declare(t, root, "")

	got := undeclared(t, root)
	if !strings.Contains(got, "evidence/junit.xml") {
		t.Fatalf("a file no declaration names was not reported:\n%s", got)
	}
	if strings.Contains(got, "evidence/.") {
		t.Errorf("the set carried an entry keyed on a basename of nothing:\n%s", got)
	}
}

// Section 4 forbids no subdirectory under evidence/. A declaration names the file it names
// and nothing at another depth.
func TestADeclarationUnderASubdirectoryDeclaresThatFileAlone(t *testing.T) {
	root, _ := corpus(t)
	evidenceFile(t, root, "logs/a.txt", "nested\n")
	evidenceFile(t, root, "a.txt", "top level\n")
	declare(t, root, "evidence/logs/a.txt")

	got := undeclared(t, root)
	if strings.Contains(got, "evidence/logs/a.txt") {
		t.Errorf("the declared file was reported undeclared:\n%s", got)
	}
	if !strings.Contains(got, "evidence/a.txt") {
		t.Errorf("a file of the same basename at another depth was taken as declared:\n%s", got)
	}
}

// The directory holding a declared file is not itself something a declaration can name, so
// widening what is read must not widen what is reported.
func TestADirectoryOfDeclaredFilesIsNoFinding(t *testing.T) {
	root, _ := corpus(t)
	evidenceFile(t, root, "logs/a.txt", "nested\n")
	declare(t, root, "evidence/logs/a.txt")

	if got := undeclared(t, root); got != "" {
		t.Fatalf("a directory whose files are all declared was reported:\n%s", got)
	}
}

// A path that names nothing inside evidence/ declares nothing inside it, which is what a
// basename key did by accident and this does by construction.
func TestAPathOutsideEvidenceDeclaresNothingInside(t *testing.T) {
	root, _ := corpus(t)
	evidenceFile(t, root, "digest.md", "not the phase digest\n")
	declare(t, root, "digest.md")

	if got := undeclared(t, root); !strings.Contains(got, "evidence/digest.md") {
		t.Fatalf("a path outside evidence/ was taken to declare a file inside it:\n%s", got)
	}
}
