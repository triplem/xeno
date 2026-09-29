---
intent: github.com/triplem/xeno#134
phase: 00-intake
created: "2026-09-29T17:56:32Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+1dc1483.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d4454ecf44a40c1c32fee77b03f5559c72d52b7d5df03882b30c5ffba63dabea
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

`xeno intent status` without an intent prints sixty unlabelled rows.

Four columns: a date, a key, a state, and a phase with its verdict. The first two
explain themselves. The third is a word this project defines in the runner rather than
in section 5, so a reader who wonders what `complete` means has nothing to look it up
by, and the fourth is two values in one column with no sign that it is.

The one-intent form has the same absence, and it is read more often: three columns, a
phase, a state and a verdict, where `finished` and `green` sit side by side and mean
different kinds of thing.

Sixty rows is also the wrong default. Fifty of them stopped at the intake and will not
move again, which XENO-0206 made visible: the listing now says `00-intake` for each of
them rather than hiding it behind `in-progress`. What a reader wants on an ordinary day
is the work that is recent.

<!-- xeno:section:scope -->
## Scope

Headings on both forms of the command.

The listing shows the last ten intents by creation, and `--all` shows every one. The
order stays ascending, which #118 settled, so the ten are the tail and the newest row is
the last printed.

A truncated listing says how many it left out and names the flag that shows them.

Not a change to the order, to the columns, or to what any of them contain.

Not paging, filtering or a count of its own. `--all` is one flag and the default is one
number.

Not a heading on anything else the tool prints. `gate run` and `phase finish` print a
verdict and its findings, which are sentences rather than a table.

<!-- xeno:section:context-rationale -->
## Why this context

**Ten, because the tail is what a listing is for.** Fifty intents here stopped at the
intake and are not coming back; the default should answer what is happening rather than
what has ever happened. Ten is a screen, and the flag is there for the other question.

**Ascending order stays, so the default is a tail and not a reversed head.** #118
settled that the listing reads in the order the work happened, and reversing it for the
default would mean the two forms of one command disagreed about time. The last ten of an
ascending list is the newest ten with the newest last, which is what `git log --reverse`
does and what a reader of this listing has already learned.

**A truncated listing that says nothing is worse than a long one.** Ten rows and silence
reads as a repository with ten intents. That is a wrong fact rather than a missing one,
and it is the error this change would introduce if the count were left out, so the count
and the flag are printed together.

**Headings are the cheapest fix for the column this project invented.** `complete` is
defined in the runner and appears in no document, which XENO-0206 recorded as a small
debt. A heading does not repay it, and it does tell a reader that the column is a state
rather than a verdict, which is the confusion available in a row that carries both.

**The one-intent form gets them too, although the issue was about the other.** It is the
form every phase of every intent here has been read with, and it puts `finished` next to
`green`: one is a position and the other is a judgement. Labelling them costs one line
and removes a question.

**Nothing about the data changes.** No new field, no new order, no new computation. This
is a presentation change and the design says so, because a presentation change that
quietly altered what was shown would be the worse kind.
