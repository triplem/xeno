---
intent: github.com/triplem/xeno#147
phase: 03-implementation
created: "2026-09-30T21:22:14Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0a51653.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2be5394f259df3e9a33421f74c07da13f0bb1735582aa7a218f91ca418d41e25
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

**`internal/host/gitlab`, new, and one entry in `branchRules`.** That entry and an import are the
whole of the change to `internal/host/host.go`. `BranchRules` itself is byte identical to `main`'s,
which is criterion 2 and the finding this intent exists to produce.

**Three requests, each failing locally.** `protected_branches/:branch` answers `allow_bypass`,
`projects/:id` answers `required_pipeline`, and `approvals` with `approval_rules` answer the two
approvals rows. A request that could not be read degrades its own requirements to `not-available` with
the status in the reason, and the others are answered anyway.

**`allow_bypass` reads four things and reports the largest.** An `unprotect_access_levels` grant
first, then `allow_force_push`, then a `push_access_levels` grant above zero, then met. Access level 0
is "no one" on this host, so an explicit entry at 0 is not a bypass. The words name which was found.

**A rejected token is an error for the whole adapter, consistently.** A 401 from any of the three
requests. The first draft swallowed it on the approvals path into a `not-available` reason, which
contradicted this intent's own design section and would have reported a bad credential as a tier
limit.

**`only_allow_merge_if_pipeline_succeeds` is a `*bool`.** This is the defect the live run found and it
is the one worth reading this section for. This host omits the merge settings rather than sending
`false` where the caller may not read them, so decoded into a `bool` an absent setting becomes off and
the requirement is reported **unmet** — a setting nobody could read reported as a setting nobody made.
That is the exact collapse the three states exist to prevent, and every unit test passed with it in
place. It is `not-available` now, with a reason naming the field.

**Tests, twelve cases.** The four `allow_bypass` findings and the level-zero case as a table; the
pipeline setting on and off; the omitted setting; a tier refusing the approvals endpoints at 403, 404
and 402, asserting both that the two rows are `not-available` with a reason and that the other
requirements survive; the approvals rows compared where the tier has them; an unprotected branch; a
rejected token; a missing project; no invented names.

**A real payload as a fixture.** `testdata/protected-branch.json` is `gitlab-org/gitlab`'s protected
`master`, fetched unauthenticated. A hand written fixture would assert the field names this adapter
already believes.

**Two live tests, skipped unless `XENO_LIVE` is set.** One asks gitlab.com unauthenticated about the
protected branch and asserts the pointer defect stays fixed; the other runs the whole adapter and needs
a credential.

<!-- xeno:section:deviations -->
## Deviations from the design

**Criteria 3 and 5 are not met in this phase, and they are the two that need a credential.** Criterion
3 is that `xeno enforcement check` answers for a real GitLab project, and criterion 5 is that the
tier-gated status codes are observed rather than guessed. Both need a gitlab.com token, which this
machine has none of and which was requested. Everything not needing one was written first, which is
what P0 said the intent would be arranged for, and P4 is where the two are settled. They are recorded
here as outstanding rather than quietly deferred, because an intent that sealed its implementation
while claiming an observation it had not made would be the failure this process exists to catch.

What is verified live and unauthenticated: `allow_bypass` against `gitlab-org/gitlab`, which reports
unmet with "2 grant(s) may push without a merge request", and the absence of the merge settings from
an unauthenticated project read. Both are real answers from the host.

What is verified against payloads only: the approvals rows in their met and unmet states, and the
three refusal codes. The codes are handled uniformly, so criterion 5's risk is not that a code is
mishandled but that the reason text names a status whose meaning was never confirmed.

**The 401 handling changed mid phase and the design section is right rather than the first draft.**
P2 decided a rejected token is an error for the whole adapter. The approvals path swallowed it into a
reason, which the live run surfaced: an unauthenticated probe reported "the token was rejected" as the
reason two requirements were unavailable, which is true of no token and would be misleading of a real
one. Now consistent.

**The live tests are skipped by default and are not a criterion's evidence.** They are a way to ask
the host, not part of the suite. A test that needed a network would put the network where A42 spent a
CI check keeping it out, and their being skipped is why they may exist at all.

**Nothing else deviates.** No change to `BranchRules`, the domain, the GitHub adapter, the wrapper, the
scaffold or the gate path. No shared code with the other adapter, which P1's non-goals forbade.
