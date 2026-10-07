---
intent: github.com/triplem/xeno#153
phase: 00-intake
created: "2026-10-07T14:22:25Z"
schema_version: "1.0"
runner_version: dev+c1f6405.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8a48246c302d57d6df4178b48dd5f9694606b523d9b06ac9ddc21364470b2eaf
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
#151 documented the symbol index format in `internal/index/testdata/example-symbols.yaml`,
annotated for a reader who does not write Go and read by `index_test` so it cannot drift from
what the reader accepts. It is published nowhere, so a project that wants to produce an index
discovers the format by cloning this repository and opening a file under `testdata`.

Read rather than assumed: every item on #153's list of what a page needs to say is already in
that file's comments — the three required provenance fields with their reason, the five symbol
fields and no more, `container` omitted where nothing encloses, `kind` free and compared
against nothing, `index.path` and the gitignored convention, `index.max_age_hours` with its
default of 24 and stale treated as absent, and that Xeno ships no indexer. What is missing is
not the description but where it sits.

So the obvious page is the wrong page. #153 says a second copy of a format description is wrong
within a release and that this one has a test holding the first copy honest, and #231 landed
the pattern earlier today: `docs/commands.md` carries the `usage` constant with a test that
fails if the two part.

The maintainer was put three options and chose publishing the example held by a test. In scope:
`docs/symbol-index.md` carrying the file verbatim, a test beside the one that reads it, and one
entry in `docs/README.md`. Out of scope: a prose reference page, editing the example, the
documentation site, a link checker, and moving the file out of `testdata`.
