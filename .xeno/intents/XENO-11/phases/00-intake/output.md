---
intent: github.com/triplem/xeno#11
phase: 00-intake
created: 2026-09-24T20:04:36Z
schema_version: "1.0"
runner_version: 0.1.0-dev+c7abd5d
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: cf691719751de3bde3e9a44b7a957ad71bda3f7f31e0586023ebead91895818d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #11. Trivy reports what the release carries and names the fixed version where
there is one.

## Scope

A scan of the **built binary** on pull request and on a schedule, with the threshold in
`trivy.yaml` in the repository, the tool pinned like everything else and recorded in
`SUPPLY-CHAIN.md`.

The target was measured rather than argued: only a binary scan reports `stdlib` with its
version. A source scan sees the declared dependency and nothing of the Go standard
library, which is the half this project moved to 1.27 for.

## Non goals

Any gate that reads the report. A scanner is evidence, and no gate evaluates its
findings against a threshold. There is no G-Trivy and there will not be one.

Secret and misconfiguration scanning, both their own decision.

Declaring the report as `kind: scan` evidence. That lives in the frontmatter of a
phase's `output.md`, and only P0 runs here until WP8. The job produces the report; the
declaration waits for a phase allowed to write it.

## The threshold, and the half worth arguing

HIGH and CRITICAL, fixed only. A finding with no fix available blocks work that nobody
can do, and a check that fires without a possible response is one people learn to route
around. Those findings stay in the report and do not stop the job.
