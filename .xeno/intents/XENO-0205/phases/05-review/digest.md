---
intent: github.com/triplem/xeno#65
phase: 05-review
created: "2026-09-29T11:42:12Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+edd2b2e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 55ae26c89b639d60882a03d61373ca67883e80defaca369d69a76f22f121f631
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The risk worth carrying is specific rather than general: the hook's command is a relative ./xeno, resolved
against whatever directory the harness runs hooks from, and if that is not the repository root the hook
silently does nothing and no gate will say so, because a phase without a cost record is complete. That is
the shape of everything optional and unmonitored, and it is the first thing the next session can check.
