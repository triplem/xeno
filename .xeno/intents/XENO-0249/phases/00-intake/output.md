---
intent: github.com/triplem/xeno#201
phase: 00-intake
created: "2026-10-05T12:47:23Z"
schema_version: "1.0"
runner_version: dev+f2f65d5.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1a371a30a09d5386056df231bc930934ab318f051e9feebeedd0f7d863d0b8ea
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

The release writes the plugin into the binaries and takes the digest over a different path,
and nothing compares the two trees.

`.github/workflows/release.yml` does three things in order: it stamps the version into
`.xeno/plugin/.claude-plugin/plugin.json`, copies `.xeno/plugin/.` into
`internal/plugin/embedded/plugin/`, and then runs `scripts/plugin-digest.go`, which hashes
`.xeno/plugin`. So the digest describes the tree at one path and the binaries carry the tree
at the other. They are the same bytes because `cp -R` ran a moment earlier, and that is the
whole of the guarantee.

`internal/plugin/embedded.go` already states the property as though it were enforced: "the
copy is therefore the same bytes the digest was taken over". Nothing checks that sentence,
which makes it the same kind of claim A90 catalogues — a rule written down with no reader —
except that this one is load bearing for every adopter rather than for a reader.

What a wrong copy costs lands on somebody else entirely. `xeno init --vendor` gives an
adopter the embedded tree; G-Supply then hashes what they vendored and compares it against
the digest the binary carries. If the two trees differed at release time, every phase of
every project installing that release fails G-Supply, with a finding naming two hex strings
and nothing the adopter can do about either. The release that made the mistake passes its
own checks.

`cp -R` under `set -eu` catches a copy that fails. What it cannot catch is a copy that
succeeds and produces a different tree: a stale file left by an earlier run in the same
checkout, a path the hasher and `cp` treat differently, or a later refactor that changes one
of the two paths and not the other.

Underneath it is a property of the hash. `plugin.Hash` keys each line on the file's path
relative to the repository root, so the same bytes under `.xeno/plugin` and under
`internal/plugin/embedded/plugin` hash to different values by construction. The two trees
cannot be compared with the definition the gate uses, which is why the release compares them
with nothing at all.

<!-- xeno:section:scope -->
## Scope

In scope is making the hash a function of a tree rather than of a tree at one path.
`plugin.HashTree(dir)` hashes the tree rooted at `dir` with each line's path relative to
`dir`, and `plugin.Hash(root)` becomes that function applied to `root/.xeno/plugin`. The
same bytes at two locations then hash the same, which is what lets the two trees be compared
with the definition the gate uses rather than with a second one.

In scope is the comparison itself, in `scripts/plugin-digest.go`, which is the only caller
the release has. With `--require-shipped` it refuses to print a digest unless the embedded
tree exists and hashes equal to the vendored one, so the anchor cannot be produced from a
tree the binaries do not carry. The workflow passes the flag, beside the copy it guards.

In scope is saying which paths differ when they do. A refusal naming two hex strings is the
defect this intent exists to prevent one level up, and repeating it in the fix would be
careless: the script lists the paths that are in one tree and not the other, and those whose
contents differ.

In scope is a test over the location independence, which is the property everything else
rests on: the same tree copied to a second path hashes the same, and a tree that differs by
one byte does not.

Out of scope is the digest's value staying what it was. Dropping the `.xeno/plugin/` prefix
from every line changes the hash, and that is the point rather than a side effect. No release
compares a digest computed by one version of this code against one computed by another: a
binary carries the digest taken at its own release and the code that recomputes it, so each
release is self-consistent and nothing cross-version is affected. A register row records it.

Out of scope is checking a published binary against its tag. The issue names it as the other
thing nothing confirms, and this intent makes it possible rather than doing it; it needs a
released artifact to fetch and belongs with #198's verification.

Out of scope is G-Supply. The gate is correct and unchanged; what changes is that the thing
it judges can now be produced wrong only by a release that refused to build.

Out of scope is `cp -R`. Replacing the copy with something that cannot produce a different
tree is the other way to close this, and it is rejected in the design: the check is cheap and
general, and a cleverer copy would still be unverified.

No normative document is touched. Section 13 describes the anchor and says nothing about how
the two trees are compared, so no specification commit precedes this.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the hash definition, its two callers, and the gate that compares what it
produces.

`internal/plugin/plugin.go` is read for `Hash`, which is where the problem is: the line
format is a sum and a path relative to the repository root, and the relativity is the whole
defect. Reading it rather than the issue's summary of it is what settles that the fix is one
argument and a `filepath.Rel` base, not a new hashing scheme.

`internal/plugin/embedded.go` is read because it contains the claim being enforced. Its
comment says the embedded copy "is therefore the same bytes the digest was taken over", and
it also explains why a development build carries no tree at all, which is what decides that
the comparison has to be opt-in rather than unconditional: the script is useful in a tree
where the embedded copy does not exist and should not fail there.

`.github/workflows/release.yml` is read for the three steps in order. The order is what makes
the bug possible and also what makes the fix cheap — the copy already precedes the digest, so
the guard goes between them without reordering anything.

`internal/gates/gates.go` is read for G-Supply, to confirm that nothing else depends on the
digest's value and that the comparison is against `ExpectedDigest` alone. That is what makes
changing the hash basis safe, and it is checkable rather than assumable: a test that pinned a
literal digest would have made this a different intent, and none does.

Both normative documents are read for section 13's anchor and for anything fixing the hash's
definition. Section 13 says the runner carries the digest of its own plugin and that nothing
in the repository states the expected value; it does not define the line format, so the basis
is this package's to choose. Reading them also settles that no specification commit precedes
this, which the last three intents each had to establish rather than assume.

`docs/assumptions.md` is read for A86 and A87, which record the arrangement being changed:
A86 that G-Supply compares against `ExpectedDigest`, A87 that the release stamps, embeds and
then takes the digest. A87 describes the order this intent inserts a step into, so the row
and the workflow have to stay in agreement.
