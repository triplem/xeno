---
intent: github.com/triplem/xeno#107
phase: 04-verification
created: "2026-09-28T18:57:33Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+907c220.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: e4538f47f814d6268ff0c24ef7ef6af07d47aa6bb36676d04b898ae336ab2ac0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
evidence:
  - kind: test-report
    job: go-test
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | What proves it | Fails without the change |
|---|---|---|
| AC1 `Status` refuses, naming the gate | `TestAFailWithoutAFindingIsRefusedRatherThanRead`: an error, an empty status, and `G-External` in the message | yes |
| AC2 the callers carry the refusal out | `TestAMalformedCheckInAStoredVerdictRefusesTheDecision`, through `Decide` and `rewriteStatus`. `phase finish` and `gate run` reach `Status` through `evaluate`, which returns its error unchanged, and `IntentClose` likewise; both are read rather than tested | yes |
| AC3 `fail` with a finding is still red | `TestAFailWithAFindingIsStillRed` | no, and must not |
| AC4 the three no-finding results are undisturbed | `TestTheResultsThatCarryNoFindingAreUndisturbed`, over `pass`, `pending` and `not-implemented` | no, and must not |
| AC5 the id collision refusal is unchanged | `TestTheIdCollisionRefusalIsUnchanged`, asserting its own wording | no, and must not |
| AC6 everything green stays green | `go test ./...` and `./xeno gate verify` over 67 verdicts | no, and must not |
| AC7 a hand edited `gate.yaml` is refused | `TestAMalformedCheckInAStoredVerdictRefusesTheDecision`, which writes the malformed check into the stored verdict | yes |
| the ordering, that a malformed check is not masked | `TestAMalformedCheckIsNotMaskedByAWellFormedOne`, which is P2's decision about position in the loop rather than an AC | yes |

Four rows fail without the change and four must pass with and without it. The right hand
column is the part of this mapping worth reading.

<!-- xeno:section:results -->
## Results

`go test ./...` passes on every package. `gofmt -l .` outside `vendor/` prints nothing,
`go vet ./...` is silent, and `./xeno gate verify` recomputes 67 verdicts and matches,
which is the check that no gate of this runner reaches the new refusal.

Against the tree with `internal/gates/gates.go` reverted, three of the nine tests fail
and the transcript carries that run: the refusal returning `green`, a malformed check
accepted beside a red one, and a decision written onto a verdict carrying a `fail` with
no finding. The other six pass there, because they guard what the rule must not disturb.

Two criteria are proved by reading. AC2 names three callers and one of them is tested:
`phase finish` and `gate run` both reach `Status` through `evaluate`, which returns its
error to the caller unchanged, and `IntentClose` does the same at its own call. A test
per caller would assert the same two lines three times.

The one number worth recording: `Status` now refuses a state nothing in this repository
can produce, and every verdict in it still matches. That is what a latent hole closed
without collateral looks like.

<!-- xeno:section:gaps -->
## Gaps

**The state has no producer.** No gate of this runner can emit `fail` without a finding,
so the case exists only in the tests and in whatever an external gate or an editor
writes. That is the ordinary condition of a guard placed before its data arrives, and it
means the fix is correct against the shape and unproven against a real foreign gate. WP4
brings the first one.

**Two of the three callers are read, not tested.** `phase finish`, `gate run` and
`intent close` all pass `Status`'s error outward, and only the `rewriteStatus` path has
a test. The untested two share the same two lines with the tested one, which is an
argument and not a proof.

**A refusal stops the pipeline where a red verdict would record a failure in it.** For a
malformed check that is the point, and for a project whose external gate is occasionally
malformed it will read as Xeno breaking rather than as the gate being wrong. The error
names the gate, which is the whole mitigation. Whether that is enough is something only
a project with a real external gate can say.

**Nothing addresses a gate that reports `pass` while something failed.** Undetectable
here and everywhere else in the local runner, and unchanged by this.

**The evidence is a local run.** Third intent in a row for the same reason: the pipeline
uploads no artifact, so the declared test report is a transcript bound by hash through
the stand-in directory, and `attached.yaml` cannot distinguish that from a pipeline
artifact.
