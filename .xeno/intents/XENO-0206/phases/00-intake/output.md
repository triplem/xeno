---
intent: github.com/triplem/xeno#132
phase: 00-intake
created: "2026-09-29T16:14:05Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+1dc1483.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2f76a601f45ae37c9a704d98a9f505b023d91631c9e9ecce719d3c40cdf88f2a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

`xeno intent status` without an intent prints a `status` column that reads `in-progress`
for every intent in this repository. Sixty rows of it. The obvious reading is that
nobody ever closed anything, and that reading is what prompted the question.

Nothing is stuck. Section 5 gives `intent.yaml`'s `status` exactly two values,
`in-progress` and `abandoned`, and section 8 says why there is no third:

> A merged intent writes no record at this level. It already carries one per phase,
including P5, > which is where the retrospective of a finished piece of work belongs.
`xeno intent close` exists > for the abandoned case … A merged intent needs no such
command: G-Complete has already run as > part of P5.

So the field means "not abandoned" and completion is the P5 verdict, which the listing
already shows in the next column as `05-review green`.

The defect is the listing, which #118 added and this intent's predecessor built. It
prints the stored field instead of the state, so the column a reader looks at first is
the only one that cannot tell a finished intent from one that stalled in P0. With
nothing abandoned, it is sixty identical words.

It is also the one column that contradicts how the rest of the tool works. `PhaseState`
carries the comment that it is computed, never stored, "because there is no position
that could go stale", and the listing's own next column is computed from the verdicts.
One row, two philosophies.

<!-- xeno:section:scope -->
## Scope

The `status` column becomes computed: `abandoned` from the field, which is the one thing
it can say; `complete` where P5 holds a decided verdict; otherwise the phase the work
has reached.

A decided verdict is anything but `red` or `provisional`, which is the test
`predecessorAllowsStart` already applies before it lets the next phase start. Reusing it
rather than listing the accepted statuses again is what keeps the listing and the
sequence agreeing about what a settled phase is.

Tests for the three states, and for the case that has none: an intent with no phases at
all.

Not the schema. `intent.yaml` keeps two values and `intent close` keeps writing one of
them.

Not a third status value. It would duplicate the verdict and could contradict it, which
is the thing computing rather than storing exists to prevent.

Not `intent close` for a merged intent. Section 8 refuses that in words, and running it
on the sixty intents here would mark them abandoned, which is false and is refused a
second time.

Not the phase column beside it, which already says what it says.

<!-- xeno:section:context-rationale -->
## Why this context

**The field is not wrong and the column is.** Section 5 decided that an intent records
only whether it was abandoned, and section 8 gave the reason: a merged intent's record
is its P5 phase, and a gate that reported after the merge could not gate it. Printing
that field as though it were a lifecycle is the error, and it is one this project made
about its own artifact.

**Computed, because the tool already says why.** `PhaseState` carries the sentence
"computed, never stored: there is no position that could go stale", and the listing's
neighbouring column is derived from the verdicts. A stored `finished` flag would be a
second answer to a question the verdict already answers, free to disagree with it, and
the disagreement would be invisible: nothing recomputes a flag.

**A decided verdict is defined once, in `predecessorAllowsStart`.** It refuses `red` and
`provisional` and accepts everything else, which is exactly what "this phase is settled"
means for the sequence. The listing asks the same question, so it uses the same answer
rather than a second list of accepted statuses that could drift from the first.

**`complete` rather than `closed`.** The word the specification does not use is
`closed`, because `intent close` means abandonment here, and a column reading `closed`
beside an intent that shipped would invite exactly the command that would falsify it.
`complete` says what a green P5 means and borrows no verb from the command surface.

**An abandoned intent keeps its own word.** It is the one state the field carries, it is
not derivable from the phases — an intent dropped in P1 has no P5 to ask — and
G-Complete's second mode exists for it. The computed column reads the field there and
computes nothing.

**What this does not fix.** The listing still cannot say whether a complete intent was
merged. Its P5 verdict says the work passed its gates, and whether a maintainer then
pressed the button is a fact about the host, which section 8 puts outside the trail.
Nothing here changes that, and a reader who wants it looks at the merge commit that
names the intent.
