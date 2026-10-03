---
name: xeno-requirements
description: Run P1, the requirements phase of a Xeno intent — turn the intake's scope into acceptance criteria, non-goals and constraints. Use after 00-intake is green, or when `xeno intent status` shows 01-requirements as the next phase.
---

# Requirements

The second phase. It turns the intake's scope into sentences that can be checked, before
anything is designed. A criterion that cannot be checked is a hope, and this is the
phase where that is caught cheaply.

## What the phase owes

Three sections, from the `requirements` template:

- `acceptance-criteria` — each one a state of the tree and a verdict over it. Write the
  case the specification forbids as its own criterion: a resolution is implemented from
  the pattern, a refusal only if something asks for it by name.
- `non-goals` — what this intent will not do, each with its reason. Scope that is merely
  unstated gets built by accident.
- `constraints` — what the work has to hold: the sections of the specification that
  bind, the conventions of the project, the dependency rule, the gates that must stay
  green.

## The commands

    xeno phase start   --intent KEY --phase 01
    xeno section set   acceptance-criteria --intent KEY --phase 01 --file PATH
    xeno section set   non-goals           --intent KEY --phase 01 --file PATH
    xeno section set   constraints         --intent KEY --phase 01 --file PATH
    xeno phase finish  --intent KEY --phase 01 --summary PATH

Where the phase needs something to be true that nobody has confirmed, record it rather
than assume it:

    xeno assumption record --intent KEY --phase 01 --text TEXT --origin WHERE --confidence HOW

An open assumption blocks the phase, which is the point: it is released by a person with
`xeno assumption confirm` or `xeno assumption reject`.

The one field the runner cannot know is `tool_version`, because section 7 forbids it
branching on the harness. G-Schema requires it, so report it on the first `section set`
of the phase and the digest takes it from the artifact rather than asking again:

    --tool-version 2.1.276

## What the gate refuses

Everything the intake's gate refuses, plus:

- **G-Assumptions** — an assumption recorded and neither confirmed nor rejected.
- **G-Freshness** — a predecessor that changed after this phase started, which means the
  intake moved under it.

## What it hands on

The criteria. P2 designs against them and P4 verifies against them, one by one. A
criterion nobody can map to a check in P4 was written loosely here.
