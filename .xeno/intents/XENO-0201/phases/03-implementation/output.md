---
intent: github.com/triplem/xeno#120
phase: 03-implementation
created: "2026-09-28T20:32:15Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6e72fed.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: b1dca3a04bb791ea387f0c60f532127194d4d60ee61407dbc841d163d7d547c4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: by-hand
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`internal/model`. `Project` gains the `agent` block of section 12, with the tool and the
default model. The struct modelled the language block alone, which is what made a
present source look absent.

`internal/runner`. `agent()` returns the two fields, and empty strings where the file,
the block or a field is missing. `writeDigest()` writes the summary with the frontmatter
section 5 requires of a file produced in a session, reusing `frontmatter` and
`frontmatterOrder` so that a digest and its `output.md` carry one field order. `Finish`
takes a summary and writes the digest before it evaluates. `SectionSet` writes the two
fields into `output.md` from the same accessor.

The digest carries `intent`, `phase`, `created`, `schema_version`, `runner_version`,
`plugin_version`, `language`, `context_hash`, `model` and `tool`. It carries no
`template` and no `strings_hash`, because a digest is not rendered, and no `rules_hash`,
because no rule set covers it. It carries no `secrets_hash`, and the writer's comment
says why at length: section 5 gives the runner filtering and writing, the effective
filter is a shipped pattern file that does not exist in this tree, and a hash over
nothing would assert the one claim in a digest a reader would take on trust.

`cmd/xeno`. `--summary` on `phase finish`, taking a path or `-` for stdin. Absence
cannot mean stdin as it does for `section set`, because a summary is optional here;
`readSummary` is therefore its own three lines rather than a reuse of `readMessage`. An
empty summary is refused with the reason, since a digest whose body is nothing would
pass G-Schema.

`ASSUMPTIONS.md`. A35 amended to the three fields that still have no writer.

Existing behaviour: `Finish` gained a parameter, and the two call sites in the tests
pass an empty summary. Every test passes unchanged, which is the evidence that a phase
finished without a summary is judged exactly as before.

This intent's own digests are written by the runner from here on, which is the first
time any digest in this repository was not typed by hand.

<!-- xeno:section:deviations -->
## Deviations from the design

None in the shape. The writer, its placement before the gates, the absent
`secrets_hash`, the optional summary and the two fields from the config are as P2
decided them.

One thing P2 left to the implementation and it went the other way than expected. The
requirements said `--summary` alone would read stdin, in the manner of `section set`.
That cannot work: for `section set` an absent `--file` means stdin because content is
always required, and here absence has to mean no digest at all. So a path or `-`, which
is the ordinary convention, and `readSummary` exists rather than reusing `readMessage`
for exactly that reason. AC1's wording is the thing that was wrong, not the criterion
behind it.

The order of the record is the order of the work. P0 to P2 before any code, the code
inside P3.
