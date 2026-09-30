---
intent: github.com/triplem/xeno#143
phase: 00-intake
created: "2026-09-29T20:26:57Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+530ce03.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 48a5f7a77d1577b326325f6e71b2c3decf9ceb2e347f51fae3684c4846a22f44
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

`internal/enforcement` is the domain package for what a project declares of its host, and it
knows exactly one host. Four things in it are GitHub's and not the domain's, which #97
enumerated: the URL shape `%s/repos/%s/branches/%s/protection`, the
`Accept: application/vnd.github+json` header, the JSON field names `enforce_admins`,
`required_status_checks` and `required_pull_request_reviews`, and the meaning of a 403, which
this code reads as "the plan does not offer branch protection" rather than as an error.

The fourth is why this is not a URL swap. A 403 meaning "not available on this plan" is a fact
about GitHub's pricing, and a second host answers the same question with a different status or
not at all.

`internal/runner/enforcement.go:50` completes the assumption: where `project.yaml` names no
base URL, the runner supplies `https://api.github.com`. The configuration to select a host
exists and is read — `tracker.adapter` and `tracker.base_url` are in the scaffold and in this
repository's own config — and the code defaults past it.

`Protection` is the shape of the problem rather than a part of it. It is a struct of host
neutral facts whose field names are GitHub's, with `ReviewsExpressible` as a device for the one
field where the host might say nothing. A65 settled that this is the wrong shape: an adapter
returns `[]enforcement.Requirement` and `Compare` shrinks to the waive logic, the
declared-or-not filter and section 13's unchecked `merge_method` line.

So the vocabulary is decided and the port is not written. That is the gap this intent closes,
and it is the piece #104's other two remaining pieces both build on: the enforcement mapping
needs somewhere to put GitLab's answers, and the wrapper's tracker side needs the package to
exist.

<!-- xeno:section:scope -->
## Scope

**In scope.** `internal/host` with `BranchRules`, returning `[]enforcement.Requirement` per
A65. `internal/host/github` implementing it and owning the URL, the header, the field names and
the status code meanings. `internal/runner.EnforcementCheck` selecting the adapter from
`project.yaml` rather than defaulting in code, and refusing with a reason that names the field
where the configuration is incomplete. The retirement of `Protection`, `ReviewsExpressible` and
the `evaluate` half of the `requirements` table.

The existing cases in `enforcement_test.go` keep their assertions. A65 names them as the proof
that this is a move and not a rewrite, so where a case has to change package its assertions go
with it unchanged, and a case whose wording would have to be edited to pass is a finding rather
than a test to update.

**Out of scope.** `Tracker`. #104 names it beside `BranchRules` in prose and leaves it out of
the done-when this intent answers, and WP12 is not built, so declaring it here produces an
interface with no adapter and no caller — which is what #97 declined to do, one step further
along. It arrives with the work package whose contract it is.

Out of scope also: the GitLab adapter and the five row enforcement mapping, which is #104's
third piece and needs the Premium fixture; the wrapper, which is the fourth; and any change to
`project.yaml`'s shape, since the two fields this needs already exist.

**Not touched, and checked rather than assumed.** A42's property, that the gate path reaches no
network. A new package under `internal/` that imports `net/http` is exactly the kind of change
that could break it by accident, so the CI check is part of this intent's verification rather
than background.

<!-- xeno:section:context-rationale -->
## Why this context

The information base is `internal/enforcement/enforcement.go` in full, because every part of it
either moves, stays or is retired and the three have to be told apart; its test file, because
criterion 5 makes those assertions the acceptance; and `internal/runner/enforcement.go`, which
is the caller and holds the defaulted line.

`ASSUMPTIONS.md` A65 is the decision this implements and is read as binding rather than as
advice. A42 is read for the property that must survive. A13 and A39 are read to confirm they are
untouched: the pipeline artifact stand in and the report's gitignored location have nothing to
do with where the host knowledge lives.

#97's closing comment is read for the four things it says move and the two ports it decided, and
#104's first piece for the done-when this intent answers. Both are read as constraints.

`.xeno/config/project.yaml` and `internal/scaffold/files/project.yaml` are read to confirm that
`tracker.adapter` and `tracker.base_url` are already present in both, which is what makes
criterion 3 a removal rather than a migration: nothing has to be written into an existing
repository's configuration for the default to be safe to drop.

Section 13 is read for the requirement names, which stay in the domain. Section 8 is read for
the edition, and is the reason nothing here is verified against GitLab: this intent writes the
port, and the host it is verified against is the one already configured.
