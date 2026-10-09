---
intent: github.com/triplem/xeno#94
phase: 03-implementation
created: "2026-10-08T20:38:55Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ea885d39824a3553cf939a4bcde8a885b923ee5be785d9f5d211350b04472af8
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

Twenty-four files, by concern. Five new, two deleted, seventeen changed; 125 lines in
and 61 out across the changed ones, 57 and 42 of those under `internal/` and `cmd/`.

**The policy: `.golangci.yml` and `.gosec.json`.** The first enables `govet`,
`staticcheck` and `errcheck` with `default: none`, the `std-error-handling` preset and
`warn-unused`, and its header says why gosec is not in it. The second is gosec's
configuration: the three file-mode rules set to the modes a git checkout has, `0755`
for a directory, `0644` for a file and a chmod, with the reasoning in a `description`
key the tool ignores, as `renovate.json` does. Both verified by their tools:
`golangci-lint config verify` passes, and gosec reads the file.

**Three workflows for one: `lint.yml`, `gosec.yml`, `govulncheck.yml`.** Each in
`trivy.yml`'s eight-step shape with the shared reasoning cited rather than repeated.
`lint` runs the publisher's action pinned by sha with `version: v2.14.0` and two
outputs, text to the log and JSON to `.xeno/local/scan/golangci-lint.json`. `gosec`
installs `v2.29.0` with `go install`, excludes `G304` and `G703` on the command line
with the reason in the header, scans with `-tests=false`, requires a rule and a
justification on every `#nosec`, and writes `gosec.json` with the text to the log; the
`-quiet` flag is deliberately absent, because with it a clean scan writes no report and
the manifest step would then fail on a tree with nothing wrong. `govulncheck` installs
`v1.8.0`, scans with `-format json` into `govulncheck.json`, and a step of its own
reads the stream with `jq -s`, prints the database's date and every finding with its
fixed version and reachability, and fails on a reachable one with a fix; it runs on a
daily cron at `29 6`, between trivy's and audit's. Job ids `lint`, `gosec`,
`govulncheck`.

**Four acceptances in shipped code, one simplification, one selector.** `#nosec G204`
with its reason on the three `git` invocations in `internal/git/git.go`, and `#nosec
G122` on the read inside the walk that hashes the plugin in `internal/plugin/plugin.go`.
`internal/cost/cost.go` drops an embedded-field selector staticcheck named, and
`internal/learning/propose.go` writes a header with `fmt.Fprintf` instead of
`WriteString(Sprintf(…))`. No other line of shipped code changes.

**Twelve lines of tests check what they ignored.** `runner_test.go`'s fixtures pass
their `WriteFile`, `Remove` and `Rename` through `f.must`; `predicates_test.go` fails
on a `git` command that did not run; `index_test.go` and `sealed_test.go` check the
`Close`; `shipped_test.go` loses a double negation. The manifest-shape test gains the
three new workflows as rows beside `trivy.yml`, and the evidence fixtures that used
`semgrep` as a job name now say `lint` with `golangci-lint run` as what produced it.

**semgrep is gone.** `.github/workflows/semgrep.yml`, 115 lines, and
`.semgrep/golang.yaml`, 2,443 lines, deleted. `CONTRIBUTING.md` names the seven checks
and says three run on a schedule; `gitleaks.yml`'s header contrasts its exit code with
`trivy.yaml`'s instead of with semgrep's flag; `xeno.yml` cites `trivy.yml` for the
reasoning it cited `semgrep.yml` for; `scripts/github-settings.sh` requires `verify`,
`audit`, `gitleaks`, `lint`, `gosec`, `govulncheck`, `trivy`.

**The pin table and its test.** Two rows out, five in: the action with sha and tag, the
three tool versions, and govulncheck's database as a row that cannot be pinned. The
paragraph about semgrep's vendored rules is replaced by one about the two policy files
and one about govulncheck's database carrying its own date; the test's exempt map loses
`semgrep's rules`, gains the database row with its reason, and three of its comments
that used semgrep as the example of a digest pin now describe the shape without it.
`renovate.json` gains three regex managers, for the action's `version:` input and the
two `go install` lines, and its description no longer says semgrep is unmanaged.

**The register.** A104 and A105 in `docs/assumptions.md`, the first rewritten once
during this phase for the reason the deviations give.

**Measured on the branch.** `go test ./...` green, `gofmt` and `go vet` clean,
`golangci-lint run` 0 findings, gosec 0 findings and 4 accepted over 43 files twice,
govulncheck 0 vulnerabilities against a database last modified 2026-10-07, `xeno gate
verify` 552 verdicts.

<!-- xeno:section:deviations -->
## Deviations from the design

Three, each named against the design's `decisions`, and the first is the one that
matters.

**gosec runs as its own binary, not through golangci-lint.** The design's first
decision put all four linters in `.golangci.yml`. Implementing it, three runs of the
committed configuration on one commit of this tree reported three different sets of
gosec findings, sixteen each: `attach.go` in one, `cost.go` in another, `propose.go` in
the third, `hashing.go` against `scaffold.go` against `index.go` for the one `G304`.
With the cache cleaned, with it warm, and at `concurrency: 1`, the sets still differed.
The gosec binary at `v2.29.0` reported fifty-eight findings on the same tree, the same
fifty-eight twice, where golangci-lint's reported eleven. A required check that moves
between two runs of one tree would be the opposite of what this project asks of a
verdict, so gosec moved to its own workflow and policy file, and `.golangci.yml` keeps
the three linters whose findings were identical across two cold runs. The required
check count goes from two new to three.

**The acceptances are two rules configured, two excluded and four inline, not four
inline.** The design's second decision accepted four findings on their lines and
rejected a global exclusion. That decision rested on the first measurement, which the
first deviation shows was wrong: the binary found twenty-five permission findings and
twenty path findings in shipped code. Writing forty-five `#nosec` lines for
constructions that are what a runner over a repository is would be noise, and the
design's own reason against a global exclusion, that the next occurrence should meet a
finding, is answered differently for each class: the file modes are configured rather
than excluded, so anything looser than a git checkout's still fires; `G304` and `G703`
are excluded with the reason in the workflow, because every file this runner reads is at
a computed path and nothing arrives from the network; and the three `G204` and one
`G122` are accepted inline, with `-nosec-require-rules` and
`-nosec-require-justification` making the rule and the reason mandatory on every such
line from now on. The `nolint` comments the first draft put on `gitlab.go` and
`rules.go` are gone: `G704` was never reported by the binary, and `G304` is excluded.

**Tests are outside gosec by a flag, not by an exclusion rule.** The design's third
decision wrote the rule into `.golangci.yml`; with gosec elsewhere it is `-tests=false`
on the command line, with the same reason in the workflow's comment. The eleven errcheck
and staticcheck findings in tests are fixed as decided, plus two more `WriteFile` lines
the first run had not reported.

Nothing else departs. The two workflows the design described are as decided; the
judgement step, the daily schedule, the pinned `go install`, the removal in one change,
the rows, the register rows and the renovate managers are as written, with a third
manager for gosec's line.
