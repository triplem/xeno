---
intent: github.com/triplem/xeno#334
phase: 03-implementation
created: "2026-10-10T15:49:39Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 73ff8fc7618ab88c655d7e3d6804caf9a79c311899c5e45c886ce8206efb56bb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

One file, four hunks, 334 lines added and 9 removed. `docs/orchestrator-evaluation.md`
grows from 28,602 bytes to 51,875.

**Frontmatter.** `revision: 3` to `revision: 4`, `date: 2026-10-08` to `date: 2026-10-10`.

**Section 1, the paragraph that indexes the page's later sections, replaced whole.** It now
names sections 9 and 10 together as two further decisions of the same kind about different
kinds of tool, keeps the sentence about the form and about a declined tool being useful only
to a reader who can find the reason, and adds the one thing that distinguishes the new
section: the orchestrator was chosen against a platform's criteria and the framework against
one section of one template, while a methodology compares with the process itself. The last
sentence, that the page is the record of external tools this project evaluated, is kept word
for word, because it was already true of three.

**Section 10, new, between 9.6 and Sources.** 293 lines of the 334 the diff adds. Six subsections
and one five-row table.

10.1 pins both repositories. The tag is annotated, so the entry says that the ref resolves
to tag object `4079edbe` and that the reading is at the commit it points at,
`6a378b53c0a4fe0641ed7d8de8dfff94264d5b6a`; the figures that move are dated to the day they
were read. It names the shape of the thing — five phases, 33 stages, eleven scope profiles,
fourteen agents, one `core/` projected onto seven harnesses — and ends with the hosted
sample as a line, pinned at `v2.2.0`, commit `bc988d0e`, held to what its own README says.

10.2 is the table and six paragraphs under it, one per row plus one that says the first four
are 9.2's and that the fifth is the addition. The third row is settled against the code and
the two contradicting sentences are quoted with their addresses. The enforcement account is
three qualifications deep — gate only, opt-in only, overridable — and the override's three
refusals are named, because an enforcement claim without its refusal path is the flattering
half. The fifth row says this project is behind, and says where it is ahead: a merge request
against the rule set is the stronger review and a tick at a gate is the one that happens.

10.3 decides. Layer 2 inside the harness, which section 2 and section 4.4 already declined,
and the fourth row deciding what the second leaves open, as 9.3 had it. A paragraph says
what is deliberately not part of the reason, with the Sigstore bundle, the two honest "not
yet" notes from AI-DLC's own plugin contract, and the fifth row as the evidence.

10.4 holds #323's three shapes with what each costs. The plugin shape fails on five specific
clauses of the plugin contract rather than on a general objection. The bridge shape is named
as the cheap route and explicitly left as something nobody has to decide, since nothing
stops it today.

10.5 is five conditions. The first is the alternative the maintainer did not take, with
`docs/process-definition.md:480` named as what it would change and the case that would bring
it back. The last is the ideation and operation question, pointed at WP21 and #337 rather
than answered.

10.6 names four mechanisms. Three point at #340, at the gap in this project's own learning
route, and at the LLM-in-the-refusal-path check, which is named and declined on section 7's
opening sentence. The fourth describes the reviewed-source fingerprint in full and records
the gap, with the plain sentence the question on #334 produced: Xeno seals the record of a
review; it does not seal the identity of the source that was reviewed.

**Sources, one new entry.** Both repositories, both annotated tags with both shas and which
is which, the date, and each file named by what it was read for. It also records the tree
search for tracker vocabulary together with its control search, because that is what the
first row's negative claim rests on, and it records that the sample's own five rows have not
been read by anybody.

Nothing else in the tree changes. No Go file, no test, no other document.

<!-- xeno:section:deviations -->
## Deviations from the design

Two, and the second is the one worth reading.

**The section is 23,273 bytes rather than the 16,000 the design estimated.** The design's
impact section put the growth "of the order of 16,000 bytes", landing the file near 54,000;
it is 51,875, so the file figure was close and the growth figure was not. The reason is the
third row: settling the contradiction against the code took the two quoted sentences, the
two functions with their line numbers, the hook's comment, the six manifests and the three
refusals of the override, which is five paragraphs where 9.2's third row needed one cell and
a clause. The budget was set at 140,000 bytes against a scope of 37,939 and absorbs it
without a finding, which is what the generous figure was for.

**The reason at `docs/process-definition.md:480` is cited twice, by two different means, and
that is deliberate rather than an oversight.** 10.5's first condition names the file and the
line, because what it says is that a change to that line comes first and a reader has to be
able to find it. 10.6 quotes the sentence and attributes it to the process definition as the
normative document without the line number, because that is how this page cites its own
documents everywhere else — 9.6 says "section 6 of [the process definition]" — and because
what 10.6 needs from the sentence is its authority and not its address. The design said
AI-DLC is quoted with file and line and said nothing about this project's own documents, so
this is a gap in the design filled rather than a departure from it, and it is written down
because the two citations look inconsistent to anybody who meets them in the other order.

Everything else is the design. Four hunks in the one file named in the impact section, in
the order it named them. The section-1 paragraph replaced whole and read back as a paragraph
rather than as a diff. No subsection for the hosted sample, which is a line in 10.1 and a
clause in the second row. No mechanism taken. `git diff main` touches no other file.
