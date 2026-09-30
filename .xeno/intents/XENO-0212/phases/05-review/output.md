---
intent: github.com/triplem/xeno#143
phase: 05-review
created: "2026-09-29T20:38:57Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+530ce03.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6fc0a7c8cd6c116166e5b8f891c7775060cb9409ae5f616d61fb2af0e4ddccbd
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
| The three standing rules | held. No document edited; `ASSUMPTIONS.md` is not touched by this intent at all, since A65 already says what was decided and this implements it. No field, gate, tool or rule invented: the requirement names are section 13's and are now constants rather than string literals, which narrows the surface rather than widening it. The change belongs to #143, labelled `wp12`, on `143-the-host-port` |
| A65 implemented rather than reinterpreted | yes. An adapter returns requirements, `Compare` is the filter, the waive logic and the `merge_method` line, and `Protection` and `ReviewsExpressible` are deleted rather than moved into the domain under another name |
| #104's first done-when | met. `project.yaml` selects the host, and no host name or host default appears outside `internal/host/github` — including in comments, which took two rewordings |
| The report is unchanged | proved against `main`'s own binary, byte identical apart from `checked_at` |
| A42 | held and checked, not assumed. `internal/gates` reaches no network and the CI check is untouched |
| A11 | held. A configuration that cannot be read is a `runner.Refusal`, so exit 1 rather than 2, and `report` already had that path |
| Criterion 5, the assertions | held. Every case from `main` survives with its wording; one split across the new boundary, six added for new behaviour |
| Eighty-eight columns, SPDX, no copyright line | held on all three new files |
| The check suite | build, vet, test, gofmt, `gate verify` all clean |
| Commit and reference convention | Conventional Commits subject with no issue reference, `Closes #143` in the footer and in the merge request description |

**What a reviewer should push back on.** Three things, in order of how much they deserve it.

The **line count**: 223 lines more across three files for a report that is byte for byte the same.
The defence is in P4's results and is honest about being a cost. A reviewer who thinks a port for
one adapter is not worth 223 lines is making #97's own argument, and the answer is that A65 settled
the vocabulary on evidence from a second host, so this is the last cheap moment to draw it.

The **`Declared` parameter** on `BranchRules`. It is the thing that keeps the port from being
`Protection` with extra steps, and it is also the thing that makes the interface less abstract than
it looks: an adapter has to understand the declaration, not just the host.

The **unavailable answer being four named requirements**. It could have been an absence with the
reason carried some other way. The regression in `deviations` is why it is not, and a reviewer may
prefer a different fix to the same problem.

<!-- xeno:section:release-notes -->
## Release notes

No change to any report, any artifact or any command's output. The enforcement report for a GitHub
project is what it was, in the same words.

One change a project can see: `xeno enforcement check` now refuses where `.xeno/config/project.yaml`
carries no `tracker.adapter` or no `tracker.base_url`, naming the missing field, instead of assuming
GitHub's hosted address. Every repository the scaffold wrote carries both, so this reaches only a
configuration that was relying on the assumption.

For a reader of the repository: the host lives behind `host.BranchRules` in `internal/host`, with
`internal/host/github` holding every GitHub specific thing — the URL, the header, the field names
and the meaning of a 403. A second host is a package beside it and an entry in one table.

<!-- xeno:section:residual-risk -->
## Residual risk

**The port may be GitHub's shape with an interface in front of it.** This is the risk P4 names as
the claim the piece cannot test, and it is the one that matters. One adapter cannot distinguish a
boundary from an indirection. The evidence that it is a boundary is A65's, from outside this intent:
GitLab answers `allow_bypass` with three lists of grants, which is why an adapter evaluates rather
than reports, and `Declared` being a parameter is the consequence. If #104's third piece has to
change `BranchRules` to fit GitLab, that is this intent being wrong and not a refinement, and it
should be recorded as such.

**A refusal now reaches a configuration that used to work.** Any `project.yaml` without
`tracker.adapter` or `tracker.base_url` got GitHub's hosted address and now gets a refusal. The
scaffold has written both since WP9 and this repository carries both, so the population is
repositories configured by hand before WP9, of which there is one and it is this one. The refusal
names the field, which is the mitigation; there is no migration because there is nothing to migrate.

**`protection` is duplicated in spirit if a second adapter copies it.** The struct is unexported and
this host's own, which is A65's point, and the honest risk is that a GitLab adapter starts by
copying it and then diverges field by field until the two are a shared type nobody named. The answer
when it happens is that they are different types and should look it.

**The error paths are uncovered.** 93.5% and 88.9%, and what is missing is transport failure. A bug
there surfaces as a confusing message rather than a wrong verdict, which is why it was not worth a
fake client; it is still uncovered code on the path a user hits when a network is broken.

**Nothing verifies that one adapter type can implement both ports.** #97 chose two ports partly on
that basis and `Tracker` does not exist, so the claim is untested. WP12 tests it or contradicts it.

**Not a risk.** That the domain no longer uses the network while its package comment used to say it
was the only package that did. The comment was rewritten, and A42's property — that the gate path
reaches nothing — is checked in CI and held.
