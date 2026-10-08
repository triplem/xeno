---
intent: github.com/triplem/xeno#95
phase: 01-requirements
created: "2026-10-08T14:42:59Z"
schema_version: "1.0"
runner_version: dev+5f12412
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 18cd77dcb86e4d32410b44f208c4df21fddac0242c7e65848d1a08e6521a21ee
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Six criteria, and the decision that was owed.

Q-1 is resolved: a person re-signs each renovate proposal and the bot does not sign off. The
DCO claim is about the right to submit, the maintainer keeps that claim with a person, and
nothing enforces sign-off today so a signing bot would have been the first writer nobody
checks. The cost is accepted — one human act per update, and an unsigned pull request is the
expected state rather than a failing one.

The criteria follow the three answers one by one, plus the two things the previous intent
missed: the `customManagers` that reach versions written inside workflow text, which is most
of what this project pins, and the supply-chain row the new pinned action needs or #317 tests
fail.

One section and one decision, of the four this template defines.
