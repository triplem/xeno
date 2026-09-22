---
intent: github.com/triplem/xeno#3
phase: 00-intake
created: 2026-09-22T11:59:31Z
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 5a9c0a456dcbf627bd89f3eda5f3f0161a119057d183dcffc2d4ba9af2faf617
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
open_questions:
  - key: Q-1
    text: Where do the WP1 fixtures live, and in what form?
    options:
      - text: A testdata directory per package, read by the existing Go tests
        consequence: closest to the tests that exist; the corpus stays a Go concern and
          nothing outside the language can read it
        recommended: true
      - text: One corpus at the repository root, each case a directory with its inputs
          and its expected gate.yaml
        consequence: readable without Go and reusable by WP17, at the price of a runner
          that has to be driven from outside its own tests
      - text: Both, with the root corpus generated from the package fixtures
        consequence: one source, two readers, and a generator to keep correct
      - text: Something else
        free: true
  - key: Q-2
    text: How is the artifact schema version recorded, given that existing artifacts are
      never rewritten?
    options:
      - text: A schema_version field in the common field set, written from the runner
        consequence: explicit and readable, but it is a new field and therefore a change
          to the process definition first
      - text: Derived from runner_version, which every artifact already carries
        consequence: no new field, at the price of tying the schema to the tool that
          wrote it, which are not the same thing
      - text: Not recorded in v1, and the question is answered before the schema first
          changes
        consequence: honest while one schema exists, and the plan says this has to be
          settled before the first release, which has now happened
      - text: Something else
        free: true
---

# Intake

WP1, artifact schema and core gates. Most of the package was built by hand before M0;
this intent closes the distance between what stands and what the package asks for.

## Scope

The three open items of issue #3: a decision must never be carried forward for a
finding from an external gate, the fixture corpus the package owes and which is meant
to be written before the code that reads it, and a recorded schema version so that an
older artifact stays readable without being rewritten.

## Non goals

The template engine, the rule engine, the external gates themselves and the platform
matrix. This package owes the behaviour an external gate's finding must receive, not an
external gate to produce one.

## Open questions

Both are recorded above as structured entries. Q-1 decides where the fixtures live,
which WP17 inherits. Q-2 decides how a schema version is recorded, and it cannot be
answered inside this package alone: every option but one is a change to the process
definition, which is a person's decision and not the agent's.
