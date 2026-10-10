---
intent: github.com/triplem/xeno#355
phase: 04-verification
created: "2026-10-10T13:13:54Z"
schema_version: "1.0"
runner_version: dev+b9beef5.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1512f0934411abb9802d0f802922d94ef0db3ea0a705b2e9467ac0897ca0f2a9
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Thirteen criteria, and they divide three ways rather than two. Six are answered here, four
by a test that runs in CI, and three only by a run of the thing: a release that reaches the
mirror, or the dispatch being dispatched. The third column says which, because a reader of
a sealed phase cannot otherwise tell a criterion that was checked from one that was
believed.

| criterion | checked by | answered |
|---|---|---|
| 1, the base is pulled from `ghcr.io` | the `load metadata` line of a release run's log | by a release run |
| 2, the copy preserves the digest | the second `imagetools inspect` in the step, which fails with the reference it could not resolve | by a release run on a miss |
| 3, the digest is written once | `grep -rn 'sha256:[0-9a-f]\{64\}' .github/workflows/` and the `sed` that reads `ARG BASE` | here |
| 4, `ARG BASE` names `docker.io` still | `git diff main -- Dockerfile`, and #351's `matchStrings` run in its own dialect | here |
| 5, Docker Hub is reached on a miss and nothing else | the `if`/`else` of the mirror block, read; the run log of two consecutive releases | here, then by a release run |
| 6, the dispatch reaches the image step | the parsed `if` and `VERSION` of the step | here, then by the dispatch |
| 7, a dispatch publishes the bytes that release published | the download block run against v0.60.2's real assets | here |
| 8, a dispatch for a version with no release fails | the same block run for `v9.9.9` and for a mismatched asset name | here |
| 9, a dispatch cuts no tag | the `release` step's `if`, and the four conditions it leaves false | here |
| 10, the image step reads nothing a dispatch skips | every variable of the step, against the step that writes each | here |
| 11, the page says where the base is fetched from | the row and the two paragraphs | here |
| 12, every pin is on the page, both directions | `TestEveryPinInTheTreeIsOnTheSupplyChainPage`, `TestEveryRowOfTheSupplyChainPageIsAPinTheTreeCarries` | CI |
| 13, the gate suite passes | `gofmt -l`, `go vet`, `go test -count=1 ./...`, `xeno gate verify` | CI |

**Why criterion 2 is not a test and is not an assumption either.** The property belongs to
buildx: that `imagetools create` with one digest-pinned source writes the manifest index
through unmodified. There is no buildx plugin and no docker daemon on this machine, so
nothing here can run it, and a Go test asserting what buildx does would assert a belief.
What the implementation does instead is name the mirror by that digest and inspect it after
the copy, so a copy that re-serialised the index fails at that line with the reference it
could not resolve. The criterion is therefore enforced on every miss by the pipeline rather
than measured once by this phase, which is a stronger arrangement than a passing test would
have been and is why it sits in the table as "by a release run" rather than as a gap.

**Why criterion 4 was run in node and not in Go or Python.** `matchStrings` are JavaScript
regular expressions and use `(?<name>…)`. Python's `re` rejects that syntax outright —
`unknown extension ?<d` — and a translation to `(?P<name>…)` would be a check of the
translation. Node runs renovate's own string against the real `Dockerfile`, which is the
same reasoning XENO-0285's P4 gave for not writing a Go test over these patterns.

**What nothing in this repository will re-check.** Criteria 1, 2, 5 and 6 are properties of
a run, and no test can hold them. The gap that follows is recorded below rather than
absorbed.

<!-- xeno:section:results -->
## Results

Ten met here, three outstanding on a run that has not happened and cannot happen before the
merge. None failed. The figures are from the tree as this phase found it, with the change
applied rather than reverted out of.

**Criterion 3, the digest is written once.** `grep -rn 'sha256:[0-9a-f]\{64\}'
.github/workflows/` returns **0 lines**. The step reads it instead, at `release.yml:346`:
`base=$(sed -n 's/^ARG BASE=\(.*\)$/\1/p' Dockerfile)`, then `base_digest=${base##*@}`.
Nothing else in the repository carries the value but the `Dockerfile` and the page's
abbreviated `sha256:a29215f6…`.

**Criterion 4, `ARG BASE` is untouched.** `git diff main -- Dockerfile` is **0 lines**.
#351's `matchStrings`, run as the JavaScript regular expression it is against the real
file, resolves to:

| group | resolved |
|---|---|
| `depName` | `docker.io/library/debian` |
| `currentValue` | `13-slim` |
| `currentDigest` | `sha256:a29215f6a35e51e22adffa17f89e9d2ef06214e64a2bad10d765c46aea49f11f` |

which is the tree's value in all three, so the manager still proposes the next digest for
the right image.

**Criteria 7 and 8, the dispatch's download.** Run outside the repository with the real
token and v0.60.2's real assets:

| run | result |
|---|---|
| `gh release download v0.60.2 --pattern xeno-0.60.2-linux-amd64` | the asset, mode **0644** |
| `sha256sum -c` against that release's own `SHA256SUMS` line | `xeno-0.60.2-linux-amd64: OK` |
| `chmod 0755` | `-rwxr-xr-x` |
| the build context after the two removals | `dist/` holding the binary and nothing else |
| the binary itself | `xeno 0.60.2` |
| `gh release download v9.9.9` | exit **1**, `release not found` |
| an asset pattern naming a version the release does not carry | exit **1**, `no assets match the file pattern` |

