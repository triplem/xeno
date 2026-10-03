---
intent: github.com/triplem/xeno#169
phase: 02-design
created: "2026-10-03T09:02:42Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+93322c5.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 506e0d791a8d9de24c73df72315ae8858cbc343c3732c04ba4e2c0d7eec48f0b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The distribution root in this repository is `.xeno/plugin/`, because section 13's tree and that
directory are the same set of names and `--plugin-from` already points there; so `skills/` and the
manifest join the templates, the rules and the filter. Two paths diverge from the document and both are
forced by the client: the plugin manifest sits in `.claude-plugin/` inside the tree, and the marketplace
wrapper sits at the repository root outside it, because that is where a client adding this repository
looks. Section 13 describes a third party's format and the third party disagrees, which is the one case
where "the specification wins" cannot be followed as written — a literal tree does not load. A77
records it and the document is a person's to correct. Seven skills with section 13's names, each the
same five things: what the phase is for, the sections it owes, the commands in order, what its gate
refuses, and what it hands on. A skill names commands and never operations, because there is no server
and because the command path has to keep working. The hook comes from the plugin and calls `xeno` on
the path, so an installing project records cost without writing a settings file. Rejected: shipping a
tree that does not load, moving three directories to the repository root, folding learning into six
skills, slash commands, and a stub `mcp.json`.
