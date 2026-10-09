---
intent: github.com/triplem/xeno#337
phase: 00-intake
created: "2026-10-09T15:48:44Z"
schema_version: "1.0"
runner_version: dev+9590797.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 448dfbbc2e99029e1356080d979abf8231efdc9762fd35d246c3e16474ae5f82
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
open_questions:
    - key: Q-1
      options:
        - consequence: the trail stays intact and the record is verifiable by the machinery that verifies a phase. The cost is that section 6 gains a record type that is not a phase, Appendix B gains its path, a new package owns it, and the two keys XENO- and IDEA- are a second sequence to keep apart
          reason: it keeps the trail intact, it has a thing to seal at a time when no intent and no issue exist yet, and the issues it produces are what the existing approval act then takes one by one
          recommended: true
          text: 'A sealed record beside the intents: a directory beside .xeno/intents/, for example .xeno/ideas/IDEA-0001/, holding output.md, digest.md, learning.yaml and gate.yaml exactly as a phase does, sealed by artifacts_hash and recomputed by gate verify like any phase; its sections are the vision and the plan, its output is issues on the tracker that cite the key, and each intent''s P0 quotes the idea it came from.'
        - consequence: 'no new record type. The cost is that a phase belongs to an intent and an intent to an issue, so the record would need an intent started from no issue, which section 12 now refuses; and the number contradicts the position, which #224 names as the permanent confusion'
          text: 'A phase appended as P6 and used first: phases/06-ideation/ inside an intent, with the existing phase machinery unchanged.'
        - consequence: nothing in the specification's record model moves. The cost is that the input and output documents are not gated, which the approval comment asks for; the only record of why the issues exist is the issues themselves, and a reader of the trail cannot verify what produced them
          text: 'A command or skill with no sealed record: xeno idea or a skill reads the vision, writes work packages and issues to the tracker through the adapter, and the trail begins as now with the first approved issue.'
        - free: true
          text: Something else, entered by the person deciding
      text: 'Where does the ideation record live? The approval comment on #337 says the idea becomes a vision and a plan, from which Xeno derives work packages and issues, with the input and output documents gated and decisions human. A gated document is a sealed record, and #224 shows a phase at the front of the list renumbers every artifact path and so every artifacts_hash in the trail.'
    - key: Q-2
      options:
        - consequence: one place says when ideation is done, and the size table and the sequence place it as one thing. The cost is a twenty-second package in a plan whose v1 was twenty; the size table and the sequence each gain a row, and a package touching four others is sized Medium at least
          reason: the record needs a template, a writer and a sealed verdict of its own, which is one done-when and not three
          recommended: true
          text: 'A new package, WP21 Ideation: section 2 gains a package with its own done-when, covering the record type in section 6 and Appendix B, a seventh template, the commands that write and judge it, the issues it produces through the adapter, and the P0 sentence that quotes it.'
        - consequence: no new package. The cost is that no single place says when ideation is done, three finished packages reopen, and WP12's contract of four operations widens to five, which its own text calls a decision
          text: 'Split across WP3, WP7 and WP12: the template to WP3, the record and its commands to WP7, the issue writing to WP12, each extending its done-when.'
        - consequence: one owner, an existing one. The cost is that WP7 is already the largest package and is done; its done-when would reopen for a feature with its own milestone shape, and the adapter has no write-issue operation today, so WP12 reopens anyway
          text: 'WP7 the runner, alone: the runner owns the record, its template and its commands as it owns intents and phases, and the issue writing is one more use of the adapter''s existing contract.'
        - free: true
          text: Something else, entered by the person deciding
      text: 'Which work package owns the ideation record? No package owns anything before P0 today, which is what made #337 a finding about the plan.'
    - key: Q-3
      options:
        - consequence: the plan says where the package sits and v1 is unchanged. The cost is that this repository's issues go on being written by hand through v1, and the first dogfooding of ideation waits one release
          reason: the scope section's one reason for deferring surface applies to this item as to the others, and the issue's done-when is met by the plan saying where the package sits, not by building it
          recommended: true
          text: '1.1, specified now and built then: WP21 is written into section 2 with its done-when, the size table and section 8 list it as specified and not built, and v1''s scope paragraph names the greenfield entry point among what waits.'
        - consequence: ideation is dogfooded before the release. The cost is that the scope section's reason for deferring surface has to be amended to say why this item is different, v1's release waits on a Medium package, and the proportionality figures it would be judged by are not in yet
          text: 'v1, sequenced after WP19: WP21 becomes step 14 of the sequence, v1''s in-scope list gains the entry point, and the 1.1 breadth could itself be the first idea run through it.'
        - consequence: 'a record can exist before the tooling. The cost is an artifact specified with no writer, which A82 and #217 show stays unproduced for a hundred intents, and a package split across two releases with two done-whens'
          text: 'v1 for the record type only, commands in 1.1: section 6, Appendix B and a seventh template in v1, so a record can be written by hand and verified, with the commands and the issue writing in 1.1.'
        - free: true
          text: Something else, entered by the person deciding
      text: 'Which release carries WP21? The plan''s scope section says every item deferred to 1.1 was deferred for one reason: it adds surface before dogfooding has said whether the process is proportionate. WP18 is the precedent, specified in section 2 and listed in section 8 as not built.'
