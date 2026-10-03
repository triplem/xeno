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
---
`ExpectedDigest` is empty by default with section 13's reason beside it; the gate table's entry
becomes `supply`, three branches, with a doc comment carrying what it must not read — not the lock's
`plugin.sha256`, which records what a phase was given, and no version, because section 13 settles
the downgrade through the digest. `internal/plugin`'s tests become external, with the cycle written
into the package comment. `scripts/plugin-digest.go` is new, `//go:build ignore` as
`scripts/go-symbols.go` already is, printing the digest and failing loudly where there is no plugin
rather than compiling in an empty anchor. The release computes it with that script and echoes it, so
a release records what it anchored to. Seven tests with two helpers — one setting and restoring the
linker variable, because there is no `t.Setenv` for one, and one writing a tree and returning its
digest so a test anchors to what it wrote rather than to this repository's tree. Everything measured
with a binary built the way the release builds one: a match passes, an appended comment fails naming
both digests, the plugin moved away fails saying `init --vendor`, the dev build reports
`not-implemented`, and the released binary verifies this whole trail at exit 0. Three deviations: the
ldflags line split because one line with both `-X` paths runs to 101 columns; `ASSUMPTIONS.md`'s
stale list corrected for G-Rules and G-Policy as well, since a list wrong about three things and
fixed for one is trusted less; and two sealed verdicts rewritten and restored during the work, the
same mistake twice, caught both times by `gate verify` because a modification is loud.
