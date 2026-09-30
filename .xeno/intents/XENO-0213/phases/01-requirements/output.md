---
intent: github.com/triplem/xeno#145
phase: 01-requirements
created: "2026-09-30T05:44:35Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4e41c6b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 57ffc2d30abd323a5f1b780d375e91ed2c52cba666420f37ba1cf3731241b25c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. `xeno init --host gitlab` writes `.gitlab-ci.yml`. It calls `xeno gate verify` and
   `xeno gate run` and nothing else: no build, no test, no install beyond the pinned runner the
   GitHub wrapper also installs.

2. The job is scoped to merge request pipelines, `$CI_PIPELINE_SOURCE == "merge_request_event"`,
   because `CI_MERGE_REQUEST_DIFF_BASE_SHA` exists in no other pipeline type. The generated file
   passes `--base $CI_MERGE_REQUEST_DIFF_BASE_SHA` and `--head $CI_COMMIT_SHA`.

3. The generated `project.yaml` names the host the wrapper was generated for:
   `tracker.adapter: gitlab` and `tracker.base_url: https://gitlab.com/api/v4` for `--host gitlab`,
   and `github` with `https://api.github.com` for `--host github`.

4. `--host` has no default in `cmd/xeno`. Absent, `xeno init` refuses with a message that lists the
   hosts there are, and the refusal is a `runner.Refusal` so the exit code is 1 by A11.

5. `WrapperHosts()` derives from `wrapperHosts` and is sorted, so a third host is one row and no
   second list to update. Asserted rather than left to inspection.

6. `xeno init --host github` writes the same `.github/workflows/xeno-gate.yml` and the same
   `project.yaml` as before this intent, save for the tracker block's two values, which are the same
   strings arriving from a different place. Verified against `main`'s binary.

7. The manual list names the squash commit message template: GitLab's must contain `%{description}`
   or a footer does not survive a squash, which is the dependency `CONTRIBUTING.md` already records
   for GitHub.

8. `go build`, `go vet ./...`, `go test ./...` clean; `gofmt -l .` outside `vendor/` silent;
   `xeno gate verify` exit 0. A test asserts the GitLab wrapper contains both range expressions and
   the merge request scoping, because a wrapper whose range is wrong produces verdicts about the
   wrong commits and nothing downstream would say so.

<!-- xeno:section:non-goals -->
## Non goals

**No tracker adapter.** #104 assigns the tracker contract to WP12 and it is unbuilt for GitHub too.
Writing GitLab's half first would be building the second of two things neither of which exists, and
the adapter contract is WP12's to draw.

**No `BranchRules` adapter for GitLab.** That is #104's third piece and needs the Premium namespace
for two of its five rows. The consequence is stated rather than hidden: a repository initialised for
GitLab by this intent has `tracker.adapter: gitlab`, and `xeno enforcement check` there refuses with
the list of adapters that exist. Since #143 that is a refusal naming the field rather than a wrong
answer, which is the behaviour that makes shipping this piece before the third one safe.

**No change to the GitHub wrapper's content.** A second host arriving is not a reason for the first
one's output to move, and criterion 6 is how that is held.

**No `.gitlab-ci.yml` for this repository.** The wrappers are templates for other people's projects
and not pipelines for this one, which the plan says in as many words. This repository stays on GitHub
Actions and gains a scaffold file, not a pipeline.

**No new configuration field.** `tracker.adapter` and `tracker.base_url` exist; what changes is where
their values come from. A new field would be a spec change first by the second standing rule.

**No attempt to verify the wrapper by running it.** There is no GitLab project to run it in, and
criterion 8's assertion is over the generated text. What a real pipeline would add is the confirmation
that the two variables are populated as documented, and that belongs with the first repository anybody
initialises for GitLab.

<!-- xeno:section:constraints -->
## Constraints

**GitLab's variable availability binds the wrapper's shape.** `CI_MERGE_REQUEST_DIFF_BASE_SHA` is
documented as available only in merge request pipelines, and only while the merge request is open.
The scoping is therefore not a stylistic choice matching GitHub's `on: pull_request`; it is what
makes the base reachable, and a wrapper without it computes a range from an empty string.

**#98 binds the method.** A wrapper is a scaffold file plus a table row. If this intent finds itself
adding a branch to a generator, the design was wrong rather than the host unusual, and that is worth
saying because it is the claim #98 made and this is its first test.

**The plan binds what a wrapper may contain.** One command and nothing else. A build or test step
here would make it the project's own pipeline, which is a different file with the opposite purpose.

**A11 binds the refusal.** A missing `--host` is a refusal with a reason, exit 1, not a
could-not-run. The precedent for the wording is `Wrapper`'s existing unknown host error and
`BranchRulesFor`'s from #143, and the three should read alike.

**Section 12 binds the tracker block.** `adapter`, `project`, `base_url` and `auth` are section 12's
fields and this intent adds none.

**A16 and the file conventions bind the new files.** SPDX on Go, no per file copyright line, 88
columns. The scaffold is YAML with a generated header naming its source, as `ci-github.yml` has.

**Criterion 6 binds by comparison, not by reading.** The GitHub output is checked against `main`'s
binary, because a scaffold is a template and a change to how it is rendered can move whitespace
nobody would notice by eye.
