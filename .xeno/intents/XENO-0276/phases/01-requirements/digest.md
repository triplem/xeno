---
intent: github.com/triplem/xeno#325
phase: 01-requirements
created: "2026-10-08T13:51:01Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5f9b8f51b0c47c24c41f19bad2673859bbe6b438b2aaa573570ef3fd4e048143
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Five criteria for two lines of a skill. The fourth is the one that makes the change hold:
`internal/plugin` already requires every command a skill names to exist in the dispatch
table, so the new lines are held against the binary rather than being prose that can drift.

The fifth says what this is not: no gate, no field, no runner behaviour. The defect is a
missing sentence.

One section of the four.
