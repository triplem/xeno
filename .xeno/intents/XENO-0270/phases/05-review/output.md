---
intent: github.com/triplem/xeno#231
phase: 05-review
created: "2026-10-07T09:50:33Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ec88b96575e78b36934a55c398da078fc1a245e0478084e7a97b0f1f52b5ce0b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'The rule applies to 03-implementation and that phase recorded four deviations, all naming what they depart from. Two are corrections to P2''s decisions, named: the reference page does carry per command prose, against P2''s ''no per command prose, no examples'', because five paragraphs of the old README had no other home and criterion 10 forbids losing them; and the no-network-call paragraph is kept rather than dropped, because section 12 carries its first half and not the sentence about xeno enforcement check. The third names P2''s requirement that the moved material move verbatim, and the two exceptions the move forces, a heading that would repeat the page title and two paths written from the repository root. The fourth names no decision because it departs from none: the old learning record paragraph was two paragraphs with no blank line, so an extraction by paragraph boundary carried one into the wrong file, and it was split. Marked not-applicable rather than met because the rule is scoped to the implementation phase and this is the review answering about it.'
      result: not-applicable
      rule: deviations-are-traceable
    - note: 'No interface changes. No artifact gains or loses a field, no gate gains or relaxes a check, no command changes its behaviour, its flags or its refusals, and no exit code moves: usage is read by the new test and not altered, and xeno gate verify reports 499 verdicts verified. The only Go change is one test function. What an adopter meets is a shorter README and three new files under docs/, so there is nothing to migrate. A reader who had bookmarked the README''s command list finds it at docs/commands.md with six more commands than it used to have.'
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: None added. go.mod and go.sum are untouched and nothing compiles differently. The new test uses os, regexp and strings, all already imported by the file it sits in.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** No normative document is touched: the purpose and the name are
pointed at and restated in the README's own words, and the implementation plan's refusal to
record the name stays. Nothing is invented — every fact in the four files was already in the
README or in a normative document, which is what the nothing-lost check also demonstrates in
the other direction. The change belongs to WP16 by subject and to intent XENO-0270 for issue
#231.

**What the issue asked for, and what the maintainer corrected.** #231's done-when says the
README opens with what the tool is for, says where the name comes from, links the documents
with the process definition named normative, and either regenerates the command list or cuts
it down. All four are done. The maintainer's comment changed one thing: the issue says "the
existing material stays ... further down" and lists the no-network-call property among it,
and the comment said that is deep information and should move. It moved, and the half of it
that section 12 does not carry moved with it rather than being dropped, which P3 records as a
deviation from P2's reading.

**What a reviewer should look at first.** The per command notes at the bottom of
`docs/commands.md`. They are prose about behaviour sitting next to a block a test holds, and
nothing holds them. This intent's whole argument against the old README was that a
hand-maintained copy drifts; these notes are a hand-maintained description that can go stale in
exactly the same way, and the only defence is that they are now beside the authoritative text
rather than instead of it. P4's gaps say so. The alternative was dropping five paragraphs of
real information, which criterion 10 forbids, so the choice was between an unheld note and a
loss.

**The second thing: the coverage table moved without being re-read.** Twelve rows naming tests
against the plan's acceptance criteria. P1 ruled re-measuring out of scope with a reason — a
table that moves and changes in one commit is unreadable both ways — and the consequence is a
page whose most authoritative-looking element has an unknown age.

**The question no rule asks: is four commands the right number?** `init`, `intent start`,
`phase start`, `phase finish`. It leaves out `section set`, which the sequence cannot be run
without, and the paragraph underneath names it in prose instead. A reader who copies the four
lines gets a phase that will not finish. The alternative was five lines, which stops being a
block somebody reads and starts being a list; the prose sentence is the compromise and it is
the place this could be wrong.

**And: does the index earn its place before WP16 exists?** It is a file that duplicates what
`ls docs/` shows, plus a sentence each. What makes it worth having is the grouping and the
normative marking, neither of which a directory listing carries, and the maintainer asked for
it by name. If WP16 generates navigation from front matter later, this becomes the page it
generates from or the page it replaces, and either is cheap.

