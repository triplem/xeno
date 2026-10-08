---
intent: github.com/triplem/xeno#94
phase: 04-verification
created: "2026-10-08T20:39:58Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 39f63b26b06d3f8f13811427f0f9fb9d3af44a88fbcb37a9bc414f6fe6c2575f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| criterion | how it was checked | result |
|---|---|---|
| 1. the policy is a file and fails on any finding | `.golangci.yml` read: `version: "2"`, `default: none`, three linters; `golangci-lint config verify` passes; `golangci-lint run ./...` exits 0 with 0 findings on the branch | met as to the file and the threshold; gosec is not in it, see the first deviation of P3 and the results below |
| 2. four accepted findings in shipped code, each with rule and reason, no more | `grep -rn '#nosec' --include='*.go' internal cmd` finds four lines, three `G204` in `git.go` and one `G122` in `plugin.go`, each with `--` and a reason; `grep -rn nolint` finds none; gosec with `-nosec-require-rules -nosec-require-justification` reports `Nosec: 4` | met, with the rules and files different from the criterion's four, which named a measurement P3 shows was wrong |
| 3. tests outside gosec, inside errcheck and staticcheck, their findings fixed | `-tests=false` on gosec's line; `golangci-lint run ./...` over the tests reports 0; the thirteen fixed lines are in the diff of `runner_test.go`, `predicates_test.go`, `index_test.go`, `sealed_test.go`, `shipped_test.go` | met, by a flag where the criterion said a rule |
| 4. the linter's workflow in the scanners' shape | `lint.yml` read against `trivy.yml` step by step: checkout, setup, make room, action by sha with `# v9.3.0` and `version: v2.14.0`, print, declare `kind: scan` `job: lint` `file: golangci-lint.json`, upload `scan-lint`, fail last; `TestTheScanWorkflowsWriteTheManifestThisExpects` holds the three strings | met |
| 5. the vulnerability scan's workflow in the same shape | `govulncheck.yml` read the same way: `go install …@v1.8.0`, `-format json` to `govulncheck.json`, a judging step, declare, upload `scan-govulncheck`, fail last, cron `29 6 * * *`; the manifest test holds its strings | met |
| 6. the judgement is the text mode's, written down | the `jq` of the judging step run on the branch's report, 0 vulnerabilities, and on the fixture module requiring `golang.org/x/text v0.3.0`: 5 vulnerabilities, 1 reachable, 1 with a fix, where `govulncheck` text mode exits 3 and JSON mode exits 0; `db_last_modified` printed in both | met |
| 7. semgrep gone from the tree and from every sentence outside the trail | the two files absent; `grep -rn semgrep` with `.xeno/` and `vendor/` excluded finds three lines: `.golangci.yml` and `gosec.yml` naming what they replaced, and A105 | met for the tree and for every file that named the check; three sentences name it as the thing replaced, by design |
| 8. the pin table held in both directions with the new rows | `go test ./internal/model/` green; rows for the action, `golangci-lint`, `gosec`, `govulncheck` and its database present; semgrep's two gone; the exempt map has the database row and not `semgrep's rules` | met, with a `gosec` row the criterion did not foresee |
| 9. the required checks name the new jobs | `scripts/github-settings.sh` lists `verify`, `audit`, `gitleaks`, `lint`, `gosec`, `govulncheck`, `trivy`; the host's protection is changed by running it before the merge, and the pull request records the `gh api` read-back | met in the tree; the host half is done at the merge and recorded on the pull request |
| 10. renovate sees the new pins | `renovate.json` parses; three regex managers, one per `version:` input or `go install` line; the description no longer names semgrep | met as to the file; whether the managers match is proved by the next scheduled run, see gaps |
| 11. suite and gates green, no behaviour change | `go test ./...` exit 0, `gofmt -l` empty, `go vet` clean, `xeno gate verify` 552 verdicts; the diff of shipped code is four `#nosec` comments, one selector and one `Fprintf` | met |
| 12. the register carries the two decisions | A104 and A105 in `docs/assumptions.md` | met |

<!-- xeno:section:results -->
## Results

Twelve criteria, twelve met, four of them differently from how they were written, and the
reason is one measurement.

    go test ./...                       ok, 21 packages, exit 0
    gofmt -l . | grep -v vendor/        nothing
    go vet ./...                        clean
    golangci-lint run ./...             0 issues, config verified, 2 cold runs identical
    gosec … ./...                       0 findings, 4 accepted, 43 files, 2 runs identical
    govulncheck -format json ./...      0 vulnerabilities, database 2026-10-07T14:10:51Z
    ./xeno gate verify                  verified 552 verdicts

