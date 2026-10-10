---
intent: github.com/triplem/xeno#334
phase: 00-intake
created: "2026-10-10T15:33:04Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 098667565f56ca8a3f905fe75d55831918384adb2f4cf4827ec95dae149bbcfa
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: 'A Xeno verdict does not bind the source it judged by identity: the section records the gap and takes nothing'
      rationale: 'Put to the maintainer on #334 on 2026-10-10T13:41:17Z with three options and a recommendation, and answered at 14:59:46Z with "Use the recommendation". AI-DLC binds a review to the bytes it approved: a per-unit REVIEW_COMPLETED receipt carries a Unit Source Fingerprint over a committed path-to-blob-OID listing, and aidlc attest resolve --diff base..head then classifies every changed path as verified, drifted, unattested, unverifiable or indeterminate from (base, head) alone in any clone. Nothing in this project binds a review result to git object identity; the five hashes are all over the record own content and G-Freshness binds the inputs a phase read, by freshness, rather than the source a review approved, by identity. The reason that decides is docs/process-definition.md:480, "A verdict says what it judged, by content and not by commit", which is the normative document and therefore a decision already taken rather than an oversight. Taking the mechanism would be a change to that line in the maintainer own commit before any code, and 9.6 already set the discipline that a mechanism is named in the evaluation and taken by an issue of its own. The two alternatives were taking it in full and taking the weaker form where a verdict records the base and head it was computed over; the third is written into the conditions for revisiting rather than into the tree, which is where it belongs until a case arrives. What this costs is that the gap stays open, and the honest answer to what a green P5 proves about the code that was merged remains that the record was sealed, not that the diff was the diff.'
      decided_by: triplem
      proposed_by: claude-opus-5
---

# Intake

<!-- xeno:section:problem -->
## Problem

Approved by @triplem on 2026-10-10T14:59:46Z: Use the recommendation

> **github.com/triplem/xeno#334** — The evaluation has no section for AI-DLC, which compares with the process itself
>
> Asked for on #323, where the reading was done on 2026-10-08 and the comparison is laid
> out; this issue is the work of writing it down.
>
> ## What is to be written
>
> A section of `docs/orchestrator-evaluation.md` in the form section 9 has for OpenSpec:
> what was evaluated, pinned at a tag; the four rows of 9.2 and a fifth that AI-DLC
> needs and OpenSpec did not, rules and learning; the decision and the reason; the
> conditions for revisiting; and the one or more mechanisms worth taking. The
> collaborative sample, `aws-samples/sample-collaborative-ai-dlc`, is a line in it, since
> it is the same methodology hosted and compares the same way.
>
> The reading on #323 was done at `main` with the two preview tags and `v2.11.0` cut
> within three days of it. The section has to be written from one pinned tag, with the
> table re-read at that tag, and the one point where the two documents disagree settled
> there: the rules guide says sensors "are advisory in this release", the sensor
> reference says a gate-fired sensor with `default_severity: blocking` stops stage entry.
> Which is current at the pinned tag decides the third row.
>
> ## What the decision is
>
> Not an extension. #323 weighs the three shapes, a plugin, a way of working a phase,
> and taking the mechanisms, and the second costs a page of an adopter's documentation
> and nothing in this repository. The section records that reasoning so that the next
> reader meets a decision and not an open comparison, and so that the conditions for
> revisiting are written where the ones for OpenSpec are.
>
> ## Done when
>
> The section is in the document at a pinned tag, the Sources list carries what it read,
> and the evaluation's table of contents names it beside section 9. Nothing in the
> process definition changes for it.
>
> Refs #323
>

The issue as it stood at 2026-10-10T15:33:04Z, read by `xeno phase start` and quoted rather than summarised. What this phase concludes about it belongs below.

<!-- xeno:section:scope -->
## Scope

This intent writes one new section of `docs/orchestrator-evaluation.md` and nothing else.
It is **section 10**, in the form section 9 has for OpenSpec, for the AI-DLC methodology
pinned at `awslabs/aidlc-workflows` `v2.11.0`. It carries six subsections, which are the
things #334 asks for: what was evaluated and at which sha, the four rows of 9.2 plus the
fifth that AI-DLC needs, the decision and the reason that decides it, the three shapes
#323 weighed with what each costs, the conditions for revisiting, and the mechanisms
worth taking. `aws-samples/sample-collaborative-ai-dlc` is a line inside it, pinned at a
tag of its own, because it is the same methodology hosted and compares the same way.

Three smaller things change in the same file, and each follows from the section existing
rather than being chosen beside it.

**The paragraph in section 1 that names section 9 is replaced.** It currently reads that
section 9 records a second decision of the same kind and that the page is the record of
external tools this project evaluated. With a third such decision on the page the
sentence has to say so, and the paragraph is being changed for the second time, so it is
replaced whole rather than edited into, which is what `CLAUDE.md` asks for.

