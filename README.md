# xeno, core

The network free core of the Xeno runner: hashing, gates, the phase sequence and
evidence attachment. Built by hand before M0; the assumptions made on the way are in
`ASSUMPTIONS.md` and want confirming.

    go build ./cmd/xeno
    go test ./...

    xeno phase start    --intent KEY --phase NN [--evidence-from DIR]
    xeno phase finish   --intent KEY --phase NN
    xeno gate run       --intent KEY --phase NN [--evidence-from DIR]
    xeno evidence attach --intent KEY --phase NN --from DIR
    xeno intent status  --intent KEY

Coverage against the implementation plan, by acceptance criterion:

| Package | Covered here | Test |
|---|---|---|
| WP1 | hashing to Appendix B, finding ids, derived status, decisions carried forward by id, the sealed invariant | `hashing_test`, `TestDecisionCarriesForward…`, `TestNoCommandButFinish…` |
| WP5 | question resolution by decision or confirmed assumption | `TestQuestionResolved…` |
| WP6 | pending declarations, pulled attachment, provisional verdicts, tampered attachments rejected | `TestPendingEvidence…`, `TestNextStartAttaches…`, `TestTampered…` |
| WP7 | sequence enforcement, run marker, G-Questions from P5, gate run attaching for P5 | `TestOutOfOrder…`, `TestSecondStart…`, `TestGateRunAttaches…` |

Two deliberate mutations were run against the suite, removing the sequence check and
letting attachment write into a sealed artifact; both were caught.
