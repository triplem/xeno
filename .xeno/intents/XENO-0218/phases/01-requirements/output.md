---
intent: github.com/triplem/xeno#160
phase: 01-requirements
created: "2026-10-01T15:31:52Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4c277a7a93cc20ccaa8840b8de2c1db8510caf6ba7dd8e57874041ad9b497875
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Every criterion is an artifact, a rule tree, and a verdict over the pair.

**A repository with no rule tree is green at every phase.** No rules, no entries required, no
findings. This is the state of every repository today and of this one.

**A review rule that applies to P5 and has no entry is red.** The finding names the P5
artifact, the rule, and that an answer is owed. This is the criterion the whole gate rests on:
without it a checklist with nothing in it passes.

**An entry with no result is red**, naming the rule it answers.

**An entry whose result is outside the three is red.** `met`, `deviation`, `not-applicable`,
and nothing else; the finding names the value it found.

**A `deviation` or a `not-applicable` with no note is red.** A `met` needs none, which is the
asymmetry section 9 writes: the note exists to record why a rule was passed over.

**A lens entry neither completes nor breaks the count.** An entry with `source: lens` and no
rule id leaves a missing answer missing, so a checklist of four lens entries and no rule entry
is as red as an empty one. And it carries a result like any other entry, because section 9
requires that of every entry; a lens entry with no result is red.

**An entry naming a rule that is not in the effective set is red.** It asserts an answer to a
rule that does not apply, and the finding says so rather than ignoring the entry.

**A checked rule applying to the phase whose predicate type has no implementation is red**, at
whichever phase it applies to, naming the rule and the type and saying what to do: write it as
`review` until the type ships, or take it out. No predicate is evaluated by this piece, so this
is the only thing it says about a checked rule — and it says it rather than passing.

**G-Policy applies from P0 and is quiet where it has nothing to say.** At P0 to P4 with no
checked rules it passes, and the checklist is judged at P5 where the artifact carrying it
exists. A phase that is quiet because nothing applied is reported as a pass and not as
`not-implemented`, which is the distinction the row in `gate.yaml` now carries.

**The tree is resolved once per gate run.** G-Policy calls the same `rules.Load` and
`rules.Effective` as G-Rules and neither re-implements resolution, which is what section 7's
note about the two gates sitting together requires in practice.

**`G-Policy` is no longer `not-implemented`**, and the five rows a phase reports as unrun
become four — three after this.

**Nothing else moves.** `rules_hash` is untouched, because none of this changes the set. No
sealed artifact's `artifacts_hash` changes. `./xeno gate verify` stays at exit 0 over the whole
trail, and the whole suite stays green.

<!-- xeno:section:non-goals -->
## Non goals

**No predicate is evaluated.** The registry this piece consults is empty, and that is the
point: every checked rule reports its type as unimplemented until the next piece fills it.

**No rendered section is derived.** The entries live in the frontmatter and the
`review-checklist` section stays prose the agent writes, exactly as `open-questions`,
`decisions` and the evidence section are today. A3's claim that the sections are derived is
unimplemented for all four, and fixing one of four would leave the assumption more wrong rather
than less.

**No command is added.** An entry is written into the frontmatter by whoever answers the rule,
which is how every structured list in this repository is written while WP11 is unbuilt. A
command belongs with the MCP operation it stands behind.

**No judgement of an answer.** Whether `met` is true, whether a note is a good reason, whether a
`not-applicable` is honest: none of it is a question a deterministic gate may ask, and section 9
says so. The gate counts answers.

**No rule ships.** `given/builtin/` stays empty, so the first checklist this gate judges will be
in the piece that writes the shipped set.

**No change to `rules_hash` or to anything sealed.** The set is read and not written here.

**No new honesty for the four quiet rows.** G-Supply, G-Secret and G-Test stay
`not-implemented`; this piece removes one row from that list and touches none of the others.

<!-- xeno:section:constraints -->
## Constraints

**Section 9 limits the gate and the limit is binding.** "Every entry carries a result, and
nothing more" rules out judging the answer, and the paragraph's purpose rules out counting only
the entries present. Both halves of that sentence are constraints on this piece, in opposite
directions.

**Section 12 decides how completeness is computed.** A lens cannot add to or subtract from the
counted set, so the count runs over the effective rule set and asks the entries about it, never
the other way round.

**Section 7's `From` column is P0 and the checklist is a P5 artifact.** The gate has to be
correct and quiet at four phases without reporting itself unrun.

**The resolver is called, not reimplemented.** Section 7 says G-Rules and G-Policy resolve the
same tree and should do it once; two resolutions that could disagree about the set would make
`rules_hash` meaningless as a record of what applied.

**The frontmatter key follows the three that exist.** Hyphenated section id,
snake_case frontmatter key, which is what `open-questions` and `open_questions` already do. The
choice is recorded, since section 9 writes the example under the section's own name.

**Deterministic, network free, one dependency.** As everywhere here, and `internal/gates` says
it in its own package comment.

**88 columns, SPDX, `gofmt`, `go vet`, the suite and `./xeno gate verify` at exit 0.**

**One intent, one branch, one issue.** `160-a-review-rule-is-a-suggestion`, #160, labelled wp4,
with the remaining pieces named there and not carried here.
