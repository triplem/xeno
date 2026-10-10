---
intent: github.com/triplem/xeno#359
phase: 04-verification
created: "2026-10-10T16:01:34Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0b4b5e12d3b9c8cd93279d8fa488f3239b78aaaa19af5ced1c438a4916770fae
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
Twelve criteria, all passing, and two of them pass against a figure the artifact that
states them got wrong. The host carries 33 labels and criterion 1 says 34; `00-intake`'s
decision record names four issues carrying `xeno-needs-decision` and the host says seven,
one of which that record omits and one of which it names wrongly. Both figures were
written from the session rather than read from a command, both are sealed, both are
corrected here, and the cause is the learning record of `03-implementation`. The set
comparison the criterion actually asks for passes in both directions. Every negative claim
carries a positive control: the one label comparison in `internal/model/identity.go`
against the same search for `ApprovedLabel`; the absence of any label write against an
enumeration of every HTTP method both adapters use, which is three GETs and one POST to
the comments endpoint each; the empty diff against the two normative documents against the
same command on a file this intent did change; the absence of `xeno-start` against the
five occurrences of `xeno-approved`; and the green strict site build against a deliberate
broken link that aborted it with exit 1. The label on #359 and the question beside it were
read back from the host rather than taken from an exit code, which is the practice the page
prescribes, and the carried-in claim that `gh issue edit --add-label` exits 0 without
applying did not reproduce under controlled conditions. The gaps are honest: nothing keeps
the page and the tracker in agreement, the question on #359 is open and no gate can see
that it is, and a redo of `03-implementation` sits between its first verdict and this
phase.
