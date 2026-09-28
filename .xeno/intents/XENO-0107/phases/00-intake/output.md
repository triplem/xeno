---
intent: github.com/triplem/xeno#107
phase: 00-intake
created: "2026-09-28T18:50:02Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+907c220.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: b57eade4bafd4e2dc5b281ce4b4139660e4924a9db2db935f68e9a0ee21a9222
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

`Status` derives a phase status from findings alone. It reads `ch.Result` for one value,
`pending`, and never for `fail`, so a check reporting a failure while carrying no
finding contributes nothing to the status it is supposed to decide:

    Status([]model.Check{{Gate: "G-External", Result: "fail", Provenance: "external"}})
    // → "green", nil

No gate of this runner can produce that check. Every one of them returns through
`result()`, which writes `fail` only where there are findings, so the hole is closed by
construction and by nothing else. A second writer reopens it.

The second writer is the one the hole is about. Section 14 makes external gates the
extension point, section 5 gives them `provenance: external` and a rule of their own
about decisions, and an external gate is precisely the code whose findings Xeno does not
word and cannot predict. A scanner wrapper reporting a failure it cannot attribute to a
file is an ordinary thing, not a malformed one, and today it would be read as green.

There is a second writer already in the tree, and it is not hypothetical. `gate.yaml` is
excluded from `artifacts_hash`, which is what lets a later run attach evidence without
breaking a seal, and the price is that a verdict can be edited without making any phase
hash stale. `rewriteStatus` then derives a status again from the checks it finds in that
file, so a `fail` with no finding written there by hand reaches `Status` without passing
through any gate.

This is the same argument `Invariants` makes and A25 records: what breaks a promise
about how findings reach a verdict is a second path into it rather than a change to the
first, so the verdict is checked and not the path. A `fail` carrying no finding is that
second path's other shape.

<!-- xeno:section:scope -->
## Scope

`Status` refuses a check that reports `fail` and carries no finding, returning an error
rather than a status. The refusal sits beside the one already there for a finding id
collision, which is the same kind of rule: a verdict whose shape makes it unusable is
not turned into a status anybody can act on.

Tests for the refusal, for the three callers that reach it, and for the states it must
not disturb: `pending` with no finding stays provisional, `not-implemented` stays what
it is, and `pass` with no finding stays green, which is the ordinary case for every gate
in the tree.

Not `Invariants`. It is the stated home for verdict shape rules and it is called from
one place, so the rule would not reach a status derived again from a stored verdict. The
reason it is not the choice is written down in the design rather than left as an
omission.

Not a finding. Refusing is not reporting: no new finding id, no new cause, and nothing
added to what a verdict carries.

Not the external gate contract. What an external gate must send is section 14's, and
this change decides only what happens to a verdict that arrives without it.

<!-- xeno:section:context-rationale -->
## Why this context

**Refused rather than read as red, because red names something that is not there.** A
red phase is one carrying a failure nobody has decided, and the way somebody decides one
is by approving or overriding a finding by its id. A `fail` with no finding offers
nothing to decide, so a red verdict from it would be a dead end: no cause, no next step,
and no way to clear it except editing the file the status was derived from. Refusing
says what is true, that a check which failed without saying what failed is not a verdict
anybody can act on.

**In `Status` rather than in `Invariants`, because of who calls each.** `Invariants` is
called once, from `evaluate`, where the checks have just been produced. `Status` is
called from there too and from `rewriteStatus`, which derives a status again from the
checks stored in `gate.yaml` after a person decides something, and from `IntentClose`.
Since `gate.yaml` lies outside `artifacts_hash` by design, it is the one file in a
judged phase a second writer can edit without making a hash stale, and `rewriteStatus`
is the path that reads it back. The rule belongs where every derivation passes, not only
the first.

**Beside the id collision refusal, because it is the same rule.** `Status` already
refuses rather than returns for one shape it cannot interpret: two findings with one id,
where a single decision would cover two things. Both refusals say the same thing about
the same function, which is that deriving a status from a verdict that cannot mean what
it says would launder the problem into a word somebody trusts.

**Checked at the verdict, not at the producer.** The consistent move of this package,
and A25's answer to the same question. A rule about what a check may look like cannot be
enforced in `result()`, because the checks that will not go through `result()` are
exactly the ones the rule is for.

**What stays broken on purpose.** An external gate reporting `pass` while something did
fail is not detectable here or anywhere, and nothing in this change reaches the contents
of a foreign finding. This closes the shape that is locally visible and says so.
