---
intent: github.com/triplem/xeno#235
phase: 05-review
created: "2026-10-06T18:05:09Z"
schema_version: "1.0"
runner_version: dev+5276f4b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ba2d9902705a43f04934f7f4e5ca45626a1e32452a3e911a06142e68532e34ff
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: deviation
      note: 'Five, in P3. (1) Status was not in P2''s decisions and had to change: section 5 says the check is pass and the phase is green, which is two statements with two readers, and with result alone a phase whose only finding was a budget overrun came out red with every check passing. Found by running it, not by reading it. (2) P1''s criterion 7 is the only one written as a behaviour — the phase green and the next phase starting — and it is what caught that; a criterion written as ''the check result is pass'' would have passed over it. (3) an advisory finding is undecided for ever, which is now load-bearing in Status rather than incidental. (4) the bound is tested by counting occurrences of a string in a source file, which is a weak and unusual reader that will false-red on a rename, and is in because section 5''s ''nothing else writes the field'' has no other. (5) the end-to-end fixture writes the lock''s files by hand, because no P0 lock records one (#267) and the runner cannot reach that state. No criterion was falsified.'
    - rule: interface-change-needs-a-migration-note
      result: deviation
      note: 'One change an adopter meets and it needs no action of theirs: a project that has declared a context budget finds an overrun reported rather than blocking, from the next release, because internal/gates and internal/model ship in the binary. Nothing migrates — model.Finding''s new field carries omitempty, so no gate.yaml already written changes shape and 0 of 469 sealed verdicts carry the key; the four check results, the five phase statuses and hashing.FindingID are untouched, so nothing that reads a verdict learns a new state. Worth a note rather than a migration: a reader of any verdict now has a distinction to hold, since a finding in gate.yaml may be one the phase was not failed for; and the advisory key is in the artifact schema, so a project''s own external gate under section 14 can write it into a check of its own with nothing to notice.'
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: None added and go.mod is untouched. The whole change is one boolean field, two comparisons and a one-line helper. The only unusual import is in a test, which reads gates.go with os.ReadFile and counts occurrences to hold section 5's bound — in the standard library, and the alternative was no reader for that sentence at all.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** No normative document is touched: the clause was committed in
XENO-0263 and this is the code that follows it, which is the order the first standing rule asks for.
Nothing is invented — `advisory` on a finding is what section 5 enumerates, the four check results
and the five phase statuses are untouched, and `hashing.FindingID` is unchanged. The work belongs to
#235 under `wp8`, on a branch carrying one intent, and the commit references the issue without
closing it: the clause is honoured and #267 still means the finding cannot arise from a real scope.

**The clause had two consequences and the design planned for one.** "The check carrying it is
`pass`, the phase is `green`." P2 named `model.Finding`, `result` and `budget`. `Status` derives the
phase from undecided findings without consulting the check result, so with `result` alone a phase
whose only finding was a budget overrun came out red with every check passing — half the clause, and
the half that was missing is the half that stops work. Found by running it. The P3 learning is that a
clause stating two consequences has its readers traced before the design is written.

**One criterion was written as a behaviour and that is what caught it.** Criterion 7: the phase
green *and* the next phase starting. A criterion written as "the check result is `pass`" would have
passed over the fault and every other criterion would still have been met.

**Every assertion was confirmed able to fail.** Five mutations, five failures, including the one that
matters most: a second check reaching for the field makes the bound test fail with "3 ... want the
budget clause's 2". The bound is section 5's sentence and that test, and no code can do better —
nothing can tell a clause that legitimately asked for the field from a check somebody found
inconvenient. A90 is the reason it is named as a risk rather than treated as solved.

**The reason the contradiction survived was measured, not guessed.** Nine tests call `budget`
directly and none asserts what the check result should be, which is the one thing section 5 is
explicit about. So the behaviour that broke the specification had no reader in the suite either, and
the new tests go through `result` and `Status`.

**Every negative result was re-measured with the thing present.** #263's convention. "Nothing else
writes the field" is a count of two call sites and one assignment plus a search outside both files
returning zero. "Nothing in the trail moved" is 0 of 469 sealed verdicts carrying the key and the
same verdict count before and after. "The enumerations are untouched" is a `grep` that finds both of
them still in section 5.

**The end-to-end fixture is named as a fixture.** The lock's `files` were written by hand, because
the runner cannot produce that state: no P0 lock records one, which is #267. The evidence says so
rather than presenting it as a run of the runner.

