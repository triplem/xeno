---
intent: github.com/triplem/xeno#242
phase: 02-design
created: "2026-10-05T12:11:34Z"
schema_version: "1.0"
runner_version: dev+3f5ad34.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8b81d972ee520502e9fedac2a5355586d1947ada78138ba21546bdb3444d4c37
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

The writer is a new file, `internal/runner/review.go`, rather than a third function in
`exchange.go`. That file is the exchange section 8 defines, a question and the decision that
settles it, and its package comment says so; the checklist is section 9's and answers a rule
rather than a question. It uses `exchange.go`'s two helpers, which is the reuse that matters.

`ReviewAnswer(key string, e model.ChecklistEntry) (*ReviewAnswered, error)`, with no phase
parameter. The gate judges the checklist only on the last phase, so the writer resolves
`model.Phases[len(model.Phases)-1]` itself. This follows `ScopeSet`, which takes no phase for
the same reason in the other direction: the scope is P0's artifact and the checklist is P5's.

`ReviewAnswered` carries the entry, whether it replaced an existing answer, and the ids still
unanswered. It is not recorded anywhere, for `ScopeResolved`'s reason: the artifact records
the answers, and a list of what was outstanding at one moment describes neither the phase nor
the verdict.

Replace keeps the entry's index. `slices` is not needed: the loop that finds the rule writes
through the index it found, so the block's order is the order the rules were first answered,
and a corrected answer is a one-line diff rather than a move.

The three refusals come before the artifact is read. A bad `result` or a missing note is
knowable without touching the file, and refusing early means a malformed call cannot leave a
half-amended artifact. The rule check needs the rule tree but not the artifact, so it is
second, and only then is `output.md` opened.

The unanswered report counts from the rules, exactly as G-Policy does. Walking the entries
would call an empty checklist complete, which is the failure the gate's own comment warns
about, and a writer that reported "nothing left" on an empty block would be worse than
silent.

`--note` is a new flag. `--result` is reused from `evidence declare` with its help text
rewritten to name both meanings, because one global flag set is the shape `main.go` has and
a second name for the same word would be the confusing half of both options.

<!-- xeno:section:alternatives -->
## Alternatives

Refusing on a written verdict was #242's proposal and is rejected, with the issue's own
reasoning found wrong. It claimed the decision commands do this; they do not, and
`exchange.go` says why in its package comment. D-6 is where the confusion came from: it
settles that a decision command refuses on a stale verdict, and it is about `gate approve`
and `gate override`, which carry `against`, the `artifacts_hash` the decision was taken on.
A checklist entry carries no such field, so there is nothing for it to assert untruthfully.

An optional `--phase` was #242's other proposal and is rejected because the gate's guard
makes it a choice of one.

Extending `section set` was considered. The checklist is frontmatter, and `SectionSet`
re-renders the body from the template, which is exactly what `exchange.go` gives as the
reason the other two frontmatter writers exist. Nothing changes that here.

Appending on a duplicate was the alternative the maintainer decided against. It would keep
every answer a rule ever had, which reads as a history and is not one: nothing records when
an entry was written or which came first, so two entries for one rule are an ambiguity
rather than a record. `section set` overwrites a section for the same reason.

Writing the block from a document on standard input, as `question record` and `scope set`
take theirs, was weighed and rejected. Those two carry nested structures that do not fit
flags; an entry is three scalars, which is `decision record`'s and `assumption record`'s
case, and they take flags.

A `--source` flag for lens entries was considered and refused as an invented field in
practice if not in the schema: the only value this command could honestly write is the
absent one, so the flag would exist to be left alone.

Teaching `next.go` to report unanswered rules was considered and rejected as too wide. The
`running` suggestion is shared by every phase and reports missing sections; adding rules to
it would put a rule-set resolution in the path of every command that prints a next step, to
say something only the review phase can be missing. The command prints it instead, which is
where the information already is.

<!-- xeno:section:impact -->
## Impact

One new file and three touched. `internal/runner/review.go` is the writer.
`cmd/xeno/main.go` gains a command table row, a `--note` flag, a usage line and a `cmdReviewAnswer`;
`--result`'s help text is rewritten to serve both commands that take it. A test file joins
the runner's. Nothing else changes: not the model, not the rule engine, not a gate, not a
template, not an artifact shape.

Nothing existing changes behaviour. The writer is new surface, `review_checklist` was already
in `frontmatterOrder`, and `amendFront` already wrote any field it is given. No checklist
that exists is touched, so the 357 verdicts stand.

What closes is the last of a family. #188, #195, #208, #217 and #220 each gave an artifact a
writer it was specified to have and lacked; this is the one remaining field of the four
frontmatter lists, so after it no artifact content in this repository is produced by hand
except `tool_version`, which has a flag precisely so that it is not.

What a P5 author gains is a refusal instead of a red verdict. Today a wrong rule id or a
`deviation` with no note is found by `phase finish`, after the sections are written and the
phase is sealed, and the repair is another edit inside `artifacts_hash` followed by another
judgement. The writer moves all three to the moment of the mistake.

The honest limit is A74. The writer compares against the rule set that resolves now, and
G-Policy judges against the set the artifact recorded in `rules_hash`. Where the rule tree
changes between `section set` and `review answer`, the writer can accept a rule the gate will
reject or refuse one it would have accepted. Neither is silent — the gate still decides — but
the refusal is advice rather than a guarantee, and this phase says so rather than letting a
reader infer more from a command that refuses than it knows.

The second limit is that nothing makes the writer compulsory. A hand edit still works and
still passes, because the field is ordinary frontmatter and no gate can tell which wrote it.
What this intent removes is the necessity, not the possibility, and the statement that the
hand edit is now avoidable belongs to the `CLAUDE.md` change XENO-0246's learning asked for,
which section 10 routes through review rather than through this commit.
