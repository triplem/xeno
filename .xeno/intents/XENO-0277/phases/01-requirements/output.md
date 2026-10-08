---
intent: github.com/triplem/xeno#95
phase: 01-requirements
created: "2026-10-08T14:42:49Z"
schema_version: "1.0"
runner_version: dev+5f12412
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 18cd77dcb86e4d32410b44f208c4df21fddac0242c7e65848d1a08e6521a21ee
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      resolves: Q-1
      chosen: A person re-signs each renovate proposal; the bot does not sign off
      rationale: 'The DCO''s claim is about the right to submit a change, and the maintainer keeps that claim with a person rather than delegating it to a process. Nothing enforces sign-off today, so a bot signing would have been the first writer nobody checks. The cost is accepted: one human act per update, and an unsigned pull request in the queue is the expected state rather than a failing one.'
      decided_by: Markus M. May
      proposed_by: Claude Opus 5
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. **Renovate runs as this repository's own job.** A scheduled workflow runs renovate, pinned
   to a commit sha like every other action in this repository, and its run appears in this
   repository's Actions log rather than on a third party's schedule. That is the self-hosted
   shape the maintainer chose, and the deciding detail behind it is `postUpgradeTasks`.

2. **A `go.mod` bump arrives vendored.** The workflow's `postUpgradeTasks` runs `go mod vendor`,
   so a proposal that moves the one module dependency moves `vendor/` with it. A proposal that
   left `vendor/` behind would fail `verify` on a tree the bot had made inconsistent.

3. **Pull requests are opened beside the dashboard, and nothing merges itself.** The dashboard
   lists everything, pull requests are opened without per-entry approval, and `automerge` is
   false everywhere. Those are the maintainer's three answers, and a reader of
   `renovate.json` can check each against the file.

4. **The versions inside workflow text are seen.** `customManagers` reach
   `semantic_version`, trivy's version, semgrep's image version and `cyclonedx-gomod`'s
   version, none of which live in a manifest. Without this renovate proposes updates for the
   action shas and `go.mod` alone, which is a minority of what this project pins.

5. **No commit claims a sign-off nobody made.** The decision recorded against Q-1 is that a
   person re-signs each proposal, so `renovate.json` carries no `Signed-off-by` and the
   workflow adds none. The arrangement is written where whoever meets an unsigned pull request
   will read it.

6. **The new pinned action has its row.** `docs/supply-chain.md` carries the renovate action's
   sha and version, and `go test ./internal/model/` passes — which since #317 is what says the
   tree and the page agree.
