---
intent: github.com/triplem/xeno#136
phase: 02-design
created: "2026-09-29T18:07:14Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c05acae.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d5a57d361b60ae67daf7d075060e27dd592ce9ad5a674105e0a7040cc4ee48c7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: by-hand
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The literals change and nothing else does.** Two calls, three words each, and the
format strings they are passed to are untouched, so the alignment property XENO-0207
established holds without being re-established.

**The comment above the format string says which convention applies.** XENO-0207's
reason was a citation of the wrong rule, and a comment naming the right one is what
stops the next reader restoring the lower case for the reason that was given the first
time. It names `ps` rather than arguing from taste.

**The two tests are updated, not duplicated.** They assert that each heading word sits
over a non-blank column of the widest row; the property is unchanged and only the
expected strings move. A second pair would test the same thing twice and both would have
to be kept in step.

**The words stay as they are, only their case moves.** `created`, `intent`, `state`,
`phase`, `verdict`. Renaming any of them is a different question and this intent has
one.

<!-- xeno:section:alternatives -->
## Alternatives

**Edit XENO-0207's design to say upper case.** One line and the record would then read
as though the decision had always been that. Its phase carries a verdict, so section 11
refuses it: what is sealed is never rewritten, and a verdict over a file that no longer
says what was judged is worse than a reversal recorded elsewhere.

**Re-run XENO-0207 from P2.** The other way to keep one intent. It rewrites four
verdicts to change three words, and G-Freshness would report the later phases stale
until their locks were rewritten too, so the cost is four phases of record for a change
of case.

**Fold it into #135 before it merges.** Tempting, since that pull request is open and
the commits are on a branch nobody has merged. It amounts to the same rewrite of sealed
phases, and the branch's verdicts were produced by the same gates that would have to be
re-run; "not merged yet" changes who would notice, not whether the record would match
what was judged.

**Upper case everywhere the tool prints a word.** Consistent, and it would shout
verdicts and states that are values rather than labels. `ps` prints `PID` and then a
number, not `12345` in capitals.

**Leave it lower case and record the convention in `CLAUDE.md` instead.** It would
settle the question for the next table and leave this one departing from what every
comparable tool does, which is the thing that prompted the issue.

<!-- xeno:section:impact -->
## Impact

`cmd/xeno`: two calls and a comment.

`cmd/xeno/list_test.go`: the expected strings in two tests.

Nothing else. No runner, no model, no format string, no width, no data.

What a reader gains: a table that looks like every other table in a terminal, so the
line that is not data is recognisable as such without reading it.

What this costs: a full intent's record for three words, which is the figure #117 is
collecting rather than a complaint. Seventeen sections and about twelve hundred lines
against a diff of two lines, against XENO-0202's four hundred and ninety-four and
XENO-0121's eighteen. The record's size is flat and this is the clearest instance of it
this project has produced.
