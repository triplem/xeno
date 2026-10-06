---
intent: github.com/triplem/xeno#235
phase: 02-design
created: "2026-10-06T15:21:43Z"
schema_version: "1.0"
runner_version: dev+8decfdc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bbdd77200932d2b574c601d3bca23e907ba251c599e119991315df1883bca2d7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
D-1: `advisory` is a key on a finding and not a result on a check, a boolean and not a severity, and
the bound is its own paragraph because no code will enforce it. It goes after the four-values
paragraph rather than after the one the draft named, which taken literally splits the status
derivation in half. Seven alternatives weighed, among them implementing `drift` — the one that reads
best from a distance and is unimplemented in a shape that does not fit a byte count.
