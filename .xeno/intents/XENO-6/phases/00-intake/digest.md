---
intent: github.com/triplem/xeno#6
phase: 00-intake
created: 2026-09-23T19:23:18Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: b7c654cd8b2756ef97bb796e56f3d0e66abc4b89ec088104f2f82ae8c1beae37
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The session was asked what was left of #6 and found more than the issue described: no
action is pinned, `cycjimmy/semantic-release-action@v4` resolves to a branch rather than
a tag, and the versions of semantic-release and its plugins are decided by npm at run
time. It also found that this was a regression introduced four commits earlier, when
`scripts/sbom.sh` and its `CYCLONEDX_GOMOD_VERSION=v1.9.0` were replaced by an action
without the pin being carried across.

Versions were then read rather than chosen: the workflow log of the v0.4.0 run for the
action shas and semantic-release 24.2.9, and the published bill of materials for
cyclonedx-gomod v1.12.0, which it records together with its own hashes.

No secret filter exists, so nothing filtered this text.
