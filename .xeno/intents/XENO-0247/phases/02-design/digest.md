---
intent: github.com/triplem/xeno#242
phase: 02-design
created: "2026-10-05T12:12:07Z"
schema_version: "1.0"
runner_version: dev+3f5ad34.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8b81d972ee520502e9fedac2a5355586d1947ada78138ba21546bdb3444d4c37
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
A new file using `exchange.go`'s two helpers, no phase parameter because the gate's guard
leaves one phase, replace at the found index, and the three refusals ordered cheapest first
so a malformed call cannot half-amend an artifact.

#242's two wrong bullets are rejected with the reasoning that found them, and D-6 is named as
where the confusion came from.
