---
intent: github.com/triplem/xeno#156
phase: 01-requirements
created: "2026-10-01T14:37:15Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 28f6e302105ba421aed26f73ab65468cb7db24f0e5d5985a25cd3d5e58e88632
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The acceptance is that each corrected sentence is checkable against the thing it describes: the "Not
built" paragraph names only what is not built and says what G-Supply and G-Secret wait on, "Built"
accounts for what postdates its list with the assumption that governs each, the README's command block
equals the dispatch table in `cmd/xeno/main.go`, the trail section states the two shapes of the trail
with `XENO-0107` as the boundary and carries no count, and the coverage table has a row per package
with tests in the tree, every named test existing. Nothing regresses: build, tests, `gofmt`, `go vet`
and `gate verify` all pass. Out of scope by rule rather than by choice: no document under `docs/`, no
owner named for the two files, and no check built that would have caught the drift — the first is the
standing rule, the second is a plan change and a person's commit, the third belongs beside WP16's
generated reference and is left in the residual risk. The figures go to #117 as a comment, counted
from the tree, with the `gate run` timing taken on a throwaway copy so no sealed phase is rewritten.
