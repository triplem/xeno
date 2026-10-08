---
intent: github.com/triplem/xeno#95
phase: 05-review
created: "2026-10-08T13:44:01Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 861a20ece48f7912da41534445cdeab41944654c3170e3ac8a8956fcf4d2944c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: not-applicable
      note: 'The implementation follows the design without deviation: the design named three managers, the dashboard and the per-rule reason, and renovate.json carries exactly those.'
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: No interface changes. A configuration file is added and no command, flag, artifact field or template moves.
    - rule: new-dependency-needs-a-rationale
      result: met
      note: 'No module dependency is added: go.mod is untouched. Renovate is a host app rather than a dependency of the build, which is the distinction this rule is about, and the rationale is the issue.'
    - result: met
      note: 'Supply chain lens: the change adds no code and no dependency, and its one risk is the one it names — a bot proposing a pin bump that CI then refuses on the page. Routed to a dashboard a person works from, which is the mitigation and not a deferral.'
      source: lens
---

# Review

<!-- xeno:section:release-notes -->
## Release notes

**Dependency updates are proposed on a dashboard, not as pull requests.** `renovate.json`
configures renovate for the three kinds of pin this repository holds — the action shas in
`.github/workflows/`, the npm tools the workflows install, and the one Go module — and
routes every proposal to the dependency dashboard for a person to take, rather than opening
a pull request.

The reason is this repository's own arrangement and it is written in the file: every pinned
sha and version is also a row of `docs/supply-chain.md`, two tests hold the two against each
other in both directions, and renovate cannot move a Markdown row. A pull request that bumped
a pin alone would fail CI on the page rather than on the bump.

One exception: a vulnerability alert does not wait for the dashboard, because an advisory
arrives without anything in the repository moving.

Nothing in this change bumps a pin, and switching dependabot off remains a host setting
that no file here reaches.
