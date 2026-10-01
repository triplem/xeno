---
intent: github.com/triplem/xeno#151
phase: 04-verification
created: "2026-10-01T14:01:29Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c54db93.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cf2e58665419a3d8807a64acbecdee845b5a719da7308383c08335de623fcdc4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| criterion | how it is verified |
|---|---|
| 1, the format is the eight fields | the types; `TestTheWorkedExampleIsWhatTheReaderAccepts` asserts all three provenance fields and a symbol with no container are present in the documented shape |
| 2, documented where a project will look | **half met.** The worked example is annotated and read by a test. Nothing is published, and that is WP16's. Recorded in `deviations` |
| 3, reading from `project.yaml` | `TestTheIndexPathIsRelativeToTheRepository`, with `max_age_hours: 48` and an index 30 hours old, so the configured value is read rather than defaulted |
| 4, four causes, no error | `TestNothingAboutReadingAnIndexIsAnError` as a table over the four, each asserting a nil index and a reason that names its cause; `TestAStaleIndexIsTreatedAsAbsent` is the fifth |
| 5, lookup returns every location | `TestLookupReturnsEveryLocationForAName` over a name with a method and a field in different containers; `TestAnAbsentNameAndAnAbsentIndexAreBothAnswers` for the two nils |
| 6, this repository produces one | `TestTheFormatRoundTripsThroughThisProjectsProducer` runs `scripts/go-symbols.go` over this tree, reads it back, and asserts the index can find `Lookup` on `Index` |
| 7, the gate path reaches no index | `go list -deps ./internal/gates` names no `internal/index`, and the check is a step in `.github/workflows/xeno.yml` beside A42's, not a test |
| 8, no dependency and the suite | `git diff` on `go.mod` and `go.sum` is empty; below |

The round trip in criterion 6 is the test worth keeping. Everything else reads a fixture that somebody
in this intent wrote, which means it asserts that the reader accepts what the author believed. That one
runs a real tool over a real tree and asserts the reader accepts **its** output, and it is also the only
test that would catch the producer and the reader disagreeing about the format after a change to either.

The staleness boundary is tested from both sides on purpose, which is the reason `Load` takes `now`
rather than reading the clock: an index an hour old is accepted inside a two hour window and rejected
inside a thirty minute one, from the same fixture.

<!-- xeno:section:results -->
## Results

Seven criteria hold in full and one in half.

| check | result |
|---|---|
| `go build -o xeno ./cmd/xeno` | builds |
| `go vet ./...` | clean |
| `go test ./...` | every package ok |
| `gofmt -l .` outside `vendor/` | prints nothing |
| `xeno gate verify` | 168 verdicts, exit 0 |
| `go.mod`, `go.sum` | unchanged, no second dependency |
| `go list -deps ./internal/gates` | no `net`, no `net/http`, no `internal/index` |
| `internal/index` coverage | 96.3% |

**The producer's output, read back and queried.** 997 symbols from this repository. Four line numbers
checked by hand against the files they name, each landing on the definition rather than near it:

```
Compare        func       internal/enforcement/enforcement.go:144  container=""
Requirements   field      internal/enforcement/enforcement.go:49   container="Report"
Requirements   method     internal/host/github/github.go:70        container="Adapter"
Requirements   method     internal/host/gitlab/gitlab.go:93        container="Adapter"
Requirements   method     internal/host/host.go:43                 container="BranchRules"
BranchRules    interface  internal/host/host.go:42                 container=""
```

That is the component doing the thing it exists for. `Requirements` has five definitions across four
files — a struct field, three methods on two different types, and an interface method — and a repository
search for the string returns every mention of it instead. The `container` field is what makes the five
distinguishable, which is why the producer emits contained symbols although the acceptance did not ask.

**Criterion 2, stated precisely.** The format's shape is documented in
`internal/index/testdata/example-symbols.yaml`, annotated for a reader who does not read Go, and a test
asserts the reader accepts it, so description and implementation cannot drift. It is not published
anywhere a project would look before cloning this repository. The gap is WP16's and is in `deviations`
and in `gaps`.

**The index is gitignored**, under `.xeno/local/`, so nothing this intent produces enters the trail.
`gate verify` recomputes the same verdicts over the same artifacts as before the intent.

<!-- xeno:section:gaps -->
## Gaps

**Nothing consumes the index, so the format is unproven where it matters.** Every test here reads it
back through the same package that defines it. The round trip adds a real producer, which is better, but
the consumer is still a test: no phase reads an index, no agent has queried one, and the sixth MCP
operation is behind WP11. The format will be judged the first time somebody with a real question uses it
to answer that question, and nothing before that point can substitute. If five fields turn out to be
four or six, this is where it shows.

**The format has never been written by a tool Xeno did not see.** `scripts/go-symbols.go` was written by
the same intent as the reader, so the two agree because one author held both. The claim #149 rests on is
that a project's own toolchain can write this, and a ctags or tree-sitter emitter would be the test of
that. It is cheap and it is not in this piece.

**Criterion 2's published half is missing**, named in `deviations`, belonging to WP16. Until then a
project discovers the format by cloning this repository and opening a `testdata` file, which is not a
documented interface.

**The provenance is trusted entirely.** `produced_at` is whatever the producer wrote, so a tool with a
wrong clock, or one that stamps the time it started a long run rather than finished it, makes staleness
meaningless in the direction that matters: an index reported fresher than it is. Nothing can check this
from inside Xeno — a file mtime would be a different claim and is also forgeable — and the documents
permit the index to be wrong, so this is a limit rather than a defect. It is worth stating because
staleness is the one judgement Xeno makes about the index and it is made on an unverifiable number.

**Lookup is linear over every symbol.** 997 symbols here and a brownfield Java repository will have
orders of magnitude more, so a map by name is the obvious next thing. It is not built because nothing
queries it yet and a map built at load time on an index nobody reads is work spent on a guess about the
caller.

**Not a gap.** That `runner.symbolIndex` has no caller. The record is the next piece and P2 said so in
advance.
