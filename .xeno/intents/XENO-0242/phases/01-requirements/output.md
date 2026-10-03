---
intent: github.com/triplem/xeno#188
phase: 01-requirements
created: "2026-10-03T21:17:41Z"
schema_version: "1.0"
runner_version: dev+efc44a9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 686062591ec40cb405416cf8ec3a2fd36b2f34d6c6cd75b5ea79d17b03750e94
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
  - id: D-1
    resolves: Q-1
    chosen: >-
      The two writing commands and the figure. question record and decision record write the
      two frontmatter blocks the way assumption record writes the register, and intent status
      prints a line for an intent that reached P5 with no question and no decision.
    rationale: >-
      The two halves answer each other. A figure that reports how little was asked is worth
      having only where recording the answer is a command rather than a hand edit of a sealed
      artifact, and a command is worth having only where somebody notices it was not used. The
      figure alone leaves the field a gate reads one hand edit away from being wrong, which is
      the gap #208 files against evidence for the same reason. Strictness about the two person
      fields is a rule no section states, so the first standing rule makes it a commit the
      maintainer has not made, and this intent would then wait on a document rather than close
      a loop that has been open since the shapes were built.
    decided_by: triplem
    proposed_by: claude-opus-5
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**A question is recorded by a command, into the frontmatter of the running phase.**
`xeno question record --intent K --phase NN` reads the entry on stdin or from
`--file` and writes it into `open_questions` of that phase's `output.md`, in the
position section 5 gives the field.  The body is not touched, no other frontmatter
field changes, and a second question appends rather than replaces.

**It refuses, before writing, everything `questionShape` would find afterwards.**
A question with no key; a key already used in that intent, in any phase; fewer than
two or more than four options that are not the free entry; anything other than exactly
one free entry. `no_options: true` is the exception the section names, and it passes
with no options at all. A refusal says which of these it is and writes nothing.

**A decision is recorded by a command, with the person on it.** `xeno decision record
--intent K --phase NN --id D-1 --chosen T --rationale T --by WHO [--resolves Q-1]
[--proposed-by WHO]` writes a `decisions` entry into the same frontmatter. `--withdraw`
is section 8's third exit: a withdrawal carries its rationale and its person and no
chosen option.

**Neither command invents a person.** `--by` missing is a refusal in the register's
words, that a decision is a statement by a person. Nothing defaults it, to the git
user or to anything else.

**A write after a verdict invalidates the verdict, as a section write does.**
Neither command refuses a judged phase and neither rewrites `gate.yaml`. The next
`phase finish` says the phase changed after its verdict was written, which is the
mechanism that already exists and the one a reader of #216 expects.

**A question raised in one phase and decided in a later one is resolved for
G-Questions.** Shown by this intent's own trail: Q-1 is in P0's frontmatter, D-1 is
in P1's with `resolves: Q-1` and `decided_by`, and P5's G-Questions passes with a
question in it rather than with none.

**`xeno intent status` says when an intent reached P5 having asked nothing.** For
an intent whose work has reached 05-review with no `open_questions` entry and no
`decisions` entry in any phase, the listing marks the row and the one-intent form
prints a line after the table. The wording names both halves, because a trail with
a decision and no question is a different state from a trail with neither.

**It is a figure and not a finding.** No gate changes, nothing goes red, and the
line appears on an intent that is otherwise green. A phase with nothing to ask is
the ordinary case; what is being reported is an intent reaching the merge having
asked nothing at all.

**The line appears on no other intent.** Not on one that raised a question, not on
one that recorded a decision, and not on one that has not reached P5, where an open
question is still normal and G-Questions does not run either.

**The figure is computed once.** Both forms read the same value from the same
walk over the phases, because two counts of one thing drift, which is the argument
`Decided` already settled in this tree.

**Nothing else changes.** No field of section 5, no gate, no rule, no template,
no document, and no existing artifact. `./xeno gate verify` stays at exit 0 over
every verdict in the repository.

