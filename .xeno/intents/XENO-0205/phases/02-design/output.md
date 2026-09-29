---
intent: github.com/triplem/xeno#65
phase: 02-design
created: "2026-09-29T11:33:20Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+edd2b2e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d6a5acfc191e1da39dad880bbbf8d97a4105dc501fc4d664e6a205fd6c1305b2
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

**`internal/cost`, and `xeno cost turn` on top of it.** The package reads a transcript
and sums usage, reads and appends the ledger, and totals a phase. The command is the
hook's entry point: JSON in on standard input, a line appended, exit zero whatever
happens.

**A `Stop` hook, which fires once when a turn ends.** It is the only event that
corresponds to work having been done, it carries `transcript_path`, and it is per turn
rather than per tool call, so the ledger has one line per turn instead of hundreds.

**Cumulative totals per line.** Each line records the sum over the whole transcript at
that moment, so a line is verifiable against the file it came from and a turn's cost is
the difference between two lines. A recorded delta would be neither.

**The phase from `.xeno/phase.env`, or `none`.** No inference. The runner writes that
file at `phase start` and removes it at `phase finish`, so it is the only statement
about which phase was open that anybody made.

**`phase finish` totals before it writes the digest and the verdict.** The order matters
for one reason: `cost.yaml` is outside `artifacts_hash`, so it may be written at any
point without touching the seal, and writing it first keeps the finish a single pass.

**A phase with nothing attributed writes no file.** Section 11 says a phase without one
is complete, so absence is the honest value and a zero would be a claim.

**Everything in the command fails open.** Every error path returns exit zero. The hook's
contract with the harness is that it is invisible.

**Counts and identifiers only.** The reader takes `usage` and `sessionId` from a
transcript record and never touches `message.content`, which is the one rule that keeps
a conversation out of a repository.

<!-- xeno:section:alternatives -->
## Alternatives

**Filter the transcript by each phase's window at `phase finish`.** The obvious reading
of section 11 and no hook at all. Measured: 7.1% of the output tokens across fifty-four
windows, because the work happens before `phase start`. Rejected on the measurement, and
it is recorded rather than merely rejected, because it is the design anybody would reach
for first.

**A `PostToolUse` hook on the xeno commands only.** Fewer lines in the ledger and it
fires where a phase boundary is crossed. It attributes nothing: the tokens were spent in
the turn that called the command, and a tool-call hook sees the transcript only up to
the last completed message.

**Snapshot cumulative counters at `phase start` and `phase finish` and subtract.** No
hook, no ledger, and exactly the same number as the window, for exactly the same reason.
Named because it sounds like a different mechanism and is the same measurement.

**Distribute a turn's cost across the phases it touched.** It would give every phase a
figure. The split would be invented, and an invented number in a file a decision reads
is worse than a gap the decision can see.

**Attribute an unattributed turn to the phase finished most recently.** Plausible, since
a turn that composes and then writes a phase is doing that phase's work. It is also a
guess that would have attributed this session's composing turns correctly and a week of
investigation turns wrongly, with no way to tell the two apart. Rejected, and the
ledger's `none` is what replaces it.

**Have the hook write `cost.yaml` directly.** One fewer step. The hook would then write
inside a phase directory on every turn, which is the runner's job and nobody else's, and
a phase already sealed would be written to.

**A shell or Python hook script instead of a subcommand.** No new command in the
surface. It puts JSONL parsing in a second language, untested, in a script that runs on
every turn, which is the place least able to afford either.

<!-- xeno:section:impact -->
## Impact

`internal/cost`: the transcript reader, the ledger, the phase total. Its own tests.

`cmd/xeno`: `cost turn`, and a line in the usage text.

`internal/runner`: `Finish` writes `cost.yaml` where the ledger has something for the
phase.

`.xeno/plugin/hooks/`: the shipped hook fragment, and `.claude/settings.json` wiring it
here.

`ASSUMPTIONS.md`: the new subcommand, the attribution rule, and what the ledger's `none`
means.

What a reader gains: the first of the three figures #117 needs that is not derivable
from the tree, for phases worked the way section 6 describes, and a visible count of the
turns that were not.

What they do not gain: a figure for the nine intents already finished, a gate over the
file, a Codex reader, or money.

What this costs: a hook on every turn in this repository, a gitignored ledger that grows
by a line a turn, and a number that will be small until the way of working changes —
which is the finding rather than a defect.

The risk worth naming now: the ledger will show that most of this project's cost is
unattributed, and that is an uncomfortable measurement about the agent rather than about
the tool. Recording it is the point.
