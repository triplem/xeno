---
intent: github.com/triplem/xeno#156
phase: 02-design
created: "2026-10-01T14:37:49Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 97dcf60b24b8ea084fff10d5c727ae8048bd1fe60afb752e482d8a5645809618
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: by-hand
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**Three paragraphs are replaced whole.** The register's "Not built", the README's opening and
the README's trail section have each been changed before, and the convention for a second
change is to replace rather than edit into, then read the result back as a paragraph. Editing
into the trail paragraph is in fact how it got into its present state: the clause about WP8
survived a change that should have carried it away.

**What postdates the register's "Built" list is a new paragraph, not an edit into it.** That
list is a record of what was built when it was written, and splicing six months of work into
its sentences would make it a record of nothing. The new paragraph says in its first line
that the list above reads as complete and is not, which is the honest relation between them.

**The README's trail section states the two shapes and no count.** Fifty-one intents hold one
phase and twenty hold six today, and both numbers change with the next intent. The boundary
`XENO-0107` does not change, so the boundary is what the paragraph carries, with
`xeno intent status --all` named as the place to read the trail itself.

**The command block is written by hand against the dispatch table, not generated.** WP16
specifies a generated command reference, and a generator written here would be the second
place to maintain it until that one exists. What this intent does instead is check every line
against `cmd/xeno/main.go` and say in the acceptance that this is the authority.

**The coverage table keeps its existing order and appends.** Its rows run WP1, WP2, WP3, WP9,
WP10, WP5, WP6, WP7, which is the order the packages were covered in rather than the plan's
order. Sorting it would be a larger diff than the addition and would lose that history; the
new rows go after WP7 in package order, which leaves the table readable in both directions.

**The heading becomes `# xeno`.** It read `# xeno, core`, and the paragraph under it no longer
describes only the core. A heading that contradicts its own first sentence is the smallest
version of exactly the drift this intent is correcting.

**The finding is an issue labelled `wp16`, with the label disclaimed in the body.** Every issue
carries the label of its package and no package owns these two files, so the nearest one is
used and the body says it is nearest rather than covering. Inventing a label for ownerless
work would be inventing a rule, which the second standing rule forbids as plainly as it
forbids inventing a field.

<!-- xeno:section:alternatives -->
## Alternatives

**Delete the trail section rather than replace it.** It would end the false claim in one line
and lose a true one: the P0-only shape of the first fifty-one intents is a real property of
this repository that A20 explains and that reads as abandonment to anyone who has not found
`M0.md`. The README is where that reader is, so the paragraph stays and gets corrected.

**Stamp each record with the commit it was last checked against.** A line reading "current as
of 2dd4dc9" is cheap and tells a reader the age of what they are holding. Rejected because a
date is not a check: it would have been written on the same day as the sentence about WP8 and
would have aged with it, and it invites exactly the trust the next stale paragraph needs to
not have. The honest version of that idea is a check, which is named in the residual risk.

**Generate the command block from the usage text now.** `cmd/xeno/main.go` holds both the
dispatch table and the usage string, and a test already asserts they agree, so extending that
into the README is a short step. Rejected as the wrong package: WP16 specifies the generated
reference and says anything derivable is derived there. A generator in this intent would make
two generators, and the second one gets deleted by the package that was meant to own it.

**Fold the two files into one.** The register and the README overlap in what they claim about
the tree, and one file could not disagree with itself. Rejected because they have different
readers: the register is for whoever is building this and wants to know what was decided, the
README is for whoever arrives. The duplication is not the problem; nobody owning either is.

**Write the ownership finding into `ASSUMPTIONS.md` as a new assumption.** Tempting, since the
register is where decisions taken while building the core live. Rejected because this is not a
decision taken here, it is a gap in the plan, and the standing rule routes it to an issue. An
assumption row would also leave it where only a builder looks, when the people who could close
it are the ones reading the plan.

**Sort the coverage table into plan order.** Rejected for the diff, which would touch every
row and show as a rewrite of coverage that did not change. Worth revisiting the next time a row
is added, because the argument against it weakens as the unsorted part shrinks.

<!-- xeno:section:impact -->
## Impact

**No code, no gate, no verdict.** Nothing under `cmd/` or `internal/` changes, no gate reads
either file, and no `artifacts_hash` in the tree depends on them. `./xeno gate verify` is run
to prove that rather than to establish it: a records-only change that moved a verdict would
mean one of them was inside a hash, which would be a finding of its own.

**Orientation changes, and the authority moves.** The README's opening now points at the
register's last two sections for what is built and what is not, which makes those two sections
the single statement and the README a pointer to it. One statement, one place, which is the
measure WP16 names for prose nothing can check.

**The claim a reader can no longer draw.** That this tool has never judged a phase beyond
intake. That was the most damaging sentence in either file, because it is read as the tool's
own admission by someone deciding whether to look further.

**The register's "Not built" becomes shorter and more specific**, and two of its entries now
carry what they wait on. A reader can see that G-Supply and G-Secret are blocked on the plugin
and the hook rather than unstarted, which is the difference between a gap and a queue.

**WP16 gains the data point rather than the work.** The package's own done-when says no
generated page is maintained by hand; this intent adds the observation that two hand-kept
records outside its scope drifted for twenty intents, which is an argument for the ownership
question being settled when that package is planned rather than after.

**What is not improved is the thing that caused it.** After this, both files are accurate and
still unowned, so the next twenty intents can repeat it exactly. The finding is the whole
mitigation, and a finding is a record and not a guard.
