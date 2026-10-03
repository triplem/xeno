---
intent: github.com/triplem/xeno#183
phase: 02-design
created: "2026-10-03T17:31:49Z"
schema_version: "1.0"
runner_version: dev+9140756.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fb9fb1fca4ca5bb218239f24876d9fb541242b14b323efba3044d0ce400dc90f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`plugin.ExpectedDigest` is a package variable the release sets through ldflags, empty by default,
sitting beside the `Hash` that computes what it is compared with so a reader of either finds the
other. `supply` is three branches and no fourth: no anchor is `notImplemented`, no plugin is a
finding, a difference is a finding naming both digests, otherwise pass — and nothing reads a version,
a manifest or the lock. The no-anchor branch reuses `notImplemented` rather than inventing a result
value, since section 5 already covers a check a runner did not perform. The release computes the
digest with `go run scripts/plugin-digest.go`, the same code the gate compares with, because a digest
computed differently is a red verdict on every project that installs the release; the script fails
loudly rather than printing an empty digest, which would compile in an empty anchor and turn the gate
off for a whole release. `internal/plugin`'s tests become external, because the gate imports the
plugin package and those tests import the gate — and the alternative, duplicating the hash inside
`internal/gates`, is what the single-implementation argument forbids. Nine alternatives refused,
three of them by one sentence of section 13: an expected hash stored beside the thing it describes
proves only that both were written by the same hand. The lock's `plugin` block is the most plausible
wrong answer, because the previous intent had just added it.
