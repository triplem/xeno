---
intent: github.com/triplem/xeno#215
phase: 01-requirements
created: "2026-10-03T20:22:45Z"
schema_version: "1.0"
runner_version: dev+b626f1a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 817e5574286ffc47477569a73e4dcc4a8dde34c8af9ff3b719bf8d426475feef
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

`xeno phase start` refuses a phase that has a verdict, and the refused call leaves the lock
byte-identical — a refusal that had already written something would be the bug with a
message attached.

The refusal names the alternative. Redoing the work is `section set` and `phase finish`,
which never needed a second start, and starting over from nothing is removing the verdict
first. A refusal that only says no is one somebody works around by deleting whatever is in
the way.

Removing the verdict deliberately still works. The clause being enforced is about sealed
artifacts, and a phase with no verdict has nothing sealed, so the escape is not a loophole.

It reads like the three refusals `Start` already makes, and is tested the way they are.

The issue's first claim is corrected on the issue itself, not only here, because a wrong
claim left standing in a closed issue is read later as fact.

<!-- xeno:section:non-goals -->
## Non goals

No mechanism for reporting what changed. `ChangedSince` and its printer exist, are derived
rather than recorded for reasons their own comments give, and the duplicate started against
them was deleted rather than finished.

No specification change. Section 11 already says what is sealed is never rewritten; this
gives that clause a reader at the point of the act rather than only afterwards.

No change to `gate verify`. It stays the last line, and it is what caught this twice.

Nothing about the context profile, whose absence makes `ChangedSince` silent throughout this
repository. Its own issue.

No new field, gate, tool or rule.

<!-- xeno:section:constraints -->
## Constraints

A refusal is a behaviour change to a command the whole trail is produced with, so the
question is not only whether it is right but whether anything legitimate relied on the old
behaviour. One thing did, and it is kept: `Start`'s existing refusal tells a reader whose
run died to remove the marker and start again, and that path leaves a phase with a lock and
no verdict. The new refusal keys on `gate.yaml` rather than on the lock, which is what keeps
the recovery open.

The clause being enforced is in the process definition, so the wording of the refusal is
constrained by it: it may not invent a rule, only name the one that already applies.

Prose at 88, tables exempt. Go at whatever `gofmt` produces, per #210.
