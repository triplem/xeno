---
intent: github.com/triplem/xeno#247
phase: 05-review
created: "2026-10-05T19:47:55Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4307f6085a6208c4e0d448f83fd16fc38b1bc1a7a09a3d030c12757cc144332a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Three clauses, three places a person-reader is recorded, one intent. The mapping row keeps saying
`nothing` because an unenabled rule fails nothing, which is #254's correction applied one row away
rather than repeated.

This phase and P4 were started over after the first P4 bound a test report's hash while the suite
was still writing it — the escape #225 made explicit, used in anger for the first time.
