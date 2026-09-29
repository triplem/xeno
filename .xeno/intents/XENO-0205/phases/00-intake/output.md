---
intent: github.com/triplem/xeno#65
phase: 00-intake
created: "2026-09-29T11:31:26Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+edd2b2e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 05940a5d937a835ba77f5b302123bf61766b1a5ff28628ca4738ff1d6ab658d9
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

Section 4 lists `cost.yaml` among the six files a phase holds and section 11 gives its
schema. Nothing writes it. A19 records why no gate requires it — a file nothing produces
would fail every phase — and #65 says what has to change: a writer, then a gate, then
the assumption closed rather than left standing.

It is also the figure #117 is waiting for. The plan settles the proportionality question
on three numbers, cost per intent among them, and that one has no source at all, so nine
intents have now been measured on the two that are derivable from the tree.

**What the harness makes available.** Claude Code writes a session transcript as JSONL,
and every assistant message carries a `usage` object with `input_tokens`,
`output_tokens`, `cache_read_input_tokens` and `cache_creation_input_tokens`, beside a
`sessionId` and a timestamp. Every field section 11 asks for exists, including
`totalCostUSD` for the optional one. Hooks receive `transcript_path`, `session_id` and
`cwd` on every event and receive no token counts, so a hook reads the transcript rather
than being told.

**What the obvious attribution does.** Section 11 assumes one session per phase. In
practice a session spans many: this one has carried every intent since #106. Filtering
the transcript by each phase's own window, from the `created` of its lock to the
`run_at` of its verdict, and summing what falls inside, captures **7.1% of the output
tokens and 5.4% of the cached** over the fifty-four phase windows of this session.
Fifty-one minutes of windows against days of work.

**And why no rule fixes that.** The tokens are spent composing a phase's content, which
in this session happens before `phase start` is called: the phase is opened, the prose
is written into it from what already exists, and the phase is finished, all inside one
turn. A window, a delta of counters, or any other rule over the same transcript measures
the same 7%, because the work genuinely happened outside the phase. Section 6's working
sequence has the agent start the phase and then work; that is the assumption cost
recording rests on, and it has not been held to here.

So a writer built from the obvious reading would understate the cost by about fourteen
times, in one direction, and feed that into the decision #117 exists to inform.

<!-- xeno:section:scope -->
## Scope

A ledger, a hook that appends to it, and `phase finish` writing `cost.yaml` from it.

`xeno cost turn` reads a hook's JSON on standard input, reads the transcript it names,
sums the usage of every assistant message in it, and appends one line to a ledger under
`.xeno/local/`: the session, the cumulative totals, the phase that was live, and when.
The difference between consecutive lines of a session is what that turn cost, so the
arithmetic needs no notion of where a turn begins.

The phase that was live comes from `.xeno/phase.env`, which `phase start` already writes
and `phase finish` already removes. A turn with no phase live is recorded as
unattributed rather than dropped, which is what makes the gap visible instead of
silently absent.

`phase finish` sums the ledger for its own phase and writes `cost.yaml` with the fields
section 11 names, `evidence: self-reported` among them, and the sessions it drew on. A
phase with nothing attributed writes no file, because section 11 says a phase without
one is complete.

The hook is shipped under the plugin and wired for this repository in
`.claude/settings.json`, so that the thing being built is used by the project building
it.

Not a gate. #65 asks for one eventually and A19 is what it would change; requiring the
file is a separate step once there is something to require, and a gate demanding a
figure a batching workflow cannot attribute would be the earlier mistake with a verdict
attached.

Not Codex. Section 14 names both clients and this reads Claude Code's transcript; the
reader is selected by the `agent.tool` the project records, so a second one is additive.

Not `cost_usd`. Section 11 makes it optional and says the authoritative money view is
v2's.

Not the retention rule of section 12, which touches this and is its own work.

<!-- xeno:section:context-rationale -->
## Why this context

**The hook exists because nothing else can see a token.** Hooks receive
`transcript_path` and no counts, the runner is invoked between turns and sees nothing of
them, and the gate path is refused the network. A hook firing once per turn, reading the
file the harness already wrote, is the only place the number is available at all.

**Cumulative totals rather than per turn deltas, because a turn is not a unit the
transcript marks.** Summing the whole transcript on every turn and recording the total
makes each line self-contained: the delta is arithmetic between two lines, and a missed
turn leaves a larger gap rather than a lost one. It also means a ledger line is
verifiable against the transcript at any time, which a recorded delta would not be.

**The live phase decides attribution, and a turn without one is recorded as such.** This
is where the measurement of the problem becomes the design. A rule that guessed —
nearest phase, last finished phase, the window — would produce a number for every turn
and be wrong for most of them. Recording the truth, that a turn was spent with no phase
open, puts the 93% where a reader can see it, and makes the figure #117 needs honest
even while it is incomplete.

**Which is a statement about how to work, not only about a writer.** Section 6 has the
phase started before the work begins. Attribution is exact for an agent that does that
and empty for one that composes first, so the ledger measures adherence as much as cost.
That is the useful half of this intent and it is not a feature.

**`evidence: self-reported` is not a hedge here, it is accurate.** Section 11 says the
v1 numbers are indicative and must not enter certification records without
re-measurement. A count read out of a local file the harness wrote, attributed by a
marker the runner wrote, is exactly that.

**The ledger is gitignored, the record is committed.** `.xeno/local/` is where a job
status belongs, by A39, and `cost.yaml` is a phase file that sits outside
`artifacts_hash` so that a figure arriving after the verdict cannot invalidate it, which
section 11 says outright.

**No gate yet, deliberately.** A19 stands until there is a figure a phase can be
required to have. Requiring one now would fail every phase whose work happened outside
its window, which is every phase this session has produced.
