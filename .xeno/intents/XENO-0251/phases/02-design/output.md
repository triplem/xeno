---
intent: github.com/triplem/xeno#212
phase: 02-design
created: "2026-10-05T15:49:26Z"
schema_version: "1.0"
runner_version: dev+09e2aa6
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 517ef183b4bf2cc0e1f13281732d88611d05a848f1619d7002e6a244bae7acf1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

`declaredResults(c, kind, label)` is the shared body and `build` and `testReport` are two
callers of it. The two gates differ by a constant and a word in the finding, so the comparison,
the pending-and-attached path and the `pending` result live once. `BuildKind`'s own comment is
the argument: the gate matched no conformant declaration for as long as the spelling was
written into the comparison, and two copies of a comparison is the same bet taken twice.

`TestKind = "test-report"` sits beside `BuildKind`, with the comment pointing at that one rather
than repeating it. Both are the spellings section 4 fixes and G-Schema judges against
`model.EvidenceKinds`, so the constant is the gate's end of a set defined elsewhere.

The finding says "test report <job> did not succeed" and the next step says to fix the suite and
let the pipeline produce a new result. It mirrors G-Build's wording because the two failures are
the same shape from the reader's side: something the pipeline ran did not pass, and the repair
is not in the artifact.

G-Test reports `pending` when a declared report is awaiting its artifact and nothing else failed,
which is G-Build's rule and A60's: a pending item owes no result. This is the decision that
makes the trail survive the change, and it is inherited rather than chosen.

The audit row names both halves in one line: the reader of the result half, and that the mapping
half has none. One row rather than two, because section 7 gives G-Test one clause and splitting
it in the audit would imply the document splits it.

The count in the audit's table moves 35 to 36 and a sentence says this row arrived after the
pass, beside the sentence already there about the row that arrived after the pass. The file's own
precedent, and the reason it exists is that a pass which absorbs its own additions stops being a
measurement.

The mapping half's issue carries the numbers rather than the conclusion. "46 of 55 P1 artifacts
have no numbered criteria" is what makes the difficulty concrete; "the mapping half is hard" is
what the issue would say without it, and #212 is the example of what that costs — it asserted
the opposite in good faith and nothing in it was checkable.

<!-- xeno:section:alternatives -->
## Alternatives

Leaving G-Test as `not-implemented` and closing #212 with the measurement was the option the
maintainer weighed this against. It is the one with nothing to criticise: the gate says honestly
that it judged nothing, and section 5 has that state precisely so a gap surfaces instead of a
green verdict meaning less than it appears to. It was rejected because the result half is a real
check that costs nothing and catches a real failure — a declared test report that did not pass —
and a gate that catches nothing at all is what the issue was filed about.

Implementing both halves was the third option and is not viable as one intent. The mapping check
would fire across 46 of 55 P1 artifacts, which is not #229's single sealed question but most of
the trail, and identifying a criterion mechanically needs a numbering convention that is a
specification change before it is code.

A `review` rule for the mapping half was considered: put it in the rule set, have a person answer
it in the P5 checklist, and the clause gets the only reader it can honestly have. Rejected for
this intent rather than rejected outright — it is a good idea and it belongs to the issue that
decides the mapping half, because a rule asserting completeness once per intent is a different
instrument from a gate asserting it per criterion and the choice between them is the issue's.

Copying `build` into a second function was the obvious implementation and is what the probe did.
Rejected for the reason `BuildKind` records one level down: the gate read the wrong spelling for
as long as the comparison carried it inline, and two bodies that must agree about pending,
attachment and the `pending` result will eventually not.

Keying the finding on the declaration's own `result` was considered and is what the first audit
script did. It is wrong, and the probe is what showed it: seventeen declarations carry no result
and are pending with attached results, so a gate reading the declaration alone would fail
seventeen sealed artifacts while believing it was reading section 7's clause.

Reporting the two halves as two results in one check was considered and has nowhere to go.
Section 5 gives a check one result from a fixed set, and inventing a second field would be an
addition to the artifact shape.

<!-- xeno:section:impact -->
## Impact

One file of code and one document. `internal/gates/gates.go` gains `TestKind`,
`declaredResults` and `testReport`, and `build` becomes a caller; the table row changes from
`notImplemented` to `testReport`. `docs/clause-readers.md` gains a row and a count. Tests join
`internal/gates`.

Every P4 verdict written from here reports G-Test as `pass`, `fail` or `pending` where it said
`not-implemented`. Nothing already written changes: the probe ran the real comparison over the
whole trail at exit 0 over 381 verdicts, which is criterion 5 and was measured before this
intent was planned rather than after it was built.

What is gained is a reader for half of section 7's row. A declared test report with
`result: fail` is now a red P4 where it was a green one, which is the failure the gate exists
for and which no artifact in this trail has yet contained.

What is lost is the honesty of `not-implemented`. G-Test will report `pass` having judged the
result half and not the mapping half, so a reader who takes a green G-Test for "section 7's row
is satisfied" is now wrong in a way they were not before — before, the gate said it had not
judged. That is the real cost of this change, it is not mitigable in the verdict, and the audit
row is the whole of the mitigation: `docs/clause-readers.md` is the one document whose purpose
is holding exactly this, and its count moving to 36 is what a later pass reads.

`docs/clause-readers.md` becomes a document with two late rows out of 36, both marked. That is
the second time its dated-pass framing has had to absorb an addition, which is worth noticing: a
measurement that keeps being amended is on its way to being a register, and nothing has decided
whether it should be.

The mapping half stays unread and now has company in a different sense: `acceptance-criteria`
and `test-mapping` go on being required sections written in every intent for a reader that
exists for one of them and not the other. The templates' cost is unchanged and its justification
is now half true.
