// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
)

// The property WP1's done-when asks for and #304 settles: a command aimed at a phase that
// already holds a verdict leaves that phase exactly as the verdict found it. What it wrote
// is not lost, it waits under .xeno/local/staged/ until the next finish applies it.
func TestAWriteToAJudgedPhaseLeavesTheSealWhereItWas(t *testing.T) {
	f := newFixture(t)
	g := f.run("00-intake", "")
	sealed := g.ArtifactsHash
	before, err := os.ReadFile(f.r.abs(model.PhaseDir(key, "00-intake") + "/output.md"))
	f.must(err)

	mustAsk(f, "who decides this?")
	if f.hash("00-intake") != sealed {
		t.Fatal("the write moved the hash the verdict covers")
	}
	after, err := os.ReadFile(f.r.abs(model.PhaseDir(key, "00-intake") + "/output.md"))
	f.must(err)
	if string(after) != string(before) {
		t.Fatal("the sealed artifact was written into")
	}
	staged := filepath.Join(model.StagedDir(f.root, key, "00-intake"), "output.md")
	b, err := os.ReadFile(staged)
	if err != nil {
		t.Fatalf("nothing was staged: %v", err)
	}
	if !strings.Contains(string(b), "who decides this?") {
		t.Fatal("the staged copy does not carry what was written")
	}
}

// A second write amends the first rather than starting again from the seal, which is the
// case that would silently drop work: both writes are part of one redo.
func TestASecondWriteToAJudgedPhaseAmendsTheFirst(t *testing.T) {
	f := newFixture(t)
	f.run("00-intake", "")
	mustAsk(f, "the first question?")
	mustAsk(f, "the second question?")
	b, err := os.ReadFile(filepath.Join(model.StagedDir(f.root, key, "00-intake"), "output.md"))
	f.must(err)
	for _, want := range []string{"the first question?", "the second question?"} {
		if !strings.Contains(string(b), want) {
			t.Fatalf("the staged copy lost %q", want)
		}
	}
}

// Every content file of a phase goes the same way. The scope is P0's own artifact and the
// learning record is the phase's, and both lie inside artifacts_hash.
func TestTheScopeAndTheLearningRecordAreStagedToo(t *testing.T) {
	f := newFixture(t)
	g := f.run("00-intake", "")
	sealed := g.ArtifactsHash
	if _, err := f.r.ScopeSet(key, []byte("include:\n  - internal/**\n")); err != nil {
		t.Fatal(err)
	}
	// no_finding, because that is what the fixture's record already states and an
	// observation would contradict it. What is under test is where the write lands.
	if _, err := f.r.RecordLearning(key, "00-intake", true, model.LearningEntry{}); err != nil {
		t.Fatal(err)
	}
	if f.hash("00-intake") != sealed {
		t.Fatal("a write moved the hash the verdict covers")
	}
	staged := f.r.staged(key, "00-intake")
	if len(staged) != 2 || staged[0] != model.ContextScope || staged[1] != "learning.yaml" {
		t.Fatalf("staged %v, want the scope and the learning record", staged)
	}
}

// And the finish is what applies them: one command moves the content and the verdict, so no
// state exists in which the repository holds content its own verdict does not cover.
func TestTheFinishAppliesTheRedoAndJudgesIt(t *testing.T) {
	f := newFixture(t)
	g := f.run("00-intake", "")
	mustAsk(f, "the redone question?")
	g2, err := f.r.Finish(key, "00-intake", "")
	f.must(err)
	if g2.ArtifactsHash == g.ArtifactsHash {
		t.Fatal("the finish did not apply the redo")
	}
	if f.hash("00-intake") != g2.ArtifactsHash {
		t.Fatal("what was sealed is not what was judged")
	}
	b, err := os.ReadFile(f.r.abs(model.PhaseDir(key, "00-intake") + "/output.md"))
	f.must(err)
	if !strings.Contains(string(b), "the redone question?") {
		t.Fatal("the applied artifact does not carry the redo")
	}
	if left := f.r.staged(key, "00-intake"); left != nil {
		t.Fatalf("the staging directory outlived the finish: %v", left)
	}
}

// A redo nobody finished leaves the trail as it was, and is visible in the one place it can
// be: staged content is not committed, so the listing is what section 6 points a reader at.
func TestAnAbandonedRedoLeavesTheTrailAloneAndIsReported(t *testing.T) {
	f := newFixture(t)
	g := f.run("00-intake", "")
	mustAsk(f, "a question nobody applied?")

	states, err := f.r.Status(key)
	f.must(err)
	if states[0].State != "finished" || states[0].Status != g.Status {
		t.Fatalf("the phase stopped reading as judged: %+v", states[0])
	}
	if len(states[0].Staged) != 1 || states[0].Staged[0] != "output.md" {
		t.Fatalf("the staged redo is not reported: %+v", states[0])
	}
	if sug := f.r.Next(key); sug == nil || !strings.Contains(sug.Text, "redo staged") {
		t.Fatalf("the next step does not name the redo: %+v", sug)
	}
	// And `gate verify` has nothing to report, because nothing in the repository moved.
	if f.hash("00-intake") != g.ArtifactsHash {
		t.Fatal("the trail moved")
	}
}

// mustAsk records one open question into the fixture's intake, which is a write into
// output.md's frontmatter. This fixture writes output.md by hand and carries no template, so
// `question record` rather than `section set` is the writer these tests can use; both go
// through the same resolution, and exchange_test covers the frontmatter shape.
func mustAsk(f *fixture, text string) {
	f.t.Helper()
	if _, err := f.r.RecordQuestion(key, "00-intake", []byte("text: "+text+"\nno_options: true\n")); err != nil {
		f.t.Fatal(err)
	}
}

// A judged phase reads as judged even where this machine still holds the marker of a run
// that was never finished through it. The marker is local and expires at thirty days (A9),
// and `phase start` refuses a phase that holds a verdict, so there is no run it could be
// describing. Found on this repository, where one marker from 2026-10-03 made a six-phase
// intent read as five and took it out of the population `report figures` counts.
func TestALeftoverRunMarkerDoesNotHideAVerdict(t *testing.T) {
	f := newFixture(t)
	g := f.run("00-intake", "")
	if err := os.MkdirAll(filepath.Dir(f.r.marker(key, "00-intake")), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(f.r.marker(key, "00-intake"), []byte("stale\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	states, err := f.r.Status(key)
	f.must(err)
	if states[0].State != "finished" || states[0].Status != g.Status {
		t.Fatalf("a leftover marker hid the verdict: %+v", states[0])
	}
}
