---
intent: github.com/triplem/xeno#235
phase: 01-requirements
created: "2026-10-06T15:21:29Z"
schema_version: "1.0"
runner_version: dev+8decfdc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3e0f1be21d0ef4d659a011a839af6efb87e6df627526ce2694506bccf499e915
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

1. **The `gate.yaml` findings block carries `advisory`.** In section 5's yaml block, between `next`
   and `decision`, with the comment convention the block already uses: `# absent in the ordinary
   case`, aligned at the column the two comments beside it use.

2. **A paragraph says what an advisory finding is and what it does not do.** The check carrying it
   is `pass`, the phase is `green`, and the finding is in `gate.yaml` with its id, its cause and its
   remedy like any other. Those three are what make it "visible" and "decidable" in section 5's own
   words, and each is stated rather than implied.

3. **The same paragraph says why the exception exists.** Blocking against a number nobody has
   experience with would be the wrong way round, and a check that fires with nothing to do about it
   is one people learn to route around — which is A44's reasoning and A90's, generalised to the
   clause rather than left in an assumption row.

4. **A second paragraph bounds it.** A finding is the thing that fails; an advisory one is readable
   only because it is rare; the clause that asks for one says so where its check is described, and
   nothing else writes the field. Without that, the field is available to any check that finds
   itself inconvenient.

5. **The paragraph is placed after the status derivation is finished, not inside it.** Section 5
   explains the phase status over two paragraphs, and the second — "The four values are not
   interchangeable" — completes the first. The new paragraph follows both.

6. **The Context economy subsection says the budget's finding is advisory.** One sentence, in the
   paragraph that carries the clause the code contradicts, so that a reader arriving at the budget
   check is told there rather than inferring it from a yaml block forty lines earlier.

7. **No fifth check result and no change to the four.** `pass`, `fail`, `pending`,
   `not-implemented` are untouched, so A4 and A42 are not reopened and no reader of a check result
   learns a new state.

8. **`drift` and section 16's ninth limitation are untouched.** The route was weighed and rejected;
   recording that in this intent's artifacts is the whole of what it gets.

9. **The commit carries no code.** No `Advisory` on `model.Finding`, no change to `result`, no
   change to `budget`, no tests. `git diff --stat` names `docs/process-definition.md` and nothing
   else outside the trail.

10. **`gate verify` exits 0, `go test ./...` passes, `go vet` is clean, `gofmt -l` prints nothing.**
    No Go source is touched and the specification is in no hash, so no verdict in the trail moves.

11. **No new line over 88 columns outside tables and code blocks.** Measured before and after as
    multisets of lengths, because the document has pre-existing long lines.

12. **The agent's edit of a normative document is recorded where a reader meets it.** P0's problem
    section, the design decision and the commit message, as XENO-0262 did for #258's amendments.

<!-- xeno:section:non-goals -->
## Non goals

**Not `Advisory` on `model.Finding`.** Not the field, not `result` choosing `fail` only on a
non-advisory finding, not `budget` marking its own, not the tests. The first standing rule puts the
code after the specification commit, and in a separate intent, so that the two are distinguishable
in a history that squashes.

**Not implementing `drift`.** It stays unimplemented and section 16's ninth limitation stays unread.
Implementing it here would be building a specified-but-absent mechanism for its second purpose
before its first, and the row's shape does not fit a byte count anyway. It is weighed in the
artifacts and not touched in the document.

**Not a fifth check result.** `advisory` is a property of a finding and not of a check. Putting it
on the check would make a future second advisory finding in G-Schema drag the whole check with it,
and `not-implemented` shows how far a non-verdict state ramifies — section 5 had to say what a phase
carrying one means.

**Not changing section 5's budget clause itself.** The sentence that says the overrun is "deliberately
a finding and not a red gate" is correct and is the thing the code violates. It gains a sentence
beside it and loses none.

**Not `docs/clause-readers.md`.** Section 5's budget clause has no row there and does not get one
here. The row is worth adding when a reader exists; a document that is a measurement of what fails
should not acquire an entry on the day a clause becomes sayable.

**Not #267.** `budget` sums `f.Bytes` over `lock.Files` and no P0 lock in this trail records one, so
the check reads nothing at the one phase that declares a budget. An advisory finding that is never
produced is still never produced. Separate issue, deliberately not merged into this one.

**Not #258's two amendments.** They merged before this branch began. Same document, same session,
different issues, different work packages, and the third standing rule gives each its own branch and
its own intent.

**Not a general advisory mechanism for other checks.** The amendment names the field and bounds it to
clauses that ask for it. Which other checks should be advisory is a question per check, and answering
it in the abstract is how a bounded exception becomes a default.

<!-- xeno:section:constraints -->
## Constraints

**"Not a red gate in the sense of stopping work" has to be written against a behaviour.** In this
runner stopping work is `predecessorAllowsStart` refusing the next phase on an undecided failure. So
the paragraph says `pass` and `green` rather than paraphrasing section 5's own phrase, because a
clause that restates the thing being clarified clarifies nothing.

**The field must be a property of the finding and not of the check.** Section 5 enumerates the four
results and A4 and A42 fixed them. The amendment has to add to `findings` and leave `result` alone,
or it reopens a decision this one was taken to avoid.

**The exception has to be bounded in the clause and not in the code.** A90: a reader that cannot fail
is worse than none, and a field that silences a finding is a way to make every inconvenient check
unable to fail. Only the code could enforce the bound and no code will; so the bound is a sentence,
and the sentence has to be explicit that it is one.

**Placement is constrained by a two-paragraph argument.** The status derivation runs over "Decisions
sit on findings, not on phases" and "The four values are not interchangeable", and the second
completes the first. Inserting between them would split it, which is what the drafted wording's
"after 'Decisions sit on findings, not on phases'" would literally have done.

**The yaml block has a comment convention.** Both existing comments start at column 40 and both say
when the key is absent. A new key that aligned differently or explained itself differently would read
as having arrived from somewhere else.

**The specification cites no repository and no issue.** Checked over the document: zero matches for
an issue number, an intent key or "this trail", against four in `CLAUDE.md`. So the clause carries
its reasoning in general terms, and A44's and A90's names stay out of it even though they are where
the reasoning came from.

**88 columns outside tables and code blocks**, and the document has pre-existing over-long lines, so
the measure compares multisets before and after rather than counting.

**One commit, and the reason is in it.** Every merge here is a squash, which XENO-0262 found while
asking the same question; the first standing rule's substance is that the specification change does
not arrive mixed with its code, and that is what the commit has to show.
