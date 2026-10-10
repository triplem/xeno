---
intent: github.com/triplem/xeno#334
phase: 05-review
created: "2026-10-10T16:02:43Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 50105e941fdae0c61a07eccba06dc82ae3d218650f8a5aa526f69ecc739dbb51
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Two deviations, both written in P3 with their cause, and one gap written in P4 with the reason it was not repaired. The section came out 23,273 bytes against the design estimate of about 16,000, because settling the third row against the code took five paragraphs where 9.2 needed one cell and a clause; the budget was set at 140,000 against a scope of 37,939 and absorbed it with no finding. And docs/process-definition.md:480 is cited twice by two different means, by file and line in 10.5 where a reader has to find it and by quotation attributed to the normative document in 10.6, which is how this page cites its own documents everywhere else; the design said nothing about the second case, so that is a gap in the design filled rather than a departure from it. The gap P4 records is that 10.6 does not name AI-DLC reviewer agent, which #339 holds: it was found by the grep meant to confirm criterion 9, after P3 was decided and the document committed, and repairing it would have staled P3 while P4 was running and could not be started again. Everything else is the design: four hunks in the one file in the order the impact section named them, the section 1 paragraph replaced whole, no subsection for the hosted sample, no mechanism taken.'
      result: met
      rule: deviations-are-traceable
    - note: No interface change. One document gains a section; no command, flag, artifact field, template, gate or rule moves, and the generated CI wrapper is untouched, so nothing in an adopter repository changes. The one thing a reader of the page will notice is that section 1 paragraph now names two later sections instead of one, which is the page own index and not an interface. The page frontmatter revision goes to 4 and its date to 2026-10-10, which is what a whole new section did the last time one landed.
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: No new dependency. go.mod and vendor/ are untouched and nothing is installed; git diff main over the Go tree is empty. What the reading needed was gh and a tarball of one commit extracted to /tmp, neither of which enters this repository, and the section cites AI-DLC own files rather than vendoring anything from them. The whole of the change is 334 added and 9 removed lines of Markdown in one file.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
    - note: 'Negative-results lens. This section rests on four negative claims and the project rule is that a tool asked about something absent answers as it does about something that does not match. Each was probed with a positive control in the same command and both halves are in P4 results: nothing in AI-DLC core binds an intent to a tracker, one hit which is a prose citation of its own issues, against a control of five files for intents.json; nothing in core hashes the record, no hits, against three files for Unit Source Fingerprint; no file here is misformatted, zero lines, against the same binary printing a deliberately misformatted file path; and nothing normative moved, a zero-line diff, against 334 9 for the file that did change. What this does not cover is the one claim that is negative about an absence of prose rather than of code: that #323 sentence about later agents revising earlier artifacts does not reproduce at the pin. Its control is that the same search found three other things it was looking for, which is weaker than the others, and the section handles it by not making the claim at all and writing the two sentences that are there instead.'
      result: met
      source: lens
    - note: 'Sealed-finding lens. Every claim in this section is about somebody else repository at one commit, and a sealed finding is permanent. Two things follow and both were done rather than intended. The reading was taken from a tarball of the exact pinned commit extracted to disk, not from main and not from the API per file, so no claim can have come from a tree that had moved; and every claim has an address in P4 cell-by-cell table, fifty-odd of them, so a later reader checks rather than trusts. What remains is the one thing no discipline here can fix: three figures from #323 were already wrong three days after it was written, so three figures in this section will be wrong three days after it merges. The ones that move are dated in the prose rather than stated flatly, which is the most a document can do about it.'
      result: met
      source: lens
---

# Review

<!-- xeno:section:release-notes -->
## Release notes

**The evaluation now records a third decision, and it is the one that compares with the
process rather than with the platform.** `docs/orchestrator-evaluation.md` gains section 10
for AI-DLC, the AI-Driven Development Life Cycle, pinned at `awslabs/aidlc-workflows`
`v2.11.0`. #323 did the reading at `main` and weighed the three shapes; a reader of the page
met an open comparison in an issue thread and now meets a decision with its reason, its
conditions for revisiting, and the mechanisms worth taking, in the form section 9 has for
OpenSpec.

