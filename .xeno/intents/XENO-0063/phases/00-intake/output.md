---
intent: github.com/triplem/xeno#63
phase: 00-intake
created: "2026-09-26T21:40:56Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0e9f4b9
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 70af949aea675f37d3e48cc145851e5d01a140ae981cd69a97a09e49f38e188e
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

G-Freshness implemented the first half of its rule, the context hash against the
predecessor. The second half, that no file a preceding phase read has changed since it
read it, needed the information base, and `context.lock.yaml` carried only the
predecessor hash and the evidence source.

A6 accepted that as temporary and drew the consequence rather than hiding it: no phase
beyond P0 runs in this repository, because until now a pass on this gate said more than
was checked. That hold was this gap.

<!-- xeno:section:scope -->
## Scope

`files` in the lock, resolved from the context profile when a phase starts and never
refreshed, and the comparison in G-Freshness against the tree for every preceding phase.
A changed file is a finding naming it; a file that is gone is a finding of its own,
because there is nothing left to read again.

Not the rest of WP8. The budget is declared and not yet reported against, the symbol
index is WP15, and reading outside the profile is not recorded, which needs a harness
that says what it read.

<!-- xeno:section:context-rationale -->
## Why this context

**The profile is P0's and applies to the whole intent.** Section 4 puts
`context-profile.yaml` in P0 alone and section 12 says a phase reads what it names, so
one budget per intent is the reading that leaves both sentences true. Six profiles would
need a rule for which one a phase obeys, and the documents give none. A53 records it.

**The pattern syntax is not defined anywhere, so it is recorded rather than assumed.**
Section 12 writes `src/payment/**` and `**/testdata/**` and defines neither. The matcher
is what those two mean to a reader: `**` spans any number of segments including none,
and a segment otherwise follows `filepath.Match`. Fifteen lines in `internal/model`,
because one dependency is a decision and a second is a bigger one.

**The comparison reads the tree rather than the commit range.** The gate's sentence says
"changed by the change under review", and the range is an input no gate reads until WP4
brings commit predicates; it is also gone after a squash, while the tree is what a
verifier has. Reading the tree reports a file changed outside the range as well, which
is more than the sentence asks for and never less. A55.

**Nothing in this repository changes, and that is the honest outcome.** No intent here
has a profile, so every lock records nothing and the comparison has nothing to say. The
check is not weaker for it: it has been told nothing, which is a smaller claim than an
empty profile would make.