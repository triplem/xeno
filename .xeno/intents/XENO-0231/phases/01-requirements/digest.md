---
intent: github.com/triplem/xeno#190
phase: 01-requirements
created: "2026-10-03T14:02:09Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+74553ad.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ff172d799c19141e606ff9af60a47100bfdc96234f102a61cc91733f95d175f3
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
A pull request whose title is not a Conventional Commit fails the required check, with a message
saying the title becomes the subject and a prose one releases nothing silently; a title carrying the
host's reference passes, because the plain pattern accepts both and A71 defines the other pattern for
the form the host produces rather than the one a person types. The check calls `xeno check
commit-message`, so the rule has one implementation and cannot drift from `CLAUDE.md`'s. A pull
request that removes or renames any path under `.xeno/intents/` fails naming the paths; one that only
adds passes; a directory created and dropped inside one branch passes, because the comparison is a
tree diff against the base and not a walk of commits; and a base missing from the clone fails rather
than passes, since a comparison against something absent reports what a clean one reports — an error
made three times in one day here. Both steps sit in the `xeno` job, whose name is the required
context, because a check beside it is advisory and #186 is what that costs; both are skipped on a
push to main, where there is no pull request to read. Nothing touches the gate path. Out of scope:
modification, which `verify` names better; the re-seal hole, which the working sequence provides for
deliberately and which is stated rather than closed; the fifteen missed releases; G-Complete's
P5-only scope, which is what let XENO-0230 reach main two phases short; and a hook, whose result is
advisory and whose machine never sees the subject that matters.
