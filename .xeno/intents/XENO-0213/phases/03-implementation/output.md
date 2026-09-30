---
intent: github.com/triplem/xeno#145
phase: 03-implementation
created: "2026-09-30T05:51:23Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4e41c6b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b2e517c05c97cd3f1ad9bb0b36e37b1b4abbec665de85be14e5d8c180723ab0a
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

**`WrapperHost` describes a host completely.** Three fields added: `Adapter` and `APIBase` for the
tracker block, and `Squash` for the manual item. The last was a function branching on `h.ID` for one
draft, which is the shape this design exists to avoid: a third host would have added a branch to code
rather than a row to the table. It is a row now, and `grep` for a host name in non-test Go finds only
table keys and the adapter registry from #143.

**`wrapperHosts` gains the gitlab row.** `.gitlab-ci.yml`,
`$CI_MERGE_REQUEST_DIFF_BASE_SHA` and `$CI_COMMIT_SHA`, both confirmed against GitLab's reference,
with a comment on the row saying why the pipeline scoping is not a stylistic match to the other host.

**`WrapperHosts()` is derived from the table and sorted.** It returned `[]string{"github"}`.

**`internal/scaffold/files/ci-gitlab.yml`, new.** The GitHub wrapper's three comment paragraphs
unchanged, because their reasons are the host's neither time, then GitLab's syntax: `stages`, one job,
`rules` scoped to `merge_request_event`, `GIT_DEPTH: 0` for the same reason the other wrapper asks for
`fetch-depth: 0`, and a `script` of the install placeholder, `xeno gate verify` and `xeno gate run`.
The generated file parses as YAML and the folded `script` entry resolves to one command with both
range flags, which was checked rather than assumed.

**`project.yaml`'s tracker block is templated.** `adapter` and `base_url` become `{{.Adapter}}` and
`{{.APIBase}}`. The `base_url` comment changed with them: "a default, overridden for Enterprise
Server" described a literal, and the value is now written per host, so it says it is a setting and
what the other case is.

**`scaffold.Project` gains `Adapter` and `APIBase`.**

**The host is resolved once, before either file is generated.** `Init` called `projectYAML` at one
point and `Wrapper` forty lines later; the lookup moved up and `projectYAML` takes the row. This is
what makes the two files agree by construction rather than by both reading the same flag.

**Both defaults are gone.** `cmd/xeno`'s `--host` default and `init.go`'s `if host == "" { host =
"github" }`. `Wrapper` refuses an empty host with the list, beside its existing refusal of an unknown
one, so there is one sentence for the two cases rather than a second written for the occasion.

**Tests.** Six new, in `internal/runner`: the refusal and that it is a `Refusal` by A11, both range
expressions present in every host's wrapper, the GitLab scoping, the tracker block naming the
wrapper's host for every host, `WrapperHosts` derived and sorted, and the squash item present for
every host. Four of the six iterate `WrapperHosts()`, so a third host is tested by being added.

<!-- xeno:section:deviations -->
## Deviations from the design

**Twelve existing tests were edited, and the distinction from criterion 5 of #143 matters.** Every
`InitOptions{...}` in `init_test.go` gained `Host: "github"`, because `--host` is now required and
those twelve call sites relied on the default. None of them is a test of the host: they are the
gitignore, the version mismatch, the vendoring, the manual list, the round trip. Their assertions are
untouched and what changed is setup made explicit by a deliberate break, which criterion 4 asked for.
The rule the previous intent wrote — an assertion that will not hold is a finding, not a test to
update — is about expected values, and none moved here. Recorded because "twelve tests changed" looks
like exactly the thing that rule forbids, and a reviewer should be able to see which it was.

**The `base_url` comment in the scaffold changed, which criterion 6 said would not happen.** Criterion
6 asked that `--host github` write the same `project.yaml` as before "save for the tracker block's two
values". The values are identical; the comment beside one of them is not. "A default, overridden for
Enterprise Server" was a true sentence about a literal and is a false one about a templated value that
is written per host, so it now says the address is a setting and names the self managed case. Leaving
it would have been a comment that described the code it replaced, which is the convention P5 of the
last intent already caught once.

**`squashSetting` was a function and is now a table field.** The first draft branched on `h.ID ==
"gitlab"`, which passed every criterion and was the wrong shape: #98's design is that a host is a row,
and a per host sentence in a function means a third host edits code. Caught by grepping for host names
in non-test Go, which criterion 3 made a habit of.

**The wrapper is not run anywhere.** P1's non-goals say so and it stays true: there is no GitLab
project to run it in. The generated YAML was parsed and the folded `script` line resolved, so the file
is well formed and the range flags reach one command; whether GitLab populates the two variables as
documented is taken from the reference and not observed. This is the largest unverified claim in the
intent and it is in `gaps` as well.

**Nothing else.** No spec change, no new field in `project.yaml`, no dependency, no change to the
GitHub wrapper's body, no change to the gate path.
