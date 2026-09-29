---
intent: github.com/triplem/xeno#141
phase: 00-intake
created: "2026-09-29T19:50:49Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6aa6f9b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8d9c57b0a1d3a243dd862cdd493fa3cbba863e7030cf13e400c0184c5b1d5661
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

#97 split `Fetch` into transport and decode and recorded a design for the host port, and
left one thing open on purpose: whether `Protection` gains per-field expressibility,
generalising `ReviewsExpressible`, or whether an adapter returns `[]Requirement` and
`Compare` shrinks to the waive logic. Its reason for leaving it was explicit. Choosing
between the two readings without a second host in front of us "is how a neutral interface
acquires fields nobody needs", so the decision was scheduled for when GitLab is written.

#104 carries that open question as the second of its four pieces and adds information #97
did not have: #99 turned `Compare` into a table of requirements, which makes the second
reading the smaller change rather than the larger one.

The problem is that the other three pieces of #104 wait on this one. The port in piece 1 is
`BranchRules`, and the vocabulary *is* that port's method signature; drawing it before the
vocabulary is settled risks drawing it against a reading that is then rejected. The
enforcement mapping in piece 3 is five rows of declared-to-host correspondence, and what a
row can express depends on what an adapter is allowed to return. So the piece that #97
scheduled last is the one the sequence needs first, and the condition #97 set for deciding
it — a second host in front of us — has to be met some other way than by writing GitLab.

<!-- xeno:section:scope -->
## Scope

One row in `ASSUMPTIONS.md` recording the vocabulary decision, the reading it rejects, the
evidence that made it decidable, and the cost it accepts.

**In scope.** The decision and its record. Naming which of the two readings #97 left open is
taken, on evidence rather than on preference, and naming what the rejected one would have
cost. Recording the central name list as the thing the domain keeps, so that the second
standing rule still binds an adapter that now supplies its own words.

**Out of scope, and deliberately.** Every line of code. `internal/host` and the retirement
of `Protection` are piece 1 of #104; the five row enforcement mapping is piece 3; the
wrapper and the tracker are piece 4. A decision that arrives together with the refactor that
expresses it cannot be reviewed as a decision, which is the same reason #97 recorded its
design in a comment and built none of it.

Also out of scope: the two rows of #104's table that need an authenticated Premium
namespace. What can be verified without a token is verified here and what cannot is left to
piece 3 with the gap named, rather than guessed at to make this phase look complete.

<!-- xeno:section:context-rationale -->
## Why this context

The information base is #97's closing comment, which is where the open question is stated in
its own words and where the design it belongs to is recorded; #104, which carries the
question forward with #99's table as new information; and `internal/enforcement/enforcement.go`,
which is the code the two readings describe — `Protection` with its `ReviewsExpressible`
field, the `requirements` table, and `Compare`.

The fourth source is the host, and it is the one #97 said was missing. GitLab answers
`GET /api/v4/projects/:id/protected_branches` on a public project without a token, so the
shape of a second host's answer can be read from gitlab.com as it is rather than from
documentation about it. That is what turns this from the abstract choice #97 declined into a
decision against evidence, and it is why this phase can run before piece 1 rather than after.

Nothing here reads the process definition or the plan beyond section 13, which fixes the
requirement names, and section 8, which fixes that the GitLab meant is the Enterprise
variant. Both are read as constraints on the decision, not as things this intent may change.
