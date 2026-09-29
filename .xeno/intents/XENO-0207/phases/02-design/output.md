---
intent: github.com/triplem/xeno#134
phase: 02-design
created: "2026-09-29T17:57:36Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+1dc1483.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fdaac67b555c4924789e60ae439db18ebb58402cd20f4a62672112e6f5a2e41e
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

**One format string per table, used for the heading and for every row.** That is what
makes AC7 hold by construction rather than by inspection: a heading cannot drift from a
column it shares a width with.

**The headings are words, not field names.** `created`, `intent`, `state`, `phase` for
the listing; `phase`, `state`, `verdict` for the one-intent form. Lower case, because
they label columns rather than begin sentences, and this project's headings name their
section in words.

**Ten is a constant in `cmd/xeno` with its reason beside it.** Not configurable, not in
`project.yaml`, not a `--limit`. The number answers "what is happening" and the flag
answers "what has ever happened", and a third knob would be a third question.

**The tail is taken after the sort, in the command.** `Intents()` returns everything in
order and the command shows a slice of it. The runner keeps knowing nothing about
presentation, which is what makes AC8 true by construction: the data cannot change
because nothing about the data was touched.

**The notice names the count and the flag, and prints only when something was left
out.** One line after the table, so that it is the last thing read and cannot be
mistaken for a row.

**A repository with no intents prints nothing.** Not a heading over an empty table,
which would be a label for an absence. The heading is printed with the first row rather
than before the loop.

<!-- xeno:section:alternatives -->
## Alternatives

**Truncate in `Intents()` and pass a limit into the runner.** It would keep the command
simpler and make the runner answer a question about presentation. It also means the one
function both forms share would return different data depending on how it was called,
which is the thing AC8 exists to prevent.

**Reverse the order for the default, newest first.** What most listings do, and it would
put the row a reader wants at the top. It would also make the two forms of one command
disagree about time, and #118 chose ascending on the argument that the record reads in
the order the work happened. A tail is the same information without contradicting that.

**Print the heading before the loop.** One line shorter. It labels an empty table in a
repository with no intents, which is a heading for nothing.

**A `--limit N` instead of `--all`.** More flexible and it asks the reader to choose a
number they have no basis for. `--all` is the only other question anybody has: this, or
everything.

**Repeat the heading every twenty rows.** Useful in a long listing and a table of ten
does not need it. It would also mean a heading could appear between two rows, which is
where a reader least expects a label.

**Leave the one-intent form alone**, since the issue was about the other. Its columns
are `finished` and `green` side by side, a position and a judgement, and it is the form
every phase here has been read with. One line each was the wrong place to economise.

<!-- xeno:section:impact -->
## Impact

`cmd/xeno`: one format string and a heading in each of the two printers, the `--all`
flag, the tail, and the notice.

Tests: the headings present in both forms, the default showing ten of more, `--all`
showing all, the notice appearing only when something was left out, a repository with no
intents printing nothing, and the alignment at the widest state.

`internal/runner` is not touched, which is the point: the data cannot have changed.

What a reader gains: a table they can read without knowing the tool, and a default that
answers what is happening rather than what has ever happened.

What they lose: the sixty-row view by default, which is one flag away and was mostly
fifty intents that stopped at the intake.

What this costs: a number in the source that somebody will eventually want to configure,
and a notice line that will look like noise to whoever has fewer than ten intents and
never sees it.
