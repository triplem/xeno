---
intent: github.com/triplem/xeno#359
phase: 02-design
created: "2026-10-10T15:39:57Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 07629c2e1bb22f407de5f3d9578eabcfdc832a9f4ee25581d68769a45281b930
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The page is `docs/labels.md`, a document of its own, indexed in `docs/README.md` under
"Running it" and named in `zensical.toml`'s navigation beside `commands.md`.** The
alternatives and what each costs are below; this is what follows from choosing it.

**It is organised by what reads a label, not by what a label is about.** Three groups,
and the order is the order of how binding they are. First the one label the runner reads,
`xeno-approved`, with the comment it is useless without. Then the labels this repository
has given itself, which no code reads: `xeno-needs-decision` and `wp0` to `wp20`. Then the
host's own set, which has no process meaning here at all. A reader who wants to know
whether a label can stop something gets the answer from which group it is in, before
reading a word of the row. Grouping by topic instead — approval, questions, work
packages — would have put `xeno-approved` and `xeno-needs-decision` side by side as two
`xeno-` labels, which is exactly the resemblance that misleads: they share a prefix and
nothing else.

**Each group carries a table whose columns are #359's questions.** Label, who sets it, who
clears it, what reads it, what it blocks. Five columns, because "who acts on a label and
when it is removed" is two questions that have different answers for the same label — the
person who sets `xeno-needs-decision` is the agent and the person who clears it is
whoever answers — and a single "who" column would have had to pick one. "What it blocks"
is the fifth because it is the question a reader actually arrives with, and for every
label but one the answer is nothing.

**The state of an intent is not in a label, and the page says so first.** This is the one
place the page risks making things worse. #359 observes that the labels have come to
represent states, and they have, but only to people: the intent's state is in
`.xeno/intents/`, it is what `xeno intent status` prints, and no label is derived from it
or checked against it. A page that tabulated labels as states without saying that would
invite the reading that setting a label changes something. So the page opens with what a
label is here — a person's note to other people, with one exception that is a
precondition — and the tables follow from that sentence rather than standing on their own.

**Nothing is written in the imperative.** The page describes; it does not instruct a tool.
Where a label's meaning is fixed by section 12, the page says what the runner does and
cites the section rather than paraphrasing the clause, so there is no second copy of a
normative sentence to drift. Where a meaning is this repository's habit, the page says
whose habit and since when, with the issue number, so a reader can see that it is a
convention and not a rule.

**The open points are named in the page with the issue they are on.** The lifecycle of
`xeno-approved` after a decision, whether a fresh approval is owed, the spelling of
`xeno-needs-decision`, and triggering. Four open points, each one sentence, each with its
issue. They are in the page rather than only in this trail because the page is where
somebody will look, and a page that quietly omits what it does not know reads as complete.

**The operational paragraph is about reading back, not about a defect.** Applying a label
to an issue on this host was carried in as a defect — `gh issue edit --add-label` exiting
0 without applying — and the probe did not reproduce it. So the page gives both forms,
says that the `gh api` form is the one used when the outcome matters, and gives the
read-back as the practice, on the ground that an exit code is a report about a request and
not evidence about an issue. What the page does not do is record a defect one session saw
and the next could not: the finding is in this intent's learning record, where a reader
can see both observations, and the page carries only what a probe today supports.

<!-- xeno:section:alternatives -->
## Alternatives

Four places the page could go, and the argument is about who the reader is.

**A section of `CONTRIBUTING.md`.** It already explains the commit footer, the DCO and
the host settings the conventions depend on, and a label is the same kind of thing: a
convention of this repository that a contributor has to keep. *Consequence:* no new page,
nothing to index, and the one file an outside contributor is pointed at would hold all
the tracker conventions together. *Cost:* two readers get the wrong document. A
contributor reading `CONTRIBUTING.md` would find `wp0` to `wp20` explained, which is the
maintainer's construction plan and nothing a contributor can act on; and somebody asking
what `xeno-approved` means to the runner would find the answer in a file about how to
send a change. `CONTRIBUTING.md` is also not on the published site — `docs/README.md`
links it to the host for exactly that reason — so the page would be the one piece of
this documentation a reader of the site cannot reach. It is 7,763 bytes and this is on
the order of its length again, which would make the labels the largest thing in a file
about contributing.

**A section of `docs/commands.md`.** That page already carries the user-facing sentence
about the approval pair, under `xeno intent start`. *Consequence:* the one label the
runner reads would be documented exactly where the command that reads it is, with no
duplication at all. *Cost:* it answers a quarter of #359. `commands.md` is "every command
with its flags, which is the text `xeno --help` prints", and a test holds it against the
usage constant; a labels section in it would be the only part of the page not
corresponding to a command, and `wp0` to `wp20`, `xeno-needs-decision` and the host's own
set correspond to no command at all. The structure would have to be bent for three
quarters of the content.

**A section of `docs/the-trail-in-this-repository.md`.** That page is about how this
repository governs its own construction, which is precisely what a label convention here
is. *Consequence:* the honest home for the half of the content that is this project's own
habit, and it sits beside the other things that are true of this repository and of no
other. *Cost:* it splits the page in two. `xeno-approved` is not a fact about this
repository, it is what Xeno reads on any tracker, and putting it on a page titled "the
trail in this repository" would say it is local when it is not. The split is the thing
#359 is asking to be undone: labels in one place, with their differences stated.

