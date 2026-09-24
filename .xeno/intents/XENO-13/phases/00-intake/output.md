---
intent: github.com/triplem/xeno#13
phase: 00-intake
created: 2026-09-24T05:14:25Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 4189acbc0eca5090e9b993a96210e3dba4be0a0e8b8c0ceaabd6a10e6b22bd7c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #13. Three comments explain a construction by contrasting it with one that was
replaced during the same session, and a reader has no way to know what the contrast
refers to.

## Scope

The three found: the step order in the release workflow, the placements in
`scripts/github-settings.sh`, and a past tense in `internal/gates/gates.go` about a
state that no longer exists. Each is rewritten to say what the construction is and why,
without needing what it replaced.

The convention is written into `CLAUDE.md`, briefly, so the next comment is written that
way rather than corrected afterwards.

## Non goals

Commit messages, which are exactly where the replaced thing belongs: a reader looking
for history is reading the log, and the log is the only place that carries it honestly.
Nothing is removed from one.

The comments checked and found not to be in this class. "bin rather than app" compares
two subcommands that both exist, "written outside dist and moved in" explains a
construction by the trap it avoids, and "@semantic-release/npm is deliberately absent"
explains a present absence. A rule that removed those would cost more than it saved.
