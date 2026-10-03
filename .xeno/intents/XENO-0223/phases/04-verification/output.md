---
intent: github.com/triplem/xeno#171
phase: 04-verification
created: "2026-10-03T09:57:12Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2ffde2738cc6f48e5b66e923e8c95ba433cd2b00fa1bba1f0c8aedb9bc0173e2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | Checked by |
|---|---|
| A repository with no profile behaves as today | `TestChangedSinceIsSilentWithoutAProfile`, `TestNoProfileAndNoBudgetAreNoFinding`, and the whole trail still verifying |
| The lock carries `repo_commit`, absent where there is none | `TestTheLockRecordsTheCommitWhereThereIsOne`, and the demonstration, which recorded the real head |
| The lock carries `rules_applied` with path and version | `TestTheLockRecordsWhichRulesApplied` |
| `rules_applied` and `rules_hash` agree | the same test, comparing the list against `rules.Effective` |
| Absent rather than empty where no rule is in force | the same test's second half |
| A context over its file budget is a finding naming both numbers | `TestAContextOverItsFileBudgetIsAFinding` |
| A context over its byte budget is a finding naming both numbers | `TestAContextOverItsByteBudgetIsAFinding`, and the demonstration |
| Within budget, nothing; no budget declared, nothing | both budget tests' second halves |
| The finding comes from G-Schema and does not block | `TestTheBudgetFindingComesFromGSchema`, and the demonstration, where the phase was decidable and was decided |
| A declared link's document is in the base, with a hash | `TestADeclaredLinksDocumentIsInTheBase` |
| The `files` order follows the profile's `include` order | `TestTheBaseFollowsTheProfilesIncludeOrder`, reversing the profile; `TestGivenFilesAreComparedAgainstTheTree` |
| `phase start` names what changed, and is silent otherwise | `TestChangedSinceNamesWhatMovedAndNothingElse`, and the demonstration's printed report |
| Nothing is written that section 5 does not enumerate | read back: the lock gains two fields from its list, and the changed set is output |
| The mechanism is demonstrated against this repository | on a copy, below |
| Nothing else moves | the suite, `gofmt`, `go vet`, `./xeno gate verify` |

One criterion has no test and could not have one: that a link naming a document which does not
exist is a finding against the profile. It is not implemented — `informationBase` skips a link it
cannot hash — and that is in the gaps as a criterion this piece did not meet.

<!-- xeno:section:results -->
## Results

**The suite is green.** 428 cases pass, nothing fails, `gofmt -l` outside `vendor/` lists nothing,
`go vet ./...` is silent, `./xeno gate verify` is at exit 0 over 216 verdicts.

**The whole mechanism was run against this repository on a copy**, with a profile written for it:
`docs/**`, `ASSUMPTIONS.md`, `internal/runner/**`, `internal/gates/**`, excluding `**/testdata/**`,
one link from `internal/gates` to the process definition, and section 5's example budget of 30 files
and 400,000 bytes.

*The base resolved to 23 files, in the profile's order.* The four documents first, then
`ASSUMPTIONS.md`, then the runner, then the gates — stable before volatile, as written, and nothing
alphabetical about the whole list.

*`repo_commit` recorded `ea0cb1c3dfbd…`*, the actual head of the copy.

*`rules_applied` recorded all four shipped rules with `version: 1` each*, which is the first time a
lock in this project has said which rules applied rather than that some did.

*The byte budget fired immediately.* 709,921 bytes against the 400,000 of section 5's example, as a
G-Schema finding naming both numbers. The phase was red, decidable, and decided.

*Three edits produced three stale-read findings*, one per file of the base that moved —
`ASSUMPTIONS.md`, `internal/gates/gates.go`, `internal/runner/runner.go` — each naming the phase
that was given the file.

*After the findings were released, `phase start` on the next phase printed exactly those three
paths*, in the base's order, under one line naming the phase it was measured against. That is WP8's
first clause working end to end: a repeated phase is told what to reread and nothing else.

The copy was then deleted. Nothing in this repository's own trail carries a profile.

**The adoption figure the decision needs.** This intent changed six files outside the trail, and
**three of them are in that base**. So a profile of that shape costs roughly three stale-read
findings per intent of this kind, each needing a person's release, plus one budget finding until the
figure is raised. Every one of them is correct — the earlier phases did read those files, and the
files did change.

**`rules_applied` agrees with `rules_hash` by construction**, and the test asserts it rather than
trusting the shared call: one resolution, two writers.

**Two of section 5's four missing fields are now written**, and the two that remain have an owner
each: `plugin` with G-Supply, `tools` with WP15. A reader comparing the section against a lock now
finds two absences, both explained.

<!-- xeno:section:gaps -->
## Gaps

**A criterion this piece did not meet.** 01-requirements asked that a link naming a document which
does not exist be a finding against the profile. It is not: `informationBase` cannot hash a file
that is not there, so it skips the link silently, and the declaration a project most likely to
mistype is the one nothing reports. Writing it means a finding that belongs to the profile rather
than to the base, which is G-Schema's, and it did not get written. It is the only criterion of the
eleven that is unmet and it is unmet rather than deferred.

**The saving is gated behind the decision.** With a profile in force, editing the base turns the
earlier phase red, a red predecessor stops the next phase, and the report that says what to reread
is only reachable once the findings about those same files are released. Four approvals in the
demonstration. Both halves are specified and correct; their composition is worse than either.

**The byte budget moves after the phase is sealed.** The count is taken from the tree when the
check runs, because the lock carries hashes and no sizes, so a phase inside its budget today can be
over it next week and `gate verify` will say so about a sealed phase. A74 solved the same class of
problem for rules by reading the hash the artifact recorded; this field has no recorded number to
read, and giving it one means a field section 5 does not enumerate.

**Section 5's example budget is wrong for a repository of this size**, by a factor of nearly two.
Raised as a finding and not fixed, and it is the figure a project will copy first.

**The changed set is printed and therefore ephemeral.** A phase that ran yesterday leaves no record
of what it was told to reread, so the decision an agent made on that basis cannot be reconstructed.
That is the deliberate consequence of not inventing a field, and it means WP20 cannot measure the
saving from the trail — it would have to measure it live.

**Nothing checks that an agent read only what changed.** The report is advice. Section 5 says the
lock records the context that was declared and not everything that was read, and that treating it
as a measurement would be wrong, so the clause "a repeated phase reads only what changed" is now
implementable and still unverifiable. What was built is the knowledge; the behaviour is the agent's.

**The link's position in the base is a guess.** Links go last, on the argument that a link is the
most specific thing in a profile and therefore the most volatile. Nothing measured that, and the
assembly order exists for a cache whose behaviour nobody here has observed.

**This repository still has no profile**, so every mechanism in this piece is exercised by tests and
by one demonstration on a copy, and by nothing in the trail. That is deliberate and it is the same
position the rule set was in before #165: the first real use is the first outside reading.
