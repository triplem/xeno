---
intent: github.com/triplem/xeno#243
phase: 04-verification
created: "2026-10-05T12:40:28Z"
schema_version: "1.0"
runner_version: dev+f2f65d5.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5fb90ee3c7f5b5e93ed47e400a2332e988837a104b3d0a8b48febd3387893b42
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

No test is added and none could be. The change is one file of prose that nothing reads: no
gate, rule, template or Go file opens the register, and `artifacts_hash` does not cover it. A
test would be a test of this commit's words against a copy of themselves.

What stands in for it is reading, plus four counts:

| criterion | what answers it |
|---|---|
| 1, no closure claim | one grep over the file for "record is closed", "Nothing is added here" and "A77 is the last"; the only hit is inside A23's row, about a release loop |
| 2 to 5, what the opening says | reading the five paragraphs back whole, which is also what found the one deviation |
| 3, the quotation | the two quoted phrases compared against `docs/implementation-plan.md` lines 1441 and 1443 |
| 6, the fourth paragraph | reading it, and the diff hunk showing its framing sentence replaced and its state-column explanation untouched |
| 7, no row touched | `grep -c '^| A[0-9]'` before and after: 94 and 95, and the diff's four hunk headers, none inside the table but the A95 insertion |
| 8, A95 | the last row, and the only insertion in the table |
| 9, replaced whole and within 88 | an `awk` pass over prose lines outside tables and code blocks; the nine over 88 are at line 174 and beyond, all pre-existing |
| 10, `CLAUDE.md` | read; line 73 names this file and is made true rather than false, so no change, which P3's deviations records as a checked non-change |
| 11, the suite and the verdicts | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l`, `./xeno gate verify` |
| 12, one commit | the commit itself |

The suite is evidence about nothing here and is run because it is the gate's condition rather
than this change's. It would pass identically had the opening been left as it was, which is
the honest statement of what a test suite can say about a convention whose reader is a person.

<!-- xeno:section:results -->
## Results

Eleven criteria met, one pending until the commit.

1. Met. No sentence claims the register is closed. The grep's only hit is inside A23, where
   "the loop is closed" is about a release workflow.

2. Met. The first line is "This register is open", with what it holds in the same paragraph.

3. Met. The second paragraph quotes "the record is a file in the branch rather than an
   artifact under `.xeno/`" and "the same loop runs through the runner, and the hand held
   record stops", both checked against `docs/implementation-plan.md` lines 1441 and 1443, and
   says the stand-in is what stopped.

4. Met. The learning rule is its own paragraph, with section 10 named and the reason kept.

5. Met. The fifth paragraph gives durability as the test, with A74, A86 and A94 as examples.

6. Met. The fourth paragraph no longer frames the file as the pre-M0 record; it says which
   rows came from which era and that nothing distinguishes them, because the test is the same.

7. Met. 94 rows before, 95 after. The four diff hunks are at the opening, the second
   paragraph, the fourth, and the A95 insertion; none touches an existing row.

8. Met. A95 is the last row and carries the misreading, the sentence it relied on, the
   seventeen rows as evidence, the rejected alternative, and that nothing reads the register.

9. Met. Every changed paragraph was replaced whole, and reading them back is what caught "the
   rows above". No prose line of the new opening exceeds 88 columns.

10. Met, with no change. `CLAUDE.md` line 73 names this file as where assumptions and
    decisions are written down, which the correction makes true. Recorded in P3's deviations
    so that a later reader can tell a checked non-change from an unchecked one.

11. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints
    nothing, `go test ./...` is `ok` across all eighteen packages that have tests, and
    `./xeno gate verify` exits 0 over 367 verdicts: the 363 that existed are intact and the
    four are this intent's own judged phases.

12. Pending. The commit comes after this phase is judged, because this phase's artifacts are
    part of what it commits.

<!-- xeno:section:gaps -->
## Gaps

Nothing reads the register, and this change adds claims to it rather than readers. The opening
now asserts that the file is open, what belongs in it, that a learning never arrives here, and
that a row outlives its intent. Not one of those is checked by anything. Two could be in
principle — whether a `learning.yaml` proposal ended up as a row, whether a row is cited by an
intent other than the one that wrote it — and neither is, so the file's claim surface grew and
its enforcement stayed at zero. A90's finding is that this is the shape of the problem rather
than an oversight, and A95 says so in the row itself.

The banner survived twelve contradictions because nothing compared the file against practice,
and nothing does now either. What changed is that the stated rule and the practice agree, so
the next drift will be practice moving away from a correct rule rather than a wrong rule
standing. That is a better failure but it is the same mechanism, and the thing that would
catch it is a person reading the file, which is what failed twelve times.

The citation that caused this is not guarded. #174 cited section 4 and misread it, and the new
opening quotes the sentence so the inference can be checked — but nothing stops the next
document from citing a normative sentence it has not read. P0's learning proposes the
convention; no rule or gate expresses it, and a `checked` rule could not, because comparing a
claim about a sentence against the sentence is the parsing problem section 7's budget rules
out.

The seventeen rows are not individually revisited. The correction says they belonged here, and
whether each earns a row by the durability test it now states is seventeen questions this
intent declined. It is possible that some do not, in which case the file states a test that
its own history partly fails, and nothing will report that.

Whether `CLAUDE.md` should say more is left open. It names the file and now agrees with it,
which was the criterion; whether it should also carry the durability test, so that the
distinction is read before a change rather than after, is the kind of addition section 10
routes through review and not a thing this intent decides.
