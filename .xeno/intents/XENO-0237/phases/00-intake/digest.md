---
intent: github.com/triplem/xeno#202
phase: 00-intake
created: "2026-10-03T19:24:08Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c7f6399e505aff7f89232c4eea39a470c6a7a70076d5d84f4fbb0c52c4fecaa9
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake fixes the shape of the deliverable before any clause is read, because the
shape is the decision: a list of clauses against their readers, with the fixing left
out. The four kinds come from #202's own caution that a clause whose reader is a person
is not unread, which rules out the simpler framing of read against unread.
