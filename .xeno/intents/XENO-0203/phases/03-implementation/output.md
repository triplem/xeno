---
intent: github.com/triplem/xeno#127
phase: 03-implementation
created: "2026-09-29T06:54:52Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57acf68.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 196193da1d177021c8b777c80ce0b543fded522f2084eac903561e6687532fc0
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

`.xeno/plugin/secrets.yaml`. 222 patterns and 493 lines. The header records gitleaks
v8.30.1, the sha256 of the upstream `config/gitleaks.toml` it was taken from, the
licence and the attribution, the reason a wider match costs less here than upstream, the
bar, and the two rules the bar excluded with what each one matched.

219 patterns are translated from upstream, id and regex preserved so that a redaction
marker names the rule a reader can look up. Three are this project's own and carry a
`house-` prefix: `house-private-key-header`, `house-slack-token-prefix`,
`house-authorization-value`.

`internal/secrets/corpus_test.go`. The bar, as a test: the shipped filter over every
Markdown, Go, YAML, JSON, shell and TOML file of this repository outside `vendor/`,
asserting that no line is redacted unless it carries one of five samples the repository
plants on purpose. 672 files, 0 lines flagged. Redacting per file rather than per line,
and skipped under `-short`, both for the same twenty-three seconds.

`internal/secrets/secrets_test.go`. `TestTheShippedFilterCatchesWhatItNames` now asserts
the five shapes by their new ids, and its comment says why three of them are house
patterns. `TestTheShippedSetIsTheImportedOne` asserts the count and that exactly three
are house, so a set that quietly shrank would fail.

`NOTICE`. The attribution and the licence, and that three patterns are this project's
own.

`SUPPLY-CHAIN.md`. A row for the rules, and a sentence in the paragraph that already
argues why semgrep's rules are vendored rather than fetched, because this is the same
argument.

`internal/secrets/secrets.go` is unchanged. The loader, the union, the canonical hash
and the redaction were written for a filter of any size.

What the numbers came to: 221 upstream rules parsed, 219 kept, 2 excluded, 3 added, 0
lines of this repository falsely redacted, `secrets_hash` moved from `ebf8354b…` to
`8002ba2d…` once.

<!-- xeno:section:deviations -->
## Deviations from the design

Two deviations, one of them about the order of the work.

**The measurement ran before the phases.** The intake of this intent could not be
written without knowing whether the import was possible at all: whether the regexes
compile under the engine the runner uses, whether they fire on this repository, and what
the selection would cost. So the ruleset was fetched, translated and measured, and the
file was written, before P0 existed. That is the deviation XENO-0108 recorded and it is
a different one here in one respect: what preceded the intake was the investigation that
produced its content, and the design's decisions were taken against numbers rather than
re-derived after the fact. It is still the phases following the work.

**The selection rule changed twice while it was being measured**, which the design
presents as a decision and was a sequence. First every rule carrying an allowlist was
excluded, which dropped `aws-access-token` and `github-pat`. Then all 221 were admitted
and the prose bar applied, which excluded `aws-access-token` again, this time on the
sample key in a test file. Only then did the carve-out for planted samples appear, and
with it the criterion that survived. The design reads as though the last of those was
reasoned to; it was arrived at.

Nothing else. The house patterns, the vendored output, the `-short` skip and the
recorded provenance are as P2 decided.
