---
intent: github.com/triplem/xeno#239
phase: 03-implementation
created: "2026-10-05T08:58:14Z"
schema_version: "1.0"
runner_version: dev+3f5ad34
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d5862d3d7eea0cb616dcbcc7407059867764dfea06cc55723a411a4b82451d34
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

Four `git mv`s, each recorded as a rename so `git log --follow` still reaches the history:

| from | to |
|---|---|
| `ASSUMPTIONS.md` | `docs/assumptions.md` |
| `SUPPLY-CHAIN.md` | `docs/supply-chain.md` |
| `CLAUSE-READERS.md` | `docs/clause-readers.md` |
| `M0.md` | `docs/m0-gate-job.md` |

The reference sweep covers eleven files and 27 occurrences. `README.md` has four, three of
them reflowed because the paragraph went over 88 once the path grew by five characters.
`CLAUDE.md` has one, in the paragraph that names where things are written down, reflowed
for the same reason. `CONTRIBUTING.md` has one. Five workflow files have six between them,
all in comments. `internal/runner/exchange.go` has one, in a comment. `docs/assumptions.md`
has eight, in the `Where` and `Why` columns of A20, A24, A26, A28, A44, A90 and A91.

`docs/supply-chain.md` and `docs/clause-readers.md` change by zero lines. Both are
referenced by others and reference none of the four themselves, which is why the rename is
their whole diff.

Inside `docs/`, a sibling is named by its full repository-relative path rather than by its
bare filename. That is the convention already in the tree: `docs/assumptions.md` names
`docs/process-definition.md` and `docs/implementation-plan.md` that way, so the sweep
follows it rather than introducing a second style.

`docs/m0-gate-job.md` takes two changes beyond the sweep. Its layout block now lists
`docs/assumptions.md` and `docs/m0-gate-job.md` in the `docs/` group, each annotated with
what it was renamed from, in the style the block already used for the four renames it
recorded; the root line keeps `cmd/`, `internal/`, `vendor/`, `go.mod`, `README.md` and
`.gitignore`. And two lines under the title say it records how M0 was reached once rather
than giving a procedure to follow.

`docs/assumptions.md` gains A94, one row for every reference this move cannot correct: the
274 sealed ones and the plan's line 2009.

<!-- xeno:section:deviations -->
## Deviations from the design

One deviation from the acceptance criteria, declared in them before the work began rather
than discovered after: criterion 2 exempts `docs/implementation-plan.md` line 2009, which
still reads "its other choices are recorded in its `ASSUMPTIONS.md`". The plan is normative
and the first standing rule puts it beyond the agent, so the correction is a person's
commit. A94 records it.

Two repairs were taken beyond the sweep because the sweep caused them. Three paragraphs in
`README.md`, one in `CLAUDE.md` and one in `docs/m0-gate-job.md` went past 88 columns when
`ASSUMPTIONS.md` became `docs/assumptions.md` and `M0.md` became `docs/m0-gate-job.md`, and
each was replaced as a paragraph and read back as prose rather than rewrapped at the break,
which is what this project's conventions ask for where a paragraph changes twice. The
`README.md` passage that named `docs/m0-gate-job.md` twice in four lines now names it once
and says "that guide" the second time, because repeating a path that long reads worse than
the pronoun.

One pre-existing defect was found and left. `docs/m0-gate-job.md` line 8 breaks a sentence
mid-phrase, "This takes one intent, one phase and one CI job. Every / block below was run
end to end", and has done since the file was written. It is not this intent's and fixing it
would put an unrelated prose repair in a move commit, where a reviewer reading the diff for
renames would have to stop and judge it. It belongs to WP16 with the link checker.

Two counts in `docs/assumptions.md` and in P2's impact table are this tree's rather than
#239's. The issue counted 189 references on 2026-10-04, of which 17 were correctable; today
it is 301, of which 27 are. Nothing was wrong with the issue's arithmetic — five intents
landed in between, each citing documents — and the figures are restated rather than quoted
because a count that moves is evidence only with its date.

No sealed artifact was touched. Nothing under `.xeno/intents/` appears in this diff but
this intent's own directory.
