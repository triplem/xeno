---
intent: github.com/triplem/xeno#184
phase: 05-review
created: "2026-10-03T12:58:35Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4e2a7340521c1508a3948349fa8fb20e5da7a28ae37dabc13d31e91b2055d60c
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
      Three, each naming what it departs from: none from the design; one from A80, which is the
      intent's purpose rather than a departure in it, marked superseded in part with the half that
      failed named; and the re-alignment of section 7's block, typography inside a document the
      agent may only touch by approval.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      A variable is added and nothing is taken away. `--tool-version` keeps working and now wins
      over the variable, which A80 predicted in those words, and a repository that sets neither
      behaves exactly as before, which is tested.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added; `os` was already imported.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Three, each naming what it departs from: none from the design;
one from A80, which is the intent's purpose rather than a departure in it, marked superseded in part
with the half that failed named; and the re-alignment of section 7's block, which is typography
inside a document the agent may only touch by approval and is recorded so the diff needs no
explaining.

**`interface-change-needs-a-migration-note` — met.** A variable is added and nothing is taken away.
`--tool-version` keeps working and now wins over the variable, which A80 predicted in those words. A
repository that sets neither behaves exactly as before, which is tested. Nothing to migrate.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added; `os` was already
imported.

**Beyond the three rules.**

*Was the first standing rule followed where it mattered most?* This is the first commit in this
repository to change the process definition at the agent's hand, so the ordering is the thing to
check and it is checkable: `e8f68b1` is two documents and no code, and every line of code came
after it. The approval was explicit and is in the issue, the commit message and A80's amendment.

*Did the approval get used for more than was approved?* The question is worth asking because a
document opened once is a document that can be edited twice. What changed is the four-entry block,
the paragraph that justifies it, and the plan's sentence naming the same list. Nothing in section 5,
nothing in Appendix A, no gate, no rule. The re-alignment of three existing lines is the only thing
in that diff nobody asked for, and it is recorded as a deviation for exactly that reason.

*Is A35 still intact after two channels?* Yes, and it is the row this pair of intents has leaned on
hardest. Reported is produced; neither channel invents a value; absence is still absence, in both
files, and is now tested in a way that cannot pass by accident on a machine whose profile exports
the variable.

*Could this change a verdict behind it?* No. No gate, no rule and no existing artifact is touched;
`gate verify` is 253 at exit 0.

*Is the problem solved?* One export per session instead of six flags per intent, and this intent's
own twelve artifacts carry the field with no flag passed anywhere. Not solved: something still has
to type the export. That is #183 and P0 said so before the code existed, which is the difference
between a smaller claim and an overstated one.

<!-- xeno:section:release-notes -->
## Release notes

**`XENO_HARNESS_VERSION` reports the harness's version for a whole session.** It is the fourth entry
in section 7's normalised environment, recorded only, beside `XENO_HARNESS`.

    export XENO_HARNESS_VERSION=2.1.276

Every `section set` of that session writes `tool_version` into `output.md`, and every `phase finish`
copies it into the digest. No flag anywhere.

**`--tool-version` still works and now overrides it**, for a person running one phase by hand or a
script doing one step.

**Said by neither, the field is absent**, G-Schema reports it missing, and nothing is invented —
unchanged, and the rule the whole arrangement rests on (A35).

**The specification changed first, in its own commit.** Section 7 lists the variable and says why a
version belongs among the things recorded and never branched on; the plan's WP7 sentence names the
same four.

**What this does not do.** `XENO_PLUGIN_ROOT`, `XENO_PLUGIN_DATA` and `XENO_HARNESS` are still read
by nothing, and nothing yet sets `XENO_HARNESS_VERSION` for you. The entry point that would is #183.

<!-- xeno:section:residual-risk -->
## Residual risk

**Somebody still has to type the export.** One per session instead of six flags per intent, which is
an improvement and not a solution. A session that exports nothing and passes nothing is as red as
before #181. #183 is what closes it, and P0 stated this limit before any code was written rather
than discovering it at review.

**An export in a profile outlives the version it names.** A machine upgraded without the profile
being touched reports the old version confidently, for every artifact of every phase. The flag
cannot drift this way because it is typed per run, and the entry point will not because it would
read the version from the client. It is the price of a session-wide channel and it is new with this
change.

**The value is unverifiable, and now from two directions.** `9.9.9` is recorded as readily by the
variable as by the flag — the surface test depends on it — as is a wrong `model` or `tool` from
`project.yaml`. A35's territory: the process records what a session reports about itself.

**Hermeticity rests on one line in one fixture.** A test in `internal/runner` that builds a runner
without `newFixture` inherits the environment again, and nothing prevents one being written. A lint
for it would be a rule and rules go through section 10.

**Three provenances for one field.** Typed before #181, flag-reported in XENO-0228, variable-reported
from here, with nothing in any artifact saying which. The commits and `runner_version` are what a
reader dates them by, and backfilling is neither available nor desirable: the hashes are sealed.

**What is not a risk.** Any verdict behind this intent, since no gate, rule or existing artifact is
touched; and any repository that sets neither channel, whose behaviour is asserted unchanged.
