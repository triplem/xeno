---
intent: github.com/triplem/xeno#21
phase: 00-intake
created: 2026-09-24T18:31:17Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 066d619ee6c01772fae3ed17db2ccbb4741835337ada9c51f0afc6127e8c7625
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #21. Three of the four pinned actions are behind their current major, and the
upgrade was deferred when they were pinned so that a failure afterwards would have one
answer rather than two.

## Scope

`actions/checkout` to v7.0.1, `actions/setup-go` to v7.0.0,
`cycjimmy/semantic-release-action` to v6.0.0, each to the commit sha of that exact tag,
with `SUPPLY-CHAIN.md` following.

## Non goals

`CycloneDX/gh-gomod-generate-sbom`, which is at v2.0.0 and current.

**semantic-release itself.** It stays at 24.2.9 while the action around it moves to v6.
The action's `semantic_version` input exists to choose it, and moving both would put this
change back in the position it was created to avoid: two majors and one answer.

## What this cannot prove

`checkout` and `setup-go` run in the verify workflow, so a pull request exercises both
and the Node 20 deprecation should disappear from the summary.

`semantic-release-action` runs only in the release workflow, and only on a commit that
produces a release. The action with the most surface is therefore the one whose upgrade
stays unproven longest. Recorded here rather than found later.
