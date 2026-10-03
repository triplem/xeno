---
intent: github.com/triplem/xeno#1
phase: 01-requirements
created: "2026-10-03T10:19:17Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 56f1b6d5f120ad165a8ff13aaa4ead65703bdd2170c02c72345e6a3b2f5cbdd1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**The evidence is written down, checked rather than asserted.** The branch protection as the host
reports it: the required check, and whether administrators are included. The job that runs
`gate verify` and what its verdict covers. The count of intents that ran all six phases, the count
of verdicts, and how many gates report `not-implemented` against the plan's expectation of six.

**`ASSUMPTIONS.md` is closed rather than deleted.** Its header says the register stops at M0, names
the intent that closed it and the date, and says where a decision and a learning go instead. Every
row stays as it is, with its state column, because the file remains the record of what was decided
while the core was built by hand.

**Nothing is added to the register in this intent or after it.** A77 is the last row. This intent
records its own decisions in its phase artifacts, which is the loop the plan says replaces the
file.

**Issue #1 is closed with what was checked**, not with a sentence saying it is done: the settings,
the counts, the two readings of M0 and which of them each piece of evidence answers.

**What M0 triggers and this intent does not do is named.** The `Xeno-Intent:` trailer, and the two
dated paragraphs of the plan. A closure that left either unsaid would be the same kind of silence
it is correcting.

**No document under `docs/` is touched**, and the corrections owed there are named for a person.

**Nothing else moves.** No gate, no rule, no artifact schema. `./xeno gate verify` stays at exit 0,
the suite stays green, and the trail behind this intent is untouched.

<!-- xeno:section:non-goals -->
## Non goals

**No trailer adoption.** The plan says this repository carries `Xeno-Intent:` once M0 exists, which
is after this intent rather than during it. Copying the example into `given/project/` changes how
every later commit and every squash message is written, and that belongs to its own intent where
the convention can be stated and tested.

**No edit to the plan.** Two paragraphs are dated and `docs/` is beyond the agent. They are named
in the closure and owed to a person.

**No claim about M1.** Its ingredients are built and one clause of WP8's done-when is unverifiable
by construction. A milestone's reading is a person's.

**No deletion from the register.** Closing it is not pruning it. A reader of A6 or A62 learns what
was decided and when, and a row removed because its subject moved on would be the rewriting this
process avoids everywhere else.

**No new mechanism to watch milestones.** The learning says a milestone should be closed by the
piece that meets its conditions. Automating that is a proposal for the plan, not a thing to build
in the intent that noticed it.

<!-- xeno:section:constraints -->
## Constraints

**The plan's own definition decides what counts.** The milestone table's row and section 4's
sentence about the hand held record are the conditions; `M0.md` is what establishes that the name
means two things. This intent closes against those words and not against a sense of being far
enough along.

**The one fact outside the tree is read from the host.** Whether the pipeline is required and
whether administrators are included is a setting, not a claim, and `xeno enforcement check` exists
for exactly that question.

**`docs/` is beyond the agent**, so every correction owed there leaves as a finding.

**The register is closed in its own header**, because that is what a reader meets before the rows,
and a note at the end of a long table is a note nobody reads.

**This intent is small and its artifacts should say so.** The process does not have a light path
and the figures for that are on #117; what it does have is the honest option of short sections. A
closure padded to look like a work package would be worse than the overhead it was padding.

**One intent, one branch, one issue.** `1-m0-is-closed`, #1, which is M0's own issue, off a `main`
that carries WP4 and the agent layer.