**The pin says which sha is which.** `v2.11.0` is an annotated tag: the ref resolves to tag
object `4079edbe`, which points at commit `6a378b53c0a4fe0641ed7d8de8dfff94264d5b6a`, and
the reading is at the commit. A reader who resolves the tag himself gets the other sha, so
the section and the Sources entry both say it. The hosted sample,
`aws-samples/sample-collaborative-ai-dlc`, is a line in 10.1 pinned the same way at `v2.2.0`,
commit `bc988d0e`, and is held to what its own README states.

**Five rows, because AI-DLC needs one OpenSpec did not.** The four rows of 9.2 all change
and the fourth reverses. The fifth, rules and learning, is the row on which this project is
behind: AI-DLC has the mechanism section 10 of the process definition describes and this
project has not built. Where they differ is the review — a tick at a gate against a merge
request against the rule set, the stronger review against the one that happens.

**The one point the issue asked to have settled is settled.** Two of AI-DLC's own documents
read literally disagree about whether a sensor can refuse, and both sentences are still
there at the tag: the rules guide says a sensor result is advisory in this release,
unqualified, and the sensor reference says blocking is enforced for `fire_on: gate`. The
reference is current, and the code is why — `fireGateSensors` narrows to gate-fired blocking
sensors, `enforceBlockingGateSensors` calls `error()`, the `gate-start` path calls it, and the
write-path hook exits 0 always by its own comment. The guide's sentence sits two lines under
a paragraph that describes gate-fired sensors, so the file disagrees with itself inside one
screen.

The row then states the enforcement three qualifications deep, because the short version
flatters it: gate only, never a write; opt-in only, with every one of the six shipped sensors
declaring `advisory`, so out of the box nothing refuses anything; and overridable by a
person, against three refusals — autonomous mode, a wrong answer, and a missing fresh
authorisation receipt — which is a better override than most things have and still the
opposite end of section 7 from "red is red".

**The decision is not an extension, and the reason is on the page.** AI-DLC answers layer 2
inside the harness, which section 2 and section 4.4 already declined, and the specification
is a stage output drafted by an agent, which the first standing rule is the opposite of. The
difference from OpenSpec is that a person stands in front of every version, so this is two
defensible answers to who owns the specification rather than an answer against the absence of
one — and only one of them can hold in one repository. 10.4 carries #323's three shapes with
what each costs, and names the cheap one as something nobody has to decide, since an AI-DLC
sensor whose `command` is `xeno gate run` is one Markdown file in somebody else's project and
nothing here stops it today.

**Nothing is taken, and the one mechanism that needed a decision got one.** AI-DLC binds a
review to the bytes it approved: a receipt carrying a fingerprint over a committed listing of
path to blob OID, and `aidlc attest resolve --diff base..head` answering verified, drifted,
unattested, unverifiable or indeterminate from `(base, head)` alone in any clone. The
question of whether a Xeno verdict should do the same was put to the maintainer on #334 with
three options and answered: record the gap, take nothing. So 10.6 describes the mechanism in
full and records the plain sentence it produced — **Xeno seals the record of a review; it
does not seal the identity of the source that was reviewed** — with "a verdict says what it
judged, by content and not by commit" as the reason, and the weaker alternative written into
10.5's conditions rather than into the tree.

The other three mechanisms are named and not taken: a sensor firing on write, which is #340
and a section 7 change before it is a line of code; the learning gate's ritual, whose
`RULE_LEARNED` audit row is the write-back this project's learnings do not get; and the
admission conflict check, which is **named and declined**, because it is the one place in
that tree where a model sits in a refusal path and section 7 opens by saying all gates are
deterministic and model free.

**Nothing normative moved.** `git diff` over `docs/process-definition.md` and
`docs/implementation-plan.md` is zero lines. The change is 334 added and 9 removed lines in
one document: the new section, the Sources entry, the frontmatter to revision 4, and the
section 1 paragraph that indexes the page's later sections, replaced whole so that it now
names 9 and 10 together and says what distinguishes the second.

