---
intent: github.com/triplem/xeno#141
phase: 01-requirements
created: "2026-09-29T19:52:19Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6aa6f9b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d20e166614fd97c2b5eea74e7b24bd0b00b515a28f3b9be20851ae750bba65d7
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

1. `ASSUMPTIONS.md` carries one new row, `A65`, in the assumption table. It names the
   reading taken — an adapter returns `[]enforcement.Requirement` and `Compare` shrinks to
   the waive logic — and the reading rejected, `Protection` gaining per-field
   expressibility.

2. The row's reason is the deciding case and not a preference. `allow_bypass` is answered by
   one boolean on GitHub, `enforce_admins`, and by three lists of grants on GitLab, and
   whether a grant amounts to a bypass is a judgement about that host's access levels. The
   row says that the rejected reading puts that judgement either in `internal/enforcement`,
   which is the package meant to know no host, or in a field per host.

3. The row cites what made it decidable: `GET /api/v4/projects/:id/protected_branches`
   answers an unauthenticated request on a public project, so the shape was read from the
   host rather than from documentation. A reader can tell it was settled against a second
   host and not in the abstract, which is the condition #97 set.

4. The row records what the domain keeps: the requirement names stay enumerated centrally,
   because they are section 13's vocabulary and the second standing rule makes that a
   budget. An adapter supplies words and states, never a name.

5. The row records the cost rather than arguing it away: two adapters can drift in the words
   they use for one requirement name, and this undoes part of #99's table in the same
   release.

6. The state column says `approved`, and the row is cross referenced from the `Where` column
   to `internal/enforcement` and to #97 and #104.

7. No file outside `ASSUMPTIONS.md` and this intent's own phase directory changes. `go build`
   and `go test ./...` pass, `gofmt -l .` outside `vendor/` prints nothing, and
   `./xeno gate verify` exits 0.

<!-- xeno:section:non-goals -->
## Non goals

**The port is not written here.** `internal/host`, `BranchRules`, `Tracker`, and the
retirement of `Protection` are piece 1 of #104. This intent decides the vocabulary the port
will express and writes none of it, for the reason #97 recorded its own design in a comment
and built nothing: a decision that arrives with its refactor is reviewed as a refactor.

**The `api.github.com` default is not moved.** `internal/runner/enforcement.go:50` is named in
#104's first piece and stays where it is until that piece runs.

**The five row enforcement mapping is not confirmed.** Two of its rows need an authenticated
Premium namespace, which is deliberately not arranged yet. What this intent verified is the
protected branch shape, because that is what the decision turned on; the merge settings
(`only_allow_merge_if_pipeline_succeeds`, `merge_method`, `squash_option`) are absent from an
unauthenticated project payload rather than false, and `/approvals` answers 401. Those gaps
belong to piece 3 and are named rather than filled.

**No spec change.** Section 13's requirement names and section 8's Enterprise variant are read
as constraints. Nothing here proposes editing the process definition or the plan, and the
first standing rule puts such a change in its own commit before the code anyway.

**No second dependency and no new field in an artifact.** The second standing rule applies:
this decision adds no gate, tool or rule, and `Requirement` already exists with the `State`
set the decision relies on, including `not-available`.

<!-- xeno:section:constraints -->
## Constraints

**Section 13 fixes the requirement names.** `required_pipeline`, `allow_bypass`,
`approvals.required`, `approvals.not_by_author` and `merge_method` are the declared
vocabulary, and the second standing rule makes that list a budget. Whatever an adapter is
allowed to return, it is not allowed to name a requirement the specification does not have.
This is the constraint that decides what the domain keeps under either reading, so it is
what keeps the chosen one from being a licence.

**Section 8 fixes the edition.** The GitLab that is meant is the Enterprise variant, because
it has the approval rules the free edition does not. The decision therefore has to be taken
against the stronger tier's vocabulary even where this intent could only reach the weaker
one's endpoints. The consequence is a real limit on what was verified and is recorded as
such: the protected branch shape is tier independent and was read directly, and the approvals
rows were not.

**A42 keeps the gate path free of hosts.** No gate touches the network or a host, #97
confirmed nothing in the gates changed or needed to, and neither reading may alter that. The
chosen one makes it easier to hold rather than harder, since a host's words stop passing
through a shared struct.

**Appendix B makes the decision unrepeatable in place.** `artifacts_hash` covers the phase
directory and the intent key sits inside it, so this row cannot be revised by editing the
phase later; a change of mind is a new intent. That is the reason the acceptance asks for the
rejected reading and the cost to be written down now rather than summarised.

**`ASSUMPTIONS.md` is the only file that may change.** The first standing rule puts the
documents out of reach, and this record is not one of them: `CLAUDE.md` names it as where
assumptions and decisions taken while building the core go. The next free row is `A65`.

**Eighty-eight columns, in prose as in code.**
