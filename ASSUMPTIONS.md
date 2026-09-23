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
| A7 | G-Learning counts a record as filled when it carries any key beyond the header fields, such as an observation or `no_finding: true` | "present and filled" is not defined further, but the **shape** is: section 10 fixes `learnings:` as a mapping with `category` from a closed set, `observation`, `proposal` and `target` | `internal/gates` | **open, now a known gap**: XENO-3's first learning record invented `observations:` with a free category and passed, so the gate says filled where it means non-empty. The shape belongs to G-Learning and therefore to WP14 |
| A8 | An evidence item with a `uri` is bound by its declared hash only | resolving a uri needs a network, which the gate path never has | `internal/gates` | approved |
| A9 | The run marker lives under `.xeno/local/runs/`, gitignored | v1 is single developer; the v2 delta moves it into the repository for a central service | `internal/runner` | approved |
| A10 | `context.lock.yaml` carries `predecessor_artifacts_hash` and `evidence_source` only | enough for sequence, freshness and staleness; the context profile and read files come with WP8 | `internal/runner` | approved |
| A11 | Exit codes: 0 done and not red, including provisional and overridden; 1 red or refused with a reason; 2 could not run | the plan's staircase, applied to every command | `cmd/xeno` | approved |
| A12 | The next phase starts only on a decided predecessor: green, approved or overridden start, red and provisional are refused | building on a red verdict moves the failure downstream; the process definition said the local verdict was advisory and now says it binds the sequence | `internal/runner` | **changed and approved**; process definition and plan updated |
| A13 | The pipeline artifact fetch is stood in for by a directory with `manifest.yaml` (`--evidence-from`) | the real fetch belongs to the GitLab adapter and needs a host; the attach logic is testable without one | `internal/evidence` | open, explained |
| A15 | `.gitattributes` fixes LF for every text file | the hash normalises CRLF, other tools do not; WP0 asks for fixed line endings without saying which | `.gitattributes` | open |
| A16 | Source files carry only `SPDX-License-Identifier`, no per file copyright line | WP0 names the identifier and nothing more; the copyright line lives in `NOTICE` | `*.go` | open |
| A17 | The copyright wording in `NOTICE` and the security contact are placeholders | both are for legal or the maintainer to fix, not for the build | `NOTICE`, `SECURITY.md` | open |
| A18 | The module path moves to the self managed GitLab instance once it exists, and that move is a package of its own rather than a side effect of A1 | the plan targets GitLab CE, and a Go module path is its host: the instance's own path replaces `github.com/triplem/xeno` in `go.mod` and in every import, and a private instance needs either a vanity import path or `GOPRIVATE` and the `go-import` meta tag, neither of which is free | `go.mod`, every import, `.github/workflows/*.yml`, `M0.md` | **open, scheduled**: to be done with the host move, not before |
| A19 | G-Schema requires `digest.md` and nothing else beyond `output.md`; `context.lock.yaml` stays with G-Freshness and `learning.yaml` with G-Learning, and `cost.yaml` is required by no gate | section 4 lists six files per phase, but two of them cannot be required today: `cost.yaml` may arrive after the gate ran and its writer is WP13, and a second gate reporting a file another one already names would give one cause two findings | `internal/gates` | open, explained; `context.lock.yaml` is unguarded at P0, where G-Freshness returns early and there is no predecessor to compare against |
| A20 | An intent that runs P0 only is the intended shape between M0 and WP8, and nothing marks it finished | `M0.md` runs one intake per work package and stops there deliberately, because G-Freshness is half implemented until WP8 (A6) and no phase beyond P0 can be judged honestly; `in-progress` is in any case what every merged intent carries, since the process definition has no `merged` status by design and says so | `.xeno/intents/*/intent.yaml`, `M0.md` | explained, intended, no action; the same honesty as `not-implemented` and `by-hand`, and the one cost is that these intents read as stalled to anyone who has not found `M0.md`, which is why the README now points at it |
| A21 | The release runs on semantic-release, brought in as a GitHub Action, with `.releaserc.json` in the repository | WP0 wants version, tag, changelog and release derived from the commit history, and a standard tool does that better than three shell scripts; as an action it is not a dependency of this project, so a Go repository with one vendored library grows no `package.json` and no Node tree | `.github/workflows/release.yml`, `.releaserc.json` | approved; `.releaserc.json` moves to GitLab unchanged through the common-ci-tasks semantic-release job, which is why it lives in the repository rather than in workflow inputs (A18) |
| A22 | How a version is derived is semantic-release's rule now, not ours | the previous script raised the minor for a breaking change below 1.0; semantic-release raises the major, so the first breaking change takes this project to 1.0.0 whether or not anyone means it as a statement about stability | `.releaserc.json` | **open, a changed behaviour rather than a kept one**: the difference is recorded so that the first major is not read as a decision nobody made |
| A23 | `CHANGELOG.md` is written back to `main` by `@semantic-release/changelog` and `@semantic-release/git` | WP0 wants the changelog derived from the history; the loop is closed as before, since a push made with `GITHUB_TOKEN` triggers no workflow, and the commit carries `[skip ci]` besides | `.releaserc.json` | approved; the commit is still never verified by CI, because it triggers nothing, and a protected `main` refuses it unless the token is on the bypass list |
| A24 | One bill of materials for all five binaries, generated by the `CycloneDX/gh-gomod-generate-sbom` action | build constraints decide module selection, so one document per target is the accurate form in general; this module has a single pure Go dependency and no platform dependent code, so all five were byte identical but for the `goos` and `goarch` in the purl | `.github/workflows/release.yml`, `SUPPLY-CHAIN.md` | **open, true under a condition**: the day a build constrained dependency enters the tree this goes back to one document per target, and nothing currently checks that condition |
| A25 | `carryForward` is the only place a decision is attached to a finding, and the rule that an external finding never keeps one binds exactly the checks that pass through it | the filter has to sit where the attachment happens, and making that one function is what turns a rule into a property; external gates do not exist yet, so nothing else currently produces a check at all | `internal/gates` | **open, a constraint on WP4**: whoever adds external gates routes their checks through this function or the rule stops holding, and a test that passes would say otherwise |
| A26 | Three observations from XENO-3's P0 are recorded here rather than in its learning record: the first artifact to carry a real `model`, `tool` and `tool_version` instead of `none`, `manual` and `0`; a runner reporting `0.1.0-dev` while v0.1.1 was released, which nothing catches because G-Supply is not implemented and no `project.yaml` pins a version; and Appendix B defining `artifacts_hash` and the finding id to the byte while `context_hash`, `secrets_hash`, `rules_hash` and `strings_hash` are named in section 5 without a definition | a learning in this process is narrow on purpose, since it "produces a merge request against the rule set" and therefore has to be expressible as a rule or template change; an observation that proposes neither is a finding and belongs where findings are kept | `ASSUMPTIONS.md` | open, explained; the third one wants a decision, since a hash that is written but not defined is checkable by nobody |
| A27 | The default branch is not protected and cannot be on this host; the guarantee is taken as given until the GitLab move | first step 2 of section 4 calls it "the setting the whole tool claims to rest on", and a ruleset exists on the repository but is not enforced: branch protection and rulesets need a paid plan for a private repository, and the API answers `403 Upgrade to GitHub Pro or make this repository public` to every question about either. On the self managed Community Edition this project is aimed at, a protected branch with a required pipeline is free | the repository, `.github/workflows/` | **open, a property of the interim host**: until the move (A18) the sequence guarantee rests on discipline rather than on a setting, and a reader would otherwise take the documentation at its word |
| A28 | Every action is pinned to a commit sha and every tool to an exact version, taken from the run that produced v0.4.0 rather than from what is newest | a tag can be repointed by whoever owns it and `cycjimmy/semantic-release-action@v4` is a branch, so without this two runs a month apart can use different tooling and nothing records which; what is wanted is the set that demonstrably works, not the set that is current | `.github/workflows/`, `SUPPLY-CHAIN.md` | approved; **this repaired a regression rather than adding a property**: `scripts/sbom.sh` pinned `cyclonedx-gomod` to v1.9.0 and the action that replaced it was given `version: v1`, and neither the commit that made the change nor the two assumptions it rewrote said so |
| A14 | A question has two to four options plus exactly one free entry, or `no_options: true` | the plan's wording, made countable | `internal/gates` | open, explained |

