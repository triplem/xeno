---
intent: github.com/triplem/xeno#247
phase: 00-intake
created: "2026-10-05T19:33:45Z"
schema_version: "1.0"
runner_version: dev+8e3b29b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a67b4d712e51825c75d8e41747c03e3975ae6ed8fe0f72a474ee2e4137962870
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

Three clauses have no reader, and for each of them the best available reader is a person. None
of the three can get a mechanical one without a change to a document the agent may not edit, so
all three have been open with nothing happening to them.

**Section 8's third requirement cannot be written down, let alone read.** It asks for "the
agent's recommendation with a reason". After #229 the recommendation has a reader in the writer;
the reason has no field on `model.Option`, so it can only live inside the question's text or
inside a consequence, where it is indistinguishable from what it shares the field with. #247.

**Nothing says questions are put one at a time.** Neither normative document contains the rule,
so there is no clause to read. The behaviour it would prevent is a batch, and XENO-0243 did it:
Q-1 on scope, Q-2 on which fields, Q-3 on inference, where the later options assumed an answer
to the first that nobody had given yet. #248.

**G-Test's mapping half has no reader and may never have a mechanical one.** Section 7 asks for
"mapping of acceptance criteria complete"; a gate would have to identify an acceptance criterion,
and 46 of 55 P1 artifacts carry no numbered criteria, so a completeness check would re-judge most
of the trail and would first need a numbering convention in section 9. #250.

What they have in common is the shape A90 describes and `docs/clause-readers.md` exists to
catalogue: a requirement whose reader was expected and never arrived. What is new is that for
these three the honest answer is that the reader is a person, and the repository has places for
that — a row in the audit, a convention in `CLAUDE.md`, a rule under `examples/rules/` — none of
which any of the three issues has used, because each was written as a request for a mechanism.

So the three have been waiting on a specification commit for something a person could have been
reading all along.

<!-- xeno:section:scope -->
## Scope

In scope is one row in `docs/clause-readers.md` for section 8's reason, saying it has no field
and no reader and that its home is the question's text. That is a measurement of what is, not a
rewording of the clause, so it needs no specification commit. The tool-requirement count moves
38 to 39. #247.

In scope is one paragraph in `CLAUDE.md` saying a decision put to a person is put one at a time,
with the reason. `CLAUDE.md` is the conventions file, not a normative document — #210 changed it
and the first standing rule binds the process definition and the plan — and it is read before a
question is asked, which is where the issue says the rule would do most good. #248.

In scope is one review rule under `examples/rules/`, that every acceptance criterion is answered
in the mapping, so a project that wants the mapping half read can have a person read it in the
P5 checklist. Under `examples/` and not `given/builtin/`, following A72: a review rule in the
shipped set reaches every adopter and would be answered at every review for ever. #250.

In scope is leaving the mapping row in the audit saying **nothing**. An example rule nobody has
enabled fails nothing, and the reader column means the thing that would fail if the clause were
violated. The paragraph under the table gains a sentence saying the example exists, which is the
honest treatment and the one most likely to be got wrong by a later reader counting coverage.

Out of scope is a field for the reason, a clause about sequence, and a numbering convention for
acceptance criteria. Each is an addition to a normative document, which the second standing rule
makes a specification change and the first makes a person's commit. All three issues say so and
this intent does not pretend otherwise.

Out of scope is enabling the example rule anywhere, including in this repository. Its own header
says how to adopt it, as the four beside it do.

Out of scope is any code. Nothing about what reads any clause changes; three descriptions and one
convention are added.

Out of scope is a gate over any of the three. A90's finding is the reason and the audit row is
where it is recorded rather than argued again.

No normative document is touched. All three clauses stay exactly as they are; what changes is
that what reads them is written down where a person will meet it.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the three clauses, the places a person-reader can be recorded, and the precedent for
each place.

`docs/process-definition.md` is read for section 8's sentence and section 7's G-Test row, because
both rows this intent writes have to quote what is asked rather than what the issues summarise.
Reading section 8 is also what settles that the reason's home can be the question's text without
rewording anything: the section asks for a recommendation with a reason and does not say where the
reason goes, so saying it goes in the text is a reading rather than an amendment.

`docs/clause-readers.md` is read for the reader column's definition — the thing that would fail if
the clause were violated — which is what decides that the mapping row stays `nothing` even after
an example rule exists. An example nobody enabled fails nothing, and a row that named it would
report coverage that is not there, which is the error #254 just corrected in the row beside it.

`CLAUDE.md` is read for its length and its own last line, "Keep it short", because this intent
adds to a file that asks not to be added to. It is 78 lines; the addition has to be a paragraph
that earns its place against that instruction, and the test is whether a session reading it
before asking a question would behave differently.

`examples/rules/documentation-follows-the-change.yaml` is read as the template for the fourth
example rule, and its header is read as the argument: it explains why a review rule in
`given/builtin/` would reach every project everywhere and be answered not-applicable at every
review for ever, which is A72's reasoning in the file that embodies it.

`internal/model/model.go` and `internal/gates/gates.go` are read for what is true rather than
remembered: `Option` has four fields and none holds a reason, `QuestionAsked` reads `Recommended`
and `QuestionShape` does not, and G-Test's row is half implemented. Each row this intent writes
asserts one of those, and asserting from memory is how the row #254 corrected came to be wrong.

`docs/assumptions.md` is read for A72 and A90. A72 is the precedent for `examples/` and A90 is
the finding that a reader which cannot fail is worse than none — which is the reason these three
get a person rather than a mechanism, and the sentence this intent is applying rather than
repeating.
