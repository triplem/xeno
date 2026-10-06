---
intent: github.com/triplem/xeno#258
phase: 02-design
created: "2026-10-06T15:10:05Z"
schema_version: "1.0"
runner_version: dev+90593d0
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ab336dd02e1ba2a3d6a73836c136266d22d1cdbe52e465a7d74f595aca28ee22
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: 'The agent writes the two specification amendments, in two commits, with the wording drafted on #258 unchanged: section 8 says the recommendation''s reason sits on the option the recommendation names, and section 5 gains a subsection making an acceptance criterion identifiable from requirements@1.1.0, anchored on the template version.'
      rationale: 'The first standing rule says the documents are not editable by the agent, and its purpose is that a person decides the specification. The person did decide: they were handed the drafted wording on #258 twice, read it, declined to write the commits themselves and instructed the agent to make them. So the rule''s substance holds and its letter does not, and the exception is recorded in P0''s problem section, in this decision and in both commit messages, because an exception with no trace is indistinguishable from the rule never having applied. Two commits rather than one, because the rule makes each specification change its own commit and the two clauses share nothing but the issue that tracks them. The wording is committed unchanged, because rewriting it here would commit text the person had not approved.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**Two commits, not one.** The first standing rule says a change to a normative document "is its own
commit made before the code that follows from it". Two changes, two commits. They are also two
clauses with two decisions taken on two days for two reasons, and a reader looking for why the reason
moved onto the option should not have to read past a paragraph about acceptance criteria to find it.

**Section 8's sentence is replaced and the reason follows it in the same paragraph.** The clause and
its justification in one place, because the thing being fixed is a clause that said what without
saying where, and a reader who finds the where in a different paragraph has the same problem one step
along.

**Section 5 gets a subsection and not a sentence in an existing one.** The convention has three parts
— the numbered list, the mapping that cites it, and the template version that bounds it — and none of
the three subsections it could have joined is about any of them. Language is about structure and
strings, Context economy about reading cost, Rendering about anchors. A heading naming what it is in
words is also what `CLAUDE.md` asks of a heading.

**The anchor is named as the template version and nothing else.** Not a date, not a runner version,
not a new field. Every P1 artifact already declares `requirements@1.0.0`, so the anchor is read
rather than added, which is the second standing rule satisfied rather than argued with. It is also
the reason this decision cost nothing where the two routes #258 first offered cost 57 verdicts or a
new specification question.

**The subsection says what is *not* judged, in the same breath as what is.** "An artifact declaring
`requirements@1.0.0` is judged as it always was." A forward-only clause that states only the new
requirement leaves a reader to work out whether the trail is in scope, and the whole value of this
route is that it is not.

**Section 7 is left alone.** Its G-Test row already asks for the mapping's completeness. Section 5
says what makes a criterion identifiable and section 7 says what reads it, which is the division the
document already uses and the one `docs/clause-readers.md` had to split a row over when it was
broken.

**The drafted wording is committed unchanged.** It was offered on #258 to be changed freely and the
person asked for it as it stood. Rewriting it here to sound more like the surrounding prose would
make the committed text something nobody had approved, which is the opposite of what the first
standing rule is protecting.

**The exception is recorded in the commit messages and not only in the trail.** The agent wrote a
normative document. `CLAUDE.md` says it does not, the person instructed it to, and the commit is
where a stranger with only the current tree will be standing when they ask why.

<!-- xeno:section:alternatives -->
## Alternatives

**One commit for both amendments.** Shorter history, one message, one diff. The first standing rule
says each change to a normative document is its own commit, and the two clauses have nothing to do
with each other beyond the issue that tracks them. Rejected on the rule and on the reader: a bisect
over a specification change should land on one clause.

**Put the reason on `Question` instead.** "with a reason" reads as one reason per question and it
costs no field that is empty on every option but one. #247 argued it and #258 settled against it: a
recommendation that later moves to another option would leave the reason describing the option it
used to be about, with nothing able to detect it. Rejected as a decision already taken, and recorded
here because the alternative is the better-sounding one.

**Make the criteria convention a `schema_version` bump.** It is the other version in the frontmatter
and it is what a reader reaches for first. Section 5 says `schema_version` "is read, never enforced
backwards", so a check keyed on it would be the one thing that section forbids, and the convention is
about what a template's section must contain rather than about the shape of the artifact. Rejected,
and it is the mistake the subsection's second paragraph exists to prevent.

**Put the criteria convention in section 6, beside P1's row in the phase table.** That row already
says P1 produces "specification with acceptance criteria and non goals", so a reader is there. It
would put a requirement about a section's content in the subsection about the phase sequence, and the
required fields of a template live in section 5 — which is the whole correction XENO-0260 made.
Rejected, and it is the near miss that section number was.

**Apply the criteria convention to the trail and let the overrides record it.** The honest route and
the one the process has a mechanism for. 57 P4 verdicts, 57 reasons, written in one sitting by one
person: A90's hazard in bulk, and a record weaker than no check. Rejected in #258 and not reopened.

**Write the amendments in the document's voice rather than the draft's.** The surrounding prose has a
particular rhythm and the drafts were written to be read in an issue. Rewriting them would commit
text the person did not see. Rejected: the draft was approved as the wording, and a specification
whose committed text differs from the approved text is a worse fault than an unevenness of tone.

**Leave both clauses unsayable and record them as person-read in `docs/clause-readers.md`.** #258's
"Done when" allows it and it costs nothing. It was the third option on both decisions and the person
chose the mechanism on both. Rejected as a decision taken, and named because it remains the cheap
answer if either implementation turns out worse than expected.

<!-- xeno:section:impact -->
## Impact

**`docs/process-definition.md`, two places.** Section 8's sentence under "Open questions have to be
resolved", replaced as a paragraph. A new subsection after "Rendering" at the end of section 5. No
other line of either section changes, and no other section is touched.

**What becomes buildable.** `reason` on `model.Option` with `QuestionAsked` requiring it where
`recommended: true`; `requirements@1.1.0` and `verification@1.1.0` in the plugin; the completeness
half of G-Test keyed on the declared template version. Each was a second-standing-rule violation an
hour ago and is now a step.

**What becomes readable that was not.** Why the reason is on the option rather than the question —
the drift argument, which lived in a comment on #247 and then in a comment on #258, neither of which
a reader of the specification reaches. And that the criteria convention does not reach the trail,
which is the single fact that decides whether somebody is about to be asked for 57 approvals.

**Nothing in the trail moves.** The specification is in no `artifacts_hash` and in no `rules_hash`,
so every verdict stands and `gate verify` reports the same count before and after. That is also the
limit: nothing fails if a later commit contradicts either amendment, and the documents are read by
people and by the agent through `CLAUDE.md`'s pointer.

**`docs/clause-readers.md` goes stale by one sentence.** Its two decision paragraphs, written in
XENO-0260, say each clause "waits on the section 8 commit" and "waits on the section 5 commit". After
this they wait on the code. That is a paragraph somebody has to replace a third time, which XENO-0260
predicted and is why it was replaced whole rather than corrected in place.

**No Go code, no gate, no template, no rule, no field.** Nothing in `cmd/`, `internal/` or
`.xeno/plugin/`, so the suite and `gate verify` assert absence of accident. Section 7's gate list and
the four check results are untouched.

**Adopters are reached by a document and not by a release.** `docs/process-definition.md` is not
shipped by `xeno init --vendor`; what an adopter gets is the plugin and the binary, neither of which
changes here. A project tracking the specification sees two clauses it cannot yet satisfy, which is
the ordinary state of a document that leads its implementation.
