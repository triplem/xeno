// SPDX-License-Identifier: Apache-2.0

// Package plugin reads the vendored plugin: where it is, what version it declares, and
// the hash over the tree. Section 5 puts the first two in every artifact's header and
// the lock, and says what they are for — "the frontmatter names what was used,
// context.lock.yaml proves it with a hash. Where the two disagree, the hash wins and
// G-Supply fails."
//
// The directory is the vendored one and nothing else, which is now what section 7 says.
// It described a resolution order above this — a --plugin-root argument, then
// XENO_PLUGIN_ROOT, then here, then a client's own variable — and that order is gone
// rather than unimplemented: internal/gates reads internal/rules and internal/template,
// both of which resolve from this directory, so a root taken from the environment would
// make rules_hash, strings_hash and a rendered artifact depend on it. Measured on this
// repository, a valid rule tree that differs leaves `gate verify` at exit 0 over 273
// verdicts while G-Policy silently stops judging all 75 phases that recorded the
// previous hash — A74 judges a phase only against the set its own artifact names, so a
// changed set does not disagree with the trail, it stops judging it (#183).
package plugin

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/triplem/xeno/internal/hashing"
)

// ExpectedDigest is the digest of the plugin this binary was released with, compiled in
// by the release through -ldflags, and empty in a build that was not made by one.
//
// It is the anchor section 13 describes, and the reason it is here rather than in the
// repository is that section's own: "Nothing in the repository states what the expected
// value is, which is the point: an expected hash stored beside the thing it describes
// proves only that both were written by the same hand."
//
// Empty means this build carries no anchor, and G-Supply then reports not-implemented —
// the state section 5 defines for a check a runner did not perform, written so that the
// gap surfaces instead of a green verdict meaning less than it appears to. A
// development build of this repository is in that state, which is every artifact in
// this trail.
var ExpectedDigest = ""

// Dir is the vendored plugin, relative to the repository root. The same directory
// section 13 lists at the distribution root, which is where templates/,
// rules/given/builtin/ and secrets.yaml already are.
const Dir = ".xeno/plugin"

// manifestPath is where the client reads a plugin's manifest from, which A77 settled
// against section 13's tree after checking the installed client:
// <root>/.claude-plugin/plugin.json.
const manifestPath = ".claude-plugin/plugin.json"

// Version is what the vendored plugin declares about itself, or empty where there is no
// plugin or no manifest.
//
// Empty rather than a default, which is A35's rule and the one this field most needed:
// it was a constant, `model.PluginVersion`, whose comment said there was no plugin yet
// to have a version — and a release set it to the runner's own through ldflags, so a
// 0.28.0 runner working in a project whose plugin was 0.26.0 recorded 0.28.0. A number
// that cannot disagree with the runner can never be proved wrong, which is the one
// thing section 5 wants it for.
func Version(root string) string {
	var m struct {
		Version string `json:"version"`
	}
	b, err := os.ReadFile(filepath.Join(root, Dir, manifestPath))
	if err != nil || json.Unmarshal(b, &m) != nil {
		return ""
	}
	return m.Version
}

// Hash is the sha256 over the vendored tree, or empty where there is no plugin. It is
// the "proves it with a hash" half of section 5's sentence and what G-Supply compares
// when that gate exists; until then the lock records it and nothing reads it back.
//
// Appendix B does not define it. It defines artifacts_hash, context_hash and
// strings_hash, and says of the other two that "secrets_hash and rules_hash are defined
// by the package that first writes them" — so this is that, for this one, and it is
// written here to the byte because the appendix's reason applies unchanged: a hash
// described rather than defined is how two implementations end up one byte apart.
//
// **The definition.** Every file under the plugin directory, at any depth, with no
// exclusions. Each contributes one line, "<sha256 of its normalised content>  <path
// relative to the repository root, slash separated>\n", and the lines are sorted by
// path in byte order. The hash is the sha256 of that stream. It is Appendix B's
// artifacts_hash computation with the descent that one deliberately omits, because a
// phase directory's subdirectory is evidence and a plugin's subdirectories are the
// plugin.
//
// Symbolic links are not followed and directories contribute nothing of their own, so
// an empty directory is invisible to it — which is what git records too, and the tree
// this hashes is one git carries.
func Hash(root string) string {
	dir := filepath.Join(root, Dir)
	if _, err := os.Stat(dir); err != nil {
		return ""
	}
	type line struct{ path, sum string }
	var lines []line
	err := filepath.WalkDir(dir, func(abs string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !d.Type().IsRegular() {
			return nil
		}
		rel, err := filepath.Rel(root, abs)
		if err != nil {
			return err
		}
		b, err := os.ReadFile(abs)
		if err != nil {
			return err
		}
		lines = append(lines, line{filepath.ToSlash(rel), hashing.Hex(hashing.Normalise(b))})
		return nil
	})
	if err != nil {
		return ""
	}
	sort.Slice(lines, func(i, j int) bool { return lines[i].path < lines[j].path })
	var stream strings.Builder
	for _, l := range lines {
		fmt.Fprintf(&stream, "%s  %s\n", l.sum, l.path)
	}
	return hashing.Hex([]byte(stream.String()))
}
