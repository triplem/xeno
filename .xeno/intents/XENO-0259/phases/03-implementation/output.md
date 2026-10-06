---
intent: github.com/triplem/xeno#260
phase: 03-implementation
created: "2026-10-06T09:20:42Z"
schema_version: "1.0"
runner_version: dev+081da51.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c56b4c137ce9e2ddbc9872a9ea5cd5597a74f0792882f1c67b69dbde9efbc812
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

Four files, no Go source.

**`.github/workflows/audit.yml` — the install.** The step that resolved the version now resolves the
action's sha beside it, out of the same file, with `test -n` on both: a `sed` matching nothing yields
an empty string, and an empty ref checks out a default branch, which is the moving target A28 exists
to prevent. A second `actions/checkout`, at the sha already pinned in this repository, places the
action under `release-toolchain/` in the workspace. The install step then runs the action's own two
commands in that directory — `npm ci --omit=dev`, then `npm install semantic-release@<version>` —
and the audit step reads the tree from there instead of from `/tmp`.

The comparison script, the manifest the job declares, the artifact it uploads, its triggers and its
`contents: read` permission are untouched. The check is the same check against a different tree.

**`.github/workflows/audit.yml` — the header.** Replaced rather than edited into, as the convention
asks for a paragraph being changed a second time. The claim that the install is "exactly as the
release installs it" is gone and what replaced it says what the action does, why a fresh resolution
sees less than a pinned one, and that one flag differs deliberately: the action passes `--only=prod`,
npm 12 deprecates it in favour of `--omit=dev`, the two select the same set and the replacement is
used. The note that the counts can now only move on a commit here is new, because that is the
property #260 asked for and nothing else in the file would say it.

**`.github/npm-audit-baseline.json`.** `critical` 0 → 2, `high` 10 → 23, `low` 0 → 1, `moderate`
unchanged at 2. The comment is rewritten and leads with which tree it describes, because a reader
arriving at a raised number needs that before anything else. It carries the 531-of-532 figure, the
twenty-eight advisory names by severity so that a reader can diff by hand what the check cannot, and
the measurement that neither upgrading nor fixing is available. The sentence that held #260's open
half is gone: the open half was whether to pin the tree, and the measurement answers it.

**`docs/supply-chain.md`.** The watched-and-not-watched paragraph is replaced: it said the audit
installs "the three pinned semantic-release packages exactly as the release does", which was wrong
twice over — one package since #121, and not as the release does. The table's semantic-release row
read 24.2.9 against a pin of 25.0.9 and now reads 25.0.9. The trivy-and-audit boundary above it is
untouched, and it is what makes two criticals an audit finding rather than a release one.

**`docs/assumptions.md`.** A44's status column records that its evidence moved. The assumption's
shape is unchanged and so is its reasoning — none of the twenty-eight is this project's to fix, since
the sixteen new ones move only when the action's sha does, which is a commit here. What is new is
that it now covers two criticals, and the approval was not asked about those.

## The three identifiers that now live in one place

`release.yml` holds the version and the sha; `audit.yml` reads both and holds neither. `npm-audit-
baseline.json` holds the counts those two produce. A bump of either identifier moves the counts, the
job fails, and the commit that bumps it raises the baseline with the figures in its message — which
is the file's own rule for itself, and is the first time that rule has had anything to do with a
commit in this repository rather than with npm's publishing schedule.

<!-- xeno:section:deviations -->
## Deviations from the design

**`--omit=dev` where the action passes `--only=prod`.** Named in P1's constraints, decided in P2 and
carried out here, so it is a deviation from the action and not from the design. Verified rather than
assumed: both flags were run against the same checkout on 2026-10-06 and produced the same 28
advisories in the same split, 2 critical, 23 high, 2 moderate and 1 low.

**`low` rises from 0 to 1 and P1 did not say it would.** Criterion 3 names the figures as 2/23/2/1
and the criterion is met, but the design's impact section wrote the change as "0/10/2/0 to 2/23/2/1"
without remarking that a severity class with no entries acquires one. The advisory is `diff`, a
denial of service in `parsePatch`, reached through npm-as-a-library. It makes no difference to the
check, which compares each class independently, and it is recorded because the baseline's own rule is
that a raised number belongs in the commit that raises it — all four classes are named in the message
rather than the three anybody would think to look at.

**The action is checked out into the workspace rather than into `/tmp`.** The old job built its tree
in `/tmp/release-toolchain` and `actions/checkout` refuses a path outside the workspace, so the tree
now sits at `release-toolchain/` beside the repository. Nothing reads it but the two steps that
build and audit it, and the job uploads only `.xeno/local/scan/`. It is worth naming because the
workspace is no longer only this repository for the length of this job, which is a difference a
later reader of the job could be surprised by.

**No deviation on the sed for the sha.** It was written expecting `- uses:` and the line in
`release.yml` is a bare `uses:` under a named step, so the pattern was tested against the file before
it was committed rather than after: `^ *uses: cycjimmy/...@\([0-9a-f]\{40\}\)` resolves
`b12c8f6015dc215fe37bc154d4ad456dd3833c90` and the `- uses:` form resolves nothing. Both were run,
because a `sed` that matches nothing and a `sed` that matches the wrong thing look identical in a
diff and the first is exactly what `test -n` exists to catch.

**Nothing was recorded from a check over a tree that had not been built.** The convention added in
#263 applies to every negative claim in this intent: `handlebars` being absent from the old audit's
tree, the 531-of-532 comparison and the equality of the two npm flags were each taken from an
install performed on 2026-10-06, not from reading a manifest. The full reproduction — resolve both
identifiers out of `release.yml`, check the action out, `npm ci --omit=dev`, install the pinned
version, audit, compare against the baseline — was run end to end and the comparison exited 0.
