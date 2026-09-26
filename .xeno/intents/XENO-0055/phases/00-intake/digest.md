---
intent: github.com/triplem/xeno#55
phase: 00-intake
created: "2026-09-26T12:40:09Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2cb8638
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 4fcea10f04e68b1e3c5b4933dd040b0318267c8b8dbcf3e2a71387021853af40
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The issue had already done the comparison field by field, so the work was reading
section 8 against the code and not deciding what the record should be. The registers
were checked for entries before anything changed, which is what made the migration
criterion answerable in one command: every one of them is empty.