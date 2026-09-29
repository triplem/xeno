---
intent: github.com/triplem/xeno#136
phase: 04-verification
created: "2026-09-29T18:10:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4b87535.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a0d67fcd97c28381eb9a81f9e89157f3fbf9d6690d2ef690cec4472c23f84455
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
evidence:
  - kind: test-report
    job: go-test
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | What proves it |
|---|---|
| AC1 the listing's heading | the transcript: `CREATED  INTENT  STATE  PHASE` |
| AC2 the one-intent form's heading | the transcript: `PHASE  STATE  VERDICT` |
| AC3 nothing else changed | the transcript prints the heading and the widest row from the same format string, and the widths are XENO-0207's; the diff outside `cmd/` and the records is empty |
| AC4 the alignment holds | `TestTheHeadingLinesUpWithTheWidestRow` and `TestThePhaseHeadingLinesUpToo`, the same two tests asserting the same property against the same widest rows |
| AC5 no other output changed case | the diff: two literals in `cmd/xeno/main.go` and four strings in its test |
| AC6 the reversal is recorded, not edited | XENO-0207's phases are not in the diff; this intent's intake and design name the decision and the rule that made a new intent the route |
| AC7 everything green stays green | `go test ./...`, `gofmt`, `go vet`, `gate verify` over 127 |

Seven rows. AC3 and AC6 are the two that matter, and neither is about the case: one says nothing moved
with the edit, the other says nothing sealed was rewritten to make it.

<!-- xeno:section:results -->
## Results

The three tests pass, the suite passes, `gofmt` and `go vet` are clean, and `gate
verify` matches 127 verdicts.

Both headings print in upper case at XENO-0207's widths, which is the whole of the
intended change. Upper case is no wider than lower, so the columns sized for the data
still fit their labels, and the transcript shows the heading and the widest row rendered
from the same format string for a reader to compare.

Nothing outside `cmd/` and this intent's own records is in the diff, which is AC3 and
AC5 together: two literals, one comment and four strings in a test.

XENO-0207's phases are untouched. That is AC6 and it is the reason this is a separate
intent rather than an edit: its design carries a verdict, and a verdict over a file that
no longer says what was judged would be worse than a reversal recorded elsewhere.

Read rather than executed: the comment that now says which convention governs a column
label, which is the part of this change that will still matter when the case is
forgotten.

<!-- xeno:section:gaps -->
## Gaps

**The case is now right and the rule is still only in a comment.** `CLAUDE.md` says a
heading names its section in words, and nothing in it distinguishes prose from tabular
output. The next table this project prints will face the same question with the same
citation available, and the answer will be in a comment in `cmd/xeno` rather than where
conventions are kept. The learning proposes moving it; this intent does not touch
`CLAUDE.md`.

**Nothing tests the case.** The two alignment tests would pass with lower-case headings,
since they assert that a heading word sits over its column and not what the word looks
like. A test for the case itself would assert a literal against a literal, which is a
test of the test. So the case is verified by the transcript and by reading, as the
wording of the notice was in XENO-0207.

**Seventeen sections for two literals.** The record is about thirty times the size of a
diff of two lines. That is the honest cost of section 11's rule here and it is the
clearest instance #117 has, which is where the figure goes.

**The reason XENO-0207 gave is still in the repository, still sealed, and still wrong.**
That is what section 11 asks for: the intake and design of this intent name it, and a
reader who finds the old design first will read a decision that was reversed without the
file saying so. The reference runs one way.

**The evidence is a local run**, thirteenth in a row.
