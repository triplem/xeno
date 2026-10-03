---
intent: github.com/triplem/xeno#215
phase: 01-requirements
created: "2026-10-03T20:22:45Z"
schema_version: "1.0"
runner_version: dev+b626f1a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 817e5574286ffc47477569a73e4dcc4a8dde34c8af9ff3b719bf8d426475feef
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The criterion that shaped the code is that the refusal keys on the verdict rather than on
the lock: `Start`'s existing message tells a reader whose run died to remove the marker and
start again, and that path leaves a lock with no verdict, so keying on the lock would have
closed a documented recovery. The refusal also has to name the alternative, because one that
only says no is worked around by deleting whatever is in the way.
