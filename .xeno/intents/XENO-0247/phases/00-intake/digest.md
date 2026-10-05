---
intent: github.com/triplem/xeno#242
phase: 00-intake
created: "2026-10-05T12:10:42Z"
schema_version: "1.0"
runner_version: dev+3f5ad34.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8707d01c613b6d3a6849d419382a2aa21cae8f1a47187fe54b4ffb5e8bdb157e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`review_checklist` is the one field a gate refuses the phase without and the only one no
command writes, so every P5 in this repository was hand-edited inside `artifacts_hash`.

Reading `exchange.go` rather than copying #242's bullets settled two of them against the
issue: no `gate.yaml` check, and no phase argument.
