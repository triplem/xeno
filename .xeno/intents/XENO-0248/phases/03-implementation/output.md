---
intent: github.com/triplem/xeno#243
phase: 03-implementation
created: "2026-10-05T12:38:08Z"
schema_version: "1.0"
runner_version: dev+f2f65d5
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c9507277dc24e33b4e3c7bb25d67c8ba0e37d855a129b5b3f4f1f5debc7633ee
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

One file, four hunks, 39 insertions against 17 deletions.

The two paragraphs that closed the record are replaced by five. The first says the register
is open and what it holds. The second says what closed at M0, quoting section 4's own two
phrases, and that this file is not that stand-in. The third says the file claimed otherwise
from #174 until #243, names the seventeen rows and twelve commits, and says why the sentence
is quoted rather than summarised. The fourth keeps the learning rule almost verbatim. The
fifth gives the test a row meets: the fact outlives the intent that found it, with A74, A86
and A94 as the examples the rows below have in common.

The fourth paragraph's framing sentence is replaced. It said this was the record kept before
M0; it now says the rows up to A77 were written before M0 and those after it by intents that
ran six phases, and that nothing about a row distinguishes them because the test is the same.
Its explanation of the state column, `approved`, `accepted until` and an `open` row is
untouched.

The sentence "Every entry is a place where the process definition or the implementation plan
says what, and the how had to be chosen" moved from that paragraph into the first, where it
now says what the register holds rather than what the closed record held.

A95 is added, the last row, recording the correction: the misreading and the sentence it
relied on, the seventeen rows as evidence about the banner rather than about the twelve
commits, the rejected alternative and why, durability rather than significance as the test,
and that nothing reads the register so the convention's reader is a person.

The paragraph saying every row stays with its state column is unchanged, and no existing row
is edited, renumbered or deleted: 94 rows before, 95 after.

`CLAUDE.md` is read and not changed, which criterion 10 expected. Its line 73 says assumptions
and decisions taken while building the core are in this file, which the banner contradicted
and the correction makes true.

<!-- xeno:section:deviations -->
## Deviations from the design

One deviation from the design, found by reading the new opening back as a paragraph rather
than as a diff. The fifth paragraph said the durability test "is what the rows above have in
common", and the rows are below it: the table follows the prose. It is the error the
conventions predict for a paragraph assembled from parts, and the method they ask for is what
caught it, two minutes after the edit and before the phase was judged.

One addition beyond the design, which named four paragraphs and produced five. The design put
the claim, the M0 distinction, the learning rule and the durability test in one paragraph each
and left the history of the banner inside the second. Separating it makes the second paragraph
about what section 4 says and the third about what #174 made of it, which is the distinction
the whole correction rests on; a reader who cannot tell the plan's sentence from the inference
drawn off it is the reader this change exists for.

No deviation on `CLAUDE.md`. Criterion 10 said to read it and change it only if it disagreed,
and the expected outcome was no change; that is the outcome. Recorded because a criterion
whose expected result is "nothing" is one a later reader cannot tell was checked.

Nothing else departs. No existing row is touched, the normative documents are untouched, the
row count goes 94 to 95, and the file is still read by nothing.
