---
intent: github.com/triplem/xeno#95
phase: 04-verification
created: "2026-10-08T14:52:20Z"
schema_version: "1.0"
runner_version: dev+5f12412.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 04236458deb9133e1f8adc6ed73617dadab7ac44ea0e605a82673ffdb35baf7b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| criterion | how it was checked | result |
|---|---|---|
| 1. renovate runs as this repository's own job | `.github/workflows/renovate.yml` exists, parses as YAML, and pins `renovatebot/github-action` to `230ce922…` with its `# v46.3.7` comment, read off the tag today | met |
| 2. a `go.mod` bump arrives vendored | `postUpgradeTasks.commands` is `["go mod vendor"]` with `fileFilters: vendor/**`, and the workflow permits exactly that command through `RENOVATE_ALLOWED_POST_UPGRADE_COMMANDS` and runs `setup-go` from `go.mod` first so the command has a toolchain | met, by reading the configuration — renovate itself has not run, see below |
| 3. pull requests beside the dashboard, nothing merges itself | read out of the file: `dependencyDashboard: true`, no `prCreation` key at all, no `dependencyDashboardApproval` on either rule, `automerge: false` | met |
| 4. the versions inside workflow text are seen | three `customManagers`, each applied to the file it reads: `semantic_version` gives `25.0.9`, the CycloneDX `version:` gives `v1.12.0` and trivy's gives `v0.74.0`, each matched recursively through its own action's `uses:` line | met |
| 5. no commit claims a sign-off nobody made | `commitBody` is `Refs #95` and carries no `Signed-off-by`; the configuration's `description` and the workflow's header both say a person re-signs | met |
| 6. the new pinned action has its row | `docs/supply-chain.md` carries it, and `go test ./internal/model/` passes. It failed first: *renovatebot/github-action is pinned in .github/workflows/renovate.yml and no row of the page names it* | met |

<!-- xeno:section:results -->
## Results

Six criteria, six met.

    go test ./internal/model/      ok          #317's two pin tests
    go vet ./...                   clean
    gofmt -l . | grep -v vendor/   nothing
    json.load(renovate.json)       parses
    yaml.safe_load(renovate.yml)   parses

Criterion 6 is the one that behaved like evidence rather than like a claim: adding a workflow
that pins one action made `TestEveryPinInTheTreeIsOnTheSupplyChainPage` red, naming the action
and the file, until the row was written. That is #317 catching on this intent the exact
interaction XENO-0274 had only described in prose.

Criterion 4's regexes were each run against the file they read, and one of them was wrong on
the first attempt in a way worth recording: applied to the whole file rather than to the span
its first pattern matches, the bare `version:` pattern captured `node-version: '24'`. Renovate
chains a recursive `matchStrings` — the second pattern applies to what the first matched — and
checked that way both give the version they should. The near miss is the reason the two
managers are anchored on their action's `uses:` line instead of on the key.

**What was not checked, and the honest bound on all six.** Renovate has not run. It cannot:
the job needs `RENOVATE_TOKEN` in the repository's secrets, which only a maintainer can add,
and no run means no dashboard, no proposal and no `go mod vendor`. So criteria 1 to 5 are
checked against the committed configuration and the published schema, not against behaviour.
The first scheduled run after the token is added is the first evidence of behaviour, and the
first `go.mod` proposal is the first evidence that the vendor step works at all.
