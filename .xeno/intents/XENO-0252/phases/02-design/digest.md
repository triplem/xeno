---
intent: github.com/triplem/xeno#205
phase: 02-design
created: "2026-10-05T16:27:41Z"
schema_version: "1.0"
runner_version: dev+97d07da.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 786279273cf1c46c69a6777c831b41a65405c774cb1e7fdccb6444211a948e4b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`LocalDir` and `LocalPath` in `internal/model`, the only package all six callers import and one
that imports nothing but the standard library, so `cost` reaches it without depending on the
runner. Absolute values are taken as given, because the entry point exports an absolute path;
empty reads as unset, because resolving it to the root would put the ledger at the top of the
tree.

`Ledger` and `ReportPath` become file names rather than paths, and the register records why this
variable is read where `XENO_PLUGIN_ROOT` was removed.
