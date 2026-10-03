---
name: xeno-intake
description: Run P0, the intake phase of a Xeno intent — state the problem, fix the scope, say why this context. Use when an intent is starting, when the user names an issue to work on, or when `xeno intent status` shows a phase of 00-intake not started.
---

# Intake

The first phase. It states what is wanted and what is deliberately not, before anything
is designed. An intent that skips it has no record of what it was for, and every later
phase is judged against this one.

## What the phase owes

Three sections, from the `intake` template:

- `problem` — what is wrong today, in the tree rather than in the abstract. Name the
  files, the fields, the gates or the documents that disagree. A problem nobody can
  locate is a wish.
- `scope` — what this intent does, and what it does not, with a reason for each
  exclusion. An exclusion with no reason reads as an oversight to the next reader.
- `context-rationale` — which files and which sections of the specification were read,
  and why those. This is what the gate compares the phase against later.

`open-questions` and `decisions` are optional and belong here where the intake raises
one.

## The commands

    xeno phase start   --intent KEY --phase 00
    xeno section set   problem           --intent KEY --phase 00 --file PATH
    xeno section set   scope             --intent KEY --phase 00 --file PATH
    xeno section set   context-rationale --intent KEY --phase 00 --file PATH
    xeno phase finish  --intent KEY --phase 00 --summary PATH

`phase start` writes `context.lock.yaml` and fixes what the phase was given. `section
set` renders the whole artifact again on every write, so the anchors are never yours to
type. `phase finish` writes the digest from the summary, filters it, and judges the
phase.

Write a `learning.yaml` in the phase directory before finishing: every phase owes one,
and `no_finding: true` is the honest empty record.

The one field the runner cannot know is `tool_version`, because section 7 forbids it
branching on the harness. G-Schema requires it, so report it on the first `section set`
of the phase and the digest takes it from the artifact rather than asking again:

    --tool-version 2.1.276

## What the gate refuses

- **G-Schema** — a missing required section, a missing frontmatter field, a hash that is
  neither a sha256 nor the placeholder.
- **G-Trace** — an artifact whose intent id does not match `intent.yaml`.
- **G-Learning** — a missing `learning.yaml`, an invented key, a category outside the
  four section 10 fixes, an entry without a proposal.
- **G-Assumptions** — an assumption recorded here and left open.
- **G-Freshness** — a context hash that does not match the lock the phase was given.

## What it hands on

The scope. P1 turns it into acceptance criteria and may not widen it; where P1 needs
something the intake excluded, that is a finding about the intake rather than a quiet
addition.
