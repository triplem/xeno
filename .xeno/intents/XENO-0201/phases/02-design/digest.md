---
intent: github.com/triplem/xeno#120
phase: 02-design
created: "2026-09-28T20:30:17Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6e72fed.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 5c93b8fbddef9da37397f458c7273664bb9df6ecd66fa7cf7c611fc18fb62208
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The ordering decision is the one that matters: the digest is inside artifacts_hash, so
it is written before the gates run, and the rejected variant of writing it afterwards
would seal a hash over a tree that lacked the file. The alternative worth naming is
writing secrets_hash over an empty filter, which makes every field present and every
gate green and asserts a filter that does not exist. It is the option the schema invites
and the one this intent exists to refuse.
