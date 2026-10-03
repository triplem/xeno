// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/plugin"
)

// withExpected sets the digest this binary claims to carry and restores it afterwards. It is
// the anchor section 13 puts in the binary, so a test that wants one has to pretend to be a
// release; t.Setenv has no equivalent for a linker variable.
func withExpected(t *testing.T, digest string) {
	t.Helper()
	was := plugin.ExpectedDigest
	plugin.ExpectedDigest = digest
	t.Cleanup(func() { plugin.ExpectedDigest = was })
}

// vendor writes a plugin tree under root and returns its digest, so a test can anchor to what
// it just wrote rather than to a literal that would pin this repository's own tree.
func vendor(t *testing.T, root, content string) string {
	t.Helper()
	dir := filepath.Join(root, plugin.Dir, "rules", "given", "builtin")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "a-rule.yaml"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return plugin.Hash(root)
}

// A build that carries no digest reports not-implemented. Section 5 defines that state for a
// check a runner did not perform — "it is not a failure, and the derived status treats it as
// neither pass nor fail" — and a development build is in it, which is every artifact in this
// repository's own trail. Passing would make a green verdict mean less than it appears to.
func TestWithoutAnAnchorSupplyIsNotImplemented(t *testing.T) {
	root := t.TempDir()
	vendor(t, root, "id: a\n")
	withExpected(t, "")

	got := supply(ctxFor(root))
	if got.Result != "not-implemented" {
		t.Errorf("result %q, want not-implemented", got.Result)
	}
	if len(got.Findings) != 0 {
		t.Errorf("a check that did not run reports %d finding(s)", len(got.Findings))
	}
}

// The plugin the runner was released with passes, and the digest is recomputed over the tree
// rather than read from anywhere in it.
func TestTheVendoredPluginMatchingTheAnchorPasses(t *testing.T) {
	root := t.TempDir()
	withExpected(t, vendor(t, root, "id: a\n"))

	if got := supply(ctxFor(root)); got.Result != "pass" {
		t.Errorf("result %q with findings %+v, want pass", got.Result, got.Findings)
	}
}

// One byte is enough, which is what a digest over a tree means. Section 7 runs this gate ahead
// of everything "because a plugin that does not match its expected digest makes every later
// verdict a statement about unknown rules and unknown templates".
func TestAChangedPluginFails(t *testing.T) {
	root := t.TempDir()
	withExpected(t, vendor(t, root, "id: a\n"))
	vendor(t, root, "id: a\n# one byte more\n")

	got := supply(ctxFor(root))
	if got.Result != "fail" {
		t.Fatalf("result %q, want fail", got.Result)
	}
	if len(got.Findings) != 1 {
		t.Fatalf("want one finding, got %+v", got.Findings)
	}
	// Both digests, because a finding that says only "does not match" leaves the reader to
	// compute the two values this gate already has.
	if !strings.Contains(got.Findings[0].Cause, plugin.ExpectedDigest) {
		t.Errorf("the finding does not name the expected digest: %q", got.Findings[0].Cause)
	}
	if !strings.Contains(got.Findings[0].Cause, plugin.Hash(root)) {
		t.Errorf("the finding does not name the digest found: %q", got.Findings[0].Cause)
	}
	if got.Findings[0].File != plugin.Dir {
		t.Errorf("the finding names %q, want the plugin directory", got.Findings[0].File)
	}
}

// Section 13 settles the downgrade through the digest alone: "A vendored plugin from an earlier
// release has a different digest than the one this runner expects, so it fails without any
// separate version check." So an older plugin is the same finding as a changed one, and the
// gate reads no version to decide it.
func TestAnEarlierPluginFailsOnTheDigestAlone(t *testing.T) {
	root := t.TempDir()
	withExpected(t, vendor(t, root, "id: a\nversion: 2\n"))
	vendor(t, root, "id: a\nversion: 1\n")

	if got := supply(ctxFor(root)); got.Result != "fail" {
		t.Errorf("an earlier plugin gives %q, want fail", got.Result)
	}
}

// No plugin at all, with a runner that carries the digest of one, is a failure and not a
// not-implemented: the check ran and the answer is that what it was anchored to is absent.
func TestNoVendoredPluginFailsWhenTheRunnerExpectsOne(t *testing.T) {
	root := t.TempDir()
	withExpected(t, "c0ffee")

	got := supply(ctxFor(root))
	if got.Result != "fail" {
		t.Fatalf("result %q, want fail", got.Result)
	}
	if len(got.Findings) != 1 || !strings.Contains(got.Findings[0].Next, "init --vendor") {
		t.Errorf("the finding does not say how to put one there: %+v", got.Findings)
	}
}

// The gate compares against the binary's own value and against nothing in the tree. Section
// 13: "Nothing in the repository states what the expected value is, which is the point: an
// expected hash stored beside the thing it describes proves only that both were written by the
// same hand." A file in the repository naming the right digest must not rescue a wrong tree.
func TestTheAnchorIsNotReadFromTheRepository(t *testing.T) {
	root := t.TempDir()
	withExpected(t, vendor(t, root, "id: a\n"))
	right := plugin.ExpectedDigest
	vendor(t, root, "id: a\n# changed\n")
	// Every plausible place somebody might put it, including the lock's own field name.
	for _, rel := range []string{
		"expected-digest", ".xeno/expected-digest", plugin.Dir + "/digest",
		plugin.Dir + "/.claude-plugin/digest",
	} {
		p := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(right), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if got := supply(ctxFor(root)); got.Result != "fail" {
		t.Errorf("result %q, want fail: a digest in the tree must not be the anchor", got.Result)
	}
}

// G-Supply applies from P0, which section 7's table says and the order depends on: it runs
// ahead of everything because a mismatched plugin makes every later verdict a statement about
// unknown rules.
func TestSupplyAppliesFromTheFirstPhaseAndRunsFirst(t *testing.T) {
	ids := Applicable("00-intake")
	if len(ids) == 0 || ids[0] != "G-Supply" {
		t.Fatalf("G-Supply is not the first gate of P0: %v", ids)
	}
}
