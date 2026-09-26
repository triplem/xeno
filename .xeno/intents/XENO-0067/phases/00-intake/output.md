---
intent: github.com/triplem/xeno#67
phase: 00-intake
created: "2026-09-26T21:54:49Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0ea4806
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: a6a642317f28a9bfb4ae0d42eb0629e23ed5f653d4ff1c58330252ede0e580cc
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

G-Complete ran in one of its two modes. The abandoned one, invoked from `xeno intent
close`, checks that the intent carries a reason and that the closing learning record
exists. The mode a merging intent meets as part of P5 was written `not-implemented` in
every verdict, which was honest and is what A32 recorded.

The issue that carried the row said the mode needed a release record and an obligation
ledger. Section 7 says nothing of the kind: it checks that the artifacts of all
preceding phases are present and green, approved or overridden, and nothing about the
merge, which has not happened when the gate runs. The dependency was inferred from the
code's shape rather than read.

<!-- xeno:section:scope -->
## Scope

The review mode, reading each preceding phase's committed verdict and requiring present
and decided. Three states and three repairs: a phase with no verdict was never judged, a
provisional one waits for evidence, a red one needs its findings decided.

Not recomputation. `gate verify` recomputes every verdict already, so a gate doing it
again would report one divergence twice under a second name.

<!-- xeno:section:context-rationale -->
## Why this context

**Reading the sentence was the work.** The gate is thirty lines. What kept it unbuilt
was a belief about its dependencies that the specification does not support, recorded in
an issue and then in the register, where it looked like a scheduling fact. This is the
second issue this week written from the code and the register without the sentence that
governs the field, after #68 and `merge_method`.

**The states are kept apart because the repairs differ.** Collapsing them into "not
green" would be shorter and would tell somebody to decide findings in a phase that was
never run. A phase with no verdict, one that is provisional and one that is red are
three different pieces of work, and a gate that names which is worth more than one that
counts.

**It reads rather than recomputes, and that is a boundary rather than an optimisation.**
A gate that recomputed its predecessors would judge what another gate already judged,
and `gate verify` covers every phase including the ones this gate looks at. Where they
disagree the divergence belongs to verify, which is where a reader already looks for it.

**Nothing in this repository exercises it.** No intent here reaches P5, because M0 runs
one intake per package deliberately, so the mode is proven by a test that carries a
fixture intent through all six phases. That test also does what no command can: it edits
a sealed verdict to red, since `phase start` refuses to begin a phase on a red
predecessor and the gate exists for the case where somebody went around it.