Both failures are before the `docker login` under `set -eu`, so a wrong input pushes
nothing. The 0644 is the defect P2 found this way and the `chmod` is the answer to it.

**Criteria 6, 9 and 10, read out of the parsed workflow rather than off the diff.** The
image step's `if` is `steps.release.outputs.new_release_published == 'true' ||
github.event_name == 'workflow_dispatch'` and its `VERSION` is
`${{ steps.release.outputs.new_release_version || inputs.version }}`. The `release` step's
`if` is `github.event_name == 'push'`, and the four conditions below it are unchanged
`steps.release.outputs.new_release_published == 'true'` — false on a dispatch because a
step that did not run has no outputs. Every variable the image step reads from outside
itself is `GITHUB_EVENT_NAME`, which the host sets on every run, `IMAGE`, and its own four
`env` entries; `IMAGE` is written by `where the image goes`, which carries no `if` at all.
So nothing the step reads is produced by a step either trigger skips.

**Criterion 5, the half that is readable.** The mirror block is an `if` on
`imagetools inspect "$mirror@$base_digest"` with the copy in its `else`, so the only
`docker.io` reference a release evaluates on a hit is the string it read out of the
`Dockerfile`. Whether two consecutive releases then show one request and no request is the
half a run answers.

**Criteria 11 and 12, the page.** The row reads `ghcr.io, mirrored from docker.io` and
keeps `sha256:a29215f6…`, the manifest index, by digest. Five tests over the page, the
workflow and the `Dockerfile` pass by name:
`TestEveryPinInTheTreeIsOnTheSupplyChainPage`,
`TestEveryRowOfTheSupplyChainPageIsAPinTheTreeCarries`,
`TestTheReleaseStampsTheVariablesThisPackageDeclares`,
`TestTheDockerfileCopiesTheBinaryTheReleaseBuilds` and
`TestTheImageRunsAsTheUidTheMountedWorkspaceBelongsTo`.

**Criterion 13, the gate suite.** `gofmt -l .` outside `vendor/` empty, `go vet ./...`
clean, `go test -count=1 ./...` clean, `xeno gate verify` exit 0 over **594 verdicts**, and
`xeno check commit-message` exit 0 on both commits of the change. The workflow parses as
YAML and its steps and conditions were read back out of the parse.

**Criteria 1, 2 and 6's second half, outstanding.** A release has to run, and the dispatch
has to be dispatched. Both are after the merge: the first release after this lands fetches
the base from `ghcr.io` and fills the mirror on its first miss, and the dispatch for
v0.60.2 is the recovery #356 decided. Nothing here could bring either forward, and the
verdict says so rather than claiming them.

<!-- xeno:section:gaps -->
## Gaps

**Three criteria wait on a run, and the run is after the merge.** Criteria 1, 2 and the
second half of 6 are properties of a release and of a dispatch, not of a tree. The first
release after this lands answers 1 and 2 on its first miss; the dispatch for v0.60.2
answers 6 and is the recovery recorded on #356. Until then the change is checked as far as
a tree can be checked and no further, which is why they are listed as outstanding in the
results rather than as met.

**Nothing re-checks that the mirror is reached instead of Docker Hub.** The `if` reads
correctly today. A later edit that moved the `--build-arg BASE=` back to the `Dockerfile`'s
own reference, or dropped it, would leave a release pulling from Docker Hub again, and
every test in this repository would still pass: `internal/model/supply_chain_test.go` holds
the page to the pins and says nothing about which registry a build fetches from, and
nothing else reads the image step at all. The failure would reappear as the next spent
quota, which is months of silence followed by the same 429. This is the same shape of gap
XENO-0285 recorded about its managers — a configuration whose job is to match text
elsewhere, with nothing holding it — and the answer there was the same: write it down.

**Nothing re-checks the dispatch's `chmod`.** If it were dropped, the dispatch would
publish an image whose `/usr/local/bin/xeno` is 0644, and no test would notice because no
test builds the image. It would arrive as a gate job failing on a permission in somebody
else's repository, which is as far from this file as a failure can get.
`TestTheImageRunsAsTheUidTheMountedWorkspaceBelongsTo` asserts three lines of the
`Dockerfile` for exactly this class of reason and could be the model for an assertion over
the image step, but a test that greps a workflow for `chmod 0755` checks the string rather
than the mode.

**What would close both, and why it is not here.** A job that builds the image and runs
`xeno version` in it would catch the mode, the base's resolution and the `COPY` path in one
check, and it is a change to the verify workflow rather than to the release. That is work
for another intent against another package, and it belongs to no work package named in
`docs/implementation-plan.md` on the reading that put this one in WP0, so it is a finding
about the plan as much as a piece of work. Written here and not absorbed.

**The two assumptions remain assumptions.** That the copy preserves the index is checked by
the pipeline on every miss, which is the best available and not a measurement. That a
container package `GITHUB_TOKEN` creates is readable by the same token on a later run is
not checked at all; its failure mode is a copy on every release rather than a broken one,
and it would show as Docker Hub appearing in every release log instead of in one per bump.
