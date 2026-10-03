---
name: xeno-learning
description: Write the learning record a Xeno phase owes, and the one an intent owes when it closes. Use at the end of every phase before `xeno phase finish`, and before `xeno intent close`. Not a phase of its own.
---

# Learning

Every phase owes a learning record, and so does an intent that closes. It is not a
seventh phase: learning happens at the end of each of the six and once more at the
close, which is why it is one skill rather than an instruction repeated six times.

## The command

One per observation, before `xeno phase finish`:

    xeno learning record --intent KEY --phase NN \
        --category <template|prompt|context-rule|project-convention> \
        --observation "what happened, in the phase, with what it cost" \
        --proposal "what should be done differently, so somebody can act on it" \
        --target <the file or directory the proposal is about>

A second call appends, so a phase that learned two things says so twice rather than
rewriting the first.

Where the phase genuinely produced nothing, say it rather than leave the file out:

    xeno learning record --intent KEY --phase NN --no-finding

And without `--phase`, the record is the intent's own, which is what `xeno intent close`
reads:

    xeno learning record --intent KEY --no-finding

**Do not write the file by hand.** The four keys are yours and the header is the
runner's — it is the same six fields the runner writes into `output.md`, `digest.md` and
`gate.yaml`, and a typed one records `0.1.0-dev` where those three record the commit the
binary came from. That was the last artifact of this process nobody wrote (#195).

A category outside the four is refused before anything reaches the file, as is an entry
missing one of its keys, as is `--no-finding` on a record that already carries an
observation.

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
