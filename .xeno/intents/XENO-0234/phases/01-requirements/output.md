---
intent: github.com/triplem/xeno#183
phase: 01-requirements
created: "2026-10-03T17:31:04Z"
schema_version: "1.0"
runner_version: dev+9140756.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bfcc88d7238601c5b5088ad8a0cd3c57625ecd4e88c7c2c71569a93bc8c0cb3a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**A vendored plugin matching the digest the binary carries passes**, and the digest is recomputed
over the tree rather than read from anywhere in it.

**A plugin that differs by one byte fails**, with a finding naming both digests — the one found and
the one expected — because a finding saying only that they differ leaves the reader to compute two
values the gate already has.

**A plugin from an earlier release fails on the digest alone.** The gate reads no version, which is
what section 13 settles: "A vendored plugin from an earlier release has a different digest than the
one this runner expects, so it fails without any separate version check."

**No vendored plugin, with a runner that carries a digest, fails** rather than reporting
`not-implemented`: the check ran and the answer is that what it was anchored to is absent. The
finding says `xeno init --vendor`.

**A binary carrying no digest reports `not-implemented`** and no findings. That is a development
build, it is this repository, and it is the state section 5 defines for a check a runner did not
perform.

**The anchor is not readable from the repository.** A file in the tree naming the right digest — in
any plausible place, including inside the plugin — does not rescue a changed tree. This is the
property section 13 says is the point, so it is asserted rather than assumed.

**The gate runs first and from P0**, which section 7's table and the order's stated reason both
require.

**The release computes the digest with the same code the gate compares with.** A digest the checker
computes differently is a red verdict on every phase of every project that installs the release, so
there is one implementation and the release calls it.

**A released binary over this repository's trail reports `pass`** where the dev build reports
`not-implemented`, and `gate verify` stays at exit 0 across that difference, because both derive to
green.

**A42 is intact.** The gate path still reaches no network and no symbol index, which the `verify`
job checks, and the new import does not change either.

**Nothing in the repository that said this gate is unimplemented still says so.**

<!-- xeno:section:non-goals -->
## Non goals

**No resolution order.** `--plugin-root` and `XENO_PLUGIN_ROOT` stay read by nothing. This gate was
recommended ahead of them so that the containment and the thing it contains are not built in one
change, and the order remains a section 7 decision with the measurements in A84.

**No embedding of the plugin in the runner.** A released binary cannot yet vendor the plugin whose
digest it carries. A58 records the adjacent decision about `init`'s scaffold files and this is a real
gap of its own; the gate judges a plugin that is present.

**No reading of the lock's `plugin` block.** Section 13 excludes it: the anchor must be something the
checked side cannot influence, and the lock is written by the same runner into the same repository.
The block stays a record with no reader, which is the way round that leaves the question answerable.

**No comparison of `plugin_version` against the lock's version.** The other half of section 5's
sentence, and the same objection: both halves are in the tree.

**No signature and no tamper protection.** Section 13 says so in as many words — "This is not tamper
protection and does not pretend to be. Whoever can rewrite the vendored plugin can rewrite rules,
artifacts and the CI wrapper as well."

**No change to any other gate**, and none to `Applicable`, the table's order or the derived status.

**No backfill.** Every sealed verdict keeps the `not-implemented` it recorded, which was true when it
was written and is still true of the binary that wrote it.

<!-- xeno:section:constraints -->
## Constraints

**The anchor is the binary and nothing else.** Section 13's reason is the whole value of the gate:
"an expected hash stored beside the thing it describes proves only that both were written by the same
hand." Any implementation that reads an expectation from the tree is not this gate.

**The digest and no version.** Section 13 settles the downgrade through the digest, and reading a
version would turn a single act — new runner, `xeno init --vendor`, one commit — into a compatibility
matrix.

**`not-implemented` is the state for a check that did not run**, and section 5 defines it for exactly
the reason that applies here: a silent skip makes a green verdict mean less than it appears to.

**One computation of the digest.** The release and the gate must not implement it twice. The
definition is `internal/plugin`'s, written to the byte for #177, and the release calls the same code.

**A42.** The gate path reaches no network and no symbol index; `internal/plugin` reads files and
`internal/hashing` only.

**The import graph.** `internal/plugin`'s tests import `internal/gates`, so a gate importing the
plugin package makes the internal test a cycle. The tests become an external package rather than the
gate reaching around the problem.

**Sealed verdicts do not move.** The gate's result changes for a released binary, and `gate verify`
compares a phase's status rather than its per-check results, so a dev-built trail verifies under a
release. That is existing behaviour and this intent relies on it rather than changing it.

**88 columns, SPDX, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0, with the exit code
captured and not piped.**

**One intent, one branch, one issue** — `183-g-supply-has-an-anchor`, #183, labelled wp7, on main.
