---
intent: github.com/triplem/xeno#151
phase: 00-intake
created: "2026-10-01T13:52:27Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c54db93.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8ea5fe383f0ff29112adcff6c0877726977491462547ab619636ce789ca65f17
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

#149 changed what WP15 is. Xeno no longer builds the index: section 5 now says it "is produced by the
project and read by Xeno, which ships no indexer", and that "what Xeno fixes is the format of the
answer and not the means of getting it". So the format is the deliverable, and until it exists there is
nothing for a project to write and nothing for Xeno to read.

Nothing in the tree reads an index or knows what one looks like. `index.path` and
`index.max_age_hours` are specified in section 5 and read by no code, which is the state
`index.grammars_dir` was in before #149 replaced it: a field the specification carries and the runner
has never looked at.

Three things the documents now require and the tree cannot express.

**The content.** Symbols with name, kind, file, line and enclosing container, and explicitly not call
relationships or type resolution. That is a shape, and no type in `internal/model` or anywhere else
holds it.

**The provenance.** #149 made this a requirement rather than a convenience: the index "names the tool
and version that produced it and the time it was produced, because a stale index is worse than none and
an index that cannot say how old it is cannot be judged." Nothing can judge an age it cannot read.

**The degradation.** Both documents call this the property everything else follows from — the index may
be missing, stale or wrong without the trail suffering, because a verdict never depends on it. An
absent index is already the state every repository is in, so the code path that has to work first is
the one where there is nothing to read.

There is also a smaller problem that blocks the acceptance rather than the work. WP15's done-when asks
for "an index this repository produces for its own language", and nothing produces one. This repository
is Go, so that is a producer this project writes for itself, and it is the only way the format gets
tested against something other than a fixture written by whoever wrote the reader.

<!-- xeno:section:scope -->
## Scope

**In scope.** `internal/index`: the format as a type, a reader for it, its provenance, and a lookup by
symbol name. The degradation path, where absent, unreadable, malformed or stale all become no index and
no error. Reading `index.path` and `index.max_age_hours` from `project.yaml`. A producer in `scripts/`
for this repository's own Go source, built on `go/parser` and `go/ast`, so the format is exercised
against something a tool wrote rather than a fixture.

**Out of scope, and each for its own reason.**

The `tools` entry in `context.lock.yaml`. It needs a decision this intent must not take: `tools` is
specified as "name, version, response hash" and Appendix B defines no such hash, which A62 shows is
delegated to the package that writes it. Worse, "the hash of its answer" has two readings — the index
Xeno read, or the response a query returned — and they answer different questions. That decision and
its byte definition belong to the piece that writes the record.

The MCP operation. `ASSUMPTIONS.md` lists the MCP server under "Not built", so there is nothing to add
a sixth operation to. The lookup here is a function, and what makes it the sixth tool is WP11's.

An indexer inside Xeno. That is what #149 removed and this intent must not reintroduce. The producer in
`scripts/` is this repository's own tooling, beside `github-settings.sh` and `module-set.sh`, and the
distinction is stated in the acceptance because a reader finding a Go indexer in the tree could
reasonably think #149 was reversed.

The saving. #149 moved the figure to WP20 because the runner cannot observe what the agent reads. This
intent measures nothing and claims nothing about tokens.

<!-- xeno:section:context-rationale -->
## Why this context

The information base is the two documents as they now stand after `c54db93`, because this intent is the
first code written against the changed specification and reading the previous shape would be reading a
tree-sitter package that no longer exists. Section "Context economy" fixes the content and the
provenance; section 5 fixes `index.path` and `index.max_age_hours` and what their absence means; WP15
fixes the kind alongside the other four symbol fields, where the index lives, and the done-when.

`internal/model/model.go` is read for how an artifact's types are declared and for `ContextFile`, which
is the nearest existing shape to a path with a hash. `internal/runner/runner.go` is read for
`informationBase`, because it is the one place that already walks the tree against a configured budget
and degrades silently when the configuration is absent — the same pattern this needs.

A42 is read for the property that must hold and for how it is held: a check in CI over the import
graph rather than a test. A62 is read as the precedent for a hash defined by the package that writes
it, which is why the record is a separate piece rather than a field added here.

`scripts/` is read for what this project's own tooling looks like, so the producer matches it rather
than inventing a convention.

Nothing outside the tree is needed. The format is Xeno's to fix, so there is no host to ask and no
reference to confirm, which makes this the first intent in a while whose information base is entirely
inside the repository.
