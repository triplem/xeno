---
intent: github.com/triplem/xeno#212
phase: 02-design
created: "2026-10-05T15:49:26Z"
schema_version: "1.0"
runner_version: dev+09e2aa6
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 517ef183b4bf2cc0e1f13281732d88611d05a848f1619d7002e6a244bae7acf1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`declaredResults` is the shared body, `build` and `testReport` its two callers, and `TestKind`
sits beside `BuildKind` for the reason that one's comment records. The pending path is inherited
from G-Build rather than chosen, and it is why the trail survives.

The cost is that `not-implemented` becomes `pass` over half a clause, which cannot be mitigated
in the verdict; the audit row is the mitigation, and its count moving to 36 is what a later pass
reads.
