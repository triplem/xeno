---
intent: github.com/triplem/xeno#181
phase: 05-review
created: "2026-10-03T12:42:25Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6cbeac4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cc9189bcfd94bf66cd70aef71d2f38264c9d0443272f0c7331439d25e6acb751
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
  - rule: deviations-are-traceable
    result: met
    note: >-
      One deviation, and a narrowing rather than a departure: the usage documents the flag on the
      `common:` block instead of the `section set` line, which would have had to give up its note
      about reading stdin to carry it inside 88 columns. Recorded in P3.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      The surface gains an optional flag and takes nothing away. A repository that never passes it
      behaves exactly as before, the field absent and the gate reporting it missing, which is a
      criterion of this intent and is tested as one. Nothing to migrate.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** One deviation, and it is a narrowing rather than a departure:
the usage documents the flag on the `common:` block instead of the `section set` line, because that
line would have had to give up its note about reading stdin to carry it inside 88 columns. Recorded
in P3 where a reader comparing the usage against the design would look for it.

**`interface-change-needs-a-migration-note` — met.** The surface gains an optional flag and takes
nothing away. A repository that never passes it behaves exactly as before — the field absent, the
gate reporting it missing — which is a criterion of this intent and is tested as one. Nothing to
migrate, and the artifacts already written keep what they have.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added.

**Beyond the three rules.**

*Does this satisfy A35 or contradict it?* Satisfies. The row's reason is that a plausible value in a
field nobody produced is worse than an absent one, and a version the harness reports is produced.
The runner still derives nothing, still asks no harness anything, and still writes nothing where
nothing was reported. The row was amended rather than overruled, and its claim that two fields
remained writerless was already one out of date.

*Was the specification respected where it had something to say?* Yes, and it is what shaped the
design. Section 7 forbids the runner branching on the harness, which rules out the obvious fix.
Section 7 also designs the better channel, which the first standing rule puts out of the agent's
reach — so the choice was put to the maintainer before any code was written. That is the rule
working rather than being worked around.

*Could this change a verdict behind it?* No. No gate, no rule and no existing artifact is touched,
and `gate verify` is 247 at exit 0.

*Is the problem actually solved?* For the agent that passes the flag, yes, and this intent's own six
phases are the demonstration: twelve artifacts, no hand edit, P0 green on its first finish for the
first time here. For an agent that forgets, no — the phase is as red as before. That is the gap P4
records and the reason A80 says the flag is a channel rather than the channel.

*What did it cost to find?* One grep. The problem had been written down twice, as a writerless
field, and `grep os.Getenv` turned it into a missing mechanism, which is a different and larger
finding than either record had.

<!-- xeno:section:release-notes -->
## Release notes

**`--tool-version` reports the harness's own version.** It is the one field of section 5 the runner
cannot know: section 7 forbids it branching on the harness, so the value is an input.

    xeno section set problem --intent KEY --phase 00 --file PATH --tool-version 2.1.276

**Report it once per phase, on the first `section set`.** A later write that does not repeat it
keeps what was recorded, because the field describes the session and not the invocation.

**`phase finish` copies it into the digest from the phase's `output.md`.** No second argument, and a
phase judged twice loses nothing — which is what made this a hand edit paid again on every finish.

**Left out, nothing changes.** The field is absent and G-Schema reports it missing, exactly as
before. Neither writer invents a value.

**The six phase skills now say to report it**, so an agent driving the process through the plugin is
told without having to be told.

**What this does not do.** `XENO_PLUGIN_ROOT`, `XENO_PLUGIN_DATA` and `XENO_HARNESS` are still read
by nothing, so the entry point section 7 specifies remains unbuilt and a fourth variable for the
harness version remains the better answer. The flag would become its override.

<!-- xeno:section:residual-risk -->
## Residual risk

**The flag has to be remembered.** A skill instructs and does not enforce, so a phase written
without it is as red as it was before and the hand edit returns for that phase. This is the main
residual risk and it is the whole argument for the environment variable the entry point would set
once per session. Recorded in P4's gaps and in A80.

**The value is unverifiable.** `--tool-version 9.9.9` is recorded as readily as the truth. That is
already true of `model` and `tool`, which come from `project.yaml`, and it is the shape of A35's
territory: the process records what a session reported about itself. Nothing here makes it worse and
nothing here could make it better.

**Two provenances for one field, undistinguished.** The artifacts written before this carry a typed
value inside a sealed hash and the ones after carry a reported one, with nothing in either saying
which. A reader has the commit and `runner_version` to date them by, and backfilling is not
available: the hashes are sealed and that is the rule this process rests on.

**The digest's copy depends on `output.md` being readable.** A phase whose artifact is malformed
gets an absent field rather than an error, which is the same answer as a phase that reported
nothing. It is deliberate — the alternative is a digest that fails to be written over a frontmatter
problem G-Schema is about to report anyway — and it means a corrupt artifact degrades quietly in
this one field.

**What is not a risk.** Any verdict behind this intent, since no gate, rule or existing artifact is
touched; and any repository that never passes the flag, whose behaviour is asserted to be unchanged.
