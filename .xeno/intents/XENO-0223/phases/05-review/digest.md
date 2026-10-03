---
intent: github.com/triplem/xeno#171
phase: 05-review
created: "2026-10-03T09:59:02Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 58870f0b0694383abe62e2c5111da1a71507a55d92a48d2057cc1e231fa25dd0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three shipped review rules are answered, and the question this review turns on is an acceptance
criterion this intent did not meet: a link naming a document that does not exist is skipped silently.
P4's mapping found it after P3 was sealed, no test had been written for it, and what is sealed is never
rewritten — so it leaves as a finding with a follow-up issue rather than being patched behind the phase
that owned it. Releasing this means accepting that. Nothing is recorded that section 5 does not
enumerate: two fields from its own list, and the changed set as output, because the section's list has
no entry for it and the file states what was declared rather than what was read. The demonstration
earned its place by producing three things no test did — section 5's example budget is too small for a
repository of this size by nearly a factor of two, the saving is reachable only after four approvals,
and a profile of that shape costs about three releases per intent of this kind. What a reader should
not conclude is that a repeated phase now reads only what changed: it is told what changed, the clause
is implementable and unverifiable, and what was built is the knowledge rather than the behaviour.
