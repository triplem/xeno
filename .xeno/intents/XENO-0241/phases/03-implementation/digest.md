---
intent: github.com/triplem/xeno#206
phase: 03-implementation
created: "2026-10-03T20:52:41Z"
schema_version: "1.0"
runner_version: dev+8574810.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3c6954234d6ace9479022d9b50217c960ef164addbcec49df8350e499aa52e42
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
One function in `internal/git` reads the paths of a range, one reader in `internal/runner`
turns them into intent keys and reports the ones that are neither complete nor abandoned,
one command exits 1 on any, and four files call it: this repository's verify job and both
shipped wrappers. Twelve tests, and the whole gate suite green over 324 verdicts. The state
is read from `summarise` and defined once. What this phase could not do is exercise the step
against its own intent, because the comparison is between committed trees.
