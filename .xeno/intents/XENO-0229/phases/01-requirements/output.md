---
intent: github.com/triplem/xeno#184
phase: 01-requirements
created: "2026-10-03T12:54:42Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5b221a1f0bd97eb1a152932f25c7bb73149cde8cc0329df694eab03f3610ce02
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

**The specification change is its own commit, made before any code.** Section 7's block lists
`XENO_HARNESS_VERSION` with the reason beside it; the plan's WP7 sentence names the same list and
has it too. No code in that commit.

**`export XENO_HARNESS_VERSION=2.1.276` is enough for a whole session.** Every artifact of every
phase carries `tool_version` with no flag anywhere — asserted on a throwaway tree and demonstrated
by this intent's own six phases.

**`--tool-version` written over the top of the variable wins.** The variable is how a session was
set up; a flag is somebody saying it about this run.

**Neither one said leaves the field absent, in both files.** Unchanged from #181 and still A35's
rule: the row's reason is that a plausible value in a field nobody produced is worse than an absent
one, and neither channel produces one from nothing.

**The variable is read once, where the runner is constructed.** A value that holds for a session is
read at the start of one, not re-read per command.

**The tests are hermetic.** They clear the variable rather than inheriting whatever the machine
running them exports, because a developer with it in their profile would otherwise watch the absence
tests pass a value and report green.

**The runner reads this variable and no other.** Not the other three in section 7's list, not
`CLAUDE_PLUGIN_ROOT` or any client variable. Normalising those is the entry point's work by the
section's own account.

**The flag is not removed.** A person driving the runner by hand, or a script doing one phase, keeps
a way to say what wrote it.

**Nothing else changes.** No field of section 5, no gate, no rule, no existing artifact, and nothing
for a repository that sets neither channel. `gate verify` stays at exit 0 and the suite stays green.

<!-- xeno:section:non-goals -->
## Non goals

**No entry point.** `XENO_PLUGIN_ROOT`, `XENO_PLUGIN_DATA` and `XENO_HARNESS` stay read by nothing,
and the `--plugin-root` argument and the resolution order stay unbuilt. #183 holds all of it. After
this intent the runner reads exactly one of section 7's four variables, which is that issue's
finding stated more precisely rather than resolved.

**No client specific variable.** Section 7 says the entry point normalises a client's environment
into `XENO_*` and that the runner does not know which harness it runs under. A runner reading
`CLAUDE_PLUGIN_ROOT` would be doing the entry point's job in the one place the section forbids it.

**No removal of `--tool-version`.** It becomes the override A80 said it would become.

**No `XENO_HARNESS`.** Recording the harness name instead of taking `tool` from `project.yaml` is a
separate argument with a separate answer, and A35's second amendment is where that one lives.

**No default and no fallback.** Absent stays absent. This is the third intent in a row to say so and
it is the one line that must not move: A35's reason is the whole argument behind every channel added
here.

**No per-command re-reading of the environment.** The value describes a session. Reading it per
command would make one phase's artifacts disagree with another's for a change nobody made.

**No backfill.** XENO-0228's artifacts carry a flag-reported value, which is correct and sealed.

<!-- xeno:section:constraints -->
## Constraints

**The specification change comes first and is its own commit.** The first standing rule, and the
shape XENO-0226 used: the clause written to answer the question the code will ask, then the code.
The approval for it was given explicitly; without that this intent could not exist.

**Section 7 bounds what may be read.** Four variables, recorded only, and no knowledge of which
harness is running. That sentence is why the runner reads one variable and not a client's.

**A35 is the test for the channel.** Reported is produced; derived would not be. The variable
satisfies the row for the same reason the flag does.

**The value holds for a session, so it is read once.** `New` is where the runner is constructed and
where a session's facts belong. This decides both the reading and the fact that `parse` must stop
assigning the field unconditionally.

**A runner that reads the environment makes the test suite depend on the machine.** Whatever asserts
absence has to clear the variable, or the suite is green on a developer's machine and green for the
wrong reason.

**The flag stays and must win.** Two channels for one field need an order, and the stronger
statement is the one made about this run.

**88 columns, SPDX, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0.**

**One intent, one branch, one issue** — `184-the-harness-version-variable`, #184, labelled wp11,
part of #183 — and one commit before the code, which is the specification change.
