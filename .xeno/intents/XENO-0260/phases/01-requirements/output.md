---
intent: github.com/triplem/xeno#258
phase: 01-requirements
created: "2026-10-06T11:32:55Z"
schema_version: "1.0"
runner_version: dev+7885661
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 95206ded104567fe592129ad62f0761629f426b1d43ebfa0dc5e0f9d35528f39
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

Numbered, and the mapping in P4 cites the numbers. This intent is about the convention that makes
that possible, so writing its own artifacts without it would be the clearest argument against it.

1. **The passage explaining why the mapping half has no reader names section 5.** Replaced as a
   paragraph, not edited into: the section number is one of three things wrong in it, and the
   convention about a second change to a paragraph exists for exactly the case where a small diff
   hides the words left around it.

2. **Both halves of the measurement are stated, with the date and the method.** 64 P1 artifacts, 19
   numbering their criteria as a list; 64 P4 artifacts, 7 citing a criterion by number. The cost of
   applying a check to the trail is given as 57 verdicts, which is the P4 figure, because that is
   the population a completeness check judges.

3. **The 2026-10-05 figures stay as a dated measurement.** Not overwritten. The trail grew between
   the two dates and only part of the difference is method, so a reader has to be able to see both
   rather than be told the older one was wrong.

4. **The section 8 row gains the explanation it has never had.** The mapping row carries three
   paragraphs and the reason-field row carries a bolded cell and an issue number. A reader learns
   from the table that one clause is unread and nothing about what would change it.

5. **Both rows say that the clause has been decided and what it waits on.** The document's closing
   section sends the triage of its unread rows to whoever reads it; a row that has been triaged and
   does not say so sends the next reader to do it again.

6. **Neither row's reader column changes.** Section 8's still reads that there is no field; the
   mapping's still reads nothing. A decision is not a reader, and this table's own convention is
   that a reader is "the thing that would fail if the clause were violated". A90 is the reason: a
   row made to look answered by a check nobody has written is worse than one that says it is empty.

7. **No normative document is touched, and no field, template version or gate.** `git diff --stat`
   names `docs/clause-readers.md` and nothing else outside the trail.

8. **`gate verify` exits 0, `go test ./...` passes, `go vet` is clean, `gofmt -l` prints nothing.**
   No Go source is touched, so these assert nothing was touched by accident.

9. **The document still wraps at 88 columns outside tables and code blocks.** Every line this intent
   writes is checked, because a paragraph replaced by hand is the common way that rule breaks.

<!-- xeno:section:non-goals -->
## Non goals

**Not the specification commits.** Section 8's wording for the reason field and section 5's for
numbered criteria are drafted in a comment on #258 and are a person's to commit under the first
standing rule, before any code. Nothing here anticipates either.

**Not `reason` on `model.Option`, the template bump, the G-Test check or their tests.** All four
follow a commit that does not exist. Writing them first is the second standing rule's invention, and
it would also be the worse order: a draft is cheap to change and code written against it is not.

**Not a change to either row's reader column.** They report what fails today and today nothing does.

**Not a row for the sequencing convention.** #248 decided it stays in `CLAUDE.md` and #255 declined
it a row, because this document is a pass over the two normative documents and a convention in
`CLAUDE.md` is a clause in neither. Declining it a second time is what makes that a rule rather than
a mood.

**Not a re-run of the pass.** The 2026-10-03 walk over 219 candidate sentences is not repeated and
the counts of the four kinds are untouched. This intent corrects one passage and adds one; a second
pass is #202's shape of work and would cost an intent of its own.

**Not correcting #258 or its comments.** An issue is not a file in the tree, and the correction is
already recorded on it.

**Not repairing the sealed artifacts that carry the wrong section number.** XENO-0258's phases and
earlier ones restate it where they restate the issue. Section 11 means they stay, and that is
recorded as residual risk rather than worked around.

**Not fixing the one numbered artifact that already fails a check by number.** The 2026-10-05
measurement found it and this intent neither re-finds nor repairs it: it is a sealed P1, and what
would read it is the check that waits on the specification commit.

<!-- xeno:section:constraints -->
## Constraints

**The document may be changed by the agent and the specification may not.** `docs/clause-readers.md`
says of itself that it "is a measurement, not a specification. Where it and the documents disagree,
the documents win." That is what puts this change on this side of the first standing rule and the two
wordings on the other. The line is worth stating because every sentence in the changed passage is
*about* the specification, and being about it is not being it.

**The document's ageing is visible rather than repaired, and that binds the figures.** Its own
preamble says a reader named by symbol "is wrong the day the symbol is renamed with nothing to say
so, which is the most a document can do about its own ageing". So the 2026-10-05 figures stay and the
2026-10-06 ones are added beside them with their method, rather than silently replacing them. Between
the two dates the trail grew from 55 P1 artifacts to 64, so the difference is not only method and a
reader has to be able to tell which part is which.

**A90 constrains what the rows may be made to look like.** A reader that cannot fail is worse than
none, so a row cannot be upgraded by a decision. What a decision changes is the explanation under the
table, not the reader column in it.

**Two figures from the same trail measure different populations.** 64 P1 artifacts and 64 P4
artifacts are the same intents, and a completeness check is judged at P4. Quoting the P1 figure as the
cost understates it, which is what the passage being replaced does; quoting both and saying which one
is the cost is the only honest form.

**The 88-column rule applies outside tables and code blocks.** The passage carries a table, so the
check has to distinguish the two — it is run over the document and the only lines above 88 are table
rows that were already there.

**The convention about a paragraph changed a second time applies to the whole passage.** It has been
amended before, by #212 and again by #255, which is how it came to carry three paragraphs and a
section number nobody re-read. It is replaced whole and read back as prose.
