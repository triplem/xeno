---
intent: github.com/triplem/xeno#132
phase: 01-requirements
created: "2026-09-29T16:15:15Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+1dc1483.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 417789584a2d0142ef65c75f2ccbb433f177867f358d76d7e4dcdb30eda7b9a7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**AC1.** An intent whose `intent.yaml` says `abandoned` is listed as `abandoned`,
whatever its phases hold.

**AC2.** An intent whose `05-review` carries a decided verdict is listed as `complete`.
Decided means anything but `red` and `provisional`, which is the test
`predecessorAllowsStart` applies.

**AC3.** An intent in flight is listed by the phase the work has reached, and a `red` or
`provisional` P5 is in flight rather than complete.

**AC4.** An intent with no phases at all is listed, and says so rather than reading as
complete or crashing.

**AC5.** The word is `complete`, not `closed`. `intent close` means abandonment in this
tool and a column reading `closed` would invite the command that falsifies it.

**AC6.** `intent.yaml` is unchanged, `intent close` is unchanged, and no new status
value exists anywhere.

**AC7.** The sixty intents this repository holds read `complete` where their P5 is
green, which is all of the ten carried past the intake, and by their phase where they
stopped, which is the rest.

**AC8.** `--intent KEY` prints what it printed before, to the column.

**AC9.** Everything green stays green and `gate verify` matches every verdict.

<!-- xeno:section:non-goals -->
## Non goals

The schema. Two values, section 5's, unchanged.

A third status value, which would duplicate the verdict and be free to contradict it.

`intent close` for a merged intent. Section 8 refuses it in words and it would write
`abandoned`.

Whether a complete intent was merged. Its P5 verdict says the work passed its gates;
whether a maintainer pressed the button is a fact about the host, which section 8 puts
outside the trail.

The phase column beside it, which already says which phase and which verdict.

Any other listing order. #118 settled that and this changes one column of it.

<!-- xeno:section:constraints -->
## Constraints

Nothing under `docs/` changes. Section 5 defines the field and section 8 explains why
there is no third value, and this intent agrees with both.

The definition of a decided verdict lives in one place. `predecessorAllowsStart` refuses
`red` and `provisional`, and the listing asks the same question, so it must not carry a
second list that can drift from the first.

`abandoned` is read, never computed. An intent dropped in P1 has no P5 to derive
anything from, which is the case section 8 says `intent close` exists for.

The one-intent form is a contract in practice, since every phase of every intent here
has been read with it.

No new field and no new command. One column of one listing changes.

One intent, one issue. The commits reference #132.
