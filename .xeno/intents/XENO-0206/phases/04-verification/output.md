---
intent: github.com/triplem/xeno#132
phase: 04-verification
created: "2026-09-29T16:19:33Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+1dc1483.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6862c68252188277d007ffbfc3b77600697f138d43644a86a0043284ea684295
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
| AC1 an abandoned intent reads `abandoned` | `TestTheListingComputesAbandonedCompleteAndInFlight`, second part, over an intent whose field says so and which has no phases at all |
| AC2 a decided P5 reads `complete` | the same test, first part, and the transcript's listing: ten intents read `complete` |
| AC3 an undecided P5 is in flight | `TestAnUndecidedFinalPhaseIsNotComplete`, with a P5 turned red by a question without options |
| AC4 no phases says so | `TestAnIntentWithNoPhasesSaysSo` |
| AC5 the word is `complete` | read, and the tests assert the string |
| AC6 no schema change | read: `intent.yaml`, `model.Intent` and `IntentClose` are not in the diff |
| AC7 the repository reads correctly | the transcript: ten `complete`, the rest at the phase they stopped in, and this intent in flight above its last decided phase |
| AC8 the one-intent form is unchanged | the transcript runs `intent status --intent XENO-0205` and prints six phase lines as before |
| AC9 everything green stays green | `go test ./...`, `gofmt`, `go vet`, `gate verify` over 115 |

Seven rows measured, two read. AC3 is the one that would have been wrong under either rejected
alternative, since both counted finished phases rather than reading their verdicts.

<!-- xeno:section:results -->
## Results

Four tests pass, the suite passes, `gofmt` and `go vet` are clean, and `gate verify`
matches 115 verdicts.

The listing now answers the question that prompted the intent. Ten intents read
`complete`, fifty read `00-intake` because that is where they stopped, and this one
reads its running phase beside its last decided one. Before the change all sixty read
`in-progress`.

`Decided` is asserted against all five verdict values, which matters more than it looks:
the sequence and the listing now share it, so a test that fixes its behaviour fixes both
callers at once. The rejected alternatives would have derived completion from the number
of finished phases, and `TestAnUndecidedFinalPhaseIsNotComplete` is the case that would
have caught either of them.

The one-intent form prints what it printed before, which is AC8 and the thing most
easily broken by a change to the type both forms share.

Read rather than executed: that no schema change happened, and the choice of the word.

<!-- xeno:section:gaps -->
## Gaps

**`complete` is a rendering and lives nowhere but the runner.** Section 5 has no such
value, so a reader who goes looking for it in the schema will not find it, and the
design says so. The word is defined by one function and printed by one column.

**It does not mean merged.** A decided P5 says the work passed its gates. Whether a
maintainer pressed the button is a fact about the host that section 8 puts outside the
trail, and a reader who reads `complete` as merged will be right almost always and
unable to tell when they are not.

**Fifty intents read `00-intake` and that is honest rather than informative.** They
stopped there because this project's practice before this session was to write an intake
and go no further. The column now says so plainly, which is a more uncomfortable reading
of the repository than `in-progress` was.

**An intent whose `intent.yaml` cannot be read has no state at all.** `summarise`
returns early with its problem, so the state column is empty beside the reason, which is
correct and was not designed: it falls out of the early return the listing already had.

**The phases of this intent have no cost record.** The hook shipped one intent ago is
read when a session starts and this session predates it, so the first real test of that
mechanism is still outstanding, exactly as XENO-0205's verification said.

**The evidence is a local run**, eleventh in a row.
