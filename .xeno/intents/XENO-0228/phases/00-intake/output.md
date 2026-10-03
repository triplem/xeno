---
intent: github.com/triplem/xeno#181
phase: 00-intake
created: "2026-10-03T12:37:33Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6cbeac4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b80e09fa0248b51d78079575cca8ef77914d046bf79b23ae78c021df2acbee25
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

`tool_version` is required by G-Schema and written by no command, so every phase of every intent in
this repository has gone red on it once and been corrected by hand. Twice per phase: once in
`output.md`, once in `digest.md`. Thirty intents, six phases each, is the order of three hundred
hand edits to files that sit inside `artifacts_hash`.

**The channel the specification designed for it was never built.** A35 assigns the field to the
harness. Section 7 says how a harness fact reaches the runner — a thin entry point that normalises
the environment, where "the runner only ever sees `XENO_*`" — and lists three variables:
`XENO_PLUGIN_ROOT`, `XENO_PLUGIN_DATA` and `XENO_HARNESS`. The runner reads none of them.
`grep os.Getenv` over `internal/` and `cmd/` returns one hit and it is the enforcement token. So the
one harness fact a gate requires has nowhere to arrive from, and A35's choice of absent over guessed
is correct while leaving nothing able to produce a value.

**The cost is paid twice because the two writers differ.** `SectionSet` carries the frontmatter of
an existing `output.md` over, so an edit there survives later section writes. `writeDigest` builds
its frontmatter from scratch every time, so the edit to `digest.md` is destroyed by every
`phase finish` and has to be made again — and again whenever a phase is judged a second time.

**A hand edit to a hashed file is the thing this process exists to prevent.** The field sits inside
`artifacts_hash`, so the correction is sealed with the phase and `gate verify` would report a later
fix as a divergence. XENO-0227 closed exactly this argument for three fields of `intent.yaml` and
recorded this one as its implementation learning in the same breath, which is the second time it has
been written down and the first time with a count attached.

**A35's row is also one field out of date.** It reads "two fields remain writerless, `tool_version`
and `rules_hash`", and `SectionSet` has written `rules_hash` since the shipped rule set gave it a
set to hash. `tool_version` is the last one.

<!-- xeno:section:scope -->
## Scope

**In scope.** `--tool-version` on the runner's surface, written into `output.md` by `section set`.
`phase finish` copying the value out of the phase's `output.md` into the digest rather than taking
it as a second input. Absence unchanged: no report means the field is left out, exactly as now. The
six phase skills telling the agent to report it. Tests for each, and A35 amended rather than
contradicted.

**Out of scope, and each for its own reason.**

A fourth `XENO_*` variable. `XENO_HARNESS_VERSION` beside `XENO_HARNESS` is the architecturally
correct home, and it is a change to the list section 7 enumerates. The first standing rule makes
that a person's commit, made before the code that follows from it. The flag does not block it and
would become its override; this intent does not write it.

Building the entry point. `XENO_PLUGIN_ROOT`, `XENO_PLUGIN_DATA` and `XENO_HARNESS` are specified
and unread, which is a WP11 gap larger than this field and wants its own issue. What is fixed here
is the field a gate already requires, not the mechanism it was supposed to arrive by.

Deriving the version in the runner. Asking the harness for its version is the most agent specific
question there is, and section 7 forbids the branching in as many words: the moment the runner
behaves differently per harness, the tools stop being interchangeable.

A `tool_version` in `project.yaml`. It would be an Appendix A addition, so a specification change
by the rule above, and wrong on its own terms: a harness version is a fact about a session and a
project file would state it once and then be stale.

Backfilling the artifacts already written. They carry the hand-edited value, which is correct, and
they sit inside hashes that are sealed.

<!-- xeno:section:context-rationale -->
## Why this context

Section 5 is read for the session field group — `language`, `secrets_hash`, `context_hash`, `model`,
`tool`, `tool_version` — because the field already exists there and this intent gives it a writer
rather than a definition. Nothing in section 5 changes.

Section 7's hooks subsection is read as the authority on where a harness fact comes from: the
normalising entry point, the three `XENO_*` variables, and the sentence that `XENO_HARNESS` is
recorded and never branched on. It is read twice — once for the channel it designs, which is what
makes a fourth variable the right long term answer, and once for the prohibition, which is what
rules out the runner finding the version out for itself.

A35 is read in full, including its two amendments, because this intent is its third and the row has
been misread as a list of members twice before. The reason it gives — a plausible value in a field
nobody produced is worse than an absent one — is the test this design has to pass.

`internal/runner/runner.go` is read for the two writers and the difference between them:
`SectionSet`, which carries over the frontmatter of an existing artifact, and `writeDigest`, which
builds its own and is rerun by every `phase finish`. That difference is the whole reason the cost
was paid twice, and it is what decides that the digest copies rather than asks.

`internal/gates/gates.go` is read for `sessionFields`, which is where the requirement lives, to
confirm that nothing about the gate needs to change.

XENO-0227's `03-implementation/learning.yaml` is read as the record that proposed this, and its
`04-verification/output.md` for the gap entry that counted the edits.

Nothing outside the repository is needed. The one question that would have changed the shape of this
— flag or environment variable — was put to the maintainer and answered before any code was written,
because one of the two answers required a commit the agent is not allowed to make.
