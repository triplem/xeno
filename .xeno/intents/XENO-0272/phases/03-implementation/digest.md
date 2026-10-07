---
intent: github.com/triplem/xeno#277
phase: 03-implementation
created: "2026-10-07T13:34:42Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 64154405566a46768fbead6ccd6c07c1b0be176b0d3dc9824e33c57f7ae1aa80
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`honest` no longer reads the `tool` field: the `byHand` variable is gone and the expression
keeps the two terms that are facts about the tree. The comment above it says what was removed,
that its reason was sound and its premise was not. `hashShape`'s doc comment is replaced whole
and names what the two remaining cases have in common rather than counting them.

Two test tables, because the removed behaviour was specified twice under one name. In
`internal/gates/schema_test.go` the case asserting the finding is inverted; in
`internal/runner/runner_test.go` the case of the same name, which asserts the colour a phase
comes out as, moves from green to red. Only the first was found before this phase was first
judged, and `go test ./internal/gates/` passed, so the package under change was green while
the suite was not. This phase was rewritten and judged again.

Two `gate.yaml` files move from green to approved, carrying four findings — `F-0c8115`,
`F-6ec27f`, `F-187a0e`, `F-e0c34d`, all `context_hash says by-hand where a writer exists` —
each with a decision, a reason and the name of the person who made it. Approvals, so nothing is
owed. The four artifacts themselves are untouched and still read `context_hash: by-hand`.

A101 in the register carries the decision, the measurement, the rejected replacement, why the
findings were released rather than fixed, and that the gate is now stricter.

Two deviations. The replacement term keyed on a missing `context.lock.yaml` was measured and
found to exempt nobody, and that count was taken while writing the design rather than before
it, so P2 reads as though the conclusion came first; the lock count is the load bearing fact
for anybody revisiting this. And the behaviour was specified in two test tables, which is why
this phase was judged twice; the learning record carries it, and the `no_finding` this phase
first wrote was withdrawn by hand to make room for it.

Build, suite, `gofmt`, `vet` pass. `gate verify` exits zero over 498 verdicts; it exited 1
between the code change and the approvals.
