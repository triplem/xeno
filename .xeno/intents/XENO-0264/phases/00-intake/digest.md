---
intent: github.com/triplem/xeno#258
phase: 00-intake
created: "2026-10-06T17:29:07Z"
schema_version: "1.0"
runner_version: dev+d3983d3
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 22639752e3827890b6dd12e29d7821e6ac156cbb3047eab683c70e30fde49511
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two clauses now enumerate what the artifacts may carry and nothing reads either: `model.Option` has
no `Reason`, and the plugin ships `requirements@1.0.0` where section 5 says 1.1.0. The gate comment
explaining why the mapping half is unimplemented is itself out of date, naming section 9 and a
figure of 2026-10-05. Measured on a scratch copy: bumping a template version silently removes the
`strings_hash` reader from every artifact declaring the old one — red before, green after — though
tampering is still caught by `artifacts_hash` and the successor"'s freshness, which exits 1.
