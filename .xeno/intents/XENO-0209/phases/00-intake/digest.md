---
intent: github.com/triplem/xeno#109
phase: 00-intake
created: "2026-09-29T18:35:45Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+3ec2429.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d7c7c2ca84561cbcd96b6b8880bd20480a56902723790393aef9e0b244a1132c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The check belongs where the hash is computed, and that is the abandoned case alone: an intent level
artifacts_hash exists only in the gate.yaml intent close writes, because a merged intent's record is its P5
phase. So CompleteOnClose is both the mode and the only invocation, and the two cannot fall out of step.
What this does not make true is worth naming: checked means reported, so a stray file still enters the hash
of the run that reports it, which is the compromise the phase level has lived with since it was written.