**Two wrong claims were written and corrected inside the intent**, both about documents this
one links: which commands may use the network, and what the two v2 drafts are. Both were
caught by opening the file. P4's gaps name that as a method nothing enforces, and P3's learning
record proposes the mechanical half of it for moved prose.

<!-- xeno:section:release-notes -->
## Release notes

`README.md` now opens with what Xeno is for. Requirements, design, code and evidence are
produced separately, at different times, partly by people and partly by models, and the only
thing holding them together is their binding to an intent; Xeno makes that binding provable by
running one change through six phases, writing an artifact per phase into the repository the
change lives in, and judging each with gates that make no network call. What it produces is
evidence, and the ISO/IEC 42001 sentence is there rather than 2284 lines into a document.

The name has a paragraph. It is a recording technique: two tracks that were never played
together, and never at the same tempo, laid over one another afterwards, where the coherence is
in the assignment rather than in the recording. Appendix C of the process definition says it
properly, and until now nothing a newcomer read said it at all.

`docs/` has an index. [docs/README.md](docs/README.md) carries every document with one line on
what it is for, grouped by the errand a reader arrives with, and marks the process definition
and the implementation plan as the normative two in their own entries.

The command list is no longer maintained by hand in two places. The README names four commands
in the order somebody runs them and points at `xeno --help` and
[docs/commands.md](docs/commands.md), whose block is the `usage` constant itself with a test
that fails if the two part. The old list had drifted by six commands — `decision record`,
`evidence declare`, `intent verify`, `question record`, `review answer` and `scope set` were
missing — and the page carries all twenty-seven.

What the README used to say about this repository's own trail is now
[docs/the-trail-in-this-repository.md](docs/the-trail-in-this-repository.md): what
`runner_version`, `plugin_version` and `tool_version` say here and why, the two shapes of the
intents directory, the coverage against the plan's acceptance criteria, and the two deliberate
mutations. The per command notes that were in the README are at the bottom of the reference.

Nothing is gone. Fourteen of the old README's nineteen paragraphs are in the new files word for
word, and the other five are the replaced opening, the superseded list, one paragraph that was
two, one heading the move renamed and one path the move rewrote.

<!-- xeno:section:residual-risk -->
## Residual risk

**The per command notes are unheld prose about behaviour.** Five paragraphs at the bottom of
`docs/commands.md` describing what `--for` accepts, what `learning record` writes, what
`intent status` lists, what `enforcement check` does on the network, and when the next step is
printed. If any of those changes, the note is wrong and the suite is green. It is the same
defect the intent was filed about, one level down and smaller: the list beside them is now
safe, and they are not. The alternative was losing five paragraphs of real information.

**Three of the four files are held by nothing.** Only the reference's block has a test. A file
added under `docs/` tomorrow will not appear in the index and nothing will say so; a link that
rots will rot quietly. P1 ruled a link checker out of scope as its own change with its own
argument, and this is the gap that argues for it.

**The coverage table's age is unknown.** It moved verbatim, as P1 required, and nobody checked
that the twelve rows still name tests that exist and still assert what each row claims. It is
the most authoritative-looking thing on the new page and the least recently verified.

**Four command lines will not run a phase.** `init`, `intent start`, `phase start`,
`phase finish` — and a phase does not finish without `section set`, which is named in the
paragraph underneath rather than in the block. Somebody who copies the block and stops reading
gets a refusal. The block was kept to four because five stops being read; the prose is the
compromise and it is where this is likeliest to annoy somebody.

**WP16 may inherit the guard as the design.** The plan says the reference is generated from the
declarations. What is here is a test asserting that a committed page matches a hand written
constant, which is a stopgap. P2's alternatives say so in as many words, and nothing stops a
later reader treating the test as the answer and leaving `usage` hand maintained.

**`docs/README.md` is the wrong name for a generator.** Chosen for GitHub, which renders it
when somebody opens the directory. A static site wants `index.md`, so WP16 has a rename to
make, decided here without a generator to test it against.

**Two wrong claims were written inside this intent and caught by reading.** Which commands may
use the network, and what the two v2 drafts are. Both were about documents this change links,
both were wrong in the confident direction, and what found them was opening the file. The next
such claim has the same defence and nothing more.

**For a person, not for the code.** Whether the first screen does its job. A stranger deciding
whether Xeno applies to them is the reader the whole change is for, and nobody in that position
has read it.
