---
intent: github.com/triplem/xeno#167
phase: 04-verification
created: "2026-10-01T16:58:32Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15693cf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f8aae81434f84d87589f815a574a6a3c07d808a3c752fe10d40e15e727079997
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
407 cases pass, `gofmt` and `go vet` clean, `gate verify` at exit 0 over 204 verdicts. The external
package runs in 1.2 seconds where it took thirty before `WaitDelay`: the process was killed on time and
the read waited for a pipe a grandchild still held, which is the measurement behind A75. A modified gate
did not run — the test replaces the file with one that would touch a marker, and the marker does not
exist, which is WP4's done-when with the second half asserted rather than assumed. An external check
reaches the verdict through the runner, carries `provenance: external`, takes a release, and the next
run produces the same finding with the same id and no decision: #66's rule finally has a second path
into the invariants, which is what the comment there has said since it was written. Four ways a foreign
command can behave badly are each a fail with a cause, and the structural one is the non-zero exit with
nothing to say, because `Status` refuses a failing check with no finding. Gaps: there is no sandbox and
none is pretended; `gate verify` runs the commands too; no external gate has ever run outside a test,
which makes this the first WP4 piece with no demonstration against the real tree; A76 is a promise with
one holder; the timeout is one number for every gate; and a tool's standard error goes nowhere.
