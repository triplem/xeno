---
intent: github.com/triplem/xeno#162
phase: 00-intake
created: "2026-10-01T15:50:24Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9f9c5ff207244f7fd2dc88adb5ca186a639c591f3ab078ad0945c5a1b08ed728
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
#160 left a registry with no entries, so `kind: checked` is worse than useless today: a checked rule
turns every phase it applies to red for want of an implementation. Five types are named and none
exists, `section-implies-section` plus the four that read the commit history, and section 9 gives each
one's input in a table, so none of them is a design question. One of the two shipped patterns is also
missing — the Conventional Commits form carrying an issue reference in the subject, which a project
needs when the reference has to survive a squash merge. Nothing here reads a commit: `Base` and `Head`
sit on the runner with a comment saying no gate reads them until WP4, and they do not reach the gate
context. That collides with `internal/gates`'s own comment, which says gates never run anything, where
what it was written to mean is no build, no test and no model; the sentence is wrong as it stands and
this piece corrects it rather than quietly violating it. Two inputs, two failure modes: the range is
passed in and never inferred, so a rule needing a range it did not get is its own red case, and
`approver-not-author` reads a previous run's `gate.yaml`, which makes it the only predicate whose
verdict depends on who ran a command rather than on what the tree holds.
