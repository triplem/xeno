---
intent: github.com/triplem/xeno#247
phase: 05-review
created: "2026-10-05T19:47:26Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4307f6085a6208c4e0d448f83fd16fc38b1bc1a7a09a3d030c12757cc144332a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Three, and the third is the costly one. Against P2 wording: the CLAUDE.md paragraph is nine lines in a file whose last instruction is to keep it short, and a three-line version was written first and discarded for saying what to do without saying why. A correction inside P3: the sentence under the audit table first repeated figures the paragraph below already carried, replaced whole and read back. And P4 and P5 were started over — the first P4 declared its test report while a background go test was still writing the file, so the hash bound a partial run; the suite finished, the file grew, and G-Evidence reported the mismatch with the exact cause. DeclareEvidence refuses a second declaration of the same kind and job, so removing both phases was the honest route rather than reconstructing a truncated file or approving a finding about evidence that cannot be trusted. P4 results and its learning carry the method.'
      result: deviation
      rule: deviations-are-traceable
    - note: No interface changes. A row in a document nothing reads, a paragraph in a file read by a model rather than a program, and a rule file nothing loads — rules.Load walks .xeno/plugin/rules/ and the project config, not examples/. No code, no command, no field, no artifact shape, and no change to the effective rule set, which an unchanged rules_hash across every verdict is the evidence for. The nearest thing to a dependant is a project choosing to adopt the example, and its header says how and what it buys.
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: 'None added and no code changed: go.mod is untouched. One rule file is added and is deliberately a dependency of nothing — under examples/ it binds nobody, which is A72 reasoning and the choice documentation-follows-the-change.yaml made. The thing that would need something new is a reader for any of the three clauses, and all three need a normative commit rather than a dependency, which is #258.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. **The first is the whole shape of this intent**: all three clauses asked
for a mechanism needing a normative commit, the agent may not write one, and what is done instead is
recording what reads each clause today. No normative document is touched. Nothing is invented — no
field, gate, tool or rule in the shipped set; the example rule sits under `examples/` where
`rules.Load` cannot reach it, which an unchanged `rules_hash` across every verdict proves rather
than argues. The branch carries one intent, the three issues each carry their own label, and the
footer takes a keyword each.

The acceptance criteria. Ten met, one met by the commit this phase precedes.

The non-goals held. No field for the reason, no clause for sequence, no numbering convention — all
three a person's commit, all three tracked by #258. Nothing enabled anywhere, including here. No
gate for any of the three. No second pass over the audit. The mapping row's reader column still says
`nothing`.

What a reviewer should check is one prohibition and one count. The mapping row must not name the
example — an unenabled rule fails nothing, and naming it would be the error #254 corrected one row
away in this same table — and `git diff` is how that is read, since the row does not appear in it.
And the count must be 39 against 39 rows, because a count drifting from its table is how this
document starts lying in the other direction.

**What these close on, and what they do not.** Each closes on what reads its clause today. None
closes on whether the mechanism should exist, and because closing them would have taken that
question with them, #258 was filed before this phase rather than left to a commit message. That was
P4's gap on the first attempt and it is addressed.

**This phase and P4 were started over**, and P4's results carries why: the first attempt bound a
test report's hash while a background suite was still writing the file, G-Evidence caught it, and
`DeclareEvidence` refuses a second declaration of the same pair — so removing both phases was the
honest route rather than reconstructing a truncated file or approving a finding about untrustworthy
evidence. The escape used is the one #225 made explicit a few intents ago, which is the first time
this session's own work has been needed by a later intent.

What this cost against what it changed: three files, one of them nine lines, behind seventeen
sections — one intent rather than three, which is this session's proportionality measurement applied
to itself and argued in P2 rather than assumed.

<!-- xeno:section:release-notes -->
## Release notes

Three clauses that had no reader now have the one they can honestly have: a person, with what they
read written where they will meet it.

**Section 8's reason has nowhere to be written, and the audit says so.** A new row:

