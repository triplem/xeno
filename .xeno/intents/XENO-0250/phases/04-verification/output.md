---
intent: github.com/triplem/xeno#229
phase: 04-verification
created: "2026-10-05T15:29:46Z"
schema_version: "1.0"
runner_version: dev+2d997e0.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4cc6a77f756c881ca802ba38c1d5ebffa1d7fb299af568511353c7bd0bc9cd1d
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

Eleven tests across two packages against the eleven criteria.

| criterion | what answers it |
|---|---|
| 1, the consequence check and the free entry's exemption | `TestAnOptionWithoutAConsequenceIsAFinding` for the finding and its counts, `TestTheFreeEntryIsNotAskedForAConsequence` for the exemption |
| 2, the writer requires the recommendation | `TestTheGateIsSilentOnTheRecommendation` for both halves at once, `TestTheWriterWantsExactlyOneRecommendation` for the count, and `TestAQuestionThatRecommendsNothingIsRefusedBeforeTheWrite` through the command |
| 3, one function pair rather than two checkers | `TestTheWriterAddsToTheGatesChecksRatherThanReplacingThem`: a question wrong in both ways produces both findings from one call |
| 4, the reason is in the code | read, in three places: both functions and the call site. Not a test, and the next criterion is what makes moving the check fail loudly |
| 5, `gate verify` at exit 0 | the command, over 378 verdicts — the 375 that existed plus this intent's three judged phases |
| 6, the writer refuses a bare option too | `TestAnOptionWithNoConsequenceIsRefusedBeforeTheWrite`, which also asserts the artifact was not touched |
| 7, `askable` still passes | `TestTheFreeEntryNeedsNoConsequence`, which is that fixture written through the command unchanged |
| 8, the eight cases | the six in `internal/gates/questions_test.go` and five in `internal/runner/exchange_test.go` |
| 9, the two specification gaps | issues filed, named in the review phase |
| 10, the suite | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l` |
| 11, one commit | the commit itself |

`TestTheGateIsSilentOnTheRecommendation` is the test this intent exists to leave behind. It
asserts an absence, which is usually a weak thing to assert, and here it is the strongest: the
reason the gate does not read the recommendation is a measurement over the trail that nobody
re-running the suite will repeat, so the test carries it and fails if somebody moves the check
up. Its comment says what to check before trusting the failure.

`TestTheWriterAddsToTheGatesChecksRatherThanReplacingThem` is how criterion 3 is testable at
all. "They share one definition" is a claim about the code's shape; "one call reports both
kinds" is a behaviour, and it is false for any implementation that restates the gate's checks
and drifts.

The two fixtures in the writer's tests are chosen rather than invented. `bare` is `askable` with
the consequences removed, so the difference between accepted and refused is exactly the field
under test; `unrecommended` is XENO-3's Q-2's shape, so the test that the writer refuses it is
also the demonstration that the trail holds one.

<!-- xeno:section:results -->
## Results

Ten criteria met, one pending until the commit.

1. Met. `QuestionShape` reports "question Q-1 has 1 of 2 options with no consequence", and a
   bare free entry is not a finding.

2. Met. `QuestionAsked` reports "question Q-1 recommends 0 of its options" and the same for 2,
   and `xeno question record` refuses on both; the gate reports neither.

3. Met. `QuestionAsked` calls `QuestionShape`, and a question wrong in both ways produces two
   findings from one call, which is the behaviour that a drifting second checker would fail.

4. Met. Both functions and the call site in `exchange.go` carry the reason, naming XENO-3's Q-2
   and the `DIVERGENT` line.

5. Met. `./xeno gate verify` exits 0 over 378 verdicts: the 375 that existed before this intent
   are intact and the three are its own judged phases. Re-run against the real change rather
   than against P0's hand-edited probe.

6. Met. `TestAnOptionWithNoConsequenceIsRefusedBeforeTheWrite`, which also asserts the artifact
   and its body are untouched after the refusal.

7. Met. `askable` passes through the command unchanged, which is what
   `TestTheFreeEntryNeedsNoConsequence` is.

8. Met. Eleven tests, six in `internal/gates/questions_test.go` and five in
   `internal/runner/exchange_test.go`, covering every case the criterion lists.

9. Met. #247 for the recommendation's reason having no field, #248 for nothing saying questions
   are put one at a time. Both carry the evidence and both say why #229 could not do them.

10. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints
    nothing, and `go test ./...` is `ok` across all eighteen packages that have tests — after a
    fixture correction that P3's deviations records, which three tests failed on before they
    passed.

11. Pending. The commit comes after this phase is judged, because this phase's artifacts are
    part of what it commits.

<!-- xeno:section:gaps -->
## Gaps

The gate cannot see the recommendation and that is permanent until somebody decides otherwise.
A hand-written artifact carrying a question that recommends nothing passes every gate, so the
guarantee holds only for questions written through `xeno question record`. That is the same
asymmetry #242 left behind for the review checklist, and the same one #229 was filed to
complain about, narrowed rather than closed: two of three requirements now have a reader in the
gate, and the third has one only in the writer.

Nothing will notice if the trail stops holding XENO-3's Q-2. The measurement that justifies the
split is a fact about today's trail, and `TestTheGateIsSilentOnTheRecommendation` asserts the
code rather than the reason. If that question were ever released or the intent closed
differently, the check could move into the gate and nothing would prompt anybody to try.

The reason for the recommendation is unreadable and unwritable, which is #247. Every
recommendation in this trail either carries its reason in prose or carries none, and nothing can
tell which, so the clause's third requirement has no reader and no field — the worst of the
three states a requirement can be in.

Sequence is unchecked and may be uncheckable, which is #248. Nothing records when a question was
asked, so no gate can tell a batch from three questions asked over three days; a reader would
need a timestamp, which is a new field and its own decision.

Whether the other exported shape functions are in the position `QuestionShape` was in is not
checked. `internal/gates` had no test for questions at all until this intent, which is how a
fixture carrying the forbidden shape survived since WP5; `EvidenceShape` and `DecisionShape` are
exported for the same reason and tested the same way, through the writer in another package, and
nobody has looked.
