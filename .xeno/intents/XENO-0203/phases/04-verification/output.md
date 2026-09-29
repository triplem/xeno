---
intent: github.com/triplem/xeno#127
phase: 04-verification
created: "2026-09-29T06:56:36Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57acf68.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f65e2a1b119c3a11a68de815d8f264f51ff5c6f57c71b693f18f81860d4352a9
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
| AC1 named provenance | the file header: gitleaks v8.30.1, the sha256 of the upstream `config/gitleaks.toml`, the licence and the copyright, echoed in the transcript |
| AC2 the selection is recorded | the header names the bar, the 221 it started from, and the two exclusions with what each matched |
| AC3 the bar is a test | `TestTheShippedFilterLeavesThisRepositoryAlone`: 222 patterns over 672 files, 0 lines flagged |
| AC4 coverage did not regress | `TestTheShippedFilterCatchesWhatItNames`, asserting all five shapes by id, three of them house |
| AC5 the set did not shrink | `TestTheShippedSetIsTheImportedOne`, asserting at least two hundred patterns and exactly three house |
| AC6 attribution recorded | `NOTICE` and `SUPPLY-CHAIN.md`, read |
| AC7 `secrets_hash` moves once and nothing else | the transcript shows `ebf8354b…` before and `8002ba2d…` after, and `gate verify` matches 97 verdicts |
| AC8 the suite stays usable | the bar skips under `-short`; 23 s with it, 0.14 s without, and 28 ms to redact a real summary |
| AC9 a bad regex still behaves | `TestABrokenRegexDropsItsPatternAndNotTheFilter`, unchanged and passing against the new set |

Nine rows, nine of them evidence. The before and after of the whole change is the pair of hashes in
AC7: one number moved and every verdict still matches.

<!-- xeno:section:results -->
## Results

`go test ./...` passes with the bar included, `gofmt -l .` outside `vendor/` prints
nothing, `go vet ./...` is silent, and `./xeno gate verify` recomputes 97 verdicts and
matches.

The bar is the result. 222 patterns over 672 files of this repository's Markdown, Go,
YAML, JSON, shell and TOML, and 0 lines redacted that do not carry a sample the
repository plants on purpose. The corpus is a harder one than most projects would offer,
because these records discuss secret shapes constantly, and it is the same corpus the
selection was made against, which is worth saying plainly: the bar cannot find a rule
that the selection already removed for failing it.

The two rules the bar removed are in the file's header with what they matched.
`generic-api-key` on `secrets_hash: ebf8354b…`, which is the field that says which
filter was used. `sourcegraph-access-token` on a git commit sha, which every evidence
manifest here carries.

Coverage held, and only because it was checked. The import alone missed three of the
five shapes the old filter caught, and the three house patterns are what closed that;
the test asserts each by id.

`secrets_hash` moved once, from `ebf8354b…` to `8002ba2d…`, which A62 said it would.

Twenty-three seconds for the bar, a hundred and forty milliseconds for the package
without it, twenty-eight for `phase finish` to redact a real summary. The middle number
is why `-short` exists here and the last is why the runner is unaffected.

<!-- xeno:section:gaps -->
## Gaps

**The bar and the selection share a corpus, so the bar cannot fail.** Every rule that
fired on this repository was removed before the test existed, so the test asserts a
property the set was constructed to have. It is still worth having: it fails when
somebody adds a pattern later, or when the records grow a line a pattern catches. What
it does not do is validate the selection independently.

**Zero false positives here is not zero elsewhere.** The corpus is one project. A rule
that is quiet against these files may fire against another project's fixtures, vendored
code or generated data, and the project file is the documented remedy — which widens a
filter and cannot narrow one. A project that needs a shipped pattern switched off has no
route, by section 4's design.

**Nobody has read 219 regexes line by line.** The provenance says where they came from
and the bar says they are quiet here. Neither says they are correct. A rule that matches
too little is invisible to every test in this change.

**The set will go stale.** Upstream adds rules and this file will not notice, which is
the trade `SUPPLY-CHAIN.md` argues for deliberately, and it means the honest description
is "gitleaks' rules as of v8.30.1" rather than "gitleaks' rules".

**Scanning a tree with this set costs twenty-three seconds.** Redaction is unaffected,
because a summary is small. G-Secret is specified to read the same file, and whoever
writes it meets that number.

**The evidence is a local run.** Eighth intent in a row.
