---
intent: github.com/triplem/xeno#108
phase: 02-design
created: "2026-09-28T16:58:24Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+023193e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: b544e2ccc83490b91cce9a1b9c0a90161c68b6c92dd72c6ff5999086bd9b7781
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The decision that took the work was where the missing lock is reported. Refusing the
section write is the loudest option and the wrong one: it makes the prose writer answer
for a file it does not own. Writing on every render rather than once fell out of one
fact about `SectionSet`, that it carries an existing frontmatter over whole, which also
disposed of the variant where `phase start` writes the field.
