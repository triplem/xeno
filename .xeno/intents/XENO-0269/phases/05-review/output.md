---
intent: github.com/triplem/xeno#207
phase: 05-review
created: "2026-10-07T09:12:03Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b979a1fbf6e1060ddd6e4986ede04c02c8a1f2217121525a543b4daaeb8c00e4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'The rule applies to 03-implementation, and that phase recorded three deviations. The first names P1''s criterion 2: the approved wording said model comes from project.yaml or the harness, agent() reads it from project.yaml alone, and the criterion asks that each of the three be described as the code has it, which the paragraph as it stands meets. The second, dropping two em dashes, names nothing because it departs from nothing: the meaning is unchanged and punctuation was never what was approved. The third, the row being A100 rather than A98, names criterion 10, which asked for one row without saying which number. Marked not-applicable rather than met because the rule is scoped to the implementation phase and this is the review answering about it; all three are traceable and none is a bare note.'
      result: not-applicable
      rule: deviations-are-traceable
    - note: 'No interface changes. The diff against main is one document and no Go file, and xeno gate verify reports 481 verdicts verified. No artifact gains or loses a field, no gate gains or relaxes a check, no command changes its behaviour or its refusals, and a green G-Schema means exactly what it meant before. What an adopter meets is one paragraph of specification saying what three existing fields are worth. There is nothing to migrate, and the one thing worth knowing is in the release notes rather than a migration note: a provider register built on the triple is reading declarations.'
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: None added. go.mod and go.sum are untouched and nothing compiles differently. The diff is one paragraph in docs/process-definition.md, one count and two paragraphs in docs/clause-readers.md, and one row in docs/assumptions.md.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** A normative document is changed, which the first standing rule
bars the agent from doing. The exception is the one the maintainer made on instruction: three
options with their consequences and costs, the choice theirs, the wording drafted and put to
them as a separate question, and "write it as drafted" acted on. Both were put one at a time.
The specification change is its own commit and comes first. Nothing is invented: no field, no
gate, no rule. The change belongs to WP11 and to intent XENO-0269 for issue #207.

**What #207 asked for, and what it got.** Either something corroborates at least one of the
three, or the documents say the triple is self-reported and the register is read in that light.
The first was examined and is unavailable inside the boundary this project drew, for three
reasons now written in the document rather than only in the issue. The second is written, in
section 12, in the issue's own closing words. The issue closes.

**What a reviewer should look at first, and it is not the paragraph.** `hashes` reads
`tool == "manual"` and accepts the `by-hand` hash placeholder in every field on the strength of
it. That is a check relaxed on a self-reported value, which is a sharper form of #207's concern
than the provider register is, and the paragraph does not mention it. It was found while
verifying criterion 3 and is deliberately not written into the specification: it is outside
#207's "done when" and outside the wording that was approved, and writing it in would have been
the agent extending an approved sentence. P4's results and gaps carry it, the learning record
carries the general form, and it wants an issue of its own.

**The one thing in the paragraph that could be read as more than it is.** "The one record from
another hand is the gateway's." It is a record of routing, as the sentence after it says, and a
reader skimming could take the first half as a promise that a gateway would settle the
question. The two sentences have to be read together, which is a weakness of prose and the
reason the plan's own phrase is reused rather than paraphrased: if the two documents disagree
later, they disagree in the same words.

**The question no rule asks: does stating a limitation make the register worse?** Section 12
reads weaker afterwards, and that is the intent. A provider register whose provenance is stated
is worth more than one a reader assumes was measured, and the alternative on offer was silence,
which is what #207 was filed about. Worth naming because the paragraph will read as a
concession to somebody who comes to it without the issue.

**What is deliberately not here.** No check against `project.yaml`, refused in the document
with the reason; no gateway comparison, which needs v2 and could only be a figure, named and
assigned to nothing; no new gate result; no field added or removed; no change to section 7 or
to `plugin_version`, both cited rather than revisited. P1's non-goals carry each reason.

**Three deviations, all recorded in P3.** The approved wording was wrong about where `model`
comes from and the specification commit was amended rather than corrected afterwards, because
it had not been pushed. The paragraph was replaced a second time to drop two em dashes the
document uses once in 132 kilobytes. And the register row is A100, because A98 and A99 are on
unmerged branches and a row's number is cited from other rows.

