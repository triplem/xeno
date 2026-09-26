// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/template"
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

func (f *fixture) must2(_ *model.Gate, err error) { f.t.Helper(); f.must(err) }

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
		{"rejected", "green"},
	} {
		f := newFixture(t)
		f.write(model.IntentDir(key)+"/assumptions.yaml", register+tc.status+"\n")
		if g := f.run("00-intake", ""); g.Status != tc.want {
			t.Errorf("status %q gave %s, wanted %s", tc.status, g.Status, tc.want)
		}
	}
}

// Confirmation is what the gate reads, so a confirmation with nobody behind it is a
// finding of its own rather than a pass.
func TestConfirmedByNobodyIsAFinding(t *testing.T) {
	f := newFixture(t)
	f.write(model.IntentDir(key)+"/assumptions.yaml", register+"confirmed\n")
	if g := f.run("00-intake", ""); g.Status != "red" {
		t.Fatalf("a confirmation with no person behind it passed: %s", g.Status)
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
