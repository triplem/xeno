---
intent: github.com/triplem/xeno#225
phase: 02-design
created: "2026-10-05T17:15:46Z"
schema_version: "1.0"
runner_version: dev+9eec4ed.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8b3c46afc2d20ac0c27b18769288c96da231efb0818fab5bbf18cbac398b47bc
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

The condition is "the artifact disagrees with the lock", not "the marker is gone". A missing
marker is the ordinary state on a second machine and after thirty days, so it cannot be the
test; what only a phase already under way can have is an `output.md` recording a `context_hash`,
and the question that matters is whether the lock still answers to it. If it does, nothing is
lost by starting again; if it does not, starting again is the act that loses it.

The comparison is the one G-Schema makes. The artifact's `context_hash` against the hash of
`context.lock.yaml` as it stands, by the same `hashing.FileHash` the gate uses, so the refusal
and the finding cannot disagree about whether a phase is consistent. A second definition of
"the lock the artifact was written against" would be the thing #201 and #242 each removed.

The guard goes beside the verdict guard, after the marker check and before anything is written.
The order is the states' order in time — running, under way, judged — and placing it there means
a phase in two of those states still gets the message about the one it is actually in.

The message names both ways on and borrows #215's shape: carry on with `section set` and
`phase finish` without a second start, or remove `output.md` to start the phase over. A refusal
that only forbade would leave somebody with a phase they cannot enter and no way to find out
how.

An unreadable artifact is not this refusal's business. Where `output.md` cannot be parsed or
carries no `context_hash`, `Start` proceeds as it does today: G-Schema reports an artifact it
cannot read, and a refusal that guessed would be refusing on the strength of a file it just
failed to understand.

The test moves the clock rather than mocking the lock. `newFixture` pins `Now`; this test
replaces it with a counter, which is the smallest change that makes `created` move and therefore
the only thing needed to make the defect visible. Pinning is right for every other test and
wrong for exactly this one.

`next.go` is not changed. Its suggestion for a phase with sections outstanding already offers
`xeno section set`, which is what the refusal tells somebody to do, so the two agree already;
criterion 6 is a test rather than an edit.

<!-- xeno:section:alternatives -->
## Alternatives

Refusing whenever `output.md` exists was the simplest condition and is wrong. A marker lost with
nothing else changed leaves the lock still describing what the phase was given, and starting
again costs nothing; refusing it would turn a harmless re-entry into an error and push people
towards removing artifacts, which is the destructive habit this refusal exists to prevent.

Rewriting the lock but preserving `created` was considered, so that a second start produced the
same bytes. It makes the common case silent and the uncommon one worse: the moment anything else
in the lock legitimately differs — a new commit, a changed rule set, a different information
base — the artifact goes stale with no warning at all, and the phase has lost the record of what
it was actually given while appearing fine.

Warning instead of refusing was considered. The write has already happened by the time anything
could warn, and the loss is the write; a warning after it is a description of damage.

Making `phase start` idempotent — detect the state and simply not rewrite — was the attractive
one and is rejected on honesty. A command that silently does nothing is a command that reports
success for an act it declined, and the person's mental model stays wrong: they think they
restarted the phase. The refusal corrects the model, which is what the printed next step is for
throughout this tool.

Writing the marker back without touching the lock was considered as a narrow repair: the state
is recoverable and the only thing missing is machine-local. It is tempting and it is the wrong
layer — `Start` would be half-running, and `phase finish` does not need the marker to work, so
restoring it buys nothing the refusal does not.

Teaching `intent status` to show the inconsistency was considered and is complementary rather
than alternative. It would report the state after somebody is already in it; the refusal is
about not getting there. #225's own listing of shapes puts the refusal first for that reason.

<!-- xeno:section:impact -->
## Impact

One function and two tests. `Start` gains a third guard between the marker check and the verdict
check; `internal/runner/runner_test.go` gains the moving-clock reproduction and the cases that
keep the refusal narrow. Nothing else in the tool changes, no artifact shape changes, and
`next.go` is read rather than edited.

Nothing existing is affected. A first start has no artifact, so the condition cannot fire, and
criterion 3 asserts that rather than arguing it. The trail is untouched because this changes what
a command refuses and not what any artifact contains.

What is gained is that the damage becomes impossible by the route people actually take. Today
the obvious way to resume a phase — run `phase start` again — invalidates the work and reports
it two commands later as a hash mismatch. After this it refuses at the moment of the act and
says what to do instead.

What a person loses is the ability to re-enter a phase by restarting it after something changed
underneath. That is the point rather than a cost, but it is a real change in what the tool
permits: somebody whose rule set or commit moved while a phase was open must now either finish
the phase as it stands or remove the artifact deliberately. The second branch of the message is
what keeps that from being a dead end.

The narrow scope is the honest limit. #225 names three gaps and this closes one. Nothing still
says how long an intent has been sitting, and nothing tells a resumed phase what changed under
it; both remain exactly as the issue describes them, and the two shapes the issue offered for
them — a figure and a command — were weighed and declined rather than forgotten.

One thing this does not improve is discoverability. A person who has been away from a tree for a
week still finds their way back through `xeno intent status` and the printed next step, which
#225 concludes is a resume in all but name. This change does not make that easier; it makes the
one path that destroyed state stop destroying it.
