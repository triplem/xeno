// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"go.yaml.in/yaml/v3"

	"github.com/triplem/xeno/internal/cost"
	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/gates"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/plugin"
	"github.com/triplem/xeno/internal/rules"
	"github.com/triplem/xeno/internal/secrets"
	"github.com/triplem/xeno/internal/template"
)

const key = "PROJ-1"

// hex64 is a sha256 shaped value for a fixture: sixty four of one hex digit.
func hex64(c byte) string { return strings.Repeat(string(c), 64) }

type fixture struct {
	t    *testing.T
	root string
	r    *Runner
}

func newFixture(t *testing.T) *fixture {
	// Cleared so that every test is hermetic: New reads XENO_HARNESS_VERSION, and a
	// developer whose shell exports it would otherwise see the absence tests pass a value.
	t.Setenv(HarnessVersionEnv, "")
	f := &fixture{t: t, root: t.TempDir()}
	f.reopen()
	// A vendored plugin, because every repository has one after `xeno init --vendor` and
	// section 5 requires plugin_version in every process file. Without it the field is
	// absent and G-Schema says so, which is correct and is asserted on its own below
	// rather than made the condition of every other test (#177).
	f.write(plugin.Dir+"/.claude-plugin/plugin.json", `{"name":"xeno","version":"9.9.9"}`)
	// A context scope, for the same reason as the plugin above: from #217 a P0 cannot be
	// finished without one, so every intent has one and a fixture without one is a case to
	// assert on its own rather than the condition of every other test. The patterns match
	// nothing in a bare fixture, so the information base stays empty unless a test writes
	// its own scope over this.
	f.write(model.PhaseDir(key, model.Phases[0])+"/"+model.ContextScope, "include:\n  - src/**\n")
	f.write(model.IntentDir(key)+"/intent.yaml", "intent: \"git.example/group/proj#1\"\nkey: PROJ-1\nstatus: in-progress\n")
	f.write(model.IntentDir(key)+"/assumptions.yaml", "assumptions: []\n")
	return f
}

// reopen builds the runner over the same tree. Separate from newFixture because New reads the
// environment, so a test that changes what it would read has to build it again.
func (f *fixture) reopen() {
	f.r = New(f.root)
	f.r.Now = func() time.Time { return time.Date(2026, 9, 20, 10, 0, 0, 0, time.UTC) }
}

func (f *fixture) write(rel, content string) {
	f.t.Helper()
	p := filepath.Join(f.root, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		f.t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		f.t.Fatal(err)
	}
}

// output writes a complete, well formed phase result; extra is appended to the
// frontmatter for questions, decisions and evidence.
func (f *fixture) output(phase, extra string) {
	d := model.PhaseDir(key, phase)
	common := "intent: git.example/group/proj#1\nphase: " + phase + "\ncreated: 2026-09-20T10:00:00Z\nrunner_version: 0.1.0-dev\nplugin_version: 0.1.0-dev\n"
	// The hash fields carry sha256 shaped values because Appendix B fixes the shape, and
	// context_hash carries the hash of the lock file this phase was started with, because
	// the gate recomputes it. A fixture with `s`, or with a hash of nothing, would be a
	// finding rather than a phase.
	session := "language: en\nsecrets_hash: " + hex64('1') + "\ncontext_hash: " + f.lockHash(phase) +
		"\nmodel: m\ntool: claude-code\ntool_version: 1\n"
	f.write(d+"/output.md", "---\n"+common+session+"template: t@1\nstrings_hash: "+hex64('3')+
		"\nrules_hash: "+hex64('4')+"\n"+extra+"---\n\n# Result\n")
	f.write(d+"/digest.md", "---\n"+common+session+"---\nsummary\n")
	f.write(d+"/learning.yaml", common+"no_finding: true\n")
}

// lockHash is what Appendix B computes over the lock file the phase was started with,
// which is what the artifact's context_hash has to carry. Before a phase is started there
// is no lock, and a value of the right shape is enough for the tests that never start one.
func (f *fixture) lockHash(phase string) string {
	h, err := hashing.FileHash(filepath.Join(f.root, model.PhaseDir(key, phase), "context.lock.yaml"))
	if err != nil {
		return hex64('2')
	}
	return h
}

// frontField reads one frontmatter field of a phase's output.md as it lies on disk.
func (f *fixture) frontField(phase, field string) string {
	f.t.Helper()
	b, err := os.ReadFile(filepath.Join(f.root, model.PhaseDir(key, phase), "output.md"))
	f.must(err)
	front, _, err := fm.Split(b)
	f.must(err)
	var m map[string]any
	f.must(yaml.Unmarshal(front, &m))
	s, _ := m[field].(string)
	return s
}

func (f *fixture) must2(_ *model.Gate, err error) { f.t.Helper(); f.must(err) }

func (f *fixture) must(err error) {
	f.t.Helper()
	if err != nil {
		f.t.Fatal(err)
	}
}

func (f *fixture) finish(phase string) *model.Gate {
	f.t.Helper()
	g, err := f.r.Finish(key, phase, "")
	f.must(err)
	return g
}

// run does start, write, finish for one phase.
func (f *fixture) run(phase, extra string) *model.Gate {
	f.t.Helper()
	f.must(f.r.Start(key, phase))
	f.output(phase, extra)
	return f.finish(phase)
}

func (f *fixture) hash(phase string) string {
	h, err := f.r.hash(key, phase)
	f.must(err)
	return h
}

func check(g *model.Gate, gate string) *model.Check {
	for i := range g.Checks {
		if g.Checks[i].Gate == gate {
			return &g.Checks[i]
		}
	}
	return nil
}

// ---- WP7: the sequence is enforced

func TestOutOfOrderStartIsRefused(t *testing.T) {
	f := newFixture(t)
	var ref *Refusal
	if err := f.r.Start(key, "01-requirements"); !errors.As(err, &ref) {
		t.Fatalf("P1 before P0 was not refused: %v", err)
	}
}

func TestSecondStartIsRefused(t *testing.T) {
	f := newFixture(t)
	f.must(f.r.Start(key, "00-intake"))
	var ref *Refusal
	if err := f.r.Start(key, "00-intake"); !errors.As(err, &ref) {
		t.Fatalf("a running phase was started twice: %v", err)
	}
}

// A judged phase is not started again. Start writes context.lock.yaml unconditionally, and the
// lock is inside artifacts_hash, so a second start rewrote a sealed artifact and left gate verify
// to report the divergence afterwards. It also overwrote the only record of what the phase was
// given, which is what ChangedSince compares against (#215).
func TestStartIsRefusedWhereThePhaseHasAVerdict(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.run("00-intake", "")

	lock := filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "context.lock.yaml")
	before, err := hashing.FileHash(lock)
	f.must(err)

	var ref *Refusal
	if err := f.r.Start(key, "00-intake"); !errors.As(err, &ref) {
		t.Fatalf("a judged phase was started again: %v", err)
	}
	// The refusal names the way to redo the work, because a refusal that only says no is one
	// somebody works around by deleting something.
	// The alternative is redoing the work in place, and the way out is the phase directory —
	// from #225 the verdict alone is not enough, because the artifact left behind is itself a
	// phase under way.
	if !strings.Contains(ref.Reason, "section set") ||
		!strings.Contains(ref.Reason, model.PhaseDir(key, "00-intake")) {
		t.Errorf("the refusal does not name the alternative: %s", ref.Reason)
	}
	after, err := hashing.FileHash(lock)
	f.must(err)
	if after != before {
		t.Error("the refused start rewrote the lock it was refused for")
	}
}

// Starting a phase over stays deliberate and stays possible, and from #225 it takes the phase
// directory rather than the verdict alone: an artifact left behind is itself a phase under way,
// so removing gate.yaml now lands on the second refusal instead of the first.
func TestStartingOverTakesThePhaseDirectory(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.run("00-intake", "")
	dir := filepath.Join(f.root, model.PhaseDir(key, "00-intake"))

	f.must(os.Remove(filepath.Join(dir, "gate.yaml")))
	err := f.r.Start(key, "00-intake")
	if !isRefusal(err) || !strings.Contains(err.Error(), "already under way") {
		t.Fatalf("removing the verdict alone should now leave a phase under way: %v", err)
	}

	f.must(os.RemoveAll(dir))
	f.must(f.r.Start(key, "00-intake"))
}

func (f *fixture) redIntake() *model.Gate {
	f.t.Helper()
	f.must(f.r.Start(key, "00-intake"))
	f.output("00-intake", "")
	os.Remove(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "learning.yaml"))
	g := f.finish("00-intake")
	if g.Status != "red" {
		f.t.Fatalf("expected red, got %s", g.Status)
	}
	return g
}

func TestRedPredecessorBlocksStart(t *testing.T) {
	f := newFixture(t)
	f.redIntake()
	var ref *Refusal
	if err := f.r.Start(key, "01-requirements"); !errors.As(err, &ref) {
		t.Fatalf("a phase started on a red predecessor: %v", err)
	}
}

func TestDecidedPredecessorAllowsStart(t *testing.T) {
	for _, decision := range []string{"approved", "overridden"} {
		f := newFixture(t)
		g := f.redIntake()
		d := &model.DecisionOnFinding{Type: decision, By: "m.example", At: "2026-09-20T10:00:00Z", Against: g.ArtifactsHash, Reason: "fixture"}
		if decision == "overridden" {
			d.Obligation = "open"
		}
		check(g, "G-Learning").Findings[0].Decision = d
		f.must(fm.WriteYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "gate.yaml"), g))
		if g = f.finish("00-intake"); g.Status != decision {
			t.Fatalf("expected %s, got %s", decision, g.Status)
		}
		if err := f.r.Start(key, "01-requirements"); err != nil {
			t.Fatalf("a %s predecessor blocked the next phase: %v", decision, err)
		}
	}
}

func TestChangedPredecessorIsRefused(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	f.write(model.PhaseDir(key, "00-intake")+"/output.md", "edited\n")
	var ref *Refusal
	if err := f.r.Start(key, "01-requirements"); !errors.As(err, &ref) {
		t.Fatalf("a predecessor changed after its verdict was accepted: %v", err)
	}
}

// ---- WP1: what is sealed is never rewritten

func TestNoCommandButFinishTouchesAFinishedPhase(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	before := map[string]string{}
	for _, p := range model.Phases[:4] {
		before[p] = f.hash(p)
	}
	_, _ = f.r.GateRun(key, "03-implementation")
	_, _ = f.r.Status(key)
	f.must(f.r.Start(key, "04-verification"))
	for _, p := range model.Phases[:4] {
		if f.hash(p) != before[p] {
			t.Fatalf("%s changed without phase finish", p)
		}
	}
}

// ---- WP6: evidence comes from CI, pulled, never pushed

const pendingTest = "evidence:\n  - kind: test-report\n    job: unit\n"

func (f *fixture) pipeline(result string) string {
	dir := filepath.Join(f.t.TempDir(), "artifacts")
	_ = os.MkdirAll(dir, 0o755)
	os.WriteFile(filepath.Join(dir, "junit.xml"), []byte("<testsuite failures=\"0\"/>\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(
		"- kind: test-report\n  job: unit\n  result: "+result+"\n  file: junit.xml\n  pipeline: \"4711\"\n  commit: abc123\n"), 0o644)
	return dir
}

func TestPendingEvidenceFinishesWithoutWaiting(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	f.must(f.r.Start(key, "04-verification"))
	f.output("04-verification", pendingTest)
	g := f.finish("04-verification")
	if g.Status != "provisional" || check(g, "G-Evidence").Result != "pending" {
		t.Fatalf("expected provisional with pending evidence, got %s", g.Status)
	}
}

func TestNextStartAttachesAndCarriesTheVerdictForward(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	f.must(f.r.Start(key, "04-verification"))
	f.output("04-verification", pendingTest)
	f.finish("04-verification")
	sealed := f.hash("04-verification")

	var ref *Refusal
	if err := f.r.Start(key, "05-review"); !errors.As(err, &ref) {
		t.Fatalf("P5 started while P4 still waited for the pipeline: %v", err)
	}
	f.r.EvidenceFrom = f.pipeline("pass")
	f.must(f.r.Start(key, "05-review"))

	g, err := f.r.readGate(key, "04-verification")
	f.must(err)
	if g.Status != "green" {
		t.Fatalf("predecessor verdict not carried forward: %s", g.Status)
	}
	if f.hash("04-verification") != sealed || g.ArtifactsHash != sealed {
		t.Fatal("attaching changed the artifacts_hash of the sealed phase")
	}
	var att []model.Attached
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "04-verification"), "evidence/attached.yaml"), &att))
	if len(att) != 1 || att[0].Pipeline != "4711" || att[0].Commit != "abc123" {
		t.Fatalf("provenance not recorded: %+v", att)
	}
}

func TestGateRunAttachesForTheLastPhase(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:5] {
		f.run(p, "")
	}
	f.must(f.r.Start(key, "05-review"))
	f.output("05-review", pendingTest)
	if g := f.finish("05-review"); g.Status != "provisional" {
		t.Fatalf("expected provisional, got %s", g.Status)
	}
	f.r.EvidenceFrom = f.pipeline("pass")
	g, err := f.r.GateRun(key, "05-review")
	f.must(err)
	if g.Status != "green" {
		t.Fatalf("gate run did not attach for P5: %s", g.Status)
	}
}

func TestTamperedEvidenceIsRejected(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	f.must(f.r.Start(key, "04-verification"))
	f.output("04-verification", pendingTest)
	f.finish("04-verification")
	f.r.EvidenceFrom = f.pipeline("pass")
	_, err := f.r.GateRun(key, "04-verification")
	f.must(err)
	f.write(model.PhaseDir(key, "04-verification")+"/evidence/junit.xml", "<testsuite failures=\"3\"/>\n")
	g, err := f.r.GateRun(key, "04-verification")
	f.must(err)
	if g.Status != "red" {
		t.Fatalf("an altered attachment passed: %s", g.Status)
	}
}

// ---- WP5 and WP7: open questions are resolved, and only P5 insists

// The options carry their consequence because section 8 asks for one and #229 made the gate
// read it. These tests are about a question being resolved and not about its shape, so the
// fixture is the shape the section asks for rather than the least the gate once accepted.
const question = "open_questions:\n  - key: Q-1\n    text: which error behaviour?\n    options:\n      - text: fail fast\n        consequence: the caller retries\n        recommended: true\n      - text: retry\n        consequence: the caller never sees it\n      - text: free entry\n        free: true\n"

func TestUnresolvedQuestionTurnsP5RedAndLeavesP1Green(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	if g := f.run("01-requirements", question); g.Status != "green" {
		t.Fatalf("an open question must not hold up P1: %s", g.Status)
	}
	for _, p := range model.Phases[2:5] {
		f.run(p, "")
	}
	g := f.run("05-review", "")
	if g.Status != "red" || check(g, "G-Questions").Result != "fail" {
		t.Fatalf("an unresolved question passed P5: %s", g.Status)
	}
}

func TestQuestionResolvedTwoPhasesLater(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	f.run("01-requirements", question)
	f.run("02-design", "")
	f.run("03-implementation", "decisions:\n  - id: D-1\n    resolves: Q-1\n    chosen: fail fast\n    rationale: the caller retries\n    decided_by: m.example\n")
	f.run("04-verification", "")
	if g := f.run("05-review", ""); check(g, "G-Questions").Result != "pass" {
		t.Fatal("a decision two phases later was not recognised as resolving the question")
	}
}

func TestQuestionResolvedByConfirmedAssumption(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	f.run("01-requirements", question)
	f.write(model.IntentDir(key)+"/assumptions.yaml",
		"assumptions:\n  - id: A-001\n    assumption: fail fast is acceptable\n    origin: user-input\n"+
			"    confidence: medium\n    status: confirmed\n    confirmed_by: m.example\n    resolves: Q-1\n")
	for _, p := range model.Phases[2:5] {
		f.run(p, "")
	}
	if g := f.run("05-review", ""); check(g, "G-Questions").Result != "pass" {
		t.Fatal("a confirmed assumption did not resolve the question")
	}
}

func TestQuestionShape(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	g := f.run("01-requirements", "open_questions:\n  - key: Q-9\n    text: bare\n")
	if g.Status != "red" {
		t.Fatal("a question without options passed")
	}
	f2 := newFixture(t)
	f2.run("00-intake", "")
	if g := f2.run("01-requirements", "open_questions:\n  - key: Q-9\n    text: bare\n    no_options: true\n"); g.Status != "green" {
		t.Fatalf("an honest no_options was rejected: %s", g.Status)
	}
}

