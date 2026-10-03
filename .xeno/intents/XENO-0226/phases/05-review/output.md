---
intent: github.com/triplem/xeno#176
phase: 05-review
created: "2026-10-03T11:42:50Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+bdf4e26.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f51f457fc2b28337952a43bf3ebcc8f4a173efade9a3f11034e9d9e8694a538a
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
      Three deviations, each naming what it departs from: a test that was right about the old
      behaviour, the walk's single read against a link's two, and the budget's comment against the
      convention for a paragraph changed twice.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      The lock gains a field and `omitempty` makes an older lock a valid newer one, so there is
      nothing to migrate. What changes for a reader is which number a budget is judged against, and
      the specification clause committed before the code is the note.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Three deviations, each naming what it departs from: a test
that was right about the old behaviour, the walk's single read against a link's two, and the
budget's comment against the convention for a paragraph changed twice.

**`interface-change-needs-a-migration-note` — met.** The lock gains a field and `omitempty` makes
an older lock a valid newer one, so there is nothing to migrate. What changes for a reader is which
number a budget is judged against, and the specification clause committed before the code is the
note.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added.

**Beyond the three rules.**

*Was the ordering the first standing rule asks for actually followed?* Yes, and visibly: the
specification clause is its own commit, `bdf4e26`, before any of this intent's code, which is what
#155 did for the tools entry and what the rule says. The clause was written to answer the question
the code would ask, and the intake says what that bought — a design with nothing to infer.

*Did the trail stay safe?* That was the first criterion written and the one a careless
implementation gets wrong quietly. A lock with no sizes says nothing about bytes; seventy-nine
intents are in that state and the test that pins it also checks the file budget still speaks, so
the silence is about the byte half alone.

*Is anything judged that cannot be?* No. The sizes bound the declaration, not the reading, which
section 5 has always said about this file and which this intent does not improve.

*What arrived while this ran?* An observation from the maintainer that `plugin_version` has never
said anything — a constant where the vendored plugin now has a manifest of its own. It is #177,
opened rather than absorbed, and it is the fourth of section 5's lock fields: `plugin: { version,
sha256 }` is what would make the frontmatter's claim checkable. That field needs no document change,
unlike the size.

<!-- xeno:section:release-notes -->
## Release notes

**A budget is judged against what the phase was given.** `context.lock.yaml` records the size of
each file beside its hash, and G-Schema's byte budget sums those recorded sizes rather than
measuring the tree. A file that grows after a phase is sealed no longer moves that phase's
standing, and a file that disappears no longer quietly brings it inside its budget.

Measured: a profile over `docs/**` recorded 320,480 bytes inside a 400,000 budget; a document then
grew by 300,000 bytes and the phase was judged again with no finding. Before this, the same growth
reported a sealed phase a quarter over budget.

**A lock written before this records no sizes, and says nothing about bytes.** Not a finding against
zero and not a pass: the seventy-nine intents behind this change cannot be judged against a byte
budget, and the check is silent rather than reassuring. The file-count budget is unaffected — it was
always judged against the lock.

**Nothing in the gate path measures a file any more.** A verdict is computed from artifacts and the
content they name, and the byte budget was the one place a gate asked the tree for a number.

**The specification changed first.** Section 5's `context.lock.yaml` block carries `bytes` in its
`files` entries, with the sentence that says why, in its own commit before this code — which is
what the first standing rule requires.

Closes #176. Refs #171, refs #172.

<!-- xeno:section:residual-risk -->
## Residual risk

**Nothing checks a recorded size against the file it names.** G-Freshness compares hashes, which is
stronger, but a hand-written lock can claim any number and the budget will believe it. That is the
same trust this process already places in a hand-written `by-hand` hash, and it is worth saying
once rather than discovering.

**A link's document is hashed and sized by two reads**, because it is resolved by path outside the
walk. A document replaced between them would be recorded as one version's hash and another's size.
Narrow, and the cost of links being paths.

**The budget still bounds the declaration and not the reading.** A phase can be inside its byte
budget and have read twice as much, because section 5 says the lock records what was declared. This
intent makes the bound judgeable; it does not make it a measurement, and nothing in WP8 can.

**Seventy-nine intents can never be judged against a byte budget.** Their sizes were not recorded
and nothing will backfill them, so the context economy of this project's own history is
unanswerable from the trail. One more reason the reading half of #117 will have to be measured live.

**Both conditions on the profile experiment are now met and the experiment has not been run.** #172
closed the first, this closes the second. Until it runs, every mechanism in WP8 — the base, the
order, the links, the budget, the changed set — is exercised by tests and by four demonstrations on
copies, and by nothing in the trail.

**And one thing this intent learned about the fields beside it.** #177 was opened while this ran:
`plugin_version` is a constant and has never distinguished anything, which also means section 5's
sentence about the frontmatter naming what was used and the lock proving it with a hash has never
been testable. Two of the lock's enumerated fields are still missing, and after this morning only
one of them needs a document change.
