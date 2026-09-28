---
intent: github.com/triplem/xeno#108
phase: 03-implementation
created: "2026-09-28T16:59:08Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+023193e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 2b63f62ae16154a7729c82c293ca0655219828240b727999738bb72405281438
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: by-hand
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`internal/runner/runner.go`. `SectionSet` hashes the lock beside the artifact and writes
`context_hash` into the frontmatter it is about to render. The path is built from
`model.PhaseDir`, the same expression that builds the path to `output.md` two lines
above, and the hash comes from `hashing.FileHash`, which is what G-Schema recomputes
with. Where `FileHash` returns an error the field is not written and nothing is refused.

The doc comment gains a paragraph. It says the field is written on every render because
an existing frontmatter is carried over whole, and it says the value cannot drift inside
a phase because `phase start` writes the lock once and nothing refreshes it. Both
sentences are there because the behaviour looks like something that could move and is
not.

`internal/runner/runner_test.go`. Two tests and one fixture helper.
`TestSectionSetWritesTheContextHashOfTheLockBesideIt` reads the field back, compares it
against `lockHash`, writes a second section and compares again.
`TestTheSupportedWalkIsNotRedOnTheContextHash` starts a phase, writes its three required
sections, evaluates, and fails on any finding whose cause or next step names the field.
`frontField` reads one frontmatter field of a phase's `output.md` from disk.

`ASSUMPTIONS.md`. A35 records that `context_hash` was never one of the five and does
have a writer. A51 records that its `context_hash` arm now runs against a produced
value. Neither row loses a sentence.

Both tests were run against the tree without the change and both fail there, one on an
empty field and one on the missing-field finding itself.

<!-- xeno:section:deviations -->
## Deviations from the design

None in what was built. The writer, the placement, the absent-lock case and the two
amended rows are as P2 decided them.

One deviation in the order the record was assembled. The code and the two tests were
written and committed before P1 to P5 existed; the intake was written first, and the
rest of the phases were written afterwards against a change that was already in the
tree. So P2 documents a decision it did not gate, and P4 reports on tests that were
already passing when it was written.

That is a finding about how this intent was run and not about the change. Section 4's
note that the phase order is the order the record is assembled rather than the order the
work happens covers writing tests before the phase that proves them; it does not stretch
to writing the design after the implementation. The honest record is that the sequence
was not followed here, which is why it is written down rather than absorbed.

Nothing in the change depends on the order: each phase's content was reconstructed from
the code, the issue and the assumption rows, all three of which a reader has.
