---
intent: github.com/triplem/xeno#96
phase: 00-intake
created: "2026-09-27T09:31:15Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+8dbd563.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 10a71c8f873a07d8a35056c5d96afc282395a2edf03069d49f80ba314434e27a
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

The four documents sat under `.xeno/docs/`. `.xeno/` is what the tool keeps in a
repository: intents, configuration, the vendored plugin, the gitignored local state. So
the same path meant two different things depending on which repository a reader was in,
and an adopter following `M0.md` would have looked for the process definition inside
their own `.xeno/`.

It also collided with where WP16 has to put things. The documentation build and its
Pages deployment want a root directory, and `docs/` for the site beside `.xeno/docs/`
for the normative pair guarantees that somebody edits the wrong one.

<!-- xeno:section:scope -->
## Scope

`git mv .xeno/docs docs`, the fourteen live references that point at the old path, and
each document's own `location` field. `docs/` at the root, as the issue asks, rather
than `spec/` or a subdirectory: it is where WP16 generates from, so the documents and
the site are one tree.

`M0.md` and `ASSUMPTIONS.md` stay at the root. They are this project's record rather
than the tool's, and moving `ASSUMPTIONS.md` would cost more references than it saves
confusion.

<!-- xeno:section:context-rationale -->
## Why this context

**The blast radius was measured before anything moved.** Sixteen references in tracked
text, two of them inside sealed artifacts, none in code, and no hash affected: the
documents lie outside every phase and intent directory, so `artifacts_hash` never
covered them. 46 verdicts before and 46 after, which is the check that this was a text
move and not a change to the trail.

**Two references stay behind on purpose.** The learning records of XENO-0035 and XENO-5
name `.xeno/docs/implementation-plan.md` as their target. Section 11's rule is that what
is sealed is never rewritten, and those are historical statements about where a document
was when somebody wrote the learning. Repairing them would be the convenient phrasing
the same section warns about.

**The normative documents are edited here, which the standing rule reserves for a
person.** What changed in them is four `location` fields, three prose mentions of the
directory and one sentence about where they live. The instruction to move them is the
maintainer's, recorded in #96, and the commit says so rather than letting a reader
wonder why the agent touched a document.

**`git mv` rather than a copy and a delete.** History per file matters more here than
anywhere else in the repository: these four files are the project's decision record, and
following a sentence back to the commit that wrote it is how a disagreement gets
settled.