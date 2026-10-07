---
intent: github.com/triplem/xeno#277
phase: 04-verification
created: "2026-10-07T13:36:50Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 87103ee9eda9d4f77bc202e4fb4e10aeb0a21d7c7ceadca8c4c916846cb207bb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Thirteen criteria, all passing: nine checks and four read in the files.

`honest` no longer reads the `tool` field, shown with a control — the grep returns the line on
`main` and nothing here. No gate reads any of the triple to decide anything: the one remaining
mention across `internal/gates` is `sessionFields`, which asserts presence and reads no value.
That claim is an absence and was established by the grep with its control.

What the term did was measured before it went. Four artifacts in 499 phases declare
`tool: manual`; `WriterlessHash` already covered `secrets_hash` and `rules_hash`, and
`goneBundle` their `strings_hash`, both carrying `template: intake@0.1.0` against a shipped
`1.0.0`. That left `context_hash`, which is the four findings. The replacement keyed on a
missing lock was measured and found empty: every phase in the trail carries one.

The four releases read as four decisions of `type: approved`, each naming the person, with no
obligation anywhere. `gate verify` exits zero again, having exited 1 between the change and the
releases.

Both test tables were checked the other way round. Nothing in the four artifacts is edited.
The suite, build, `gofmt`, `vet` and `gate verify` all pass.

The gaps are the honest half. Two sealed phases read `approved` permanently, for a change made
after the fact. The gate is stricter and twice in 499 is the whole sample. The four approvals
are a person's statement typed by the agent, which is further from what `next.go` protects than
the specification passages were. No test asserts that the gate reads no declared field, so the
term could come back. `goneBundle` is still coarser than it looks. And this phase was started
over, so its lock postdates the implementation it verifies.
