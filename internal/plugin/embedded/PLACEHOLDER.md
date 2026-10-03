<!-- SPDX-License-Identifier: Apache-2.0 -->

# Not the plugin

This directory is where the release puts the plugin tree before it builds the binaries, so
that `xeno init --vendor` has something to copy in a project that has no plugin yet. A
development build carries only this file, and `init --vendor` then says so and asks for
`--plugin-from`.

The release writes the tree to `plugin/` beside this file, after stamping the version into
its manifest and before computing the digest G-Supply compares. The order matters: the
digest has to be taken over the bytes an adopter will receive, or every adopter fails the
gate.

It is not `.gitignore`d, because `go:embed` needs the directory to exist at compile time and
a build of this repository has to work from a clean checkout.
