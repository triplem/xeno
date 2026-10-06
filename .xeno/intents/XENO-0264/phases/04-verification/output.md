---
intent: github.com/triplem/xeno#258
phase: 04-verification
created: "2026-10-06T17:45:53Z"
schema_version: "1.0"
runner_version: dev+d3983d3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ad1da3dabd1c43d5802a9635adc88c037488a057e20085206af5c23a352f5632
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
      sha256: 5df57bf71b6fa29606eb85a8e6df3cc8627c6f005aae7d02c49fbcea469e0a83
    - format: other
      job: checks
      kind: other
      path: evidence/checks.txt
      produced_by: both checks end to end, the mutations, and the counts
      sha256: be37a4f07d4acbc9226b3c3ac551f5207eef24fdf0f344225307f2e3273c56e0
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Thirteen criteria, by number. Twelve pass and one does not, which is criterion 9.

| # | what it asserts | how it is checked | result |
|---|---|---|---|
| 1 | `Option` carries `Reason` with `omitempty` | `git diff`; and `gate verify` at the same verdict count, which is what `omitempty` buys | pass |
| 2 | `QuestionAsked` refuses a recommendation with no reason | `TestAQuestionWhoseRecommendationHasNoReasonIsRefused`, and `TestAReasonOnAnotherOptionDoesNotSatisfyTheClause` for the half that distinguishes the two designs | pass |
| 3 | `QuestionShape` unchanged | `git diff` over the function | pass |
| 4 | both templates at 1.1.0, nothing else changed | `git diff`, two lines | pass |
| 5 | G-Schema refuses an unnumbered P1 at 1.1.0 | `TestAtTheNewTemplateUnnumberedCriteriaAreReported`, and end to end on a scratch copy | pass |
| 6 | G-Test refuses an incomplete mapping at 1.1.0 | `TestAMappingThatMissesACriterionIsReported`, and end to end | pass |
| 7 | both read the declared version, not the current one | `TestAtTheOldTemplateUnnumberedCriteriaAreNotReported`, `TestAMappingIsNotJudgedAgainstAnOlderRequirementsPhase`, and 69 P1 and 68 P4 artifacts still at 1.0.0 with `gate verify` green | pass |
| 8 | the comparison is numeric | `TestTenthMinorVersionIsLaterThanTheNinth`, which asserts its own premise | pass |
| 9 | **both checks fire on this intent** | its P1 declares `requirements@1.0.0` | **fail** |
| 10 | `testReport`'s comment is corrected | a person reads it; `grep` for "section 9" and "46 of 55" | pass |
| 11 | two rows name their readers | a person reads the table | pass |
| 12 | the measured cost of the bump is recorded | a person reads the document, the artifacts and the pull request | pass |
| 13 | `go test`, `go vet`, `gofmt`, `gate verify` and the same verdict count | run; `evidence/go-test.txt` | pass |

Criterion 9 fails for a structural reason and not an oversight, which P3's deviations states: the
bump lands in the implementation phase, after the requirements phase has sealed its artifact.

<!-- xeno:section:results -->
## Results

## Both checks fire end to end, on a scratch copy at the bumped templates

A fresh intent, run to P4 with the new templates in place:

    P1 at requirements@1.1.0, criteria unnumbered
      01-requirements: red / G-Schema fail
      acceptance-criteria carries no numbered criterion and requirements@1.1.0
      requires a numbered list

    the same P1, criteria numbered 1 2 3
      01-requirements: green / G-Schema pass

    P4 at verification@1.1.0, mapping naming 1 and 3
      G-Test fail
      test-mapping does not name acceptance criteria 2

    the same P4, all three named in a table   G-Test pass
    the same P4, all three named in prose     G-Test pass

The last two matter as much as the failures: a check that only ever goes red is not a reader, and
section 5 says the mapping "names each criterion by its number" without saying how, so a table and
a sentence both have to pass.

## Every assertion was confirmed able to fail

| mutation | what failed |
|---|---|
| `templateAtLeast` made lexical | `TestTenthMinorVersionIsLaterThanTheNinth` |
| `numberedCriteria` ignoring the declared version | `TestAtTheOldTemplateUnnumberedCriteriaAreNotReported` |
| `mappingComplete` reporting nothing | `TestAMappingThatMissesACriterionIsReported` |
| `namesNumber` made a substring match | `TestANumberIsNamedAsATokenAndNotAsASubstring` |

