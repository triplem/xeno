---
intent: github.com/triplem/xeno#359
phase: 05-review
created: "2026-10-10T16:02:09Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8a60c9716fe5ae60679f3908cb5bc9464985150fd4fe26636314cf66f2b7000d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Three deviations, each recorded where it happened and carried forward rather than re-read as satisfied. 03-implementation records that it was redone: it was judged green with the page naming the spelling of xeno-needs-decision as an open point, the maintainer then settled the name, and the page was written again, D-2 recorded, the phase finished again and 04-verification started afresh on the redone predecessor after its unsealed and uncommitted directory was removed. The same phase records that the page came out 12,146 bytes against the designs estimate of 18,000, and that the built navigation was read back locally with a generator this repository does not pin, which the design had left to CI. 04-verification carries two unmet figures rather than fixed ones: criterion 1 says the host lists 34 labels and it lists 33, and 00-intakes decision record names four issues carrying the label where the host says seven. Both are sealed and both are corrected in the results section with the cause in 03-implementations learning record. One further mechanical step is worth naming here because no section owns it: this phase was gate run before its sections were written, in order to read which review rules were owed, which gave it a red verdict and turned every later write into a staged redo; the finish applies them and judges in one step, so the trail carries one verdict covering the content it was computed over, which is the arrangement section 6 describes rather than a departure from it.'
      result: met
      rule: deviations-are-traceable
    - note: 'No interface change. One new document and two one-line additions to the files that index it, and no Go file edited: internal/model/identity.go, internal/host/host.go and internal/host/github/tracker.go are in the scope to be read and were read and not touched. No command, flag, artifact field, template, gate or rule moves. docs/process-definition.md and docs/implementation-plan.md are both at zero lines of diff against main, so nothing normative moves and there is nothing to migrate from. The one act on the host is a label applied to #359 and a comment, neither of which any code reads.'
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: No new dependency. go.mod, go.sum and vendor/ are untouched and nothing was added to any workflow. zensical 0.0.69 was installed into a throwaway virtual environment outside the repository in order to read back the built navigation and to run the positive control on the strict build; it is the version .github/workflows/docs.yml already installs, nothing in the tree references it and nothing in the pipeline changed. The environment was not committed and is not in the repository.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
    - note: 'Documentation lens, which is where this change can go wrong in a way no rule asks about. The page states facts about a tracker that no commit can hold it to. Three kinds are mixed in it and they age differently. What is read off the tree — one label comparison in internal/model/identity.go, three GETs and one POST per adapter, the three operations of the Host interface — is as durable as the code and a change to it would be visible in a diff beside the page. What is cited — section 12 on approval, the plan on work package labels, A103 and A107 — is pinned by being cited rather than restated, so a clause changing leaves no stale copy. What is read off the host is neither: the 33 labels, the nineteen issues carrying xeno-approved, the seven carrying xeno-needs-decision, the three default labels in occasional use and their counts. Every one of those is true of 2026-10-10 and nothing in this repository can notice when it stops being. The page carries the date and the command rather than implying permanence, and 04-verification records the gap as accepted with the design argument for not closing it, which is that the two ways to close it are a test that calls a code host and a page that is a build product. Named as a deviation rather than met, because the honest reading is that a third of this page is a dated observation and a reader who skims past the date will not know which third.'
      result: deviation
      source: lens
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** Held. No document edited: `docs/process-definition.md` and
`docs/implementation-plan.md` are both at zero lines of diff against main, and the page
cites section 12 and the plan rather than restating either. No invented field, gate, tool
or rule: nothing is added to any artifact's schema, no label the runner reads is defined,
and the one label it does read is `internal/model/identity.go`'s constant, unchanged. The
half of #359 that would have been an addition — what `xeno-approved` means after the work
stops — is a question on the issue with the maintainer's options and not a sentence in a
document. The change belongs to #359, on `359-labels-and-the-states-they-stand-for`, and
the issue carries no work package label: WP16 is the nearest fit, nothing in the plan owns
this repository's tracker conventions, and the gap is `00-intake`'s learning record rather
than a label the agent applied.

**What this intent did and what it drafted.** It documented. The page says what exists:
which labels the host carries, who sets each, who clears it, what reads it, and what it
blocks. Nothing on it is new semantics. What was drafted rather than committed is the
answer to #359's first question, put on the issue as three options with their consequences
and their costs and a recommendation with its reason; if the maintainer takes the
recommendation, the sentence section 12 would need is drafted after that, in their commit
and not this one. The second question is not asked yet because every option on the first
changes what it asks.

