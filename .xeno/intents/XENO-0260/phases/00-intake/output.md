---
intent: github.com/triplem/xeno#258
phase: 00-intake
created: "2026-10-06T11:31:31Z"
schema_version: "1.0"
runner_version: dev+7885661
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cf4bfcbb4c2038db8ea51a282848c2e9cb4b79b835bc9198156956566ac49503
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

`docs/clause-readers.md` tells a reader where the specification commit for a clause has to go, and
for one of its two unread clauses it names the wrong section.

The passage under the clause table reads: a completeness check for G-Test's mapping half "would
first need a numbering convention, which is an addition to section 9 and therefore a specification
change". Section 9 is the Rule model. `grep -n "acceptance" docs/process-definition.md` returns 572
and 692 — both inside section 5, the Language and Rendering subsections, where `template.yaml` holds
the section ids, their order and the required fields — then 710 and 713 in section 6's phase table
and 942 in section 7's G-Test row. Nothing between 1255 and 1485, which is all of section 9.

#258 repeats the same number, twice, and so does the comment on it that recorded the decision. The
document is the thing a person reaches for before writing the commit, so a wrong section number in it
is wrong at exactly the moment it is read.

## The figures the passage carries measure half of what the check needs

The same passage reads: "of the 55 P1 artifacts in this trail, 46 carry no numbered criteria at all,
and of the nine that do, one already fails a check by number. Measured on 2026-10-05."

A completeness check reads a P4 `test-mapping` and asks whether it covers every criterion P1 raised.
It therefore needs both halves: criteria that can be named, and a mapping that names them. Only the
first was measured. Measured on 2026-10-06 across the trail as it now stands:

| | artifacts | identifiable | not |
|---|---|---|---|
| P1 `acceptance-criteria`, as a numbered list | 64 | 19 | 45 |
| P4 `test-mapping`, citing a criterion by number | 64 | 7 | 57 |

**The unmeasured half is the worse one.** Even where criteria are numbered the mapping usually does
not cite the number, so the cost the passage states — that a check "would re-judge most of the
trail" — is 57 verdicts and not 46. The direction of the conclusion is unchanged and its size was
understated by a quarter, which matters because the size is the whole of the argument against
applying the check to the trail.

## What the two rows wait on, and what a reader cannot tell

Two rows of the clause table report no reader: section 8's "the recommendation carries a reason",
which has no field to carry it, and section 7's "mapping of acceptance criteria complete", which has
nothing at all. #258 has now decided both, and the decisions are recorded in the issue and nowhere a
reader of the document reaches. So the rows say a clause is unread and leave open whether that is
a gap nobody has looked at, a gap somebody is working on, or a deliberate answer — and the document's
own closing section says that triage "belongs to whoever reads this".

The reason field was decided onto `model.Option`, required where `recommended: true`. The mapping
check was decided onto the template version: `requirements@1.1.0` and `verification@1.1.0` carry the
numbering, and G-Test judges only artifacts declaring 1.1.0 or later, which needs no new anchor
because all 64 P1 artifacts declare `requirements@1.0.0` today and `template` is already a
frontmatter field. Both wait on a specification commit, which the first standing rule makes a
person's.

<!-- xeno:section:scope -->
## Scope

In scope is the passage in `docs/clause-readers.md` that explains why the mapping half has no
reader. It is replaced rather than edited into: the section number is wrong, the figures measure half
of what the check needs, and the sentence about what it waits on is now answerable. Three changes to
one paragraph is exactly the case the convention is about.

In scope is a short passage for the section 8 row, which has none. The mapping row has three
paragraphs explaining itself and the reason-field row has a bolded cell and an issue number, so a
reader learns from the table that one clause is unread and nothing about why or what would change
it.

In scope is saying, for both rows, that the clause has been decided and what the decision waits on.
The document's closing section says the triage of these rows "belongs to whoever reads this"; a row
that has been triaged and does not say so sends the next reader to do it again.

