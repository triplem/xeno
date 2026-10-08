---
intent: github.com/triplem/xeno#325
phase: 03-implementation
created: "2026-10-08T13:52:05Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a96e165f89ee9b6678fa792e8dd20201e359f67ffe44078d65dcb5b421b5b247
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
One file in two places: the two writers in the command block with the sentence that they
resolve P5 themselves and take no `--phase`, and the reason there are two commands in the
paragraph that already describes the entries.

The `--phase` sentence is worth its line: these two are the only commands in that block that
resolve the phase themselves, so a reader copying the pattern would add a flag they do not
take.

`internal/plugin` surface test passes, which is what now holds the lines to the dispatch
table.

One section of the three, `changes`.
