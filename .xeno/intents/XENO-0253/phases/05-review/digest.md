---
intent: github.com/triplem/xeno#225
phase: 05-review
created: "2026-10-05T17:30:07Z"
schema_version: "1.0"
runner_version: dev+6adc0f9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: afca0b01e8db7260910dd042a887f1c9b704c8df2098cbfc1389fa7111b97b91
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`phase start` refuses a phase already under way, keyed on the artifact because the marker is
machine-local and expires, and a phase with an artifact now reports `running` everywhere — which
closes #225's first gap as a consequence of making the suggestion agree with the refusal.

Starting over takes the phase directory now, which changes what #215 documented; the messages,
its comment and its test were corrected together.
