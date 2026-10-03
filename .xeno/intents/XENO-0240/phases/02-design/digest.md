---
intent: github.com/triplem/xeno#215
phase: 02-design
created: "2026-10-03T20:23:09Z"
schema_version: "1.0"
runner_version: dev+b626f1a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7abd887b3a02d4e2eec1a9d7e1a0be5a824952a49223e5fe5efbb02764faa9aa
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two decisions carry it: key on the verdict rather than the lock, which keeps the dead-run
recovery open and is the better test on the merits, and name both ways forward in the
message. The alternatives section records the one that would have shipped a regression —
refusing on the lock — and the one that was already half-written, a second derivation of the
changed set, which was deleted because two derivations of one answer can disagree.
