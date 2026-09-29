---
intent: github.com/triplem/xeno#134
phase: 04-verification
created: "2026-09-29T18:00:53Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c05acae.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4dd5ba832fb5e26e0cf350ad7b5983319fc8283bd2d9ab82459c91b290c6a510
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
| AC1 the listing has a heading | the transcript: `created  intent  state  phase` above the rows |
| AC2 the one-intent form has one | the transcript: `phase  state  verdict` above its six lines |
| AC3 the last ten by default, ascending | the transcript's default run: ten rows, `XENO-0207` last, and `TestTheDefaultIsTenAndTheNoticeCountsTheRest` over five counts |
| AC4 `--all` shows every one | the transcript: 64 lines, being 63 intents and a heading |
| AC5 the notice, only when truncating | `53 older, --all to see them` in the default run, absent under `--all`, and asserted at 0, 3 and 10 intents |
| AC6 ten or fewer shows all with no notice | the same test, at 3 and 10 |
| AC7 the headings line up at the widest data | `TestTheHeadingLinesUpWithTheWidestRow` and `TestThePhaseHeadingLinesUpToo`, which compare a heading against a row of the widest values each column can take |
| AC8 nothing about the data changed | `internal/runner` is not in the diff |
| AC9 no intents prints nothing | the transcript: zero lines against an empty root |
| AC10 everything green stays green | `go test ./...`, `gofmt`, `go vet`, `gate verify` over 121 |

Ten rows, all evidence. AC7 is the one the tests were shaped around, because it is the criterion a
heading can fail in a way that makes output worse than none.

<!-- xeno:section:results -->
## Results

The three tests pass, the suite passes, `gofmt` and `go vet` are clean, and `gate
verify` matches 121 verdicts.

The default prints a heading, ten rows with the newest last, and `53 older, --all to see
them`. `--all` prints 64 lines, which is 63 intents and the heading. An empty root
prints nothing at all, not a heading over an empty table.

The alignment is asserted rather than eyeballed. Both tests build a heading and a row
from the same format string and check that each heading word sits over a non-blank
column of the widest row that column can hold — `03-implementation` for the listing's
state, `changed-after-verdict` for the phase form's. That is the property, and it cannot
drift while the format string is shared.

`internal/runner` is not in the diff, which is AC8 by construction: the data cannot have
changed because nothing that produces it was touched.

Read rather than executed: the words chosen for the headings, and the reason ten is a
constant rather than a setting.

<!-- xeno:section:gaps -->
## Gaps

**Ten is a number in the source and somebody will want to configure it.** The design
says why it is not a setting — one flag answers the other question — and the first
person with two hundred intents will disagree. When they do, the place it belongs is
`project.yaml` under section 12, and that is a decision rather than an edit.

**The notice says `older` and means `created earlier`.** For this repository those
coincide. An intent created long ago and worked yesterday sorts by its creation, which
#118 chose deliberately and XENO-0200's own gaps section recorded; such an intent would
be called older while being the most recent work. Nothing here makes that worse and the
word inherits it.

**No test covers the rendering, only the arithmetic and the alignment.** Nothing asserts
that the notice reads as it does or that a row's date is truncated to ten characters,
because that would mean capturing output, which the deviations explain was avoided by
separating the two. The rendering is verified by the transcript and by reading.

**`cmd/xeno` now has a test file with three tests in it.** The command surface has no
test behind its exit codes, which is #110, and three tests about a table do not begin
that work. They do make the file exist, which is the smaller half of the problem.

**The evidence is a local run**, twelfth in a row.
