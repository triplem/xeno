// SPDX-License-Identifier: Apache-2.0

package external

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
)

// gate writes a command into a repository and returns the declaration for it, with the hash the
// file actually has. A test that wants a mismatch declares something else.
func gate(t *testing.T, root, id, script string, phases ...string) model.ExternalGate {
	t.Helper()
	rel := filepath.Join("tools", id)
	p := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(script), 0o755); err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(script))
	if len(phases) == 0 {
		phases = []string{"03-implementation"}
	}
	return model.ExternalGate{ID: id, Path: rel, SHA256: hex.EncodeToString(sum[:]), Phases: phases}
}

func run1(t *testing.T, root string, d model.ExternalGate) model.Check {
	t.Helper()
	checks := Run(root, "03-implementation", "deadbeef", "git.example/p#1",
		".xeno/intents/PROJ-1/phases/03-implementation", []model.ExternalGate{d})
	if len(checks) != 1 {
		t.Fatalf("%d checks, want 1", len(checks))
	}
	return checks[0]
}

func causes(c model.Check) string {
	var b strings.Builder
	for _, f := range c.Findings {
		b.WriteString(f.Cause + " | " + f.Next + "\n")
	}
	return b.String()
}

const passing = `#!/bin/sh
cat > /dev/null
printf '{"findings":[]}'
`

// Every project today, and Appendix A's default: none declared, none run.
func TestNothingDeclaredRunsNothing(t *testing.T) {
	if got := Run(t.TempDir(), "03-implementation", "h", "i", "d", nil); got != nil {
		t.Fatalf("%d checks from no declaration, want none", len(got))
	}
}

func TestADeclaredGateRunsAndPasses(t *testing.T) {
	root := t.TempDir()
	c := run1(t, root, gate(t, root, "house-linter", passing))
	if c.Result != "pass" {
		t.Fatalf("result %s: %s", c.Result, causes(c))
	}
	if c.Gate != "house-linter" {
		t.Errorf("the check is named %q, want the declaration's id", c.Gate)
	}
	if c.Provenance != Provenance {
		t.Errorf("provenance %q, want %q: the mark is the point of section 14", c.Provenance, Provenance)
	}
}

// Section 14: the exit status decides. A tool that reports something it does not consider fatal
// passes and carries its findings.
func TestExitZeroWithFindingsIsAPassCarryingThem(t *testing.T) {
	root := t.TempDir()
	c := run1(t, root, gate(t, root, "warner", `#!/bin/sh
cat > /dev/null
printf '{"findings":[{"file":"a.go","cause":"a thing worth mentioning","next":"mention it"}]}'
`))
	if c.Result != "pass" {
		t.Fatalf("result %s, want pass: the exit status decides", c.Result)
	}
	if len(c.Findings) != 1 || c.Findings[0].Cause != "a thing worth mentioning" {
		t.Fatalf("findings %v, want the one it reported", c.Findings)
	}
}

func TestANonZeroExitIsAFail(t *testing.T) {
	root := t.TempDir()
	c := run1(t, root, gate(t, root, "failer", `#!/bin/sh
cat > /dev/null
printf '{"findings":[{"file":"a.go","cause":"the house style is not followed","next":"follow it"}]}'
exit 1
`))
	if c.Result != "fail" {
		t.Fatalf("result %s, want fail", c.Result)
	}
	if len(c.Findings) != 1 || c.Findings[0].File != "a.go" {
		t.Fatalf("findings %v, want the one it reported", c.Findings)
	}
}

// Status refuses a fail with no finding, so the absence is filled here rather than allowed to
// make a verdict unrepresentable.
func TestANonZeroExitWithNothingToSayStillProducesAFinding(t *testing.T) {
	root := t.TempDir()
	c := run1(t, root, gate(t, root, "silent", `#!/bin/sh
cat > /dev/null
exit 3
`))
	if c.Result != "fail" {
		t.Fatalf("result %s, want fail", c.Result)
	}
	if len(c.Findings) != 1 || !strings.Contains(c.Findings[0].Cause, "exited 3") {
		t.Fatalf("findings %v, want one naming the exit status", c.Findings)
	}
}

