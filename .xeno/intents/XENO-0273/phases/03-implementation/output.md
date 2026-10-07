---
intent: github.com/triplem/xeno#153
phase: 03-implementation
created: "2026-10-07T14:26:36Z"
schema_version: "1.0"
runner_version: dev+6b48c17.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3da3c8afa52f3ea66b7cbcf06697f0e978f5ccd7c0e72186318c47575b36dc47
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`docs/symbol-index.md`, new. A title, three paragraphs, and
`internal/index/testdata/example-symbols.yaml` verbatim in a fenced block tagged `yaml`.

The first paragraph says what the index is for, that a project produces it and Xeno ships no
indexer with the four tools the example names, that what Xeno fixes is the shape of the answer,
and where the format is fixed — section 5 under "A symbol index, not a graph", with
`index.path` and `index.max_age_hours` in Appendix A's `project.yaml` rather than in section 5.
The second says nothing depends on the index being there: absent, unreadable, malformed and
stale are four causes with one outcome, no gate reads it, and a verdict never depends on it.
The third says what the block is and that a test holds it against the file.

No sentence names a field, a requiredness, a default or a value. That was the test applied to
each one: whether removing it would leave a reader unable to use the block.

`internal/index/index_test.go`, one test.
`TestThePublishedPageCarriesTheWorkedExample` reads `../../docs/symbol-index.md`, takes its
first ` ```yaml ` block and compares it against `testdata/example-symbols.yaml` with the
trailing newline trimmed. It sits beside `TestTheWorkedExampleIsWhatTheReaderAccepts`, which
reads the same file through `Load`, so the file is the authority for both and the dependence
runs one way. The import block gains `regexp`.

`docs/README.md`, one entry under "Running it", beside `commands.md`, saying what the page is
and that Xeno ships no indexer and a phase runs without one.

Nothing else. `example-symbols.yaml` is untouched, no normative document changes, and no Go
file outside the test.

`go build`, `go test ./...`, `gofmt -l .` outside `vendor/` and `go vet ./...` all pass.
`xeno gate verify` matches 516 verdicts. All fourteen relative links in the two touched pages
resolve against the tree.

<!-- xeno:section:deviations -->
## Deviations from the design

None. P2's decisions were all carried out as written: the page carries the file, the prose sits
above the block and says only what the block cannot, the test compares the first fenced block
against the file on disk rather than an embedded copy, it lives in `internal/index/index_test.go`
beside the reader's own test, the entry is under "Running it", the page is named
`symbol-index.md`, and nothing in the example points back at it.

One thing is worth naming although it departs from nothing. The fenced block is tagged `yaml`
and P2 did not say so; `docs/commands.md`'s block is untagged, because what it carries is a
usage text and not a language. The test's pattern matches ` ```yaml ` rather than ` ``` `, which
is a stricter anchor than the commands test uses and ties the page's rendering to the test. If
somebody drops the tag for the sake of a renderer, the test fails with "carries no yaml block"
rather than with a mismatch, which is the right message for what went wrong.