In scope is the re-measurement as a dated figure beside the original rather than over it. The pass
itself was made on 2026-10-03 and says a reader named by symbol "is wrong the day the symbol is
renamed with nothing to say so, which is the most a document can do about its own ageing". Replacing
a dated measurement with a newer one without saying so would be the same fault in the other
direction.

Out of scope is the reader column of either row. Section 8's still says there is no field and the
mapping's still says nothing, because that is what is true until the specification commits exist. A
row naming a decision is not a row naming a reader, and the table's own convention is that a reader
is "the thing that would fail if the clause were violated".

Out of scope is the specification commit for either clause. The first standing rule makes both a
person's and makes them come before the code, and the wording for each is drafted on #258 rather
than written here.

Out of scope is `reason` on `model.Option`, the template version bump, the check in G-Test and their
tests. All four follow a specification commit that does not exist, and starting any of them is the
second standing rule's invention.

Out of scope is a third row for the sequencing convention. #248's decision is that it stays in
`CLAUDE.md`, and #255 declined it a row because this document is a pass over the two normative
documents and a convention in `CLAUDE.md` is a clause in neither. Declining it again is what keeps
that a rule.

Out of scope is correcting #258's own text. An issue is not a file in the tree and the comment that
records the correction is already on it.

No normative document is touched. `docs/clause-readers.md` says of itself that it "is a measurement,
not a specification. Where it and the documents disagree, the documents win."

<!-- xeno:section:context-rationale -->
## Why this context

The input is the document being corrected, the document it misnames a section of, and the three
files that decide whether anything else has to change with it.

`docs/clause-readers.md` is read as the thing being changed and for its own rules about itself. Three
of them bind this intent: that it is a measurement and not a specification, so the agent may change
it where a normative document would wait for a person; that its figures are dated and its ageing is
visible rather than repaired, which is why the re-measurement sits beside the original; and that the
triage of its unread rows belongs to whoever reads it, which is what makes "this has been decided"
worth writing into a row.

`docs/process-definition.md` is read to find the section the passage should have named. Section 5's
Language subsection is where `template.yaml` holds the section ids, their order and the required
fields, and its Rendering subsection is where the `acceptance-criteria` anchor is shown; section 9 is
the Rule model and contains no occurrence of the word. It is also read for section 7's G-Test row,
which already asks for the completeness of the mapping and therefore needs no change under either
decision, and for section 8's sentence about the recommendation carrying a reason, which is the
clause the other row is about.

`docs/implementation-plan.md` is read to confirm that neither clause belongs to a step of it that
would have to move. It does not govern the content of a template section or the fields of a question,
so nothing in the plan is a finding here.

`docs/assumptions.md` is read for A90 — a reader that cannot fail is worse than none — which is the
reason the two rows keep saying they have no reader rather than being made to look answered by a
decision that no gate yet enforces.

`CLAUDE.md` is read for the three standing rules, which decide the boundary this intent is drawn
along: the documents are not the agent's to edit, so the specification wording is drafted on the
issue and the measurement is corrected here; nothing is invented, so no field, template version or
gate is touched; and the change belongs to #258 under `wp5` on a branch carrying one intent. It is
also read for the convention about replacing a paragraph being changed a second time, which the
mapping passage is, and for the convention on verifying a negative, which is why "section 9 does not
mention acceptance criteria" was taken from a `grep` over the whole file and a read of 1255–1485
rather than from recalling what section 9 is about.

The figures were measured and not recalled. Both halves were counted over all 64 P1 and all 64 P4
artifacts in the trail on 2026-10-06, by extracting each artifact's section between its anchor and
the next and testing for a numbered list item and for a cited number; the template version of every
P1 artifact was read the same way. The earlier 55-and-46 figures are of 2026-10-05 and are left
standing as a dated measurement rather than corrected, because the trail has grown between the two
dates and only part of the difference is method.
