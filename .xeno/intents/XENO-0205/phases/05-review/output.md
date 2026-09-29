---
intent: github.com/triplem/xeno#65
phase: 05-review
created: "2026-09-29T11:42:12Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+edd2b2e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 55ae26c89b639d60882a03d61373ca67883e80defaca369d69a76f22f121f631
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**Nothing under `docs/` changed.** Section 11 gives the schema and section 4 the file's
place, and both described a record nothing wrote.

**Nothing was invented in the artifact.** `cost.yaml` carries the fields section 11
names, without `cost_usd`, which that section makes optional and v2 authoritative. The
two decisions the specification does not take are A63, the attribution rule, and A64,
the subcommand.

**The writer records what it could attribute and says so about the rest.** A turn with
no phase open is `none` in the ledger. That is the measured 7% problem made visible
instead of distributed, and it is the one design choice this intent would be wrong
without.

**The hook cannot break a session.** Four broken payloads, four exits of zero, run
rather than argued. Section 11 makes the record optional, so nothing it fails at is
worth a turn.

**No conversation left the transcript.** The reader takes four integers and two
identifiers. A ledger in a repository is not a place for a conversation, even a
gitignored one.

**No verdict moved.** `gate verify` over 109 with a `cost.yaml` inside a phase
directory, which is the check that section 11's placement outside `artifacts_hash` is
real.

**No gate requires the record**, so A19 stands, and a phase without one is complete by
section 11.

**The order of the work was right.** The 7% measurement answered whether the writer was
possible, P0 to P2 followed it, and the code followed them. First time since XENO-0201.

**A repair rode along and is admitted rather than hidden.** The duplicated `GateRun`
comment from
#125, in its own commit, with the deviations section saying why it is here and not in an issue.

**What is unproven is named.** The hook has never fired by itself, because settings are
read at session start.

<!-- xeno:section:release-notes -->
## Release notes

`cost.yaml` has a writer. `xeno phase finish` writes the record section 11 defines —
`tokens_in`, `tokens_out`, `tokens_cached`, `evidence: self-reported` and the sessions
it drew on — for phases that have something attributed to them.

The counts come from a hook. `xeno cost turn` reads a hook's JSON on standard input,
sums the token usage of the transcript it names, and appends one line to
`.xeno/local/cost-ledger.yaml` naming the phase `phase start` said was open. Hook input
carries no token counts, so reading the transcript the harness already wrote is the only
way to see one. The hook ships under `.xeno/plugin/hooks/` and this repository is wired
to it in `.claude/settings.json`.

It reads counts and identifiers and never a message. It exits zero on every failure,
because it runs on every turn and a phase without a cost record is complete.

**A turn spent with no phase open is recorded as such and counted nowhere.** That is
deliberate, and it is worth knowing before reading the numbers: attributing by each
phase's own window instead was measured at 7.1% of a session's output tokens, because
the work of composing a phase happens before `phase start` is called. So the record is
exact for a phase worked the way section 6 describes and absent for one composed first
and written afterwards. A63 carries the measurement.

A phase with nothing attributed writes no file, and nothing was backfilled: the nine
intents finished before the ledger existed have no cost record and will not get one.

No gate requires the file. A19 stands.

<!-- xeno:section:residual-risk -->
## Residual risk

**The wiring is unproven and one doubt in it is specific.** The hook has never fired on
its own, since settings are read when a session starts, and the command is `./xeno cost
turn` — a relative path, resolved against whatever directory the harness runs hooks
from. If that is not the repository root, the hook silently does nothing, and nothing
will say so: a phase with no cost record is complete.

**Which is the shape of the real risk here.** There is no gate, so a hook that stopped
working and a phase worked outside its window produce the same artifact, namely none.
The figure is optional by section 11 and unmonitored by construction.

**The numbers will read as mostly unattributed**, and that is a measurement about how
this project works rather than about the writer. It also means #117 gets its third
figure in a form that says "seven per cent of what was spent, for the phases that were
worked inside", which is honest and much less useful than a cost per intent.

**The ledger grows without bound.** A line a turn, gitignored, and nothing prunes it.
Section 12 ties retention to the presence of `cost.yaml` and that rule is unwritten, so
the file that should trigger pruning exists and the pruning does not.

**`tokens_cached` adds a cache read to a cache write.** Section 11 names one field and
the transcript carries two that cost differently. Conforming loses that, no test can see
the loss, and the phase that noticed it is the only record of it.

**Codex records nothing**, and section 14 names two clients.

**Accepted with the six named.** The state it replaces is a file section 4 lists,
section 11 specifies, #65 has been open about since the beginning, and nothing wrote.
