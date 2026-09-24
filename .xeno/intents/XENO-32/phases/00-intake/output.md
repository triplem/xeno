---
intent: github.com/triplem/xeno#32
phase: 00-intake
created: 2026-09-24T21:09:05Z
schema_version: "1.0"
runner_version: 0.1.0-dev+c7abd5d
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: efd2d4eab12efbba5df4cb5d0de8b9c9b53f8b78100e133b8fa288a752f3549f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #32, step 3 of the sequence. WP2 and WP3 are one piece of work: the engine cannot
be tested without a set, and a set means nothing without the engine.

## Scope

`template.yaml` and the strings bundles, the renderer that writes anchors the agent
never touches, resolution where a project template beats the shipped one, and
`template_source` in `context.lock.yaml`. Then the six templates and both bundles.

## Non goals

`review-checklist` rendered from rules, which needs WP4. It renders present and empty,
which WP2 says is enough for M0.

The frontmatter fields a session produces. `model`, `tool` and `tool_version` come from
the harness, which is WP11.

## The one this package is really about

The section count is the proportionality lever, not the gates, and this is the package
that decides whether the process feels bearable. Three to four required sections per
template is the budget; a fifth means arguing another away.

There is a measurement for it already, from this repository rather than from a guess:
XENO-9 was ten lines of shell carried through an issue, an intent, an intake of four
files, a branch and a pull request. Whatever this set costs is on top of that.
