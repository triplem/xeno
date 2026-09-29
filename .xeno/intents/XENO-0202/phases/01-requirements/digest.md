---
intent: github.com/triplem/xeno#120
phase: 01-requirements
created: "2026-09-29T05:58:06Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ad0a764.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 16f274be80b0e51ad08d5df00a705f3fe95ead57f76063694e72f658c33da07c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Ten criteria, and the two that will decide the code are AC3 and AC4: the hash covers the effective set
rather than the bytes of the files, so a comment must not move it and an order must not either. AC7 is
the one an implementation gets wrong quickly, since a repository with no vendored plugin has no filter
and must still finish a phase.
