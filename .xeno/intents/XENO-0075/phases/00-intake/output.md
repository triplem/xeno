---
intent: github.com/triplem/xeno#75
phase: 00-intake
created: "2026-09-26T14:02:04Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+850b546
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: fbe079d02d2f7bcf60dc61b7fe1ce2e06311458618b0bd9588d1da050ffe5e6d
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

G-Schema counted a field as present when its value was neither absent, nil nor the empty
string, so `by-hand` passed as a hash and so would a typo or a truncated digest. No hash
field was shape checked anywhere, while `schema_version` was, twenty lines above, with a
regexp: the gate validated the shape of a version and not the shape of a hash.

The placeholder is legitimate in the trail, which is why a plain hex check would have
been wrong. Eighteen sealed intakes name `intake@0.1.0` and carry `by-hand` in
`strings_hash`, and two of them, M0's first, carry it in every hash field.

<!-- xeno:section:scope -->
## Scope

Appendix B says what a hash field may carry and when the placeholder is honest, and
G-Schema checks it. Three honest cases: a field with no writer, an artifact that
declares `tool: manual`, and a `strings_hash` whose bundle version the repository no
longer carries.

Recomputing a hash against the file it covers is #73 and not here. The order is
deliberate: the shape is cheap and should not wait for the comparison.

<!-- xeno:section:context-rationale -->
## Why this context

**The third case came out of the trail, not out of the rule.** The first draft accepted
the placeholder where an artifact declared `tool: manual`, which kept M0's two intents
green and turned sixteen others red. Reading every artifact's `template` and
`strings_hash` together showed why: the value is `by-hand` in exactly the eighteen that
name `intake@0.1.0`, a version the plugin no longer carries, and it is the real hash in
every artifact naming `intake@1.0.0`. A bundle that is gone cannot be hashed by anybody,
and nothing archives templates.

**What the check reads is the loader, not a list.** Whether the named version is the one
here is a question the template loader already answers, and asking it there means the
rule cannot drift from the resolution `section set` uses.

**The fixture was part of the finding.** The runner tests wrote `secrets_hash: s` and
`context_hash: c`, values the specification does not allow, and every phase they built
passed. A fixture that would be a finding is not a phase, and the tests now say so.