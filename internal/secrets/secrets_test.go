// SPDX-License-Identifier: Apache-2.0

package secrets

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func write(t *testing.T, root, rel, body string) {
	t.Helper()
	p := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

const shipped = `patterns:
  - id: aws-access-key
    regex: '\b(?:AKIA|ASIA)[0-9A-Z]{16}\b'
  - id: github-token
    regex: '\bgh[posru]_[A-Za-z0-9]{36,}\b'
paths_never_digested:
  - "**/*.pem"
`

func load(t *testing.T, root string) Filter {
	t.Helper()
	f, err := Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return f
}

// Section 4: a project adds and removes nothing, because a filter a project can switch off is
// not a filter. Matching is a disjunction, so a project pattern reusing a shipped id leaves the
// shipped one in force rather than replacing it.
func TestAProjectAddsAndRemovesNothing(t *testing.T) {
	root := t.TempDir()
	write(t, root, Shipped, shipped)
	write(t, root, Project, "patterns:\n  - id: aws-access-key\n    regex: 'never-matches-anything'\n"+
		"  - id: house-token\n    regex: '\\bHOUSE-[0-9]{4}\\b'\n")

	f := load(t, root)
	if len(f.Patterns) != 4 {
		t.Fatalf("the effective filter has %d patterns, want 4: %v", len(f.Patterns), f.Patterns)
	}
	// The shipped pattern still fires although the project reused its id.
	if got := f.Redact("key AKIAIOSFODNN7EXAMPLE here"); !strings.Contains(got, "[redacted: aws-access-key]") {
		t.Errorf("a project reusing a shipped id switched it off: %q", got)
	}
	if got := f.Redact("ticket HOUSE-1234 here"); !strings.Contains(got, "[redacted: house-token]") {
		t.Errorf("the project's own pattern did not fire: %q", got)
	}
}

// The hash covers the effective set and not the bytes of the files: a comment must not move it.
func TestACommentDoesNotChangeTheHash(t *testing.T) {
	a, b := t.TempDir(), t.TempDir()
	write(t, a, Shipped, shipped)
	write(t, b, Shipped, "# a comment nobody should be able to feel\n"+shipped+"\n# and a trailing one\n")

	if load(t, a).Hash() != load(t, b).Hash() {
		t.Fatal("editing a comment changed which filter a digest says it passed through")
	}
}

// Nor the order the patterns were written in, nor which of the two files they came from.
func TestTheOrderAndTheFileDoNotChangeTheHash(t *testing.T) {
	a, b := t.TempDir(), t.TempDir()
	write(t, a, Shipped, shipped)
	// The same effective set, written in the other order and split across both files.
	write(t, b, Shipped, "patterns:\n  - id: github-token\n    regex: '\\bgh[posru]_[A-Za-z0-9]{36,}\\b'\n"+
		"paths_never_digested:\n  - \"**/*.pem\"\n")
	write(t, b, Project, "patterns:\n  - id: aws-access-key\n    regex: '\\b(?:AKIA|ASIA)[0-9A-Z]{16}\\b'\n")

	if load(t, a).Hash() != load(t, b).Hash() {
		t.Fatalf("the same effective filter hashed differently:\n%s\n%s",
			load(t, a).Hash(), load(t, b).Hash())
	}
}

// A pattern added anywhere changes it, or the field would not say which filter was in force.
func TestAddingAPatternChangesTheHash(t *testing.T) {
	a, b := t.TempDir(), t.TempDir()
	write(t, a, Shipped, shipped)
	write(t, b, Shipped, shipped)
	write(t, b, Project, "patterns:\n  - id: house-token\n    regex: 'HOUSE'\n")

	if load(t, a).Hash() == load(t, b).Hash() {
		t.Fatal("a pattern was added and the hash did not move")
	}
}

// A path is part of the effective set, so it is part of the hash even though nothing reads it yet.
func TestAddingAPathChangesTheHash(t *testing.T) {
	a, b := t.TempDir(), t.TempDir()
	write(t, a, Shipped, shipped)
	write(t, b, Shipped, shipped+"  - \"**/id_rsa\"\n")

	if load(t, a).Hash() == load(t, b).Hash() {
		t.Fatal("a never digested path was added and the hash did not move")
	}
}

// The state of a repository before its plugin is vendored: no filter, no hash, nothing redacted,
// and no error. The field is then absent rather than a hash over an empty set.
func TestNoFilterIsEmptyRatherThanAnError(t *testing.T) {
	f := load(t, t.TempDir())
	if !f.Empty() {
		t.Fatal("a repository with no filter produced one")
	}
	if f.Hash() != "" {
		t.Errorf("an empty filter has the hash %q, want none", f.Hash())
	}
	const text = "key AKIAIOSFODNN7EXAMPLE stays as it is"
	if got := f.Redact(text); got != text {
		t.Errorf("an empty filter redacted something: %q", got)
	}
}

// The redaction names the pattern and leaves the line, because a digest is read by a person and
// a silent deletion would make the sentence around it a lie.
func TestRedactionNamesThePatternAndKeepsTheLine(t *testing.T) {
	root := t.TempDir()
	write(t, root, Shipped, shipped)

	got := load(t, root).Redact("the run used AKIAIOSFODNN7EXAMPLE against the bucket")
	if strings.Contains(got, "AKIAIOSFODNN7EXAMPLE") {
		t.Fatalf("the secret survived: %q", got)
	}
	if !strings.Contains(got, "[redacted: aws-access-key]") {
		t.Errorf("the redaction does not name the pattern: %q", got)
	}
	for _, keep := range []string{"the run used", "against the bucket"} {
		if !strings.Contains(got, keep) {
			t.Errorf("the line around the match was destroyed: %q", got)
		}
	}
}

// A project file with a regex that will not compile drops that pattern rather than stopping every
// phase in the repository, and the pattern still enters the hash so what was in force is
// reconstructable.
func TestABrokenRegexDropsItsPatternAndNotTheFilter(t *testing.T) {
	root := t.TempDir()
	write(t, root, Shipped, shipped)
	write(t, root, Project, "patterns:\n  - id: broken\n    regex: '([unclosed'\n")

	f := load(t, root)
	if f.Hash() == "" {
		t.Fatal("one bad pattern emptied the filter")
	}
	if got := f.Redact("key AKIAIOSFODNN7EXAMPLE here"); !strings.Contains(got, "[redacted: aws-access-key]") {
		t.Errorf("one bad pattern stopped the others firing: %q", got)
	}
	var named bool
	for _, p := range f.Patterns {
		if p.ID == "broken" {
			named = true
		}
	}
	if !named {
		t.Error("the broken pattern left the effective set, so the hash no longer says what was in force")
	}
}

// The shipped file this repository carries, against the shapes it has to catch.
//
// Asserted by id rather than by "something fired", because a rule set imported from upstream
// is not a superset of what was here before: private keys, Slack prefixes and Authorization
// values are caught by patterns of this project's own, since upstream is tuned to find a whole
// credential in a file and a digest carries prose where one is quoted or truncated. Losing
// those three silently is what this test exists to prevent.
func TestTheShippedFilterCatchesWhatItNames(t *testing.T) {
	f, err := Load("../..")
	if err != nil {
		t.Fatal(err)
	}
	if f.Empty() {
		t.Fatal("this repository ships no filter")
	}
	for _, tc := range []struct{ id, sample string }{
		{"aws-access-token", "AKIAIOSFODNN7EXAMPLE"},
		{"github-pat", "ghp_" + strings.Repeat("a", 36)},
		{"house-private-key-header", "-----BEGIN OPENSSH PRIVATE KEY-----"},
		{"house-slack-token-prefix", "xoxb-" + strings.Repeat("1", 12)},
		{"house-authorization-value", "Authorization: " + strings.Repeat("A", 24)},
	} {
		got := f.Redact("before " + tc.sample + " after")
		if !strings.Contains(got, "[redacted: "+tc.id+"]") {
			t.Errorf("%s did not catch %q: %q", tc.id, tc.sample, got)
		}
	}
	// Ordinary prose passes through untouched, or a reader learns to distrust the redaction.
	const prose = "The runner filters the summary against the effective filter and writes secrets_hash."
	if got := f.Redact(prose); got != prose {
		t.Errorf("the shipped filter redacted prose: %q", got)
	}
}

// The import is a selection, so the count is part of what shipped: a set that quietly shrank to
// a handful would pass every test above.
func TestTheShippedSetIsTheImportedOne(t *testing.T) {
	f, err := Load("../..")
	if err != nil {
		t.Fatal(err)
	}
	if len(f.Patterns) < 200 {
		t.Errorf("%d patterns in force, want the imported set of about two hundred", len(f.Patterns))
	}
	var house int
	for _, p := range f.Patterns {
		if strings.HasPrefix(p.ID, "house-") {
			house++
		}
	}
	if house != 3 {
		t.Errorf("%d house patterns, want the three upstream does not cover", house)
	}
}

// A match reports the pattern and the line and nothing of the value. The gate writes this into
// `gate.yaml`, which is committed, so a value carried out here would be the secret published a
// second time in the file nobody thinks of as holding it.
func TestAMatchReportsThePatternAndTheLineAndNotTheValue(t *testing.T) {
	root := t.TempDir()
	write(t, root, Shipped, shipped)

	hits := load(t, root).Matches("one\ntwo\nthe run used AKIAIOSFODNN7EXAMPLE here\nfour\n")
	if len(hits) != 1 {
		t.Fatalf("want one hit, got %+v", hits)
	}
	if hits[0] != (Hit{PatternID: "aws-access-key", Line: 3}) {
		t.Errorf("got %+v, want the aws pattern on line 3", hits[0])
	}
}

// Two patterns on one line are two hits, and one pattern firing twice on one line is one: a
// finding is a place to look, and one pattern naming one line twice would be two findings a
// person has to decide separately about the same line.
func TestOnePatternPerLineAndEveryPatternThatFires(t *testing.T) {
	root := t.TempDir()
	write(t, root, Shipped, shipped)
	text := "a AKIAIOSFODNN7EXAMPLE and AKIAIOSFODNN7EXAMPL2 and ghp_" + strings.Repeat("a", 36) + "\n"

	hits := load(t, root).Matches(text)
	if len(hits) != 2 {
		t.Fatalf("want one hit per pattern on the line, got %+v", hits)
	}
	if hits[0].PatternID != "aws-access-key" || hits[1].PatternID != "github-token" {
		t.Errorf("hits are not sorted by pattern within a line: %+v", hits)
	}
}

// The matcher and the redaction are the same compiled set, which is section 4's "G-Secret and
// the digest writer read the same file" seen one level further in. Anything the redaction
// replaces the matcher finds, and the subject is this repository's own shipped filter rather
// than a fixture, because the claim is about what ships.
func TestTheMatcherAndTheRedactionAgreeOnTheShippedFilter(t *testing.T) {
	f, err := Load("../..")
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range samples {
		text := "before " + s + " after"
		redacted := f.Redact(text) != text
		matched := len(f.Matches(text)) > 0
		if redacted != matched {
			t.Errorf("%q: the redaction says %v and the matcher says %v", s, redacted, matched)
		}
	}
}

// A filter rewritten under one root is compiled again. The compiled sets are memoised, because a
// verification loads the same two files once per verdict, and a memo keyed on anything but the
// content would hand back the set the files used to have.
func TestARewrittenFilterIsNotServedFromTheMemo(t *testing.T) {
	root := t.TempDir()
	write(t, root, Shipped, shipped)
	if hits := load(t, root).Matches("ticket HOUSE-1234"); len(hits) != 0 {
		t.Fatalf("the fixture already matches: %+v", hits)
	}
	write(t, root, Shipped, "patterns:\n  - id: house-token\n    regex: '\\bHOUSE-[0-9]{4}\\b'\n")

	if hits := load(t, root).Matches("ticket HOUSE-1234"); len(hits) != 1 {
		t.Errorf("the rewritten filter did not take effect: %+v", hits)
	}
}
