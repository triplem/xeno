---
intent: github.com/triplem/xeno#184
phase: 02-design
created: "2026-10-03T12:55:25Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9867bfc13262bbaa46161c01b04a6ff2276ff8c11b9d8cf699d7a4a68d013ee7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**`New` reads the variable.** `Runner.ToolVersion` is initialised from `os.Getenv` where the runner
is constructed, because the value describes the session and a session has one runner per process.
Reading it per command would let two phases of one session disagree about what wrote them.

**`HarnessVersionEnv` is an exported constant.** The name appears in three places — the reader, the
flag's help text and the tests — and a string literal in three places is how one of them comes to
be misspelled. Exported because `cmd/xeno` names it in the help and the surface test asserts
against it.

**`parse` assigns the field only when the flag is non-empty.** It assigned unconditionally, which
would have wiped what `New` read. The precedence is therefore flag, then variable, then absent, and
it lives in one `if`.

**The two writers do not change.** `SectionSet` and `writeDigest` already read
`Runner.ToolVersion`, so the new channel reaches both without either knowing a channel exists. That
is the reason the field was put on the runner rather than passed to the writers in #181, and it is
what makes this change four lines.

**The runner reads one variable and no client specific one.** Section 7 says the runner does not
know which harness it runs under; a reader for `CLAUDE_PLUGIN_ROOT` here would contradict the
sentence that justifies the list existing. The comment on the constant says so and names #183.

**`newFixture` clears the variable, and a `reopen` helper builds the runner again.** Clearing makes
every existing test hermetic for free. `reopen` exists because the one test about the variable has
to set it and then construct the runner, and duplicating the fixture's construction for one test is
how two constructions drift.

**The flag's help text names the variable it overrides.** Somebody reading `--tool-version` should
learn that a session-wide form exists without reading section 7.

<!-- xeno:section:alternatives -->
## Alternatives

**Reading the variable per command rather than in `New`.** Refused: the value describes a session,
and a per-command read invites one phase of a session to record a different version from the next
for a change nobody made. It would also put the read in `parse`, which is the CLI, where a runner
fact does not belong.

**The variable winning over the flag.** Refused. Both are statements about what wrote a phase, and
the narrower one is the better evidence: a variable is how the session was started, possibly weeks
ago by a profile, and a flag is somebody saying it about this run. A80 also already said the flag
would become the override, so inverting it would have made that row wrong twice.

**Dropping `--tool-version` now that the variable exists.** Refused for the case the variable does
not serve: a person running one phase by hand, a script doing one step, a pipeline that sets no
environment. Two channels with a stated order cost one `if`.

**Reading `XENO_HARNESS` in the same change, for the `tool` field.** Refused as a different
argument. `tool` comes from `project.yaml`'s agent block by A35's second amendment, which has a
reason of its own, and changing where a field comes from is not the same act as giving a writerless
field a channel. It belongs to #183.

**Normalising a client variable as a fallback**, `CLAUDE_CODE_VERSION` or its kind. Refused by the
sentence that makes section 7's list coherent: the runner does not know which harness it runs under.
A fallback chain ending in a client's variable is harness detection with extra steps.

**Leaving the tests to inherit the environment.** Refused as the quiet failure. The absence tests
would pass a value on any machine whose profile exports the variable, report green, and be wrong
about the one rule A35 exists to protect.

<!-- xeno:section:impact -->
## Impact

**`docs/process-definition.md` and `docs/implementation-plan.md`.** Already committed, before this
code and on its own: section 7's block gains the fourth entry and a paragraph saying why a version
belongs in a list of things recorded and never branched on, with the `XENO_HARNESS` paragraph
replaced rather than edited into; the plan's WP7 sentence names the same four.

**`internal/runner/runner.go`.** `HarnessVersionEnv`, a new exported constant with the comment that
this is the only `XENO_*` variable the binary reads and that the other three are specified and read
by nothing. `New` initialises `ToolVersion` from it. The `ToolVersion` field's comment is replaced,
because it said the value is an input rather than something read and that is no longer the whole
truth.

**`cmd/xeno/main.go`.** The unconditional assignment becomes conditional, the flag's help names the
variable, and the usage's `common:` block gains a line for it.

**`internal/runner/runner_test.go`.** `newFixture` clears the variable and delegates construction to
a new `reopen`; `TestTheHarnessVersionIsReadFromTheEnvironment` asserts the same thing the first
`tool_version` test asserts, reached through the variable instead of the field.

**`cmd/xeno/main_test.go`.** `TestTheFlagBeatsTheHarnessVersionVariable`, which writes one section
with the variable alone and a second with the flag over it, and reads the artifact after each.

**`.xeno/plugin/skills/`.** The same paragraph in all six phase skills, now naming both channels and
showing the export.

**`ASSUMPTIONS.md`.** A80 marked superseded in part, saying which half of it failed — it treated the
document as fixed rather than as something to ask about — and A81 for the two channels and their
order.

**`README.md`.** Both channels in the paragraph, and the two new test names in the WP7 row.

**No gate and no rule changes**, so nothing behind this recomputes differently. **No change to the
two writers**, which is what keeps this small.
