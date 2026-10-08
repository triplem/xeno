---
intent: github.com/triplem/xeno#95
phase: 00-intake
created: "2026-10-08T14:38:37Z"
schema_version: "1.0"
runner_version: dev+5f12412
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2d0fe57bf21259ff40c12425220caa7dc1d0c58036e65a071721522135f5dd3e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
open_questions:
    - key: Q-1
      options:
        - consequence: |
            renovate.json carries commitBody with Signed-off-by: renovate[bot], every proposal is mergeable, and the DCO's claim that the signer has the right to submit the change is made by a process rather than by a person. Nothing checks sign-off today, so the bot would be the first writer nobody verifies.
          reason: |
            What a dependency bump certifies is the origin of somebody else's released artifact, which no human signature makes truer. The alternative leaves every proposal unmergeable, which makes the whole arrangement decorative.
          recommended: true
          text: Yes, renovate signs off as itself
        - consequence: |
            renovate opens the pull request and a person amends the commit with their own sign-off before it can merge. Every update then costs one human act, which is the thing automation was for, and the dashboard becomes a list of work rather than a list of proposals.
          text: No, a person re-signs each proposal
        - free: true
          text: Something else, in the words of whoever decides
      text: May a bot sign off under the Developer Certificate of Origin, so that renovate's commits carry Signed-off-by and are mergeable at all?
---

# Intake

<!-- xeno:section:problem -->
## Problem

> **github.com/triplem/xeno#95** — Replace dependabot with renovatebot
>
> Renovatebot does offer a dependency dashboard. This could be used instead of prs.

The issue as `xeno phase start` read it. What it does not carry, and what this intent is
really about, is the analysis in the issue's own comments and the three answers the
maintainer gave on it: **self-hosted, no automerge, dashboard plus pull requests.**

XENO-0274 shipped a `renovate.json` that contradicts two of the three and misses what the
analysis names as the substantive argument for renovate here. Against the answers given:

- **self-hosted** — nothing self-hosted was built. No scheduled workflow, no token, no
  `postUpgradeTasks`, so the file is the App shape by omission.
- **no automerge** — `"automerge": false`, which is right.
- **dashboard plus pull requests** — `prCreation: "approval"` with
  `dependencyDashboardApproval: true` on every rule is dashboard *instead of* pull requests:
  nothing opens until a person approves each entry.
- **the DCO** — not mentioned at all, so a bot commit would carry no `Signed-off-by` and
  `CONTRIBUTING.md` says a commit without one is not merged.

And the gap that matters most: **the versions written inside workflow text are invisible to
it.** `semantic_version: 25.0.9` in `release.yml`, trivy's `v0.74.0`, semgrep's image version,
`cyclonedx-gomod`'s version — none of those sit in a manifest, so the `npm` manager finds
nothing and the configuration reaches the action shas and `go.mod` alone. That is most of
what this project pins. `customManagers` with a regex over the workflow files is what reaches
them.

`vendor/` is the same shape from the other side: `gomodTidy` does not re-vendor, and only the
self-hosted form can run `go mod vendor` as a `postUpgradeTasks` step — which is the deciding
detail behind the answer "self-hosted".

<!-- xeno:section:scope -->
## Scope

This intent brings `renovate.json` to the answers given on #95 and adds what the self-hosted
shape needs:

- the configuration corrected: pull requests opened beside the dashboard rather than held
  behind per-entry approval, no automerge anywhere, and the DCO sign-off carried in
  `commitBody` once the open question above is answered;
- `customManagers` for the versions that live inside workflow text, which is most of what
  this project pins and what the plain managers cannot see;
- a scheduled workflow running renovate as this repository's own job, pinned by sha like every
  other action, with `postUpgradeTasks` running `go mod vendor` so a `go.mod` bump arrives
  vendored;
- the row in `docs/supply-chain.md` that the new pinned action needs, without which #317's
  tests fail.

What it does not do.

**It does not install or run anything.** The workflow exists and needs a token in a secret
that only the maintainer can add. Until then the job fails to authenticate and nothing is
proposed, which is the honest state of a repository that has decided the shape and not yet
switched it on.

**It does not answer the DCO question.** That is the open question this intake records, and
G-Questions will not let P5 be decided until it is resolved into a decision, a withdrawal or a
confirmed assumption. The `commitBody` line is written for the recommended answer and is the
one thing in the change that the other answer would delete.

**It does not switch dependabot off**, for the reason the previous intent gave: there is no
dependabot configuration here and whatever runs is a host setting no file reaches.
