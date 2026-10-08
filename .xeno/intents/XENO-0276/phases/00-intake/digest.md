---
intent: github.com/triplem/xeno#325
phase: 00-intake
created: "2026-10-08T13:50:47Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 70a042ee1a779054b96b747fbd617fcb1d2f2a3a6133c8b50751a8fe301ae2f2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake of #325: the review skill lists the commands P5 is carried out with and omits the
two that write the checklist it is judged on.

Confirmed in the tree rather than taken from the issue: the command block names `phase start`,
three `section set` calls, `phase finish`, `gate approve` and `gate override`, and `grep` for
`review answer` or `review lens` over the file returns nothing. The skill describes the
entries shape correctly, which makes the omission harder to notice: the reader is told what an
entry is and not how to write one.

The consequence is a phase that passes having answered no rule — G-Policy counts entries, and
a checklist written as prose has none, which is A73 shape.

Noticed by running this intent own P5 through the two commands and finding the skill the wrong
place to look them up.

One section of the five.
