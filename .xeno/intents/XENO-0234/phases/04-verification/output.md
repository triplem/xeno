---
intent: github.com/triplem/xeno#183
phase: 04-verification
created: "2026-10-03T17:33:12Z"
schema_version: "1.0"
runner_version: dev+9140756.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1cf9b338ccb2a59043a6b1ecfcd73396764429e651ad108967aa6f7c6cb5a991
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Each acceptance criterion of P1 against the test that holds it.

**A matching plugin passes** — `TestTheVendoredPluginMatchingTheAnchorPasses`, anchored to a tree the
test wrote rather than to a literal, which would have pinned this repository's own plugin and needed
editing on every change to a skill.

**One byte fails, naming both digests** — `TestAChangedPluginFails`, which asserts the result, that
there is exactly one finding, that the cause contains the expected digest, that it contains the one
found, and that the file named is the plugin directory.

**An earlier plugin fails on the digest alone** — `TestAnEarlierPluginFailsOnTheDigestAlone`, where
the two trees differ only in a `version:` line, so a gate that read the version could have answered
differently and this one does not.

**No plugin with an anchored runner fails** —
`TestNoVendoredPluginFailsWhenTheRunnerExpectsOne`, asserting the result and that the finding says
`init --vendor`, because the finding's job here is to say how to put one there.

**No anchor reports `not-implemented` with no findings** —
`TestWithoutAnAnchorSupplyIsNotImplemented`. Both halves: a check that did not run must not carry a
finding, or a reader sees a complaint about something nobody checked.

**The anchor is not readable from the repository** — `TestTheAnchorIsNotReadFromTheRepository`,
which writes the correct digest into four plausible places, including inside the plugin's own
manifest directory, and requires the gate to fail anyway. This is the criterion that separates this
gate from the obvious wrong implementation; every other test passes under one that reads the lock.

**The gate runs first and from P0** — `TestSupplyAppliesFromTheFirstPhaseAndRunsFirst`, against
`Applicable`, which is the one list the order is read from.

**The release uses the gate's own code** — read in the diff, and held by the script existing: the
release calls `go run scripts/plugin-digest.go` and that file imports `internal/plugin`. A test
asserting it would be asserting a workflow's text.

**A released binary reports pass where the dev build reports not-implemented** — measured, below.

**A42 intact** — `go list -deps ./internal/gates` greps, run locally as the `verify` job runs them.

**Nothing still says the gate is unimplemented** — `grep -rn G-Supply` over the records, with the
two status paragraphs corrected and the README explaining what a reader of this trail meets.

<!-- xeno:section:results -->
## Results

**`go test ./...`** — eighteen packages ok, seven cases new.

**`gofmt -l .` outside `vendor/`** and **`go vet ./...`** — nothing.

**`./xeno gate verify`** — `verified 279 verdicts`, exit 0, no divergence, exit code captured into a
variable rather than piped.

**A42, run as the `verify` job runs it.** `go list -deps ./internal/gates` contains neither `net`
nor `net/http`, and not `internal/index`. The new import adds `internal/plugin`, which reads files
and `internal/hashing`.

**Measured with a binary built the way the release builds one**, `-X ...RunnerVersion=0.31.0 -X
...ExpectedDigest=$(go run scripts/plugin-digest.go)`:

| the tree | the gate |
|---|---|
| the plugin the digest was taken over | `G-Supply pass` |
| one comment appended to `secrets.yaml` | `fail`, naming `0aa3df0a…` found against `c116e0cc…` expected |
| `.xeno/plugin` moved away | `fail`, "no vendored plugin, and this runner carries the digest of one" |
| the same tree, development build | `not-implemented` |

**And the released binary over this repository's whole trail:** `verified 279 verdicts`, exit 0.
Sealed phases recorded `not-implemented` and recompute as `pass`, and `gate verify` compares a
phase's status rather than its per-check results, so the difference does not diverge. Both derive to
green.

**The digest itself.** `go run scripts/plugin-digest.go` and `plugin.Hash` are the same call, so
there is nothing to cross-check between them — which is the point. The definition was cross-checked
against a shell pipeline when it was written for #177, and that check still holds because the
computation did not change.

**This intent's own trail reports `not-implemented` on all six phases**, which is the honest state of
a `go build` and is what the README now explains.

<!-- xeno:section:gaps -->
## Gaps

**A released binary cannot vendor the plugin it carries the digest of.** `xeno init --vendor` copies
from `--plugin-from`, which defaults to the repository's own `.xeno/plugin`, so an adopter running a
released binary in a fresh repository has nothing to copy from. The gate is anchored to a plugin the
runner cannot install. A58 records the adjacent decision about `init`'s scaffold files being
embedded; the plugin is not, and until it is, this gate is anchored in a release that cannot deliver
what it anchors to. That is the largest remaining piece and it wants its own issue.

**Nothing verifies that the release's digest is the plugin it shipped.** The release computes the
digest over its own checkout, which is the right tree by construction, and nothing afterwards
confirms that the binaries carry the value the tag's tree produces. A check recomputing it from the
tag against a published binary would close it and needs the binary.

**The gate is untested against a real release.** Every measurement above used a binary built locally
with the release's flags. The first real test is the next release, and a mistake in the workflow's
quoting would compile in an empty anchor and switch the gate off for that release silently — which
is exactly what the script's non-zero exit prevents for a missing plugin and does not prevent for a
mis-assembled ldflags line.

**`gate verify` compares status and not per-check results.** It is why a dev-built trail verifies
under a released binary, which this intent relies on. It also means a verdict whose per-check
results changed for any other reason would not diverge either, and nothing here changes that.

**The resolution order is still not built**, which was the point of doing this first. With the gate
in place a swapped plugin becomes a red verdict rather than silence — for a released binary. Under a
development build it is still silence, because the anchor is absent, so the containment Shape 3
would need is only as good as the binary running it.

**No signature.** Section 13 disclaims it and says why, and says when it changes: "signing comes with
publication and not before".
