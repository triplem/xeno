---
intent: github.com/triplem/xeno#160
phase: 01-requirements
created: "2026-10-01T15:32:11Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4c277a7a93cc20ccaa8840b8de2c1db8510caf6ba7dd8e57874041ad9b497875
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The criteria are artifact, rule tree and verdict over the pair. Green: a repository with no rule tree,
at every phase. Red: a review rule applying to P5 with no entry, which is the criterion the gate rests
on, since without it an empty checklist passes; an entry with no result; a result outside `met`,
`deviation` and `not-applicable`; a `deviation` or `not-applicable` with no note, where `met` needs
none; an entry naming a rule not in the effective set; and a checked rule applying to the phase whose
predicate type has no implementation, which is the only thing this piece says about a checked rule and
it says it rather than passing. A lens entry neither completes nor breaks the count — four of them and
no rule entry is as red as an empty checklist — and still owes a result. The gate applies from P0 and is
quiet where nothing applies, reported as a pass and not as `not-implemented`, and it calls the resolver
#158 wrote rather than resolving again, because two resolutions that could disagree would make
`rules_hash` meaningless. Non-goals: no predicate evaluated, no derived section, no command, no
judgement of an answer, no rule shipped, nothing sealed touched.
