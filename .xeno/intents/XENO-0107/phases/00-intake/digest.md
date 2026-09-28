---
intent: github.com/triplem/xeno#107
phase: 00-intake
created: "2026-09-28T18:50:02Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+907c220.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: b57eade4bafd4e2dc5b281ce4b4139660e4924a9db2db935f68e9a0ee21a9222
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The issue left one thing open, whether such a check is red or refused, and named the
second writer as hypothetical. Reading the callers found it already in the tree:
gate.yaml sits outside artifacts_hash, so a verdict can be edited without staling a
hash, and rewriteStatus derives a status again from exactly that file. That decided the
location as well as the answer, and the choice of Status over Invariants follows from
which paths each one covers.
