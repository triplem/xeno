---
intent: github.com/triplem/xeno#72
phase: 00-intake
created: "2026-09-26T13:44:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ba0c9c7
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: a4225286b002212b6d9e032ffdfd99509d6d1be27203b2854b0fc554e29d7778
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

Appendix B defines two values to the byte and says why: whoever verifies the trail
recomputes them. Section 5 names four more hashes with a phrase and nothing further, so
for four of the six an artifact carries there was nothing to recompute against.
`strings_hash` had a definition in `internal/template/template.go` rather than in the
document a verifier reads, and `context_hash` had none at all: every value in the
repository was produced by hand with `sha256sum`.

<!-- xeno:section:scope -->
## Scope

Appendix B defines `context_hash` and `strings_hash`, as the sha256 of one file's
normalised content, and states that `secrets_hash` and `rules_hash` are defined by the
package that first writes them. A26 closes with it.

No code. That nothing recomputes any of these values is #73, and it is a separate change
because the definition is what makes the check possible rather than the other way round.

<!-- xeno:section:context-rationale -->
## Why this context

**Two of four, not four of four.** `secrets_hash` and `rules_hash` each cover an
effective set rather than one file, and how a set of several files reduces to one value
is a decision that belongs with the code that assembles the set. Defining them now would
be deciding it in the abstract, and the appendix saying they are deferred is the honest
form of the same sentence.

**The definition was chosen to match what is already sealed.** 29 of the 31 artifacts in
the repository match the new definition byte for byte, checked by recomputing every one
of them before the words were written. The two that do not are M0's first intents, which
carry `by-hand` because nothing computed the value when they were written. So nothing is
re-sealed, and the two exceptions are visible rather than quietly wrong.

**`.gitattributes` is why the definition can be the normalised one.** A15 fixes LF for
every text file, so `sha256sum` over a working tree here and the normalised computation
agree. On a checkout that did not have that they would not, which is the one way a value
already written could stop matching.