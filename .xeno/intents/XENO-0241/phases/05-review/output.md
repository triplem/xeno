---
intent: github.com/triplem/xeno#206
phase: 05-review
created: "2026-10-03T20:55:34Z"
schema_version: "1.0"
runner_version: dev+8574810.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7384078bed85492ab88cf111692674f0656debaf72db1f1c07b275d0a4b07314
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
      One, in the implementation phase's deviations section, and it names the acceptance
      criterion it departs from: the criteria asked for a test that a range touching no
      intent passes and said nothing about a path under .xeno/intents/ that is not an
      intent directory. The reader skips it, which is the only reading that cannot invent
      a key, and the case does not occur in a trail.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      Two surfaces change and neither needs a migration. xeno intent verify is new, so
      nothing depends on an older form of it. The shipped CI wrapper templates gain a
      call, which reaches an adopter at their next xeno init --vendor; a wrapper generated
      before this keeps working and simply does not carry the check, which is the state
      every wrapper was in until now. Section 13's rule that the plugin and the runner move
      together is what makes that an update rather than a break.
  - rule: new-dependency-needs-a-rationale
    result: met
    note: >-
      None added. One dependency remains, go.yaml.in/yaml/v3, vendored. The range is read
      with git diff through internal/git, which already starts a subprocess for section 9's
      commit predicates and is the only place in the tree that does.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No document is touched: section 7's G-Complete is unchanged and
runs where it ran, and what was added is a second reader of section 8's sentence about how
an intent ends, so nothing precedes the code. No field, gate, tool or rule — the gate list
is not extended and the new command is a reader, not a verdict. The branch carries one
intent, the commits reference #206, and the issue carries `wp1`.

The issue's three answers were put to a person before anything was built, because the issue
says the shape is open and names the first question as whose job it is. The first was
chosen. The third would in any case have been a document change and so a person's commit.

The acceptance criteria are mapped one by one in the verification phase, each to the test
or the exercise that answers it. The four that are about what CI does were answered by
running the command over a clone in the job's exact form, which is also where the one
defect the tests missed was found.

The non-goals held. G-Complete is untouched, nothing writes, no migration was needed
because the trail was measured complete first, the rule that every change belongs to an
intent was left to #120 rather than absorbed, and `CLAUSE-READERS.md` is unchanged.

What the work is not. It does not make an incomplete trail impossible. It makes one
arriving at a merge loud, which is where the process says a merge is decided. An intent
stopped and left on a branch nobody opens is still invisible, and so is one whose phases
are all present and all wrong.

<!-- xeno:section:release-notes -->
## Release notes

`xeno intent verify --base REF --head REF` is new. It answers whether every intent a change
touches has reached one of the two endings an intent has: a decided final phase, or `xeno
intent close`. An intent that simply stopped is neither, and it exits 1 naming the intent
and the phase the work reached. A range that touches no intent exits 0 and says so. A
missing or unresolvable ref exits 2, because a comparison made on no evidence reports the
same thing as a clean one.

The check runs in CI. This repository's `verify` job calls it on every pull request, with
the same base as the trail guard, so a branch whose trail stops in the middle no longer
merges with every gate green on the phases it did write. Both shipped CI wrapper templates
carry the same call, so an adopter's generated pipeline gets it too.

Why it was missing: G-Complete runs at P5 and at `intent close`, and an intent that reaches
neither has no phase for the gate to run in. XENO-0230 merged with its verification and
review phases absent and nothing red. The state this reads is the one `xeno intent status`
has printed since #132 — what was missing was something that treated it as a condition.

One consequence to expect: a pull request is red until its intent reaches a decided final
phase. That is what every other gate already does, and it is the point.

<!-- xeno:section:residual-risk -->
## Residual risk

**The step is first exercised by the pull request this intent opens.** The command is
tested and was run by hand in the job's exact form over a clone, but the wiring — the base
expression, the `if`, the step's place in the job — is asserted by nothing in the
repository. If it is wrong it is wrong on this pull request, loudly, which is the cheapest
place for it to be wrong.

**An intent can still be left incomplete where no merge follows.** A branch abandoned
without a pull request is invisible to a check that runs at the merge, and that is the
residue of choosing the merge as the moment. The alternative, checking the whole trail on
every pull request, was rejected in the design: it would make one unclosed intent
everybody's red check, and the cheapest way to pass it would be to close intents nobody
had finished.

**A branch that touches no intent passes.** The rule that every change belongs to an intent
is real and unenforced, and #120 owns it. Until then, the way around this check is to write
no trail at all rather than an incomplete one.

**The shipped wrappers are read, not run.** Nothing in this repository executes a generated
wrapper on either host; an adopter's first run is their own first pull request. WP10's
standing gap, not made worse here.

**A provisional final phase now blocks a merge.** Section 6 says a provisional verdict is
reported rather than failed and becomes a hard condition at the merge request, so this is
the condition rather than a new rule — but it is the first thing in this repository that
makes a provisional verdict stop something, and the first intent to declare evidence will
be the one that feels it.
