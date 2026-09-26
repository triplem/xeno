---
intent: github.com/triplem/xeno#64
phase: 00-intake
created: "2026-09-26T21:11:32Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+d9f41c4
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: dbc1091643b112c52b0b429c779817bcb069fe15ca9aa648188dd3e83c1d055e
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

G-Learning counted a record as filled when it carried any key beyond the header fields
every process file has. Section 10 defines more than that: `learnings` is a list whose
entries carry `category` from a closed set of four, `observation`, `proposal` and
`target`. XENO-3's first record invented `observations:` with a category of its own and
passed, which is how A7 was found rather than predicted.

The gap had grown quietly. Ten records across nine intents carried `category: process`,
a word section 10 does not have, and six of them were written today.

<!-- xeno:section:scope -->
## Scope

The shape, in G-Learning: the closed set for `category`, the four keys of an entry, an
unknown key named rather than ignored, and `no_finding: true` still passing as the
honest empty case. The ten records are reclassified and their phases re-sealed.

Not `target`. What that field may name was settled in #77 and the gate does not check it
here, which leaves four records naming a document rather than the file a project reads
before it acts. Checking it is a separate change and would flag those four.

<!-- xeno:section:context-rationale -->
## Why this context

**All ten readings came out the same, and that decided the question.** Every one of the
records was a statement about how this project works, which is what the twenty seven
records using `project-convention` say; none proposed a template, a prompt version or a
context selection rule. The alternative was widening the set to include `process`, and a
set with two members meaning the same thing drifts again as soon as somebody picks by
feel.

**Reclassifying was cheap here and would not be everywhere.** `learning.yaml` sits in
the phase directory, so `artifacts_hash` covers it and nine phases had to be sealed
again. That is safe in this repository because the whole trail carries no decision, so
no release and no `against` hash depends on those verdicts. Where a decision existed it
would have been lost, which is the difference between this and the rename in XENO-41.

**Grandfathering was the option not taken.** An exception for records written before the
gate would have no expiry and no reason: unlike `by-hand` in a `strings_hash`, whose
bundle is gone, a category outside the set is recoverable by reading the record. The
register has spent this week removing exceptions that outlived their reason.