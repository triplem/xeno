---
intent: github.com/triplem/xeno#334
phase: 01-requirements
created: "2026-10-10T15:40:16Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f0075f71570c036897f9af29080d13162683c73017b1014f63b3f50429808933
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
Fourteen acceptance criteria for one new section of one document, and the shape of them is
the point: most are about the reading rather than about the prose, because a section of an
evaluation is a claim about somebody else's repository and the only thing that makes it
worth keeping is that the claim was checked where it says it was.

Four of the fourteen are not about the section's content at all. One asks that every AI-DLC
cell was re-read at the pinned sha and that the figures which moved since #323 say the new
number. One asks that every absence the section claims was probed with a positive control in
the same command, with both answers recorded, which is `CLAUDE.md`'s rule and #201's case.
One asks that the pin names which sha is the annotated tag object and which is the commit,
because both answer to "v2.11.0" and a reader who resolves the tag himself lands on the
other one. One asks that nothing normative moves, checked by two empty diffs.

The third row carries the one thing #334 names as undecided. Two of AI-DLC's own documents
read literally disagree about whether a sensor can refuse, and both sentences are still
there at the tag. The criterion is not that the section picks a side but that it says which
document is current, names what implements it, and states how the override is refused,
because an enforcement claim without its refusal path is the flattering half.

The fifth row is the addition section 9 did not need. On rules and learning AI-DLC is ahead
of this project, and a criterion that only permitted the other thing to come out worse would
have produced a document not worth writing.

Ten non-goals, of which three matter. Not an extension in any of its three shapes, which
#334 settles. No mechanism taken, because each of the three #323 named already has an issue
and a sensor firing on write would be a specification change before it was anything else. No
source binding in any form, because D-1 is the maintainer's and says record the gap and take
nothing.

Nine constraints, of which the binding one is that #323's reading is a lead and not a source.
It states its own limits — two large references read by grep, no upstream issues searched, no
diff of `main` against the tag, the sample's rows not read by anybody — and what this section
may assert is what was re-read at the sha. Where #323 said something the re-reading did not
reproduce, it is left out rather than softened.

One learning, about the page rather than about this intent: a Sources entry records that a
reading happened but does not oblige the next writer to redo it, so a section written from an
earlier comment can look pinned while resting on figures read somewhere else.

Three sections, of the five this template defines. No open question and no decision in this
phase; the one decision of this intent is P0's D-1.
