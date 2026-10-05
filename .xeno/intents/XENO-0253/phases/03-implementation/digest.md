---
intent: github.com/triplem/xeno#225
phase: 03-implementation
created: "2026-10-05T17:26:26Z"
schema_version: "1.0"
runner_version: dev+6adc0f9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: acfd289f9257410672f75b9f7213ca782efe2e1171c052adadfb76ccf39d5ea5
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The guard is the artifact's existence, not a comparison against the lock: before a second start
the two still agree, so P2's condition would have passed exactly when the damage was about to be
done. The reproduction is what found it.

A phase with an artifact now reads `running` without its marker, which closes #225's first gap
and is why the suggestion no longer offers the refused start. Starting over takes the phase
directory, which changes #215's documented escape.
