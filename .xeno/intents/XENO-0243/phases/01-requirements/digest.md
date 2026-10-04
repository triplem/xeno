---
intent: github.com/triplem/xeno#208
phase: 01-requirements
created: "2026-10-04T08:26:49Z"
schema_version: "1.0"
runner_version: dev+49f2794
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f88a71e86c72c348270949328dbf227bc0e9cb85426abffbcce273ba402c1894
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The requirements phase turns the three answers into criteria and records them as D-1 to
D-3, which is the first use in this repository of the two commands #188 built. D-1 takes
the whole of the scope: the writer, `produced_by` and `format` added to the model, four
pending items declared in this intent's own P4, and a test report published by
`xeno.yml` so that one exists to declare. D-2 leaves `id` out, because the attach and
G-Evidence key an item by kind and job and a second identity both readers ignore is how
two keys for one thing come to disagree. D-3 leaves "declared, never inferred" exactly
where `CLAUSE-READERS.md` has it, a property with no reader, rather than contradicting
it in code. The criteria are written so that each is a test: the three forms of a
declaration, the six refusals, the computed hash no flag can supply, the single
authority on shape, the pair declared once, and G-Evidence failing for a real reason
when the copied file is edited or removed. The constraints name what the command may not
do, and the narrowest of them is the write path: the declaration goes inside
`artifacts_hash` and the attachment outside it, and neither command touches the other's
file. Two assumptions are recorded rather than absorbed. A-001: the `--file` form cannot
tell a report pulled from an artifact store from one a developer typed, so under
`evidence.source: ci` the regime in `context.lock.yaml` is as far as the tool can go.
A-002: `produced_by` on a pending item names the command the job will run, because there
is no local run to name and the manifest carries no such field. Read in this phase:
section 4's field list and closed sets again against each criterion, section 6's note on
what a phase start may write, section 7's exit code staircase,
`internal/evidence/attach.go` for the words of the uri refusal, and
`internal/runner/exchange.go` for the shape a refusal before a write takes.
