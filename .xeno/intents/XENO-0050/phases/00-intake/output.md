---
intent: github.com/triplem/xeno#50
phase: 00-intake
created: "2026-09-25T21:01:25Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+9f4b620
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 925cdfd2f1bf16de7807f755734b75f0a146b16ccbff5fe39728631614ef91ef
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

Both scanners run, and both enforce their threshold by failing the job. Neither writes
a file. The findings go to the log and nothing outlives the run, so the last acceptance
criterion of #11 and of #12 is unmet in the same way in both: no report exists that a
later phase could declare as `kind: scan`.

<!-- xeno:section:scope -->
## Scope

Each scan job writes its report to a stable path and a `manifest.yaml` beside it, and
uploads the pair as a CI artifact. Trivy's database metadata travels with its report as
a second entry, since `UpdatedAt` is in the cache and in no output.

The declaration itself is not in scope. It needs a phase beyond P0, and no phase beyond
P0 may run before WP8 closes the G-Freshness gap.

<!-- xeno:section:context-rationale -->
## Why this context

**Nothing is designed here.** `internal/evidence/attach.go` already reads a source
directory holding `manifest.yaml` with `kind`, `job`, `result`, `file` or `uri`
plus `sha256`, `pipeline` and `commit`; `xeno evidence attach --from DIR` consumes
it and hashes on attachment. The work is producing what the reader already expects.

The failing run is the one worth reading, and it is the one a naively written step
misses: the scan fails by design, `exit-code: 1` for Trivy and `--error` for semgrep,
so a report step placed after it never runs. The exit code has to be captured and
re-raised, and `result` in the manifest is that code rather than the absence of an
error. This is to be shown by making a scan fail, not by reading the workflow.

The database age is a workaround and is recorded as one. The finding routed here from
#42 says the value belongs beside `produced_by` and `result`, which is a field the
process definition does not have. Shipping `metadata.json` as its own item puts the
value in the trail without inventing one.

<!-- xeno:section:open-questions -->
## Open questions

- Q-1: whether the upload of a CI artifact is a `uri` item or a `file` item at
  attachment time. A JSON report is small and textual, which argues for `file`; it
  lives in the artifact store, which is what `uri` is for and what makes evidence
  uncheckable to write by hand.
