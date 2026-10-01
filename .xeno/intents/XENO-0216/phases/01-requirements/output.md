---
intent: github.com/triplem/xeno#156
phase: 01-requirements
created: "2026-10-01T14:36:38Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 28f6e302105ba421aed26f73ab65468cb7db24f0e5d5985a25cd3d5e58e88632
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**The register's "Not built" paragraph says only what is not built.** The secret filter and
the digest writer are gone from it, because both are in `internal/secrets` and wired into
`phase finish`. WP12 appears there as its tracker half alone, with the branch rules port and
both adapters named as built. G-Supply and G-Secret appear with what they wait on, the
plugin and the hook, so a reader knows why they stand as `not-implemented` and does not read
it as work nobody started. The verification points name the marketplace URL as the one still
open and `phase start --export` as having answered the other.

**The register's "Built" accounts for everything in the tree.** `section set`, the three
`assumption` commands, the digest and its filter, the cost record with `cost turn`, the
branch rules port and the index reader each appear, with the assumption that governs them
where one does: A62, A63, A64, A65.

**The README's command block is the command surface.** Every command in
`cmd/xeno/main.go`'s dispatch table appears, with the flags its usage text carries, and
nothing appears that the binary does not have. A reader copying a line from the README gets
the behaviour the binary has.

**The README's trail section is true of the trail.** It states the two shapes the intents
have and `XENO-0107` as the boundary between them, and it points at `xeno intent status
--all` for the trail itself rather than carrying a count that is stale at the next intent.
Nothing in it asserts that a phase beyond intake cannot be judged.

**The coverage table has a row per package with tests in this tree.** WP8, WP11, WP12, WP13,
WP15 and WP17 are added, and every test named in a new row exists under the name given. The
WP7 row names the digest and its filter, which it predates.

**Nothing regresses.** `go build`, `go test ./...`, `gofmt -l`, `go vet ./...` and
`./xeno gate verify` all pass, the last at exit 0, because a correction to two records that
broke the suite would be a change to the tool disguised as a change to its description.

**The gap is written down once.** #156 carries the finding that no package owns either file,
and neither file absorbs it as a remark about itself.

**The figures reach #117.** The derivable proportionality figures for this intent are posted
there as a comment, counted from the tree rather than estimated, with the `gate run` timing
taken on a throwaway copy so that no sealed phase is rewritten to measure it.

<!-- xeno:section:non-goals -->
## Non goals

**No document under `docs/` is touched.** The first standing rule puts them beyond the
agent, and this intent has no disagreement with them to settle: every correction is a record
catching up with the tree.

**No owner is named for the two files.** It is the fix the finding points at and it is a
plan change, so it is a person's commit made before the code that follows from it. Writing
it here would be the absorption the third standing rule forbids, and it would also guess at
a shape the plan may want differently — a package, a line in the working method, or a check.

**No check is built that would have caught this.** A test that reads the README and compares
it against the dispatch table is a plausible answer and a bigger one than this intent, and
it belongs beside the generated reference WP16 specifies rather than beside a corrected
paragraph. It is named in the residual risk and left there.

**No artifact field is invented for the figures.** The process definition enumerates what
the artifacts carry, so a field for a proportionality figure would be a specification change
first. They go to #117 as a comment, which is where the standing request puts them.

**No earlier state of the register is rewritten.** A closed entry records when it was
answered, and correcting its history would lose that. What is corrected is the present tense
of the two summary sections.

**No paragraph that is not stale is rewrapped.** Four lines in the register's "Built" run
past the column limit and predate this change; touching them would mean editing paragraphs
that are correct, and the convention for a second change to a paragraph is to replace it,
which is a worse trade than leaving a long line alone.

<!-- xeno:section:constraints -->
## Constraints

**The documents are not editable by the agent**, which is what makes the ownership gap a
finding rather than a fix. It also decides where the finding lives: an issue, because that
is the one place this intent may write a statement about the plan.

**Prose wraps at 88 columns**, and a paragraph being changed for the second time is replaced
rather than edited into. Both summary sections of the register and the README's trail section
have been changed before, so all three are rewritten whole and read back as paragraphs.

**Every claim is counted from the tree, not estimated.** The intents are counted by phase
directory, the gates read from the gate table in `internal/gates/gates.go`, the commands from
the dispatch table in `cmd/xeno/main.go`, and every test named in the coverage table is
checked to exist under that name. A record corrected with a guess is worse than one left
stale, because the staleness at least has a date behind it.

**The figures are taken without disturbing a sealed phase.** `gate run` rewrites the
`gate.yaml` of the phase it judges, so timing one on a real phase would move its `run_at` and
dirty a committed artifact. The timing is taken on a throwaway copy of the tree;
`gate verify` writes nothing and is timed in place.

**The commit references the issue and closes it.** Conventional Commits subject without an
issue reference, `Closes #156` in the footer and in the pull request description, as
`CONTRIBUTING.md` requires of the commit that finishes the work.

**One intent, one branch.** `156-the-records-kept-by-hand`, off `main` at `2dd4dc9`, which is
the merge of #155.
