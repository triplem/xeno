---
intent: github.com/triplem/xeno#145
phase: 00-intake
created: "2026-09-30T05:43:53Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4e41c6b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3843c51c73f5b0ff2f21e956d1c0c2e97dd28f4681626fafd4e1084ce71c08cf
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

`xeno init` takes `--host` and can generate one wrapper. `wrapperHosts` is a table with a single
row and #98 built it that way on purpose, so that a second host is an entry and a scaffold file
rather than a second generator. The claim has never been tested, because there has only ever been
one row.

Two things are wrong beyond the missing row.

`project.yaml` is generated with `adapter: github` and `base_url: https://api.github.com` as
literals, while the wrapper is generated for whatever `--host` says. Those are the same choice made
twice, once from a flag and once from a constant, so `--host gitlab` produces a repository whose
pipeline is GitLab's and whose tracker configuration is GitHub's. Since #143 that configuration
refuses rather than silently asking the wrong host, which turns a wrong answer into a stopped
command and is still a repository `xeno init` should not have written.

`--host` defaults to `github` in `cmd/xeno/main.go`. #104's first done-when is that no default host
name appears in code, and #143 removed the one in the enforcement path. This is the other one, and
it is the one that decides what an adopter gets when they do not say.

There is also a fact about GitLab that #104 recorded as unconfirmed and that changes the wrapper's
shape. `CI_MERGE_REQUEST_DIFF_BASE_SHA` is the base of the range, and it exists only in merge
request pipelines and only while the merge request is open. GitHub's wrapper is scoped with
`on: pull_request` for reasons of its own; GitLab's has to be scoped to `merge_request_event` or the
base is empty and the runner judges a range nobody wrote. That is the kind of thing #104 meant by
the shapes being close rather than the same.

<!-- xeno:section:scope -->
## Scope

**In scope.** `internal/scaffold/files/ci-gitlab.yml`, scoped to merge request pipelines and calling
`xeno gate verify` and `xeno gate run` and nothing else. One row in `wrapperHosts` with the path and
the two confirmed range expressions. `WrapperHosts()` derived from the table instead of repeating
it. The scaffold's `tracker.adapter` and `tracker.base_url` becoming values of the host `xeno init`
was given. The removal of `--host`'s default in `cmd/xeno`, which is the last default host name in
code and the remainder of #104's first done-when. The squash commit message template added to the
manual list `xeno init` prints.

**Out of scope.** The tracker adapter. #104 assigns "resolve an intent, read an issue, write a
comment, resolve credentials" to WP12, and none of it is built for GitHub either, so writing
GitLab's half first would be building the second of two things neither of which exists. What this
intent takes from the tracker side is the one item that is a host setting rather than code.

Out of scope also: the enforcement mapping, which is #104's third piece and needs the Premium
namespace; and any GitLab adapter for `BranchRules`, which is the same piece. A repository
initialised for GitLab by this intent gets a working wrapper and a `tracker.adapter` naming a host
no adapter answers for yet, so `xeno enforcement check` refuses with the list of adapters there
are. That is the honest state and it is named here rather than discovered.

**Not touched.** The GitHub wrapper's content. Criterion 6 asks that `xeno init --host github`
writes what it wrote before, byte for byte, because a second host arriving is not a reason for the
first one's output to move.

<!-- xeno:section:context-rationale -->
## Why this context

The information base is `internal/runner/wrapper.go` for the table and its two range expressions;
`internal/scaffold/files/ci-github.yml` as the shape a wrapper has, since the GitLab one is the same
file for a different host and not a new design; `internal/scaffold/files/project.yaml` for the two
literals that become values; `internal/runner/init.go` for where the host is resolved and where the
manual list is built; and `cmd/xeno/main.go` for the flag's default.

`internal/scaffold/scaffold.go` is read for `RenderWrapper` and the `Wrapper` and `Project` structs,
because the host has to reach the project scaffold and does not today.

Outside the tree, GitLab's predefined variable reference is read for the two names #104 marked "to
confirm". Both are confirmed and one carries a constraint the issue did not have: the base exists
only in merge request pipelines. That is why the documentation was read rather than the names taken
from #104, and it is the second time on this work package that asking the host changed the shape of
the answer.

#98 is read for why the wrapper is a scaffold file, #104 for the done-when, and A65 and #143 for
what a `tracker.adapter` naming an unadaptered host now does. `CONTRIBUTING.md` is read for the
existing record of the same squash dependency on GitHub, so that the manual item names the setting
the way the project already names it.
