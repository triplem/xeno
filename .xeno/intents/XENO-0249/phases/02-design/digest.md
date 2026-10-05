---
intent: github.com/triplem/xeno#201
phase: 02-design
created: "2026-10-05T12:49:29Z"
schema_version: "1.0"
runner_version: dev+f2f65d5.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b377d6a2765e66ff3555636654789a1798ffd07ec543c6f1cab6724197eff685
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
One definition with two entry points: `HashTree(dir)` is the primitive, `Hash(root)` a wrapper
over `root/.xeno/plugin`. The basis changes rather than a second scheme being added, because
two definitions of tree equality is the thing being removed.

The comparison lives in the script, since refusing a release is the release's business; the
flag is opt-in because a development tree carries no embedded copy by design.
