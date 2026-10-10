---
intent: github.com/triplem/xeno#334
phase: 02-design
created: "2026-10-10T15:41:57Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4a116381f0666bacf836f8b5b1182c8d9d4b7ca308030fc68a5fbf8ee99af393
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

These are the construction decisions of the section, not decisions put to a person. The one
of those is P0's D-1.

**It is section 10 of the page, not section 9.7.** Section 9 is about a specification
framework; this is about a development lifecycle methodology, which compares with the
process rather than with one template section. The two decisions are of the same kind and
about different kinds of tool, which is exactly what section 1's paragraph says of the
orchestrator and section 9, so the page grows a third top-level section rather than a
subsection inside the second. Its heading names it in words.

**Six subsections, in section 9's order, with 10.4 in 9.4's place.** 9.4 is the
measurement, and the measurement in this case is #323's weighing of three shapes with what
each costs, so 10.4 holds that. The order matters for the same reason 9 has it: the
decision before the arithmetic, because the decision does not rest on the arithmetic and a
reader who meets the numbers first will think it does.

**Five rows and four columns.** The fifth row, rules and learning, is marked in the prose
under the table as the addition OpenSpec did not need, with what it costs to add it: the
fourth column's vocabulary has to admit a row where the answer is not "replaced" or
"weaker".

**The sample is a line in 10.1 and a clause in the second row, not a subsection.** #334
says it is a line. What it contributes is one difference that matters on the rows already
there, and a subsection would have to repeat four cells to say one thing.

**AI-DLC is quoted, with file and line, where the two documents disagree.** A paraphrase of
a disagreement is where the disagreement goes missing: #334 names this as the one point the
section has to settle, and a reader who wants to check which document is current needs the
two sentences and the two addresses. Elsewhere the section describes rather than quotes,
because a section made of quotations is a reading and not a decision.

**The third row rests on the code and not on either document.** The guide and the reference
disagree; the implementation decides which is current, so the row names the two functions
and the hook, with their line numbers in an 8,520-line file, and states the three ways the
override is refused. Without the refusal path an enforcement claim is the flattering half of
the truth.

**Negative claims are made only where a positive control was run in the same command.** One
claim in the section is load-bearing and negative: that nothing in AI-DLC's core binds an
intent to a tracker issue. It is made from a search of the whole tree at the pinned sha
whose positive control was a search of the same tree for a term that is there, and the one
hit the negative probe did return is accounted for in P4 rather than dropped.

**Both tags are named with both shas.** Each pin is an annotated tag, so the ref resolves to
a tag object and the commit is one more dereference away. The Sources entry says which sha
is which, because a reader who resolves `v2.11.0` himself gets `4079edbe` and would
reasonably think the section cited the wrong thing.

**The conditions for revisiting carry the alternative D-1 did not take.** That is where the
maintainer put it, and writing it there rather than into an issue keeps the reason and the
alternative in the same place, which is what makes a condition for revisiting readable by
somebody who was not in the conversation.

<!-- xeno:section:alternatives -->
## Alternatives

**A subsection of its own for the hosted sample.** Rejected because #334 says it is a line,
and because the four rows it would fill are the same four with one cell changed: its record
lives in a deployed database — Neptune and DynamoDB, per its own README — rather than in a
git tree, which is a sharper version of the second row and not a fifth comparison. A
subsection would have restated three cells to say one thing.

**Writing the third row as "advisory", following the guide.** Rejected. The guide's sentence
at `docs/guide/09-rules-and-the-learning-loop.md:140` is unqualified and the paragraph two
lines above it already describes gate-fired sensors, so the file disagrees with itself
within one screen. The reference and the implementation agree with each other against it,
and the implementation is what runs. The row therefore says the reference is current, and
says the guide's sentence reads as one that was true before gate-fired blocking landed.

**Taking the side of the guide and calling the contradiction unresolved.** Rejected for a
different reason: #334 asks the section to settle it, and "the documents disagree" is what
#323 already said. An evaluation that forwards the ambiguity has done the reading and not
the work.

**Recording any mechanism as taken.** Rejected on the first standing rule. A sensor firing
on write is a change to the process definition's section 7 before it is a line of code;
#340 holds it, and it already has the use case the maintainer gave on #323. 9.6's discipline
is that the evaluation names a mechanism and an issue of its own takes it, and a section
that quietly took one would be a specification change wearing an evaluation's clothes.

