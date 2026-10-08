---
intent: github.com/triplem/xeno#330
phase: 02-design
created: "2026-10-08T17:40:02Z"
schema_version: "1.0"
runner_version: dev+e22a533.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3a3176d4f108faaceb078adb617326dc62b9efeafa3460f402b9ba97a60ce9d5
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

**The answer grows, the port does not.** `model.Issue` gains `Labels`, `Milestone`,
`Comments` and `OpenMilestones`, and `host.Issues` keeps its one method. Section 12 says
reading an issue stays one operation and is up to three calls, so the adapter makes the
calls and the port's signature is untouched. The new fields are the host's own words only
where the two hosts share them — a label is a name on both, a milestone a title and a due
date, a comment an author, a time and a body — which is the reading A65 fixed: the adapter
answers in the host's vocabulary and the domain judges. `Milestone` is a pointer, because
an issue with none and an issue whose milestone the adapter did not read are different
answers, and `OpenMilestones` is read only where the issue carries one.

**The judgement lives in `model`, beside `Qualified` and `Locate`.** `Issue.Approval()`
returns who, when and the reason, or the list of what is missing; `Issue.Ahead()` returns
the open milestone ahead of the issue's, or nil. Both are pure functions over the answer,
so a test needs no host and the two adapters cannot disagree about what approval means.
The comment that counts is the last one whose first line, trimmed, is `approved` with case
ignored; the reason is the rest of the comment. The milestone ahead is the first of the
open milestones ordered by due date, undated ones last, then by number, where that is not
the issue's own; an issue whose milestone is closed is held by nothing, since its turn has
passed.

**One read for both commands.** `readIssue` in `tracker.go` becomes `issueOf(qualified)`,
taking the id rather than the key, and `IntentStart` and `StartPhase` both call it. It
keeps its three answers — the issue, a sentence saying why there is none, an error — and
the two callers read the middle one differently: `phase start` prints it and starts the
phase, as Appendix A says; `intent start` turns it into a refusal, because a start that
could not read the approval has nothing to stand on. That includes the missing token, which
`phase start` has always carried on without: a clone without a token is the ordinary state
of somebody reading the trail, and starting an intent is not reading.

**`--now` is an argument of `IntentStart`.** A fourth parameter rather than a field on the
runner, because it belongs to one call and not to the session; the CLI reads it as a
boolean beside `--for`. It is consulted only where a milestone is ahead, so passing it on
an issue with no milestone changes nothing and is not refused.

**The sentence is recomputed at P0, not carried.** `intent start` writes `intent.yaml`,
whose fields section 5 enumerates, and there is no place in it for who approved or for the
flag. So `phase start` at P0 reads the issue again — it does already — and writes what it
finds at that moment above the quote: `Approved by @who on 2026-10-08T15:10:28Z: reason`,
and `Started ahead of milestone 1.1 while 1.0 is open` where that still holds. What this
loses is the case where the label was removed, or the milestone moved, between the two
commands, and the intake then says what is true at P0 rather than what was true at the
start; that is the honest sentence, since the intake is the record and the start wrote
none. Where no approval is found at P0, the sentence says so.

**The refusals are literal and name the thing.** `issue 330 in triplem/xeno does not carry
the label approved`, `... has no comment whose first line is approved; the label is there`,
`issue 330 in triplem/xeno is for milestone 1.1, and 1.0 is open ahead of it; --now starts
it anyway`, and for the reads that cannot happen, the sentence `readIssue` already produces
with `so the intent is not started` after it.

**Comments are paged.** Both hosts page at 100, and an approval is usually the last
comment, so the adapter follows pages until a short one, up to ten. Ten is a bound rather
than a limit anybody expects to meet, and an issue with a thousand comments is a different
problem.

<!-- xeno:section:alternatives -->
## Alternatives

**A fifth port operation, `Approval(project, key)`.** Cleaner to test and declined because
section 12 and WP12 both fix the contract at four and say a second host never changes it;
the clause says in words that reading an issue is the operation and grows. The cost of the
reading taken is that `phase start` now makes two or three calls where it made one.

**A `tracker.approval` block in `project.yaml` naming the label and the word.** Declined:
Appendix A would gain a field, which is a specification change, and the clause fixes both
names so that a reader of any trail knows what `approved` meant without reading a
configuration that may have changed since.

**Carrying who approved into `intent.yaml`.** The natural place and a forbidden one:
section 5 enumerates the file's fields and this intent adds none. The sentence in the
intake is what the clause names, and the cost of recomputing it is in the decision above.

**Judging the comment's author.** GitHub's `author_association` would let the adapter
require a maintainer; GitLab's notes carry nothing like it and a per-author lookup is a
fourth call in the host's own access vocabulary, which is the judgement A65 keeps out of
the domain. The label's right is the standing check instead, by the decision on the issue.

**A gate reading a copy of the host's answer.** Rejected in the issue and in the clause: a
gate never reaches the host, and a copy is what the agent wrote about itself.

**Refusing on `--now` where no milestone is ahead.** Would catch a flag passed by habit;
declined because a refusal for a flag that changed nothing teaches nothing, and the intake
says whether the intent was started ahead regardless of the flag.

<!-- xeno:section:impact -->
## Impact

**For a project with a tracker block.** `xeno intent start` now needs the token and the
issue's approval, and refuses otherwise; `xeno phase start` at P0 makes up to three calls
where it made one. A project without a block sees nothing.

**For the tests.** Every runner and CLI test that starts an intent against the fixture's
tracker block gets a fake host that answers an approved issue, since a block with no host
behind it is now a refusal; the fake host in `tracker_test.go` learns the comments and
milestones paths. Seven tests move, none weakens.

**For the trail.** Nothing sealed changes. Intakes written from here carry one or two
sentences above the quote, inside the problem section that already exists, so no template
and no gate moves; `intake@1.0.0` stays. This intent's own intake was written before the
code and records the authorisation by hand, which its problem section says.

**For a reader.** `docs/commands.md` gains `[--now]` and a paragraph on what `intent start`
refuses; the README's one-line sequence is unchanged, since the flag is the exception.
`docs/assumptions.md` gains A103.

**For 1.1.** A third host implements the same one operation with its own three calls. Jira
has labels and versions rather than milestones, and whether a version is "ahead" is the
question 1.1 answers; the domain's ordering by due date is where that would change.
