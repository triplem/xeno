---
intent: github.com/triplem/xeno#143
phase: 00-intake
created: "2026-09-29T20:26:57Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+530ce03.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 48a5f7a77d1577b326325f6e71b2c3decf9ceb2e347f51fae3684c4846a22f44
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Four things in internal/enforcement are GitHub's and not the domain's — the URL shape, the Accept header, the
JSON field names and the meaning of a 403 — and the fourth is why this is not a URL swap, since reading 403 as
"not available on this plan" is a fact about one host's pricing. internal/runner/enforcement.go:50 completes
the assumption by defaulting to api.github.com past configuration that already exists and is already read.
A65 settled the target, so this intake partitions one file into what moves, what stays and what is retired
rather than choosing a shape. Tracker is deliberately excluded: #104 leaves it out of the done-when this
intent answers, and an interface with no adapter and no caller is what #97 declined to write.
