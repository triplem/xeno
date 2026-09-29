---
intent: github.com/triplem/xeno#110
phase: 01-requirements
created: "2026-09-29T19:02:49Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+454cfbf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 59a69283634bd0c19cd5369128f8d166e998ebf75e9cfea2accc8a9893e9659d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
AC9 is the criterion the refactor needs and the issue did not ask for: fifty-five edits fail by sending
something to the wrong stream, and an exit code would not notice, so every assertion names the stream it
read. AC3 is the one the intent exists for, since a provisional verdict exiting 0 is what the CI wrapper
depends on and what section 6 says must not fail.
