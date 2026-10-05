---
intent: github.com/triplem/xeno#225
phase: 05-review
created: "2026-10-05T17:29:29Z"
schema_version: "1.0"
runner_version: dev+6adc0f9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: afca0b01e8db7260910dd042a887f1c9b704c8df2098cbfc1389fa7111b97b91
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Four, each naming what it departs from. Against P2 design: the condition is the artifact existence and not a comparison against the lock, because before a second start the two still agree — the staleness is what the start causes rather than what it finds, so the designed condition would have passed in exactly the damaging case. Against P2 claim that nothing outside Start needed touching: the suggestion offered the refused start, fixed in the state computation. Against acceptance criterion 4, which asked for a harmless second start to stay allowed: no such state survives a real clock, and the test now asserts the opposite. And one behaviour change beyond the issue: starting over takes the phase directory rather than the verdict, with #215 message, comment and test corrected together.'
      result: deviation
      rule: deviations-are-traceable
    - note: 'Two interfaces change for a person rather than for a caller, and both are documented in the messages themselves. Starting a phase over took the verdict and now takes the phase directory: anyone following #215 instruction would remove gate.yaml, hit the new refusal, and be told the correct thing by it, so the migration note is carried by the refusal a person actually meets. And a phase with an artifact but no marker now reports running where it reported not-started, which changes what xeno intent status says across machines and after retention expires the marker; nothing consumes that value outside this repository, and inside it the two consumers are the suggestion and the reached-phase computation, both of which were wrong before and are right now.'
      result: deviation
      rule: interface-change-needs-a-migration-note
    - note: 'None added. go.mod is untouched and the change uses fm.Exists and model.PhaseDir, both already imported by runner.go; the test helper uses time, already imported by the test file. No tool was weighed either: the thing that would have needed one is detecting an artifact already damaged by a past restart, which is named in P4 gaps and needs no dependency, only a decision about whether intent status should report it.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No normative document is touched: section 6 already says `phase start`
refuses a second start of a running phase, and this makes the tool able to tell that a phase is
running when the marker cannot say so, so the code moves towards the specification. Nothing is
invented — no field, gate, command or artifact shape, and the one state value used is one the
set already had. The branch carries one intent, the commit references #225, and the issue
carries `wp7`.

The acceptance criteria. Eight met, one pending until the commit, and **one falsified by its own
test**: criterion 4 asked that a phase whose lock still matches be allowed to start, and no such
state survives a real clock. P4's results carries it.

The non-goals held. No `xeno intent resume`. No age figure. No change to G-Schema, to the
marker's home, or to what a resumed phase is told about its predecessor. No repair of artifacts
already damaged.

What a reviewer should look at hardest is the escape hatch. Starting a phase over now takes the
phase directory, where #215 documented the verdict; removing the verdict alone leaves the
artifact and lands on the new refusal. Both messages, #215's comment and #215's test were
corrected together so nothing sends a person from one refusal into another, but it is a change
to documented behaviour that #225 did not ask for and the criteria did not name.

What this intent got wrong is in P3's deviations, and both errors were caught by tests the
criteria required rather than by review. The design's condition compared the artifact against
its lock, which passes in exactly the case the damage is about to be done; and the design said
nothing outside `Start` needed touching, while the suggestion was offering the very start being
refused.

What is unplanned and good: #225's first gap closed as a consequence. A phase begun on one
machine no longer reads as `not-started` on another, because the state keys on the artifact
rather than on the machine-local marker.

G-Test judges this intent. P4 declares `test-report/go-test` bound by hash and its sealed verdict
carries `G-Test: pass`.

<!-- xeno:section:release-notes -->
## Release notes

`xeno phase start` refuses a phase that is already under way, and `xeno intent status` reports
one correctly on a machine that did not start it.

Before this, a phase with its sections written, no verdict, and no run marker could be started
again. The start rewrote `context.lock.yaml`, which is inside `artifacts_hash` and is the only
record of what the phase was given; the artifact's `context_hash` then named a lock that no
longer existed, and `phase finish` reported **red on G-Schema two commands later** — nowhere
near where the mistake was made.

The marker goes missing in the ordinary course. A9 puts it under `.xeno/local/`, gitignored,
because it describes a machine rather than the trail, so a phase begun on one machine has no
marker on another, and `retention: local_days: 30` expires it on the same one.

So the refusal keys on the artifact, not the marker:

    00-intake is already under way; carry on with section set and phase finish,
    which need no second start, or remove .xeno/intents/KEY/phases/00-intake
    to start the phase over

**Starting a phase over now takes the phase directory, where #215 said the verdict.** Removing
`gate.yaml` alone leaves the artifact, which is itself a phase under way, so both refusals now
name the directory and #215's comment and test were corrected with them. Anyone following the
old instruction meets the new refusal, which tells them the right thing.

**A phase with an artifact and no verdict now reports `running`**, whether or not this machine
holds the marker. It read `not-started` before, which is the first gap #225 names and the reason
the printed next step was offering `phase start` for a phase that had already been started.

What #225 asked and this does not do: there is no `xeno intent resume`, no figure for how long
an intent has been sitting, and nothing new telling a resumed phase what changed under it. The
issue's own conclusion is that `intent status` plus the printed next step is a resume in all but
name; what was wrong was that the obvious way to re-enter a phase destroyed its record, and that
is what this fixes.

<!-- xeno:section:residual-risk -->
## Residual risk

The escape hatch moved and somebody may be carrying the old instruction. #215 said removing the
verdict starts a phase over; it now leaves the artifact and meets the new refusal. The refusal
names the right thing, both messages agree, and #215's comment and test were corrected — so the
failure mode is one confusing step rather than a wrong outcome. It is still a documented
behaviour changed by an intent that was not asked to change it, and it is the first thing to
object to in this diff.

An artifact already damaged by a past restart is not detected. Anyone who hit this before today
has a phase that will go red on G-Schema, `intent status` will report it as running, and the only
route back is removing the phase directory. Nothing says so at the point of confusion; the
finding arrives at `phase finish` as it always did.

The marker is now nearly vestigial and nobody has decided that. It still produces the "already
running" refusal on the machine that holds it, but the artifact produces a stronger one
everywhere; what the marker uniquely covers is the window between `phase start` and the first
`section set`, when no artifact exists. Worth knowing before anybody removes it, and outside this
intent.

The state value `running` now means two things that were distinct: this machine is in the middle
of the phase, and some machine has been. Nothing reads the difference today — the two consumers
are the suggestion and the reached-phase computation, and both want the second meaning — but a
later reader could reasonably expect the first.

The pinned clock will hide the next defect of this shape too. `newFixture` fixes `Now` for good
reasons, and `f.moving()` exists in one test file for one issue; anything else whose damage is a
timestamp moving will pass its tests the same way this did. That is a property of the fixture
rather than of this change, and it is the most transferable thing this intent found.

What is not a risk: the 393 pre-existing verdicts, confirmed at exit 0; a first start, which has
no artifact and cannot meet the new guard; and the two older refusals, whose conditions are
unchanged and asserted.
