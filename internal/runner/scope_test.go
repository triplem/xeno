// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
)

// The artifact section 5 puts at P0 had no writer for a hundred intents, and the readers took
// its absence for a project that had declared nothing (#217). These are the two halves that
// close it: a command that produces it, and a refusal that asks for it.

// TestTheScopeIsWrittenAndItsPatternsAreReported is the writer's point: the figure comes back
// so that a budget is counted rather than guessed, which is what this intent's own P0 had to
// do by hand and got wrong the first time.
func TestTheScopeIsWrittenAndItsPatternsAreReported(t *testing.T) {
	f := newFixture(t)
	f.write("src/payment/card.go", "package payment\n")
	f.write("src/payment/testdata/golden.json", "{}\n")
	f.write("docs/adr/0012-payments.md", "# payments\n")

	res, err := f.r.ScopeSet(key, []byte("include:\n  - src/payment/**\n  - docs/adr/*.md\n"+
		"exclude:\n  - \"**/testdata/**\"\nbudget:\n  files: 10\n  bytes: 4000\n"))
	f.must(err)
	if res.Files != 2 || res.Bytes != int64(len("package payment\n")+len("# payments\n")) {
		t.Fatalf("the patterns resolve to %d files and %d bytes", res.Files, res.Bytes)
	}

	var s model.Scope
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, model.Phases[0]), model.ContextScope), &s))
	// The header is the runner's, as it is in every other artifact.
	if s.Intent != "git.example/group/proj#1" || s.Phase != model.Phases[0] ||
		s.SchemaVersion == "" || s.RunnerVersion == "" {
		t.Fatalf("the header was not written: %+v", s.Common)
	}
	if len(s.Include) != 2 || s.Budget.Files != 10 {
		t.Fatalf("the content was not written as given: %+v", s)
	}
}

// The reported figure and the recorded one come from one resolution, because the budget is set
// from the first and judged against the second.
func TestTheReportedFigureIsTheOneTheLockRecords(t *testing.T) {
	f := newFixture(t)
	f.write("src/a.go", "package a\n")
	f.write("src/b.go", "package b\n")

	res, err := f.r.ScopeSet(key, []byte("include:\n  - src/**\n"))
	f.must(err)
	f.must(f.r.Start(key, model.Phases[0]))

	var lock model.ContextLock
	f.must(fm.ReadYAML(filepath.Join(f.root, model.PhaseDir(key, model.Phases[0]), "context.lock.yaml"), &lock))
	if len(lock.Files) != res.Files {
		t.Fatalf("the writer reported %d files and the lock records %d", res.Files, len(lock.Files))
	}
	var bytes int64
	for _, c := range lock.Files {
		bytes += c.Bytes
	}
	if bytes != res.Bytes {
		t.Fatalf("the writer reported %d bytes and the lock records %d", res.Bytes, bytes)
	}
}

// Three refusals, each for a reason the standing rules or A35 give.
func TestTheScopeWriterRefusesWhatItCannotMean(t *testing.T) {
	for _, tc := range []struct{ name, entry, want string }{
		{"an invented field", "include:\n  - src/**\nbudgets:\n  files: 1\n", "could not be read"},
		{"no include", "exclude:\n  - src/**\n", "needs include"},
		{"a header a person wrote", "intent: someone/else#1\ninclude:\n  - src/**\n",
			"written by the runner and not given"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newFixture(t)
			_, err := f.r.ScopeSet(key, []byte(tc.entry))
			if err == nil {
				t.Fatalf("%s was accepted", tc.name)
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("the refusal says %q, want %q", err, tc.want)
			}
		})
	}
}

// D-2: the requirement is in the command, because Verify recomputes a verdict and compares it,
// so a check in a gate would re-judge every phase already sealed.
func TestP0CannotBeFinishedWithoutAScope(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.must(f.r.Start(key, model.Phases[0]))
	f.output(model.Phases[0], "")
	f.must(os.Remove(filepath.Join(f.root, model.PhaseDir(key, model.Phases[0]), model.ContextScope)))

	_, err := f.r.Finish(key, model.Phases[0], "the summary of a refused finish")
	if err == nil {
		t.Fatal("a P0 with no scope was judged")
	}
	if !strings.Contains(err.Error(), model.ContextScope) ||
		!strings.Contains(err.Error(), "xeno scope set") {
		t.Fatalf("the refusal does not name the artifact and the remedy: %s", err)
	}
	// Nothing is written by a refused finish, which is what putting the check first buys. The
	// digest is asserted by its content rather than its presence, because the fixture writes
	// one when it writes the artifact.
	if fm.Exists(filepath.Join(f.root, model.PhaseDir(key, model.Phases[0]), "gate.yaml")) {
		t.Fatal("a refused finish wrote a verdict")
	}
	b, err := os.ReadFile(filepath.Join(f.root, model.PhaseDir(key, model.Phases[0]), "digest.md"))
	f.must(err)
	if strings.Contains(string(b), "the summary of a refused finish") {
		t.Fatal("a refused finish wrote the digest")
	}
}

// The requirement is P0's alone: a later phase has no scope of its own, so asking it for one
// would be asking for an artifact section 5 does not give it.
func TestAPhaseAfterP0IsNotAskedForAScope(t *testing.T) {
	f := newFixture(t)
	f.templated()
	f.run(model.Phases[0], "")
	f.must(f.r.Start(key, model.Phases[1]))
	f.output(model.Phases[1], "")

	// P1 has no scope in its own directory and never will: informationBase reads P0's, so
	// there is one budget per intent rather than six.
	if fm.Exists(filepath.Join(f.root, model.PhaseDir(key, model.Phases[1]), model.ContextScope)) {
		t.Fatal("the fixture gave P1 a scope of its own")
	}
	if _, err := f.r.Finish(key, model.Phases[1], "a summary"); err != nil {
		t.Fatalf("a later phase was asked for a scope: %s", err)
	}
}
