---
intent: github.com/triplem/xeno#95
phase: 02-design
created: "2026-10-08T14:43:31Z"
schema_version: "1.0"
runner_version: dev+5f12412
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2e41f4e9338b818bd734026e243e95ad9b4b0bf3439e4e28dfc2ac53050fa1df
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:impact -->
## Impact

Three files: the configuration corrected, a workflow that runs it, and the page row the new
pinned action needs.

**Why self-hosted is a workflow and not a setting.** The deciding detail is `go mod vendor`.
Renovate's App cannot run a command in the repository, so a `go.mod` bump from the App arrives
with `vendor/` stale and fails `verify` on a tree the bot made inconsistent. The action form
runs `postUpgradeTasks`, which is the only shape that can re-vendor — and it has a second
property this project cares about more than it expected: the run is a job here, pinned by sha,
and therefore a row of `docs/supply-chain.md` like everything else.

**`renovatebot/github-action@230ce922b08968d0a4f6f70f295601daff09ef1d`, v46.3.7**, read from
the tag today. `allowedPostUpgradeCommands` has to permit `go mod vendor` explicitly, because
the action refuses an unlisted command rather than running it.

**What the three answers become.** Dashboard plus pull requests is `dependencyDashboard: true`
with no `dependencyDashboardApproval` and no `prCreation: approval` — both of which the
previous configuration had, and both of which turn the dashboard into a gate instead of a
list. No automerge is `automerge: false`, which was already right. Self-hosted is the workflow.

**What Q-1's answer removes.** No `commitBody` carrying a sign-off. The workflow's comment and
the dashboard's own text say instead that a proposal arrives unsigned and a person re-signs it
before merging, so that an unsigned pull request reads as the expected state rather than as a
broken one.

**`customManagers`, which is the substance.** Four of this project's pins live inside workflow
text rather than in a manifest: `semantic_version`, trivy's version, semgrep's image version
and `cyclonedx-gomod`'s. A regex manager over `.github/workflows/*.yml` reaches them; without
it renovate sees the action shas and `go.mod` and nothing else.

**The token.** `RENOVATE_TOKEN` in a secret, which only the maintainer can add. Until it is
there the job fails to authenticate and nothing is proposed — the honest state of a repository
that has decided the shape and not switched it on.
