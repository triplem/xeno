---
intent: github.com/triplem/xeno#183
phase: 02-design
created: "2026-10-03T18:39:15Z"
schema_version: "1.0"
runner_version: dev+0462eb7.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f7e7c9733fd4c43dc18250cf45c0b2622ba241d2b5b03859ed1da5e5713bb220
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

**The paragraph is replaced, not annotated.** This project's convention for a passage changed a
second time, and a removal annotated in place reads as a mechanism with a caveat rather than a
mechanism that is gone.

**It carries both measurements rather than the conclusion.** The silent-unjudging figure — exit 0
over 273 verdicts, 75 phases unjudged — and the un-vendored one — no template resolves, so no phase
renders. A removal somebody can reinstate by disagreeing with an absence is a removal that will be
reinstated; one they have to disagree with evidence to reinstate is a decision.

**`XENO_PLUGIN_ROOT` leaves the list and the other three stay.** The list is what the runner may
see, and the runner may not see a plugin root any more. Leaving it listed while the paragraph below
said there was no order would be the half-agreeing document this change exists to avoid.

**One sentence is kept about the condition.** Not the order, not a plan to add one: the condition any
override would have to meet, which was discovered by measurement and is worth more written down than
rediscovered. It also marks the removal as a decision rather than an oversight for whoever reads the
section next.

**The reason the client fallback is unreachable is kept too.** A reader who finds a mechanism missing
is better served by the reason than by the silence, and this one is counter-intuitive: the fallback
looks useful and cannot work.

**The plan's sentence says what replaced the clause**, rather than quietly carrying three variables
where it carried four. The two documents are both normative and this is the one place they overlapped
on the subject.

**Nothing in the code changes.** The first commit is documents only, which the first standing rule
requires and which is also simply true: nothing implemented the order.

**The comments follow in the second commit.** `internal/plugin`'s package comment and the entry
point's both argued for not implementing something section 7 described; they now describe a section
that does not describe it.

**The test keeps its assertion and changes its reason.** The thing worth preventing is a root
arriving from the environment, which does not stop being worth preventing when the document stops
naming a way to do it. A test deleted with the clause would have left the entry point free to grow
the behaviour back.

**`internal/plugin`'s comments are rewrapped to 88.** They were written at 95 in an earlier intent of
mine and no check enforces the width, so it went unnoticed; the file is being edited anyway.

<!-- xeno:section:alternatives -->
## Alternatives

**Implementing the order with the override honoured only where the binary carries an anchor.** The
alternative put to the maintainer beside this one. It would have made section 7 true as written for
an adopter and safe for a development build, at the cost of a condition the section does not state,
eight call sites threaded, and a mechanism to maintain. Declined: a clause with no reader is cheaper
to delete than a mechanism with a guard is to keep, and nothing in 85 intents has wanted an override.

**Implementing it as written and amending A42.** Declined for what it would cost rather than what it
would build: A42 is the property the verdicts rest on, and weakening it to accommodate a mechanism
nobody uses is the wrong direction for the one claim this process makes about reproducibility.

**The split by purpose — the gate path always vendored, everything else honouring the order.** This
was recommended twice and is retracted. It breaks the one case the client fallback existed for: a
project relying on a client's installed plugin could not resolve templates, so the gate path pinned
to the vendored tree makes that project unable to render at all. And then the fallback turned out to
be unreachable for exactly that reason anyway, which is what retired the design.

**Keeping the clause and marking it unimplemented.** Declined: the gate list is a budget and a
declaration, and `not-implemented` is a state a gate can report. A resolution order has no such
state — a document that describes a mechanism nothing reads is the class of defect four of this
session's findings belong to.

**Removing the client fallback's reasoning along with the mechanism.** Declined. The fallback looks
useful and cannot work, which is worth a reader's time; silence would invite its reinstatement.

**Leaving the comments as they were.** Declined — they argue for not implementing something the
document no longer asks for, which is the stale half-agreement this repository has already met twice
in `ASSUMPTIONS.md`.

<!-- xeno:section:impact -->
## Impact

**`docs/process-definition.md`.** Section 7's variable list loses `XENO_PLUGIN_ROOT`. The resolution
order paragraph is replaced by four: what the plugin is, why the order was removed with the
measurement, why the client fallback was unreachable with the other measurement, and the condition an
override would have to meet. Already committed, alone.

**`docs/implementation-plan.md`.** WP7's sentence names three variables and says the root and the
order were on it. Same commit.

**`internal/plugin/plugin.go`.** The package comment described the order as not implemented on
purpose and pending a decision; it describes a section that no longer has it. The file's comments are
rewrapped to 88 while it is open.

**`.xeno/plugin/bin/xeno-env.sh`.** The paragraph saying `XENO_PLUGIN_ROOT` is deliberately not set,
and why, becomes the paragraph saying there is none to set, and why. Rewrapped with it.

**`internal/plugin/plugin_test.go`.** The assertion stays; its comment says why it stays.

**`ASSUMPTIONS.md`.** A84's open clause closed. A89 for the decision, with both measurements and the
three shapes.

**No change to any other file.** No gate, no rule, no artifact, no field, and nothing about how the
plugin is found — which is the point: the code was already what the document now says.

**What a reader should expect.** Nothing to behave differently. `gate verify` at 285 and exit 0
before and after, the suite unchanged, and the only observable difference is that a reader of section
7 no longer finds a mechanism they cannot use.
