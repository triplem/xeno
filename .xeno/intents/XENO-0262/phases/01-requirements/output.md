---
intent: github.com/triplem/xeno#258
phase: 01-requirements
created: "2026-10-06T15:09:16Z"
schema_version: "1.0"
runner_version: dev+90593d0
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 989902d3cc574f41c128fbd9d4f2c9043d2e368f264a2cc0fdd945cf01def8b2
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

Numbered, and P4's mapping cites the numbers.

1. **Section 8 says where the recommendation's reason lives.** On the option the recommendation
   names, not on the question. The sentence under "Open questions have to be resolved" is replaced as
   a paragraph rather than edited into.

2. **Section 8 says why it lives there.** A recommendation moved to another option takes its reason
   with it or the writer refuses; a reason held beside the question would go on describing the option
   it used to be about with nothing able to notice. The field being empty on every option but one is
   named as the cost rather than left for a reader to discover.

3. **Section 5 gains a subsection making an acceptance criterion identifiable.** After "Rendering".
   From `requirements@1.1.0` the `acceptance-criteria` section is a numbered list; from
   `verification@1.1.0` the `test-mapping` section names each criterion by its number.

4. **That subsection says the anchor is the template version and not the schema.** An artifact
   declaring `requirements@1.0.0` is judged as it always was, because what is sealed is never
   rewritten and a check reaching backwards would re-judge every mapping written before the
   convention existed.

5. **Each amendment is its own commit, and neither carries any code.** The first standing rule makes
   a specification change its own commit before the code that follows. `git log` over this branch
   shows two commits, each touching `docs/process-definition.md` and nothing else outside the trail.

6. **The wording is the wording drafted on #258, unchanged.** It was put there to be changed freely.
   Where it is not changed, that is the decision and not an oversight, and the artifacts say which.

7. **Nothing in sections 5 or 7 that already covers this is duplicated.** Section 7's G-Test row
   already asks for "mapping of acceptance criteria complete" and is untouched; the four check
   results are untouched; `schema_version` is untouched.

8. **`docs/clause-readers.md` is not changed.** Both rows still report no reader, which is true until
   the code exists.

9. **`xeno gate verify` exits 0, `go test ./...` passes, `go vet` is clean, `gofmt -l` prints
   nothing.** No Go source is touched, so these assert nothing was touched by accident. The
   specification is in no artifact's hash, so no verdict in the trail moves.

10. **The document still wraps at 88 columns outside tables and code blocks**, and the over-long
    lines after the change are the ones that were over-long before it.

11. **The agent having written a normative document is recorded where a reader meets it.** P0's
    problem section, this intent's learning, and the commit messages. The first standing rule makes
    this the exception; an exception with no trace is the rule never having applied.

<!-- xeno:section:non-goals -->
## Non goals

**Not `reason` on `model.Option`.** Not the field, not `QuestionAsked` requiring it where
`recommended: true`, not the tests. The first standing rule puts the code after the specification
commit, and putting it in the same intent would make the two indistinguishable in the history, which
is the thing the rule is for.

**Not `requirements@1.1.0` or `verification@1.1.0` in the plugin.** The specification naming a
template version that does not exist yet is deliberate: that is what makes the anchor forward-only,
and the template is created by the commit that implements the check. A plugin bump also moves the
shipped digest, which is a release-visible change and has no business in a document commit.

**Not the completeness check in G-Test.** It is the reader the second amendment exists to make
possible, and it is the next intent's.

**Not `docs/clause-readers.md`.** Its two rows report what fails today and nothing fails today. A
clause that can be carried is not a clause that is read, and A90 is why the distinction is kept:
moving a row on a decision rather than on a reader is what would make that document misleading.

**Not #258's third clause.** #248 decided the sequencing convention stays in `CLAUDE.md` permanently,
and nothing follows from it — the paragraph there already ends "Nothing checks this (#248)".

**Not #235.** Its section 5 amendment for an advisory finding is decided and drafted, and it is a
different issue under a different work package. The third standing rule gives it its own branch and
its own intent, and mixing the two would put three unrelated specification changes behind one
reference.

**Not a fifth check result, and not a change to the result set.** Neither amendment needs one. A4 and
A42 fixed that set and #235's decision was taken partly to avoid touching it.

**Not a numbering convention applied to the trail.** 45 of 64 P1 artifacts carry no numbered criteria
and 57 of 64 P4 mappings cite no number. The template version is the anchor precisely so that none of
them is re-judged, and this intent does not touch a sealed artifact.

**Not a change to any other clause the two sections contain.** The amendments are one replaced
paragraph and one new subsection. Section 5's Language, Context economy and Rendering subsections
keep their text; section 8's assumption register, its decisions-are-not-assumptions subsection and
the three exits for a question keep theirs.

<!-- xeno:section:constraints -->
## Constraints

**The specification is not in any hash, and that cuts both ways.** `artifacts_hash` covers a phase
directory and `rules_hash` the rule set; `docs/process-definition.md` is in neither. So amending it
moves no verdict in the trail and `gate verify` cannot notice a change to it at all. The document is
read by people and by the agent through `CLAUDE.md`'s pointer, and by nothing that fails.

**A specification sentence has to be true of a runner that does not exist yet.** Both amendments
describe behaviour no code implements. That is the normal direction for this document and it is why
each names the mechanism it leans on — the option carrying the reason, the template version carrying
the convention — rather than a function or a field name that a later rename would falsify.

**The template version is already in every artifact and must stay the only anchor named.** All 64 P1
artifacts declare `requirements@1.0.0` and `model.Template` reads it. Naming anything else — a date,
a runner version, a new field — would be inventing an anchor where one exists, which the second
standing rule forbids and which was the reason #258's earlier framing stalled.

**Section 7's G-Test row must not be touched.** It already asks for "mapping of acceptance criteria
complete". Restating the completeness requirement in section 5 would give one clause two homes, and
`docs/clause-readers.md` already carries a row that exists because one sentence of section 7 had to
be split over two readers. Section 5 says what makes a criterion identifiable; section 7 says what
reads it.

**Replacing rather than editing applies to section 8's sentence.** It has been amended before — the
free-entry clause and the two-to-four bound arrived separately — so it is the case the convention is
about, and it is replaced whole and read back as prose.

**88 columns outside tables and code blocks**, and the document has pre-existing over-long lines, so
the measure has to distinguish the ones this intent writes from the ones it finds.

**Two commits on one branch, each referencing #258.** The first standing rule makes each specification
change its own commit; the third makes the branch carry one intent and the commit reference its
issue. `Closes #258` belongs on neither, because the issue's third clause closed without a commit and
its first two are not finished until their code exists.
