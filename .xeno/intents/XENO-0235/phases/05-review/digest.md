---
intent: github.com/triplem/xeno#199
phase: 05-review
created: "2026-10-03T18:25:56Z"
schema_version: "1.0"
runner_version: dev+a897f2a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a99f2eca6567f7882c45d8f3b507c8c418f32d1ec829a4c706b7e24d226c7c15
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three shipped review rules are answered, and the migration note is the substance: `init --vendor`
changes what it copies, from a named list to everything, and where it copies from, preferring the
binary over a directory — while an existing repository is untouched, because `create` keeps what
exists, which is the property that made the walk safe to substitute. Beyond the rules: the gate is
satisfiable now and was not when it shipped one change ago, measured rather than argued, and that
state is worse than unimplemented — `not-implemented` says a check did not run, where a failing gate
blames the adopter for a tree the release could not deliver. The ordering is recorded between the
workflow's steps rather than in a document, because its failure lands on somebody else: a digest
taken before the stamp passes for the release and fails every project that installs it. Nothing was
invented; `PluginFrom` is a field of an in-memory struct that `init` prints. The verification phase
found what the design had not — the embed target is tracked, so a failed release could leave a
generated tree in a dirty checkout — and it is fixed here, before shipping rather than after. For the
decision this was done in service of: the anchor is now deliverable and the gate satisfiable for an
adopter, and under a development build the gate is still inert, which was said before and is
unchanged. No verdict behind this can change: 285 at exit 0.
