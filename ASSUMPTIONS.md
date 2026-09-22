# Assumptions made while building the core

Before M0, Xeno does not govern its own construction and the record is kept by hand
(implementation plan, section 4). This is that record. Every entry is a place where the
process definition or the implementation plan says what, and the how had to be chosen.
None is confirmed yet; each wants a yes, a no or a replacement from a person.

| # | Assumption | Why | Where | State |
|---|---|---|---|---|
| A1 | Module path `github.com/triplem/xeno`, matching where the repository currently lives | a placeholder namespace breaks `go get` and every import path reads as a stand-in; the plan leaves the namespace unrecorded, so the honest choice is the host the code is actually on | `go.mod`, every import | **approved for now, to be redone at the GitLab move**, see A18 |
| A2 | `gopkg.in/yaml.v3`, vendored | the only dependency; vendored matches the plan's rule of vendoring rather than fetching | `vendor/` | approved |
| A3 | Open questions, decisions and evidence declarations are structured frontmatter lists in `output.md` (`open_questions`, `decisions`, `evidence`); the rendered sections are derived from them | the documents fix stable keys and sections, not where the structured form lives, and a gate cannot parse prose | `internal/model` | open, explained; once approved it belongs in the process definition as a format rule |
| A4 | Gate result `pending` and status `provisional`, loudness red, overridden, provisional, approved, green | the evidence section decides a provisional state; the `gate.yaml` schema in section 5 is corrected alongside | `internal/gates` | approved |
| A5 | A gate is omitted from `gate.yaml` before its "From" phase rather than listed as `not-applicable` | the spec fixes the "From" column but not whether an inapplicable gate appears; both readings fit it | `internal/gates` | open |
| A6 | G-Freshness implements its first half only, the context hash against the predecessor | the second half needs the files a phase read, which `context.lock.yaml` records only from WP8 | `internal/gates` | **known gap, accepted as temporary**: no schema change; to be closed with WP8 and before any phase beyond P0 runs in Xeno's own process, since until then a pass on this gate says more than was checked |
| A7 | G-Learning counts a record as filled when it carries any key beyond the header fields, such as an observation or `no_finding: true` | "present and filled" is not defined further | `internal/gates` | open, explained |
| A8 | An evidence item with a `uri` is bound by its declared hash only | resolving a uri needs a network, which the gate path never has | `internal/gates` | approved |
| A9 | The run marker lives under `.xeno/local/runs/`, gitignored | v1 is single developer; the v2 delta moves it into the repository for a central service | `internal/runner` | approved |
| A10 | `context.lock.yaml` carries `predecessor_artifacts_hash` and `evidence_source` only | enough for sequence, freshness and staleness; the context profile and read files come with WP8 | `internal/runner` | approved |
| A11 | Exit codes: 0 done and not red, including provisional and overridden; 1 red or refused with a reason; 2 could not run | the plan's staircase, applied to every command | `cmd/xeno` | approved |
| A12 | The next phase starts only on a decided predecessor: green, approved or overridden start, red and provisional are refused | building on a red verdict moves the failure downstream; the process definition said the local verdict was advisory and now says it binds the sequence | `internal/runner` | **changed and approved**; process definition and plan updated |
| A13 | The pipeline artifact fetch is stood in for by a directory with `manifest.yaml` (`--evidence-from`) | the real fetch belongs to the GitLab adapter and needs a host; the attach logic is testable without one | `internal/evidence` | open, explained |
| A15 | `.gitattributes` fixes LF for every text file | the hash normalises CRLF, other tools do not; WP0 asks for fixed line endings without saying which | `.gitattributes` | open |
| A16 | Source files carry only `SPDX-License-Identifier`, no per file copyright line | WP0 names the identifier and nothing more; the copyright line lives in `NOTICE` | `*.go` | open |
| A17 | The copyright wording in `NOTICE` and the security contact are placeholders | both are for legal or the maintainer to fix, not for the build | `NOTICE`, `SECURITY.md` | open |
| A18 | The module path moves to the self managed GitLab instance once it exists, and that move is a package of its own rather than a side effect of A1 | the plan targets GitLab CE, and a Go module path is its host: the instance's own path replaces `github.com/triplem/xeno` in `go.mod` and in every import, and a private instance needs either a vanity import path or `GOPRIVATE` and the `go-import` meta tag, neither of which is free | `go.mod`, every import, `.github/workflows/xeno.yml`, `M0.md` | **open, scheduled**: to be done with the host move, not before |
| A19 | G-Schema requires `digest.md` and nothing else beyond `output.md`; `context.lock.yaml` stays with G-Freshness and `learning.yaml` with G-Learning, and `cost.yaml` is required by no gate | section 4 lists six files per phase, but two of them cannot be required today: `cost.yaml` may arrive after the gate ran and its writer is WP13, and a second gate reporting a file another one already names would give one cause two findings | `internal/gates` | open, explained; `context.lock.yaml` is unguarded at P0, where G-Freshness returns early and there is no predecessor to compare against |
| A20 | An intent that runs P0 only has no defined way to close | `intent close` writes the closing record for an *abandoned* intent, and a merged one meets G-Complete in P5; an intent that by design stops after intake is neither, so `status: in-progress` is the only honest value and stays | `.xeno/intents/*/intent.yaml` | **open, a finding about the plan rather than about the code**: M0.md asks every work package to run its intake through Xeno, and the process definition offers no closing path for that shape |
| A21 | Release automation, SBOM and checksums are built on GitHub Actions now rather than waiting for the GitLab move | WP0's "Done when" is not met without them, and a package left short because its final host does not exist yet stays short indefinitely; the logic sits in `scripts/`, so the move replaces the workflow that calls them and not the derivation itself | `.github/workflows/release.yml`, `scripts/` | approved; the workflow is thrown away at the move (A18), the three scripts are not |
| A22 | Below 1.0.0 a breaking change raises the minor, not the major | the usual reading of semver while a project is unstable; WP0 names Conventional Commits as the source and not which rule applies under 1.0 | `scripts/version.sh` | open |
| A23 | The release notes go into the release, and no `CHANGELOG.md` is committed | WP0 names GitLab's changelog API, which commits the file; a pipeline that commits into the repository it just verified starts the next pipeline doing the same, which the process definition names as the reason CI writes nothing | `scripts/changelog.sh`, `.github/workflows/release.yml` | open, explained; the GitLab move may switch to the API and accept the commit |
| A24 | The SBOM is derived from the built binary through `go version -m`, not from a bill of materials tool | what a release ships is what the binary contains, and reading it there needs no network, no `go.sum` and no fetched toolchain, which is what WP0 asks of an instance whose runners have no route out | `scripts/sbom.sh` | open, explained; swapping in `cyclonedx-gomod` is one line if a standard tool is wanted |
| A14 | A question has two to four options plus exactly one free entry, or `no_options: true` | the plan's wording, made countable | `internal/gates` | open, explained |

