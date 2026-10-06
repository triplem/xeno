---
intent: github.com/triplem/xeno#267
phase: 03-implementation
created: "2026-10-06T20:05:16Z"
schema_version: "1.0"
runner_version: dev+3f1fab3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: dfab8198d8b029fc105ce5936462d8e66740886106aca8188a438f44c94eaa5a
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

**`docs/process-definition.md`, three paragraphs and one comment**, committed first and on its
own, as `CLAUDE.md`'s first standing rule requires. Section 5's lock subsection gains **At P0 it
names no files**, after the two paragraphs about what the lock records and that it is not
refreshed. The schema block's `files` line gains `# empty at P0, see below`. Section 5's budget
clause gains a paragraph naming the comparison it makes. Section 7's staleness subsection gains a
paragraph after the two limits, saying it is not a third one.

**`internal/runner/runner.go`, `informationBase`'s doc comment.** The sentence citing #217 said
an empty list is "a phase started before that rule rather than a project opting out", with no
phase named. It now names P1 on, where that is true, and a second paragraph says the P0 case is
not an exception: `ScopeSet` refuses before the phase has started, so at the moment `Start`
resolves the base there is nothing to resolve, and the lock is never written again.

**`internal/gates/gates.go`, `staleReads`'s doc comment.** The same sentence, the same
correction, placed after the two limits it already quotes from section 7 and before the paragraph
about an absent commit range. It says the two limits are trades against noise and P0's exclusion
is not, which is the distinction the specification paragraph makes.

**`internal/gates/gates.go`, `budget`'s doc comment.** Not in P1's criteria, and added because
the function's own text argues against the new clause without meaning to: it already explains
that an entry with no recorded size is "silence rather than a context of zero bytes", which is
the distinction A74 draws, and says nothing about an entry count of zero being the ordinary state
at P0. A reader reaching the `recorded` flag with #267 in mind would conclude the flag is what
suppresses the check there, when the empty list is.

**`docs/clause-readers.md`, a paragraph and one count.** See the deviations below.

No test. Nothing changes behaviour; `go test ./...` passes as it stood.

<!-- xeno:section:deviations -->
## Deviations from the design

**Criterion 12 asked for three rows in `docs/clause-readers.md` and got a paragraph instead.**
P1 said the three clauses "carry what reads each, which for all three is a reader of the document
and no code". That is the wrong column of the wrong table. The enumerated table is **tool
requirements**, which the document defines as clauses "the runner, a gate or CI can satisfy or
violate"; none of the three can be violated by code, because each states a consequence of the
order in which `phase start` and `scope set` run. In the pass's four kinds they are explanations,
which are counted and not enumerated, so the explanation count moves from 126 to 129 and a
paragraph says which clauses and why.

The paragraph also records the one of the three that is reader-shaped and is deliberately left
without a reader: "An intake's `files` is empty" becomes false the day either of #267's candidates
is adopted, and nothing would notice. The sentence is phrased so that adopting one means deleting
it, which makes the specification commit that opens the change visible as a deletion.

Three rows would have claimed the opposite — that something in the repository fails if the
clauses are broken — and the table is the one document in this repository whose whole value is
that its reader column is honest.

**`budget`'s doc comment was added and is not in P1's criteria.** Criteria 10 and 11 name the two
comments that assert the wrong reason; `budget`'s asserts nothing wrong. It was added because its
existing text leads a reader the wrong way: it explains at length that an entry with no recorded
size is silence rather than zero bytes, which is A74's distinction, so somebody reaching the
`recorded` flag with #267 in mind concludes the flag is what suppresses the check at P0. It is the
empty list. Three sentences there cost nothing and the phase is where the comment belongs.

**Nothing else.** The three paragraphs are as approved except for the three wordings recorded in
the specification commit's message, each of which was found by reading the paragraph in the file
rather than in the draft — the budget paragraph's final sentence also had to be rewrapped after
the edit, which is the fourth time a draft needed adjusting on contact with the document.
