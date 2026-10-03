---
intent: github.com/triplem/xeno#190
phase: 04-verification
created: "2026-10-03T14:05:33Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+74553ad.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c0719b7b288bb990a4f900d8b293d2f89b1bed15ed15645b017a717475c97a86
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

Both checks are workflow steps, so most of these are held by a run or by a command somebody can
repeat, not by a test.

**A prose title fails** — `printf '%s' "The audit gate binds" | ./xeno check commit-message
--pattern conventional-commits` exits non-zero. Run before the step was written. This pull request
is the live case: its title is a Conventional Commit, and the step's output names the version it
accepted.

**A title with the host's reference passes** — the same command on `feat(runner): a thing (#191)`,
which is one of six titles checked against both shipped patterns before the pattern was chosen. The
table in P0's context is that measurement.

**The check uses the shipped command** — read in the diff. `grep -c 'check commit-message'
.github/workflows/` was 0 before this intent and is 1 after.

**A deletion fails, naming the paths** — in a throwaway worktree off main, `rm -rf
.xeno/intents/XENO-0225` committed, then the step's own invocation: 31 paths reported.

**An addition passes** — the same worktree, one new intent directory committed: 0 paths.

**A directory created and dropped inside the branch passes** — the same worktree, two commits, add
then remove: 0 paths. This is the case the two-dot form exists for and the one a commit range would
have failed.

**A missing base fails rather than passes** — held by construction, `git cat-file -e` under `set
-eu` with its own message. Not run against a real missing base, because arranging one means a
pull request whose base has been garbage collected; the branch is three lines and the alternative is
the vacuous pass this intent is about.

**Both steps are in the `xeno` job** — read in the diff, and confirmed by the check list on this
pull request, where the five contexts are unchanged: nothing new appears, which is the point.

**Both skipped on a push to main** — `if: github.event_name == 'pull_request'`, and confirmed by
main's own run after this merges carrying neither step's output.

**Nothing touches the gate path** — `git diff --stat` names two files, neither of them Go, and
`gate verify` reports 265 verdicts at exit 0.

**The conventions are in `CONTRIBUTING.md`** — read rather than tested.

<!-- xeno:section:results -->
## Results

**Before anything was written, in a throwaway worktree off `main`:**

| case | removed or renamed paths | step |
|---|---|---|
| `rm -rf .xeno/intents/XENO-0225`, committed | **31** | fails, correctly |
| one new intent directory, committed | 0 | passes |
| a directory added, then removed, two commits | 0 | passes, correctly |

**And the title, against both shipped patterns**, six titles including `feat(runner): a thing`,
`feat(runner): a thing (#191)`, `chore(deps)!: breaking` and `The audit gate binds`. The plain
pattern accepts the first three and refuses the last; `conventional-commits-with-issue` accepts only
the one carrying the reference, which is A71's definition and the reason the plain one was chosen.

**`go build`, `gofmt -l .` outside `vendor/`, `go vet ./...`** — all clean, and no Go file is
touched.

**`go test ./...`** — eighteen packages ok.

**`./xeno gate verify`** — `verified 265 verdicts`, exit 0.

**This pull request is the first live case of both steps.** Its title is a Conventional Commit and
every path it adds under `.xeno/intents/` is an addition, so both are expected to pass — which is
the weaker half of the evidence. The strong half is the worktree table above, where each step was
shown to fail on the case it exists for.

**The evidence this intent cannot produce.** Neither step can be shown failing on this pull request
without making the pull request wrong, and a deliberately broken one would be a commit whose purpose
is to be refused. The worktree is the substitute and it is named here as such.

<!-- xeno:section:gaps -->
## Gaps

**Neither step is shown failing in CI.** They are shown failing in a worktree, by the same
invocation, and shown passing in CI. Closing the gap means a pull request that is deliberately wrong,
which is a commit whose purpose is to be refused; the next prose title somebody writes by accident is
the real first test, and it will be a contributor meeting an error message rather than an experiment.

**The re-seal hole is open, and is now the only way to rewrite the trail quietly.** Edit a sealed
artifact and run `phase finish` again: the verdict is written over the new content and `verify`
passes for that phase. G-Freshness turns the successor red, so hiding one edit means re-judging the
rest of the intent, and the diff that does it is in front of a reviewer. Deliberately not closed,
because refusing modification would refuse a re-judgement the working sequence provides for.

**G-Complete still runs only at P5.** An intent that never reaches P5 is never checked for
completeness, which is how XENO-0230 reached main without its verification or its review while every
gate was green. Nothing in this intent touches it and it wants its own issue — the third time today
that the thing which should have noticed could not.

**The fifteen releases remain unreleased and undescribed.** `v0.29.0` covers them by version and
names one commit, because the other fifteen subjects cannot be parsed. #190 leaves the decision open
and this intent does not take it.

**A title can be changed after the check passes.** The step runs on `pull_request`, which fires on
`edited` only if the workflow asks for that trigger type, and this one does not. A title edited
between the last check and the merge is not re-checked. GitHub re-runs required checks on a new
commit and not on a retitle, so the window is real; it is narrow, it needs somebody to edit a title
and merge without pushing, and closing it means adding `types: [opened, synchronize, reopened,
edited]` — which would also re-run the whole job on every description edit.
