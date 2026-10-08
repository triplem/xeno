---
intent: github.com/triplem/xeno#330
phase: 00-intake
created: "2026-10-08T17:36:16Z"
schema_version: "1.0"
runner_version: dev+e22a533.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d361d6b153f0589e0420ffb05c15f53cab105a9d5502040d6c7f5968ebec1965
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: A comment whose first line is the word approved, the rest being the reason, beside the label approved; both required
      rationale: One call on either host and no lookup of who the author is; the label, which takes a right the host grants, is what keeps a stranger's comment from counting, and the comment is what carries the reason the intake quotes. The latest maintainer comment was set aside because GitLab notes carry no author association and the latest one is usually a reply; the label event because it records no reason and needs two endpoint shapes behind one port.
      decided_by: the maintainer
      proposed_by: the agent
    - id: D-2
      chosen: An issue with a milestone is held until that milestone is the earliest open one, by due date with undated ones last; --now starts it anyway and the intake says so
      rationale: It is the one reading in which earlier means something the host can answer, and the cost, one more call to list the milestones, is paid only by a project that uses them. Holding on any open milestone would refuse the issue in the current milestone too; holding until the due date would hold work until the day it was meant to be done.
      decided_by: the maintainer
      proposed_by: the agent
    - id: D-3
      chosen: The agent makes the specification commit, after the clause was placed in section 12 and approved there
      rationale: The first standing rule exists so that a person decides the specification, not so that a person types it; the wording was read in the document as a diff and approved as placed, and the exception is recorded here and in the commit message so that it stays one.
      decided_by: the maintainer
      proposed_by: the agent
---

# Intake

<!-- xeno:section:problem -->
## Problem

> **github.com/triplem/xeno#330** — An intent can be started from an issue nobody approved
>
> Found by doing it, 2026-10-08. #95 is the case and it is reopened.
>
> `xeno intent start --for ISSUE` takes any issue. An issue that proposes
> something and an issue that is approved work look identical to it, and to
> every phase after it: the intake quotes the issue's body, the phases are
> judged against that intake, and the pull request that finishes the work
> closes the issue with `Closes #N`.
>
> So an issue that said "this **could** be used instead of prs" was carried
> through six phases and closed by the merge, having had its central
> question — adopt this tool at all? — answered by the agent that ran the
> intent. The maintainer's objection is the whole of the defect: *the issue
> was not approved.*
>
> ## What the process already has, and why it did not fire
>
> The intake template has an `open-questions` section, and G-Questions
> refuses a P5 whose questions are not resolved into a decision, a
> withdrawal or a confirmed assumption. That mechanism exists exactly so
> that a question reaches a person before an intent closes, and it works —
> the corrective intent for #95 uses it, and its P5 cannot be decided until
> the question it raised is answered.
>
> But it fires only when a phase chooses to raise a question. Nothing makes
> a phase ask whether the work was wanted, and nothing can: the agent that
> decides there is no question to ask is the agent that already decided the
> answer.
>
> ## Where a check can live, and where it cannot
>
> **Not in a gate.** `xeno gate ...` never touches the network, so no gate
> can ask the host whether an issue carries a label, a milestone or an
> approving comment. A gate that read a local copy of that answer would be
> judging what the agent wrote down about its own authorisation.
>
> **In `xeno intent start`.** Starting an intent may use the network —
> section 12 puts it on that side deliberately, and since #288 it already
> reads the issue through the adapter. Reading one more field of the same
> issue is not a fifth adapter operation; it is the same read.
>
> So the shape is a refusal at the start: an issue that does not carry the
> agreed marker does not become an intent, and the refusal names what is
> missing.
>
> ## What the marker should be is the open question
>
> Candidates, each with a cost:
>
> - **A label**, say `approved`. Cheap, visible, and the host already has
>   labels for the work packages. It is also one click, which is either the
>   right weight or too little.
> - **An assignee.** "Somebody has taken this" is close to "somebody wants
>   this", and it is not the same thing.
> - **A milestone**, which says when and not whether.
> - **An approving comment in a form the adapter can read.** Heavier, and it
>   carries a reason — which is what the trail wants to record, since the
>   intake would then quote why the work was approved and not only that it
>   was.
>
> Whatever it is, the intake should record it: a phase that says it was
> authorised, and by what, is a phase whose authorisation a reader can
> check. That is a sentence in the problem section rather than a new field.
>
> ## The sibling finding, already recorded
>
> `phase start` carries the issue's **body** into the intake and not its
> comments. On #95 the analysis and the three answers that settled the
> design were comments, so the intent that needed them never saw them, and
> got two of three wrong. That is recorded as XENO-0277's P0 learning
> against `internal/runner/tracker.go` and it is the other half of the same
> morning: an issue is what it says plus what was said about it.
>
> ## Done when
>
> `xeno intent start` refuses an issue that does not carry the marker a
> person decided on, with a refusal that names it; the intake records that
> the intent was authorised and by what; and the decision about which marker
> is recorded, since every candidate above is a different claim about what
> approval means.
>
> Nothing here is a gate, and nothing new enters an artifact beyond a
> sentence in a section that already exists.

