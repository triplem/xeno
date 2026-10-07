---
intent: github.com/triplem/xeno#228
phase: 04-verification
created: "2026-10-07T06:40:11Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2086af031f4c57a41fa056ebc850343c8a7ef9a71e7402709d1c38891639c7ec
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Eleven criteria, all passing: two held by a test, two read off the binary's own output, four by
a grep run with a positive control, three by the project's check suite.

`go test ./...` passes across 20 packages. `go build`, `gofmt` and `go vet` are clean.
`xeno gate verify` recomputes and matches 481 verdicts, so nothing already sealed moved.

Both sentences were read off a run rather than from the strings: this intent's own P3 finish
printed the fresh-session suggestion, and a decided P5 on a throwaway copy printed the
end-of-intent one. Both tests were checked the other way round — each string reverted, both
tests failed naming the phrase that had gone, then restored.

No harness is named, checked with a grep and a positive control in the same file, and both
greps `xeno.yml` runs were run locally and are clean.

The gaps are the honest half. Nothing checks that anybody follows the hint, and nothing can
without reading the cost ledger, which is local gitignored data outside `artifacts_hash`; that
figure is WP20's and A98 says so. The absence of a harness command is checked over `internal/`,
`cmd/` and `.xeno/plugin/` and not over `docs/`. No test asserts that the suggestions for an
open phase stayed silent, so a later addition there would pass. And the end-of-intent sentence
was read on a copy, because the run that prints it rewrites a `gate.yaml`.
