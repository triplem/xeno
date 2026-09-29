---
intent: github.com/triplem/xeno#110
phase: 04-verification
created: "2026-09-29T19:13:36Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+454cfbf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 227b7fabe08be017cc21a999e4a05a7977fd86702beaed737e882296fe8e9e7d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Three defects is the answer to whether the package needed this: two found by the refactor, a shadowed
receiver and an opts built without its writers, and one by the tests, version being named in the usage and
absent from the table. None of the three could have been noticed by a package at zero coverage. The gap that
remains is honest — twenty printers untested, and no test anywhere builds a provisional phase and verifies it
end to end, which the deviation explains and does not close.
