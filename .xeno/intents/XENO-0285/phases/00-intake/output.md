---
intent: github.com/triplem/xeno#95
phase: 00-intake
created: "2026-10-09T16:43:00Z"
schema_version: "1.0"
runner_version: dev+d2bc411.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 529394b65a69b529fb75660837924fcfe7496c59e8ae5cb1dfd4598b927d7989
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: '#95 closes when the configuration is complete, and switching renovate on becomes its own issue'
      rationale: 'The issue asks which shape to adopt, and that is decided, built and on main. Whether RENOVATE_TOKEN exists is an operational step only the maintainer can take, not the adoption question, so holding #95 open for it would leave the issue as a standing reminder of a secret. The step is not lost: it is #350, opened before this intent started. The cost accepted is that #95 closes with renovate still inert, so nothing demonstrates the configuration works until the token is added.'
      decided_by: triplem
    - id: D-2
      chosen: 'The agent set the xeno-approved label on #95, on the maintainer''s explicit instruction'
      rationale: 'The label is the machine-readable half of the approval gate and is read as a person''s act. On #95 it was reported twice as not landing: the issue carried no labels and its timeline held no labeled event in its whole history, while the same query returned wp10 for #350, so the negative was real. The maintainer approved in words, in this issue''s comment and twice in session, and instructed the agent to set it. Recorded because xeno intent start''s check is unfalsifiable once an agent may set its own input, which is the property #330 was written to create and #95 is the case it was written from. The approval is the maintainer''s; only the label write is the agent''s.'
      decided_by: triplem
      proposed_by: agent
---

# Intake

<!-- xeno:section:problem -->
## Problem

Approved by @triplem on 2026-10-08T19:46:46Z: This only depends on switching host and/ or plan. Therefore we will go ahead, even without another plan or host.

> **github.com/triplem/xeno#95** — Replace dependabot with renovatebot
>
> Renovatebot does offer a dependency dashboard. This could be used instead of prs.

The issue as it stood at 2026-10-09T16:43:00Z, read by `xeno phase start` and quoted rather than summarised. What this phase concludes about it belongs below.

<!-- xeno:section:scope -->
## Scope

This intent finishes the configuration #95 decided and makes the issue closable. The
shape is already on main and already correct: XENO-0277 brought `renovate.json` and
`.github/workflows/renovate.yml` to the three answers given in this issue's comments —
self-hosted, no automerge, dashboard plus pull requests. Nothing here revisits those.

What is left is the argument the shape was chosen on, left unfinished. #95 prefers
renovate over dependabot because renovate can see versions written inside workflow text,
and three such pins are still invisible to it:

- the image's base, `ARG BASE=docker.io/library/debian@sha256:…` in the `Dockerfile`;
- `GITLEAKS_VERSION: 8.30.1` in `.github/workflows/gitleaks.yml`;
- `pip install zensical==0.0.69` in `.github/workflows/docs.yml`.

All three are rows of `docs/supply-chain.md`, so each is a pin this project has written
down as deliberate and cannot be told has moved. `enabledManagers` is `github-actions`,
`gomod` and `custom.regex`, and the regex managers target five workflows, none of them
these. So this is not a manager that is misconfigured; it is three pins with no manager
at all.

This intent adds a regex manager for each, and one line to the `Dockerfile`: the base
image reference gains the `13-slim` tag that `docs/supply-chain.md` already calls it by.
The tag is not decoration — a digest alone names no stream to follow, so the page names a
tag the tree does not carry, and that is also what a manager needs in order to propose a
newer digest for the right image.

It also drops `commitBody: "Refs #95"`. That footer points at the issue this intent
closes, and `CONTRIBUTING.md` puts in the footer the issue the work belongs to, which for
a dependency bump is none. It is also discarded at the merge: this repository sets
`squash_merge_commit_message=PR_BODY`, so a branch commit's body never reaches main.

What it does not do.

**It does not switch renovate on.** That needs `RENOVATE_TOKEN` in the repository's
secrets, which only the maintainer can add, and it is now #350 rather than a step of this
intent. Measured today, this repository has no secrets at all and the renovate workflow
has never run: it is registered and active, `id=378609475`, with no rows from
`gh run list --workflow=renovate.yml`, against `lint.yml` which answers the same query
with its last runs.

**It does not revisit the DCO.** Q-1 of XENO-0277 is decided — a person re-signs each
proposal — and an unsigned renovate pull request is the expected state. Dropping
`commitBody` does not touch that: the key never carried a sign-off, only the footer.

**It does not cover the Node and Python toolchain versions.** Those are major-only by
choice (`node-version: '24'`, `python-version: '3.14'`), resolved to the newest of the
line by `setup-node` and `setup-python`, and `docs/supply-chain.md` records them that
way. A manager proposing 25 or 3.15 would be proposing a decision, not an update.

<!-- xeno:section:context-rationale -->
## Why this context

Thirteen files, 103,892 bytes, against a budget set from those figures rather than from a
round number. XENO-0277 declared ten and resolved eleven, and carried two advisory
findings into every phase for it; the scope here is enumerated file by file instead of
reaching for `.github/workflows/**`, which would pull in `audit.yml` and `xeno.yml` for
nothing.

They fall into three groups, and the third is the one worth defending.

**What changes.** `renovate.json` and the `Dockerfile`. Both are edited, and the
`Dockerfile` only in the one line that carries the base image reference.

**What says whether the change is right.** `CONTRIBUTING.md` holds the footer rule that
decides whether `commitBody` should carry `Refs #95`, and it is the document the answer is
read off rather than argued from. `docs/supply-chain.md` is the list of what this project
pins deliberately, so it is the only place that says which pins a manager still has to
reach. `internal/model/supply_chain_test.go` holds that page against the tree in both
directions, and it is the test the `Dockerfile` edit could break — the digest it extracts
is matched by a pattern that stops at `:`, so a tag inserted before `@sha256:` is either
tolerated or it is not, and reading the pattern is how that is known before the edit
rather than after.

**The five workflows already reached by a manager.** `release.yml`, `trivy.yml`,
`lint.yml`, `govulncheck.yml` and `gosec.yml` are not edited and nothing proposed here
touches them. They are in scope because the claim this intent makes is a claim about all
of the managers, not only the new ones: that each resolves to the version the tree
actually carries. A regex manager that silently matches nothing is the failure mode being
fixed, so the six that already exist have to be shown still matching in the same run that
shows the three new ones matching. Six passing beside three is what makes a `NO MATCH`
evidence rather than a harness that never worked — which is the rule about negative
results, applied to the thing this intent is about.

`gitleaks.yml`, `docs.yml` and `renovate.yml` sit in the first two groups by their
contents: the first two hold the pins being reached, and the third is the workflow whose
`postUpgradeTasks` and schedule the configuration has to stay consistent with.

What is deliberately out. `.xeno/intents/**`, because the trail is the record and not
input. And `go.mod`, which XENO-0277 included: the `gomod` manager is untouched here, its
one module is already seen, and nothing in this intent reads or moves it.
