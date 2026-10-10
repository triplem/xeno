# Labels and the states they stand for

The labels on this repository's tracker have come to stand for states of an issue and of
the intent it becomes. This page records which labels exist, who sets each one, who
clears it, what reads it and what it blocks, so that the meanings are written down
somewhere rather than carried in the habits of whoever is working.

## What a label is here

**A label does not hold an intent's state.** The state is in `.xeno/intents/`: which
phases are complete, what each one's verdict was, whether a phase holds staged content,
whether the intent is in progress, complete or abandoned. `xeno intent status` prints
it, and it is derived from the repository and from nothing else. No label is computed
from that state, nothing compares a label against it, and setting or clearing one
changes no artifact and no verdict. So a label here is a person's note to other people
about an issue, with one exception, and the exception is a precondition rather than a
state.

**Nothing in Xeno ever writes a label.** The adapter contract in `internal/host/host.go`
is three operations — read what a branch enforces, read an issue, write a comment — and
a label write has nowhere to live in it. Every label on an issue is somebody's act on
the host, through the web interface or the command line.

**Xeno reads exactly one label.** Both adapters carry every label on the issue into
`Issue.Labels`, and one function looks at that list: `Issue.Approval()` in
`internal/model/identity.go`, which compares each entry against the constant
`ApprovedLabel`. Every other label on this repository, including the twenty-one work
package labels, reaches that code and is passed over.

The set below was read from the host on 2026-10-10 with

    gh api repos/triplem/xeno/labels --paginate --jq '.[].name'

which lists 33 labels in three groups. Nothing keeps this page and that command in
agreement; the command is given so that a reader can repeat it rather than trust the
page's age. That is the same arrangement the implementation plan chose for the work
package labels, whose creation it calls "a one time act with no script behind it".

## The one label Xeno reads

| Label | Who sets it | Who clears it | What reads it | What it blocks |
|---|---|---|---|---|
| `xeno-approved` | a person with the right the host grants to set labels | nobody, today | `xeno intent start`, and the intake when it writes its problem section | `xeno intent start` refuses an issue without it |

Approval is two things on the issue and neither is enough alone: this label, and a
comment whose first line is `/xeno approved`, the rest of which is the reason. Section
12 of [the process definition](process-definition.md) fixes both and says why — the
label is what a list of issues shows and setting it takes a right, while the comment
carries a reason and can be written by anybody. Both names carry the tool's own, because
a brownfield tracker may already have an `approved` label meaning something else.

What reads it, exactly. `Issue.Approval()` compares the label with case ignored and
surrounding whitespace trimmed, and takes the reason from the **last** comment whose
first line matches, because an approval is withdrawn and given again by writing another
one. The label is read twice in the life of an intent: once by `xeno intent start` as
the precondition, and once more at P0, where the intake writes who approved, when and
with what reason into its problem section. It is never read again. A103 in
[assumptions.md](assumptions.md) says why the second read happens rather than the answer
being carried from the first: the intake records what is true when it is written, which
is the record.

The label is created once per repository, by hand or by
`.xeno/plugin/bin/xeno-labels.sh`, which `xeno init` names among the host settings it
does not make. It is the one label Xeno asks a project's tracker for; everything below
is this repository's own.

Nineteen issues carried it on 2026-10-10, against more than a hundred and forty
intents in the trail. The reason is A107: the two names moved on #332 with no
transition, so issues approved before that carry a label that no longer exists on
this repository. It is worth knowing because it shows what the label
is and is not. The durable record of an approval is the sentence the intake sealed, not
the label; the label is what makes an issue findable as one where work may start.

**Open point.** What should become of `xeno-approved` when the work it authorised stops
— a decision against it, or an issue that is no longer of interest — is open on #359,
with the options put to the maintainer. It would be a section 12 change before it was
anything, which is why this page records the reading above rather than a better one.

## The labels this repository has given itself

Neither of these is read by any code. They are conventions of this tracker, and an
adopter of Xeno needs neither.

| Label | Who sets it | Who clears it | What reads it | What it blocks |
|---|---|---|---|---|
| `xeno-needs-decision` | the agent, when it has put a question to the maintainer and is waiting | whoever answers, or the agent once the answer is recorded in a phase's decisions | nobody | nothing mechanically; it is how a reader of the issue list sees that the agent is blocked rather than working |
| `wp0` … `wp20` | a person, when the issue is created | nobody; a package label outlives the issue's closing | nobody | nothing |

### A question waiting on a person

`xeno-needs-decision` was created on 2026-10-10 on the maintainer's instruction, in the
comment that approved #359, with the description "A question is open and waiting on the
maintainer; the agent is blocked on it, not working it". It says what the issue list
cannot otherwise show: that an issue is not progressing because somebody owes an answer,
rather than because nobody has got to it.

