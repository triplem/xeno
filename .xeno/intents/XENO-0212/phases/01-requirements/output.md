---
intent: github.com/triplem/xeno#143
phase: 01-requirements
created: "2026-09-29T20:27:46Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+530ce03.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ba1a872477adce4f1ee25484fb3db595a2b73ebab64e8353e3a1dc29bb58aa95
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

1. `internal/host` holds `BranchRules`, whose method returns `[]enforcement.Requirement` and an
   error. The package imports `internal/enforcement` and nothing imports it back.

2. `internal/host/github` implements `BranchRules`. The string `/repos/`, the header
   `application/vnd.github+json`, the field names `enforce_admins`, `required_status_checks`
   and `required_pull_request_reviews`, and the mapping of 403, 404 and 401 appear in that
   package and nowhere else. `grep -rn 'api\.github\.com\|vnd\.github' --include=*.go` outside
   that package and outside the scaffold finds nothing.

3. `internal/runner.EnforcementCheck` selects the adapter by `tracker.adapter` and passes
   `tracker.base_url`. Where either is empty it refuses with a message naming the field and the
   file, and the refusal is a `runner.Refusal`, so `cmd/xeno` reports exit 1 by A11 rather than
   2. An unknown adapter name refuses and lists what there is, as `Wrapper` does for hosts.

4. `enforcement.Protection`, `ReviewsExpressible` and the `evaluate` half of the `requirements`
   table are gone. `Compare` takes `[]Requirement` from an adapter and applies the
   declared-or-not filter, the waive logic and section 13's `merge_method` line. The
   requirement names stay enumerated in `internal/enforcement`, and a name an adapter returns
   that the domain does not have is dropped and reported, not passed through.

5. Every assertion in the existing `enforcement_test.go` still holds, unchanged in wording.
   Cases that test the transport and the decode move to `internal/host/github` with their
   assertions intact; cases that test waiving, the states and the `merge_method` line stay.
   A test whose expected string had to be edited to pass is a finding recorded in `deviations`,
   not a test updated.

6. A42 holds: `go list -deps ./internal/gates` names neither `net` nor `net/http`, and the CI
   check in `.github/workflows/xeno.yml` is untouched.

7. `go build`, `go vet ./...`, `go test ./...` clean; `gofmt -l .` outside `vendor/` silent;
   `xeno gate verify` exit 0; and `xeno enforcement check` against this repository still reports
   `required_pipeline` met and `allow_bypass` met, which is what #84 decided and A27 measured.

<!-- xeno:section:non-goals -->
## Non goals

**`Tracker` is not written.** #104 names it beside `BranchRules` in prose and omits it from the
done-when this intent answers. WP12 is not built, so the interface would have no adapter and no
caller, and #97's own reason applies with more force than it did there: a port with one adapter
is a sentence in an issue, and a port with none is a sentence about a sentence. It arrives with
the package whose contract it is, and this is recorded in `deviations` rather than left for a
reader of #104 to notice.

**No GitLab.** No second adapter, no five row mapping, no `ci-gitlab.yml`. Those are #104's
third and fourth pieces and the third needs the Premium fixture that is deliberately not
arranged. The port having one adapter at the end of this intent is the expected state, not an
incomplete one: what makes it a port rather than an indirection is that the second adapter needs
no change to it, and that claim is tested when the second arrives.

**No change to `project.yaml`'s shape.** `tracker.adapter` and `tracker.base_url` already exist
in the scaffold and in this repository's configuration, so criterion 3 removes a default rather
than migrating a format. The second standing rule would make a new field a spec change first.

**No new behaviour in the report.** The same requirements with the same words and the same
states. A65 says the GitHub phrasings move rather than get rewritten, and criterion 5 is how that
is held. If the report for this repository changes at all, the intent has failed.

**No new dependency.** One vendored library, and `net/http` is standard.

<!-- xeno:section:constraints -->
## Constraints

**A65 binds the signature.** An adapter returns requirements. This intent may not revisit that
choice, and Appendix B is why: the row sits inside a sealed phase's `artifacts_hash`, so
disagreeing with it is a new intent and not an edit here.

**A42 binds the import graph.** No gate reaches the network. `internal/host/github` imports
`net/http`, so the risk is a path from `internal/gates` to it, and the CI check is what holds
the property rather than a test. It is checked in P4 rather than assumed.

**Section 13 binds the names.** `required_pipeline`, `allow_bypass`, `approvals.required`,
`approvals.not_by_author` and `merge_method`, and the second standing rule makes that list a
budget. The domain keeps enumerating them, which is what makes an adapter's freedom over words
safe.

**A11 binds the exit codes.** A configuration that cannot be read is a refusal with a reason,
which is 1, and not a could-not-run, which is 2. The distinction is why criterion 3 asks for a
`runner.Refusal` specifically.

**The report's own shape is fixed by section 7 and A39.** `Report` keeps its fields and its
gitignored location, and nothing about where a requirement's words come from reaches the file's
format.

**Criterion 5 is a constraint on method, not only a test.** The existing assertions are the
specification of the GitHub phrasings. Working by making tests pass is ordinary; working by
editing an expected string until it matches new output would destroy the one piece of evidence
that this was a move, so an assertion that will not hold is a finding.

**Eighty-eight columns, and SPDX on every new file, with no copyright line by A16.**