// A phase finished without its digest is not complete, and G-Schema is the gate that
// says so: no other one looks at the file.
func TestMissingDigestIsAStructuralFailure(t *testing.T) {
	f := newFixture(t)
	f.must(f.r.Start(key, "00-intake"))
	f.output("00-intake", "")
	f.must(os.Remove(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "digest.md")))
	g := f.finish("00-intake")
	if g.Status != "red" {
		t.Fatalf("a phase without a digest passed: %s", g.Status)
	}
	if c := check(g, "G-Schema"); c == nil || len(c.Findings) != 1 || !strings.Contains(c.Findings[0].Cause, "digest.md is missing") {
		t.Fatalf("G-Schema did not name the missing digest: %+v", check(g, "G-Schema"))
	}
}

// ---- WP1: decisions survive an unchanged finding and not a changed one

func TestDecisionCarriesForwardOnlyForTheSameFinding(t *testing.T) {
	f := newFixture(t)
	f.must(f.r.Start(key, "00-intake"))
	f.output("00-intake", "")
	f.write(model.PhaseDir(key, "00-intake")+"/stray.txt", "x")
	g := f.finish("00-intake")
	ch := check(g, "G-Schema")
	ch.Findings[0].Decision = &model.DecisionOnFinding{Type: "approved", By: "m.example", At: "2026-09-20T10:00:00Z", Against: g.ArtifactsHash, Reason: "fixture"}
	f.must(fm.WriteYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "gate.yaml"), g))

	if g = f.finish("00-intake"); g.Status != "approved" {
		t.Fatalf("decision lost on an unchanged finding: %s", g.Status)
	}
	os.Rename(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "stray.txt"), filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "other.txt"))
	if g = f.finish("00-intake"); g.Status != "red" {
		t.Fatalf("a decision covered a different finding: %s", g.Status)
	}
}

// ---- computed status and freshness

func TestStatusIsComputedAndStalenessDetected(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	f.run("01-requirements", "")
	f.must(f.r.Start(key, "02-design"))

	// Reopen P1 by hand: new content, new verdict.
	f.write(model.PhaseDir(key, "01-requirements")+"/output.md", strings.Replace(
		readFile(t, filepath.Join(f.root, model.PhaseDir(key, "01-requirements"), "output.md")), "# Result", "# Result, revised", 1))
	f.finish("01-requirements")

	states, err := f.r.Status(key)
	f.must(err)
	if states[2].State != "running" || !states[2].Stale {
		t.Fatalf("staleness not computed: %+v", states[2])
	}
	f.output("02-design", "")
	if g := f.finish("02-design"); check(g, "G-Freshness").Result != "fail" {
		t.Fatal("G-Freshness passed against a changed predecessor")
	}
}

func readFile(t *testing.T, p string) string {
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

// ---- CI: recompute and compare, never write

func TestVerifyMatchesAndWritesNothing(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	gatePath := filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "gate.yaml")
	before := readFile(t, gatePath)
	res, err := f.r.Verify("")
	f.must(err)
	if res.Checked != 1 || len(res.Divergences) != 0 || len(res.Red) != 0 {
		t.Fatalf("unexpected verify result %+v", res)
	}
	if readFile(t, gatePath) != before {
		t.Fatal("verify wrote gate.yaml")
	}
}

func TestVerifyDetectsAChangeAfterTheVerdict(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	f.write(model.PhaseDir(key, "00-intake")+"/digest.md", "edited after the verdict\n")
	res, err := f.r.Verify(key)
	f.must(err)
	if len(res.Divergences) != 1 {
		t.Fatalf("a change after the verdict went unnoticed: %+v", res)
	}
}

func TestVerifyReportsProvisionalWithoutFailing(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	f.must(f.r.Start(key, "04-verification"))
	f.output("04-verification", pendingTest)
	f.finish("04-verification")
	res, err := f.r.Verify(key)
	f.must(err)
	if len(res.Provisional) != 1 || len(res.Divergences) != 0 || len(res.Red) != 0 {
		t.Fatalf("pending evidence was not reported as provisional: %+v", res)
	}
}

func TestVerifyOnARepositoryWithoutIntents(t *testing.T) {
	r := New(t.TempDir())
	res, err := r.Verify("")
	if err != nil || res.Checked != 0 {
		t.Fatalf("a repository without intents must verify cleanly: %v %+v", err, res)
	}
}

func TestMissingIntentIdIsNamedNotGuessed(t *testing.T) {
	f := newFixture(t)
	f.write(model.IntentDir(key)+"/intent.yaml", "intent: github.com/o/r #1\nkey: PROJ-1\n")
	err := f.r.Start(key, "00-intake")
	var ref *Refusal
	if !errors.As(err, &ref) || !strings.Contains(ref.Reason, "comment") {
		t.Fatalf("an id cut off by a YAML comment was not named: %v", err)
	}
}

func TestKeyMustMatchItsDirectory(t *testing.T) {
	f := newFixture(t)
	f.write(model.IntentDir(key)+"/intent.yaml", "intent: \"github.com/o/r#9\"\nkey: PROJ-9\n")
	var ref *Refusal
	if err := f.r.Start(key, "00-intake"); !errors.As(err, &ref) {
		t.Fatalf("an intent.yaml naming another key was accepted: %v", err)
	}
}

// ---- WP7: the commands that write a decision

// findingOf returns the first failing finding of a phase, which is what a person would
// be deciding about.
func (f *fixture) findingOf(phase string) string {
	f.t.Helper()
	g, err := f.r.readGate(key, phase)
	f.must(err)
	for _, ch := range g.Checks {
		if len(ch.Findings) > 0 {
			return ch.Findings[0].ID
		}
	}
	f.t.Fatal("no finding to decide about")
	return ""
}

// redPhase produces a phase with one failing finding: a file Xeno did not write.
func (f *fixture) redPhase(phase string) string {
	f.t.Helper()
	f.must(f.r.Start(key, phase))
	f.output(phase, "")
	f.write(model.PhaseDir(key, phase)+"/stray.txt", "x")
	if g := f.finish(phase); g.Status != "red" {
		f.t.Fatalf("expected a red phase, got %s", g.Status)
	}
	return f.findingOf(phase)
}

func TestApprovalTurnsRedIntoApprovedAndRecordsWhatItJudged(t *testing.T) {
	f := newFixture(t)
	id := f.redPhase("00-intake")
	before := f.hash("00-intake")

	g, err := f.r.Decide(key, "00-intake", id, "approved", "a person", "assessed and accepted")
	f.must(err)
	if g.Status != "approved" {
		t.Fatalf("status %s, want approved", g.Status)
	}
	d := findFinding(g, id).Decision
	if d == nil || d.By != "a person" || d.Reason != "assessed and accepted" {
		t.Fatalf("decision not recorded: %+v", d)
	}
	if d.Against != before {
		t.Fatalf("against is %s, want the hash that was judged, %s", d.Against, before)
	}
	if d.Obligation != "" {
		t.Fatal("an approval owes nothing and must carry no obligation")
	}
	// gate.yaml is outside artifacts_hash, so deciding must not move it.
	if after := f.hash("00-intake"); after != before {
		t.Fatalf("deciding changed the artifacts_hash, %s to %s", before, after)
	}
}

func TestOverrideCarriesAnOpenObligationUntilItIsClosed(t *testing.T) {
	f := newFixture(t)
	id := f.redPhase("00-intake")

	g, err := f.r.Decide(key, "00-intake", id, "overridden", "a person", "production incident")
	f.must(err)
	if g.Status != "overridden" {
		t.Fatalf("status %s, want overridden", g.Status)
	}
	if o := findFinding(g, id).Decision.Obligation; o != "open" {
		t.Fatalf("obligation %q, want open", o)
	}
	g, err = f.r.CloseObligation(key, "00-intake", id)
	f.must(err)
	if o := findFinding(g, id).Decision.Obligation; o != "closed" {
		t.Fatalf("obligation %q, want closed", o)
	}
	if g.Status != "overridden" {
		t.Fatalf("status %s: closing what was owed does not unmake the override", g.Status)
	}
}

// D-6. The decision carries the hash it was made on, so there has to be one.
func TestDecidingOnAStaleVerdictIsRefused(t *testing.T) {
	f := newFixture(t)
	id := f.redPhase("00-intake")
	f.write(model.PhaseDir(key, "00-intake")+"/output.md", "---\nintent: x\n---\nedited outside the runner\n")

	_, err := f.r.Decide(key, "00-intake", id, "approved", "a person", "why not")
	var refusal *Refusal
	if !errors.As(err, &refusal) {
		t.Fatalf("expected a refusal, got %v", err)
	}
	if !strings.Contains(err.Error(), "changed after it was judged") {
		t.Fatalf("the refusal does not say why: %v", err)
	}
}

// An obligation is closed after the artifacts exist, so the phase has moved by then.
// The check that refuses a stale decision must not refuse this.
func TestClosingAnObligationSurvivesAChangedPhase(t *testing.T) {
	f := newFixture(t)
	id := f.redPhase("00-intake")
	f.must2(f.r.Decide(key, "00-intake", id, "overridden", "a person", "incident"))
	f.write(model.PhaseDir(key, "00-intake")+"/output.md", "---\nintent: x\n---\nthe artifacts that were owed\n")

	g, err := f.r.CloseObligation(key, "00-intake", id)
	f.must(err)
	if o := findFinding(g, id).Decision.Obligation; o != "closed" {
		t.Fatalf("obligation %q, want closed", o)
	}
}

func TestADecisionIsMadeOnceAndNeedsAPersonAndAReason(t *testing.T) {
	f := newFixture(t)
	id := f.redPhase("00-intake")

	if _, err := f.r.Decide(key, "00-intake", id, "approved", "", "a reason"); err == nil {
		t.Fatal("a decision without a person was accepted")
	}
	if _, err := f.r.Decide(key, "00-intake", id, "approved", "a person", ""); err == nil {
		t.Fatal("a decision without a reason was accepted")
	}
	f.must2(f.r.Decide(key, "00-intake", id, "approved", "a person", "assessed"))
	if _, err := f.r.Decide(key, "00-intake", id, "overridden", "somebody else", "changed my mind"); err == nil {
		t.Fatal("a second decision replaced the first")
	}
}

// ---- WP7: an intent that is dropped rather than merged

func (f *fixture) intentLearning(body string) {
	f.write(model.IntentDir(key)+"/learning.yaml",
		"intent: git.example/group/proj#1\ncreated: 2026-09-20T10:00:00Z\n"+
			"runner_version: 0.1.0-dev\nplugin_version: 0.1.0-dev\n"+body)
}

func TestClosingAnAbandonedIntentRecordsItAndJudgesIt(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	f.intentLearning("no_finding: true\n")

	g, err := f.r.IntentClose(key, "the requirement was withdrawn")
	f.must(err)
	if g.Status != "green" {
		t.Fatalf("status %s, want green: the reason and the record are both there", g.Status)
	}
	if g.Phase != "" {
		t.Fatalf("an intent level verdict carries no phase, got %q", g.Phase)
	}
	if len(g.Checks) != 1 || g.Checks[0].Gate != "G-Complete" {
		t.Fatalf("expected one G-Complete check, got %+v", g.Checks)
	}

	var in model.Intent
	f.must(fm.ReadYAML(filepath.Join(f.root, model.IntentDir(key), "intent.yaml"), &in))
	if in.Status != "abandoned" || in.Reason != "the requirement was withdrawn" {
		t.Fatalf("intent.yaml not recorded: %+v", in)
	}
}

// The gate reads the files, not the argument. Without the closing record it is red, and
// the intent is abandoned all the same: the record says what is missing.
func TestClosingWithoutTheLearningRecordIsRedRatherThanRefused(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")

	g, err := f.r.IntentClose(key, "dropped")
	f.must(err)
	if g.Status != "red" {
		t.Fatalf("status %s, want red", g.Status)
	}
	var in model.Intent
	f.must(fm.ReadYAML(filepath.Join(f.root, model.IntentDir(key), "intent.yaml"), &in))
	if in.Status != "abandoned" {
		t.Fatal("the intent was not recorded as abandoned")
	}
}

func TestClosingNeedsAReasonAndHappensOnce(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	f.intentLearning("no_finding: true\n")

	if _, err := f.r.IntentClose(key, "  "); err == nil {
		t.Fatal("an intent was abandoned without a reason")
	}
	f.must2(f.r.IntentClose(key, "dropped"))
	if _, err := f.r.IntentClose(key, "dropped again"); err == nil {
		t.Fatal("an abandoned intent was abandoned a second time")
	}
}

// The intent level hash covers the files lying directly in the intent directory and
// does not descend, which is what keeps phases/ out: each phase already has a verdict.
func TestTheIntentHashDoesNotDescendIntoPhases(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	f.intentLearning("no_finding: true\n")
	g, err := f.r.IntentClose(key, "dropped")
	f.must(err)

	f.write(model.PhaseDir(key, "00-intake")+"/output.md", "---\nintent: x\n---\nchanged\n")
	after, err := hashing.DirHash(f.root, model.IntentDir(key), hashing.IntentExcluded)
	f.must(err)
	if after != g.ArtifactsHash {
		t.Fatalf("changing a phase moved the intent hash, %s to %s", g.ArtifactsHash, after)
	}
}

// ---- WP2: the renderer reached through a command

// templated copies the shipped set into the fixture, so that these tests exercise the
// set this repository ships rather than one written to make them pass.
func (f *fixture) templated() {
	f.t.Helper()
	src := filepath.Join("..", "..", ".xeno", "plugin", "templates")
	ids, err := os.ReadDir(src)
	f.must(err)
	for _, id := range ids {
		files, err := os.ReadDir(filepath.Join(src, id.Name()))
		f.must(err)
		for _, file := range files {
			b, err := os.ReadFile(filepath.Join(src, id.Name(), file.Name()))
			f.must(err)
			f.write(filepath.Join(".xeno/plugin/templates", id.Name(), file.Name()), string(b))
		}
	}
}

func TestStartRecordsWhichTemplateItWillRenderFrom(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))

	var lock model.ContextLock
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "context.lock.yaml"), &lock))
	if lock.TemplateSource != "plugin" {
		t.Fatalf("template_source is %q, want plugin", lock.TemplateSource)
	}
}

// A repository without a vendored plugin records nothing and finds out at the first
// section write, which is where it matters.
func TestStartWithoutATemplateRecordsNothingAndDoesNotRefuse(t *testing.T) {
	f := newFixture(t)
	f.must(f.r.Start(key, "00-intake"))

	var lock model.ContextLock
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "context.lock.yaml"), &lock))
	if lock.TemplateSource != "" {
		t.Fatalf("template_source is %q, want nothing", lock.TemplateSource)
	}
	if _, err := f.r.SectionSet(key, "00-intake", "problem", "x"); err == nil {
		t.Fatal("a section was written with no template to render it from")
	}
}

func TestSectionSetRendersAnchorsTheCallerNeverWrites(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))

	_, err := f.r.SectionSet(key, "00-intake", "problem", "The thing that is wrong.")
	f.must(err)
	_, err = f.r.SectionSet(key, "00-intake", "scope", "What is being done about it.")
	f.must(err)

	b, err := os.ReadFile(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "output.md"))
	f.must(err)
	front, body, err := fm.Split(b)
	f.must(err)

	for _, want := range []string{"problem", "scope", "context-rationale"} {
		if !strings.Contains(string(body), template.AnchorPrefix+want+" -->") {
			t.Errorf("required section %s has no anchor", want)
		}
	}
	if strings.Contains(string(body), "open-questions") {
		t.Error("an empty optional section was rendered")
	}
	sections := template.Parse(string(body))
	if sections["problem"] != "The thing that is wrong." {
		t.Errorf("the first section did not survive the second write: %q", sections["problem"])
	}
	if !strings.Contains(string(front), "template: intake@1.0.0") {
		t.Errorf("the frontmatter does not name the template:\n%s", front)
	}
	if !strings.Contains(string(front), "strings_hash: ") {
		t.Error("the frontmatter carries no strings_hash")
	}
	// The order of section 5, not the alphabetical order a map would give.
	if i, j := strings.Index(string(front), "intent:"), strings.Index(string(front), "created:"); i > j {
		t.Errorf("the frontmatter is not in the order section 5 lists:\n%s", front)
	}
}

