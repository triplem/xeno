---
intent: github.com/triplem/xeno#267
phase: 02-design
created: "2026-10-06T20:00:02Z"
schema_version: "1.0"
runner_version: dev+1d61fa2
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cf0cf484893ca92d101aa88859f55d6f0c593a8fb2974fc2b4c9fdfe2ae1f55e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: 'Record in the specification why P0''s lock names no files, rather than making it record them: a paragraph where the lock is described in section 5, a paragraph under section 5''s budget clause, a paragraph under section 7''s staleness clause, and the correction of two code comments that explain the empty list as a phase older than #217''s rule.'
      rationale: 'The scope is one artifact per intent and every later phase resolves the same one, so the budget is judged five times per intent and a moved file is found against the locks of P1 onwards; what goes unjudged is the intake against the budget the intake declares, and that is a smaller loss than either alternative costs. Completing the lock in scope set needs an ordering guard nobody would expect from the description, because SectionSet writes context_hash from the lock on every render and a scope set after the sections leaves a stale hash and a red G-Schema two commands later, which is #225; the guard then makes the scope a thing declared before the work, which is a quieter form of the second candidate rather than a different one. Taking the scope at phase start makes it an input to starting rather than P0''s product, which reaches section 6''s phase table and changes the phase model. Both stay written down in #267 with their costs, and the new paragraph is phrased so that adopting one means deleting it.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**Where the paragraph about the lock goes, and what it is next to.** Section 5 already has a
paragraph saying the lock records what was declared and not what was read, and another saying it
is written once and not refreshed. The new one goes after both, because it is a consequence of
the second and would read as a qualification of the first if it came earlier. The paragraph that
follows, about version numbers appearing in two places, begins a different subject, so nothing is
split.

**The schema block gets a comment and not a second paragraph.** `files:` in the example block now
carries `# empty at P0, see below`, in the register `tools:` already uses two lines down.
Somebody reading the shape of the artifact rather than the prose around it is who the warning is
for, and six words there save a reader who would otherwise write a tool against a field that is
populated for five phases out of six.

**The budget paragraph names the comparison, not the exclusion.** "The comparison is against the
lock of the phase being gated" is the whole of why P0 is outside it, so the mechanism carries the
statement and the exclusion follows from it. The alternative, asserting that P0 is exempt and
giving the reason afterwards, states a rule the reader then has to hold; this way there is no rule
to hold.

**The staleness paragraph says what it is not.** Section 7 opens its two limits with "Two limits
keep the check useful rather than noisy", and both are deliberate trades against false positives.
P0's exclusion is neither deliberate nor about noise, so the paragraph says so in its first
sentence and the count of limits stays two. It also names what is lost — a file that moved while
the intake itself ran — because a reader who is told a check does not apply somewhere will ask
what escapes, and an unanswered version of that question is how a check becomes decoration.

**Both wrong comments are corrected in the commit after the specification, not in it.** They
follow from the paragraphs, and `CLAUDE.md` puts the specification change first and the code that
follows from it after.

<!-- xeno:section:alternatives -->
## Alternatives

**`scope set` completes the lock's `files` when it writes the scope.** The smallest of the three
code changes: one call to `resolveScope` is already there, and `ScopeSet` would write its result
into the lock beside the scope. It keeps what #215 actually protects, because the lock would be
completed before the phase is judged rather than rewritten after.

What it costs is an ordering guard nobody would expect from the description. `SectionSet` writes
`context_hash` from the lock on every render, so a scope set after the sections leaves the
frontmatter hashing bytes that have since changed, and `phase finish` reports red on G-Schema two
commands later — which is #225 exactly, and #225 is the issue about a failure surfacing where the
mistake was not made. So `scope set` would have to refuse once `output.md` exists, which makes
the scope a thing declared before the work rather than an artifact of the phase, and that is a
quieter version of the second candidate rather than a different one. It also leaves a window in
which the lock exists and is empty, and makes one command the writer of two artifacts.

**`phase start` takes the scope as an argument at P0.** One write, in the right order, no window
and no guard. It is the honest form of what the first candidate half does.

What it costs is what P0 is. Section 5 says "P0 produces it as an artifact of its own, and a
phase reads what it names"; the scope would become an input to starting instead, which reaches
section 6's phase table, where P0's input is the issue and its output includes the context scope.
That is a change to the phase model and not to a command.

**Say nothing and leave the comments as they are.** Free, and the state #267 found: three readers
that read nothing, a measurement in an issue, and two comments in the code asserting a reason
that is false for the phase it is read at. It was chosen against because the next person to read
`staleReads` would believe the empty list is a phase older than a rule, which is what this
intent's author believed until the count came out 0 of 117.

<!-- xeno:section:impact -->
## Impact

`docs/process-definition.md`: three paragraphs and one comment in an example block. Section 5's
lock subsection, section 5's context economy and section 7's repetition subsection. 26 lines
added, one changed.

`internal/runner/runner.go`: `informationBase`'s doc comment. The sentence about a phase older
than #217's rule is true of P1 to P5 and is kept for them; the P0 case is named as what it is.

`internal/gates/gates.go`: `staleReads`'s doc comment, which carries the same sentence and gets
the same treatment, next to the two limits it already quotes from section 7.

`docs/clause-readers.md`: three rows. All three clauses are read by a person and by no code,
which the table has a shape for, since XENO-0260 and XENO-0265 both added clauses whose only
reader is a reader.

Nothing else. No artifact gains a field, no gate gains a check, no command changes its behaviour,
and `go test ./...` passes unchanged — which is the test of this intent as much as the absence of
a new assertion is.

**What a later reader inherits.** The two candidates are in #267 with what each costs, and the
paragraph in section 5 is written so that recording files at P0 would require deleting it rather
than editing around it: it says `files` is empty and why, not that it may be empty. If the figures
later argue for one of the candidates, the specification change that opens it is visible as a
deletion, which is the trace this intent owes the next one.
