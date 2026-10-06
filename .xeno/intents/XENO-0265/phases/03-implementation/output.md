---
intent: github.com/triplem/xeno#235
phase: 03-implementation
created: "2026-10-06T18:00:45Z"
schema_version: "1.0"
runner_version: dev+5276f4b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 33d64ca462a0f7ab409609c7d2b9b2163cab59df173e173a7d638cd123a4254e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

Three files. One field, two functions, one helper and a row.

**`internal/model/model.go`.** `Finding` gains `Advisory bool` with `omitempty`, and a comment
carrying both halves of section 5's clause — what it means, and the bound that no code can enforce.
It also records why the field is absent from `hashing.FindingID`: ids do not depend on the run, and
an id that moved when the flag moved would not be stable.

**`internal/gates/gates.go`, `result`.** A check fails where any finding is not advisory, rather than
where there is any finding. It is the only constructor of a check and `schema` appends from six
checks before calling it, so nothing upstream knows which finding came from where and this has them
all in front of it.

**`internal/gates/gates.go`, `Status`.** An advisory finding with no decision no longer counts as
undecided. Section 5 says two things — "the check carrying it is `pass`, the phase is `green`" — and
they have two readers; `result` answers the first and this answers the second.

**`internal/gates/gates.go`, `advisory` and `budget`.** A one-line helper that marks a finding, and
both budget findings passed through it. The helper exists so that there is one place to count, which
is what the test of the bound reads.

**`docs/clause-readers.md`.** A row for section 5's budget clause, which the table has never carried,
naming `gates.advisory`, `result` and `Status` — because violating the clause means any of the three
failing to do its part. The count goes to forty-one.

## The gap the end-to-end run found, after `result` was already changed

With `result` reading the field and `Status` not, a scratch phase whose only finding was a budget
overrun came out:

    00-intake: red
      G-Schema       pass
          F-2af1d8 ...: the recorded context is 9999 bytes and the budget is 10

Every check passing and the phase red. `Status` counts undecided findings and an advisory one is
undecided for ever, because nothing asks anybody to decide it. Half the clause honoured, and the half
that was missing is the half that stops work.

With both changed, the same phase:

    00-intake: green
    01-requirements started
    gate approve F-2af1d8 -> 00-intake: approved

Green, the next phase starts, and the finding is still decidable — which is section 5's sentence in
three parts.

## What was found about the tests that let this survive

Every existing budget test calls `budget` directly and reads its findings. **None asserts what the
check result should be**, which is the one thing section 5 is explicit about. So the behaviour that
contradicted the specification had no reader in the suite either, and the new tests go through
`result` and `Status` rather than through `budget`.

<!-- xeno:section:deviations -->
## Deviations from the design

**`Status` was not in P2's decisions and had to change.** P2 named `model.Finding`, `result` and
`budget`, and said the clause would be honoured by the first two. It is not: "the check carrying it
is `pass`, the phase is `green`" is two statements with two readers, and `Status` derives the phase
from undecided findings without consulting the check result. The gap was found by running the thing
rather than by reading it, after `result` alone had been changed and the scratch phase came out red
with every check passing.

**P1's criterion 7 is what caught it**, and it is the only criterion in this intent that is about a
behaviour rather than a value. "A phase whose only finding is a budget overrun comes out `green`, end
to end, and the next phase is allowed to start." A criterion written as "the check result is pass"
would have passed over the fault.

**An advisory finding is undecided for ever, and that is now load-bearing.** Nothing asks anybody to
decide one, so `Status` has to skip it rather than wait for it. The consequence is that `approved` is
reachable for an advisory finding only if somebody chooses to approve it — which section 5 permits
and the test covers — and is never required.

**The bound is tested by counting occurrences of a string in a source file.** `strings.Count` over
`gates.go` for `advisory(finding(` and for `Advisory = true`. It is a weak reader and an unusual one:
it will fail if the helper is renamed or the call is reformatted, which is a false red for a true
property. It is in because the alternative was nothing at all — no code can tell a clause that
legitimately asked for the field from a check somebody found inconvenient — and because the moment
worth catching is a second check reaching for it, which it does catch.

**The end-to-end fixture writes a lock by hand.** The budget check has no input in this repository:
no P0 `context.lock.yaml` records a `files` list, which is #267. So the scratch intent's lock was
given one file with a size, which is a state the runner cannot currently reach on its own. The
evidence says so rather than presenting the fixture as a run.