// The field the runner knows, written rather than left to a hand that edits the
// frontmatter afterwards. The second write is what shows the value does not drift inside a
// phase: the lock is written once at start and nothing refreshes it.
func TestSectionSetWritesTheContextHashOfTheLockBesideIt(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))

	_, err := f.r.SectionSet(key, "00-intake", "problem", "The thing that is wrong.")
	f.must(err)
	first := f.frontField("00-intake", "context_hash")
	if first != f.lockHash("00-intake") {
		t.Fatalf("context_hash is %q, want the hash of the lock beside it", first)
	}
	_, err = f.r.SectionSet(key, "00-intake", "scope", "What is being done about it.")
	f.must(err)
	if second := f.frontField("00-intake", "context_hash"); second != first {
		t.Errorf("context_hash moved within the phase: %q then %q", first, second)
	}
}

// The walk section 5 supports, without a hand editing the frontmatter: the gate still
// reports what A35 leaves out and reports nothing about context_hash.
func TestTheSupportedWalkIsNotRedOnTheContextHash(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	for section, content := range map[string]string{
		"problem":           "The thing that is wrong.",
		"scope":             "What is being done about it.",
		"context-rationale": "Why these files.",
	} {
		_, err := f.r.SectionSet(key, "00-intake", section, content)
		f.must(err)
	}

	g, err := f.r.evaluate(key, "00-intake")
	f.must(err)
	for _, c := range g.Checks {
		for _, fd := range c.Findings {
			if strings.Contains(fd.Cause, "context_hash") || strings.Contains(fd.Next, "context_hash") {
				t.Errorf("%s reports on context_hash: %s / %s", c.Gate, fd.Cause, fd.Next)
			}
		}
	}
}

func TestAnUnknownSectionIsRefusedWithWhatThereIs(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))

	_, err := f.r.SectionSet(key, "00-intake", "not-a-section", "x")
	if err == nil {
		t.Fatal("an unknown section was accepted")
	}
	if !strings.Contains(err.Error(), "context-rationale") {
		t.Errorf("the refusal does not say what the template has: %v", err)
	}
}

// ---- WP6: a scan report from the pipeline attaches

// scanArtifact builds the directory the scan workflows upload: the report, the
// database metadata that travels with it, and the manifest naming both.
func (f *fixture) scanArtifact(result string) string {
	dir := filepath.Join(f.t.TempDir(), "scan-trivy")
	_ = os.MkdirAll(dir, 0o755)
	os.WriteFile(filepath.Join(dir, "trivy.json"), []byte("{\"Results\":[]}\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "db-metadata.json"),
		[]byte("{\"UpdatedAt\":\"2026-09-25T06:17:00Z\"}\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(
		"- kind: scan\n  job: trivy\n  result: "+result+"\n  file: trivy.json\n"+
			"  pipeline: \"4711\"\n  commit: abc123\n"+
			"- kind: other\n  job: trivy-db\n  file: db-metadata.json\n"+
			"  pipeline: \"4711\"\n  commit: abc123\n"), 0o644)
	return dir
}

const pendingScan = "evidence:\n  - kind: scan\n    job: trivy\n  - kind: other\n    job: trivy-db\n"

func TestAScanReportFromThePipelineAttachesWithItsDatabaseAge(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	f.must(f.r.Start(key, "04-verification"))
	f.output("04-verification", pendingScan)
	if g := f.finish("04-verification"); check(g, "G-Evidence").Result != "pending" {
		t.Fatalf("a declared scan that the pipeline has not produced is not pending")
	}

	// A failing scan is still evidence: the job's own result is recorded, and no gate
	// judges it.
	f.r.EvidenceFrom = f.scanArtifact("fail")
	f.must(f.r.Start(key, "05-review"))

	g, err := f.r.readGate(key, "04-verification")
	f.must(err)
	if g.Status != "green" || check(g, "G-Evidence").Result != "pass" {
		t.Fatalf("attached scan did not resolve: %s, %s", g.Status, check(g, "G-Evidence").Result)
	}

	var att []model.Attached
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "04-verification"),
		"evidence", "attached.yaml"), &att))
	if len(att) != 2 {
		t.Fatalf("expected the report and its database metadata, got %d", len(att))
	}
	if att[0].Result != "fail" {
		t.Fatalf("the scanner's own result was not carried: %q", att[0].Result)
	}
	for _, a := range att {
		if a.SHA256 == "" || a.Path == "" {
			t.Fatalf("attached item is not bound: %+v", a)
		}
	}
}

// The manifest above is written by the workflows, not by the runner, so the shape the
// test relies on is read back out of them rather than assumed to still match.
func TestTheScanWorkflowsWriteTheManifestThisExpects(t *testing.T) {
	for _, c := range []struct{ file, job, report string }{
		{"trivy.yml", "trivy", "trivy.json"},
		{"semgrep.yml", "semgrep", "semgrep.json"},
	} {
		b, err := os.ReadFile(filepath.Join("..", "..", ".github", "workflows", c.file))
		if err != nil {
			t.Fatal(err)
		}
		for _, want := range []string{"kind: scan", "job: " + c.job, "file: " + c.report} {
			if !strings.Contains(string(b), want) {
				t.Errorf("%s no longer writes %q", c.file, want)
			}
		}
	}
}

// ---- WP7: the next step of the working sequence, read off the state

// redByRendering finishes P0 from the renderer rather than from the fixture's
// frontmatter, which is red on G-Schema: section set writes the fields the runner knows
// and leaves the rest out, which is A35. It is the shortest honest way to a red verdict.
func (f *fixture) redByRendering() *model.Gate {
	f.t.Helper()
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	for _, id := range []string{"problem", "scope", "context-rationale"} {
		_, err := f.r.SectionSet(key, "00-intake", id, "written")
		f.must(err)
	}
	g := f.finish("00-intake")
	if g.Status != "red" {
		f.t.Fatalf("this fixture is meant to be red, got %s", g.Status)
	}
	return g
}

func TestTheSuggestionAsksForTheSectionsThatAreMissing(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))

	s := f.r.Next(key)
	if !strings.Contains(s.Text, "problem") || !strings.Contains(s.Command, "xeno section set problem") {
		t.Fatalf("a running phase did not name a missing section: %+v", s)
	}
	if !strings.Contains(s.Command, "--phase 00") || !strings.Contains(s.Command, "--intent "+key) {
		t.Errorf("the command cannot be run as printed: %q", s.Command)
	}

	for _, id := range []string{"problem", "scope", "context-rationale"} {
		_, err := f.r.SectionSet(key, "00-intake", id, "written")
		f.must(err)
	}
	if s := f.r.Next(key); !strings.Contains(s.Command, "xeno phase finish") {
		t.Fatalf("a complete phase was not sent to be judged: %+v", s)
	}
}

func TestTheSuggestionMovesOnAndBackWithTheVerdict(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")

	s := f.r.Next(key)
	if !strings.Contains(s.Command, "xeno phase start") || !strings.Contains(s.Command, "--phase 01") {
		t.Fatalf("a decided phase did not point at its successor: %+v", s)
	}

	f.write(model.PhaseDir(key, "00-intake")+"/output.md", "edited\n")
	if s := f.r.Next(key); !strings.Contains(s.Command, "xeno phase finish") ||
		!strings.Contains(s.Text, "changed after its verdict") {
		t.Fatalf("a changed phase was not sent back to be judged: %+v", s)
	}
}

// A red verdict is the one place a suggestion could push somebody into a decision, so it
// names the finding and both ways out and offers neither as a command.
func TestTheSuggestionNamesAFindingAndDecidesNothing(t *testing.T) {
	f := newFixture(t)
	g := f.redByRendering()

	s := f.r.Next(key)
	if !strings.Contains(s.Text, "is red on F-") {
		t.Fatalf("the finding was not named: %+v", s)
	}
	if s.Command != "" {
		t.Fatalf("a red verdict offered a command to run: %q", s.Command)
	}
	for _, want := range []string{"gate approve", "gate override", "second person"} {
		if !strings.Contains(s.Text, want) {
			t.Errorf("the way out through a decision does not mention %q: %q", want, s.Text)
		}
	}
	if !strings.Contains(s.Text, firstFinding(g).Next) {
		t.Errorf("the finding's own repair is not passed on: %q", s.Text)
	}
}

// An override lets the work go on, so the obligation it leaves is not the next step. It
// is listed, because section 6 puts it at "later, whoever owes it", and later is
// otherwise never.
func TestAnOverrideIsOwedAndDoesNotBlockTheSequence(t *testing.T) {
	f := newFixture(t)
	// Every finding, because one decision on a verdict with several leaves it red, and
	// this test is about what an override does once the phase is decided.
	var ids []string
	for _, c := range f.redByRendering().Checks {
		for _, fn := range c.Findings {
			ids = append(ids, fn.ID)
		}
	}
	for _, id := range ids {
		f.must2(f.r.Decide(key, "00-intake", id, "overridden", "a second person", "shipping"))
	}

	s := f.r.Next(key)
	if !strings.Contains(s.Command, "xeno phase start") {
		t.Fatalf("an overridden phase did not let the sequence go on: %+v", s)
	}
	if len(s.Owed) != len(ids) {
		t.Fatalf("%d obligations open, %d listed: %+v", len(ids), len(s.Owed), s.Owed)
	}
	if !strings.Contains(s.Owed[0], "xeno obligation close "+ids[0]) {
		t.Fatalf("the obligation is not runnable as printed: %q", s.Owed[0])
	}

	for _, id := range ids {
		f.must2(f.r.CloseObligation(key, "00-intake", id))
	}
	if s := f.r.Next(key); len(s.Owed) != 0 {
		t.Fatalf("a closed obligation is still owed: %+v", s.Owed)
	}
}

func firstFinding(g *model.Gate) *model.Finding {
	for i := range g.Checks {
		if len(g.Checks[i].Findings) > 0 {
			return &g.Checks[i].Findings[0]
		}
	}
	return nil
}

// The pipeline owes the evidence, and nothing here can produce it. A suggestion that
// offered a command would be offering one that cannot help.
func TestAProvisionalVerdictSuggestsNoCommand(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	f.must(f.r.Start(key, "04-verification"))
	f.output("04-verification", pendingTest)
	f.finish("04-verification")

	s := f.r.Next(key)
	if s.Command != "" {
		t.Fatalf("a provisional verdict offered %q", s.Command)
	}
	if !strings.Contains(s.Text, "provisional") || !strings.Contains(s.Text, "pipeline") {
		t.Fatalf("it does not say who owes what: %q", s.Text)
	}
}

// The merge is a step of the sequence and not a subcommand, and the last phase has to
// say so rather than pointing at a phase that does not exist.
func TestAfterP5TheNextStepIsNotACommand(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases {
		f.run(p, "")
	}
	s := f.r.Next(key)
	if s.Command != "" {
		t.Fatalf("something was offered after P5: %q", s.Command)
	}
	if !strings.Contains(s.Text, "merge") {
		t.Fatalf("the merge is not named: %q", s.Text)
	}
}

func TestWithoutAnIntentTheSuggestionSaysWhatIsMissing(t *testing.T) {
	f := newFixture(t)
	s := f.r.Next("PROJ-404")
	if s.Command != "" || !strings.Contains(s.Text, "no intent PROJ-404") {
		t.Fatalf("an intent that does not exist got %+v", s)
	}
}

// A state the working sequence does not cover says so. A confident wrong suggestion is
// worse than none, because it is followed.
func TestAnUncoveredStateIsSaidToBeOne(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	g, err := f.r.readGate(key, "00-intake")
	f.must(err)
	g.Status = "sideways"
	f.must(fm.WriteYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "gate.yaml"), g))

	s := f.r.Next(key)
	if s.Command != "" || !strings.Contains(s.Text, "does not cover") {
		t.Fatalf("an unknown status got %+v", s)
	}
}

// ---- The assumption register, section 8

const register = "assumptions:\n  - id: A-001\n    assumption: the cache is warm\n" +
	"    origin: repo-convention\n    confidence: low\n    status: "

// An open assumption turns the gate of its phase red, and a rejected one does not: it was
// examined and dropped, which an empty confirmed_by cannot say.
func TestAssumptionStatusDecidesTheGate(t *testing.T) {
	for _, tc := range []struct {
		status, want string
	}{
		{"open", "red"},
		{"", "red"},
		{"confirmed\n    confirmed_by: m.example", "green"},
		{"rejected\n    rejected_by: m.example", "green"},
	} {
		f := newFixture(t)
		f.write(model.IntentDir(key)+"/assumptions.yaml", register+tc.status+"\n")
		if g := f.run("00-intake", ""); g.Status != tc.want {
			t.Errorf("status %q gave %s, wanted %s", tc.status, g.Status, tc.want)
		}
	}
}

// Section 8 gives each decided state its person, so a decision with nobody behind it is a
// finding of its own rather than a pass, in both states.
func TestDecidedByNobodyIsAFinding(t *testing.T) {
	for _, status := range []string{"confirmed", "rejected"} {
		f := newFixture(t)
		f.write(model.IntentDir(key)+"/assumptions.yaml", register+status+"\n")
		if g := f.run("00-intake", ""); g.Status != "red" {
			t.Errorf("%s with no person behind it passed: %s", status, g.Status)
		}
	}
}

// The person goes into the field that belongs to the status, and the other stays absent:
// a rejected assumption was not confirmed by anybody.
func TestRejectionNamesItsPersonInItsOwnField(t *testing.T) {
	f := newFixture(t)
	f.must2nd(f.r.RecordAssumption(key, "02-design", "the cache is warm", "rules", "high", ""))
	a, err := f.r.DecideAssumption(key, "A-001", "rejected", "m.example")
	f.must(err)
	if a.RejectedBy != "m.example" || a.ConfirmedBy != "" {
		t.Fatalf("a rejection recorded %+v", a)
	}
	if a.DecidedBy() != "m.example" {
		t.Fatalf("the person behind the rejection reads as %q", a.DecidedBy())
	}
}

func TestRecordAssumption(t *testing.T) {
	f := newFixture(t)
	a, err := f.r.RecordAssumption(key, "02-design", "the cache is warm", "rules", "high", "Q-1")
	f.must(err)
	if a.ID != "A-001" || a.Status != "open" {
		t.Fatalf("recorded %+v", a)
	}
	b, err := f.r.RecordAssumption(key, "02-design", "and stays warm", "rules", "high", "")
	f.must(err)
	if b.ID != "A-002" {
		t.Fatalf("the second id was %s", b.ID)
	}
	var reg model.Assumptions
	f.must(fm.ReadYAML(filepath.Join(f.root, model.IntentDir(key), "assumptions.yaml"), &reg))
	if reg.Updated == "" {
		t.Error("the register was not stamped as updated")
	}
	if len(reg.Assumptions) != 2 || reg.Assumptions[0].Origin != "rules" {
		t.Fatalf("register is %+v", reg.Assumptions)
	}
}

// Origin and confidence are closed sets, and the command is where that is enforced: a
// register full of spellings the reader has to guess at is worse than an empty one.
func TestRecordRefusesValuesOutsideTheSets(t *testing.T) {
	f := newFixture(t)
	for _, tc := range [][2]string{{"", "high"}, {"invented", "high"}, {"rules", ""}, {"rules", "certain"}} {
		if _, err := f.r.RecordAssumption(key, "02-design", "text", tc[0], tc[1], ""); err == nil {
			t.Errorf("origin %q confidence %q was accepted", tc[0], tc[1])
		}
	}
	if _, err := f.r.RecordAssumption(key, "02-design", "", "rules", "high", ""); err == nil {
		t.Error("an assumption with no statement was accepted")
	}
}

func TestDecideAssumption(t *testing.T) {
	f := newFixture(t)
	f.must2nd(f.r.RecordAssumption(key, "02-design", "the cache is warm", "rules", "high", ""))
	a, err := f.r.DecideAssumption(key, "A-001", "rejected", "m.example")
	f.must(err)
	if a.Status != "rejected" || a.ConfirmedBy != "" {
		t.Fatalf("a rejection recorded %+v; confirmed_by belongs to a confirmation", a)
	}
	if a.RejectedBy != "m.example" {
		t.Fatalf("the rejection named %q", a.RejectedBy)
	}
	if _, err := f.r.DecideAssumption(key, "A-001", "confirmed", "m.example"); err == nil {
		t.Error("a decided assumption was decided again")
	}
	if _, err := f.r.DecideAssumption(key, "A-002", "confirmed", "m.example"); err == nil {
		t.Error("an assumption that does not exist was confirmed")
	}
	if _, err := f.r.DecideAssumption(key, "A-001", "confirmed", ""); err == nil {
		t.Error("a decision with nobody behind it was accepted")
	}
}

