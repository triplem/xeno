---
intent: github.com/triplem/xeno#106
phase: 00-intake
created: "2026-09-27T14:31:59Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0b3a547.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: a618da7e7909a912dd9b300251b714d0ae402318dab6113756a1144a6a07f198
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The call sites were read before the scope was written, and they decided its shape: `Attach`
is called from three places and writes no output, so a refusal it swallowed would leave the
item pending with no explanation anywhere. That turned a one line guard into a returned
result, and the second half, the gate check, came from asking what covers a file that sits
outside the `artifacts_hash` by design.
