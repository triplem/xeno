---
intent: github.com/triplem/xeno#95
phase: 03-implementation
created: "2026-10-08T13:40:07Z"
schema_version: "1.0"
runner_version: dev+042c9bc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 712f210d6a5ae60d551299ee40074a7ccf24ef86f36bed4dee6da26081df429a
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

One file added, `renovate.json`, and nothing else touched.

**Three managers and no more.** `github-actions`, `npm`, `gomod` — the three kinds of pin
this repository holds. Everything else renovate can detect would propose updates to
nothing.

**`prCreation: "approval"` and `dependencyDashboardApproval` per manager.** This is the
issue's own request — the dashboard instead of pull requests — and it is also what keeps a
proposal from failing CI on the wrong file: a bumped sha and its row in
`docs/supply-chain.md` move together, and only a person can move the row.

**The reason is in the file, in each rule's `description`.** Renovate reads those and so
does the next person to open it, which is the point: the constraint is not obvious from the
absence of automerge, and a reader who does not know about #317's tests would otherwise
switch automerge on and meet a red `verify` they cannot place.

**One exception, `vulnerabilityAlerts`.** It does not wait for the dashboard. An advisory
arrives without anything in the repository moving, and `CONTRIBUTING.md` already says the
work it names comes before the next merge rather than after the next feature.

**`pinDigests` for actions and `rangeStrategy: pin` for npm**, because this repository pins
rather than ranges, and a bot that proposed a range would be proposing a change of policy.

Nothing is bumped. No sha, version or digest in the tree moves in this commit, which is
acceptance criterion 4, and `go test ./internal/model/` is what says the pins and the page
still agree.
