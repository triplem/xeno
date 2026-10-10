---
intent: github.com/triplem/xeno#334
phase: 01-requirements
created: "2026-10-10T15:36:47Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f0075f71570c036897f9af29080d13162683c73017b1014f63b3f50429808933
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. **The section exists, in section 9's form, and is section 10.** `docs/orchestrator-evaluation.md`
   gains one top-level section whose heading names it in words. It carries six subsections
   answering what section 9's six answer: what was evaluated, the rows, the decision, the
   measurement — here the three shapes and what each costs — the conditions for revisiting,
   and the mechanisms worth taking. Checked by reading the heading list of the file against
   section 9's.

2. **The pin is one tag, named with the sha, and the sha is the commit the annotated tag
   points at.** The section names `awslabs/aidlc-workflows` `v2.11.0` and commit
   `6a378b53c0a4fe0641ed7d8de8dfff94264d5b6a`, and says that it is the commit the annotated
   tag object points at, because the ref resolves to tag object `4079edbe` and both answer
   to "v2.11.0". The hosted sample is pinned the same way. Checked by resolving both refs
   through the API and comparing what the section says to what comes back.

3. **Five rows, and the fifth is rules and learning.** The table carries the four rows of
   9.2 — the unit of work, where the record lives, who enforces it, who may edit the
   specification — plus a fifth, rules and learning, which OpenSpec did not need and on
   which AI-DLC is not behind. Each row says what changes, in the fourth column, the way
   9.2's does.

4. **Every AI-DLC cell was re-read at the pinned sha, and the rows that moved since #323
   say the new figure.** #323 read at `main` three days before the tag. Three figures in
   that reading are stale and the section carries the current ones: 115 audit event types
   rather than 117, the stars as read on the day rather than 5,084, and the third row's
   enforcement account, which #323 left as a contradiction between two documents and this
   section settles. Checked by the reading being recorded in P4 cell by cell, with the file
   and line each cell rests on.

5. **The third row settles the contradiction #334 names, and says which document is
   current.** The rules guide's unqualified "a sensor result is advisory in this release"
   and the sensor reference's "blocking is enforced for `fire_on: gate`" both still stand at
   the tag. The section says that the reference is current, names the implementing functions
   that make it so, and states the three ways the override is refused, because an enforcement
   claim without the refusal path is the half that flatters.

6. **Every absence the section claims was probed with a positive control in the same
   command, and both answers are recorded.** The section's load-bearing negative is that
   nothing in AI-DLC's core binds an intent to a tracker issue. A probe of an absent thing
   answers as one of a thing that does not match, so each such claim is made only where a
   search of the same tree for something present returned it. Recorded in P4 with both
   halves.

7. **The decision is "not an extension", and the reason is on the page rather than on the
   issue.** The section states the decision, gives the reason that decides it, and weighs
   the three shapes #323 named with what each costs, so that a reader who disagrees can see
   what the disagreement is about without reading an issue thread.

8. **The conditions for revisiting are written where section 9.5's are, and carry the
   alternative the maintainer did not take.** The weaker form of the source-binding
   mechanism — a verdict recording the base and head it was computed over — is named there
   as the thing to revisit when a case arrives, because that is where D-1 put it.

9. **The mechanisms are named and not taken, each pointing at the issue that holds it.**
   #338, #339 and #340 already carry the three #323 named; the section points at them. The
   reviewed-source fingerprint is recorded as a gap, with `docs/process-definition.md:480`
   as the reason, and nothing is taken.

10. **`aws-samples/sample-collaborative-ai-dlc` is a line in the section and compares on the
    same rows.** It is named, pinned, and placed: the same methodology hosted, whose record
    lives in a deployed database rather than in a git tree, which is the row-two difference
    that matters here.

11. **The page's own index names the new section beside section 9.** The paragraph in
    section 1 that currently names section 9 is replaced whole, not edited into, and reads
    as a paragraph.

12. **Sources carries what was read.** An entry naming both repositories, both tags with
    both shas, the date of the reading, and each file by what it was read for, in the form
    the two existing entries have.

13. **Nothing normative moves.** `git diff main -- docs/process-definition.md` and
    `git diff main -- docs/implementation-plan.md` are both empty.

14. **The gates are green and the conventions hold.** `gofmt`, `go vet`, `go test ./...` and
    `./xeno gate verify` as `CLAUDE.md` names them; prose wrapped at 88 characters with
    tables and code blocks exempt.

<!-- xeno:section:non-goals -->
## Non goals

