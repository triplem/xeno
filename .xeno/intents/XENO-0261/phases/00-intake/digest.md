---
intent: github.com/triplem/xeno#237
phase: 00-intake
created: "2026-10-06T12:02:20Z"
schema_version: "1.0"
runner_version: dev+7885661.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bc38359aca6d5e793eb2fc4275850dd774c78530dced031448c31fa01b4a1a68
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Three staleness remedies tell a person to read the phase again, which `Start` refuses for a judged
phase, and the first route the refusal offers — `section set` and `phase finish` — does not clear the
finding, because the lock is written only by `phase start` and #215 refuses a second one. Section 7
already names the two routes that work: re-run, or explicitly approve as still valid. #236 is closed,
so the findings the check raises are now only the kind that instruction fits. The unreadable-file
remedy already works and stays, because that fault is the machine's and not the trail's.
