// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/secrets"
)

// The value planted in an artifact. It is AWS's own documentation key, which is the shape this
// repository already carries on purpose elsewhere, so a reader of this file does not have to
// wonder whether a real credential leaked into a test.
const plantedKey = "AKIAIOSFODNN7EXAMPLE"

// A filter of two patterns rather than the shipped two hundred. The gate's subject is which
// files it opens and what it writes down, and a test against the shipped set would be a test of
// gitleaks' translation as well, which internal/secrets owns.
const testFilter = `patterns:
  - id: aws-access-key
    regex: '\b(?:AKIA|ASIA)[0-9A-Z]{16}\b'
  - id: github-token
    regex: '\bgh[posru]_[A-Za-z0-9]{36,}\b'
paths_never_digested:
  - "**/*.pem"
  - ".env*"
`

// phaseWith lays out a root with a filter in place and the named files in the phase directory,
// and returns the context a gate runs under. The files are given as a map because the point of
// most of these tests is which of them the gate opened.
func secretPhase(t *testing.T, files map[string]string) (Ctx, string) {
	t.Helper()
	root := t.TempDir()
	write := func(rel, body string) {
		p := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write(secrets.Shipped, testFilter)
	dir := model.PhaseDir("PROJ-1", model.Phases[0])
	for name, body := range files {
		write(dir+"/"+name, body)
	}
	return ctxFor(root), dir
}

// The finding says which pattern fired and where, and carries no part of the match. gate.yaml is
// committed and stays in the trail for as long as the intent does, so a value copied into a
// finding is the secret published a second time, in the file nobody thinks of as holding it.
func TestASecretInAnArtifactIsReportedByPatternAndLineAndNeverByValue(t *testing.T) {
	c, dir := secretPhase(t, map[string]string{
		"output.md": "---\nintent: x\n---\n\n## Results\n\nThe bucket key is " + plantedKey + " today.\n",
	})

	got := secret(c)
	if got.Result != "fail" {
		t.Fatalf("result %q, want fail: the planted key is in output.md", got.Result)
	}
	if len(got.Findings) != 1 {
		t.Fatalf("want one finding, got %+v", got.Findings)
	}
	f := got.Findings[0]
	if f.File != dir+"/output.md" {
		t.Errorf("the finding names %q, want the artifact it is in", f.File)
	}
	if !strings.Contains(f.Cause, "aws-access-key") {
		t.Errorf("the cause does not name the pattern that fired: %q", f.Cause)
	}
	if !strings.Contains(f.Cause, "line 7") {
		t.Errorf("the cause does not name the line the pattern fired on: %q", f.Cause)
	}
	if strings.Contains(f.Cause+f.Next, plantedKey) {
		t.Errorf("the finding carries the match itself: %q / %q", f.Cause, f.Next)
	}
}

// A clean phase passes, which is the state the whole trail is in and the reason this gate can
// stop reporting not-implemented without a person deciding anything.
func TestAPhaseWithNoSecretInItPasses(t *testing.T) {
	c, _ := secretPhase(t, map[string]string{
		"output.md":     "---\nintent: x\n---\n\nOrdinary prose about secrets_hash and nothing else.\n",
		"learning.yaml": "intent: x\nlearnings: []\n",
	})

	if got := secret(c); got.Result != "pass" || len(got.Findings) != 0 {
		t.Errorf("result %q with findings %+v, want a clean pass", got.Result, got.Findings)
	}
}

// Every artifact of the phase, not output.md alone. A secret reaches a trail through whatever
// the agent or the person wrote, and `learning.yaml` carries prose the agent supplied as much as
// `output.md` does.
func TestEveryArtifactOfThePhaseIsRead(t *testing.T) {
	c, dir := secretPhase(t, map[string]string{
		"output.md":         "---\nintent: x\n---\n\nclean\n",
		"learning.yaml":     "intent: x\nnote: " + plantedKey + "\n",
		"context.lock.yaml": "intent: x\nfiles: []\n",
	})

	got := secret(c)
	if got.Result != "fail" {
		t.Fatalf("result %q, want fail: the key is in learning.yaml", got.Result)
	}
	if got.Findings[0].File != dir+"/learning.yaml" {
		t.Errorf("the finding names %q, want learning.yaml", got.Findings[0].File)
	}
}

// The two names artifacts_hash excludes are excluded here too. `gate.yaml` is the verdict this
// gate is writing, so reading it would make a finding about the file carrying the finding, and
// `cost.yaml` is figures the runner wrote. Both lie outside the sealed set, which is the set a
// reader compares a verdict against.
func TestTheGateReadsTheSealedSetAndNotTheVerdictBesideIt(t *testing.T) {
	c, _ := secretPhase(t, map[string]string{
		"output.md": "---\nintent: x\n---\n\nclean\n",
		"gate.yaml": "status: green\nnote: " + plantedKey + "\n",
		"cost.yaml": "tokens_in: 1\nnote: " + plantedKey + "\n",
	})

	if got := secret(c); got.Result != "pass" {
		t.Errorf("result %q with findings %+v: the gate read a file outside artifacts_hash",
			got.Result, got.Findings)
	}
}

// `paths_never_digested` is not a skip list. It names files whose content may never be quoted
// into a digest, which is a rule for a writer; read here as an exclusion it would invert into
// its opposite and make the files most likely to hold a key the only ones a gate does not open.
func TestAPathThatIsNeverDigestedIsStillRead(t *testing.T) {
	c, dir := secretPhase(t, map[string]string{
		"output.md": "---\nintent: x\n---\n\nclean\n",
		".env":      "AWS_ACCESS_KEY_ID=" + plantedKey + "\n",
	})

	got := secret(c)
	if got.Result != "fail" {
		t.Fatalf("result %q, want fail: .env matches paths_never_digested and is still an artifact", got.Result)
	}
	if got.Findings[0].File != dir+"/.env" {
		t.Errorf("the finding names %q, want the .env file", got.Findings[0].File)
	}
}

// Section 4: "G-Secret and the digest writer read the same file, which is what makes the claim
// that they share rules verifiable rather than stated." A project pattern is the way to see it:
// nothing in this package knows about it, so it can only fire here because the gate resolved the
// effective filter through the same package the digest writer resolves it through.
func TestAProjectPatternFiresInTheGate(t *testing.T) {
	c, _ := secretPhase(t, map[string]string{
		"output.md": "---\nintent: x\n---\n\nticket HOUSE-1234 here\n",
	})
	p := filepath.Join(c.Root, secrets.Project)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	body := "patterns:\n  - id: house-token\n    regex: '\\bHOUSE-[0-9]{4}\\b'\n"
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	got := secret(c)
	if got.Result != "fail" {
		t.Fatalf("result %q, want fail: the project added the pattern that matches", got.Result)
	}
	if !strings.Contains(got.Findings[0].Cause, "house-token") {
		t.Errorf("the cause does not name the project's pattern: %q", got.Findings[0].Cause)
	}
}

// A repository before its plugin is vendored has no patterns. Section 5 gives that state a name
// — "it is not a failure, and the derived status treats it as neither pass nor fail" — and a
// pass over no patterns would claim a judgement nobody made.
func TestWithoutAFilterSecretIsNotImplemented(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, model.PhaseDir("PROJ-1", model.Phases[0]))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "output.md"), []byte(plantedKey+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	got := secret(ctxFor(root))
	if got.Result != "not-implemented" {
		t.Errorf("result %q, want not-implemented with no filter in force", got.Result)
	}
	if len(got.Findings) != 0 {
		t.Errorf("a check that did not run reports %d finding(s)", len(got.Findings))
	}
}

// Section 7's table puts G-Secret at P0, and the gate list is where a verdict's set of checks
// comes from, so a phase that carries the check has to be every phase.
func TestSecretAppliesFromTheFirstPhase(t *testing.T) {
	for _, p := range model.Phases {
		found := false
		for _, id := range Applicable(p) {
			if id == "G-Secret" {
				found = true
			}
		}
		if !found {
			t.Errorf("%s does not carry G-Secret", p)
		}
	}
}

// The bound of the result, for a reader who was not there. It names the half the digest filter
// does instead and the two things nothing covers, because the state this replaces was one
// `not-implemented` standing for both halves at once.
func TestTheCoverageOfSecretNamesWhatItDoesNotRead(t *testing.T) {
	note := Coverage("G-Secret")
	for _, want := range []string{"digest.md", "evidence/", "between two gate runs"} {
		if !strings.Contains(note, want) {
			t.Errorf("the coverage note does not mention %q: %q", want, note)
		}
	}
	if Coverage("G-Schema") != "" {
		t.Error("a gate with no bound to state carries a note")
	}
}
