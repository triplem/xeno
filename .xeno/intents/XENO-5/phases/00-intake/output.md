---
intent: github.com/triplem/xeno#5
phase: 00-intake
created: 2026-09-23T18:46:00Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 6756c73ba0804f3dd5e4302debb666acba0b2eb668b82a3d0d4441d7b5c7f77f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #5. Two of the six first steps of section 4 are not done, and they are not done
because they belong to no work package: the steps come before the packages, so nothing
in the plan ever picks them up.

## Scope

`CLAUDE.md` at the root, holding what step 3 names and nothing beyond it, and the state
of the protected branch recorded rather than assumed.

The file stays short, which is a requirement and not a preference: it is sent with every
request of every session for the life of the project. It is also not where learnings
accumulate. Section 10 sends those to the rule set through a merge request, and a file
that took them directly would make learning take effect unreviewed, which is the one
thing that section forbids.

## Non goals

`AGENTS.md`. The plan names both; one is wanted, and a second file with the same content
is a second thing to keep true.

The protected branch itself, which this host cannot give: rulesets and branch protection
need a paid plan for a private repository. What is in scope is writing that down.
