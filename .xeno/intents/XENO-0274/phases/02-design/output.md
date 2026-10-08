---
intent: github.com/triplem/xeno#95
phase: 02-design
created: "2026-10-08T13:39:22Z"
schema_version: "1.0"
runner_version: dev+042c9bc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b1b439ff8a0db4d5c60242771dfa9990da41903e6b5351affffd1e0ba2dd50a5
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

One file is added, `renovate.json`, and nothing existing changes. What it tells renovate:

**Three managers, because this repository holds three kinds of pin.** `github-actions` for
the shas in `.github/workflows/`, `npm` for the tools the workflows install, and `gomod`
for the one module dependency. A manager for anything else would propose updates to
nothing.

**The dashboard, and no pull request of its own accord.** `dependencyDashboard` is on,
`automerge` is off everywhere, and nothing is scheduled. The reason is this repository's
own arrangement rather than a preference: since #317, two tests hold every pinned sha and
version against the rows of `docs/supply-chain.md`, and renovate has no reader for a
Markdown table. A pull request that bumped a sha would therefore fail `verify` on the page,
not on the bump, and the person reading that failure would be debugging the wrong file.

**The consequence is written in the file.** `renovate.json` carries the sentence that a pin
and its row move together, so whoever works the dashboard reads it there rather than
discovering it from a red check.

What this does not affect: no pin moves, so `go test ./internal/model/` answers the same
as before, and the trail's verdicts are untouched because nothing they cover changes.