The issue as `xeno phase start` read it, and what the read did not carry is this time the
point: the maintainer's answer sits in a comment, and `phase start` quotes the body alone.
The answer, 2026-10-08T15:10:28Z, is one line — *a comment as well, if applicable, a
milestone. if there is a milestone, the issue should not be worked on any earlier except on
an explicit query* — and it chose the fourth candidate on top of the first: the issue
carries the label `approved` already, and approval is that label plus a comment. Two
readings were still open, the form of the comment and what a milestone holds back, and both
were put to the maintainer one at a time with options, consequences and a recommendation;
the choices are recorded on the issue and in this phase's `decisions`.

**Where this intent stands against the rule it builds.** #330 carries `approved` and
`wp12`. Its one comment does not begin with the word `approved`, because the word was fixed
by the clause this intent exists to enforce, on the same afternoon; by that clause this
issue would be refused today. The intake therefore says what it can: the intent was
authorised by the label and by the maintainer's answer of 2026-10-08T15:10:28Z, read by hand
from the comment the runner does not yet read. It carries no milestone, and the repository
has none.

**The specification moved first.** Section 12 gained the paragraph *Starting an intent* in
commit `8fb365d` on this branch, written by the agent on the maintainer's instruction after
the wording was placed in the document and approved there — an exception to the first
standing rule, recorded here and in the commit message so that it stays one. The clause
fixes the label name and the word, so nothing enters `project.yaml`, and it fixes what the
intake records as a sentence in a section that already exists, so nothing enters any
artifact's field set.

<!-- xeno:section:scope -->
## Scope

This intent makes `xeno intent start` refuse an issue nobody approved, as section 12 now
says, and makes the intake say by what the intent was authorised.

- **The refusal.** With a tracker block configured, `intent start` reads the issue and
  refuses unless it carries the label `approved` and a comment whose first line is the
  word `approved`; the refusal names which of the two is missing. With a milestone on the
  issue, it reads the project's open milestones and refuses while an earlier one is open,
  naming it, unless `--now` is passed. The reads that cannot happen — no token, no such
  issue, an issue on a host the block does not reach — refuse too, since none of them can
  say the work was approved.
- **Both adapters.** Reading an issue stays one operation and grows to the issue's labels,
  milestone and comments, and the project's open milestones where the issue carries one:
  up to three calls on GitHub and on GitLab, in each host's own words, answered into
  `model.Issue`.
- **The sentence in the intake.** At P0, `phase start` reads the same issue again and
  writes above the quote who approved it, when and why, and that it was started ahead of
  its milestone where that is still the case.
- **The flag, the doc and the tests.** `--now` in the usage constant and in
  `docs/commands.md`, which a test holds identical; runner tests on a fake host for the
  refusals, the override and the sentence; adapter tests for the new reads on both hosts.
- **The register.** One row in `docs/assumptions.md` for what the clause left to the
  implementation: that a read which cannot happen is a refusal and not a start, and that
  the intake's sentence is recomputed at P0 rather than carried from the start.

What it does not do.

**It does not carry the issue's comments into the intake.** XENO-0277's P0 learning asks
for that against `internal/runner/tracker.go`, and the comments are now read, so the step
is small; it is still a change to what the problem section quotes, which is its own decision
and not a side effect of this one.

**It does not make the label or the word configurable.** The clause fixes both. A project
that wants another word is asking for a specification change, and a field in
`project.yaml` would be an addition to Appendix A nobody asked for.

**It does not touch a gate.** Section 12 says why in the clause itself, and the import
graph test that keeps `internal/gates` from the host adapters is the mechanical form of it.

**It does not read who applied the label**, or whether the comment's author may approve
anything. The label needs a right the host grants, which is the one check of standing the
design relies on; a comment's author is quoted and not judged.

**It does not check Jira, assignees or due dates.** 1.1 is the plan's place for a third
host, and the other two candidates were weighed on the issue and set aside.

<!-- xeno:section:context-rationale -->
## Why this context

Read for this phase, and why.

- `docs/process-definition.md` section 12, including the new *Starting an intent*
  paragraph, Appendix A's tracker block, and section 8's rule on how a question is put —
  the clause being implemented, the configuration it reads, and the way the two decisions
  behind it had to be asked.
- `docs/implementation-plan.md` WP12 — the adapter contract's four operations and the two
  requirements that keep 1.1 cheap, which bound how the read may grow.
- `internal/runner/runner.go` `IntentStart` and `StartPhase`, and `internal/runner/
  tracker.go` `readIssue` and `intake` — where the refusal goes, where the issue is read
  today, and where the sentence is written.
- `internal/host/host.go`, `internal/host/github/tracker.go`, `internal/host/gitlab/
  tracker.go` and `internal/model/identity.go` — the port, the two adapters and the shape
  the answer travels in.
- `internal/runner/tracker_test.go` and `runner_test.go` around `IntentStart`, both
  adapters' `tracker_test.go`, and `cmd/xeno/main_test.go`'s intent start tests — what
  the existing tests assume, since every one of them starts an intent against a tracker
  block with no host behind it and will have to be given one.
- `docs/commands.md` and the `usage` constant — the two copies a test holds identical.
- `docs/assumptions.md` A82 and A100 — the rows nearest to this change in shape.
- Issue #330 with its comment, #95, and XENO-0277's P0 `learning.yaml` — the case, the
  answer, and the sibling finding this intent leaves alone.
