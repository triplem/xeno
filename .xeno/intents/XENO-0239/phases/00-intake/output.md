---
intent: github.com/triplem/xeno#213
phase: 00-intake
created: "2026-10-03T20:03:34Z"
schema_version: "1.0"
runner_version: dev+427eeb7
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5c66c94437899c52f97ee6140237bd0aec239fe32c8cb4302aff820123fe51c6
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

Two normative documents said opposite things about the same file.

WP8's completion condition read "every read outside the profile appears in
`context.lock.yaml`". Section 5 of the process definition forbids it in a paragraph with
its own heading: the lock "records the context that was declared, not everything that was
read", and "treating it as a measurement of what was read would be wrong".

So WP8 could not be finished without breaking section 5, and `ContextLock` in
`internal/model/model.go` matches section 5 exactly — the code was right and the plan was
unsatisfiable. Neither document is wrong by default, which is why this was a decision and
not a patch.

Both reasons are load-bearing. Section 5's lock is written before the agent starts and
`context_hash` seals it, "and a lock rewritten at the end would describe nothing"; a
record of what was read can only be written afterwards, which would make one file a
measurement taken at two different times. WP8's intent is equally sound: G-Schema judges
the budget against `files`, so a read outside the profile costs nothing today and is
invisible, which makes the budget one in name for anything an agent chooses to open.

It is also a miss by the clause audit of #202. That pass checked every clause against the
code and never checked clauses against each other, and `CLAUSE-READERS.md` has no column
for two documents disagreeing.

<!-- xeno:section:scope -->
## Scope

In scope is WP8's paragraph and its done-when: the recording moves from
`context.lock.yaml` to the phase's digest, which is written when the phase closes, covered
by `artifacts_hash` and read at P5. The plan says plainly that it is the agent's own
account, because nothing in the harness reports what was opened. One row in
`ASSUMPTIONS.md` for the resolution and the two alternatives.

The specification change is its own commit, made before this intent's work, as the first
standing rule requires. The maintainer chose the resolution from three put to them.

Out of scope is section 5, which is untouched and wins — that is the rule, and here it
also happens to carry the better reasoning.

Out of scope is a gate. Nothing judges the digest's prose, and a check that an
out-of-profile read was declared would need to know what was read, which is the thing
nothing reports.

Out of scope is WP8's first half, "a repeated phase reads only what changed". It is still
unmet and belongs to its own issue once this is settled.

<!-- xeno:section:context-rationale -->
## Why this context

The input is WP8, section 5's paragraph on what the lock records, the lock's shape in
section 5 and `ContextLock` in `internal/model/model.go`, which were compared field by
field: `files`, `rules_applied`, `template_source`, `plugin`, `tools`, and nothing for a
read outside the profile in either.

The digest's properties are read rather than assumed, because the resolution rests on
them: written at `phase finish` from the agent's summary, inside `artifacts_hash` and so
sealed once written, outside `context_hash` and so not implicated in the lock's seal, and
rendered into the P5 checklist, which is what makes a person its reader.
