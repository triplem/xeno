---
intent: github.com/triplem/xeno#343
phase: 00-intake
created: "2026-10-09T13:19:47Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cc24daf6fe3f9d889ef74dadd4e9ef63e76786005760b7c358945203d4255ab1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake for #343: gosec 2.29.0 cannot read Go 1.27.2 export data, every merge is red since this morning, and the fix is a pin moved to the commit on gosec main that carries x/tools 0.51.0, verified against this tree with the same four accepted findings and zero issues. Scope: the pin line, the pin table row, A104. Not Go, not renovate, not the policy. Key XENO-0281 because XENO-0280 is live on #342.
