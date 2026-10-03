---
intent: github.com/triplem/xeno#171
phase: 05-review
created: "2026-10-03T09:58:33Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 58870f0b0694383abe62e2c5111da1a71507a55d92a48d2057cc1e231fa25dd0
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
      The five deviations each name what they depart from: a test from an earlier intent, section
      5's example budget against the size of this repository, the sequence the two specified halves
      compose into, the shape of the changed-set report, and the byte count against the field A74
      would have used if one existed.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      The lock gains two fields, additively: `omitempty` means an older lock is a valid newer one,
      and nothing reads the fields yet except a person. No migration exists and none is needed.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      No dependency was added. Head is one more call in the package that already owns the git
      subprocess.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter. What follows is the reasoning and
the questions this change raises beyond them — including the one that should decide whether this
intent is released as it stands.

**`deviations-are-traceable` — met.** The five deviations each name what they depart from: a test
from an earlier intent that was right about the runner and wrong about section 5, the example
budget in section 5 against the size of this repository, the sequence the two specified halves
compose into, the shape of the changed-set report against the alternatives, and the byte count
against the field A74 would have used if one existed.

**`interface-change-needs-a-migration-note` — met.** The lock gains two fields, which is additive:
a reader of an older lock sees two absences where a newer one has values, and nothing reads the
fields yet except a person. The note is that no migration exists and none is needed, because
`omitempty` means an old lock is a valid new lock.

**`new-dependency-needs-a-rationale` — not-applicable.** No dependency was added. `Head` is one
more call in the package that already owns the git subprocess.

**Beyond the three rules.**

*One acceptance criterion is not met, and that is the question for this review.* 01-requirements
asked that a link naming a document which does not exist be a finding against the profile.
`informationBase` skips it silently. The mapping in P4 found it, no test had been written for it,
and P3 is sealed — so it cannot be fixed inside the phase that owned it without rewriting what is
sealed. It is written down in three places and carried out of the intent as a finding, which is
what the process is for. Releasing this intent means accepting that, and the alternative is a
second intent for one `if`.

*Is anything recorded that section 5 does not enumerate?* No. Two fields from its own list, and
the changed set as output. The temptation was a `changed_since_predecessor` field, and the section
answers it twice: its list has no such entry, and the file states what was declared rather than
what was read.

*Did the demonstration earn its place?* It produced three things no test did: the example budget in
section 5 is too small for a repository of this size by nearly a factor of two, the saving is
reachable only after four approvals, and the adoption cost for an intent of this kind is about
three releases. The first and third are figures the plan asked for and nobody had.

*What should a reader of this not conclude?* That a repeated phase now reads only what changed. It
is told what changed. Section 5 says the lock is not a measurement of what was read and that
treating it as one would be wrong, so the clause is implementable and unverifiable, and what was
built is the knowledge rather than the behaviour.

<!-- xeno:section:release-notes -->
## Release notes

**`context.lock.yaml` says what section 5 says it says**, two fields further than before.
`repo_commit` is the commit the phase was started against, absent where the repository has none.
`rules_applied` is the effective rule set with each rule's path and its own version counter, absent
where no rule is in force — which answers what `rules_hash` cannot: which rules, and which revision
of each.

**The information base is assembled in the profile's own order.** A project writes its `include`
patterns from stable to volatile and the lock records that order, which is what section 5 asks for
and what a harness's prompt cache rewards. Paths are sorted inside each pattern, so two runs over
one tree stay identical.

**A declared link's document is part of what a phase was given.** Until now the field was parsed
and ignored.

**A context over its budget is a finding.** G-Schema reports the recorded file count against
`budget.files` and the measured size against `budget.bytes`, naming both numbers. The phase stays
decidable: section 5 says deliberately a finding and not a block, because blocking against a number
nobody has experience with yet is the wrong way round.

**`xeno phase start` says what to reread.** Starting a phase whose predecessor declared a base
prints the files of that base whose hashes no longer match the tree, and says nothing where nothing
moved. It is derived from the predecessor's lock and the tree, and recorded nowhere.

**A project with no context profile is unaffected**, which is every project today and this one.

**Known and not fixed.** A link naming a document that does not exist is skipped rather than
reported — an acceptance criterion of this intent that it did not meet, carried out as a finding.

Closes #171. Refs #1.

<!-- xeno:section:residual-risk -->
## Residual risk

**An unmet criterion is being released rather than fixed.** A link naming a document that does not
exist is skipped silently. It is one condition and one finding, and it is not in this intent
because P4 found it after P3 was sealed, and the process this project runs on says what is sealed
is never rewritten. So it leaves as a finding, with a follow-up issue, and anybody who reads
01-requirements against the result will see the gap. That is the intended behaviour and it is
still a release with a known hole in it.

**The saving is behind the decision, and nothing in the documents says so.** With a profile in
force, editing the base turns the earlier phase red; a red predecessor stops the next phase from
starting; so the report that says what to reread is unreachable until the findings about those same
files have been released. The demonstration needed four approvals to get there. Both halves are
specified and each is right. Their composition is the kind of thing that only shows when somebody
walks the sequence, and it will decide whether anybody adopts a profile at all.

**The byte budget is the only number here that moves after a phase is sealed.** It is measured from
the tree, because the lock carries no sizes, so `gate verify` can report a budget finding against a
phase that was inside its budget when it ran. A74 solved the same class of problem by reading what
the artifact recorded; this has nothing recorded to read, and giving it something means a field
section 5 does not enumerate.

**Section 5's own example will put a project over budget on its first phase.** 400,000 bytes
against the 709,921 a plausible profile resolves here. The mechanism is right and the illustration
is misleading, and the illustration is what gets copied.

**The changed set is ephemeral by design**, so the trail cannot show what a phase was told to
reread. WP20 will have to measure the saving live rather than from the record, and a decision an
agent took on that basis cannot be reconstructed afterwards.

**Nothing verifies the clause the piece is named after.** "A repeated phase reads only what
changed" is now possible and still unverifiable: section 5 forbids reading the lock as a
measurement of what was read. What exists is a runner that knows and an agent that is told.

**The adoption question is yours, with the figure attached.** A profile of the shape demonstrated
costs about three stale-read releases per intent of this kind, plus one budget finding until the
number is raised. Every one of them is a true statement about the work. Whether that is a process
worth running is the decision, and it cannot be taken by the agent that would be writing and
releasing its own findings.
