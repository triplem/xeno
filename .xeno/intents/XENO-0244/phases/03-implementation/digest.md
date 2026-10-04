---
intent: github.com/triplem/xeno#221
phase: 03-implementation
created: "2026-10-04T09:39:35Z"
schema_version: "1.0"
runner_version: dev+24becc3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d001d5c46f78496c10bb57a7ca07db2b99176f67c91de60c44f34fe68a5d83ad
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The implementation does the two steps in the order the decision requires, in one commit,
so no tree exists in which the check is present and the fifteen still read `success`.
First the correction: one word in fifteen `evidence/attached.yaml` files, all fifteen
read beforehand and identical in shape, fifteen insertions and fifteen deletions, with
`pipeline: local` left as it is because it is accurate about what happened. Then the
check: `attachedResult` in `internal/gates/gates.go` judges a recorded value against
`model.EvidenceResults` and returns its finding on `attached.yaml`, called from the
branch of `evidence` where an attachment has been resolved, and judging a value while
leaving an absence alone, which is what keeps `other/trivy-db` green. Then the refusal:
`unrecordable` in `internal/evidence/attach.go`, running for every entry and before the
two binding checks, because a result is wrong whichever way an entry is bound. Six new
tests, three per package, with the accepted values read from
`testdata/evidence-declaration.yaml` rather than from the constants, which is that
file's own reason for existing. The figure the whole order exists for was taken three
times, each at exit 0 over 342 verdicts: before the change, after the correction, and
after the gate — so the reading of Appendix B held, `artifacts_hash` does not descend
into `evidence/`, and no committed verdict was contradicted. Build, `gofmt`, `go vet`
and the suite are clean. Three deviations are recorded. `Result.Unbindable` became
`Result.Declined`, which the design did not call for and which writing it showed to be
the reason two callers each printed a remedy that was right for one reason out of three;
one expected substring in a test follows the reworded refusal; and the refusal could not
be transcribed against this repository's own P4 yet, because the phase had no artifact,
which is the exact error #208 was filed about and is taken in the verification phase
instead. Read in this phase: all fifteen files, `evidence` and `Collect` in `gates.go`,
`attach.go` whole, the two callers of the renamed field, and both test files for the
idiom their fixtures already have.
