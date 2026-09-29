---
intent: github.com/triplem/xeno#134
phase: 05-review
created: "2026-09-29T18:01:29Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c05acae.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 35efeeffa7d8eb77c790ffd091d16e6718168ec4cc3cd31f217346249fdb478f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**Nothing under `docs/` changed.** This is presentation and the documents describe
artifacts.

**Nothing about the data changed.** `internal/runner` is not in the diff, so the order,
the columns and the computation behind them are what XENO-0206 left. A presentation
change that quietly altered what was shown would be the worse kind, and the absence of
that package from the diff is the proof rather than a claim.

**The alignment cannot drift.** One format string per table, shared by the heading and
every row, and two tests that compare a heading against a row of the widest values each
column can hold.

**A truncated listing says what it left out.** `53 older, --all to see them`, after the
table, printed only when something was. Ten rows and silence would have been a wrong
fact rather than a missing one.

**A repository with no intents prints nothing**, not a heading over an empty table,
which is a label for an absence.

**The order is untouched.** Ascending, as #118 settled, so the default is the tail and
the two forms of one command still agree about time. Reversing it for the default was
considered and rejected on that.

**The one-intent form got headings although the issue was about the other**, because it
puts a position beside a judgement and it is the form every phase here has been read
with.

**Two helpers exist for the tests and the deviations say so.** `tail` and `sprintRow`,
three lines each, and without them the arithmetic and the alignment would have been
checked by capturing output and by reading.

**The phases were written in order.**

<!-- xeno:section:release-notes -->
## Release notes

`xeno intent status` prints column headings, in both of its forms. The listing names
`created`, `intent`, `state` and `phase`; the one-intent form names `phase`, `state` and
`verdict`, which is the pair most worth labelling, since one is a position and the other
a judgement.

The listing shows the last ten intents by creation. `--all` shows every one, in the same
order. The order is unchanged: ascending, so the newest row is the last printed and the
default is the tail of the record rather than a reversed head.

A truncated listing says how many it left out and how to see them, after the table. A
listing that shows everything says nothing extra, and a repository with no intents
prints nothing at all.

Nothing about the data changed: same order, same columns, same values.

<!-- xeno:section:residual-risk -->
## Residual risk

**Ten will be wrong for somebody.** It is a constant in `cmd/xeno` with its reason
beside it, and the first project with two hundred intents will want it configurable. The
place it would belong is section 12's `project.yaml`, which makes it a decision rather
than an edit, and nothing here prepares that.

**`older` means `created earlier` and those will diverge.** The order is over creation,
which #118 chose and XENO-0200 recorded the cost of: an intent created months ago and
worked yesterday sorts early, so the notice would count it among the older ones while it
holds the newest work. Inherited rather than introduced, and now printed in a word that
makes the claim sound stronger.

**Nothing tests the rendering.** The arithmetic and the alignment are asserted; that the
notice reads as it does, and that a date is cut to ten characters, are verified by a
transcript and by reading. That was the price of keeping the tests off standard output,
and it is the right price, but it means the wording of the one new line in the output
has no guard.

**`cmd/xeno` has three tests where it had none.** #110 is about the command surface
having no test behind its exit codes, and three tests about a table neither close it nor
pretend to. What they do is make the file exist.

**Accepted with the four named.** The state it replaces is sixty unlabelled rows, of
which fifty had stopped at the intake, above a column whose invented word had nothing to
identify it by.
