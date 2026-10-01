---
intent: github.com/triplem/xeno#151
phase: 01-requirements
created: "2026-10-01T13:53:05Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c54db93.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: dde464e9b230bdeee958110c72948291dcf333bfca30d997d35c8104f497ed04
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. `internal/index` holds the format as a type: per symbol, name, kind, file, line and enclosing
   container; per index, the tool that produced it, that tool's version, and the time it was produced.
   Nothing else. Call relationships and type resolution are absent and their absence is deliberate.

2. The format is documented where a project writing one will look, and a project can write a valid
   index from that description without reading Go.

3. Reading an index at `index.path` returns its symbols and its provenance. `index.path` and
   `index.max_age_hours` come from `project.yaml`, with `max_age_hours` defaulting to 24.

4. **Absent, unreadable, malformed and stale all return no index and no error**, each with a reason
   saying which it was. Stale means produced more than `index.max_age_hours` ago. No input makes the
   reader fail, because a phase must run without an index and the absent case is the state every
   repository is in today.

5. A lookup by symbol name returns every location for that name, each with its kind and container. A
   name that is not in the index returns nothing, which is not an error either: an index is allowed to
   be incomplete.

6. `scripts/` holds a producer for this repository's own Go source, built on `go/parser` and `go/ast`.
   A test produces an index from this repository, reads it back and queries it, so the format is
   exercised against a tool's output and not only against a fixture the author wrote.

7. **Nothing in `internal/gates` reaches `internal/index`.** Asserted the way A42's property is, by a
   check over the import graph rather than by a test, because a verdict may not depend on something the
   documents permit to be missing, stale or wrong.

8. `go.mod` gains no dependency. `go build`, `go vet ./...`, `go test ./...` clean; `gofmt -l .` outside
   `vendor/` silent; `xeno gate verify` exit 0.

<!-- xeno:section:non-goals -->
## Non goals

**The `tools` record is not written.** Two reasons and the second is the binding one. It is a separate
piece of work, and it needs a decision this intent may not take: `tools` carries a "response hash"
Appendix B does not define, and the phrase has two readings — the index Xeno read, or the response a
query returned. The first is reproducible from the index file; the second records what the agent
received and cannot be recomputed. A62 is the precedent for defining such a hash in the package that
writes it, so the definition belongs with the record and an implementation that picked one here would
make the choice invisible.

**No MCP operation.** The server is not built. The lookup is a function, and what makes it the sixth
tool is WP11's work.

**No indexer in Xeno.** #149's whole substance. The producer in `scripts/` is this repository's own
tooling, beside `github-settings.sh` and `module-set.sh`, and criterion 6 says so explicitly so that a
reader finding a Go indexer does not conclude #149 was reversed.

**No incremental rebuilding, no freshness management.** #149 replaced that paragraph: staleness is read
and reported, never managed. Xeno has no opinion about when a project should rebuild.

**No measurement.** #149 moved the saving to WP20. This intent claims nothing about tokens or bytes.

**No validation beyond the format.** An index may name a file that does not exist or a line past its
end. Checking that would be a second walk of the tree on every phase start, and the documents permit
the index to be wrong. A lookup returns what the index says; whether it is true is the agent's problem
and costs it one read.

<!-- xeno:section:constraints -->
## Constraints

**The documents fix the content exactly, and the second standing rule makes it a budget.** Five symbol
fields and three provenance fields. A sixth symbol field — a signature, a visibility, a docstring — is
a specification change first, however useful it would be, and the index exists to answer "where is X"
rather than to describe X.

**The degradation is the load bearing property.** Both documents say a verdict never depends on the
index, and WP15 calls that "the property everything else follows from". So the reader has no failure
mode. That is a constraint on its signature and not only on its behaviour: a function returning an
error invites a caller to treat an absent index as a problem.

**A42 binds the import graph, and this is the first package that could break it by being useful.** The
network packages are obviously out of the gate path. An index is a tempting thing for a gate to consult
and it is exactly what must not happen, so the property is checked the way A42's is.

**One dependency.** `go/parser` and `go/ast` are the standard library, so the producer adds nothing to
`go.mod`. Anything else would be a decision and this intent has no case for one.

**`.xeno/local/` is where the index lives**, uncommitted and under retention. Nothing this intent writes
enters the trail, and the producer's output is gitignored like the rest of the local data.

**Appendix B is not extended.** No hash is defined here. The one the record needs is the next piece's,
by A62's precedent.

**Eighty-eight columns, SPDX on every new file, no copyright line by A16.**
