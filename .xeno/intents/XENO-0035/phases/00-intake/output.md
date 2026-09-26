---
intent: github.com/triplem/xeno#35
phase: 00-intake
created: "2026-09-26T07:32:23Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+5b66fcd.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: dc9fcb7a525e4b9b662c7f8d7c58ce3fd528bfc7a0160e70b32b63226ddbe982
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

Somebody working an intent has to hold the working sequence in their head. The runner
knows which phase is running, what its verdict says, whether the directory still matches
that verdict and which sections the template wants, and says none of it. The question it
cannot answer today is the one that gets asked most: what now.

<!-- xeno:section:scope -->
## Scope

Every command that changes something ends by naming the next step, derived from the
state it left behind rather than the one it found, and `xeno intent status` says it
too, because that is the command somebody runs when they have lost the thread.

`--no-next` leaves it out. The commands a pipeline or a git hook runs never print it:
`gate verify`, `check commit-message`, `enforcement check`.

<!-- xeno:section:context-rationale -->
## Why this context

**It reads section 6 and adds nothing.** No field, no gate, no rule: the working
sequence already fixes who does what with which command and in which order, and every
input is state the runner computes today.

Two properties are the whole point and both are tested.

**On a red verdict it offers no command.** It names the finding, the repair the finding
itself carries, and that a second person can release it with `gate approve` or
`gate override` — and offers none of them to run. A runner that put a decision in the
reader's hand would be nudging towards the thing the process exists to record. The test
that holds this was mutated: adding the command back fails it.

**Where the next step is nobody's subcommand it says so.** The merge after P5, the
commit and the push, the second person, the pipeline. Offering a command that does not
exist is worse than offering none, because it gets tried.

An override does not block the sequence. Section 6 puts its obligation at "later,
whoever owes it", so the suggestion points on and lists what is owed, because later is
otherwise never.

A state the table does not cover is said to be one. A confident wrong suggestion is
followed, which is worse than silence.

<!-- xeno:section:open-questions -->
## Open questions

- Q-1: whether a harness reads the same thing. WP11 says every MCP operation has a
  command behind it, and a machine readable next step would let an agent follow the
  sequence rather than a person reading a sentence. That decides whether this is a
  convenience or the process's control loop, and it belongs with WP11 rather than here.
- Q-2: the suggestion follows section 6 and therefore offers the next phase, while this
  repository holds every intent at P0 until WP8 closes the G-Freshness gap (A6). That
  hold is a decision in `M0.md` and not a state the runner can see, so the suggestion
  is right about the process and wrong about this repository. Whether a project can
  declare such a hold is a question for WP8.
