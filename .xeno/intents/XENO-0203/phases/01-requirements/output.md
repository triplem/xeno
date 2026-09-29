---
intent: github.com/triplem/xeno#127
phase: 01-requirements
created: "2026-09-29T06:53:39Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57acf68.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: da61321910a0ea922876e66500af66faebe8c44e6263dd2f5063293c0aed9b4e
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

**AC1.** The shipped filter's patterns come from gitleaks at a named version, and the
file records the version, the upstream file's sha256, the licence and the attribution.

**AC2.** The file records the selection: what the bar was, how many rules it started
from, which ones it excluded and why each one failed.

**AC3.** The bar is a test rather than a claim. No pattern in the shipped set fires on
this repository's own text, over every Markdown, Go, YAML, JSON, shell and TOML file
outside `vendor/`, excluding the lines that carry a sample the repository plants on
purpose.

**AC4.** Every shape the original five patterns caught is still caught, asserted by id:
an AWS key, a GitHub token, a private key header, a Slack token prefix and an
`Authorization` value.

**AC5.** The set did not quietly shrink. A test asserts the order of two hundred
patterns and the three house ones.

**AC6.** `NOTICE` carries the attribution and the licence, and `SUPPLY-CHAIN.md` names
the rules beside semgrep's, with the same argument for vendoring rather than fetching.

**AC7.** `secrets_hash` changes once, and nothing else does. The sixty intents that
carry `by-hand` are untouched and every verdict still matches.

**AC8.** The suite stays usable. The bar is skipped under `-short`, and the runtime cost
of redaction is measured rather than assumed.

**AC9.** A regex that will not compile still behaves as it did: dropped from matching,
kept in the hash.

<!-- xeno:section:non-goals -->
## Non goals

The mechanism. Section 4's file, read in process by the runner, with an id and a regex
and nothing else.

A generator in the repository. The selection is described precisely enough to repeat,
and the output is vendored the way `.semgrep/` is. A script would add a language to a
repository that has one, and it is named here so that its absence is a decision.

Automatic currency. Upstream will add rules and this file will not notice. That is the
same trade `SUPPLY-CHAIN.md` already argues for semgrep's rules: policy wants a commit.

Entropy, keywords and allowlists. Section 4 has no room for them and this intent does
not add room.

G-Secret, and the CI scan of #128.

A bar over anything but this repository. The corpus is the text here, which is a harder
one than most because it discusses secret shapes constantly, and it is still one
project.

<!-- xeno:section:constraints -->
## Constraints

Nothing under `docs/` changes. Section 4 fixes the file's shape and this fills it.

The selection must be reproducible from what the file records. A reader who fetches the
named version and applies the stated bar has to arrive at the same set, or the
provenance is decoration.

Coverage may not regress. The five shapes the shipped filter caught before have to be
caught after, whichever rule catches them, because an import that appears to widen while
narrowing is the failure mode of this change.

The bar's carve-out must be explicit. Samples the repository plants are excluded by
name, in the test, rather than by a rule that quietly forgives whole files.

`secrets_hash` moves once and that is expected. Every digest written after this
disagrees with every digest written before, which is what the field is for and what the
release note has to say.

The suite stays fast by default. A bar that makes `go test ./...` unpleasant gets
deleted by somebody in six months.

One intent, one issue. The commits reference #127.
