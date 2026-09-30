---
intent: github.com/triplem/xeno#147
phase: 04-verification
created: "2026-09-30T21:25:39Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0a51653.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ca10f938b68327d4c3123e97a49db46d18d03fe976f6fdf47cf5c356d74a1166
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
| 1, the adapter implements and is registered | it compiles; `Adapters()` returns two, and `xeno enforcement check` selected `gitlab` from a generated `project.yaml` |
| 2, `BranchRules` unchanged | `diff` of the interface block against `main`'s: unchanged, not one field |
| 3, the command answers for a real project | `xeno enforcement check --branch master` against `gitlab-org/gitlab`, run from a repository `xeno init --host gitlab` generated |
| 4, a tier limit is `not-available` with a reason | observed in that run, and unit tested at 403, 404 and 402 with an assertion that the other requirements survive |
| 5, the status codes observed | observed with a real token, and the observation changed the code. Below |
| 6, `allow_bypass` names its finding | the live run reports "2 grant(s) may push without a merge request"; the four findings and the level-zero case are a unit table; a test asserts the word "administrator" never appears |
| 7, no invented name | `TestTheAdapterInventsNoRequirementName`, and the live test checks every returned name against `enforcement.Names()` |
| 8, the GitHub path unchanged | `xeno enforcement check` against this repository, diffed against `main`'s binary's report |
| 9, the suite | build, vet, test, gofmt, `gate verify` |

Thirteen unit cases, two live tests skipped unless `XENO_LIVE` is set. The live ones are not a
criterion's only evidence and they are not in the suite: a test that needed a network would put the
network where A42 spends a CI check keeping it out.

<!-- xeno:section:results -->
## Results

All nine criteria hold.

**Criterion 2 is the result this intent was arranged to produce.** `host.BranchRules` is byte
identical to `main`'s. The second adapter has a structurally different shape — three requests with
independent failures against the first's one — and the interface took it without a field. A65 decided
the vocabulary on evidence from this host, and the host has now borne it out.

**Criterion 3, the command, against `gitlab-org/gitlab`:**

```
met           required_pipeline    declared true   actual only_allow_merge_if_pipeline_succeeds is on
unmet         allow_bypass         declared false  actual 2 grant(s) may push without a merge request
met           approvals.required   declared 1      actual 1
not-available approvals.not_by_author declared true   actual not available
```

Three states in one report, in this host's own words. Declaring the waiver turns the fourth row to
`waived` and drops the unmet count from 2 to 1, and declaring `merge_method` adds the `unknown` line
from the domain. So the whole of A65's split was exercised end to end with a second host: the adapter's
states, the domain's filter, the waive logic and section 13's unchecked line.

**Criterion 5, and the observation changed the code.** With a real token on gitlab.com:

| endpoint | anonymous | valid token, not a member |
|---|---|---|
| `/projects/:id` | 200, merge settings **absent** | 200, merge settings **present** |
| `/projects/:id/approvals` | 401 | **403** |
| `/projects/:id/approval_rules` | 401 | **200** |
| `/protected_branches/:branch` | 200 | 200 |

Two things follow and neither was in #104.

The merge settings being absent anonymously and present with a token confirms the `*bool` from both
directions. This is the defect P3 records, and it is now verified in the state that would have hidden
it as well as the state that exposed it.

The 403 is not the tier. `gitlab-org/gitlab` is on the richest tier there is, and the refusal is that
this token is not a member with enough access. So a single status code is ambiguous between a tier
without approval rules and a token without permission, which invalidates the table-of-codes approach
#104's row implied and vindicates P2's decision to distinguish only readable from unreadable. The
reason text asserted the tier and now names all three possibilities, because sending somebody to a
billing page over a permission is the kind of wrongness a report is judged on.

**Criterion 8.** The GitHub report is identical to `main`'s, `checked_at` aside.

| | |
|---|---|
| `internal/host/gitlab` | 90.2% |
| `internal/host` | 93.3% |
| `internal/host/github` | 93.5% |
| build, vet, test, gofmt | clean |
| `gate verify` | exit 0 |
| A42 | `internal/gates` reaches neither `net` nor `net/http` |

<!-- xeno:section:gaps -->
## Gaps

**The tier itself is still not observed, and this is the honest residue of the fixture question.** What
was observed is a permission refusal on a project on the richest tier. A Free project whose owner asks
about its own approval rules would answer something, and whether that is 403, 404 or 402 is unknown.
The adapter handles all three identically, and the reason text no longer claims which cause it was, so
the gap costs precision in a sentence rather than correctness in a state. That is a better place for it
than #104's table, which would have encoded one guess as a mapping.

Closing it needs a Free project the token owns. The account this token belongs to has none, and
creating one is a change to somebody's GitLab account rather than a step of this intent.

**`approvals.required` sums the rules, and the sum is a reading.** Two rules requiring none and one
give one, which the live run confirms. Where rules overlap in their eligible approvers the sum
overstates what a merge request actually needs, because one person may satisfy two rules. GitLab's own
model is the authority and this adapter does not consult it; the alternative, reporting the largest
single rule, understates instead. Neither is wrong enough to matter at the counts a declaration uses,
and it is recorded rather than left for somebody to find at a count where it does.

**`merge_access_levels` is decoded and unused.** It is in the payload and in the struct and no
requirement reads it. GitLab's merge permission is a third kind of thing the declaration has no word
for, and inventing one is forbidden by the second standing rule. Keeping the field is a small
dishonesty of the kind a reader trips over; removing it loses the record that the host answers it.

**Two rows are verified in states no real project showed.** `approvals.required` unmet and
`approvals.not_by_author` in both its states come from payloads. The live project had them met and
refused respectively.

**The wrapper and the adapter have still not been exercised together.** #145 generated a
`.gitlab-ci.yml` nobody has run, and this intent's command was run from a shell. A real merge request
on a real GitLab project is what tests the pair, and it remains the one thing no amount of work here
substitutes for.

**Not a gap.** That `xeno enforcement check` was run against a project the token does not own. It is a
real project, a real answer and the real command, which is what criterion 3 asked for.
