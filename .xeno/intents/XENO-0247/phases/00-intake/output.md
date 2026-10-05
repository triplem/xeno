---
intent: github.com/triplem/xeno#242
phase: 00-intake
created: "2026-10-05T12:10:08Z"
schema_version: "1.0"
runner_version: dev+3f5ad34.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8707d01c613b6d3a6849d419382a2aa21cae8f1a47187fe54b4ffb5e8bdb157e
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

`review_checklist` is a list in the P5 artifact's frontmatter, and no command writes it.

G-Policy refuses the phase without it. A69 requires every `review` rule of the effective
set to be answered whatever its `applies_to`, and `internal/gates/gates.go` says it in as
many words: "answer it in review_checklist with met, deviation or not-applicable". The gate
counts from the rules to the entries, so an empty checklist is a finding per unanswered
rule rather than a pass.

Every other artifact field has a writer. `open_questions` has `question record`,
`decisions` has `decision record`, `evidence` has `evidence declare`, the sections have
`section set`, the register has the three `assumption` commands, the learning record got
one in #195 and the context scope in #217. `review_checklist` appears in
`internal/runner/runner.go` exactly once, in `frontmatterOrder`, which means the runner
preserves it across a `section set` and never produces it.

So the one field a gate refuses the phase without is the one field written by editing the
artifact by hand. That is the practice this project forbids elsewhere, and the edit lands
inside `artifacts_hash`, in a window that opens when the sections are written and closes
when `phase finish` seals the phase.

It is also the only artifact content nothing validates on the way in. G-Policy catches an
omission, a bad `result` and a rule outside the set, but only after the fact and only when
the phase is judged. Every P5 in this repository went through that window, and the reason
none broke is that a person was careful each time.

<!-- xeno:section:scope -->
## Scope

In scope is one writer, `xeno review answer RULE --result R --note TEXT`, which amends the
`review_checklist` block of the review phase and nothing else, through the `artifact` and
`amendFront` helpers `question record` and `decision record` already use.

In scope are the checks a writer can make that a hand edit cannot: a `result` outside
section 9's three is refused, a `deviation` or `not-applicable` without a note is refused,
and a rule that is not a `review` rule of the effective set is refused by id. These are
G-Policy's own three, moved to the moment of writing, where the answer is a refusal with a
reason rather than a red verdict on a sealed phase.

In scope is replacing the entry for a rule already answered rather than appending a second,
which is the maintainer's call on the one question #242 left to this package. It matches
`section set`, which overwrites a section rather than accumulating versions.

In scope is reporting which `review` rules of the set are still unanswered, printed by the
command after each write, so the checklist can be finished without reading the gate's
findings to discover what is missing.

Out of scope is the checklist's shape. A68 fixes `rule`, `result`, `note` and `source`, and
nothing here changes them. This is a writer for what is already specified, which is the
second standing rule: no invented fields.

Out of scope is a lens entry. Section 12 gives it `source: lens` and no rule id, and this
command takes a rule as its one positional argument, so it cannot express one. A lens writes
its own entries and G-Policy already keys on the missing id rather than on the source.

Out of scope is any change to G-Policy. The gate keeps being the authority on a sealed
phase; the writer only makes it harder to reach it with something the gate will reject.

No normative document is touched. Sections 9 and 12 specify the checklist and this intent
implements a writer for it, so the code moves towards the specification and no specification
commit precedes it.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the specification of the field, the gate that reads it, and the three writers
this one is modelled on.

Both normative documents are read rather than assumed, because this intent's scope claims
neither needs changing and the last intent made that claim and was half wrong. It holds
here: section 9 specifies the checklist and section 7 the gate, and both describe a field
that exists and is judged. What is missing is a writer, which is code, so the specification
is ahead of the tree rather than behind it.

`internal/runner/exchange.go` is the closest precedent and is read for its reasoning as
much as its code. It explains why a frontmatter writer amends one field instead of going
through `SectionSet`, which would re-render the body from the template, and why neither
question nor decision looks at `gate.yaml`. That second point is the one this intent had
expected to copy from `gate approve` instead, and reading the file is what settled which
precedent applies.

`internal/runner/scope.go` is read for the shape of a writer added to an artifact that had
none, which is this intent's own shape, and for how it reports a figure the caller needs
without recording it.

`internal/gates/gates.go` is read for what the writer must agree with. The checks are in
`reviewChecklist`, and the guard above it is what settles whether the command takes a phase
argument: the gate judges the checklist only on the last phase, so the field belongs to one
phase and a flag offering a choice would offer one the format does not have.

`internal/rules/rules.go` and `internal/model/model.go` are read for the two sets the writer
compares against: `Effective` filtered to `Review`, and `ChecklistResults` with
`ChecklistNeedsNote`. Both are already exported and already the gate's authority, so the
writer shares them rather than restating them, which is what keeps a refusal on the way in
and a finding after the fact from disagreeing.
