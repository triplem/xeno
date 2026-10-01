// SPDX-License-Identifier: Apache-2.0

package index

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var now = time.Date(2026, 10, 1, 12, 0, 0, 0, time.UTC)

func write(t *testing.T, body string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "symbols.yaml")
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

const fresh = `tool: go-symbols
tool_version: 0.1.0
produced_at: "2026-10-01T11:00:00Z"
symbols:
  - name: Compare
    kind: func
    file: internal/enforcement/enforcement.go
    line: 144
  - name: Requirements
    kind: method
    file: internal/host/gitlab/gitlab.go
    line: 93
    container: Adapter
  - name: Requirements
    kind: field
    file: internal/enforcement/enforcement.go
    line: 49
    container: Report
`

// The property everything else follows from: an index may be missing, stale or wrong, so
// nothing about reading one is a failure. Four causes, one outcome, and a reason for each —
// because an error return would invite a caller to treat the ordinary state of every
// repository as a problem.
func TestNothingAboutReadingAnIndexIsAnError(t *testing.T) {
	for _, tc := range []struct {
		name string
		path string
		says string
	}{
		{name: "no path configured", path: "", says: "index.path"},
		{name: "configured and absent", path: filepath.Join(t.TempDir(), "nope.yaml"), says: "the project produces it"},
		{name: "not an index at all", path: write(t, "this: [is not\n"), says: "not a symbol index"},
		{name: "no provenance", path: write(t, "symbols:\n  - name: X\n    kind: func\n"), says: "age cannot be judged"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			i, why := Load(tc.path, DefaultMaxAge, now)
			if i != nil {
				t.Error("an index was returned")
			}
			if !strings.Contains(why, tc.says) {
				t.Errorf("the reason is %q, want it to mention %q", why, tc.says)
			}
		})
	}
}

// A stale index is worse than none, which is why its age is part of every answer. The
// boundary is tested from both sides, which is why Load takes the time rather than reading
// the clock.
func TestAStaleIndexIsTreatedAsAbsent(t *testing.T) {
	p := write(t, fresh) // produced at 11:00, read at 12:00

	if i, why := Load(p, 2*time.Hour, now); i == nil {
		t.Fatalf("an hour old index inside a two hour window was rejected: %s", why)
	}

	i, why := Load(p, 30*time.Minute, now)
	if i != nil {
		t.Error("an hour old index was accepted inside a thirty minute window")
	}
	if !strings.Contains(why, "stale") {
		t.Errorf("the reason is %q, want it to say the index is stale", why)
	}
}

func TestMaxAgeFallsBackToTheDefault(t *testing.T) {
	p := write(t, fresh)
	if i, _ := Load(p, 0, now); i == nil {
		t.Error("a zero max age rejected an index an hour old, so the default did not apply")
	}
	if DefaultMaxAge != 24*time.Hour {
		t.Errorf("the default is %s; section 5 says 24 hours", DefaultMaxAge)
	}
}

func TestLoadReturnsTheProvenance(t *testing.T) {
	i, why := Load(write(t, fresh), DefaultMaxAge, now)
	if i == nil {
		t.Fatal(why)
	}
	if i.Tool != "go-symbols" || i.ToolVersion != "0.1.0" {
		t.Errorf("provenance is %q %q", i.Tool, i.ToolVersion)
	}
	if got := i.Age(now); got != time.Hour {
		t.Errorf("age is %s, want an hour", got)
	}
}

// Where is X, answered with every place X is defined. The case that matters is a name with
// several definitions of different kinds, which is the one a search over the repository
// answers worst.
func TestLookupReturnsEveryLocationForAName(t *testing.T) {
	i, why := Load(write(t, fresh), DefaultMaxAge, now)
	if i == nil {
		t.Fatal(why)
	}
	got := i.Lookup("Requirements")
	if len(got) != 2 {
		t.Fatalf("%d locations for Requirements, want 2", len(got))
	}
	kinds := map[string]string{}
	for _, s := range got {
		kinds[s.Kind] = s.Container
	}
	if kinds["method"] != "Adapter" || kinds["field"] != "Report" {
		t.Errorf("the kinds and containers did not survive: %v", kinds)
	}
	if got := i.Lookup("Compare"); len(got) != 1 || got[0].Line != 144 {
		t.Errorf("Compare resolved to %v", got)
	}
}

// An index is allowed to be incomplete, so a name it does not hold is an answer. And a
// caller with no index needs no branch, which is what the nil receiver is for.
func TestAnAbsentNameAndAnAbsentIndexAreBothAnswers(t *testing.T) {
	i, _ := Load(write(t, fresh), DefaultMaxAge, now)
	if got := i.Lookup("NoSuchSymbol"); got != nil {
		t.Errorf("an absent name returned %v", got)
	}
	var none *Index
	if got := none.Lookup("Compare"); got != nil {
		t.Errorf("a nil index returned %v", got)
	}
	if got := none.Age(now); got != 0 {
		t.Errorf("a nil index has age %s", got)
	}
}

// The worked example is the format's documentation, so a test reads it: a description a
// project writes against which nothing checks is a description that drifts from the reader
// within a release.
func TestTheWorkedExampleIsWhatTheReaderAccepts(t *testing.T) {
	at := time.Date(2026, 10, 1, 12, 0, 0, 0, time.UTC)
	i, why := Load("testdata/example-symbols.yaml", DefaultMaxAge, at)
	if i == nil {
		t.Fatalf("the documented example is not readable: %s", why)
	}
	if i.Tool == "" || i.ToolVersion == "" || i.ProducedAt.IsZero() {
		t.Error("the example does not show all three provenance fields")
	}
	// It has to show the case the index exists for: one name, several definitions.
	if got := i.Lookup("Requirements"); len(got) < 2 {
		t.Errorf("the example does not show a name with several definitions: %v", got)
	}
	// And a symbol with no container, since that is a value and not an omission to guess at.
	top := false
	for _, s := range i.Symbols {
		if s.Container == "" {
			top = true
		}
	}
	if !top {
		t.Error("the example does not show a top level symbol")
	}
}

// The format has to work against a tool's output and not only against a fixture written by
// whoever wrote the reader. This produces an index for this repository with this project's
// own producer, reads it back and queries it, which is WP15's done-when about an index this
// repository produces for its own language.
func TestTheFormatRoundTripsThroughThisProjectsProducer(t *testing.T) {
	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("no go toolchain")
	}
	root := filepath.Join("..", "..")
	out := filepath.Join(t.TempDir(), "symbols.yaml")
	f, err := os.Create(out)
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", filepath.Join("scripts", "go-symbols.go"))
	cmd.Dir = root
	cmd.Stdout = f
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		f.Close()
		t.Fatalf("the producer failed: %v", err)
	}
	f.Close()

	i, why := Load(out, DefaultMaxAge, time.Now())
	if i == nil {
		t.Fatalf("the producer's own output could not be read: %s", why)
	}
	if i.Tool != "go-symbols" {
		t.Errorf("tool is %q", i.Tool)
	}
	if len(i.Symbols) < 100 {
		t.Errorf("%d symbols from this repository, which is too few to be right", len(i.Symbols))
	}

	// Lookup has to find itself, in this file's own package, with its own container.
	got := i.Lookup("Lookup")
	found := false
	for _, s := range got {
		if s.Container == "Index" && strings.HasSuffix(s.File, "index.go") {
			found = true
		}
	}
	if !found {
		t.Errorf("the index of this repository cannot find Lookup on Index: %v", got)
	}
}
