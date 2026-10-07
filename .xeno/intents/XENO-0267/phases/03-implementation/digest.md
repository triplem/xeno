---
intent: github.com/triplem/xeno#228
phase: 03-implementation
created: "2026-10-07T06:34:37Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3f920b60852ab2c93fbe492852418050bc810966579e4e750f52e466938fd3f6
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two suggestions in `internal/runner/next.go`. The `not-started` one says the phase is started
in a fresh session so that its context is what `context.lock.yaml` says it was given, and keeps
its command. The last phase's one says nothing of the intent's context is read again, so the
next begins best in a fresh session, and keeps having no command. Each carries a comment naming
the clauses it rests on and why no harness is named.

Two assertions in `internal/runner/runner_test.go`, beside the ones already on those two
suggestions. They hold the reason as well as the phrase, and the first also asserts that
`/clear` has not appeared.

One row in `docs/assumptions.md`, A98, carrying the two decisions beyond the sentence with the
clause behind each, and the reason a gate was refused: the only signal is the cost ledger,
which is gitignored local data outside `artifacts_hash` by design.

One deviation from P2, recorded: the row is `open` and not `approved`, because the state column
reports what a person has said and nobody has said anything about either half yet. The cost
half carries `accepted until WP20` inside it.

Build, tests, `gofmt` and `go vet` pass.
