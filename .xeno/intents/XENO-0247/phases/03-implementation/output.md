---
intent: github.com/triplem/xeno#242
phase: 03-implementation
created: "2026-10-05T12:17:40Z"
schema_version: "1.0"
runner_version: dev+8f4b75b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 429a9c7b3a1e27d1f31f22249bbe1ac198493973300bee176ae95e524ea33e03
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

`internal/runner/review.go`, new. `ReviewAnswer(key, model.ChecklistEntry)` validates, reads
the review phase's artifact through `exchange.go`'s `artifact`, upserts the entry and writes
it back through `amendFront`. Two helpers beside it: `upsertChecklist`, which replaces at the
index it finds or appends, and `unanswered`, which counts from the rules to the entries.
`ReviewAnswered` carries the entry, whether it replaced, and what is left.

`cmd/xeno/main.go`: a command table row for `review answer` with `needsKey` and not
`needsPhase`, a `--note` flag, a `note` field on `opts`, a usage line, and `cmdReviewAnswer`.
`--result`'s help text now names both meanings it serves, because the flag set is global and
`evidence declare` had it first.

`internal/runner/review_test.go`, new. Eight tests: the write and the report, replacement in
place, eight refusals as a table, `met` without a note, the body byte for byte, the answer
landing in P5 while another phase is open, an answer before the artifact exists, and a rule
set holding no review rule.

The refusals are ordered by what they need. Rule present, no source, then the three about
`result`, all of which are knowable without touching the disk; then the rule against the
effective set, which needs the rule tree; then the artifact is opened. A malformed call
therefore cannot leave a half-amended file, and the test asserts the checklist is still empty
after each refusal.

Two sets are imported rather than restated. `model.ChecklistResults` and
`model.ChecklistNeedsNote` are the gate's authority on a result and
`rules.Effective` filtered to `rules.Review` on a rule, so a refusal on the way in and a
finding after the fact cannot disagree.

`model.OneOf` is used for both membership checks. A local `contains` was written first and
removed: `OneOf` already does it, and this repository has no `slices` import anywhere, so a
second helper would have been a third spelling of one idea.

This intent's own P5 checklist is written with the command, which is criterion 9 and the only
check that the writer works where it is meant to be used.

<!-- xeno:section:deviations -->
## Deviations from the design

Two deviations from #242, both declared in P0 and P1 before the code was written, and both
against the issue rather than against this intent's own plan.

The issue asked the command to refuse a phase whose verdict is written, "the way the decision
commands do". It is not implemented, because they do not. `exchange.go`'s package comment says
neither `question record` nor `decision record` looks at `gate.yaml`, and gives the reason: a
write after a verdict changes the hash, and the next `phase finish` reports the phase as
changed after its verdict, which is #216's mechanism. D-6 is what the issue was reading, and
it settles the same question for `gate approve` and `gate override`, which carry `against`, the
`artifacts_hash` the decision was taken on. A checklist entry has no such field, so it can
assert nothing a recompute would contradict. The file's comment records this, because the next
reader will have the same question.

The issue offered `[--phase NN]`. There is no phase argument. G-Policy calls `reviewChecklist`
only when the phase is the last one, so the field belongs to one phase and a flag would offer a
choice of one. This follows `ScopeSet`, which takes no phase because the scope is P0's.

Three additions beyond what #242 listed, each small and each from reading the gate rather than
the issue.

A `source` is refused. Section 12 gives `source: lens` to a lens entry, which answers no rule,
and this command answers one; a flag whose only honest value is absent would be a field
existing to be left alone, so the field is refused if set rather than offered.

A rule set with no `review` rule at all is its own refusal. The general message lists what the
set holds, and listing nothing would read as though the rule were misspelled rather than as a
tree with nothing to answer.

An empty `--result` is separated from a wrong one. "`--result` is required" and "`partly` is no
checklist result" are different mistakes, and one message for both would name the three without
saying which of the two happened.

One thing was written and removed rather than deviated from: a local `contains`. `model.OneOf`
already is it. Noted here because the first version of this file carried both, and the reason it
does not is the convention rather than the correctness.
