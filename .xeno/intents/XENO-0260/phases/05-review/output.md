---
intent: github.com/triplem/xeno#258
phase: 05-review
created: "2026-10-06T11:41:53Z"
schema_version: "1.0"
runner_version: dev+7885661.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8d6592f1dfe79041e243225ce35f0467bca6768ade1d8b7328aefe5c08aaf7be
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: deviation
      note: Three, all recorded and all caught by measuring rather than reading. (1) eleven lines of the replacement came out at 89 or 90 columns and were reflowed with a wrapper over the whole block, which is what replacing a paragraph whole means in practice. (2) that reflow split a code span across a line break; CommonMark joins it, so it would have rendered correctly and read as a mistake in the file, and the sentence was rewritten to name no command. (3) P4's re-run of the figures returned 65 and 20 against the 64 and 19 written in P0 and P3, because this intent's own P1 entered the trail in between — the document now names its population rather than the figure being rounded, which is the honest repair of a count taken from inside the thing it counts. No criterion was falsified; criterion 2 is the one that moved and P4's results say how.
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: No interface changes. One document, and git diff --stat names it alone outside the trail. No Go source, no command, no flag, no artifact field, no gate, no template and no rule; the clause table and both reader columns are byte-identical. The document is this repository's own measurement of its own clauses and is not shipped by xeno init --vendor, so no adopter sees anything. What changes for a maintainer is which section a specification commit goes to, which is the whole point.
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: None added, none weighed, and none could help. go.mod is untouched and the diff is Markdown. The fault repaired was a section number copied from the document that had it wrong, and the thing that catches that is resolving the reference against the specification's headings when it is written, which is the P0 learning and costs a grep.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** No normative document is touched: `docs/process-definition.md` and
`docs/implementation-plan.md` are unchanged, and `docs/clause-readers.md` says of itself that it "is
a measurement, not a specification. Where it and the documents disagree, the documents win." Nothing
is invented — no field, no gate, no template version, no rule — and the two specification wordings
that would add one are drafted in a comment on #258 and left for a person. The work belongs to #258
under `wp5`, on a branch carrying one intent, and the commit references it.

**This is the half of #258 an agent can do, and the boundary is the point.** The issue's "Done when"
asks each of three clauses to be decided and the decision recorded. All three are now decided and the
third was decided this session with a route neither earlier option offered. What remains is two
specification commits and the code behind each, and the first standing rule makes the commits a
person's and makes them come before the code. Drafting the wording and correcting the document that
misdirects it is the work available on this side of that line; starting the code is not.

**The decision was put one at a time, with its options, their consequences and a recommendation.**
One question, three options, each costed, the recommendation named with its reason. The third option
— leave it to a person and record the reason — was included because `docs/clause-readers.md` exists
to record exactly that answer and a question that omitted it would have been asking for permission
rather than for a decision.

**The recommended option was found rather than offered.** The earlier comment on #258 put two
sub-routes: re-judge the trail, or invent a forward-only anchor for a gate's own checks. The third
costs nothing because the anchor already exists — every P1 artifact declares its template version and
`model.Template` already reads it. It was found by reading an artifact's frontmatter, which is the
P2 learning: a field with one reader reads as a field with no reader when it is taken from the
specification's prose.

**Every negative result was re-measured with the thing present.** #263's convention. "Section 9 says
nothing about acceptance criteria" rests on a search that returns five hits elsewhere in the same
document, and on reading 1255 to 1485 directly; a search returning nothing everywhere would have
meant the pattern was wrong. Both counts were taken with two patterns that agreed, and the two
over-88 lines were established as pre-existing by stashing the change and re-measuring rather than
by assuming.

**The verification contradicted the implementation and the document changed, not the figure.** P4's
re-run returned 65 and 20 against the 64 and 19 written in P0 and P3, because this intent's own P1
had entered the trail in between. A count taken from inside the thing it counts moves while it is
taken, and the measuring intent always looks like the exception it is arguing for. The document now
names its population, which is the honest repair; rounding the figure would have been the other one.

**Nothing in the clause table moved, and that is deliberate rather than an omission.** A90: a reader
that cannot fail is worse than none. Both rows still report no reader, because nothing fails today if
a question recommends without a reason or a mapping covers three of five criteria. A decision is not
a reader, and the one alternative that would have made this document misleading rather than
incomplete was the one that moved those columns.

<!-- xeno:section:release-notes -->
## Release notes

