---
intent: github.com/triplem/xeno#237
phase: 03-implementation
created: "2026-10-06T12:07:31Z"
schema_version: "1.0"
runner_version: dev+7885661.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 65fa90703ca035ac3241d6828ae97652970b090f05d5609d73c40b7c2a07e7cb
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

Two files, both in `internal/gates`.

**`gates.go`: the two remedies about a sealed lock become one function.** `sealedLockRemedy(stale,
gated)` holds section 7's two routes, named in the clause's own order, with the earlier phase in the
re-run and the gated phase in the approval. The file-is-gone case and the file-has-changed case both
call it; the old strings differed and had no reason to, and the second of them offered "record why
the file left", which is not an act the runner carries out either.

    start 01-requirements over, which discards it and every phase after it, or a second
    person approves this finding on 02-design as still valid; section set and phase
    finish will not clear it, because the lock keeps what the phase was given and only
    phase start writes one

**`gates.go`: the unreadable case keeps its remedy and says why it differs.** "make it readable and
run the gate again; the tree is wrong here and not the lock, so nothing has to be approved or started
over." It was the one of the three that already worked, and the added clause is for the reader who
meets all three and would otherwise take the difference for an oversight.

**`gates.go`: a comment above the helper carries the reasoning.** Why the re-run is named first, why
`section set` and `phase finish` do not clear the finding, why the approval is a second person's act
rather than a command — `Runner.red` prints `gate approve` and says it hands over neither way out —
and why `gate.yaml`'s reader is the one that put the route in the string.

**`staleness_test.go`: three tests, one per case.** `assertSealedLockRemedy` holds what the shared
remedy must name in one place, so the two cases that share it fail together if a route is dropped.
The unreadable case asserts the opposite: that its remedy names "make it readable" and does *not* ask
for an approval or a start-over.

The check's conditions are untouched. `git diff` over `staleReads` shows changes inside the three
`finding` calls and nowhere else: the loop bound, the `touched` guard, the hashing, the causes and
the sort are as #236 left them.

## What was found while proving it, and is not fixed here

The end-to-end check needed a lock with files in it, and P0's had none. **No `context.lock.yaml` of
any P0 in this trail records a `files` list — 0 of 117.** `phase start` writes the lock, and P0's
scope is written afterwards by `scope set`, so P0's lock is born empty and nothing fills it. Of the
P1 locks, 17 of 66 carry files. So the staleness check has had almost nothing to read for most of
this trail, which is why a remedy naming an act the runner refuses survived four intents.

That is a finding about the plan and not about this issue, so it is written down rather than absorbed,
and it is filed rather than fixed: it is a change to when the lock is written, which is #215's
territory and section 5's.

<!-- xeno:section:deviations -->
## Deviations from the design

**The composed suggestion now names the release twice, and that is accepted.** `Runner.red` wraps the
remedy as "…: <next>. Then judge it again. Or a second person releases it, with gate approve or gate
override, naming <id> and a reason", so the reader of `xeno gate run` sees the approval mentioned in
the remedy and again in the sentence after it. Visible in the scratch run and recorded rather than
smoothed over. The alternative was to leave the release out of `next`, which would have been shorter
for that reader and would have left the reader of `gate.yaml` — who has no sentence beside the field
— with one route where section 7 gives two. The duplicate is for the reader who has the suggestion;
the omission would have been for the one who does not.

**"Then judge it again" reads oddly after a remedy whose second route is an approval.** The order in
the composed sentence is the remedy, then "judge it again", then the release. For the re-run route
that is right. For the approval route, judging again is what follows the approval rather than what
precedes it, and the sentence does not say so. It is `Runner.red`'s sentence and not this intent's,
and changing it was a non-goal; it is named here because the two readings of one sentence are only
visible once the remedy carries two routes.

**P0's empty lock was found by the end-to-end check and changed what it had to be built on.** The
first attempt made the ground move under P0 and gated P1, and G-Freshness passed: P0's lock records no
files. The check was rebuilt on P1's lock, gating P2, which is where the proof came from. The
original plan would have reported a green gate as evidence that nothing was wrong, which is the
failure shape #263's convention is about — a tool asked about something absent answering as it does
about something that does not match.

**The remedy does not name `gate approve` by name, and `gate.yaml`'s reader therefore has one hop.**
"a second person approves this finding on 02-design as still valid" tells them what and who, not
which command. That is the cost of respecting `next.go`'s stated reason for not handing the command
over, and it is a smaller hop than the two refusals it replaces.

**The tests were confirmed able to fail.** Each assertion was checked by mutation rather than by
being written and seen to pass: dropping the release route, dropping the re-run's cost, and giving
the unreadable case the shared remedy each fail the test that covers them. A90's objection applies to
a test as much as to a gate, and three assertions over a string are exactly the kind that pass
because they match something incidental.
