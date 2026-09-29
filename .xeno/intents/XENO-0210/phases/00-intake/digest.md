---
intent: github.com/triplem/xeno#110
phase: 00-intake
created: "2026-09-29T19:02:18Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+454cfbf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a39abc00d216eef899527c10ff906983ef0a1c506836483669a15fb520da73a5
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Twenty-eight of thirty functions at zero coverage, and the two exceptions are the helpers XENO-0207
extracted so a table's arithmetic could be tested — which is the shape of the whole problem: what was
testable got tested and the rest writes to package level globals. The case that justifies the intent is the
provisional one, where cmdGateVerify returns 0 because section 6 says a P4 waiting on a pipeline must not
fail verification, and that is one if away from turning every push out of a P4 red.
