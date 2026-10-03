---
intent: github.com/triplem/xeno#202
phase: 00-intake
created: "2026-10-03T19:23:52Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c7f6399e505aff7f89232c4eea39a470c6a7a70076d5d84f4fbb0c52c4fecaa9
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

Eight times in this project a rule was stated in one document and nothing in the
repository compared itself against it. Each was found by accident, late, and cost about
an intent to close: `plugin_version` was a constant for weeks while every artifact
reported it; the deletion half of "what is sealed is never rewritten" had no guard until
#193; `vendorPlugin` copied two of thirty-four files, so G-Supply would have failed for
every adopter.

The documents are normative and long. A clause in them is one of four things: something
the tool can enforce, something true only by how the code happens to be arranged,
something addressed to a person, or prose that merely uses a normative word. The first
two are the ones worth knowing about, and nothing distinguishes them.

What is missing is not a gate. It is the list: every normative clause against what reads
it, so that the unread ones are a known set rather than the next accident.

<!-- xeno:section:scope -->
## Scope

In scope is one pass over `docs/process-definition.md` and `docs/implementation-plan.md`,
both normative, producing a list in the repository: each tool requirement with the gate,
refusal or CI step that would fail if it were violated, and each architectural property
with nothing asserting it named as such.

Out of scope is fixing anything the pass finds. That is the point of #202: a pass that
fixed as it went would stop at the first finding and the other seven would stay unknown.
Findings become issues or a recorded decision afterwards.

Out of scope are `CLAUDE.md` and `CONTRIBUTING.md`, which carry conventions rather than
the process, and no gate reads them by design.

No field, gate, tool or rule is added, so this is not a specification change.

<!-- xeno:section:context-rationale -->
## Why this context

The process definition and the plan are the whole input, and both are already in the
context the phase seals. The readers are looked up in `internal/gates`, `internal/runner`,
`internal/rules`, `internal/hashing` and `.github/workflows`, which is where everything
that can fail lives.

The lookup is done against the code and not from memory, because a reader recalled is how
`plugin_version` stayed in the enforced set for weeks: the field was in G-Schema's
required list, so it looked read, and what filled it was a constant.
