---
intent: github.com/triplem/xeno#109
phase: 01-requirements
created: "2026-09-29T18:36:12Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+3ec2429.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c80696e6aa4ffc00a40210bbcffaebe364bc2d4169b8ef10a51240d26e403f80
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**AC1.** `model.KnownIntentFiles` names exactly what section 4 lists at the intent
level: `intent.yaml`, `assumptions.yaml`, `learning.yaml`, `gate.yaml`.

**AC2.** A file in the intent directory that is not one of those is a finding from
G-Complete when an intent is closed, naming the file and what to do.

**AC3.** A directory in the intent directory other than `phases` is a finding of its
own, with its own wording, as the phase level check distinguishes them.

**AC4.** `phases/` and the four known files produce no finding, so an ordinary
abandonment is judged on its reason and its learning record alone.

**AC5.** The finding appears in the verdict `intent close` writes, which is the one
place an intent level `artifacts_hash` exists, and it turns that verdict red as any
undecided finding does.

**AC6.** The phase level check is unchanged, and so are the hash, its exclusions and its
normalisation.

**AC7.** An intent directory that cannot be read produces no panic and no false finding:
whatever `intent close` did before, it still does.

**AC8.** Everything green stays green. `gate verify` matches every verdict, and no
intent here is abandoned so none gains a finding.

<!-- xeno:section:non-goals -->
## Non goals

The hash. What it covers, what it excludes and how it normalises are Appendix B's, and
this makes the property it claims true rather than changing the value.

The phase level check, which exists and works.

A new gate. G-Complete has the mode this belongs in.

Prevention. Section 4's arrangement is that a verdict records what a directory
contained, so a stray file is reported and still enters the hash of the run reporting
it. Refusing to hash would be a different design and a specification change.

`cost.yaml` at the intent level, which section 4 puts in a phase.

The merged case. A merged intent computes no intent level hash, by section 8, so there
is nothing to guard there.

<!-- xeno:section:constraints -->
## Constraints

Nothing under `docs/` changes. Section 4 enumerates the files and Appendix B claims the
property; this makes the claim true at the second level.

The check must mirror the phase level one. Same question, same shape, same two wordings
for a file and a directory, so that a reader of either recognises the other and neither
drifts.

It goes in `CompleteOnClose`, which is where the intent level hash is computed and the
only occasion it exists.

Section 4's list is the authority for what is known, and it is copied rather than
extended.

No new gate, no new field, no change to the hash.

One intent, one issue. The commits reference #109.
