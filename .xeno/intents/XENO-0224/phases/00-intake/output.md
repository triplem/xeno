---
intent: github.com/triplem/xeno#1
phase: 00-intake
created: "2026-10-03T10:18:39Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5b930cab3f2f897ec064bf6de8f8864fbedb4bfb005ef9b965848b56d7ad2321
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

M0 is reached and the record does not say so. The plan uses the name for two things and `M0.md`
says which: the milestone table's M0 is the gate job, "from which point Xeno checks its own
repository and the hand held record in section 4 stops", and section 6's is the walking skeleton,
one intent, one phase, one agent end to end.

**The gate job has been in place for weeks.** `verify` is a required status check on the default
branch, administrators included, so no merge reaches `main` without a recomputed verdict. The job
runs `xeno gate verify`, which is green over the whole trail. Twenty-four intents have run all six
phases with artifacts, verdicts and decisions, and every change to the runner since has gone
through them.

**The walking skeleton's blocker went with #169.** Section 6 puts WP11 at step 5 and M0 at step 6.
The agent layer now exists for one harness: a plugin that installs, seven skills named as section
13 names them, hook wiring, and a measured cost. What it does not have is the MCP server, and
WP11's own done-when requires the command path to work without one — which is what twenty-four
intents did and what the skills now write down.

**What has not happened is the bookkeeping.** Issue #1 is open. `ASSUMPTIONS.md` is still taking
rows: A72 to A77 this week, where the plan says the hand held record stops at M0 and the loop runs
through the runner instead. This repository does not carry the `Xeno-Intent:` trailer the plan says
it adopts once M0 exists. And two paragraphs of the plan describe a state two months old: "at M0
six of the fourteen do not exist" where three do not, and "M0 proves one phase" where twenty-four
have proved six.

**So the problem is a record that trails its subject, which is the thing this project exists to
prevent.** The same shape as #156, where the two hand-kept files described a runner that had moved
on — and the same cost: a reader orienting from the plan would conclude the project is at the
beginning of its sequence rather than past M0 and one clause from M1.

<!-- xeno:section:scope -->
## Scope

**In scope.** Closing M0: the evidence written down where a reader will find it, the hand held
record closed at the point the plan says it closes, and issue #1 closed with what was checked.

**Out of scope, and each for its own reason.**

The two dated paragraphs in the plan. The first standing rule puts `docs/` beyond the agent, and
the gate count and the phase count are prose a person changes. They are named in the closure so
that the correction is owed rather than forgotten.

Adopting the `Xeno-Intent:` trailer. The plan says this repository carries it "once M0 exists", so
it follows the closure rather than accompanying it: the rule is copied into `given/project/`, the
squash message has to carry it as well as the commits, and both are a change to how every later
commit is written. Its own intent, next.

The MCP server, the second harness and the lenses. WP11's remaining half, which M0 does not need
and which the plan sequences after it.

Any claim about M1. Its ingredients are built — the rule engine, the assumption register, the
context profile — and one clause of WP8's done-when remains unverifiable by construction. That is a
reading for a person and this intent does not take it.

**One boundary worth naming.** Closing the register is not deleting it. The rows stay, their state
columns stay, and the file stays the place a reader learns what was decided while the core was
built by hand. What stops is adding to it: from here a decision belongs to a phase artifact and a
learning routes through section 10's merge request against the rule set.

<!-- xeno:section:context-rationale -->
## Why this context

The plan's milestone table is read for the M0 row, which is the definition this intent closes
against: the gate job, the point from which Xeno checks its own repository, and the hand held
record stopping. Section 4's working method is read for the sentence that says what stops and what
replaces it: "From M0 the same loop runs through the runner, and the hand held record stops."
Section 6 is read for the walking skeleton and for the sequence that puts WP11 before M0.

`M0.md` is read entire, because it is the only document that says the name means two things and
which of them its own guide builds. Its reasoning is what makes a single closure defensible rather
than ambiguous.

The host is read rather than assumed: the branch protection on the default branch, for whether the
pipeline is required and whether administrators are included. A claim that Xeno checks its own
repository rests on that setting, and it is the one fact in the closure that lives outside the
tree.

`.github/workflows/xeno.yml` is read for the job that runs `gate verify`, and the trail is read for
what that verdict covers: how many intents, how many phases, how many verdicts.

`ASSUMPTIONS.md` is read for its own header, which states that it is the record kept by hand before
M0 and cites the plan's section 4 for why. That header is where the closure belongs, because it is
what a reader meets before the rows.

The plan's paragraph on squash merges and the trailer is read for what M0 triggers and is
deliberately not acted on here.

Nothing outside the repository and the host is needed. The one thing that would have changed the
answer — a plan amended to say M0 needs an MCP server — does not exist, and the plan's own
done-when for WP11 says the opposite.
