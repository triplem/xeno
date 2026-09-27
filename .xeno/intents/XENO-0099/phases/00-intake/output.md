---
intent: github.com/triplem/xeno#99
phase: 00-intake
created: "2026-09-27T12:39:28Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+915578b
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 74fc27e3c738d4d9dcff9f1dc2520486f1e90ff9a65c6b7c7c26763e90f9f3d8
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

Eleven of 121 functions were over forty lines, and two of them were the shape the issue
names: `enforcement.Compare` at 89 lines with four near identical switches, and
`cmd/xeno/main.go:run` at 225 with a flag set, two switches listing their exceptions and
a twenty way dispatch.

The cost was not reading them once. It was that a branch of `Compare` could only be
reached by building a whole report and reading one line out of it, and that adding a
flag to one command in `run` meant finding where a case ended among two hundred lines.

<!-- xeno:section:scope -->
## Scope

Five refactorings with no behaviour change: `run` to a command table, `Compare` to a
requirement table, `Fetch` into transport and decode, `schema` into the four things it
judges, and A12's predecessor decision out of `Start`. Plus one duplication removed,
`oneOf`, which existed in the gates and in the runner.

Not `Init` and `Attach`, which are long because they are sequences and read top to
bottom. Not a length limit either: that waits for #94's tool decision, and a limit set
before the refactorings would have been raised rather than met.

<!-- xeno:section:context-rationale -->
## Why this context

**Measured before and after, because a review of this kind is otherwise taste.** Eleven
functions over forty lines became six, and the function count went from 121 to 147: the
same code in smaller pieces. `run` went from 225 lines to 39.

**The two switches that listed exceptions were the actual defect.** `run` decided
whether a command needs `--intent` and whether it needs a phase by naming the exceptions
in two places, so a new command could be forgotten from either, and #58's `--export` was
added by reading all 225 lines to find where a case ended. Both are a field on the
command now, read once.

**A table is worth it where the branches repeat and not otherwise.** `Compare`'s four
blocks had one shape written four times, so the table removes duplication rather than
adding indirection. `Init` and `Attach` are equally long and are sequences: a table
there would name functions after line ranges.

**The splits were chosen for what they make testable.** `Fetch` into transport and
decode buys two tests that could not exist before, the status mapping without a body and
the body without a server, and it is also #97's step 1, which is the part that does not
depend on how that port is eventually shaped.

**The duplication was mine, twice.** `oneOf` was written into the gates for the learning
categories and into the runner for the assumption sets, in different weeks, four lines
each time. It is `model.OneOf` now, beside the enumerations it tests, and section 10's
categories and Appendix B's hash fields moved there with it: they are the
specification's, not one gate's.

**What a refactoring may not do is change a test.** Every one of these kept the suite as
it was, and the only additions are the two tests the `Fetch` split exists for. The exit
code staircase was checked by hand against A11 afterwards, because it is behaviour no
test asserts end to end.