---
intent: github.com/triplem/xeno#243
phase: 05-review
created: "2026-10-05T12:41:52Z"
schema_version: "1.0"
runner_version: dev+f2f65d5.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0fe55d166b7aa451d1da43a0c5f7c72dde5b3ef15f4671a6f6d3d987d3ec970e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'One, found by reading the new opening back as a paragraph rather than as a diff: the durability paragraph said the examples are what ''the rows above'' have in common, and the table follows the prose. It departs from P2''s design, which named the paragraph''s content and not its wording. One addition beyond the design is recorded with it: four paragraphs were planned and five written, splitting what section 4 says from what #174 made of it, because that distinction is what the correction rests on.'
      result: deviation
      rule: deviations-are-traceable
    - note: 'No interface changes. One file of prose: a banner replaced by five paragraphs, one framing sentence corrected, one row added. No command, field, gate, rule, template or artifact shape, and nothing reads the register, so there is nothing outside this intent that could depend on it. The nearest dependant is CLAUDE.md''s line 73, which names the file and is made true rather than broken.'
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: 'None added, and none could be: the change touches no code. go.mod is untouched. The one tool this correction would have benefited from is a reader for the register, and A95 and P4''s gaps both record that it has none and that A90''s finding is why one is not being added here.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. The first is engaged and held in the way that matters most here:
this intent's whole claim is about what the plan's section 4 means, and the plan is not
touched. The correction is that #174 misread it, not that it is wrong, so no specification
commit precedes this and the sentence is quoted rather than amended. No field, gate, tool or
rule is added — the register gains a row and loses a banner. The branch carries one intent,
the commit references #243, and the issue carries `wp5`.

The acceptance criteria. Eleven met, one met by the commit this phase precedes. Criterion 10's
expected outcome was no change to `CLAUDE.md` and that is what happened, recorded in P3 so a
reader can tell it was checked rather than skipped.

The non-goals held. No row deleted, renumbered or rewritten: 94 before, 95 after. No judgement
on the seventeen. No change to either normative document. No gate, no checker, no significance
threshold. No marking of the rows written while the banner said otherwise, because the
correction's claim is that they were right.

What a reviewer should check is the inference, not the prose. Everything here rests on reading
section 4's sentence as being about the stand-in for phase artifacts rather than about this
register. The two quoted phrases are at `docs/implementation-plan.md` lines 1441 and 1443 and
the paragraph around them is about how a package was worked while there was no runner. If that
reading is wrong then the banner was right and seventeen rows are a drift, which is the issue's
other option and was weighed in P2.

What this intent got wrong is in P3's deviations: the durability paragraph said "the rows
above" for rows that are below. Found by reading the result back as a paragraph, which is what
the conventions ask for and what the design had already committed to.

What stays unguarded is in P4's gaps, and the register's own A95 says it too: the file now
makes four claims and nothing checks any of them.

<!-- xeno:section:release-notes -->
## Release notes

`docs/assumptions.md` said the record was closed at A77 and ran to A94. It is now stated as
open, and says what it is for.

The register holds assumptions and decisions about how this repository is built. A row belongs
in it when the fact outlives the intent that found it; a decision taken inside one intent and
sealed with it belongs to that intent's design phase. A learning never comes here at all — it
routes through the phase's `learning.yaml` and a merge request against the rule set, as section
10 describes, which is the half of the old banner that was right.

What actually closed at M0 is now distinguished from this file by name. The implementation
plan's section 4 describes the record kept while Xeno could not govern its own construction and
says "the record is a file in the branch rather than an artifact under `.xeno/`", and that from
M0 "the hand held record stops". That record was the stand-in for a phase's artifacts and it
did stop: every change since runs six phases and writes real ones. The register is not that
stand-in, and neither normative document closes it.

So this is a correction rather than a reversal. #174 applied a sentence about artifacts to a
file that is not one. Seventeen rows were added afterwards across twelve commits, by intents
that each read "nothing is added here" and recorded their row anyway, and A95 now records that
this is evidence about the banner rather than about the twelve.

Nothing else changes. No row is deleted, renumbered or rewritten — 94 before, 95 after — and
the paragraphs explaining the state column, `approved`, `accepted until` and an `open` row are
untouched. `CLAUDE.md` already named this file as where assumptions and decisions are written
down; that line needed no edit and is now true.

For anyone who read the banner and put a decision elsewhere: it belongs here if a later intent
will need the fact, and in the intent's P2 if it does not.

<!-- xeno:section:residual-risk -->
## Residual risk

The whole change rests on one reading, and if the reading is wrong the change is backwards.
Section 4's sentence is about a record kept "as a file in the branch rather than an artifact
under `.xeno/`", in a paragraph about how a package was worked while no runner existed, which
is why this intent reads it as the stand-in for artifacts. A reader who takes it as covering
every hand-kept record, this register included, gets #174's conclusion, and it is not an absurd
reading — it is the one twelve commits' worth of intents did not bother to check either way.
The mitigation is that the sentence is quoted in the file, so the next reader checks the
inference rather than inheriting it. Low likelihood, total impact, and the cheapest thing a
reviewer can do with this PR is read those two lines of the plan.

The register still has no reader and now says more. Four claims, none checked, two checkable in
principle. The file's claim surface grew while its enforcement stayed at zero, which is A90's
own finding about itself; A95 and P4's gaps both say so rather than leaving it implied.

The next drift will look like the last one. Nothing compares the file against practice, so what
changed is only that the stated rule is now the correct one: a future divergence will be
practice leaving a right rule rather than a wrong rule standing. Better, and the same mechanism.

The seventeen rows were not revisited against the test the file now states. If some do not
outlive the intent that found them, the register states a rule its own history partly fails and
nothing will report it. Declined deliberately as seventeen questions, and a reviewer who wants
them asked is asking for a different intent rather than for more of this one.

Whether the durability test should also be in `CLAUDE.md` is open. It is read before every
change, which is where the distinction would do the most good, and leaving it out means the
rule lives only in the file a person reaches after deciding they have something to record.

What is not a risk: the 363 pre-existing verdicts, which `gate verify` confirms at exit 0, and
anything executable. No gate, rule, template or Go file opens this register, `artifacts_hash`
does not cover it, and the suite would pass identically had nothing been written.
