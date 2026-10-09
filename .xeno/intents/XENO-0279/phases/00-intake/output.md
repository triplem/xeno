---
intent: github.com/triplem/xeno#94
phase: 00-intake
created: "2026-10-08T20:10:48Z"
schema_version: "1.0"
runner_version: dev+30b1dea
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f7942c793409a86798848db26f58f4758c59e7d84267d0c9fe50af4ae692cb99
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

Approved by @triplem on 2026-10-08T19:40:07Z: For details see comment above

> **github.com/triplem/xeno#94** — Replace semgrep
>
> Some tools:
>
> govulncheck https://go.dev/blog/vuln
> Staticcheck https://staticcheck.dev/
>
> Golint https://golangci-lint.run/
>

The issue as it stood at 2026-10-08T20:10:48Z, read by `xeno phase start` and quoted rather than summarised. What this phase concludes about it belongs below.

The issue as `xeno phase start` read it, and the first intent started through the
clause #330 built: the label and the comment were both there, `intent start` read them
from the host and wrote the sentence above the quote, which is the positive path
XENO-0278's gaps said no real issue had yet exercised.

The issue's body names three tools. What decides the work is the analysis of
2026-09-27 in its comments and the two answers of 2026-10-08, which `phase start` reads
and does not quote, so they are restated here as the problem actually set.

**What semgrep does today.** `.github/workflows/semgrep.yml` runs the vendored `p/golang`
pack, 42 rules in `.semgrep/golang.yaml`, from an image pinned by digest, with `--error`
so that a finding fails the job, and writes a JSON report with a manifest for later
attachment. Fourteen of the 42 rules name libraries this project does not import,
gorilla, jwt-go, grpc, aws-lambda, and cannot fire; what is carried is the crypto, net,
filepath and injection audit, and it has reported nothing against this tree since it
went in.

**The mapping the analysis gave, and the maintainer approved.** The audit half of the
pack is gosec, run through golangci-lint, with staticcheck and errcheck beside it and
`go vet` as it is; known vulnerabilities are govulncheck, which matches call paths where
trivy matches versions; the one project rule semgrep could express, A42's gate path free
of the network, was to go to depguard. The three answers: go with the replacements,
since it is a project decision with no effect on projects using Xeno; gosec at "some
sensible values", with zero findings wanted if reachable; trivy stays for what the
release ships and govulncheck takes the source; and the report plumbing stays, because
the report is evidence of clean code and a thing a later intent can declare.

**What the tree says before any of it is written.** Measured on 2026-10-08 at `30b1dea`
with golangci-lint 2.14.0 and govulncheck 1.8.0: the five linters named report 51
findings, errcheck 21, gosec 27, staticcheck 3; of the gosec 27, four are in shipped
code and 23 in tests, where fixtures are written with ordinary permissions, subprocesses
are launched on purpose and Amazon's documented example key is the secret filter's own
test input. govulncheck reports no vulnerability, and its JSON carries the database's
own date, `db_last_modified`, which is the field trivy's report lacks and #11 is about.
Zero is reachable: four findings in shipped code with their reason written beside them,
tests outside gosec's scope, and the rest fixed.

**The one step of the plan that does not survive measurement.** depguard reads the
direct imports of the files it is pointed at, and A42 is a property of the closure:
`go list -deps ./internal/gates` is what the `verify` workflow runs, and it would catch
a gate importing a package that imports `net/http` where depguard would not. Moving the
grep into depguard would make the rule weaker under the name of giving it one home, so
the grep keeps the home it has and depguard is not enabled; the design says so against
the comment's step two.

<!-- xeno:section:scope -->
## Scope

This intent replaces semgrep with golangci-lint, gosec, staticcheck, errcheck and
govulncheck, as the analysis on the issue laid out and the maintainer approved, and
removes semgrep in the same change rather than after a period of running both.

