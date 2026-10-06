---
intent: github.com/triplem/xeno#257
phase: 00-intake
created: "2026-10-06T07:49:35Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4d35683286c9cd1595c29ec4840a84549e9de4c96a026f1546bda0911401b8f8
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

#117 now has nine intents. The record is a fixed cost — 17 required sections, 32 to 33 files,
1,350 to 1,499 lines — against changes spanning 13 to 440 lines. The spread on the record is 11%
and on the change a factor of thirty-four, so proportionality is not a ratio that a shortcut could
improve by a percentage: the numerator does not move.

The plan forbids anticipating a shortcut's shape and says the shape should follow from measurement.
The measurement now exists and points at one lever: which of the seventeen a small change may omit.

**Two things make that answerable now rather than after a design.** The mechanism is already data —
each section is a line in `template.yaml` carrying `required: true|false`, and each template's own
comment says a project "puts its own under `.xeno/config/templates/<id>/`, which beats this one for
that id alone". And `Missing` reports only sections marked required that carry nothing, so the flag
is the whole of the lever; nothing has to be deleted.

**What is not obvious, and is what this intent measured, is how little reads a section.** No Go
file outside the tests names any of the seventeen by id — the one apparent hit, `Scope` in
`internal/rules/rules.go`, is a rule's own scope field and not the `scope` section. Across the
shipped rules and the examples exactly one section id appears: `release-notes`, named by
`release-notes-are-filled` through the `section-non-empty` predicate. `sectionNonEmpty`'s own
comment says the rest out loud: "a required section carrying nothing is read only by the next-step
suggestion and by no gate (A73)."

So of seventeen required sections, **one is read by a gate and sixteen are read by a suggestion**.
G-Policy's checklist requirement is not a counter-example: it reads the `review_checklist`
frontmatter list, not the `review-checklist` section.

What is missing is not an argument about which to drop. It is a candidate expressed as data, that a
project can adopt by copying and that the next intent can be measured against — which is what the
plan means by letting the shape follow from measurement rather than from anticipation.

<!-- xeno:section:scope -->
## Scope

In scope is a candidate reduced section set under `examples/templates/`, six template directories,
each carrying `template.yaml` and `strings.en.yaml` — both, because `Load` treats a missing bundle
as an error and never falls back to another language, so a directory with only a `template.yaml` is
not adoptable.

In scope is the candidate keeping six of the seventeen required: `problem`,
`acceptance-criteria`, `changes`, `test-mapping`, `results` and `release-notes`. Eleven become
`required: false`, which leaves them defined and writable and stops `Missing` asking for them.

In scope is each file saying what it drops and why, and the two that may not be dropped saying what
reads them: `release-notes`, read by `release-notes-are-filled` through `section-non-empty`, the
only section any rule names; and the `review_checklist` frontmatter, read by G-Policy, which is
required whatever the `review-checklist` section's flag says.

In scope is `acceptance-criteria` and `test-mapping` staying required by decision rather than by a
reader. Nothing reads them today — G-Test's mapping half has none, which is #250 — and #258 has
since settled that section 9 will make a criterion identifiable and a gate will read the mapping.
Dropping them would foreclose a decision already taken.

In scope is a README under `examples/templates/` saying how to try it, that trying it is how the
question gets answered, and what to measure: the same figures #117 collects, against the same
seventeen-section baseline.

Out of scope is adopting it anywhere, including here. Copying it into `.xeno/config/templates/`
would change this repository's own process mid-session and make the next intent's figures
incomparable with the nine already recorded.

Out of scope is changing any shipped template. Nothing under `.xeno/plugin/templates/` is touched,
so no adopter's process changes and the plugin digest is unaffected.

Out of scope is a recommendation that the shortcut be taken. The candidate is the thing that makes
trying one cheap; the measurement that follows is what would decide, and the plan is explicit that
the shape should not be settled in advance of it.

Out of scope is a second language bundle. `strings.en.yaml` only, which the README says, because a
project on another language would have to write its own and the fixture should not pretend
otherwise.

No normative document is touched. Section 9 fixes what the sections are; the `required` flag is the
template layer's and the plan leaves the call to the project.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the six shipped templates, the resolver that lets a project override one, the gate
that holds the only mechanical section reader, and the two documents that say what the sections are
for.

`.xeno/plugin/templates/*/template.yaml` are read for the seventeen and their flags, and for the
comment each carries about being overridden per id. That comment is the whole mechanism and it is
quoted rather than recalled, because the fixture's claim is that a shortcut needs no new machinery.

`internal/template/template.go` is read for three things that shape the fixture. `Load` prefers
`.xeno/config/templates/<id>` over the plugin's and falls back per id, so a candidate can be
partial; a missing strings bundle is an error and never a language fall back, so each directory
needs `strings.en.yaml` as well; and `Missing` reports only required sections that are empty, which
is what makes `required: false` the lever rather than deletion.

`internal/gates/gates.go` is read for what reads a section, and it is where the measurement came
from. `sectionNonEmpty` is the one predicate that names a section, and its comment states the
general case: a required section carrying nothing is read by the next-step suggestion and by no
gate, citing A73. G-Policy is read too, to confirm the checklist requirement is over the
`review_checklist` frontmatter and not over the `review-checklist` section — otherwise the
candidate would be wrong about the one section it most needs to be right about.

`docs/implementation-plan.md` is read for why `acceptance-criteria` and `test-mapping` exist, which
is G-Test mapping onto them, and for the sentence that no shortcut is defined on purpose because
the shape should follow from measurement. Both bear on which sections the candidate keeps and on
whether a candidate is allowed at all.

`docs/process-definition.md` is read for section 9, so that the candidate changes a flag and not
what a section is, and `docs/assumptions.md` for A72 and A73 — the first is the precedent for
`examples/` over the shipped set, the second is the row `sectionNonEmpty` cites.

`examples/rules/documentation-follows-the-change.yaml` is read as the shape a fixture's header
takes in this repository: what it is, why it is not shipped enabled, and how to adopt it.
