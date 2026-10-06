---
intent: github.com/triplem/xeno#258
phase: 05-review
created: "2026-10-06T15:16:13Z"
schema_version: "1.0"
runner_version: dev+90593d0
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 74311df3a4b4a3038ee19e55041159266092317375c53fab240f53aeff6e6818
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: deviation
      note: 'Four, in P3, and three are against this intent''s own earlier phases. (1) criterion 5 asked for two commits and there is one: the repository squash-merges, so two on a branch arrive as one and two on main would have meant two pull requests and, under the third standing rule, two intents for two paragraphs. Reported failed in P4 rather than reinterpreted. (2) the drafted wording was committed with two changes where P2 decided it would be unchanged — the section 8 heading said ''replacing the third sentence'' and the replacement did not contain that sentence, so committing it as headed would have deleted an existing normative clause; and the section 5 text cited ''57 of the 64 in this trail'' where the specification cites this repository nowhere, checked by a search returning zero there and four in CLAUDE.md. (3) the section 8 amendment is a second paragraph and not a replacement, so the convention about a paragraph changed twice gave way to not losing a sentence; both paragraphs were read back as prose. (4) the specification now names two template versions that do not exist, which is the forward-only anchor working and is named so the next reader finds it deliberate.'
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: 'No interface changes, and deliberately no code at all. Twenty lines of docs/process-definition.md and nothing else outside the trail: no Go source, no field, no command, no flag, no gate, no template, no rule. The specification is in no artifacts_hash and no rules_hash, so every verdict in the trail stands and gate verify reports the same count before and after. It is not shipped by xeno init --vendor either, so an adopter sees nothing until the code follows. What a reader of the document gets is two clauses naming template versions that do not exist yet, which P3''s deviations records.'
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: 'None added and go.mod is untouched; the whole diff is Markdown. Nothing was weighed either: both clauses lean on mechanisms this process already has — a field on an existing struct, and the template version every artifact already declares — which is the second standing rule satisfied rather than argued with, and is why the chosen anchor cost nothing where the two routes #258 first offered cost 57 verdicts or a new specification question.'
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The first standing rule is broken deliberately, and this is the record of it.** "The documents are
not editable by the agent." The agent edited one. The rule's purpose is that a person decides the
specification, and the person did: the wording was drafted on #258, put in front of them twice, and
they declined to write the commit and instructed the agent to make it. So the substance held and the
letter did not. It is written in P0's problem section, in D-1, in this checklist and in the commit
message, because the rule makes this the exception and an exception nobody can find afterwards is
the same as the rule never having applied.

**The other two standing rules hold as usual.** Nothing is invented — the amendments add what the
artifacts may carry, which is what the second rule asks a specification change to do, and the anchor
is a field every artifact already has. The work belongs to #258 under `wp5`, on a branch carrying
one intent, and the commit references the issue without closing it: its first two clauses are not
finished until their code exists.

**No code is in this commit, which is the half of the first rule that is not being broken.** "A
change to it is its own commit made before the code that follows from it." `git diff --stat` names
`docs/process-definition.md` and nothing else outside the trail: no field, no template, no gate, no
test.

**A criterion is reported failed and not reinterpreted.** Criterion 5 asked for two commits and there
is one. P1 wrote it without asking how this repository merges, which was a question about the tree
and answerable with one `git log`; every merge on `main` is a squash, so two commits on a branch
arrive as one. The temptation was to read "its own commit" as satisfied by the branch and mark it
green. P4 says fail.

**Two sentences of the committed text differ from what the person approved, and they are being told
which.** The draft's section 8 heading said "replacing the third sentence", and the replacement did
not contain that sentence, so committing it as headed would have deleted an existing normative
clause. The draft's section 5 text cited "57 of the 64 in this trail", and the specification cites
this repository nowhere — checked, not assumed: the search returns zero there and four in
`CLAUDE.md`. Both were found by putting the text into the document and reading it there, which is
the one thing an issue comment cannot be read in.

**Every negative result was re-measured with the thing present.** #263's convention. "The
specification never cites this repository" rests on a search that finds four matches one file away.
"Nothing was deleted" rests on a diff of 20 insertions and 0 deletions rather than on the absence of
a complaint. "The clause binds nothing yet" rests on all 65 P1 artifacts being read for their
declared template version.

**The convention about replacing a paragraph gave way to not losing a sentence.** Section 8's
paragraph has been changed before, so the convention asks for a replacement; the replacement on
offer would have dropped a clause. The existing sentences are kept, the new material follows them,
and both paragraphs were read back as prose, which is the part of the convention that was actually
at stake.

