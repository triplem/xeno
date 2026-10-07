---
intent: github.com/triplem/xeno#153
phase: 04-verification
created: "2026-10-07T14:29:47Z"
schema_version: "1.0"
runner_version: dev+6b48c17.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 48f3d510201dec69493b140981a9c6c2c8d553b882e9ccf3cde3145d828a0cbd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Ten criteria, all passing: six checks and four read in the files.

The test holds the pair, checked three ways: one line of the page's block changed, which fails
naming both files; the fence's `yaml` tag removed, which fails with "carries no yaml block"
rather than with a mismatch; and restored, which passes.

The prose was checked claim by claim against the line that carries it — section 5, Appendix A,
the example's comments and `internal/index/index.go` — rather than asserted. Where the format
is fixed was checked rather than remembered: line 659 is inside section 5 and lines 2293–2294
inside Appendix A, so the page sends a reader to the right place for each. No field name
appears as a backticked token in the prose; a looser grep returns one line where the match is
the English word "file", recorded because that is the grep somebody would run first.

The example is untouched and the three files touched are the page, the index and the test.
Suite, build, `gofmt`, `vet` pass, `gate verify` matches 517 verdicts, and all fourteen links
resolve.

The gaps are the honest half. The page's three paragraphs are held by nothing and go stale
silently if section 5 or the reader changes. The format is published and still not an interface
description, which is the half #153 recorded as unmet and the maintainer chose to leave. Nothing
notices a fourth page missing from the index. Two guards now stand where WP16 wants a
generator, and a later reader could take the pattern as the design. The test couples a
package's assertion to a document's fence tag. And nobody searching the web arrives anywhere,
because WP16 has published nothing.