func (f *fixture) must2nd(_ *model.Assumption, err error) { f.t.Helper(); f.must(err) }

// Appendix B fixes what a hash field may carry, and G-Schema turns that into a verdict.
// The cases are exhaustive in internal/gates; what this asserts is the colour a phase comes
// out as, which is what a developer sees.
func TestHashFieldShape(t *testing.T) {
	const ofTheLock = "<the lock's hash>"
	for _, tc := range []struct {
		name, tool, context, secrets, want string
	}{
		{"the lock's own hash passes", "claude-code", ofTheLock, hex64('b'), "green"},
		{"by-hand passes where no writer exists", "claude-code", ofTheLock, "by-hand", "green"},
		{"by-hand passes a manual artifact", "manual", "by-hand", "by-hand", "green"},
		{"by-hand fails where a session wrote", "claude-code", "by-hand", hex64('b'), "red"},
		{"a hash of the wrong content fails", "claude-code", hex64('a'), hex64('b'), "red"},
		{"a truncated hash fails", "claude-code", "abc123", hex64('b'), "red"},
		{"an invented value fails", "claude-code", "todo", hex64('b'), "red"},
	} {
		f := newFixture(t)
		d := model.PhaseDir(key, "00-intake")
		// Started first, because the lock file is what the hash covers and phase start is
		// what writes it.
		f.must(f.r.Start(key, "00-intake"))
		context := tc.context
		if context == ofTheLock {
			context = f.lockHash("00-intake")
		}
		common := "intent: git.example/group/proj#1\nphase: 00-intake\ncreated: 2026-09-20T10:00:00Z\n" +
			"runner_version: 0.1.0-dev\nplugin_version: 0.1.0-dev\n"
		session := "language: en\nsecrets_hash: " + tc.secrets + "\ncontext_hash: " + context +
			"\nmodel: m\ntool: " + tc.tool + "\ntool_version: 1\n"
		f.write(d+"/output.md", "---\n"+common+session+"template: t@1\nstrings_hash: "+hex64('3')+
			"\nrules_hash: by-hand\n---\n\n# Result\n")
		f.write(d+"/digest.md", "---\n"+common+session+"---\nsummary\n")
		f.write(d+"/learning.yaml", common+"no_finding: true\n")
		g, err := f.r.Finish(key, "00-intake", "")
		f.must(err)
		if g.Status != tc.want {
			t.Errorf("%s: got %s, wanted %s", tc.name, g.Status, tc.want)
		}
	}
}

// A verdict's gate set is compared against the table rather than against another run,
// because a recomputation uses the same table and leaves out the same gate. Without that
// comparison an absence says both "does not apply yet" and "did not run".
func TestVerifyComparesTheGateSet(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	rel := model.PhaseDir(key, "00-intake") + "/gate.yaml"

	var g model.Gate
	f.must(fm.ReadYAML(filepath.Join(f.root, rel), &g))
	full := len(g.Checks)
	if full != len(gates.Applicable("00-intake")) {
		t.Fatalf("a fresh verdict carries %d checks and %d apply", full, len(gates.Applicable("00-intake")))
	}

	// A gate that silently did not run leaves a verdict one check short, and every hash
	// still matches, because gate.yaml is not covered by artifacts_hash.
	dropped := g
	dropped.Checks = nil
	for _, ch := range g.Checks {
		if ch.Gate != "G-Trace" {
			dropped.Checks = append(dropped.Checks, ch)
		}
	}
	f.must(fm.WriteYAML(filepath.Join(f.root, rel), dropped))
	res, err := f.r.Verify(key)
	f.must(err)
	if !divergenceNames(res, "G-Trace applies") {
		t.Fatalf("a missing gate was not reported: %+v", res.Divergences)
	}

	// The same defect the other way: a verdict claiming a gate whose phase has not come.
	early := g
	early.Checks = append(append([]model.Check{}, g.Checks...), model.Check{Gate: "G-Questions", Result: "pass"})
	f.must(fm.WriteYAML(filepath.Join(f.root, rel), early))
	res, err = f.r.Verify(key)
	f.must(err)
	if !divergenceNames(res, "G-Questions is in the verdict") {
		t.Fatalf("a gate before its phase was not reported: %+v", res.Divergences)
	}
}

func divergenceNames(res *VerifyResult, want string) bool {
	for _, d := range res.Divergences {
		if strings.Contains(d.What, want) {
			return true
		}
	}
	return false
}

// WP8's half of G-Freshness: a phase records the information base it was given, with a
// hash each, and a file that changes afterwards is a finding. A6 accepted the gap as
// temporary and said a pass here said more than was checked.
func TestGivenFilesAreComparedAgainstTheTree(t *testing.T) {
	f := newFixture(t)
	f.write("src/payment/card.go", "package payment\n")
	f.write("src/payment/testdata/golden.json", "{}\n")
	f.write("docs/adr/0012-payments.md", "# payments\n")
	f.write("src/shipping/box.go", "package shipping\n")
	f.write(model.PhaseDir(key, "00-intake")+"/"+model.ContextScope,
		"include:\n  - src/payment/**\n  - docs/adr/*.md\nexclude:\n  - \"**/testdata/**\"\nbudget:\n  files: 120\n")

	f.must(f.r.Start(key, "00-intake"))
	var lock model.ContextLock
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "context.lock.yaml"), &lock))

	var paths []string
	for _, c := range lock.Files {
		paths = append(paths, c.Path)
	}
	// The order is the scope's include order, not the alphabet: section 5 asks for an order of
	// volatility and says the lock records the assembly order rather than only the set, so the
	// project writes its patterns from stable to volatile and this follows (#171).
	want := []string{"src/payment/card.go", "docs/adr/0012-payments.md"}
	if strings.Join(paths, " ") != strings.Join(want, " ") {
		t.Fatalf("the information base is %v, want %v: the scope's include order, its exclude, and the files it does not name", paths, want)
	}

	f.output("00-intake", "")
	if g := f.finish("00-intake"); g.Status != "green" {
		t.Fatalf("a phase whose files are untouched is %s", g.Status)
	}

	// The file the phase was given changes after it was given, and P0 stays green: section 5
	// says the check "looks at the locks of the preceding phases, never at the lock of the
	// phase being gated", and P0 has none. This test asserted the opposite until #236, which
	// is how the defect lived: "a phase that changes the files it read is not stale, it is
	// working". What the check does catch now has its own tests beside staleReads, where the
	// two limits are, because both need a commit range and this fixture has no repository.
	f.write("src/payment/card.go", "package payment // and more\n")
	g, err := f.r.GateRun(key, "00-intake")
	f.must(err)
	if g.Status != "green" {
		t.Fatalf("a phase that changed a file its own lock lists is %s, and it is working rather than stale", g.Status)
	}
}

// A phase with no scope records no information base, which is a smaller claim than an
// empty one: nothing was declared rather than nothing read.
func TestNoProfileRecordsNoInformationBase(t *testing.T) {
	f := newFixture(t)
	f.must(f.r.Start(key, "00-intake"))
	var lock model.ContextLock
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "context.lock.yaml"), &lock))
	if len(lock.Files) != 0 {
		t.Fatalf("a repository without a scope recorded %v", lock.Files)
	}
	f.output("00-intake", "")
	if g := f.finish("00-intake"); g.Status != "green" {
		t.Fatalf("a phase without a scope is %s", g.Status)
	}
}

func causeOf(g *model.Gate, gate string) string {
	var b strings.Builder
	for _, ch := range g.Checks {
		if ch.Gate != gate {
			continue
		}
		for _, f := range ch.Findings {
			b.WriteString(f.Cause)
			b.WriteString("; ")
		}
	}
	return b.String()
}

// Section 7's first mode of G-Complete: a merging intent reaches review with every
// preceding phase present and green, approved or overridden. A32 recorded that only the
// abandoned mode was implemented, and the phase table said so in every verdict.
func TestCompleteInReviewReadsEveryPrecedingPhase(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:5] {
		f.run(p, "")
	}
	g := f.run("05-review", "")
	if g.Status != "green" {
		t.Fatalf("an intent with five judged phases is %s: %s", g.Status, causeOf(g, "G-Complete"))
	}
	if check(g, "G-Complete").Result != "pass" {
		t.Fatalf("G-Complete is %s, want pass", check(g, "G-Complete").Result)
	}

	// A phase whose verdict is gone was never judged, which is a different repair from a
	// phase that was judged and failed.
	f2 := newFixture(t)
	for _, p := range model.Phases[:5] {
		f2.run(p, "")
	}
	f2.must(os.Remove(filepath.Join(f2.root, model.PhaseDir(key, "02-design"), "gate.yaml")))
	g = f2.run("05-review", "")
	if g.Status != "red" || !strings.Contains(causeOf(g, "G-Complete"), "02-design holds no verdict") {
		t.Fatalf("a missing verdict gave %s: %s", g.Status, causeOf(g, "G-Complete"))
	}

	// A red predecessor cannot be reached through phase start, which refuses it, so the
	// gate is what catches a verdict edited after the fact.
	f3 := newFixture(t)
	for _, p := range model.Phases[:5] {
		f3.run(p, "")
	}
	rel := filepath.Join(f3.root, model.PhaseDir(key, "03-implementation"), "gate.yaml")
	var edited model.Gate
	f3.must(fm.ReadYAML(rel, &edited))
	edited.Status = "red"
	f3.must(fm.WriteYAML(rel, &edited))
	g = f3.run("05-review", "")
	if !strings.Contains(causeOf(g, "G-Complete"), "03-implementation is red") {
		t.Fatalf("a red predecessor gave: %s", causeOf(g, "G-Complete"))
	}
}

// The plan asks phase start to export the intent and the phase into the environment a
// harness reads for request headers. A child cannot set its parent's environment, so it
// writes them where a later process reads them, and a phase that has ended attributes
// nothing.
func TestPhaseEnvIsWrittenAndRemoved(t *testing.T) {
	f := newFixture(t)
	env := filepath.Join(f.root, ".xeno/local/phase.env")

	f.must(f.r.Start(key, "00-intake"))
	b, err := os.ReadFile(env)
	f.must(err)
	got := string(b)
	for _, want := range []string{`export XENO_INTENT="git.example/group/proj#1"`, `export XENO_PHASE="00-intake"`} {
		if !strings.Contains(got, want) {
			t.Fatalf("the phase environment is %q, want a line %q", got, want)
		}
	}
	// The qualified id and not the directory name: a request tagged with a guess is worse
	// than one tagged with nothing.
	if strings.Contains(got, key) {
		t.Fatalf("the environment carries the directory name: %q", got)
	}

	f.output("00-intake", "")
	f.finish("00-intake")
	if _, err := os.Stat(env); !os.IsNotExist(err) {
		t.Fatal("the phase environment outlived the phase")
	}

	// The same content is what --export prints, so the file and the print cannot drift.
	f.must(f.r.Start(key, "01-requirements"))
	printed, err := f.r.PhaseEnv(key, "01-requirements")
	f.must(err)
	b, err = os.ReadFile(env)
	f.must(err)
	if printed != string(b) {
		t.Fatalf("print is %q and the file is %q", printed, string(b))
	}
}

// brokenPipeline publishes the unit test result with a uri and no hash, which is a
// pipeline that ran and published wrongly rather than one that has not answered.
func (f *fixture) brokenPipeline() string {
	dir := filepath.Join(f.t.TempDir(), "artifacts")
	_ = os.MkdirAll(dir, 0o755)
	os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(
		"- kind: test-report\n  job: unit\n  result: pass\n  uri: https://ci.example/a/7\n"), 0o644)
	return dir
}

// The refusal has to say which entry and why. An item the pipeline published wrong keeps
// the phase provisional exactly as a missing one does, and "start again once it has run" is
// the wrong advice for it: the job has run, and nothing arrives by waiting.
func TestStartNamesAnEntryThePipelinePublishedWrong(t *testing.T) {
	f := newFixture(t)
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	f.must(f.r.Start(key, "04-verification"))
	f.output("04-verification", pendingTest)
	f.finish("04-verification")
	sealed := f.hash("04-verification")

	f.r.EvidenceFrom = f.brokenPipeline()
	var ref *Refusal
	err := f.r.Start(key, "05-review")
	if !errors.As(err, &ref) {
		t.Fatalf("P5 started on an attachment nothing binds: %v", err)
	}
	for _, want := range []string{"test-report/unit", "sha256", "Waiting will not help"} {
		if !strings.Contains(ref.Reason, want) {
			t.Errorf("the refusal does not mention %q:\n%s", want, ref.Reason)
		}
	}
	// Nothing was recorded, so the phase is exactly where it was and a corrected pipeline
	// still attaches cleanly.
	if f.hash("04-verification") != sealed {
		t.Fatal("declining an entry changed the artifacts_hash of the sealed phase")
	}
	f.r.EvidenceFrom = f.pipeline("pass")
	f.must(f.r.Start(key, "05-review"))
	g, err := f.r.readGate(key, "04-verification")
	f.must(err)
	if g.Status != "green" {
		t.Fatalf("the corrected pipeline did not carry the verdict forward: %s", g.Status)
	}
}

// ---- a check that fails without saying what failed, through a caller

// gate.yaml lies outside artifacts_hash by design, so it is the one file of a judged phase
// that can be edited without staling a hash. A malformed check written there reaches Status
// through rewriteStatus, which is why the rule is not in Invariants: that runs on checks a
// gate just produced and never on these.
func TestAMalformedCheckInAStoredVerdictRefusesTheDecision(t *testing.T) {
	f := newFixture(t)
	g := f.run("00-intake", "open_questions:\n  - key: Q-9\n    text: bare\n")
	if g.Status != "red" {
		t.Fatalf("the fixture phase is %s, want red so that there is a finding to decide", g.Status)
	}
	var id string
	for _, c := range g.Checks {
		if len(c.Findings) > 0 {
			id = c.Findings[0].ID
			break
		}
	}
	if id == "" {
		t.Fatal("a red phase carried no finding to decide")
	}

	// Written as a person with an editor would, into the one file no hash covers.
	g.Checks = append(g.Checks, model.Check{Gate: "G-External", Result: "fail", Provenance: "external"})
	f.must(fm.WriteYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "gate.yaml"), g))

	_, err := f.r.Decide(key, "00-intake", id, "approved", "a.person", "assessed")
	if err == nil {
		t.Fatal("a decision on a verdict carrying a fail with no finding was written")
	}
	if !strings.Contains(err.Error(), "G-External") {
		t.Errorf("the refusal does not name the gate that failed without saying what: %v", err)
	}
}

// ---- the listing, in the order the work happened

// intentAt writes a bare intent directory with a created value, which is all the listing reads.
func (f *fixture) intentAt(key, created string) {
	f.t.Helper()
	body := "intent: \"git.example/group/proj#1\"\nkey: " + key + "\nstatus: in-progress\n"
	if created != "" {
		body += "created: \"" + created + "\"\n"
	}
	f.write(model.IntentDir(key)+"/intent.yaml", body)
}

func keysOf(t *testing.T, got []IntentSummary) []string {
	t.Helper()
	var out []string
	for _, s := range got {
		out = append(out, s.Key)
	}
	return out
}

// The case #118 is about: keys that sort one way and work that happened in another. The key
// carried the issue number, and issues are filed in a different order than work is done.
func TestIntentsAreListedInTheOrderTheyWereCreated(t *testing.T) {
	f := newFixture(t)
	f.intentAt("XENO-0108", "2026-09-28T09:00:00Z")
	f.intentAt("XENO-0111", "2026-09-28T18:00:00Z")
	f.intentAt("XENO-0107", "2026-09-28T19:00:00Z")
	f.intentAt("XENO-0121", "2026-09-28T20:00:00Z")

	got, err := f.r.Intents()
	f.must(err)
	want := []string{"PROJ-1", "XENO-0108", "XENO-0111", "XENO-0107", "XENO-0121"}
	if strings.Join(keysOf(t, got), ",") != strings.Join(want, ",") {
		t.Fatalf("listed %v, want %v", keysOf(t, got), want)
	}
}

// Two intents of one day are ordered by their time. Truncating the value to a date before
// sorting would fall through to the key tie break and reproduce the order being corrected.
func TestTwoIntentsOfOneDayAreOrderedByTheirTime(t *testing.T) {
	f := newFixture(t)
	f.intentAt("XENO-0300", "2026-10-01T18:00:00Z")
	f.intentAt("XENO-0201", "2026-10-01T19:00:00Z")

	got, err := f.r.Intents()
	f.must(err)
	if got[len(got)-1].Key != "XENO-0201" {
		t.Fatalf("the later intent of the day is %s, want XENO-0201: %v", got[len(got)-1].Key, keysOf(t, got))
	}
}

