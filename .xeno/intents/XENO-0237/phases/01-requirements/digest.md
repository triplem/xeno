---
intent: github.com/triplem/xeno#202
phase: 01-requirements
created: "2026-10-03T19:24:36Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bb7dbda3783f9c05f328efc51b287be157611f2a85f07c52879bca14ad06413c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The criteria restate #202's "done when" and add one thing it left implicit: a reader that
cannot fail is a worse answer than none, so the list has to be able to say that about a
clause rather than ticking it. The constraints admit what the deliverable cannot do,
which is stay true on its own, and settle for dating the pass.
