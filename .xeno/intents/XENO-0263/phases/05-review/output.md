---
intent: github.com/triplem/xeno#235
phase: 05-review
created: "2026-10-06T15:25:35Z"
schema_version: "1.0"
runner_version: dev+8decfdc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 59d0271082767c2c6562ff3f416b727944f0a5464c9768453efd7e6562c376c0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: deviation
      note: One, in P3. The paragraph is placed after the four-values paragraph where the draft said 'after Decisions sit on findings, not on phases'. The status derivation runs over two paragraphs and the second completes the first, so taking the draft literally would have inserted a clause about a finding that does not fail into the middle of the argument about findings that do. The prose is committed as drafted; the placement is the document's. It is the third time in two intents that a clause drafted in an issue comment needed adjusting on contact with the file, after a heading that would have deleted an existing sentence and a figure citing this repository where the specification cites none, and the P3 learning proposes showing the person the diff rather than the prose. No criterion was falsified.
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: 'No interface changes and deliberately no code. Sixteen lines added and one replaced in docs/process-definition.md, and nothing else outside the trail: no Go source, no field, no command, no gate, no template, no rule. The four check results are untouched, so no reader of a verdict learns a new state yet. The specification is in no artifacts_hash and no rules_hash, so every verdict stands and gate verify reports the same count; it is not shipped by xeno init --vendor, so an adopter sees nothing until the code follows. What a reader of the document gets is a key no runner writes, which P3''s deviations records as the second such key this session added.'
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: None added and go.mod is untouched; the whole diff is Markdown. Nothing was weighed either. The one route that would have used an existing mechanism was drift, which is in the specification and in no code at all — 0 Drift in model.go against 2 for the fields beside it, 0 writers under internal/ or cmd/, 0 of 457 sealed gates — and its row carries sha256 hashes where a budget overrun is two byte counts.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The first standing rule is broken deliberately, for the second time in this session, and this is
the record of it.** "The documents are not editable by the agent." The agent edited one. The wording
was drafted on #235, put to the person, and committed at their instruction rather than written by
them. The rule's purpose — that a person decides the specification — held; its letter did not. It is
in P0's problem section, in the design decision, here, and in the commit message.

**Twice in one session is worth naming, because an exception that recurs is on its way to becoming
a habit.** XENO-0262 made the same record for #258's two amendments an hour earlier and recorded a
learning proposing a convention for it. That learning sits where every learning sits, routed through
section 10 to a merge request nobody has opened. The honest statement is that nothing stops the
third time being unremarked.

**The other two standing rules hold.** Nothing is invented — the amendment adds a key to what a
finding may carry, which is what the second rule asks a specification change to do, and it leaves
the four check results alone. The work belongs to #235 under `wp8`, on a branch carrying one intent,
and the commit references the issue without closing it: the issue is not finished until the code
exists.

**No code, which is the half of the first rule that is not bent.** `git diff --stat` names
`docs/process-definition.md` and nothing else outside the trail. No `Advisory` on `model.Finding`,
no change to `result`, no change to `budget`, no test.

**The drafted placement was wrong and the document decided.** The draft said the paragraph goes
"after 'Decisions sit on findings, not on phases'". The status derivation runs over two paragraphs
and the second completes the first, so taking the draft literally would have split an argument in
half. That is the third time in two intents that wording drafted in an issue comment needed its
placement adjusted on contact with the file, and it is the P3 learning.

**The exception the field creates is bounded in the clause, because nothing else will bound it.** A
field that lets a finding not fail is the most useful thing in this runner for anybody who finds a
check inconvenient. No code will limit it; a paragraph does, and it is a paragraph of its own rather
than a caveat at the end of another, so that it reads as a rule.

**Every negative result was re-measured with the thing present.** #263's convention. "`drift` is
unimplemented" rests on a read of `model.go` that finds `RunAt` and `ArtifactsHash` beside the
absent `Drift`, on a search for a `drift` key across `internal/` and `cmd/`, and on 0 of 443 sealed
gates — not on one grep returning nothing. "The specification cites this repository nowhere" rests
on a search returning four matches in `CLAUDE.md`.

**#267 is deliberately not folded in.** The budget check reads nothing at P0 because no P0 lock
records a `files` list, so an advisory budget finding is still a finding nobody will see. Two
issues, two causes, and merging them would have made one intent that fixed neither cleanly.

