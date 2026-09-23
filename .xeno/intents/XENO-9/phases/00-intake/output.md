---
intent: github.com/triplem/xeno#9
phase: 00-intake
created: 2026-09-23T19:36:21Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: ba254e36de69cbd446fb4505656b679d3287649e4830b9dc2328c3a56b3246a7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #9. The repository offers three merge methods and only one of them keeps the
convention.

## Scope

Squash becomes the only method the repository offers, applied by
`scripts/github-settings.sh` alongside the two settings it already carries, and printed
before and after like they are.

`CONTRIBUTING.md` gains the width at which a pull request description is written, since
GitHub rewraps it to 72 characters when it builds the squashed message. Measured on the
first pull request rather than assumed.

## Non goals

The other merge methods elsewhere. This is a setting of this repository, not a statement
about how a project using Xeno should merge.

## On proportion

Ten lines of shell and one paragraph, carried through an issue, an intent, an intake of
four files, a branch and a pull request. The implementation plan asks the dogfooding
stretch whether the process is too heavy for small changes, and this intent is one
answer to it. Skipping the ceremony because the change is small would have removed the
measurement along with the effort.
