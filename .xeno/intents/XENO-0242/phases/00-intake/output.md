---
intent: github.com/triplem/xeno#188
phase: 00-intake
created: "2026-10-03T21:15:12Z"
schema_version: "1.0"
runner_version: dev+efc44a9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4963a3601877c4ea1fde451f9cd88cd5eabdd3c473fba3f35a3aab488b0c5fd4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
open_questions:
  - key: Q-1
    text: >-
      #188 asks for three things: a figure that notices nothing was asked, the decision
      landing in the field rather than in prose, and strictness about proposed_by against
      decided_by. How much of that does this intent build?
    options:
      - text: >-
          The two writing commands and the figure. question record and decision record write
          the two frontmatter blocks the way assumption record writes the register, and intent
          status prints a line for an intent that reached P5 with no question and no decision.
        consequence: >-
          the fields become writable without a hand edit of a sealed artifact, the omission
          becomes a figure rather than a finding, and this intent's own decision is the first
          decisions entry in the trail; no gate and no document changes
        recommended: true
      - text: >-
          The figure only. open_questions and decisions stay hand written in the frontmatter,
          as review_checklist is today.
        consequence: >-
          the smallest change that makes the omission visible, and the loop stays carried by
          discipline; the field a gate reads stays one hand edit away from being wrong, which
          is the gap #208 files against evidence for the same reason
      - text: >-
          The commands, the figure, and strictness: G-Questions treats a decision whose
          decided_by equals its proposed_by as a finding.
        consequence: >-
          the trail cannot be satisfied by the agent naming itself as the decider, but no
          section says that today, so the first standing rule makes it a document change in
          its own commit before any code and this intent would wait on a commit it may not
          make
      - text: Something else, entered by the person deciding
        free: true
---

# Intake

<!-- xeno:section:problem -->
## Problem

Section 5 carries open questions and decisions structurally, and section 8 gives a
question three exits through a person. The shapes are built, G-Questions is built, and
across the whole trail the loop has never once run end to end. Four intents raised a
structured `open_questions` block, all of them in P0 and all of them early. No intent
has ever written a `decisions` block, and `decided_by` does not occur anywhere under
`.xeno/`.

**The gate cannot catch it, by construction.** G-Questions verifies that a question
which *was* raised reaches one of its three exits. Nothing declares how many
there should have been, so an intent that asks nothing passes it, and the greenest
possible trail is the one that asked nothing. Every other way this process fails is
an omission a gate names: a missing `learning.yaml`, an absent required field, an
assumption without an origin. Here the thing omitted is the raising of the question,
and no gate can see an absence nothing counts.

**And the field has no writer.** `assumption record` writes the register, `learning
record` writes the learning file, and `gate approve` writes a decision on a finding
with the person on it. `open_questions` and `decisions` are frontmatter blocks
nothing writes, so a decision taken in conversation lands wherever prose is cheapest:
a scope section, a design section called decisions, an `ASSUMPTIONS.md` row, a commit
message. Findable by a reader, invisible to a query, and with nobody's name in the
field built to carry it. The three intents run before this one each had a real decision
in them, two of them were put to the maintainer and answered, and all of it is prose.

**The two halves are one problem.** A figure that reports how little was asked is
worth having only where recording the answer is a command rather than a hand edit of
a sealed artifact's frontmatter, and a command is worth having only where somebody
notices it was not used.

<!-- xeno:section:scope -->
## Scope

**In scope.** Two writing commands and one figure.

`xeno question record` writes an `open_questions` entry into the phase artifact's
frontmatter, with its key, its text, its two to four options and their consequences,
the recommendation and the free entry. `xeno decision record` writes a `decisions`
entry, with `resolves`, `chosen`, `rationale`, `proposed_by` and `decided_by`. Both
refuse what the gates already refuse, in the shape `assumption confirm` refuses:
a decision without a person is not a decision, and a question with one invented
option is not a question.

`xeno intent status` gains one line for an intent that has reached P5 with no question
and no decision in any phase. A line, not a finding: a phase with nothing to ask is
the ordinary case and most phases are that. What is being reported is the figure
across an intent, which is what #117 does for proportionality and what would have
surfaced this years earlier.

**Out of scope, each for its own reason.**

Strictness about `proposed_by` against `decided_by`. It was Q-1's third option and
it was not chosen. No section says that a decision whose proposer is its decider is
a finding, so by the first standing rule it is a document change in its own commit
before any code, and that commit is the maintainer's. The fields are written here;
what reads them is a later argument.

A quota, or any finding for a phase that asked nothing. The issue says it in as many
words: a phase with no open question is legitimate. A gate that demanded one would
be answered with invented questions, which is the failure section 8's third exit
exists to prevent.

Rewriting the sealed phases. XENO-0227 to XENO-0229 are inside `artifacts_hash`
with verdicts written, and a retrofitted question would be a record of an exchange
that did not happen in that phase. D-1 to D-7 stay in `ASSUMPTIONS.md` for the reason
that file gives.

The other two frontmatter blocks nothing writes. `evidence` has the same shape of
gap and is #208 in WP6; `review_checklist` is written by hand into P5's frontmatter
and is WP7's. Three writers at once would be one command generalised past what any
of the three asks for.

Section 5 and section 8. Both already say what is needed, which the issue checked
before filing. Nothing here is a document change.

**What this intent owes its own trail.** Q-1 is raised in this phase's frontmatter
and settled by a `decisions` entry in P1 carrying `decided_by`. The done-when of
#188 is that the count of `decided_by` in the trail is not zero, so the first thing
the commands have to be able to record is the decision that scoped them.

<!-- xeno:section:context-rationale -->
## Why this context

#188 is read in full, because it is the measurement and this intent is only its
consequence.  Its counts were recomputed rather than trusted: four intents with an
`open_questions` block, zero `decisions` blocks, zero occurrences of `decided_by`
under `.xeno/`.

Section 8's open questions subsection is read as it stands: a question asked with
options and never open, two to four of them with a consequence each, the recommendation,
the free entry as the normal case for a domain expert rather than an escape hatch,
the three exits and why the third is not optional. It is the specification for what
`question record` and `decision record` are allowed to write, down to the counts
the gate already enforces.

Section 8's neighbouring subsection, that decisions are not assumptions, is read for
where a decision lives: in the artifact of the phase that settled it, sealed with it,
and not at intent level, because only an assumption has a state that changes. That
is why these two commands write frontmatter and the register's commands write a file.

`internal/model/model.go` is read for `Question`, `Option`, `Decision` and `Output`,
which are the shapes being written and are not changing. `Decisions` is `omitempty`,
which is the half of the issue's diagnosis that the code carries: prose satisfies
the template while the field stays empty and nothing is violated.

`internal/gates/gates.go` is read for `questionShape`, `decisionShape` and `questions`,
because a command that writes what a gate would reject is worse than no command. The
shape checks are where the refusals come from.

`internal/runner/runner.go` is read for `SectionSet`, which is the one writer of
`output.md` and shows how frontmatter is carried over and re-rendered, and for
`RecordAssumption` and `DecideAssumption`, which are the pattern a writing command
follows in this tree. Also for `Intents`, `summarise` and `IntentSummary`, where
the figure has to be computed, and for `Status`, which already walks the phases the
figure has to read.

`cmd/xeno/main.go` is read for `cmdIntentList` and `listRow`, because the figure is
a line in that output and the column widths are a decision #136 already took.

`docs/v2-delta.md` is read for the one sentence that fixes what the two person fields
mean: `proposed_by` carries the agent identity and the person's field carries the
person. It is the convention this trail's first `decisions` entry will set.

Nothing outside the repository is needed.
