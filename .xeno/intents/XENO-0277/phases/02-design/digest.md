---
intent: github.com/triplem/xeno#95
phase: 02-design
created: "2026-10-08T14:43:42Z"
schema_version: "1.0"
runner_version: dev+5f12412
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2e41f4e9338b818bd734026e243e95ad9b4b0bf3439e4e28dfc2ac53050fa1df
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Three files: the configuration corrected to the three answers, a workflow that runs renovate as
this repository own job, and the page row the new pinned action needs.

Self-hosted is a workflow and not a setting, and the deciding detail is `go mod vendor`: the App
cannot run a command in the repository, so a `go.mod` bump from it arrives with `vendor/` stale
and fails `verify` on a tree the bot made inconsistent. The action form runs `postUpgradeTasks`,
and it has a second property this project cares about — the run is a job here, pinned by sha,
and therefore a row of the supply chain page like everything else.

Dashboard plus pull requests means removing the two things the previous configuration had,
`prCreation: approval` and `dependencyDashboardApproval`, which turn the dashboard into a gate
rather than a list. Q-1 answer removes the `commitBody` sign-off and replaces it with the
sentence that a proposal arrives unsigned and a person re-signs it.

`customManagers` is the substance: four of this project pins live inside workflow text rather
than in a manifest, and without a regex manager renovate sees the action shas and `go.mod` alone.

One section of the four, `impact`.
