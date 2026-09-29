---
intent: github.com/triplem/xeno#127
phase: 00-intake
created: "2026-09-29T06:52:41Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57acf68.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f62fb0f00238ff490c56429da4e509c26393a95c23ce1ab066c8444f4915f5fe
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

The shipped filter carried five patterns written by hand while XENO-0202 was built, and
that intent's own review named them as the security surface of the change. #127 records
the objection in one sentence: secret detection is a domain with maintained rule sets in
it, and five regexes somebody wrote in an afternoon are not one.

What the objection does not reach is the mechanism. Section 4 makes the filter a file
with an id and a regex, the runner reads it in process, and `secrets_hash` has to be
reconstructable from the repository. A tool in place of the file would have to pin and
hash its own rules to keep that field meaning anything, which is vendoring under another
name, and the gate path is forbidden the network by a check the verify workflow runs. So
the file stays and its contents change.

The contents were the weak part and they were mine.

<!-- xeno:section:scope -->
## Scope

The shipped patterns are translated from gitleaks' rules at a named version, selected
against a bar, and vendored. Provenance, the upstream file's hash, the selection and
what was left out are recorded in the file itself.

Three patterns of this project's own stay, marked `house-`, for the shapes upstream does
not reach.

A test asserts the bar: the filter leaves this repository's own text alone. Another
asserts the set did not quietly shrink, and a third that every shape the original five
caught is still caught.

`NOTICE` records the attribution and the licence, `SUPPLY-CHAIN.md` gains a row and a
sentence beside the one about semgrep's rules, which makes the same argument for the
same reason.

Not the mechanism. Section 4's file, read in process, unchanged.

Not a generator in the repository. The translation is described precisely enough to
repeat and the output is vendored, which is what `.semgrep/` already does. A script
would add a language to a repository that has one.

Not entropy, keywords or allowlists. Section 4's file has no room for them, so rules
relying on them are either taken as they are, which widens them, or dropped by the bar.

Not G-Secret, and not the CI scan, which is #128.

<!-- xeno:section:context-rationale -->
## Why this context

**The bar is empirical, because the alternative was my judgement again.** Upstream
carries three devices section 4 cannot express: an entropy threshold, a keyword
prefilter and an allowlist. The tempting selection rule was to drop every rule that uses
one, and it fails immediately: it discards `aws-access-token`, which allowlists the AWS
documentation key, and `github-pat`, which carries an entropy threshold. Those are the
two most valuable rules in the set. So the criterion became the one thing that can be
measured instead of argued: does the rule fire on this repository's own text.

**A wider match costs something different here than upstream.** Those thresholds exist
mostly so a scanner does not cry wolf on a placeholder. This filter redacts a summary on
its way into a digest, so the same match replaces a documentation example with a marker.
Upstream's precision is tuned against an alert somebody has to dismiss; here the cost is
a redacted example in prose. That asymmetry is why the entropy thresholds were not
grounds for exclusion, and it is the argument the whole selection rests on.

**Two rules failed the bar and both would have been damaging.** `generic-api-key`
matched `secrets_hash: ebf8354b…`, so it would have redacted the field that says which
filter was used. `sourcegraph-access-token` matched a git commit sha, forty hex
characters, so it would have redacted every commit this repository names. Neither is a
hypothetical: both were found by running the candidate set over 672 files.

**A bigger set is not a superset, which is the finding worth keeping.** The import alone
lost three of the original five: upstream's `private-key` requires the closing line
sixty four characters later, `slack-bot-token` requires both numeric groups, and an
`Authorization` header is caught only inside a `curl` command line. Each is correct for
a scanner reading a file and wrong for a digest, which carries prose where a credential
is quoted, truncated or mentioned. So three house patterns stay and a test asserts they
do, because losing coverage while appearing to gain it is exactly what an import does
quietly.

**The samples this repository plants are not counted against the bar.** The fixtures
contain an AWS example key and a placeholder token on purpose, and a rule matching one
of those is working. That is the carve-out gitleaks makes upstream with its `.+EXAMPLE$`
allowlist, made explicit here because section 4's file cannot hold it.

**The cost is a slower test and not a slower runner.** Two hundred and twenty-two
patterns over a megabyte and a half take about twenty-three seconds, and redacting a
summary takes twenty-eight milliseconds, measured. The bar is therefore skipped under
`-short` and runs in CI, and the per pattern cost is worth knowing before G-Secret comes
to scan a tree with the same set.
