---
intent: github.com/triplem/xeno#120
phase: 03-implementation
created: "2026-09-28T20:32:55Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6e72fed.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: b1dca3a04bb791ea387f0c60f532127194d4d60ee61407dbc841d163d7d547c4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two of the three fields A35 called harness fields were in project.yaml all along, in the block
section 12 defines for them, and the runner was reading that file through a struct that modelled one
block of it. The writer itself was smaller than its comment: what took the thinking is what it must
not write, since section 5 gives the runner filtering as well as writing and there is no filter, so
the digest this produces carries no secrets_hash and says why in the code rather than only here.

The other discovery is that AC1's wording was wrong and the criterion behind it was right. A bare
--summary cannot mean stdin when absence has to mean no digest, so the flag takes a path or a dash.
