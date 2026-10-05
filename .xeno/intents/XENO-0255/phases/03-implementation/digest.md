---
intent: github.com/triplem/xeno#247
phase: 03-implementation
created: "2026-10-05T19:39:30Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a8e4506fd5b5135e62a53f3b2022ea4825236e835d4c939ab78b6092dc95c257
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
One audit row, one `CLAUDE.md` paragraph, one example rule. The mapping row keeps saying `nothing`
and the example is named in prose, which is the prohibition criterion 2 set.

The example is inert by construction and by measurement: `rules.Load` cannot reach `examples/`,
and `gate verify` is at exit 0 over 408 verdicts with `rules_hash` unchanged throughout.
