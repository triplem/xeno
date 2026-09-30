---
intent: github.com/triplem/xeno#147
phase: 02-design
created: "2026-09-30T21:15:20Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0a51653.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3dbf42d0e686b9199939fa59ecdf6b67957fe45955a44283466b94414bfa7802
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

**Three requests, and each one's failure is local to the requirements it answers for.**

| request | answers |
|---|---|
| `GET /projects/:id/protected_branches/:branch` | `allow_bypass`, and whether the branch is protected at all |
| `GET /projects/:id` | `required_pipeline`, via `only_allow_merge_if_pipeline_succeeds` |
| `GET /projects/:id/approvals` and `/approval_rules` | the two approvals rows |

The GitHub adapter makes one request, so one status code decides everything. Here a tier limit on the
approvals endpoints must not stop the other three requirements being answered, so a failed request
degrades the rows it was for to `not-available` with its own reason and the rest are answered anyway.
That is the structural difference between the two adapters and it is why no code is shared.

**`allow_bypass` reads four things and names the first it finds.**

An `unprotect_access_levels` grant is reported first, because removing the protection is the largest
bypass and reporting a smaller one alongside would bury it. Then `allow_force_push`, then a
`push_access_levels` grant above zero. None of the three is met. GitLab's access level 0 is "no one",
so a grant at 0 is not a bypass, and the numbers are named constants with the reference in a comment
rather than inlined.

The words say which was found — "3 can unprotect the branch", "the branch allows force push", "2 can
push without a merge request" — because `allow_bypass` met on GitHub means administrators are included
and means something else here, and a shared sentence would be false for one of them.

**A tier limit is `not-available` whatever status carries it.** 401, 403, 402 and 404 are all plausible
from the approvals endpoints and #104 guesses none. The decision is that the adapter distinguishes only
two cases: an answer it can read, and an answer it cannot, the latter being `not-available` with the
status in the reason. A token that is simply wrong produces the same state as a tier that lacks the
feature, which is a real loss of precision and is the right trade: the alternative is a table of status
codes written from guesses, and a wrong guess turns a tier limit into a failed command.

`401` is the one exception and it is handled at the transport: a rejected token is an error for the
whole adapter, because a token that cannot read the project cannot answer anything and reporting five
requirements as not available would hide a misconfiguration behind the tier.

**`project` is URL encoded once, in one helper.** `namespace/project` becomes `namespace%2Fproject`.

<!-- xeno:section:alternatives -->
## Alternatives

**Mapping each status code to a state from a table.** Rejected, and this is the closest call in the
intent. It is what the GitHub adapter does and it is right there, because that adapter makes one
request whose codes are documented and observed. Here the codes are neither: `/approval_rules` on a
Free project answers 401 to no token and something unknown to a real one, and #104 marks the row "to
verify" precisely because nobody has looked. A table written from four plausible codes would encode
three guesses, and the cost of a wrong one is a command that fails where it should report. Collapsing
to "readable or not" throws away precision the adapter does not have.

What would change this: the observation in P4. If the real code turns out to be a single clean answer,
a later intent can narrow the mapping, and the report will have been correct in the meantime.

**Reading `allow_bypass` as a count rather than a finding.** "4 principals may bypass" is one number
and comparable across hosts. Rejected: the three lists mean different things, and a total implies they
are commensurable. Somebody who can unprotect the branch and somebody who can push to it are not two
of the same thing.

**Sharing a transport helper with the GitHub adapter.** Rejected, and P1's non-goals say why. The
request shapes differ, the identity encoding differs, the status meanings differ and the number of
requests differs. What is left in common is `http.NewRequest` and a header, which is the standard
library.

**Copying `protection` from the GitHub adapter as a starting struct.** Rejected for the same reason
A65 rejected the shared one. This adapter's intermediate shape is three payloads with three
independent failures, which is not GitHub's shape with different field names.

**Treating a Free tier as unmet rather than not available.** Rejected by criterion 4 and by the
domain's own comment: a setting the host does not have is nobody's oversight. It would also make every
Free project's report red forever, which is the standing complaint the waive logic exists to prevent
and which would be reintroduced one tier down.

**Requiring Premium and refusing otherwise.** Rejected. Section 8 says the GitLab meant is Enterprise,
which is a statement about what the adapter must be able to express, not a licence to refuse the tier
an adopter has.

<!-- xeno:section:impact -->
## Impact

**#104 closes.** All four pieces: the vocabulary in #141, the port in #143, the wrapper in #145, the
adapter here. The work package `wp12` has no open issue after this.

**A65 gets its test.** The port is unchanged or it is not, and either way the answer is recorded. If it
is unchanged, the decision A65 took on evidence from this host is borne out by the host itself, which
is the strongest form the claim can take. If it is not, `ASSUMPTIONS.md` gains a row against A65 and
this intent stops rather than widening the interface.

**`internal/host` gains its second entry and stops being a table with one row.** `Adapters()` returns
two, so `BranchRulesFor`'s error message becomes useful rather than decorative, and #145's refusal
listing hosts now lists two hosts everywhere it appears.

**A GitLab project becomes checkable.** `xeno init --host gitlab` writes a configuration that works end
to end for the first time: a wrapper that runs the gates, and an `enforcement check` that answers. What
it answers depends on the tier, and on Free two of five requirements are `not-available`, which a
project records as waived once and then does not see again — the waive logic doing exactly what it was
built for, one tier rather than one plan down.

**The report gains a shape nobody has seen.** Every report so far has been GitHub's. A reader comparing
two projects' reports will see the same five names with different words and, on Free, different states
for reasons that are the tier and not the project. The names being section 13's is what makes them
comparable at all, which is the part of A65 that was about the domain keeping something.

**Three requests where there was one.** The command takes longer and can partly fail, which the GitHub
path cannot. A daily schedule now has a way to report four requirements and miss one.

**Nothing in the gate path, the domain, the artifacts or the hashes.** No new field, gate, template or
dependency. A42 holds with a third `net/http` package under `internal/`.
