---
intent: github.com/triplem/xeno#235
phase: 00-intake
created: "2026-10-06T17:51:18Z"
schema_version: "1.0"
runner_version: dev+5276f4b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 67975d39df0a5d8af4550e705ccf6a8e00d945f87ec51bfa7a3fc57bc5ab7120
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`budget` is called from `schema` and `result` fails a check on any finding, so a context over its
budget turns G-Schema red and stops the phase, where section 5 now says twice that it must not. The
shape arrived in XENO-0263 and nothing writes or reads it: `model.Finding` has no `Advisory`, and
`result` counts findings without looking at one. The check also reads nothing at P0, which is #267
and a different fault — both have to close before the clause is exercised once.
