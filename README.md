# xeno, core

The network free core of the Xeno runner: hashing, gates, the phase sequence and
evidence attachment. Built by hand before M0; the assumptions made on the way are in
`ASSUMPTIONS.md` and want confirming.

    go build ./cmd/xeno
    go test ./...

    xeno phase start     --intent KEY --phase NN [--evidence-from DIR]
    xeno phase finish    --intent KEY --phase NN
    xeno gate run        --intent KEY --phase NN [--evidence-from DIR]
    xeno gate verify     [--intent KEY]
    xeno evidence attach --intent KEY --phase NN --from DIR
    xeno intent status   --intent KEY
    xeno version

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
| WP5 | question resolution by decision or confirmed assumption | `TestQuestionResolved…` |
| WP6 | pending declarations, pulled attachment, provisional verdicts, tampered attachments rejected | `TestPendingEvidence…`, `TestNextStartAttaches…`, `TestTampered…` |
| WP7 | sequence enforcement, run marker, G-Questions from P5, gate run attaching for P5 | `TestOutOfOrder…`, `TestSecondStart…`, `TestGateRunAttaches…` |

Two deliberate mutations were run against the suite, removing the sequence check and
letting attachment write into a sealed artifact; both were caught.
