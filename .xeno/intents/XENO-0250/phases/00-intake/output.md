---
intent: github.com/triplem/xeno#229
phase: 00-intake
created: "2026-10-05T15:20:18Z"
schema_version: "1.0"
runner_version: dev+325b27a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: af96e587b03f859e4fb73bc5736bd1ecbb9ce6893f3c6d14fad7820ab513c8d1
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

Section 8 asks three things of a question's options, and one of them is read.

> **A question is asked with options, not open.** Each one carries two to four options with
> their consequence, the agent's recommendation with a reason, and always a free entry as a
> further option.

`gates.QuestionShape` counts the options and finds the free entry. It does not look at
`Consequence`, which is `omitempty` on `model.Option` and therefore absent in a well-formed
question as far as anything can tell; it does not look at `Recommended`, which is a boolean no
option is required to carry; and the reason for the recommendation has no field at all, so it
can only live inside the question's text or inside a consequence.

So a question with four bare options and no recommendation is well formed today, which is
exactly the shape the section was written to forbid. The part of the clause that carries the
work — what each option leads to, which one the agent would take and why — is the part with no
reader, and the part that is mechanical is checked. That is #208's and #221's asymmetry again,
and `docs/clause-readers.md` catalogues the family.

Measured on this tree rather than assumed. Ten artifacts carry `open_questions`, holding
questions whose options number in the dozens, and **every proper option in the trail already
carries a consequence**. Adding that check finds nothing and leaves `gate verify` at exit 0
over 375 verdicts, which was run rather than reasoned about.

The recommendation is different and the difference is the whole of this intent's shape. One
question in the trail has no recommended option: XENO-3's Q-2, on how the schema version is
recorded, with three options each carrying a consequence and none marked. Adding a
recommendation check to `QuestionShape` makes G-Schema report a finding on that sealed P0,
because the shape is reached through `phaseResult`, and `gate verify` then reports
`DIVERGENT XENO-3 00-intake: committed status green, recomputed red` and exits 1. That was
measured, not predicted.

What is not written down anywhere is that questions are put in sequence, one decision at a
time. Nothing in either document says it. The behaviour it would prevent is a batch, which
XENO-0243 did with three questions whose later options depended on how the first was answered.

<!-- xeno:section:scope -->
## Scope

In scope is the consequence check in the gate. `QuestionShape` reports a question whose proper
options do not all carry a consequence, which costs nothing against history: every option in
the trail has one, and `gate verify` stays at exit 0 over 375 verdicts.

In scope is the recommendation check in the writer and not in the gate, decided by the
maintainer against the measurement above. `xeno question record` refuses a question that
recommends no option, so every question written from here carries one; the gate keeps the
checks history passes, so no sealed verdict is re-judged. The shape function therefore splits
into what the gate asks of any artifact and what the writer asks of a new question.

In scope is saying in the code why the split exists, where both halves are. A reader who finds
the writer stricter than the gate will otherwise read it as an oversight, and the next person to
tidy it would close the gap by moving the check into the gate and break `gate verify` on
XENO-3.

In scope is the free entry being exempt from the consequence check. Section 8 asks for "two to
four options with their consequence" and then "always a free entry as a further option", and the
free entry's consequence is unknowable by construction: it stands for an answer nobody has
written yet. The shipped fixture in `exchange_test.go` carries no consequence on it, which is
the existing reading and this intent keeps it.

Out of scope is the reason for the recommendation. It has no field, and adding one to
`model.Option` is an addition to what section 8 enumerates, which the second standing rule
makes a specification change first and the first standing rule makes a person's commit. This
intent records the gap rather than inventing the field.

Out of scope is the sequence rule. That questions are put one at a time is in neither normative
document, so there is nothing to check against; writing it down is a specification change and
therefore a person's. The finding is recorded with the example that prompted it.

Out of scope is re-judging XENO-3. Its Q-2 genuinely fails the clause, and the honest options
were to override the finding or to keep the gate silent; the maintainer chose the latter, and
this intent does not touch the verdict, the artifact or the row.

No normative document is touched. Section 8 already asks for the consequence and the
recommendation, so the code moves towards the specification and no specification commit
precedes this.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the clause, the function that half-reads it, the writer that shares that function,
and the trail the check would be applied to.

`docs/process-definition.md` is read for section 8's own sentence rather than for the issue's
quotation of it, because the whole question here is which of three requirements the code reads.
Reading it also settles that the consequence and the recommendation are already asked for, so
adding their checks needs no specification commit — and that the reason for the recommendation
is asked for in the same breath, with no field to put it in, which is where the line between
this intent and a specification change falls.

`internal/gates/gates.go` is read for `QuestionShape` and for `phaseResult`, and the second is
what made the measurement necessary. The shape is reached through G-Schema, which runs from P0,
rather than only through G-Questions, which runs from P5 — so a stricter shape check reaches
every artifact in the trail and not just the ones that reached review. An intent that assumed
G-Questions was the only caller would have predicted no effect on XENO-3 and been wrong.

`internal/runner/exchange.go` is read because the writer calls the same function through
`shapeRefusal`, which is what makes a writer-only check possible at all: the refusal path and
the gate path already share one definition, so splitting it is a matter of which half each
calls rather than of writing a second checker.

`internal/runner/exchange_test.go` is read for the `askable` fixture, which is the project's own
example of a well-formed question. It carries a consequence on each proper option, a
recommendation, and a bare free entry, which is the reading this intent keeps rather than
invents.

`internal/model/model.go` is read for `Option` and `Question`: four fields, two of them unread,
and no field for the recommendation's reason. The absence is the finding, and it is checkable in
the file rather than inferable from the issue.

The trail itself is read by a script over every `output.md`'s frontmatter, because "every option
already carries a consequence" and "one question has no recommendation" are claims that decide
the shape of this intent, and a wrong count either way would have produced the wrong design.
Both were then confirmed by building the check and running `gate verify`, which is what turned
the second one from an audit result into a measured divergence.
