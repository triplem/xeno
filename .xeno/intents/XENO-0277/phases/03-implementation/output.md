---
intent: github.com/triplem/xeno#95
phase: 03-implementation
created: "2026-10-08T14:47:32Z"
schema_version: "1.0"
runner_version: dev+5f12412
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c52061df5af1c14b5706da12cb021f8e9753d8013c840d60d87503add8c8ab72
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

Three files.

**`renovate.json`, corrected to the three answers.** `prCreation: "approval"` and the
per-rule `dependencyDashboardApproval` are gone — those made the dashboard a gate rather than
a list, which is dashboard *instead of* pull requests. What is there now is
`dependencyDashboard: true` with pull requests opened on a Monday schedule,
`prConcurrentLimit: 3`, and `automerge: false`, which was already right.

`commitBody` carries `Refs #95` and **no sign-off**, which is Q-1's answer: a person re-signs
each proposal. The file's own `description` says so, so that whoever meets an unsigned pull
request reads it there rather than treating it as broken.

**`customManagers`, three of them, which is the substance the first attempt missed.** Four of
this project's pins live inside workflow text: `semantic_version: 25.0.9` in `release.yml`,
`version: v1.12.0` under the CycloneDX action, `version: v0.74.0` under the trivy action, and
semgrep's image digest. The two `version:` keys are matched recursively through the action's
own `uses:` line, because the bare key appears more than once per file and a regex that took
the first one would have proposed a Node version as a trivy release. Each was checked against
the file it reads: `25.0.9`, `v1.12.0`, `v0.74.0`.

semgrep is deliberately not managed: it is pinned by digest with no version to read, so the
docker manager would propose a digest and no human-readable bump. The `description` says that
rather than leaving its absence to be noticed.

`enabledManagers` drops `npm`, because there is no npm manifest here at all — the npm tools
live in workflow text and are reached by the regex instead.

**`.github/workflows/renovate.yml`**, the self-hosted run. `renovatebot/github-action` pinned
to a sha, `setup-go` before it so `go mod vendor` runs against the toolchain `go.mod` names,
and `RENOVATE_ALLOWED_POST_UPGRADE_COMMANDS` permitting that one command, because the action
refuses a command it was not told about. `postUpgradeTasks` in the configuration runs it with
`fileFilters: vendor/**`.

Nothing runs until `RENOVATE_TOKEN` is in the repository's secrets, which only a maintainer
can add.

**`docs/supply-chain.md`** gains the row for the new pinned action — and did so because the
test demanded it: `TestEveryPinInTheTreeIsOnTheSupplyChainPage` failed with *renovatebot/github-action
is pinned in .github/workflows/renovate.yml and no row of the page names it*. That is #317
catching the interaction XENO-0274 predicted, on the intent that corrects it.
