---
intent: github.com/triplem/xeno#267
phase: 01-requirements
created: "2026-10-06T19:56:56Z"
schema_version: "1.0"
runner_version: dev+1d61fa2
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7ef8ab81da7d67d07dee7b8d0622bf81a40c89370e130463728856d88d9aac84
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Numbered, and P4's mapping cites the numbers.

1. **Section 5 says, where the lock is described, that at P0 `files` is empty.** A reader who has
   just read "The file is written before the agent starts, from the context scope" learns in the
   same place that at the intake there is no scope yet to read.

2. **It says why the order is forced**, in terms of the two commands: `scope set` writes into the
   phase directory, so it cannot run before the phase has started, and the lock is therefore
   written before the thing it resolves.

3. **It says why nothing fills the lock afterwards** — a lock written a second time is a sealed
   artifact rewritten and the answer to "what changed" destroyed, which is #215's reason stated
   once more where it has this consequence.

4. **It names the two checks and says they are judged from P1 on**, and says that each of them
   says so where it is described, so a reader following either direction arrives at the other.

5. **It says what P0 still binds.** The scope is one artifact per intent and every later phase
   resolves the same one, so the declaration is not decoration even though the phase that makes it
   is not measured against it.

6. **The schema block's `files` line carries a comment pointing at that paragraph.** Somebody
   reading the shape of the artifact rather than the prose around it is who needs the warning, and
   the block is where they are.

7. **Section 5's budget clause says the check is judged from P1 on**, naming the comparison it
   makes — the lock of the phase being gated against the one budget the intent declares — so that
   the exclusion follows from the mechanism rather than being asserted beside it.

8. **Section 7's staleness clause says P0 stands outside the check**, and says this is not a third
   limit chosen against noise but a consequence of when the lock is written. The two limits that
   are about noise keep their reasons and stay two.

9. **It says what the arrangement loses**: a file that moved while the intake itself ran, since the
   next phase's lock records it as it was by then and that is the state the comparison starts from.

10. **`informationBase`'s comment no longer says an empty list is a phase older than #217's
    rule.** It says which phases that holds for and which phase is always empty.

11. **`staleReads`'s comment no longer says it either**, in the same terms, since it carries the
    same sentence today.

12. **`docs/clause-readers.md` carries the three clauses** with what reads each, which for all
    three is a reader of the document and no code, since nothing is being checked.

13. **No behaviour changes.** `go test ./...` passes without a new or amended assertion, because
    the three paragraphs describe the runner as it already is.

14. **The specification change is its own commit, before the one that corrects the comments.**
    `CLAUDE.md`'s first standing rule, and the exception under which the agent typed it is named
    in that commit's message.

<!-- xeno:section:non-goals -->
## Non goals

**P0's lock does not start recording files.** That is the decision #267 resolved and not a thing
postponed inside this intent. The two candidates that would record them are written down in the
issue with what each costs, and either stays available if the figures later argue for it.

**No new check result and no new finding at P0.** A `budget` that reported "not applicable here"
would need a result outside the four A4 and A42 fix, and a finding saying a check did not apply
would fire on every intake forever. #235 is the issue about a check that did not run, and
XENO-0265 closed the half of it that had a shape.

**No change to the two limits of the staleness half.** They are about noise, they have their
reasons in section 7 already, and the paragraph being added is careful to say it is not a third
one. Renumbering "Two limits" to three would absorb a consequence into a design choice.

**No test.** Nothing changes behaviour. The measurement in #267 was taken over the trail with a
shell loop and is already in the issue; turning it into an assertion would be a test that P0's
lock is empty, which is a fact about today's order and not a requirement anybody wants held.

**No new row in `docs/assumptions.md`.** Nothing here is an assumption taken while building. The
one open judgement was put to a person and answered, and a decision is recorded at P2 where
decisions are recorded.

**`ChangedSince` is not changed and gets no clause.** It prints the predecessor's `files`, so at
P1 it prints nothing, which is the same cause and a third place to say it. Section 5's
re-reading-follows-change sentence is about the mechanism and stays as it is; the paragraph where
the lock is described covers the one artifact all three readers share.

<!-- xeno:section:constraints -->
## Constraints

**The first standing rule.** `docs/process-definition.md` is not editable by the agent. The
wording of all three paragraphs was drafted and put to the maintainer before anything was typed
into the file, and the instruction to write them as drafted is what this intent runs under. The
exception is named in the specification commit's message, because a trace is what stops it
becoming a habit.

**The specification commit comes first**, before the commit that corrects the two comments, since
those follow from it.

**What is sealed is never rewritten.** Nothing in this intent touches a verdict or an artifact of
another intent, so the three paragraphs cannot move a figure anywhere in the trail. That is also
why the two wrong comments are corrected rather than the artifacts that quote them: XENO-0265's P0
states the 0-of-117 measurement and stays as it is.

**Prose wraps at 88 characters**, tables and code blocks do not, and a paragraph being changed for
the second time is replaced rather than edited into. The budget clause under context economy is in
that state — XENO-0263 rewrote it for the advisory sentence — so the addition goes beside it as
its own paragraph and the existing one is read back whole.

**Headings name their section in words.** The three paragraphs are bold lead-ins in the register
section 5 and section 7 already use, not numbered subsections.

**No invented fields, gates, tools or rules.** Nothing is added to an artifact. The change is
three paragraphs of prose, one YAML comment in a non-normative example block, two code comments
and three index rows.

**One dependency.** Untouched; nothing here compiles differently.