**The page answers "who acts on a label" in two halves, and the second is the one worth
reviewing.** The first half is the tables: who sets each label and who clears it. The
second is that the host cannot tell the agent's acts from the maintainer's, because there
is no bot identity and the agent runs under the maintainer's token — `gh api user --jq
.login` returns `triplem`. So the authorship field of a label or a comment answers the
question wrongly, and the written convention "the decision is the maintainer's; the
transcription is not" is what carries the truth into a place a verdict seals.
`XENO-0286`'s intake and this intent's D-2 both use it, which is what the page cites. The
same property is noted against the open question on `xeno-approved`, because an option
that recorded a withdrawal as a comment would record it as unfalsifiably as an approval;
that is a property of the arrangement rather than an argument against the option, and the
recommendation on the issue does not turn on it.

**#338 was not absorbed.** The page says triggering is #338's, says why section 12 keeps
issue commands out, and lists no trigger label, because none exists. The maintainer's
answer on that issue — one label, no analysis trigger — is a specification change nobody
has made, and a page listing `xeno-start` would have been documenting a decision as a state
of the repository. The word does not occur on the page.

**One question, not three.** Three were open when this intent started: #359's two and the
spelling of the new label. One was asked. The spelling was then settled by the maintainer
in session and is recorded as D-2 of `03-implementation` with `decided_by: triplem`, which
is why the page states a name rather than a discrepancy. #359's second question is named in
the record and on the page as following from the first.

**Negative results.** Four claims of absence, each probed with a positive control in the
same command and both results reported in `04-verification`: the one label comparison, the
absence of any label write, the empty diff against the normative two, and the absence of
`xeno-start` from the page. The claim that `gh issue edit --add-label` exits 0 without
applying a label was carried into this intent from an earlier session, was probed under the
conditions that make a failure visible, and did not reproduce; the page prescribes the
read-back on its own merits instead of recording a defect today's probe cannot support.

**Two redos, and what the trail does and does not show.** `03-implementation` was judged,
then redone when the maintainer settled the label's name, and `04-verification` was started
again on the redone predecessor after its unsealed and uncommitted directory was removed.
This phase was judged and then redone when the authorship fact arrived. Each redo moved the
artifact and its verdict together, which is section 6's mechanism, so the trail carries one
verdict per phase covering the content it was computed over and does not carry the
intermediate states. The residual risk section says which figures in earlier phases the
last change left behind and why they were not chased.

**What a reviewer should look at first.** The paragraph in `docs/labels.md` under "The one
label Xeno reads" that says nineteen issues carry it against more than a hundred and forty
intents in the trail, and the inference drawn from it. It is the page's strongest claim —
that the durable record of an approval is the sealed intake sentence and not the label —
and it rests on a count, on A107, and on `approved` being absent from the host's label set.
`04-verification`'s gaps section says what it does not evidence.

**Conventions.** Prose at 88 characters, tables and code blocks exempt, checked over both
changed markdown files; headings name their sections in words; commit subjects are
Conventional Commits with `Refs #359` in the footer and `Closes #359` on the last; the
pull request title is a Conventional Commit because it becomes the squashed subject.

<!-- xeno:section:release-notes -->
## Release notes

**A page that says what each label on this tracker means, who sets it, who clears it and
what reads it.** `docs/labels.md`, linked from the documents index and from the published
site's navigation.

The labels on this repository's tracker had come to stand for states of an issue and of
the intent it becomes, without anywhere saying so. The page records the 33 labels the host
carries, in three groups ordered by how binding they are: the one label Xeno reads on any
tracker, `xeno-approved`, with the `/xeno approved` comment it is useless without; the two
this repository has given itself, `xeno-needs-decision` and `wp0` to `wp20`, which no code
reads; and the host's own set, which has no process meaning here. Each group carries a
table whose columns are the questions that prompted the page — who sets it, who clears it,
what reads it, what it blocks.

It opens with three things worth knowing before any table. A label does not hold an
intent's state: the state is in `.xeno/intents/`, it is what `xeno intent status` prints,
and nothing is derived from a label or checked against one. Nothing in Xeno ever writes a
label, because the adapter contract is three operations and none of them is a label write.
And who acted cannot be read off the host at all: there is no bot identity here, so a
label the agent applies and a comment it posts arrive under the maintainer's own account,
which is why writing "the decision is the maintainer's; the transcription is not" into an
artifact carries a distinction the tracker cannot.

Three questions are named in the page as open, each with its issue, so that its silence is
not read as an answer: what becomes of `xeno-approved` when the work it authorised stops,
whether a fresh approval is owed after a decision reverses it, and whether a label should
ever start work rather than permit it, which is #338's. The first is on #359 with its
options, and the second follows from it and is asked after.

Nothing in the specification changed and no code changed. The page cites section 12 for
what approval is rather than restating it, so there is no second copy of a normative
sentence to drift.