The first mutation was run twice. The first attempt was crude enough that the test failed on a
malformed ref rather than on `1.10.0`, which is a pass earned for the wrong reason; a test asserting
the claim directly was added, including its own premise that a string compare would order the two
wrongly, and the mutation was made precise.

## The trail is not re-judged

    gate verify: verified 463 verdicts
    69 P1 artifacts, all template: requirements@1.0.0
    68 P4 artifacts, all template: verification@1.0.0

Both checks key on the declared version, so none of the 137 is in scope and no verdict moved. The
count rose only by this intent's own phases.

## The writer requirement, and the seven tests it refused

`QuestionAsked` now refuses a recommended option with no reason, and seven existing tests wrote
exactly that. Both fixtures gained a reason rather than the check being narrowed; the diff touches
seven test files and none of them is a test weakened to pass.

`TestAReasonOnAnotherOptionDoesNotSatisfyTheClause` is the one that earns the design: a reason on an
option nobody recommended leaves the clause unsatisfied, which is what putting the field on `Option`
rather than on `Question` buys and the only test that can tell the two apart.

## Two figures found wrong

The clause table carried **thirty-nine** rows at HEAD while the prose said **thirty-eight**, before
this intent touched either. With the new row it is forty, and the paragraph says both numbers and
which was wrong. Counted by reading the rows.

`testReport`'s comment named section 9 — the Rule model — and carried "46 of 55 P1 artifacts", a
figure superseded twice. `grep` for either now returns nothing in that function.

## The parser is not copied

    bodies that read an artifact back into sections: 1

`renderedSections` delegates to `parseSections`, so the section a rule reads and the section a gate
reads cannot diverge.

## The suite

`go test ./...` exits 0 across every package, in `evidence/go-test.txt`. `go vet` is clean, `gofmt
-l` outside `vendor/` prints nothing, `xeno gate verify` exits 0.

<!-- xeno:section:gaps -->
## Gaps

**130 sealed artifacts lost a reader and nothing announces it.** `goneBundle` treats a bumped
template version as a bundle that is gone, so `strings_hash` is no longer recomputed for any P1 or P4
artifact declaring 1.0.0. Measured: red at 1.0.0, green at 1.1.0, with `gate verify` reporting the
same count either way. Tampering is still caught — `artifacts_hash` and the successor's G-Freshness
report a divergence and exit 1 — so it is one of two readers. There is no honest repair and the
document now says so.

**Criterion 9 is unmet and cannot be met by this intent.** The first artifacts either check judges
belong to the next one. So both checks are proved by unit tests and by a scratch intent, and not by
the trail they were written for, until something else runs.

**The mapping check can report a false green.** A number appearing anywhere in `test-mapping` as a
token satisfies it, so a mapping that mentions "3" in a sentence about something else covers
criterion 3. Deliberate: section 5 does not say how a criterion is named, and a stricter reading
would fail a format the clause permits. The cost is named in the code, in the document and here.

**Nothing requires the numbering to be sane.** `1. 1. 1.` satisfies `numberedCriteria`, and a mapping
naming `1` satisfies the completeness check for all three. Consecutiveness and uniqueness are not in
section 5, and inventing them is what the second standing rule forbids — but the gap is real and a
future clause is where it would close.

**The reason has no gate, so a hand-written question needs none.** `QuestionAsked` is the writer.
Section 5 says G-Schema is the backstop for whoever writes by hand, and for this clause it is not:
one sealed question in this trail already fails the recommendation requirement, so a gate reading
either half would re-judge the trail. #229 settled it and this intent inherits it.

**The plugin digest moved and nothing compares it.** Two template files changed, so a new
`context.lock.yaml` records a different `plugin.sha256` from the sealed ones. G-Supply is
`not-implemented`, so nothing notices; when it is implemented it will have to read the lock's own
recorded value rather than the current tree's.

**A project overriding `requirements` at 1.0.0 opts out silently.** The per-id override section 5
describes means a project's own template keeps whatever version it declares, so an adopter who forked
the template before this release gets neither check and nothing says so.
