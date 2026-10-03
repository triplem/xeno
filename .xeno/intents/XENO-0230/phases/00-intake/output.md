---
intent: github.com/triplem/xeno#186
phase: 00-intake
created: "2026-10-03T13:21:02Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 55df07b8aec14a356f3bf9328891f3dbdd5ad219268b3fb95eaddbb0021c4d81
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

Six commits reached `main` in one morning while the `audit` gate was red, and nothing could have
stopped them.

| commit | subject |
|---|---|
| `ea0cb1c` | The agent layer, for one harness (#170) |
| `8c1b647` | The lock says what section 5 says, and a phase is told what moved (#173) |
| `f001058` | M0 is closed, and the record kept by hand closes with it (#174) |
| `d882c44` | A link naming a document that is not there is a finding (#175) |
| `4d472b6` | A budget is judged against the size the lock recorded (#178) |
| `6c4a1aa` | The intent is created by a command (#180) |

Every one of those pull requests showed green checks and every one was honest about it.

**A check that cannot run before a merge cannot gate one.** `audit` triggered on `push` to `main`, on
a schedule, and on `pull_request` filtered to three paths — its own file, the baseline and
`.releaserc.json`. A pull request touching none of those never ran it, so the job reported after the
merge, on the push. That is a property of the trigger and not of anybody's diligence.

**And only one check was required.** `main`'s protection required exactly one context, `verify`.
`audit`, and the three scanners, were advisory: a pull request could be merged with all of them red
or with none of them having run.

**The check that exists to catch this reported met.** `project.yaml` declares
`enforcement.required_pipeline: true`; the GitHub adapter answers met when *any* status check is
required. So the repository declares that a merge needs a green pipeline, asks the host, and is told
yes, while four of five workflows bind nothing and one of them had been red for a day. Appendix A's
`enforcement` block has rows for `required_pipeline`, `allow_bypass`, `merge_method` and `approvals`
and nothing for *which* checks, so the declaration cannot express what was violated.

**Three of the five checks could not have been required even if somebody had tried.** The job id is
the name of the check GitHub reports, and `gitleaks.yml`, `semgrep.yml` and `trivy.yml` all named
their job `scan`. A required check is matched by that name. Three identical names cannot be required
individually, and `scan` names no particular workflow.

**What was red.** `audit` was last green on 2026-10-02 13:06Z and has failed on every run since,
seven of them. Nothing in the repository changed: advisories arrived against the pinned release
toolchain and took it from the 15 high the baseline recorded to 35. The daily schedule found it,
which is what the schedule is for, and the finding bound nothing.

<!-- xeno:section:scope -->
## Scope

**In scope.** `audit` on every pull request, with no path filter, so the job that judges the release
toolchain runs before a merge and not after it. Distinct check names for the three scanners, so each
can be required. The advisory drift read and answered: the pin moved to the current
`semantic-release` and the baseline measured against it rather than raised to meet it. The version
written once and read rather than copied, because the audit's premise is that it installs what the
release installs. Node pinned in both workflows, because the newer toolchain declares an engine and
neither workflow said which node it ran under. And the rule written where a person reads it, in
`CONTRIBUTING.md`, with the reason being that it failed.

**Out of scope, and each for its own reason.**

The branch protection itself. It is a host setting and not a commit — the category `xeno init`
already prints as "Xeno cannot make these settings" — and it cannot be applied before this change
merges: a required context that never reports blocks every merge, and the three renamed contexts do
not exist on `main` until then. The command is in the release notes and the ordering with it.

`required_pipeline` naming the set. That is the part that was supposed to notice and it is an
Appendix A addition, so a specification change and its own commit first by the first standing rule.
This intent is the finding; #186 carries it and says so.

A hook. Section 7 settles it twice: a hook's "result is advisory, the binding result is the CI run,
because hook configuration lives on the developer machine", and "a hook only invokes checks the
runner already implements" — asking a host for the status of its check runs is neither a check this
runner implements nor something a developer machine can be relied on to do. The fast feedback a hook
gives is not the gate, and this intent is about the gate.

The thirty-one findings that remain after the upgrade. They are read and recorded, not fixed: npm's
own suggestion for thirty of them is to downgrade `semantic-release` to 15, which is not a fix.

<!-- xeno:section:context-rationale -->
## Why this context

Section 7's enforcement subsection is read for what a project declares about its host and for the
sentence the adapter quotes back when the declaration is unmet — "a red gate then produces a report
that a merge walks past" — which is exactly what happened and is the clearest statement of why this
matters that either document contains.

Section 7's hooks subsection is read for the two constraints, because the question of whether a hook
could enforce this was asked and both constraints answer it: advisory by construction, and no
exclusive logic.

Appendix A's `enforcement` rows are read to establish what the declaration can and cannot say. Four
rows, none of them a set of checks, which is why `xeno enforcement check` reported met.

`internal/host/github/github.go` is read for `requiredPipeline`, which is the eight lines that turn
one required context into a met declaration.

`.github/workflows/audit.yml` is read in full: its triggers, which are the defect; its header, which
claims the tree is "installed here exactly as the release installs it"; and its baseline comparison,
which is right and was never the problem.

`.github/workflows/release.yml` is read for `semantic_version`, which is the number the audit's claim
depends on and which was written twice with nothing checking that the copies agreed.

`.github/npm-audit-baseline.json` is read for the rule it sets about itself: "Raising a number here
is a deliberate act and belongs in the commit message that raises it." That sentence is why this
intent measured a newer toolchain before touching the number.

The run log of `37123738064` and the `npm audit --json` output for both `24.2.9` and `25.0.9` are
read as the evidence: twenty distinct advisories against twelve, and which ones move.

`CONTRIBUTING.md` is read for where the merge procedure already lives, because that is where a rule
about merging belongs rather than in a new file.
