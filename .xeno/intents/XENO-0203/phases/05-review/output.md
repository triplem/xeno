---
intent: github.com/triplem/xeno#127
phase: 05-review
created: "2026-09-29T06:57:18Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57acf68.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d4a236361f7f4ca7df74115480e6790dddb1d6831b5bc0852e892e304c258a62
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**Nothing under `docs/` changed.** Section 4 fixes the file's shape and this fills it.
The mechanism question the issue's title raises is answered by four constraints already
in the documents, and the design names all four.

**Nothing was invented.** No field, no gate, no dependency, no command.
`internal/secrets` is unchanged: it was written for a filter of any size and a filter of
222 exercises it.

**The provenance is in the artifact, not only in the record.** The file's header carries
the version, the upstream hash, the licence, the bar and the exclusions, so a reader who
never finds this intent can still tell where the patterns came from and what was
dropped.

**Coverage was checked rather than assumed, and it had regressed.** The import missed
three of the five shapes the old filter caught. Three house patterns close it and a test
asserts each by id.

**The bar is zero over 672 files**, and the review says where that claim stops: same
corpus as the selection, one project, and nobody has read 219 regexes.

**`secrets_hash` moved once**, `ebf8354b…` to `8002ba2d…`, and 97 verdicts still match.

**The suite stays usable.** 23 s with the bar, 0.14 s under `-short`, 28 ms to redact a
real summary. All three measured, and the last is the one that matters for the runner.

**The order of the work is admitted.** The measurement produced the intake, so it
preceded it, and the deviations say the selection rule was arrived at rather than
reasoned to.

**Attribution is recorded** in `NOTICE` and `SUPPLY-CHAIN.md`, the latter beside the
paragraph that makes the same argument for semgrep's rules.

<!-- xeno:section:release-notes -->
## Release notes

The shipped secret filter carries 222 patterns instead of five. 219 are translated from
gitleaks' rules at v8.30.1, MIT licensed, with the ids preserved so a redaction marker
names a rule that can be looked up. Three are this project's own and carry a `house-`
prefix.

The three exist because an imported set is not a superset. Upstream is tuned to find a
whole credential sitting in a file: its `private-key` rule requires the closing line,
its Slack rule requires both numeric groups, and an `Authorization` header is caught
only inside a `curl` command. A digest is prose, where a credential is quoted, truncated
or mentioned.

Two upstream rules were left out, and the file says which and why: `generic-api-key`,
which matched the `secrets_hash` field itself, and `sourcegraph-access-token`, which
matched a git commit sha. The bar was that no pattern may fire on this repository's own
text, and it is a test rather than a claim.

`secrets_hash` changes once. Every digest written from here disagrees with every digest
written before, which is the field saying that the filter changed.

`NOTICE` and `SUPPLY-CHAIN.md` record the source, the version and the licence. The rules
are vendored and not fetched, for the reason `SUPPLY-CHAIN.md` already gives about
semgrep's: facts want currency, policy wants a commit.

`go test ./...` is about twenty-three seconds slower, which is the bar running over the
whole repository. `-short` skips it. Nothing about redacting a digest is slower.

<!-- xeno:section:residual-risk -->
## Residual risk

**The bar cannot fail on what shipped.** It runs over the corpus the selection was made
against, so every rule that fired was already removed. It guards what comes next and
validates nothing now, which the verification says and the test's own comment does not.
That is the weakest part of this change and it is weak by construction rather than by
oversight.

**Nobody has read 219 regexes.** The provenance says where they came from, the bar says
they are quiet here, and neither says they are correct. A rule matching too little is
invisible to every test in this change, and that is the shape of the remaining risk: the
filter now looks authoritative.

**Zero false positives here is not zero elsewhere**, and a project cannot switch a
shipped pattern off, by section 4's design. The first project this catches wrongly has
no remedy but an issue.

**The set will go stale**, deliberately. "gitleaks' rules as of v8.30.1" is the honest
description and nothing notices when upstream moves. Whether a periodic re-import is
worth a recurring intent is a real question this leaves open.

**Scanning a tree costs twenty-three seconds with this set.** G-Secret is specified to
read the same file and will meet that number, and it is better known now than discovered
then.

**A wider filter redacts more of a digest.** Two hundred patterns over prose that
discusses credentials will replace more text than five did, and the records of this
project are exactly such prose. Nothing here measures how often that will happen in
practice.

**Accepted with the six named.** The state it replaces was five patterns with no
provenance, and the one thing this change buys that is beyond argument is that a reader
can now check where the filter came from.
