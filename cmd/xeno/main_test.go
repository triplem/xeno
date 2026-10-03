// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/runner"
)

// invoke drives the command the way main does, with writers a test can read. The package
// comment promises one staircase of exit codes and this is what asserts it (#110).
func invoke(t *testing.T, args ...string) (code int, out, errw string) {
	t.Helper()
	var o, e bytes.Buffer
	code = run(args, &o, &e)
	return code, o.String(), e.String()
}

// repo is a repository with one intent, ready for a phase to be started in it.
func repo(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, model.IntentDir("PROJ-1"))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	write := func(name, body string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("intent.yaml", "intent: \"git.example/group/proj#1\"\nkey: PROJ-1\nstatus: in-progress\n"+
		"created: \"2026-09-20T10:00:00Z\"\n")
	write("assumptions.yaml", "assumptions: []\n")
	// The shipped templates, so that section set can render.
	src := filepath.Join("..", "..", ".xeno", "plugin", "templates")
	ids, err := os.ReadDir(src)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range ids {
		files, err := os.ReadDir(filepath.Join(src, id.Name()))
		if err != nil {
			t.Fatal(err)
		}
		for _, f := range files {
			b, err := os.ReadFile(filepath.Join(src, id.Name(), f.Name()))
			if err != nil {
				t.Fatal(err)
			}
			dst := filepath.Join(root, ".xeno/plugin/templates", id.Name(), f.Name())
			if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(dst, b, 0o644); err != nil {
				t.Fatal(err)
			}
		}
	}
	return root
}

// 2 is "could not run at all", and it is what an unknown command, a missing argument and an
// unresolvable phase all get. The usage goes to standard error, or a shell redirecting output
// would lose the only thing it says.
func TestTwoIsForWhatCouldNotRunAtAll(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		says string
	}{
		{"nothing at all", nil, "usage:"},
		{"one word that is not init", []string{"phase"}, "usage:"},
		{"a command that does not exist", []string{"gate", "frobnicate"}, "usage:"},
		{"a command needing an intent, without one", []string{"phase", "start", "--phase", "00"}, "--intent is required"},
		{"a phase that does not resolve", []string{"phase", "start", "--intent", "PROJ-1", "--phase", "99"}, "unknown phase"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			code, out, errw := invoke(t, tc.args...)
			if code != 2 {
				t.Errorf("exit %d, want 2", code)
			}
			if !strings.Contains(errw, tc.says) {
				t.Errorf("standard error does not say %q: %q", tc.says, errw)
			}
			if out != "" {
				t.Errorf("something reached standard output: %q", out)
			}
		})
	}
}

// 1 is a refusal with a reason, and the reason goes to standard error. 0 is the command having
// done what was asked.
func TestOneIsARefusalAndZeroIsSuccess(t *testing.T) {
	root := repo(t)

	code, out, errw := invoke(t, "phase", "start", "--root", root, "--intent", "PROJ-1", "--phase", "01", "--no-next")
	if code != 1 {
		t.Errorf("starting P1 before P0 exits %d, want 1", code)
	}
	if !strings.Contains(errw, "refused:") {
		t.Errorf("a refusal does not say so on standard error: %q", errw)
	}
	if out != "" {
		t.Errorf("a refusal wrote to standard output: %q", out)
	}

	if code, _, errw = invoke(t, "phase", "start", "--root", root, "--intent", "PROJ-1", "--phase", "00", "--no-next"); code != 0 {
		t.Fatalf("starting P0 exits %d, want 0: %s", code, errw)
	}
}

// 1 again on a red verdict, and the verdict itself on standard output: it is what the command was
// asked for, not an error.
func TestARedVerdictExitsOneAndPrintsToStandardOutput(t *testing.T) {
	root := repo(t)
	if code, _, e := invoke(t, "phase", "start", "--root", root, "--intent", "PROJ-1", "--phase", "00", "--no-next"); code != 0 {
		t.Fatalf("could not start the phase: %s", e)
	}
	// Finished with no output.md at all, so G-Schema reports it.
	code, out, _ := invoke(t, "phase", "finish", "--root", root, "--intent", "PROJ-1", "--phase", "00", "--no-next")
	if code != 1 {
		t.Errorf("a red verdict exits %d, want 1", code)
	}
	if !strings.Contains(out, "red") {
		t.Errorf("the verdict did not reach standard output: %q", out)
	}
}

