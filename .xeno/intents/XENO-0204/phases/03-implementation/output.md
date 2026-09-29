---
intent: github.com/triplem/xeno#128
phase: 03-implementation
created: "2026-09-29T09:31:29Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+38430d4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 068e2dc389009e0ab70a799540a4bd6a77a01c36aebf9d1bde7cb25d0ebe62a0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: by-hand
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`.github/workflows/gitleaks.yml`. The workflow, in the shape `semgrep.yml` established:
install, scan with `continue-on-error`, print what was found, assert a report exists,
raise the scan's own result. On pull requests, on pushes to `main`, and weekly, with
`contents: read`.

The binary is fetched from its release by version and checked with `sha256sum -c` before
it runs, so the pin is on the artifact and not only on the name. The version is the one
#127 translated its rules from, which is why a redaction marker in a digest names a rule
this scan can also report.

`.gitleaks.toml`. `useDefault = true`, then one allowlist with one path and four
regexes. The path is `.xeno/plugin/secrets.yaml`, which contains the literals its own
patterns match on. The regexes are the artifact hash fields, the two samples the tests
plant, and `key, "` in Go source.

`SUPPLY-CHAIN.md`. A row for the binary, with its version and the first bytes of its
checksum, beside the row for the rules.

**What the scan reported before the configuration existed.** Forty findings. Thirty-nine
were `generic-api-key`, and almost all of them matched a line of the form `secrets_hash:
8002ba2d…`: sixty-four hex characters after a field name ending in `_hash` reads as a
key named `secrets_hash`. The fortieth was `aws-amazon-bedrock-api-key-short-lived`
firing on the pattern file, because that rule's regex contains the base64 literal it
matches on.

**What it reports now.** Nothing, against a tracked tree, which was checked by
extracting `git archive HEAD` and scanning that rather than the working directory: a
working directory carries an ignored `.env`, which a checkout does not, and two findings
in it would otherwise have read as a failure of the configuration rather than as the
scan doing its job.

**And it still fails on a real one.** A file carrying a GitHub token of genuine entropy,
added to that same extracted tree, is reported as `github-pat` and the command exits 1.
The token is not reproduced in the record or in the evidence, for the reason the
workflow itself gives about logs, and because this record is scanned by the job it
describes.

<!-- xeno:section:deviations -->
## Deviations from the design

The workflow and the configuration were written before the phases, and the reason is the
same one XENO-0203 recorded: the intake's subject was what an unconfigured scan reports
against this repository, which is not knowable without running it. The forty findings,
their shapes and the two exit codes are measurements, and the design was written against
them.

One thing P2 did not foresee. The path allowlist was first anchored with `^`, which
matches the relative path a scan run from the repository root reports and not the
absolute one a scan of an extracted tree reports. The anchor is gone, so the rule holds
whichever way the scan is invoked. That surfaced only because AC3 is checked against an
extracted tree rather than the working directory.

A second, smaller: the evidence file for this intent cannot carry the planted token,
because it is committed into a tree this very workflow scans. So the verification
records the rule, the file, the line and the exit code, and not the string. That
constraint is new — no previous intent's evidence had to avoid being its own subject.