## Decisions

Not assumptions, and kept apart from them for the reason the process definition gives:
an assumption stands in for missing knowledge and can turn out wrong, a decision is a
choice between workable paths and can only be regretted. Decisions belong in the
`decisions` section of the phase they were made in, sealed with it. These two were taken
after the P0 of XENO-3 was sealed, and the phase that would carry them is P1, which
cannot run before WP8 closes the G-Freshness gap (A6). So they are recorded here, by
hand, like the rest of the work in this window, and Q-1 and Q-2 stay open in the trail
because nothing that could resolve them is allowed to run yet.

| # | Question | Decision | By |
|---|---|---|---|
| D-1 | XENO-3 Q-1, where the WP1 fixtures live | `testdata` per package, read by the tests of that package. The corpus for G-Schema is `internal/gates/testdata/`: a well formed artifact of each kind, and `required-fields.yaml` written out of section 5 so that the tests assert behaviour against the specification rather than against the constants in `gates.go` | the maintainer |
| D-2 | XENO-3 Q-2, how the schema version is recorded | `schema_version` in the common field set, `major.minor`, written by the runner. Absent means the schema that predates the field and is read rather than failed; present has to be well formed. Process definition changed first, in its own commit | the maintainer |
| D-3 | XENO-3, where the cross platform assertion of the hashes belongs | WP17, not WP1. `Normalise`, `TestCRLFIsNormalised` and `TestPathIsPartOfTheHash` hold in this package; running them on a second operating system needs a CI matrix, which the implementation plan already assigns to WP17. The criterion was written into the wrong issue at intake and is handed over rather than dropped | the maintainer |
| D-4 | XENO-5, where a learning goes | into the rule set through `learning.yaml` and a merge request, per section 10, and never into `CLAUDE.md`. That file carries what first step 3 names and stops there: it is sent with every request of every session, so its length is a standing cost, and a learning written into it would take effect unreviewed, which is what section 10 exists to prevent | the maintainer |

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
Commits, `CHANGELOG.md` written back to the branch, five platform binaries with the
version injected through ldflags, a CycloneDX bill of materials per target and
`SHA256SUMS`. Signing waits for publication, as WP0 says. `LICENSE` is to be taken unchanged from the Apache Software
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
