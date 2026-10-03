---
intent: github.com/triplem/xeno#188
phase: 02-design
created: "2026-10-03T21:20:37Z"
schema_version: "1.0"
runner_version: dev+efc44a9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bfd2d13a3817d1fad98faea469c685df66186e9d525a8341d580fea9c3061a5e
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

Nothing here is a decision in the sense section 8 gives the word, so nothing here
has a `decisions` entry. A design choice an agent makes with a reason is design;
a choice between workable paths that needed a person is D-1, and it is in P1's
frontmatter where the person who made it is named. That the template's section and
the frontmatter field share a word is half of what #188 is about, and this phase is
the first one in which the two are visibly different things.

**Two commands, each taking the channel its shape asks for.** `xeno question record`
reads the entry as YAML, on stdin or from `--file`, because up to five options
with a consequence each, a recommendation and a free entry do not fit flags without
inventing a separator. `xeno decision record` takes flags, because its entry is six
scalars. `section set` already reads stdin and `assumption record` already takes
flags, so neither channel is new to this surface.

**The input is the artifact's own shape.** What `question record` reads is an
`open_questions` entry as section 5 defines it: `text`, `options` with `text`,
`consequence`, `recommended` and `free`, or `no_options: true`. No translation layer,
so the gate's shape checks are the only authority on what a question is, and nothing
can be well formed on the way in and malformed in the file.

**The key and the id are assigned, not given.** `question record` writes the next
`Q-n` of the intent and `decision record` the next `D-n`, each continuing the highest
the intent already holds rather than counting entries, which is `nextAssumptionID`'s
reason and holds for the same case: a question withdrawn from a draft must not hand
its key to the next one. Both are per intent, as the register's `A-001` is, and the
`D-` table in `ASSUMPTIONS.md` is a separate sequence in a separate file exactly as
its `A80` is separate from a register's `A-001`.

**The refusals are the gate's own checks, called before the write.** `questionShape`
and `decisionShape` become exported and the runner calls them on the entry it is about
to append; a finding becomes a refusal with the gate's sentence in it. Uniqueness
is the one check that widens: the gate asks it of a file and the command asks it of
the intent, because a key is referenced by `resolves` from any later phase and two
phases with a `Q-1` each would make `resolves: Q-1` ambiguous.

**`--reason` is the rationale, and `--by` is the person.** Both words already mean
this on this surface: `gate approve --by --reason` and `assumption confirm --by`. A
decision's field is `rationale` and a withdrawal's is a reason, and section 8 calls
them the same thing, so one flag carries it. `--chosen` is new and `--proposed-by`
is new, and `--withdraw` is section 8's third exit: a withdrawal carries its reason
and its person and no chosen option, which is what `decisionShape` already permits.

**The writer amends the frontmatter and leaves the body alone.** A new unexported
`amendFront` reads `output.md`, splits it, unmarshals the frontmatter into the
same map `SectionSet` uses, replaces the one key with the typed slice, and writes
`frontmatter(front)` back with the body byte for byte. Not through `SectionSet`:
re-rendering the body is the template's business, and a command that has nothing to
say about the body should not be the reason it is rendered again.  `frontmatterOrder`
already places both keys, so the position is not decided here.

**It refuses a phase with no `output.md`.** The artifact is created by its first
section write, and a question belongs to a phase that has one. The refusal names
`section set` rather than creating a file whose frontmatter would carry whatever
the runner could work out.

**A write after a verdict is the mechanism that already exists.** Neither command
looks at `gate.yaml` and neither refuses a judged phase. The artifact changes, its
hash changes, and the next `phase finish` says the phase changed after its verdict
was written, which is what `section set` and `learning record` already do and what
#216 built the state for.

**The figure is computed where the state is.** `IntentSummary` gains the two counts,
filled on the walk `summarise` already does over the phases, and an `AskedNothing`
predicate that is true where the intent has reached `05-review`, is not abandoned,
and both counts are zero. Abandoned is excluded for G-Questions' own reason: abandoning
means giving up, and a figure that reproached it would only produce invented questions.

**Both forms of `intent status` read that one predicate.** The listing marks the row
after its columns, where the unreadable-record note already goes, and the one-intent
form prints a line after the table, where the truncation note already goes. Neither
changes a column width, which is what #136 settled.

<!-- xeno:section:alternatives -->
## Alternatives

