---
intent: github.com/triplem/xeno#325
phase: 04-verification
created: "2026-10-08T13:53:18Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a4be30412ac22b4c844a1a0b72b8826281f5ff30e8242f265933321f1241f60f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Five criteria, five met, and the fourth took two attempts to check.

The first sabotage renamed the command to `xeno review lensX` and the surface test passed,
which looked like the criterion being false. The check was wrong, not the test: skills
commands are read with `^ {4}xeno ([a-z-]+)(?: ([a-z-]+))?`, so an uppercase letter ends the
match and `lensX` was read as `lens`, which the table has. `xeno review note` fails it as it
should.

So the lines are held to the dispatch table, and the near miss is a property of the pattern.
A command differing from a real one only by case would pass unnoticed — recorded as this
phase learning and not fixed here.

Two sections, which is what this template requires.
