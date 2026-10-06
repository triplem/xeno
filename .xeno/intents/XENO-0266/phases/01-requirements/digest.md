---
intent: github.com/triplem/xeno#267
phase: 01-requirements
created: "2026-10-06T19:57:28Z"
schema_version: "1.0"
runner_version: dev+1d61fa2
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7ef8ab81da7d67d07dee7b8d0622bf81a40c89370e130463728856d88d9aac84
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Fourteen criteria. Section 5 gains a paragraph where the lock is described, saying that at P0
`files` is empty, why the order is forced, why nothing fills it afterwards and what the intake
still binds; the schema block's `files` line points at it. Section 5's budget clause and section
7's staleness clause each say the check is judged from P1 on, the latter being explicit that this
is not a third limit against noise. Two code comments explaining the empty list as a phase older
than #217's rule are corrected. Nothing changes behaviour, so there is no test; P0's lock does not
start recording files, which is the decision and not a postponement.
