---
intent: github.com/triplem/xeno#235
phase: 00-intake
created: "2026-10-06T15:21:29Z"
schema_version: "1.0"
runner_version: dev+8decfdc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 11db835658e8a5253ec394ab3188e2599e97e2f997d3f485810eb4cdabad07d1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`budget` is called from `schema` and `result` fails a check on any finding, so a recorded context
over its budget turns G-Schema red and stops the phase. Section 5 says that check is "deliberately a
finding and not a red gate in the sense of stopping work", and the code quotes the sentence it
breaks. No shape exists for it: the four check results are fixed and a finding fails its check.
`drift` was the one route that would have used an existing mechanism and is itself unimplemented —
0 of 443 sealed gates carry one. #235 settled on an advisory finding, which needs section 5 first.
