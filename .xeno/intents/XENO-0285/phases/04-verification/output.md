---
intent: github.com/triplem/xeno#95
phase: 04-verification
created: "2026-10-09T16:54:15Z"
schema_version: "1.0"
runner_version: dev+d2bc411.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0396db2a4b80d486cf5aea37057af4f55e49d852f0ecd97528d678b4e3c57709
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Seven criteria. Four are checked by something that runs in CI, three by a check run here and
recorded as evidence, and the table says which is which because the difference is what a
reader of a sealed phase cannot otherwise tell.

| criterion | checked by | runs in CI |
|---|---|---|
| 1, every movable pin is reached | the manager run below, read against `docs/supply-chain.md` row by row | no |
| 2, every manager resolves to the tree's value | the manager run below, nine entries | no |
| 3, the reference names the tag the page names | `sed -n 43p Dockerfile` against the page's row | no |
| 4, the page is still held against the tree | `go test ./internal/model/` — `TestEveryPinInTheTreeIsOnTheSupplyChainPage` and `TestEveryRowOfTheSupplyChainPageIsAPinTheTreeCarries` | yes |
| 5, no footer points at a closed issue | `commitBody` absent from `renovate.json` | no |
| 6, the decided shape has not moved | the keys compared against `HEAD`'s copy of the file | no |
| 7, the gate suite passes | `go build`, `go test ./...`, `gofmt -l`, `go vet`, `xeno gate verify` | yes |

**Why criterion 2 is not a test.** The regex dialect is renovate's: `managerFilePatterns` are
`/…/`-delimited and `matchStrings` use JavaScript named groups. A Go test would compile a
translation of these strings, and a translation that passes while the original fails is the
undetected gap this intent exists to close, so it would check the wrong thing confidently.
There is no renovate binary in this repository and vendoring one is a second dependency, which
this project treats as a decision rather than a step. So the strings are run as renovate's own,
by a script, and its output is this phase's evidence.

**What that leaves unchecked after this intent.** Nothing re-runs that script. A workflow line
rewritten next month leaves its manager silently matching nothing, which is the same defect
this intent found and fixed by hand, and the next person would have to think to look. That is a
gap and it is recorded as one below rather than absorbed.

<!-- xeno:section:results -->
## Results

All seven met. The figures below are from the tree as this phase found it, with the change
applied — not from a tree the change was reverted out of, which is criterion 7's wording and
the reason it is worded that way.

**Criterion 2, the one that decides whether this intent did anything.** Nine managers, nine
resolved values, the six that already existed run beside the three that are new:

| manager | file | resolved | in the tree |
|---|---|---|---|
| semantic-release | `release.yml` | 25.0.9 | 25.0.9 |
| CycloneDX/cyclonedx-gomod | `release.yml` | v1.12.0 | v1.12.0 |
| aquasecurity/trivy | `trivy.yml` | v0.74.0 | v0.74.0 |
| golangci/golangci-lint | `lint.yml` | v2.14.0 | v2.14.0 |
| golang.org/x/vuln | `govulncheck.yml` | v1.8.0 | v1.8.0 |
| github.com/securego/gosec/v2 | `gosec.yml` | v2.29.1-0.20261009120814-7b1b5cebe007 | the same |
| **gitleaks/gitleaks** | `gitleaks.yml` | **8.30.1** | 8.30.1 |
| **zensical** | `docs.yml` | **0.0.69** | 0.0.69 |
| **docker.io/library/debian** | `Dockerfile` | **13-slim**, digest `sha256:a29215f6…` | the same |

The six unchanged rows are the control. Before the change the same script printed `NO MATCH`
for the last three while printing these six, which is what makes the absence a finding rather
than a script that never worked — the rule about negative results, applied to the thing this
intent is about.

**Criterion 1.** The page's thirty pipeline rows and one module row, classified individually:
13 actions by sha and 1 module and the `go` directive reached by the two ordinary managers, 6
tool versions by the managers that existed, 3 by the managers added here. Seven remain
unreached and each is explained by the page: four cannot be pinned at all, two are a toolchain
major this project chose, and gitleaks' rules are a hand translation into `secrets.yaml` that
nothing fetches.

**Criterion 3.** `Dockerfile:43` reads
`ARG BASE=docker.io/library/debian:13-slim@sha256:a29215f6…`. The page's row calls that pin
`debian` 13-slim. They now agree; before this change the page named a tag the tree did not
carry.

**Criterion 4.** `go test ./internal/model/` passes, including both directions of the
page-against-tree check. The digest the test extracts is unchanged: its pattern's first group
stops at `:` so it now captures `13-slim` instead of the image name, and the second group,
which is the only one stored, is byte for byte what it was.

**Criterion 5.** `commitBody` is in `HEAD`'s copy of `renovate.json` and absent from this one.

**Criterion 6.** Every other key compared against `HEAD`: `$schema`, `automerge`,
`dependencyDashboard`, `enabledManagers`, `extends`, `packageRules`, `pinDigests`,
`postUpdateOptions`, `postUpgradeTasks`, `prConcurrentLimit`, `schedule`, `semanticCommits` and
`vulnerabilityAlerts` are identical, and the six existing `customManagers` are byte-identical
as a list. One key moved and it is `commitBody`. So `automerge` is still false, the dashboard
is still on beside pull requests, `go mod vendor` still runs, and nothing gained a sign-off.

**Criterion 7.** `go build` ok. `go test ./...` exit 0, no `FAIL`. `gofmt -l` excluding
`vendor/` prints nothing. `go vet ./...` silent. `xeno gate verify` exit 0.

<!-- xeno:section:gaps -->
## Gaps

**Nothing re-runs the manager check.** This intent found three pins with no manager by running
every `customManager` regex against the tree and reading what each resolved to. Nothing in the
repository does that again. A workflow line rewritten next month — `GITLEAKS_VERSION` moved
into a `with:` block, the `go install` line split across two — leaves its manager silently
matching nothing, and the configuration still looks complete, because a manager that resolves
nothing is indistinguishable from one that was never run. That is the defect this intent fixed
by hand, and it can recur the same way.

Three things were weighed and none is taken here. A Go test would check a translation of the
regexes rather than the regexes, for the reason the test mapping gives. A committed script,
beside `scripts/plugin-digest.go` and `scripts/go-symbols.go`, would be reproducible and is the
most likely answer; it was not taken because it is a tool this intent's design did not propose,
and adding one while nobody is looking is how a scope stops meaning anything. Vendoring
renovate itself is a second dependency, which this project treats as a decision.

So it is written down rather than absorbed, which is the third standing rule's instruction for
something needed that belongs to no package. It is not a blocker for #95: the three pins are
reached today and the page and the tree agree today, which is what this intent claimed.

**Nothing has run renovate.** Every statement here is about what the configuration says, not
about what renovate does with it. The regexes are run as renovate's own dialect by a script
that is not renovate, so a behaviour the script and renovate disagree about would not show up.
The datasources in particular are unexercised: that `pypi` answers for `zensical` and that
`github-releases` plus `extractVersionTemplate` gives `8.30.1` from tag `v8.30.1` are read off
renovate's documented behaviour, not observed. #350 carries what to check on the first real
run, and this is the half of that list which only a real run can settle.
