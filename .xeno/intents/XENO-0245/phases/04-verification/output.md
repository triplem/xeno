---
intent: github.com/triplem/xeno#217
phase: 04-verification
created: "2026-10-04T21:20:34Z"
schema_version: "1.0"
runner_version: dev+fbf8a72.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4f20aae59686cb170806a7342c8ee3859732261e74c5b10053cdea647cc6840b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - kind: test-report
      result: pass
      produced_by: go test -json ./...
      format: go-test-json
      sha256: 08d483c5d33bcb3b10cd0808918768d1223d27fee7e59ecea26d415253f815d8
      path: evidence/go-test.json
      job: go-test
    - kind: build-log
      result: pass
      produced_by: go build -o xeno ./cmd/xeno
      format: other
      sha256: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
      path: evidence/build.txt
      job: go-build
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Each acceptance criterion of P1 against what asserts it. A criterion with no test beside
it is named in the gaps section rather than left to be noticed.

**An intent's P0 can produce a scope with a command.**
`TestTheScopeIsWrittenAndItsPatternsAreReported` writes one and reads it back: the
content as given, and the header written by the runner rather than by the entry.
`TestTheScopeCommandWritesP0sArtifactAndReportsTheFigure` is the same through the CLI,
with `--file` and no `--phase`.

**The writer says what the patterns resolve to.** The same two tests assert the figure,
and `TestTheReportedFigureIsTheOneTheLockRecords` asserts the thing the criterion is
actually for: the count and the byte total the writer reports equal the ones `phase
start` records in the lock, because they come from one resolution.

**A scope written after its phase has a verdict says so.** Not a new test. The behaviour
is `section set`'s, which `TestAnArtifactWrittenAfterItsVerdictIsSaidToHaveChanged`
already covers, and this intent's own trail exercised it four times: every re-judge of
P0 in this intent went through that path.

**`phase finish` at P0 refuses where there is no scope.**
`TestP0CannotBeFinishedWithoutAScope` asserts the refusal, that it names both the
artifact and `xeno scope set`, that no verdict is written, and that the digest does not
carry the summary the refused call passed. `TestAPhaseAfterP0IsNotAskedForAScope`
asserts the other side, that the requirement is P0's alone.

**The refusal reaches no sealed phase.** `TestGateVerifyOverThisRepositoryExitsZero`
runs `gate verify` over this repository's real trail, which is the assertion that
matters: 349 verdicts at exit 0, with the 100 historical P0 phases recomputing exactly
as before. The figure is in the results section.

**A phase after P0 is given what the scope names.**
`TestGivenFilesAreComparedAgainstTheTree` keeps the half of its assertions about
resolution: the include order, the exclude, and the files the scope does not name.
`TestTheBaseFollowsTheScopesIncludeOrder` and `TestADeclaredLinksDocumentIsInTheBase`
cover the order of volatility and the link. This intent's own P1 to P4 locks carry 43
files each, which is the first non-empty base in the repository and is asserted by the
trail rather than by a fixture.

**A repeated phase is told what moved.** `TestChangedSinceReportsTheEditAndTheRemoval`
covers it against a fixture, and `TestChangedSinceIsSilentWhereTheScopeMatchesNothing`
covers the empty case. It also ran for real: `phase start` for 04-verification printed
`cmd/xeno/main_test.go` and the other files P3 changed, which is the first time the
mechanism has produced output in this repository.

**A declared link naming a file not in the tree is a finding.**
`TestALinkToAMissingDocumentIsAFinding` in `internal/gates/budget_test.go`, both
directions, which existed before this intent and had never had a scope to read.

**An absent scope is not a finding on a phase that already has a verdict.** The same
`gate verify` run over the real trail, and `TestAPhaseAfterP0IsNotAskedForAScope` for
the unit case.

**The two limits of section 5.** `TestTheLockOfThePhaseBeingGatedIsNotRead` for the
second, `TestOnlyWhatTheChangeUnderReviewTouchedIsReported` for the first,
`TestAPrecedingPhasesMovedInputIsAFinding` for the direction the check is for, and
`TestWithNoCommitRangeTheStalenessHalfReportsNothing` for the gap this intent names
rather than hides. All four use the `gitRepo` and `gitCommit` helpers, so they skip
where git is not on the path.

**This intent's own P1 is green on the merits once the limits are in.** Asserted by the
trail: P1's `gate.yaml` carries no approval and no override, and `gate verify`
recomputes it green. That is the criterion written to check that the fix addresses what
actually happened rather than something adjacent.

