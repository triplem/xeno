---
intent: github.com/triplem/xeno#195
phase: 05-review
created: "2026-10-03T14:24:20Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+b54626e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 64c4bbcd6e7f1edce1647fff04e78309c703269b55f87751aac097a74bf1f25f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
  - rule: deviations-are-traceable
    result: met
    note: >-
      Two: ASSUMPTIONS.md's "Built" sentence amended rather than appended to, because a list that
      names the last of something and then grows is one a reader stops trusting; and
      `checkLearningArgs` made a method for no reason beyond matching its neighbours.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      A command is added and nothing is taken away; a hand-written record still passes, which thirty
      intents' worth of sealed records depend on. What changes for a contributor is the skill, whose
      YAML block is gone and whose place the invocation takes.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Two, each naming what it departs from: `ASSUMPTIONS.md`'s
"Built" sentence amended rather than appended to, because a list that names the last of something and
then grows is one a reader stops trusting; and `checkLearningArgs` made a method for no reason beyond
matching its neighbours, recorded because a reviewer looking for the reason would otherwise find
none.

**`interface-change-needs-a-migration-note` — met.** A command is added and nothing is taken away. A
hand-written record still passes every gate, which is a criterion and is what thirty intents' worth
of sealed records depend on. What changes for a contributor is the skill: the YAML block is gone and
the invocation is in its place, so the thing that used to be documented is no longer.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added.

**Beyond the three rules.**

*Is the thing actually fixed?* Yes, and the trail is the evidence rather than the test. A phase
directory of this intent shows `output.md`, `digest.md`, `gate.yaml` and `learning.yaml` all carrying
`0.1.0-dev+b54626e.dirty`, which no phase directory in this repository could show this morning.

*Is the split right?* The content is the agent's and the header is the runner's, which is `section
set`'s split reused and not a new one. Four of the seven alternatives were refused for leaving a
field with the wrong author, and the question that separated them was always who authors this field
rather than what is the shortest change — which is the design learning.

*Does anything about the gate change?* No, and it should not have. G-Learning was right and was never
the gap; the writer refuses the same closed set earlier, both reading
`model.LearningCategories`, which is two readers of one definition rather than two definitions.

*What does this leave that a reader should not mistake for closed?* `plugin_version` is still a
constant and now the only field in the header that is wrong, and it is wrong in all four artifacts
together rather than in one of them — which is #177 and is arguably more visible now that the other
three agree. Nothing makes the agent call the command, because G-Learning checks that the header
fields are present and not that they are true. And `--no-finding` cannot be told from a phase nobody
thought about.

*Is this the last of them?* For section 5's fields, yes: `intent.yaml` had no writer and now has one,
`learning.yaml` was the other and now has one, and what remains is a constant rather than an absence.
The four writerless things A35 once listed are down to none.

*Could this change a verdict behind it?* No. No gate, no rule, no existing artifact; `gate verify`
is 271 at exit 0.

<!-- xeno:section:release-notes -->
## Release notes

**`xeno learning record` writes the record section 10 owes.** It was the last artifact of this
process that no command wrote.

    xeno learning record --intent KEY --phase NN \
        --category <template|prompt|context-rule|project-convention> \
        --observation "what happened, with what it cost" \
        --proposal "what should change because of it" \
        --target <the file or directory it is about>

The four keys are yours and the header is the runner's — the same six fields it writes into
`output.md`, `digest.md` and `gate.yaml`, so all four artifacts of a phase now say which build
produced them. A record typed by hand said `0.1.0-dev` where those three said the commit.

**A second call appends**, so a phase that learned two things says so twice rather than retyping the
first.

**The empty record is said rather than left out:**

    xeno learning record --intent KEY --phase NN --no-finding

**Without `--phase` the record is the intent's own**, which is what `xeno intent close` reads.

**A category outside section 10's four is refused before anything reaches the file**, as is an entry
missing one of its keys, and as is either statement contradicting the other on a record that already
exists.

**Do not write the file by hand.** The learning skill no longer shows its YAML.

<!-- xeno:section:residual-risk -->
## Residual risk

**Nothing makes the agent call it.** G-Learning checks that the header fields are present, not that
they are true, so a hand-written record carrying `0.1.0-dev` still passes. The skill says to use the
command and a skill instructs. This is the same shape as `--tool-version` before the variable
existed, and the same answer: cooperation until something reads the value. A gate that compared a
record's `runner_version` against the `gate.yaml` beside it would close it, and would also turn every
one of the thirty sealed records red, which is A74's argument for why it must not.

**`plugin_version` is now the only wrong field in the header, and it is wrong in all four at once.**
Making the other three agree has made this one more conspicuous rather than less, which is an
improvement in visibility and no change in correctness. #177.

**Two provenances for one field, undistinguished.** Thirty intents' records carry a typed header and
the ones from here carry a stamped one, with nothing in either saying which. The commit is the
boundary. Third time today, and the answer is the same: the hashes are sealed and backfilling would
describe a write that never happened.

**`--no-finding` is the cheapest way to satisfy the gate.** One flag, and section 10 calls it
honest, so nothing distinguishes a phase that produced nothing from one whose author did not look.
Unchanged by this intent, which only moves where the flag is typed.

**The skill's YAML block is gone, so a reader who knew the file no longer finds its shape
documented.** Deliberate — a skill documenting both would let the reader pick the one that was
wrong — but somebody debugging a hand-written record now has section 10 and the gate's messages
rather than the skill.

**What is not a risk.** Any verdict behind this intent, and any record already written.
