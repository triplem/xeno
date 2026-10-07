# The trail in this repository

Xeno governs its own construction, so `.xeno/` here is both the tool's output and its
longest worked example. Three things about it read oddly to somebody who expects a
released tool's trail, and each has a reason.

## What the version fields say here

**Every artifact in this trail was written by a development build.** `runner_version`
reads `dev+<commit>.dirty` on a build made here and the release sets the number through
ldflags, so the trail says plainly that a released binary did not write it. It used to
read `0.1.0-dev`, a literal that stayed put through twenty-nine minor releases — not a
stale number but a wrong one — and a development build cannot do better, because
`debug.ReadBuildInfo` reports `(devel)` for the main module and only git knows the
tag (A85). Artifacts behind that change carry the old shape; the commit is the
boundary. #177 asks for this where section 16 lists what the record does not mean,
which is a change to the process definition and not this file's to make.

`plugin_version` names the vendored plugin an artifact was rendered from, read from its
own manifest and absent where no plugin is vendored (A83). **G-Supply reports
`not-implemented` throughout this trail**, and that is the gate working rather than
missing: its anchor is a digest the release compiles into the binary, a `go build`
carries none, and section 5's `not-implemented` is the state for a check a runner did
not perform. A released binary over the same tree reports `pass` (A86).

The manifest is stamped by the release into the copy it ships, so section 13's one
shared version number holds for anybody who receives a release, and this repository's
own literal is `0.0.0-dev` and never drifts (A87). `xeno init --vendor` copies the
plugin out of the binary, which is what lets a release deliver the tree its digest was
taken over — it used to copy from a directory, so a released binary had nothing to copy
and the gate was anchored to something the release could not produce (A88).

The harness reports its own version, which is the one field of section 5 the runner
cannot know: section 7 forbids it branching on the harness, and a value nobody produced
is worse than an absent one (A35, A80, A81). `XENO_HARNESS_VERSION` says it for a whole
session, which is the channel section 7 lists, and `--tool-version` overrides it for one
run. Either reaches `output.md` on a phase's first `section set`, and `phase finish`
copies it into the digest from the artifact beside it. Said by neither, the field is
absent and G-Schema reports it missing, which is what every artifact written before this
carries.

## The two shapes of the intents directory

`.xeno/intents/` holds Xeno's own process artifacts, in two shapes. The intents up to
`XENO-0106` run P0 and stop there, which was deliberate rather than abandoned:
`m0-gate-job.md` beside this file runs one intake per work package and goes no further,
so an intent whose P1 to P5 read `not-started` is complete for what it set out to
record. That guide has the reasoning and the sequence, and `assumptions.md` has it as
A20.

From `XENO-0107` the intents run all six phases and finish green, which is what the
second half of G-Freshness made judgeable (A6, #63): a later phase can be compared
against what its predecessor read, so a phase beyond P0 can be judged honestly. Read
`xeno intent status --all` for the whole trail rather than this paragraph, which is a
statement about two shapes and not a count.

Coverage against the implementation plan, by acceptance criterion:

| Package | Covered here | Test |
|---|---|---|
| WP1 | hashing to Appendix B, finding ids, derived status, decisions carried forward by id, the sealed invariant, the four field sets over frontmatter and YAML alike, `schema_version` read rather than enforced backwards, no decision carried forward for an external finding, a development build naming the commit it came from | `hashing_test`, `TestDecisionCarriesForward…`, `TestNoCommandButFinish…`, `TestDecisionSurvivesForXenoAndNeverFor…`, `gates/schema_test` against `gates/testdata`, `model/version_test` |
| WP2 | template resolution with project over plugin, anchors the caller never writes, a second language that changes only the headings, a missing bundle that fails rather than falling back | `template/template_test`, `TestSectionSetRendersAnchors…` |
| WP3 | six shipped templates, both strings bundles, the required section budget | `TestEveryShippedTemplate…`, `TestTheRequiredSectionBudgetHolds` |
| WP9 | init that changes nothing on a second run, appends to .gitignore without replacing it, stops on a version mismatch, and prints what it cannot do | `init_test` |
| WP10 | a wrapper that passes both ends of the range, not available told apart from unmet, and a waived requirement that stops being a daily complaint | `enforcement_test`, `TestTheWrapperPassesBothEnds…` |
| WP5 | question resolution by decision or confirmed assumption | `TestQuestionResolved…` |
| WP6 | pending declarations, pulled attachment, provisional verdicts, tampered attachments rejected, a scan report from the pipeline attaching with its database age | `TestPendingEvidence…`, `TestNextStartAttaches…`, `TestTampered…`, `TestAScanReportFromThePipeline…`, `TestTheScanWorkflowsWrite…` |
| WP7 | sequence enforcement, run marker, G-Questions from P5, gate run attaching for P5, the three commands that write a decision, a stale verdict refused, an abandoned intent closed and judged, the shipped commit-message pattern, the next step of the working sequence read off the state and deciding nothing, the digest written from the agent's summary and filtered before it is hashed, the intent created by a command rather than by hand, the harness reporting the one field the runner cannot know, through the variable section 7 lists or a flag over it, the learning record written by a command | `TestOutOfOrder…`, `TestSecondStart…`, `TestGateRunAttaches…`, `TestApprovalTurnsRed…`, `TestOverrideCarries…`, `TestDecidingOnAStaleVerdict…`, `TestClosingAnAbandonedIntent…`, `gates/patterns_test`, `TestTheSuggestion…`, `TestAnOverrideIsOwed…`, `TestFinishWritesTheDigestFromTheSummary`, `TestTheDigestIsFilteredAndSaysWhichFilter`, `secrets/secrets_test`, `TestStartingAnIntentDerivesEverythingButTheIssue`, `model/identity_test`, `TestTheReportedToolVersionReachesBothArtifacts`, `TestASecondFinishKeepsTheToolVersionWithoutBeingToldAgain`, `TestTheHarnessVersionIsReadFromTheEnvironment`, `TestTheFlagBeatsTheHarnessVersionVariable`, `TestTheLearningRecordCarriesTheRunnersHeader` |
| WP8 | the information base resolved from the context profile with a hash each, a repository without a profile recording none rather than an empty one, and the context hash of the lock written beside a section | `TestNoProfileRecordsNoInformationBase`, `TestSectionSetWritesTheContextHashOfTheLockBesideIt` |
| WP11 | the intent and the phase exported into the environment a harness reads for request headers, and nothing else printed on that path | `TestExportPrintsTheEnvironmentAndNothingElse` |
| WP12 | the branch rules port with no default host, an adapter chosen by configuration, and the GitHub and GitLab adapters behind it, where forbidden and absent are answers rather than errors and no requirement name is invented | `host_test`, `host/github`, `host/gitlab` |
| WP13 | a turn's counts read from the transcript, a phase totalled from the differences between its lines, two sessions kept apart, and a turn spent with no phase open recorded as attributed nowhere | `cost/cost_test` |
| WP15 | the index format read and round-tripped against this project's own producer, every location returned for a name, and a missing, unreadable or stale index treated as absent | `index/index_test`, `TestNoIndexBlockMeansNoIndex`, `TestTheIndexPathIsRelativeToTheRepository` |
| WP17 | the exit code staircase asserted through `run` rather than by running the binary, and the dispatch table checked against the usage text | `cmd/xeno/main_test` |

Two deliberate mutations were run against the suite, removing the sequence check and
letting attachment write into a sealed artifact; both were caught.
