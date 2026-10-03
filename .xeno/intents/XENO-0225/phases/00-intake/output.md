---
intent: github.com/triplem/xeno#172
phase: 00-intake
created: "2026-10-03T10:33:58Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 355b353ef38ec51d71acb083e9c539ff5f834f646d955646c053872510e63b05
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

A context profile declares links between code and documentation, and #171 made a declared link's
document part of the information base. A link naming a document that is not in the tree is skipped
in silence: `informationBase` cannot hash a file that is not there, so it moves on, and the one
field in a profile that names a specific path rather than a pattern is the one field whose mistakes
nobody reports.

**It is an acceptance criterion of #171 that #171 did not meet.** Its 01-requirements asked for "a
finding against the profile rather than a silently missing file". P3 took a different path, no test
had been written for it, so the suite was green; P4's criterion table found the gap after P3 was
sealed, and what is sealed is never rewritten. The intent was released with the hole written into
its release notes, its residual risk and this issue.

**Why it matters more than its size.** Everything else in a profile is a pattern, and a pattern
that matches nothing is indistinguishable from a pattern whose files have not been written yet. A
link is a specific path, declared because section 5 forbids inferring one: "an inferred mapping is
an assumption, and assumptions in this process are either registered or absent". So a mistyped
link is the profile's only unambiguous error, and it is the one that produces no finding.

**And it is the sharp edge in the way of adopting a profile here.** The decision about writing a
profile for this repository was deferred until this is fixed and until the byte count stops being
measured after the fact, because adopting a mechanism whose first mistake is invisible means the
first mistake is invisible.

**The byte count is the other half and it is not this intent's to fix.** G-Schema measures the
recorded context's size from the tree when the check runs, because the lock carries paths and
hashes and no sizes — so a phase inside its budget when it ran can be over it a week later, and
`gate verify` will say so about a sealed phase. Recording the number means a field in the lock, and
section 5 enumerates the lock's fields: that is a change to the specification and therefore a
person's commit, made before the code that follows from it. This intent names it and stops.

<!-- xeno:section:scope -->
## Scope

**In scope.** A link naming a document that is not in the tree becomes a G-Schema finding against
the profile, naming the component, the path it declared and that the file is not there. A link
whose document exists still joins the information base, as #171 left it. Tests for both.

**Out of scope, and each for its own reason.**

Recording the byte count. It needs a field in the lock and section 5 enumerates the lock's fields,
so it is a specification change and a person's commit before any code. Named in the problem and in
the review, with the wording offered, and not taken.

Whether a missing link's document should stop a phase from starting. Today it does not: the base is
resolved at `phase start` and the finding arrives at the gate, which is the same shape as every
other profile mistake. Making it a refusal would be a second mechanism for one error and is not
what the criterion asked for.

Adopting a profile for this repository. That is the experiment this fix unblocks and it comes
after, on one intent, with the releases counted.

Any other field of the profile. `include` and `exclude` are patterns and a pattern matching nothing
is a legitimate state. `budget` has its own finding since #171.

**One boundary worth naming.** The finding belongs to the profile and the profile is a P0 artifact,
so it is reported wherever G-Schema runs with that profile in view — which is every phase of the
intent, because the budget check established that the profile is read from P0 regardless of which
phase is being judged. A mistyped link therefore shows at the first gate run after it is written
and keeps showing until it is corrected, which is what a configuration error should do.

<!-- xeno:section:context-rationale -->
## Why this context

Section 5's context economy subsection is read for the sentence that makes this a finding rather
than a tolerance: links are declared, never inferred, because an inferred mapping is an assumption
and assumptions here are either registered or absent. A declaration whose target is missing is
neither registered nor absent — it is a claim about a file that does not exist.

#171's own 01-requirements is read as the specification for this piece, because the criterion it
wrote is what this intent exists to meet, and its 04-verification is read for how the gap was
found: by the criterion table rather than by a test.

`internal/runner/runner.go` is read for `informationBase` as #171 left it, which is where the link
is skipped, and for the ordering decision around it, since the fix must not change which files
reach the base or in what order.

`internal/gates/gates.go` is read for `budget`, which is the precedent: a check that reads the
profile from P0 and the lock from the phase, reports against G-Schema, and does not block. The link
check is the same shape with a cheaper question, and reading it first is what decided that this
belongs beside it rather than in the runner.

`internal/gates/budget_test.go` is read for the fixture that writes a profile and a lock, because
the same fixture answers this with one more field.

The process definition's gate list is read once more to confirm that G-Schema is where a finding
about an artifact's own content belongs, and that no other gate has a claim on it.

Nothing outside the repository is needed. The fix is one condition and the only question it raised
— whether the finding belongs to the runner or to a gate — was answered by the shape of the check
that landed last week.
