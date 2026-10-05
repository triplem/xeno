---
intent: github.com/triplem/xeno#212
phase: 04-verification
created: "2026-10-05T16:12:35Z"
schema_version: "1.0"
runner_version: dev+97d07da.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f5912bbe39f793a271e5e4b651f6011865b2c1b5cb4041777388583a8aee0ab1
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
      sha256: fe3d2dc71262bd84cd327296b3a6f38f18e156432f734361be506dcd8ec56c2c
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Five tests in `internal/gates/evidence_test.go`, plus one declaration in this phase so that the
gate judges its own intent.

| criterion | what answers it |
|---|---|
| 1, G-Test is implemented at P4 | `TestGTestIsImplementedAtP4`, over the table rather than over a call: the row's phase is 4 and the function does not report `not-implemented` |
| 2, pass, fail and pending | `TestGTestJudgesTheDeclaredTestResult` for the first two, `TestAPendingTestReportIsPendingAndNotAFailure` for the third |
| 3, the body is shared | read, plus `TestGTestAndGBuildReadTheirOwnKind`, which fails for any implementation where one gate reads the other's kind |
| 4, `TestKind` named and pinned | the `declaration` struct and `TestEveryDocumentedKindIsAcceptedAndOthersAreNot`'s neighbour assertions, which compare the constant against `testdata/evidence-declaration.yaml` |
| 5, the trail | `./xeno gate verify`, at exit 0 over 384 verdicts |
| 6, the audit row and its count | read: two rows, the count at 37, and 37 rows in the table counted by script |
| 7, the mapping half's issue | #250 |
| 8, the eight cases | the five tests above, `TestNeitherGateFiresWithoutADeclaration` for an artifact declaring neither, and the two directions of criterion 3 |
| 9, the suite | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l` |
| 10, one commit | the commit itself |

**This phase declares a test report, which is the check no unit test can make.** `go test ./...`
was run, its output written to `evidence/go-test.txt`, and declared with
`xeno evidence declare --kind test-report --job go-test --result pass`. So G-Test judges this
intent's own P4 against a real declaration rather than against a fixture, and the gate that had
never judged anything judges the artifact that implements it. That mirrors #242's P5 writing its
own checklist with the command it added.

`TestKind` is pinned to `testdata/evidence-declaration.yaml` rather than asserted inline, which
is the one test here that is about the document rather than about the code. `BuildKind` exists
because the gate compared against a spelling nothing checked and matched no conformant
declaration the whole time; the new constant inherits the same guard rather than the same
defect.

The pending case is the test that matters most and was wrong twice before it was right. An audit
script read the declaration's own `result` and predicted the result half was unimplementable; a
first version of the test supplied a sha256 through the `sealed` fixture and so asserted a
sealed item with no result, which is a failure. Both are in P3's deviations, and the code was
right both times.

<!-- xeno:section:results -->
## Results

Nine criteria met, one pending until the commit.

1. Met. The table row is `{"G-Test", 4, testReport}` and `TestGTestIsImplementedAtP4` asserts
   both the phase and that the function no longer reports `not-implemented`.

2. Met. A declared `test-report` with `result: fail` is a finding naming the job; with `pass` it
   is green; declared with no hash and nothing attached it is `pending`, which is A60's rule and
   G-Build's path rather than a new one.

3. Met. `declaredResults` is the one body; `build` and `testReport` are two lines each. Neither
   reads the other's kind, which is asserted in both directions.

4. Met. `TestKind` is beside `BuildKind` and compared against
   `testdata/evidence-declaration.yaml`, so a rename in the document fails the suite rather than
   making the gate inert.

5. Met. `./xeno gate verify` exits 0 over 384 verdicts: the 381 that existed are intact and the
   three are this intent's own judged phases.

6. Met. Two rows for section 7's one sentence, the count at 37, and 37 rows counted in the table
   by script. Two paragraphs say the rows arrived in #212, that a green G-Test now means half its
   row, and why the mapping half has no reader, with the figures and the date.

7. Met. #250, carrying the measurement and the three things deciding it involves.

8. Met. Six tests cover every case the criterion lists, including an artifact declaring neither
   kind.

9. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints
   nothing, and `go test ./...` is `ok` across all eighteen packages that have tests — the run
   whose output this phase declares.

10. Pending. The commit comes after this phase is judged.

**And the result that is not a criterion.** `gate run` on this phase reports **G-Test: pass**
against the declaration this phase made: `kind: test-report`, `job: go-test`, `result: pass`,
bound to `evidence/go-test.txt` by `fe3d2dc7`. The gate that had never judged anything in this
trail judges the artifact of the intent that implements it, on a real declaration rather than a
fixture. That is the check a unit test cannot make, and it is the same shape as #242's P5 writing
its own checklist with the command it added.

<!-- xeno:section:gaps -->
## Gaps

Half of section 7's G-Test row has no reader and a green G-Test now means less than it looks
like. Before this change the gate said `not-implemented`, which was useless and honest; now it
says `pass` having judged the declared result and not the completeness of the mapping. Nothing
in a verdict can express that, because section 5 gives a check one result from a fixed set, so
the only place it is written is `docs/clause-readers.md` and the only thing that reads that is a
person. This is the cost the maintainer chose over the status quo and it is not mitigated, only
recorded.

The mapping half may not be reachable as a gate at all, which #250 carries. 46 of 55 P1
artifacts have no numbered criteria, so a completeness check re-judges most of the trail, and a
gate matching numbers in prose is near the thing A90 warns about. A `review` rule answered by a
person may be the honest ceiling, which would leave this clause permanently without a mechanical
reader.

Nothing checks that a P4 declares a test report at all. G-Test counts from the declarations, so
an artifact declaring none is green — `TestNeitherGateFiresWithoutADeclaration` asserts exactly
that. Every P4 in this trail but eighteen declares nothing, and this gate will pass all of them
in silence. Whether a P4 owes a test report is not something section 7 says, so this is a gap in
what the gate can mean rather than a defect in it.

`docs/clause-readers.md` is now a dated measurement carrying three late rows, and this intent
found a fourth row that is already stale: line 84's clause about a question's options was given
a larger reader by #229 and still names only G-Questions. P3's learning proposes that an intent
changing a clause's reader updates its row in the same commit; until something does that, the
document's accuracy depends on each intent remembering it exists, and two intents in a row did
not.

The declaration this phase made is a file in the trail rather than a pipeline artifact. It was
produced by `go test ./...` on this machine and bound by hash, which is what `--file` is for and
is weaker than the pipeline path the plan describes: nothing proves the run happened in CI or
that the output matches the commit. A60 and the evidence model allow it; it is worth knowing that
the gate's first real judgement in this trail was over locally produced evidence.
