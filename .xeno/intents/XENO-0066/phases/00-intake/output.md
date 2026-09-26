---
intent: github.com/triplem/xeno#66
phase: 00-intake
created: "2026-09-26T21:31:31Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+9a233d5
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: fcb286c9735c1a1b0626b266c9d6d436fa05ee28f819147df1fbb97613672c62
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

A decision on a finding is attached in one place, `carryForward`, and the rule that an
external finding never keeps one holds because every check that can produce a finding
goes through that function. Nothing enforced the routing. A25 recorded it as a
constraint on WP4 rather than as a property of the code: whoever adds external gates
routes their checks through the function or the rule stops holding, and a test that
passes would say otherwise.

The test cannot say otherwise, which is the part worth stating. Every test of
`carryForward` exercises the function, and what would break the rule is a second path
into the verdict that never calls it.

<!-- xeno:section:scope -->
## Scope

The properties the routing was supposed to guarantee, checked where a verdict is
produced: every finding carries the id its content hashes to, and no finding of an
external check carries a decision. A run that produces anything else refuses.

Not the external gates themselves. WP4 builds those, and what this changes is that it
cannot break the rule quietly while doing so.

<!-- xeno:section:context-rationale -->
## Why this context

**The verdict is checked rather than the path.** A guard on `carryForward` would protect
the code that already obeys the rule. What the rule needs protecting from is code that
does not call it, so the check reads the finished set of checks in `compute`, which is
the one place every verdict passes through, and fails the run there.

**An id that is derived can be verified as derived.** The id is a hash over gate, rule,
file and cause, so a finding either carries the hash of what it reports or it was
written by something that did not derive it. That makes the routing observable from the
result, which is what the constraint could not be before.

**It cannot be triggered by data, only by code, and that is the point.** A hand edited
`gate.yaml` marking an external finding approved is already dropped by the next run,
because `carryForward` reads this run's provenance and skips it. The guard exists for
the path that does not exist yet, which is what a constraint on a work package is, and
the tests reach it by constructing the checks directly rather than through a command.

**What WP4 still owes.** Routing its checks through `Run` so they are judged with the
others. A verdict assembled beside `Run` would now fail rather than pass, which turns
the constraint from a sentence somebody has to read into a run somebody cannot ignore.