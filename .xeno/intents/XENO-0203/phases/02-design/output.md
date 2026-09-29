---
intent: github.com/triplem/xeno#127
phase: 02-design
created: "2026-09-29T06:54:17Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57acf68.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0a6dcc237594662227b047b6a8865f642438e6a73016e6d9ddaf585f5121937e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: by-hand
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The bar is "does it fire on this repository's text", and nothing else.** One
criterion, applied to all 221 upstream rules uniformly. The alternatives were criteria
over the schema — no entropy, no allowlist — and each of those discards a rule everybody
would name first.

**Samples the repository plants are named in the test, not forgiven by file.** Five
strings, listed, because excluding whole files would hide a real match in the rest of
them.

**Three house patterns stay, and their ids say so.** `house-private-key-header`,
`house-slack-token-prefix`, `house-authorization-value`, for the three shapes the import
misses. The prefix is the point: a reader of a redaction marker can tell whether the
rule came from upstream or from here.

**The output is vendored, not generated at build time.** Provenance, the upstream hash
and the selection go in the file's header, so the set is reproducible by repeating a
described procedure. That is what `.semgrep/` already does, and `SUPPLY-CHAIN.md` gains
a sentence saying the two follow the same rule for the same reason.

**The bar is skipped under `-short`.** Twenty-three seconds of regex over the corpus,
against twenty-eight milliseconds to redact a real summary, both measured. CI does not
pass `-short`, so the bar runs where it has to.

**`secrets_hash` moves once and the release note says so.** Every digest written after
this disagrees with one written before, which is the field working.

<!-- xeno:section:alternatives -->
## Alternatives

**Keep the mechanism and call a tool instead of reading a file.** The objection's first
reading. Rejected on four independent grounds, any one of which would be enough: section
4 makes the filter a file; `secrets_hash` has to be reconstructable from the repository,
which a tool's moving rule set defeats unless it is pinned and hashed, which is
vendoring; the verify workflow refuses `net` and `net/http` among the gate path's
dependencies; and section 14 has the harness hook filtering on every write, where a
subprocess is the wrong shape. Recorded because the issue's title invites it.

**Exclude every rule that uses entropy, keywords or an allowlist.** The schema-shaped
criterion, and the one that reads as principled. It drops `aws-access-token` and
`github-pat`. Rejected on the measurement.

**Import all 221 without a bar.** Two rules then redact `secrets_hash` itself and every
git commit sha in the repository. The bar is not a refinement of the import; it is what
makes the import safe.

**Drop the three house patterns and take upstream's coverage as the definition.**
Tidier, one source, one provenance. It silently loses private key headers, Slack
prefixes and `Authorization` values, because upstream is tuned to find a whole
credential in a file. Rejected on the coverage test that exists to catch exactly this.

**Write a generator and run it in CI.** Currency without a manual step, and the set
would then change underneath a project without a commit, which is the failure
`SUPPLY-CHAIN.md` argues against for semgrep's rules at length. Rejected on that
argument, and it would add a language.

**Add entropy to section 4's file so the rules translate faithfully.** The honest way to
import everything. It is a specification change, so it is a person's commit before any
code, and the case for it is weaker than it looks: entropy exists to spare a scanner
from placeholders, and a redactor can afford to redact one.

<!-- xeno:section:impact -->
## Impact

`.xeno/plugin/secrets.yaml`: 222 patterns instead of five, 493 lines, most of it the set
and the rest a header recording where it came from and what was left out.

`internal/secrets/corpus_test.go`: the bar. `internal/secrets/secrets_test.go`: the
coverage assertions and the count.

`NOTICE`, `SUPPLY-CHAIN.md`: attribution, licence, a row and a sentence.

`internal/secrets/secrets.go` does not change. The loader, the union, the canonical hash
and the redaction were built for a filter of any size, and a filter of two hundred and
twenty-two exercises them rather than needing them altered.

What a reader gains: a filter with a maintained provenance and a measured false positive
rate of zero over this repository. What they still do not gain: currency, since upstream
will add rules and this file will not notice, and coverage of anything but a digest.

What this costs: `secrets_hash` changes once, `go test ./...` gains twenty-three seconds
unless `-short` is passed, and the per pattern cost is now a fact to plan around rather
than a detail — scanning a whole tree with this set takes twenty-three seconds, which is
what G-Secret will meet.
