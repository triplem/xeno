---
name: xeno-learning
description: Write the learning record a Xeno phase owes, and the one an intent owes when it closes. Use at the end of every phase before `xeno phase finish`, and before `xeno intent close`. Not a phase of its own.
---

# Learning

Every phase owes a learning record, and so does an intent that closes. It is not a
seventh phase: learning happens at the end of each of the six and once more at the
close, which is why it is one skill rather than an instruction repeated six times.

## What a record is

`learning.yaml` in the phase directory, beside `output.md`. The header is the phase's
own identity; the body is a list of entries, or the honest empty record.

    intent: "<qualified id>"
    created: <iso8601>
    phase: <phase id>
    schema_version: "1.0"
    runner_version: <as the artifact records it>
    plugin_version: <as the artifact records it>
    learnings:
      - category: <template|prompt|context-rule|project-convention>
        observation: what happened, in the phase, with what it cost
        proposal: what should be done differently, stated so somebody can act on it
        target: the file or directory the proposal is about

`no_finding: true` instead of `learnings` is the empty record, and it is honest where a
phase genuinely produced nothing. An invented key, a category outside the four, or an
entry without a proposal is red.

## What makes an entry worth writing

An observation about the work, not about the subject. The phase found something it did
not expect; a convention cost more than it returned; an instruction was followed and
produced the wrong thing; a template asked for something the phase could not supply. If
the entry would be true of any project, it is probably about a rule rather than a
learning.

The entry nobody writes and everybody needs is the one where the process itself got in
the way.

## Where a learning goes afterwards

Nowhere automatically. Section 10 routes a learning through a merge request against the
rule set, so that it takes effect after review rather than on being noticed. A learning
is a proposal and a rule is a decision, and the gap between them is deliberate.

## At the close

An intent that ends, finished or abandoned, owes a record of its own:

    xeno intent close --intent KEY --reason TEXT

An abandoned intent cannot close without one, which is the one place the process insists
on learning from work that produced nothing else.
