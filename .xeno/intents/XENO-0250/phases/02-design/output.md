---
intent: github.com/triplem/xeno#229
phase: 02-design
created: "2026-10-05T15:21:53Z"
schema_version: "1.0"
runner_version: dev+325b27a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d26a96b1e869d93886f1b828e7892406f3355d3da99dc09e0e721462d2efb26e
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

`QuestionShape` keeps its name and gains the consequence check. It is the gate's question, and
every caller of it today is asking the gate's question, so the function that already exists
goes on being the one the gate uses.

`QuestionAsked(file, o)` is the writer's, and calls `QuestionShape` first. The name says the
difference: the gate reads a question that was written, and the writer judges one being asked.
Calling through rather than beside is what criterion 3 asks for — the writer is the gate's check
plus one, so the two cannot disagree about the part they share.

The recommendation check counts rather than tests presence. Exactly one, so both nothing
recommended and three options recommended are refused, and the message names the count it
found. Section 8 says "the agent's recommendation", singular, and a question with three
recommendations has made no recommendation.

The free entry is exempt from the consequence check and counts for the recommendation. A free
entry cannot carry a consequence, and a recommendation that lands on it would be the agent
recommending that the person think of something else, which is the shape section 8's first
sentence exists to prevent — so it is not special-cased for the recommendation, and
recommending it stays possible and stays its author's problem.

Both functions carry the reason for the split, and so does `exchange.go` at the call site.
Three places, which is more than this project usually repeats itself, and the reason is that
the asymmetry reads as an oversight from any one of them: the gate looks lax, the writer looks
arbitrary, and the call site looks like it picked the wrong function. XENO-3's Q-2 and the
`DIVERGENT` line are named, because the next person to tidy this will otherwise move the check
and discover the divergence themselves.

The consequence check's message names the counts, "has 2 of 3 options with no consequence",
rather than the first offending option. A question is read whole and a count tells its author
how much work is left, where a single name invites fixing one and running again.

No register row. The split is a fact about this code and is written where the code is; what
outlives this intent is the two specification gaps, and those become issues, which is the
register's own test from #243 applied honestly rather than as a habit of adding rows.

<!-- xeno:section:alternatives -->
## Alternatives

Putting the recommendation check in the gate and overriding XENO-3's finding was the option the
maintainer weighed this against, and it is the one with the better end state: the clause gets a
real reader and the single historical violation is recorded as an accepted deviation rather
than left unexamined. It was rejected on cost. It rewrites a pre-M0 intent's `gate.yaml`, needs
a second person's reason, and D-6 settles that a decision command refuses a verdict it finds
stale — so it is a multi-step operation on the oldest part of the trail in exchange for one
question nobody will ask again. The writer gets the same forward guarantee for one function.

Shipping only the consequence check and deferring the recommendation entirely was the third
option. Rejected because the recommendation is the half that carries the meaning: an option list
without consequences is unreadable, and a question without a recommendation is the agent handing
back the whole decision, which is what section 8's first sentence names as the failure.

Version-gating the check was considered and found to have nothing to gate on. G-Policy's
`judgedUnder` keys on `rules_hash`, which exists because a rule set is data; a gate's own checks
have no such anchor, and `schema_version` is "1.0" on XENO-3 and on everything written today, so
it cannot tell them apart. Inventing a marker would be a new field.

Keeping `QuestionShape` as the writer's function and giving the gate a narrower one was
considered, and is the same split with the names reversed. Rejected because every existing
caller of `QuestionShape` is a gate path except one, so the narrower function would be the one
with more callers and the name would be doing the opposite of its job.

Writing the recommendation check as a `review` rule rather than code was considered. It would
put the clause in the rule set where a person answers it in the P5 checklist, which is honest
about who reads it — and it would answer it once per intent rather than once per question, which
is the wrong grain for a shape.

Requiring a consequence on the free entry was considered and rejected on section 8's own
wording and on the project's own fixture, both of which treat the free entry as a further
option rather than as one of the two to four.

<!-- xeno:section:impact -->
## Impact

Two files changed and two test files. `internal/gates/gates.go` gains the consequence check in
`QuestionShape` and `QuestionAsked` beside it. `internal/runner/exchange.go` calls
`QuestionAsked` instead of `QuestionShape`, one line. Tests in `internal/gates` for the gate's
half and in `internal/runner` for the writer's.

Nothing is added to the model, the rule set, the gate list or any artifact shape. `Consequence`
and `Recommended` were already there; what changes is that two fields stop being decoration.

The trail is unaffected, which is the point and is measured rather than claimed: the consequence
check finds nothing in ten artifacts carrying questions, and `gate verify` stays at exit 0 over
375 verdicts. Had that come out otherwise this intent would have had a different shape.

What a question's author gains is a refusal at the moment of asking. A question with a bare
option or no recommendation is refused by `xeno question record` with the count, where today it
is written, sealed, and read by a person who may or may not notice. What they do not gain is
help with the reason for the recommendation, which still has nowhere to go.

What the gate gains is one of the three requirements. It now reads the count, the free entry and
the consequence, and is knowingly silent on the recommendation. So `docs/clause-readers.md`'s
entry for section 8 improves and does not close, and the audit's next pass should say so rather
than mark the clause read.

The asymmetry is the cost and it is permanent until somebody decides otherwise. A hand-written
artifact can carry a question with no recommendation and pass every gate, which is exactly the
class of defect this repository keeps finding — a writer that enforces and a gate that does not,
so the guarantee holds only for work that went through the command. #242 left the same shape
behind for the review checklist and said so; this says so too.

The two specification gaps leave this intent as issues rather than as code: no field for the
recommendation's reason, and no rule about sequence. Both need a person's commit to a normative
document before anything can read them, and filing them is the whole of what this intent can
honestly do about either.
