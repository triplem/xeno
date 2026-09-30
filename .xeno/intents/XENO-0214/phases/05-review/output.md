---
intent: github.com/triplem/xeno#147
phase: 05-review
created: "2026-09-30T21:26:30Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0a51653.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ab56f5016b91a781768c831a0b5916d95a8b2e977d057619259edf8e8a247083
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

| item | state |
|---|---|
| The three standing rules | held. No document edited. No invented name: four requirements from `enforcement.Names()`, asserted twice. The change belongs to #147, labelled `wp12`, on `147-the-gitlab-adapter` |
| **A65 is borne out** | `host.BranchRules` byte identical to `main`'s after a second adapter with a different request shape. This is the finding, and it is the reason the four pieces were ordered as they were |
| #104's second done-when | met. `enforcement check` answers for a GitLab project and every requirement it cannot express is `not-available` with a reason |
| #104 closes | all four pieces done: the vocabulary in #141, the port in #143, the wrapper in #145, the adapter here. `wp12` has no open issue after this |
| No shared code with the other adapter | held, and P2 records why it would have been wrong: the two differ in how many requests they need and therefore in whether a failure is partial |
| A42 | `internal/gates` reaches neither `net` nor `net/http` with a third `net/http` package under `internal/` |
| The credential | nowhere in the tree; the only `glpat` strings are the secret filter's own patterns. The live tests read it from the environment and skip without it |
| Eighty-eight columns, SPDX, no copyright line | held on both new files |
| The check suite | build, vet, test, gofmt clean; `gate verify` exit 0; GitHub report identical to `main`'s |
| Commit and reference convention | Conventional Commits, `Closes #147`, `Closes #104` on the same commit since this finishes the package |

**What a reviewer should push back on.**

**`merge_access_levels` is decoded and read by nothing.** It is in `gaps`. Either the field goes, losing
the record that the host answers it, or it stays as a struct member no code uses. A reviewer may
reasonably want it gone.

**`approvals.required` sums the rules.** Defensible, confirmed against one real project, and wrong in a
direction: overlapping eligible approvers mean one person can satisfy two rules, so the sum can
overstate. The alternative understates. Neither matters at the counts a declaration uses and a reviewer
may disagree about that.

**The reason text names three possible causes rather than one.** It is longer and vaguer than the
GitHub adapter's sentences, and that is deliberate after the 403 turned out to be ambiguous. Somebody
who prefers a confident sentence should read P4's table first.

<!-- xeno:section:release-notes -->
## Release notes

**`xeno enforcement check` answers for GitLab.** A project with `tracker.adapter: gitlab` in
`.xeno/config/project.yaml` is now compared against its host rather than refused.

What is read: `only_allow_merge_if_pipeline_succeeds` for `required_pipeline`; the protected branch's
push, force push and unprotect grants for `allow_bypass`; the project's approval rules and approval
settings for the two approvals requirements. `merge_method` is reported unchecked, as on every host,
because section 13 says so.

**The words are GitLab's.** `allow_bypass` says which bypass it found — a grant that may unprotect the
branch, a force push allowance, or a grant that may push without a merge request — rather than the
other host's sentence about administrators.

**A requirement the tier or the token cannot reach is `not-available`, never unmet.** Record it as
waived in `project.yaml` with a reason and a date and it becomes a decision in the repository instead
of a line on every run.

**With #141, #143 and #145 this completes GitLab support** for the gate wrapper and the enforcement
check. There is no tracker adapter for any host; that is WP12.

<!-- xeno:section:residual-risk -->
## Residual risk

**A Free tier's own answer is still unobserved.** The adapter handles 403, 404 and 402 identically and
names no cause, so the risk is not a wrong state but that somebody on Free reads a sentence offering
three explanations when their situation has one. It is the last of the fixture dependency that ran
through all four pieces, and it is now confined to the wording of one note.

**The sum in `approvals.required` can overstate.** Overlapping eligible approvers across rules mean one
person may satisfy two, so a project with two rules requiring one each is reported as needing two and
may in practice need one. It reports a requirement as met where it is met and can report unmet where a
declaration asked for more than the sum understates — the direction that fails loudly rather than
quietly, which is why it was chosen.

**`allow_bypass` reports the largest finding and hides the others.** A branch with an unprotect grant
and a force push allowance reports only the first. Fixing the unprotect grant then changes the report
rather than clearing the requirement, which reads as the tool moving the goalposts. The alternative was
a count, which P2 rejected because the three lists are not commensurable, and the note says that
unprotecting subsumes the rest.

**Three requests where the other host makes one.** A daily schedule against a slow or rate limited
deployment can now partly fail, and a partial failure reports `not-available` for the rows it lost,
which is indistinguishable in the report from a tier limit. The reason names the status, so the
distinction survives in the note and not in the state.

**The wrapper and the adapter have never run together.** #145's `.gitlab-ci.yml` has never executed and
this intent's command was run from a shell. The one thing that would test the pair is a merge request
on a real GitLab project, and the risk #145's P5 named still stands: an empty base SHA would make
`gate run` judge the wrong commits rather than fail.

**Not a risk.** `merge_access_levels` being unused. It is untidy and it is in the checklist for a
reviewer to decide; nothing reads it, so nothing can be wrong about it.
