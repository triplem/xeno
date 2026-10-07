---
intent: github.com/triplem/xeno#277
phase: 03-implementation
created: "2026-10-07T13:30:54Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 64154405566a46768fbead6ccd6c07c1b0be176b0d3dc9824e33c57f7ae1aa80
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

`internal/gates/gates.go`. The `byHand` variable is gone and `honest` reads

    honest := model.OneOf(f, model.WriterlessHash) || (f == "strings_hash" && goneBundle)

with a comment above it saying what was removed, that its reason was sound and its premise was
not, and that the two remaining terms are checkable. `hashShape`'s doc comment is replaced
whole: it no longer says the placeholder is honest where the artifact calls itself manual, it
names what the two remaining cases have in common, and it says that what an artifact declares
about the tool that produced it does not enter into a hash check, citing the issue.

Two test tables, because the removed behaviour was specified in two places and both are called
`TestHashFieldShape`. In `internal/gates/schema_test.go` the case "by-hand passes a manual
artifact" is now "by-hand is a finding even where the artifact says it was manual", with
`wantFinding: true`. In `internal/runner/runner_test.go` the case of the same name is now "a
manual artifact does not excuse its context hash", with the expected verdict moving from green
to red; that table asserts the colour a phase comes out as, which is what a developer sees,
where the gates one asserts the finding. Each carries a comment saying the field is a
declaration and decides nothing. The other six cases in the gates table and the other six in
the runner table are untouched and pass.

`.xeno/intents/XENO-1/phases/00-intake/gate.yaml` and `.xeno/intents/XENO-2/phases/00-intake/gate.yaml`.
Each goes from green to approved, carrying its two findings with a decision, a reason and the
name of the person who made it. The four findings are `F-0c8115` and `F-6ec27f` on XENO-1,
`F-187a0e` and `F-e0c34d` on XENO-2, all `context_hash says by-hand where a writer exists`.
Approvals, so nothing is owed.

`docs/assumptions.md`. A101 carries the decision, the measurement, what the rejected
replacement would have exempted, why the findings were released rather than fixed, that
approvals were chosen over overrides and why, and that the gate is stricter than it was.

The four artifacts themselves are untouched. Each still reads `context_hash: by-hand`; what
changed is the verdict over it.

`go build`, `go test ./...`, `gofmt -l .` outside `vendor/` and `go vet ./...` all pass.
`xeno gate verify` exits zero over 498 verdicts — it exited 1 between the code change and the
approvals, which is the state the four releases exist to resolve.

<!-- xeno:section:deviations -->
## Deviations from the design

One, and it is a change to the design's reason rather than to what was built.

**The replacement term was not merely rejected on judgement; it was measured and found empty.**
P2 records the decision to remove the `tool` term rather than replace it with one keyed on a
phase having no `context.lock.yaml`, and the reason given there is the measurement: all 495
phases in the trail carry a lock, so the replacement exempts nobody. That measurement was taken
while writing P2, not before it. The design was drafted expecting the replacement to be the
answer — it is the shape the issue's second option implies and the one that closes the gap
without any verdict moving — and the count is what ruled it out. P2 reads as though the
conclusion came first; it did not, and the order matters because a reader deciding whether to
revisit this needs to know the lock count is the load bearing fact and not a supporting one.

**The removed behaviour was specified twice, and P3 was judged before the second was found.**
`TestHashFieldShape` exists in `internal/gates/schema_test.go` and again in
`internal/runner/runner_test.go`, the first asserting the finding and the second the colour a
phase comes out as. Only the first was inverted when this phase was first finished, and
`go test ./internal/gates/` passed, so the package under change was green while the suite was
not. The second surfaced on `go test ./...` during verification. This section and the changes
section were rewritten and the phase judged again, which is what the `changed-after-verdict`
state is for; the alternative was to let P4 report a change P3 did not describe.

It is worth the record because of what hid it. Running the tests of the package being changed
is the obvious check and it is the one that passed. P1's criterion 5 said "the five other cases
of `TestHashFieldShape` are unchanged", in the singular, which is the assumption the criterion
was written under and the criterion is now met twice over.

Nothing else departs from P2. The term is removed, the comment keeps the reason, `hashShape`'s
comment names the shared property rather than counting cases, four approvals were recorded one
per finding with dates that differ by intent and a sentence saying the finding is correct, the
row is in the register rather than the code, and section 12 is untouched.
