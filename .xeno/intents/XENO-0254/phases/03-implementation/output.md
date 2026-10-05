---
intent: github.com/triplem/xeno#254
phase: 03-implementation
created: "2026-10-05T19:00:24Z"
schema_version: "1.0"
runner_version: dev+e471bbb.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3f80a35639c7c87563dfbe209c6a1945d1239e7382245a6beed2ecefd457f13b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

One file, thirteen lines added and two removed.

`docs/clause-readers.md` line 84's row becomes two, in section 8's own order:

| § | clause | reader |
|---|---|---|
| 6 | a question carries two to four options, each with its consequence, and one free entry | G-Schema's shape check, G-Questions |
| 8 | a question recommends exactly one of its options | `QuestionAsked`, the writer only; no gate reads it |

G-Schema is named first because it is the surprise: a reader who knows the gate table expects
G-Questions, which runs from P5, and the shape check is reached through `phaseResult` and runs
from P0. That ordering is the correction rather than a presentational choice.

The second row's reader column says "the writer only; no gate reads it" rather than naming
`QuestionAsked` alone, because a bare symbol in a column of gate names reads as a gate. The
asymmetry #229 left deliberately now has a home outside two function comments and a call site.

The tool-requirement count moves 37 to 38, and the table counts 38 rows.

A paragraph above the table says these two are a correction and not a late addition, with why:
#229 gave two of section 8's requirements a reader and did not touch the row describing them,
and the pass "was not incomplete here; it was overtaken, and then wrong". The document already
explained one late row from #202's own miss and two from #212, and collapsing a correction into
those would hide that the table had been wrong rather than short.

Nothing else moves. No clause is added, removed or reclassified, the other three kinds keep
their counts, the four-kinds framing is untouched, and no code changes — `git diff --stat`
names one file.

The gitignore half of what prompted this is not here, because there was nothing to change: the
entry exists at `.gitignore:10` from #200, #201's finding was false, and the correction is in
#254 and in a comment on PR #246.

<!-- xeno:section:deviations -->
## Deviations from the design

No deviation from the design. The two rows, their order, the wording of the second reader
column, the correction paragraph and the count all landed as P2 specified them.

One correction inside the phase, caught by reading rather than by a tool. The correction
paragraph came out at 89 columns on two lines, one over the Markdown prose rule, and was
replaced whole and read back rather than rewrapped at the break — which is what the
conventions ask for a paragraph being changed a second time, and which is also how the
sentence "naming what each is read by" became "naming what reads each", a small improvement
that a rewrap at the break would not have produced.

One thing is worth recording about what this intent cost. It is thirteen lines of one document,
and it has seventeen sections and about 1,400 lines of record behind it, because the process
has one shape whatever the change. That is the figure #117 has now collected eight times, and
this intent is the clearest instance of it so far: the ratio is around a hundred to one. It is
not a complaint about the work — the row was wrong in a document whose whole value is being
right — but it is the case the plan's proportionality worry is actually about, and the ninth
data point for it.

Nothing else departs. No code, no normative document, no other row, and the gitignore half of
the prompting request turned out to need no change at all.