<!-- xeno:section:release-notes -->
## Release notes

Section 12 now says what the `model`, `tool` and `tool_version` triple is worth. All three are
declarations: `tool_version` is what the caller passed on `--tool-version` or the session
exported as `XENO_HARNESS_VERSION`, `model` is the default `project.yaml` declares, and `tool`
is what `project.yaml` names unless `XENO_HARNESS` says otherwise. G-Schema checks that each is
present and well shaped and nothing more.

Nothing corroborates them, and the paragraph says why rather than leaving a reader to wonder.
Section 7 keeps the runner from branching on the harness, so it cannot ask the harness anything
beyond what the harness volunteered. The only other copy of `tool` and `model` in the repository
is `project.yaml`, which is where the runner read them, so a check against that would make the
gap invisible instead of closing it. The one record from another hand is a gateway's, and the
plan says what it is worth: it names the deployment a request was routed to rather than what a
provider attests. It is outside the repository and outside the gate path, both by design.

So a project that keeps a register of its ICT service providers and draws on these fields is
reading what an agent declared. That is the sentence the paragraph ends on, and it is the one
thing such a project could not get from this document before.

No check changes. A green G-Schema means what it meant yesterday, no artifact gains or loses a
field, and no verdict in the trail moves. `docs/clause-readers.md` counts the new clause among
the explanations, 129 to 130, and says why it has no reader and what would falsify it.
`docs/assumptions.md` gains A100 with the decision, who took it and when, and the options not
taken.

One thing found in passing and not fixed here: `hashes` accepts the `by-hand` hash placeholder
on the strength of `tool: manual`, which is one of the self-reported fields. It is raised as
its own issue.

<!-- xeno:section:residual-risk -->
## Residual risk

**One gate acts on a self-reported field and the new paragraph does not say so.** `hashes`
accepts the `by-hand` placeholder in every hash field when the artifact says `tool: manual`.
An artifact that claimed to be manual would have its placeholders accepted and nothing
distinguishes the two cases. This is the sharpest instance of what #207 is about and it is
left open deliberately: it is outside the issue's "done when" and outside the wording that was
approved, and extending an approved sentence is the thing the exception to the first standing
rule does not cover. It wants an issue, and until it has one the fact lives in this intent's P4
and in a learning record, which is two places a reader of section 12 will not be.

**Stating the limitation is all that happened.** Nobody can trust the triple more than before.
If a project was going to put these fields in front of an auditor, the paragraph tells it not
to without saying what to do instead, because there is nothing to do inside v1. That is the
honest end of #207's two branches and it is still the weaker one.

**The gateway sentence can be skimmed as a promise.** "The one record from another hand is the
gateway's" reads, on its own, as though a gateway would settle the question. The sentence after
it says it is a routing record and not an attestation. Two sentences that have to be read
together are a weakness of prose, and the mitigation is only that the plan's own phrase is
reused rather than paraphrased, so the two documents will not drift apart in wording even if a
reader drifts.

**The gateway claim is inherited.** It rests on section 9's measurement against LiteLLM
1.102.1 for #56 and was not re-measured here. A different gateway may record something else,
and the paragraph speaks of "the gateway's" record as though there were one kind.

**Nothing guards the clause, by design.** It is an explanation, and
`docs/clause-readers.md` says what would falsify it and that both are outside v1. The
likeliest thing somebody builds, the comparison against `project.yaml`, leaves it true. So the
clause is as durable as the next reader's attention.

**The approved wording was wrong once.** It said `model` comes from `project.yaml` or the
harness, which the issue says too, and the code reads it from `project.yaml` alone. Drafting
and approval both passed it; a reading of `agent()` caught it. P0's learning record proposes
that drafted wording about the code cite the function, and nothing enforces that, so the next
such sentence is as exposed as this one was.

**The register skips two numbers on this branch.** A97 to A100, because A98 and A99 are on the
unmerged branches for #228 and #238, and the key sequence has the same property: three intents
open at once, each needing `--intent` passed by hand because `intent start` counts from one
branch. P0 of XENO-0268 carries that finding; nothing has scheduled the runner change.
