---
intent: github.com/triplem/xeno#172
phase: 02-design
created: "2026-10-03T10:35:18Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d3d721517b3e89eec1361e4b07f2dec3b2a1586d9e9ec2bb1cfac200dd1902db
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

**The check is in G-Schema, beside the budget.** Both read the same two files — the profile from
P0, the lock from the phase — and both report a profile mistake without blocking. Putting the link
check anywhere else would mean a second place that knows how to find a profile, and the one that
already exists is three weeks old.

**The finding names the profile, not the base.** The base is the consequence; the profile is the
claim. A finding on the lock would point a reader at a file that is correct about what it was
given.

**It is checked where the profile is read, which is every phase.** `budget` established the shape:
the profile is P0's artifact and is read from P0 whatever phase is being judged. So a mistyped link
is reported at the first gate run after it is written and at every one after that until it is
corrected, which is what a configuration error should do and what a one-phase check would not.

**The runner still skips the link in silence, and that is correct.** `informationBase` resolves
what it can and reports nothing: it is the writer of the lock and the lock describes what the phase
was given. A file that is not there was not given. The report belongs to the gate, which is the
part of this process whose job is to say what is wrong.

**Nothing is added to the register, because there is no register to add to.** M0 closed it this
morning. The decision above — the check's home and why — is recorded here, in the design phase of
the intent that took it, which is the loop the plan switched to. This is the first intent to use
it rather than to describe it.

<!-- xeno:section:alternatives -->
## Alternatives

**Report it from the runner, at `phase start`.** It is where the missing file is discovered, and
the agent would learn about the mistake before writing anything. Rejected because the runner writes
the lock and the lock describes what the phase was given — a file that is not there was not given,
and the runner saying so would make the writer of a record also its judge. It would also mean a
profile mistake reported in one place when the phase starts and nowhere afterwards.

**Make it a refusal.** A profile that names a file it does not have is a configuration error, and
`phase start` already refuses a phase whose predecessor is red. Rejected: #171's criterion asked for
a finding, the budget's precedent is a finding, and section 5's reason for the budget not blocking
applies here too — a project learning the mechanism should see what is wrong rather than be stopped
by it.

**Check it once, at P0, since that is where the profile lives.** Narrower, and it reports the
mistake where the file is. Rejected because a profile written at P0 and mistyped there would then
be reported once, in a phase that may already be sealed, and never again — so the finding would
exist in a verdict nobody re-reads rather than in the phase somebody is working on.

**Also check that the component exists.** It would catch the other half of a mistyped link.
Rejected because a component is a prefix in the project's own vocabulary: it may name a directory, a
module, or something the repository does not model as a path at all, and a check that demanded a
directory would reject a legitimate declaration.

**Fix the byte count in the same intent by storing sizes in the lock.** The user asked for both and
they are one condition and one field. Rejected on the second standing rule: section 5 enumerates
the lock's fields, `files` is written there as a path and a hash, and a `bytes` key is a field the
specification does not have. It is a change to the document and therefore a person's commit made
before the code that follows from it. The wording is offered in the review.

<!-- xeno:section:impact -->
## Impact

**The profile's only unambiguous error becomes visible.** After this, every way a profile can be
wrong that the specification treats as an error is reported: a budget exceeded, and a link whose
document is not there. Everything else in the file is a pattern, and a pattern that matches nothing
is a state rather than a mistake.

**One sharp edge is gone from the profile adoption decision.** The experiment that was deferred —
write a profile on one intent and count what it costs — can now be run without its first mistake
being invisible. The other edge, the byte count measured after the fact, needs the specification
change this intent names.

**#171's unmet criterion is met, by a second intent.** That is the process working as designed and
it is also the cost the figures on #117 record: one condition, one test, and a second intent's worth
of overhead, because the phase that owned the criterion was sealed before the gap was found.

**The first decision after M0 lands in a design phase rather than in a register row.** Where the
check lives and why is recorded in this intent's P2. A reader looking for it in `ASSUMPTIONS.md`
will find the closure paragraph that says where to look instead, which is the mitigation the
closure wrote for itself and the first time it is needed.

**Nothing changes for a project without a profile**, which is every project and this repository.
