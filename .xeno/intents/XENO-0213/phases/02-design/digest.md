---
intent: github.com/triplem/xeno#145
phase: 02-design
created: "2026-09-30T05:45:26Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4e41c6b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c9c2f1629f6b4237b0927638ade8126dea77f9a250d76efe2d642794ec9bd89e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
WrapperHost gains Adapter and APIBase, so one row describes a host completely: the range expressions and the
tracker address are the same fact arriving at two files, and splitting them is how they drift. scaffold.Project
gains the same two, and the scaffold's literals become fields. --host loses its default in cmd/xeno and in
init.go, which both filled it in — the string github was in three places for one reason and the flag was the
only visible one. WrapperHosts derives from the table instead of repeating it. The GitLab wrapper is scoped to
merge_request_event because the base SHA exists nowhere else. The closest rejected alternative was letting the
adapter export its canonical address as a fallback, which is a host default in code and is silently wrong for a
self managed deployment.
