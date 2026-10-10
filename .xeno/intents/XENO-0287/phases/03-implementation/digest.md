---
intent: github.com/triplem/xeno#336
phase: 03-implementation
created: "2026-10-10T13:42:21Z"
schema_version: "1.0"
runner_version: dev+5f645cb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e487a47233690e0cc961143851c9d1e0e896128554446538de4022ad3f03d7cf
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
One file, 23 lines added, two hunks, and `git diff -U0 main` confirms they are at 4.1 and
at Sources and nowhere else.

The clause the paragraph turns on is the one that came from the tree rather than from the
issue: the receiver is a hosted automation "and not anything in the repository — whose tree
carries no workflows at all." Section 12 says "something has to receive them", and that
clause is the sentence where it stops being a figure of speech and becomes an observation
about a specific tree.

One deviation, against my own criterion rather than against the issue. Criterion 9 said one
sentence and the paragraph is two. Written as one, the four claims came to about ninety
words with three nested subordinate clauses, and the #330 contrast — the clause whose whole
job is a distinction a reader can get wrong — arrived last, after two others had to be held.
The design's reason for not quoting section 12 was that the pointer is load-bearing; the
same reasoning says the contrast has to be legible, and it was not. What the issue's "one
sentence" was for is preserved: not a section, not a table, not a list.

The learning generalises it. A criterion that fixes the form of prose rather than a property
of it should be drafted before it is sealed, or written about what the prose must not become:
a sentence count is a proxy for "not a section", and the proxy can fail while the thing it
stands for holds.

`git diff main -- docs/process-definition.md` is 0 lines. `gofmt`, `go vet` and
`xeno gate verify` at exit 0 over 600 verdicts, and no line over 88 outside tables.

Two sections, no open question, no decision.
