---
intent: github.com/triplem/xeno#330
phase: 03-implementation
created: "2026-10-08T17:53:10Z"
schema_version: "1.0"
runner_version: dev+8fb365d.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 87ecd0afdd234e50cf67bd24504c044f9772f406075b4f834a79429bae5e4c0b
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

Fifteen files, by concern.

**The answer grows: `internal/model/identity.go`.** `Issue` gains `Labels`, `Milestone`,
`Comments` and `OpenMilestones`, with `Milestone` and `Comment` as the two shapes both
hosts share and `ApprovedLabel` and `ApprovedWord` as the constants section 12 fixes.
`Issue.Approval()` returns who, when and the reason off the last comment whose first line
is the word, trimmed, case ignored and trailing `.`, `:` or `-` dropped, or the list of
what is missing in the words a refusal prints; `Issue.Ahead()` returns the open milestone
whose turn comes first where that is not the issue's own, ordering by due date, undated
last, then by number, and holding nothing for a closed milestone. Six table-driven cases
on each in `identity_test.go`.

**Both adapters make the calls: `internal/host/github/tracker.go` and `gitlab/tracker.go`.**
`Issue` reads the labels and the milestone off the issue, then the comments — GitHub at
`/issues/N/comments`, GitLab at `/issues/N/notes?sort=asc&order_by=created_at` with
`system: true` notes left out — following pages of a hundred until a short one, ten at
most; and, only where the issue carries a milestone, the project's open milestones,
`state=open` on one host and `state=active` on the other. Each host's milestone shape is
a private struct with its own field names, `due_on` and `number` against `due_date` and
`iid`, converted to the model's. The port, `host.Issues`, is unchanged. Tests on each
host assert the paths asked in order, the fields read, that an issue without a milestone
costs two calls, and on GitHub that a second page is followed.

**One read for both commands: `internal/runner/tracker.go`.** `readIssue(key)` is now a
thin caller of `issueOf(tracker, qualified)`, which holds the body that was there and
takes the id rather than the key, so `intent start` can call it for an intent that does
not exist yet. `intake()` writes the sentence above the quote: `Approved by @who on
stamp: reason`, or `with no reason beside the word`, or `No approval was found on the
issue when this phase started: it is missing …`; and `Started ahead of its milestone,
"1.1", while "1.0" is open.` where `Ahead()` says so.

**The refusal: `internal/runner/runner.go`.** `IntentStart` takes a third argument,
`now`, and after the cheap refusals — no issue named, no key, a directory that exists —
calls `approved()` where a tracker block is there. `approved` reads the issue and refuses
in three shapes: a read that did not happen, with the sentence `issueOf` gave and the
host's reason folded in — the case the first run of the tests caught, since a 404 arrives
on the issue and not beside it; an issue missing one or both halves; and a milestone
ahead, unless `now`. The literal wordings are the ones the design wrote down, with the
read-failure one reordered so that it starts with what did not happen.

**The flag: `cmd/xeno/main.go`.** `--now` as a boolean beside `--for`, passed through, and
`[--now]` on the usage line; `docs/commands.md` carries the same line and a paragraph on
what is refused, what the flag does and what the intake records.

**The tests that had to move.** Every test that started an intent against a tracker block
with no host behind it — five in `runner_test.go`, one in `main_test.go`, the MCP
fixture — now points the block at a fake host answering an approved issue, because a
block with nothing behind it is a refusal from here and a block naming `api.github.com`
would have reached the network from a test. `fakeHost` in `tracker_test.go` learned the
labels, the milestone, the comments and the milestones; `approved()` builds the common
case; `hostOf()` derives the id's host from a fake server's address. The MCP fixture
drops the token again after the start, so that its two routes keep leaving identical
artifacts. Six new tests in `tracker_test.go` cover the refusals, `--now`, the
milestone cases, the reads that cannot happen, and the two sentences.

**The register.** A103 in `docs/assumptions.md`.

Against the real host, before the artifacts were written: `xeno intent start --for 330`
with a token refuses *missing a comment whose first line is approved*, without one *no
token in the environment*, and `--for 99999` *there is no issue 99999 in triplem/xeno*.
None wrote a directory.

<!-- xeno:section:deviations -->
## Deviations from the design

Two, both named against the design's `decisions`.

**The read-failure refusal is worded the other way round.** The design wrote *the sentence
`readIssue` already produces with `so the intent is not started` after it*. Against the
real host that read as a sentence with two tails — *no issue was read: no token …; set it
…, so the intent is not started: … nothing was read* — so the refusal now begins with what
did not happen and carries the read's sentence after it: *the intent is not started,
because starting one reads the issue's approval and no issue was read: …*. Same content,
one tail.

**A host's refusal is folded into the same shape.** The design's `issueOf` has three
answers, and the adapters put a 404 or 403 on the issue's `Reason` rather than in the
middle answer, which `phase start` relied on and the design did not mention. The first
run of `TestAReadThatCannotHappenDoesNotStartTheIntent` caught it: a 404 was judged as an
issue missing both halves. `approved` now treats `Reason` as the note it is. Nothing
else departs; the port, the constants, the pure functions, the recomputed sentence, the
paging bound and the flag are as decided.
