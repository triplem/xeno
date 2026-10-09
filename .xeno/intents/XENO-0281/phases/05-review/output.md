---
intent: github.com/triplem/xeno#343
phase: 05-review
created: "2026-10-09T13:24:41Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: be57fee88c30d8f89dc74c7ff3d0fc6efd02c1f8e6dc9567fc19fc0717d79db9
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: met
      note: P3 records no deviation and says why.
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: The check keeps its name and its policy; nothing outside the intent moves.
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: The same tool at a different commit, with the row saying why the commit.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**deviations-are-traceable — met.** P3 records none and says why.

**interface-change-needs-a-migration-note — not-applicable.** The check keeps its name and its policy; a pull request sees it turn green.

**new-dependency-needs-a-rationale — not-applicable.** The same tool at a different commit; the row says why the commit.

**What no rule asks.** Whether a pseudo-version on the pin table is a pin in the sense the page means: it is exact, checksummed and names a commit, which is what the action rows do with a sha.

<!-- xeno:section:release-notes -->
## Release notes

**gosec reads Go 1.27.2's export data again.** The scanner is pinned to the commit on its main branch that carries `x/tools` 0.51.0, as a pseudo-version, until release 2.30.0; policy, exclusions and the four accepted findings are unchanged. Every merge was red since 1.27.2 reached the runner this morning; the toolchain stays pinned by major, and the tool moved.

<!-- xeno:section:residual-risk -->
## Residual risk

**The next Go patch can do this again** to any tool built against an older `x/tools`; govulncheck and golangci-lint passed this time and are not immune. The remedy is the same: the tool moves. A check in the pipeline that a tool reads the toolchain's export data before the scan would say so in one line rather than through a wall of type errors, and nobody has asked for it.

**A pseudo-version pin outlives its release.** Renovate proposes 2.30.0; if nobody merges it, the row says a commit indefinitely, truthfully.
