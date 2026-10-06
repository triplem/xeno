---
intent: github.com/triplem/xeno#235
phase: 04-verification
created: "2026-10-06T15:25:19Z"
schema_version: "1.0"
runner_version: dev+8decfdc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 369c903eb7fa71d39b6ee8ddcf233c0f7946491cd5c7e612f5a45f588d6dc0d5
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
      sha256: 52afe8727c55e7be8ddc872a520f220d2c58286c4d7656e78a8a8fb7ccc008c8
    - format: other
      job: checks
      kind: other
      path: evidence/checks.txt
      produced_by: the checks over the amendment, re-run in P4
      sha256: cb35724cda4a0867d526d6d70c96a8ec5dedc371144f066b4067fae3d70742f7
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Twelve criteria, by number, all passing.

| # | what it asserts | how it is checked | result |
|---|---|---|---|
| 1 | the findings block carries `advisory`, aligned | `grep -n` over the block, beside the two existing comments | pass |
| 2 | a paragraph says what it is and does not do | a person reads it; the three facts are `pass`, `green`, and the finding in the file | pass |
| 3 | the same paragraph says why | a person reads it | pass |
| 4 | a second paragraph bounds it | a person reads it | pass |
| 5 | it is placed after the status derivation, not inside it | `git diff`, which shows it before "Finding ids are stable" and after the four-values paragraph | pass |
| 6 | Context economy says the budget finding is advisory | `git diff`, the one replaced line | pass |
| 7 | no fifth check result, and the four unchanged | `grep` for the result enumeration; the diff removes one line and it is not that one | pass |
| 8 | `drift` and section 16's ninth limitation untouched | `git diff`, which names neither region | pass |
| 9 | the commit carries no code | `git diff --stat`, which names one file outside the trail | pass |
| 10 | `gate verify`, `go test`, `go vet`, `gofmt` | run; `evidence/go-test.txt` | pass |
| 11 | no new line over 88 columns | the measure before and after, as multisets | pass |
| 12 | the agent's edit is recorded where a reader meets it | P0's problem, D-1, the commit message | pass |

<!-- xeno:section:results -->
## Results

## Sixteen lines added and one replaced

    docs/process-definition.md | 17 ++++++++++++++++-
    1 file changed, 16 insertions(+), 1 deletion(-)

The single deletion is the Context economy line that was extended, which is criterion 6. Nothing
else is removed, which is criteria 7 and 8: the result enumeration at line 394 still reads `result:
<pass|fail|pending|not-implemented>`, and neither the `drift` block nor section 16's ninth
limitation appears in the diff.

## The key sits where the block's convention puts it

    401:        advisory: true                  # absent in the ordinary case
    402:        decision:                       # absent while undecided
    408:          obligation: <open|closed>     # only with type overridden

Column 40 for all three, and all three say when the key is absent.

## `drift` is still unimplemented, measured rather than recalled

| | |
|---|---|
| `Drift` in `model.go` | 0 |
| `RunAt` and `ArtifactsHash` in `model.go` | 2 |
| a `drift` yaml key written under `internal/` or `cmd/` | 0 |
| sealed `gate.yaml` files carrying a `drift` list | 0 of 457 |

The second row is what makes the first evidence. A search returning nothing for `Drift` and nothing
for everything else would mean the search was wrong; it finds the two fields beside it in the same
struct. This is the measurement the decision rested on and it was re-run here rather than carried
from P0.

## The specification still cites no repository

    grep -cE '#[0-9]{2,3}\b|XENO-|this trail' docs/process-definition.md   → 0
    grep -cE '#[0-9]{2,3}\b' CLAUDE.md                                      → 4

The amendment carries A44's and A90's reasoning in general terms without naming either, which is
the document's habit. XENO-0262 had to remove a drafted sentence for breaking it; this one did not,
because the draft was written after that.

## Placement

`git diff` puts the two paragraphs after "artifacts, which is why it carries an obligation and
`approved` does not" and before "**Finding ids are stable across runs.**" — that is, after both
paragraphs of the status derivation rather than between them. Criterion 5, and the one departure
from the drafted wording.

## The suite

`go test ./...` exits 0 across every package, in `evidence/go-test.txt`. `go vet` is clean, `gofmt
-l` outside `vendor/` prints nothing, and `xeno gate verify` exits 0. No Go source is touched and
the specification is in no hash, so these say nothing was touched by accident and nothing more.

Over-88 lines outside table rows: 27 before and 27 after, as the same multiset of lengths.

<!-- xeno:section:gaps -->
## Gaps

**The clause has no reader and the check it describes reads nothing either.** `Advisory` is in no
struct, `result` is unchanged, `budget` still fails its check — and even once all three are fixed,
`budget` sums `f.Bytes` over `lock.Files` and no P0 `context.lock.yaml` in this trail records one
(#267). Two issues have to close before this amendment is exercised once.

**The bound is a sentence and will stay one.** No code can tell a clause that legitimately asked to
be advisory from a check somebody found inconvenient, so "it is the exception and stays one" is
enforced by whoever reads it. That is the category `docs/clause-readers.md` counts, and this
amendment adds to it.

**`docs/clause-readers.md` is one row short and this intent does not add it.** Section 5's budget
clause has no entry at all. Defensible while the clause had no mechanism; a gap now. It belongs to
the implementing intent because the column it would carry is the reader, and nothing records the
obligation outside this artifact.

**The specification now carries two keys nothing writes.** `advisory` here and
`requirements@1.1.0` from XENO-0262, both added today. The document leading its implementation by
one intent is the intended shape; two open at once with nothing bounding how long is a fact about
this session rather than a design.

**The agent edited a normative document twice today and nothing checks that it stays exceptional.**
Both are recorded in their intents and commit messages. The convention proposed for it sits in
XENO-0262's P0 learning, routed by section 10 through a merge request nobody has opened.

**What the person approved and what landed differ, for the third time.** The prose is theirs; the
placement is the document's. They are told, in a report. The P3 learning proposes showing the diff
rather than the prose next time, and that learning has the same reader as every other one.
