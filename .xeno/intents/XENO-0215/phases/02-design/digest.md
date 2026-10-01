---
intent: github.com/triplem/xeno#151
phase: 02-design
created: "2026-10-01T13:53:29Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c54db93.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 489e8a546322002d5f3c16a101919fc191776abe0396a2ec24701faf471395de
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
YAML, one file, provenance first, because provenance decides whether the symbols are read at all and a reader that
parses a hundred thousand symbols to discover the index is stale has done the work it was avoiding. kind and
container are free strings and not an enumeration: once #149 put the choice of indexer on the project, Xeno cannot
enumerate what that tool thinks a symbol is, and what makes the open set safe is that nothing compares the value.
Load returns a nil index and a reason, never an error, so four causes have one outcome and the signature says that
absence is ordinary; now is a parameter because staleness is the one boundary worth testing on both sides. Lookup
is exact on name, since anything looser is a search over the index and a search is what the index replaces. The
closest rejected alternative was a line oriented format, cheaper to produce and read, and the first file here that
would not be YAML.
