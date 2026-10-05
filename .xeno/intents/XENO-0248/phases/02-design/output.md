---
intent: github.com/triplem/xeno#243
phase: 02-design
created: "2026-10-05T12:36:00Z"
schema_version: "1.0"
runner_version: dev+f2f65d5
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e52792b6bf25da183dd2271cbbe1698da64c1df633b9c4e9741ec4b93aae73f1
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

The register is open, and the opening says so in its first line rather than arriving at it.
The banner's first line was its claim, and a correction that buried the new position under
the history of the old one would be read the same way the banner was: by its first sentence.

The two roles are separated by name. One paragraph says what closed at M0, quoting section
4's "a file in the branch rather than an artifact under `.xeno/`", and says that was the
substitute for a phase's artifacts. Another says what this file is and that it continues.
Naming both is the whole correction: #174's error was available only because the file had
two roles and one name.

The misreading is recorded rather than quietly repaired. The opening says the record was
read as closed and names #174 and the seventeen rows, because a reader who has seen the
banner in the history needs to know which reading won and why, and because a file that
corrects itself silently is a file that can drift again without anyone noticing it had.

The banner's second half is kept almost verbatim. A learning routes through `learning.yaml`
and a merge request against the rule set, per section 10, and practice has observed that
without exception. It was never the part in dispute and rewording it would invite the
question of whether it changed.

The distinction a row must meet is "outlives the intent that found it", not a significance
threshold. A threshold would need a number or a list and would be the invented rule the
second standing rule forbids; durability is a property of the thing itself and is the
property the seventeen rows actually share — A74 about rule sets, A86 about G-Supply's
anchor, A94 about sealed references, each read by intents that did not write them.

A95 records the correction, and the row is in the register rather than only in this intent's
P2 by its own new standard: where a construction decision belongs is a fact every later
intent reads, so it outlives this one. The row is the first use of the rule it establishes.

The fourth paragraph is corrected rather than deleted. It explains the state column, what
`accepted until` means and what an `open` row waits on, all of which stay true; only its
framing sentence, that this is the record kept before M0, is replaced.

<!-- xeno:section:alternatives -->
## Alternatives

Keeping the register closed and recording the seventeen rows as a drift was the issue's other
option and was rejected by the maintainer on this evidence. It asks twelve commits' worth of
judgement to have been wrong when the sentence they were judged against turns out not to say
what the banner claimed; and "every row stays" forecloses deleting them, so the file would
keep seventeen rows it says should not be there. It also re-accepts the unfindability the
banner itself called the cost of the switch, which is the thing that drove the twelve.

Removing the banner and saying nothing was the third option and is rejected for throwing away
the half that was right. The learning rule is correct and observed; a register with no
statement of what it is for is how this one drifted into having two roles and one name.

Rewriting the plan's section 4 was never on the table and is worth saying so. It would make
the banner retrospectively correct, which is backwards: the plan's sentence is about
artifacts and is right about them, and the first standing rule puts it beyond the agent in any
case.

A significance threshold for a row was considered and rejected. "A choice worth recording" or
a list of kinds would be a rule nobody could apply the same way twice, and the second standing
rule makes an invented rule a specification change first. Durability is the test because it is
checkable: either a later intent needs the fact or it does not.

Marking A78 through A94 as rows written while the banner said otherwise was considered and
rejected. The correction's claim is that they were right, so annotating them would record the
opposite, and it would put seventeen edits into a commit whose point is one paragraph.

A gate over the register was considered and rejected for A90's reason. Nothing can compare a
repository against "a row here is for what outlives the intent that found it", and a checker
that could would be checking something else. The opening joins `docs/clause-readers.md`'s list
of rules whose reader is a person, knowingly, which is where the width rules went for the same
reason in #210.

<!-- xeno:section:impact -->
## Impact

One file, one paragraph replaced, one sentence of another corrected, one row added. No code,
no command, no gate, no rule, no artifact shape, and nothing under `.xeno/intents/` but this
intent's own directory.

Nothing executable is affected, and nothing can be: the register is not covered by
`artifacts_hash`, no gate reads it, and nothing opens it by path. The 363 verdicts stand, the
build and the suite are indifferent, and the five workflows do not know the file exists. That
is the same property that let the banner drift for seventeen rows, and it is why this
correction is prose and why the gaps section will say what still has no reader.

What changes for a reader is the first sentence they meet. Today it says the record is closed
and the file goes on for sixty-nine rows past the one it names as last; after this it says the
register is open, what for, and what did close. The second thing it says is now the plan's own
sentence, so the inference #174 made is visible as one.

What changes for an agent is that `CLAUDE.md`'s pointer becomes true. It says assumptions and
decisions taken while building the core are in this file, which the banner contradicted; the
expected outcome of criterion 10 is that `CLAUDE.md` needs no edit, and if it does need one
that is a finding about this design rather than a step.

What changes for the next intent is that A95 exists and says where a construction decision
goes. The seventeen rows had no stated home and were written against a banner; the
eighteenth has one, which is the whole point of the change and the only part a later reader
will care about.

The cost is that the register still has no reader but a person, and now says more that nobody
checks. Two sentences of it are testable in principle — whether a row outlives its intent,
whether a learning ended up here — and neither is tested. The file grows its claim surface
while keeping its enforcement at zero, which A90 would call the shape of the problem it
documents, and this intent accepts knowingly rather than quietly.
