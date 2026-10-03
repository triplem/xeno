---
intent: github.com/triplem/xeno#202
phase: 05-review
created: "2026-10-03T19:29:00Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 004be2171719ba338ab9a6791ef78e9acadf41665d9f8c284d3f7c17f71b6e4a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
  - rule: deviations-are-traceable
    result: met
    note: >-
      One, in the implementation phase's deviations section: section 11's "what is sealed is
      never rewritten" splits into modification and deletion, which have different readers, so it
      is two rows rather than the one the design assumed. A single row would have had to name one
      reader and would then have been half wrong, which is the error this file exists to find.
  - rule: interface-change-needs-a-migration-note
    result: not-applicable
    note: >-
      No interface changes. One new document and one row in ASSUMPTIONS.md; no code, no field, no
      gate, no rule, no command and no artifact shape, so there is nothing anybody could have
      depended on.
  - rule: new-dependency-needs-a-rationale
    result: met
    note: >-
      None added. The pass was a search and a reading, and the file names its own method so that
      it is repeatable without tooling; go.mod still carries go.yaml.in/yaml/v3 alone.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. The documents are untouched: this intent reads them and writes
a third file. No field, gate, tool or rule is added, so the second rule needs no
specification change ahead of it. The branch carries one intent, the commit references
#202, and the issue carries its package label.

The acceptance criteria. Every clause has a named reader or is listed as having none; the
list is in the repository; clauses addressed to a person are separated from clauses with
no reader; a reader that cannot fail is called out as one, in the `tool_version` finding;
all four counts are stated.

The non-goals held. Nothing was fixed, no gate was added, no clause was judged as
deserving a reader, and the conventions files are untouched.

The cautions #202 states. "A clause whose reader is a person is not unread" — those are a
counted kind, not a gap. "A reader that cannot fail is worse than one with none" — the
`plugin_version` case is the reason the file has four kinds rather than two, and section
12's triple is listed under that heading. "The plan is normative here too" — both
documents were extracted from.

Conventions. Prose at 88, tables exempt, SPDX on the file, and the commit subject a
Conventional Commit with `Closes #202` in the footer.

<!-- xeno:section:release-notes -->
## Release notes

`CLAUSE-READERS.md` lists every normative clause of the process definition and the
implementation plan against what reads it: thirty-four tool requirements with their gate,
refusal or workflow step, and eleven properties that hold because of how the code is
arranged, nine of them with nothing that would notice a change. Forty-eight clauses
addressed to a person and 126 explanatory sentences are counted and not listed.

Nothing is fixed by this change, which is the point of it. Three findings have no issue
yet and are named in the file: `XENO_PLUGIN_DATA` is specified and read by nothing,
G-Complete runs only at P5 so an intent that stops short is never checked for
completeness, and section 12's model-tool-version triple is self-reported with nothing
checking it is true.

`ASSUMPTIONS.md` gains A90 for choosing a dated survey over a coverage gate.

<!-- xeno:section:residual-risk -->
## Residual risk

The file will age and nothing will say so. A reader named by symbol is wrong the day the
symbol is renamed; the file dates its pass and names the commit, which bounds the damage
to a reader who checks the date. The alternative was the gate A90 rejects.

The placement of 219 sentences into four kinds is a judgement. The two counted kinds are
the soft boundary — a second pass would move a handful of sentences between "addressed to
a person" and "explanation" — and the two enumerated kinds were each checked
individually, so the figures that carry weight are 34 and 11.

A list of unguarded properties is a list of the places a change would be silent, which is
useful to whoever maintains the project and to anybody else who reads the repository. It
describes the absence of a check and no way to exploit one, and the repository is already
public.

The three findings stay open until they are issues. Until then this file is where they
live, and a finding recorded in a document nobody is obliged to read is weaker than an
issue — which is the next step and is not part of this intent.
