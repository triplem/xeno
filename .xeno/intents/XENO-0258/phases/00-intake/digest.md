---
intent: github.com/triplem/xeno#256
phase: 00-intake
created: "2026-10-06T08:34:36Z"
schema_version: "1.0"
runner_version: dev+5163d1b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8e8a33c45af8d19a088cd51c8a17c446489e675d7130bd8c93159b80ef470454
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
A false negative written into a phase's artifact is sealed with it, and this session produced one: a
`git check-ignore` run against a path that had just been deleted reached two sealed phases, a pull
request and two reports before anything tested it with the path present.

What makes it this project's convention rather than general advice is where a finding goes. Section
11 says what is sealed is never rewritten, so the mistake becomes permanent rather than corrected.
