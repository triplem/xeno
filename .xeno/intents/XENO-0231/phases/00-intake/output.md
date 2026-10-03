---
intent: github.com/triplem/xeno#190
phase: 00-intake
created: "2026-10-03T14:01:06Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+74553ad.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b3fda305f67b232e62b645bdada281d943abd01d66119e731303dde41999d457
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

Two conventions this repository states and does not check, each of which failed in a way nothing
reported.

**The pull request title is the commit subject, and nothing read it.** `squash_merge_commit_title`
is `PR_TITLE` and squash is the only method on offer, so the subject written on a branch is
discarded at the merge and the title is the only place `CLAUDE.md`'s "commit subjects are
Conventional Commits" has any effect. Fifteen merges since `v0.28.0` carried prose titles,
semantic-release read sixteen commits and concluded "no release", and no release ran for five days.
The work in that window is WP4 complete, the rule engine, external gates, the agent layer, the
context lock and the symbol index. `xeno check commit-message` exists, ships the pattern, and is
invoked nowhere — it appears in the README, in `ASSUMPTIONS.md` and in two documents, and in no
workflow.

**A sealed intent can be deleted and nothing notices.** Section 7 states "what is sealed is never
rewritten" as a rule holding throughout. Measured on a copy of this repository: editing a sealed
`output.md` is caught, `gate verify` reporting `DIVERGENT` and exiting 1; deleting
`.xeno/intents/XENO-0225` entirely takes the count from 259 verdicts to 253 and exits **0**.
Nothing in the repository records how many verdicts there should be, so a verdict that is gone is
indistinguishable from one that never existed. That is a worse violation of the rule than an edit,
because an edit at least diverges.

**Both are the shape the last three intents have been about.** #186 was a check that could not run
before the merge; #187 was the drift it found; #190 is a convention that cannot survive the merge;
#193 is a rule with no reader. Each time something correct is stated in one place and nothing
compares the repository against it, and each time the thing that should have noticed could not.

**The second one has already cost this repository a phase.** XENO-0230 reached main with its
verification and review missing, because #189 was merged while its last commit was arriving. Every
gate was green. That is the same hole seen from the other side: the trail is not compared against
anything that knows what it should contain.

<!-- xeno:section:scope -->
## Scope

**In scope.** Two steps in the `xeno` workflow, so both land inside the one check that gates the
merge rather than beside it.

The title, fed to `xeno check commit-message --pattern conventional-commits` — the command that
already exists, the pattern that already ships, and the title that the ci-skip step already reads
into its environment for a different purpose.

The trail, compared against the base of the pull request, failing where any path under
`.xeno/intents/` is removed or renamed. A tree comparison and not a commit range, so a directory
created and dropped inside one branch is not reported. The checkout gains the history that
comparison needs.

And the two conventions written where a contributor reads them, in `CONTRIBUTING.md`, beside the
squash setting the footer already depends on.

**Out of scope, and each for its own reason.**

Modification of a sealed artifact. `gate verify` already catches it and reports which phase
diverged, which is better than a path list.

Blocking a re-seal — editing an artifact and running `phase finish` again. The working sequence
provides for re-judging deliberately, through the `changed-after-verdict` state and the suggestion
that names it, so a guard that refused modification would refuse a designed act. It is left as the
remaining hole, stated in the verification phase.

The fifteen releases that did not happen. Their types are unknowable from prose subjects and
nothing can reconstruct what they would have cut. #190 leaves that as a maintainer's decision and
this intent does not take it.

G-Complete not running before P5, which is how XENO-0230 reached main with two phases missing. It is
a real finding, it is about the gate set rather than about CI, and it wants its own issue.

A commit-msg hook for the subject. Section 7: a hook's result is advisory because its configuration
lives on a developer machine, and the subject that matters is not written on the machine at all.

<!-- xeno:section:context-rationale -->
## Why this context

Section 7 is read for the sentence this intent enforces — "what is sealed is never rewritten" — and
for the paragraph around it, which says what the rule permits instead: a separate file that
references a phase, or a comparison that is computed. A deletion is neither.

Section 7's hooks subsection is read again, for the second time in three intents, because the
question of whether a hook could carry either check has a standing answer: advisory by
construction, and no exclusive logic.

Appendix B is read for the path being part of `artifacts_hash`, which is why a rename counts as a
rewrite and is refused beside a deletion.

A71 is read for what the second shipped pattern is for: a Conventional Commits subject ending in a
parenthesised reference, which is exactly the form the host produces from a title. It settles which
of the two patterns the title is checked against, and the answer is the other one.

`.github/workflows/xeno.yml` is read in full. It already reads
`github.event.pull_request.title` into a step's environment to scan it for ci-skip markers, with
the comment that the title and the body come through the environment rather than by interpolation
because both are text somebody else can write. That decision is reused rather than retaken, and the
job is where both steps belong because its name is the one required context.

`scripts/github-settings.sh` is read for `squash_merge_commit_title=PR_TITLE`, which is the fact the
whole title half rests on, and for the merge settings group that was already the record of it.

`CONTRIBUTING.md` is read for where the squash setting is already explained, because the title half
of that setting belongs in the same place and was never written there.

The release workflow's run log on `eba5afc` is read as the evidence that a conventional title is
sufficient: the first conventional subject in sixteen commits cut `v0.29.0` within two minutes.

Nothing outside the repository is needed. Both checks were run against real cases before anything
was written, which the verification phase records.
