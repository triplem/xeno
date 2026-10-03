---
name: xeno-design
description: Run P2, the design phase of a Xeno intent — record the decisions, the alternatives rejected and the impact, before writing code. Use after 01-requirements is green, or when `xeno intent status` shows 02-design as the next phase.
---

# Design

The third phase. It records what was decided and what was rejected, so that a reviewer
can overturn a decision cheaply rather than reconstruct it. The phase with the highest
ratio of reading to writing.

## What the phase owes

Three sections, from the `design` template:

- `decisions` — one paragraph per decision, each saying what it is and why it is that
  way. Where a decision follows a precedent in the tree, cite it: an argument from first
  principles for something already settled is the commonest waste in this phase.
- `alternatives` — what was not done, and what each rejection costs. Write out the
  reading the specification supports and this phase did not take: a decision recorded
  against its alternative is one a reviewer can overturn, where a decision recorded
  alone reads as settled.
- `impact` — what changes for a reader, for a project, for the trail. Include what the
  change does to artifacts that already exist, which is the question this phase forgets.

## The commands

    xeno phase start   --intent KEY --phase 02
    xeno section set   decisions    --intent KEY --phase 02 --file PATH
    xeno section set   alternatives --intent KEY --phase 02 --file PATH
    xeno section set   impact       --intent KEY --phase 02 --file PATH
    xeno phase finish  --intent KEY --phase 02 --summary PATH

Where the design settles a question an earlier phase raised, say so in a `decisions`
entry with `resolves:` naming the question's key: the gate reads it from any later
phase.

## What the gate refuses

Everything the earlier gates refuse, plus:

- **G-Questions**, from P5 — an open question raised here and never resolved by a
  decision, a confirmed assumption or a withdrawal.
- **G-Policy** — a checked rule of the effective set that applies to this phase and does
  not hold, and a rule whose predicate type nothing implements.

## What it hands on

The decisions. P3 implements them and records every departure as a deviation naming what
it departed from. A deviation that names nothing is a note, which the shipped rule set
asks about at review.
