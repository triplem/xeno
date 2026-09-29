---
intent: github.com/triplem/xeno#141
phase: 04-verification
created: "2026-09-29T19:56:23Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6aa6f9b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bfb90adac650c0414da729caa3f1899a7f6fd3a73ed5635b3a79f3da6447b122
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

This intent adds no code, so nothing here maps a criterion to a new test. Six of the seven
acceptance statements are about the content of one table row and are verified by reading it;
the seventh is about the tree and is verified by the check suite. Both kinds are recorded,
because a phase that reported "no tests, nothing to verify" would leave the row unchecked.

| criterion | how it is verified |
|---|---|
| 1, one row `A65` naming both readings | `grep -c '^| A65' ASSUMPTIONS.md` is 1, and the row names the reading taken in its first clause and the rejected one after **The rejected reading is** |
| 2, the reason is the deciding case | the row carries `allow_bypass` with `enforce_admins` against the three grant lists, and says where the judgement would otherwise have to live |
| 3, the row cites what made it decidable | the row names `GET /api/v4/projects/:id/protected_branches` and that it answers unauthenticated |
| 4, the domain keeps the names | the row carries **The domain keeps the names** with section 13 and the second standing rule as the reason |
| 5, the cost is recorded | the row carries **The cost, recorded rather than argued away** with the drift and the partial undoing of #99's table |
| 6, state and cross reference | the state column reads `approved` and bounds it; the `Where` column names `internal/enforcement`, #97 and #104 |
| 7, nothing else changed, suite clean | `git diff --stat` against `main` outside this intent's phase directories is `ASSUMPTIONS.md | 2 ++`; `go build`, `go vet ./...`, `go test ./...`, `gofmt -l .` outside `vendor/`, and `xeno gate verify` |

The existing `enforcement_test.go` is named in the row and is not run against anything new
here. Its role is in piece 1: it holds the GitHub phrasings, so it is the assertion that the
move of `evaluate` into an adapter was a move rather than a rewrite. Naming it now is what
makes that checkable then.

<!-- xeno:section:results -->
## Results

All seven criteria hold.

The row is present once, `grep -c '^| A65'` returns 1, and it carries each of the five parts
the criteria name: the reading taken, the deciding case, the rejected reading with why
expressibility does not answer it, the unauthenticated endpoint that made it decidable, and
the drift cost. The state column reads `approved` and then bounds what was approved to the
decision, naming the port as #104's first piece. The `Where` column names
`internal/enforcement`, #97 and #104.

The tree is what the seventh criterion asks for. `git diff --stat` against `main`, excluding
this intent's own phase directories, is `ASSUMPTIONS.md | 2 ++`, one file and two insertions.

| check | result |
|---|---|
| `go build -o xeno ./cmd/xeno` | builds |
| `go vet ./...` | clean, no output |
| `go test ./...` | every package ok; `internal/secrets` is the slow one at 33.285s |
| `gofmt -l .` outside `vendor/` | prints nothing |
| `xeno gate verify` | `verified 145 verdicts`, exit 0 |

`gate verify` recomputing 145 verdicts and finding none divergent is the check that matters
most for a phase like this one, because it says the twelve artifacts this intent has written so
far hash to what their verdicts were taken against, and Appendix B puts the phase path inside
`artifacts_hash`. A decision recorded in a phase that did not verify would be a decision
nobody could later prove was the one judged.

<!-- xeno:section:gaps -->
## Gaps

**Two rows of #104's mapping are unverified, and this is the gap that matters.** The merge
settings — `only_allow_merge_if_pipeline_succeeds`, `merge_method`, `squash_option` — are
absent from an unauthenticated project payload rather than returned false, so their presence
and spelling are taken from documentation and not from the host. `/api/v4/projects/:id/approvals`
answers 401, so `merge_requests_author_approval`, the setting the whole
`approvals.not_by_author` argument rests on, was not observed either.

The decision does not depend on the second of those in the way it might look. What it rests on
is that the requirement is met by construction on GitHub and settable on GitLab, and the first
half is in this repository's own code while the second is what section 8 says the Enterprise
variant has. If the field turns out to be named or shaped differently, the argument stands and
the mapping row changes. If GitLab turned out to refuse an author's own approval inherently, as
GitHub does, that would weaken this case — and `allow_bypass` would still decide it alone.

**What closes the gap.** #104's third piece, against the Premium namespace section 8 asks for,
which is deliberately not arranged yet. The figures and the shape observed here are in this
phase's directory so that piece does not start from nothing.

**A second gap, smaller.** No test asserts that a GitHub adapter reproduces the current
phrasings, because there is no adapter. The row names `enforcement_test.go` as what pins them,
which is a claim about piece 1's verification rather than a check made here.

**Not a gap.** That no test was added. The deliverable is a decision, and the process's own
verifier is what checks it: 145 verdicts recomputed, none divergent.