**Flags for the question too, with the options paired positionally.** `--option`
repeated, each `--consequence` attaching to the option before it, `--recommended`
marking the current one. Go's flag package parses in order, so a custom `flag.Value`
could do it. Rejected: the pairing is invisible in the help text and wrong silently
when the order slips, and the thing being typed is a nested structure whichever way
it is spelled. Stdin costs a heredoc and the shape is then the one the gate checks.

**YAML for the decision as well, for symmetry.** Rejected because the symmetry is
the wrong one to keep: `assumption record` takes flags for a flat record and `section
set` reads stdin for a nested one, so the surface is already consistent about this,
and a decision read from stdin would put `--by` inside the document where nothing
could require it as a flag.

**One command for both, `xeno exchange record` or similar.** Rejected: the two halves
are raised in different phases by different parties, one is nested and one is flat,
and a command that did both would need a mode flag to say which. The register is
two commands for the same reason, `assumption record` and `assumption confirm`.

**A generic frontmatter writer, `xeno front set KEY`.** It would serve `evidence` and
`review_checklist` as well, which are the same gap. Rejected, and this is the one worth
stating plainly: a generic writer cannot refuse anything, because the refusals are
what each block means. `question record` exists to say that two invented options are
not a question. #208 and WP7's checklist get their own writers, with their own refusals.

**Letting the input carry its own key.** Rejected: a stable key is section 8's
requirement and choosing it is not the author's decision to make, and the hand
written trail shows what happens when it is — four intents each with a `Q-1`,
which is harmless only because none of them was ever resolved from another phase.

**A gate instead of a figure.** Section 8 and #188 both rule it out and the non
goals say why. It is recorded here because it is the first thing a reader of the
issue's table will reach for.

**A column in the listing instead of a row mark.** Two more columns of numbers on
every row, to carry a fact about a handful of intents. Rejected: #136 settled the
widths, and the listing is read to see what is happening rather than to total the
trail. The counts are available to whatever asks `summarise`, which is where a later
report would read them.

**Printing the figure in `gate run` or `phase finish` as a note.** Rejected for now:
those outputs are a verdict and a next step, and a note that is neither would be
read as one of them.  `intent status` is the command whose whole output is a figure.

<!-- xeno:section:impact -->
## Impact

**`internal/model`.** Nothing changes. `Question`, `Option` and `Decision` are
written as they stand, which is the point of the intent.

**`internal/gates`.** `questionShape` and `decisionShape` become exported,
unchanged in behaviour, so that the runner refuses on the same checks the gate would
report. `questions`, G-Questions itself, does not change. No gate is added and no
gate's verdict moves.

**`internal/runner`.** Two methods, `RecordQuestion` and `RecordDecision`, and
the unexported `amendFront` they share. `IntentSummary` gains two counts and the
`AskedNothing` predicate, and `summarise` fills them. `SectionSet` is not touched;
`frontmatter` and `frontmatterOrder` are read and not changed.

**`cmd/xeno`.** Two entries in `commands`, both needing a key and a phase; `--chosen`,
`--proposed-by` and `--withdraw` in `parse`; two command functions; two lines in
`usage`. The row mark in `cmdIntentList` and the line in `cmdIntentStatus`.

**The artifacts.** No existing artifact changes. This intent's own P0 and P1 carry
a hand written `open_questions` and `decisions` block, which is the bootstrap
case the P0 learning records: the commands cannot write the question that scoped
them. Everything after P2 goes through the commands, and `gate verify` stays at exit
0 over every verdict in the repository because nothing sealed is touched.

**The documents.** None. Section 5 defines both blocks, section 8 defines the exchange,
and the plan's WP5 already says that a question needs no file of its own and that a
`decisions` entry two phases later resolves it. This intent is that sentence being
true of the tool and not only of the schema.

**The trail, which is the measurable part.** After this intent the count of
`decided_by` under `.xeno/` is one rather than zero, one intent has a `decisions`
block, and G-Questions has something to verify in its P5 rather than nothing. The
figure then reports on the intents that follow, which is the half of #188 that cannot
be closed by the intent that builds it.

**What stays open.** `evidence` and `review_checklist` are still written by hand
or not at all, which is #208 and WP7's checklist. Nothing compares `proposed_by`
with `decided_by`. Nothing counts questions across intents, which is #117's shape
of report rather than this one's.
