---
intent: github.com/triplem/xeno#1
phase: 05-review
created: "2026-10-03T10:24:41Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0f66b4c3f6e9a5de9aaaf173fd647aabc1d906e34e3180a740be7f2be4b06134
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
  - rule: deviations-are-traceable
    result: met
    note: >-
      The three deviations name what they depart from: two decisions that arrived against this
      intent's scope and the register that would have held them, the column convention, and P1's
      own constraint that the artifacts be short.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      Nothing outside this intent depends on what changed; `ASSUMPTIONS.md` is prose and the issue
      is a record. What changes for a reader is where a decision goes, and the second paragraph of
      the new header is that note.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added, and there is no code in this intent.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** The three deviations name what they depart from: two
decisions that arrived against this intent's scope and the register that would have held them, a
column convention, and P1's own constraint that the artifacts be short.

**`interface-change-needs-a-migration-note` — met.** Nothing outside this intent depends on what
changed: `ASSUMPTIONS.md` is prose and the issue is a record. What does change for a reader is
where a decision goes, and that is the second paragraph of the new header, which is the migration
note.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added and there is no code in
this intent.

**Beyond the three rules.**

*Is the closure honest about which reading it closes?* Both, separately, with the evidence for each
named: the host's setting and the verdict for the gate job, the agent layer for the walking
skeleton. A reader who wants to know when each became true reads the dates rather than the
milestone, which the gaps say.

*Was anything closed that is not met?* The strict reading of section 6 — WP11 as a whole package
before M0 — is not met, and this closure does not claim it. WP11's own done-when requires a phase
to be carryable with commands alone, and twenty-seven intents did that; the MCP server, the lenses
and the second harness are named as outstanding.

*Did anything get rewritten?* No. Nineteen lines added, one sentence moved into the past tense,
seventy-seven rows untouched. Closing the register is closing it to additions.

*What is the weakest part?* That the closure rests on a host setting nothing in the trail records.
`enforce_admins` could be turned off tomorrow and M0 would stay closed. The gaps name the scheduled
check WP10 describes as what would catch it, and it is not set up.

*What did this intent learn that the plan should know?* That nothing watches a milestone, and that
M1 is one clause and one reading away under the same silence. The learning in P0 proposes closing a
milestone in the review of the piece that meets its conditions, which is the only place the
conditions are freshest.

<!-- xeno:section:release-notes -->
## Release notes

**M0 is closed.** Xeno checks its own repository, and from here the loop runs through the runner
rather than through a file kept by hand.

What was checked: `verify` is a required status check on the default branch with administrators
included, so no merge arrives without a recomputed verdict and nobody can bypass it; that check
runs `xeno gate verify`, green over 216 verdicts; twenty-seven intents have run all six phases;
eleven of fourteen gates are implemented, where the plan's M0 paragraph expected eight; and the
agent layer exists for one harness, with seven skills, validated manifests and hook wiring.

**`ASSUMPTIONS.md` is closed to additions.** A77 is the last row. Every row stays with its state
column, because they are where a reader of this tree learns why it is the way it is. From here a
decision belongs to the design phase of the intent that takes it, and a learning to that phase's
`learning.yaml` and from there to a merge request against the rule set.

**What M0 triggers and is not in this release.** This repository adopts the `Xeno-Intent:` trailer
from `examples/rules/` — its own intent, next, because it changes how every later commit and every
squash message is written. And two paragraphs of the implementation plan describe a state two
months old: "at M0 six of the fourteen do not exist", where three do not and G-Questions is named
among the missing although it has been implemented since WP1, and "M0 proves one phase", where
twenty-seven intents have proved six. Both are a person's to correct.

**Two decisions arrived while this ran** and are carried rather than dropped: the shipped rule set
is closed at four rules for v1 with everything else routed to `examples/rules/`, which the next
intent records where it acts; and the `by-hand` placeholder rule stays as A62 and A66 left it, with
the fact that two hash fields accept it for artifacts written before their writers existed to be
recorded in section 16, which is a person's edit.

**What is not claimed.** M1. Its ingredients are built and one clause of WP8's done-when is
unverifiable by construction, and a milestone's reading is a person's.

Closes #1.

<!-- xeno:section:residual-risk -->
## Residual risk

**The closure rests on a setting nothing in the trail records.** `enforce_admins: true` and the
required `verify` check are the host's state as read today. If either changes, M0 stays closed and
stops being deserved, and the only thing that would notice is somebody running
`xeno enforcement check` with a token. WP10 describes a scheduled run for exactly this and it is
not set up in this repository.

**Nothing watches a milestone, and M1 is next.** M0 was reached weeks before anything said so, and
what found it was a question about open work packages. M1's ingredients are built; whoever asks
next will find that out the same way.

**The register's closure is a paragraph, not a check.** Nothing stops another row. Two decisions
arrived looking for a home within the hour of the switch being designed, which is the test the
guard has already half failed: they were routed elsewhere because this intent noticed, not because
anything refused them.

**A decision is now harder to find than it was.** Seventy-seven numbered rows were greppable; a
paragraph in one intent's design phase is reachable only by knowing which intent. The first cost
will be somebody re-deciding something already settled, and the mitigation — an index of decisions
across intents — is a thing nobody has asked for yet and which WP16 would own.

**The two readings were closed together.** The gate job has been true for weeks and the walking
skeleton since #169. One closure, two dates, both in the evidence rather than in the milestone.

**The artifacts of this intent are the shortest of twenty-seven and still about forty thousand
bytes**, which is the floor rather than the padding. That is now a measured point for #117 and it
is also a statement about this process: a milestone closure costs six phases because everything
does, and the only part an author can make smaller is the part that was already small.
