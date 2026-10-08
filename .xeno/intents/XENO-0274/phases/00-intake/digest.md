---
intent: github.com/triplem/xeno#95
phase: 00-intake
created: "2026-10-08T13:36:55Z"
schema_version: "1.0"
runner_version: dev+042c9bc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d4322d4f833c38bd066c03cfe35aee5892929c227585f62a269b33a10e188322
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake of #95, under this repository's reduced section set.

`phase start` read the issue and wrote its text into the problem section, so the phase
began with the problem as its author stated it and added what the tree says about it:
there is no dependabot configuration in this repository at all, so this is the first
in-repository statement of how dependencies are kept current rather than a swap of one
bot for another.

The fact the design has to answer: since #317 two tests hold every pinned sha and version
against the rows of `docs/supply-chain.md`, in both directions. A bot that bumps a pin and
leaves the page alone produces a pull request that fails `verify` on the page rather than
on the bump. That makes the dependency dashboard the issue asks for the interesting half
of it — a dashboard proposes and a person moves the row.

Scope excludes switching dependabot off, which is a host setting no file here reaches, and
excludes teaching a bot to edit a Markdown table. The context scope is the workflows, the
page, its test and `go.mod`: nine files, 91 KB.

Two sections written of the five the intake template defines, which is the reduced set
this intent is measuring: `context-rationale` was left out because the scope and the lock
already say which files were read and the sentence would have restated them.
