---
intent: github.com/triplem/xeno#258
phase: 01-requirements
created: "2026-10-06T17:29:58Z"
schema_version: "1.0"
runner_version: dev+d3983d3
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f2c8e59caba230883270e549d931532a97da4bf92f0df2b415b8f12976045339
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Thirteen numbered criteria. `Reason` on `model.Option` with `QuestionAsked` requiring it on the
recommended option; both templates at 1.1.0 with nothing else changed; G-Schema reading the numbered
list on a P1 at 1.1.0 and G-Test reading the mapping on a P4 at 1.1.0, both keyed on the declared
version with a numeric compare. Both rows of `docs/clause-readers.md` name a reader. Not #235, not
the gate reading the reason, not a stricter convention than section 5 wrote.
