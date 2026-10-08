---
intent: github.com/triplem/xeno#330
phase: 04-verification
created: "2026-10-08T17:56:10Z"
schema_version: "1.0"
runner_version: dev+8fb365d.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3ae64436e3dc18c6c096402a5b7b95bae0397f526d9e22b0d737fd6ffefdab60
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| criterion | how it was checked | result |
|---|---|---|
| 1. an unlabelled issue does not become an intent | `TestAnIssueNobodyApprovedDoesNotBecomeAnIntent/neither` and `/the_comment_alone`: a refusal naming *the label approved*, and no directory left | met |
| 2. a label without the word is not enough | the same test, `/the_label_alone` and `/a_comment_that_only_contains_the_word`: a refusal naming *a comment whose first line is approved*; `TestApprovalNeedsTheLabelAndTheWord` on the six shapes of the comment | met |
| 3. both present, the intent starts as before | `TestAnApprovedIssueBecomesAnIntent`, and the five existing intent start tests moved onto the approved fake host and still asserting what they asserted | met |
| 4. a milestone holds the issue until its turn | `TestAMilestoneHoldsTheIssueUntilItsTurn`, six cases: the second and the undated held naming `"1.0"`, `--now` starting the second, the earliest, a closed one and none starting; `TestAheadIsTheEarliestOpenMilestoneThatIsNotTheIssues` on the ordering, two undated by number | met |
| 5. a read that cannot happen is a refusal | `TestAReadThatCannotHappenDoesNotStartTheIntent`: no token with no call made, 404, 403, another host with no call made, and a 401 as an error rather than a refusal; against the real host, `--for 330` without a token and `--for 99999` | met — the 404 case failed first, see results |
| 6. without a tracker block nothing changes | `TestWithoutATrackerBlockTheWholeIdIsGivenByHand`, unchanged but for the argument | met |
| 7. the intake says by what it was authorised | `TestTheIntakeSaysByWhatItWasAuthorised`: who, when, the reason, the milestone sentence, and the sentence above the quote; `TestTheIntakeSaysWhereNoApprovalWasFound` | met |
| 8. both adapters read the five things in their own words | `TestAnIssueIsReadWithWhatTheHostSaysAroundIt` on GitHub and `TestAnIssueIsReadWithWhatThisHostSaysAroundIt` on GitLab, asserting the paths in order, the field names, `state=open` against `state=active`, and the system note left out; `TestAnIssueWithoutAMilestoneListsNone` and `TestCommentsBeyondTheFirstPageAreFollowed` | met |
| 9. the flag is in both copies | `TestTheDispatchTableAndTheUsageAgree` and the test holding `docs/commands.md` to the `usage` constant, both green on the edited line | met |
| 10. the suite and the gates stay green | `go test ./...`, `gofmt -l`, `go vet ./...`, `xeno gate verify`, figures in results | met |
| 11. the register carries the two decisions | A103 in `docs/assumptions.md` | met |

<!-- xeno:section:results -->
## Results

Eleven criteria, eleven met.

    go test ./...                   ok, 20 packages, exit 0
    gofmt -l . | grep -v vendor/    nothing
    go vet ./...                    clean
    ./xeno gate verify              verified 547 verdicts, exit 0

Thirteen new tests across four files — six in `internal/runner/tracker_test.go`, three
in `internal/host/github/tracker_test.go`, one in `internal/host/gitlab/tracker_test.go`,
three in `internal/model/identity_test.go` — most of them table-driven, and 515 passing
tests and subtests in the touched packages with `-count=1`. Seven existing tests moved
onto a fake approved host and assert what they asserted.

Three things the suite caught before any reader did, in the order they happened. The
first run of `TestAReadThatCannotHappenDoesNotStartTheIntent` judged a 404 as an issue
missing both halves, because the adapters put a host's refusal on the issue's `Reason`
rather than in the read's middle answer; `approved` now folds it in, and P3 records the
deviation. The MCP route comparison, `TestAPhaseDrivenThroughTheOperationsIsTheSameArtifact`,
went red twice: once on the intake quoting an issue from a per-fixture port, then on the
intent id itself carrying that port — one shared fake host per package is the answer, and
P3's learning records the shape. And `TestEveryCommandLeavesASealedPhaseWhereItWas`, the
walk over every command, refused `intent start` for a tracker block with no credential,
which is the walk doing what #242 built it for: catching the command nobody was thinking
about.

Against the real host, by hand on 2026-10-08, with nothing written in any case:

    xeno intent start --for 330      refused: … is not approved: it is missing a comment
                                     whose first line is approved. …
    (no token) --for 330             refused: the intent is not started, because starting
                                     one reads the issue's approval and no issue was read:
                                     no token in the environment: XENO_TRACKER_TOKEN …
    --for 99999                      refused: … there is no issue 99999 in triplem/xeno, or
                                     the token cannot see the repository …

**Timings, for #117.** `gate run` on P3 of this intent, on a scratch copy: 0.093 s.
`gate verify` over 547 verdicts: 34–36 s, three runs. That is not this change: the binary
built from `main` takes 34–36 s over this tree and over its own, and this binary takes
the same over `main`'s tree. XENO-0266 measured 1.6–1.9 s over 477 verdicts on
2026-10-06, so something between then and `e22a533` moved the per-verdict cost by an
order of magnitude, and it is a finding for #117 rather than for this intent.

<!-- xeno:section:gaps -->
## Gaps

**No run against a real approved issue.** The three refusals were run against the real
host; the one path that writes was not, because no issue of this repository carries a
comment whose first line is `approved` yet — this issue's own answer predates the word.
The positive path is proved by the fake hosts, which serve what the real ones document,
and by nothing else until the maintainer writes the first such comment.

**The milestone ordering is proved on fakes only.** The repository has no milestones, so
the ordering by due date, undated last and by number has met no host's actual listing;
GitHub's `due_on` and GitLab's `due_date` are compared as text within one host, which
holds for the formats both document and has not been seen to hold on a self managed
deployment.

**Paging beyond one page is proved on GitHub only**, and the ten-page bound is asserted
nowhere: a test for it would serve a thousand comments to prove a number nobody expects
to meet.

**A comment by anybody counts.** By the decision on the issue, the comment is quoted and
not vetted, and the label's right is the check of standing. Nothing here asserts that a
host actually restricts the label, because nothing here can.

**The recomputed sentence is tested at one moment.** The case A103 describes, a label
removed or a milestone moved between `intent start` and `phase start`, is reasoned about
and not exercised: the sentence says what P0 found, and a test for the gap would assert
that the intake does not say something.

**The real-host checks are not repeatable.** They were run by hand on 2026-10-08 against
`triplem/xeno#330` and `#99999` and are recorded in P3's `changes`; nothing in CI makes
that call, by design.
