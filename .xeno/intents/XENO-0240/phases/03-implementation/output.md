---
intent: github.com/triplem/xeno#215
phase: 03-implementation
created: "2026-10-03T20:23:37Z"
schema_version: "1.0"
runner_version: dev+b626f1a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7a682565e6b1cc90cf12cdce3d8292f8ee8efcb704bbb06a9ed2276f0557b41e
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

`internal/runner/runner.go`: `Start` gains a fourth precondition. Where
`<phase>/gate.yaml` exists it refuses with "has a verdict; redo the work with section set
and phase finish, or remove <path> to start it over". It sits with the other
preconditions, before anything is computed or written.

`internal/runner/runner_test.go`: two tests.
`TestStartIsRefusedWhereThePhaseHasAVerdict` asserts the refusal, that its message names
both ways forward, and that the lock's hash is unchanged after the refused call.
`TestStartAfterRemovingTheVerdictIsAllowed` asserts the deliberate escape still works.

`ASSUMPTIONS.md` gains A93, with the reasoning and the correction to the issue.

Nothing else. The duplicate `ChangedSince` begun in this intent was deleted before the
implementation phase started, so it is in no commit.

<!-- xeno:section:deviations -->
## Deviations from the design

**The issue was half wrong and the work shrank accordingly.** #215 said nothing tells a
repeated phase what changed. `Runner.ChangedSince` has done exactly that since #171,
`cmd/xeno/main.go` prints it at the end of every `phase start`, and two tests cover it,
including one for the no-profile case. The issue is corrected in a comment on itself rather
than only here.

I had written a second derivation — a `Changed` method and a list on the suggestion — before
finding the first. It was deleted. Had it shipped, two derivations of one answer would have
printed in two places and could have disagreed, which is worse than the gap it was written
for, and the existing comment says so in its own words: a printed derivation cannot drift
from its inputs, a recorded one can.

**The cause is procedural and worth naming.** The issue was filed from a reading of WP8 and
of `staleReads`, without searching for an implementation of WP8's first half. Both halves of
that package were discussed in the same breath as #213, and only one of them was unbuilt.

**One finding came out of the reading and is not fixed here.** This repository has no
`context-profile.yaml`, so `informationBase` returns nothing, every lock's `files` list is
empty, and `ChangedSince` is silent for every phase of every intent in the trail. The
mechanism is built and has never had an input. It gets its own issue.
