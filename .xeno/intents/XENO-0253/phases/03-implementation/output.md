---
intent: github.com/triplem/xeno#225
phase: 03-implementation
created: "2026-10-05T17:26:03Z"
schema_version: "1.0"
runner_version: dev+6adc0f9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: acfd289f9257410672f75b9f7213ca782efe2e1171c052adadfb76ccf39d5ea5
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

`internal/runner/runner.go`, two changes.

`Start` gains a third guard, after the marker check and the verdict check: a phase whose
`output.md` exists is already under way and is not started again. The comment carries why the
condition is the artifact rather than a comparison against the lock, and why a missing marker
cannot be the test.

The phase state gains a case. An artifact with no verdict reads `running` whether or not this
machine holds the marker, where before it read `not-started`. That is #225's first named gap —
a phase begun on one machine reading as not-started on another — and it is what let the printed
next step offer `phase start` for a phase that had already been started.

`internal/runner/runner_test.go` gains six tests: the reproduction with a moving clock; that
there is no harmless second start of a phase under way; that a phase with an artifact reads
`running` without its marker; that a first start is unaffected; that the two older refusals keep
their conditions; and that the suggestion no longer offers the start that is refused.

`f.moving()` replaces the fixture's pinned clock with a counter. It is one helper and it is the
reason this defect was findable: under the pinned clock a rewritten lock is byte-identical, so
the artifact stays consistent and a test in the package's usual style passes against the broken
code.

**Starting a phase over now takes the phase directory, not the verdict.** `TestStartAfter­RemovingTheVerdictIsAllowed`
became `TestStartingOverTakesThePhaseDirectory`, and both refusals name the directory rather
than `gate.yaml` or `output.md`. Removing the verdict alone leaves the artifact, and the artifact
is itself a phase under way, so the old instruction would have sent somebody from the first
refusal into the second. #215's comment is corrected to say so.

Two assertions in existing tests moved with the message, from the literal `gate.yaml` to
`model.PhaseDir`, which is what they were checking for in the first place.

<!-- xeno:section:deviations -->
## Deviations from the design

**The design's condition was wrong and the test is what showed it.** P2 settled on "the artifact
disagrees with the lock", compared with G-Schema's own hash. The reproduction failed against it:
before a second start the artifact and the lock still agree, because the staleness is what the
start *causes* and not what it finds. The condition passed in exactly the case the damage was
about to be done. It is now the artifact's existence, which is simpler, and the helper P2
specified is gone.

That also falsifies P2's alternatives section, which rejected "refuse whenever `output.md`
exists" as too broad on the grounds that a marker lost with nothing else changed leaves the lock
still describing what the phase was given. With a real clock `created` moves, so the rewrite is
never byte-identical and there is no harmless second start. The test asserting the harmless case
was replaced by one asserting there is none, under the moving clock that makes it true.

**The design said `next.go` needed no change, and the agreement test failed.** The suggestion
offered `phase start` for the very state being refused, because the phase state keyed `running`
on the marker alone and read a phase under way as `not-started`. `next.go` is still not edited —
the fix is in the state computation, which is where the wrong answer came from — but P2's claim
that nothing outside `Start` needed touching was wrong, and criterion 6 existing as a test rather
than an assumption is the only reason it was caught.

**#215's escape hatch changed and that is a behaviour change beyond this issue.** Removing the
verdict was the documented way to start a phase over; it now leaves the artifact, which the new
guard refuses. Both refusals name the phase directory instead, #215's comment is corrected, and
its test is rewritten to assert the new contract — that removing the verdict alone lands on the
second refusal and removing the directory starts over. A person following the old message would
have been sent from one refusal into another, which is worse than either.

One thing went wrong mechanically and is worth recording. A scripted edit removed the span
between two string anchors and the second anchor matched earlier than intended, mangling
`runner.go` into a file that redeclared a type. It was caught by the build, reverted with
`git checkout`, and reapplied as two independent replacements. Nothing was lost, because the
file held only this intent's changes.
