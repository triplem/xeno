---
intent: github.com/triplem/xeno#107
phase: 05-review
created: "2026-09-28T18:58:17Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+907c220.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 068b4a27a1df187c2b26afba9b5631841ce9ea2313866ed92a9365c2c27d2d91
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**Nothing under `docs/` changed.** Section 14 already makes external gates the extension
point and section 5 already gives them their provenance. This adds a rule about what the
runner does with a malformed verdict, which neither document had to say.

**Nothing was invented.** No field, no gate, no finding, no dependency, no signature
change. The rule reads two values a check already carries and returns the error `Status`
could already return.

**The rule spares what it must.** Four of the eight mapped rows must pass with and
without the change, and they do: `pass`, `pending` and `not-implemented` with no
finding, and `fail` with one. A rule phrased over the findings alone would have refused
every passing gate in the tree, which is what AC4 exists to catch.

**Refusing rather than reddening is argued, not assumed.** Red promises something can be
decided and there is no finding to decide; the design says so and the alternatives
section records what red would have produced.

**The location is argued from the callers.** `Invariants` would have been the tidier
home and covers one path; `Status` covers that one and two more, including the stored
verdict a person can edit because `gate.yaml` sits outside `artifacts_hash` on purpose.
Guarding twice was rejected because one caller set contains the other, which is the one
place the #114 precedent does not transfer.

**Every verdict in the repository still matches.** `gate verify` over 67, which is the
evidence that no gate of this runner reaches the refusal.

**The phases were written in their order.** P0 to P2 before any code, the code inside
P3, the tests run against the reverted file inside P4.

**The comment carries the reason and not the history.** It says who the rule is for and
why it is in this function; that the issue left the choice open belongs to this record
and to the commit.

<!-- xeno:section:release-notes -->
## Release notes

A check reporting `fail` while carrying no finding is now refused rather than counted as
nothing. `xeno phase finish`, `xeno gate run`, a decision that rewrites a stored status,
and `xeno intent close` all stop with an error naming the gate, instead of deriving a
status from a verdict that says a failure occurred and cannot say what it was.

Nothing changes for any gate of this runner. Every one reports through the same helper,
which writes `fail` only where there are findings, so no phase in this repository is
affected and every verdict still matches.

What changes for whoever writes an external gate: a failure has to name what failed. A
scanner wrapper that reports a failure it cannot attribute to a file meets an error
rather than a green phase.

A verdict edited by hand is covered as well. `gate.yaml` lies outside `artifacts_hash`
so that evidence can be attached later without breaking a seal, and the same exclusion
lets a verdict be edited; a `fail` with no finding written there is refused the next
time a status is derived from it.

Unchanged: a `fail` carrying a finding is red, `pass`, `pending` and `not-implemented`
with no findings are what they were, and the refusal for two findings sharing an id
keeps its own words.

<!-- xeno:section:residual-risk -->
## Residual risk

**A refusal stops the pipeline where a red verdict would have recorded a failure in
it.** For a malformed check that is the intent, and for a project whose external gate is
occasionally malformed it will read as Xeno breaking rather than as the gate being
wrong. The error names the gate and that is the whole mitigation. Only a project with a
real external gate can say whether it is enough, and WP4 brings the first one.

**The guarded state has no producer.** It exists in the tests and nowhere else in this
repository, so the rule is correct against the shape a check can take and unproven
against a foreign gate that actually takes it. A guard placed before its data arrives is
worth having and is not the same as a tested one.

**Two of the three callers are read rather than tested.** All three pass `Status`'s
error outward through the same two lines, and only the path through a decision has a
test. An argument, not a proof, and named so that the next person adding a caller knows
the guarantee rests on a convention.

**The opposite shape is untouched and undetectable.** A gate reporting `pass` while
something failed produces a green phase, here and anywhere else in a local runner.
Nothing in this change reaches it and nothing can.

**One sentence in the same function is now inconsistent.** The refusal beside this one
runs past eighty-eight characters on one line; the new one is split. Left as it is
rather than reformatted into this intent, which is a small debt recorded rather than
paid.

**Accepted with the five named.** None blocks the change. Each is smaller than the green
verdict it replaces, which said a phase was fine on the word of a check that had said it
was not.
