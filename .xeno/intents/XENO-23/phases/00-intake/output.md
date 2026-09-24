---
intent: github.com/triplem/xeno#23
phase: 00-intake
created: 2026-09-24T18:38:58Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 5818b27dd69d3a149f97bceeadb1178ced4a0591e234513605667310d312fafd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #23. `SUPPLY-CHAIN.md` states where its pins come from twice, and the two
statements disagree since the actions were upgraded.

## Scope

One statement of provenance, written so that it does not go stale the next time a pin
moves: the principle rather than a release number, with a pointer to where the evidence
actually lives, which is in every release.

The paragraph that narrates the upgrade goes. #13 removed that habit from the comments
and did not look in this file.

And the illustration of why a tag is not enough, which names a version this repository
no longer uses.

## Non goals

The pins themselves, which are current and correct.

The sentence naming v0.4.0 is not wrong about the two tools it still covers. It is being
replaced because a document that has to be re-read against reality after every change is
the thing being fixed, not because it lies.
