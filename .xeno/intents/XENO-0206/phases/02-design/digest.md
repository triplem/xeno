---
intent: github.com/triplem/xeno#132
phase: 02-design
created: "2026-09-29T16:15:50Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+1dc1483.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a16c233f5bf1af4d78060193466df38c0a1faeb886702abcd3e8cf851402d4ba
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The alternative worth the record is adding complete to section 5's status values, because it is what a
reader expects and section 8 argues against it directly: a stored flag would be a second answer to a
question the verdict already answers, free to disagree, and nothing recomputes a flag. The smallest
decision is that the new word is defined in the runner and not in the schema, which the design says outright
so that nobody later looks for it in section 5.