**The Sources list gains an entry** naming what was read, at which sha, and what each
file was read for. The existing entries for section 9 and for 4.1 are the form.

**The frontmatter's `revision` goes to 4 and `date` to 2026-10-10.** A whole new section
is what bumped it from 2 to 3 in `a010cd1`; the one-sentence change in `5044a7a` left
both alone, which is the line this follows.

What this intent does not do.

**Nothing in `docs/process-definition.md` or `docs/implementation-plan.md` changes.**
#334 says so in as many words, and the first standing rule would make either a person's
commit before any of this. The section quotes both and edits neither.

**It takes no mechanism.** #334 settles that the decision is not an extension, and the
question put to the maintainer on the issue settles that the one mechanism worth a
decision of its own — a verdict bound to the source it judged by identity — is recorded
as a gap and nothing is taken. The three mechanisms #323 named each already have an
issue (#338, #339, #340); the section names them and points at the issues rather than
re-deciding them here. A section that recorded a mechanism as taken would be a
specification change wearing an evaluation's clothes.

**It opens no new question.** The one open question on #334 was answered by the
maintainer at 2026-10-10T14:59:46Z, before this intent started, and it is recorded in the
decisions of this phase rather than carried as an assumption.

**It does not re-run the market exploration.** #323 keeps that deliberately, for the
reason it gives: the methodology side moves at a tag a day and a document read at `main`
is stale by the time it is reviewed. This section is one pinned reading, and the pin is
the point.

**Which package.** #334 carries no work-package label, and neither did its sibling #336,
whose intent shipped as `5044a7a` without one. The third standing rule asks that a change
belong to a package and that a gap be written down rather than absorbed: the documents of
this repository are not a deliverable of any package in the plan, which is a property of
the plan and not of this issue. It is recorded as a learning of this phase rather than
settled here, and #334 is left unlabelled rather than labelled wrongly.

<!-- xeno:section:context-rationale -->
## Why this context

Three files, 37,939 bytes as the scope was set. The budget is 8 files and 140,000 bytes,
which is above both figures on purpose: the work of this intent is to make one of those
three files substantially longer, and the budget is judged against the lock of each phase
from P1 on, so a figure measured before the writing would be a figure the writing itself
breaks. Section 9 is about 180 lines and 14,000 bytes, this section is of that order, and
the room left over is deliberate rather than an estimate.

**What changes.** `docs/orchestrator-evaluation.md`, in the four places the scope section
names.

**What says whether the change is right.** `CLAUDE.md`, for the three standing rules this
section is constrained by and for the three conventions it is easiest to break here: that
a heading names its section in words rather than borrowing a number, that a paragraph
changed for the second time is replaced and read back as a paragraph, and that a negative
result is evidence only when the thing checked was there to be found. The last of those
decides how the section may state what AI-DLC does not have.

**What tells the reader where the page sits.** `docs/README.md`, whose entry for this file
describes it as the record of which agent platform v2 delegates a phase to. That
description was already narrower than the page after section 9 landed and is narrower
again after this one. Read so that the question is answered rather than left; #334 asks
for the evaluation's own index entry and not for this one, so the line is read and left
alone, and whether the page's description should widen is a thing for whoever next
changes that file.

Two files are deliberately out, and both are quoted here instead so that no later phase
needs them.

**`docs/process-definition.md`, 132 KB, excluded.** The section leans on four things in
it, and all four are short enough to carry.

Line 480, which the maintainer's question turns on: "**A verdict says what it judged, by
content and not by commit.**" That sentence appears exactly once across `docs/*.md`, and
in the normative document, which is what makes it a decision rather than an observation.

Section 7's gate table, for the count and the character of the gates: fourteen gates,
"All gates are deterministic and model free", and `G-Freshness` checks "each phase's
context hash points at the current state of its predecessor, and no file a preceding phase
read was changed by the change under review afterwards" — the inputs a phase read, by
freshness, and not the source a review approved, by identity.

Section 10, for the fifth row: "Learning never changes behaviour directly. It produces a
merge request against the rule set. Only after merge does it take effect", and "never an
edit made where it was noticed." Against AI-DLC's tick at a gate that is what the fifth row
compares.

Section 6's `Override` and the first standing rule, both of which section 9.6 already
quotes on this same page, so the page carries them for its own reader.

**`docs/implementation-plan.md`, 130 KB, excluded.** Two packages are named and each is one
line. WP18 is the across-intents dashboard, deferred to 1.1, which 9.6 already cites on this
page. WP21 is ideation, "The stage before there is a request", deferred to 1.1 and described
there as "A sealed record beside the intents, and not a phase" — which is the one place
AI-DLC's lifecycle reaches past this project's six phases and the plan already has an answer
for.

`.xeno/intents/**` is out because the trail is the record and not input.