Section 8 of the process definition is what the label sits beside rather than replaces.
A question raised inside a phase goes in that phase's `open-questions` section, is
sealed with it, and `G-Questions` demands from P5 that it was resolved as a decision, as
a confirmed assumption or as a withdrawal. A question about a clause only a person can
write is a different thing: it belongs to the issue and not to the intent's own work, so
it is put on the issue as a comment with its options and this label is set beside it.
Nothing in the artifacts distinguishes the two kinds, which is recorded as a finding
against section 8 rather than settled here.

Seven issues carried it on 2026-10-10: #120, #224, #234, #329, #334, #338 and #359.

**Open point.** The approving comment on #359 spells the label `xeno-need-decision`; the
label that exists and that seven issues carry is `xeno-needs-decision`. The spelling is
the agent's and the discrepancy is named here rather than smoothed over, because a label
is a string and two spellings of one state is the fault this page exists to prevent.
Whether it is renamed to the singular is a question for the maintainer, asked after the
one already open on #359.

### The work packages

`wp0` to `wp20`, one per work package of the
[implementation plan](implementation-plan.md), created once at setup from the list in
its section 2. The plan says what they are for: "A work package is a label, an intent
is an issue. […] Labels group, issues are the unit of work, and a package's progress is
a filtered issue list, which every host provides and epics do not."

They are project labels and not group labels, and the plan is explicit that the package
numbers "are this project's construction plan and mean nothing in any repository that
merely uses Xeno". The third standing rule in `CLAUDE.md` is why an issue carries one:
every change belongs to a work package and to an intent. Where something needed belongs
to no package, that is a finding about the plan and gets written down rather than
absorbed — #359 itself is such a case, and so was #216 before it.

Nothing reads them. Searching the whole of `internal` and `cmd` for a work package label
finds three test fixtures, where `wp12` sits in a label list that an adapter carries
into `Issue.Labels` and `Issue.Approval()` passes over; the same search finds
`xeno-approved` in a dozen places including the one comparison. A `wp` label is read by
people and by the host's own filters.

## The host's own labels

`bug`, `documentation`, `duplicate`, `enhancement`, `good first issue`, `help wanted`,
`invalid`, `question`, `wontfix` and `accessibility`. The first nine are what GitHub
creates with a repository; `accessibility` was added here. They carry no process meaning
at all, nothing reads them, and three of them are in occasional use — `question` on
seven issues, `bug` on six, `documentation` on three, as read on 2026-10-10.

They are listed for completeness, because a page that enumerates a tracker's labels and
silently omits ten of them is a page a reader cannot check against the host.

## Applying a label, and reading it back

    gh api -X POST "repos/triplem/xeno/issues/<n>/labels" -f "labels[]=<name>"
    gh api "repos/triplem/xeno/issues/<n>/labels" --jq '[.[].name]|join(",")'

The second line is the point. An exit code is a report about a request, not evidence
about an issue: it says the host accepted something, and what the issue carries
afterwards is a separate question with a separate answer. Whoever applies a label where
the outcome matters reads it back.

`gh issue edit <n> --add-label <name>` does the same job and was probed on #359 on
2026-10-10 under the conditions that make a failure visible — the label present on the
repository, absent from the issue beforehand — and it applied the label and exited 0. An
earlier session saw that form exit 0 without applying one. That observation is in the
trail of this intent rather than on this page, because it did not reproduce, and a page
that recorded it as a defect would be asserting something today's probe does not
support. The read-back is prescribed on its own merits and holds whichever form is used.

Creating a label is `gh label create`, and the one label Xeno requires has a script:
`.xeno/plugin/bin/xeno-labels.sh github triplem/xeno`, which is idempotent and reports a
label that already exists rather than failing.

## What this page does not say

Four things are open, each on an issue, so that a reader does not take this page's
silence for an answer.

- What becomes of `xeno-approved` when the work it authorised stops, on #359.
- Whether a fresh approval is owed after a decision reverses the work, which follows
  from the previous question and is asked after it, on #359.
- Whether `xeno-needs-decision` is renamed to the maintainer's singular spelling.
- Whether a label should ever *start* work rather than permit it. That is #338, and it
  is the other half of the title #359 was raised under. Section 12 holds that "issue
  commands are deliberately absent" because something has to receive them, and #338
  holds the use case for admitting a trigger, the question of what a trigger attaches to
  and the maintainer's answer to it. No trigger label exists on this repository, and a
  section 12 change comes before one does, so this page lists none.

Nothing mechanically checks any of what is above. No gate reads a label, no rule
mentions one, and nothing notices an issue whose labels no longer match its state —
including the one label that is a precondition, which is refused on when absent and
never checked for afterwards. These are conventions people keep, and the one place that
is different is `xeno intent start` refusing to begin.
