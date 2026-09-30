---
intent: github.com/triplem/xeno#143
phase: 04-verification
created: "2026-09-29T20:38:02Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+530ce03.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 61a97308668fc7eff486aa6533ca29141ba06d7831e2c02e29aa253d5910de0e
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

| criterion | how it is verified |
|---|---|
| 1, the port returns requirements | it compiles: `BranchRules` has one method returning `[]enforcement.Requirement`, and `internal/host` imports the domain while nothing imports `internal/host` back except the runner |
| 2, every GitHub spelling in one package | `grep -rn 'api\.github\.com\|vnd\.github\|/repos/\|enforce_admins\|required_status_checks\|required_pull_request_reviews' --include=*.go .` outside `vendor/` and outside `internal/host/github/` returns nothing |
| 3, selection with no default | `TestThereIsNoDefaultHost`, `TestTheAddressIsRequiredRatherThanAssumed`, `TestAnUnknownAdapterSaysWhatThereIs`, `TestTheConfiguredAdapterIsSelected`, `TestAdaptersIsSortedAndNotEmpty`; each refusal is asserted to name the `project.yaml` field |
| 4, `Protection` gone, `Compare` shrunk | `Protection` and `ReviewsExpressible` are absent from the package; `TestANameTheDomainDoesNotKnowIsDropped` and `TestTheAdapterInventsNoRequirementName` hold the name budget; `TestNamesHoldsEveryRequirementTheDomainKnows` holds that the table and `Names()` agree |
| 5, existing assertions unchanged | every case from `main`'s `enforcement_test.go` is present, in the domain or in the adapter, with its `t.Fatalf` wording intact. Eleven cases on `main`; nineteen now, ten in the domain and nine in the adapter, of which one is a split and six are new |
| 6, A42 | `go list -deps ./internal/gates | grep -xE 'net|net/http'` finds nothing, and `.github/workflows/xeno.yml` is untouched in this diff |
| 7, the suite and the real host | below, and the report compared against `main`'s own binary |

The check that carries the most weight is not in the table's first six rows. It is that
`xeno enforcement check` against `triplem/xeno` produces a report **byte identical to the one
`main`'s binary produces**, apart from `checked_at`. A65 says the GitHub phrasings move rather
than get rewritten, and this exercises the real host rather than a fixture, so it tests the
words, the states and the order at once.

Method: `git worktree add` at `main`, build both binaries, run both with the same token, diff the
reports with `checked_at` filtered out.

<!-- xeno:section:results -->
## Results

All seven criteria hold.

| check | result |
|---|---|
| `go build -o xeno ./cmd/xeno` | builds |
| `go vet ./...` | clean, no output |
| `go test ./...` | every package ok |
| `gofmt -l .` outside `vendor/` | prints nothing |
| `xeno gate verify` | exit 0 |
| `go list -deps ./internal/gates` | neither `net` nor `net/http` |
| GitHub spellings outside `internal/host/github/` | none |

**The report is unchanged.** `diff` of `main`'s report against this branch's, with `checked_at`
filtered, is empty.

```
triplem/xeno, branch main
  met           required_pipeline    declared true   actual required
  met           allow_bypass         declared false  actual administrators are included
```

Which is what #84 decided and A27 measured, in the same words and the same order.

| | `main` | here |
|---|---|---|
| `internal/enforcement/enforcement.go` | 331 lines | 216 |
| `internal/host/host.go` | — | 88 |
| `internal/host/github/github.go` | — | 250 |
| test cases | 11 | 19 |
| coverage, domain | — | 88.9% |
| coverage, `internal/host` | — | 100.0% |
| coverage, `internal/host/github` | — | 93.5% |

The line count went up, by 223 across the three files. That is the honest figure and it is what a
port costs: the interface, its selection with two refusals that name their field, and four
evaluation branches that were anonymous arms of a table and are now named functions with a
comment each. The domain lost a third of its length and all of its host.

<!-- xeno:section:gaps -->
## Gaps

**The claim this piece makes is the one it cannot test.** A port is an indirection until a second
adapter fits it without changing it, and there is one adapter. Everything here is consistent with
`BranchRules` being exactly GitHub's shape with an interface in front of it, and nothing available
today distinguishes the two. #104's third piece is the test, and A65's evidence — that GitLab
answers `allow_bypass` with three lists of grants — is the reason to expect it to pass rather than
a demonstration that it does.

The one thing that would have caught a GitHub shaped port is already spent: `Declared` is passed
to the adapter because the host decides met from unmet, and that is the parameter a struct
returning port would not have needed.

**Coverage is 93.5% in the adapter and 88.9% in the domain, not 100.** What is uncovered is the
error returns: a request that cannot be constructed, a body that cannot be read, a client whose
`Do` fails. Reaching them needs a transport that fails on command, which is a fake client and a
larger test surface than the paths are worth. This is the same judgement `internal/secrets` makes
and it is recorded rather than silently accepted.

**The report was compared on one repository, in one state.** `triplem/xeno` on a public plan with
protection present, reviews not required and `merge_method` undeclared, so the comparison exercises
two of the five rows end to end. The other three are covered by unit assertions and by the moved
fixtures, not against a host. A repository with required approvals would exercise the two
approvals rows against a real answer, and this project has none to point at: #84 decided zero
required approvals for a reason that has not changed.

**`Tracker` is not written, so WP12's half of #97's design is unverified.** Nothing here shows
that one adapter type can implement both ports, which is the claim #97 made when it chose two.
It is cheap to satisfy and it is not this intent's to satisfy.

**Not a gap.** The new tests outnumbering the moved ones. Six of the nineteen are new because the
selection, the name budget and the unanswered requirement are new behaviour that A65 and criterion
3 introduced, and behaviour without a test is what the criteria were written to prevent.
