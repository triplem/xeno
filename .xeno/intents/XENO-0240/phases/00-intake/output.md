---
intent: github.com/triplem/xeno#215
phase: 00-intake
created: "2026-10-03T20:22:27Z"
schema_version: "1.0"
runner_version: dev+b626f1a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0a3bf2a9723afb8b14919cf905451ebc36cdab251de49b4cf897fd66c7a6bdc4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

`Runner.Start` checks the run marker and then writes `context.lock.yaml` unconditionally.
Starting a phase that already has a verdict therefore rewrites a file inside
`artifacts_hash`, which section 11 forbids: "what is sealed is never rewritten".

Nothing refused it. `gate verify` reported the divergence afterwards, which it did twice
during the past day's intents, and the recovery was `git checkout --` on a sealed
artifact. A guard that catches a forbidden act after it happened is the right last line
and the wrong first one.

The overwrite also destroys the record the project's own change-driven re-reading depends
on. That lock is the only statement of what the phase was given, and `ChangedSince`
compares against it.

**What this problem is not.** The issue it comes from claimed that nothing tells a repeated
phase what changed. That is wrong and the issue is corrected: `xeno phase start` has
printed the changed set since #171, derived by `ChangedSince` from the predecessor's lock
and the tree, with two tests over it. WP8's first half is built. The claim was filed
without looking, and a second derivation of the same comparison had been written before the
existing one was found — which would have been worse than the gap it was meant to fix,
since two derivations of one answer can disagree.

<!-- xeno:section:scope -->
## Scope

In scope is one refusal in `Runner.Start`: a phase that has a verdict is not started again,
and the refusal names the way to redo the work. Two tests, one that the refusal happens and
leaves the lock untouched, one that removing the verdict deliberately still allows a fresh
start.

Out of scope is any mechanism for reporting what changed. It exists, it is derived rather
than recorded for stated reasons, and the work begun against it was deleted rather than
finished.

Out of scope is the context profile. This repository has none, so every lock's `files` list
is empty and `ChangedSince` is silent throughout the trail. That is a finding of its own and
gets its own issue rather than being absorbed here.

No field, gate, tool or rule. The refusal enforces a clause the process definition already
carries, which is why no specification change precedes it.

<!-- xeno:section:context-rationale -->
## Why this context

The inputs are `Runner.Start`, section 11's clause on sealed artifacts, `hashing.PhaseExcluded`
for what `artifacts_hash` covers, and the three refusals `Start` already makes — the run
marker, a predecessor with no completed verdict, and missing declared evidence — because a
fourth has to read like them.

`ChangedSince` and its printer in `cmd/xeno/main.go` were read because the issue's first claim
was about them, and reading them is what corrected the claim. The lesson is in the record: the
search should have come before the issue.
