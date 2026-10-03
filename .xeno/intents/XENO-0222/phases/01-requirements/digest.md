---
intent: github.com/triplem/xeno#169
phase: 01-requirements
created: "2026-10-03T09:01:28Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+93322c5.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 65e8fec5334d47a343f367ee8db23932901eaf4524a8662ad752704144e97ece
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Both manifests validate against the client, with no warning left unaddressed, and the marketplace is
added and read back rather than asserted. Seven skills with the names section 13 fixes — six phases
plus `xeno-learning`, not six with learning folded in. Three criteria are mechanical checks on prose:
every command a skill names exists in the dispatch table, every section it asks a phase to write is in
that phase's template, and every gate it mentions is in the table and not one reporting
`not-implemented`. Nothing in a skill only makes sense inside this repository, which is the bar
`given/builtin/` has, and each skill names the command path rather than an MCP operation, because
there is no server and because the command path is what WP11's done-when requires to work alone. The
hooks come from the plugin, so a project that installs it records cost without writing a settings file
— today this repository's own `.claude/settings.json` is the only place that knows. `init --vendor`
carries what exists of section 13's list and invents nothing. Non-goals: no `mcp.json` without a
server, no lenses, no second harness, no slash commands or subagents, no edit to `docs/`, and no
closing of M0 — the blocker goes, the reading is yours.
