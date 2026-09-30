---
intent: github.com/triplem/xeno#147
phase: 00-intake
created: "2026-09-30T21:14:02Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0a51653.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 24c37d56e0aa8ed1efbe2efbc77d3cd4cd6a1c23c0d2216b6ade5e4db525aa55
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

Since #145 a repository can be initialised for GitLab: it gets a wrapper that runs the gates and a
`project.yaml` naming `adapter: gitlab`. Nothing answers for that name. `xeno enforcement check`
refuses there and lists the adapters that exist, which is honest and is the whole of what GitLab
support currently means for the half of #104 that is about the host.

The port those rows would arrive through is #143's, and #143 could not test its own claim. A port is
an indirection until a second adapter fits it without changing it, and there was one adapter. P4 of
XENO-0212 recorded that as the gap and P5 recorded the consequence: if this piece has to change
`BranchRules`, #143 was wrong rather than refinable.

So this intent has two subjects. The first is the adapter, which is #104's remaining done-when. The
second is the port, which is not this issue's deliverable and is what the issue is evidence about.

What makes the adapter more than a transcription of GitHub's is the second row of #104's table.
GitHub answers `allow_bypass` with `enforce_admins`, one boolean. GitLab has no such field: it has
`push_access_levels`, `merge_access_levels` and `unprotect_access_levels`, three lists in which each
grant carries an access level and a user, a group or a member role, plus `allow_force_push`. Whether
any of that amounts to "somebody may bypass" is a judgement about this host's permission model, and
A65 exists because that judgement cannot live in the domain.

The other half of the difficulty is the tier. Section 8 fixes that the GitLab meant is the Enterprise
variant, because approval rules are Premium. `/projects/:id/approvals` and
`/projects/:id/approval_rules` answer 401 to an unauthenticated request on Free and paid projects
alike, so what a Free project's token sees at those endpoints is the fact this intent has to observe
rather than assume. It decides whether two of five requirements report `not-available` with a reason
or an error that stops the command.

<!-- xeno:section:scope -->
## Scope

**In scope.** `internal/host/gitlab` implementing `host.BranchRules`, one entry in `branchRules`, and
the five row mapping. The reading of `allow_bypass` from the protected branch's three grant lists and
`allow_force_push`, with words that name which of them was found. The tier-gated endpoints observed
with a real token, and whatever they answer mapped to `not-available` with a reason rather than to an
error.

**In scope as a finding rather than a change.** Whether `host.BranchRules` survives unchanged. The
interface is not this intent's to redesign; if it does not fit, the intent stops and records that
against A65, because #143's P5 said so in advance and a port quietly widened to fit its second
adapter is the failure this whole sequence was arranged to avoid.

**Out of scope.** The tracker adapter, still WP12's. `merge_method` remains uncompared, per section 13
and the same reasoning the GitHub adapter records. No change to the GitHub adapter, the domain, the
wrapper or the scaffold.

**Bounded by the fixture.** Two rows can be seen in one state only. `approvals.required` and
`approvals.not_by_author` are Premium settings, and on a Free project the honest answer is
`not-available`, which is what #104's second done-when asks for. Their met and unmet paths are unit
tested against payloads and not observed against a host. This is named in the acceptance rather than
discovered in review, and it is not a reason to hold the piece: an adapter that reports correctly on
the tier an adopter has is more useful than one that waits for a tier nobody here pays for.

<!-- xeno:section:context-rationale -->
## Why this context

The information base is `internal/host/host.go` for the interface this must satisfy without changing,
and `internal/host/github/github.go` as the one existing implementation — read for its shape, and
deliberately not as a template, since copying its `protection` struct would reintroduce the neutral
facts A65 rejected under a different name.

`internal/enforcement` is read for the name constants and the states, because the adapter builds its
answers from them and may not invent either.

A65 is read as the decision being tested, and #143's P4 and P5 for the claim this intent is evidence
about. #104's table is read as the starting point and not as fact: two of its five rows are marked
"to verify" and one of them, the range expressions in the fourth piece, already turned out to carry a
constraint the issue did not have.

Section 8 is read for the edition, which is the reason the tier is a subject here at all, and section
13 for the names and for `merge_method` staying uncompared.

Outside the tree: GitLab's protected branches API, observed live and unauthenticated on a real public
project, which is where the three grant lists and `allow_force_push` come from. GitLab's project and
approvals endpoints, whose unauthenticated status codes are known and whose authenticated ones are
not. The reference for the permission model's access levels, since reading a grant requires knowing
what 0, 30 and 40 mean.

A token is needed for the parts a public read cannot reach, and the intent is arranged so that
everything not needing one is written first.
