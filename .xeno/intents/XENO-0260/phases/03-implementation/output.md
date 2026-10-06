---
intent: github.com/triplem/xeno#258
phase: 03-implementation
created: "2026-10-06T11:37:40Z"
schema_version: "1.0"
runner_version: dev+7885661.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fa9cd6c7b458b65e6efa299035158b8c18f7e4be31d6ff22a0baa6fe43c9dd67
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

One file, `docs/clause-readers.md`, and nothing outside the trail.

**The paragraph that named section 9 is replaced.** It now names section 5, says which subsections —
Language and Rendering, where `template.yaml` holds the section ids, their order and the required
fields — and says what section 9 actually is, with the search that settles it: "acceptance" occurs at
572 and 692 in section 5, 710 and 713 in section 6's phase table, 942 in section 7's G-Test row, and
nowhere in section 9. It also records that #258 and its first comment had copied the wrong number,
because the fault was propagation and not a typo, and a reader who sees only the correction cannot
tell that the next writer took it from here.

**Both halves of the measurement are now stated, as a table with its populations.** 64 P1 artifacts,
19 numbering their criteria as a list; 64 P4 artifacts, 7 citing a criterion by number. The cost of
applying a check to the trail is given as 57 verdicts, with the reason it is the P4 figure: a
completeness check is judged at P4.

**The 2026-10-05 figures stay, with their date and with what they counted.** 55 P1 artifacts, 46
without numbered criteria, one of the nine already failing a check by number. The paragraph says why
they stay: nine intents arrived between the two dates, so part of the difference is growth and part is
population, and a reader shown one number cannot separate them.

**A paragraph records what the mapping clause was decided to and what it waits on.** The template
version as the anchor, `requirements@1.1.0` and `verification@1.1.0`, the check judging only artifacts
declaring 1.1.0 or later, and the fact that makes it cost nothing: every P1 artifact already declares
`requirements@1.0.0` and `template` is already read by `model.Template`, so G-Policy's `judgedUnder`
has an equivalent after all. It says the reader column still reads `nothing` and why — A90, and that
a decision is not a reader.

**A paragraph is added for the section 8 row, which had none.** It states the clause, the four fields
`model.Option` carries, why a reason can only hide inside one of them, both homes #247 argued, and
what #258 settled: `reason` on `Option`, required by `QuestionAsked` wherever `recommended: true`,
with the drift argument that decided it and `ChecklistEntry.Note` as the precedent for a
conditionally required field. It too says the row goes on reporting no field until one exists.

## What was not touched

The clause table and both reader columns. The four-kinds counts, the thirty-eight row count, the
architectural-property table, the findings section and the closing section. The first paragraph of
the mapping explanation, about `examples/rules/mapping-is-complete.yaml`, which was correct and is
where the person-reader route still lives.

<!-- xeno:section:deviations -->
## Deviations from the design

**The replacement was written over 88 columns and had to be reflowed.** Eleven lines between 89 and
90 columns after the paragraphs were written by hand. They were reflowed with a wrapper over the
whole block rather than trimmed line by line, which is what replacing a paragraph whole means in
practice; the two lines still over 88 are in the paragraph above the replacement and pre-date this
intent, which was checked by stashing the change and re-running the measure.

**The reflow split a code span across a line break and the sentence was rewritten.** `grep -n
"acceptance"` came out as a span broken over two lines. CommonMark joins it back, so it would have
rendered correctly, and it read as a mistake in the file — which for a document whose whole subject
is what a later reader can tell is the wrong trade. The sentence now says "searching the
specification for" and names no command, which costs the reader the exact invocation and buys a
sentence that reads as written rather than as wrapped.

**No deviation on the figures.** Both halves were counted over all 64 P1 and all 64 P4 artifacts by
extracting each artifact's section between its anchor and the next, and the template version of every
P1 artifact was read the same way. The first count of the P1 half also matched numbered table rows
and returned 19 and 45; the stricter count, matching only a numbered list item, returned the same 19
and 45, and the loose one was discarded. Both were run, because a looser pattern and a correct one
agreeing is worth more than either alone and they do not always.

**The claim that section 9 says nothing about acceptance criteria was verified with the thing
present.** #263's convention: a search returning nothing is the same answer for an absent term and an
absent section. So the search was run over the whole document and returned five hits elsewhere, and
section 9's range was read directly. A search that had returned nothing anywhere would have meant the
pattern was wrong.
