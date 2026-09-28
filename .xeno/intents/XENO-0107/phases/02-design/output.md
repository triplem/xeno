---
intent: github.com/triplem/xeno#107
phase: 02-design
created: "2026-09-28T18:51:26Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+907c220.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 2611fc4b103092354187b19a1a5c296cbb6081a33cc8757795e56f206d3b39ae
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

**The rule is a refusal, not a status.** `Status` returns an empty status and an error.
A verdict whose shape makes it uninterpretable is not turned into a word a person acts
on, and the existing refusal for a finding id collision is the precedent: both are cases
where deriving a status would launder the problem.

**It lives in `Status`, and the reason is the callers.** `Invariants` runs once, from
`evaluate`, on checks that were just produced. `Status` runs from there, from
`rewriteStatus` after a person decides something, and from `IntentClose`. `gate.yaml`
lies outside `artifacts_hash` by design, so it is the one file of a judged phase that a
second writer can edit without staling a hash, and `rewriteStatus` is what reads it
back. A rule that must meet a second writer belongs where every derivation passes.

**The condition reads both fields.** `ch.Result == "fail" && len(ch.Findings) == 0`, and
nothing about the other three results. `pass`, `pending` and `not-implemented` carry no
finding as their ordinary state, so a rule phrased over the findings alone would refuse
every passing gate in the tree. This is AC4 written as a condition.

**The error names the gate and describes the verdict.** Not the phase and not a file: a
person meeting it has a check to fix in whatever produced it, and the wording has to
keep them from looking for an artifact. It is the runner's own sentence, since no gate
produced it.

**Placed at the top of the loop, before the findings are walked.** A malformed check is
refused before its findings contribute anything to the status, so one malformed check
cannot be masked by a well formed one earlier in the list. Order in a loop that returns
early is a decision, not an accident.

**Nothing new in what a verdict carries.** No finding, no id, no cause. The verdict does
not record its own malformation, because it is not written at all.

<!-- xeno:section:alternatives -->
## Alternatives

**Read it as red.** Simplest, never refuses, and the phase stays in the pipeline.
Rejected because red is a promise about what can be done next: a red phase carries a
failure somebody decides by approving or overriding a finding by its id, and there is no
finding here. The verdict would name no cause, offer no next step, and clear only by
editing the file it was derived from. A status that cannot be acted on is the same
defect as the green it replaces, one word louder.

**Put it in `Invariants`.** The stated home for verdict shape rules, and the issue
suggests it sits beside them. Rejected on reach, not on fit: `Invariants` is called from
`evaluate` alone, so a `fail` with no finding in a stored `gate.yaml` would pass through
`rewriteStatus` untouched, and that file is editable without staling a hash by
deliberate design. The rule would then be enforced against the writer least likely to
break it.

**Both, guarded twice.** The shape #114 used, where two writers can produce a state and
the second guard is not redundancy. Rejected because the two guards there covered two
different producers; here `Status` covers every path `Invariants` covers and three more,
so the second guard would be the same check on a subset. The #114 argument does not
transfer to a strict superset, and saying so is worth more than the duplicate line.

**Check it in `result()`.** Where `fail` is written. Rejected as the argument of the
issue itself: the producers that do not go through `result()` are exactly the ones the
rule is for, so a guard there is correct and unreachable.

**Refuse any disagreement between `result` and `findings`.** Would also catch a `pass`
carrying findings. Rejected as a different rule with a different argument, and not this
issue's. Named in the non-goals so that its absence is a decision.

<!-- xeno:section:impact -->
## Impact

`internal/gates/gates.go`. `Status` gains three lines and a paragraph on its comment. No
signature change, no new import.

`internal/gates/*_test.go`. The refusal, the three legitimate no-finding results, the
unchanged `fail` with a finding, and the unchanged id collision.

`internal/runner/*_test.go`. One test through a caller, so that AC2 is proved where a
person meets it rather than only at the function. The hand edited `gate.yaml` of AC7 is
a runner level test, since `rewriteStatus` is reached through a decision.

Behaviour on everything that exists: identical. Every gate returns through `result()`,
which writes `fail` only with findings, so the condition is unreachable from the tree.
`gate verify` over every verdict is the check on that claim.

What a later reader gains: `Status` says what a check may look like, so the next
external gate meets a rule instead of a gap. What they still do not gain is any
protection against a foreign gate that reports `pass` while something failed.

One risk worth naming now rather than in P5: a refusal is louder than a red verdict, and
it stops the pipeline instead of recording a failure in it. For a malformed check that
is right, and for a project whose external gate is flaky it will read as Xeno breaking
rather than as the gate being wrong. The error's wording is the only mitigation, which
is why it names the gate.