**The rename reaches every name the artifact has.** A grep for the old identifiers over
the tree outside `vendor/`, reported in the results section, and the suite itself: the
fixtures and the tests were renamed with the code, so a surviving `Profile` would be a
compile error rather than a silent miss.

<!-- xeno:section:results -->
## Results

Every figure below was taken on this working tree after the change, and the ones that
have a before are stated as a pair rather than as a claim.

**The suite.** `go test ./...` at exit 0, eighteen packages reporting `ok` and two with
no test files. Declared as evidence of kind `test-report`, job `go-test`, with the
command that produced it.

**`gofmt` and `go vet`.** Both silent. `gofmt -l .` outside `vendor/` lists nothing, and
`go vet ./...` produced no output.

**The build.** `go build -o xeno ./cmd/xeno` at exit 0, declared as `build-log`.

**`gate verify`, the figure the whole of D-2 rests on.** Exit 0 over 349 verdicts after
the change. Before the change it was exit 0 over 346, and the three added are this
intent's own P0, P1 and P2; P3 makes 349 and P4 will make 350. No verdict in the 100
historical intents moved, and no divergence was reported at any point, which is what a
requirement living in a command rather than in a gate was chosen to buy.

**The rename is complete.** A grep for `ContextProfile`, `model.Profile` and
`context-profile` over the tree outside `vendor/` and outside the trail returns nothing,
and a grep for `context profile` over the two normative documents returns nothing. The
trail's own sealed artifacts still carry the old name where they described the state
before `fbf8a72`, which is correct: they are records of what was true when they were
written.

**The refusal's reach, measured rather than reasoned.** Adding it refused 88 tests
before the two fixtures were given a scope. That is the count of tests that finish a P0,
and it is the clearest statement available of what the gate-based alternative would have
cost: no fixture change answers a gate, because a gate recomputes sealed phases.

**This intent's own trail, which is the dogfooding result.** P1 to P4 each resolve 43
files and 904,078 bytes against the declared budget of 50 and 1,000,000, the first
non-empty information base in 349 verdicts. P1 carries no approval and no override.
`phase start` for this phase printed the files P3 changed, which is `ChangedSince`
producing output for the first time since it was built.

**The staleness half, in the state the design named.** With both limits in place and no
commit range on a local run, it reports nothing, so P4 and P5 of this intent are not
troubled by P3 having edited files their predecessors read. That is the fix working and
it is also the silent pass #235 is about; it is a gap and not a result, so it is in the
gaps section.

<!-- xeno:section:gaps -->
## Gaps

**The staleness half has no local coverage, by design and not by omission.** Both limits
need a commit range; section 12 forbids inferring one; so on a developer's machine the
half reports nothing and only CI exercises it. The four tests give it a real repository
and a real range, so the logic is covered, but nobody running `xeno gate run` locally
will ever see it fire. Named in P2's decisions, in the gate's own comment, and here.

**And its silence cannot be reported.** A check that did not run has no result to say so
with, which is #235. Three findings in this intent want that same mechanism: this one,
the budget overrun that turns a phase red against section 5's words, and a P0 with no
scope, which is now refused by a command and therefore appears in no verdict. The
learning record at P2 says this rather than filing a third issue about a symptom.

**`xeno scope set` has not been run against a real intent.** It is covered by six tests
in two packages and was deliberately not used to rewrite this intent's own scope,
because the fresh `created` in the header it writes would change P0's `artifacts_hash`
and restart the cascade P3 recorded. The first real use is the next intent, which P1
named as the measure of whether the writer is usable. Until then the only scope in
existence was written by hand.

**The budget has never been exceeded, so the budget finding is still unexercised.** This
intent's own declared budget has headroom on purpose, to keep clear of #235. So the one
reader of the scope that has still produced nothing is the one whose behaviour is
disputed between the code and section 5.

**G-Supply, G-Secret and G-Test remain `not-implemented`** through all six phases of
this intent, as they are for every intent in the trail. Unchanged by this work and named
so that the green above is read for what it is.

**`CLAUSE-READERS.md` is now stale in the way it predicted of itself.** It names readers
by symbol and says "a reader named by symbol is wrong the day the symbol is renamed with
nothing to say so". This change renamed `model.Profile`, `model.ContextProfile` and
`informationBase`'s shape. Not corrected here: the audit is a measurement with a date
and a commit on it, and re-running it is its own act.
