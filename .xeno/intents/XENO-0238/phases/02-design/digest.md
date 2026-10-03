---
intent: github.com/triplem/xeno#210
phase: 02-design
created: "2026-10-03T19:40:35Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 76ba9b4e2d9bbf3d53dcc085a77cbb8c7fc9db358f6a85673b1d032523707c0f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The split is the decision and the rest follows from it. The alternative worth recording is
the Go width check, which is cheap to write and expensive to agree: it would need a ruling
on long string literals and on which existing lines get rewrapped, so it is deferred as a
conversation rather than rejected as a bad idea. Keeping the number in `CLAUDE.md` and the
reason in `CONTRIBUTING.md` is the one place this intent adds to the approved wording.
