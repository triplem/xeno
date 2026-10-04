---
intent: github.com/triplem/xeno#217
phase: 02-design
created: "2026-10-04T19:56:19Z"
schema_version: "1.0"
runner_version: dev+dce3eab
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 55d12062332bcb3a6bdd4151b20f9d59091413375abeac7e56a095ed0a59671e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - chosen: 'Widen XENO-0245 to carry #236 as well. P0''s scope and P1''s acceptance criteria grew, both phases were re-judged, and the branch closes both issues.'
      decided_by: triplem
      id: D-3
      proposed_by: claude-opus-5
      rationale: 'The refusal this intent adds is not safe without the fix, which is what makes them one piece of work rather than two: a mandatory scope at P0 plus a staleness check that reads the gated phase''s own lock makes every implementation intent report itself out of date for doing its job. P1 already recorded that the writer and the refusal ship in one commit, and the fix joins that constraint for the same reason. The commit convention takes a keyword per issue, ''Closes #217, closes #236'', so one branch carrying both is what CONTRIBUTING.md describes rather than an exception to it. Against its own intent run first: a full six-phase intent for close to a two-line fix, this one sitting open across it, and its P1 approved against documents that would have moved again by the time it resumed. That option had a genuine advantage, that limit 1 would have been in place to silence the rebase noise when this intent rebased, and it was not worth the ceremony. Against shipping the writer and deferring the refusal: it withdraws a constraint P1 recorded, which costs the same re-judge for less result, and leaves #217''s done-when unmet with a writer nothing asks for, which is how the artifact came to be specified and unwritten in the first place. The cost accepted and paid: re-judging P0 staled P1''s recorded predecessor hash. Rather than a third release, P1''s verdict was removed and the phase re-started so its lock re-resolved against the committed tree, which also cleared the two staleness findings legitimately and discarded the releases they lived in.'
      resolves: Q-3
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The rename goes first, and it is mechanical.** `model.ContextProfile` becomes
`ContextScope` and `model.Profile` becomes `Scope`, with the five readers at
`runner.go:527`, `gates.go:455`, `gates.go:492`, `gates.go:681` and the constant at
`model.go:146`. This intent's own artifact is renamed with `git mv` and P0 re-judged,
which is free while nothing is merged. It goes first because every name below is the
file's, and `docs/v2-delta.md` is deliberately untouched: it is a different document and
not in this change.

**The writer is `xeno scope set`, and it reads its entry whole.** The grammar is
`section set`'s, which is the command it most resembles — a person supplies content, the
runner supplies provenance. There is no `--phase`: `informationBase` reads the scope
from P0 and from nowhere else, so a flag would offer a choice the format does not have.
The entry arrives on stdin or through `--file`, decoded into `model.Scope` with
`KnownFields(true)`, which is `RecordQuestion`'s shape and for its reason: include
patterns, exclude patterns, a link per component and two budget numbers are a nested
structure however they are spelled, and the shape read is then the shape written. A typo
is refused rather than dropped, which is the standing rule about invented fields.

**The command writes the header, never the person.** `intent`, `phase`, `created`,
`schema_version`, `runner_version` and `plugin_version` come from `r.common`, as every
other artifact's do. A35's rule is the reason: a plausible value in a field nobody
produced is worse than an absent one, and these fields are provenance rather than
content.

**`include` is required.** A scope whose include is empty resolves to no files, and in
every one of the readers that is indistinguishable from having no scope at all — which
is the state this intent exists to end. A74's distinction does not rescue it here,
because the distinction it draws is between an empty list and an absent one, and both of
those arrive at the same empty `files` list in the lock.

**The resolution is extracted and shared.** `informationBase` currently resolves the
patterns inside the phase-start path. The walk, the first-match bucketing and the
hashing move to a function the writer and `phase start` both call, so the figure the
writer reports and the figure the lock records come from one place. Two resolutions
would be two answers to one question, and the budget is set from the first and judged
against the second.

**The writer reports the count and the byte total.** It is the criterion P1 took from
this intent's own P0, where the figure was established with `wc` by hand and the first
budget was wrong. The report is what the patterns select now, so the budget is a counted
figure, and the command says in the same breath that the budget cannot be revised once
P0 is judged.

**The refusal is the first statement in `Finish`, before the digest is written.** At
`runner.go:630`, where `phase == model.Phases[0]` and the scope file is absent, it
returns a refusal naming section 5 and the writer. Placed before `writeDigest` and
`writeCost` so that a refused finish writes nothing at all, which is what every other
refusal in the runner does. It is not in `phase start` of P1: a P0 can only be left
scopeless if it was allowed to finish, so the check belongs where the phase closes.

