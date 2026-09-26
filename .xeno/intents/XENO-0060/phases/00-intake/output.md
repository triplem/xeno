---
intent: github.com/triplem/xeno#60
phase: 00-intake
created: "2026-09-26T12:49:59Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+374fea6
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 3f713178f5a8c72c999fdea73b4b1ec46f25d6da4277da7fc56c3b958bab1528
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

Section 8 gave the assumption register `confirmed_by` and no counterpart. The third exit
could be recorded, `status: rejected`, but the person taking it could not, and
`confirmed_by` is not the place for them: a rejected assumption was not confirmed by
anybody. The register could say that an assumption had been dropped and not who dropped
it, while the same trail names the person behind every decision on a finding.

<!-- xeno:section:scope -->
## Scope

`rejected_by` in section 8 beside `confirmed_by`, with exactly one of the two present,
and the code that follows: the field on the record, `assumption reject` writing the
person into it, and the gate asking both decided states for their person rather than
only confirmation.

No reason field. A decision on a finding carries one and an assumption does not, in
either state, and adding one here would be a second change wearing the first one's
clothes.

<!-- xeno:section:context-rationale -->
## Why this context

**The specification changed first, in its own commit.** The asymmetry was a finding
about section 8 rather than something to work around, which is why #55 recorded it as
A49 and left it rather than inventing the field. The order here is the one the standing
rule fixes: the section, then the code that follows from it.

**What the field is for is not the gate.** Both decided states pass, so nothing
downstream reads the name. It is read by somebody asking months later who dropped an
assumption, which is the question the trail exists to answer without asking anybody to
remember. That is also why the commit message was not good enough as the record: a
squash rewrites the commits and leaves the trees alone, so what is in the register
survives and what is only in a message may not.

**The gate reads the field that belongs to the status.** `DecidedBy` returns
`rejected_by` where the status is rejected and `confirmed_by` otherwise, so a record
with the wrong field filled in reads as decided by nobody and is a finding, rather than
passing because some name was present somewhere.