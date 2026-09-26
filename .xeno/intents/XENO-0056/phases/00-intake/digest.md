---
intent: github.com/triplem/xeno#56
phase: 00-intake
created: "2026-09-26T11:14:43Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f5114fa
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 8aaa5d10685e4dae2389f196abd112a9ea923e9194362ff66ef3520a8a1c8f2f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The proxy was running before this started, so the questions were answered by probing it
rather than by reading its documentation. The one probe that changed the answer is the
one that failed: custom header names are dropped silently, which an export written from
the plan's wording would have hit only once somebody queried the rows. The temporary
deployment the probes needed was removed afterwards and the model list is as it was.