// A record that cannot be dated is listed with the reason, because one missing from a listing
// is worse than one that looks wrong in it. The empty value sorts it to the front, where
// somebody sees it.
func TestAnIntentWithoutACreatedIsListedWithTheReason(t *testing.T) {
	f := newFixture(t)
	f.intentAt("XENO-0202", "2026-10-02T10:00:00Z")
	f.intentAt("XENO-0203", "")

	got, err := f.r.Intents()
	f.must(err)
	var undated *IntentSummary
	for i := range got {
		if got[i].Key == "XENO-0203" {
			undated = &got[i]
		}
	}
	if undated == nil {
		t.Fatal("an intent with no created was left out of the listing")
	}
	if undated.Problem == "" {
		t.Error("the row gives no reason for its missing date")
	}
	// Before the dated one, not necessarily first: the fixture's own intent carries no
	// created either, so the undated ones sort together at the front.
	var undatedAt, datedAt int
	for i := range got {
		switch got[i].Key {
		case "XENO-0203":
			undatedAt = i
		case "XENO-0202":
			datedAt = i
		}
	}
	if undatedAt > datedAt {
		t.Errorf("the undated intent sorts after a dated one, so it is not where somebody looks: %v",
			keysOf(t, got))
	}
}

// Ties break on the key, so two runs over one tree print the same thing.
func TestTheOrderIsStableForIdenticalTimes(t *testing.T) {
	f := newFixture(t)
	f.intentAt("XENO-0205", "2026-10-03T10:00:00Z")
	f.intentAt("XENO-0204", "2026-10-03T10:00:00Z")

	first, err := f.r.Intents()
	f.must(err)
	second, err := f.r.Intents()
	f.must(err)
	if strings.Join(keysOf(t, first), ",") != strings.Join(keysOf(t, second), ",") {
		t.Fatalf("two runs disagree: %v then %v", keysOf(t, first), keysOf(t, second))
	}
	if first[len(first)-2].Key != "XENO-0204" || first[len(first)-1].Key != "XENO-0205" {
		t.Errorf("identical times did not break on the key: %v", keysOf(t, first))
	}
}

// A repository that has started no intent is not in error.
func TestNoIntentDirectoryListsNothingAndDoesNotFail(t *testing.T) {
	r := New(t.TempDir())
	got, err := r.Intents()
	if err != nil {
		t.Fatalf("an empty repository reported an error: %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("listed %d intents in an empty repository", len(got))
	}
}

// The row carries where an intent got to, which is what saves opening it.
func TestTheRowNamesTheFurthestPhaseWithAVerdict(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")

	got, err := f.r.Intents()
	f.must(err)
	var own *IntentSummary
	for i := range got {
		if got[i].Key == key {
			own = &got[i]
		}
	}
	if own == nil {
		t.Fatal("the fixture intent is not in the listing")
	}
	if own.Phase != "00-intake" || own.Verdict != "green" {
		t.Fatalf("the row says %q %q, want 00-intake green", own.Phase, own.Verdict)
	}
}

// ---- the runner writes the digest (#120)

// project writes the agent block section 12 defines, which is where model and tool live.
func (f *fixture) project(body string) {
	f.t.Helper()
	f.write(".xeno/config/project.yaml", body)
}

const agentBlock = "language:\n  artifacts: en\nagent:\n  tool: claude-code\n  model:\n    default: a-model\n"

func digestFront(t *testing.T, root, phase string) (map[string]any, string) {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(root, model.PhaseDir(key, phase), "digest.md"))
	if err != nil {
		t.Fatal(err)
	}
	front, body, err := fm.Split(b)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := yaml.Unmarshal(front, &m); err != nil {
		t.Fatal(err)
	}
	return m, string(body)
}

// Section 5: the agent supplies the summary text, the runner writes the file. What it carries is
// where it came from, and three groups rather than two: no template and no strings bundle.
func TestFinishWritesTheDigestFromTheSummary(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	_, err := f.r.SectionSet(key, "00-intake", "problem", "what is wrong")
	f.must(err)

	f.must2(f.r.Finish(key, "00-intake", "the summary of the exchange"))

	front, body := digestFront(t, f.root, "00-intake")
	if !strings.Contains(body, "the summary of the exchange") {
		t.Errorf("the digest does not carry the summary: %q", body)
	}
	for _, want := range []string{"intent", "phase", "created", "schema_version",
		"runner_version", "plugin_version", "language", "context_hash", "model", "tool"} {
		if front[want] == nil || front[want] == "" {
			t.Errorf("the digest carries no %s", want)
		}
	}
	// A digest is neither rendered nor covered by a rule set.
	for _, unwanted := range []string{"template", "strings_hash", "rules_hash"} {
		if _, ok := front[unwanted]; ok {
			t.Errorf("the digest carries %s, which belongs to a rendered file", unwanted)
		}
	}
	if front["model"] != "a-model" || front["tool"] != "claude-code" {
		t.Errorf("model and tool are %v and %v, want the project's", front["model"], front["tool"])
	}
}

// The honest half of the writer. Section 5 gives the runner filtering as well as writing, and
// there is no filter in this tree, so the field is absent rather than a hash over nothing. The
// consequence is that the phase is red on it, which is the true state and not a regression.
func TestTheWrittenDigestCarriesNoSecretsHashAndTheGateSaysSo(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	_, err := f.r.SectionSet(key, "00-intake", "problem", "what is wrong")
	f.must(err)

	g, err := f.r.Finish(key, "00-intake", "a summary")
	f.must(err)

	front, _ := digestFront(t, f.root, "00-intake")
	if _, ok := front["secrets_hash"]; ok {
		t.Error("the digest asserts a filter that does not exist")
	}
	var named bool
	for _, c := range g.Checks {
		for _, fd := range c.Findings {
			if strings.Contains(fd.File, "digest.md") && strings.Contains(fd.Cause, "secrets_hash") {
				named = true
			}
		}
	}
	if !named {
		t.Error("G-Schema does not report the missing secrets_hash of the digest it was given")
	}
}

// Section 5 says the supported path is not an enforced one, so a phase finished without a summary
// is judged exactly as before and no digest appears.
func TestFinishWithoutASummaryWritesNoDigest(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	_, err := f.r.SectionSet(key, "00-intake", "problem", "what is wrong")
	f.must(err)

	f.must2(f.r.Finish(key, "00-intake", ""))

	if _, err := os.Stat(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "digest.md")); err == nil {
		t.Fatal("a finish with no summary wrote a digest")
	}
}

// One source, two writers: the fields appear in output.md as the digest carries them.
func TestSectionSetWritesTheModelAndToolTheProjectRecords(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	_, err := f.r.SectionSet(key, "00-intake", "problem", "what is wrong")
	f.must(err)

	if got := f.frontField("00-intake", "model"); got != "a-model" {
		t.Errorf("model is %q, want the project's", got)
	}
	if got := f.frontField("00-intake", "tool"); got != "claude-code" {
		t.Errorf("tool is %q, want the project's", got)
	}
}

// A35's argument, kept: a plausible value in a field nobody produced is worse than an absent one,
// and the only tool this repository has used would be exactly such a value.
func TestNoAgentBlockLeavesBothFieldsAbsent(t *testing.T) {
	for _, tc := range []struct{ name, body string }{
		{"no project file", ""},
		{"no agent block", "language:\n  artifacts: en\n"},
		{"an empty agent block", "language:\n  artifacts: en\nagent: {}\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newFixture(t)
			if tc.body != "" {
				f.project(tc.body)
			}
			f.templated()
			f.must(f.r.Start(key, "00-intake"))
			_, err := f.r.SectionSet(key, "00-intake", "problem", "what is wrong")
			f.must(err)
			f.must2(f.r.Finish(key, "00-intake", "a summary"))

			if got := f.frontField("00-intake", "model"); got != "" {
				t.Errorf("output.md guessed a model: %q", got)
			}
			front, _ := digestFront(t, f.root, "00-intake")
			for _, k := range []string{"model", "tool"} {
				if _, ok := front[k]; ok {
					t.Errorf("the digest guessed %s: %v", k, front[k])
				}
			}
		})
	}
}

// The order section 5 lists, one order for both artifacts, so a reader comparing them does not
// have to know which writer produced which.
func TestTheDigestCarriesTheFieldOrderOfSectionFive(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	_, err := f.r.SectionSet(key, "00-intake", "problem", "what is wrong")
	f.must(err)
	f.must2(f.r.Finish(key, "00-intake", "a summary"))

	b, err := os.ReadFile(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "digest.md"))
	f.must(err)
	front, _, err := fm.Split(b)
	f.must(err)
	at := func(field string) int { return strings.Index(string(front), field+":") }
	for _, pair := range [][2]string{{"intent", "phase"}, {"phase", "created"},
		{"created", "schema_version"}, {"language", "context_hash"}, {"context_hash", "model"},
		{"model", "tool"}} {
		if at(pair[0]) > at(pair[1]) {
			t.Errorf("%s comes after %s:\n%s", pair[0], pair[1], front)
		}
	}
}

// ---- the digest passes through a filter (#120)

const shippedFilter = `patterns:
  - id: aws-access-key
    regex: '\b(?:AKIA|ASIA)[0-9A-Z]{16}\b'
paths_never_digested:
  - "**/*.pem"
`

// Section 16: the runner holds the text between the summary and the file, so the filtering is
// deterministic and outside the model's reach. This is the test of that sentence.
func TestTheDigestIsFilteredAndSaysWhichFilter(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.write(secrets.Shipped, shippedFilter)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	_, err := f.r.SectionSet(key, "00-intake", "problem", "what is wrong")
	f.must(err)

	f.must2(f.r.Finish(key, "00-intake", "the run used AKIAIOSFODNN7EXAMPLE against the bucket"))

	front, body := digestFront(t, f.root, "00-intake")
	if strings.Contains(body, "AKIAIOSFODNN7EXAMPLE") {
		t.Fatalf("the secret reached the digest: %q", body)
	}
	if !strings.Contains(body, "[redacted: aws-access-key]") {
		t.Errorf("the redaction does not name the pattern: %q", body)
	}
	if !strings.Contains(body, "against the bucket") {
		t.Errorf("the line around the match was destroyed: %q", body)
	}
	want, err := secrets.Load(f.root)
	f.must(err)
	if front["secrets_hash"] != want.Hash() {
		t.Errorf("secrets_hash is %v, want the hash of the effective filter %s", front["secrets_hash"], want.Hash())
	}
}

// output.md carries the field and is not itself redacted: section 16 puts the filtering on the
// digest, and a filter over the agent's prose would redact a discussion of a pattern by that
// pattern. This repository's own records are the case that proves it.
func TestSectionSetWritesTheHashAndDoesNotRedactTheProse(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.write(secrets.Shipped, shippedFilter)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	const prose = "A key looks like AKIAIOSFODNN7EXAMPLE, which is what the pattern catches."
	_, err := f.r.SectionSet(key, "00-intake", "problem", prose)
	f.must(err)

	want, err := secrets.Load(f.root)
	f.must(err)
	if got := f.frontField("00-intake", "secrets_hash"); got != want.Hash() {
		t.Errorf("output.md carries secrets_hash %q, want %q", got, want.Hash())
	}
	b, err := os.ReadFile(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "output.md"))
	f.must(err)
	if !strings.Contains(string(b), prose) {
		t.Error("the agent's prose was redacted, which section 16 does not ask for")
	}
}

// A repository before its plugin is vendored. No filter, no field, nothing redacted, and the
// phase still finishes: absence is the value for an empty set, and by-hand is reserved for a
// person.
func TestWithoutAFilterNothingIsRedactedAndNoHashIsWritten(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	_, err := f.r.SectionSet(key, "00-intake", "problem", "what is wrong")
	f.must(err)

	const summary = "the run used AKIAIOSFODNN7EXAMPLE against the bucket"
	f.must2(f.r.Finish(key, "00-intake", summary))

	front, body := digestFront(t, f.root, "00-intake")
	if !strings.Contains(body, "AKIAIOSFODNN7EXAMPLE") {
		t.Error("something redacted with no filter in the repository")
	}
	if _, ok := front["secrets_hash"]; ok {
		t.Errorf("a digest that passed through no filter carries secrets_hash: %v", front["secrets_hash"])
	}
	if got := f.frontField("00-intake", "secrets_hash"); got != "" {
		t.Errorf("output.md carries secrets_hash %q with no filter", got)
	}
}

// ---- the listing computes the state (#132)

// abandonedIntent writes an intent whose own status carries the one value that field can carry.
func (f *fixture) abandonedIntent(key, reason string) {
	f.t.Helper()
	f.write(model.IntentDir(key)+"/intent.yaml",
		"intent: \"git.example/group/proj#9\"\nkey: "+key+"\nstatus: abandoned\nreason: "+
			reason+"\ncreated: \"2026-09-20T10:00:00Z\"\n")
}

func stateOf(t *testing.T, r *Runner, key string) string {
	t.Helper()
	got, err := r.Intents()
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range got {
		if s.Key == key {
			return s.State
		}
	}
	t.Fatalf("%s is not in the listing", key)
	return ""
}

// The three states, and the reason the column is computed: intent.yaml's status can only say
// whether an intent was abandoned, so printing it read every finished intent as unfinished.
func TestTheListingComputesAbandonedCompleteAndInFlight(t *testing.T) {
	f := newFixture(t)
	f.templated()
	for _, p := range model.Phases {
		f.run(p, "")
	}
	if got := stateOf(t, f.r, key); got != "complete" {
		t.Errorf("an intent whose P5 is decided reads %q, want complete", got)
	}

	f.abandonedIntent("PROJ-9", "the requirement went away")
	if got := stateOf(t, f.r, "PROJ-9"); got != "abandoned" {
		t.Errorf("an abandoned intent reads %q", got)
	}

	// In flight: phases up to P2 only.
	g := newFixture(t)
	g.templated()
	g.run("00-intake", "")
	g.run("01-requirements", "")
	if got := stateOf(t, g.r, key); got != "01-requirements" {
		t.Errorf("an intent in flight reads %q, want the phase it reached", got)
	}
}

// A P5 that is red or provisional is in flight, not complete: it is the state the sequence refuses
// to build on, so the listing must not call it finished.
func TestAnUndecidedFinalPhaseIsNotComplete(t *testing.T) {
	f := newFixture(t)
	f.templated()
	for _, p := range model.Phases[:len(model.Phases)-1] {
		f.run(p, "")
	}
	// A question without options turns the phase red.
	g := f.run("05-review", "open_questions:\n  - key: Q-1\n    text: bare\n")
	if g.Status != "red" {
		t.Fatalf("the fixture P5 is %s, want red", g.Status)
	}
	if got := stateOf(t, f.r, key); got == "complete" {
		t.Error("a red P5 was reported complete, which the sequence would refuse to build on")
	}
}

// An intent directory with nothing in it says so rather than reading as complete beside a blank
// verdict.
func TestAnIntentWithNoPhasesSaysSo(t *testing.T) {
	f := newFixture(t)
	if got := stateOf(t, f.r, key); got != "no phases" {
		t.Errorf("an intent with no phases reads %q", got)
	}
}

// The predicate the listing uses is the one the sequence enforces, so the two cannot disagree about
// whether a phase is settled.
func TestDecidedIsTheSequencesOwnTest(t *testing.T) {
	for _, tc := range []struct {
		status string
		want   bool
	}{
		{"green", true}, {"approved", true}, {"overridden", true},
		{"red", false}, {"provisional", false},
	} {
		if got := Decided(tc.status); got != tc.want {
			t.Errorf("Decided(%q) is %v, want %v", tc.status, got, tc.want)
		}
	}
}

// Section 5's index block, read by the runner. An absent block means no index, which is the
// state of every repository that has not produced one, including this one.
func TestNoIndexBlockMeansNoIndex(t *testing.T) {
	f := newFixture(t)
	f.write(".xeno/config/project.yaml", "runner_version: "+model.RunnerVersion+"\n")
	i, why := f.r.symbolIndex(time.Now())
	if i != nil {
		t.Error("an index was returned for a project that configured none")
	}
	if !strings.Contains(why, "index.path") {
		t.Errorf("the reason is %q, want it to name the field", why)
	}
}

