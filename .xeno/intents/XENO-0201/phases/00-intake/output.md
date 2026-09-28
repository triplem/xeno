---
intent: github.com/triplem/xeno#120
phase: 00-intake
created: "2026-09-28T20:29:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6e72fed.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 27f9bd71ff159236dbda07d988c3f750316110ceae0d6a82533228b1d6821738
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

Section 5 says who writes a digest, and it is not the agent:

> Three groups, not two, because `digest.md` is neither rendered nor written by the
agent. The > agent supplies the summary text, the runner filters, hashes and writes the
file, so a digest > carries no template and no strings bundle. What it does carry is
where it came from: model, > tool, language and the filter it passed through.

The working sequence of section 6 lists `digest.md` among what `xeno phase finish`
produces. Nothing in the runner writes it. `phase finish` judges the file and reports it
missing, which is the only part of the arrangement that exists, so every digest in this
repository was written by hand, frontmatter included.

The same hand writes three of the five fields A35 leaves out of `output.md`. A35's
reason is that `model`, `tool` and `tool_version` come from the harness, and for
`tool_version` that is true. For the other two it is not: `.xeno/config/project.yaml`
carries them, in the block section 12 defines.

    agent:
      tool: claude-code
      model:
        default: claude-opus-5

The runner reads that file already, for the artifact language, into a struct that models
the language block and not the agent one. So two fields with a source in the repository
are left absent on the grounds that they come from somewhere the runner cannot see.

This is what #120 asks about, from the other end. An agent cannot be asked to stop
working around the tool while the tool cannot produce what a phase requires: six intents
in the last two days filled the digest and those fields by hand, not as a shortcut but
because no command does it. The scoping comment on that issue named this the gap that
has to close before any enforcement is worth discussing.

`secrets_hash` is a different case and stays one. The effective filter is a shipped
pattern file plus a project's additions, and no such file exists anywhere in the tree,
so there is nothing to hash and nothing to filter against.

<!-- xeno:section:scope -->
## Scope

`phase finish` takes a summary and writes `digest.md` from it: the frontmatter section 5
requires of a file produced in a session, then the text. Without a summary it behaves
exactly as it does today, judging whatever digest is there.

The runner reads the `agent` block of `project.yaml`, and `model` and `tool` are written
into both artifacts it produces, `output.md` through `section set` and `digest.md`
through `phase finish`.

A35 is amended to name the three fields that still have no writer instead of five.

Not `secrets_hash`, and not the filtering. Section 5 says the runner filters, and the
file it would filter against does not exist. A digest written now passes through no
filter and must not claim otherwise, so the field stays absent rather than being written
as a hash of nothing.

Not `tool_version`. It is a property of the session rather than of the project, and the
config block that carries `tool` does not carry it. A35 keeps it.

Not `rules_hash`. WP4.

Not a required summary. A phase finished without one behaves as before. Section 5 says
the tool path is supported and not enforced, so a command that refused a phase for not
using it would contradict the sentence this intent is implementing.

Not the enforcement of #120. This closes the gap that made enforcement unarguable;
whether to enforce anything is still that issue's question and still needs a
specification change.

<!-- xeno:section:context-rationale -->
## Why this context

**This is the code catching up with section 5, not a new feature.** The sentence about
who writes a digest is normative and has been since before the runner existed; the
working sequence names `phase finish` as the producer. So no document changes, and the
first standing rule is satisfied by reading rather than by editing.

**The summary is optional because section 5 says the tool path is not enforced.** A
required summary would refuse a phase for not using the operation, which is the opposite
of what the same section says two pages earlier: the supported path is not an enforced
one, and G-Schema is the backstop. Optional also means nothing existing breaks, and a
digest written by hand stays a digest written by hand, judged as it is today.

**`model` and `tool` come from the project file because that is where section 12 puts
them, and the runner is already reading that file.** The gap was never a missing source,
it was a struct that modelled one block of the file. A35's reason held for three fields
and was applied to five, which is the same error the `context_hash` intent found in the
same row: a list of reasons read as a list of members.

**`secrets_hash` stays absent, and that is the load bearing restraint of this intent.**
Section 5 says the runner filters and hashes; a digest written today is filtered by
nothing. Writing the field would assert a filter that does not exist, and writing the
digest without it is honest: the file is produced by the runner, and the one thing it
cannot yet claim is the one thing it does not do. G-Schema reports the field missing,
which is the true state.

**Why the digest is written at finish rather than by a command of its own.** The working
sequence puts it there, and the reason survives inspection: a digest summarises a phase,
so the earliest honest moment is when the phase is complete, and that is also when the
runner already computes the hash the file is sealed with. A separate command would let a
digest be written before the phase it summarises had finished.

**What this does not settle about #120.** Nothing here forces anybody to use the tool,
and section 5 forbids that shape anyway. What changes is that the hand written digest
becomes a choice rather than the only route, which is the precondition the scoping named
and not the enforcement itself.
