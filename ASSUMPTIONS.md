# Assumptions made while building the core

Before M0, Xeno does not govern its own construction and the record is kept by hand
(implementation plan, section 4). This is that record. Every entry is a place where the
process definition or the implementation plan says what, and the how had to be chosen.
None is confirmed yet; each wants a yes, a no or a replacement from a person.

| # | Assumption | Why | Where | State |
|---|---|---|---|---|
| A1 | Module path `example.com/xeno` | the namespace is deliberately unrecorded in the plan; renaming is one line in `go.mod` and a search | `go.mod` | open until the namespace is fixed |
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

## Not built

Everything that needs a host, a harness or a network: the GitLab adapter, the CI
wrapper, `enforcement check`, the MCP server, hooks, the plugin, the documentation site.
Templates, rules, the secret filter and the digest writer. `xeno init`, approve and
override, intent close. Five of the plan's verification points stand before those parts.
