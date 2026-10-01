---
intent: github.com/triplem/xeno#151
phase: 03-implementation
created: "2026-10-01T14:00:58Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c54db93.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a75f160c4fa038e78e28937afba152a1b4211dc241f1c2221999e2470679afd1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: by-hand
---

# Implementation

<!-- xeno:section:changes -->
## Changes

**`internal/index`, new.** `Symbol` with the five fields, `Index` with the three provenance fields and
the symbols, `Load`, `Lookup` and `Age`. `Kind` and `Container` are free strings with the reason at the
type, since that is where somebody will want to add an enumeration.

`Load(path, maxAge, now) (*Index, string)` returns a nil index and a reason, never an error. Five
reasons: no path configured, configured and absent, unparseable, no provenance to judge an age by, and
stale. `maxAge` of zero falls back to `DefaultMaxAge`, 24 hours, which is section 5's default.

`Lookup` matches exactly on name and returns every location. A nil receiver returns nothing, so a
caller without an index needs no branch, and so does a name the index does not hold.

**`model.Project` gains the `index` block**, `path` and `max_age_hours`, which section 5 already
specifies and nothing read.

**`runner.symbolIndex(now)`** reads it, resolves a relative path against the repository root, and
returns what `Load` returned. The reason is a value rather than a log line because it is what the
`tools` entry will record, which is the next piece.

**`scripts/go-symbols.go`,** this repository's own producer, `//go:build ignore` so the module never
builds it and the release never sees it. It parses with `parser.SkipObjectResolution` and skips no
bodies explicitly — the AST walk only looks at declarations, so bodies cost parsing and nothing else. It
indexes functions, methods with their receiver as container, types, struct fields and interface methods
with the type as container, and consts and vars. A file that does not parse is skipped with a line on
stderr rather than failing the run, because an index is allowed to be incomplete and refusing to index a
repository over one generated file would be worse.

It produces **997 symbols** for this repository, and the line numbers were checked by hand against four
definitions.

**`internal/index/testdata/example-symbols.yaml`,** the worked example, annotated, and read by a test so
it cannot drift from what the reader accepts.

**`.github/workflows/xeno.yml`** gains a second import check beside A42's, asserting
`internal/gates` reaches no `internal/index`. Same step, same reason: a verdict may not depend on
something the documents permit to be missing, stale or wrong.

**Tests.** Eight in `internal/index`: the four no-error causes as a table, the staleness boundary from
both sides, the default, the provenance, lookup over a name with several definitions, the two absent
cases, the worked example, and a round trip through the real producer. Two in `internal/runner` for the
config and the relative path. Coverage 96.3%.

<!-- xeno:section:deviations -->
## Deviations from the design

**Criterion 2 is half met, and the missing half is not mine to supply.** It asks that the format be
"documented where a project writing one will look". `docs/` holds the four documents, two of them
normative, and standing rule 1 puts them out of an agent's reach; a published description belongs to
WP16, whose documentation site is not built. So what exists is the worked example in `testdata/`,
annotated for somebody who does not read Go and machine checked by a test, plus the package comment.

That is the shape documented and kept honest, and it is not published. Where a project would actually
look, there is still nothing. This belongs on WP16's list rather than being quietly counted as done,
and it is the one criterion in this intent I would not claim in full.

**The producer indexes more kinds than the acceptance asked for.** Criterion 6 asks only for a producer.
This one emits struct fields and interface methods with their type as container, and consts and vars,
beyond the functions and types that would have satisfied it. The reason is in the format: `container`
exists so that a common name is findable, and a producer that emitted no contained symbols would have
left that field untested against real output. It is more than was asked and it is what made the
`Requirements` case — one name, a field and three methods across four files — a real demonstration
rather than a fixture.

**The design said bodies are skipped and the code does not skip them.** P2 said the producer parses "in
a mode that skips function bodies". `parser.SkipObjectResolution` does not do that, and the mode that
would, parsing declarations only, is not what `go/parser` offers for a file walk of this shape. The AST
walk looks only at `f.Decls`, so bodies are parsed and then ignored: the saving P2 claimed is not taken,
and for a repository of this size the producer runs in well under a second anyway. Recorded because the
design sentence is wrong rather than imprecise.

**A `*Index` with no caller in the product.** `runner.symbolIndex` exists and nothing calls it, because
the record is the next piece and the MCP operation is behind WP11. P2's impact section says this is
expected; it is restated here because an unused method is exactly what a reviewer should ask about, and
the answer is that a format has to exist before anything can read one.

**Nothing else.** No new dependency, no artifact field, no gate, no template, no spec change, and no
change to any existing command's output. `gate verify` recomputes the same verdicts.
