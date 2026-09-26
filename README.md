# xeno, core

The network free core of the Xeno runner: hashing, gates, the phase sequence and
evidence attachment. Built by hand before M0; the assumptions made on the way are in
`ASSUMPTIONS.md` and want confirming.

    go build ./cmd/xeno
    go test ./...

    xeno init            [--vendor] [--project OWNER/REPO] [--model ID] [--language TAG]
    xeno phase start     --intent KEY --phase NN [--evidence-from DIR]
    xeno phase finish    --intent KEY --phase NN
    xeno gate run        --intent KEY --phase NN [--base REF --head REF] [--evidence-from DIR]
    xeno gate approve    FINDING --intent KEY --phase NN --by WHO --reason TEXT
    xeno gate override   FINDING --intent KEY --phase NN --by WHO --reason TEXT
    xeno obligation close FINDING --intent KEY --phase NN
    xeno gate verify     [--intent KEY]
    xeno enforcement check [--branch NAME]
    xeno evidence attach --intent KEY --phase NN --from DIR
    xeno section set     SECTION --intent KEY --phase NN [--file PATH]
    xeno intent status   --intent KEY
    xeno intent close    --intent KEY --reason TEXT
    xeno check commit-message [--pattern NAME] [--file PATH]
    xeno version

Every command that changes something ends by naming the next step of the working
sequence, and so does `xeno intent status`. `--no-next` leaves it out; the commands a
pipeline or a git hook runs never print it.

Xeno collects no telemetry. Nothing it writes leaves the repository it writes in, and
there is no endpoint for it to leave towards.

## The trail in this repository

`.xeno/intents/` holds Xeno's own process artifacts, and until WP8 they stop after P0.
That is deliberate rather than abandoned: `M0.md` runs one intake per work package and
goes no further, because G-Freshness is half implemented until then and a phase beyond
P0 could not be judged honestly. An intent whose P1 to P5 read `not-started` is
therefore complete for what it set out to record. `M0.md` has the reasoning and the
sequence; `ASSUMPTIONS.md` has it as A6 and A20.

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
| WP7 | sequence enforcement, run marker, G-Questions from P5, gate run attaching for P5, the three commands that write a decision, a stale verdict refused, an abandoned intent closed and judged, the shipped commit-message pattern, the next step of the working sequence read off the state and deciding nothing | `TestOutOfOrder…`, `TestSecondStart…`, `TestGateRunAttaches…`, `TestApprovalTurnsRed…`, `TestOverrideCarries…`, `TestDecidingOnAStaleVerdict…`, `TestClosingAnAbandonedIntent…`, `gates/patterns_test`, `TestTheSuggestion…`, `TestAnOverrideIsOwed…` |

Two deliberate mutations were run against the suite, removing the sequence check and
letting attachment write into a sealed artifact; both were caught.
