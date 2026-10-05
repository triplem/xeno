---
intent: github.com/triplem/xeno#242
phase: 01-requirements
created: "2026-10-05T12:11:19Z"
schema_version: "1.0"
runner_version: dev+3f5ad34.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7069385d8e026615b8f40dad35dd739d01ecb40161d2d6a1690eedfb182aaf9b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Ten criteria. Two of them are the ones #242 got wrong and now state the opposite: no phase
argument, and no `gate.yaml` read. Criterion 9 is the one that matters most and costs
nothing: this intent writes its own P5 checklist with the command.
