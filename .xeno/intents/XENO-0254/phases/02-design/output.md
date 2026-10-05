---
intent: github.com/triplem/xeno#254
phase: 02-design
created: "2026-10-05T18:57:38Z"
schema_version: "1.0"
runner_version: dev+e471bbb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: eb64b2b1ed0909c4c278dc6df1ac182def62d59b76b6b18243d5fea675196fe6
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

Two rows rather than one amended. Section 8's sentence has parts with different readers, and a
single row would have to name G-Schema, G-Questions and a writer in one cell and leave the
reader to work out which reads which. The audit's value is that the second column can be read
at a glance; a cell listing three readers for three requirements destroys that more quietly
than a stale row does.

The rows keep section 8's order — options and their consequence, then the recommendation —
rather than ordering by how well read each is. The section is the authority and a reader
checking this document against it should find the same sequence.

The first row names G-Schema first and G-Questions second, because that is the order they fire
in and because the first is the surprise. A reader who knows the gate table expects
G-Questions, which runs from P5; the shape check runs from P0 through `phaseResult`, and
putting it first is the whole correction.

The second row's reader column says the writer and says no gate, in those terms. "`QuestionAsked`"
alone would read as a gate to anyone scanning a column full of gate names, which is the specific
misreading this row exists to prevent.

The amendment is described as a correction and not as an addition. The document already
explains one late row from #202's own miss and two from #212; a fourth paragraph saying "and
these two replaced a row that an intent made wrong" is a different claim about the pass's
accuracy, and collapsing it into the others would hide that the document has been wrong rather
than incomplete.

The count moves 37 to 38 and nothing else moves. One row becomes two; no clause is added,
removed or reclassified, so the other three kinds keep their counts and the four-kinds framing
is untouched.

No register row. What this records is a correction to a measurement, which the measurement now
carries itself; A90 already holds the general fact that this document has no reader, and the
convention that would have prevented the staleness is #212's learning and belongs to the rule
set rather than to the register.

<!-- xeno:section:alternatives -->
## Alternatives

Amending the one row in place was the obvious option and is rejected on the reader column. It
would have to carry G-Schema, G-Questions and a writer for three requirements of unequal
coverage, and the one thing this document does well is that a reader can scan that column; a
cell that needs parsing is worse than a cell that is wrong, because a wrong one can at least be
noticed.

Leaving the row and adding a note under the table was considered. It keeps the pass's rows
exactly as made, which has something to be said for it in a dated measurement — but the row
would still assert a reader that does not read it, and somebody scanning for coverage of
section 8 would find the wrong answer in the place the document is organised around.

Marking the row stale rather than correcting it, as `~~strikethrough~~` does for a void
assumption in the register, was considered. That convention exists for a row whose *claim* was
superseded; here the claim was never true of the reader, so striking it through would record
history that did not happen.

Correcting every row the last four intents touched was considered and is rejected for #202's
reason, which this document states about itself: a pass that fixed as it went would stop at the
first finding. This corrects the one an intent is known to have made wrong, and P4's gaps says
plainly that nobody has looked at the rest.

Adding the convention from #212's learning to `CLAUDE.md` in this commit was considered and
rejected under section 10: a learning routes through a merge request against the rule set after
review, not through the intent that happens to be nearby. It stays a learning.

Checking the document mechanically was considered, briefly. A tool that verified "this symbol
exists and this gate calls it" would catch a renamed reader and not a wrong one, which is the
failure here; A90's finding is that a reader which cannot fail in the interesting case is worse
than none.

<!-- xeno:section:impact -->
## Impact

One file, one row replaced by two, one count and one paragraph. `git diff --stat` names
`docs/clause-readers.md` and nothing else.

Nothing executable is affected and nothing could be. No gate, rule, template or Go file reads
this document; it is prose about prose. The suite and the 399 verdicts are untouched, and
running them proves only that this change is what it appears to be.

What a reader gains is a true answer to the question the document exists for. Section 8's three
requirements are now covered by two rows whose reader column is accurate: the shape by two
gates, the recommendation by a writer and by nothing that judges. Before, one row claimed
G-Questions for a third of the clause, which is the kind of error A90 names as worse than a
gap — reported coverage that is not there.

What it also records, in the only place that can, is #229's asymmetry. The recommendation is
enforced when a question is written and not when one is read, so a green verdict says nothing
about it. That fact had three homes until now — two function comments and a call site — all of
them in code that a person auditing clause coverage would not open.

The honest limit is that this corrects one row and licenses nothing about the others. The pass
is dated 2026-10-03 and four intents have since changed what reads a clause; two of them added
their rows, one did not, and whether any other row aged the same way is unknown because nobody
has re-read them. P4's gaps says so rather than leaving the corrected count to imply accuracy.

The second limit is that the document is still read by nobody but a person, and is now longer.
Each correction makes it more useful and more expensive to keep true, and #212's learning —
that the intent changing a reader fixes the row — is the only thing that would hold it, with
nothing enforcing it.
