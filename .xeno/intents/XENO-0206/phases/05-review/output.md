---
intent: github.com/triplem/xeno#132
phase: 05-review
created: "2026-09-29T16:20:12Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+1dc1483.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 521040ca364145efa1510e7f481f7ab330e5fad33f72d9d77e67b90d3e0e8ae2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**Nothing under `docs/` changed, and the intent is about agreeing with what they say.**
Section 5 gives the field two values and section 8 explains why there is no third. The
defect was reading that field as a lifecycle.

**No schema change and no new status value.** `intent.yaml`, `model.Intent` and
`IntentClose` are not in the diff. A third value would duplicate the verdict and be free
to contradict it, which nothing recomputes.

**The predicate is defined once and used by both callers.** `Decided` replaces
`predecessorAllowsStart`'s own comparison, so the sequence and the listing cannot
disagree about whether a phase is settled. A test fixes both.

**Completion is read from the verdict, not from a count of phases.** Both rejected
alternatives counted finished phases, and both would have called a red P5 complete,
which is the state the sequence refuses to build on.

**The word was chosen against the command surface.** `closed` is the natural English and
it is the verb of the command that writes `abandoned` here, so it would have invited the
thing that falsifies the row.

**The rename makes two listings agree.** `State` was already the computed field of
`PhaseState` and `Status` its verdict; the intent listing had used the two words the
other way round.

**Every verdict still matches**, 115, and the one-intent form prints what it printed
before.

**The phases were written in order**, the measurement of the defect being a single
command rather than an investigation.

**What the fix exposes is said rather than left.** Fifty rows now read `00-intake`,
which is where this project's intents stopped before this session.

<!-- xeno:section:release-notes -->
## Release notes

`xeno intent status` without an intent now says what state each intent is in, computed
rather than read: `abandoned` where `intent.yaml` says so, `complete` where the review
phase carries a decided verdict, otherwise the phase the work has reached, and `no
phases` where there are none.

Before this it printed `intent.yaml`'s own status, which section 5 gives two values,
`in-progress` and `abandoned`. Nothing that shipped was ever marked otherwise, so every
intent read `in-progress` and the listing could not tell a finished one from one that
stalled in its first phase.

Nothing about the schema changed. There is no third status value, `intent close` still
writes `abandoned` and still exists for the abandoned case alone, and a merged intent
still needs no command: its record is its review phase, which is where a decided verdict
already says the work passed its gates.

`complete` does not mean merged. Whether a maintainer pressed the button is a fact about
the host, which the process definition puts outside the trail.

A decided verdict means anything but `red` and `provisional`, which is the test that
already decides whether the next phase may start. The two now share one predicate.

<!-- xeno:section:residual-risk -->
## Residual risk

**`complete` exists in the runner and in no document.** A reader who looks for it in
section 5 will not find it, because it is a rendering of a verdict rather than a value
of the schema. That is the right place for it and it is one more word this project
defines outside the documents it calls normative.

**It will be read as merged.** A decided review phase says the work passed its gates,
and for every intent here that was followed by a merge, so the reading will be right
until the first time it is not: an intent whose P5 went green and whose pull request was
never merged reads `complete` and is not. Nothing in the trail can distinguish those, by
section 8's design.

**Fifty rows now say `00-intake`.** That is honest and it is a different picture of the
repository than the one the column showed yesterday. Anybody reading it will ask why
fifty intents stopped after one phase, and the answer is this project's practice before
this session, which nothing records except those fifty directories.

**The state of an unreadable intent is empty.** `summarise` returns early with its
problem, so the column is blank beside the reason. Correct, and it falls out of an early
return rather than a decision, so nothing would notice if that changed.

**No cost record for this intent.** The hook shipped one intent ago is read at session
start and this session predates it, so #65's mechanism is still unproven in the way its
own review said.

**Accepted with the five named.** The state it replaces is a column that answered a
question nobody asked with a word that made ten finished intents look abandoned.
