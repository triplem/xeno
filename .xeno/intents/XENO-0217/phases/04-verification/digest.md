---
intent: github.com/triplem/xeno#158
phase: 04-verification
created: "2026-10-01T15:15:50Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2ea9174303a58847db930b3d5825924e254a41f7f6b8cced3a4c6ccdf4dcad33
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
203 test cases pass across the three packages touched, 17 packages report `ok` or no test files, nothing
fails, `gofmt` and `go vet` are clean, and `./xeno gate verify` is at exit 0 over 175 verdicts with no
sealed `artifacts_hash` moved. G-Rules reports a verdict in this repository — pass, over a tree with no
rules, which is the case the criteria put first — and its row no longer reads `not-implemented`, so a
phase here reports ten gates running rather than nine. The collision was exercised against the real
binary on a throwaway copy: two binding rules with one id turned a phase red with a finding naming both
files, removing one returned it to green, and the copy was discarded. The hash a phase recorded on that
copy, `c721ad48…`, was reproduced by piping A66's rendering through `sha256sum`, which is what byte
exactness actually claims. The division of the field's two meanings falls inside this intent: its first
three phases carry `by-hand`, its last two carry the hash of the empty rendering. Gaps: nothing checks
`rules_hash` after it is written, the placeholder is still accepted, a `check` is unvalidated until the
predicates land, `applies_to` is unchecked against the phases, and no rule tree here has been written by
anyone but the test suite.
