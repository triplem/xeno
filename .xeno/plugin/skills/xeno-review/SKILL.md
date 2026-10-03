---
name: xeno-review
description: Run P5, the review phase of a Xeno intent — answer the review rules, write the release notes and the residual risk. Use after 04-verification is green, or when `xeno intent status` shows 05-review as the next phase.
---

# Review

The last phase. It answers the rules a gate cannot evaluate, says what the change is for
a reader, and states what is still wrong. After it, the merge.

## What the phase owes

Three sections, from the `review` template, and one structured list:

- `review-checklist` — the section, carrying the reasoning behind each answer and the
  questions the change raises that no rule asks.
- `release-notes` — what changed, for somebody who did not follow the work. The shipped
  rule set requires this section to carry text, which is the one mechanical check on a
  review.
- `residual-risk` — what is still wrong, in order of weight, including what only a
  person could fix and what belongs to the specification rather than to the code.

And in the frontmatter, `review_checklist`: one entry per review rule of the effective
rule set, each with `rule`, a `result` of `met`, `deviation` or `not-applicable`, and a
`note` where the result is not `met`. An entry from a lens carries `source: lens` and no
rule id, and answers no rule.

Read the effective set before writing the entries; a rule with no entry is red.

## The commands

    xeno phase start   --intent KEY --phase 05
    xeno section set   review-checklist --intent KEY --phase 05 --file PATH
    xeno section set   release-notes    --intent KEY --phase 05 --file PATH
    xeno section set   residual-risk    --intent KEY --phase 05 --file PATH
    xeno phase finish  --intent KEY --phase 05 --summary PATH

Where a finding is accepted rather than fixed, a second person releases it and the
reason is recorded:

    xeno gate approve  FINDING --intent KEY --phase 05 --by WHO --reason TEXT
    xeno gate override FINDING --intent KEY --phase 05 --by WHO --reason TEXT

An override leaves an obligation, closed by `xeno obligation close` when the work is
done.

The one field the runner cannot know is `tool_version`, because section 7 forbids it
branching on the harness. G-Schema requires it, so report it on the first `section set`
of the phase and the digest takes it from the artifact rather than asking again:

    --tool-version 2.1.276

## What the gate refuses

Everything the earlier gates refuse, plus:

- **G-Questions** — any open question of any phase of this intent left unresolved.
- **G-Policy** — a review rule of the effective set with no checklist entry, an entry
  with no result, a result outside the three, a `deviation` or `not-applicable` with no
  note.
- **G-Complete** — a preceding phase missing, or red without a decision on its findings.

## What it hands on

The merge, which is not a command of this runner: commit, push, and let the review and
the pipeline run. The pipeline recomputes every verdict with `xeno gate verify` and
writes nothing.