decisions:
    - chosen: 'A sealed record beside the intents: a directory beside .xeno/intents/ holding the same four files a phase holds, sealed by artifacts_hash and recomputed by gate verify, whose sections are the vision and the plan and whose output is issues that cite its key.'
      decided_by: triplem
      id: D-1
      proposed_by: claude-fable-5-1
      rationale: 'The approval comment asks for the input and output documents to be gated, which rules out a command without a record, and a phase at the front of the list renumbers every artifact path in the trail (#224), while a P6 used first needs an intent that section 12 refuses to start from no issue. A record beside the intents is the one shape that has something to seal before any intent or issue exists and leaves every existing hash alone. The cost accepted: a record type that is not a phase in section 6 and Appendix B, a package to own it, and a second key sequence.'
      resolves: Q-1
    - chosen: 'A new package, WP21 Ideation, owns the record: its type in section 6 and Appendix B, a seventh template, the commands that write and judge it, the issues it produces through the adapter, and the P0 sentence that quotes it.'
      decided_by: triplem
      id: D-2
      proposed_by: claude-fable-5-1
      rationale: 'The record needs a template, a writer and a sealed verdict of its own, which is one done-when and not three; splitting it reopens three finished packages and widens WP12''s four-operation contract without a place that says when the whole is done, and WP7 alone reopens the largest package and WP12 besides. The cost accepted: a twenty-second package, a row in the size table and in the sequence, sized Medium at least for touching four others.'
      resolves: Q-2
    - chosen: '1.1: WP21 is specified in section 2 now with its done-when, listed in the size table and in section 8 as specified and not built, and named in v1''s scope among what waits, with WP18 as the precedent.'
      decided_by: triplem
      id: D-3
      proposed_by: claude-fable-5-1
      rationale: 'The scope section defers surface to 1.1 for one reason, that dogfooding has to say first whether the process is proportionate, and ideation is surface like the rest; the issue''s done-when is met by the plan saying where the package sits. Building it in v1 would need the scope section amended to say why this item differs and would hold the release on a Medium package, and specifying the record without its writer is the shape A82 and #217 show stays unproduced. The cost accepted: this repository''s issues are written by hand through v1 and ideation is first dogfooded in 1.1.'
      resolves: Q-3
---

# Intake

<!-- xeno:section:problem -->
## Problem

Approved by @triplem on 2026-10-09T11:33:43Z: There is an idea and this idea needs to get translated into a vision and implementation plan. With this, xeno can then start to.build issues. An issue could also be a place to start, so the issue is just an idea and xeno then generates work packages and according issues. Since there are input and output documents all of those should also be gated and agents can and should support this process, but the decisions need to be human on the loop.

> **github.com/triplem/xeno#337** — Nothing says how a greenfield project reaches its first issue
>
> Asked for on #323, where AI-DLC's lifecycle was read: five phases and 33 stages from
> ideation to operation, cut by profile. Xeno's six phases begin at P0 with an issue that
> exists, which is the brownfield case the plan is built for. For a greenfield project
> there is no issue yet, and nothing in the process says how the first ones come to be.
>
> ## What ideation would be here
>
> The stage before there is a request: a sentence, a conversation, becomes the set of
> issues that then carry the label and become intents one by one. It produces issues and
> not artifacts, so on the face of it it sits upstream of P0 and outside the phase model,
> beside the pre-intent act #330 built and #224 discusses, and the ideation output is
> what that act then approves or declines.
>
> Whether it is a phase at all is the first question. Section 6 keeps the phase list open
> at the end only, and #224 says what inserting one at the front costs: every path under
> `.xeno/intents/` and so every `artifacts_hash` in the trail. Appended as P6 and used
> first, it has the numbering contradiction #224 names. Outside the model, it is a
> command or a skill that writes to the tracker and leaves no sealed record, and the
> trail begins, as now, with the first approved issue.
>
> ## What this issue is
>
> A finding about the plan, by the third standing rule: no work package owns anything
> before P0, and `docs/implementation-plan.md` does not say whether Xeno covers the ends
> of the lifecycle, ideation and operation, or stops at the merge. That is the
> maintainer's decision, in the plan, before anything is designed. The decision has three
> shapes, each with its cost, and the question to put first is whether a greenfield
> entry point is Xeno's at all; the shape follows from the answer.
>
> ## Done when
>
> The plan says whether an entry point before the first issue is in scope, and if it is,
> which package owns it and whether it is a phase, a command or a skill; and #224's
> reasoning about the front of the phase list is cited rather than repeated.
>
> Refs #323, refs #224
>

