---
intent: github.com/triplem/xeno#84
phase: 00-intake
created: "2026-09-26T16:04:01Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+00838d7
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 71901bd29695913144cff5b2e0522e5c003880fc3f542bdb15163d87ef69c895
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

The repository went public, so the host could answer about branch protection, and `xeno
enforcement check` reported what it found: no status check required and administrators
able to bypass. Two of the three requirements `project.yaml` declares were unmet.

The first is what first step 2 of section 4 calls the setting the whole tool claims to
rest on: the `verify` job recomputes every verdict on every pull request, and nothing
required it to pass, so a red gate produced a report that a merge walked past.

The second had become load bearing rather than a gap. One approval was required and a
single maintainer cannot approve their own pull request, so every merge went through the
administrator bypass, including the one that landed the change which measured this.

<!-- xeno:section:scope -->
## Scope

The three settings on `main`, in `scripts/github-settings.sh` so that they are
re-appliable, and the declaration in `project.yaml` that they are compared against.
`verify` required, administrators included, approvals at zero, force pushes and
deletions still refused, and `strict` false so that a branch need not be rebased before
it merges.

Not the scanners. Both report as `scan`, from two workflows, so the name cannot identify
one check to require, and the project treats a scan as a report rather than a barrier.

<!-- xeno:section:context-rationale -->
## Why this context

**Zero approvals reads as a loosening and is the opposite.** A requirement nobody can
satisfy is not a guarantee: it forced every merge through the bypass, which then covered
the required check as well, because a bypass covers all of them. Dropping the approval
is what allowed administrators to be included, and the machine gate binds the owner of
the repository as a result.

**Four eyes is not being claimed anywhere.** The process definition wants an approval
where a project has people to give one; this project has one maintainer, so the honest
declaration is zero rather than one recorded as waived, since `waived` is for a
requirement the host cannot express and this host expresses it.

**`strict` is false deliberately.** Requiring a branch to be current before it merges
would have made today's stacked branches unmergeable without a rebase apiece, and
`verify` still runs on the branch. What it gives up is the case where two mergeable
branches are individually green and red together, which is a real gap and cheaper to
catch on `main` than to prevent by friction.

**The escape hatch is gone, and that is the cost.** With administrators included, a
`verify` that never reports blocks every merge until somebody edits the protection. The
repair is a setting rather than a commit, which is why this is acceptable, and it is
worth knowing before it happens rather than during.