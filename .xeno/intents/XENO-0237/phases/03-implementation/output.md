---
intent: github.com/triplem/xeno#202
phase: 03-implementation
created: "2026-10-03T19:25:49Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2da3489c9fe667e82ccbe7724780bc10d2f3201245e9d99ecdc175fb61e60c6f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`CLAUSE-READERS.md`, new. One pass over both normative documents: 219 candidate
sentences extracted by searching for `never`, `always`, `must`, `cannot`, `has to`,
`is required`, `refuses`, `fails red` and `is recorded` outside tables and code blocks,
each read and placed in one of four kinds.

Thirty-four tool requirements are tabled with the gate, refusal or workflow step that
would fail if the clause were violated. Eleven architectural properties are tabled with
what would notice, which for nine of them is nothing.

Forty-eight clauses addressed to a person and 126 explanatory sentences are counted and
not listed. The counts are stated so the pass is falsifiable.

Every symbol named as a reader was looked up rather than recalled, and all of them
exist: `model.ResultRequiredKinds`, `hashing.PhaseExcluded`, `hashing.FindingID`,
`predecessorAllowsStart`, `rewriteStatus`, `notImplemented`, `IntentClose`, `Decided`.
`exec.Command` appears outside tests in `internal/git` and `internal/external` only.

Three findings have no issue yet and are written down as such: `XENO_PLUGIN_DATA` is
specified and read by nothing, G-Complete runs only at P5 so an intent that stops short
is never checked for completeness, and section 12's model-tool-version triple is
self-reported with nothing checking it is true.

`ASSUMPTIONS.md` gains A90 for the choice of a survey over a gate.

<!-- xeno:section:deviations -->
## Deviations from the design

The design said the file would sit beside `ASSUMPTIONS.md` and it does, at the
repository root.

One thing the design did not anticipate. The pair of clauses under section 11, "what is
sealed is never rewritten", splits into modification and deletion, and the two have
different readers: `gate verify` has reported a divergence all along and the deletion
guard arrived with #193. They are listed as two rows rather than one, because a single
row would have to name one reader and would then be half wrong — which is the error this
file exists to find.

Nothing else deviates. No code, no field, no gate, no rule, so no specification change
precedes this.
