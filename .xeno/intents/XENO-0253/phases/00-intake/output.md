---
intent: github.com/triplem/xeno#225
phase: 00-intake
created: "2026-10-05T17:14:05Z"
schema_version: "1.0"
runner_version: dev+9eec4ed.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6551e21cd2b7dfc7302b2459e67f04649adb4c424d22916241a0708651c3cbc6
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

`phase start` on a phase that is already under way rewrites `context.lock.yaml`, and the lock is
inside `artifacts_hash`. The artifact's `context_hash` then names a lock that no longer exists,
and the phase goes red on G-Schema at the next `phase finish`.

Two guards already stand in the way and neither covers this. `Start` refuses a phase whose run
marker exists, "already running"; and from #215 it refuses a phase that has a verdict, because a
second start would rewrite a sealed artifact and destroy the record of what the phase was given.
What is unguarded is the case between them: sections written, no verdict yet, and the marker
gone.

The marker goes missing in the ordinary course. A9 puts it under `.xeno/local/`, gitignored,
because it describes a machine's current state rather than the trail — so a phase begun on one
machine has no marker on another, and `retention` is `local_days: 30`, so it eventually has none
on the same machine either.

Reproduced rather than argued. With the clock moving, as it does outside a fixture: start a
phase, write its sections, remove the marker, start again. The second start returns no error,
the lock's hash changes, the artifact's recorded `context_hash` no longer matches it, and
`phase finish` returns **red on G-Schema**. In a fixture with a pinned clock the lock is
rewritten byte-identically and nothing shows, which is why this is easy to miss: the damage is
the `created` stamp moving, and only a real clock moves it.

So the failure is not that a resume is missing. It is that the obvious way to resume — run
`phase start` again — silently invalidates the work already done, and the tool reports it two
steps later as a schema finding about a hash, which is not where the mistake was made.

#225 asks a wider question, whether a resume would be useful, and answers most of it: the state
is readable through `xeno intent status` and every command that changes state prints the next
step. It names three gaps around that, and the maintainer chose the one that corrupts state
rather than the two that are inconvenient.

<!-- xeno:section:scope -->
## Scope

In scope is one refusal. `phase start` declines a phase whose artifact exists and whose lock the
artifact no longer matches, which is the state a lost marker leaves, and names the way on:
`section set` and `phase finish`, without a second start. That is the same answer #215's refusal
already gives for a judged phase, offered one step earlier.

In scope is the condition being about the artifact rather than about the marker. The marker's
absence is normal and says nothing — A9 makes it machine-local and `retention` expires it — so
refusing on a missing marker would refuse every legitimate first start. What distinguishes the
case is that `output.md` exists and `context_hash` names a lock, which only a phase already
under way can have.

In scope is the message naming both ways out, as the verdict refusal does: carry on writing, or
remove the artifact to start the phase over. A refusal that only forbids leaves a person with a
phase they cannot enter.

In scope is a test with a moving clock. The defect is invisible under the fixture's pinned one,
because the lock is then rewritten byte-identically, and a test that pinned the clock would pass
against the broken code.

Out of scope is `xeno intent resume`. #225 names the objection itself: it is a second entry
point into the computation `intent status` and the printed next step already perform, which is
how two answers to one question appear. The maintainer chose the refusal over it.

Out of scope is the age figure. #225 calls it decoration unless something acts on it, and
nothing would.

Out of scope is telling a resumed phase what changed under it. That is the third gap #225 names,
it is half-built already — a phase whose predecessor moved is reported `Stale` — and it is about
re-reading a predecessor rather than about re-entering a phase.

Out of scope is making the marker durable or moving it into the trail. A9 settles where it
lives, and recording a machine's current state in the trail is what that row refuses.

Out of scope is changing G-Schema. The red verdict it produces today is correct; what is wrong
is that nothing stopped the act that caused it.

No normative document is touched. Section 6 already says `phase start` refuses a second start of
a running phase, and this makes the tool able to tell that a phase is running when the marker
cannot say so, so the code moves towards the specification and no specification commit precedes
this.

<!-- xeno:section:context-rationale -->
## Why this context

The input is `Start`, the two guards it already has, the artifact they are about, and the
suggestion a refusal has to agree with.

`internal/runner/runner.go` is read for `Start` and for the comment #215 left on the verdict
guard. That comment is the argument for this one: it says a second start rewrites a sealed
artifact and destroys the record of what the phase was given, and both halves are true one step
earlier, before the verdict exists. The guard was written against the case somebody hit, not
against the condition that produces it.

`internal/runner/runner_test.go` is read for the fixture, and reading it is what found why this
survived. `newFixture` pins `Now` to a fixed time, so a lock rewritten in a test is
byte-identical and nothing changes; the defect is the `created` stamp moving, and only a clock
that moves moves it. A test written in the usual style would have passed against the bug.

`internal/runner/next.go` is read because a refusal and a suggestion can disagree. The suggestion
for a running phase names the sections still wanted and offers `xeno section set`; the refusal
has to point the same way, or a person follows the printed next step into the thing just
refused.

`internal/gates/gates.go` is read for what reports the damage today. G-Schema compares the
artifact's `context_hash` against the lock, so the existing failure mode is a schema finding two
commands after the mistake — which is the measurement of how far the report is from the cause.

Both normative documents are read for section 6's list of what `phase start` refuses and for
section 5's account of the lock. Section 6 already says a second start of a running phase is
refused, so this intent implements a clause rather than adding one, and no specification commit
precedes it. Reading section 5 settles that the lock is the only record of what the phase was
given, which is what makes the rewrite a loss rather than an inconvenience.

`docs/assumptions.md` is read for A9 and A12. A9 is why the marker is machine-local and
therefore why its absence cannot be the condition; A12 is the sequence rule the two existing
guards serve, which this one joins.
