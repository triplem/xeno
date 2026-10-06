---
intent: github.com/triplem/xeno#258
phase: 00-intake
created: "2026-10-06T15:07:41Z"
schema_version: "1.0"
runner_version: dev+90593d0
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5bcec7c7e0822e5fcd51379f077543e21df66e301844eef68f0de4da82a007d1
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

Two clauses of the specification ask for something the artifacts have no way to carry, and #258
decided both. Neither decision can become code until the specification enumerates the thing, because
the second standing rule makes an addition to what the artifacts carry a specification change, and
the first makes that commit a person's and makes it come before the code.

**The recommendation's reason has nowhere to live.** Section 8 asks for "the agent's recommendation
with a reason". `model.Option` carries `Text`, `Consequence`, `Recommended` and `Free`, so the reason
can only sit inside the question's text or inside a consequence, where it is indistinguishable from
what it shares the field with. `QuestionAsked` requires exactly one recommended option and nothing
requires a reason, because there is no field to require. #247 argued both homes without settling and
#258 settled it on `Option`, required where `recommended: true`.

**An acceptance criterion cannot be named, so the mapping cannot be judged complete.** Section 7 asks
G-Test for "declared test result successful, mapping of acceptance criteria complete". The first half
has a reader from #212 and the second has none, which makes it the only row in
`docs/clause-readers.md` reporting a gate that judges part of what it is asked for. A completeness
check needs criteria that can be named and a mapping that names them, and measured on 2026-10-06
over the trail before XENO-0260: 19 of 64 P1 artifacts number their criteria and 7 of 64 P4 mappings
cite a number. #258 settled the anchor on the template version, because all 64 P1 artifacts declare
`requirements@1.0.0` and `template` is already a frontmatter field read by `model.Template`.

## What stands in the way, precisely

Nothing in the specification enumerates a reason on an option, a numbered criterion, or a numbered
mapping. The documents are what the artifacts may carry, so each decision is a sentence away from
being buildable and no distance at all from being arguable.

This is the step the two earlier intents of this session stopped at. XENO-0260 corrected the section
number the convention belongs in — section 5, not section 9 as the issue had copied twice — and
recorded both decisions under the clause table with what each waits on. What it waits on is this.

## Why it is this commit and not the code

The first standing rule: a change to a normative document "is its own commit made before the code
that follows from it". So the two amendments are made here, alone, with nothing of the implementation
in the same commit or the same intent. `model.Option` gains no field in this branch, no template
version is bumped, and G-Test grows no check.

The rule also says the documents are not the agent's to edit. That is the one thing this intent does
anyway, and it does so because the person the rule defers to asked for it after being handed the
drafted wording twice and declining to write it themselves. The decision is recorded on #258 and in
this artifact rather than inferred, because an agent editing a normative document is the exception
the rule exists to make visible, and an exception nobody can find afterwards is the same as none.

<!-- xeno:section:scope -->
## Scope

In scope are two amendments to `docs/process-definition.md`, in two commits, because the first
standing rule makes a specification change its own commit and these are two changes.

In scope for section 8 is the sentence under "Open questions have to be resolved" that asks for "the
agent's recommendation with a reason". It gains the field's home and the reason for it: the reason
sits on the option the recommendation names, so that a recommendation moved to another option takes
its reason with it or the writer refuses.

In scope for section 5 is a subsection of its own, after "Rendering", saying that from
`requirements@1.1.0` the `acceptance-criteria` section is a numbered list and from
`verification@1.1.0` the `test-mapping` section names each criterion by its number, and that the
requirement is carried by the template version rather than by the schema so that nothing already
sealed is re-judged.

In scope is each amendment carrying its own reason in the document. A specification sentence that
says what without why is the thing #258's three clauses all turned out to be: readable, decided
somewhere else, and arguable again next year.

In scope is the wording being the wording drafted on #258. It was put there to be changed freely and
was not changed, which is recorded rather than passed over: the person had it in front of them and
asked for it as drafted, so the draft is the decision and not a starting point somebody has yet to
read.

Out of scope is `reason` on `model.Option`, `QuestionAsked` requiring it, `requirements@1.1.0` and
`verification@1.1.0` in the plugin, the completeness check in G-Test, and the tests for any of it.
All of it follows these commits and the first standing rule puts it after them, in its own intent.

Out of scope is `docs/clause-readers.md`. Its two rows still report no reader, and they are right
until the code exists: this intent makes the clause sayable, not read. XENO-0260 already recorded
under the table what each waits on, and the rows move when their readers do.

Out of scope is the sequencing convention, #258's third clause. #248 decided it stays in `CLAUDE.md`
permanently and needs no specification change, and #255 declined it a row for a reason this intent
does not reverse.

Out of scope is #235's advisory finding. It is the other decision of this session that needs a
section 5 commit, it belongs to a different issue and a different work package, and the third
standing rule gives it its own branch and its own intent.

Out of scope is making the sections say anything about `schema_version`. Neither amendment changes
the artifact schema's version: section 5 already reads `schema_version` and never enforces it
backwards, and the forward-only anchor here is the template version, which is a different number
doing a different job.

Out of scope is bumping the plugin's shipped template versions. The specification naming
`requirements@1.1.0` before one exists is deliberate and is what makes the anchor forward-only; the
template is created by the commit that implements the check.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the document being amended, the document that says where, and the three files that
decide what the amendments may and may not say.

`docs/process-definition.md` is read for the two places being changed and for the mechanisms the new
wording has to lean on rather than invent. Section 8's sentence under "Open questions have to be
resolved" is one of them. Section 5's Language subsection is the other, because it is where
`template.yaml` holds "stable section ids…, their order and the required fields" — which is why a
convention about what a section must contain belongs in section 5 and not in section 9, the Rule
model, where #258 had put it. It is also read for section 5's gate result, where the four check
results are enumerated, to confirm that neither amendment needs a fifth; for section 7's G-Test row,
which already asks for the mapping's completeness and therefore needs no change; and for section
11's sealing rule, which is the reason the criteria convention is anchored on a template version
rather than applied to the trail.

`docs/clause-readers.md` is read for the two rows this makes buildable and for what XENO-0260 wrote
under them. It is the measurement of which clauses have readers, and it is read here to confirm that
neither row changes yet: a clause that can be carried is not a clause that is read, and the table's
own definition of a reader is the thing that would fail if the clause were violated.

`docs/assumptions.md` is read for A90 — a reader that cannot fail is worse than none — which is why
the amendments are written so that each has exactly one reader when the code follows: `QuestionAsked`
for the reason, G-Test for the mapping. And for A4 and A42, which fixed the result set, to confirm
that neither amendment touches it.

`docs/implementation-plan.md` is read to confirm that neither clause belongs to a step that has to
move. It governs the order the runner is built in and not the content of a template section or the
fields of a question, so no step changes and nothing in the plan is a finding here.

`CLAUDE.md` is read for the three standing rules, which are the whole shape of this intent: the first
makes these commits a person's and makes them precede the code, the second is why they are needed at
all, and the third puts #235's commit on a different branch. It is also read for the convention on
replacing a paragraph changed a second time, which applies to section 8's sentence, and for the one
on verifying a negative, which is why "section 9 says nothing about acceptance criteria" was taken
from a search that returns five hits elsewhere in the same file rather than from recalling what
section 9 contains.

The figures were taken from the trail and not from the issue. 19 of 64 P1 artifacts numbering their
criteria and 7 of 64 P4 mappings citing a number were counted by extracting each artifact's section
between its anchor and the next; every P1 artifact declaring `requirements@1.0.0` was read the same
way, and that count is what makes the template version a usable anchor rather than a hope.
