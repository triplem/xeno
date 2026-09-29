---
intent: github.com/triplem/xeno#134
phase: 01-requirements
created: "2026-09-29T17:57:02Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+1dc1483.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a4f85ece3907c816e7e74c53cc46773a166ee21ee71f05cc44a6b519095b96d2
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

**AC1.** `xeno intent status` prints a heading row above the listing, naming the four
columns: when the intent was created, its key, its state, and the phase and verdict it
has reached.

**AC2.** `xeno intent status --intent KEY` prints a heading row above its phase lines,
naming the phase, the state and the verdict.

**AC3.** Without `--all` the listing shows the last ten intents by creation, in
ascending order, so the newest is the last row.

**AC4.** `--all` shows every intent, in the same order.

**AC5.** A truncated listing says how many intents it left out and names `--all`. That
line is not printed when nothing was left out.

**AC6.** A repository with ten or fewer intents prints all of them and no such line,
under either form.

**AC7.** The headings line up with the columns under them at every width the data takes,
including the widest state, `03-implementation`, and a key as long as the longest here.

**AC8.** Nothing about the data changes: same order, same columns, same values, same
computation. `Intents()` is untouched.

**AC9.** A repository with no intents prints nothing at all, not a heading over an empty
table.

**AC10.** Everything green stays green and `gate verify` matches every verdict.

<!-- xeno:section:non-goals -->
## Non goals

The order, which #118 settled and this keeps.

The columns, their contents and the computation behind them. `Intents()` does not
change.

Paging, filtering, sorting by anything else, or a `--limit N`. One flag and one number,
and a second knob would be a second decision nobody has asked for.

Headings anywhere else. `gate run` and `phase finish` print a verdict and its findings,
which are sentences and not a table.

Colour, alignment characters or a box. Plain columns, as now.

The ten itself as a configurable thing. It is a number in the source with a reason
beside it.

<!-- xeno:section:constraints -->
## Constraints

Nothing under `docs/` changes. This is presentation and the documents describe
artifacts.

The headings must line up with what is under them at every width the data takes. A
heading that drifts from its column is worse than none, because it labels the wrong
thing.

The truncation notice belongs in the output and not in the flag's help. A reader of a
short listing has to learn from the listing that it is short.

`Intents()` is untouched, so the change cannot alter what is shown. A presentation
change that quietly altered the data would be the worse kind.

No new flag beyond `--all`, and it is a boolean, which is what the three existing
booleans are.

One intent, one issue. The commits reference #134.
