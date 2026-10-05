---
intent: github.com/triplem/xeno#201
phase: 03-implementation
created: "2026-10-05T12:55:08Z"
schema_version: "1.0"
runner_version: dev+325b27a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5d5810587b460276bb638ecc34590b2e1526c4cd9b7d6242062dda7114583ff1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`Hash` becomes a wrapper over `HashTree`, which keys each line on the path relative to the tree
root, so a tree hashes to what it is. `treeLines` is shared with `Differences` so the
comparison inherits the hash's normalisation rather than reimplementing it.

The release sequence was rehearsed locally: a stale file and a one-byte change both refuse, and
both name the path. This phase was opened before the work, which is the first time in four
intents that its elapsed figure measures anything.
