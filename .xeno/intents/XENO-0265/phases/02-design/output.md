---
intent: github.com/triplem/xeno#235
phase: 02-design
created: "2026-10-06T17:52:59Z"
schema_version: "1.0"
runner_version: dev+5276f4b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b184f6ec704417ec518016e05b9a687a18f9840f55100367e85440fd3cdfdf8f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: Advisory bool with omitempty on model.Finding; result fails only where a finding is not advisory; budget sets it on both of its findings. The bound section 5 writes is carried as a comment beside the field and as a test counting the writers.
      rationale: 'result is the only constructor of a check and schema appends from six checks before calling it, so the property has to be on the finding rather than on the check or the caller — which is also what section 5 enumerates. A check carrying one advisory finding and one ordinary one still fails, which is the case a count of findings cannot express and the real test of the change. Nothing about the field enters hashing.FindingID, because section 5 says finding ids do not depend on the run and an id that moved when a flag moved would not be stable. The bound can only be a comment and a weak test: no code can tell a clause that legitimately asked to be advisory from a check somebody found inconvenient, and A90 makes that the risk worth naming rather than hiding. The clause is honoured end to end only if a phase whose sole finding is the overrun is green and the next phase starts, because stopping work in this runner is predecessorAllowsStart.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**`result` reads the finding rather than counting findings.** One line changes meaning: a check fails
where any finding is not advisory, and passes where none is or where there are none at all. That is
the only place it can go — `schema` appends from six checks before calling `result`, so "G-Schema
ignores budget findings" would have to know which finding came from where, and `result` already has
the findings in front of it.

**The field is on the finding and the bound is a comment beside it.** Section 5: "the clause that asks
for it says so where the check is described, and nothing else writes the field." No code can tell a
clause that legitimately asked from a check somebody found inconvenient, so the comment carries the
sentence and a test counts the writers. A test over a count is a weak reader and it is the only one
available; it fails when a second check reaches for the field, which is the moment worth catching.

**`budget` marks both findings, not one.** The file count and the byte total are the same clause
measured two ways, and section 5's sentence is about the check rather than about one of its numbers.

**Nothing about the field enters the finding's id.** `hashing.FindingID` hashes gate, rule id, file
and cause, and section 5 says ids do not depend on the run. Adding `advisory` to that hash would
change every budget finding's id the day it became advisory, and an id that moves when a flag moves
is not stable.

**The row in `docs/clause-readers.md` names `result` and `budget` together.** The table's own
definition is the thing that would fail if the clause were violated, and violating this clause means
either `budget` not marking its finding or `result` failing on one that is marked. Naming one would
leave the other unguarded in the document.

**The end-to-end check constructs a lock with sizes in it.** The budget check has no input in this
repository — 0 of 117 P0 locks record a `files` list, which is #267 — so a scratch intent cannot
produce the finding from a real scope. The lock is written by hand for the check, and the evidence
says so rather than presenting a fixture as a run.

**A phase whose only finding is the overrun is checked for being green *and* for letting the next
phase start.** The second is the one that matters: "not a red gate in the sense of stopping work" is
`predecessorAllowsStart`, and a green check with a phase that still refuses its successor would
satisfy the letter of the clause and none of it.

<!-- xeno:section:alternatives -->
## Alternatives

**Have `schema` drop the budget findings into a separate advisory list of its own.** It keeps
`result` untouched and the field off `model.Finding`. It needs somewhere to put them, which is a new
shape in `gate.yaml` — and section 5 enumerates `advisory` on a finding, so this would be inventing
a second mechanism for the clause that already has one. Rejected on the second standing rule.

**Give `budget` its own gate, whose failure does not bind the sequence.** The third shape #235 named.
There is no gate like that: every check's result feeds one phase status and a red status refuses the
next phase. It would need a new gate and a new rule about what a verdict means, which is section 7's
and WP1's. Rejected as the largest route to the same place.

**Put `advisory` on the check rather than on the finding.** One field instead of one per finding, and
`result` would not change. A second advisory finding in G-Schema would then drag the whole check with
it — G-Schema carries six checks' worth of findings — and XENO-0263 rejected it in the specification
for that reason. Rejected again here, consistently.

**Make `result` take the advisory set as a second argument.** It would keep the property out of the
artifact entirely, so nothing is added to `gate.yaml`. Then a reader of a verdict cannot tell which
finding was advisory, which is the thing section 5 asks to be visible. Rejected: the field exists for
the reader, not for the constructor.

**Suppress an advisory finding from `gate.yaml` and print it only.** Shorter, and it keeps "finding"
meaning one thing. Section 5 says the finding "is in `gate.yaml` with its id, its cause and its
remedy like any other", and a printed line is not decidable — `gate approve` takes an id. Rejected on
the clause.

**Let the budget finding be approved at P0 every time.** No field, no change to `result`: the
mechanism exists and a person approves it. It makes a routine measurement a governance statement, and
an approval everybody makes every time is a record nobody reads — A90 from the other side. Rejected,
and it is what the code effectively demands today.

**Add a test that no other check is advisory by enumerating the checks.** Stronger than counting the
writers, and it would name any new one. It has to be updated whenever a check is added, so it fails
for the wrong reason and gets weakened the first time somebody is in a hurry. Rejected in favour of
the count, which fails only when the field spreads.

<!-- xeno:section:impact -->
## Impact

**`internal/model/model.go`.** `Finding` gains `Advisory bool` with `omitempty`, and a comment
carrying section 5's two sentences — what it means and that it stays the exception.

**`internal/gates/gates.go`.** `result` fails only where a finding is not advisory. `budget` sets the
field on both of its findings. Nothing else.

**What a reader of a verdict meets.** A check that passed, carrying findings. That is new: until now
a finding in `gate.yaml` meant the check failed, and "finding" now means two things. It is the cost
section 5's clause was always going to charge somebody, and it is charged to the reader.

**What stops happening.** A P0 whose declared scope resolves over its own budget no longer stops the
intent. Today it turns G-Schema red, `predecessorAllowsStart` refuses P1, and the budget cannot be
revised because the scope is inside P0's `artifacts_hash` — so the only ways out are an approval or
starting the intent over. XENO-0245 kept clear of this by giving its budget headroom and recorded
the workaround as one.

**What does not start happening.** The check still reads nothing at P0, because no P0 lock records a
`files` list (#267). So the clause is honoured and the finding it describes is still not produced by
this repository. Both have to close.

**Nothing in the trail moves.** `omitempty` means no existing `gate.yaml` changes when it is
rewritten, and `gate verify` reports the same verdict count. The four check results, the five phase
statuses and `hashing.FindingID` are untouched.

**`docs/clause-readers.md` gains a row.** Section 5's budget clause has never had one, which is a gap
this intent is in a position to notice because it is the intent that gives the clause a reader. The
count goes from forty to forty-one, and the paragraph that carries the count says so — it is the
figure that was already wrong once today.

**Adopters get it with the runner.** `internal/gates` and `internal/model` ship in the binary, so a
project on the next release finds its budget advisory with no action of its own. A project that has
declared no budget sees nothing, which is all of them but this one.
