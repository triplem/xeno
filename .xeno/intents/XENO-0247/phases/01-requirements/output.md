---
intent: github.com/triplem/xeno#242
phase: 01-requirements
created: "2026-10-05T12:10:57Z"
schema_version: "1.0"
runner_version: dev+3f5ad34.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7069385d8e026615b8f40dad35dd739d01ecb40161d2d6a1690eedfb182aaf9b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. `xeno review answer RULE --intent KEY --result R [--note TEXT]` exists, the usage text
   lists it, and it appends an entry to `review_checklist` in the review phase's
   `output.md`, creating the block where there is none and leaving the body byte for byte.

2. The frontmatter it writes is in `frontmatterOrder`'s order, which already places
   `review_checklist` last of the four lists, and an artifact amended by it re-reads as the
   same `model.Output` the gate reads.

3. A rule that is not a `review` rule of the effective set is refused, naming the id and
   saying what the set holds. A `result` outside `model.ChecklistResults` is refused naming
   the three. A `deviation` or `not-applicable` with no note is refused, and `met` without
   one is accepted, which is `model.ChecklistNeedsNote`'s split.

4. Answering a rule already answered replaces that entry in place, keeping its position in
   the list, and the command says it replaced rather than appended.

5. The command prints the `review` rules of the set still unanswered after the write, and
   says the checklist is complete when none are.

6. No argument for the phase. The checklist is the review phase's artifact, as the gate's
   own guard says, so the command resolves it and offers no choice.

7. No `gate.yaml` is read. A write after a verdict is what `phase finish` reports as a phase
   changed after its verdict, which is `exchange.go`'s stated design and #216's mechanism.

8. Tests cover: append to an absent block, append to a present one, replace on duplicate,
   each of the three refusals, `met` without a note accepted, the unanswered-rules report,
   and that the body survives byte for byte.

9. This intent's own P5 checklist is written with the command and not by hand, which is the
   one check that the writer works where it is meant to be used.

10. `go build`, `go test ./...`, `go vet ./...` pass, `gofmt -l` outside `vendor/` prints
    nothing, and `./xeno gate verify` exits 0 with the 357 verdicts that exist now intact
    and the growth this intent's own six phases.

<!-- xeno:section:non-goals -->
## Non goals

Not a change to the checklist's shape. A68 fixes `rule`, `result`, `note` and `source`, and
the second standing rule makes an addition a specification change first.

Not a change to G-Policy. The gate stays the authority on a sealed phase. The writer
duplicates three of its checks at the moment of writing, deliberately, and sharing
`ChecklistResults`, `ChecklistNeedsNote` and `rules.Effective` is what keeps the two from
disagreeing.

Not a lens entry. Section 12 gives a lens `source: lens` and no rule id; this command takes
a rule and cannot express one.

Not a `--source` flag. Every entry this command writes is a person answering a rule, so a
flag whose only honest value is the absent one is a field offering a lie.

Not a reader. Nothing lists or prints the checklist as a document; the command reports what
is unanswered because the writer knows it anyway, and a view of the whole is WP18's.

Not a migration of the checklists that exist. Every P5 in this repository has a hand-written
block that G-Policy passes, and rewriting one would change its `artifacts_hash` and stale its
verdict. The writer applies from here.

Not a statement in `CLAUDE.md` about hand edits. XENO-0246's learning asked for one, and
section 10 routes a learning through a merge request against the rule set rather than letting
the intent that noticed it apply it; that is a separate change and this intent is the code
half it was waiting on.

<!-- xeno:section:constraints -->
## Constraints

The field sits inside `artifacts_hash`. The writer amends `output.md`, so it changes the
hash by construction, and that is why it must not try to be clever about verdicts: the
honest behaviour is to write and let `phase finish` report a phase that changed after its
verdict, which is what the other two frontmatter writers do.

The two sets belong to other packages. `model.ChecklistResults` and
`model.ChecklistNeedsNote` are the gate's authority on a result, and `rules.Effective`
filtered to `rules.Review` is its authority on a rule. The writer imports both rather than
restating either, because a second copy of a three-element set is a thing that drifts.

A74 binds what the writer can promise. G-Policy judges a phase only against the rule set
the artifact records in `rules_hash`, and the writer compares against the set that resolves
now. The two agree whenever the rule tree has not changed since `section set` wrote the
hash, and where they diverge the gate is right and the writer's refusal is advice. Worth
stating because a refusal that claims more than it can know is the defect this project keeps
finding.

One global flag set in `cmd/xeno/main.go`. `--result` already exists for `evidence declare`
with a different meaning, so its help text has to serve both or the flag needs a second
name; `--note` is new. The positional argument goes through `o.finding`, which is how
`gate approve` and `obligation close` take theirs.

Replace rather than append on a duplicate, decided by the maintainer, and the position in
the list is kept so that a re-answer does not reorder the block and show up as a larger
diff than it is.

One intent, one branch, `Closes #242`, and the issue carries `wp4`.
