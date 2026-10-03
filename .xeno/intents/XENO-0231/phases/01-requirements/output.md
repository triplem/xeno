---
intent: github.com/triplem/xeno#190
phase: 01-requirements
created: "2026-10-03T14:02:09Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+74553ad.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ff172d799c19141e606ff9af60a47100bfdc96234f102a61cc91733f95d175f3
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**A pull request whose title is not a Conventional Commit fails the required check**, and the
failure says that the title becomes the subject and that a prose one releases nothing silently.

**A title with the host's reference on it passes.** `feat(x): a thing (#3)` is what a renamed or
re-opened pull request can carry, and the plain pattern accepts it; the author still writes it
without one.

**The check uses the command that ships.** `xeno check commit-message --pattern
conventional-commits`, not an expression written into a workflow, so the rule has one
implementation and `CLAUDE.md`'s subject rule and this check cannot drift apart.

**A pull request that removes or renames any path under `.xeno/intents/` fails**, naming the paths.

**A pull request that only adds passes**, which is every ordinary intent.

**A directory created and dropped inside one branch passes.** It never reached main. The comparison
is a tree diff against the base and not a walk of the commits.

**A base commit missing from the clone fails rather than passes.** A comparison against something
that is not there reports the same thing as a clean one, and that is the failure this whole
sequence of intents has been about.

**Both steps are in the `xeno` job.** Its name is the required context; a check beside it is
advisory and #186 is what that costs.

**Both are skipped on a push to main**, where there is no pull request to read. The title has
already become a subject by then and the trail comparison has no base.

**Neither touches the gate path.** No Go file, no gate, no rule, no artifact, so `gate verify` is
unchanged and nothing behind this recomputes differently.

**The conventions are written in `CONTRIBUTING.md`** — the title beside the squash setting the
footer already depends on, and the append-only rule beside the intent it belongs to.

**The suite, `gofmt`, `go vet` and `gate verify` stay green.**

<!-- xeno:section:non-goals -->
## Non goals

**No new command, pattern, gate or rule.** The title check calls what ships. The trail check calls
git. Both are steps in a workflow, which is the same category as the four checks already there.

**No blocking of modification under `.xeno/intents/`.** `gate verify` catches it and names the
phase, which is better than a path list, and refusing it would refuse a re-judgement.

**No attempt to close the re-seal hole.** Editing an artifact and running `phase finish` again
writes a verdict over the new content. The working sequence provides for that deliberately, and
G-Freshness turns the successor red, so hiding one edit means re-judging an intent in a diff a
reviewer is looking at. Stated rather than closed.

**No decision about the fifteen missed releases.** #190 leaves it open on purpose.

**No change to G-Complete.** That it runs only at P5, so an intent which never reaches P5 is never
checked for completeness, is what let XENO-0230 reach main with two phases missing. It is a finding
about the gate set and wants its own issue.

**No hook.** Advisory by construction, and the subject that matters is not written on a developer's
machine at all.

**No `required_pipeline` naming the set.** Still an Appendix A addition, still #186's open clause.

<!-- xeno:section:constraints -->
## Constraints

**A check that is not required is advisory.** #186's whole content. Both steps go in the job whose
name is the required context, and not into a workflow of their own, however much tidier that would
read.

**A vacuous check reports what a clean one reports.** So the trail step refuses a base it cannot
find rather than comparing against nothing. This is the same error as a monitor reading an earlier
commit's results, made three times in one day in this repository, and it is cheap to make again.

**The tree, not the commits.** Section 7's rule is about what is sealed, and what is sealed is what
reached main. A commit range would refuse a branch that corrected itself.

**The path is part of the hash.** Appendix B, so a rename is a rewrite and belongs with the
deletion rather than with the modifications.

**One implementation per rule.** The subject rule is `CLAUDE.md`'s and the pattern is the plugin's;
a regular expression in a workflow would be a second statement of it that nothing compares.

**Text somebody else can write goes through the environment.** The neighbouring step says so about
the title and the body, and the reason does not change for a second reader of the same string.

**The plain pattern, not the one with the issue.** A71 defines the second for the form the host
produces, which is the subject on main and not the title a person types.

**88 columns, SPDX where it applies, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0.**

**One intent, one branch, two issues** — `190-the-conventions-become-checks`, closing #190 and
#193, labelled wp0 and wp1. One branch based on main and not on another branch, which is the rule
that came out of this morning.
