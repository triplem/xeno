---
intent: github.com/triplem/xeno#145
phase: 05-review
created: "2026-09-30T05:53:24Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4e41c6b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 70bda9af3edbaaa4e4bebd9ad79d6baa60133df2671485ad73fd9ca8c241a2d7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

| item | state |
|---|---|
| The three standing rules | held. No document edited. No invented field: `tracker.adapter` and `tracker.base_url` are section 12's and this changes where their values come from. The change belongs to #145, labelled `wp12`, on `145-the-gitlab-wrapper` |
| #104's first done-when, completed | yes. #143 removed the enforcement path's default; this removes the flag's and `init.go`'s. Every host name in non-test Go is now a table key |
| #104's third done-when | yes. `xeno init --host gitlab` writes a wrapper that calls two xeno commands and nothing else |
| #98's design | held, and tested for the first time. One scaffold file, one table row, no generator branch |
| The GitHub output | byte identical for the wrapper; one comment changed in `project.yaml` and no value, recorded in `deviations` |
| A11 | the missing host is a `*Refusal`, asserted, so exit 1 rather than 2 |
| The twelve edited tests | setup only, assertions untouched, and the distinction from #143's criterion 5 is written out in `deviations` because it looks like the thing that rule forbids |
| Eighty-eight columns, SPDX, no copyright line | held; the new file is YAML with the generated header naming its source, as the other wrapper has |
| The check suite | build, vet, test, gofmt clean; `gate verify` 156 verdicts, exit 0 |
| Commit and reference convention | Conventional Commits subject with no issue reference, `Closes #145` in the footer and in the merge request description |

**What a reviewer should push back on.**

**`image: alpine:latest`.** The wrapper pins the runner version and in the next line names an unpinned
image, which is the contradiction the pin exists to prevent. It is in `gaps` as a finding about the
scaffold. A reviewer may reasonably say a generated file should not ship the word `latest` at all, and
the counter is that any specific image is a guess about the adopter's registry. Neither answer is
obviously right and the current one is the one that does not pretend to know.

**That this ships without an enforcement adapter.** A repository initialised for GitLab gets a working
gate pipeline and a `tracker.adapter` nothing answers for. The defence is that #143 made that a
refusal naming the field rather than a wrong answer, so the half-built state is honest. A reviewer who
would rather ship the third piece first is arguing for a bigger change with a fixture dependency, and
should say so.

**The `base_url` comment.** Criterion 6 said the GitHub `project.yaml` would change in two values and
nothing else, and a comment changed. It is recorded rather than quietly allowed, and a reviewer who
holds criteria literally is right that this is a miss in the criterion's wording.

<!-- xeno:section:release-notes -->
## Release notes

**`xeno init` now requires `--host`.** It had defaulted to `github`. A script calling `xeno init`
without it gets a refusal naming the flag and listing the hosts, exit 1. The reason is that the host
decides the CI wrapper and the tracker configuration both, and a tool that picks one for you is a tool
with one host.

**`xeno init --host gitlab` generates `.gitlab-ci.yml`.** It runs on merge request pipelines, calls
`xeno gate verify` and `xeno gate run`, and passes both ends of the commit range from
`CI_MERGE_REQUEST_DIFF_BASE_SHA` and `CI_COMMIT_SHA`. The generated `project.yaml` names
`adapter: gitlab` with GitLab's API address.

**What a GitLab project does not get yet.** `xeno enforcement check` has no GitLab adapter, so it
refuses there and lists the hosts it can ask. That is #104's third piece. There is no tracker adapter
for any host.

**One more manual item.** `xeno init` now names the host's squash commit message setting, because a
squash rewrites the message and a `Refs` or `Closes` footer survives only where the template carries
the description through. GitLab's needs `%{description}`.

<!-- xeno:section:residual-risk -->
## Residual risk

**An empty base SHA would produce verdicts about the wrong commits rather than an error.** This is the
risk that outranks the rest. `gate run` takes `--base` and `--head` and computes a range; given an
empty base it does not obviously fail, it judges something else. The scoping to `merge_request_event`
is what prevents it, and the evidence that the scoping is sufficient is GitLab's reference. The first
real merge request settles it, and the failure mode if the reference is wrong is a green verdict about
the wrong change, which is worse than a red one. Worth watching for on the first GitLab adoption
specifically, rather than trusting that a passing pipeline means the range was right.

**`alpine:latest` moves under the project.** An unpinned image beside a pinned runner version. The
runner is pinned because a different one produces a different verdict; an image that changed could stop
the job running at all, which is loud, or change the shell the script runs under, which is not.

**The generated pair can still disagree if somebody adds a row wrongly.** `Adapter` and `APIBase` are
free strings in the table, so a row naming a wrapper for one host and an address for another compiles
and generates. The tests iterate the table and would catch a missing value, not a wrong one. What would
catch it is the adapter registry from #143 and the wrapper table being one list, which they are not, and
making them one is a change neither piece asked for.

**A repository initialised for GitLab is half configured by design, and nothing in it says so.**
`project.yaml` names an adapter that does not exist and the file does not mention that. The refusal from
`enforcement check` explains itself when somebody runs it, which may be weeks later. A comment in the
generated tracker block would have said it at the right moment, and adding one means the scaffold
carries a statement about which adapters exist, which dates.

**`internal/scaffold` at 0% coverage** is unchanged by this intent and now has one more template in it.

**Not a risk.** Requiring `--host`. It breaks scripts written against the default, the refusal names the
flag, and there is no population to migrate.
