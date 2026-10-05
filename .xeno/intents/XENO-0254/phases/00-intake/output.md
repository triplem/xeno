---
intent: github.com/triplem/xeno#254
phase: 00-intake
created: "2026-10-05T18:55:52Z"
schema_version: "1.0"
runner_version: dev+e471bbb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 33b31309b5f15745460e2d648e32e31e19821e51502a2fe58c3ca17772107ddc
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

`docs/clause-readers.md` line 84 says a clause is read by a gate that does not read it, and
describes a third of what the clause asks.

| § | clause | reader |
|---|---|---|
| 6 | a question carries two to four options and one free entry | G-Questions |

Both halves were accurate when the pass was made on 2026-10-03, and #229 made them wrong three
days later without touching the row.

**The clause is a third of section 8's sentence.** The section asks for two to four options
*with their consequence*, the agent's recommendation with a reason, and always a free entry.
The row names the count and the free entry — what was read at the time. #229 gave the
consequence a reader in the gate's shape check and the recommendation a reader in the writer,
and left the row describing neither.

**The reader is not G-Questions, and was not when the row was written.** `QuestionShape` is
reached through `phaseResult`, so **G-Schema** calls it from P0. That is not a detail: it is
precisely what #229 turned on, because a recommendation check added to the shape function
reaches every artifact in the trail and makes `gate verify` report XENO-3's sealed P0 as
`DIVERGENT`. The row names the gate the clause sounds like it belongs to rather than the gate
that would fail if the clause were violated, which is this document's own stated definition of
a reader.

So the one row in the audit that is about questions is wrong about what is checked and wrong
about what checks it, in a document whose first paragraph says it is a measurement and whose
value is that a reader can trust the second column.

This is also the second time in four intents that the document has needed a row it did not get.
#212 added two for G-Test and found this one stale while doing it; #212's own learning proposed
the convention that would have prevented both, and nothing enforces it.

<!-- xeno:section:scope -->
## Scope

In scope is line 84, replaced by two rows because section 8's sentence has two parts with
different readers. The first names the options, their consequence and the free entry, read by
G-Schema's shape check and by G-Questions; the second names the recommendation, read by the
writer and by no gate at all.

In scope is naming G-Schema where the row said only G-Questions. The document defines a reader
as the thing that would fail if the clause were violated, and for the shape of a question that
is G-Schema from P0 — which is the fact #229 turned on and the one a later reader most needs.

In scope is the tool-requirement count, from 37 to 38, and a sentence saying these rows were
corrected rather than added, so the document's own account of its late rows stays true.

In scope is saying that the recommendation's reader is a writer and not a gate. That is the
asymmetry #229 left deliberately, nothing in a verdict can express it, and this document is the
only place in the repository whose purpose is holding it.

Out of scope is the reason for the recommendation, which has no field and no reader and is
#247. The rows describe what is read; a clause with nowhere to be written is that issue's.

Out of scope is the sequence rule, which is #248 and is not a clause yet.

Out of scope is changing any code. Nothing about what reads these clauses changes; this
corrects a description of what already reads them.

Out of scope is a general audit. The pass is dated and #202's reasoning stands: a re-reading
that fixed as it went would stop at the first finding. This corrects the row an intent made
wrong, and does not look for others — though P4's gaps will say that nobody has.

Out of scope is the gitignore finding #201 recorded. It was false, the entry exists at
`.gitignore:10` from #200, and the correction is in #254 and in a comment on PR #246 rather
than in this intent's changes, because there is nothing to change.

No normative document is touched. The audit is a measurement and says so; correcting it moves a
description towards the code rather than the other way.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the row, the clause it describes, and the two places that read parts of it.

`docs/process-definition.md` is read for section 8's own sentence rather than for the row's
summary of it, because the row's first fault is that it describes a third of the sentence.
Reading it gives the three requirements in the order the section states them, which is the
order the replacement rows follow.

`internal/gates/gates.go` is read for `QuestionShape`, `QuestionAsked` and the call in
`phaseResult`. The third is the one that matters and is invisible from the gate table: the
shape check is reached through G-Schema, so a clause about a question's shape is enforced from
P0 and not from P5, and the row's single named reader was wrong about which gate fails.

`internal/runner/exchange.go` is read for the call site that chooses `QuestionAsked` over
`QuestionShape`. That is where the recommendation's reader actually is, and reading it confirms
the asymmetry is in the writer rather than in a gate, which the second row has to state
precisely.

`docs/clause-readers.md` is read for its own definition of a reader — "the thing that would
fail if the clause were violated" — which is what makes the existing row wrong rather than
merely imprecise, and for its handling of rows that arrived after the pass, which this intent
has to extend rather than invent.

`docs/implementation-plan.md` and `docs/assumptions.md` are read for A90, which is this
document's reason to exist, and to confirm nothing else records what reads a question's shape.
A90 says a clause is one of four kinds and that the interesting ones are those something could
read; a row naming the wrong reader is worse for that purpose than a row naming none, because
it reports coverage that is not there.
