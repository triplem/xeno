---
intent: github.com/triplem/xeno#12
phase: 00-intake
created: 2026-09-24T20:34:38Z
schema_version: "1.0"
runner_version: 0.1.0-dev+c7abd5d
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: ce52dcd78a2b436b95932f446aa56a356fd13038320048e89795eb847a6b3814
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #12. Static analysis of this project's own code, in the same process shape as the
vulnerability scan: a scanner is evidence, no gate reads the report, and the threshold
lives in the tool.

## Scope

semgrep with the `p/golang` rule set **vendored into the repository**, run on pull
request, on push and on a schedule, with every finding failing the job.

## Non goals

`p/secrets`, which reported nothing even against a planted key. Secret detection is
adjacent to G-Secret and is its own decision.

Any gate that reads the report, and the evidence declaration, which waits for a phase
allowed to write one, exactly as in #11.

## Why the rules are vendored when the vulnerability database is not

The two landed opposite ways in the same week and the reason is what each one reads.

Trivy's database is facts about the world that other people discover, so it goes stale
by time passing and wants a mirror with a cadence. A semgrep rule set is patterns
somebody chose to enforce; it goes stale only when nobody changes it, and a set that
changes underneath a project can fail a build that nothing in the repository touched.

Facts want currency. Policy wants a commit.

## What was measured before any of this was decided

The canary first, because a scan that reports nothing and a scan that did not run look
the same in a log. A file with `crypto/md5` and `math/rand` is found by both candidate
sets; this repository is found clean by all of them; and the vendored pack was run with
the network removed to see whether it needed it.
