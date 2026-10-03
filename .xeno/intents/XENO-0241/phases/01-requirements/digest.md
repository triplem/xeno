---
intent: github.com/triplem/xeno#206
phase: 01-requirements
created: "2026-10-03T20:44:50Z"
schema_version: "1.0"
runner_version: dev+8574810
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ceb3f25885139b68b7eb77825a75163b65ed2d62921e3902b18c422feda20525
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The criteria fix one command, one reader, one step and the two shipped wrappers, and fix
the exit codes against A11's staircase: 1 for an intent that stopped, 2 for a range that
could not be read. The constraint worth stating is the cost: a pull request is red until
its trail reaches a decided P5, which is what every other gate already does. The state is
read from `state` rather than derived again, which is the lesson XENO-0240 recorded.
