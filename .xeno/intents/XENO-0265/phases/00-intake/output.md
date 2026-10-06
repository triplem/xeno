---
intent: github.com/triplem/xeno#235
phase: 00-intake
created: "2026-10-06T17:51:18Z"
schema_version: "1.0"
runner_version: dev+5276f4b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 67975d39df0a5d8af4550e705ccf6a8e00d945f87ec51bfa7a3fc57bc5ab7120
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

`budget` is called from `schema`, and `result` sets a check to `fail` where it carries any finding at
all. So a recorded context over its declared budget turns G-Schema red, and a red gate stops the
phase: `predecessorAllowsStart` refuses to begin the next one on an undecided failure.

Section 5 says the opposite, and now says it twice. Of the same check: "The finding is `advisory`, so
G-Schema stays `pass` and the phase stays green. That is deliberately a finding and not a red gate in
the sense of stopping work: it is visible, it can be decided like any other finding, and blocking
against a number nobody has experience with yet would be the wrong way round." And in the gate result
subsection: "**An advisory finding does not fail its check.** `advisory` marks a finding that is
reported rather than held against the phase."

The second of those arrived in XENO-0263 and is what makes this a step rather than an argument.
Before it there was no shape for what section 5 asked: the four check results are fixed by A4 and
A42, `result` is the only constructor, and a finding fails its check. **A finding that is visible,
recorded and decidable without stopping work did not exist in this runner.** Now the artifacts may
carry one and nothing writes or reads it.

## What is missing, precisely

`model.Finding` carries `ID`, `File`, `Cause`, `Next` and `Decision`. There is no `Advisory`.
`result` counts findings and does not look at them. `budget` produces its two findings like any
other check. So the clause is enumerated and unimplemented, which is the state the second standing
rule puts a specification change in and the first puts its code after.

## And the check reads nothing anyway

`budget` sums `f.Bytes` over `lock.Files` and compares the count of entries against
`p.Budget.Files`. **No P0 `context.lock.yaml` in this trail records a `files` list** — 0 of 117 when
#267 was filed — because `phase start` writes the lock and `scope set` writes P0's scope afterwards.
So `recorded` stays false, the entry count is zero, and the check reports nothing at the one phase
that declares a budget.

That is #267 and it is a different fault with a different cause. Making the finding advisory does not
produce it, and making the lock record files does not make it advisory. Both have to close before
this clause is exercised once, and this intent closes one.

## What `docs/clause-readers.md` says about it

Nothing. Section 5's budget clause has no row in the table at all — it was not among the 219
sentences the pass of 2026-10-03 extracted, or it was and fell outside the four kinds. The clause
that section 5 states most explicitly about a gate's behaviour is the one the document does not
list.

<!-- xeno:section:scope -->
## Scope

In scope is `Advisory bool` on `model.Finding`, with `omitempty`, so that no artifact already written
changes and nothing in the trail re-hashes.

In scope is `result` choosing `fail` only where a finding is not advisory. A check carrying nothing
but advisory findings is `pass`, and the findings are in `gate.yaml` with their ids, causes and
remedies — which is what section 5 means by visible and decidable.

In scope is `budget` marking both of its findings advisory, the file count and the byte total. Section
5's clause is about the budget check and names no other, so nothing else in the runner sets the field.

In scope is a row in `docs/clause-readers.md` for section 5's budget clause, which has none. The
clause the document does not list is the one section 5 states most plainly about a gate's behaviour,
and this intent is what gives it a reader worth naming.

In scope is the bound section 5 writes being visible in the code: "It is the exception and stays one…
the clause that asks for it says so where the check is described, and nothing else writes the field."
A comment where `Advisory` is defined and a test that nothing else sets it.

In scope is a phase whose only finding is the budget overrun coming out green, end to end, and the
next phase being allowed to start. That is the whole of what "not a red gate in the sense of stopping
work" means in this runner, and it is a behaviour rather than a field.

In scope is an advisory finding still being decidable. `gate approve` names a finding by id and
section 5 says this one "can be decided like any other finding"; nothing in the field may prevent it.

Out of scope is #267. The budget check reads nothing at P0 because no P0 lock records a `files` list,
and that is a change to when `phase start` writes the lock, against #215's reasons for writing it
once. Separate issue, separate cause, and both have to close before this clause is exercised.

Out of scope is making any other check advisory. Which others should be is a question per check, and
answering it in the abstract is how a bounded exception becomes a default — which is what section 5's
second paragraph exists to prevent.

Out of scope is `drift` and section 16's ninth limitation. It stays unimplemented. XENO-0263 recorded
why it was not the route and this intent does not revisit it.

Out of scope is a fifth check result. `advisory` is a property of a finding and the four results are
untouched, so no reader of a verdict learns a new state.

Out of scope is the derived phase status gaining a value. A phase whose only findings are advisory is
`green`, which is the existing derivation reading a check that passes; nothing is added to section
5's five statuses.

No normative document is touched. The clause was committed in XENO-0263 and this is the code that
follows it.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the clause, the three functions that have to change to honour it, and the two files that
bound how far it may go.

`docs/process-definition.md` is read for both halves of the clause as committed in XENO-0263. The
gate result subsection says what an advisory finding is — the check is `pass`, the phase is `green`,
the finding is in `gate.yaml` with its id, cause and remedy — and bounds it: "A finding is the thing
that fails, and an advisory one is readable only because it is rare; the clause that asks for it says
so where the check is described, and nothing else writes the field." The Context economy subsection
is the clause that asks for it. Both are read because the bound is the part no code can enforce and
therefore the part the code has to be written against deliberately.

`internal/gates/gates.go` is read for the three places. `budget`, which produces the two findings and
whose comment already quotes the sentence it breaks. `result`, which is the only constructor of a
check and the reason there was no shape for this: it counts findings and never looks at one.
And `schema`, which calls `budget` among six others, so the change has to leave a check carrying one
advisory finding and one real one failing.

`internal/model/model.go` is read for `Finding`, which gains the field, and for the four check
results, to confirm none of them moves. A4 and A42 fixed that set and XENO-0263's decision was taken
partly to leave it alone.

`internal/runner/runner.go` is not in the scope and is the reason the end-to-end check exists:
"stopping work" in this runner is `predecessorAllowsStart` refusing the next phase on an undecided
failure, which is a behaviour and not a field. The clause is satisfied when a phase whose only
finding is the budget overrun comes out green and the next phase starts, so that is run rather than
reasoned.

`docs/clause-readers.md` is read for the row that is not there. Section 5's budget clause is absent
from the table, which is the gap this intent is in a position to notice and fill, and the document's
own rule is that a reader is the thing that would fail if the clause were violated — so the row names
`result` and `budget` together rather than one of them.

`docs/assumptions.md` is read for A44, A4, A42 and A90. A44 is the baseline reasoning — a check that
fires with no action available is one people learn to route around — which is the same argument
section 5 makes for the budget. A90 is why the bound matters: a field that lets a finding not fail is
the most useful thing in this runner for anybody who finds a check inconvenient, and nothing but a
sentence will limit it.

`CLAUDE.md` is read for the three standing rules — the first because this code follows a
specification commit and comes after it, the second because the field was enumerated first, the third
because #258's code was a different branch — and for the convention on verifying a negative, which is
why "no P0 lock records a files list" and "nothing else writes the field" are each a search whose
pattern was shown to match something before its silence was believed.
