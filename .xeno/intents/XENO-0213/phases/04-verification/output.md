---
intent: github.com/triplem/xeno#145
phase: 04-verification
created: "2026-09-30T05:52:30Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4e41c6b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 58b0be75926dc8575f8c2bb8bae43de7bfc764accb13d68fd24a96b36b855662
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| criterion | how it is verified |
|---|---|
| 1, `--host gitlab` writes a wrapper that calls one command | `TestEveryWrapperPassesBothEndsOfTheRange` iterates `WrapperHosts()` and asserts `xeno gate verify` and `xeno gate run` are present; the generated file was read and parsed by hand as well, and carries no build or test step |
| 2, scoped to merge request pipelines, both range flags | `TestTheGitLabWrapperRunsOnlyOnMergeRequests` asserts the `rules:` condition; the same iterating test asserts `$CI_MERGE_REQUEST_DIFF_BASE_SHA` and `$CI_COMMIT_SHA` |
| 3, the tracker block names the wrapper's host | `TestTheTrackerBlockNamesTheHostTheWrapperWasGeneratedFor`, over every host, reading the generated YAML back through `fm.ReadYAML` rather than grepping it |
| 4, `--host` has no default, refusal is a `Refusal` | `TestInitRefusesWithoutAHost`, asserting the message lists both hosts and that `errors.As` finds a `*Refusal`, which is what makes it exit 1 by A11 |
| 5, `WrapperHosts()` derived and sorted | `TestWrapperHostsIsDerivedFromTheTableAndSorted`, comparing against `wrapperHosts` itself so the two cannot drift |
| 6, GitHub output unchanged | compared against `main`'s binary, below |
| 7, the squash setting in the manual list | `TestTheSquashSettingIsNamedForEveryHost`, over every host |
| 8, the suite | below |

Four of the six new tests iterate `WrapperHosts()` rather than naming a host. That is deliberate and
it is the part worth keeping: a third host is tested by being added to the table, and a host added
without a wrapper file, without range expressions, without a tracker address or without a squash
sentence fails one of them. The table is the thing under test, not GitLab.

The generated GitLab YAML was also parsed with a YAML reader and its `script` list inspected, because
the `xeno gate run` entry is a folded plain scalar over three lines. It resolves to one command with
both flags. A test asserting the substrings would not have caught a folding error that produced two
commands or a broken one.

<!-- xeno:section:results -->
## Results

All eight criteria hold.

| check | result |
|---|---|
| `go build -o xeno ./cmd/xeno` | builds |
| `go vet ./...` | clean |
| `go test ./...` | every package ok |
| `gofmt -l .` outside `vendor/` | prints nothing |
| `xeno gate verify` | 156 verdicts, exit 0 |
| host names in non-test Go | only table keys, in `wrapperHosts` and #143's adapter registry |
| `internal/runner` coverage | 83.6% |

**The GitHub wrapper is byte identical to `main`'s**, with the version stamp normalised, since this
tree is `dirty` and `main`'s is not.

**The GitHub `project.yaml` differs in one comment and no value.** `adapter: github` and
`base_url: https://api.github.com` are the same strings arriving from the table instead of from
literals. The comment on `base_url` changed, which `deviations` records and explains: it described a
literal that no longer exists.

**`xeno init --host gitlab`** writes `.gitlab-ci.yml`, a `project.yaml` with `adapter: gitlab` and
`base_url: https://gitlab.com/api/v4`, and prints five manual items, the fifth naming
`%{description}`.

**`xeno init` with no host** refuses: "no --host; there is github, gitlab" with the reason, and the
error is a `*Refusal`.

The generated GitLab file parses, and its `script` resolves to three entries, the third being
`xeno gate run --intent "$XENO_INTENT" --phase "$XENO_PHASE" --base $CI_MERGE_REQUEST_DIFF_BASE_SHA
--head $CI_COMMIT_SHA` on one line.

**#98's claim held.** A second host was one scaffold file and one table row. No generator gained a
branch, and the one branch that appeared in a draft — the squash sentence on `h.ID` — was moved into
the row, which is the same claim being enforced rather than an exception to it.

<!-- xeno:section:gaps -->
## Gaps

**The wrapper has never run, and this is the gap that matters.** Everything asserted about
`.gitlab-ci.yml` is about its text: it parses, it carries both range expressions, it is scoped to
merge request pipelines, it calls two commands. What is taken from GitLab's reference rather than
observed is that `CI_MERGE_REQUEST_DIFF_BASE_SHA` is populated in a merge request pipeline with the
value the wrapper needs. The reference says so plainly and the variable is long standing, so the risk
is low and it is not zero: a base SHA that arrived empty would make every CI run judge a range from an
empty string, and `gate run` would produce a verdict about the wrong commits rather than fail.

What closes it is the first repository anybody initialises for GitLab, and one merge request in it.
That is cheaper than anything this intent could have built, which is why nothing was built for it.

**`image: alpine:latest` is a guess.** The GitHub wrapper says "install xeno here" and leaves the
install to the adopter, and GitLab needs an image named for the job to run at all. `alpine:latest` is
the smallest thing that will hold a static binary and it is unpinned, which contradicts the reason the
runner version beside it is pinned. A project will change it; the generated file should probably not
be the place that says `latest`, and that is a finding about this scaffold rather than about the
intent.

**`internal/scaffold` is at 0% coverage** and always was. Its templates are exercised through
`internal/runner`'s tests, which is why the number is misleading rather than alarming, but nothing
tests `RenderWrapper` or `RenderProject` directly and a template that failed to parse would surface as
a runner test failing for a reason two packages away.

**The tracker half of #104's fourth piece is untouched**, by P1's scope. So "GitLab support" after this
intent means: a wrapper that runs gates, a configuration that names the host, and neither an
enforcement adapter nor a tracker. A reader of the issue list could reasonably think #104 is three
quarters done and the remaining quarter is small. It is not: the third piece needs the Premium
namespace and carries the five row mapping.

**Not a gap.** That `--host` breaking existing scripts is untested against a real script. The break is
the criterion, the refusal names the flag and lists the hosts, and there is no population to migrate.
