---
intent: github.com/triplem/xeno#128
phase: 00-intake
created: "2026-09-29T09:29:37Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+38430d4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 92d71d185fc92d909e8f6c5759636b9bdf932a7e5780bef3d8b78da72256420c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The scan's first run reported forty findings and thirty-nine were this repository's own hash fields:
sixty-four hex characters after a name ending in _hash reads to generic-api-key as a key. That is the
same false positive that kept the rule out of the shipped filter in #127, arriving from the other
direction, and it turned a one line workflow into a workflow with a configuration file whose allowlist
names four shapes this repository contains.
