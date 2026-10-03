---
intent: github.com/triplem/xeno#190
phase: 02-design
created: "2026-10-03T14:02:55Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+74553ad.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: adb7d3cc26947b41663591761a817795bee68537284fdfa03ebfae8d07bd0fae
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

**Both steps go in `xeno.yml`'s `verify` job.** Its id is the context `main` requires, so a step
there binds and a workflow beside it would not. #186 is the cost of the other choice, measured at
six merges.

**The title step runs after `build`**, because it uses the binary that step produces. Calling the
command rather than writing an expression is the point: `CLAUDE.md` states the subject rule, the
plugin ships the pattern, and a regular expression in a workflow would be a third statement of one
rule with nothing comparing the three.

**The trail step runs before `format`**, early, because it is the one check here whose subject is a
destructive mistake and the feedback is worth having before a minute of Go.

**The trail comparison is `git diff --diff-filter=DR --name-only "$BASE" HEAD`.** Two-dot, so it
compares trees: a directory created and dropped inside one branch never reached main and is not
reported. `HEAD` is the merge result on a `pull_request` event, which is the tree the merge would
produce and therefore the right subject rather than the branch tip.

**`D` and `R`, not `M`.** `gate verify` already catches modification and names the phase, which is
more use than a path. Rename is included because Appendix B puts the path inside `artifacts_hash`.

**A base that is not in the clone is a failure.** Stated as its own branch with its own message,
because the alternative is a step that compares against nothing and prints the same line a clean
run prints.

**The checkout gains `fetch-depth: 0`.** The base commit has to be present. A targeted fetch of one
commit would be shorter and would need the step to know how the action configured the remote; the
whole history of this repository is small and the cost is seconds.

**The title and the base come through `env`.** The neighbouring step already does this for the title
with the reason written down — both are text somebody else can write — and the reason does not
change for a second reader of the same string.

**Both steps are `if: github.event_name == 'pull_request'`.** On a push to main the title has
already become a subject and there is no base to compare against.

**The two conventions go into `CONTRIBUTING.md`** rather than into the workflow comments alone: a
contributor reads that file and does not read a workflow, and the squash setting the footer depends
on is already explained there.

<!-- xeno:section:alternatives -->
## Alternatives

**A regular expression in the workflow for the title.** Refused. It would be a third statement of
one rule, beside `CLAUDE.md`'s prose and the plugin's pattern, with nothing comparing them. The
command exists, ships, and was invoked nowhere, which is the defect rather than a reason to write
around it.

**`conventional-commits-with-issue` for the title.** Refused on A71, which defines that pattern for
a subject ending in a parenthesised reference — the form the host produces at the merge, not the
one a person types. Requiring it would ask for a suffix nobody writes.

**A separate workflow for the two checks.** Refused: a new workflow is a new context, and a context
that is not in the protection rule is advisory. It would also have to be added to
`scripts/github-settings.sh` and to the protection, which is three places for a tidiness gain.

**A commit range for the trail, `$BASE..HEAD`.** Refused. It reports a directory that was created
and dropped inside the branch, which never reached main and which section 7's rule does not reach
either. The rule is about what is sealed.

**Refusing modification under `.xeno/intents/` as well.** Refused twice over: `gate verify` already
reports it, better, and refusing it would refuse a re-judgement, which the `changed-after-verdict`
state and the suggestion that names it exist to permit.

**A hash over the set of intents, recorded in the repository.** It would catch a deletion without
git, which is the only way a gate could — and it is a new field, so a specification change, and a
field whose only purpose is to make deletion detectable where the host already knows.

**A commit-msg hook for the subject.** Refused on section 7's two constraints, and more simply: the
subject that reaches main is not written on a developer's machine, so there is no commit for a hook
to see.

**Letting the base default to `origin/main`.** Refused. A pull request's base is whatever it says it
is, and this morning's stacked branches are exactly the case where assuming main would have
compared against the wrong tree.

<!-- xeno:section:impact -->
## Impact

**`.github/workflows/xeno.yml`.** `fetch-depth: 0` on the checkout, with the reason. A step "the
trail is not rewritten" after the ci-skip step, carrying the measurement that produced it — one
intent removed, six fewer verdicts, exit 0 — and refusing a base it cannot find. A step "the title
is the commit subject, so it is a Conventional Commit" after `build`, carrying why it is the plain
pattern and what a prose title costs. Both guarded by `github.event_name == 'pull_request'`.

**`CONTRIBUTING.md`.** Under "Every change is an intent": the trail only grows, what catches an edit,
what caught nothing until #193, and that a branch which corrects itself is not refused. Under
"Commit messages": the title is the subject, the host appends the reference, which pattern is which,
and that a prose title releases nothing silently with the five days as the evidence.

**No Go file.** No gate, no rule, no pattern, no artifact schema. `gate verify` is unchanged and no
verdict behind this intent recomputes differently.

**No document change.** Section 7 already states the rule this enforces, and `CLAUDE.md` already
states the other.

**No change to `scripts/github-settings.sh`.** The contexts it records do not move: both steps are
inside `verify`, which is why they were put there.

**What a reader should expect to break.** A pull request with a prose title now fails. That is the
point, and it is the first thing a contributor will meet; the message says what to do and why, and
`CONTRIBUTING.md` says it at greater length.
