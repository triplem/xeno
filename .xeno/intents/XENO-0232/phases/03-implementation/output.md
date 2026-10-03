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
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

**`internal/model/model.go`.** `Learning`, with `Common` inline and `NoFinding` and `Learnings` both
`omitempty` so exactly one is written, and `LearningEntry` with `Category`, `Observation`, `Proposal`
and `Target` in the order `LearningKeys` lists them. The comment on `Learning` says why both fields
are optional: a record carrying neither is what G-Learning calls empty, and one carrying both would
claim there was nothing to say beside something said.

**`internal/runner/runner.go`.** `learningPath(key, phase)`, the one place the phase's absence
branches, returning the intent directory's record where it is empty. `RecordLearning`, which calls
the argument check first, reads an existing record tolerating `os.IsNotExist`, refuses the two
contradictions with the path and the count in the message, sets `rec.Common` from `common()` so the
header describes this write, and appends or sets `no_finding`. `checkLearningArgs`, a switch over
the four states of `--no-finding` against an empty entry, then the category against
`model.LearningCategories` and the three text keys against `strings.TrimSpace`.

**`cmd/xeno/main.go`.** Two usage lines, one per form. Five flags with the comment that one of the
two statements has to be said. The table entry with `needsKey` true and `needsPhase` false, carrying
the reason the phase is optional. `cmdLearningRecord`, which resolves the phase where one was given,
reports through `o.report` for the staircase, and prints the count with the last category or that the
record states no finding.

**`.xeno/plugin/skills/xeno-learning/SKILL.md`.** "What a record is" is replaced by "The command" —
the invocation, the appending second call, the empty case, the intent level form, and **do not write
the file by hand** with the reason stated as the version comparison. The YAML block is gone, because
a skill showing both would let the reader choose the one that was wrong.

**`.xeno/plugin/skills/xeno-intake/SKILL.md`.** The instruction to write a `learning.yaml` names the
command instead.

**Tests.** Six in `internal/runner` and one at the surface, with an `assertRefusal` helper and a
`mustLearning` beside the two that already exist. The refusal table is seven cases, each also
checking that no file was left behind.

**`README.md`**, the command and a paragraph and the WP7 row. **`ASSUMPTIONS.md`**, A82, and the
"Built" list, which claimed `xeno intent start` was the last of these and now names this one.

**Measured on a throwaway tree before the tests were written**: the record written by the command
carries `0.1.0-dev+b54626e.dirty` where a hand-written one carries `0.1.0-dev`, both forms work, all
seven refusals fire, and G-Learning passes.

<!-- xeno:section:deviations -->
## Deviations from the design

**None from the design.** Every decision P2 recorded is what the code does.

**One widening in `ASSUMPTIONS.md` that the design did not list.** The "Built" section's sentence
ended "and `xeno intent start`, which takes the issue and derives the other six fields of
`intent.yaml`", which read as the end of that line of work and was true for four hours. It is
amended rather than appended to, because a list that names the last of something and then grows is a
list a reader stops trusting.

**One thing the implementation decided and the design had not.** `checkLearningArgs` is a method on
the runner rather than a free function, for no reason beyond matching its neighbours; it reads no
state and could be either. Recorded because a reviewer looking for why would otherwise find none.