**A page of its own, `docs/labels.md`.** *Consequence:* one place with all of it, grouped
so that the difference between what Xeno reads and what this repository has agreed is the
page's own structure rather than a caveat inside somebody else's page. It is linkable,
which matters because the answer to "what does this label mean" is a link somebody pastes
into an issue, and a link to a page beats a link to an anchor two thirds down a page about
something else. *Cost:* a page is a thing to maintain and to find. It must be added to
`docs/README.md` and to `zensical.toml`'s `nav` or it is unreachable from the site — two
edits that are easy to forget and that the strict build catches for only one of them. And
one more document is one more place a reader has to be told about: the index gains an
entry under "Running it", which is the section a reader looking for conventions reads.

**Chosen: the page of its own.** The reason is the one the alternatives keep running
into. The content has two halves with different standing — what Xeno reads on any
tracker, and what this repository has agreed among itself — and every existing page is
about one half or the other, so putting the content in any of them makes one half look
like the other. That confusion is the fault #359 reports: labels that came to mean
something without anybody saying so. A page whose first job is to separate the two cannot
be a section of a page that has already chosen a side.

Two further choices, with what they cost.

**The page reads the label set from the host and does not generate it.** The alternative
is a script, as `.xeno/plugin/bin/xeno-labels.sh` is for the one label Xeno asks for, or a
test holding the page against `gh api .../labels` as `supply_chain_test.go` holds the
supply chain page against the tree. *Consequence of not doing it:* the page is written by
hand and can go stale, and nothing will say so. *Cost of doing it:* a test in a Go
repository that calls a code host, which is the one thing section 12 keeps out of the gate
path, or a script producing a page, which makes the page a build product. The plan's own
answer for the `wp` labels is the precedent — "a one time act with no script behind it,
because a tool that parses the plan to automate something that happens once is more
apparatus than it returns" — and 34 labels that change a few times a year are the same
case. What the page does instead is say when it was read and from where, so a reader can
repeat the one command.

**The page names `xeno-needs-decision`'s spelling discrepancy rather than resolving it.**
The alternative is to rename the label to the maintainer's singular, which is five issues
re-associated on the host and any saved query naming the old string failing silently, or
to say nothing. *Consequence:* the page carries one sentence of friction that a later
decision will remove. *Cost of the alternatives:* a rename on nobody's instruction is the
agent deciding the vocabulary, and saying nothing leaves a reader to find the mismatch
between the maintainer's comment and the label themselves, which is the kind of thing this
page exists to stop happening.

<!-- xeno:section:impact -->
## Impact

Three files change and no code does.

**`docs/labels.md`**, new. Three tables and the prose around them, on the order of
18,000 bytes. It is the whole of the deliverable.

**`docs/README.md`**, one entry under "Running it", beside `commands.md` and
`symbol-index.md`. That section is where a reader looking for how the thing is operated
goes, and a label convention is operation rather than specification. The entry says what
the page is and, in the same sentence, what it is not: a record of what is in use, not a
definition of what a tool reads.

**`zensical.toml`**, one line in `nav`, in the "Working with it" list after "Commands".
The nav is hand-written, so this is the only thing that makes the page reachable from the
published site; the index entry in `docs/README.md` makes it reachable from the
repository. Both are needed and they fail differently: a missing nav entry publishes an
orphan page nobody can navigate to, and a missing index entry leaves a reader of the
repository not knowing it exists. Only the second is caught by the strict build, and only
indirectly — the build fails on a link to a page that does not exist, not on a page that
nothing links to.

**What does not change, and why each is worth saying.**

`docs/process-definition.md` and `docs/implementation-plan.md` are untouched. The page
cites section 12 for what approval is and the plan for what a work package label is, and
restates neither. This is the constraint the whole intent is shaped by, and the empty diff
against both is an acceptance criterion rather than an observation.

`.xeno/plugin/bin/xeno-labels.sh` is untouched, and its comment's claim to create "the
one label Xeno asks a project's tracker for" stays true. The question the task set was
whether the second label belongs in that script. It does not, and the reason is what the
script is for: it creates the label that is a precondition of `xeno intent start` in any
repository that adopts Xeno, which is why `xeno init` names it among the settings it does
not make. `xeno-needs-decision` is read by nothing, is a habit of this tracker, and an
adopter who never puts a question to a maintainer never needs it. Adding it would widen
the script's promise from "what Xeno requires" to "what this project happens to use",
and the next adopter would get a label that means nothing to their process. Had the label
been something the runner read, the script would have been the right place and the change
would have been a section 12 amendment first.

`internal/model/identity.go` is untouched, and so is every other Go file. No gate, no
rule, no test is added: nothing mechanically checks that an issue carries the label its
state implies, which the page says in as many words so that its tables are not read as
enforcement.

**What this makes possible and what it does not.** After this, a reader has one place that
says what each label means, who acts on it and what reads it, which is what #359 asked
for. What it does not give is any guarantee that the tracker and the page agree tomorrow:
the page records the set as read on a date and names the command that reads it, and the
only thing keeping it current is somebody noticing. That is the same arrangement the plan
chose for the `wp` labels and it is stated on the page rather than left as a surprise.

**The one thing a later change will have to touch.** When the lifecycle question on #359
is answered, section 12 gains a sentence by the maintainer's commit, and this page's row
for `xeno-approved` changes in the "who clears it" column and loses one of its four open
points. The page is built so that this is one cell and one sentence rather than a
rewrite, which is the reason the rows are narrow and the reasoning is in the prose above
them.