**`docs/clause-readers.md` named the wrong section for a specification commit, and now names section
5.** The passage explaining why G-Test's mapping half has no reader said a numbering convention is
"an addition to section 9". Section 9 is the Rule model. The word "acceptance" occurs six times in
the specification — 572 and 692 in section 5, where `template.yaml` holds the section ids and the
required fields; 710 and 713 in section 6's phase table; 942 in section 7's G-Test row; 1775 in
section 12 — and not once in section 9. #258 and its first comment had both copied the number, which
is why the correction says where it came from rather than only what it is.

**The figures measured one of the two halves the check needs.** A completeness check reads a P4
`test-mapping` and asks whether it covers every criterion P1 raised. Only the criteria were counted.
Measured over the trail as it stood before this intent: 19 of 64 P1 artifacts number their criteria
as a list, and **7 of 64 P4 mappings cite a criterion by number**. So the cost of applying the check
to the trail is 57 verdicts, not the 46 the passage stated. The 2026-10-05 figures stay with their
date and their population, because nine intents arrived between the measurements and a reader shown
only the newer number cannot separate growth from method.

**Both unread clauses have been decided, and the document says so and what each waits on.** The
recommendation's reason becomes `reason` on `model.Option`, required by `QuestionAsked` wherever
`recommended: true`. The mapping's completeness is anchored on the template version: section 5
requires numbered criteria from `requirements@1.1.0` and a numbered mapping from
`verification@1.1.0`, and the check judges only artifacts declaring 1.1.0 or later.

**That anchor needs nothing new.** G-Policy has `judgedUnder` and a gate's own checks were said to
have no equivalent, so a forward-only check looked like a specification question of its own. Every
P1 artifact already declares `template: requirements@1.0.0` and `model.Template` already reads it,
so the anchor is read rather than added, nothing is re-judged, and the second standing rule's budget
is untouched.

**The section 8 row gets the explanation it has never had.** The mapping row has carried three
paragraphs since #212; the reason-field row carried a bolded cell and an issue number, which made the
cheaper of the two clauses look like the one nobody had thought about.

**Neither reader column changes.** Both rows still report that nothing reads the clause, because
nothing does. A decision is not a reader, and the table's own definition of one is the thing that
would fail if the clause were violated.

**What is still owed, and to whom.** Two specification commits, drafted in a comment on #258 and a
person's to make under the first standing rule, and the code behind each: `reason` on `model.Option`
with its gate and tests; `requirements@1.1.0`, `verification@1.1.0` and the completeness check in
G-Test. Each moves one reader column, and this passage will be replaced a third time when they do.

<!-- xeno:section:residual-risk -->
## Residual risk

**#258 is not closed by this.** Two of its three clauses wait on specification commits that are a
person's, and the issue's "Done when" asks for "a specification commit and the code that follows it"
or a row recording a deliberate person-reader. Neither clause has either yet. What this intent
delivers is the decisions recorded where a reader of the document meets them, and the section number
corrected before anybody writes a commit against it.

**A reader could take "decided in #258" for "implemented".** Each paragraph ends on the reader column
being unchanged and on what the clause waits on, and that sentence is the only thing preventing the
misreading. It is a person-reader guarding a document about clauses whose readers are people, which
is either fitting or the problem, depending on who is reading.

**The figures go stale on the next intent that numbers its criteria.** They are dated and their
population is named, which is the most the pass's own rule about its ageing allows, and nothing
recomputes them. Three documents in this repository now carry trail-wide counts and none of them has
a reader that would notice the count being wrong.

**The document's prose and its table can disagree and nothing notices.** The table says 64 and 19 and
the paragraph says what the 64s are; whoever updates one has to update the other. XENO-0259 recorded
the identical gap in the audit baseline's comment a few hours earlier, which makes this a shape
rather than an instance and a candidate for the next thing somebody writes down.

**The wrong section number stands in sealed artifacts.** XENO-0258's phases and the earlier ones that
restated it carry section 9, and section 11 means they stay. A reader of those meets it with nothing
beside it. The third time this trail has recorded that residual risk, by the same mechanism each
time.

**One of the nine numbered P1 artifacts already fails a check by number**, from the 2026-10-05
measurement, and this intent neither re-found nor repaired it. It is carried in the document's
figures and nowhere else, and the thing that would read it waits on the section 5 commit.

**This passage is written to be replaced a third time.** Each specification commit makes one of its
two decision paragraphs historical and moves one reader column. Nothing records that obligation
outside this artifact and #258.