**This intent is the first subject of the checks written in the one before it.** Its P1 renders at
`requirements@1.1.0` and its P4 at `verification@1.1.0`, so `numberedCriteria` and `mappingComplete`
judged real artifacts for the first time. Both were exercised rather than assumed: removing a row
from the mapping turned G-Test red naming criterion 11, and restoring it turned it green. XENO-0264's
criterion 9 failed because it could not do this; this intent is where it happened.

<!-- xeno:section:release-notes -->
## Release notes

**A budget overrun no longer stops the phase.** `budget` is called from `schema` and `result` failed
a check on any finding at all, so a recorded context over its declared budget turned G-Schema red and
`predecessorAllowsStart` refused the next phase. Section 5 says of that same check that it is
"deliberately a finding and not a red gate in the sense of stopping work", and the code's own comment
quoted the sentence it broke.

**`model.Finding` gains `Advisory` with `omitempty`.** `result` fails a check only where a finding is
not advisory; `Status` no longer counts an advisory finding as undecided. Two readers, because
section 5 says two things in one sentence — the check is `pass` and the phase is `green` — and
changing only the first left a phase red with every check passing.

**`budget` marks both of its findings**, the file count and the byte total, through a one-line helper
that is the only place in the runner setting the field. Section 5 bounds it there — "nothing else
writes the field" — and a test counts the writers, because that sentence has no other reader.

**End to end:** a phase whose only finding is a budget overrun comes out `green`, the next phase
starts, and `gate approve` on the finding gives `approved`. Those are section 5's three claims —
visible, not stopping work, decidable like any other finding — and each was run rather than inferred.

**Nothing in the trail moved.** `omitempty` means no sealed `gate.yaml` changes, 0 of 469 carry the
key, and `gate verify` reports the same count. The four check results, the five phase statuses and
`hashing.FindingID` are untouched — the flag is deliberately outside the id, because section 5 says
a finding's id does not depend on the run.

**`docs/clause-readers.md` gains a row for section 5's budget clause**, which the table had never
carried: the clause section 5 states most plainly about a gate's behaviour was the one the document
did not list. It names `gates.advisory`, `result` and `Status`, because violating the clause means
any of the three failing to do its part. The count is forty-one.

## Why it survived this long

Nine tests call `budget` directly and **none asserts what the check result should be** — the one
thing section 5 is explicit about. The behaviour that contradicted the specification had no reader in
the suite either. The new tests go through `result` and `Status`.

## Not fixed by this

**#267.** `budget` sums bytes over a lock's `files` and no P0 `context.lock.yaml` in this trail
records one, so the check reads nothing at the phase that declares the budget. The end-to-end
evidence rests on a lock written by hand, and both issues have to close before the runner produces
an advisory finding once.

<!-- xeno:section:residual-risk -->
## Residual risk

**The clause is honoured and the finding still cannot arise.** No P0 lock records a `files` list
(#267), so `budget` reads nothing where a budget is declared. The scratch evidence wrote the lock by
hand. This intent closes one of the two faults and the other is open.

**"Finding" means two things now.** A finding in `gate.yaml` may be one the phase was not failed for,
and every reader of a verdict carries that distinction. The cost section 5's clause was always going
to charge, charged to the reader.

**The bound is a sentence and a string count.** `TestOnlyTheBudgetClauseWritesTheAdvisoryField` reads
`gates.go` and counts two call sites, so it fails if the helper is renamed or a call reformatted — a
false red for a true property — and it is in because nothing better exists. No code can tell a clause
that legitimately asked for the field from a check somebody found inconvenient, and a project's own
external gate under section 14 can write the key into a check of its own with nothing to notice.

**An advisory finding is undecided for ever and nothing asks about it.** It never becomes an
obligation and never reaches the next-step suggestion. A reader who wants to know whether anybody
looked at a budget overrun has the artifact and nothing else.

**`docs/clause-readers.md`'s count moved twice today.** Forty in XENO-0264, forty-one here, having
been wrong by one before either. It is a figure in prose beside the rows it describes; the learning
proposing it be derived from the table is recorded and has the same reader as every other learning —
section 10's merge request, which nothing has opened.

**Three learnings from this session now propose conventions about the same thing**: a clause with two
consequences, a criterion written as a behaviour, and a figure in prose beside its rows. All sit in
phase `learning.yaml` files. Section 10 is explicit that this is the route and equally explicit that
nothing takes effect on being noticed.