## Built

Hashing per Appendix B, byte exact and checked against a `sha256sum` pipeline. Finding
ids. The runner commands `phase start`, `phase finish`, `gate run`, `evidence attach`,
`intent status`. Sequence enforcement, the run marker, computed status and staleness.
Evidence pulled from the pipeline into `evidence/attached.yaml` with provisional
verdicts. Decisions carried forward by finding id.

Gates implemented: G-Schema, G-Trace, G-Assumptions, G-Questions, G-Learning,
G-Freshness (partial, A6), G-Evidence, G-Build.
Gates written as `not-implemented`: G-Supply, G-Secret, G-Test, G-Rules, G-Policy,
G-Complete.

WP0 in part: `.gitattributes`, SPDX identifiers, `NOTICE`, `CONTRIBUTING.md` with the DCO
procedure, `SECURITY.md`, the module path (A1), the format and vet jobs the package asks
for on day one, and the release pipeline: version and notes derived from Conventional
Commits, five platform binaries with the version injected through ldflags, a CycloneDX
bill of materials and `SHA256SUMS`. Signing waits for publication, as WP0 says. `LICENSE` is to be taken unchanged from the Apache Software
Foundation rather than transcribed. SBOM, checksums and signing wait for a release.
Where the sign-off is enforced is open question Q-1 of intent XENO-2, now recorded in
that intent's `output.md` rather than only here. Release automation, SBOM and checksums are
built (A21). What WP0 still owes is the container and package registry side, which needs
the instance.

## Not built

Everything that needs a host, a harness or a network: the GitLab adapter, the CI
wrapper, `enforcement check`, the MCP server, hooks, the plugin, the documentation site.
Templates, rules, the secret filter and the digest writer. `xeno init`, approve and
override, intent close. Five of the plan's verification points stand before those parts.
