---
intent: github.com/triplem/xeno#242
phase: 05-review
created: "2026-10-05T12:19:40Z"
schema_version: "1.0"
runner_version: dev+8f4b75b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ca7055fea74e0baf3d7361afdcf6477fbf909f0dd3d845bd40e5b1b947419cd0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Two, both against #242 rather than against this intent''s own plan, and both declared in P0 before the code was written: no gate.yaml refusal, because exchange.go''s precedent says the decision commands do not do it and D-6 is about gate approve and gate override; and no --phase, because G-Policy''s guard judges the checklist on the last phase only. Three additions beyond the issue''s list are named in P3 too, each from reading the gate rather than the issue.'
      result: deviation
      rule: deviations-are-traceable
    - note: No interface changes. The command is new surface, review_checklist and its four keys are A68's and untouched, and no existing call, field, gate, rule or artifact shape changes. Nothing outside this intent depended on the field having no writer.
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: 'None added. go.mod is untouched and the writer imports only internal/model and internal/rules, both of which the gate that judges the field already uses, which is the point: the two sets are shared rather than restated so a refusal on the way in cannot disagree with a finding after the fact. A local contains was written and removed in favour of model.OneOf, since this repository imports slices nowhere.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No normative document is touched: sections 9 and 12 specify the
checklist and section 7 the gate that reads it, and this intent implements a writer for what
they already describe, so the code moves towards the specification and no specification commit
precedes it. No field, gate, tool or rule is added — `review_checklist` and its four keys are
A68's and unchanged, and the gate list is untouched. The branch carries one intent, the commit
references #242, and the issue carries `wp4`.

The acceptance criteria. Nine met, one met by this section existing: criterion 9 asked that this
intent's own checklist be written with the command, and the entries below were written with
`xeno review answer`, which is the only check that the writer works where it is meant to be
used. Criterion 7 is met by construction rather than by a test, which P4's test mapping says
plainly rather than counting it covered.

The non-goals held. No change to the checklist's shape, to G-Policy, or to any checklist that
exists. No lens entry and no `--source` flag. No reader. No `CLAUDE.md` statement about hand
edits, which is XENO-0246's learning and section 10's to route.

Two of #242's five bullets are not implemented, and that is the substance a reviewer should
check rather than the code. The issue asked for a `gate.yaml` refusal on the precedent of the
decision commands, which do not do it; D-6 is about `gate approve` and `gate override`, which
carry `against`. And it offered an optional `--phase` where the gate's own guard leaves one
phase. Both are argued in P2's alternatives and P3's deviations, and `review.go`'s comment
carries the first where the next reader will meet it.

What this phase got wrong is in P4's results with the figures: the test counts were taken from a
filtered `go test -run` and said twelve functions for a file holding eight. Corrected before P4
was judged, and its learning records the method rather than the number.

Three things are thin and named in P4's gaps rather than here: `cmdReviewAnswer` has no test,
A74's divergence has none, and nothing makes the writer compulsory.

<!-- xeno:section:release-notes -->
## Release notes

`xeno review answer` writes the P5 review checklist, the last artifact field in this
repository that no command produced.

    xeno review answer RULE --intent KEY --result met|deviation|not-applicable [--note TEXT]

It appends an entry to `review_checklist` in the review phase's artifact, or replaces the
entry for a rule already answered, keeping its place in the list. After each answer it says
which review rules of the effective set are still unanswered, and says so plainly when none
are, so a checklist can be finished without reading `phase finish`'s findings to discover
what is missing.

Four things it refuses, which are G-Policy's own checks moved from the verdict to the moment
of writing:

- a rule that is not a `review` rule of the effective set, naming what the set holds
- a `result` outside section 9's three
- a `deviation` or `not-applicable` with no note; `met` is the only result that needs none
- a `source`, which section 12 gives to a lens entry, and a lens writes its own

There is no `--phase`. G-Policy judges the checklist only on the last phase, so the command
resolves it rather than offering a choice of one.

Why it matters: the checklist is the one field a gate refuses the phase without, and until
now the only way to produce it was to edit the artifact by hand, inside `artifacts_hash`,
after the sections were written and before `phase finish` sealed the phase. Every P5 in this
repository went through that window. A wrong rule id, a bad result or a missing note was
found by the gate, after the fact, and repaired by another edit and another judgement.

Nothing existing changes. No checklist already written is touched, the field was already in
`frontmatterOrder`, and the 357 verdicts that stood before this intent stand after it. This
closes the family #188, #195, #208, #217 and #220 worked through: after it, no artifact
content in this repository is produced by hand except `tool_version`, which has a flag so
that it need not be.

<!-- xeno:section:residual-risk -->
## Residual risk

The writer is not compulsory and cannot be made so by this change. A hand edit still produces
a valid checklist, because the field is ordinary frontmatter and no gate can tell which wrote
it. What is removed is the necessity. The honest completion is the `CLAUDE.md` statement
XENO-0246's learning asked for, and section 10 routes that through review rather than letting
this commit apply it, so for now the ban on hand edits still has an exception nothing states.

A74's divergence is untested and the refusal can mislead. The writer compares a rule against
the set that resolves now; G-Policy judges against the set the artifact recorded in
`rules_hash`. Where the rule tree changes between `section set` and `review answer`, the writer
can accept a rule the gate will reject, or refuse one it would have accepted. The gate decides
either way, so nothing silent happens, but a person reading a refusal may take it for more than
advice. Low, and written into the file's comment where that reader will be.

`cmdReviewAnswer` has no test. The behaviour is one `Fprintf` and a call, and `cmd/xeno` has a
test file where the other writers' commands are reached, so this is thinness rather than a
boundary. It is the most reasonable thing in this change for a reviewer to ask for.

The trail now carries two provenances for one field. Eighteen checklists were written by hand
and the nineteenth by a command, and nothing marks which is which. Rewriting the eighteen is
foreclosed — it would change each `artifacts_hash` and stale each verdict — so this is permanent
and is the same shape as A94's sealed references: a difference a reader can only infer from the
dates.

What no writer can check is whether an answer is true. Every refusal is about form, and whether
`met` is the honest answer to a rule is what the checklist exists to record. A90's finding is
why the command does not try: a reader that cannot fail is worse than none.

What is not a risk: the 357 pre-existing verdicts, which `gate verify` confirms at exit 0, and
the existing commands, none of which changed. `--result`'s help text is the only shared surface
touched, and it gained a clause rather than a meaning.
