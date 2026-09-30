---
intent: github.com/triplem/xeno#147
phase: 01-requirements
created: "2026-09-30T21:14:47Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0a51653.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0b699131a10e85e88f931272d73a4d78f37b8cbef15e0898855a2b256c265bfe
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. `internal/host/gitlab` implements `host.BranchRules` and is one entry in `branchRules`.
   `xeno init --host gitlab` already writes `adapter: gitlab`, so nothing in the scaffold changes.

2. **`host.BranchRules` is unchanged.** Not one field, not one parameter. This is the criterion the
   intent is evidence for, and #143's P5 wrote in advance what a failure means: A65 was wrong rather
   than refinable, and the intent stops and records it.

3. `xeno enforcement check` answers for a real GitLab project. Not a fixture: the command, against
   gitlab.com, with a token.

4. Every requirement the tier cannot express is `not-available` carrying a reason, never `unmet` and
   never an error that stops the command. A Premium setting absent on Free is nobody's oversight,
   which is the distinction the report exists to make.

5. The tier-gated status codes are **observed**, and the mapping is written from what was seen. 401,
   403, 402 and 404 are all plausible answers from `/approvals` and `/approval_rules` and #104 guesses
   none of them; whichever arrives is handled as `not-available` rather than as a transport failure.

6. `allow_bypass` reports which of four findings it made, in GitLab's words: a push grant above "no
   one", a force push allowance, an unprotect grant, or none of them. "Administrators may bypass" is
   GitHub's sentence and must not appear.

7. The adapter invents no requirement name. Asserted the way the GitHub adapter's is, against
   `enforcement.Names()`.

8. The GitHub path is unchanged, checked by running `xeno enforcement check` against this repository
   and diffing the report against `main`'s binary.

9. `go build`, `go vet ./...`, `go test ./...` clean; `gofmt -l .` outside `vendor/` silent;
   `xeno gate verify` exit 0.

<!-- xeno:section:non-goals -->
## Non goals

**No change to `host.BranchRules`.** Criterion 2 makes this an outcome rather than a preference. It is
listed here as well because the temptation in a second adapter is to widen the interface by one field
to make the awkward row fit, and doing that quietly is the failure the whole ordering of #104 was
arranged to prevent.

**No change to the GitHub adapter.** Not to share code with it, either. The two hosts answer the same
five questions and nothing about how one reads its own payload helps the other; a shared helper
between them would be the neutral struct arriving by the back door, one function at a time.

**No `merge_method` comparison.** Section 13 says unchecked, the GitHub adapter reports it as such
from the domain, and GitLab's richer answer — `merge_method` plus `squash_option` — is not a reason to
start. Comparing it is a spec change first.

**No tracker adapter.** WP12's, for both hosts.

**No Premium fixture.** The two approvals rows are verified in their `not-available` state against a
real Free project and in their met and unmet states against payloads. That is the bound P0 set and
criterion 4 is written to match it: what a Free tier reports is the honest answer for a Free tier, and
an adapter that refused to answer until somebody paid would be worse than one that says what it cannot
see.

**No retry, no caching, no rate limit handling.** The command is run by a person or by a daily
schedule, and a second request is a second run.

<!-- xeno:section:constraints -->
## Constraints

**`host.BranchRules` binds absolutely.** The signature is A65's decision made concrete by #143, and it
sits inside two sealed phases' `artifacts_hash`. Changing it here is not an edit, it is a contradiction
that has to be recorded.

**`enforcement.Names()` binds the names.** Section 13's five, and the second standing rule makes that a
budget. The adapter chooses words, never names.

**The states are the domain's four.** `Met`, `Unmet`, `NotAvailable` and, by not being the adapter's to
give, never `Waived`. A tier limit is `NotAvailable`; a setting somebody could have made and did not is
`Unmet`. Collapsing them is what makes the report worthless, which the domain says in as many words.

**Section 8 binds the reading of the tier.** The GitLab meant is the Enterprise variant. An adapter
written to satisfy Free alone would be written against the weaker of the two, so the Premium paths have
to be implemented from the reference even though they can only be tested from payloads.

**GitLab's permission model binds `allow_bypass`.** Access levels are numbers with meanings — 0 is no
one — and a grant carries a user, a group or a member role. The adapter has to read the numbers, so the
meaning it assigns them is a decision to write down rather than a constant to inline.

**A42 binds the import graph.** A third package under `internal/` using `net/http`, and the gate path
still reaches none of it.

**Project identity is URL encoded.** GitLab takes `namespace/project` as `namespace%2Fproject`, which
is a difference from GitHub's path that the adapter owns and nothing else sees.

**Eighty-eight columns, SPDX, no copyright line.**
