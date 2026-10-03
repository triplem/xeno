---
intent: github.com/triplem/xeno#188
phase: 03-implementation
created: "2026-10-03T21:41:42Z"
schema_version: "1.0"
runner_version: dev+efc44a9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9d4ce1497faf8247d67a0f700420d69bdc21f56890c1b886db349d20c543a2a1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The two commands and the figure are built. `internal/gates` exports `QuestionShape` and
`DecisionShape`, unchanged, so that the refusals are the gate's own checks rather than
a second opinion about shape; G-Questions is not touched. `internal/runner/exchange.go`
holds `RecordQuestion`, which decodes the entry with unknown fields refused rather
than dropped, and `RecordDecision`, which refuses without a person, without a reason,
without an option, a withdrawal carrying an option, and a given id; around them
`exchange` over every phase, `nextExchangeID` continuing the intent's own numbering,
`artifact` refusing a phase with no `output.md`, and `amendFront` replacing one field
with the body written back byte for byte.  `IntentSummary` gains the two counts and
`AskedNothing`, `summarise` becomes `Summarise` so that both forms of the status
command read one definition, and `cmd/xeno` gains the two commands, `--chosen`,
`--proposed-by`, `--withdraw`, the row mark and the line after the table. Sixteen new
tests: the assigned ids, every refusal, the loop end to end with G-Questions passing
on a question that was recorded rather than hand written, a withdrawal resolving as
well, the body untouched, and the figure's three answers. Four deviations recorded,
the substantive one being a refusal the design did not name: `--resolves` has to
name a question the intent raised, because a typo is otherwise silent until P5 and
leaves something that looks like an answer in the trail.  The commands were used
on this intent as they were built: Q-2 is in this phase's frontmatter, put there by
`question record`, which refused its first draft for an invented field.
