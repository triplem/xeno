---
intent: github.com/triplem/xeno#95
phase: 05-review
created: "2026-10-08T14:52:47Z"
schema_version: "1.0"
runner_version: dev+5f12412.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 31a386dcd007e20afb557e2211881e305a7debba86d12aafd228e4b43d4214d4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: met
      note: 'The implementation is the design: the corrections named, the three customManagers, the workflow with its one permitted post-upgrade command, and the page row. The one thing the design left to a person -- the token -- is named as absent rather than assumed.'
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: No interface change. A configuration and a workflow change; no command, flag, artifact field or template moves, and the previous configuration had no behaviour to migrate from because nothing ever ran it.
    - rule: new-dependency-needs-a-rationale
      result: met
      note: 'One new pinned action, renovatebot/github-action, with its rationale in the issue and its row on the supply-chain page. No module dependency: go.mod is untouched.'
    - result: deviation
      note: 'Supply chain lens: this adds a third party with write access to the repository, by a token a maintainer issues, and the first thing it will do is propose moving pins that A28 says are pinned to what a release demonstrably used. The deviation is accepted by design -- no automerge, so every proposal is read -- but the tension is real and is named rather than resolved: an update bot proposes the newest and A28 wants the demonstrated.'
      source: lens
---

# Review

<!-- xeno:section:release-notes -->
## Release notes

**Renovate runs as this repository's own scheduled job.** `.github/workflows/renovate.yml`
runs `renovatebot/github-action`, pinned by sha like every other action here, on Monday
mornings and on demand. It proposes pull requests and keeps a dependency dashboard; nothing
automerges, and every proposal waits for a person.

The self-hosted shape is what allows `go mod vendor` to run after a `go.mod` bump, so a module
update arrives with `vendor/` in step rather than failing CI on a tree the bot left
inconsistent. That is the reason this is a workflow rather than an installed app.

**A proposal arrives unsigned, on purpose.** The Developer Certificate of Origin's claim stays
with a person: amend the commit with your own `Signed-off-by` before merging. An unsigned
renovate pull request is the expected state and not a broken one.

**It is off until somebody switches it on.** The job needs `RENOVATE_TOKEN` in the
repository's secrets. Until that exists it authenticates against nothing and proposes nothing.

Four of this project's pins live inside workflow text rather than in a manifest —
semantic-release's version, trivy's, cyclonedx-gomod's — and are reached by regex managers.
semgrep stays unmanaged because it is pinned by digest with no version to read. And a bumped
sha is two edits: the pin and its row in `docs/supply-chain.md`, which a test holds together.
