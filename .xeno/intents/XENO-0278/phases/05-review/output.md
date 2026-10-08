---
intent: github.com/triplem/xeno#330
phase: 05-review
created: "2026-10-08T18:04:48Z"
schema_version: "1.0"
runner_version: dev+8fb365d.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 567dda5c8ca6e6d466def46e810e0b5be50cc8815b27409d0c8c2534015e0bf1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: met
      note: 'Two deviations in P3, each naming the design decision it departs from: the wording of the read-failure refusal, and the third shape of a read that did not happen, which the design''s issueOf did not mention and a test caught.'
    - rule: interface-change-needs-a-migration-note
      result: deviation
      note: 'xeno intent start with a tracker block now needs the token the block names and an approved issue; a project that started intents without a token is refused on its first start after upgrading. The move is one sentence, in docs/commands.md and in the refusal: set the variable, label and comment the issue. Runner.IntentStart''s third argument is internal, since the runner is not a published library.'
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: 'No dependency: two more HTTP calls through the client the adapters had, and nothing imported that was not.'
    - result: deviation
      note: 'Security lens: an approval is a comment anybody may write, by the decision on the issue, with the label''s right as the standing check; the intake quotes the author so a reader sees who, and the alternative of judging the author is named in the design and the residual risk rather than resolved.'
      source: lens
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**deviations-are-traceable — met.** P3 records two deviations and each names the design
decision it departs from: the literal wording of the read-failure refusal, and the three
answers of `issueOf`, which the design described without the adapter's third shape. Both
are wording and shape; nothing decided was undone.

**interface-change-needs-a-migration-note — deviation.** Two interfaces move and one of
them is outside this intent. `Runner.IntentStart` gains a third argument, which every
caller in the tree was given and which no adopter calls, since the runner is not a
published library; that is internal. `xeno intent start` is not: with a tracker block it
now needs a token and an approved issue, and a project that started intents without a
token will be refused on the first start after upgrading. The migration is one sentence,
in `docs/commands.md` and in the refusal itself — set the variable the block names, and
label and comment the issue — and the result is `deviation` rather than `met` because the
note is in the command's documentation and not in this artifact's own text until this
entry, which is where the rule wants it.

**new-dependency-needs-a-rationale — not-applicable.** No dependency. The adapters make
two more HTTP calls through the client they had, and nothing is imported that was not.

**What no rule asks.** Whether the intake's sentence, written from a second read at P0,
is the record the clause meant: A103 says why it is written that way and when it could
disagree with the start, and the residual risk below carries the case.

<!-- xeno:section:release-notes -->
## Release notes

**An intent starts only from an approved issue.** With a tracker block configured, `xeno
intent start` reads the issue and refuses one nobody approved. Approval is two things on
the issue: the label `approved`, and a comment whose first line is the word `approved`,
the rest of which is the reason. The refusal names which of the two is missing. The label
needs a right the host grants; the comment carries why.

**A milestone holds work until its turn.** An issue in a milestone is refused while an
earlier open milestone exists — by due date, undated ones last — and `--now` starts it
anyway. A closed milestone, or the earliest open one, holds nothing.

**The intake says by what it was authorised.** `phase start` at P0 writes, above the
quoted issue, who approved it, when and why, and that the intent was started ahead of its
milestone where it was. Where it finds no approval at that moment, it says so.

**What this costs a project.** `intent start` now needs the token the tracker block names
and makes up to three calls — the issue, its comments, and the project's open milestones
where the issue carries one — on GitHub and on GitLab alike; `phase start` at P0 makes the
same calls. A project without a tracker block sees nothing of this. A read that cannot
happen — no token, no such issue, an issue on a host the block does not reach — refuses
the start rather than starting on nothing.

Section 12 of the process definition carries the clause, *Starting an intent*, committed
before the code. The reading `xeno gate` makes is unchanged: none of this is a gate, and
no gate reaches the host.

<!-- xeno:section:residual-risk -->
## Residual risk

In order of weight.

**A comment by anybody counts, by decision.** The word is the marker and the label is the
standing check; a stranger's `approved` comment on a labelled issue is an approval with a
stranger's name on it in the intake. The host's rights on the label are what the design
leans on, and the intake quotes the author so that a reader sees who. The remedy, if it
is ever needed, is judging the author, which GitHub can answer in one field and GitLab
cannot — the alternative the design declined and named.

**The sentence is written from a second read.** A label removed or a milestone moved
between `intent start` and `phase start` leaves an intake that says what P0 found, not
what the start found. A103 records why; the risk is a reader who takes the sentence as a
record of the start. The sentence says *when this phase started*, which is the one
defence.

**The positive path has not met a real approved issue.** Every refusal was run against
the host; the start that writes was proved on fakes that serve what the hosts document.
The first `approved` comment the maintainer writes is the first evidence, and this
issue's own answer predates the word, so this intent's intake was authorised by hand.

**Milestone ordering is a text comparison.** `due_on` and `due_date` sort as text within
one host for the formats both document; a self managed deployment that writes dates
otherwise would order them wrongly and hold or release an issue for the wrong reason.
The refusal names both milestones, so a wrong order is visible rather than silent.

**The clause fixes the word in English.** A project working in German labels its issues
`approved` and writes `approved` all the same; that is the specification's choice and
the one place this change is not localisable without a specification change.

**Belongs to the specification, not to this code.** Whether a closed milestone should
hold nothing, as implemented, or whether an issue in a closed milestone is itself a
mistake worth refusing, the clause does not say; and the sibling finding, comments in the
intake, is read now and not quoted, which is XENO-0277's and this intent's P0 learning
and nobody's issue yet.
