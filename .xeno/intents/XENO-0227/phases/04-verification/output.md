---
intent: github.com/triplem/xeno#179
phase: 04-verification
created: "2026-10-03T12:18:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4d472b6.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 81831dc2f7a3eda3864164095b732d94b564a96ead5672e80aafc6217335fe26
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

**The seven fields and no eighth** — `TestStartingAnIntentDerivesEverythingButTheIssue` compares the
record read back off disk against a whole `model.Intent` by value, so a field added or dropped fails
it. It is not a field-by-field check for that reason.

**The id is the host, the repository and the key** — `TestTheQualifiedIdIsCompletedFromTheConfiguration`,
row "a bare number": `--for 176` against this repository's block gives
`github.com/triplem/xeno#176`. The same test asserts it end to end at the runner level.

**`--for` takes a longer prefix** — the same table's rows "the whole id", "the repository without
the host", "another repository on the same host" and "surrounding space", and
`TestWithoutATrackerBlockTheWholeIdIsGivenByHand` for the case that has no configuration to fall
back on, which also asserts that a bare key is refused there.

**`created` and the version fields are the runner's** — the first test's expected record carries the
fixture's clock, `model.SchemaVersion`, `model.RunnerVersion` and `model.PluginVersion`.
`TestANewIntentRecordsTheVersionTheArtifactsRecord` then compares the written `intent.yaml` against
the `gate.yaml` of a phase run by the same binary, which is the #177 criterion stated as the
comparison it is about.

**The key continues the sequence** — `TestTheNextKeyContinuesTheSequenceOnDisk` over eight rows: one
key, the highest rather than the last read, an unpadded key of the older scheme losing, three
paddings including a width that has to grow, a prefix holding a hyphen, and names that are not keys
being skipped. At the runner level the first test asserts `PROJ-2` after the fixture's `PROJ-1`.

**`--intent` names the key instead** — `TestAGivenKeyIsUsedAsGiven`, and
`TestASequenceThatCannotBeContinuedIsRefused` for the four cases where there is nothing to continue
and the refusal has to ask for the flag.

**Twice refuses and the first id survives** — `TestStartingAnIntentTwiceIsRefused`, which asserts
the refusal type and then reads the file to confirm the second `--for` did not reach it.

**Without `--for` it refuses naming the flag** — `TestStartingAnIntentNeedsTheIssue`, which also
asserts that the refused run left no directory behind, and
`TestIntentStartWithoutAnIssueIsRefused` at the surface for the exit code.

**A hand-written `intent.yaml` still works** — the whole existing suite, whose fixture writes one by
hand, and `./xeno gate verify` over the seventy-nine in this repository.

**The listing and the phase start** — `TestANewIntentIsListedAndItsPhaseStarts`, which asserts the
row, its date, its state and that no problem is reported, then starts P0 on it.

**The staircase** — `TestIntentStartNeedsNoKeyAndSaysWhichOneItChose`: 0 and the key on standard
output, then 1 with `refused:` on standard error and nothing on standard output.

**The refusals name a field or a flag** — `TestAnIdThatCannotBeCompletedIsRefused` asserts the text
of each: `--for`, `tracker.project`, `tracker.base_url`, and the space.

<!-- xeno:section:results -->
## Results

**`go build ./cmd/xeno`** — builds.

**`go test ./...`** — eighteen packages, all `ok`, none skipped. Thirty test cases are new, counting
table rows as the cases they are: twenty in `internal/model`, eight in `internal/runner`, two in
`cmd/xeno`.

**`gofmt -l .` outside `vendor/`** — nothing.

**`go vet ./...`** — nothing.

**`./xeno gate verify`** — `verified 241 verdicts`, exit 0, no divergence, nothing red, nothing
provisional. The count was 237 before this intent and the four are this intent's own phases.

**The command against this repository.** `xeno intent start --for 179` wrote
`.xeno/intents/XENO-0227/intent.yaml` with `intent: github.com/triplem/xeno#179`, key `XENO-0227`,
`status: in-progress`, `created: "2026-10-03T12:12:01Z"` and
`runner_version: 0.1.0-dev+4d472b6.dirty`. That last string is the result the issue asks for: it is
what the same binary wrote into every `gate.yaml` of this intent, where the seventy-nine intents
before it record `0.1.0-dev` beside artifacts recording the commit.

**The key it chose.** `XENO-0227`, derived, with no `--intent` given, after a directory whose highest
key is `XENO-0226` and whose lowest are unpadded keys of the older scheme.

**Evidence is self-reported and local.** The gate path makes no network call and these runs are on
this machine; the pipeline runs the same commands on a push, which is where the evidence for the
merge request comes from.

<!-- xeno:section:gaps -->
## Gaps

**No test asserts the host derivation against a real GitHub Enterprise or self-managed GitLab
address.** The rows are written from the addresses the wrapper table already carries and from the
form a self-managed deployment takes; nobody here has one to point at. The derivation is one line
and A78 states it, so a deployment that disagrees is a finding against the row rather than a bug
hiding in a branch.

**Nothing checks that the issue exists.** `--for 999999` writes an id for an issue nobody filed.
That is a non-goal of this intent and the open half of WP12, and the trail does not depend on it;
what it costs is that a typo in the issue number is sealed exactly as it was before.

**The key derivation is not tested against a directory holding both schemes at their real sizes.**
The table covers an unpadded key losing to a padded one, which is the property; the seventy-nine
real directories are covered only by the command having chosen `XENO-0227` once, by hand, in this
repository.

**`tool_version` is still a hand edit, twice per phase.** The runner has no source for the field and
the digest is rewritten by every `phase finish`, so both artifacts of all four phases of this intent
went red once and were corrected by hand. It is recorded as the implementation phase's learning and
it is a gap in the process rather than in this change, which did not make it worse.

**The consolidation of the tracker block is covered only by the existing tests.**
`EnforcementCheck`'s behaviour is unchanged and its tests pass untouched, which is the argument;
nothing new was written for it, because nothing new was intended.
