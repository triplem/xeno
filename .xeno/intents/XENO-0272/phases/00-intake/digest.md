---
intent: github.com/triplem/xeno#277
phase: 00-intake
created: "2026-10-07T13:28:05Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5436dd20ed5cf7566de5405d0183f1d7f12c4232bf50d3802350373cb25af122
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`hashes` read the artifact's own `tool` field and accepted the `by-hand` placeholder in every
hash field when it said `manual`. The comment's reason was sound — nothing produced a manual
artifact, so none of its hashes had a writer — and it assumed the field is true. Section 12
says the triple is a declaration and that nothing corroborates it, so the one gate that acted
on the triple was relaxing itself on a declaration, which decided a verdict rather than how a
register reads.

Measured: four artifacts in the trail declare `tool: manual`, the intakes of XENO-1 and XENO-2
written by hand before M0. Dropping the term turns exactly those two phases red on four
`context_hash says by-hand where a writer exists` findings and moves nothing else, because
`WriterlessHash` already covers `secrets_hash` and `rules_hash` and `goneBundle` covers their
`strings_hash`. So the term had one effect: it exempted `context_hash` for two pre-M0 intakes.

Both phases do have a `context.lock.yaml`, so the finding is correct rather than spurious, and
section 11 forbids rewriting a sealed artifact to satisfy it. The findings can only be released.

The maintainer was put the issue's three shapes with the measurement and chose the narrower
exemption, then approved all four findings with a reason naming each artifact's date and why it
cannot be corrected.

In scope: the term, the stale doc comment, one inverted test case, the four approvals, and a
register row. Out of scope: the specification, correcting the artifacts, `model` and
`tool_version`, `goneBundle`, and the `tool` field itself.
