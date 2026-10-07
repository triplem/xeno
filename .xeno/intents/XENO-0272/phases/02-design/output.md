---
intent: github.com/triplem/xeno#277
phase: 02-design
created: "2026-10-07T13:29:14Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0a79d51c039222b73f53ad713dd5f41b5be77f56de1b1f307ed82eaf0ea0916b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: Remove the tool field from the hash placeholder exemption in internal/gates/gates.go, so that what makes a by-hand placeholder honest is only a fact about the tree, and release the four findings that produces with gate approve rather than gate override.
      rationale: '#277 put three shapes with the measurement: a clause in section 12 saying the exemption rests on a declared field, a narrower exemption, or accepting it with the reason written down. The measurement is that four artifacts in the trail declare tool: manual, the intakes of XENO-1 and XENO-2 written by hand before M0, and dropping the term turns exactly those two phases red on four context_hash findings and moves nothing else, because WriterlessHash already covers secrets_hash and rules_hash and goneBundle covers their strings_hash. Section 12 now calls the triple a declaration that nothing corroborates, so the one gate that acted on it was relaxing itself on a declaration, which decides a verdict rather than how a register reads. The clause was declined because it documents the gap instead of closing it. The replacement that looked right, keying on the absence of a context.lock.yaml, exempts nobody: all 495 phases carry one. The four findings are correct rather than spurious, since the lock beside each artifact is hashable today, and section 11 forbids rewriting a sealed artifact to satisfy them, so they are released. Approvals rather than overrides, because an override carries an obligation and there is nothing anybody can close.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The term is removed rather than replaced.** The obvious alternative was to key the exemption
on something that is not a declaration, and the candidate was the absence of a
`context.lock.yaml`: a phase with no lock has nothing to hash, which is a fact about the tree.
Measured, and it exempts nobody. All 495 phases in the trail have a lock, including the two
pre-M0 intakes. So the replacement would have been a term that never fires, written to look
like a safeguard, and the honest version of the same change is to take the old term out.

**The comment keeps the reason it removed.** "Nothing produced a manual artifact, so none of
its hashes had a writer" is true, and a reader who finds the obvious exemption missing deserves
to know it was considered and what was wrong with it. A comment that only said the field is
self-reported would leave the next person to rediscover the argument for it.

**The two remaining terms are described by what they have in common**, in the `hashShape`
comment: both are facts about the tree. That is the rule the change establishes, and stating it
is what stops somebody adding a third term from a field. The old comment's "two cases and wrong
in a third" counted cases; the new one names the property.

**The release is `gate approve`, four times, one per finding.** Not two, although they come in
pairs: each finding names its own file, `output.md` and `digest.md`, and the decision is
recorded against the finding. Approving a phase rather than a finding is not a thing the
process offers, and it should not be.

**The reason differs between the two intents by one date.** XENO-1's artifacts were written on
2026-09-21 and XENO-2's on 2026-09-22, and the reason names the date because the whole of it is
that nothing could compute the value then. A single reason for all four would have been one
sentence shorter and would have blurred the fact it rests on.

**The reason says the finding is correct.** "The lock beside it is hashable today, which is why
the finding is correct rather than spurious" — because the easy reading of an approval is that
the gate was wrong, and here it is right and the artifact cannot be changed. An approval that
implied the check was over-eager would invite somebody to narrow it again.

**One row in `docs/assumptions.md`, not a note in the code.** The decision outlives this intent
in two directions: the next person who finds a check that would be simpler with a declaration
needs the reason it was taken out, and the four approvals need to be findable from somewhere
other than two `gate.yaml` files nobody opens. The code comment says what the term was; the row
says what was decided and who decided it.

**No sentence in section 12.** It was the issue's first shape and the maintainer chose the
second. The paragraph already says the triple is a declaration and that nothing corroborates
it; adding that one gate used to read it would be recording a defect's history in a normative
document, and the defect is gone.

<!-- xeno:section:alternatives -->
## Alternatives

**A clause in section 12 and no code change.** #277's first shape and the cheapest: the
paragraph on the triple would say that one gate acts on `tool` and what that buys, and the
relaxation would stay. It was put to the maintainer with the measurement and declined. What it
costs is that the gap stays open while being documented, which is the answer #207 took for the
provider register and is weaker here, because a verdict depends on the field rather than a
reader's confidence.

**Accept it and say why, which is the same thing one step further.** #277's third shape. The
comment already carried the reason and the reason is sound; writing it out again in the register
would have been the clause above without even the normative place. Declined with the first.

**Key the exemption on the absence of a `context.lock.yaml`.** The replacement that looked
right: a phase with nothing to hash is a fact about the tree, not a declaration. Measured and
it exempts nobody — all 495 phases in the trail carry a lock, the two pre-M0 intakes included —
so it would have been a term that never fires. It would also have been generous in the wrong
direction later: a phase whose lock had been deleted would have its `context_hash` excused.

**Bound the exemption by date or by template version.** Accept `by-hand` only for artifacts
created before M0, or only where the template ref is `intake@0.1.0`. It would keep the two
phases green with no approval needed. It puts a date or a version literal in a gate, which is
the kind of thing that is correct for exactly as long as nobody looks at it, and it would make
the gate's behaviour depend on when an artifact claims to have been written — another declared
field, which is the defect under a different name.

**Write the real `context_hash` into the four artifacts.** The only path that keeps every phase
green without an approval, and the one section 11 forbids: rewriting a sealed artifact destroys
the answer to "what changed", and it would move the `artifacts_hash` of two verdicts. It is
also the path that would make the trail say the hashes were computed when they were not.

**Release the four findings as overrides instead of approvals.** Same verdict change, different
record: an override says the work goes on with the finding outstanding. `xeno intent status`
would then list four `obligation close` calls owed forever, on artifacts nobody can correct,
which is how a list of obligations stops being read.

**Leave the two phases red.** Honest, and it makes `gate verify` exit non-zero, so CI fails on
main and every pull request after it. A permanently red trail is a trail nobody reads a verdict
from.

<!-- xeno:section:impact -->
## Impact

`internal/gates/gates.go`: the `byHand` variable removed, one term gone from `honest`, a
comment above it saying what was removed and why, and `hashShape`'s doc comment replaced.

`internal/gates/schema_test.go`: one case inverted, with a comment.

`.xeno/intents/XENO-1/phases/00-intake/gate.yaml` and the same for `XENO-2`: the verdict moves
from green to approved, with two findings each carrying a decision, a reason and the name of
the person who made it.

`docs/assumptions.md`: one row.

**For a verdict.** Two phases in the trail now read `approved` rather than `green`. That is a
visible difference in `xeno intent status` and in anything that reads a status, and it is the
honest record: a finding was raised against them and a person released it.

**For the gate.** `hashes` no longer reads anything the artifact says about its own production.
What makes a placeholder honest is now entirely a fact about the tree, and the comment states
that as the rule rather than as a list.

**For a project that is not this one.** An artifact written by hand with `context_hash:
by-hand` is now a finding. That is the intended consequence and it is a stricter gate than
yesterday's, which by section 13's table is what a major version means: a check that was green
yesterday is red today. The two phases here are the only evidence of how often it happens, and
it is twice in 495.

**For the trail's own history.** Two `gate.yaml` files are rewritten. That is the mechanism for
a verdict that changes rather than an exception to section 11, which is about artifacts; the
four artifacts the verdicts are about keep every byte.

**What is left standing.** Nothing about `model` and `tool_version`, which no gate reads, and
nothing about whether a recorded `tool` value is true. The claim is narrower than it looks: no
verdict depends on the triple any more, and the triple is still a declaration.
