// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"example.com/xeno/internal/fm"
	"example.com/xeno/internal/model"
)

const key = "PROJ-1"

type fixture struct {
	t    *testing.T
	root string
	r    *Runner
}

func newFixture(t *testing.T) *fixture {
	root := t.TempDir()
	r := New(root)
	r.Now = func() time.Time { return time.Date(2026, 9, 20, 10, 0, 0, 0, time.UTC) }
	f := &fixture{t, root, r}
	f.write(model.IntentDir(key)+"/intent.yaml", "intent: \"git.example/group/proj#1\"\nkey: PROJ-1\nstatus: in-progress\n")
	f.write(model.IntentDir(key)+"/assumptions.yaml", "assumptions: []\n")
	return f
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
	session := "language: en\nsecrets_hash: s\ncontext_hash: c\nmodel: m\ntool: claude-code\ntool_version: 1\n"
	f.write(d+"/output.md", "---\n"+common+session+"template: t@1\nstrings_hash: s\nrules_hash: r\n"+extra+"---\n\n# Result\n")
	f.write(d+"/digest.md", "---\n"+common+session+"---\nsummary\n")
	f.write(d+"/learning.yaml", common+"no_finding: true\n")
}

func (f *fixture) must(err error) {
	f.t.Helper()
	if err != nil {
		f.t.Fatal(err)
	}
}

func (f *fixture) finish(phase string) *model.Gate {
	f.t.Helper()
	g, err := f.r.Finish(key, phase)
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

const pendingTest = "evidence:\n  - kind: test\n    job: unit\n"

func (f *fixture) pipeline(result string) string {
	dir := filepath.Join(f.t.TempDir(), "artifacts")
	_ = os.MkdirAll(dir, 0o755)
	os.WriteFile(filepath.Join(dir, "junit.xml"), []byte("<testsuite failures=\"0\"/>\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(
		"- kind: test\n  job: unit\n  result: "+result+"\n  file: junit.xml\n  pipeline: \"4711\"\n  commit: abc123\n"), 0o644)
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

const question = "open_questions:\n  - key: Q-1\n    text: which error behaviour?\n    options:\n      - text: fail fast\n        recommended: true\n      - text: retry\n      - text: free entry\n        free: true\n"

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
		"assumptions:\n  - id: A-1\n    text: fail fast is acceptable\n    resolves: Q-1\n    confirmed_by: m.example\n")
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
