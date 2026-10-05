// SPDX-License-Identifier: Apache-2.0

package plugin_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/plugin"
)

// The digest is taken over the tree at plugin.Dir and the binaries carry the tree at
// plugin.ShippedDir. Before #201 the hash keyed each line on the path relative to the
// repository root, so the same bytes at those two places hashed differently and the release
// compared them with nothing at all. These cover the property that replaced that: a tree
// hashes to what it is rather than to where it sits.

// tree writes a small tree of files given as path to content, relative to dir.
func tree(t *testing.T, dir string, files map[string]string) string {
	t.Helper()
	for rel, content := range files {
		abs := filepath.Join(dir, rel)
		if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(abs, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

// files is the shape of a plugin as far as the hash is concerned: nested, and carrying a
// dot-prefixed directory, because the manifest lives under .claude-plugin/ and a walk that
// skipped dot names would hash a different tree than the one shipped.
var files = map[string]string{
	".claude-plugin/plugin.json": `{"name":"xeno","version":"9.9.9"}`,
	"templates/intake.md":        "# Intake\n",
	"rules/given/builtin/a.yaml": "id: a\nversion: 1\n",
}

// TestATreeHashesTheSameAtTwoPaths is the property the release's check rests on, and the one
// the old definition made impossible.
func TestATreeHashesTheSameAtTwoPaths(t *testing.T) {
	root := t.TempDir()
	a := tree(t, filepath.Join(root, "vendored"), files)
	b := tree(t, filepath.Join(root, "somewhere/else/entirely"), files)

	ha, hb := plugin.HashTree(a), plugin.HashTree(b)
	if ha == "" {
		t.Fatal("a tree that is there hashed to nothing")
	}
	if ha != hb {
		t.Fatalf("the same tree at two paths hashed differently:\n  %s is %s\n  %s is %s", a, ha, b, hb)
	}
}

func TestATreeThatDiffersHashesDifferently(t *testing.T) {
	root := t.TempDir()
	a := tree(t, filepath.Join(root, "a"), files)

	changed := map[string]string{}
	for k, v := range files {
		changed[k] = v
	}
	changed["templates/intake.md"] = "# Intake\nand one more line\n"
	b := tree(t, filepath.Join(root, "b"), changed)

	if plugin.HashTree(a) == plugin.HashTree(b) {
		t.Fatal("trees differing in a file's contents hashed the same")
	}

	// And a file present in one and not the other, which is the stale-copy case #201 names.
	c := tree(t, filepath.Join(root, "c"), files)
	if err := os.WriteFile(filepath.Join(c, "STALE.txt"), []byte("left over\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if plugin.HashTree(a) == plugin.HashTree(c) {
		t.Fatal("a tree with an extra file hashed the same as one without it")
	}
}

// Hash stays the vendored tree's hash, so G-Supply and xeno init are unaffected by the split.
func TestHashIsTheTreeAtTheVendoredPath(t *testing.T) {
	root := t.TempDir()
	tree(t, filepath.Join(root, plugin.Dir), files)

	if got, want := plugin.Hash(root), plugin.HashTree(filepath.Join(root, plugin.Dir)); got != want {
		t.Fatalf("Hash is %s and the tree at %s is %s", got, plugin.Dir, want)
	}
}

// Empty rather than a hash of nothing, which is what G-Supply reads as a tree it cannot judge.
// A hash of nothing would compare unequal instead, and that is a verdict where there was none.
func TestATreeThatIsNotThereHashesToEmpty(t *testing.T) {
	if got := plugin.HashTree(filepath.Join(t.TempDir(), "absent")); got != "" {
		t.Fatalf("an absent tree hashed to %q", got)
	}
	if got := plugin.Hash(t.TempDir()); got != "" {
		t.Fatalf("a repository with no vendored plugin hashed to %q", got)
	}
}

// Differences decides nothing — the hashes do — and exists so a refusal can be acted on. #201
// ends with an adopter holding two hex strings and no way forward; a fix whose own failure said
// only "the digests differ" would have reproduced that one level up.
func TestDifferencesNamesWhatTheTreesDisagreeAbout(t *testing.T) {
	root := t.TempDir()
	a := tree(t, filepath.Join(root, "a"), files)

	other := map[string]string{}
	for k, v := range files {
		other[k] = v
	}
	other["templates/intake.md"] = "# Intake\ndifferent\n"
	delete(other, "rules/given/builtin/a.yaml")
	b := tree(t, filepath.Join(root, "b"), other)
	if err := os.WriteFile(filepath.Join(b, "STALE.txt"), []byte("left over\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	got := strings.Join(plugin.Differences(a, b), "\n")
	for _, want := range []string{
		"contents differ: templates/intake.md",
		"only in " + a + ": rules/given/builtin/a.yaml",
		"only in " + b + ": STALE.txt",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("the differences do not name %q:\n%s", want, got)
		}
	}
}

func TestDifferencesIsEmptyWhereTheTreesAgree(t *testing.T) {
	root := t.TempDir()
	a := tree(t, filepath.Join(root, "a"), files)
	b := tree(t, filepath.Join(root, "b"), files)

	if got := plugin.Differences(a, b); len(got) != 0 {
		t.Fatalf("two identical trees differ by %v", got)
	}
}

// An absent tree is the release's own failure mode — the copy did not happen — so it is said
// plainly rather than reported as every file being missing.
func TestDifferencesSaysWhichTreeIsNotThere(t *testing.T) {
	root := t.TempDir()
	a := tree(t, filepath.Join(root, "a"), files)
	absent := filepath.Join(root, "absent")

	if got := plugin.Differences(a, absent); len(got) != 1 || !strings.Contains(got[0], "is not there") {
		t.Fatalf("want one line saying the tree is not there, got %v", got)
	}
	if got := plugin.Differences(absent, a); len(got) != 1 || !strings.Contains(got[0], "is not there") {
		t.Fatalf("want one line saying the tree is not there, got %v", got)
	}
	if got := plugin.Differences(absent, absent); len(got) != 1 || !strings.Contains(got[0], "neither") {
		t.Fatalf("want one line saying neither is a tree, got %v", got)
	}
}
