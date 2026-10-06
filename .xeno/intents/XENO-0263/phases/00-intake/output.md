---
intent: github.com/triplem/xeno#235
phase: 00-intake
created: "2026-10-06T15:21:29Z"
schema_version: "1.0"
runner_version: dev+8decfdc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 11db835658e8a5253ec394ab3188e2599e97e2f997d3f485810eb4cdabad07d1
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

`budget` is called from `schema` (`internal/gates/gates.go`), and `result` sets a check to `fail`
where it carries any finding at all. So a recorded context over its declared budget turns G-Schema
red, and a red gate stops the phase: `predecessorAllowsStart` refuses to begin the next one on an
undecided failure.

**Section 5 says the opposite, and gives the reason.** Of the same check: "That is deliberately a
finding and not a red gate in the sense of stopping work: it is visible, it can be decided like any
other finding, and blocking against a number nobody has experience with yet would be the wrong way
round. What it prevents is the budget quietly becoming decoration."

The code's own comment on `budget` quotes that sentence, and the placement contradicts it. By the
first standing rule the specification wins, so this is the code being wrong rather than a question
about what is wanted.

## There is no mechanism to move the check into

`model.Check` carries a result from the set A4 and A42 fixed — `pass`, `fail`, `pending`,
`not-implemented` — and `result` is the only constructor: a finding fails its check, and a failing
check makes the phase red. **A finding that is visible, recorded and decidable without stopping work
does not exist in this runner**, and section 5 asks for exactly that here.

So the repair is not a call site. It is the thing the artifacts may carry, which the second standing
rule makes a specification change and the first makes a person's commit before any code.

## What was ruled out before deciding, and why it is worth recording

Section 5 already enumerates a channel for something marked and not blocked: the `drift` list in
`gate.yaml`, with `file`, `field`, `artifact` and `gate_run`. Section 16's ninth limitation states
the principle in as many words — "Rule drift is marked, not blocked… the difference appears in
`drift` and as a line in the P5 checklist."

**It is entirely unimplemented.** No `Drift` field on `model.Gate`; no writer of a `drift` key
anywhere in `internal/` or `cmd/`; 0 of 443 sealed `gate.yaml` files carry one. The same search over
`model.go` does find `RunAt` and `ArtifactsHash`, so the absence is the answer and not a bad pattern.

It was rejected on two counts. Implementing it for the budget would be building a specified-but-absent
mechanism for its second purpose before its first, leaving rule drift unmarked afterwards. And the
row does not fit: `drift` carries `artifact` and `gate_run` as sha256 hashes, where a budget overrun
is two byte counts.

## What the decision was

#235 settled on an advisory finding: a field on `model.Finding`, with `result` choosing `fail` only
on a finding that is not advisory. It is what section 5's sentence literally describes — "deliberately
a finding and not a red gate", "visible", "can be decided like any other finding" — and it leaves the
check's result `pass`, which is one of the four the specification already defines.

That decision is recorded on the issue and nowhere the specification's reader will be, and it cannot
become code until section 5 enumerates the field.

<!-- xeno:section:scope -->
## Scope

In scope are three edits to `docs/process-definition.md`, all in section 5 and all one change: the
`advisory` key in the `gate.yaml` findings block, a paragraph saying what it means and why the
exception exists, and one sentence in the Context economy subsection saying the budget's finding is
one.

In scope is the paragraph saying the exception stays an exception. A finding is the thing that fails
in this process, and an advisory one is readable only while it is rare; a clause that introduces the
field without bounding it invites every noisy check to reach for it.

In scope is the paragraph being placed where the derivation it qualifies has finished. Section 5
explains the phase status over two paragraphs — "Decisions sit on findings, not on phases" and the
one about the four values not being interchangeable — and a clause about a finding that does not
fail belongs after both rather than between them.

In scope is the Context economy sentence. The clause the code contradicts is there, so a reader who
arrives at the budget check has to be told in the same paragraph that the finding is advisory;
leaving it to be inferred from a field in a yaml block forty lines earlier is what left the code and
the document able to disagree for this long.

Out of scope is `Advisory` on `model.Finding`, `result` ignoring an advisory finding, `budget`
marking its own, and the tests. All of it follows this commit and the first standing rule puts it
after, in its own intent.

Out of scope is `drift`. It stays unimplemented and section 16's ninth limitation stays unread. This
intent records in its artifacts that it was weighed, and does not touch the clause.

Out of scope is a fifth check result. A4 and A42 fixed that set and the decision was taken partly to
leave it alone; the amendment adds a field to a finding and changes nothing a check may report.

Out of scope is `docs/clause-readers.md`. Section 5's budget clause has no row there at all, which
is a gap worth a row once the check reads something; nothing fails today and a row naming a decision
is not a row naming a reader.

Out of scope is #267, the empty P0 lock. `budget` sums `f.Bytes` over `lock.Files` and no P0 lock in
this trail records one, so the check reads nothing at the one phase that declares a budget. Making
the finding advisory does not change that and making the lock record files does not change this. Two
issues, deliberately not merged.

Out of scope is #258's two amendments, which merged before this branch began. Same document, same
session, different issues and different work packages, and the third standing rule gives each its
own branch.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the section being amended, the code that contradicts it, and the files that fix what the
amendment may say.

`docs/process-definition.md` is read for three places and for two mechanisms it must not disturb.
Section 5's gate result subsection is where the findings block and the status derivation live, and
where the four check results are enumerated — read to confirm the amendment needs no fifth. The
Context economy subsection is where the budget clause the code contradicts sits. Section 16's ninth
limitation is read for the principle "marked, not blocked" and for `drift`, the channel it names,
because ruling that route out honestly meant reading what the document promises of it rather than
recalling it.

`internal/gates/gates.go` is read for `budget`, for `schema` calling it, and for `result`, which is
the three lines that make this a specification question rather than a call site. `result` returns
`fail` for any non-empty finding list, so there is no shape in the runner for what section 5 asks,
and the comment on `budget` quotes the very sentence it breaks.

`internal/model/model.go` is read for `Check` and `Finding`, to see exactly what a field would be
added to and to confirm `Drift` is not there. The same read finds `RunAt` and `ArtifactsHash`, which
is what makes the absence of `Drift` evidence rather than a search that matched nothing — the
convention from #263, applied to a struct.

`internal/runner/runner.go` is read for `predecessorAllowsStart`, because "stopping work" in this
runner is a specific thing: a red predecessor refusing the next phase. Without that, "not a red gate
in the sense of stopping work" is a phrase rather than a behaviour, and the amendment has to be
written against the behaviour.

`docs/assumptions.md` is read for A4 and A42, the rows that fixed the result set, to confirm the
amendment does not reopen them; and for A90, a reader that cannot fail is worse than none, which is
why the paragraph bounds the exception rather than offering the field to anything that wants it.

`CLAUDE.md` is read for the three standing rules. The second is why this commit is needed, the first
is why it is a person's and why it carries no code, and the third is why #258's amendments were a
different branch. It is also read for the convention on verifying a negative, which governs every
claim about `drift` here.

The `drift` figures were measured and not recalled: no `Drift` field in `model.go`, no `drift` key
written anywhere under `internal/` or `cmd/`, and 0 of 443 sealed `gate.yaml` files carrying one,
counted over the trail.
