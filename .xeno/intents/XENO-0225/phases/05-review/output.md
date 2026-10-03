---
intent: github.com/triplem/xeno#172
phase: 05-review
created: "2026-10-03T10:40:29Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4153057120ef92ef429e5b32d645bef56f6933f8cbf55bd89ab56471eb571588
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
      One deviation, and it is not a departure: the awkward sentence a link with no component
      produces, recorded with the alternative it was chosen over — skipping such a link and leaving
      its document unchecked.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      Nothing outside this intent depends on what changed. A project whose profile has a correct
      link sees no difference; one whose link is mistyped sees a finding where it saw nothing, which
      is the point rather than a migration.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met**, with one deviation that is not a departure: the awkward
sentence a link with no component produces, which names the alternative it was chosen over.

**`interface-change-needs-a-migration-note` — met.** Nothing outside this intent depends on what
changed. A project whose profile has a correct link sees no difference; one whose profile has a
mistyped link sees a finding where it saw nothing, which is the point rather than a migration.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added.

**Beyond the three rules.**

*Was the criterion met as written?* Yes, quoted rather than reinterpreted: a finding against the
profile rather than a silently missing file. #171's gap is closed by #172, which is the process
working as designed and costing what the figures on #117 say it costs.

*Was the other half honestly refused?* The byte count needs one clause in section 5 and this intent
says so in its problem, its scope, its changes, its gaps and here. The sentence a person would add:
in the `context.lock.yaml` block, `files: [{ path: ..., sha256: ..., bytes: ... }]`, with a line
saying the size is recorded because the budget is judged against what the phase was given and not
against what the tree holds now.

*Did the first decision after M0 find a home?* Yes, in 02-design, which is where the closed register
says decisions go. It cost nothing because the decision was about this intent's own subject, and
02-design's learning says what the harder case will be: a decision about the project at large,
arriving while an intent about something else runs. The proposal there is to open the issue for the
piece it belongs to, which is what happened this morning with the shipped set and the placeholder.

*What should a reader not conclude?* That a profile's links are validated. A document's existence is
checked; the component is not, and whether the document documents the component is a question
section 5 forbids inferring.

<!-- xeno:section:release-notes -->
## Release notes

**A context profile's link is checked.** A declared link whose document is not in the tree is a
G-Schema finding naming the profile, the component the link claims and the path as written, with
the next step to correct the path or take the link out. Before this it was skipped in silence, and
a link is the only field in a profile that names a specific file: everything else is a pattern, and
a pattern that matches nothing is a state rather than a mistake.

It is reported at every phase, like the budget, and it does not block.

A link whose document exists is unchanged: it joins the information base, last, with its hash.

**This closes an acceptance criterion #171 did not meet**, which is why it is a second intent for
one condition.

**Not in this release.** The byte count a budget is judged against is still measured from the tree
when the check runs rather than recorded when the phase ran, so a phase inside its budget can be
over it later. Fixing that needs one clause in section 5's `context.lock.yaml` block, which is a
change to the specification and a person's commit; the review carries the wording.

Closes #172. Refs #171.

<!-- xeno:section:residual-risk -->
## Residual risk

**Half of a link is validated.** The document is checked for existence; the component is not,
because it is a prefix in the project's own vocabulary and may not be a directory. So a profile can
declare that a component nothing in the repository has is documented by a file that exists, and
nothing says a word. That is deliberate and it means the check catches the half of a mistyped link
that is unambiguous.

**A document's existence is all this proves.** It is not read, and whether it documents what the
component contains is a question section 5 forbids inferring. A project with a stale document that
still exists passes.

**The byte count is still the one number that moves after sealing**, and it is the reason the
profile adoption experiment is still one step away. One clause of section 5, and it is not mine.

**A link with no component reads badly.** "The link for a link names …". Nobody writes that profile
on purpose and the better message was worth less than the document going unchecked.

**No profile exists here**, so WP8's three mechanisms — the base, the budget, the links — have been
exercised by tests and by three demonstrations on copies. The experiment that was deferred is now
unblocked on one of its two conditions.

**And the cost of this intent is the record of what the sequence charges for a one-condition fix.**
One function, three tests, 55 lines, six phases. The figures go to #117 beside XENO-0224's floor
measurement, and the two together are the clearest statement this project has of what a small
change costs: not because the sections were padded, but because there are thirty-one of them.
