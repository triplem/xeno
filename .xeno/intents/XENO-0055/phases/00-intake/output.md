---
intent: github.com/triplem/xeno#55
phase: 00-intake
created: "2026-09-26T12:40:09Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2cb8638
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 4fcea10f04e68b1e3c5b4933dd040b0318267c8b8dbcf3e2a71387021853af40
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

`model.Assumption` and the record section 8 of the process definition defines were not
the same thing. The code had `text` where the specification has `assumption`, and no
`origin`, `confidence` or `status` at all. G-Assumptions therefore read the right
question off the wrong field: it inferred openness from an empty `confirmed_by`, which
gives the same answer in the two ordinary cases and the wrong one in the third. A
rejected assumption has no `confirmed_by` either, so an assumption examined and dropped
was reported as unconfirmed, and there was no way to record the dropping at all. Section
8 insists that third exit exists, because without it people invent assumptions to get a
gate green.

<!-- xeno:section:scope -->
## Scope

The register's record, the gate that judges it, and commands to write one. `assumption`,
`origin`, `confidence` and `status` under the names section 8 uses; G-Assumptions red on
`status: open` and not on an empty `confirmed_by`; `assumption record`, `assumption
confirm` and `assumption reject`, so that WP11's `record assumption` operation has a
command behind it as every operation must.

Question resolution stays as it is. G-Questions asks for `status: confirmed` where it
asked for a non-empty `confirmed_by`, which is the same rule read off the right field,
and nothing else about WP5 is touched here.

<!-- xeno:section:context-rationale -->
## Why this context

**The specification wins, so this is a code change.** Section 8 fixes the shape and says
what the gate does with it; the code disagreed with both. Nothing in the documents was
edited for this, and the divergence is not a reading either way: the field names are
written out in the section.

**Nothing migrates, and that was checked rather than assumed.** Every `assumptions.yaml`
in this repository is `assumptions: []`, so no record has to be rewritten, and the
register of this intent is the first that could carry one. The single fixture in the
runner tests that used the old shape is rewritten, which is how the tests now assert the
specification's shape rather than the one the code had.

**Three things the section does not settle, recorded rather than absorbed.** `resolves`
is on the record and not in the schema, because resolution by confirmed assumption needs
the question's key to be checkable; an empty `status` is read as open, so a register
written before the field existed stays readable; and a rejection records the status
without naming who made it, because section 8 has `confirmed_by` and no counterpart. The
third is a finding about the section rather than a choice: a rejection is a statement by
a person and the register cannot name them. All three are in `ASSUMPTIONS.md`, A47 to
A50.