**#235's amendment is deliberately not here.** It is decided and drafted, it is a different issue
under a different work package, and the third standing rule gives it its own branch and its own
intent. Three unrelated specification changes behind one reference would have been the batch the
conventions keep warning about.

<!-- xeno:section:release-notes -->
## Release notes

**Two clauses the artifacts could not carry are now enumerated.** Both were decided on #258 and
neither could become code, because the second standing rule makes an addition to what the artifacts
carry a specification change and the first makes that commit a person's, before the code.

**Section 8 says where the recommendation's reason lives.** On the option the recommendation names,
not on the question. A recommendation that later moves to another option takes its reason with it or
the writer refuses; a reason held beside the question would go on describing the option it used to
be about with nothing able to notice. The field is empty on every option but one and the clause says
so. #247 argued both homes and the argument lived in two issue comments, neither of which a reader
of the specification reaches.

**Section 5 gains "Acceptance criteria are identifiable".** From `requirements@1.1.0` the
`acceptance-criteria` section is a numbered list; from `verification@1.1.0` the `test-mapping`
section names each criterion by its number. Section 7 has asked G-Test for the completeness of the
mapping since it was written, and nothing could answer it because nothing identified a criterion — a
sentence cannot be reported as covered or uncovered.

**The anchor is the template version and nothing new.** An artifact declaring `requirements@1.0.0`
is judged as it always was, so nothing sealed is re-judged: all 65 P1 artifacts in this trail declare
1.0.0, and `template` is already in every artifact's frontmatter. The two routes #258 first offered
cost 57 verdicts or a new specification question; this one costs neither.

**No code is in this commit.** No `reason` on `model.Option`, no `requirements@1.1.0` in the plugin,
no check in G-Test, no change to `docs/clause-readers.md` — whose two rows still report no reader,
which is true until the code exists.

**The agent wrote a normative document, on instruction.** `CLAUDE.md`'s first standing rule says it
does not. The wording was drafted on #258, read by the person, and committed at their instruction
rather than rewritten by them. The exception is recorded here, in the intake and in the design
decision, because a reader with only the current tree is the reader this project writes for.

**Two sentences differ from the approved draft.** The draft's section 8 heading would have deleted
"An open question without options moves the whole of the thinking onto the person"; it is kept and
the new material follows it. The draft's section 5 text cited a count from this repository's trail,
which the specification does nowhere, and now reads "every mapping a project wrote before it adopted
the convention".

**One criterion failed.** Two commits were asked for and one was made, because this repository
squash-merges. Reported rather than reinterpreted.

<!-- xeno:section:residual-risk -->
## Residual risk

**The agent edited a normative document and the only thing stopping that becoming ordinary is this
record.** Three artifacts and a commit message say it was an exception and name where the
instruction was given. Nothing checks it: no gate reads `CLAUDE.md`, and the next session has only
the tree. The P0 learning proposes a convention for it and section 10 routes that through a merge
request against the rule set, which is where it now sits with every other learning nothing has read.

**Criterion 5 stays unmet.** One commit, and after the merge the history is what it is. A reader
bisecting for the day the reason moved onto the option lands on a commit that also moved the criteria
clause. The two-pull-request alternative is still available and still costs two intents for two
paragraphs.

**The specification names two template versions that do not exist.** `requirements@1.1.0` and
`verification@1.1.0` are in no `template.yaml`, so a reader of the document today cannot find what it
names. Deliberate — it is how the anchor is forward-only — and unbounded: nothing says how long the
window stays open, and nothing will notice if it never closes.

**Nothing will check that the template and the clause describing it agree.** When 1.1.0 is created,
G-Schema will compare an artifact's anchors against its declared template version, and nothing will
compare the template against the sentence in section 5 that says what it must require. The document
and the plugin can drift apart silently, which is a new instance of the shape
`docs/clause-readers.md` catalogues and the third such instance recorded in this session.

**`docs/clause-readers.md` is stale by two sentences.** Its decision paragraphs say each clause waits
on a specification commit; they now wait on the code. The third replacement of that passage, with
nothing to remind whoever makes it.

**Two clauses that were unsayable are now unread.** That is the honest size of this intent. The field
does not exist, the templates do not exist, the check does not exist, and a specification sentence
with no reader is the thing this project has a whole document for counting.

**The person is being told which two sentences differ from what they approved, and that is a person
reading a report.** A specification written by an agent on instruction is only as good as the
person's ability to see what landed, and the mechanism for that here is a paragraph in a message.
