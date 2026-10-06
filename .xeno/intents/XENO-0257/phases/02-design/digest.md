---
intent: github.com/triplem/xeno#257
phase: 02-design
created: "2026-10-06T07:53:39Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1854ba1f9621b50a769dcddc5868407b93829f5f51862a37197afad069939e1c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Six of seventeen stay required and each of the eleven drops carries its reason in the file.
`required: false` everywhere rather than deletion, because the flag is the only thing `Missing`
reads and keeping ids defined leaves `section set` and the rules working.

`release-notes` stays required because a shipped checked rule reads it — a candidate dropping it
would be a trap — and `acceptance-criteria` and `test-mapping` stay by decision rather than by a
reader, which their files say so nobody infers a gate.