**Section 5's second limit is the loop bound.** `staleReads` loops `for i := 0; i <=
idx; i++`; it becomes `i < idx`. One character, and it is the whole of the fix that
unblocks this intent, because judging P1 then reads P0's lock alone. The specification's
reason goes in the comment, in its own words, so the next reader does not restore the
inclusive bound as a tidy-up.

**Section 5's first limit narrows to the commit range, through `git.Paths`.** `Ctx`
already carries `Base` and `Head`, and `git.Paths(root, base, head, "")` is the set of
paths the change under review touched. A lock entry outside that set is not reported.
The limit exists for noise, and the specification names the case: "every rebase onto a
moved default branch would set it off, and a check that fires constantly is one people
learn to ignore."

**Where there is no commit range, the staleness half reports nothing, and that is a
stated gap.** `git.Paths` treats an empty base or head as a caller error rather than a
range, because section 12 says the range is passed in and never inferred. So a local
`gate run` without `--base` and `--head` has no change under review and the half cannot
narrow to one; it reports nothing and the gate's first half is unaffected. This makes
the second half a CI check in practice. It is also a silent pass of exactly the kind
this intent was opened about, and it cannot be reported as anything else until there is
a result for a check that did not run — which is #235's mechanism. Recorded here rather
than papered over, and the code comment says it where a reader of the gate will find it.

<!-- xeno:section:alternatives -->
## Alternatives

**The writer taking flags instead of a document.** `--include`, repeated, with `--link
component=docs` and two budget flags. Rejected for `RecordQuestion`'s reason, written
down when the same choice came up for questions: a nested structure does not fit flags
without inventing a separator, and a separator invented here would be a second spelling
of a format `model.Scope` already fixes. `assumption record` and `decision record` do
take flags, and the difference is real — their entries are flat scalars.

**Keeping the name `profile` in the code and renaming only the file.** Cheaper by five
call sites, and rejected because it leaves the tree disagreeing with the document that
was just changed to settle the name. The standing rule puts the specification first; a
code identifier that contradicts it is the drift #202's audit was about.

**Calling the command `xeno context scope set`.** Three words, and no command in the
surface has three. `scope set` is unambiguous in a tool whose other nouns are intent,
phase, gate, section, evidence, assumption, question, decision, learning and obligation,
none of which is a near miss for it.

**Putting the refusal in `phase start` of P1 instead.** It would catch the same intents
one step later and leave a sealed P0 with no scope behind it, which is the state the
requirement exists to prevent. A phase that finished without its artifact is already
wrong by the time the next one starts.

**Putting the requirement in G-Schema after all, and releasing the 100.** This is Q-2's
third option, rejected there with its cost: 100 decisions by a second person, 100 open
obligations, and a red `gate verify` in between. Nothing in the design changed that
arithmetic, and it is recorded here because a later reader will reach for the gate
first, as #217 itself did.

**Inferring the commit range for limit 1 from the working tree, or from HEAD against the
default branch.** It would make the staleness half work locally instead of only in CI,
which is a real gain. Rejected because section 12 says the range is passed in and never
inferred, and `git.Paths` enforces that by treating an empty base or head as a caller
error. An inferred range would also make a verdict depend on which branch the runner
happened to be on, which is the class of defect A89 removed from the plugin root.

**Dropping limit 1 and implementing only limit 2.** Limit 2 alone unblocks this intent,
and it is one character. Rejected because the specification states both and gives limit
1 the reason that matters for an active repository: a check that fires on every rebase
onto a moved default branch is one people learn to ignore. Shipping half of a two-part
limit while the artifact it guards becomes mandatory is how the second half never
arrives.

**Fixing #235 here so a budget overrun and a skipped check could both be reported
without stopping work.** It is the mechanism three findings in this intent would use,
and it is genuinely wanted. Rejected as scope: a gate result the process definition does
not define, or an advisory finding inside one, is a change to what a verdict means,
which is WP1's and the document's. This intent names the three places that want it
instead.

**Making the lock re-resolvable so a staleness finding could be cleared by re-running
the phase.** That is #237, and it is the honest remedy for the finding's own wording.
Rejected here because #215 made the lock write-once deliberately — an overwritten lock
destroys the answer to "what changed" — so the fix is a decision about how to keep that
property while allowing a re-resolve, not a line in this change.

<!-- xeno:section:impact -->
## Impact

**Code.** `internal/model` for the two renamed identifiers; `internal/runner` for the
extracted resolution, the new `ScopeSet` and the refusal in `Finish`; `internal/gates`
for the two limits in `staleReads` and the three renamed readers; `cmd/xeno` for the
command, its entry in the table and two lines of help. One new file for the writer and
one for its test, which is the +2 the budget's headroom was left for.

**No new dependency.** The walk, the matching, the hashing, the YAML and `git.Paths` are
all in the tree. `git.Paths` already shells out to git, which the gate path already does
for `Commits`, so the staleness half gains no capability the gate did not have.

**The 100 sealed P0 phases are untouched, and this is the figure that proves it.** `gate
verify` is at exit 0 over 347 verdicts before the change and must be after it, with the
same count. The requirement lives in a command, so recomputing any sealed phase runs
exactly the checks it ran before.

**One existing verdict changes, and it is this intent's own.** P1 was released twice on
staleness findings raised by the inclusive loop bound. With limit 2 in place,
recomputing P1 raises neither, so those releases stop being load-bearing. P1's verdict
was since re-made green by re-resolving its lock, so the change removes a reason rather
than a record.

**The staleness half becomes a CI check in effect.** Locally there is no commit range,
so it reports nothing; in CI, where `--base` and `--head` are passed, it narrows to the
files the change touched. Before this change it ran everywhere and compared everything,
which is why it fired on this intent. The direction is less reporting, not more, and the
reporting it loses is the noise both limits were written to remove.

**Every new intent owes a scope at P0 from the refusal's commit onward.** That is D-1
taking effect. The first one after this is the measure of whether the writer is usable,
and it is the earliest point at which the budget figures mean anything, since this
intent's own were set before the writer existed.

**#235 is unchanged and now has three callers waiting on it.** A budget overrun still
turns a phase red against section 5's words; a scope-less P0 cannot be reported in a
verdict; and a staleness half with no commit range cannot say that it did not run. All
three want the same missing thing.

**No document changes.** Section 5 already puts the scope at P0, already says what it
carries, and already states both limits with their reasons. The specification commit
that renamed it is this intent's cause and is already made. Where the code and the
document disagreed, the document won in every case here.

**Tests.** The writer against a refused unknown field, a missing include and a reported
figure; the refusal at P0 with and without a scope; the extracted resolution giving one
answer to both callers; limit 2 by judging a phase that changed a file its own lock
lists and asserting green; limit 1 by a lock listing two files with one touched,
asserting one finding; and `gate verify` over this repository's own trail, which is
where the regression would show.
