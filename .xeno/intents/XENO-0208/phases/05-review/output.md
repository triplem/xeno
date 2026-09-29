---
intent: github.com/triplem/xeno#136
phase: 05-review
created: "2026-09-29T18:10:38Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4b87535.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bd3b84de73ddf17db5daaa8fd704fc278f0e5e34a1e54be64f551ba20161b587
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

**Nothing under `docs/` changed**, and nothing in `CLAUDE.md` either, which the gaps
section records as the place the rule still is not.

**Nothing sealed was rewritten.** XENO-0207's phases are not in the diff. Its design
says lower case and carries a verdict; this intent reverses the decision and names it,
which is what section 11 asks for. Three ways of avoiding a second intent were
considered and each rewrites something a verdict was taken over — including folding it
into the open pull request, which felt different and is the same act.

**Only the case moved.** Two literals, one comment, four strings in a test, and nothing
outside `cmd/`. The widths are XENO-0207's because the columns were sized for the data,
and the transcript prints the heading beside the widest row from the same format string
for a reader to check.

**The alignment property is unchanged and still asserted.** Same two tests, same widest
rows, new literals.

**The convention is named in the code.** `ps` prints `PID TTY TIME CMD`; a row of
capitals is how a reader recognises the line that is not data. And the comment says that
this project's rule about a heading naming its section in words is about prose, which is
the citation XENO-0207 borrowed.

**Every verdict still matches**, 127.

**The phases were written in order**, and the intent is honest about its own ratio:
seventeen sections for two literals, which goes to #117 as a figure rather than into an
argument here.

<!-- xeno:section:release-notes -->
## Release notes

The column headings of `xeno intent status` are upper case: `CREATED  INTENT  STATE
PHASE` for the listing and `PHASE  STATE  VERDICT` for the one-intent form.

That is the convention for a label on tabular output, which `ps` established and
everything descended from it follows, and it is how a reader scanning a terminal
recognises the line that is not data. This project's rule that a heading names its
section in words is about prose, in files and in issues; XENO-0207 cited it for a column
label and set these in lower case.

Nothing else changed. The same words, the same widths, the same order, the same default
of ten, the same notice, and the same values in every row.

<!-- xeno:section:residual-risk -->
## Residual risk

**The rule lives in a comment, not where conventions are kept.** `CLAUDE.md` says a
heading names its section in words and does not distinguish prose from tabular output,
so the next table will face the same question with the same citation available. The
answer is now in `cmd/xeno`, which is the wrong place for a convention and the right
place for a decision about one table. Moving it is a change to `CLAUDE.md` and therefore
not this intent's.

**Nothing tests the case**, because a test for it would assert a literal against a
literal. The transcript carries the output and reading it is the verification, as it was
for the notice's wording one intent ago. Two literals in this repository are now load
bearing for a convention and unguarded.

**The sealed reason stays wrong and stays readable.** XENO-0207's design says lower case
with a citation, and nothing in that file says it was reversed. Section 11 asks for
exactly this: the reference runs from the new record to the old one and not back. A
reader who finds the old design first will read a decision that no longer holds.

**Three intents are now queued on one another.** #133, #135 and this one, each built on
the last, because each was asked for while the previous was still open. They merge in
order and nothing depends on them being merged together, but a reader of any one of them
sees the others' commits until the earlier one lands.

**Seventeen sections for two literals**, which is thirty times the diff. Accepted as the
cost of section 11's rule and recorded in #117, where the question of whether that cost
is always right is being decided on figures rather than on this instance.
