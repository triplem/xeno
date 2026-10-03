---
intent: github.com/triplem/xeno#190
phase: 05-review
created: "2026-10-03T14:06:20Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+74553ad.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7f3ac872c1eeb39de806065ea92e92458a324a5a67295f1fad293c6ab1c2c07c
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
      Two, each naming what it departs from: the missing-base refusal written as its own branch
      rather than folded into the comparison, and the title pipeline split across two lines because
      one line runs to 102 columns and a shell wrap is load bearing rather than cosmetic.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      A pull request with a prose title now fails, which is the first thing a contributor will meet.
      The step's message, CONTRIBUTING.md and this description all say what to do and why. Nothing
      has to be migrated, but something that used to pass now does not.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added; both steps use git and a binary this repository builds.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Two, each naming what it departs from: the missing-base
refusal written as its own branch rather than folded into the comparison, and the title pipeline
split across two lines because one line runs to 102 columns and a shell wrap is load bearing rather
than cosmetic.

**`interface-change-needs-a-migration-note` — met, and it is the substance here.** A pull request
with a prose title now fails, which is the first thing a contributor will meet. The step's own
message says what to do and why; `CONTRIBUTING.md` says it at greater length, immediately after the
subject rule it qualifies. Nothing has to be migrated, but something that used to pass now does
not, and the three places a person could discover that — the step, the file and the pull request
description — all say the same thing.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added; both steps use git and a
binary this repository builds.

**Beyond the three rules.**

*Is either check in the place that binds?* Yes, and it is the only question that mattered. Both are
steps of the `verify` job, whose id is the context `main` requires. A workflow of their own would
have read better and bound nothing, which #186 measured at six merges.

*Does either restate a rule?* No, and four of the eight alternatives were refused for proposing it.
The title calls the command that `CLAUDE.md`'s rule and the plugin's pattern already share; the
trail calls git rather than a new hash over the set of intents, which would have been a field and
therefore a specification change.

*Was the destructive case actually tried?* Yes, in a worktree, before the step existed: a committed
deletion of one intent reports 31 paths. The case that decided the design — a directory added and
dropped inside one branch — reports 0, and reasoning about `git diff A B` against `git diff A..B`
would have got that right only by accident.

*What does this intent leave open that a reader should not mistake for closed?* Three things, all in
P4's gaps and worth naming here: the re-seal hole, which is now the only quiet way to rewrite the
trail and is left open deliberately; G-Complete running only at P5, which is how XENO-0230 reached
main two phases short and is the third instance today of the thing that should have noticed not
being able to; and a title edited after the last push, which `pull_request` does not re-check.

*Could this change a verdict behind it?* No. Two files, neither Go, no gate, no rule, no artifact.
`gate verify` reports 265 at exit 0.

<!-- xeno:section:release-notes -->
## Release notes

**A pull request title has to be a Conventional Commit.** It is the commit subject: the squashed
message is built from the title and the description, so the subject written on a branch is discarded
at the merge.

    feat(docs): a description

The host appends the number, so that arrives on main as `feat(docs): a description (#3)`. Write it
without the reference.

**A prose title releases nothing, silently.** semantic-release reads the subjects on main to decide
whether there is a version to cut, and one it cannot parse is not a small release but no release.
Fifteen merges of this repository carried prose titles and nothing was released for five days.

**The trail only grows.** A pull request that removes or renames any path under `.xeno/intents/` is
refused, naming the paths. What is sealed is never rewritten, which section 7 states as a rule
holding throughout; an edit was already caught by `gate verify`, and a deletion was caught by
nothing — removing one intent took the count from 259 verdicts to 253 and exited 0.

Adding a directory and dropping it again inside one branch is not that and is not refused: the
comparison is against the base of the pull request, so what never reached main is nobody's business.

**Both are steps of the `verify` job**, so both gate the merge rather than reporting beside it.

<!-- xeno:section:residual-risk -->
## Residual risk

**Neither check has been seen failing in CI.** Both were seen failing by the same invocation in a
worktree, and seen passing in CI on this pull request. The first real failure will be a contributor
meeting an error message, which is a poor place to discover a mistake in the message. Closing it
properly means a pull request whose purpose is to be refused.

**A title edited after the last push is not re-checked.** `pull_request` does not fire on `edited`
unless the workflow asks for that type, and GitHub re-runs required checks on a new commit rather
than on a retitle. Narrow — it needs somebody to edit a title and merge without pushing — and real.
Closing it means re-running the whole job on every description edit.

**The re-seal hole is now the only quiet way to rewrite the trail**, and this intent makes it the
obvious one by closing the other two. Editing a sealed artifact and re-running `phase finish` writes
a verdict over the new content; G-Freshness turns the successor red, so it means re-judging an intent
in a diff a reviewer is looking at. Left open because refusing modification would refuse a
re-judgement the working sequence provides for.

**An intent can still reach main incomplete.** G-Complete runs at P5, so one that never reaches P5
is never checked for completeness. XENO-0230 proved it this morning with every gate green. Nothing
here touches it.

**`fetch-depth: 0` makes the checkout fetch the whole history.** Seconds on this repository, and it
grows with the trail, which grows with every intent. The number to watch is the `verify` job's
duration rather than the clone's.

**What is not a risk.** Any verdict behind this intent — two files, neither Go — and the five check
contexts, which are unchanged, so the protection rule needs nothing.
