---
intent: github.com/triplem/xeno#6
phase: 00-intake
created: 2026-09-23T19:23:18Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: b7c654cd8b2756ef97bb796e56f3d0e66abc4b89ec088104f2f82ae8c1beae37
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #6, the half that can be done without the instance.

## Scope

Every action and every tool the pipeline fetches is pinned to something that cannot
move, and what it fetches is enumerated in the repository with its source and its pin.

The pins are taken from the run that produced v0.4.0 rather than from what is newest,
because what is wanted is the set that demonstrably works, not the set that happens to
be current. The versions were read from the workflow log and from the published bill of
materials, which records its own generator with its hashes.

## Non goals

The instance. The form each fetch takes on a runner without egress, the module proxy
question for a repository that vendors everything without a `go.sum`, and writing the
answer where the move will be carried out from all need a host that does not exist yet.
The issue stays open for them.

Upgrading anything. Three of the four actions are behind their current major, and the
Node 20 deprecation warning on two of them says so out loud. Pinning what works and
moving to a newer major are different changes and a run that does both cannot say which
of them broke it.
