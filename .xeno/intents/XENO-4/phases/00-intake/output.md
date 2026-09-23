---
intent: github.com/triplem/xeno#4
phase: 00-intake
created: 2026-09-23T18:30:41Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: e3c7ee36f170286e84828da196007fb05a55ff6cf8f81954e666d730086069d1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #4, two changes that arrived together and are kept apart here because they answer
to different parts of the plan.

## Scope

The commit convention, third and last placement: the reference leaves the subject for
the footer, which is possible because the premise underneath the earlier placements was
a host default and not a fact. Both hosts can build a squashed message from the merge
request description, so the setting that makes it so is part of the deliverable rather
than an assumption about how somebody configured their repository.

And the release tooling: semantic-release in place of four shell scripts, configured by
a file in the repository so that the same configuration serves both hosts, brought in as
a CI job rather than as a dependency of this project.

## Non goals

The rule that would check the subject form. It is `type: commit-message` and needs the
rule engine, which is WP4. Until then the convention lives in two documents and nothing
enforces it, which is the honest state of it.

## Note on this record

It was written after the work rather than before it. The intake is the first phase and
this one ran last, which is a departure from the process this repository exists to
demonstrate, recorded here rather than quietly corrected.