**The measurement that moved the design.** The intake counted gosec's findings through
golangci-lint 2.14.0 once and found four in shipped code. In P3 the same configuration
was run three times on one commit: two from a cold cache, one warm, then twice more at
`concurrency: 1`. Every run reported sixteen gosec findings over shipped code and tests
together, and no two runs reported the same sixteen: `attach.go`, `cost.go` and
`propose.go` took turns carrying the two permission findings, and `hashing.go`,
`scaffold.go` and `index.go` the one `G304`. The gosec binary at `v2.29.0` on the same
tree reported fifty-eight findings over shipped code, and the same fifty-eight on a
second run. golangci-lint without gosec reported the same single errcheck line on two
cold runs. So the three linters stay in golangci-lint, gosec runs as the binary, and
criteria 1, 2 and 3 are met in the shape the deviations describe rather than the shape
P1 wrote.

**The fifty-eight, by rule.** Twenty `G304` and eight `G703`, a file read or written at a
path from a variable, which for a runner over a repository is every read and write;
twenty-five `G301`, `G302` and `G306`, directories at `0755` and files at `0644`, which
are the modes a git checkout has; three `G204`, `git` launched with the runner's own
arguments; one `G122`, a read inside the walk that hashes the plugin, where a file
changed between listing and reading changes the digest, which is what the digest is
for. After the configuration and the two exclusions, four remain and four are accepted
on their lines.

**What the judging step says on a tree with something to find.** On a scratch module
requiring `golang.org/x/text v0.3.0`: govulncheck text mode exits 3; JSON mode exits 0
and carries eight finding messages for five OSV ids, one reaching a function; the step's
`jq` prints the five with their fixed versions and `reachable` on the one, and its
count of reachable-with-a-fix is 1, so the job would fail. On this repository both
counts are 0.

**The report that was not written.** gosec with `-quiet` on the clean tree exits 0 and
writes no `-out` file. The workflow does not pass the flag, and was run as written on
the branch with the report present at 139 bytes and `"found": 0` both times.

**Timings, for #117.** `gate run` on P2 of this intent, on a scratch copy: 0.118 s.
`gate verify` over 552 verdicts: 36.4 s, one run; XENO-0278 measured 34 to 36 s over
547 the same day, so the figure is the tree's and has not moved with this change.

<!-- xeno:section:gaps -->
## Gaps

**The host half of criterion 9 is not in the tree.** The branch protection names the
required checks and this change renames them; the script is run against the host at the
merge and nothing here can prove it was. The pull request carries the `gh api` read-back
as the evidence, and a protection left naming `semgrep` would show itself at once, as a
pull request that waits forever for a check that no longer exists.

**The three renovate managers are proved by parsing, not by matching.** No
`renovate-config-validator` runs here, and the regexes follow the trivy manager's shape
line for line; whether they propose anything is known on the first Monday after the
merge, or on a `workflow_dispatch` of `renovate.yml` before it.

**The three new workflows have run locally and not on the runner.** Every command in
them was run on the branch as written and the reports exist; the action's cache, the
`ubuntu-latest` `jq`, and the `go install` from the runner's network are the parts a
local run does not exercise. The pull request's own checks are the first run.

**Nondeterminism is shown, not explained.** Five runs of golangci-lint's gosec gave
five sets, and the binary gave one set twice; why the integration differs is not known
here, and A104 says the two jobs can be one again on the day it reports the same set
twice. Two identical runs of the binary are evidence of stability and not proof.

**The fifty-eight were read by rule, not one by one.** The `G304` and `G703` exclusion
rests on the claim that every path this runner reads is computed from the tree or given
on the command line and none arrives from the network; that was checked against the
adapters, which read bodies and not paths, and against the twenty-eight findings' files,
and not against every `os.Open` in the tree.

**The report is evidence nobody declares.** The manifests are in the shape section 4
can attach and no phase of this intent attaches one, which is every intent since #51.
The maintainer's third answer, that the report is a thing a later intent learns from,
is met by the plumbing and not by a declaration.

**The real-host checks of this intent are the first positive case for #330.** `intent
start --for 94` read the label and the comment from the host and started; that is one
run, recorded in P0, and the milestone half is still unexercised, as XENO-0278's gaps
said.
