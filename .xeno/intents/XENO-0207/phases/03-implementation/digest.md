---
intent: github.com/triplem/xeno#134
phase: 03-implementation
created: "2026-09-29T17:59:30Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c05acae.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2ec969e5d3f2810492c214a76d8f637bdf50c022e28403642b78788caa4d59ed
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two helpers exist only because the tests needed them, and both earn it: tail makes the arithmetic checkable
without capturing output, and sprintRow lets a test build a heading and a row from one format string and
compare them, which is how the alignment becomes assertable rather than inspected. This is cmd/xeno's first
test file, which says something about #110 that three tests do not fix.
