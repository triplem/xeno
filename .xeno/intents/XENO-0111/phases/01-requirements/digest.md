---
intent: github.com/triplem/xeno#111
phase: 01-requirements
created: "2026-09-28T17:32:20Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0e77cc2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 580ecc08eed0854893b081a8ad6e3a58c8ca84b3a0cc407598e016cb7a785144
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Seven criteria for two small fixes, and five of them are about what stays true. AC5 and
AC6 are the ones that constrain the design: a walk must not turn a directory into
something that has to be declared, and the flat layout Attach writes has to produce the
verdicts it produces today. The constraint that decided the implementation is smaller
than either — this package names its findings variable fs, so io/fs cannot be imported
here.
