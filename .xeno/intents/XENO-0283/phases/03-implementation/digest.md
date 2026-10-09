---
intent: github.com/triplem/xeno#346
phase: 03-implementation
created: "2026-10-09T15:20:13Z"
schema_version: "1.0"
runner_version: dev+9590797.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ceb7dfaa54dfd27113b53314d8533549915dcf6429fd43000b7baa477b8918ab
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The implementation for #346: the GitHub read searches by name and keeps the exact match, the GitLab read takes a page of a hundred, the header says what idempotence rests on, and the plugin test requires both flags and failed on main first. Suite, gofmt, vet and sh -n clean; the second run against this repository exits 0.
