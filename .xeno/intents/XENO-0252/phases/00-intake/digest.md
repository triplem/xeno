---
intent: github.com/triplem/xeno#205
phase: 00-intake
created: "2026-10-05T16:26:31Z"
schema_version: "1.0"
runner_version: dev+97d07da.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 703eded3d3a399750676b4a085fb3d290ec17617c83883f93e60b96624e3a1a5
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`XENO_PLUGIN_DATA` is specified, exported by the entry point and read by nothing; six places
write `.xeno/local/` as a literal instead, one of them twice over.

Nothing under that directory is hashed, so reading the variable cannot make a verdict depend on
the environment — which is the objection A89 raised against the plugin root and the reason it
does not reach this one.
