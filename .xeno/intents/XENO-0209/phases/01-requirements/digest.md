---
intent: github.com/triplem/xeno#109
phase: 01-requirements
created: "2026-09-29T18:36:13Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+3ec2429.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c80696e6aa4ffc00a40210bbcffaebe364bc2d4169b8ef10a51240d26e403f80
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
AC3 is the criterion that stops the obvious shortcut of skipping directories entirely: the hash does not
descend into one, so a stray directory cannot corrupt the value, and the phase level check reports it anyway
because a directory nobody wrote is a place files will appear. AC4 is the one that keeps an ordinary
abandonment quiet, which matters because the finding lands in the only verdict such an intent ever gets.
