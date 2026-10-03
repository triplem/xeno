---
intent: github.com/triplem/xeno#188
phase: 04-verification
created: "2026-10-03T21:44:04Z"
schema_version: "1.0"
runner_version: dev+efc44a9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: db60e3a63cc56198c005d022a3b6bc2dbbe84d93ead5401f5741b3724fe1c72d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - chosen: Leave it on both forms. The listing keeps the row mark and the one-intent form keeps the line.
      decided_by: triplem
      id: D-2
      proposed_by: claude-opus-5
      rationale: |-
        The mark is uniform because the habit is, and the uniformity is the finding #188
        filed: nobody noticed zero questions in ninety-eight intents by looking at one of
        them. A figure moved out of the listing in the week it first reports something is a
        figure tuned to be quiet. It also thins out by itself, since a row loses the mark
        the moment its intent records an exchange, so the cost is paid for as long as the
        fact holds and no longer, and no code has to change for it to stop.
      resolves: Q-2
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Each acceptance criterion of P1 against what judges it.

**A question is recorded by a command, into the frontmatter of the running phase**
— `TestAQuestionIsWrittenWithTheKeyTheRunnerAssigns`, which asserts the entry
reaches the frontmatter as it was given, down to the free entry being the third
option. Also live: Q-1 was written by hand because the command did not exist yet,
and Q-2 is in P3's frontmatter because `xeno question record` put it there.

**It refuses, before writing, everything `questionShape` would find afterwards**
— `TestAQuestionWithoutEnoughOptionsIsRefusedBeforeTheWrite`, which asserts the
refusal and that the artifact did not move, `TestAnHonestNoOptionsIsAccepted` for
the exception section 8 names, and `TestAKeyOnTheWayInIsRefused`. The intent-wide
uniqueness is `TestKeysContinueAcrossThePhasesOfOneIntent`, which is the check that
widened from the file to the intent.

**A decision is recorded by a command, with the person on it** —
`TestADecisionNeedsItsPersonItsReasonAndItsOption`, a table of the five refusals,
which also asserts that none of the refused entries reached the artifact.
`TestAWithdrawalResolvesTheQuestionAsWell` covers section 8's third exit.

**Neither command invents a person** — the first row of that table, and
`TestTheExchangeIsRecordedFromTheCommandLine` at the surface, which drives `decision
record` without `--by` and asserts exit 1.

**A write after a verdict invalidates the verdict, as a section write does** — not
a test of its own, because it is the absence of a behaviour: neither command reads
`gate.yaml` and neither writes one. `TestNoCommandButFinishTouchesAFinishedPhase`
is the test in this package that holds the property for every command, and the two
new ones are inside it by being ordinary writers of `output.md`. Shown live as well,
in this intent's P3: the learning record changed the phase after its verdict and
`phase finish` said so.

**A question raised in one phase and decided in a later one is resolved for
G-Questions** — `TestAQuestionRecordedAndDecidedIsResolvedForTheGate`, which records
the question in P0 through the command, the decision in P1 through the command,
runs all six phases and asserts G-Questions passes and the figure agrees with the
trail. This is the loop #188 says has never once run.

**`xeno intent status` says when an intent reached P5 having asked nothing** —
`TestTheFigureNamesAnIntentThatReachedReviewAskingNothing`, which also asserts the
figure is silent at each of the five earlier phases. The printing is in the results
below, measured on this repository.

**It is a figure and not a finding** — no gate was added or changed, which `gate
verify` over the whole trail asserts, and the marked intents in the listing are green.

**The line appears on no other intent** — `TestOneEntryOfEitherKindAnswersTheFigure`,
`TestTheFigureLeavesAnAbandonedIntentAlone`, and the five earlier phases in the
test above.

**The figure is computed once** — `summarise` became `Summarise` and both forms
call it; there is one walk and one predicate. Asserted by construction rather than
by a test, and visible in the diff as the absence of a second count.

**Nothing else changes** — `./xeno gate verify` over 331 verdicts, the existing
suite, and no document in the tree touched.

**The usual gates of this repository** — below.

<!-- xeno:section:results -->
## Results

**`go build -o xeno ./cmd/xeno`** — builds.

**`go test ./...`** — eighteen packages, all `ok`, none skipped. Seventeen cases are new
here: fifteen in `internal/runner/exchange_test.go` and one at the surface in
`cmd/xeno/main_test.go`, plus the withdrawal case, which the two packages report
together.

