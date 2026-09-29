---
intent: github.com/triplem/xeno#143
phase: 03-implementation
created: "2026-09-29T20:36:59Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+530ce03.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ad97e4880c94000dde64a2250d1e3c5bae247a79bdcdf528553ed1b15ee6ba59
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

Three packages, and one of them shrinks.

**`internal/host`, new.** `BranchRules` with one method, `Requirements(repo, branch, token
string, d enforcement.Declared) ([]enforcement.Requirement, error)`. `BranchRulesFor` selects
by the name `project.yaml` carries and refuses with the field name where either the adapter or
the address is missing. `Adapters` lists what there is, sorted, so an error does not depend on
map iteration order. The package comment restates A65's reason at the interface, because that
is where somebody deciding to add a field will be standing.

**`internal/host/github`, new.** `Adapter` with the base URL as a field. `transport` and
`decode` moved with their bodies intact, renamed from `fetchProtection` and `decodeProtection`
because the package name now carries the host. `protection` is the old `Protection` with
lowercase fields: this package's own shape, unexported, which is what A65 rejected the shared
version of. Four requirement builders hold the `evaluate` bodies that were in the domain's
table, one per requirement, so each is named and reads on its own.

Where the host offers nothing the adapter answers `not-available` for all four names carrying
the reason, rather than answering nothing. That is a change from the first draft of this
intent and the reason is in `deviations`.

**`internal/enforcement`, smaller.** `Protection`, `ReviewsExpressible`, `Fetch`,
`fetchProtection` and `decodeProtection` are gone. The `requirement` table keeps `name`,
`declared` and `waived` and loses `evaluate` and `expressible`. `Compare` takes
`[]Requirement`, indexes it by name, walks the domain's table in report order and applies the
waivers. The names are exported constants with `Names()` in report order, so an adapter builds
its answers from them and a misspelling does not compile. The package comment no longer claims
to be the only package that uses the network, because it no longer uses it at all.

**`internal/runner/enforcement.go`.** Reads `tracker.adapter`, calls `host.BranchRulesFor`, and
turns its error into a `runner.Refusal` so the exit code is 1 by A11. The three lines that
defaulted the address are gone.

**Tests.** Transport, decode and the GitHub specific evaluation move to the adapter's package
with their assertions intact. Waiving, the states, the filter and `merge_method` stay in the
domain, with `Protection` literals replaced by an `answered` helper that says what a host
answered rather than what it was configured to do. Six cases are new: three for the selection
and the two refusals, one for a declared requirement nobody answered, one that the adapter
invents no name, and one that `Names()` covers the table.

<!-- xeno:section:deviations -->
## Deviations from the design

**One found by a test, and it was a real regression.** The first version of `Requirements`
returned nothing where the host offers no protection, and let the domain read the absence as
`not-available`. The state was right and the reason was lost: the 403 branch sets "the host does
not offer branch protection for this repository on its current plan", which used to reach the
report's note through `noteOf(w, p.Reason)`, and an empty answer has nowhere to carry it. The
sentence that tells somebody there is no checkbox to go and find is the whole reason
`not-available` exists as a state, so dropping it would have kept every assertion about states
passing while removing what the states are for. The adapter now answers all four names with the
reason attached, and `TestAnUnavailableHostStillSaysWhyForEveryRequirement` holds it. The domain
still reads a genuine absence as `not-available`, which is now the case where an adapter says
nothing at all rather than the case where it cannot express something.

**Criterion 4's reporting half is not built as written.** "Dropped and reported" has nowhere to
report to. `Report` carries requirements and nothing else, and a field for names the
specification does not have would be the invented field the second standing rule forbids, in the
file that exists to be trustworthy. P2 recorded the alternative and it is what was built: the
names are exported constants, a misspelling does not compile, `Compare` drops an unknown name,
and two tests hold the property — one that the drop happens, one that this adapter produces no
such name. The wording of the criterion was written in P1 against a `Report` nobody had re-read,
which is P2's learning.

**Criterion 5 held, with one case split rather than moved.** Every assertion survives with its
wording intact. `TestNotByAuthorIsCompared` covered three host states and one domain fact, that
an undeclared requirement gets no line, and those two now live either side of the boundary: the
host states in the adapter's package, the filter as `TestAnUndeclaredRequirementGetsNoLine` in
the domain. Nothing was edited to pass and no assertion was dropped. What changed in the domain's
tests is construction, not expectation: `Protection` literals became `answered(...)`, because the
type they named is the thing this intent retires.

**Two comments were reworded for a convention rather than for this intent.** The doc comment on
`BranchRulesFor` and the one in the runner both explained the new code by naming the default it
replaced. `CLAUDE.md` puts that in the commit message: the reader has the current tree and
nothing else. Both now say what the construction is and why, and the history is where the
`api.github.com` default is named. This also took the last GitHub spelling outside the adapter
out of the tree, which criterion 2 asks for.

**`Tracker` is not written**, as P1 said. #104 names it in prose and omits it from the done-when
this intent answers, WP12 is not built, and an interface with no adapter and no caller is what
#97 declined to write.
