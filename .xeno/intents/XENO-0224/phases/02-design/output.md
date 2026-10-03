---
intent: github.com/triplem/xeno#1
phase: 02-design
created: "2026-10-03T10:20:05Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0400f3385b8f23385733b59be3d6ced7f78fe1dec23e40a463dda5c15fb599fe
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

**M0 is closed on both readings, and the closure says which evidence answers which.** The
milestone table's reading is answered by the host's setting and the job: a required check with
administrators included, running a recomputation of every verdict. Section 6's is answered by the
agent layer existing for one harness. Keeping them apart matters because somebody reading the
closure later will ask which of the two it meant, and `M0.md` already had to explain that the name
means two things.

**The register is closed in its header, with the rows untouched.** A reader meets the header before
the table, so that is where "this stops here" belongs. The rows keep their state columns because
they are the record of what was decided and when; a row deleted for having been superseded would be
the rewriting this process refuses everywhere else.

**A77 is the last row, and the closure says what replaces it.** A decision goes into the phase
artifact that took it, and a learning goes into `learning.yaml` and from there through section 10's
merge request against the rule set. That is not a new mechanism: it is the loop the plan says M0
switches to, and both halves have existed since #165.

**The two things M0 triggers are named and not done.** The trailer is its own intent because it
changes every later commit and every squash message, and a convention adopted in the same breath as
a closure is a convention nobody wrote down. The plan's dated paragraphs are a person's, and naming
them is the whole of what the agent may do about them.

**The artifacts of this intent are short on purpose.** The change is one file's header and an issue
comment. Twelve intents of figures on #117 say the cost is fixed per intent, and part of that fixed
cost is the habit of writing to the length of the last intent. This is the first one to test the
other reading, that the sections are required and their length is not.

<!-- xeno:section:alternatives -->
## Alternatives

**Close M0 without closing the register.** The two are separable: the milestone is a fact about the
repository and the register is a working habit. Rejected because the plan ties them in one sentence
— the point from which Xeno checks its own repository is the point the hand held record stops — and
a closure that kept writing rows would be a closure nobody acted on. A78 would have appeared within
a week.

**Close the register and leave #1 open until the MCP server exists.** It would hold the strict
reading of section 6, where WP11 is a whole package before M0. Rejected because WP11's own done-when
requires a phase to be carryable with commands alone and twenty-four intents did exactly that: the
server is an optimisation of a working path, and a milestone held open for an optimisation stops
being a planning instrument.

**Write the closure into `M0.md` instead of the register's header.** That file is the guide to the
gate job and the natural home for a sentence about M0. Rejected because `M0.md` is a guide somebody
reads once and the register is a file somebody is about to add a row to. The closure has to be where
the next row would have gone.

**Delete the register once the loop replaces it.** The rows are history and the trail is the record
now. Rejected outright: A6, A62, A74 and A77 explain why things in this tree are the way they are,
and a reader of the code has nowhere else to learn it. Closed means closed to additions.

**Adopt the trailer in this intent.** It is a file copy, and the plan ties it to M0 existing.
Rejected because adopting it mid-intent would mean this intent's own commits do not carry what the
rule requires of them, so the first thing the new rule would do is report a finding about the commit
that adopted it.

<!-- xeno:section:impact -->
## Impact

**The plan's sequence becomes readable again.** A reader orienting from the documents has been
seeing a project at step 4 of 13 with a hand-kept record; after this, M0 is closed, WP4 is complete,
and the gap to M1 is one clause of WP8 and a reading.

**`ASSUMPTIONS.md` stops growing, and the artifacts take the weight.** Seventy-seven rows have
carried the reasoning for a decision apiece. From here that reasoning lives in a design phase, which
is the loop the plan wants and which is also harder to read: a row in a table is findable and a
paragraph in one intent's P2 is not. That is the cost of the switch and it is the plan's choice
rather than this intent's.

**Two things become owed.** The trailer, which is the next intent, and the plan's two dated
paragraphs, which are a person's. Both are named in the closure so that neither is discovered again
by somebody reading the plan in a month.

**Nothing in the tree behaves differently.** No gate, no rule, no schema. This is a record catching
up with its subject, and the only mechanical consequence is that `ASSUMPTIONS.md` has a different
header.

**One claim gets easier to make and should still be made carefully.** With M0 closed, "Xeno is
developed through Xeno" is true without qualification for everything since XENO-0107. For the
fifty-one intents before that it is true of the intake alone, which `M0.md` and A20 explain and which
the README's trail section already says.