// WP4's done-when, in its own words: a modified external gate refuses to run rather than running
// unnoticed. The file is modified after the declaration was written.
func TestAModifiedGateRefusesToRun(t *testing.T) {
	root := t.TempDir()
	d := gate(t, root, "house-linter", passing)
	marker := filepath.Join(root, "it-ran")
	if err := os.WriteFile(filepath.Join(root, d.Path), []byte(`#!/bin/sh
cat > /dev/null
touch `+marker+`
printf '{"findings":[]}'
`), 0o755); err != nil {
		t.Fatal(err)
	}
	c := run1(t, root, d)
	if c.Result != "fail" {
		t.Fatalf("result %s, want fail", c.Result)
	}
	if !strings.Contains(causes(c), "is declared with sha256") {
		t.Fatalf("findings %q, want the mismatch named", causes(c))
	}
	if _, err := os.Stat(marker); err == nil {
		t.Fatal("the modified gate ran")
	}
}

func TestAMissingOrUnrunnableGateIsAFinding(t *testing.T) {
	root := t.TempDir()
	// Declared but not there.
	missing := model.ExternalGate{ID: "ghost", Path: "tools/ghost",
		SHA256: strings.Repeat("a", 64), Phases: []string{"03-implementation"}}
	c := run1(t, root, missing)
	if c.Result != "fail" || !strings.Contains(causes(c), "cannot be read") {
		t.Fatalf("a missing gate is %s: %s", c.Result, causes(c))
	}

	// There, hash matching, and not executable.
	d := gate(t, root, "not-executable", passing)
	if err := os.Chmod(filepath.Join(root, d.Path), 0o644); err != nil {
		t.Fatal(err)
	}
	c = run1(t, root, d)
	if c.Result != "fail" || !strings.Contains(causes(c), "could not be run") {
		t.Fatalf("an unrunnable gate is %s: %s", c.Result, causes(c))
	}
}

