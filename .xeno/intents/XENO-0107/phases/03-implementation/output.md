---
intent: github.com/triplem/xeno#107
phase: 03-implementation
created: "2026-09-28T18:56:39Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+907c220.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 90a4bc70ccfbc9757ae25b047ea0805221519946f9ee830112de42d30e2a376a
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

`internal/gates/gates.go`. `Status` gains three lines at the top of its loop and four
paragraphs on its comment.

The condition is `ch.Result == "fail" && len(ch.Findings) == 0`, and it returns an empty
status with an error naming the gate. It stands before `pending` is read and before the
findings are walked, so a malformed check is refused whatever stands beside it and a
well formed one earlier in the list cannot decide the run first.

The comment carries what the code cannot: that red is a promise something can be decided
and there is no id to decide, that no gate of this runner reaches the state because
every one returns through `result()`, that the writers it is for are an external gate
and a hand edited `gate.yaml`, and that the rule is here rather than in `Invariants`
because `Invariants` runs from `evaluate` alone while every derivation of a status
passes through here.

`internal/gates/carryforward_test.go`. Five tests. The refusal and its wording; a
malformed check beside a red one, which is the ordering; the three results that carry no
finding legitimately, where `pass` is every passing gate in the tree; a `fail` with a
finding, still red; and the id collision refusal with its own words, so the new rule did
not displace it.

`internal/runner/runner_test.go`. One test through a caller. It runs a real phase red,
appends a malformed external check to the stored `gate.yaml` as a person with an editor
would, and approves a finding. `Decide` refuses, which is AC2 and AC7 in one path and
the reason the rule is not in `Invariants`.

Three of the nine new tests fail against the tree without the change. The other six are
guards on what must not move, and they pass either way by construction; the verification
phase says which are which rather than counting all nine.

<!-- xeno:section:deviations -->
## Deviations from the design

None in the change. The condition, the location, the position in the loop and the
refused alternatives are as P2 decided them.

Two things the design did not foresee, both inside decisions it had already taken.

The error line ran past eighty-eight characters and is split across two, the way the
neighbouring refusal in the same function is not. That one predates the convention and
is left alone rather than reformatted into this intent.

The runner test needed a phase that is actually red, and the first fixture reached for,
an unresolved question in the intake, is green there: questions turn P5 red, not P0. A
question without options does turn any phase red, which is what the test uses. A detail
of writing a test, recorded because the first attempt looked like a working fixture and
was not.

The order of the record is the order of the work. P0 to P2 before any code, the code
inside P3.
