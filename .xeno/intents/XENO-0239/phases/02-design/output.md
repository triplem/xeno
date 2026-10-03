---
intent: github.com/triplem/xeno#213
phase: 02-design
created: "2026-10-03T20:05:00Z"
schema_version: "1.0"
runner_version: dev+427eeb7
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f1ebc28f9b5fd633906d06f7867ece2a3d476ae3b803ae0b13fab08caec3f2cf
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The plan changes, not the process definition.** Section 5's reasoning is the stronger of
the two and the first standing rule points the same way, so there was no case for touching
it. The plan's completion condition was the unsatisfiable half.

**The digest carries the record.** It is written at `phase finish`, sits inside
`artifacts_hash` and so cannot be altered afterwards without a divergence, sits outside
`context_hash` and so does not disturb the lock's seal, and the phases after it read it.
Those four properties are the resolution; nothing else in the trail has all four.

**The reason travels with the clause.** WP8 now states why the lock cannot carry the
record rather than pointing at section 5. A pointer would have been shorter, and the
question would have been reopened by the next reader who met WP8 first — which is how this
contradiction survived as long as it did.

**The record is called a self-report in the plan.** Nothing in the harness reports what an
agent opened, so the clause says so. A91 and #207 are the precedent: a stated convention
whose reader is a person is honest; one that implies a measurement is not.

**No gate.** A92 records that, with A90's reason: a check that an out-of-profile read was
declared would have to know what was read.

<!-- xeno:section:alternatives -->
## Alternatives

**Add `out_of_profile` to section 5's lock.** Rejected by the maintainer, and the reasoning
is worth keeping. The lock is written once, before the agent starts, and `context_hash`
seals it; a field filled at phase finish makes it two files in one, written at two times,
and the paragraph explaining that the lock states what a phase was given would have to be
rewritten. It would also be the largest of the three changes and in the document every
verdict in the trail is judged against.

**Drop the clause, as A89 dropped the plugin resolution order.** A real answer, and the
closest call. Nothing observes what an agent opened, so the record is the agent's own
account either way, and A89's precedent is exactly this: a clause with no reader removed
rather than built under a condition. It lost because WP8's intent is sound — a budget that
cannot tell it was exceeded is a budget in name — and because the digest already exists,
so keeping the clause costs a sentence rather than a mechanism.

**Leave both documents and let WP8 stay unfinishable.** Rejected. That is the state that
produced this issue, and an unsatisfiable done-when is worse than either resolution because
it reads as work outstanding rather than as a question nobody answered.

**Block an out-of-profile read instead of recording it.** Not considered seriously: WP8 says
recorded, not blocked, and the runner cannot block a read it does not observe.

<!-- xeno:section:impact -->
## Impact

One paragraph and one sentence of `docs/implementation-plan.md`, in its own commit ahead of
this intent, and one row in `ASSUMPTIONS.md`. No code: `ContextLock` already matches section
5, and the condition that changed was never implemented.

What changes is what WP8 can be finished against. Before, its done-when could not be met
without breaking section 5; now it can be met by an agent naming an out-of-profile read in
the digest it already writes.

What does not change is anybody's ability to tell whether that happened. The record is a
self-report, the budget is still judged against `files`, and an agent that reads outside the
profile and says nothing leaves no trace. The plan now says that plainly, which is the
difference between a weak record and a misleading one.

WP8's first half remains unmet and becomes its own issue. Keeping it inside this one would
have let a contradiction and a missing feature share a close, and only one of them is fixed.
