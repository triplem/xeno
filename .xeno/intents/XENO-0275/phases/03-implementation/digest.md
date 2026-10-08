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
---
Two files: six lines of comment and one commented setting in the GitHub wrapper template, and
a test beside the one that already asserts every wrapper names the image.

The test asserts the absence as well as the presence — no uncommented `options:` key, and no
`--user` in any other host wrapper — because both failure modes are a later reader being
helpful: filling the uid in for everybody, or adding the line to GitLab by symmetry.

One section of the three, `changes`.
