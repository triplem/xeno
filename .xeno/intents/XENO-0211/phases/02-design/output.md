---
intent: github.com/triplem/xeno#141
phase: 02-design
created: "2026-09-29T19:53:29Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6aa6f9b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bed35f7a140271c0e0db6c991c3584942c75f9fb629f5eef9cf112ce63b6bc79
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

**An adapter returns `[]enforcement.Requirement`, and `Compare` shrinks to the waive logic.**
This is the reading #97 left open, taken in the second of its two forms.

The division it draws. A host adapter answers with the requirements it can speak to, each
carrying the host's own words in `Actual` and its own `State`, including `NotAvailable` where
the host has no equivalent. `internal/enforcement` keeps `Declared`, `Report`, `State` and
`Requirement`, and keeps three jobs: which requirements the declaration actually asks for,
applying the waivers, and the `merge_method` unchecked line that section 13 requires be
reported rather than left silent. What it loses is `Protection` and the `evaluate` half of the
`requirements` table.

**The names stay in the domain, the words go to the adapter.** The requirement names are
section 13's vocabulary and the second standing rule makes that list a budget, so the domain
keeps enumerating them and an adapter that returned an unknown name is a bug rather than an
extension. This is the line that keeps the decision from being a licence, and it is why the
whole table does not simply move.

**The state set does not change.** `Met`, `Unmet`, `NotAvailable`, `Waived` and `Unknown`
already exist and already carry the distinction this needs: a setting nobody made, a setting
the host does not have, and a decision already taken. No artifact gains a field, which the
second standing rule requires.

**Nothing is built in this intent.** The port that expresses this is piece 1 of #104. What
this intent writes is `A65`, and the reason is the one #97 acted on when it recorded its own
design and built none of it: a decision arriving with its refactor is reviewed as a refactor,
and this one has an alternative that deserves to be visible in review.

<!-- xeno:section:alternatives -->
## Alternatives

**`Protection` gains per-field expressibility, generalising `ReviewsExpressible`.** Rejected.

It is the smaller diff today and it was the more plausible of the two when #97 was written.
`ReviewsExpressible` already exists as exactly this device for one field, so generalising it
is a pattern the file already contains, and `Compare` keeps its table. #99 changed the
arithmetic: the table is now data with an `expressible` hook per row, so both readings are
reachable from here, and the one that removes a struct is no longer the larger change.

What rejects it is `allow_bypass`. GitHub answers it with one boolean, `enforce_admins`.
GitLab answers the same question with `push_access_levels`, `merge_access_levels` and
`unprotect_access_levels`, three lists in which each grant carries an access level and a
user, a group or a member role. Expressibility does not help: the field is expressible, and
the difficulty is that reducing a list of grants to "somebody may bypass" is a judgement
about that host's access levels. Under this reading that judgement sits in
`internal/enforcement`, the one package meant to know no host, or `Protection` grows a field
shaped like GitLab's answer and stops being neutral. Both are the outcome #97 named: a
neutral interface acquiring fields nobody needs.

`approvals.not_by_author` rejects it a second time. #104's table calls this row the
interesting one, and the code says why: the GitHub branch returns met because the host
refuses an author's own approval, which is met by construction. On GitLab
`merge_requests_author_approval` is a setting, so the same requirement has a reachable
`Unmet` that GitHub's does not. Two hosts whose answers to one name differ in which states
exist cannot share a struct of facts without one of them being flattened.

**A third reading, not in #97: one `Host` port carrying both concerns.** Rejected already,
and not reopened here. #97 decided two ports because `BranchRules` and `Tracker` serve
different packages, and this intent changes nothing about that.

<!-- xeno:section:impact -->
## Impact

**On this tree, one row.** `ASSUMPTIONS.md` gains `A65`. No Go file changes, no test changes,
no artifact gains a field.

**On piece 1 of #104, the signature.** `BranchRules` returns requirements rather than a
`Protection`, so the port's method is settled before it is drawn, which is what this intent
exists to do. `Protection` is retired there rather than here, and `ReviewsExpressible` goes
with it.

**On #99, a partial reversal in the same release.** #99 turned `Compare`'s switches into a
table, and this decision moves the `evaluate` half of that table into the adapters while the
names stay. That is the cost, and the row records it: two adapters can drift in the words
they use for one requirement name, where today one table holds every phrasing. The
alternative was worse, and the honest account is that #99's table made this reading cheap
rather than that it was aimed at it.

**On the GitHub path, no behaviour change is intended.** Its current phrasings are the
adapter's phrasings, moved rather than rewritten. Piece 1 has to assert that: the existing
`enforcement_test.go` cases are what pin the words, and they should pass unchanged against
the GitHub adapter or the move was not a move.

**On A13 and A39, nothing.** The pipeline artifact stand in and the report's gitignored
location are untouched.

**On the gate path, nothing, and that is load bearing.** A42 keeps the gates free of hosts
and #97 confirmed none of them needed to change. This reading makes that easier to hold,
since a host's words no longer travel through a struct the domain owns.

**What stays owed.** Two rows of #104's five row table are unverified, because the merge
settings are absent from an unauthenticated project payload rather than false and `/approvals`
answers 401. They need the Premium namespace, which is deliberately not arranged yet, and
they belong to piece 3. The protected branch shape, which is what this decision turned on, is
tier independent and was read from the host directly.
