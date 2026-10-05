---
intent: github.com/triplem/xeno#254
phase: 01-requirements
created: "2026-10-05T18:56:53Z"
schema_version: "1.0"
runner_version: dev+e471bbb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 23e78308f334a98a8abbb922ff36cb2db2d4126488c2f8d2257837ba3e21ad58
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. Line 84's row is replaced by two, following section 8's own order: the options with their
   consequence and the free entry, then the recommendation.

2. The first row names G-Schema as well as G-Questions, because the shape check is reached
   through `phaseResult` and so runs from P0. The document defines a reader as the thing that
   would fail if the clause were violated, and that is G-Schema.

3. The second row says the recommendation's reader is the writer and that no gate reads it,
   plainly enough that somebody reading a green verdict does not take it as covered.

4. The tool-requirement count moves from 37 to 38, and the rows in the table count 38.

5. The document's account of its own late rows covers these two: it already explains one row
   that arrived after the pass and two more from #212, and this adds a correction rather than
   an addition, which is a different thing and is said so.

6. No clause text is invented. Both rows quote what section 8 asks, which is read from the
   section rather than from the existing row.

7. No code changes. Nothing about what reads these clauses is altered, and `git diff --stat`
   names one file.

8. Markdown prose stays within 88 columns; the table is exempt, as `CLAUDE.md` says and as
   this document's own widest rows already are.

9. `./xeno gate verify` exits 0 with the 399 verdicts that exist now intact, plus this intent's
   own judged phases.

10. `go build`, `go test ./...`, `go vet ./...` pass and `gofmt -l` outside `vendor/` prints
    nothing. None of them can be affected by this change, and running them is what confirms
    that rather than the reasoning.

11. One commit, `Closes #254`, and the issue carries `wp5`.

<!-- xeno:section:non-goals -->
## Non goals

Not a re-audit. The pass is dated 2026-10-03 and #202's reasoning stands: a re-reading that
fixed as it went would spend itself on the first finding. This corrects the row an intent made
wrong and looks no further, and P4's gaps will say that nobody has looked.

Not the recommendation's reason. It has no field and therefore no reader, which is #247; these
rows describe what is read, and a requirement with nowhere to be written is a different kind of
gap.

Not the sequence rule. It is not a clause yet, which is #248.

Not a change to any code. What reads these clauses is exactly what read them before this
intent; the document said otherwise and now does not.

Not a change to the row's four kinds, the counts of the other three, or the document's framing.
Only the tool-requirement count moves, and only because one row becomes two.

Not a gate over the document. Nothing will compare the repository against it, which is A90's
own finding about itself and is why the convention #212 proposed is a convention rather than a
check.

Not the gitignore correction. #201's finding was false, the entry exists from #200, and the
correction belongs in the issue and on that pull request because there is no change to make.

<!-- xeno:section:constraints -->
## Constraints

The document is a dated measurement and must go on saying so. It records a pass made on
2026-10-03 against `d19a1ca` and already explains three rows that arrived later; a correction
is a fourth kind of amendment and has to be distinguishable from the three, or the file starts
claiming a pass it did not make.

The reader column has a definition and it is not "the gate this clause belongs to". It is the
thing that would fail if the clause were violated, which for a question's shape is G-Schema
from P0 and not G-Questions from P5. Getting that wrong is what made the row misreport
coverage rather than merely age.

The second row describes an absence and that is the point. No gate reads the recommendation and
nothing in a verdict can say so, so the row's value is entirely in being read by a person; it
has to be worded so that somebody skimming the reader column does not take `QuestionAsked` for
a gate.

`CLAUDE.md` exempts tables from the 88-column rule and this document's rows are already wider,
so the constraint binds the prose and not the two rows.

Nothing here can break a test, which is a constraint on what the verification phase may claim
rather than a freedom: the suite passing says nothing about whether the rows are right, and
P4 has to say that rather than count a green suite as evidence.

One intent, one branch, `Closes #254`, and the issue carries `wp5` because #229 changed the
reader and #212's learning puts the row in the changing intent's package rather than the
document's.