<!-- xeno:section:residual-risk -->
## Residual risk

**The page can go stale and nothing will say so.** A third of it is read off the host — the
33 labels, which issues carry which, the counts — and it is true of 2026-10-10. The page
carries the date and the command that reads the set, which is the mitigation and not a fix.
The two ways to close it were weighed in the design and both were rejected: a test calling
a code host, which the gate path forbids, and a generator, which makes the page a build
product. The plan's own treatment of the work package labels is the precedent. A reader who
skims past the date will not know which third of the page is a dated observation.

**The page's central claim depends on code it does not sit beside.** That exactly one
label is read anywhere, and that nothing writes a label, are facts about
`internal/model/identity.go` and `internal/host/`. If a later change reads a second label
or adds a label write, nothing connects it to this page, and the page would then be wrong
in the way that matters most — it would understate what a label does. The mitigation is
that the page names the file and the function, so a reader checking is one grep away; there
is no mechanism.

**A reflow script merged both of the page's code blocks into prose, and the strict build
did not notice.** The page is held at 88 characters by a script that wraps paragraphs and
leaves indented blocks and table rows alone. Run a second time, it swallowed the two
command blocks into the paragraphs above them, so `gh api repos/triplem/xeno/labels
--paginate --jq '.[].name'` and the two lines under "Applying a label" rendered as
sentences. Nothing caught it: `zensical build --strict` reported no issues, because an
unindented command line is valid prose and a link that still resolves is still a link,
and the width check passed because the merged lines came out under 88. It was found by
reading the page back as a page rather than as a diff, which is what `CLAUDE.md` asks for
and what a width check cannot do. The blocks are restored and the rendered page now holds
two `<pre>` elements. The lesson is in this phase's second learning record: a tool that
reflows prose has to be run once and read back, or not used on a file that has code in
it.

**The page gained a section after `03-implementation` was sealed, and the figures in that
phase are now low.** It records 193 lines and 12,146 bytes; the page is 215 lines and
13,442. What was added is the paragraph on authorship — that there is no bot identity on
this repository, so an agent's label and an agent's comment arrive under the maintainer's
account, and the written convention is therefore load-bearing — with a sentence in the
section on applying a label and a sentence on the open question that inherits the same
property. The fact reached the agent after `05-review` had a verdict. This phase was
redone with it rather than the whole tail of the intent being redone, and the trade is
stated rather than hidden: redoing `03-implementation` would have moved the predecessor
hash in `04-verification`'s lock, and that lock can only be rewritten by a `phase start`,
which refuses a second one, so the honest repair would have meant removing two committed,
sealed phase directories and running them again. That is a larger act than the
inaccuracy warrants, and it is the same judgement `04-verification` already made about two
wrong figures it corrected in prose rather than by rewriting the phases that hold them.

**`04-verification`'s lock records an earlier hash of `docs/labels.md`.** The page
changed twice after that phase was given it, once for the paragraph on authorship and
once for the repair of the code blocks. G-Freshness's second half is what reports this, and it
reports nothing without a commit range: `.github/workflows/xeno.yml` runs `xeno gate
verify` with none, so no check in this repository will raise it. It is written here
because that is the only place it will be visible, and because a reader comparing the lock
against the tree deserves to find the reason rather than a mystery.

**The question on #359 is open and no gate can see that it is.** This intent deliberately
did not record it as an `open-questions` entry, because a question about a section 12 clause
can only be answered by a commit the agent may not make, and `G-Questions` would have held a
documentation page at a red P5 until it was. The consequence is that the trail of this intent
is silent about the question: `G-Questions` passes because nothing was raised. The
`xeno-needs-decision` label on #359 is the only thing carrying it, and nothing checks that
either. `01-requirements`'s learning record proposes the rule that would close this.

**Two sealed artifacts carry wrong figures.** The label count in `01-requirements` and the
list of labelled issues in `00-intake`. Corrected in `04-verification`, cause in
`03-implementation`'s learning record, and permanent where they are. A reader arriving at
either phase alone reads the wrong number, which is the cost this project's own rule about
sealed findings describes.

**The approval this intent ran under is itself an instance of what the page leaves open.**
#359 carries `xeno-approved` and the approval comment, and the page records that nothing
removes the label and nothing reads it after P0. So if the maintainer's answer to the open
question reverses part of this work, the issue will still carry an approval that permits a
second intent to start on it. That is the state the question is about, demonstrated by the
intent that asked it.

**The work package label is still absent from #359.** The third standing rule asks the
issue to carry the label of its package; WP16 is the nearest fit and nothing in the plan
owns this repository's tracker conventions. The gap is `00-intake`'s learning record and
the label remains the maintainer's act, so the rule is unmet on the host at merge time and
knowingly so.