**Stating the cost of the source-binding mechanism the way the question stated it.**
Rejected, and this is the one place the section departs from the wording of the decision it
records. Option two on #334 named its cost as blob OIDs not being stable across a rebase
that changes a tree, against section 6's "Rebasing is safe where the trees are". AI-DLC's
own reference claims the opposite property for its mechanism: `docs/reference/20-commit-provenance.md:138`
says attribution "keys on blob content (OIDs), not commit ancestry — a reviewed change that
lands squashed with others still verifies". A rebase that does not change a tree leaves
every blob OID as it was, and one that does change a tree reports `drifted`, which is the
right answer rather than a false one. So the section states what the mechanism does and
rests the decision on `docs/process-definition.md:480`, which is the reason that decides it
and is a decision already taken. Nothing about D-1 moves: the recommendation was the first
option and the first option's consequence is what the section writes down. What would have
been wrong is repeating a cost the source contradicts, inside an artifact that seals it.

**Leaving the three stale figures from #323 silently corrected.** Rejected. The section
carries the current figures and P4 records which moved and by how much, because a reader
comparing this section with #323 will find three differences and should be able to tell a
correction from a disagreement.

**Reading `docs/process-definition.md` into the scope rather than quoting it.** Rejected on
scope economy: 132 KB for four quotations, in a scope whose point is that one file grows.
What it costs is that a later phase cannot search the document, which is why the four
passages are in P0's context rationale in full rather than by reference.

**Numbering the subsections 10.1 to 10.6 at all.** Kept, against `CLAUDE.md`'s rule about
borrowed numbers, because these are this page's own numbers: section 9 has 9.1 to 9.6 and
the page is referred to by them from `docs/v2-delta.md` and from the plan. The headings
still say what they are about in words; the number is the page's index and not a substitute
for one.

<!-- xeno:section:impact -->
## Impact

One file changes, in four places, in this order.

**`docs/orchestrator-evaluation.md`, frontmatter.** `revision: 3` becomes `revision: 4` and
`date: 2026-10-08` becomes `date: 2026-10-10`. A whole new section is what took it from 2 to
3 in `a010cd1`; the one-sentence change in `5044a7a` left both alone, which is the line this
follows.

**Section 1, the paragraph that names section 9.** Replaced whole. It is being changed for
the second time and `CLAUDE.md` asks for a replacement read back as a paragraph rather than
an edit into a sentence, because a small diff is exactly when the words left behind are not
noticed. The replacement has to say three things the current one cannot: that there are now
two decisions of this kind beside the orchestrator, that they are about different kinds of
tool, and that the second of the two compares with the process rather than with the
platform.

**A new section 10, after 9.6 and before Sources.** Six subsections and one five-row table.

**The Sources list.** One new entry, in the form of the two that are there: both
repositories, both tags, both shas with which is which, the date of the reading, and each
file named by what it was read for.

What this does not touch, and why each is safe to leave.

**No code, so no test moves.** `git diff main -- '*.go'` is empty by construction. The gates
that have anything to say about this change are G-Schema, G-Secret, G-Learning and
G-Policy's review checklist; G-Build and G-Test read declared evidence, and the evidence
this intent declares is the four commands `CLAUDE.md` names, run against a tree whose Go
files nobody touched.

**No change to `docs/process-definition.md` or `docs/implementation-plan.md`.** Both are
quoted and neither is edited. Checked by two empty diffs in P4 rather than asserted here.

**No change to `docs/README.md`.** Its entry describes this page as the record of which
agent platform v2 delegates a phase to, which was already narrower than the page after
section 9 landed. #334 asks for the evaluation's own index entry, which is section 1's
paragraph, and widening a neighbouring document's description of it is a different change
with a different reader.

**The context budget.** The scope resolved 37,939 bytes against a budget of 140,000. The
section adds of the order of 16,000 bytes to the one file in scope that grows, which lands
near 54,000, so the budget holds with room. This is deliberate: XENO-0286 set a budget from
files measured before the writing that enlarged them and carried an advisory G-Schema
finding from P4 to the end with no repair available, because the budget sits inside the
intake's `artifacts_hash`.

Two assumptions this design works under, marked as assumptions because there is no register
for them in a new intent and because each would change the section if it turned out wrong.

**Assumption: a pinned section is the right form even though the pin will be stale within
days.** AI-DLC cut three preview tags in the three days around `v2.11.0` and its default
branch moved again on the day of this reading. So the section is accurate about one commit
and will not be accurate about `main` for long. Taken as right rather than as a cost,
because that is what section 9 did and what #323 argued for in as many words: a document
that reads at `main` is stale by the time it is reviewed, and a pinned one is at least
stale at a stated address. What the section owes in exchange is that every cell say where it
was read, which is acceptance criterion four.

**Assumption: `aws-samples/sample-collaborative-ai-dlc` compares on the same rows without
its own reading.** Nobody has read its five rows, #323 did not and this intent does not. The
line the section gives it is held to what its own README at the pinned commit states — an
early-preview AWS sample deployed into the adopter's account, which can start an intent from
a GitHub, GitLab or Jira issue, and whose record lives in Neptune and DynamoDB — and makes
no claim about its gates, its sensors or its approvals. If that line later turns out to
understate it, the section is incomplete rather than wrong, which is the trade a line is for.
