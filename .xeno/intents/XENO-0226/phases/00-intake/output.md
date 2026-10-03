---
intent: github.com/triplem/xeno#176
phase: 00-intake
created: "2026-10-03T11:34:08Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f7ba15bbe7766e6bd043ff29ce40f2058496fd833e90095ac8641f10af85fa0e
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

The budget is the one number in the trail that can change after a phase is sealed. G-Schema
compares a profile's `budget.bytes` against the size of the recorded context and takes that size
from the tree when the check runs, because until the commit before this one section 5 wrote the
lock's `files` as a path and a hash and nothing else. A phase inside its budget when it ran is over
it a week later if a file grew, and `xeno gate verify` says so about a phase nobody has touched.

**It is an exception by omission.** A66 and A74 both rest on the opposite rule: a phase is judged
against the rule set its own lock names, and a hash is recomputed against the file the artifact
names. A74 exists because writing four rule files turned twenty-four sealed verdicts red, and the
answer was to judge a phase by what it recorded. The budget asks the same question — what was this
phase given — and answers it from the tree.

**The specification now has the field.** The commit before this one adds `bytes` to the lock's
`files` entries and the sentence that says why, which is the order the first standing rule requires:
the document first, then the code that follows from it.

**What the code has to be careful about is the trail behind it.** Every lock in this repository
records no sizes, and there are seventy-nine intents' worth. A check that summed absent sizes would
read every one of them as a zero-byte context, which for a budget means permanently inside it —
quietly, which is worse than the finding it replaced. The lock that records nothing has to produce
no byte finding at all, which is the same distinction A74 drew for rules and the closure of the
register drew for decisions: absent is not empty.

<!-- xeno:section:scope -->
## Scope

**In scope.** `bytes` on each entry of the lock's `files`, written at `phase start` beside the
hash. The budget's byte check reading those recorded sizes instead of measuring the tree. A lock
that records no sizes producing no byte finding. Tests for all three, and for the file budget being
unaffected.

**Out of scope, and each for its own reason.**

The file-count budget. It is already judged against what the lock recorded — the length of `files`
— so it has the property this intent is giving the byte budget, and nothing about it changes.

Backfilling sizes into existing locks. They are sealed, their hashes cover them, and a lock
rewritten to carry a number nobody recorded would be a lock that describes a reading that never
happened. Seventy-nine intents keep what they have.

Any other use of the recorded size. A size per file is enough to sum, and nothing else asks for
one. It is not a cache, not a change detector — the hash is that — and not an input to any other
gate.

Adopting a profile here. That is the experiment this and #172 were the two conditions for, and it
comes next.

<!-- xeno:section:context-rationale -->
## Why this context

Section 5 is read as it now stands, one commit old: the lock's block with `bytes` in `files`, and
the paragraph that says the size is recorded because the budget is judged against what the phase
was given and not against what the tree holds now. That paragraph is the specification for this
intent and it was written for it.

A74 is read as the precedent and the warning. Its own measurement — four rule files turning
twenty-four sealed verdicts red — is why this intent's first question is what happens to the locks
that record nothing, and its rule for absence is the one this follows.

`internal/gates/gates.go`'s `budget` is read for what it does today: the profile from P0, the lock
from the phase, the count against `files` and a sum taken with `os.Stat` over the tree. The sum is
the only part that changes.

`internal/runner/runner.go`'s `informationBase` is read for where a file's hash is taken, because
the size is taken in the same place from the same walk, and for the ordering decision around it,
which must not move.

`internal/model/model.go` is read for `ContextFile`, which is the type the field goes on.

`internal/gates/budget_test.go` is read for the fixture that writes a profile, a lock and the files
the base names: it already writes the lock by hand, so it is the one place that has to learn to
write sizes.

Nothing outside the repository is needed, and the one thing that would have changed the shape of
this — the specification not having the field — was settled by the commit before it.
