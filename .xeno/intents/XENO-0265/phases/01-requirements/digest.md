---
intent: github.com/triplem/xeno#235
phase: 01-requirements
created: "2026-10-06T17:52:01Z"
schema_version: "1.0"
runner_version: dev+5276f4b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9fadfaab33a7554b87ba785dfd90e94945c6681c56a733e53dba1566cc267753
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Thirteen numbered criteria. `Advisory` on `model.Finding` with `omitempty`; `result` failing only on
a finding that is not advisory; `budget` marking both of its own. A check carrying one of each still
fails, which is the case a count of findings cannot express. A phase whose only finding is the
overrun comes out green and the next phase starts, which is what the clause means in this runner.
Not #267, not any other check, not `drift`, not a fifth result.
