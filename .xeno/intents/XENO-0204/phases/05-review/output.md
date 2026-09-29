---
intent: github.com/triplem/xeno#128
phase: 05-review
created: "2026-09-29T09:33:24Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+38430d4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9175b32acba0e03fba570143fde62296cb1fea2b8da85ff2eccdb954f938a29b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**Nothing under `docs/` changed.** Section 4 says G-Secret reads the effective filter,
and this is not that gate; nothing here touches what the documents define.

**Nothing was invented.** A workflow in the shape the other three use, a configuration
file for a pinned tool, and two lines of a table.

**Evidence rather than a gate.** No gate reads the report, the threshold lives in the
tool, and what stops a merge is the job failing, which is `trivy.yml`'s own sentence
about itself.

**The tool is pinned to an artifact and not to a name.** Fetched by version and checked
with `sha256sum -c` before it is executed, and the checksum is in `SUPPLY-CHAIN.md`.

**The configuration names shapes and disables nothing.** One path and four regexes, each
a thing this repository contains. `generic-api-key` produced thirty-nine of forty
findings and is still in force.

**It lands clean and it still fails on a real token.** Both proved against the same
extracted tree, and the token is not written into a record that the job scans.

**No secret reaches a log**, on any path, including the failure path.

**The order of the work is admitted.** The unconfigured scan's output was the intake's
subject, so the measurement preceded the phases, as it did in XENO-0203.

**What is still outstanding is named rather than implied.** The workflow has never run;
the first pull request is its first evidence, and the verification says so instead of
letting a transcript stand in.

<!-- xeno:section:release-notes -->
## Release notes

A secret scan runs on every pull request, on every push to `main` and weekly. It reads
the source tree rather than the built binary, with gitleaks 8.30.1, fetched by version
and verified by checksum before it runs.

It is evidence and not a gate: no gate reads its report, and what stops a merge is the
job failing.

`.gitleaks.toml` extends the rules the pinned binary carries, which are the rules
`.xeno/plugin/ secrets.yaml` was translated from, with the entropy thresholds, keyword
prefilters and allowlists that section 4's file has no room for. Running the project's
own copy here would repeat what the runner already did, with the precision devices
missing.

Its allowlist names four shapes this repository contains and disables no rule: the
artifact hash fields, which every pattern scanner reads as keys; the pattern file, whose
regexes contain the literals they match; the two samples the tests plant; and the Go
identifier `key` followed by a string.

What this covers that the digest filter does not: every file that stays in the
repository, read once per change. What neither covers: the commit history, anything
before a commit, and the artifacts that G-Secret is specified to judge, which is still
unimplemented.

A checkout carries what is tracked, so an ignored file is not scanned by the job. The
same command run locally does read one, which is the scan working rather than a gap in
it.

<!-- xeno:section:residual-risk -->
## Residual risk

**The workflow has never run.** Every measurement behind it is the same binary invoked
by hand against the same tree. Whether the steps wire together, whether the guard that
separates a failed run from a leak behaves as intended, and whether the checksum step
works on a runner are things the first pull request establishes. A green transcript is
not a green job and the verification says so.

**The allowlist will drift towards a list of exceptions.** Four shapes are named because
they occur today. The next record with a new shape that looks like a key fails the job,
and the cheapest fix is another line here. Nothing reviews that file, and a file of
exceptions nobody reads is how a scanner stops meaning anything.

**A red pull request will usually not belong to whoever opened it.** The first real
finding arrives as a failure on somebody's unrelated change, and the remedy for a
committed secret is a rotation rather than an edit. Nothing here prepares that path.

**No history, and nothing before a commit.** A secret committed and removed is still in
the log and this does not look there; `examples/hooks/` is where the earlier catch
belongs and it is not there either.

**G-Secret is unchanged and this does not shorten it.** It is bound to the effective
filter by section 4, which is a narrower rule set than this job runs, so the two will
disagree about what a secret is and that disagreement is by design.

**Accepted with the five named.** The state it replaces is a repository where the only
thing reading for secrets was a filter over one file per phase, and where nothing read
the rest at all.
