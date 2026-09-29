---
intent: github.com/triplem/xeno#109
phase: 00-intake
created: "2026-09-29T18:35:20Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+3ec2429.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d7c7c2ca84561cbcd96b6b8880bd20480a56902723790393aef9e0b244a1132c
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

Appendix B rests the normalisation of `artifacts_hash` on a property it says is checked
rather than assumed:

> Every file so covered is one Xeno wrote itself and is therefore text, which is what
allows the > normalisation below to apply without any detection. That is not an
assumption but a checked > property: G-Schema reports a file it does not recognise in
the phase directory, so anything else > that lands there is a finding before it reaches
a hash.

The check exists for a phase, as `directoryFindings` against `model.KnownPhaseFiles`,
and it reports an unknown file and an unknown directory alike, allowing only `evidence`.

The intent level hash has no equivalent. `IntentClose` computes it with
`hashing.DirHash` over the intent directory, excluding `gate.yaml`, and nothing has
judged what lies there. So a stray file — a binary, an editor backup, a half-written
note — enters the one verdict that closes an intent and is CRLF normalised on the way in
as though it were text, and the appendix's justification does not reach it.

Section 4 already enumerates what belongs: `intent.yaml`, `assumptions.yaml`,
`learning.yaml`, `gate.yaml`, and the `phases/` directory the hash does not descend
into. `model.KnownPhaseFiles` has no intent level twin, which is the whole of the gap.

It is the narrower of the two cases and the louder one. A phase has six verdicts around
it and this has one, the one a dropped intent is measured by, and section 7 says that
without the intent level hash the verdict closing an intent would be the only one not
bound to what it judged. A hash bound to something nobody checked is that objection one
step in.

No intent in this repository is abandoned, so nothing is wrong with any verdict here.
The exposure is the next abandonment.

<!-- xeno:section:scope -->
## Scope

`model.KnownIntentFiles` names what section 4 lists at the intent level, and a check in
`CompleteOnClose` reports anything else: an unknown file, and an unknown directory where
only `phases` belongs.

`CompleteOnClose` rather than an intent mode of G-Schema, which is what #109 leans to
and what the shape of the code agrees with. G-Schema is a phase gate, its `Ctx` carries
a phase, and `CompleteOnClose` is already G-Complete's second mode: the one an abandoned
intent meets, which is the only occasion an intent level hash is computed at all. The
finding then lands beside the other findings an abandoned intent gets.

Tests: a stray file, a stray directory, the four known files, and `phases/` left alone.

Not the phase level check, which exists and works.

Not `hashing.IntentExcluded` or the hash itself. What is covered and how it is
normalised are Appendix B's and unchanged; this makes its stated property true at the
intent level.

Not a new gate. G-Complete has the mode already.

Not `cost.yaml` at the intent level. Section 4 puts it in a phase and nothing writes one
here.

<!-- xeno:section:context-rationale -->
## Why this context

**The check goes where the hash is computed, which is the abandoned case and nowhere
else.** An intent level `artifacts_hash` exists only in the `gate.yaml` that `intent
close` writes, because a merged intent's record is its P5 phase and section 8 says so.
`CompleteOnClose` is G-Complete's mode for exactly that occasion, so the check and the
hash have the same one invocation and cannot fall out of step.

**A twin of the phase level function rather than something new.** `directoryFindings`
already establishes the shape: read the directory, allow the one legitimate
subdirectory, allow the enumerated files, report the rest as a finding before anything
reaches a hash. Mirroring it means the two levels answer the same question the same way,
and a reader of one recognises the other.

**Directories are reported too, although the hash does not descend into them.**
`DirHash` skips a directory, so a stray one cannot corrupt the value, and the phase
level check reports one anyway. The reason holds here: the property Appendix B claims is
that everything in a judged directory is something Xeno wrote, and a directory nobody
wrote is a place files will appear. Reporting it is the cheaper half of the same rule.

**Section 4's list is the authority and it is short.** `intent.yaml`,
`assumptions.yaml`, `learning.yaml`, `gate.yaml`, and `phases/`. No `cost.yaml`, which
section 4 puts in a phase; nothing for evidence, which is a phase's; and no room for a
note somebody meant to keep.

**What this does not make true.** Appendix B's property is now checked at both levels,
and being checked means reported rather than prevented: a stray file still enters the
hash of the run that reports it, because the hash is computed over the directory as it
is and the finding says what the directory contained. Section 4's arrangement is that
the verdict records both, which is the same compromise the phase level has lived with
since it was written.

**No intent here is abandoned, so nothing is repaired by this.** It is a guard placed
before its data, like the one XENO-0107 put in `Status`, and the honest statement is
that the first abandonment is when it earns its keep.
