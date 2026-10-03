---
intent: github.com/triplem/xeno#202
phase: 02-design
created: "2026-10-03T19:25:12Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e12ea06c99772542d0453a615454d618cb6acd241acf33fcb553943bc995f995
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The design's one real choice is four kinds rather than two, and it comes from the case
the project already has: `plugin_version` had a reader and was meaningless, so "read"
and "unread" cannot carry the answer. The rejected gate is recorded at length because it
is the obvious response to this issue and the reason against it — a gate parsing prose
fails for reasons about the parser — is not obvious.
