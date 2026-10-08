---
intent: github.com/triplem/xeno#324
phase: 03-implementation
created: "2026-10-08T13:48:08Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cb41096967a075708e1b6fec705bb8cd6215d710600633709ba3a9594ac0684f
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

Two files.

`internal/scaffold/files/ci-github.yml` — the job's `container:` block gains a commented
`options: --user 1001` and six lines of comment above it saying why: the host mounts its own
work directory in and chowns nothing, `actions/checkout` runs inside the container, a hosted
runner's directory belongs to uid 1001 and the image runs as 1001, and no setting inside the
image can fix a write where those differ.

`internal/runner/init_test.go` — `TestOnlyTheGitHubWrapperOffersTheContainerUser`, beside the
test that asserts every wrapper names the image. It asserts three things: the GitHub wrapper
carries the commented line; it does *not* carry an uncommented `options:` key, because a uid
this repository wrote would be right on one runner and wrong on every other; and no other
host's wrapper mentions `--user` at all.

Nothing else. The GitLab template is untouched and the Dockerfile is untouched: the image
already runs as 1001 and already carries the `safe.directory` exemption, which answers the
other half of the uid problem and cannot answer this one.
