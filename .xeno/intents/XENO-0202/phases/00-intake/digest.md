---
intent: github.com/triplem/xeno#120
phase: 00-intake
created: "2026-09-29T05:57:30Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ad0a764.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 8fcfd55b4b10b31e02abf6f4a6fab73b535416c78f29f67911caea87667cddc0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Section 16 turned out to carry #120's argument in the specification's own words: the filtering sits in
the runner so that a model cannot route around it, not even by accident. That makes this the one part of
the issue where prevention is real rather than refused, and it is why the review of the previous intent
called it the thing to close first. The decision that needed measuring rather than reasoning is the
honesty of by-hand: tightening it diverges all eighty-seven verdicts, so the loophole stays and is
recorded with its number.
