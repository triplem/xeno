---
intent: github.com/triplem/xeno#94
phase: 01-requirements
created: "2026-10-08T20:14:31Z"
schema_version: "1.0"
runner_version: dev+30b1dea
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 25758f0d043cb5342c96f8255e06e0fd9387349617b6752ddcdcd9144bc7e090
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Each criterion is a state of the tree and a verdict over it, read on the branch.

1. **The policy is a file and it fails on any finding.** `.golangci.yml` at the root,
   `version: "2"`, enables `govet`, `staticcheck`, `errcheck` and `gosec` and no other
   linter; `golangci-lint run ./...` with that file exits 0 on the branch, and the file
   says in a comment why each exclusion is there.

2. **Four findings in shipped code are accepted with their reason, and no more.** Each
   `nolint:gosec` in a non-test file names the rule it accepts and carries the reason on
   the same line; there are four, in `internal/git/git.go`, `internal/host/gitlab/
   gitlab.go` twice, and `internal/rules/rules.go`; `grep -rn nolint --include='*.go'
   internal cmd` finds those and nothing else.

3. **Tests are outside gosec and inside everything else.** An exclusion rule in
   `.golangci.yml` keeps gosec off `_test.go` with the reason beside it; errcheck and
   staticcheck still read the tests and report nothing, because the eleven findings they
   had are fixed by checking the error or simplifying the expression, not by a second
   exclusion.

4. **The linter has a workflow in the scanners' shape.** `.github/workflows/lint.yml`,
   job id `lint`, runs on a push to `main` and on every pull request, installs
   golangci-lint through `golangci/golangci-lint-action` pinned by commit sha with the
   tag in a comment and `version: v2.14.0`, writes `.xeno/local/scan/golangci-lint.json`,
   prints what it found, declares a manifest with `kind: scan`, `job: lint` and the
   result the run returned, uploads the directory as `scan-lint`, and raises the
   failure last.

5. **The vulnerability scan of the source has a workflow in the same shape.**
   `.github/workflows/govulncheck.yml`, job id `govulncheck`, runs on a push, on every
   pull request and on a daily schedule, installs `golang.org/x/vuln/cmd/govulncheck`
   at `v1.8.0` with `go install`, scans with `-format json` into
   `.xeno/local/scan/govulncheck.json`, judges the report in a step of its own, prints
   every finding with its fixed version, declares the manifest, uploads `scan-govulncheck`,
   and raises the failure last.

6. **The judgement is the one the scanner's text mode makes, written down.** The step
   fails on a finding whose trace reaches a function, which is what `govulncheck` exits
   3 for; a finding at module or package level is printed and does not fail the job,
   and a reachable finding with no fixed version is printed and does not fail the job,
   for the reason `trivy.yaml` gives. The report's `db_last_modified` is printed beside
   the result, so a reader can judge the scan's coverage from the log.

7. **semgrep is gone from the tree and from every sentence outside the trail.**
   `.github/workflows/semgrep.yml` and `.semgrep/` do not exist; `grep -rn semgrep`
   over the repository with `.xeno/intents/` excluded finds nothing, so `CONTRIBUTING.md`,
   `gitleaks.yml`, `xeno.yml`, `renovate.json`, `docs/supply-chain.md` and
   `scripts/github-settings.sh` have each been rewritten where they named it.

8. **The pin table is held to the tree in both directions, with the new rows.** `go test
   ./internal/model/` is green: the table carries `golangci/golangci-lint-action` at its
   sha and tag, `golangci-lint` at `v2.14.0`, `govulncheck` at `v1.8.0` and its database
   as a row that cannot be pinned, semgrep's two rows are gone, and the test's exempt
   list has lost `semgrep's rules` and gained the database row with its reason.

9. **The required checks name the new jobs.** `scripts/github-settings.sh` lists
   `verify`, `audit`, `gitleaks`, `lint`, `govulncheck`, `trivy`; the script is run
   against the host before the merge and `gh api .../protection` shows that list, since
   a required context that never reports blocks every merge, this one included.

10. **Renovate sees the two new pins.** `renovate.json` carries a custom manager for the
    action's `version:` input in the shape the trivy manager has, and one for the
    `@v1.8.0` of the `go install` line against the Go datasource, and its description no
    longer says semgrep is unmanaged.

11. **The suite and the gates stay green, and no behaviour changes.** `go test ./...`,
    `gofmt -l`, `go vet ./...` and `./xeno gate verify` exit 0; every change to a
    non-test file is a suppression comment or a simplification staticcheck named, and
    the diff shows no other line.

12. **The register carries the two implementation decisions.** `docs/assumptions.md`
    gains A104 for the linter's threshold and what the tests are exempt from, and A105
    for reading "trivy for the container scan" as `trivy.yml` unchanged.

<!-- xeno:section:non-goals -->
## Non goals

- **depguard.** It reads direct imports and A42 is about the closure; the `go list
  -deps` step in `xeno.yml` keeps the rule, and the design records the departure from
  the comment's step two.
- **A period of running both.** The comment proposed removing semgrep last, after the
  three report on the same tree; the maintainer approved the replacement outright, and
  the measured run is the comparison: semgrep reports nothing today and the new set
  reports 51, so nothing semgrep carried is lost by not overlapping.
- **Changes to trivy.** It stays as it is, scanning the built binary daily; whether it
  should scan the image #282 built is its own question.
- **Attaching a report to a phase.** The manifests stay so that a phase can declare
  one; this intent declares none, as every intent since #51.
- **gosec on test files.** Tests write fixtures with ordinary permissions, launch
  subprocesses on purpose and carry Amazon's documented example key as the secret
  filter's input; a rule set that fires on those teaches people to route around it.
- **nolintlint, or any linter beyond the four.** Each one is a threshold chosen before
  seeing its result, which the analysis said not to do.
- **The generated wrapper.** What `xeno init` writes for an adopter is untouched; this
  is this repository's own pipeline.
- **A SARIF upload or code scanning alerts.** The report is a file in the manifest's
  shape, which is what the process can attach; the host's own scanning surface is a
  second record of the same finding.

<!-- xeno:section:constraints -->
## Constraints

- Section 4's evidence shape: a scan is `kind: scan` with a job, a result, a file, a
  pipeline and a commit, and a manifest that names a file which does not exist declares
  something untrue, so each workflow tests for the report before declaring it.
- `docs/supply-chain.md`'s rule and the test behind it: every action pinned to a sha
  with the tag beside it, every tool to an exact version, and every pin on the page in
  both directions; a row that cannot be pinned is exempt in the test with its reason.
- The job id is the check name GitHub reports, and a required context that never reports
  blocks every merge; the rename lands in the script and on the host before the merge.
- One vendored dependency, and this adds none: the two tools are pipeline rows.
- A42 holds: `internal/gates` reaches no network, checked by the step that checks it.
- The scanners' shape: `continue-on-error` on the scan, the report printed, the manifest
  declared, the artifact uploaded, the failure raised last, so the run worth reading is
  the one that failed.
- `CLAUDE.md`: prose wraps at 88, commit subjects are Conventional Commits with the
  reference in the footer, a comment says what the construction is and why; a paragraph
  changed for the second time is replaced, not edited into.
- No sealed artifact changes; the trail's 549 verdicts stay as they are.
- The tool versions in the workflows are the ones the measured run used, 2.14.0 and
  1.8.0, read off the tools rather than off a release page.
