---
intent: github.com/triplem/xeno#195
phase: 03-implementation
created: "2026-10-03T14:21:36Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+b54626e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a131b0e3e191e5be4643cde69fbfd92dbcc2347506208cb6eb730387e31ba1f9
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`model.Learning` carries `Common` inline with both body fields `omitempty`, and `LearningEntry` the
four keys in `LearningKeys`' order. `learningPath` is the one place the phase's absence branches;
`RecordLearning` checks arguments first, reads an existing record tolerating `os.IsNotExist`, refuses
the two contradictions with the path and the count in the message, sets the header from `common()` so
it describes this write, and appends or sets `no_finding`; `checkLearningArgs` switches over the four
states of `--no-finding` against an empty entry, then the category against the closed set and the
three text keys against `TrimSpace`. `cmd/xeno` gains two usage lines, five flags, a table entry with
`needsPhase` false and the reason, and a command that resolves the phase where given and reports the
count with the last category. The learning skill's YAML block is replaced by the invocation and "do
not write the file by hand" with the version comparison as the reason — a skill showing both would
let the reader choose the one that was wrong — and the intake skill names the command. Six runner
tests and one at the surface, with a refusal table of seven that also checks no file was left behind.
Measured on a throwaway tree first: the command's record carries the stamped version where a typed
one carries `0.1.0-dev`, both forms work, all seven refusals fire, G-Learning passes. Two deviations:
`ASSUMPTIONS.md`'s "Built" sentence amended rather than appended to, because a list that names the
last of something and then grows stops being trusted, and `checkLearningArgs` made a method for no
reason beyond matching its neighbours.
