---
intent: github.com/triplem/xeno#16
phase: 00-intake
created: 2026-09-24T19:02:03Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: e983b01655a4c7ce68e7934c70ba64150535687934a204359fe6579098c72c6e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #16. Every artifact in this repository says it was written by `0.1.0-dev`, which
is the name every unreleased build has ever had, so the field identifies nothing.

## Scope

A development build reports the commit it was built from and whether that tree was
clean, from the `vcs.revision` and `vcs.modified` settings the Go toolchain stamps into
every binary. A release build is untouched, because ldflags set the version at link time
and the stamping only applies where the placeholder survived.

## Non goals

`PluginVersion`. The stamp identifies the binary that wrote an artifact and there is no
plugin binary, so appending the runner's commit to it would put information into a field
that cannot carry it.

Rewriting what is already sealed. Thirteen phases carry the placeholder and will keep
carrying it; changing them would change every `artifacts_hash` and invalidate every
verdict.

## A consequence worth stating before it is noticed

`context.lock.yaml` is inside `artifacts_hash` and carries this field, so a phase
started by a development build will be sealed with that build's commit in it. That is
the point rather than a side effect: the verdict then records which binary judged it.
The lock is written once and never refreshed, so the hash does not move afterwards.

Two builds from the same commit with different uncommitted changes report the same
string. `.dirty` says that the tree was not the commit, not which tree it was.
