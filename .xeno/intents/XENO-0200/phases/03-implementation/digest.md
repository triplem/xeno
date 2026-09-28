---
intent: github.com/triplem/xeno#118
phase: 03-implementation
created: "2026-09-28T20:15:46Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57e9207.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 9c382f5800735d9bbee76a2276fa7ee7ccc2390182a13d25be29bb343eb76ad2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The first version of the listing reproduced the bug it was written to fix: created was
truncated to a date before the sort, so two intents of one day fell through to the key
tie break and XENO-0107 came out before XENO-0108. AC3 named that exact pair, which is
the only reason it was caught, and both the runner and the command now carry a comment
saying why the whole value is sorted and the date printed.
