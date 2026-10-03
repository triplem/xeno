---
intent: github.com/triplem/xeno#171
phase: 00-intake
created: "2026-10-03T09:43:13Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e9bd2f93f760341fa7dc3e3041d0e9c62c6f10bbde0d79a20aff4a7dcb3e2388
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
WP8's done-when has two clauses. "Every read outside the profile appears in `context.lock.yaml`" holds
in the only sense section 5 allows, since the lock states what the phase was given and the section says
treating it as a measurement of what was read would be wrong. "A repeated phase reads only what
changed" has no mechanism at all: nothing computes what changed, nothing reports it, and a repeated
phase is told what a first run is told. Reading the section for it found four fields section 5 writes
out and the lock does not carry — `repo_commit`, `rules_applied`, `plugin` and `tools` — two of which
became writable only this month, when #165 shipped a rule set and #162 brought git into the tree. And
three things the profile specifies that nothing reads: the budget, whose own section warns against it
becoming decoration, which is what happened; the links, declared and ornamental, since a declared
document is not in the base unless `include` matches it; and the assembly order, where the section
asks for volatility and the runner sorts alphabetically. None of it is exercised here, because
twenty-three intents have written no profile — the position the rule engine was in before #165 and the
external gate invariants before #167.
