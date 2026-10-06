---
intent: github.com/triplem/xeno#260
phase: 00-intake
created: "2026-10-06T09:12:50Z"
schema_version: "1.0"
runner_version: dev+081da51
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e256d8656e0a16728fb828f6dd54e15be693bbbf84fddb5ce509cd8c84cee391
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The audit job does not install what the release installs. Measured on 2026-10-06: the audit's
tree carries 12 advisories and 0 critical, the release's carries 28 and 2 critical, and the
release's is a strict superset. The cause is that the release installs through an action which
runs `npm ci` against its own committed lockfile first, so 531 of its 532 packages are pinned at
the action's sha — the shipped tree does not drift, and the tree that drifts is the audit's own.
This intent makes the audit install the action's three steps and baselines the real figures.