The approval comment is the instruction this intent works from, and it says more than
the issue asked: an idea is to become a vision and an implementation plan, from which
Xeno builds work packages and issues; an issue may itself be the idea; the input and
output documents are gated, agents support the process, and the decisions are a
person's. That answers the issue's first question, whether a greenfield entry point is
Xeno's at all, with yes, and leaves the shape, the owner and the release.

**What the plan said before this intent.** Section 1 lists what v1 covers and what
waits for 1.1, and names "operations and maintenance phase" among what is out of scope
by decision, so the far end of the lifecycle was already settled and only the near end
was open. Section 2 has twenty packages built for v1 and WP18 specified and deferred,
the one precedent for a package that is written down and not built. Section 6 keeps
the phase list open at its end only, and #224 says what the front costs: every path
under `.xeno/intents/` and so every `artifacts_hash` in the trail. Section 12 of the
process definition, since #330 and #332, starts an intent from an approved issue and
from nothing else, which is a reason of its own against a phase used before any issue
exists. Nothing in either document names the stage before the first issue, which is the
finding the third standing rule makes of it.

**What was decided, and how.** Three questions, one at a time, each with its options,
their consequence and cost, and a recommendation with its reason, as section 8 asks
and as Q-1 to Q-3 below record. The record is a sealed directory beside the intents and
not a phase, D-1; a new package WP21 owns it, D-2; it is specified now and built in
1.1, D-3. The plan was then edited in place, read back as a diff, approved by the
maintainer as written, and committed as `ec372ea` ahead of this trail, which the first
standing rule makes an exception and the commit message records.

<!-- xeno:section:scope -->
## Scope

This intent changes one document, `docs/implementation-plan.md`, so that it says what
#337's done-when asks: whether an entry point before the first issue is in scope, which
package owns it, and what it is.

- **Section 1.** The list of what waits for 1.1 gains the entry point, WP21, as the
  one item there that adds a record rather than reach.
- **Section 2.** A package, WP21 Ideation, deferred to 1.1 in WP18's shape: what the
  stage is, why the record is beside the intents and not a phase, citing #224 rather
  than repeating it, what the record holds and produces, what building it touches,
  and when it is done.
- **Section 6 and 6.1.** The sentence that moves WP18 to 1.1 moves WP21 with it, and
  the size table's unsized row names both.
- **Section 8.** The deferred list names the record beside the dashboard.

What it does not do.

**It does not touch the process definition.** The record type, its path in Appendix B,
its template and its commands are WP21's, and each is a specification change first
when the package is built; fixing them now would be designing a 1.1 package in a v1
intent.

**It does not size WP21.** It sits in the unsized row beside WP18, because the plan
sizes what it sequences and WP21 is not sequenced; D-2's remark that a package touching
four others is Medium at least is an expectation and not a row.

**It does not name the key prefix, the directory or the commands.** The plan says "a
key sequence of its own" and "a directory beside `.xeno/intents/`", which is what the
decision fixed; the names are for the package.

**It does not answer #323.** The evaluation that raised the lifecycle question stays
open; this intent takes the one finding it produced.

<!-- xeno:section:context-rationale -->
## Why this context

- Issue #337 with its approval comment, which is the instruction; #224 for the cost of
  a phase at the front, #323 for where the question came from, #330 for the pre-intent
  act the record feeds.
- `docs/implementation-plan.md`: section 1 for what is in and out of v1 and why things
  wait, section 2 for the package shape and WP18 as the deferred precedent, section 6
  and 6.1 for the sequence and the size table, section 8 for the 1.1 scope of record.
- `docs/process-definition.md`: section 6 for the phase list being open at the end
  only, section 8 for the shape a question takes, section 12 for what an intent starts
  from.
- `docs/assumptions.md`, A107 and the rows around it, for whether a register row is
  owed; none is, because every choice here is the maintainer's decision and not an
  assumption.
- `CLAUDE.md` for the three standing rules this intent works under, the first of which
  it records an exception to.
