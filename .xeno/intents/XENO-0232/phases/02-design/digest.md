---
intent: github.com/triplem/xeno#195
phase: 02-design
created: "2026-10-03T14:21:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+b54626e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 60caf93b67d10a66489e68a52008b3a84be027b4b92fbee9bc8e957f2f663bbb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`model.Learning` carries the common header inline with `NoFinding` and `Learnings` both `omitempty`,
so exactly one is written and a record carrying both is not representable — which is the shape
G-Learning reads and the reason the gate needs no change. `RecordLearning` takes the phase as a
string where empty means the intent level record, `learningPath` being the one place that branches.
The header is rewritten on every call rather than carried over, because it describes the binary doing
this write and reading an existing record must not smuggle an old version forward. Arguments are
checked before the record is opened, so a refused call leaves no file and a bad category never
reaches disk; the gate checks the same set afterwards, two readers of one closed set in `model`
rather than two definitions. Contradictions are refused in both directions and never resolved, since
the statement already on disk is somebody's. Entries append, so the order is the order recorded. No
sealed-phase guard, matching `SectionSet` for the reason that holds: the write makes `verify` report
a divergence and `phase finish` judges the phase again, so refusing would refuse the first half of a
re-judgement. Seven alternatives refused, four of them because they would have left a field with the
wrong author — `phase finish` stamping the header, a `--file` taking the whole record, the gate
keeping sole ownership of the category check, and writing the record automatically, which is the
runner authoring an observation it cannot have made.
