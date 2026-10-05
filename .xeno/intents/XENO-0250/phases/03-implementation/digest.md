---
intent: github.com/triplem/xeno#229
phase: 03-implementation
created: "2026-10-05T15:29:28Z"
schema_version: "1.0"
runner_version: dev+2d997e0.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cbaea306c782118a36a55d7654b02e3cf8cd22bb4e27ff9df15a37c8ff8881bd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The gate reads the consequence, the writer reads the recommendation, and `QuestionAsked` calls
`QuestionShape` so the shared half cannot drift. Eleven tests across the two packages, and the
asymmetry is pinned as an assertion rather than left to a comment.

A fixture that had carried the forbidden shape since WP5 is corrected, which broke three tests
before it fixed them, and `internal/gates` had no question tests at all until this.
