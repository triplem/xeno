---
intent: github.com/triplem/xeno#176
phase: 02-design
created: "2026-10-03T11:35:42Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d45cd1ea2159b08b459e93f537886c91ae1ee3bb26b0178890fd18554ebc367a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`bytes` is an `int64` on `ContextFile`, `omitempty`, taken from the same walk as the hash so the two
describe one read of one file and cannot disagree about which version they saw. The check sums what the
lock recorded and measures nothing: a lock with no sizes has nothing to sum, and the condition is on
whether any entry carries a size rather than on the sum being zero — that one line is the whole of the
third criterion. A lock with some sizes and not others is judged on what it has, with nothing
estimated. And nothing in the gate path measures a file after this, which is a property of the path
rather than of the check: a verdict is computed from artifacts and hashes, and a size was the one
exception. Four of the five alternatives were refused by rules this project already has — a verdict
that differs on recomputation, a quiet pass replacing a loud finding, a rewrite of what an artifact
recorded, and two reads of one file — and the fifth, one total instead of a size per file, by the
specification clause and by what a reader comparing two locks can see.
