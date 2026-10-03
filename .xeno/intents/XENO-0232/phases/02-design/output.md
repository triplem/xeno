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
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**`model.Learning` and `model.LearningEntry`, beside the other section 5 shapes.** `NoFinding` and
`Learnings` are both `omitempty`, so exactly one is written and a record carrying both is not
representable by the writer — which is the shape G-Learning reads and the reason the gate needs no
change.

**`RecordLearning(key, phase, noFinding, entry)` on the runner.** The phase is a string and empty
means the intent level record, because section 10 owes one at each and `intent close` reads the
second. `learningPath` is the one place that branches on it.

**The header is rewritten on every call, not carried over.** It describes the binary doing this
write, and a second entry recorded by a newer build should say so. `rec.Common = common` after
reading, which is also why reading an existing record cannot smuggle an old version forward.

**Arguments are checked before the record is opened**, in `checkLearningArgs`, so a refused call
leaves no file and a category outside the four never reaches disk. The gate checks the same set
afterwards; two readers of one closed set, both naming `model.LearningCategories`.

**A contradiction is refused in both directions and never resolved.** An entry on a record that
states `no_finding`, and `--no-finding` on a record that carries entries. One of the two statements
is already on disk and it is somebody's; a writer that picked would be deciding which was meant.

**`--no-finding` beside any of the four keys is refused** rather than ignored, because the two say
opposite things and ignoring one would record the other silently.

**Entries append.** `rec.Learnings = append(...)`, so the order is the order recorded.

**No sealed-phase guard.** `SectionSet` has none and the reason carries: the file is inside
`artifacts_hash`, the write makes `verify` report a divergence, and `phase finish` is what judges
the phase again. A refusal would refuse the first half of a re-judgement.

**`needsPhase` is false and the command resolves the phase itself.** It is the second command whose
phase is optional, after `learning record`'s sibling cases in `gate verify` and `intent status`, and
the dispatch field already allows it.

**The skill says not to write the file.** A skill that showed the YAML and a command would leave the
reader to choose, and the one they would choose is the one that was wrong.

<!-- xeno:section:alternatives -->
## Alternatives

**`phase finish` stamping an existing record's header.** Refused. It would fix the version and leave
the rest of the file unwritten — the category still unchecked until the gate, the empty case still
a convention rather than an argument — and it would rewrite a file inside `artifacts_hash` as a side
effect of a command whose subject is the digest. The problem is that the artifact has no writer, not
that its header is stale.

**A `--file` taking the whole record as YAML.** Refused: it is the hand-written file with a command
in front of it. The header would be the author's again, which is the defect.

**Replacing the record on each call rather than appending.** Refused. A phase that learned two
things would have to pass both every time or lose the first, and the second call is usually made
because the second observation arrived later.

**Resolving a contradiction by clearing the other field.** Refused in both directions. `no_finding`
is a statement that there was nothing to record, and an entry is a statement that there was; the one
on disk was made by somebody and the writer is not the one to withdraw it.

**Letting the gate keep sole ownership of the category check.** Refused as the weaker arrangement:
G-Learning reports it after the file exists and after the phase is judged, where the writer refuses
it in the call that would have made the mistake. Both check it, which is two readers of one set in
`model` rather than two definitions.

**A separate `learning` command group with `add`, `clear` and `show`.** Refused: `clear` is the
rewrite this process forbids, `show` is `cat`, and the MCP surface is a budget. One verb.

**Writing the record automatically at `phase finish` from nothing.** Refused — it would be a record
nobody wrote, which is where this started, with the runner as the author of an observation it cannot
have made.

<!-- xeno:section:impact -->
## Impact

**`internal/model/model.go`.** `Learning` with the common header inline and the two body fields
`omitempty`, and `LearningEntry` with section 10's four keys in the order `LearningKeys` lists them.

**`internal/runner/runner.go`.** `learningPath`, which is the only place the phase's absence
branches; `RecordLearning`, which validates, reads an existing record, refuses the two
contradictions, rewrites the header and appends or sets `no_finding`; and `checkLearningArgs`, which
refuses before anything is read.

**`cmd/xeno/main.go`.** Two usage lines, the five flags, the table entry with `needsKey` true and
`needsPhase` false, and `cmdLearningRecord`, which resolves the phase where one was given and
reports what the record now holds — the count and the last category, or that it records no finding.

**`.xeno/plugin/skills/xeno-learning/SKILL.md`.** "What a record is" becomes "The command": the
invocation, the appending second call, the empty case, the intent level form, and the instruction
not to write the file with the reason — the header is the runner's and a typed one records
`0.1.0-dev` where the three artifacts beside it record the commit.

**`.xeno/plugin/skills/xeno-intake/SKILL.md`.** The line that said to write a `learning.yaml` now
names the command.

**Tests.** Six in `internal/runner`: the header agreeing with the verdict beside it, entries
accumulating in order, the empty record, both contradictions, a table of seven refusals each
checking that no file was left behind, and the intent level record carrying no phase. One in
`cmd/xeno` driving both forms and the staircase through the command line. An `assertRefusal` helper,
because every refusal test in that file wants the same two assertions.

**`README.md` and `ASSUMPTIONS.md`.** The command in the surface list with a paragraph, the WP7 row,
A82, and the "Built" list which claimed `xeno intent start` was the last of these.

**No gate, no rule, no document, no existing artifact.** `gate verify` is unchanged.
