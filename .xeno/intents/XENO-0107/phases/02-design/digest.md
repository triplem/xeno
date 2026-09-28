---
intent: github.com/triplem/xeno#107
phase: 02-design
created: "2026-09-28T18:51:26Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+907c220.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 2611fc4b103092354187b19a1a5c296cbb6081a33cc8757795e56f206d3b39ae
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Four of the five alternatives were rejected on reach rather than on fit, which is the
pattern of this design: the question is never whether a guard is correct but which
writers it meets. Guarding twice, which the predecessor intent's issue made look like
the careful option, was rejected here because Status covers a strict superset of
Invariants, and a second guard on a subset buys a line of code and no coverage.