// The case section 6 rests the evidence arrangement on, asserted as the branch it is: a
// provisional verdict exits 0, a divergence or a red phase exits 1. The generated wrapper depends
// on the first, and a regression would turn every push out of a P4 into a failed verification,
// which section 6 names as the outcome that would make people stop taking verification seriously.
func TestAProvisionalVerdictVerifiesAsZero(t *testing.T) {
	for _, tc := range []struct {
		name string
		res  runner.VerifyResult
		want int
	}{
		{"nothing at all", runner.VerifyResult{Checked: 0}, 0},
		{"all green", runner.VerifyResult{Checked: 9}, 0},
		{"provisional, which the wrapper must not fail on",
			runner.VerifyResult{Checked: 9, Provisional: []string{"PROJ-1 04-verification"}}, 0},
		{"provisional and green together",
			runner.VerifyResult{Checked: 9, Provisional: []string{"a", "b"}}, 0},
		{"a red phase", runner.VerifyResult{Checked: 9, Red: []string{"PROJ-1 00-intake"}}, 1},
		{"a divergence", runner.VerifyResult{Checked: 9,
			Divergences: []runner.Divergence{{Key: "PROJ-1", Phase: "00-intake", What: "changed"}}}, 1},
		{"provisional beside a red one, which is still a failure",
			runner.VerifyResult{Checked: 9, Provisional: []string{"a"}, Red: []string{"b"}}, 1},
	} {
		if got := verifyCode(&tc.res); got != tc.want {
			t.Errorf("%s: exit %d, want %d", tc.name, got, tc.want)
		}
	}
}

// And the whole command over this repository, which has verdicts and no divergence: 0, with the
// count on standard output.
func TestGateVerifyOverThisRepositoryExitsZero(t *testing.T) {
	code, out, errw := invoke(t, "gate", "verify", "--root", filepath.Join("..", ".."))
	if code != 0 {
		t.Fatalf("exit %d: %s%s", code, out, errw)
	}
	if !strings.Contains(out, "verified ") {
		t.Errorf("the count did not reach standard output: %q", out)
	}
	if errw != "" {
		t.Errorf("a clean verification wrote to standard error: %q", errw)
	}
}

// Every command the usage string names resolves, and every entry in the table is named. The names
// are read out of the usage rather than listed again: a third copy would be a third place to be
// wrong, which is the defect needsKey being a field rather than a switch was meant to avoid.
func TestTheDispatchTableAndTheUsageAgree(t *testing.T) {
	named := map[string]bool{}
	for _, m := range regexp.MustCompile(`(?m)^  xeno (\w+)(?: (\w[\w-]*))?`).FindAllStringSubmatch(usage, -1) {
		name := m[1]
		if m[2] != "" {
			name += " " + m[2]
		}
		named[name] = true
	}
	if len(named) < 15 {
		t.Fatalf("read %d command names out of the usage, which cannot be right: %v", len(named), named)
	}
	// version is the one name the usage carries and the table does not. run answers it before
	// the dispatch, and before the usage check, so that it works in a directory holding nothing:
	// a version is what somebody asks for when nothing else works.
	if _, ok := commands["version"]; ok {
		t.Error("version is in the table, so run answers it twice")
	}
	delete(named, "version")
	for name := range named {
		if _, ok := commands[name]; !ok {
			t.Errorf("the usage names %q and the table has no such command", name)
		}
	}
	for name := range commands {
		if !named[name] {
			t.Errorf("the table carries %q and the usage does not name it", name)
		}
	}
}

// init is one word and everything else is two, which is a property of the dispatch rather than of
// any name.
func TestInitIsOneWordAndTheRestAreTwo(t *testing.T) {
	for name := range commands {
		words := len(strings.Fields(name))
		if name == "init" {
			continue
		}
		if words != 2 {
			t.Errorf("%q is %d words, and only init is one", name, words)
		}
	}
	if _, ok := commands["init"]; !ok {
		t.Error("init is not in the table")
	}
}

// The positional argument is taken off the front before the flag set sees it, because Go's flag
// package stops at the first argument that is not a flag. Absent, it must not become a silent empty
// string that something refuses two layers down without saying why.
func TestAMissingPositionalArgumentIsNotSilent(t *testing.T) {
	root := repo(t)
	code, out, errw := invoke(t, "gate", "approve", "--root", root, "--intent", "PROJ-1",
		"--phase", "00", "--by", "a.person", "--reason", "assessed")
	if code == 0 {
		t.Errorf("approving nothing exited 0:\n%s", out)
	}
	if errw == "" {
		t.Errorf("approving nothing said nothing on standard error")
	}
}

