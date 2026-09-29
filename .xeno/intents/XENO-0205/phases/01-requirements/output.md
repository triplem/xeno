---
intent: github.com/triplem/xeno#65
phase: 01-requirements
created: "2026-09-29T11:32:39Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+edd2b2e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bb7d792a5e950440ff86acd163b3ab92f6ef0e50ea172e8ab27c7e0a2d936b69
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

**AC1.** `xeno cost turn` reads a hook's JSON on standard input, takes `transcript_path`
and `session_id` from it, and appends one line to `.xeno/local/cost-ledger.yaml`: the
session, the cumulative token totals over that transcript, the live phase or none, and a
timestamp.

**AC2.** It reads counts and identifiers and nothing else. No message content, no file
path from the conversation and no prompt text reaches the ledger.

**AC3.** The live phase comes from `.xeno/phase.env`. A turn with no phase open records
`phase: none` rather than being dropped.

**AC4.** It never fails a turn. Unreadable JSON, a missing transcript, a transcript with
no usage and an unwritable ledger all exit zero and leave the session alone, because a
hook that breaks the harness is worse than a missing figure.

**AC5.** `phase finish` writes `cost.yaml` with the fields section 11 names, `evidence:
self-reported` among them, the token totals attributed to that phase, and the sessions
they came from.

**AC6.** The totals are the difference between the first and last ledger lines that name
the phase, per session, summed across sessions. Consecutive lines make a turn's cost
arithmetic rather than a record.

**AC7.** A phase with nothing attributed writes no `cost.yaml`, since section 11 says a
phase without one is complete.

**AC8.** `cost.yaml` is outside `artifacts_hash`, which it already is, so writing one
does not change any verdict. `gate verify` matches every verdict before and after.

**AC9.** The hook is shipped under the plugin and wired in `.claude/settings.json` for
this repository, so that the next phase of this intent records its own cost.

**AC10.** No gate requires the file. A19 stands and says why.

<!-- xeno:section:non-goals -->
## Non goals

A gate over `cost.yaml`. #65 asks for one and A19 is the row it changes. Requiring the
file while attribution depends on a workflow nobody has adopted would fail every phase
for a figure the tool cannot honestly produce.

Codex. The transcript read here is Claude Code's. Section 14 names both clients, the
reader is chosen by the `agent.tool` the project records, and a second one is additive
work.

`cost_usd`. Section 11 makes it optional and puts the authoritative money view in v2.
The transcript offers `totalCostUSD` and it is a price nobody in the repository agreed
to.

Retention. Section 12 ties pruning to the presence of `cost.yaml` and that is its own
work.

Splitting a turn between phases. A turn that opens a phase, works and closes it is
attributed to the phase that was live, and a turn that touches two is attributed to one.
No fraction is invented.

A figure for the intents already finished. The ledger starts empty, so nine intents have
no cost record and will not get one. #117 says so rather than this backfilling it.

<!-- xeno:section:constraints -->
## Constraints

Section 11 fixes the schema and section 4 the file's place. Nothing under `docs/`
changes.

The hook must be harmless. It runs on every turn of every session in this repository, so
it exits zero on every failure and writes nothing it cannot write. A tool that makes the
harness unusable to record a number has made the wrong trade twice.

Only counts and identifiers leave the transcript. A transcript is the whole
conversation, and a ledger in a repository is a place a conversation must not end up,
even gitignored.

`.xeno/local/` for the ledger, by A39: a job status and never a commit. `cost.yaml` is
committed and sits outside `artifacts_hash`, by section 11.

A new subcommand is a decision. The command surface is small on purpose, section 11
names no command, and `cost turn` exists because a hook needs something to call. It is
recorded as an assumption.

Attribution is never invented. Where the runner cannot say which phase a turn belonged
to, the ledger says so.

One intent, one issue. The commits reference #65.
