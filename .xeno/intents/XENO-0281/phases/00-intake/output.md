---
intent: github.com/triplem/xeno#343
phase: 00-intake
created: "2026-10-09T13:18:48Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cc24daf6fe3f9d889ef74dadd4e9ef63e76786005760b7c358945203d4255ab1
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

Approved by @triplem on 2026-10-09T13:14:39Z: like described

> **github.com/triplem/xeno#343** — gosec cannot read Go 1.27.2's export data, and every merge is red
>
> Found on 2026-10-09 by the first `gosec` run after Go 1.27.2 was released. The run on
> `main` at `2af2c28`, the merge of #341, is red, and so is every pull request since,
> #342 first; nothing in the tree changed between the green run and the red one.
>
> ## What happens
>
> `actions/setup-go` resolves the `go 1.27` directive to the newest patch, which was
> 1.27.1 yesterday and is 1.27.2 today. gosec 2.29.0 is built with `golang.org/x/tools`
> 0.49.0, whose export data reader stops at version 4, and 1.27.2's compiler writes
> version 5. Every package the scanner loads then has "type errors", every import fails
> with *export data version 5 is greater than maximum supported version 4*, and the SSA
> analysis is skipped for all of them. The job fails on the type errors rather than on a
> finding, which is the right way round, and the manifest says `fail`.
>
> `lint` and `govulncheck` are green under the same toolchain: golangci-lint 2.14.0 ships
> its own reader and `govulncheck` 1.8.0 was built against a newer `x/tools`.
>
> ## What fixes it
>
> gosec's main branch moved to `x/tools` 0.51.0 and Go 1.27.2 in commit `7b1b5cebe007`,
> at 12:08 UTC today, and there is no release carrying it yet. `go install` takes a
> pseudo-version for a commit, so the pin in `gosec.yml` and the row in
> `docs/supply-chain.md` move from `v2.29.0` to that commit's pseudo-version, which the pin
> test reads the same way; when 2.30.0 is released the pin moves again, which renovate's
> regex manager proposes. Run locally at that commit against this tree: the same four
> accepted findings, zero issues, and no type errors.
>
> ## What this is not
>
> Not a reason to pin Go to a patch. D-5 pins the toolchain by major so that a patch
> reaches the pipeline without anybody editing a file, and this is that happening: the
> patch arrived and a tool behind it was found wanting. The tool moves, not the
> toolchain.
>
> ## Done when
>
> `gosec` is green on `main` and on #342 with the scanner pinned to a commit that reads
> 1.27.2's export data, the pin table and renovate agree with the workflow, and A104
> records that the scanner is pinned to a commit until its next release.
>

The issue as it stood at 2026-10-09T13:18:48Z, read by `xeno phase start` and quoted rather than summarised. What this phase concludes about it belongs below.

The issue as `xeno phase start` read it, approved by the maintainer within the hour of
being filed, because it blocks every merge. The analysis in it was made by reading the
failed run's log, where every import of every package reports *export data version 5 is
greater than maximum supported version 4*, and the passing run of the day before, which
ran under Go 1.27.1 where this one ran under 1.27.2; and by installing gosec at
`7b1b5cebe007`, the commit that moves it to `golang.org/x/tools` 0.51.0, and running it
against this tree: four accepted findings, zero issues, no type errors, the same answer
2.29.0 gave under 1.27.1.

This intent's key is XENO-0281 and not the derived XENO-0280, because XENO-0280 is live
on the unmerged branch of #342 and the sequence counts from what this branch can see;
the derived directory was removed before anything was written into it.

<!-- xeno:section:scope -->
## Scope

This intent moves gosec's pin from the release that cannot read Go 1.27.2's export data
to the commit that can, and nothing else.

- **The pin.** `go install github.com/securego/gosec/v2/cmd/gosec@v2.29.1-0.20261009120814-7b1b5cebe007`
  in `.github/workflows/gosec.yml`, the pseudo-version of commit `7b1b5cebe007`, with a
  comment saying why a commit and when it moves back to a release.
- **The row.** `docs/supply-chain.md`'s gosec row names the pseudo-version and says the
  same in a clause; the pin test reads `v2.29.1` out of both and holds them together.
- **The register.** A104's status cell records that the scanner is pinned to a commit
  until 2.30.0, and why.

What it does not do.

**It does not pin Go to a patch.** D-5 pins the toolchain by major on purpose, and this
is a tool behind the toolchain being found wanting, not the toolchain.

**It does not touch renovate.** The regex manager for the `go install` line matches the
pseudo-version as it matched the version, and the Go datasource proposes 2.30.0 when it
exists.

**It does not change what gosec checks.** Policy, exclusions, modes and the four
accepted findings are as #94 left them, and the run at the new commit says so.

<!-- xeno:section:context-rationale -->
## Why this context

- Issue #343 and the two runs it compares, `37934485207` red under 1.27.2 and
  `37841461938` green under 1.27.1, read from their logs.
- `.github/workflows/gosec.yml`, the line that moves and the header that explains the
  job.
- `docs/supply-chain.md` and `internal/model/supply_chain_test.go`, because the row
  moves with the pin and the test holds them together by the `v2.29.1` shape.
- `renovate.json`, to confirm the manager's regex matches a pseudo-version and nothing
  there has to move.
- `docs/assumptions.md` A104, the row this amends.
- gosec's `go.mod` at `master` and the proxy's `.info` for the commit, for the
  pseudo-version and the `x/tools` version it carries.
