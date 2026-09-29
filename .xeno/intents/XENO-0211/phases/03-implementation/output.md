---
intent: github.com/triplem/xeno#141
phase: 03-implementation
created: "2026-09-29T19:55:01Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6aa6f9b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6795b2fbff79a47ff0955f45092a59741c041cdfb5b4d4cc283127e612421c0e
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

One row, `A65`, appended to the assumption table of `ASSUMPTIONS.md`, immediately before the
`## Decisions` heading. Two lines changed: the row, and the blank line that separates the
table from that heading.

The row carries the five things the acceptance asks for. The reading taken, in its first
clause, so that a reader who stops there has the decision. The deciding evidence,
`allow_bypass` answered by one boolean on GitHub and by three lists of grants on GitLab, with
`approvals.not_by_author` as the second case, where the states a requirement can reach differ
by host. The rejected reading named as such, per-field expressibility generalising
`ReviewsExpressible`, with why expressibility does not help: the field is expressible and the
difficulty is where the judgement about access levels lives. What made it decidable without
writing GitLab, the unauthenticated protected branch endpoint, which is the condition #97 set.
And the cost, that two adapters can drift in their words for one requirement name and that
this undoes the `evaluate` half of #99's table in the release that built it.

The `Where` column names `internal/enforcement`, #97 and #104. The state column says
`approved` and then bounds what was approved: the decision only, by #104's fourth criterion,
with the port left to the first piece, the existing `enforcement_test.go` cases named as what
pins the GitHub phrasings across that move, and the two unverified rows of #104's mapping
named with the reason each is unverified.

No Go file changed. No test changed. No artifact gained a field. The row is placed as `A65`
because `A64` was the highest, and the table is read in order rather than sorted.

<!-- xeno:section:deviations -->
## Deviations from the design

**One, and it is about the timing rather than the content.** #97 said the vocabulary "is
decided when GitLab is written", and no GitLab is written here. The letter of that schedule is
not met.

What is met is the reason it gave. #97's ground for waiting was that choosing in the abstract
is how a neutral interface acquires fields nobody needs, so what it wanted was a second host's
answer in front of us and not the adapter that consumes it. gitlab.com serves
`protected_branches` to an unauthenticated request, so the answer was available for the asking
and the deferral was satisfiable a release earlier than it read. Deciding it now is what lets
#104's first and third pieces be drawn against a settled signature instead of a provisional
one, which is the cost the letter of the schedule would have charged.

The deviation is recorded here rather than silently, and the row itself says what made it
decidable, so a reader can judge the substitution rather than discover it.

**Nothing else deviates.** The acceptance's seventh statement holds as written: only
`ASSUMPTIONS.md` and this intent's own phase directories changed. The first standing rule was
not touched, since `ASSUMPTIONS.md` is not one of the documents it protects and `CLAUDE.md`
names it as where a decision taken while building the core goes. The second added no field,
gate, tool or rule.