// The path is resolved against the repository root, so a project writes the relative path it
// would write anywhere else in this file.
func TestTheIndexPathIsRelativeToTheRepository(t *testing.T) {
	f := newFixture(t)
	f.write(".xeno/config/project.yaml", "runner_version: "+model.RunnerVersion+"\n"+
		"index:\n  path: .xeno/local/index/symbols.yaml\n  max_age_hours: 48\n")
	f.write(".xeno/local/index/symbols.yaml", "tool: go-symbols\ntool_version: 0.1.0\n"+
		"produced_at: \""+time.Now().UTC().Add(-30*time.Hour).Format(time.RFC3339)+"\"\n"+
		"symbols:\n  - name: Compare\n    kind: func\n    file: a.go\n    line: 1\n")

	i, why := f.r.symbolIndex(time.Now())
	if i == nil {
		t.Fatalf("the index was not found at a relative path: %s", why)
	}
	// Thirty hours old, inside the configured forty-eight, so max_age_hours was read rather
	// than defaulted to twenty-four.
	if got := i.Lookup("Compare"); len(got) != 1 {
		t.Errorf("Compare resolved to %v", got)
	}
}

// A phase written after internal/rules existed carries a hash over the rule set it was judged
// against, where every artifact before it carries the placeholder the harness wrote. The two
// meanings of the field are divided by that commit and recorded in A66.
func TestSectionSetWritesTheHashOfTheEffectiveRuleSet(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	_, err := f.r.SectionSet(key, "00-intake", "problem", "what is wrong")
	f.must(err)

	empty := f.frontField("00-intake", "rules_hash")
	if empty == model.HashPlaceholder {
		t.Fatal("rules_hash is the placeholder where a writer exists")
	}
	read, _ := rules.Load(f.root)
	eff, _ := rules.Effective(read)
	if empty != rules.Hash(eff) {
		t.Fatalf("rules_hash is %s, want the hash of the effective set %s", empty, rules.Hash(eff))
	}

	// A rule entering the tree changes the field, which is what makes it say which set was
	// in force rather than that one was.
	f.write(rules.ConfigDir+"/given/org/migration-note.yaml",
		"id: migration-note\nversion: 1\nscope: org\nkind: review\napplies_to: [00-intake]\n"+
			"statement: >\n  A change to a published interface comes with a migration note.\n")
	_, err = f.r.SectionSet(key, "00-intake", "problem", "what is wrong")
	f.must(err)
	if got := f.frontField("00-intake", "rules_hash"); got == empty {
		t.Error("a rule entering the tree left rules_hash alone")
	}
}

// Section 14 end to end, through the one place a verdict is assembled: a declared command's
// check reaches gate.yaml with its provenance, its finding carries an id derived like every
// other, and a decision taken on it does not survive the next run — which is the rule #66 wrote
// and nothing could exercise until there was an external gate (A25, section 4).
func TestAnExternalGateReachesTheVerdictAndKeepsNoDecision(t *testing.T) {
	f := newFixture(t)
	script := "#!/bin/sh\ncat > /dev/null\n" +
		`printf '{"findings":[{"file":"tools/house-linter","cause":"the house style is not followed","next":"follow it"}]}'` +
		"\nexit 1\n"
	f.write("tools/house-linter", script)
	if err := os.Chmod(filepath.Join(f.root, "tools/house-linter"), 0o755); err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(script))
	f.project(agentBlock + "external_gates:\n  - id: house-linter\n    path: tools/house-linter\n" +
		"    sha256: " + hex.EncodeToString(sum[:]) + "\n    phases: [00-intake]\n")
	f.templated()
	g := f.run("00-intake", "")
	var ext *model.Check
	for i := range g.Checks {
		if g.Checks[i].Gate == "house-linter" {
			ext = &g.Checks[i]
		}
	}
	if ext == nil {
		t.Fatal("the declared gate produced no check")
	}
	if ext.Provenance != "external" {
		t.Errorf("provenance %q, want external: the mark is the point of section 14", ext.Provenance)
	}
	if ext.Result != "fail" || len(ext.Findings) != 1 {
		t.Fatalf("check %+v, want one failing finding", ext)
	}
	id := ext.Findings[0].ID
	if id == "" {
		t.Fatal("the external finding carries no id")
	}
	if g.Status != "red" {
		t.Errorf("the phase is %s, want red: an external fail is a fail", g.Status)
	}

	// A release is taken on it like any other.
	f.must2(f.r.Decide(key, "00-intake", id, "approved", "somebody", "accepted for this release"))
	after, err := f.r.GateRun(key, "00-intake")
	f.must(err)
	for _, ch := range after.Checks {
		if ch.Gate != "house-linter" {
			continue
		}
		if len(ch.Findings) != 1 {
			t.Fatalf("check %+v, want the same finding", ch)
		}
		if ch.Findings[0].ID != id {
			t.Errorf("the finding's id moved from %s to %s", id, ch.Findings[0].ID)
		}
		if ch.Findings[0].Decision != nil {
			t.Error("the decision survived the next run; section 4 says a release on foreign wording is taken again")
		}
	}
}

// A project that declares none keeps the closed chain, which is the default and every project
// today: no check, and the verdict carries what it carried before.
func TestNoDeclarationLeavesTheVerdictAsItWas(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	for _, ch := range f.run("00-intake", "").Checks {
		if ch.Provenance == "external" {
			t.Fatalf("an external check appeared without a declaration: %+v", ch)
		}
	}
}

// Section 5: context is assembled in order of volatility and the lock records the assembly order
// rather than only the set. The project decides which of its directories is stable by the order it
// writes its patterns in, so reversing the scope reverses the base (#171).
func TestTheBaseFollowsTheScopesIncludeOrder(t *testing.T) {
	base := func(t *testing.T, include string) []string {
		t.Helper()
		f := newFixture(t)
		f.write("src/a.go", "package a\n")
		f.write("docs/b.md", "# b\n")
		f.write(model.PhaseDir(key, "00-intake")+"/"+model.ContextScope, include)
		f.must(f.r.Start(key, "00-intake"))
		var lock model.ContextLock
		f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "context.lock.yaml"), &lock))
		var paths []string
		for _, c := range lock.Files {
			paths = append(paths, c.Path)
		}
		return paths
	}

	docsFirst := base(t, "include:\n  - docs/**\n  - src/**\n")
	if strings.Join(docsFirst, " ") != "docs/b.md src/a.go" {
		t.Errorf("docs first gives %v", docsFirst)
	}
	srcFirst := base(t, "include:\n  - src/**\n  - docs/**\n")
	if strings.Join(srcFirst, " ") != "src/a.go docs/b.md" {
		t.Errorf("src first gives %v", srcFirst)
	}
}

// A declared link's document is part of what the phase was given, whether or not include matches
// it. Section 5 has links declared and never inferred, and a declaration that put nothing in the
// base would be ornamental.
func TestADeclaredLinksDocumentIsInTheBase(t *testing.T) {
	f := newFixture(t)
	f.write("src/payment/card.go", "package payment\n")
	f.write("docs/adr/0012-payments.md", "# payments\n")
	f.write(model.PhaseDir(key, "00-intake")+"/"+model.ContextScope,
		"include:\n  - src/payment/**\nlinks:\n  - component: src/payment\n    docs: docs/adr/0012-payments.md\n")

	f.must(f.r.Start(key, "00-intake"))
	var lock model.ContextLock
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "context.lock.yaml"), &lock))

	var paths []string
	for _, c := range lock.Files {
		paths = append(paths, c.Path)
	}
	// The link goes last: it is the most specific thing in a scope and the most likely to move.
	if strings.Join(paths, " ") != "src/payment/card.go docs/adr/0012-payments.md" {
		t.Fatalf("the base is %v, want the include then the link's document", paths)
	}
	for _, c := range lock.Files {
		if c.SHA256 == "" {
			t.Errorf("%s is in the base with no hash, so G-Freshness cannot guard it", c.Path)
		}
	}
}

// Section 5 writes repo_commit into the lock and the tree never wrote it. Absent where the
// repository has none, because a lock saying which commit a phase ran against is a claim.
func TestTheLockRecordsTheCommitWhereThereIsOne(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	var lock model.ContextLock
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "context.lock.yaml"), &lock))
	if lock.RepoCommit != "" {
		t.Errorf("a fixture that is no git repository recorded commit %q", lock.RepoCommit)
	}
}

// rules_applied answers what rules_hash cannot: which rules, and which revision of each. Absent
// where no rule is in force, because an empty list says a set was resolved and came out empty.
func TestTheLockRecordsWhichRulesApplied(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.write(rules.ConfigDir+"/given/org/notes.yaml",
		"id: release-notes-say-something\nversion: 3\nscope: org\nkind: review\napplies_to: [00-intake]\n"+
			"statement: >\n  The release notes say what changed.\n")
	f.must(f.r.Start(key, "00-intake"))

	var lock model.ContextLock
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "context.lock.yaml"), &lock))
	if len(lock.RulesApplied) != 1 {
		t.Fatalf("rules_applied is %v, want the one rule in force", lock.RulesApplied)
	}
	got := lock.RulesApplied[0]
	if got.Path != rules.ConfigDir+"/given/org/notes.yaml" || got.Version != 3 {
		t.Fatalf("rules_applied carries %+v, want the path and the version counter", got)
	}

	// And the list agrees with the hash, because one resolution answers both.
	read, _ := rules.Load(f.root)
	effective, _ := rules.Effective(read)
	if len(effective) != len(lock.RulesApplied) {
		t.Errorf("%d rules in force and %d in the lock", len(effective), len(lock.RulesApplied))
	}

	// With no rule tree the field is absent, not empty.
	g := newFixture(t)
	g.project(agentBlock)
	g.templated()
	g.must(g.r.Start(key, "00-intake"))
	var bare model.ContextLock
	g.must(fm.ReadYAML(filepath.Join(g.root, model.PhaseDir(key, "00-intake"), "context.lock.yaml"), &bare))
	if bare.RulesApplied != nil {
		t.Errorf("rules_applied is %v where no rule is in force, want absent", bare.RulesApplied)
	}
}

// WP8's first clause: "a repeated phase reads only what changed". The runner cannot read for the
// agent, so what it does is say which files of the predecessor's base moved — derived from that
// lock and the tree, printed rather than recorded, because section 5's field list has no entry for
// it and the lock states what was declared (#171).
func TestChangedSinceNamesWhatMovedAndNothingElse(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.write("src/a.go", "package a\n")
	f.write("src/b.go", "package b\n")
	f.write(model.PhaseDir(key, "00-intake")+"/"+model.ContextScope, "include:\n  - src/**\n")

	// The first phase has no predecessor, so there is nothing to compare.
	if got := f.r.ChangedSince(key, "00-intake"); got != nil {
		t.Fatalf("the first phase reports %v", got)
	}
	f.run("00-intake", "")

	// Nothing moved: nothing to say.
	if got := f.r.ChangedSince(key, "01-requirements"); len(got) != 0 {
		t.Fatalf("an untouched tree reports %v", got)
	}

	// One file edited, one left alone, one of the base removed.
	f.write("src/a.go", "package a // changed\n")
	if got := f.r.ChangedSince(key, "01-requirements"); strings.Join(got, " ") != "src/a.go" {
		t.Fatalf("after one edit the changed set is %v, want src/a.go alone", got)
	}
	if err := os.Remove(filepath.Join(f.root, "src/b.go")); err != nil {
		t.Fatal(err)
	}
	got := f.r.ChangedSince(key, "01-requirements")
	if strings.Join(got, " ") != "src/a.go src/b.go (gone)" {
		t.Fatalf("the changed set is %v, want the edit and the removal, in the base's order", got)
	}
}

// A scope that names nothing in the tree declared no base, so nothing can have moved.
//
// This asserted an absent scope until #217, and that case is now unreachable: a P0 cannot be
// finished without one, so the silence belongs to a scope that resolved and came out empty
// rather than to one that was never written. A74's distinction, arriving where it belongs.
func TestChangedSinceIsSilentWhereTheScopeMatchesNothing(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.write(model.PhaseDir(key, model.Phases[0])+"/"+model.ContextScope, "include:\n  - vendor/**\n")
	f.write("src/a.go", "package a\n")
	f.run("00-intake", "")
	f.write("src/a.go", "package a // changed\n")
	if got := f.r.ChangedSince(key, "01-requirements"); got != nil {
		t.Fatalf("a scope matching nothing reports %v", got)
	}
}

// ---- the intent is created by a command (#179)

// trackerBlock is the configuration the qualified id is completed from. The address is
// GitHub's, because that is the one host whose API is not served under the id's host and
// therefore the only row where the derivation does anything.
const trackerBlock = "tracker:\n  adapter: github\n  project: triplem/xeno\n" +
	"  base_url: https://api.github.com\n"

// TestStartingAnIntentDerivesEverythingButTheIssue is the point of the command: the issue
// goes in, and the other six fields come from the runner and the configuration rather than
// from somebody's memory of them.
func TestStartingAnIntentDerivesEverythingButTheIssue(t *testing.T) {
	f := newFixture(t)
	f.project(trackerBlock)

	in, err := f.r.IntentStart("", "176")
	f.must(err)
	if in.Key != "PROJ-2" {
		t.Errorf("key %q, want PROJ-2: the sequence has PROJ-1 in it", in.Key)
	}
	var on model.Intent
	f.must(fm.ReadYAML(filepath.Join(f.root, model.IntentDir(in.Key), "intent.yaml"), &on))
	want := model.Intent{
		Intent: "github.com/triplem/xeno#176", Key: "PROJ-2", Status: "in-progress",
		Created: "2026-09-20T10:00:00Z", SchemaVersion: model.SchemaVersion,
		// From the vendored plugin's manifest, not a constant: the field names the plugin
		// an artifact was rendered from, and a number that cannot disagree with the runner
		// can never be proved wrong (#177).
		RunnerVersion: model.RunnerVersion, PluginVersion: "9.9.9",
	}
	if on != want {
		t.Errorf("intent.yaml is\n%+v\nwant\n%+v", on, want)
	}
}

// The version fields are the binary's own, which is what #177 is about at the intent
// level: a hand written intent.yaml records 0.1.0-dev beside artifacts that record the
// commit as well, and two strings for one build in one directory say nothing.
func TestANewIntentRecordsTheVersionTheArtifactsRecord(t *testing.T) {
	f := newFixture(t)
	f.project(trackerBlock)
	f.mustIntent(f.r.IntentStart("NEW-1", "176"))

	f.run("00-intake", "")
	var in model.Intent
	f.must(fm.ReadYAML(filepath.Join(f.root, model.IntentDir("NEW-1"), "intent.yaml"), &in))
	g, err := f.r.readGate(key, "00-intake")
	f.must(err)
	if in.RunnerVersion != g.RunnerVersion || in.PluginVersion != g.PluginVersion {
		t.Errorf("intent.yaml records %s/%s and gate.yaml %s/%s; one binary, one pair",
			in.RunnerVersion, in.PluginVersion, g.RunnerVersion, g.PluginVersion)
	}
}

// A key given is used as given. That is what the first intent of a repository needs, since
// there is no sequence to continue, and it is how the older naming scheme stays reachable.
func TestAGivenKeyIsUsedAsGiven(t *testing.T) {
	f := newFixture(t)
	f.project(trackerBlock)

	in, err := f.r.IntentStart("XENO-0300", "176")
	f.must(err)
	if in.Key != "XENO-0300" {
		t.Fatalf("key %q, want XENO-0300", in.Key)
	}
	if !fm.Exists(filepath.Join(f.root, model.IntentDir("XENO-0300"), "intent.yaml")) {
		t.Error("intent.yaml was not written where the key says")
	}
}

// Running it twice refuses rather than overwriting. intent.yaml sits inside the intent
// level hash, so a second write would change a verdict that named the first one.
func TestStartingAnIntentTwiceIsRefused(t *testing.T) {
	f := newFixture(t)
	f.project(trackerBlock)
	f.mustIntent(f.r.IntentStart("NEW-1", "176"))

	_, err := f.r.IntentStart("NEW-1", "177")
	var ref *Refusal
	if !errors.As(err, &ref) {
		t.Fatalf("second start: %v, want a refusal", err)
	}
	var in model.Intent
	f.must(fm.ReadYAML(filepath.Join(f.root, model.IntentDir("NEW-1"), "intent.yaml"), &in))
	if in.Intent != "github.com/triplem/xeno#176" {
		t.Errorf("the first id was overwritten: %q", in.Intent)
	}
}

