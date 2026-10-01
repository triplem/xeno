---
intent: github.com/triplem/xeno#160
phase: 00-intake
created: "2026-10-01T15:30:25Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e3d97f76568e930d23000290093b00c3983c05df2939993e2435c7a932f8d105
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
---

# Intake

<!-- xeno:section:problem -->
## Problem

#158 left a resolver and a gate over the tree. What it did not leave is any consequence for a
rule: `G-Policy` still stands as `not-implemented`, so a review rule produces nothing and a
checked rule is a line in `rules_hash` and no more. Section 9 says it in one sentence —
without the check that every entry was answered, "a review rule is a suggestion".

**The entries have nowhere to live.** `output.md` carries three structured lists in its
frontmatter, `open_questions`, `decisions` and `evidence`, each with a type in
`internal/model` and a gate that reads it. The checklist has neither. Section 9 fixes the
entry's shape — a rule id as the anchor, a result of `met`, `deviation` or `not-applicable`,
and a note required for the latter two — and nothing in the tree expresses it, so there is
nothing for a gate to count and nothing for an agent to write.

**Two different questions sit in one gate row.** Section 7 gives G-Policy "conformance against
the effective rule set, every review checklist entry answered". The second half is countable
today: the entries exist or they do not, they carry a result or they do not. The first half is
the predicates, and they are the next piece. A gate that implements only the half it can and
says nothing about the other would report green over a checked rule nobody evaluated, which is
precisely the class of claim section 16 exists to keep out of the record.

**A lens must not be able to move the count.** Section 12 is explicit: a checklist entry from a
lens carries `source: lens` and no rule id, "which is what keeps it apart from the entries
G-Policy counts", and a lens can neither add to nor subtract from that set. So completeness is
counted against the effective rule set and not against the entries present, and a gate written
the other way round — every entry has a rule, therefore the checklist is complete — would let
four lens entries stand in for a missing answer.

**What makes the omission the point rather than the result.** Section 9 says G-Policy checks
that every entry carries a result "and nothing more", and in the same paragraph says that
passing over a rule becomes a recorded deviation rather than an omission nobody sees. Those two
sentences only hold together if a missing entry is a finding: a gate that checked the entries
present and not the rules they were supposed to answer would be satisfied by a checklist with
nothing in it, which is the omission the paragraph says is now visible.

**And a rule nobody can evaluate is worse than no rule.** A `checked` rule whose predicate type
has no implementation resolves, enters `rules_hash`, applies to a phase, and is never evaluated.
After this piece there are two gates that touch rules and neither would have said so. A project
adopting such a rule today would reasonably believe its check was running.

<!-- xeno:section:scope -->
## Scope

**In scope.** The checklist entry as a type in `internal/model`, carried as a
`review_checklist` list in `output.md`'s frontmatter beside the three lists already there.
`G-Policy` implemented in `internal/gates`, from P0, over the effective set #158 resolves:
completeness against the review rules that apply to the phase, a result on every entry from the
three section 9 names, a note on `deviation` and `not-applicable`, lens entries outside the
count and inside the result requirement, and a checked rule whose predicate type is
unimplemented reported rather than passed over. Tests for each, and the register row for the
one decision the documents leave open.

**Out of scope, and each for its own reason.**

The predicate types. `section-implies-section` and the four that read a commit range are the
next piece, and this one is written so that the registry they will fill is the only thing that
has to change: a type with an implementation is evaluated, a type without one is reported.

Deriving the rendered section from the frontmatter. A3 says the rendered sections are derived
from the structured lists, and nothing in this repository derives anything — `open_questions`,
`decisions` and `evidence` are all written as prose by the agent beside their frontmatter. Doing
it for the checklist alone would make one list of four derived and leave A3 three quarters
untrue. It is a gap in its own right and gets its own issue.

A command for writing an entry. Section 9's checklist is filled by whoever answers the rules,
and every structured list in this repository is written by the harness into the frontmatter
today. `xeno section set` exists because WP11 requires a command behind every MCP operation, and
a `checklist answer` command belongs with that package rather than ahead of it.

The shipped set and the examples. Still content, still after the predicates, and now with a
second reason: the first review rule anybody writes will be judged by the gate this piece adds,
so the shipped set is where its findings get their first outside reader.

External gates. Unchanged by this and independent of it.

**One boundary worth naming.** G-Policy is `from P0`, and the checklist is a P5 artifact. So at
P0 to P4 this gate has exactly one thing to say — a checked rule applicable to that phase whose
type is unimplemented — and otherwise passes. That is not the gate being half built; it is what
section 7's column and section 9's placement mean together, and a reader of `gate.yaml` should
not take four quiet rows for four unrun ones.

<!-- xeno:section:context-rationale -->
## Why this context

Section 9's rule format and checklist paragraphs are the specification for the entry and for
what the gate may say about it, and the sentence that decides the design is the one pairing
"every entry carries a result, and nothing more" with "passing over one is now a recorded
deviation rather than an omission nobody sees". Section 7 is read for the row and its `From`
column. Section 12's lens paragraph is read for `source: lens` and for the clause that a lens
cannot add to or subtract from the counted set, which is the one external constraint on how
completeness is computed.

Section 16 is read for what a green gate does not mean, because this piece adds a case to that
list by refusing to: a checked rule nobody evaluated would belong there, and reporting it
instead keeps it out.

`internal/model/model.go` is read for `Output` and for how `open_questions`, `decisions` and
`evidence` are declared, since the checklist is the fourth of the same kind and A3 is the
assumption that put them in the frontmatter. `internal/gates/gates.go` is read for G-Questions,
which is the nearest gate in shape — it walks artifacts, collects what was raised, collects what
answers it, and reports the difference — and for the gate table's `from` column.

`internal/rules` is read as it now stands, because this piece is its second caller and the
argument for returning a set rather than a verdict was made on the strength of this gate
existing. Whether that shape actually serves it is a question this intent answers rather than
assumes.

`internal/template` and the shipped `review` template are read for the `review-checklist` section
id, so that the frontmatter key and the section id follow the convention the other three set:
hyphenated section, snake_case key.

Nothing outside the repository is needed. Both halves of the gate are defined by Xeno's own
specification, and the one thing an outside source could settle — what a lens actually writes —
is fixed by section 12 and not by any lens, since none exists.
