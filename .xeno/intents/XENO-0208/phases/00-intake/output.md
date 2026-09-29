---
intent: github.com/triplem/xeno#136
phase: 00-intake
created: "2026-09-29T18:06:16Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c05acae.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d1f6357474aa8bb20c44931ce90eee495f2ef428699cefcf032df94667214e59
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

XENO-0207 added column headings in lower case and its design gave a reason:

> Lower case, because they label columns rather than begin sentences, and this project's
headings name > their section in words.

The first clause is fine and the second is borrowed from somewhere it does not apply.
This project's convention about headings is about prose: a heading names its section in
words, in issues as much as in files, because a leading number indexes a list the reader
cannot see. That is a rule about documents.

A column label in a terminal table is not a heading in that sense. It is a label on
tabular output, and every tool a reader of this one already uses prints those in upper
case: `ps`, `docker ps`, `kubectl get`, `gh pr list`. The convention exists and this
listing is the only table in the project that departs from it.

So the case was decided by analogy to the wrong thing, and the analogy was recorded as
though it were a reason.

<!-- xeno:section:scope -->
## Scope

`CREATED  INTENT  STATE  PHASE` for the listing, `PHASE  STATE  VERDICT` for the
one-intent form.

Nothing else. The words, the columns, the widths, the order, the default and the notice
are XENO-0207's and stay as they are.

Not a rewrite of XENO-0207's phases. Its design carries a verdict, and section 11 says
what is sealed is never rewritten; the decision is reversed here and the reversal names
it.

Not this project's convention about prose headings, which is unchanged and was never the
thing in question.

Not the case of anything else the tool prints. Verdicts, states and findings are lower
case and are sentences or values rather than labels.

<!-- xeno:section:context-rationale -->
## Why this context

**The reversal is recorded, not hidden.** XENO-0207's design is sealed and its reason
was wrong in a specific way: it cited a convention about prose headings to decide the
case of a table label. Naming that is the whole content of this intent, and it is worth
more than the three words it changes, because the same borrowing is available every time
this project has a rule about one kind of writing and a question about another.

**Upper case because a reader already knows it.** Not because it is prettier: because
`ps` and everything descended from it print `PID TTY TIME CMD`, and a reader scanning a
terminal recognises a row of capitals as the line that is not data. That is the function
of the convention and this listing had opted out of it alone.

**Only the case changes.** Same words, same widths, same order, same default, same
notice. The alignment is already guaranteed by the format string the heading shares with
its rows, and upper case does not change a width, so what XENO-0207 proved stays proved.

**A new intent for three words, and the cost is the point.** This will carry the same
seventeen sections and the same twelve hundred lines of record as a change to the gate
path did, which is exactly the measurement #117 is collecting: the record's size is flat
and the work's size is not. It is the clearest data point this project has produced on
that question, and the figures go to #117 rather than into an argument here.

**Section 11's rule is what makes it a new intent.** What is sealed is never rewritten.
Editing XENO-0207's design to say the opposite of what it was judged saying would leave
a verdict over a file that no longer says it, and the alternative — re-running that
intent's phases from P2 — would rewrite four verdicts to change three words.