// The issue is the one input, so its absence is a refusal and not a default. A derived
// everything is only worth having if the one thing nobody can derive is required.
func TestStartingAnIntentNeedsTheIssue(t *testing.T) {
	f := newFixture(t)
	f.project(trackerBlock)

	_, err := f.r.IntentStart("NEW-1", "")
	var ref *Refusal
	if !errors.As(err, &ref) || !strings.Contains(ref.Reason, "--for") {
		t.Fatalf("start without an issue: %v, want a refusal naming --for", err)
	}
	if fm.Exists(filepath.Join(f.root, model.IntentDir("NEW-1"))) {
		t.Error("a refused start left a directory behind")
	}
}

// A repository with no tracker block has nothing to complete the id from, and the whole
// qualified id given by hand is the way through. The block is optional per Appendix A.
func TestWithoutATrackerBlockTheWholeIdIsGivenByHand(t *testing.T) {
	f := newFixture(t)

	in, err := f.r.IntentStart("NEW-1", "git.example/group/proj#4")
	f.must(err)
	if in.Intent != "git.example/group/proj#4" {
		t.Errorf("id %q, want it as given", in.Intent)
	}
	if _, err := f.r.IntentStart("NEW-2", "4"); err == nil {
		t.Error("a bare key with no configuration behind it was accepted")
	}
}

// The two things the issue asks for after the file exists: the new intent is listed like
// any other, and phase start reads its intent.yaml as it reads a hand written one.
func TestANewIntentIsListedAndItsPhaseStarts(t *testing.T) {
	f := newFixture(t)
	f.project(trackerBlock)
	f.mustIntent(f.r.IntentStart("NEW-1", "176"))

	got, err := f.r.Intents()
	f.must(err)
	var found *IntentSummary
	for i := range got {
		if got[i].Key == "NEW-1" {
			found = &got[i]
		}
	}
	if found == nil {
		t.Fatalf("NEW-1 is not in the listing: %v", keysOf(t, got))
	}
	if found.Problem != "" {
		t.Errorf("the listing reports %q about a file the runner wrote", found.Problem)
	}
	if found.Created == "" || found.State != "no phases" {
		t.Errorf("listed as created %q, state %q", found.Created, found.State)
	}
	f.must(f.r.Start("NEW-1", "00-intake"))
}

// mustIntent discards the record and keeps the error, as must2 does for a verdict: a test
// that only needs the intent to exist says so without naming a variable it ignores.
func (f *fixture) mustIntent(_ *model.Intent, err error) { f.t.Helper(); f.must(err) }

// ---- the harness reports the one field the runner cannot know (#181)

// set writes one section and keeps the error. These tests are about the frontmatter the write
// produces and never about the resolved template, so naming a variable for it would be naming
// one to ignore it.
func (f *fixture) set(key, phase, section, content string) {
	f.t.Helper()
	_, err := f.r.SectionSet(key, phase, section, content)
	f.must(err)
}

// TestTheReportedToolVersionReachesBothArtifacts is the point of `--tool-version`: the field
// G-Schema requires, written by a command into the two files that carry it, with no edit to
// either. The digest takes it from `output.md` rather than from a second input.
func TestTheReportedToolVersionReachesBothArtifacts(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.r.ToolVersion = "2.1.276"
	f.must(f.r.Start(key, "00-intake"))
	f.set(key, "00-intake", "problem", "what is wrong")

	if got := f.frontField("00-intake", "tool_version"); got != "2.1.276" {
		t.Errorf("output.md records tool_version %q, want 2.1.276", got)
	}
	f.must2(f.r.Finish(key, "00-intake", "a summary"))
	front, _ := digestFront(t, f.root, "00-intake")
	if front["tool_version"] != "2.1.276" {
		t.Errorf("digest.md records tool_version %v, want 2.1.276", front["tool_version"])
	}
}

// The digest is rewritten by every `phase finish`, which is where the hand edit was paid
// again and again. A second finish that reports nothing keeps what the artifact beside it
// records, because the two describe one session and `output.md` is the one that says so.
func TestASecondFinishKeepsTheToolVersionWithoutBeingToldAgain(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.r.ToolVersion = "2.1.276"
	f.must(f.r.Start(key, "00-intake"))
	f.set(key, "00-intake", "problem", "what is wrong")
	f.must2(f.r.Finish(key, "00-intake", "a summary"))

	f.r.ToolVersion = ""
	f.must2(f.r.Finish(key, "00-intake", "a summary"))

	front, _ := digestFront(t, f.root, "00-intake")
	if front["tool_version"] != "2.1.276" {
		t.Errorf("the second finish lost it: tool_version %v", front["tool_version"])
	}
}

// A reported version wins over the artifact's, which is the case of a digest written before
// any section was: there is no output.md to carry anything over from, and the harness saying
// what it is cannot be overruled by a file that does not exist.
func TestAReportedVersionBeatsTheRecordedOne(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.r.ToolVersion = "2.1.276"
	f.must(f.r.Start(key, "00-intake"))
	f.set(key, "00-intake", "problem", "what is wrong")

	f.r.ToolVersion = "2.2.0"
	f.must2(f.r.Finish(key, "00-intake", "a summary"))
	front, _ := digestFront(t, f.root, "00-intake")
	if front["tool_version"] != "2.2.0" {
		t.Errorf("the reported version did not win: tool_version %v", front["tool_version"])
	}
}

// Nothing reported is the field absent, which is A35's rule and the state of every artifact
// written before there was a way to report it. A plausible value here would be worse than
// none, so neither writer invents one and G-Schema goes on reporting it missing.
func TestWithoutAReportTheFieldStaysAbsent(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	f.set(key, "00-intake", "problem", "what is wrong")
	f.must2(f.r.Finish(key, "00-intake", "a summary"))

	if got := f.frontField("00-intake", "tool_version"); got != "" {
		t.Errorf("output.md invented a tool_version: %q", got)
	}
	front, _ := digestFront(t, f.root, "00-intake")
	if _, ok := front["tool_version"]; ok {
		t.Errorf("digest.md invented a tool_version: %v", front["tool_version"])
	}
}

// Section 7 lists XENO_HARNESS_VERSION for a value that holds for a whole session, so the
// runner reads it where no flag said otherwise. This is the same assertion as the first test
// in this section, reached through the channel the specification names rather than an
// argument (#183).
func TestTheHarnessVersionIsReadFromTheEnvironment(t *testing.T) {
	f := newFixture(t)
	// Set after the fixture cleared it, and the runner built again, because New is where
	// the variable is read and that reading is what this asserts.
	t.Setenv(HarnessVersionEnv, "2.1.276")
	f.reopen()
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	f.set(key, "00-intake", "problem", "what is wrong")
	f.must2(f.r.Finish(key, "00-intake", "a summary"))

	if got := f.frontField("00-intake", "tool_version"); got != "2.1.276" {
		t.Errorf("output.md records tool_version %q, want 2.1.276", got)
	}
	front, _ := digestFront(t, f.root, "00-intake")
	if front["tool_version"] != "2.1.276" {
		t.Errorf("digest.md records tool_version %v, want 2.1.276", front["tool_version"])
	}
}

// A later section write that reports nothing does not erase what an earlier one recorded.
// The field describes the session and not the invocation, and `SectionSet` carries the
// frontmatter of an existing artifact over, which is what makes that true.
func TestALaterSectionWriteDoesNotEraseTheVersion(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.r.ToolVersion = "2.1.276"
	f.must(f.r.Start(key, "00-intake"))
	f.set(key, "00-intake", "problem", "what is wrong")

	f.r.ToolVersion = ""
	f.set(key, "00-intake", "scope", "what is in")

	if got := f.frontField("00-intake", "tool_version"); got != "2.1.276" {
		t.Errorf("the second section write lost it: %q", got)
	}
}

// ---- the learning record has a writer (#195)

// TestTheLearningRecordCarriesTheRunnersHeader is the point of the command: the four
// artifacts of one phase agree about which build produced them. A hand-written record
// carries 0.1.0-dev where the three beside it carry the commit, which is what #179 said
// about intent.yaml and is why this exists.
func TestTheLearningRecordCarriesTheRunnersHeader(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	f.set(key, "00-intake", "problem", "what is wrong")

	rec, err := f.r.RecordLearning(key, "00-intake", false, model.LearningEntry{
		Category: "context-rule", Observation: "o", Proposal: "p", Target: "t",
	})
	f.must(err)
	if rec.RunnerVersion != model.RunnerVersion || rec.PluginVersion != plugin.Version(f.root) {
		t.Errorf("header is %s/%s, want the runner's %s and the plugin's %s",
			rec.RunnerVersion, rec.PluginVersion, model.RunnerVersion, plugin.Version(f.root))
	}
	var on model.Learning
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "learning.yaml"), &on))
	if on.Phase != "00-intake" || on.Intent == "" || on.Created == "" {
		t.Errorf("the record on disk carries %+v", on.Common)
	}
	// The same binary wrote the verdict beside it, so the two have to agree.
	f.must2(f.r.Finish(key, "00-intake", "a summary"))
	g, err := f.r.readGate(key, "00-intake")
	f.must(err)
	if on.RunnerVersion != g.RunnerVersion {
		t.Errorf("learning.yaml says %s and gate.yaml says %s; one build, one string",
			on.RunnerVersion, g.RunnerVersion)
	}
}

// Entries accumulate, because a phase learns more than one thing often enough and the
// alternative is remembering the first one and retyping it.
func TestLearningEntriesAccumulate(t *testing.T) {
	f := newFixture(t)
	for _, c := range []string{"context-rule", "prompt"} {
		_, err := f.r.RecordLearning(key, "00-intake", false, model.LearningEntry{
			Category: c, Observation: "o", Proposal: "p", Target: "t",
		})
		f.must(err)
	}
	rec, err := f.r.RecordLearning(key, "00-intake", false, model.LearningEntry{
		Category: "template", Observation: "o", Proposal: "p", Target: "t",
	})
	f.must(err)
	if len(rec.Learnings) != 3 {
		t.Fatalf("records %d entries, want 3", len(rec.Learnings))
	}
	if rec.Learnings[0].Category != "context-rule" || rec.Learnings[2].Category != "template" {
		t.Errorf("the order moved: %+v", rec.Learnings)
	}
}

// The empty case is said rather than left out, which is section 10's honest record, and
// it is written to the same file with the same header.
func TestNoFindingIsTheEmptyRecord(t *testing.T) {
	f := newFixture(t)
	rec, err := f.r.RecordLearning(key, "00-intake", true, model.LearningEntry{})
	f.must(err)
	if !rec.NoFinding || len(rec.Learnings) != 0 {
		t.Fatalf("no_finding=%v with %d entries", rec.NoFinding, len(rec.Learnings))
	}
	if rec.RunnerVersion != model.RunnerVersion {
		t.Errorf("the empty record carries %s", rec.RunnerVersion)
	}
}

// A record cannot both state there was nothing and carry something. Refused in both
// directions, because either way one of the two statements is already on disk and a
// writer that resolved it silently would be deciding which one was meant.
func TestNoFindingAndAnObservationContradict(t *testing.T) {
	f := newFixture(t)
	f.mustLearning(f.r.RecordLearning(key, "00-intake", false, model.LearningEntry{
		Category: "prompt", Observation: "o", Proposal: "p", Target: "t",
	}))
	_, err := f.r.RecordLearning(key, "00-intake", true, model.LearningEntry{})
	assertRefusal(t, err, "contradicts")

	f2 := newFixture(t)
	f2.mustLearning(f2.r.RecordLearning(key, "01-requirements", true, model.LearningEntry{}))
	_, err = f2.r.RecordLearning(key, "01-requirements", false, model.LearningEntry{
		Category: "prompt", Observation: "o", Proposal: "p", Target: "t",
	})
	assertRefusal(t, err, "contradicts")
}

// Section 10 fixes the category set and all four keys, and the writer checks them so that
// G-Learning never has to report what a command could have refused.
func TestALearningIsRefusedBeforeItReachesTheFile(t *testing.T) {
	cases := []struct {
		name  string
		entry model.LearningEntry
		noFin bool
		says  string
	}{
		{"a category outside the four", model.LearningEntry{
			Category: "nonsense", Observation: "o", Proposal: "p", Target: "t"}, false, "--category is one of"},
		{"no observation", model.LearningEntry{
			Category: "prompt", Proposal: "p", Target: "t"}, false, "--observation is required"},
		{"no proposal", model.LearningEntry{
			Category: "prompt", Observation: "o", Target: "t"}, false, "--proposal is required"},
		{"no target", model.LearningEntry{
			Category: "prompt", Observation: "o", Proposal: "p"}, false, "--target is required"},
		{"whitespace is not a target", model.LearningEntry{
			Category: "prompt", Observation: "o", Proposal: "p", Target: "   "}, false, "--target is required"},
		{"nothing at all", model.LearningEntry{}, false, "a learning needs"},
		{"no-finding beside an entry", model.LearningEntry{Category: "prompt"}, true, "takes none of"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := newFixture(t)
			_, err := f.r.RecordLearning(key, "00-intake", c.noFin, c.entry)
			assertRefusal(t, err, c.says)
			if fm.Exists(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "learning.yaml")) {
				t.Error("a refused record reached the file")
			}
		})
	}
}

// Section 10 owes a record at the end of every phase and once more when an intent closes,
// so the phase is optional and its absence means the intent level record G-Complete reads.
func TestWithoutAPhaseTheRecordIsTheIntentsOwn(t *testing.T) {
	f := newFixture(t)
	rec, err := f.r.RecordLearning(key, "", true, model.LearningEntry{})
	f.must(err)
	if rec.Phase != "" {
		t.Errorf("an intent level record carries phase %q", rec.Phase)
	}
	if !fm.Exists(filepath.Join(f.root, model.IntentDir(key), "learning.yaml")) {
		t.Error("it was not written at intent level")
	}
	if fm.Exists(filepath.Join(f.root, model.PhaseDir(key, "00-intake"), "learning.yaml")) {
		t.Error("it was written into a phase as well")
	}
}

// assertRefusal is the shape every refusal test in this file wants: the error is a
// Refusal and its reason names what the caller has to change.
func assertRefusal(t *testing.T, err error, says string) {
	t.Helper()
	var ref *Refusal
	if !errors.As(err, &ref) {
		t.Fatalf("got %v, want a refusal", err)
	}
	if !strings.Contains(ref.Reason, says) {
		t.Errorf("the refusal reads %q, which does not say %q", ref.Reason, says)
	}
}

// mustLearning discards the record and keeps the error, as must2 does for a verdict.
func (f *fixture) mustLearning(_ *model.Learning, err error) { f.t.Helper(); f.must(err) }

// ---- the plugin says its own version (#177)

// TestThePluginVersionComesFromTheVendoredPlugin is the point: the field names the plugin an
// artifact was rendered from. It was a constant, and a release set it to the runner's own
// version through ldflags, so a number that could not disagree with the runner could never be
// proved wrong — which is the one thing section 5 wants it for.
func TestThePluginVersionComesFromTheVendoredPlugin(t *testing.T) {
	f := newFixture(t)
	f.write(plugin.Dir+"/.claude-plugin/plugin.json", `{"name":"xeno","version":"0.26.0"}`)
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))
	f.set(key, "00-intake", "problem", "what is wrong")

	if got := f.frontField("00-intake", "plugin_version"); got != "0.26.0" {
		t.Errorf("output.md records plugin_version %q, want the manifest's 0.26.0", got)
	}
	if got := f.frontField("00-intake", "plugin_version"); got == model.RunnerVersion {
		t.Error("plugin_version is the runner's version again")
	}
}

// Absent where no plugin is vendored, which is A35's rule. Section 5 requires the field in
// every process file, so the absence is a G-Schema finding — a repository with no plugin
// rendered from none, and that is the honest thing for its artifacts to say.
func TestWithoutAVendoredPluginTheVersionIsAbsentAndReported(t *testing.T) {
	f := newFixture(t)
	f.must(os.RemoveAll(filepath.Join(f.root, plugin.Dir)))
	f.project(agentBlock)

	in, err := f.r.IntentStart("NEW-1", "git.example/g/p#4")
	f.must(err)
	if in.PluginVersion != "" {
		t.Errorf("plugin_version is %q, want absent", in.PluginVersion)
	}
	g := f.run("00-intake", "")
	var said bool
	for _, c := range g.Checks {
		for _, fd := range c.Findings {
			if c.Gate == "G-Schema" && strings.Contains(fd.Cause, "plugin_version") {
				said = true
			}
		}
	}
	if !said {
		t.Error("no plugin and no finding about it; section 5 requires the field")
	}
}