// A foreign tool's broken answer must not be able to produce a green.
func TestAnAnswerThatIsNotTheAgreedJSONIsAFail(t *testing.T) {
	for _, tc := range []struct{ name, answer string }{
		{"not json", `this is not json`},
		{"a list rather than an object", `[{"cause":"x"}]`},
		{"a finding with no cause", `{"findings":[{"file":"a.go","next":"do something"}]}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			// A heredoc, so the script carries the answer verbatim whatever it contains.
			d := gate(t, root, "broken", "#!/bin/sh\ncat > /dev/null\ncat <<'EOF'\n"+tc.answer+"\nEOF\n")
			got := run1(t, root, d)
			if got.Result != "fail" {
				t.Fatalf("result %s, want fail", got.Result)
			}
			if !strings.Contains(causes(got), "answered something other than the agreed JSON") {
				t.Fatalf("findings %q, want the contract named", causes(got))
			}
		})
	}
}

// Nothing in the documents bounds a foreign command, and a gate that hangs is not deterministic
// (A75). The limit is not waited out here: what is checked is that the mechanism is the context's
// deadline, by giving the command longer than the test is willing to wait and a timeout it will
// hit.
func TestAGateThatDoesNotFinishIsKilled(t *testing.T) {
	if testing.Short() {
		t.Skip("the timeout test waits")
	}
	root := t.TempDir()
	d := gate(t, root, "sleeper", "#!/bin/sh\ncat > /dev/null\nsleep 30\n")
	// A local deadline, so the suite does not wait sixty seconds for a decision the code makes
	// the same way at any limit.
	old := timeoutForTest
	timeoutForTest = 200_000_000 // 200ms
	defer func() { timeoutForTest = old }()
	c := run1(t, root, d)
	if c.Result != "fail" {
		t.Fatalf("result %s, want fail", c.Result)
	}
	if !strings.Contains(causes(c), "did not finish within") {
		t.Fatalf("findings %q, want the limit named", causes(c))
	}
}

func TestADeclarationThatIsWrongReadsAsAConfigurationMistake(t *testing.T) {
	root := t.TempDir()
	good := gate(t, root, "ok", passing)
	for _, tc := range []struct {
		name string
		d    model.ExternalGate
		want string
	}{
		{"no id", model.ExternalGate{Path: good.Path, SHA256: good.SHA256,
			Phases: []string{"03-implementation"}}, "no id"},
		{"no path", model.ExternalGate{ID: "x", SHA256: good.SHA256,
			Phases: []string{"03-implementation"}}, "no path"},
		{"no hash", model.ExternalGate{ID: "x", Path: good.Path,
			Phases: []string{"03-implementation"}}, "no sha256"},
		{"a hash that is not one", model.ExternalGate{ID: "x", Path: good.Path, SHA256: "abc",
			Phases: []string{"03-implementation"}}, "which is not a sha256"},
		{"an upper case hash", model.ExternalGate{ID: "x", Path: good.Path,
			SHA256: strings.ToUpper(good.SHA256), Phases: []string{"03-implementation"}},
			"which is not a sha256"},
		{"a phase that is not one", model.ExternalGate{ID: "x", Path: good.Path, SHA256: good.SHA256,
			Phases: []string{"03-implementation", "07-afterwards"}}, "which is not a phase"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := run1(t, root, tc.d)
			if c.Result != "fail" {
				t.Fatalf("result %s, want fail", c.Result)
			}
			if !strings.Contains(causes(c), tc.want) {
				t.Fatalf("findings %q, want one containing %q", causes(c), tc.want)
			}
			for _, f := range c.Findings {
				if f.File != model.ProjectFile {
					t.Errorf("the finding names %q, want the configuration file", f.File)
				}
			}
		})
	}
}

// A declaration naming no phase runs nowhere rather than everywhere: the generous reading would
// run foreign code at six phases because a field was forgotten.
func TestADeclarationForNoPhaseRunsNowhereAndSaysSo(t *testing.T) {
	root := t.TempDir()
	d := gate(t, root, "ok", passing)
	d.Phases = nil
	if got := Run(root, "03-implementation", "h", "i", "d", []model.ExternalGate{d}); got != nil {
		t.Fatalf("%d checks for a gate declared for no phase, want none", len(got))
	}
}

func TestAGateRunsOnlyAtThePhasesItNames(t *testing.T) {
	root := t.TempDir()
	d := gate(t, root, "ok", passing, "05-review")
	if got := Run(root, "03-implementation", "h", "i", "d", []model.ExternalGate{d}); got != nil {
		t.Fatalf("%d checks at a phase the gate does not name, want none", len(got))
	}
	if got := Run(root, "05-review", "h", "i", "d", []model.ExternalGate{d}); len(got) != 1 {
		t.Fatalf("%d checks at the phase it names, want 1", len(got))
	}
}

// What the command is given, which is the half of the contract a project's tool reads (A76).
func TestTheCommandIsGivenTheIntentThePhaseAndWhereToLook(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "seen.json")
	d := gate(t, root, "echo", "#!/bin/sh\ncat > "+out+"\nprintf '{\"findings\":[]}'\n")
	if c := run1(t, root, d); c.Result != "pass" {
		t.Fatalf("result %s: %s", c.Result, causes(c))
	}
	b, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	var in Input
	if err := json.Unmarshal(b, &in); err != nil {
		t.Fatalf("the input is not the agreed JSON: %v (%s)", err, b)
	}
	if in.Intent != "git.example/p#1" || in.Phase != "03-implementation" ||
		in.ArtifactsHash != "deadbeef" || in.Root != root ||
		in.PhaseDir != ".xeno/intents/PROJ-1/phases/03-implementation" {
		t.Fatalf("the command was given %+v", in)
	}
}

// Everything from foreign code is text, and a gate.yaml is read by people and by a diff.
func TestAForeignAnswerIsBounded(t *testing.T) {
	root := t.TempDir()
	long := strings.Repeat("x", 2000)
	d := gate(t, root, "verbose", "#!/bin/sh\ncat > /dev/null\n"+
		`printf '{"findings":[{"cause":"`+long+`"}]}'`+"\n")
	c := run1(t, root, d)
	if len(c.Findings) != 1 {
		t.Fatalf("findings %v", c.Findings)
	}
	if n := len(c.Findings[0].Cause); n > 500 {
		t.Fatalf("a cause of %d characters reached the verdict", n)
	}
	// And a finding with no file of its own is named after the gate rather than left blank.
	if c.Findings[0].File != d.Path {
		t.Errorf("the finding names %q, want the gate's own path", c.Findings[0].File)
	}
}