<!-- xeno:section:release-notes -->
## Release notes

**A finding may now be advisory, which is what section 5 has asked for since it was written.** The
budget check turns G-Schema red and the phase with it, and section 5 says of that same check: "That
is deliberately a finding and not a red gate in the sense of stopping work: it is visible, it can be
decided like any other finding." The code's own comment quotes the sentence it breaks. By the first
standing rule the specification wins, and the reason nothing had been done is that **a finding that
is visible, recorded and decidable without stopping work did not exist in this runner**: `result`
returns `fail` for any non-empty finding list, and the four check results are fixed.

**`advisory` is a key on a finding.** In section 5's `gate.yaml` block, absent in the ordinary case.
A paragraph says what it does: the check carrying it is `pass`, the phase is `green`, and the
finding is in `gate.yaml` with its id, its cause and its remedy like any other. And why: blocking
against a number nobody has experience with would be the wrong way round, and a check that fires
with nothing to do about it is one people learn to route around.

**A second paragraph bounds it.** A finding is the thing that fails; an advisory one is readable
only because it is rare; the clause that asks for one says so where its check is described, and
nothing else writes the field. No code will enforce that, which is exactly why it is written down.

**The Context economy subsection says it where the budget clause is.** One sentence, so that a
reader arriving at the budget check is told there rather than inferring it from a yaml block forty
lines earlier — which is how the code and the document came to disagree for this long.

**What was ruled out, and measured first.** Section 5 already enumerates a channel for something
marked and not blocked: the `drift` list, with section 16's ninth limitation stating the principle
in as many words. It is entirely unimplemented — no `Drift` field, no writer anywhere in `internal/`
or `cmd/`, 0 of 443 sealed `gate.yaml` files — and its row carries sha256 hashes where a budget
overrun is two byte counts. A fifth check result was rejected for putting the property on the check
rather than on the thing that is advisory. Changing section 5 to let the budget block was weighed
and declined: one budget has ever been declared, so there is still no experience with the number.

**No code is in this commit.** No `Advisory` on `model.Finding`, no change to `result` or `budget`,
no tests, no row in `docs/clause-readers.md`.

**The agent wrote this, on instruction**, as it did for #258's amendments an hour earlier.
`CLAUDE.md`'s first standing rule says it does not; the person read the drafted wording and asked
for the commit. Recorded here because the reader has the current tree and nothing else.

**Not fixed by this:** #267. The budget check sums bytes over a lock's `files`, and no P0
`context.lock.yaml` in this trail records one, so the check reads nothing at the one phase that
declares a budget.

<!-- xeno:section:residual-risk -->
## Residual risk

**"Finding" now means two things.** Every reader of a verdict has a distinction to hold: a finding
in `gate.yaml` may be one the phase was not failed for. That is the whole cost of this shape and it
is paid by the reader. It is bounded by a paragraph and by nothing else.

**Nothing will stop the field spreading.** No code can tell a clause that legitimately asked to be
advisory from a check somebody found inconvenient. The bound is a sentence in section 5, and a
sentence is a person-reader — the category `docs/clause-readers.md` exists to count.

**The agent edited a normative document twice in one session.** Both are recorded; nothing checks
that they stay exceptional. The learning XENO-0262 wrote proposes a convention and sits where every
learning sits, routed through section 10 to a merge request nobody has opened. The third time has
nothing in its way.

**The clause still has no reader and will not until the code lands.** `Advisory` is in no struct,
`result` is unchanged, and `budget` still fails its check. The specification now describes a runner
that does not exist, which is the ordinary state of this document and is the second key this session
added to it that nothing writes.

**And the check reads nothing anyway.** `budget` sums `f.Bytes` over `lock.Files`, and 0 of 117 P0
locks in this trail record a `files` list (#267). So even once advisory, the budget finding will not
be produced at the phase that declares the budget. Two issues have to close before the clause this
amendment enables is exercised once.

**`docs/clause-readers.md` is now one row short.** Section 5's budget clause has no entry; the gap
was defensible while the clause had no mechanism and is a gap now. It belongs to the implementing
intent, because the column it would carry is the reader.

**The committed prose is the drafted prose and the placement is not.** The person approved wording,
not a location, and the location changed on contact with the file. They are being told, which is a
person reading a report — the same mechanism, and the same limit, as XENO-0262 recorded for #258.
