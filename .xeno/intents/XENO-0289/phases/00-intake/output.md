---
intent: github.com/triplem/xeno#359
phase: 00-intake
created: "2026-10-10T15:32:50Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5b4221e73ab7771aaa9c890736b44a0e52630307b21a3215aec8e9316bb1a788
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: The label xeno-needs-decision exists on the maintainers instruction, and the page records that spelling as the one in use
      rationale: 'The approving comment on #359 asks for another useful label and spells it xeno-need-decision; the label created and applied on 2026-10-10 is xeno-needs-decision, with the description "A question is open and waiting on the maintainer; the agent is blocked on it, not working it", colour fbca04, and it is on #329, #334, #338, #352 and #359. The decision recorded here is the maintainers, that the label should exist; the spelling is the agents and is named rather than left for a reader to notice, because a label is a string compared by whoever compares it and two spellings of one state is the fault this page exists to prevent. The applied spelling is what the page records, since the page records what is in use and renaming a label that five issues carry is an act with a cost of its own: the host rewrites the association on every issue and any saved filter or query naming the old string stops matching silently. Whether it is renamed to the singular is a question for the maintainer and it is not asked now, because one question is already on the issue and a batch of two moves the whole of the thinking onto them, which is what CLAUDE.md records XENO-0243 for. It is named in the page as the open point instead, so it is asked from the record rather than from memory.'
      decided_by: triplem
---

# Intake

<!-- xeno:section:problem -->
## Problem

Approved by @triplem on 2026-10-10T14:36:07Z: And please add xeno-need-decision as another useful label

> **github.com/triplem/xeno#359** — Labels and trigger labels
>
> I guess we have reached a state where the issue labels are somewhat represent states of the issue and/ or intent.
> The labels are not yet defined, I suggest to have a page where all the states/ labels are defined.
> Furthermore it is required to think about who acts in a certain label and when a label is removed. What happens with xeno-approved and /xeno approved, if a decision is on and the user disapproves or the issue is not of interest anymore?
> Should there be a new approval after the decision?

The issue as it stood at 2026-10-10T15:32:50Z, read by `xeno phase start` and quoted rather than summarised. What this phase concludes about it belongs below.

<!-- xeno:section:scope -->
## Scope

#359 asks for three things and they are not the same size. A page where the labels and
the states they stand for are defined, including who acts on each one and when it is
removed. The lifecycle of approval: what becomes of `xeno-approved` and of the
`/xeno approved` comment when a decision goes against the work or the issue stops being
of interest, and whether a fresh approval is owed after one. And, from the title, trigger
labels.

**The page is this intent.** `docs/labels.md`, a record of the labels this repository's
tracker carries today: `xeno-approved`, `xeno-needs-decision`, `wp0` to `wp20`, and the
host's own set. For each one, who sets it, who clears it, what reads it, and what it
blocks. It is a record of what is in use and defers to section 12 for the one label the
runner reads, which is how it stays a document that can be written without the
specification moving first.

**The lifecycle is not this intent, and one question of it is on the issue.** What
`xeno-approved` means after a decision reverses the work is a clause about what
`xeno intent start` reads, and the second standing rule puts that in section 12 before it
is anything else: an addition is a specification change by the maintainer, in a commit of
its own. So the first of #359's two questions is put on the issue with its options, their
consequences and their costs, and `xeno-needs-decision` is applied to #359 for the reason
the label exists. The second question — whether a fresh approval is owed after a decision
— follows from the first and is asked after it is answered, never beside it, because an
answer to the second assumes one to the first that nobody has given.

The question is put on the tracker rather than recorded as an `open-questions` entry of
this phase, and the reason is worth stating because the mechanism exists and is not being
used. An `open-questions` entry is sealed with its artifact and section 8 gives it three
exits, each of which needs a person; `G-Questions` demands one from P5. The question here
is not about this intent's own work, which the page answers on its own, but about a clause
only the maintainer can write. Recording it as a question of this intent would hold a
documentation page at a red P5 until a specification change landed, which is the wrong
dependency in the wrong direction. The label is the project's own device for a question
waiting on a person, and this is what it is for.

**Trigger labels are #338 and stay there.** The title's second half is a mechanism this
project has already decided to keep out of v1 — section 12's "Issue commands are
deliberately absent" — and #338 holds the use case for admitting one, the question of
what a trigger attaches to, and the maintainer's answer to it: one label, `xeno-start`,
and no analysis trigger. That answer is a section 12 change the maintainer has yet to
make, and no such label exists on the tracker. A page listing it would be documenting a
decision as though it were a state of the repository. So `docs/labels.md` names the
labels that exist, says that triggering is #338's, and stops. The half of the title this
intent is not doing is named in the page itself, so that a reader who came from #359 is
not left looking for it.

