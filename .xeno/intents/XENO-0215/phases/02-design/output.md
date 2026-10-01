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
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: by-hand
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The format is YAML, one file, with the provenance at the top.**

```yaml
tool: go-symbols          # what produced it
tool_version: 0.1.0
produced_at: "2026-10-01T09:00:00Z"
symbols:
  - name: Compare
    kind: func
    file: internal/enforcement/enforcement.go
    line: 212
    container: ""
  - name: Requirements
    kind: method
    file: internal/host/gitlab/gitlab.go
    line: 96
    container: Adapter
```

YAML because every other file a project writes for Xeno is YAML and the one vendored dependency already
reads it. A line oriented format would be cheaper to produce and would be the first file in this project
a reader needs a separate explanation for.

Provenance first because it is what decides whether the symbols are read at all, and a reader that has
to parse a hundred thousand symbols to discover the index is stale has done the work it was trying to
avoid.

**`kind` and `container` are free strings, not enumerations.** This is the decision most likely to be
questioned, so the reason is here rather than in a comment. Xeno cannot enumerate the kinds of every
language a project may index: `func`, `method`, `class`, `interface`, `trait`, `object`, `record`,
`struct`, `impl`, `module`. A closed set would be Xeno's idea of what a symbol is, and section 5 says
every tool has its own and that is the project's decision. The agent reads these strings and the agent
can read anything; no gate reads them, so nothing has to match.

`container` is empty where a symbol is at the top level, which is a value and not a missing field.

**The reader cannot fail, and its signature says so.**

```go
func Load(path string, maxAge time.Duration, now time.Time) (*Index, string)
```

A nil index with a reason, never an error. Absent, unreadable, malformed and stale are four reasons and
one outcome, which is what the documents require and what criterion 4 asks for. The reason exists for
the record the next piece writes, and until then for a person asking why there was no index.

`now` is a parameter because staleness is compared against it and a function that read the clock could
not be tested across the boundary.

**The lookup is exact, on name, and returns every match.**

```go
func (i *Index) Lookup(name string) []Symbol
```

No prefix matching, no fuzzy matching, no case folding. "Where is X" is the question section 5 names,
and X is a name the agent already has, from an error message or a call site. Anything looser is a search
over the index, which is the thing the index replaces.

Nil receiver returns nil, so a caller holding no index needs no branch.

**The producer is `scripts/go-symbols.go`,** run with `go run`, writing to `.xeno/local/index/`. It
walks the repository with `go/parser` in a mode that skips function bodies, since nothing in the format
needs them.

<!-- xeno:section:alternatives -->
## Alternatives

**A closed set of kinds.** Rejected, and it is the closest call. An enumeration would let Xeno validate
an index, give a reader a list to write against, and let a future gate — if one ever read the index,
which none may — reason about kinds. It would also be Xeno deciding what a symbol is for Java, Kotlin,
C# and TypeScript from a Go repository, which is exactly the judgement #149 moved to the project. The
set would be wrong at the first language nobody considered, and a project whose tool emits `trait` would
have to choose between lying and being rejected.

What makes the open set safe is that nothing compares a kind. No gate reads it, the lookup matches on
name, and the consumer is an agent that reads prose for a living.

**A line oriented format**, tab separated, one symbol per line, with the provenance in a header comment.
Rejected. It is cheaper to produce and much cheaper to read for an index of any size, and ctags already
has a format of this shape that projects could emit directly. It would be the first file in this project
that is not YAML, so a reader needs a second explanation and the one vendored parser does not apply. If
index size ever becomes the problem the format solves, this is the alternative to revisit and the
provenance block is already a header.

**A `format_version` field.** Rejected, narrowly. Every artifact carries `schema_version` and this is not
an artifact: it is not hashed, not committed, not in the trail, and may be wrong without consequence. A
version on a file nothing validates is a field that will be read by nothing and set to `1` forever. The
cost of being wrong is that a second format needs a way to be told apart, and a `tool` field that
already names its producer is most of that.

**Returning an error for a malformed index and a nil for an absent one.** Rejected by criterion 4 and by
P1's learning. The distinction is real — one is a project that has not written an index and the other is
a project whose tool emitted something broken — but both must leave the phase running, so both are a
reason rather than an error. Making one an error would put the choice of whether to continue in the
caller, and the documents have already made it.

**Validating that files and lines exist.** Rejected and in P1's non-goals. It is a second walk of the
tree on every phase start, to check something the documents permit to be wrong, and the cost of a wrong
answer is one wasted read by the agent.

**Reading the clock inside `Load`.** Rejected. Staleness is the one thing in this piece with a boundary
that has to be tested on both sides.

<!-- xeno:section:impact -->
## Impact

**A new package that no existing code calls.** `internal/index` is read by nothing when this intent
lands: the record is the next piece and the MCP operation is behind WP11. That is deliberate and it is
the shape #149 implies, but it means the package's first real caller will be the thing that tests
whether the format was right. A reader reasonably asks why a package exists with no caller, and the
answer is that a format a project writes against has to exist before anything can read one.

**`project.yaml` gains two fields that the runner reads.** `index.path` and `index.max_age_hours` are in
section 5 already, so this is the runner catching up with the specification rather than a new
configuration surface. Neither is required and an absent `index` block means no index, which is every
repository's state today including this one.

**The scaffold does not change.** `xeno init` writes no `index` block, because the honest default is no
index and a commented out block in a generated file is a suggestion a project has not made. A project
that produces an index adds the two fields, which is the step #149 calls the honest cost.

**A42's check gains a second property to hold.** The CI step asserts `internal/gates` reaches no network;
it now also asserts it reaches no index. One more line in the same step, and the reason is the same
reason: a verdict may not depend on something allowed to be absent.

**`scripts/` gains a Go file**, which is a first: the two existing scripts are shell. `go run` on a file
under `scripts/` is not part of the module's build, so the release and the five binaries are untouched.

**Nothing in the trail, the hashes or the gates.** No artifact field, no gate, no template, no
dependency, no change to any existing report or command's output. `xeno gate verify` recomputes the same
verdicts over the same artifacts.

**What the next piece inherits.** A `*Index` with provenance, and a `string` reason when there is none.
Both are what the `tools` entry needs, which is why the reason is a value rather than a log line — and
the decision that piece has to take, which of the two hashes "response hash" means, is untouched by
anything here.
