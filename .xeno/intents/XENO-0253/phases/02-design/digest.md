---
intent: github.com/triplem/xeno#225
phase: 02-design
created: "2026-10-05T17:15:46Z"
schema_version: "1.0"
runner_version: dev+9eec4ed.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8b3c46afc2d20ac0c27b18769288c96da231efb0818fab5bbf18cbac398b47bc
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The condition is the artifact disagreeing with the lock, not the marker being gone, because a
missing marker is ordinary and only a phase under way has an `output.md` naming a
`context_hash`. The comparison is G-Schema's own, so the refusal and the finding cannot disagree.

Idempotence was the attractive alternative and is rejected on honesty: a command that silently
declines reports success for an act it did not perform, and leaves the person's model wrong.
