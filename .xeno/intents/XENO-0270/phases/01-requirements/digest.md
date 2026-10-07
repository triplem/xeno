---
intent: github.com/triplem/xeno#231
phase: 01-requirements
created: "2026-10-07T09:39:06Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7bdda5fa86f0286f5e20c8ad0e133dd87040477b99a644f6e2b99285af6e30bd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Twelve criteria. Four are the README's opening: what the tool is for, the problem stated before
the tool, the name in a sentence pointing at Appendix C, and the index linked with the
normative document named at the point of linking. Two are the index page: it exists and every
file under `docs/` appears with what it is for. Three are the reference: the README names the
handful a reader starts with, `docs/commands.md` carries a block byte-identical to `usage`, and
a test fails when they diverge, checked by making them diverge. Three are the rest: the moved
material is all present, nothing is lost but the no-network-call paragraph which section 12
already carries, no normative document changes, and the suite plus every relative link
resolves.

The non goals are the restraint: no normative edit, no reference generated from the flag
declarations, which is WP16's, no site, no shortening of `usage`, no touching the skills' seven
short lists, no new facts, no re-measuring the coverage table on the way past, and no link
checker in CI.

The binding constraints are the first standing rule, writing the README rather than patching it
while moving the migrating paragraphs verbatim, and keeping `usage` the single place a command
name is written down.