**The usual gates of this repository.** `gofmt`, `go vet`, the suite green, 88
columns in prose, SPDX on new files, and a test for each refusal and for each of
the figure's three answers.

<!-- xeno:section:non-goals -->
## Non goals

**No strictness about `proposed_by` against `decided_by`.** It was Q-1's third option
and the maintainer did not choose it. The fields are written; a gate that compared
them would be a rule no section states, and the first standing rule makes that the
maintainer's commit, made before any code.

**No quota, and no finding for a phase that asked nothing.** #188 says it plainly
and section 8 gives the reason: a gate that demanded a question would be answered
with invented ones, which is the failure the third exit exists to prevent.

**No change to G-Questions.** It is correct about what it checks. That it is vacuous
where nothing was asked is not a defect in the gate, it is the reason the figure is
a figure.

**No writer for `evidence` or `review_checklist`.** Both are frontmatter blocks nothing
writes and both have the same shape of gap. The first is #208 in WP6 and the second is
WP7's. One command generalised over three blocks would be designed against none of them.

**No intent-level file for questions.** Section 8 settles this: only an assumption
has a state that changes, so only the register lives at intent level. A question is
sealed with the phase that raised it and a decision with the phase that settled it.

**No backfill.** XENO-0227 to XENO-0229 are sealed with their verdicts written and
two of them merged, and D-1 to D-7 stay in `ASSUMPTIONS.md` for the reason that
file gives. A retrofitted question is a record of an exchange that did not happen
in that phase.

**No MCP surface.** Section 7's tool surface is a budget and this is not the intent
that spends it.

**No change to how a verdict is invalidated.** Recording a question or a decision after
a verdict behaves exactly as a section write does today. Making these two commands
refuse where `section set` does not would be a second rule about the same situation.

<!-- xeno:section:constraints -->
## Constraints

**Q-1 is the scope, and the maintainer chose it.** The two commands and the figure,
no gate change and no document change. The two options not taken are in the non
goals with the reason each was not taken.

**`SectionSet` is the only writer of `output.md`, and it carries frontmatter over.**
It reads the file, unmarshals the frontmatter into a map, re-renders the body from
the template and writes the whole file again. A command that writes one frontmatter
key has to leave that path exactly as capable as it is now, because the next section
write of the same phase re-marshals whatever these commands put there.

**`frontmatterOrder` already lists both keys.** `open_questions` and `decisions`
sit between `rules_hash` and `evidence` in the order section 5 lists them, so the
position is settled and nothing here decides it.

**The gates are the specification for the refusals.** `questionShape`, `decisionShape`
and `questions` already say what a well formed entry is. A command that could write
what a gate rejects would be worse than the hand edit it replaces, so the refusals
are those checks moved in front of the write rather than a second opinion about shape.

**The person is never defaulted.** A35's rule, which three intents in a row have
now restated: a plausible value in a field nobody produced is worse than an absent
one. For `decided_by` the stakes are the whole register, since the question it exists
to answer is who decided.

**A question is nested and a decision is flat.** Up to five options with a consequence,
a recommendation and a free entry do not fit flags without inventing a separator; six
scalars fit flags exactly, and `--by` is already this tool's word for the person. Each
command takes the channel its shape asks for, which is why `section set` reads stdin
and `assumption record` does not.

**The figure belongs where the state is already computed.** `summarise` reads
`intent.yaml` and the phase verdicts for every row of the listing and is the one
definition of an intent's state in this tree. The counts are read on the same walk,
and `cmdIntentList`'s column widths stay as #136 set them.

**The figure has to be cheap.** The listing reads every intent in the repository,
so counting questions means opening the phase artifacts of each one. What it opens
is what `Status` already opens, and nothing is added per phase beyond reading the
frontmatter that is read anyway.

**88 columns in prose, SPDX on new files, `gofmt`, `go vet`, the suite, and `./xeno
gate verify` at exit 0.**

**One intent, one branch, one issue.** `188-the-decision-lands-in-the-field`, #188,
labelled wp5, and no commit before the code because this intent changes no document.