// --export prints the phase environment for a shell to eval, so a suggestion meant for a person
// must not travel with it.
func TestExportPrintsTheEnvironmentAndNothingElse(t *testing.T) {
	root := repo(t)
	code, out, errw := invoke(t, "phase", "start", "--root", root, "--intent", "PROJ-1",
		"--phase", "00", "--export")
	if code != 0 {
		t.Fatalf("exit %d: %s", code, errw)
	}
	if !strings.Contains(out, "XENO_INTENT") || !strings.Contains(out, "XENO_PHASE") {
		t.Errorf("the environment did not reach standard output: %q", out)
	}
	for _, unwanted := range []string{"next:", "xeno section set"} {
		if strings.Contains(out, unwanted) {
			t.Errorf("a suggestion travelled with the environment, which a shell would eval: %q", out)
		}
	}
}

// version is one word, exits 0 and writes to standard output.
func TestVersionIsOneWordAndExitsZero(t *testing.T) {
	code, out, errw := invoke(t, "version")
	if code != 0 {
		t.Errorf("exit %d, want 0", code)
	}
	if !strings.HasPrefix(out, "xeno ") {
		t.Errorf("version wrote %q to standard output", out)
	}
	if errw != "" {
		t.Errorf("version wrote to standard error: %q", errw)
	}
}

