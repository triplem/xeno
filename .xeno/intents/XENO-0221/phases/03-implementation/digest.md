---
intent: github.com/triplem/xeno#167
phase: 03-implementation
created: "2026-10-01T16:56:36Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15693cf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3791e2004efb4cf75f1d8afcbf47632a72b6683790965ecd705623e7e853a54b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`internal/external` is 283 lines plus 330 of test: a check per declared gate, the declaration validated
so a configuration mistake reads as one, the file hashed from disk on every call because "checked before
every run" is a frequency, and a mismatch refusing to run and naming both hashes. The contract is A76 —
five fields in, findings with a required cause out, unknown fields ignored so a tool may answer a
superset. The exit status decides, a non-zero exit with nothing to say gets a synthesised finding
because `Status` refuses a fail without one, and output is clipped before it reaches a verdict.
`gates.Ctx` gains an injected `External` producer, called after the table and routed through the same
`carryForward`, which is where an external finding gets its id and is refused a decision; the gates
package still starts no process. The runner reads the declaration per run and passes nil where there is
none. The deviation worth reading: the timeout looked implemented and bounded nothing, because killing
the process leaves a grandchild holding the pipe and `Output` waits for it — the test took thirty
seconds while asserting a two hundred millisecond limit, and only the duration showed it. `WaitDelay`
fixes it and A75 records the whole of it. Also recorded: this is the first WP4 piece with no
demonstration against the real tree, because declaring a gate here would mean writing a tool for this
project to run on itself.
