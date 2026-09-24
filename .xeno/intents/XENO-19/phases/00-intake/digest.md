---
intent: github.com/triplem/xeno#19
phase: 00-intake
created: 2026-09-24T18:11:27Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: b43f326d64e7109fe30a5b1b640650298caaeb188eace6de0d0a5f182404a749
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The move was checked against the module proxy before anything was written: v3.0.5 exists,
was published on 2026-07-26 from github.com/yaml/go-yaml, and its go.mod retracts v3.0.0
and v3.0.1 because those tags came from the old path and do not match the new module. The
same file showed the `go 1.16` directive that makes an explanation in SUPPLY-CHAIN.md
stale, which is why that correction is in scope rather than discovered later.

No secret filter exists, so nothing filtered this text.