// intent start is the one command that names an intent and takes --intent optionally, so
// the dispatch's needsKey switch is off for it and the key is derived instead. This asserts
// both halves at the surface: the command runs without the flag, and the key it chose is
// what the output says it wrote (#179).
func TestIntentStartNeedsNoKeyAndSaysWhichOneItChose(t *testing.T) {
	root := repo(t)
	cfg := filepath.Join(root, ".xeno/config/project.yaml")
	if err := os.MkdirAll(filepath.Dir(cfg), 0o755); err != nil {
		t.Fatal(err)
	}
	body := "tracker:\n  adapter: github\n  project: triplem/xeno\n" +
		"  base_url: https://api.github.com\n"
	if err := os.WriteFile(cfg, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	code, out, errw := invoke(t, "intent", "start", "--root", root, "--for", "176", "--no-next")
	if code != 0 {
		t.Fatalf("exit %d, want 0: %s", code, errw)
	}
	if !strings.Contains(out, "PROJ-2") || !strings.Contains(out, "github.com/triplem/xeno#176") {
		t.Errorf("the output names neither the key nor the id it wrote: %q", out)
	}

	// Twice is a refusal, which is the staircase's 1 with the reason on standard error.
	code, out, errw = invoke(t, "intent", "start", "--root", root, "--intent", "PROJ-2",
		"--for", "176", "--no-next")
	if code != 1 || !strings.Contains(errw, "refused:") {
		t.Errorf("a second start exits %d saying %q, want 1 and a refusal", code, errw)
	}
	if out != "" {
		t.Errorf("a refusal wrote to standard output: %q", out)
	}
}

// Without --for there is nothing to derive the id from, and the refusal names the flag.
// This is the command's one required input, so it is asserted where somebody would type it.
func TestIntentStartWithoutAnIssueIsRefused(t *testing.T) {
	root := repo(t)
	code, _, errw := invoke(t, "intent", "start", "--root", root, "--intent", "NEW-1", "--no-next")
	if code != 1 || !strings.Contains(errw, "--for") {
		t.Errorf("exit %d saying %q, want 1 and a refusal naming --for", code, errw)
	}
}

// The flag is read for every command and acted on by the two that write a phase artifact, so
// it is asserted where somebody types it: one `section set` with it, then `phase finish`
// without, and both files carry the version with nothing edited by hand (#181).
func TestTheToolVersionFlagReachesBothArtifacts(t *testing.T) {
	root := repo(t)
	args := []string{"--root", root, "--intent", "PROJ-1", "--phase", "00", "--no-next"}
	if code, _, e := invoke(t, append([]string{"phase", "start"}, args...)...); code != 0 {
		t.Fatalf("could not start the phase: %s", e)
	}
	set := append([]string{"section", "set", "problem", "--file", writeTemp(t, root, "what is wrong")},
		append(args, "--tool-version", "2.1.276")...)
	if code, _, e := invoke(t, set...); code != 0 {
		t.Fatalf("section set exits %d: %s", code, e)
	}
	if got := frontOf(t, root, "output.md"); got != "2.1.276" {
		t.Errorf("output.md records tool_version %q, want 2.1.276", got)
	}

	// No --tool-version here: the digest takes it from the artifact beside it.
	summary := writeTemp(t, root, "a summary")
	if code, _, e := invoke(t, append([]string{"phase", "finish", "--summary", summary}, args...)...); code == 2 {
		t.Fatalf("phase finish could not run: %s", e)
	}
	if got := frontOf(t, root, "digest.md"); got != "2.1.276" {
		t.Errorf("digest.md records tool_version %q, want 2.1.276", got)
	}
}

// writeTemp puts content in a file under the repository and returns the path, for the flags
// that take one.
func writeTemp(t *testing.T, root, content string) string {
	t.Helper()
	p := filepath.Join(root, "in-"+strings.ReplaceAll(content, " ", "-")+".md")
	if err := os.WriteFile(p, []byte(content+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

// frontOf reads tool_version out of one artifact of PROJ-1's first phase.
func frontOf(t *testing.T, root, name string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(root, model.PhaseDir("PROJ-1", "00-intake"), name))
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range strings.Split(string(b), "\n") {
		if v, ok := strings.CutPrefix(line, "tool_version:"); ok {
			return strings.Trim(strings.TrimSpace(v), `"`)
		}
	}
	return ""
}

// The two channels and their order. Section 7 lists the variable for a value that holds for a
// whole session; a flag is somebody saying it about one run, so the flag wins. Asserted at the
// surface because `parse` is where the precedence lives (#183).
func TestTheFlagBeatsTheHarnessVersionVariable(t *testing.T) {
	root := repo(t)
	t.Setenv(runner.HarnessVersionEnv, "9.9.9")
	args := []string{"--root", root, "--intent", "PROJ-1", "--phase", "00", "--no-next"}
	if code, _, e := invoke(t, append([]string{"phase", "start"}, args...)...); code != 0 {
		t.Fatalf("could not start the phase: %s", e)
	}

	// No flag: the variable is what the runner read.
	set := append([]string{"section", "set", "problem", "--file", writeTemp(t, root, "from the variable")}, args...)
	if code, _, e := invoke(t, set...); code != 0 {
		t.Fatalf("section set exits %d: %s", code, e)
	}
	if got := frontOf(t, root, "output.md"); got != "9.9.9" {
		t.Errorf("without a flag the variable should be recorded, got %q", got)
	}

	// With the flag, over the top of the same variable.
	set = append([]string{"section", "set", "scope", "--file", writeTemp(t, root, "from the flag")},
		append(args, "--tool-version", "2.1.276")...)
	if code, _, e := invoke(t, set...); code != 0 {
		t.Fatalf("section set exits %d: %s", code, e)
	}
	if got := frontOf(t, root, "output.md"); got != "2.1.276" {
		t.Errorf("the flag should beat the variable, got %q", got)
	}
}

// The learning record at the surface, where the phase is optional alone among the commands
// that take one: section 10 owes a record per phase and one more when an intent closes (#195).
func TestLearningRecordTakesThePhaseOptionally(t *testing.T) {
	root := repo(t)
	at := []string{"--root", root, "--intent", "PROJ-1", "--no-next"}
	entry := []string{"--category", "context-rule", "--observation", "o",
		"--proposal", "p", "--target", "t"}

	code, out, errw := invoke(t, append(append([]string{"learning", "record"}, at...),
		append(entry, "--phase", "00")...)...)
	if code != 0 {
		t.Fatalf("with a phase: exit %d: %s", code, errw)
	}
	if !strings.Contains(out, "00-intake") {
		t.Errorf("the output does not name the phase: %q", out)
	}
	if !fm.Exists(filepath.Join(root, model.PhaseDir("PROJ-1", "00-intake"), "learning.yaml")) {
		t.Error("no record at phase level")
	}

	if code, _, errw = invoke(t, append(append([]string{"learning", "record"}, at...),
		"--no-finding")...); code != 0 {
		t.Fatalf("without a phase: exit %d: %s", code, errw)
	}
	if !fm.Exists(filepath.Join(root, model.IntentDir("PROJ-1"), "learning.yaml")) {
		t.Error("no record at intent level")
	}

	// And the staircase on a category the set does not have.
	code, out, errw = invoke(t, append(append([]string{"learning", "record"}, at...),
		"--phase", "01", "--category", "nonsense", "--observation", "o",
		"--proposal", "p", "--target", "t")...)
	if code != 1 || !strings.Contains(errw, "refused:") {
		t.Errorf("a bad category exits %d saying %q, want 1 and a refusal", code, errw)
	}
	if out != "" {
		t.Errorf("a refusal wrote to standard output: %q", out)
	}
}
