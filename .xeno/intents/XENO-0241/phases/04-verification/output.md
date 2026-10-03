---
intent: github.com/triplem/xeno#206
phase: 04-verification
created: "2026-10-03T20:54:10Z"
schema_version: "1.0"
runner_version: dev+8574810.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: eb5731b1881ad8fbe0f0543240779c4b9f1475f22bce9eaaceef056522b8a0bd
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

Every acceptance criterion of P1 against what answers it. Twelve tests and one exercise over
a real clone; the clone is named because three of the criteria are about what CI does and a
test in process cannot be that.

| criterion | what answers it |
|---|---|
| exits 1 for a touched intent that is neither complete nor abandoned, naming it | `TestAnIntentThatStopsShortIsReportedWithTheStateItReached`, `TestTheMergeCheckExitsOneForAnIntentThatStopped`, and the clone below |
| exits 0 where every touched intent is one of the two | `TestAnIntentThatReachedADecidedP5Passes`, `TestAnAbandonedIntentPasses`, and the clone below |
| the state is read from the one definition, not derived again | `Completeness` calls `summarise`; `TestTheListingComputesAbandonedCompleteAndInFlight` is the test over that definition and was not touched |
| a range touching no intent exits 0 and says so in words | `TestARangeThatTouchesNoIntentHasNothingToCheck`, `TestTheMergeCheckOverAnEmptyRangeSaysNothingIsTouched` |
| both ends required, an unresolvable ref exits 2 | `TestTheMergeCheckRefusesWithoutBothEndsOfTheRange`, `TestTheMergeCheckWithoutARangeIsTwo`, `TestPathsRefusesAnAbsentEndOfTheRange`, `TestPathsSaysWhichRefDidNotResolve` |
| an unreadable record exits 1 with the reason | `TestAnIntentWhoseRecordCannotBeReadIsUnfinished` |
| a provisional P5 fails | `TestAProvisionalFinalPhaseIsNotFinished` |
| the `verify` job calls it with the same base as the trail guard | read in `.github/workflows/xeno.yml`; the pull request this intent opens is the first run |
| both shipped wrappers carry the call | read in `internal/scaffold/files/ci-github.yml` and `ci-gitlab.yml`; `TestEveryWrapperPassesBothEndsOfTheRange` still passes over both |

Two tests answer no criterion and were written because the behaviour they fix is easy to
get wrong later: `TestAPathAddedAndRemovedInsideTheRangeIsNotReported`, for the tree
comparison, and `TestAMoveReportsBothPaths`, for rename detection being off.

<!-- xeno:section:results -->
## Results

`go build -o xeno ./cmd/xeno`, `go vet ./...`, `gofmt -l .` outside `vendor/` and `go test
./...` all pass. `./xeno gate verify` exits 0 over 324 verdicts. The twelve new tests pass
individually and under the whole suite.

The exercise the tests cannot be: a clone of this repository, a branch off it, and the
command invoked exactly as the `verify` job invokes it, with a base commit and `HEAD`.

Four phases of a real intent copied in under a new key, committed, and the check run:

    checked 1 intent the change touches
      UNFINISHED  XENO-9999: 03-implementation
    An intent ends at a decided P5 or at xeno intent close.
    exit=1

The remaining two phases committed and the same command run again:

    checked 1 intent the change touches
    exit=0

An intent whose `intent.yaml` says `abandoned` with a reason, on a branch of its own:
`checked 1 intent the change touches`, exit 0. A branch that changes `README.md` and
nothing else: `no intent is touched by this change`, exit 0. A base that does not resolve,
which is how the earlier attempt at this exercise went wrong: `error: git diff main HEAD:
fatal: bad revision 'main'`, exit 2.

The clone found one defect the tests did not: the count said "checked 1 intents". Fixed,
and the wording is now chosen by the count.

<!-- xeno:section:gaps -->
## Gaps

**The workflow step itself is unexercised until this pull request.** The command it calls is
tested and was run by hand in CI's exact form over a clone, but that the step is wired
correctly — the base expression, the `if`, the step's place in the job — is asserted by
nothing in the repository and is first proved by the run this intent triggers. That is the
same gap the implementation phase's learning record describes, and it is a property of
adding a step to a workflow rather than of this change.

**The shipped wrappers are templates and are read, not run.** `TestEveryWrapperPassesBothEndsOfTheRange`
checks that both carry the range, which is the property that already had a test; nothing
runs a generated wrapper on either host. An adopter's first run is their first pull
request. WP10 owns this and it is not made worse here.

**No evidence is declared.** This phase's results are a local run, and section 5's evidence
is for artefacts a pipeline produced. Nothing was declared because nothing was fetched, and
G-Evidence passes correctly over an empty set — which #208 points out has been true of
every verification phase in this trail.

**A path under `.xeno/intents/` that is not an intent directory** is skipped rather than
reported. Deliberate, recorded in the implementation phase's deviations, and untested
because the shape does not occur.
