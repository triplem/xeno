---
intent: github.com/triplem/xeno#108
phase: 00-intake
created: "2026-09-28T16:39:11Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+023193e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: c1c807db950e31204e84ad4e013d961a9b288477eabb62d7f6242e68cb6ea382
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

`SectionSet` writes `language`, `template` and `strings_hash` into the frontmatter and
leaves the remaining session fields out. A35 justifies that for `model`, `tool` and
`tool_version`, which come from the harness, and for `secrets_hash` and `rules_hash`,
which have no writer: a plausible value in a field nobody produced is worse than an
absent one.

`context_hash` was never one of those five. It has a writer, and the writer ran moments
earlier in the same directory: `phase start` writes `context.lock.yaml`, and
`recomputed` hashes exactly that file to check the field. The value was available where
the frontmatter was written and was left out anyway.

So a phase produced the way section 5 supports it — `init`, `phase start`, three
`section set`, `phase finish` — ended red on a field the runner knew:

    F-fc0280 …/00-intake/output.md: required field missing: context_hash
      next: add context_hash to the frontmatter

The only way to a green phase was to write the field by hand, which is the path section
5 calls unsupported. A51 says both recomputable hashes are checked in G-Schema rather
than one of them in G-Freshness; since only one of the two was ever written, the check
on `context_hash` had never run against a value a writer produced. It had only ever
reported the field absent.

<!-- xeno:section:scope -->
## Scope

`SectionSet` hashes the lock beside the artifact and writes `context_hash`. Where the
lock is absent the field is left out, which is a file written outside a started phase;
the lock's own absence is G-Freshness's finding and not this one's.

The comment on `SectionSet` says why the field cannot drift inside a phase, since a hash
written on every render otherwise reads as something that might.

Two assumption rows are amended rather than replaced, because both sentences were right
and only unfinished: A35's list of what is left out never included this field, and A51's
half about `context_hash` now runs against a produced value.

Not the other five fields. `model`, `tool` and `tool_version` still come from a harness
the runner does not have, and `secrets_hash` and `rules_hash` still have no writer. A35
is unchanged about them, and a phase produced by the supported path is still red on them
— it is simply no longer red on a field the runner could fill.

<!-- xeno:section:context-rationale -->
## Why this context

**Written on every render, not only on the first one.** The frontmatter of an existing
file is carried over whole, so a field written once would be whatever the first writer
left there, and a file whose frontmatter a hand had touched would keep the hand's value.
Writing it each time makes the rendered file agree with the lock it was rendered beside,
which is the only thing the field claims.

**Which lock the field covers when a phase is re-run.** The lock is written once at
`phase start` and nothing refreshes it, by design, so the hash written at the first
`section set` and the hash written at a later one are the same value. That is why the
field is stable within a phase and why re-rendering cannot move it. It follows from the
lock not moving rather than from anything the section writer does, which is exactly why
the comment says so: a hash recomputed on every write looks like a hash that could
drift, and the reader has no way to see that the file behind it is frozen.

**A35's boundary, read again rather than widened.** The five fields A35 leaves out have
no writer in this tree. This one had a writer all along, and the row's reason listed the
five without ever claiming the sixth. Amending the row records that the consequence was
drawn late; replacing it would delete a sentence that was right.

**What the fix buys beyond quiet.** A51 put both recomputable hashes in G-Schema. With
only `strings_hash` ever written, the `context_hash` arm of that check had never
compared anything — every observation of it was the missing-field finding from one layer
up. A phase now carries a value that arm can disagree with, so the half of A51 that
matters is exercised by the supported path rather than only by fixtures that write the
field by hand.
