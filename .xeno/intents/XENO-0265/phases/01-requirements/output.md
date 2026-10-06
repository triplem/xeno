---
intent: github.com/triplem/xeno#235
phase: 01-requirements
created: "2026-10-06T17:52:01Z"
schema_version: "1.0"
runner_version: dev+5276f4b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9fadfaab33a7554b87ba785dfd90e94945c6681c56a733e53dba1566cc267753
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Numbered, and P4's mapping cites the numbers — which `gates.mappingComplete` now reads, because this
intent's P1 is the first in the trail to render at `requirements@1.1.0`.

1. **`model.Finding` carries `Advisory bool` with `omitempty`.** No artifact already written changes
   shape and nothing in the trail re-hashes.

2. **`result` returns `fail` only where a finding is not advisory.** A check carrying nothing but
   advisory findings is `pass`.

3. **A check carrying one advisory finding and one ordinary one still fails.** The ordinary one is
   what fails it, and the advisory one rides along — which is the whole of the distinction and the
   case a count of findings cannot express.

4. **The advisory findings are in `gate.yaml` with their id, cause and remedy.** That is what section
   5 means by "visible": not a log line, an entry in the verdict like any other.

5. **`budget` marks both of its findings advisory**, the file count and the byte total.

6. **Nothing else in the runner sets the field.** Section 5: "the clause that asks for it says so
   where the check is described, and nothing else writes the field." A test asserts the count of
   writers, so a later check reaching for it fails rather than passes.

7. **A phase whose only finding is a budget overrun comes out `green`**, end to end, and the next
   phase is allowed to start. That is what "not a red gate in the sense of stopping work" means in
   this runner — `predecessorAllowsStart` — and it is a behaviour, not a field.

8. **An advisory finding is still decidable.** `gate approve` takes it by id, as section 5 says it
   can be decided like any other.

9. **The four check results are untouched and the five phase statuses are untouched.** No reader of a
   verdict learns a new state; a phase whose only findings are advisory is green by the existing
   derivation reading a check that passes.

10. **`docs/clause-readers.md` gains a row for section 5's budget clause**, which it has never had,
    naming what would fail if the clause were violated.

11. **The bound is in the code where the field is defined**, carrying section 5's sentence, so that a
    reader who finds the field knows it is an exception before they use it.

12. **`go test ./...` passes, `go vet` is clean, `gofmt -l` prints nothing, `xeno gate verify` exits
    0 at the same verdict count.** `omitempty` is what makes the count the check.

13. **Every new assertion is confirmed able to fail**, by mutation, as in the two intents before this.

<!-- xeno:section:non-goals -->
## Non goals

**Not #267.** `budget` reads nothing at P0 because no P0 `context.lock.yaml` records a `files` list.
That is a change to when `phase start` writes the lock, against #215's reasons for writing it once,
and it is the other wp8 issue. An advisory finding that is never produced is still never produced,
and this intent closes one of the two faults rather than pretending to close both.

**Not making any other check advisory.** Which others should be is a question per check. Section 5's
second paragraph bounds the field to the clause that asks for it, and answering the general question
here is how a bounded exception becomes a default.

**Not `drift`.** It stays unimplemented and section 16's ninth limitation stays unread. XENO-0263
recorded why it was not the route — built for its second purpose before its first, and a row carrying
sha256 hashes where a budget overrun is two byte counts — and this intent does not revisit it.

**Not a fifth check result.** `advisory` is a property of a finding. A4 and A42 fixed the result set
and the decision was taken partly to leave it alone.

**Not a sixth phase status.** A phase whose only findings are advisory is `green` because the check
passed, which is the existing derivation. Nothing is added to section 5's five.

**Not a change to how a finding is decided.** `gate approve` and `gate override` take a finding by
id and neither needs to know about the field. An advisory finding can be approved, which section 5
asks for and which costs nothing to leave alone.

**Not a warning count, a summary line or a new output.** The findings print as findings do. A reader
of `xeno gate run` sees them under a check that passed, which is the shape section 5 describes, and
inventing a second presentation would be a new form nobody wrote.

**Not a change to any normative document.** The clause was committed in XENO-0263 and this is the
code that follows it, which is the order the first standing rule asks for.

<!-- xeno:section:constraints -->
## Constraints

**`result` is the only constructor of a check, and three gates call it with findings from several
sources.** `schema` appends from six checks before calling it, so the change cannot be "G-Schema
ignores budget findings"; it has to be a property of the finding that `result` reads. That is also
what makes criterion 3 the real test: a check carrying one of each must fail.

**The field can silence any check, and only a sentence limits it.** A90's objection in its strongest
form — a reader that cannot fail is worse than none, and this is the switch that makes one. Section
5 bounds it in prose; the code can carry the bound as a comment and a test over the number of
writers, and that is the most either can do.

**"Not a red gate in the sense of stopping work" is a behaviour in `predecessorAllowsStart`.** A red
predecessor refuses the next phase. So the clause is satisfied only if a phase whose sole finding is
the budget overrun is green and the next phase starts, and that has to be run rather than inferred
from a check result.

**`omitempty` is what keeps the trail still.** A sealed artifact's `gate.yaml` is not inside
`artifacts_hash` — `hashing.PhaseExcluded` leaves it out — but a field that serialised always would
rewrite 465 verdicts' files on the next run and make every diff unreadable. The verdict count before
and after is the check that it did not.

**The budget check has no input in this repository.** 0 of 117 P0 locks record a `files` list, so
nothing here can demonstrate the finding arising from a real scope. The end-to-end check therefore
has to construct a lock that does record one, which is a fixture standing in for a state the runner
cannot currently reach — and saying so is part of the evidence rather than a footnote to it.

**An advisory finding still needs an id.** `hashing.FindingID` is a hash over gate, rule id, file and
cause, and section 5 says finding ids do not depend on the run. Nothing about the field may enter
that hash, or every budget finding's id would change the day it became advisory.

**This intent's own P1 is the first artifact judged by `numberedCriteria`.** The template bump landed
in #270, so this P1 renders at `requirements@1.1.0` and its criteria must be numbered or G-Schema
fails it; its P4 mapping must name each number or G-Test fails it. The checks written one intent ago
meet their first real subject here.
