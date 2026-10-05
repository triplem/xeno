---
intent: github.com/triplem/xeno#242
phase: 04-verification
created: "2026-10-05T12:18:27Z"
schema_version: "1.0"
runner_version: dev+8f4b75b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4b5169e23673fe79a2491c5ea893ebae1536de5efa22550c0edf3e98c4031676
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

Eight tests in `internal/runner/review_test.go`, sixteen cases counting the refusal table's
eight, against the ten criteria:

| criterion | what answers it |
|---|---|
| 1, the command exists and writes | `TestAnAnswerIsWrittenAndWhatIsLeftIsReported`, plus `xeno --help` listing the usage line |
| 2, the frontmatter re-reads | the same test, through `f.readFront`, which is the call the gates make |
| 3, the three refusals | `TestAnAnswerThatTheGateWouldRejectIsRefused`, eight cases, each asserting the message says why and that nothing was written |
| 4, replace in place | `TestAnsweringARuleTwiceReplacesTheAnswerInPlace`, on the list length, the content and the index |
| 5, the unanswered report | `TestAnAnswerIsWrittenAndWhatIsLeftIsReported`, and the complete case in its second half |
| 6, no phase argument | `TestTheChecklistIsTheReviewPhases`, which answers while P0 is open and asserts the entry lands in P5 and not in P0 |
| 7, no `gate.yaml` read | not a test, and deliberately: asserting a file is not read is asserting the absence of code. The deviations section carries the reasoning and `review.go`'s comment carries it where it will be read |
| 8, the cases listed | the eight tests above, plus `TestMetNeedsNoNote`, `TestTheBodySurvivesAnAnswer`, `TestAnAnswerBeforeTheArtifactIsRefused` and `TestAnAnswerAgainstARuleSetWithNoReviewRuleIsRefused` |
| 9, this P5 written with it | the artifact of this phase's successor, and the one check a unit test cannot make |
| 10, the suite | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l`, `./xeno gate verify` |

The refusal table asserts two things per case and the second is the one that matters: the
message contains the reason, and the checklist is still empty afterwards. A refusal that
rejected the call and wrote anyway would pass a test that only read the error.

`TestTheBodySurvivesAnAnswer` is a byte assertion rather than a parse. The writer's whole claim
against going through `SectionSet` is that the body is not re-rendered, and a test that parsed
the artifact would pass whether or not the body had been rewritten.

The checked rule is in the fixture to be left out. `reviewRules` lays down two `review` rules
and one `checked`, and two tests use it: the refusal table rejects an answer to the checked rule,
and the complete case asserts it is not counted as owing one. A fixture with only review rules
would pass a writer that ignored `Kind` entirely.

<!-- xeno:section:results -->
## Results

Nine criteria met, one pending until the next phase.

1. Met. `xeno review answer` is in the usage text and in the command table, and writes the
   entry.

2. Met. The amended artifact re-reads through `fm.ReadFront` into the same `model.Output` the
   gates read, with `review_checklist` in `frontmatterOrder`'s position.

3. Met. Eight refusals, each with its reason in the message and nothing written: a rule outside
   the set, a `checked` rule, no rule, no result, a result outside the three, `deviation`
   without a note, `not-applicable` with a blank note, and a `source`.

4. Met. Re-answering replaces at the index it was first answered in; the list stays at two
   entries and the first keeps its place with the new result and note.

5. Met. The report names what is left and says the set is answered when nothing is.

6. Met. An answer given while P0 is open lands in P5, and P0's checklist stays empty.

7. Met by construction and not by a test, which the test mapping states rather than counting it
   as covered.

8. Met. Eight test functions in `internal/runner/review_test.go`, fifteen leaf cases counting
   the refusal table's eight subtests, all passing. The figures stated here are the file's own:
   an earlier draft of this section said twelve functions and nineteen cases, which was the
   output of a `-run` pattern wide enough to match three pre-existing tests in the package.
   Corrected before this phase was judged, and worth recording because a count taken from a
   filtered test run is a count of the filter.

9. Pending. The checklist of this intent's P5 is written with the command in the next phase,
   which is the criterion's point and cannot be met before then.

10. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints
    nothing, `go test ./...` is `ok` across all eighteen packages that have tests, and
    `./xeno gate verify` exits 0 over 361 verdicts: the 357 that existed are intact and the
    four are this intent's own judged phases.

<!-- xeno:section:gaps -->
## Gaps

Nothing makes the writer compulsory. A hand edit still produces a valid checklist, because the
field is ordinary frontmatter and no gate can tell which wrote it. This intent removes the
necessity, not the possibility, and the statement that the hand edit is now avoidable is the
`CLAUDE.md` change XENO-0246's learning asked for, which section 10 routes through review.

A74 bounds the rule refusal, and no test covers the divergence. The writer compares against the
set that resolves now, G-Policy against the set the artifact recorded in `rules_hash`, and a
test for the case would have to change the rule tree between `section set` and `review answer`.
It is written down in the design's impact and in the file's comment instead, because the
behaviour is correct in both directions — the gate decides — and what would be asserted is the
wording of a refusal that turns out to be advice.

The command has no test. `cmdReviewAnswer` is covered only by the smoke runs recorded in P3:
the four refusals were exercised from the shell against this repository, and the success path by
this intent's own P5. `cmd/xeno` has its own test file and the other writers' commands are
reached through it, so this is a genuine thinness rather than a boundary, and the reason it
stayed is that the behaviour under test is one `Fprintf` and a call. A reader who disagrees is
right to.

Nothing reports a checklist that is complete but wrong. Every refusal is about form: a rule in
the set, a result among three, a note where one is owed. Whether `met` is the honest answer is
what the checklist exists for and what no writer can check, and A90's finding applies — a reader
that cannot fail is worse than none, so the writer does not try to judge an answer.

The eighteen existing checklists stay hand-written. They pass G-Policy and rewriting one would
change its `artifacts_hash` and stale its verdict, so the writer applies from here and the trail
keeps two provenances for one field with nothing marking which is which.
