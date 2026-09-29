---
intent: github.com/triplem/xeno#128
phase: 00-intake
created: "2026-09-29T09:29:37Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+38430d4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 92d71d185fc92d909e8f6c5759636b9bdf932a7e5780bef3d8b78da72256420c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

Nothing in this repository scans for secrets. The three scanning workflows answer other
questions: `trivy` runs `scan-type: rootfs` over the built binary for known
vulnerabilities, `semgrep` runs one vendored Go rule set, and `audit` reads the npm tree
the release installs.

Since #120's second half the digest is filtered on its way out of a session, and since
#127 it is filtered against a maintained set. Neither reaches a file that stays. A
secret pasted into a requirements section, a context profile or a learning record is
committed with nothing to notice it, and G-Secret, which section 4 gives that job, is
still written as `not-implemented`.

The review of XENO-0203 put the same thing from the other side: the filter now looks
authoritative, and what it covers is one file per phase.

<!-- xeno:section:scope -->
## Scope

A secret scan as its own workflow, beside `semgrep` and `trivy`: the source tree rather
than the binary, on pull requests and on a schedule, with the tool pinned by version and
verified by checksum.

`.gitleaks.toml` extends the rules the pinned binary carries and records what this
repository contains that is not a secret.

`SUPPLY-CHAIN.md` gains a row for the binary beside the one for the rules.

Evidence rather than a gate, which is the stance `trivy.yml` already states for itself:
no gate reads the report, the threshold lives in the tool's configuration, and what
stops a merge is the job failing.

Not G-Secret. Section 4 says that gate reads the same file the digest writer reads,
which is now possible and still unwritten. A scanner in CI is broader, later and not
bound to the effective filter, so it cannot stand in for it.

Not a pre-commit hook. A scan on a pull request finds a secret that is already in a
branch and the remedy is a rotation; `examples/hooks/` is where the earlier catch would
belong.

Not the filter in `internal/secrets`, which is unchanged.

<!-- xeno:section:context-rationale -->
## Why this context

**The scan uses upstream's rules and not this project's copy, and that is the whole
point of running both.** `.xeno/plugin/secrets.yaml` is the same rule set with the
entropy thresholds, the keyword prefilters and the allowlists removed, because section
4's file has no room for them. Pointing the scanner at it would repeat what the runner
already did, more slowly. The pinned binary carries the rules with those parts intact,
which is what a scanner is tuned for.

**What the scan found first was this repository's own hashes.** Run unconfigured it
reports forty findings, thirty-nine of them `generic-api-key`, almost all matching a
line like `secrets_hash: 8002ba2d…`: sixty-four hex characters after a field name ending
in `_hash` reads as a key named `secrets_hash`. That is the same false positive that
kept `generic-api-key` out of the shipped filter in #127, arriving from the other
direction, and it is the reason a configuration file exists rather than a bare
invocation.

**The fortieth was the pattern file matching itself.**
`aws-amazon-bedrock-api-key-short-lived` fires on `.xeno/plugin/secrets.yaml`, because a
pattern contains the literal prefix it matches on. A file of patterns is not a file of
secrets, and it is allowlisted by path.

**The allowlist is narrow on purpose.** Four regexes and one path, each naming a shape
this repository contains rather than switching a rule off: the artifact hash fields, the
two samples the tests plant, and the Go identifier `key` followed by a string, which is
this project's intent key variable. Disabling `generic-api-key` would have been one line
and would have removed a rule that catches real keys.

**A checkout is the right subject.** It carries what is tracked, so an ignored file is
not scanned here. Running the same command locally does read one, which is how the two
findings in a developer's `.env` appear: that is the scan working rather than a gap in
it, and the workflow comment says so.

**Failing to run and finding a leak look identical from outside.** gitleaks exits 1 for
both, so the report's existence is checked before the failure is raised again. `semgrep`
needed the opposite guard, an `--error` flag, because it exits 0 with findings present;
the two tools are wrong in opposite directions and the workflows say which.

**Nothing prints the match.** The report carries the matched text and a build log is a
poor place for it. Rule, file, line and fingerprint identify a finding without
reproducing it.
