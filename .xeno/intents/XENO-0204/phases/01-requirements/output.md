---
intent: github.com/triplem/xeno#128
phase: 01-requirements
created: "2026-09-29T09:30:06Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+38430d4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a2b8238079afc30a59aa2ff52f48672587efb2ab3f58e1be2542ab7fc5583127
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**AC1.** A workflow named `gitleaks` runs on pull requests, on pushes to `main` and on a
schedule, with `contents: read` and nothing else.

**AC2.** The binary is pinned by version and verified by sha256 before it runs, so
nothing executes that was not named. `SUPPLY-CHAIN.md` carries the version and the
checksum.

**AC3.** The scan is clean against the tracked tree. A checkout of this repository as it
stands produces no findings and the job passes.

**AC4.** A planted secret fails the job. A file carrying a credential shape upstream
catches is reported and the workflow exits non-zero.

**AC5.** The report is written to `.xeno/local/scan/`, gitignored, as the other two
scans write theirs, and nothing is committed by the job.

**AC6.** No matched text reaches the log. A finding is printed as rule, file, line and
fingerprint.

**AC7.** A run that could not scan is distinguishable from a run that found something.
The report's existence is asserted, so a failure with no report says so rather than
reading as a leak.

**AC8.** The configuration extends the rules the pinned binary carries rather than
pointing at `.xeno/plugin/secrets.yaml`, and its allowlist names shapes rather than
disabling rules.

**AC9.** Nothing else changes. `internal/secrets` is untouched, `gate verify` matches
every verdict, and the other three workflows are as they were.

<!-- xeno:section:non-goals -->
## Non goals

G-Secret. Section 4 gives it the artifacts and says it reads the same file the digest
writer reads. This reads what is committed, with a different rule set, once per change.
Implementing the gate is its own work and this does not shorten it.

A pre-commit hook. The remedy for a secret found on a pull request is a rotation, and
the earlier catch belongs in `examples/hooks/` with the limits those templates already
carry.

The filter in `internal/secrets`, unchanged.

History. The scan reads the working tree of a checkout, not every commit that reached
it. A secret removed in a later commit is still in the history and this does not look
there.

Ignored files. A checkout does not carry them. The same command run locally does, which
is the scan working.

A finding threshold or a baseline. Every finding fails the job, which is what a
repository with none can afford and a repository with a backlog cannot.

<!-- xeno:section:constraints -->
## Constraints

Nothing under `docs/` changes.

Evidence, not a gate. `trivy.yml` states the stance and this follows it: no gate reads
the report, and what stops a merge is the job failing.

The scan has to be clean on the day it lands, or it is a job somebody disables in a
week. That is AC3 and it is what the configuration file exists for.

The allowlist names shapes and never disables a rule. `generic-api-key` catches real
keys, and switching it off would be one line and the wrong one.

No secret in a log, in any step, including a failure path.

The binary is verified before it runs. A pinned version fetched over the network without
a checksum is a pinned name, not a pinned artifact.

No new vendored dependency and no Go code. This is a workflow, a configuration file and
two lines of a table.

One intent, one issue. The commits reference #128.
