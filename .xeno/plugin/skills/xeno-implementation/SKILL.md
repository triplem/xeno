---
name: xeno-implementation
description: Run P3, the implementation phase of a Xeno intent — record what was changed and every deviation from the design. Use after 02-design is green, or when `xeno intent status` shows 03-implementation as the next phase.
---

# Implementation

The fourth phase. The code is written here and the artifact records what the code is and
where it left the design. The phase that most often finds the design was wrong.

## What the phase owes

Two sections, from the `implementation` template:

- `changes` — what was changed, file by file or concern by concern, with the reason
  where the reason is not obvious from the diff. Counts and paths belong here; a reader
  who has the tree and nothing else is the audience.
- `deviations` — every departure from the design, each naming what it departed from: a
  decision of this intent, an assumption, or an acceptance criterion. A deviation
  discovered while implementing is the most valuable thing this phase produces, and the
  one most often left out.

Where implementing answered a question or produced a decision the design did not take,
record it here rather than editing the design: what is sealed is never rewritten.

## The commands

    xeno phase start   --intent KEY --phase 03
    xeno section set   changes    --intent KEY --phase 03 --file PATH
    xeno section set   deviations --intent KEY --phase 03 --file PATH
    xeno phase finish  --intent KEY --phase 03 --summary PATH

Evidence a pipeline produces is declared here and attached later:

    xeno evidence attach --intent KEY --phase 03 --from DIR

The one field the runner cannot know is `tool_version`, because section 7 forbids it
branching on the harness. G-Schema requires it, so report it on the first `section set`
of the phase and the digest takes it from the artifact rather than asking again:

    --tool-version 2.1.276

## What the gate refuses

Everything the earlier gates refuse, plus:

- **G-Evidence** — a declared item whose content changed, or an attachment with no hash.
- **G-Build** — a declared build result that is missing, failed or stale.
- **G-Policy** — a checked rule applying to this phase that does not hold.

## What it hands on

The changes and the deviations. P4 maps every acceptance criterion to the check that
proves it, and a criterion with no check is either untested or was written so that
nothing could test it.
