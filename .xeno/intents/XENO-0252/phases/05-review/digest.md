---
intent: github.com/triplem/xeno#205
phase: 05-review
created: "2026-10-05T17:01:55Z"
schema_version: "1.0"
runner_version: dev+9eec4ed.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f6e1917fc60243e475f2fe47872e66768cd68a19853d81ec2ad6862b3dfbec83
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The last variable in section 7's environment with no reader has one: `LocalDir` and `LocalPath`
resolve it, six literals are gone, and unset resolves to exactly today's paths, asserted over the
strings and over the disk.

Why this one and not `XENO_PLUGIN_ROOT` is A97: nothing under the local data location is hashed,
so no verdict can be made to depend on the environment.
