---
intent: github.com/triplem/xeno#65
phase: 04-verification
created: "2026-09-29T11:41:05Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+edd2b2e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2c0ba451002100dd08a1c74362f232f4d491391094a9552f6e3896c30a3fd165
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
evidence:
  - kind: test-report
    job: go-test
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | What proves it |
|---|---|
| AC1 `cost turn` appends a line from a hook's JSON | the transcript: a real hook payload in, and the ledger it wrote, with the session, the phase, the cumulative totals and the time |
| AC2 counts and identifiers only | read: the reader's struct has four integer fields and nothing that reaches `message.content`. `TestTranscriptTotalsSumsEveryAssistantMessage` asserts the sum, not the source |
| AC3 the live phase, or `none` | `TestTheLivePhaseIsReadFromTheRunnersOwnMarker`, and `TestAnUnattributedTurnIsRecordedAndCountedNowhere` for the other half |
| AC4 it never fails a turn | the transcript runs four broken payloads — not JSON, no transcript path, a path that does not exist, a file that is not JSONL — and all four exit 0 |
| AC5 `phase finish` writes the record | the `cost.yaml` of this intent's own 03-implementation, in the transcript, with `evidence: self-reported` |
| AC6 totals are differences between lines | `TestAPhaseTotalsTheDifferencesBetweenItsLines` and `TestTwoSessionsAreNotSubtractedFromEachOther` |
| AC7 nothing attributed writes no file | `TestAnUnattributedTurnIsRecordedAndCountedNowhere`, and the four earlier phases of this intent, which have no `cost.yaml` |
| AC8 no verdict changes | `gate verify` over 109, with a `cost.yaml` now inside a phase directory: it is excluded from `artifacts_hash`, so the seal is untouched |
| AC9 the hook is shipped and wired | `.xeno/plugin/hooks/README.md` and `.claude/settings.json`, read. **Not exercised**: settings are read at session start, so the wiring itself is unproven |
| AC10 no gate requires the file | read: A19 stands and no gate changed |

Seven rows measured, three read. AC9 is the one to watch: the mechanism ran, the wiring did not.

<!-- xeno:section:results -->
## Results

Seven tests on `internal/cost` pass, the whole suite passes, `gofmt` and `go vet` are
clean, and `gate verify` matches 109 verdicts with a `cost.yaml` now sitting in a phase
directory — which is the check that section 11's placement outside `artifacts_hash` is
real rather than stated.

The result is a number. Two invocations of the command the hook runs, against this
session's own transcript, appended two lines naming `03-implementation` of this intent,
and `phase finish` turned them into `tokens_in: 12`, `tokens_out: 13932`,
`tokens_cached: 4097272`, `evidence: self-reported`, one session. That is the first
token figure this repository has recorded and the first of #117's three numbers to come
from anywhere but the tree.

The failure paths were run rather than argued. Four broken payloads, four exits of zero,
because a hook that breaks a session is worse than a missing figure and AC4 outranks the
rest.

The four earlier phases of this intent have no `cost.yaml`, correctly: the ledger did
not exist when they ran. Nothing was backfilled, which is what a file that says
`self-reported` has to mean.

Read rather than executed: the hook wiring, the assumption rows, and that no gate
changed.

<!-- xeno:section:gaps -->
## Gaps

**The wiring is unproven.** `.claude/settings.json` is read when a session begins, so
the hook has never fired by itself: the two ledger lines came from the same command
invoked by hand with a real payload. Whether the `Stop` event fires as expected in a
working session, and whether the relative `./xeno` in the command resolves from the
harness's working directory, are things the next session finds out. That last one is a
genuine doubt rather than a formality.

**Most of this project's cost will read as unattributed, and that is the measurement.**
The ledger names a phase only where one was open, and for nine intents' worth of work
that would have been about seven per cent. The figure #117 wants therefore arrives
incomplete, and its incompleteness is a fact about how the work was done rather than
about the writer.

**A ledger line per turn, forever, in a gitignored file.** Nothing prunes it. Section 12
ties retention to the presence of `cost.yaml` and that rule is unwritten, so the ledger
grows without bound and the pruning that should read it does not exist.

**`tokens_cached` is the two cache fields added together.** Section 11 says "the cached
share" and names one field; a cache read and a cache write cost differently, and adding
them loses that. It is the reading that fits the schema and it discards information the
transcript had.

**No gate requires the record**, by decision, so a phase without one is complete and a
phase whose hook silently stopped working is indistinguishable from a phase worked
outside its window.

**Codex records nothing.** The reader is Claude Code's transcript, and section 14 names
two clients.

**The evidence is a local run**, tenth in a row, and the one thing it cannot stand in
for is the hook firing on its own.
