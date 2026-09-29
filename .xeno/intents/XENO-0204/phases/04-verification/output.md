---
intent: github.com/triplem/xeno#128
phase: 04-verification
created: "2026-09-29T09:32:50Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+38430d4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 91206ec1c7154d54ff0b456fefbf06b7e620559bfa94549466ce3c3171b1d047
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
evidence:
  - kind: test-report
    job: go-test
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | What proves it |
|---|---|
| AC1 the workflow, its triggers and its permissions | `.github/workflows/gitleaks.yml`, read: `pull_request`, `push` to `main`, a weekly cron, `contents: read` |
| AC2 pinned and verified | the workflow fetches by version and runs `sha256sum -c` before executing; the transcript prints the tool's version and the tarball's checksum, and `SUPPLY-CHAIN.md` carries both |
| AC3 clean against a tracked tree | the transcript: `git archive HEAD` extracted and scanned, no leaks found, exit 0 |
| AC4 a planted secret fails the job | the same tree with one token added: one `github-pat` finding, exit 1. The token is not in the record |
| AC5 the report goes to `.xeno/local/scan/` | read, and the path is gitignored as the other two scans' are |
| AC6 no matched text in the log | read: the reporting step prints `RuleID`, `File`, `StartLine` and `Fingerprint`, and the transcript shows that shape |
| AC7 a failure to run is distinguishable | read: the report's existence is asserted in its own step before the result is raised again |
| AC8 the configuration extends the binary's rules | `.gitleaks.toml`, read: `useDefault = true` and an allowlist naming shapes, no rule disabled |
| AC9 nothing else changes | `go test ./...`, `gofmt`, `go vet`, and `gate verify` over 103 verdicts |

Six rows are evidence and three are read. AC3 and AC4 are the pair that matters: the scan is silent
on this repository and loud on a real token, and neither is an assertion about the other.

<!-- xeno:section:results -->
## Results

The scan is clean against a tracked tree and exits 0. With one planted token of genuine
entropy added to that same tree it reports one `github-pat` finding and exits 1. Both
runs are in the transcript; the token is not, because this record is committed into the
tree the job scans, and a fixture written here would fail the next run.

The extracted tree rather than the working directory is the subject, and that is not a
formality. A working directory here carries an ignored `.env`, which a checkout does
not, and its two findings would have read as the configuration failing rather than as
the scan working.

What the configuration had to answer is in the transcript as a note: forty findings
before it existed, thirty-nine of them `generic-api-key` on lines of the form
`secrets_hash: <sixty-four hex>`, and one on the pattern file itself. Every one of those
is this repository containing something shaped like a key, and none of them is a rule
that should be switched off.

`go test ./...`, `gofmt -l .` outside `vendor/`, `go vet ./...` and `./xeno gate verify`
over 103 verdicts are unchanged, which is AC9: this intent adds a workflow and a
configuration file and touches no Go.

Read rather than executed: the triggers, the permissions, the report path, the reporting
step's fields and the guard that separates a failed run from a leak. Nothing in CI reads
a workflow file to check it, so those five are a reader's job.

<!-- xeno:section:gaps -->
## Gaps

**The workflow itself has never run.** Everything here was proved with the same binary,
the same configuration and the same tree, invoked by hand. Whether the job's steps wire
together, whether `continue-on-error` and the final step behave as intended, and whether
the checksum step works on a runner are things the first pull request finds out. That is
the ordinary condition of a CI change and it is worth stating rather than implying that
a green transcript is a green job.

**The scan reads a tree, not a history.** A secret committed and later removed is in the
history and this does not look there. The remedy for one found either way is a rotation,
so the difference is about what is discovered, not about what is repaired.

**The allowlist is a promise about the present.** Four shapes are named because they
occur today. A record written next month with a different shape that looks like a key
will fail the job, and the fix will be another line here, which is a file that will
drift towards being a list of exceptions unless somebody reads it occasionally.

**Nothing catches a secret before it is committed.** `examples/hooks/` is where that
belongs and this does not put it there.

**G-Secret is still not implemented**, and this does not shorten that work. Section 4
binds it to the effective filter, which is a different rule set from the one this job
runs.

**The evidence is a local run**, ninth in a row, and here the phrase means more than
usual: the job this intent adds is the one thing the local run cannot stand in for.