**`gofmt -l .` outside `vendor/`** — nothing. **`go vet ./...`** — nothing.

**`./xeno gate verify`** — `verified 331 verdicts`, exit 0. No divergence, nothing red,
nothing provisional, and no existing artifact touched.

**The four refusals, run against this repository rather than a fixture.** Each exits 1,
says why, and writes nothing; the artifact carried one `decided_by` before them and one
after.

```
$ xeno decision record --intent XENO-0242 --phase 04 --chosen x --reason y
refused: --by is required: taking a decision is a statement by a person

$ xeno decision record --intent XENO-0242 --phase 04 --resolves Q-9 --chosen x --reason y --by triplem
refused: XENO-0242 has raised no question Q-9; --resolves names a question of this intent, in any phase

$ printf 'text: t\noptions:\n  - text: only one\n' | xeno question record --intent XENO-0242 --phase 04
refused: question Q-3 needs two to four options and one free entry, has 1 and 0; offer the options found, or state no_options: true where there are none

$ printf 'text: t\nno_options: true\n' | xeno question record --intent XENO-0242 --phase 05
refused: 05-review has no .xeno/intents/XENO-0242/phases/05-review/output.md yet; it is written by xeno section set, which creates the artifact
```

**The fifth refusal fired before it was asked for.** The first draft of Q-2 carried a
stray `consequence_note_unused` key and `question record` refused it: `field
consequence_note_unused not found in type model.Option`. That is the unknown-field
decode doing what it is for, found by making the mistake rather than by writing the test
for it.

**The loop ran live, which is the acceptance.** Q-2 was raised in P3 by `xeno question
record`, which printed `raises Q-2`, continuing past the hand written Q-1 of P0. D-2 was
recorded in P4 by `xeno decision record`, which printed `records D-2, decided by
triplem, resolving Q-2`. Both entries are in the frontmatter of the phase that holds
them, and the person's name is in the field rather than in a paragraph.

**The figure, on this repository.** `xeno intent status --all` marks 46 rows. Every
complete intent in the trail carries the mark and not one of them is missing it, which
is #188's measurement printed by the tool instead of counted by hand. XENO-0242 is
unmarked, because it asked twice.

**The trail's own counts, before and after.** `decided_by` as a field: none before this
intent, two now, in D-1 and D-2. Artifacts carrying a `decisions` block: none before,
two now. Question entries: four before, six now.

<!-- xeno:section:gaps -->
## Gaps

**A shape refusal names a key the artifact does not carry.** The third refusal above
says "question Q-3", and no Q-3 exists: the key is assigned before the shape is judged,
because `QuestionShape` reports a question by its key and an entry without one is a
different finding of the same gate. Nothing is lost — the number is not consumed, so the
next question recorded does become Q-3, and the line begins with "refused" — but a
reader could take it for a key that was written. Left as it is rather than fixed in
verification, and recorded here rather than absorbed.

**A question cannot be replaced before its phase is judged.** This phase's learning in
P3 records it: Q-2 was recorded, sent back for sharper options, and put again with a
fourth, and the artifact carries the three it was recorded with. The right habit is to
settle the question with the person before recording it, which costs nothing. A
`--replace` path is a change to the command and would have to refuse a judged phase,
which makes it an issue of its own rather than a line here.

**The command writes YAML at a different indent than the hand written blocks.**
`yaml.Marshal` indents a list by four and every block written by hand in this repository
is indented by two, so XENO-0242's P0 and P4 do not look alike. It is the tree's
existing behaviour and not this command's: `frontmatter` marshals every field it
carries, so the next section write of any phase would reindent a hand written block the
same way. Cosmetic, inside `artifacts_hash`, and so not worth a rewrite of anything
sealed.

**The printed figure has no test, by judgement.** Recorded as a deviation in P3 and
measured above instead. A fixture building a six-phase intent to assert a suffix would
assert the suffix.

**What #188 cannot close in the intent that builds it.** The figure reports on the
intents that follow. This one asked twice, so it is the sample of one that proves the
mechanism and not the evidence that the habit changed; that evidence is the mark
disappearing from rows over the next weeks, which is the shape of report #117 already
has.

**Still unwritten, and each already filed.** `evidence` has no writing command, which is
#208.  `review_checklist` is written by hand into P5's frontmatter, which is WP7's.
Nothing compares `proposed_by` with `decided_by`, which is Q-1's third option and needs
a document change first.
