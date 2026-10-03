# xeno

The Xeno runner: hashing, gates, the phase sequence, evidence attachment, the digest and
its secret filter, the cost record, the branch rules port and the symbol index reader.
Built by hand before M0; the assumptions made on the way are in `ASSUMPTIONS.md` and
want confirming, and what is built and what is not is the last two sections of it.

The gate path makes no network call, which is the property the verdicts rest on.
`xeno enforcement check` is the one command that does, because asking a host what it
enforces is the one question the repository cannot answer about itself.

    go build ./cmd/xeno
    go test ./...

    xeno init            [--vendor] [--project OWNER/REPO] [--model ID] [--language TAG]
    xeno intent start    --for ISSUE [--intent KEY]
    xeno phase start     --intent KEY --phase NN [--evidence-from DIR] [--export]
    xeno phase finish    --intent KEY --phase NN [--summary PATH|-]   writes digest.md
    xeno gate run        --intent KEY --phase NN [--base REF --head REF] [--evidence-from DIR]
    xeno gate approve    FINDING --intent KEY --phase NN --by WHO --reason TEXT
    xeno gate override   FINDING --intent KEY --phase NN --by WHO --reason TEXT
    xeno obligation close FINDING --intent KEY --phase NN
    xeno learning record --intent KEY [--phase NN] --category C --observation T --proposal T --target P
    xeno assumption record --intent KEY --phase NN --text TEXT --origin WHERE --confidence HOW [--resolves KEY]
    xeno assumption confirm ID --intent KEY --by WHO
    xeno assumption reject  ID --intent KEY --by WHO
    xeno gate verify     [--intent KEY]
    xeno enforcement check [--branch NAME]
    xeno evidence attach --intent KEY --phase NN --from DIR
    xeno section set     SECTION --intent KEY --phase NN [--file PATH] [--tool-version V]
    xeno intent status   [--intent KEY] [--all]
    xeno intent close    --intent KEY --reason TEXT
    xeno check commit-message [--pattern NAME] [--file PATH]
    xeno cost turn
    xeno version

`xeno intent start` takes the one thing it cannot derive, which is the issue the intent
belongs to: a key, `176`, or as much of the qualified id as the configuration does not
already hold, up to the whole `github.com/triplem/xeno#176`. The key continues the
sequence the intents directory holds, and `--intent` names it instead where there is no
sequence to continue or the key is not the next one.

`xeno learning record` writes the record section 10 owes at the end of every phase, and
without `--phase` the one an intent owes when it closes. The four keys are the agent's
and the header is the runner's; `--no-finding` is the empty record said rather than left
out, and a second call appends. It was the last artifact of this process that no command
wrote (A82).

The harness reports its own version, which is the one field of section 5 the runner
cannot know: section 7 forbids it branching on the harness, and a value nobody produced
is worse than an absent one (A35, A80, A81). `XENO_HARNESS_VERSION` says it for a whole
session, which is the channel section 7 lists, and `--tool-version` overrides it for one
run. Either reaches `output.md` on a phase's first `section set`, and `phase finish`
copies it into the digest from the artifact beside it. Said by neither, the field is
absent and G-Schema reports it missing, which is what every artifact written before this
carries.

`xeno intent status` without an intent lists the last ten by creation, `--all` every one
of them. `xeno cost turn` reads a hook's JSON on stdin and attributes the turn to the
phase the run marker says is open. Every command takes `--root DIR`, which defaults to
the working directory.

Every command that changes something ends by naming the next step of the working
sequence, and so does `xeno intent status`. `--no-next` leaves it out; the commands a
pipeline or a git hook runs never print it.

Xeno collects no telemetry. Nothing it writes leaves the repository it writes in, and
there is no endpoint for it to leave towards.

## The trail in this repository

`.xeno/intents/` holds Xeno's own process artifacts, in two shapes. The intents up to
`XENO-0106` run P0 and stop there, which was deliberate rather than abandoned: `M0.md`
runs one intake per work package and goes no further, so an intent whose P1 to P5 read
`not-started` is complete for what it set out to record. `M0.md` has that reasoning and
the sequence, `ASSUMPTIONS.md` has it as A20.

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