| § | clause | reader |
|---|---|---|
| 8 | the recommendation carries a reason | **no field to carry it**; nothing reads it (#247) |

`model.Option` carries `Text`, `Consequence`, `Recommended` and `Free`, so a reason can only live
inside the question's text or inside a consequence, indistinguishable from what it shares the field
with. "No field" comes first because it is why no gate can fix it. The count moves 38 to 39.

**`CLAUDE.md` now says a decision put to a person is put one at a time.** With section 8's own
reason — an open question moves the whole of the thinking onto somebody, and a batch does that while
satisfying the letter of it — and with XENO-0243 as the instance: it asked three at once, and the
later two assumed an answer to the first that nobody had given. The free entry has the same problem,
since a person writing their own option changes what the next question should ask and a batch has
already asked it. Nothing checks this, and the paragraph says so.

**G-Test's mapping half has a person-reader available.**
`examples/rules/mapping-is-complete.yaml` asks that every acceptance criterion the requirements
phase declared is answered in the verification phase's mapping. It is enabled nowhere; adopting it
means copying it into `.xeno/config/rules/given/org/` or `given/project/`.

Under `examples/` and not `given/builtin/` on A72's reasoning: a review rule in the shipped set
reaches every project everywhere, including every project whose acceptance criteria are prose, and
would be answered not-applicable with a note at every review for ever.

**The audit's mapping row still says `nothing`**, deliberately. A rule nobody has enabled fails
nothing, and the reader column means the thing that would fail if the clause were violated — the
distinction #254 corrected one row away. The example is named in the prose beneath instead, with
what adopting it buys: a reader once per intent rather than once per criterion, which is the ceiling
until a criterion is identifiable.

Nothing executable changes. `rules.Load` walks `.xeno/plugin/rules/` and the project config, so
nothing under `examples/` can reach the effective set, and `rules_hash` is unchanged in every
verdict.

**What is still open, and now tracked.** All three issues close on what reads their clause today.
None closes on whether a mechanism should exist — a field for the reason, a clause for sequence, an
identifiable criterion — and each needs a commit to a normative document. **#258** carries all
three, so the question does not leave with the issues that raised it.

<!-- xeno:section:residual-risk -->
## Residual risk

Three issues close and the mechanisms they asked for do not exist. That is the intended outcome and
also how a question gets lost: a reader finding #247 closed will not know a field is still wanted
unless they follow to #258. Filing that before the review phase is the mitigation, and it is a
document pointing at a document — the strongest thing available, weaker than the issues staying open.
Whoever disagrees should reopen the three; nothing here depends on them being closed.

The `CLAUDE.md` paragraph is the only change with a mechanism, and the mechanism is a model reading a
file. A session that batches three questions breaks the convention with no finding, no refusal and no
verdict. It is also nine lines in a file whose last instruction is to keep it short: a few more
additions of this size and the instruction stops being true and the file stops being read closely,
which is slower than a wrong rule and harder to notice.

The example rule has never been loaded, which is what makes it safe and means a malformed field
would sit unreported until a project copies it in. The four beside it have been in that position
since A72. Whether a person can usefully answer it once per intent is also unknown — an intent with
eleven criteria compresses a real reading task into one checklist entry, and the honest answer may
always be `deviation` with a note, which would make it worse than no rule. Only adoption shows
either.

The audit now has a row whose reader column names a missing field rather than a reader. More useful
than `nothing`, and a precedent for rows describing why a clause cannot be read; a table of those
would be a different document from the one #202 made. Worth watching rather than acting on.

**The evidence mistake this intent made can recur and nothing prevents it.**
`evidence declare --file` hashes what is on disk at that moment and cannot know a writer is still
running, so a declaration over a growing file is indistinguishable at declare time from one over a
finished file. The gate catches it — immediately, and with the exact cause — two commands after the
act, and the cost was this phase and P4. The learning records the method; a guard would have to know
about a process it cannot see, and is not obviously possible.

What is not a risk: the verdicts that existed before this intent, confirmed intact at exit 0; the
effective rule set, unchanged and proved by `rules_hash`; and anything executable, since no code
changed.