**What the page does not become.** Not a specification: it states what the runner reads
and points at section 12 for why, rather than restating the clause in words that could
drift from it. Not a how-to for adopters either — `wp0` to `wp20` are this project's own
construction plan and mean nothing in a repository that merely uses Xeno, which the page
says where it lists them.

**The work package.** #359 carries `xeno-approved` and `xeno-needs-decision` and no `wp`
label. WP16 is documentation and is the nearest fit, and nothing in the plan's twenty
packages owns the project's own tracker conventions. The third standing rule asks for the
label and for a finding where something belongs to no package; the finding is this
paragraph and the learning record, and the label is the maintainer's act, not the agent's.

<!-- xeno:section:context-rationale -->
## Why this context

Ten files and 64,727 bytes resolve today; the eleventh, `docs/labels.md`, is what this
intent writes, and the budget of 14 files and 130,000 bytes is set from what the scope
will hold once it exists rather than from what it holds now. The margin is deliberate and
costs nothing: the page is prose whose length is not known until it is written, and the
budget sits inside this phase's `artifacts_hash` and cannot be raised afterwards.

**What changes.** `docs/labels.md` is new. `docs/README.md` indexes the documents and
`zensical.toml` carries the published site's navigation, and a page in neither is a page
nothing links to. `.github/workflows/docs.yml` is why that matters mechanically rather
than only as a habit: it builds with `--strict`, so a link to a page that does not exist
fails a pull request, which makes the index entry checkable and the nav entry the only
way the page is reachable from the site.

**What the page is read off rather than argued from.** `internal/model/identity.go` holds
`ApprovedLabel`, `ApprovedWord` and `Issue.Approval()`, which is the whole of what the
runner reads about any label; it is the file the page's central claim is taken from.
`internal/host/host.go` fixes the adapter contract at three operations — read branch
rules, read an issue, write a comment — which is how the page can say that nothing in
Xeno ever writes a label, and `internal/host/github/tracker.go` is where every label on
the issue is carried into `Issue.Labels` and where the one POST in the adapter goes to the
comments endpoint. `.xeno/plugin/bin/xeno-labels.sh` is the act that creates the approval
label written down so it can be run, and its own comment claims it creates "the one label
Xeno asks a project's tracker for", which this intent has to either keep true or change.
`docs/commands.md` carries the user-facing sentence about the approval pair, and the page
must not say something different from it. `CONTRIBUTING.md` is the other candidate home
for the page and is read in order to argue against it rather than from it.
`CLAUDE.md` holds the third standing rule, which is what the work-package labels are for.

What is deliberately out, and the paragraphs the page needs from each.

**`docs/process-definition.md`, 138,216 bytes.** Normative, not editable by the agent, and
needed for one paragraph of section 12, which is quoted here so that the page can defer to
it without the file being in scope: "An issue becomes an intent only where a person
approved it, and approval is two things on the issue itself, the label `xeno-approved` and
a comment whose first line is `/xeno approved`, the rest of which is the reason. The label
is what a list of issues shows, and setting it takes a right the host grants; the comment
is what the intake quotes, with who wrote it and when. Neither alone is enough: a label
carries no reason, and a comment can be written by anybody. […] The label is a person's
act on the host, which `xeno init` names among the settings it does not make and the
plugin's `bin/xeno-labels.sh` creates for either host." Also from section 12, the clause
#338 turns on: "Issue commands are deliberately absent. Something has to receive them, and
every way of doing that is a component to build and operate." The page cites the section;
it does not restate it, so nothing in it can drift from a document it cannot edit.

**`docs/implementation-plan.md`, 132,590 bytes.** Normative, and needed for the two
sentences that say what a `wp` label is: "**One issue label per work package**, `wp0` to
`wp20`, created once at setup from the list in section 2. Project labels, not group
labels: the package numbers are this project's construction plan and mean nothing in any
repository that merely uses Xeno. It is a one time act with no script behind it", and
"**A work package is a label, an intent is an issue.** […] Labels group, issues are the
unit of work, and a package's progress is a filtered issue list, which every host provides
and epics do not." Both are quoted rather than read from scope for the same reason as
above.

**`docs/assumptions.md`, 116,367 bytes.** Two rows bear on the page and neither needs the
file. A103 is why the intake writes the approval sentence from a second read at P0 rather
than carrying it from the start, and says what that buys: "where the label was removed or
the milestone moved between them the intake says what is true when it is written, which is
the record". A107 is the approval names moving to the current pair with no transition, and
says of the label that it "is created by the shipped script or by hand and never read
back". No row is added by this intent: the page records what is in use, which is not a fact
that outlives an intent in the sense that page asks for.

**`internal/host/gitlab/tracker.go`.** It carries labels into `Issue.Labels` exactly as the
GitHub adapter does, from `body.Labels` at line 64, and the page makes one claim about both
— that the adapter reads every label and compares none. One of the two is enough to read
that off, and `internal/host/host.go` is what makes it a property of the contract rather
than of one host.

**`.xeno/intents/**`.** The trail is the record and not input.
