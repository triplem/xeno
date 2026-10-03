---
intent: github.com/triplem/xeno#183
phase: 05-review
created: "2026-10-03T17:33:59Z"
schema_version: "1.0"
runner_version: dev+9140756.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a6884dd47b63307b54fc9f91659ebf97baf0581faaa428c1cb262eef500dc136
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three shipped review rules are answered. `deviations-are-traceable` with three: the ldflags line
split because one line with both `-X` paths runs to 101 columns, `ASSUMPTIONS.md`'s stale gate list
corrected for G-Rules and G-Policy as well, and two sealed verdicts rewritten and restored.
`interface-change-needs-a-migration-note` because a gate that reported `not-implemented` now reports
a judgement — a changed vendored plugin becomes red on every phase from P0 for anybody on a released
binary, where it was silent — and nothing migrates, since sealed verdicts keep what they recorded and
`gate verify` stays at exit 0 because it compares a phase's status and both results derive to green.
Beyond the rules: the gate does what section 13 asks and only that, and the test that proves it is
the one asserting the anchor is not readable from the tree, because the lock's `plugin` block added
by the previous intent was the most plausible wrong answer. Nothing was invented — the no-anchor case
reuses `not-implemented`, which section 5 defines for exactly this. A42 is intact. And the
recommendation this came from holds conditionally: the containment Shape 3 would rely on is
self-checking for a released binary and inert under `go build`, which is most of the work done here,
and that belongs in front of section 7's decision rather than behind it. The verification phase found
the largest gap while writing it: a released binary cannot vendor the plugin it carries the digest
of, which is a distribution question and wants its own issue.
