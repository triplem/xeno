---
intent: github.com/triplem/xeno#71
phase: 00-intake
created: "2026-09-26T21:21:34Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+50001d4
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 922d578b75d7004db1a5a02f31f2d765dfe3fc5d7db28ba2a2a5e4aabd02f5a9
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

A gate that does not apply to a phase yet is left out of `gate.yaml`, which A5 approved:
the phase table of section 7 is the authority on applicability and `not-applicable` is
not in the result set section 5 defines. What the shape cannot express is the difference
between a gate left out because it does not apply and one left out because it did not
run.

`xeno gate verify` did not close that gap either. It recomputes with the same table and
compares the status and the `artifacts_hash`, so a gate missing from both sides matches,
and `gate.yaml` is not covered by `artifacts_hash` at all, which means a verdict can
lose a check without any hash moving.

<!-- xeno:section:scope -->
## Scope

A comparison in `gate verify` between the gates a committed verdict carries and the
gates that apply at its phase, reported in both directions: one that applies and is
missing, and one whose phase has not come.

`gates.Applicable` is the single place applicability is read from, which is the table
itself, so the check cannot answer differently from the run.

<!-- xeno:section:context-rationale -->
## Why this context

**Against the table, not against a second run.** The obvious implementation compares the
committed verdict with a recomputed one, and it cannot work: the recomputation reads the
same table and leaves out the same gate, so the two agree exactly when the defect is
present. The comparison has to be against the applicability rule rather than against
another instance of the same code.

**In verify and not in a gate.** G-Schema judges the shape of a phase directory, which
is the argument for putting it there, but `gate.yaml` is the verdict a gate's own result
is written into, and a gate reporting on the completeness of that file is reading its
own output. A52 records the choice.

**Both directions, because the error is symmetrical.** A gate that vanished is the case
worth catching. A gate that appears before its phase is a verdict claiming a check that
could not have happened, which is the same untruth pointing the other way and costs one
loop to catch.

**Shown against the trail.** Removing `G-Trace` from one committed verdict makes verify
name it and exit 1, with every hash still matching. That is the demonstration that the
gap was real rather than theoretical, and it is also why the test writes a verdict file
directly instead of going through a command: no command can produce the state being
guarded against.