---
intent: github.com/triplem/xeno#183
phase: 03-implementation
created: "2026-10-03T17:32:23Z"
schema_version: "1.0"
runner_version: dev+9140756.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c312cb263e0e3668a8c972994051b72bed518e9d8ccd14e6579ca665546d5ff0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

**`internal/plugin/plugin.go`.** `ExpectedDigest`, empty by default, with section 13's sentence about
why the anchor is not in the repository and with the no-anchor case stated: a build carrying none
reports `not-implemented`, which is every artifact in this trail.

**`internal/gates/gates.go`.** `{"G-Supply", 0, supply}` where the table said `notImplemented`.
`supply` with three branches, and a doc comment that carries what the gate must not read — not the
lock's `plugin.sha256`, because it records what a phase was given rather than what this runner
expects, and no version, because section 13 settles the downgrade through the digest alone. An import
of `internal/plugin`.

**`internal/plugin/plugin_test.go`.** `package plugin_test`, with the cycle written into the package
comment: the tests read what a skill says a gate will refuse, which means importing
`internal/gates`, and that now imports `internal/plugin`.

**`scripts/plugin-digest.go`, new.** `//go:build ignore`, the pattern `scripts/go-symbols.go`
already uses. Prints `plugin.Hash(".")`, and exits non-zero with a message where there is no plugin
rather than printing an empty line that would compile in an empty anchor.

**`.github/workflows/release.yml`.** `DIGEST=$(go run scripts/plugin-digest.go)` with the section 13
quotation above it and the reason the release uses the gate's own code, echoed into the log so a
release records what it anchored to. The ldflags line is assembled in three steps rather than one,
because one line with both `-X` paths runs past the width this project wraps at.

**Tests, seven in `internal/gates/supply_test.go`.** `withExpected`, which sets and restores the
linker variable because there is no `t.Setenv` for one, and `vendor`, which writes a tree and returns
its digest so a test anchors to what it wrote. Then: the no-anchor case asserting both the result and
that it carries no findings; the match; a one-byte change asserting both digests appear in the cause
and that the file named is the directory; an earlier plugin, to show no version is read; no plugin
with an anchored runner; the anchor not being readable from four plausible places in the tree; and
the gate's position in `Applicable`.

**`ASSUMPTIONS.md`.** A86, and both status paragraphs corrected — G-Supply leaves the
`not-implemented` list, which also still named G-Rules and G-Policy wrongly, and the "Not built"
paragraph stops saying this gate waits on a harness.

**`README.md`.** Why a reader of this trail meets `not-implemented` on every phase and what a
released binary reports instead.

**Measured, with a binary built the way the release builds one.** A matching tree passes; one appended
comment fails and the finding names both digests; the plugin moved away fails and the finding says
`init --vendor`; the development build reports `not-implemented`; and the released binary verifies
this whole trail at exit 0.

<!-- xeno:section:deviations -->
## Deviations from the design

**One from the design.** The ldflags line is built in three assignments rather than one, with two
short local variables for the package paths, because a single line carrying both `-X` arguments runs
to 101 columns. The design said "a second `-X` on the ldflags line" and the line could not hold it.

**One widening the design did not ask for.** `ASSUMPTIONS.md`'s list of gates written as
`not-implemented` also named G-Rules and G-Policy, which have been implemented since the shipped rule
set landed. They are corrected in the same sentence, because a list that is wrong about three things
and fixed for one is a list a reader trusts less than before.

**Two sealed verdicts were rewritten during this work and restored.** A `gate run` against a sealed
phase, twice: once to see what a released binary reported, once from a mistyped key. `gate verify`
caught both as divergences and both were restored from git. Recorded because it is the same mistake
twice in one session, and because the reason it was visible is the subject of #193: a modification to
a sealed artifact is loud.
