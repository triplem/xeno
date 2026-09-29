---
intent: github.com/triplem/xeno#132
phase: 03-implementation
created: "2026-09-29T16:18:17Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+1dc1483.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c033634023c6835a2809d3790c4827516c13503d6539e9208cca22afea8cc886
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: by-hand
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`internal/runner`. `Decided` is a named predicate: not `red` and not `provisional`.
`predecessorAllowsStart` uses it in place of its own comparison against `red`, so the
sequence and the listing cannot drift apart about what a settled phase is.

`IntentSummary.Status` becomes `State`, and `summarise` stops copying `intent.yaml`'s
field into it. `state` computes the three answers in order: `abandoned` from the field,
`complete` where P5 carries a decided verdict, otherwise the phase the work has reached,
and `no phases` where there are none.

`cmd/xeno`. The column prints `State` and is wider, because `03-implementation` is
seventeen characters.

Four tests: the three states together, a red P5 that must not read as complete, an
intent with no phases, and `Decided` against all five verdict values.

**What the listing says now.** The ten intents carried past the intake read `complete`;
the fifty that stopped read `00-intake`, which is where they stopped; and this intent
reads `03-implementation` beside `02-design green`, which is a phase running above its
last decided one.

The column that prompted the question now answers it.

<!-- xeno:section:deviations -->
## Deviations from the design

None in the design. The predicate, the three answers in order, the word `complete` and
the absence of any schema change are as P2 decided.

Two notes on the work rather than on the change.

The struct field is renamed rather than kept and shadowed. `Status` meant the stored
value and `State` means the computed one, and the per-phase listing already uses `State`
for a computed thing and `Status` for a verdict, so the names now agree across both.

This intent's phases carry no `cost.yaml`. The hook that XENO-0205 shipped is read when
a session starts, and this session began before it existed, so nothing has appended to
the ledger. That is the gap XENO-0205's verification named, observed rather than worked
around: invoking the command by hand to manufacture a record would have produced a
figure no hook attributed.
