---
intent: github.com/triplem/xeno#336
phase: 03-implementation
created: "2026-10-10T13:33:02Z"
schema_version: "1.0"
runner_version: dev+5f645cb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e487a47233690e0cc961143851c9d1e0e896128554446538de4022ad3f03d7cf
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

One file, two hunks, 23 lines added and none removed. `git diff -U0 main` confirms the
hunks are at `4.1` and at `Sources` and nowhere else.

**Section 4.1, after the properties list.** The paragraph opens on the repository and the
claim that decides how everything after it is read — "is this choice running rather than an
alternative to it" — then says what the design is, where the receiver lives, and that it is
evidence for section 12's paragraph rather than against it. The clause that carries the
weight is the one taken from the tree rather than from the issue: the receiver is a hosted
automation "and not anything in the repository — whose tree carries no workflows at all."
That turns "something has to receive them" from a figure of speech into an observation about
a specific tree.

Then the #330 contrast, as a second sentence: `xeno-approved` is a precondition
`xeno intent start` reads off an issue it has already fetched, not an event anything listens
for. The wording is held to what `Issue.Approval()` actually does, which is why "reads off an
issue it has already fetched" and not "reads the issue" — the function is handed the issue
and does no fetching of its own.

**Sources, a new final paragraph.** In the form the OpenSpec entry above it uses: what it was
read for, the date, the repository, the reference, what was read, and the facts about the
repository itself. Three things in it are deliberate. The pin is a commit and the entry says
why — "because the repository publishes no releases" — so the departure from the tag above is
answered where a reader meets it. The paths are named by what each was read *for*, so the
entry is a record of the reading and not a file list. And the absence of `.github/workflows`
is recorded as a thing read, with the reason it matters, because an absence nobody wrote down
is an absence the next reader has to re-establish.

The four status labels, the issue and pull request templates and `openspec/` appear in the
Sources entry as "the rest of the design the sentence does not take". That is the line
between the two places: the sentence names the mechanism and the receiver, the entry records
everything the reading covered, and a reader who wants the rest is one click away rather than
one list longer.

**What was checked rather than asserted.** `git diff main -- docs/process-definition.md` is
**0 lines**, which is criterion 7. The two hunk headers are criterion 8. `awk 'length>88'`
over the file returns nothing outside tables, which is half of criterion 10; the other half,
that the page builds, is the docs workflow's answer and not this machine's, since zensical is
installed by the workflow and not here.

<!-- xeno:section:deviations -->
## Deviations from the design

One, and it is against my own criterion rather than against the issue.

**Criterion 9 says one sentence and the paragraph is two.** The criterion was written from
the issue's "as one sentence", and the constraints section had already named the collision:
four claims in one sentence is long, and it said the criterion wins and the sentence takes a
subordinate clause. Written out, one sentence carrying all four came to about ninety words
with three subordinate clauses nested inside each other, and the #330 contrast — the clause
whose whole job is a distinction a reader can get wrong — arrived last, after the reader had
already been asked to hold two others. The pointer being load-bearing is the reason the
design gave for not quoting section 12; the same reasoning says the contrast has to be
legible, and it was not.

So the four claims are in two sentences: three in the first, the #330 contrast in the second.

What that preserves is what the issue's "one sentence" was for. The distinction it draws is
against a section — section 9's treatment is what a candidate weighed in earnest gets, and
#334 is to give AI-DLC the same, so a third such treatment here would miscategorise the demo.
Criterion 9's own wording says that: "It is one sentence, not a section... it may not become a
table or a list." A two-sentence paragraph is none of those. What it costs is that the
criterion as written is not met, and saying it is met would be the kind of small untruth the
deviations section exists to catch.

Nothing else departs from the design. The placement is after the properties list, as the
alternatives section decided and for the reason it gave; the Sources entry is the form that
section's second paragraph already uses; and `internal/model/identity.go` was read and not
edited, which is the distinction the intake drew between what changes and what the change is
read off.
