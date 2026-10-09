---
intent: github.com/triplem/xeno#346
phase: 02-design
created: "2026-10-09T15:11:56Z"
schema_version: "1.0"
runner_version: dev+9590797
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a2ff9f917f79cf2c442457c46827e3458da43747b185ecf193d9c5f4e0311b27
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The design for #346: the GitHub read asks by name with --search and keeps the exact match, the GitLab read takes --per-page 100 with the bound on the line, the plugin test adds the two flags to its strings, the proof is a real second run recorded in P4, and no register row because the choice is a tool fact.