<!-- xeno:section:residual-risk -->
## Residual risk

**The section is a claim about one commit, and that commit is already not `main`.** AI-DLC's
default branch moved on the day of this reading and three preview tags were cut in the three
days around `v2.11.0`. This is the form section 9 chose and the form #323 argued for, so it
is a cost accepted rather than an oversight. What has no repair is the narrower case: the day
AI-DLC adds the gate-fired qualification to the sentence at `docs/guide/09-rules-and-the-learning-loop.md:140`,
10.2's third row is still correct about the pin and reads as an argument about a
contradiction that is no longer there. Nothing here notices that, and nothing could.

**Three figures will be wrong three days after this merges, in the same way three of #323's
were.** 115 audit event types, 5,123 stars, three preview tags around the release. The ones
that move are dated in the prose rather than stated as facts about the project, which is the
most a document can do. The ones that decide anything — the two contradicting sentences, the
four places in the code, the six advisory manifests, the plugin contract's two "not yet"
notes — are what the decision rests on, and those change only if AI-DLC changes its design,
which 10.5 already names as a condition for revisiting.

**10.6 does not name AI-DLC's reviewer agent, and #339 holds it.** The reviewer is invoked as
a separate sub-agent after a stage body produces its artifacts and before the gate, and it
never blocks. That is a mechanism of exactly the kind 10.6 is for. It was found by the grep
meant to confirm criterion 9, after P3 was decided and the document committed, and repairing
it then would have staled P3 while P4 was running and could not be started again. So it is a
gap in P4 rather than a paragraph in the section: one missing pointer, with the mechanism
still reachable through #323, and one paragraph owed to whoever next touches 10.6.

**The hosted sample's own rows are unread.** Not by #323, not here. The line it gets in 10.1
is held to its `README.md` at `bc988d0e` and claims nothing about its gates, its sensors or
its approvals — and that README names human gates, a typed traceability graph and per-stage
cost figures as things it has, each of which would be a cell if somebody read them. The line
is therefore incomplete rather than wrong, which is the trade a line is for, and whether the
hosted comparison earns a reading of its own is a question this section leaves open.

**One negative claim has a weaker control than the other four.** That #323's sentence about
later agents revising earlier artifacts does not reproduce at the pin is an absence of prose
rather than of code, and its only control is that the same search found three other things
it was looking for. The section handles it by not making the claim: it writes the two
sentences that are there — reviewed outputs remain frozen, and a person's jump back offers
keep, modify or redo — and the fourth row says "a person's jump back", which is weaker and
checkable. A reader comparing this row with #323's will find it softer, and the reason is in
P4's mapping rather than on the page.

**Nothing re-checks the width of this page's prose.** The first pass of section 10 was written
at 90 characters and reached a committable state only because the width was measured again by
hand, in characters rather than bytes, which is the distinction a shell check gets wrong in a
document full of em dashes. `gofmt` says nothing about Markdown and `xeno gate verify` judges
the artifacts of the trail rather than the files the trail is about. The next long document
written here will make the same mistake in the same way.

**#334 carries no work-package label.** Neither did #336, whose intent shipped as `5044a7a`
without one, and XENO-0286 met the same thing from the other side. No package in the plan
owns this repository's own documents, so the third standing rule is breached visibly every
time one of them changes. It is written down in P0's learning and in P4's gaps, and nowhere a
reader of the plan would find it.

**What a green P5 proves here.** That the record was sealed, that the four gates
`CLAUDE.md` names are green, and that every claim the section makes about AI-DLC has an
address at a pinned commit which a reader can check. Not that the reading was the right
reading, nor that the fifth row is the right fifth row, nor that the reason in 10.3 decides.
Those were read and not tested, P4's third column says which criteria are of that kind, and
the honest form of a document's verification is to say so rather than to invent a check. It
is also, as 10.6 says of this project generally, exactly the distinction the gap in that
subsection is about.