**Not an extension, in any of its three shapes.** No plugin, no sensor manifest, no bridge,
no page of adopter documentation describing how an AI-DLC stage drives `xeno section set`.
#334 settles this and #323 gives the reason. The section records why, which is a different
act from building the cheap one because it is cheap.

**No mechanism taken.** The three #323 named each have an issue of their own (#338 trigger
labels, #339 a reviewer agent outside the phase, #340 a sensor firing on write), and a
sensor firing on write would be a `section 7` change before it was anything else. Naming a
mechanism in an evaluation and taking it are two acts with two different approvals, which is
the discipline 9.6 set.

**No source-binding, in any form.** Not the fingerprint, not the weaker form that records a
base and a head. D-1 is the maintainer's and it says record the gap and take nothing; the
weaker form goes into the conditions for revisiting and not into the tree.

**No change to the process definition or to the implementation plan.** Both are normative
and neither is the agent's to edit. Where this section seems to need either, that is a
finding to write down.

**No second reading at `main`.** #323's market exploration beyond these candidates stays on
#323, for the reason it gives. This section is one pinned reading and does not try to be
current.

**No judgement of AI-DLC's quality.** Section 9.3 drew that line for OpenSpec — "It is not
declined for quality" — and the same line holds here more strongly, because AI-DLC is ahead
of this project on the fifth row and the section says so. A comparison that only found the
other thing worse would not be worth writing down.

**No claim about whether Xeno should cover ideation or operation.** AI-DLC's lifecycle
reaches past the six phases at both ends. Whether this project should is the plan's
question — WP21 holds the front of it and #337 is the finding behind that — and the section
points rather than answers.

**No re-decision of what #323 decided.** That OpenHands is not replaced by any of the four
candidates is 4.1's and 4.5's business, and #335 and #336 carried those sentences. This
section is about the methodology layer only.

**No change to `docs/README.md`.** Its entry for this page describes the page as the
orchestrator record, which was already narrow after section 9 and is narrower after this
one. The line is read in P0's context and left alone, because #334 asks for the
evaluation's own index entry and not for the directory's.

**No table of this project's selling points.** #323 lists five, stated against the
comparison. They belong to whatever document makes a case to a reader outside this
repository, and the evaluation is an architecture decision record whose reader is
inside it.

<!-- xeno:section:constraints -->
## Constraints

**The two normative documents are not editable by the agent, and this intent needs
nothing from them that it cannot cite.** `docs/process-definition.md` section 12 fixes
what approval is; `docs/implementation-plan.md` fixes what a work package label is. The
page quotes neither at length and restates neither: it says what the label does and names
the section, so that a clause changing does not leave a second copy of it saying the old
thing. This is the binding constraint on the whole intent and it is what decides the shape
of the page.

**No invented fields, gates, tools or rules.** A page that said a label meant something to
the runner would be adding to the specification by describing it. The page is written in
the past and present tense about what is in use — this label exists, this person sets it,
this code reads it — and in the conditional only where it names an open question and
points at the issue it is on.

**A label is a string compared by whoever compares it.** The one comparison in the code is
case-insensitive and trims whitespace, and that is a property of one function rather than
of labels in general: a filter in a saved query, a person's eye, and a shell script
grepping a list all compare it differently. So the page gives each label's exact spelling
and says where the one comparison happens, and the spelling discrepancy on
`xeno-needs-decision` is named rather than smoothed over.

**The site's navigation is hand-written and the build is strict.** `zensical.toml` carries
an explicit `nav`, so a page not named there is unreachable from the published site even
though it builds. `.github/workflows/docs.yml` builds with `--strict`, so a link to a page
or anchor that does not exist is a red check before a merge. The two together mean the
index entry is verified mechanically and the nav entry is not, which is why the nav is
read back from the built site rather than assumed.

**A negative result is evidence only when the thing checked was there to be found.** The
page makes four claims of the form "nothing reads this" or "nothing writes that", and each
one is probed with a positive control in the same command, both results reported. This is
the constraint this project cares about most and it bears directly here, because a page
about labels is mostly a page about what does not happen to them.

**Prose wraps at 88 characters; tables and code blocks do not.** The page is largely a
table, which is the shape #359 asks for — a row per label, a column per question — and the
table is the one part that may run long.

**One question at a time.** Three questions are open at the start of this intent: #359's
two, and the spelling of `xeno-needs-decision`. One is asked. The other two are named in
the record with the reason they are not asked yet, which is what keeps them from being
forgotten without putting them on somebody as a batch.
