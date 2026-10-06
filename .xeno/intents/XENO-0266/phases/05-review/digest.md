---
intent: github.com/triplem/xeno#267
phase: 05-review
created: "2026-10-06T20:12:31Z"
schema_version: "1.0"
runner_version: dev+3f1fab3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3f824f5c03369fcccd5f297706b775cedacaa16c84df93300f44e98174c07de4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The first standing rule's exception is the third in two days and is recorded in the commit
message, in P0 and in the checklist. The specification commit is its own and comes first. Two
deviations, both in P3: a paragraph in `docs/clause-readers.md` rather than three rows, because
none of the three clauses can be violated by code, and a third comment corrected that asserted
nothing wrong but led a reader to the wrong cause. Nothing to adopt — `gate verify` at 475
verdicts before and after, no test file in the diff — and one thing to know, that a declared
context budget is judged from P1 on and always has been. The residual risk is that the three
paragraphs have no reader; the control is that one of them states a fact a change must delete
rather than a permission a change can satisfy, which the learning recorded here proposes as a
convention.
