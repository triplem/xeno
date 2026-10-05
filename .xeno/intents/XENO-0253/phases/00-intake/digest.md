---
intent: github.com/triplem/xeno#225
phase: 00-intake
created: "2026-10-05T17:14:39Z"
schema_version: "1.0"
runner_version: dev+9eec4ed.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6551e21cd2b7dfc7302b2459e67f04649adb4c424d22916241a0708651c3cbc6
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
A phase whose marker is gone but whose sections are written can be started again. The lock is
rewritten, the artifact's `context_hash` stops matching it, and `phase finish` goes red on
G-Schema two commands after the mistake.

Reproduced with a moving clock, which is what makes it visible: the fixture pins `Now`, so the
rewritten lock is byte-identical and a test in the usual style passes against the bug.
