---
intent: github.com/triplem/xeno#183
phase: 00-intake
created: "2026-10-03T17:30:24Z"
schema_version: "1.0"
runner_version: dev+9140756.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 411b3f7c1e2d855f01d3eb145085cdd16425870fb556c1b2c945145dea399b6e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

G-Supply is the first gate section 7's table lists and the reason it gives for the order: it "runs
ahead of everything because a plugin that does not match its expected digest makes every later
verdict a statement about unknown rules and unknown templates." It was written as
`not-implemented`, so every verdict in this repository is that statement.

**It is the gate the last two intents kept arriving at.** #177 found `plugin_version` to be a
constant a release overwrote with the runner's own number, and section 5's reason for the field —
"the frontmatter names what was used, `context.lock.yaml` proves it with a hash. Where the two
disagree, the hash wins and G-Supply fails" — names this gate as the thing that would act on it.
#183 found that section 7's resolution order could not be built, and the reason was that the three
defences against a swapped plugin were absent at once: the field was a constant, the lock carried no
`plugin` block, and this gate did not exist. Two of the three were closed and this is the third.

**What the specification asks for is unusually specific, and it is not what the lock does.** Section
13: "The anchor is the runner binary. Plugin and runner are released together under one version, so
the runner carries the digest of its own plugin compiled in. G-Supply recomputes the digest over
`.xeno/plugin/` and compares. Nothing in the repository states what the expected value is, which is
the point: an expected hash stored beside the thing it describes proves only that both were written
by the same hand."

So the lock's `plugin.sha256`, added by #177, is not the anchor and must not be read as one. It
records what a phase was given; the anchor records what this runner expects, and the whole value of
the arrangement is that the checked side cannot influence it.

**And the downgrade needs no version check.** Section 13 again: "A vendored plugin from an earlier
release has a different digest than the one this runner expects, so it fails without any separate
version check." The gate therefore reads no version at all, which is what keeps it from becoming a
compatibility matrix — section 13 says G-Supply "fails on any version difference, not only a major
one".

**The case the specification does not name is a build that carries no digest.** Every artifact in
this repository was written by `go build`, which compiles in nothing, and a gate with no anchor has
nothing to compare. Reporting a pass would be the exact failure section 5 defines `not-implemented`
to prevent: "A gate that is skipped silently makes a green verdict mean less than it appears to, and
the gap never surfaces afterwards."

<!-- xeno:section:scope -->
## Scope

**In scope.** `G-Supply` implemented: recompute the digest over the vendored tree and compare
against one the binary carries. `not-implemented` where the binary carries none, which is a
development build and so is this repository. A finding naming both digests where they differ, and a
finding naming the absence where there is no plugin and the runner expects one. The release
computing the digest and compiling it in, with the same code the gate compares with. And the records
that said this gate was unimplemented corrected — in `ASSUMPTIONS.md`'s two status paragraphs and in
the README, where a reader of this trail will meet `not-implemented` and needs to know it is the
gate working.

**Out of scope, and each for its own reason.**

The resolution order, `--plugin-root` and `XENO_PLUGIN_ROOT`. This gate is what the maintainer's
recommendation put ahead of it, and building both at once would mean building the containment and
the thing it contains in one change, with no run in between where the containment was the only thing
tested. #183 keeps the order and the decision is still section 7's.

Embedding the plugin in the runner. `xeno init --vendor` copies from `--plugin-from`, which defaults
to the repository's own tree, and A58 already records that `init`'s scaffold files are embedded
while the plugin is not. A released binary that cannot vendor the plugin it carries the digest of is
a real gap and it is a different one: this gate judges a plugin that is there.

Comparing the frontmatter's `plugin_version` against the lock's, which is the other half of section
5's sentence. The lock now records both halves and nothing reads them back; making this gate read
them would make the checked side's own record part of the anchor, which section 13 excludes in as
many words.

G-Secret, the other gate the plan puts on a harness. It waits on the hook that runs on every write.

<!-- xeno:section:context-rationale -->
## Why this context

Section 13's **Integrity** and **Version coupling** paragraphs are the specification for this intent
and are read whole. They give the anchor, the reason it is not in the repository, the recomputation
over `.xeno/plugin/`, the downgrade following from the digest alone, and the limit — "This is not
tamper protection and does not pretend to be... What G-Supply gives is that a changed shipped set
cannot pass unnoticed, and for that a digest the checked side cannot influence is enough."

Section 7's gate table and the paragraph above it are read for the position and the reason: first,
because a mismatch makes every later verdict a statement about unknown rules. Also for `From P0`.

Section 5's `not-implemented` paragraph is read for what that state means and what it is for, which
is what decides the no-anchor case: "The state is written, it is not a failure, and the derived
status treats it as neither pass nor fail: a phase with unimplemented gates is green in what was
checked and says which checks were not."

Section 5's sentence about the two version fields is read to establish what this gate does *not* do
with the lock's new block.

A74 is read once more, for the shape of the failure the measurement in #183 produced, because it is
the reason this gate is worth building before the override that would need it.

`internal/gates/gates.go` is read for the table, `Applicable`, `notImplemented`, `result` and
`finding`, and for the import graph: `internal/plugin`'s tests import `internal/gates`, so a gate
importing the plugin package makes the internal test a cycle and the tests have to become an
external package.

`internal/plugin/plugin.go` is read for `Hash`, written for #177 with the definition in its doc
comment, which is what this gate recomputes.

`.github/workflows/release.yml` is read for the ldflags line and for where a digest would be
computed before it.

`scripts/go-symbols.go` is read for the `//go:build ignore` pattern this repository already uses for
a script run with `go run`.
