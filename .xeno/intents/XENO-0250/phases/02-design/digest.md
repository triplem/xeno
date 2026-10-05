---
intent: github.com/triplem/xeno#229
phase: 02-design
created: "2026-10-05T15:22:23Z"
schema_version: "1.0"
runner_version: dev+325b27a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d26a96b1e869d93886f1b828e7892406f3355d3da99dc09e0e721462d2efb26e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`QuestionShape` stays the gate's and gains the consequence check; `QuestionAsked` is the
writer's and calls it first, so the shared half cannot drift. Exactly one recommendation, with
the count in the message, and the free entry exempt from the consequence and eligible for the
recommendation.

The reason for the split is written in three places, because from any one of them the asymmetry
reads as an oversight and the next tidy-up would break `gate verify` on XENO-3.
