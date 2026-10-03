// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"embed"
	"io/fs"
)

// embedded is where the release puts the plugin tree before building, so that a released
// binary can vendor the plugin it carries the digest of. A development build carries only
// the placeholder.
//
// `all:` because the plugin's manifest lives under `.claude-plugin/`, and embed skips names
// beginning with a dot without it.
//
//go:embed all:embedded
var embedded embed.FS

// shippedRoot is where the release writes the tree, beside the placeholder rather than over
// it, so that the presence of this one path is what distinguishes a release from a build.
const shippedRoot = "embedded/plugin"

// Shipped is the plugin this binary carries, and false where it carries none.
//
// A released binary carries one because `xeno init --vendor` has to work in a project that
// has no plugin yet, and because vendoring from anywhere else would produce a tree whose
// digest is not the one this runner expects — which G-Supply would then fail, for every
// adopter, on every phase. The copy is therefore the same bytes the digest was taken over.
//
// A development build carries none and says so rather than vendoring something plausible.
// That is the same answer `ExpectedDigest` gives in the same situation: this build is not a
// release and will not pretend to be one.
func Shipped() (fs.FS, bool) {
	sub, err := fs.Sub(embedded, shippedRoot)
	if err != nil {
		return nil, false
	}
	// fs.Sub succeeds on a path that is not there, so the manifest is what settles it: a
	// tree without one is not a plugin a client would load.
	if _, err := fs.Stat(sub, manifestPath); err != nil {
		return nil, false
	}
	return sub, true
}
