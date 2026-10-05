---
intent: github.com/triplem/xeno#243
phase: 00-intake
created: "2026-10-05T12:34:31Z"
schema_version: "1.0"
runner_version: dev+f2f65d5
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c695bebe6a06cb1860e7a4c9e9183ba8562f98ef8d223c918a33bbef6a8bb728
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

`docs/assumptions.md` opens by saying the record is closed, that A77 is the last row, and
that nothing is added here. The file ends at A94.

A78 through A94 were added after the banner arrived in `f001058` (#174): seventeen rows
across twelve commits, including every documentation change of the last three days. Each of
those intents read the banner, added a row anyway, and left the banner saying otherwise.

The file that documents this project's unread rules opens with one. That is A90's own
finding — a rule stated in a document with nothing comparing the repository against it —
except here the document contradicts itself and twelve commits say which half is false.

The cost is not tidiness. A reader who believes the banner puts their row somewhere else or
writes none; a reader who disbelieves it has no way to tell which other sentence in the file
is also stale, and the file's remaining paragraphs are load bearing: they explain the state
column, why a superseded row is never deleted, and what `accepted until` means.

Underneath the contradiction is a misreading, which is what makes this a correction rather
than a reversal. The plan's section 4 says the record kept while Xeno could not govern
itself is "a file in the branch rather than an artifact under `.xeno/`", and that from M0
"the same loop runs through the runner, and the hand held record stops". What stops is the
substitute for a phase's artifacts. That did stop: every change since M0 runs six phases and
writes real ones. The register is not that substitute. It is a cross-intent record of
construction choices, and nothing in either normative document closes it.

So #174 applied a sentence about artifacts to a file that is not one, and seventeen rows of
practice have been quietly right ever since.

<!-- xeno:section:scope -->
## Scope

In scope is the opening of `docs/assumptions.md`: the two paragraphs that close the record,
and the sentence of the fourth that frames the whole file as pre-M0. They are replaced rather
than edited into, and read back as prose, because each is being changed for the second time.

What the new opening says, decided by the maintainer on the evidence above: the register is
open, for assumptions and decisions about this repository's construction. What closed at M0
is the hand held record the plan's section 4 names, the record kept as a file in the branch
instead of an artifact under `.xeno/`, and that is stated so the next reader does not make
#174's inference again.

In scope is keeping the banner's correct half. A learning never comes here: it belongs to the
phase's `learning.yaml` and from there to a merge request against the rule set, as section 10
describes. Practice has observed that without exception and the file should keep saying it.

In scope is saying what distinguishes a row here from a decision in an intent's P2, because
that is the question the banner was reaching for and answered by closing the file. A decision
taken inside one intent and sealed with it belongs to that intent. A row here is for what
outlives the intent that found it.

In scope is one row recording this correction, which is A95 and the first row written under
the reopened register rather than against its banner.

Out of scope is deleting or rewriting any existing row. The file's own third paragraph says
every row stays with its state column, and this intent does not reopen that.

Out of scope is the seventeen rows' content. Whether A78 through A94 were each the right call
is seventeen questions; this intent settles only where such a row belongs.

Out of scope is any change to either normative document. The correction is that #174 misread
the plan, not that the plan is wrong, so nothing in `docs/implementation-plan.md` needs
touching and no specification commit precedes this.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the register's opening, the plan sentence it leans on, and the three files that
tell a reader where a decision goes.

`docs/implementation-plan.md` is read for section 4 rather than quoted from the banner, and
that is what turned this from a reversal into a correction. The banner's claim is checkable
and it fails: the plan's sentence is about the record being "a file in the branch rather than
an artifact under `.xeno/`", which is the substitute for phase artifacts and not this file.
Reading it also settles that no specification commit precedes this intent, which the previous
documentation intent had to discover the hard way.

`docs/process-definition.md` is read for section 10, because the new opening keeps the claim
that a learning routes through `learning.yaml` and a merge request against the rule set. A
paragraph that points at another document for its rule is only as good as the pointer, and
this one holds.

`CLAUDE.md` is in scope because it names this file as where assumptions and decisions taken
while building the core are written down, and because it is sent with every request of every
session. If the register's opening and `CLAUDE.md` disagree about whether the file is open,
the one that is read before every change wins in practice, which is how a banner survives
twelve contradictions.

`CONTRIBUTING.md` and `README.md` are read for the same reason and found not to need changing:
`CONTRIBUTING.md` cites A6 and `README.md` cites the file and A20, and all three are citations
of rows rather than statements about whether rows may be added.

The twelve commits are evidence rather than context, and they come from `git log -S` over each
row id. They are in the issue and in P2's alternatives, counted rather than estimated, because
"the banner was ignored" and "the banner was ignored twelve times by twelve intents" are
different claims and only the second one decides anything.
