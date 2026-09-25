---
intent: github.com/triplem/xeno#41
phase: 00-intake
created: "2026-09-25T20:42:07Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+81b8426
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: b418bc256843b67ca9ed729ca284ccf21c925f68b6817148af133ceb9ffb13c9
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

Intent directories are named for their key, and a key is `XENO-` plus the issue
number. Anything that sorts by name puts `XENO-11` before `XENO-2`, so the trail is
listed in an order that is neither the order it was written in nor any other order a
reader has a use for.

<!-- xeno:section:scope -->
## Scope

The issue numbers of new intents are padded to four digits, `XENO-0049`. Existing
intents keep their names, and this intent keeps its own: the rule holds for intents
created after it is merged, which this one is not.

The rule goes in `CLAUDE.md`, next to the conventions it belongs with, because that is
the file read before an intent is created. No check enforces it.

<!-- xeno:section:context-rationale -->
## Why this context

**The obvious repair was measured and rejected.** Renaming `XENO-9` to `XENO-0009`
and re-running `gate verify` reports `DIVERGENT … artifacts changed after the verdict
was written`, because Appendix B hashes `sha256  <path>` with the path relative to the
repository root, and the key is in that path. §14 states the consequence outright:
moving or renaming an intent directory changes every verdict in it.

Re-sealing would produce verdicts again, but eleven keys are written into merge commits
already, and §14 makes the key the join from a merge commit back to its intent. Git
history does not move, so that join would break for everything past. The sortable half
of the trail is not worth an unreadable half.

That leaves two formats side by side, which is the price paid knowingly: the padded
ones sort correctly among themselves and group ahead of the old ones, and the old ones
stay reachable from the commits that name them.

The key is a project convention rather than something the tracker hands over. §3 calls
it the short tracker key, `PROJ-123`; GitHub issues have no short key at all, so
`XENO-` and the width of the number are both ours to choose.

Nothing checks the format. A check would be a rule the process definition does not
carry, and the standing rule against invented rules outranks the convenience — which
does mean this one can be got wrong again, unlike the digest that #2 turned into a
gate.
