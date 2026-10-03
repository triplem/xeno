---
intent: github.com/triplem/xeno#176
phase: 01-requirements
created: "2026-10-03T11:34:42Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a8207c073efa599bf222b3e706517cedbea311307d9f31eef834c306961594aa
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
A lock written from now on carries a size per file, taken from the same walk that took the hash so the
two describe the same read. The byte budget is judged against those recorded sizes, so a file that grows
after a phase is sealed does not change that phase's verdict — the property the specification clause
was written for and the only reason this intent exists. A lock that records no sizes produces no byte
finding: not a finding against zero and not a pass claiming the context was inside its budget, because
seventy-nine intents' worth of locks record nothing and reading them as zero-byte contexts would be a
quiet pass where there used to be a loud one. A lock with some sizes and not others is judged on what it
has, which the runner cannot produce and a hand-written lock can. The file-count budget is unchanged,
having always been judged against the lock. Non-goals: no backfill, since a lock rewritten to carry a
number nobody recorded would describe a reading that never happened; no second use of the size, because
the hash is what detects change; and no measurement anywhere in the gate path after this.