- **The policy, in the repository.** `.golangci.yml` enables `govet`, `staticcheck`,
  `errcheck` and `gosec` and nothing else, excludes `_test.go` from gosec with the reason
  in the file, and uses the standard error-handling exclusions for errcheck so that an
  unchecked `fmt.Fprintln` to standard error is not a finding. Any finding fails the
  run, which is what `--error` did for semgrep and what the maintainer asked for.
- **The tree at zero.** The four gosec findings in shipped code are suppressed inline,
  each with its reason; the errcheck findings in tests are fixed by checking the error;
  the staticcheck findings are fixed. No behaviour changes.
- **Two workflows for one.** `lint.yml`, job `lint`, runs golangci-lint through its
  action pinned by sha with the version pinned beside it, writes the JSON report to
  `.xeno/local/scan/`, declares it in a manifest and uploads it, in the shape
  `semgrep.yml` had. `govulncheck.yml`, job `govulncheck`, installs the scanner at an
  exact version with `go install`, scans the source, writes the JSON report with the
  database date it already carries, and runs on the daily schedule trivy has, since a
  vulnerability is found after the code is written.
- **semgrep goes.** `.github/workflows/semgrep.yml` and `.semgrep/` are deleted; the
  pin table, the renovate configuration, `CONTRIBUTING.md`, the comments in
  `gitleaks.yml` and `xeno.yml` that name it, and the required checks in
  `scripts/github-settings.sh` move with it. The host's branch protection is updated
  from that script before the merge, because a required context that never reports
  blocks every merge.
- **The pin table and its test.** `docs/supply-chain.md` gains rows for the action, the
  linter, govulncheck and its database, loses semgrep's two, and the paragraph that
  contrasted semgrep's vendored rules with trivy's fetched database is rewritten for a
  policy file and a scanner; `supply_chain_test.go`'s exemptions move with the rows.
- **The register.** Rows in `docs/assumptions.md` for what the answers left to the
  implementation: the linter's threshold and what tests are exempt from, and the reading
  of "trivy for the container scan" as trivy unchanged.

What it does not do.

**It does not enable depguard.** The problem section says why: it reads direct imports,
A42 is about the closure, and the `go list -deps` step already holds the rule.

**It does not change trivy.** The maintainer keeps it for what the release ships, and
`trivy.yml` scans the built binary today; whether it should scan the image #282 built
instead is a question about trivy and not about semgrep.

**It does not attach a report to any phase.** The manifests exist so that a phase can
declare one; this intent keeps that plumbing and declares nothing, as every intent since
#51 has, and the residual risk says so.

**It does not add a dependency.** golangci-lint and govulncheck are pipeline tools with
rows in the pin table, not modules; `go.mod` is untouched.

**It does not touch the generated wrapper.** `xeno init` writes a pipeline for an
adopter, and what this repository lints itself with is this repository's decision, as
the maintainer's first answer says.

<!-- xeno:section:context-rationale -->
## Why this context

Read for this phase, and why.

- Issue #94 with its three comments, which hold the analysis, the mapping and the
  three answers; the body alone names three tools and decides nothing.
- `.github/workflows/semgrep.yml` and `.semgrep/golang.yaml`, the thing being replaced,
  for the shape every scan workflow here shares and the vendoring argument its header
  makes; `trivy.yml`, `gitleaks.yml` and `xeno.yml` for the same shape, the comments
  that name semgrep, and the A42 step the comment proposed to move.
- `docs/supply-chain.md` and `internal/model/supply_chain_test.go`, because every
  pin this adds or removes is a row the test holds to the tree in both directions, and
  the exempt list has to move with the rows.
- `scripts/github-settings.sh`, `renovate.json` and `CONTRIBUTING.md`, the three other
  places that name the semgrep check or the semgrep digest.
- `docs/assumptions.md` A42 and the rows nearest in shape, A100 to A103.
- The files the measured run named: the seven in shipped code and the fourteen tests
  with a finding, so that the scope's budget covers what the implementation changes.
- `docs/commands.md`, since a workflow that calls the binary is held to it.
- golangci-lint 2.14.0's configuration reference and gosec's rule list, read at the
  tool rather than from memory, for the exclusion syntax and the four rule ids.
