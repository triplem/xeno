---
intent: github.com/triplem/xeno#184
phase: 02-design
created: "2026-10-03T12:55:25Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9867bfc13262bbaa46161c01b04a6ff2276ff8c11b9d8cf699d7a4a68d013ee7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`New` reads the variable, because the value describes a session and a session has one runner per
process; reading it per command would let two phases of one session disagree about what wrote them,
and would put a runner fact in the CLI. `HarnessVersionEnv` is an exported constant, since the name
appears in the reader, the flag's help and the tests, and a literal in three places is how one comes
to be misspelled. `parse` assigns the field only when the flag is non-empty — it assigned
unconditionally, which would have wiped what `New` read — so the precedence is flag, variable,
absent, in one `if`. The two writers do not change at all: they already read `Runner.ToolVersion`,
which is why this is four lines and is the payoff for having put the value on the runner in #181.
The runner reads this one variable and no client specific one, because section 7 says it does not
know which harness it runs under and a fallback chain ending in a client's variable is harness
detection with extra steps. `newFixture` clears the variable, which makes every existing test
hermetic for free, and a `reopen` helper exists so the one test about the variable does not duplicate
the fixture's construction. Six alternatives refused, including the variable winning over the flag —
the narrower statement is the better evidence, and A80 had already said which way round it goes.