// The lock carries section 5's plugin block, which is the other half of the sentence that
// explains the pair: the frontmatter names what was used and the lock proves it with a hash.
func TestTheLockCarriesThePluginsVersionAndHash(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock)
	f.templated()
	f.must(f.r.Start(key, "00-intake"))

	var lock model.ContextLock
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"),
		"context.lock.yaml"), &lock))
	if lock.Plugin == nil {
		t.Fatal("the lock carries no plugin block")
	}
	if lock.Plugin.Version != "9.9.9" {
		t.Errorf("the lock records plugin version %q, want the manifest's", lock.Plugin.Version)
	}
	if len(lock.Plugin.SHA256) != 64 {
		t.Errorf("the lock records a plugin hash %q, want sixty four hex characters",
			lock.Plugin.SHA256)
	}
	if got := lock.PluginVersion; got != lock.Plugin.Version {
		t.Errorf("the lock's header says %q and its plugin block says %q", got, lock.Plugin.Version)
	}
}

// Without a plugin the block is absent rather than half written: a version with no hash, or a
// hash with no version, would be half a claim.
func TestWithoutAPluginTheLockCarriesNoBlock(t *testing.T) {
	f := newFixture(t)
	f.must(os.RemoveAll(filepath.Join(f.root, plugin.Dir)))
	f.must(f.r.Start(key, "00-intake"))

	var lock model.ContextLock
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, "00-intake"),
		"context.lock.yaml"), &lock))
	if lock.Plugin != nil {
		t.Errorf("the lock carries a plugin block with no plugin: %+v", *lock.Plugin)
	}
}

// Section 7 lists XENO_HARNESS and calls it recorded only. It beats the project's declaration,
// because `agent.tool` is what a project says it uses and the variable is what is running.
func TestTheHarnessIsRecordedOverTheProjectsDeclaration(t *testing.T) {
	f := newFixture(t)
	f.project(agentBlock) // declares claude-code
	f.templated()
	t.Setenv(HarnessEnv, "codex")
	f.must(f.r.Start(key, "00-intake"))
	f.set(key, "00-intake", "problem", "what is wrong")

	if got := f.frontField("00-intake", "tool"); got != "codex" {
		t.Errorf("tool is %q, want the harness's codex over the project's claude-code", got)
	}
	// And with nothing in the environment the project's declaration stands, which is what
	// A35's second amendment settled and this does not undo.
	t.Setenv(HarnessEnv, "")
	f.set(key, "00-intake", "scope", "what is in")
	if got := f.frontField("00-intake", "tool"); got != "claude-code" {
		t.Errorf("tool is %q, want the project's claude-code when nothing is exported", got)
	}
}

// ---- the merge check: an intent the change touches has finished or been closed (#206)

// gitTree makes the fixture's root a repository and commits everything in it, returning the
// commit as the base of a range. A real repository, because what is under test is which paths
// git reports differ and a fake would be a second implementation of that.
func (f *fixture) gitTree() string {
	f.t.Helper()
	if _, err := exec.LookPath("git"); err != nil {
		f.t.Skip("git is not on the path")
	}
	for _, args := range [][]string{
		{"init", "--initial-branch=main"},
		{"config", "user.name", "A Committer"},
		{"config", "user.email", "committer@example.test"},
		{"config", "commit.gpgsign", "false"},
	} {
		f.gitRun(args...)
	}
	return f.gitCommit("the base")
}

func (f *fixture) gitRun(args ...string) string {
	f.t.Helper()
	out, err := exec.Command("git", append([]string{"-C", f.root}, args...)...).CombinedOutput()
	if err != nil {
		f.t.Fatalf("git %s: %v\n%s", strings.Join(args, " "), err, out)
	}
	return strings.TrimSpace(string(out))
}

// gitCommit commits whatever is in the tree, allowing an empty commit so that a test can make
// a range that changes nothing under .xeno/intents/.
func (f *fixture) gitCommit(message string) string {
	f.t.Helper()
	f.gitRun("add", "-A")
	f.gitRun("commit", "--allow-empty", "-m", message)
	return f.gitRun("rev-parse", "HEAD")
}

func (f *fixture) completeness(base string) *CompletenessResult {
	f.t.Helper()
	res, err := f.r.Completeness(base, "HEAD")
	f.must(err)
	return res
}

func TestAnIntentThatReachedADecidedP5Passes(t *testing.T) {
	f := newFixture(t)
	f.templated()
	base := f.gitTree()
	for _, p := range model.Phases {
		f.run(p, "")
	}
	f.gitCommit("the whole intent")

	res := f.completeness(base)
	if strings.Join(res.Touched, ",") != key {
		t.Fatalf("touched %v, want just the intent the change wrote", res.Touched)
	}
	if len(res.Unfinished) != 0 {
		t.Fatalf("a complete intent was reported unfinished: %+v", res.Unfinished)
	}
}

// The case #206 was filed for. XENO-0230 reached main with its verification and review phases
// missing and every gate green on the phases it did write, and nothing said so.
func TestAnIntentThatStopsShortIsReportedWithTheStateItReached(t *testing.T) {
	f := newFixture(t)
	f.templated()
	base := f.gitTree()
	for _, p := range model.Phases[:4] {
		f.run(p, "")
	}
	f.gitCommit("four phases and no more")

	res := f.completeness(base)
	if len(res.Unfinished) != 1 || res.Unfinished[0].Key != key {
		t.Fatalf("want the one intent that stopped, got %+v", res.Unfinished)
	}
	if got := res.Unfinished[0].State; got != "03-implementation" {
		t.Errorf("reported as %q, want the phase it reached", got)
	}
}

// An abandoned intent is an ending, not a gap. It is the case `xeno intent close` exists for,
// and the whole point of the check is that it distinguishes the two.
func TestAnAbandonedIntentPasses(t *testing.T) {
	f := newFixture(t)
	base := f.gitTree()
	f.abandonedIntent("PROJ-9", "the requirement went away")
	f.gitCommit("closed as abandoned")

	res := f.completeness(base)
	if strings.Join(res.Touched, ",") != "PROJ-9" {
		t.Fatalf("touched %v, want the intent that was closed", res.Touched)
	}
	if len(res.Unfinished) != 0 {
		t.Fatalf("an abandoned intent was reported unfinished: %+v", res.Unfinished)
	}
}

// A P5 that is provisional is not an ending either. Section 6 reports a provisional verdict
// rather than failing it and says it becomes "a hard condition at the merge request"; this is
// the merge request, which is why `gate verify` still exits 0 on the same state.
func TestAProvisionalFinalPhaseIsNotFinished(t *testing.T) {
	f := newFixture(t)
	base := f.gitTree()
	for _, p := range model.Phases[:5] {
		f.run(p, "")
	}
	f.must(f.r.Start(key, "05-review"))
	f.output("05-review", pendingTest)
	if g := f.finish("05-review"); g.Status != "provisional" {
		t.Fatalf("the fixture P5 is %s, want provisional", g.Status)
	}
	f.gitCommit("a P5 waiting on a pipeline")

	res := f.completeness(base)
	if len(res.Unfinished) != 1 || res.Unfinished[0].State != "05-review" {
		t.Fatalf("a provisional P5 counted as finished: %+v", res.Unfinished)
	}
}

// The rule that every change belongs to an intent is real and unenforced, and #120 owns it.
// Enforcing it from inside a check about something else would bury it, so a range that touches
// no intent is not this command's finding.
func TestARangeThatTouchesNoIntentHasNothingToCheck(t *testing.T) {
	f := newFixture(t)
	base := f.gitTree()
	f.write("internal/thing.go", "package thing\n")
	f.gitCommit("code and no trail")

	res := f.completeness(base)
	if len(res.Touched) != 0 || len(res.Unfinished) != 0 {
		t.Fatalf("want nothing to check, got %+v", res)
	}
}

// The listing reports an unreadable intent rather than leaving it out, because a record missing
// from a listing is worse than one that looks wrong in it. Here the answer is a verdict, so the
// same fact has to refuse.
func TestAnIntentWhoseRecordCannotBeReadIsUnfinished(t *testing.T) {
	f := newFixture(t)
	base := f.gitTree()
	f.write(model.IntentDir("PROJ-9")+"/intent.yaml", "key: [this is not an intent\n")
	f.gitCommit("a record nothing can read")

	res := f.completeness(base)
	if len(res.Unfinished) != 1 || res.Unfinished[0].Key != "PROJ-9" {
		t.Fatalf("want the unreadable intent, got %+v", res.Unfinished)
	}
	if res.Unfinished[0].Problem == "" {
		t.Error("reported with no reason, which is the one thing a refusal owes")
	}
}

// Section 9: the range is passed in and never inferred. There is no base to fall back to, so a
// missing one is an error rather than a comparison against something the tool chose.
func TestTheMergeCheckRefusesWithoutBothEndsOfTheRange(t *testing.T) {
	f := newFixture(t)
	f.gitTree()
	if _, err := f.r.Completeness("", "HEAD"); err == nil {
		t.Error("an absent base was taken for a range")
	}
	if _, err := f.r.Completeness("HEAD", ""); err == nil {
		t.Error("an absent head was taken for a range")
	}
	if _, err := f.r.Completeness("no-such-ref", "HEAD"); err == nil {
		t.Error("a ref that does not resolve passed as a clean comparison")
	}
}

// ---- #205: the local data location is read

// Criterion 6: setting XENO_PLUGIN_DATA moves the run marker and phase.env, which is what the
// entry point's export has been promising since it was written. Nothing under there is covered
// by artifacts_hash, which is why reading the variable cannot make a verdict depend on the
// environment — the distinction from XENO_PLUGIN_ROOT that A89 removed.
func TestTheLocalDataLocationMovesWithTheVariable(t *testing.T) {
	f := newFixture(t)
	elsewhere := t.TempDir()
	t.Setenv(model.LocalDataEnv, elsewhere)

	f.must(f.r.Start(key, model.Phases[0]))

	for _, rel := range []string{
		filepath.Join("runs", key, model.Phases[0]+".lock"),
		PhaseEnvFile,
	} {
		if !fm.Exists(filepath.Join(elsewhere, rel)) {
			t.Errorf("%s was not written under the redirected location", rel)
		}
		if fm.Exists(filepath.Join(f.root, model.LocalDefault, rel)) {
			t.Errorf("%s was written under the default as well", rel)
		}
	}
}

// Criterion 7, end to end rather than over the resolver: with the variable unset the paths are
// exactly where every existing repository already has them.
func TestWithTheVariableUnsetTheLocalPathsDoNotMove(t *testing.T) {
	f := newFixture(t)
	t.Setenv(model.LocalDataEnv, "")

	f.must(f.r.Start(key, model.Phases[0]))

	for _, rel := range []string{
		filepath.Join("runs", key, model.Phases[0]+".lock"),
		PhaseEnvFile,
	} {
		if !fm.Exists(filepath.Join(f.root, model.LocalDefault, rel)) {
			t.Errorf("%s is not under %s, where it has always been", rel, model.LocalDefault)
		}
	}
}

// The enforcement report and the ledger resolve through the same function, so one setting moves
// all four files rather than two of them.
func TestTheReportAndTheLedgerResolveToTheSameDirectory(t *testing.T) {
	elsewhere := t.TempDir()
	t.Setenv(model.LocalDataEnv, elsewhere)
	root := t.TempDir()

	if got, want := ReportPath(root), filepath.Join(elsewhere, ReportFile); got != want {
		t.Errorf("the enforcement report resolves to %s, want %s", got, want)
	}
	if got, want := cost.LedgerPath(root), filepath.Join(elsewhere, cost.LedgerFile); got != want {
		t.Errorf("the ledger resolves to %s, want %s", got, want)
	}
}

// ---- #225: a second start that would stale the artifact

// moving replaces the fixture's pinned clock. newFixture fixes Now, so a rewritten lock is
// byte-identical and the artifact stays consistent with it — which is why this defect survived:
// a test in the package's usual style passes against the broken code. The damage is the created
// stamp moving, and only a clock that moves moves it.
func (f *fixture) moving() {
	n := 0
	f.r.Now = func() time.Time { n++; return time.Date(2026, 9, 20, 10, n, 0, 0, time.UTC) }
}

// The reproduction. Sections written, no verdict, marker gone — which is what a phase begun on
// another machine looks like here, and what retention leaves after thirty days. Before #225 the
// second start succeeded, rewrote the lock, and phase finish went red on G-Schema.
func TestAStartThatWouldStaleTheArtifactIsRefused(t *testing.T) {
	f := newFixture(t)
	f.moving()
	f.must(f.r.Start(key, model.Phases[0]))
	f.output(model.Phases[0], "")
	lock := filepath.Join(f.root, model.PhaseDir(key, model.Phases[0]), "context.lock.yaml")
	before, err := hashing.FileHash(lock)
	f.must(err)

	f.must(os.Remove(f.r.marker(key, model.Phases[0])))

	err = f.r.Start(key, model.Phases[0])
	if !isRefusal(err) {
		t.Fatalf("a start that would stale the artifact was allowed: %v", err)
	}
	for _, want := range []string{
		"already under way", "section set and phase finish", model.PhaseDir(key, model.Phases[0]),
	} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("the refusal does not say %q: %v", want, err)
		}
	}
	after, err := hashing.FileHash(lock)
	f.must(err)
	if before != after {
		t.Fatal("the refused start rewrote the lock anyway")
	}
}

// The case P2 expected to be harmless and is not. With the fixture's pinned clock a rewritten
// lock is byte-identical, so "the marker is gone and nothing else changed" looks safe; with a
// clock that moves, `created` moves and the artifact goes stale. There is no safe second start
// of a phase under way, which is why the condition is the artifact rather than a comparison
// against the lock.
func TestThereIsNoHarmlessSecondStartOfAPhaseUnderWay(t *testing.T) {
	f := newFixture(t)
	f.moving()
	f.must(f.r.Start(key, model.Phases[0]))
	f.output(model.Phases[0], "")
	f.must(os.Remove(f.r.marker(key, model.Phases[0])))

	if err := f.r.Start(key, model.Phases[0]); !isRefusal(err) {
		t.Fatalf("a second start of a phase under way was allowed: %v", err)
	}
}

// A phase under way is reported as running whether or not this machine holds the marker, which
// is the first of #225's three gaps: a phase begun on one machine read as not-started on
// another, and retention made it read that way on the same one after thirty days.
func TestAPhaseWithAnArtifactIsRunningWithoutItsMarker(t *testing.T) {
	f := newFixture(t)
	f.must(f.r.Start(key, model.Phases[0]))
	f.output(model.Phases[0], "")
	f.must(os.Remove(f.r.marker(key, model.Phases[0])))

	states, err := f.r.Status(key)
	f.must(err)
	if states[0].State != "running" {
		t.Fatalf("a phase with sections written and no verdict reads %q", states[0].State)
	}
}

// A first start has no artifact, so the condition cannot fire. This is every normal start.
func TestAFirstStartIsUnaffected(t *testing.T) {
	f := newFixture(t)
	f.moving()
	f.must(f.r.Start(key, model.Phases[0]))
}

// The two guards that were already there keep their conditions and their wording.
func TestTheOlderTwoRefusalsAreUnchanged(t *testing.T) {
	f := newFixture(t)
	f.moving()
	f.must(f.r.Start(key, model.Phases[0]))
	if err := f.r.Start(key, model.Phases[0]); !isRefusal(err) ||
		!strings.Contains(err.Error(), "already running") {
		t.Fatalf("a present marker no longer says already running: %v", err)
	}

	f.output(model.Phases[0], "")
	f.finish(model.Phases[0])
	if err := f.r.Start(key, model.Phases[0]); !isRefusal(err) ||
		!strings.Contains(err.Error(), "has a verdict") {
		t.Fatalf("a judged phase no longer says it has a verdict: %v", err)
	}
}

// The refusal and the printed next step have to point the same way, or somebody follows the
// suggestion into the thing just refused.
func TestTheSuggestionAgreesWithTheRefusal(t *testing.T) {
	f := newFixture(t)
	f.moving()
	f.must(f.r.Start(key, model.Phases[0]))
	f.output(model.Phases[0], "")
	f.must(os.Remove(f.r.marker(key, model.Phases[0])))

	s := f.r.Next(key)
	if s == nil {
		t.Fatal("no suggestion for a phase under way")
	}
	if strings.Contains(s.Command, "phase start") {
		t.Fatalf("the next step offers the start that is refused: %q", s.Command)
	}
}
