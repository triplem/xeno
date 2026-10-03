---
intent: github.com/triplem/xeno#184
phase: 00-intake
created: "2026-10-03T12:53:50Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6e42bfd05a27291c07441bb3fdc6f5b3eab080b3d1eff39ae47ff6ab432708f7
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

#181 gave `tool_version` a flag, and A80 recorded that flag as a channel rather than the channel.
The better home for a harness fact is a variable in the list section 7 enumerates, and that row
could not reach it: the first standing rule makes the list a person's commit, made before the code
that follows from it.

**The constraint was the rule, not the design.** A80's own words are that the flag "does not block
it and would become its override", so the design was finished and waiting on a commit the agent is
not allowed to make. The maintainer has approved that commit, so what was out of reach is in reach
and the only question left is the one the row already answered.

**What the flag leaves unsolved is remembering it.** A skill instructs and does not enforce. A phase
written without the flag is exactly as red as it was before #181, which XENO-0228's P4 recorded as
its first gap and its P5 as its main residual risk. A value that holds for a whole session being
passed per command is the shape of that risk: six phases times one flag is six chances to forget,
and the artifact that results is wrong in the one field a gate checks.

**A variable is the channel for a value that holds for a session**, which is what a harness version
is. It is set once by whatever starts the session, and every command of that session carries it
without anybody naming it again.

**One thing this does not change.** The entry point section 7 designs is still unbuilt, so there is
nothing to set the variable except the session itself. This removes the remembering for anybody who
exports it once and for nobody who does not, which is a smaller claim than #183 will eventually
make and is worth stating as the smaller one.

<!-- xeno:section:scope -->
## Scope

**In scope.** The specification change, which is its own commit and already made: section 7's block
gains `XENO_HARNESS_VERSION`, recorded only, with the reason beside it, and the plan's WP7 sentence
gains it so the two documents do not disagree. Then the code: the runner reads the variable where it
is constructed, `--tool-version` overrides it, neither said leaves the field absent as before, and
the tests clear the variable so they do not inherit whatever the machine exports.

**Out of scope, and each for its own reason.**

The entry point. `XENO_PLUGIN_ROOT`, `XENO_PLUGIN_DATA` and `XENO_HARNESS` stay read by nothing,
and the `--plugin-root` argument and the resolution order stay unbuilt. That is #183 and it is
larger than this; what this intent takes is the one slice that needs no entry point to be useful.

Any other `XENO_*` variable, and any client specific variable. Normalising `CLAUDE_PLUGIN_ROOT` and
its kind into `XENO_*` is the entry point's work by section 7's own account, so a runner reading one
here would be doing the entry point's job in the one place the section says it must not.

Removing `--tool-version`. A80 said the variable would make the flag its override rather than
replace it, and that is what happens: a person driving the runner by hand, or a script doing one
phase, still has a way to say what wrote it.

Recording `XENO_HARNESS`. The field beside `tool_version` is `tool`, it comes from `project.yaml`'s
agent block by A35's second amendment, and changing where it comes from is a separate argument with
a separate answer.

Backfilling. The artifacts of XENO-0228 carry a flag-reported value, which is correct.

<!-- xeno:section:context-rationale -->
## Why this context

Section 7's environment normalisation is read as it now stands, one commit old: the four variables,
the sentence that the runner only ever sees `XENO_*` and does not know which harness it runs under,
the paragraph that both harness entries are recorded and never branched on, and the paragraph added
with the fourth that says why a version belongs in the list. That last paragraph is the
specification for this intent and it was written for it, which is the shape XENO-0226 found cheapest
and recorded as a learning.

A80 is read in full, because this intent is its correction. The row is right about the design and
wrong about one thing — it treated the document as fixed rather than as something to ask about — and
saying which half failed is more use than replacing it.

A35 is read again for the same reason it was read for #181: its test is that a plausible value in a
field nobody produced is worse than an absent one. A variable the session exports is produced, so
the row is satisfied by this channel as it was by the flag.

XENO-0228's `04-verification/output.md` and `05-review/output.md` are read for the gap and the
residual risk this closes, because the claim being made is theirs and not a new one.

`internal/runner/runner.go` is read for `New`, which is where a value that holds for a session
belongs, and for the `ToolVersion` field and its two writers, which do not change.

`cmd/xeno/main.go` is read for `parse`, which is where the precedence has to live: it already
assigns the field unconditionally, and that is what has to stop.

`internal/runner/runner_test.go` is read for `newFixture`, because a runner that reads the
environment makes every test that asserts absence depend on the machine it runs on.

Nothing outside the repository is needed.
