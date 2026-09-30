---
intent: github.com/triplem/xeno#143
phase: 02-design
created: "2026-09-29T20:28:41Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+530ce03.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ec61f75342987b4256771cd0cbea4297e0703d9a8ee90c9d4719e720e305a701
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

**The port.** `internal/host` holds one interface.

```go
type BranchRules interface {
	Requirements(repo, branch, token string, d enforcement.Declared) (
		[]enforcement.Requirement, error)
}
```

`Declared` is passed in because the adapter cannot evaluate without it: `approvals.required` is
met or not against a declared count, and only the host knows how to read its own answer. The
adapter returns `Met`, `Unmet` or `NotAvailable` and never `Waived`, because waiving is a
decision the project took and not something a host can report.

**The selection.** `host.BranchRulesFor(adapter, baseURL string, c *http.Client)` returns the
adapter or an error listing what there is, the way `runner.Wrapper` already does for wrapper
hosts. That existing function is the precedent for the error's wording, so the two read alike.

**The names become constants in the domain.** `enforcement` exports `NameRequiredPipeline`,
`NameAllowBypass`, `NameApprovalsRequired`, `NameApprovalsNotByAuthor` and `NameMergeMethod`,
and `Names()` returns them in report order. An adapter builds its answers from those constants,
so section 13's budget is expressed in the type system rather than in a comment, and the second
standing rule is held by the compiler for anything an adapter can spell wrong.

**What `Compare` becomes.**

```go
func Compare(repo, branch string, d Declared, answered []Requirement, now time.Time) Report
```

It walks the domain's names in report order. For each name the declaration asks for it takes the
adapter's answer, and a declared name the adapter did not answer for is `NotAvailable`, which is
the honest reading: the host had nothing to say about it. `Met` passes through; `Unmet` and
`NotAvailable` go through `waiveOr`. `Actual` and `Note` are the adapter's words and are carried
rather than replaced. Then section 13's `merge_method` line, unchanged.

The domain keeps `requirement` as a table of name, `declared` and `waived`. `evaluate` and
`expressible` go to the adapter, which is the half A65 named.

**What is deleted.** `Protection`, `ReviewsExpressible`, `Fetch`, `fetchProtection` and
`decodeProtection`. The last three move to `internal/host/github` with their bodies intact; the
first two have no successor, which is what retiring rather than generalising means.

<!-- xeno:section:alternatives -->
## Alternatives

**An unknown requirement name reported rather than dropped.** Criterion 4 asks for "dropped and
reported, not passed through", and the reporting half is not built as written. There is nowhere
for it to go: `Report` has `Requirements` and nothing else, and adding a field for names the
specification does not have would be the invented field the second standing rule forbids, in the
file that exists to be trustworthy. `Compare` returning an error was the other reading, and it
makes every caller handle a condition that is an adapter bug rather than a run time state.

What is built instead: the names are exported constants, an adapter that spells one wrong does not
compile, and a test in `internal/host` asserts every adapter's names are a subset of `Names()`. The
property is held in the same place A42's is — by a check outside the type — and the deviation from
the criterion's wording is recorded in P3 rather than glossed.

**`BranchRules` without `Declared`, returning facts for the domain to judge.** This is the rejected
reading of A65 arriving by the back door: an adapter that returns what it found without evaluating
it is `Protection` again, with a slice instead of a struct. Rejected for the reason the row gives.

**The adapter reading its own token from the environment.** Rejected. `enforcement.Token()` already
reads it in one place and WP12 requires credentials come from the environment and never from
`project.yaml`, which is satisfied by the caller doing it. An adapter that read `os.Getenv` itself
would make the same rule true in two places and testable in neither.

**A default base URL inside the GitHub adapter.** Tempting, and `api.github.com` is that host's
canonical address, so it is not the same offence as the runner defaulting to it. Rejected on
criterion 3's wording, "no host name and no host default appears in code", and because the
configuration already carries it in both the scaffold and this repository, so the default would
only ever mask a misconfiguration.

**One `Host` interface carrying both concerns.** Settled by #97, not reopened. Two ports, because
`BranchRules` serves `enforcement check` and `Tracker` serves WP12, and one adapter type may
implement both without the interfaces being one.

<!-- xeno:section:impact -->
## Impact

**`internal/enforcement` shrinks and stops being the only package that uses the network.** Its
package comment says it is, and that sentence becomes false: the network moves to
`internal/host/github` and the comment has to say what the package is now, which is the domain
with no host in it. The half of the comment that matters is the half about the gate path, and that
stays true.

**A42's property is unchanged but its shape is not.** The gate path still reaches no network, and
`internal/gates` imports neither the domain nor the adapter. The new risk is a future import of
`internal/host` from somewhere in the gate path, which the existing CI check catches because it
walks dependencies rather than names packages. The check needs no edit, and P4 confirms that
rather than assuming it.

**`cmd/xeno` changes in one respect only.** `enforcement.Token()` stays where it is, so the call
site keeps its shape; what changes is that a missing `tracker.adapter` now reaches the user as a
refusal with a field name. By A11 that is exit 1, and `report` already distinguishes a
`runner.Refusal` from an ordinary error, so no new exit path appears.

**The report for this repository must not change.** Same names, same words, same states. This is
criterion 7's second half and it is the strongest available check that the move preserved
behaviour, because it exercises the real host rather than a fixture.

**#104's remaining two pieces get their place.** The GitLab adapter becomes a second directory
under `internal/host` and a row in whatever `BranchRulesFor` selects from, and the five row mapping
becomes that adapter's evaluation. Neither needs the port to change, and that is the claim this
piece makes and cannot yet prove.

**The test suite redistributes rather than grows.** Transport and decode cases move to the adapter
package; waiving, states and `merge_method` stay. Criterion 5 means the count of assertions should
be the same afterwards, and a case that disappears is a finding.

**Nothing else.** No artifact field, no gate, no template, no dependency, no change to
`project.yaml`'s shape or to `ReportPath`.
