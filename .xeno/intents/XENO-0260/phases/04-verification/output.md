---
intent: github.com/triplem/xeno#258
phase: 04-verification
created: "2026-10-06T11:41:02Z"
schema_version: "1.0"
runner_version: dev+7885661.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1d69edac7d88f40eb71643e1a0586c70e22cad6c6f5841050471c199dded6121
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - format: other
      job: go-test
      kind: test-report
      path: evidence/go-test.txt
      produced_by: go test ./...
      result: pass
      sha256: 0dd072684024295d94c06da885d2ae070e0b27633f81b8f08f089a042a3977f3
    - format: other
      job: figures
      kind: other
      path: evidence/figures.txt
      produced_by: both counts and the section 9 search, re-run in P4
      sha256: 3aa655a39b7119b3dfc7b22f71cd61cb26454f8c51fc39204ad5692a7cc6fcae
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Nine criteria, by number. Four are read by a person, which the last column says rather than hides.

| # | what it asserts | how it is checked | result |
|---|---|---|---|
| 1 | the passage names section 5 | a person reads the replaced paragraph; `grep -n "section 9"` over the document | pass |
| 2 | both halves are stated with date and method | both counts re-run over all P1 and all P4 artifacts; `evidence/figures.txt` | pass |
| 3 | the 2026-10-05 figures stay | a person reads the paragraph; the sentence naming them is present | pass |
| 4 | the section 8 row gains an explanation | a person reads it | pass |
| 5 | both rows say they are decided and what they wait on | a person reads both paragraphs | pass |
| 6 | neither reader column changes | `git diff` over the clause table | pass |
| 7 | no normative document, field, gate or template | `git diff --stat`, which names one file | pass |
| 8 | `gate verify`, `go test`, `go vet`, `gofmt` | run; `evidence/go-test.txt` | pass |
| 9 | the document wraps at 88 outside tables | `awk 'length>88'` over the document, before and after the change | pass |

Criterion 2 is the one that moved under the phase, and the results section says how.

<!-- xeno:section:results -->
## Results

## The figures, re-run

Both counts were re-run from the trail in this phase rather than carried from P0, by extracting each
artifact's section between its anchor and the next:

    P1 artifacts: 65        P1 with a numbered list of criteria: 20
    P4 artifacts: 64        P4 whose mapping cites a number: 7
    distinct template declarations in P1: 65 × requirements@1.0.0

**This phase's figures disagree with the document's, and the document is right.** 65 and 20, against
the 64 and 19 written in P0 and P3. The difference is this intent: its own P1 artifact entered the
trail between the measurement and the verification, and it numbers its criteria, so it is counted.
The P4 figure is still 64 because this phase's artifact is not finished as the count runs, and will
be 65 and 8 once it is.

A count of a trail taken from inside that trail moves while it is being taken. The document now says
what its 64s are — the trail as it stood before the intent that measured it — and records that this
intent's own artifacts move both figures the right way, which is the forward-only anchor working one
intent at a time rather than an inconvenience. That is the correction criterion 2 required and it
was made in this phase, which is what the phase is for.

**The P1 count was taken twice with two patterns.** A loose one matching a numbered list item or a
numbered table row, and a strict one matching only a list item. Both returned the same split, and
the strict one reports zero artifacts numbering by table row, so no artifact is counted by the loose
pattern that the strict one misses.

## Section 9

`grep -n "acceptance"` over `docs/process-definition.md` returns 572, 692, 710, 713, 942 and 1775.
Section 9 runs from 1255 to 1485 and contains none of them. The search returning five hits elsewhere
is what makes the absence evidence rather than a pattern that matches nothing — #263's convention,
and the reason the range was also read directly.

`grep -n "section 9"` over `docs/clause-readers.md` returns nothing.

## What a person read

Criteria 1, 3, 4 and 5 are answered by reading. The replaced paragraph names section 5 in its second
sentence and says what section 9 is; the 2026-10-05 figures are present with their date and what they
counted; the section 8 row has a paragraph where it had none; both paragraphs end on what the clause
waits on and on the reader column being unchanged.

## The table, and the suite

`git diff` over the clause table is empty: thirty-eight rows, both reader columns as they were.
`git diff --stat` names `docs/clause-readers.md` and nothing else outside the trail. `go test ./...`
exits 0 in `evidence/go-test.txt`, `go vet` is clean, `gofmt -l` outside `vendor/` prints nothing, and
`xeno gate verify` reports **438 verdicts verified** at exit 0.

`awk 'length>88'` outside table rows returns two lines, 73 and 75. Both are in the paragraph above
the replacement; stashing the change and re-running returns the same two, which is how they were
established as pre-existing rather than introduced.

<!-- xeno:section:gaps -->
## Gaps

**The count will be wrong again tomorrow, and the document says so only indirectly.** Its figures are
dated and the paragraph names the population, which is the most the pass's own rule about ageing
allows. Nothing recomputes them, and the next intent to number its criteria makes them stale by one
without anything saying so. That is the same shape as every other figure in this document and it is
not repaired here.

**The two paragraphs record decisions that no gate enforces.** A reader who takes "decided in #258"
for "implemented" would be wrong, which is why each paragraph ends on the reader column being
unchanged and on what it waits on. The sentence is the only thing preventing that reading, and a
sentence is a person-reader.

**The passage will be replaced a third time.** Each specification commit, and the code behind it,
moves one reader column and makes one of these paragraphs historical. Writing it to be replaced is
why it was replaced whole now rather than corrected in place, and it is still an edit somebody has to
remember to make in two separate intents.

**The section number is corrected here and stands wrong in sealed artifacts.** XENO-0258's phases and
the earlier ones that restated it carry section 9, and section 11 means they stay. A reader of those
meets the wrong number with nothing beside it — the same residual risk XENO-0254 recorded about the
gitignore claim, by the same mechanism, and the third time this trail has recorded it.

**One of the nine numbered P1 artifacts already fails a check by number**, from the 2026-10-05
measurement. This intent neither re-found nor repaired it, and the check that would read it waits on
the section 5 commit. It is carried forward in the document's own figures and nowhere else.

**Nothing checks that the document's figures and its prose agree.** The table says 64 and 19, the
paragraph says what the 64s are, and the next person to update one has to update the other. This is
the same gap the audit baseline's comment has, recorded in XENO-0259's residual risk a few hours
earlier, and it is becoming a shape rather than an instance.
