---
intent: github.com/triplem/xeno#98
phase: 00-intake
created: "2026-09-27T09:44:22Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+5ddcaab
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: d4ea6131918e594e50d2b201c4499f46579418402a1f66e52d8673d568ee4c11
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The obvious home for these two files, the vendored plugin, is ruled out by the order in
which init does things, and noticing that first is what kept the change to one package
instead of two.
