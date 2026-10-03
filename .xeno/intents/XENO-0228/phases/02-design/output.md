---
intent: github.com/triplem/xeno#181
phase: 02-design
created: "2026-10-03T12:39:12Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6cbeac4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e8bf10ef8058c02b7658172f52832a0a2377a7acd7b4eeb47890d512f6fe1567
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

**`ToolVersion` is a field on `Runner`, set by `parse` for every command.** It joins `Base`, `Head`
and `EvidenceFrom`, which are the other inputs of a run that are not arguments of a method. Two
commands act on it and the rest ignore it, which is why it is documented on the `common:` line of
the usage rather than against one subcommand.

**`SectionSet` writes it only when reported, and a report never erases.** `front["tool_version"]` is
set when the value is non-empty and left alone otherwise, so the carry-over of an existing
artifact's frontmatter does the rest. A phase whose first section named the version and whose last
did not is one session.

**`writeDigest` copies from `output.md`, with a report winning.** The order is: reported value,
then `recordedToolVersion`, then nothing. Reported first because the harness saying what it is
cannot be overruled by a file, and because a digest can be written for a phase with no `output.md`
at all.

**`recordedToolVersion` reads the frontmatter and answers empty for every failure.** No artifact, no
frontmatter, no field: all the same answer, which is the field absent. It uses `fm.ReadFront`, which
is the reader this package already has for exactly this.

**The copy is in `writeDigest` and not in `Finish`.** `writeDigest` is the function rerun by every
`phase finish` and the one that builds the frontmatter, so putting the decision anywhere else would
leave a path that writes a digest without it.

**The flag is `--tool-version` and not `--harness-version`.** The field in section 5 is
`tool_version` and the one beside it is `tool`; a flag named after the field is a flag somebody can
find from a gate's finding, which names the field.

**The skills carry the instruction, in the same words in all six.** A flag the agent does not know
about is a field with no writer by another route, and the skills are where the plugin tells an agent
what to run.

<!-- xeno:section:alternatives -->
## Alternatives

**`XENO_HARNESS_VERSION`, read from the environment.** The right long term answer and refused here
for one reason only: section 7 enumerates three `XENO_*` variables and a fourth is a change to that
list, which the first standing rule makes a person's commit made before the code. Put to the
maintainer before any code was written, with the flag as the alternative, and the flag was chosen.
It would become this flag's fallback rather than its replacement, so nothing here has to be undone
for it.

**The runner asking the harness for its version.** Refused by section 7 in as many words: the
moment the runner behaves differently per harness, the tools stop being interchangeable. It is also
the most agent specific question there is, in a binary whose claim is that it holds no agent
specific logic.

**`agent.tool_version` in `project.yaml`.** Refused twice over. It is an Appendix A addition, so a
specification change by the rule above; and a harness version is a property of a session, so a
project file would state it once and be wrong after the next upgrade — the field would go from
absent to confidently stale, which is worse than what A35 was protecting against.

**A second input to `phase finish`.** Refused because it moves the hand edit rather than removing
it: `writeDigest` runs on every finish, so the value would have to be supplied on every one, and
the criterion about the second run would fail by design.

**Defaulting the digest's value to the runner's own version.** Refused as the exact failure A35
names. `runner_version` and `tool_version` are different facts and a field filled with the wrong
one reads as a harness nobody used.

**Teaching G-Schema to stop requiring the field.** Refused: the field is in section 5's session
group and the gate is right. The gap was the writer, and relaxing the gate would have hidden it in
every project rather than in this one.

<!-- xeno:section:impact -->
## Impact

**`internal/runner/runner.go`.** `ToolVersion` on `Runner` with the reason it is an input;
`SectionSet` writing it when reported; `writeDigest` preferring a report and falling back to the
artifact; `recordedToolVersion`, new, reading `output.md`'s frontmatter and answering empty for
every failure.

**`cmd/xeno/main.go`.** The `--tool-version` flag, the `toolVersion` field on `opts`, the assignment
to the runner in `parse`, and two lines of usage: the flag on the `common:` block, which now has a
continuation line, because it is read for every command and acted on by two.

**`.xeno/plugin/skills/`.** The same paragraph in all six phase skills, naming the flag, why the
runner cannot know the value, and what the digest does with it. The learning skill has no
`section set` and gains nothing.

**Tests.** Five in `internal/runner`: both artifacts from one report, the second finish keeping it,
a report beating the record, absence staying absence in both files, and a later section write not
erasing it. One in `cmd/xeno` at the surface, driving `section set` with the flag and `phase finish`
without. A `set` helper, because these tests are about frontmatter and never about the resolved
template.

**`ASSUMPTIONS.md`.** A35 amended a third time — it is satisfied, not contradicted, and the row's
claim that two fields remain writerless was already one out of date. A80 for the decision and for
the entry point it leaves unbuilt.

**`README.md`.** The flag in the surface list, a paragraph on what it is for, and the WP7 row.

**No gate and no rule changes**, so no verdict behind this intent recomputes differently. **No
document changes**: section 5 already has the field and section 7 already has the prohibition.

**What is not touched.** The artifacts already written, which carry a hand-typed value inside a
sealed hash, and the behaviour of any repository that never passes the flag.
