---
intent: github.com/triplem/xeno#258
phase: 02-design
created: "2026-10-06T11:33:43Z"
schema_version: "1.0"
runner_version: dev+7885661
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7b909285c7fa38334a1744637e817750cc2048539d1c2896d533b458beb0482a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: 'The completeness half of G-Test is anchored on the template version: section 5 requires numbered criteria from requirements@1.1.0 and a numbered mapping from verification@1.1.0, and the check judges only artifacts declaring 1.1.0 or later.'
      rationale: 'It is forward-only with no new field and no new anchor to invent. All 64 P1 artifacts declare template: requirements@1.0.0 today and template is already an enumerated frontmatter field read by model.Template, so the judgedUnder equivalent a gate''s own checks were said to lack is already in every artifact. Re-judging the trail is 57 P4 verdicts measured on 2026-10-06, not the 46 estimated from the P1 half, and 57 overrides written in one sitting are 57 reasons nobody reads, which is A90''s hazard in bulk. The cost accepted is that a template bump reaches the shipped plugin and its digest, and that the check is silent until the first intent after the bump.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The passage is replaced whole and leads with the section number.** It had three paragraphs and
three faults — the wrong section, half the measurement, and a sentence about what it waits on that is
now answerable — and the fault a reader acts on first is the section number, because it is the one
that sends somebody to write a commit in the wrong place. So the replacement states it in its first
sentence rather than inside the argument about cost.

**Both figures are given and the one that is the cost is named.** 19 of 64 P1 artifacts number their
criteria; 7 of 64 P4 mappings cite a number. A completeness check is judged at P4, so 57 is the cost
and 45 is not, and the passage being replaced quoted the P1 figure for the P4 population. Giving one
number without saying which population it is of is how that happened, so both are given with their
population.

**The 2026-10-05 figures stay, with their date.** The document's preamble makes its own ageing
visible rather than repaired. The trail grew from 55 P1 artifacts to 64 between the two dates, so
part of the difference is the trail and part is method, and a reader who is shown only the newer
figure cannot tell which. The alternative — a footnote saying the older measurement was wrong — would
be wrong itself: it was right about the population it counted.

**The decision goes under the table and not into it.** A90: a reader that cannot fail is worse than
none. Section 8's row still reads that there is no field and the mapping's still reads nothing,
because that is what fails today. What a decision changes is what a reader of the explanation
learns, which is that the row has been triaged and by what.

**The section 8 row gets a passage of its own.** The mapping row has three paragraphs and this one
has a bolded cell. The asymmetry is an accident of which issue was open when, and it leaves the
cheaper of the two clauses looking like the one nobody has thought about — when in fact it is the one
whose shape was argued on both sides in #247 and settled. One paragraph, in the same place and the
same form as the mapping's.

**Nothing says what the specification wording is.** The two drafts live in a comment on #258. A
document carrying draft wording for a commit nobody has made would be a specification in everything
but name, which is the first standing rule's whole concern, and it would go stale the moment the
person changed a word.

<!-- xeno:section:alternatives -->
## Alternatives

**Correct the section number and nothing else.** A one-word diff, and the smallest honest change. It
leaves the figures measuring the wrong population and leaves both rows silent about having been
decided, which is the state that sent this intent to re-derive a triage the issue had already done.
Rejected: the convention about a paragraph changed a second time exists because a small diff is
precisely when the words around it are not re-read, and this passage is the instance.

**Replace the figures with the 2026-10-06 ones.** Cleaner to read and it is what most documents do.
It also erases that the trail grew by nine intents between the measurements, so a reader cannot tell
method from growth, and it contradicts the document's own statement that its ageing is visible rather
than repaired. Rejected on the document's own rule.

**Change the reader columns to name `QuestionAsked` and G-Test.** The table would then describe the
end state everybody has agreed on. Nothing would fail if either clause were violated, which is this
table's definition of a reader, so the rows would be false in the direction A90 names — and the
document would stop being a measurement. Rejected, and it is the one alternative that would have
made the document actively misleading rather than incomplete.

**Add a fourth kind to the clause taxonomy, "decided, awaiting a commit".** It would put the state in
the table where it is read rather than under it. It is also a change to how the pass classifies
clauses, which would ask the 219 sentences to be re-walked for consistency — #202's shape of work, an
intent of its own, and a taxonomy change made to carry two rows. Rejected as out of proportion.

**Record the decisions in `docs/assumptions.md` instead.** They would sit with A90 and A44, among
rows this project already treats as the register of what was decided and why. The register is for
assumptions, and section 8 of the specification is explicit that a decision is not one: it cannot
become wrong, only regretted, and carrying it as an assumption would ask somebody to confirm what
they just decided. Rejected on that distinction, which the specification draws deliberately.

**Wait for the specification commits and change the document once.** One edit instead of two, and the
rows would move to naming their reader in the same breath. It leaves a wrong section number in the
file for as long as the commit takes — and the commit is the thing the wrong number misdirects.
Rejected on the order: the correction is worth least after the event it exists to prevent.

<!-- xeno:section:impact -->
## Impact

**`docs/clause-readers.md`, one passage replaced and one added.** The mapping explanation keeps its
place under the table and its first paragraph about `examples/rules/mapping-is-complete.yaml`
unchanged; what is replaced is the paragraph that named section 9 and quoted the P1 figure. The new
section 8 paragraph sits beside it, in the same form. The clause table, the four-kinds table, the
counts, the architectural-property table and the findings section are untouched.

**What a reader of either row now learns.** That the clause has been decided, what it was decided to,
and that it waits on a commit a person makes. Before, both rows read as gaps of unknown age, and the
document's closing sentence sent the triage of them to the reader — who would have redone work the
issue had finished.

**What does not change, and is the point.** Both reader columns. Nothing fails today if a question
carries a recommendation with no reason, or if a mapping covers three of five criteria, and the table
goes on saying so. The count of thirty-eight rows is unchanged, so the figure the implementation
phase of the pass reported still reconciles.

**No normative document, no field, no gate, no template.** `docs/process-definition.md` and
`docs/implementation-plan.md` are untouched; section 7's G-Test row already asks for the mapping's
completeness and needs no change under either decision. Nothing in `cmd/`, `internal/` or
`.xeno/plugin/` is touched, so the suite and `gate verify` assert absence of accident.

**Adopters see nothing.** The document is this repository's own measurement of its own clauses and is
not shipped by `xeno init --vendor`.

**What is still owed after this.** Two specification commits and the code behind each: `reason` on
`model.Option` with `QuestionAsked` requiring it where `recommended: true`, and
`requirements@1.1.0`, `verification@1.1.0` and the completeness check in G-Test gated on the declared
template version. The rows change to name their readers in the intent that implements each, which is
the second time this passage will be replaced and the reason it is being replaced whole now.
