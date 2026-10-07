---
intent: github.com/triplem/xeno#153
phase: 03-implementation
created: "2026-10-07T14:26:52Z"
schema_version: "1.0"
runner_version: dev+6b48c17.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3da3c8afa52f3ea66b7cbcf06697f0e978f5ccd7c0e72186318c47575b36dc47
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`docs/symbol-index.md` carries `internal/index/testdata/example-symbols.yaml` verbatim in a
fenced block, with three paragraphs above it: what the index is for and where the format is
fixed, that nothing depends on the index being there, and that a test holds the block against
the file. No sentence names a field, a requiredness, a default or a value, which was the test
applied to each one.

`TestThePublishedPageCarriesTheWorkedExample` in `internal/index/index_test.go` compares the
page's first `yaml` block against the file on disk, beside the test that reads the same file
through `Load`, so the file is the authority for both.

One entry in `docs/README.md` under "Running it", beside `commands.md`.

No deviation. One thing worth naming that departs from nothing: the block is tagged `yaml` and
the test's pattern requires the tag, which is a stricter anchor than the commands test uses and
gives the right failure message if somebody drops it for a renderer's sake.

Build, suite, `gofmt`, `vet` pass, `gate verify` matches 516 verdicts, and all fourteen
relative links resolve.
