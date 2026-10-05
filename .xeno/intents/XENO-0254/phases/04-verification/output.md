---
intent: github.com/triplem/xeno#254
phase: 04-verification
created: "2026-10-05T19:02:56Z"
schema_version: "1.0"
runner_version: dev+e471bbb.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a0533d8134280932eed3698da8ed824c366152b5776902bd939326484092663e
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
      sha256: f8121d9f24313de16aee35b15554c76f7b4d8f0038a7ccd55963b35df9aa16f0
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

No test is added and none could be. The change is two rows of a document that nothing reads:
no gate, rule, template or Go file opens `docs/clause-readers.md`, and A90 is the row recording
that this is deliberate. A test would compare this commit's words against a copy of themselves.

What stands in for it is reading, four counts and one check against the code:

| criterion | what answers it |
|---|---|
| 1, two rows in section 8's order | reading them |
| 2, G-Schema is named | `grep` for `QuestionShape` in `internal/gates/gates.go`, which finds the call in `phaseResult` as well as the definition — the second is what makes G-Schema the reader |
| 3, the recommendation's reader is a writer | the call site in `internal/runner/exchange.go`, which chooses `QuestionAsked`, and the absence of any other caller |
| 4, the count | `grep` for the count line, and a script counting rows between the two headings: both say 38 |
| 5, the correction is distinguished | reading the paragraph, which says "a correction rather than a late addition" and why |
| 6, no clause text invented | both rows read against section 8's own sentence in `docs/process-definition.md` |
| 7, one file | `git diff --stat` |
| 8, prose within 88 | an `awk` pass over lines outside tables and code blocks, which caught two at 89 and now finds none |
| 9, the trail | `./xeno gate verify` at exit 0 over 402 verdicts |
| 10, the suite | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l` |
| 11, one commit | the commit itself |

Criterion 2 is the one that needed checking against code rather than reading. "G-Questions is
wrong" is easy to assert and the evidence is specific: `QuestionShape` is called from
`phaseResult`, `phaseResult` is called from `schema`, and `schema` is `{"G-Schema", 0, schema}`
in the table — so a malformed question fails from P0. That chain is what the row now states and
it is three greps rather than an opinion.

The suite is run and proves nothing about this change, which criterion 10 says and this phase
repeats rather than letting a green suite stand as evidence. It would pass identically if both
rows were wrong.

This phase declares its test report as the previous four have, so G-Test judges it. That is the
only mechanical check in this intent that touches the change at all, and it touches it only in
the sense that the artifact carrying these words is itself judged.

<!-- xeno:section:results -->
## Results

Ten criteria met, one pending until the commit.

1. Met. Two rows, options with their consequence and the free entry first, the recommendation
   second, which is section 8's order.

2. Met. The first row names G-Schema's shape check before G-Questions, and the chain is
   checkable: `QuestionShape` is called from `phaseResult`, `phaseResult` from `schema`, and
   `schema` is the table's `G-Schema` entry from phase 0.

3. Met. The second row reads "`QuestionAsked`, the writer only; no gate reads it".

4. Met. The count line says 38 and a script counting rows between the two headings says 38.

5. Met. The paragraph above the table says these two are a correction rather than a late
   addition, and that the pass "was not incomplete here; it was overtaken, and then wrong".

6. Met. Both rows were written against section 8's sentence in `docs/process-definition.md`
   rather than against the row they replace.

7. Met. `git diff --stat` names `docs/clause-readers.md` and nothing else: 13 insertions, 2
   deletions.

8. Met, after a correction inside the phase. The new paragraph came out at 89 columns on two
   lines and was replaced whole and read back rather than rewrapped at the break; the `awk`
   pass now finds nothing over 88 outside tables and code blocks.

9. Met. `./xeno gate verify` exits 0 over 402 verdicts: the 399 that existed are intact and the
   three are this intent's own judged phases.

10. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints
    nothing, and `go test ./...` is `ok` across all eighteen packages that have tests — the run
    this phase declares, and a run that says nothing whatever about whether these two rows are
    right.

11. Pending. The commit comes after this phase is judged.

**And the half of the request that needed no change.** What prompted this intent was two
findings left by earlier intents, and only one of them was real. `internal/plugin/embedded/plugin`
is gitignored, at `.gitignore:10`, added by #200 with a comment explaining why the directory
above it is tracked and this one is not. #201's finding was false: the check ran
`git check-ignore` against the path moments after deleting the directory, and git answers
nothing for a path that is not there. Verified here by creating the directory, re-running the
check — which names `.gitignore:10` — and confirming `git status` stays silent.

<!-- xeno:section:gaps -->
## Gaps

Nobody has re-read the other thirty-six rows. This corrects the one an intent is known to have
made wrong, found because #212 happened to be adding rows beside it. Four intents since the pass
have changed what reads a clause; two added their rows, one did not, and whether any other row
aged is unknown. The document is dated and says so, which is the whole of the protection.

Nothing enforces the convention that would prevent the next one. #212's learning proposes that
an intent changing what reads a clause updates its row in the same commit; it is a learning,
section 10 routes it through review, and until something adopts it the document's accuracy
depends on each intent remembering it exists — which is exactly what failed here.

The correction cannot be checked by anything. A tool could verify that `QuestionAsked` exists
and that `exchange.go` calls it; it could not verify that the row describes the clause section 8
states, which is where the error was. A90's finding applies to its own document: a reader that
cannot fail in the interesting case is worse than none.

The second row describes an absence, and an absence has no test. "No gate reads it" is true
today because `QuestionShape` does not look at `Recommended`; if some later intent adds that
check — which #229 showed would turn XENO-3's sealed P0 red — the row goes quietly wrong in the
same way the one it replaces did, and in the same direction: claiming less coverage than exists
rather than more.

This intent's own proportions are a finding about the plan rather than about itself. Thirteen
lines of a document, behind seventeen sections and about 1,400 lines of record. #117 has the
figure eight times over and this is the sharpest instance; the plan says no shortcut is defined
on purpose and that the shape of one should follow from measurement, and the measurement now
exists.
