---
intent: github.com/triplem/xeno#132
phase: 02-design
created: "2026-09-29T16:15:50Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+1dc1483.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a16c233f5bf1af4d78060193466df38c0a1faeb886702abcd3e8cf851402d4ba
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: by-hand
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**`IntentSummary.Status` stops carrying the field and carries the state.** `summarise`
already reads `intent.yaml` and already walks the phases through `Status`, so both
halves are in hand; what changes is which of them the field holds.

**`Decided` is a predicate in one place.** `red` and `provisional` are not decided and
everything else is, which is what `predecessorAllowsStart` enforces for the sequence. It
becomes a named predicate that both call, so the listing and the sequence cannot drift
apart about what a settled phase is.

**Three answers, in order.** The field says `abandoned` and that wins, because an intent
dropped in P1 has no P5 to ask. Otherwise a decided P5 is `complete`. Otherwise the
phase the work has reached, which is what the existing `Phase` field already computes.

**An intent with no phases says so.** `no phases`, rather than an empty column that
would read as complete beside a blank verdict.

**`complete`, not `closed`.** Chosen against the command surface: `intent close` writes
`abandoned` here, so a row reading `closed` would invite the command that makes it
false.

**The verdict column stays.** It says which phase and which verdict, and the state
column says what that amounts to. One is the evidence and the other is the reading,
which is the same division the per-phase listing already makes between `State` and
`Status`.

<!-- xeno:section:alternatives -->
## Alternatives

**Leave the field and add a computed column beside it.** Both answers on one row,
nothing removed, and the row then carries a field that reads `in-progress` next to a
state that reads `complete`. The contradiction is the reason the question was asked in
the first place.

**Add `complete` to section 5's status values and have `phase finish` write it at P5.**
The reading that makes the field mean what a reader expects. It is a specification
change, and section 8 argues against it directly: a merged intent's record is its P5
phase, and a stored flag would be a second answer free to disagree with the verdict,
which nothing would recompute. Rejected on the document and on the principle the tool
states about `PhaseState`.

**Derive `complete` from the phase count rather than the verdict.** Six phases finished
means complete. It would call a phase complete whose P5 is red, which is precisely the
state the sequence refuses to build on, and it would have called every one of this
session's provisional P4s finished.

**Call it `merged`.** Truer to what a reader wants to know and not something the trail
can see: whether a maintainer pressed the button is a fact about the host, which section
8 puts outside the record. `complete` claims only what the verdict claims.

**Drop the column.** Smaller, and abandonment then looks identical to an intent that
stalled, which is the one thing the field was carrying correctly.

<!-- xeno:section:impact -->
## Impact

`internal/runner`: `Decided` as a named predicate, `summarise` computing the state, and
`predecessorAllowsStart` using the predicate rather than its own comparisons.

`cmd/xeno`: nothing, or a column width. The command prints what the summary gives it.

Tests: the three states, the intent with no phases, and that `predecessorAllowsStart`
still refuses a red and a provisional predecessor, since it now shares a predicate with
the listing.

What a reader gains: the column they look at first says whether an intent finished, was
dropped or is in flight, which is the question they were asking it all along.

What they do not gain: whether a complete intent was merged.

What this costs: one word of vocabulary that the specification does not use, `complete`,
defined in the runner rather than in section 5. It is a rendering of the verdict rather
than a state of the artifact, and the design says so in case somebody later looks for it
in the schema.
