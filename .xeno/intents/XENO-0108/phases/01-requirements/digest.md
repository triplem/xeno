---
intent: github.com/triplem/xeno#108
phase: 01-requirements
created: "2026-09-28T16:57:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+023193e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 6ea202ada72d56477f14abe6de27214bda1c8532d30444966effc808e11eb69c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Writing the acceptance criteria first is what turned a one line change into five checks:
the absent lock, the second write and the untouched five fields are each a case the code
has to answer, and only the first was obvious from the issue. AC3 is the one that
mattered, because it is the question the issue said to settle rather than discover, and
stating it as a criterion is what made the comment on `SectionSet` necessary rather than
decorative.
