---
intent: github.com/triplem/xeno#215
phase: 05-review
created: "2026-10-03T20:24:35Z"
schema_version: "1.0"
runner_version: dev+b626f1a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7b75527ddbd9c4a99ad49a756795fad76742516a3bfac3ce890fa8816c7f2da1
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
      Two, both in the implementation phase's deviations section. The issue's first claim was
      wrong — #171 had already built what it said was missing — so the work shrank from a
      mechanism to a refusal. And a second derivation of the changed set had been written here
      before the existing one was found; it was deleted rather than finished, because two
      derivations of one answer can disagree, and it is in no commit.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      One command changes behaviour: xeno phase start on a phase that has a verdict now exits
      non-zero instead of rewriting the lock. Nothing in this repository relied on the old
      behaviour, and the path that came closest — remove the run marker after a dead run and
      start again — is why the refusal keys on gate.yaml rather than on the lock. Starting a
      phase over stays possible by removing the verdict, which the message names.
  - rule: new-dependency-needs-a-rationale
    result: met
    note: >-
      None added. The change is four lines in the runner and two tests, using hashing and fm,
      which the package already imports.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No document is touched: the refusal enforces section 11's clause
rather than adding one, so nothing precedes it. No field, gate, tool or rule. The branch
carries one intent, the commit references #215, and the issue carries `wp8`.

The acceptance criteria. The refusal happens, writes nothing, names both ways forward,
reads like the three refusals already there, and the dead-run recovery stays open — each
with a test, and the refusal also run against this intent's own sealed phase.

The non-goals held. No reporting mechanism was built and the duplicate was deleted before
the implementation phase; no specification change; nothing about the context profile.

The issue's correction is on the issue, not only in this trail. A wrong claim left standing
in a closed issue is read later as fact.

What the work is not: it does not make a sealed artifact safe. It closes the one path the
runner itself owned. A hand edit still goes through, and `gate verify` is still what
catches it.

<!-- xeno:section:release-notes -->
## Release notes

`xeno phase start` refuses a phase that already has a verdict. Starting one rewrote
`context.lock.yaml`, which is inside `artifacts_hash`, so it rewrote a sealed artifact and
left `gate verify` to report the divergence afterwards — which it did, twice, during the
past day's intents. The refusal names the way to redo the work, `section set` with `phase
finish`, and the way to start over, removing the verdict.

Section 11's "what is sealed is never rewritten" gains a reader at the point of the act. It
had one that fires afterwards, and the deletion half gained its guard in #193.

The issue's other claim, that nothing tells a repeated phase what changed, was wrong: #171
built it, `ChangedSince` derives it from the predecessor's lock and `phase start` prints it.
The issue is corrected and no mechanism was added.

A93 records the reasoning, including why the refusal keys on the verdict and not on the lock.

<!-- xeno:section:residual-risk -->
## Residual risk

A refusal on a command every test and every intent calls is the kind of change that breaks
something quietly. The two paths it could have closed are the dead-run recovery, which keying
on the verdict keeps open, and the ordinary first start, which has no verdict to find. Both
are asserted, and the package's existing tests pass unchanged.

The escape is real and that is deliberate. Removing `gate.yaml` starts a phase over, and
`gate.yaml` is outside `artifacts_hash`, so removing it is not itself a divergence. Somebody
who does that loses a verdict and the decisions carried in it, and nothing warns them beyond
the sentence in the refusal.

Section 11 is better guarded and not guarded. The hand edit is the common way to rewrite a
sealed artifact, this session made two, and only `gate verify` catches it. A refusal in the
runner cannot reach what the runner has no command for, which is what #188 and #208 are
about.

The finding this intent's reading produced is the larger one and it is not addressed: no
intent has ever written a `context-profile.yaml`, so the information base is empty
throughout the trail, the budget is judged against nothing, and `ChangedSince` has never had
an input. A mechanism that works and has no input looks identical to one that works.
