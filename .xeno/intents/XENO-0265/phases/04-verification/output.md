---
intent: github.com/triplem/xeno#235
phase: 04-verification
created: "2026-10-06T18:04:04Z"
schema_version: "1.0"
runner_version: dev+5276f4b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cb2934754c063a5e8f6e78a9de4533df4eb9af8aca1f6de5e0b0961325a952e0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - format: other
      job: go-test
      kind: test-report
      path: evidence/go-test.txt
      produced_by: go test ./...
      result: pass
      sha256: 6297ceef64f0d218c0c38319c15dd42e45b4c5193db31c4a1b66964503a114d9
    - format: other
      job: checks
      kind: other
      path: evidence/checks.txt
      produced_by: the clause end to end, the five mutations and the counts
      sha256: d4c674a19476cd391234049ab79e36f6fb029a9e9ee430cbb68843ae36b9d971
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Thirteen criteria, by number, all passing. This is the first P4 in the trail whose mapping
`gates.mappingComplete` reads, because its P1 is the first at `requirements@1.1.0`.

| # | what it asserts | how it is checked | result |
|---|---|---|---|
| 1 | `Finding` carries `Advisory` with `omitempty` | `git diff`; 0 of 469 sealed `gate.yaml` files carry the key | pass |
| 2 | `result` fails only on a non-advisory finding | `TestACheckOfOnlyAdvisoryFindingsPasses`; mutation | pass |
| 3 | one advisory and one ordinary still fails | `TestACheckWithOneOrdinaryFindingStillFails`; mutation | pass |
| 4 | the findings are in `gate.yaml` with id, cause, remedy | the scratch run's output, under a passing G-Schema | pass |
| 5 | `budget` marks both of its findings | `TestABudgetOverrunLeavesGSchemaPassing`; `grep` counts 2 | pass |
| 6 | nothing else sets the field | `TestOnlyTheBudgetClauseWritesTheAdvisoryField`; 0 outside the two files | pass |
| 7 | **the phase is green and the next phase starts** | the scratch run; `TestAPhaseWhoseOnlyFindingIsAdvisoryIsGreen` | pass |
| 8 | an advisory finding is still decidable | `gate approve` on the scratch run gives `approved`; `TestAnAdvisoryFindingNeitherHidesNorBlocksTheOthers` | pass |
| 9 | the four results and five statuses are untouched | `grep` over the specification; `git diff` names no enumeration | pass |
| 10 | `docs/clause-readers.md` gains the budget row | a person reads the table; 41 rows counted | pass |
| 11 | the bound is in the code where the field is | a person reads the comment on `Advisory` | pass |
| 12 | `go test`, `go vet`, `gofmt`, `gate verify` at the same count | run; `evidence/go-test.txt` | pass |
| 13 | every new assertion confirmed able to fail | five mutations, five failures | pass |

Criterion 7 is the one that earned its place. It is written as a behaviour — the phase green *and*
the next phase starting — and it is what found that `result` alone did not satisfy the clause.

<!-- xeno:section:results -->
## Results

## The clause end to end, and the half that was missing

On a scratch copy, a phase whose only finding is a budget overrun. With `result` changed and `Status`
not:

    00-intake: red
      G-Schema       pass
          F-2af1d8 ...: the recorded context is 9999 bytes and the budget is 10

Every check passing and the phase red, because `Status` counts undecided findings and an advisory one
is undecided for ever. With both changed:

    00-intake: green
    xeno phase start --phase 01  -> started
    xeno gate approve F-2af1d8   -> 00-intake: approved

Green; the next phase starts, which is what "not a red gate in the sense of stopping work" means in
this runner; and the finding is still decidable, which section 5 asks for in the same sentence.

The fixture writes the lock's `files` by hand. The budget check has no input in this repository — no
P0 `context.lock.yaml` records a `files` list, which is #267 — so the state cannot be reached by
running the runner, and saying so is part of the evidence rather than a footnote.

## Five mutations, five failures

| mutation | what failed |
|---|---|
| `result` counting findings again | `TestACheckOfOnlyAdvisoryFindingsPasses`, `TestABudgetOverrunLeavesGSchemaPassing` |
| `result` ignoring every finding | `TestACheckWithOneOrdinaryFindingStillFails`, and the empty and ordinary cases |
| `budget` no longer marking its own | `TestABudgetOverrunLeavesGSchemaPassing`, `TestOnlyTheBudgetClauseWritesTheAdvisoryField` |
| a second check reaching for the field | `TestOnlyTheBudgetClauseWritesTheAdvisoryField`, "3 ... want the budget clause's 2" |
| `Status` counting advisory findings | `TestAPhaseWhoseOnlyFindingIsAdvisoryIsGreen` |

The fourth is the bound. It is the only test in this intent that reads a source file, and it exists
because section 5's "nothing else writes the field" has no other reader.

## Nothing else writes it

    advisory(finding( in gates.go: 2
    Advisory = true in gates.go:   1
    Advisory anywhere else:        0

## Nothing in the trail moved

    gate verify: verified 469 verdicts
    sealed gate.yaml files carrying an advisory key: 0

`omitempty` is what buys that: the field serialises only where it is set, so a rewritten verdict is
byte identical to the one before it.

The four check results and the five phase statuses are untouched — `grep` finds both enumerations in
section 5 unchanged, and `git diff` names neither.

## Why the contradiction survived, which is worth its own line

    calls to budget( in budget_test.go: 9
    assertions on a check Result there: 0

Nine tests over the budget check, none of them asserting what the check result should be — the one
thing section 5 is explicit about. The behaviour that contradicted the specification had no reader in
the suite either, which is why the new tests go through `result` and `Status` rather than through
`budget`.

## The suite

`go test ./...` exits 0 across every package, in `evidence/go-test.txt`. `go vet` is clean, `gofmt
-l` outside `vendor/` prints nothing, `xeno gate verify` exits 0.

<!-- xeno:section:gaps -->
## Gaps

**The check still reads nothing at P0.** `budget` sums `f.Bytes` over `lock.Files` and no P0
`context.lock.yaml` in this trail records one (#267). So the clause is honoured and the finding it
describes cannot arise from a real scope in this repository: the end-to-end evidence rests on a lock
written by hand. Both issues have to close before an advisory finding is produced by the runner
once.

**"Finding" now means two things.** A finding in `gate.yaml` may be one the phase was not failed
for, and every reader of a verdict has that distinction to hold. It is the cost section 5's clause
was always going to charge, and it is charged to the reader rather than to the writer.

**The bound is a sentence and a string count.** No code can tell a clause that legitimately asked
for the field from a check somebody found inconvenient. `TestOnlyTheBudgetClauseWritesTheAdvisoryField`
reads `gates.go` and counts, so it fails if the helper is renamed or a call is reformatted — a false
red for a true property — and it is in because the alternative was no reader at all.

**An advisory finding is undecided for ever and nothing asks about it.** `Status` skips it, so it
never becomes an obligation and never appears in the next-step suggestion. A reader who wants to know
whether anybody looked at a budget overrun has the artifact and nothing else.

**No gate reads whether the field is used honestly.** A future check could mark its findings advisory
and the test would catch the count, but a project with its own external gate reporting through
section 14 can write `advisory: true` into a check of its own and nothing would notice. The field is
in the artifact schema, so it is available to anything that writes one.

**`docs/clause-readers.md`'s count moved twice today and nothing watches it.** Forty in XENO-0264,
forty-one here. It is a number in prose beside the rows it describes; the learning proposing that it
be derived is recorded and has the same reader as every other learning.
