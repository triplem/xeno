---
intent: github.com/triplem/xeno#235
phase: 02-design
created: "2026-10-06T15:21:43Z"
schema_version: "1.0"
runner_version: dev+8decfdc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bbdd77200932d2b574c601d3bca23e907ba251c599e119991315df1883bca2d7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: 'Section 5 gains an advisory key on a finding: advisory: true in the gate.yaml findings block, a paragraph saying the check carrying it is pass and the phase green while the finding keeps its id, cause and remedy, a second paragraph bounding the exception, and one sentence in Context economy saying the budget''s finding is one.'
      rationale: 'It is what section 5''s own sentence describes — deliberately a finding and not a red gate, visible, decidable like any other — and it leaves the four check results alone, which A4 and A42 fixed. drift was the one route using a mechanism the specification already enumerates and was rejected on measurement: it is unimplemented, with no Drift field, no writer in internal/ or cmd/ and 0 of 443 sealed gates carrying one, so it would be built for its second purpose before its first, and its row carries sha256 hashes where a budget overrun is two byte counts. A fifth check result would put the property on the check rather than on the thing that is advisory. The bound is a paragraph of its own because no code can tell a clause that legitimately asked to be advisory from a check somebody found inconvenient. The agent writes it on the person''s instruction, which the first standing rule makes an exception and this records as one.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**`advisory` is a key on a finding, not a result on a check.** It is the shape section 5's sentence
describes — "deliberately a finding and not a red gate" — and it leaves the four check results
alone, which A4 and A42 fixed and which every reader of a check handles. Putting it on the check
would also attach the property to the wrong thing: a second advisory finding in G-Schema would drag
the whole check with it, and G-Schema carries six checks' worth of findings.

**A boolean that is absent in the ordinary case, not a severity.** `advisory: true` and nothing
else. A severity scale would ask every existing finding to declare one and would invite an argument
about where the line sits; a boolean says only that this one clause asked not to block, which is
what section 5 asks for and all it asks for.

**The paragraph states the three things that make it visible, rather than asserting visibility.**
The check is `pass`, the phase is `green`, and the finding is in `gate.yaml` with its id, its cause
and its remedy. Section 5 already says "it is visible, it can be decided like any other finding";
repeating that would clarify nothing, and naming what is true of the artifact is what a reader can
check.

**The bound is its own paragraph.** "It is the exception and stays one." A field that lets a finding
not fail is the most useful thing in this runner for anybody who finds a check inconvenient, and the
only thing that will ever limit it is a sentence, because no code will. A bound tucked into the end
of the defining paragraph reads as a caveat; a paragraph of its own reads as a rule.

**It goes after the four-values paragraph, not after the one the draft named.** The status derivation
is two paragraphs and the second completes the first. The draft said "after 'Decisions sit on
findings, not on phases'", which taken literally splits an argument in half. This is the third time
in two intents that a wording drafted in an issue comment has needed its placement adjusted when put
into the document.

**The Context economy subsection gets a sentence, not a cross-reference.** A reader arriving at the
budget clause is told there that its finding is advisory. A pointer to the gate result subsection
would be correct and would be read by nobody, and the whole fault being repaired is that the clause
and its mechanism have lived forty lines apart.

**`drift` is weighed in the artifacts and untouched in the document.** It is the one route that
would have used a mechanism the specification already enumerates, and it is rejected for reasons
that belong in a trail rather than in a specification: the document does not record what a clause
was not.

**The wording is the wording drafted on #235, with placement adjusted.** It was offered to be
changed freely and the person asked for it as it stood. The prose is committed as written; where the
draft said where to put it, the document decided.

<!-- xeno:section:alternatives -->
## Alternatives

**Implement `drift` and give it a third `field` value.** The only route that uses a mechanism the
specification already enumerates, and section 16's ninth limitation already states the principle.
It is unimplemented — no `Drift` on `model.Gate`, no writer, 0 of 443 sealed gates — so it would
mean building it for its second purpose before its first, leaving rule drift unmarked afterwards.
And the row does not fit: `drift` carries `artifact` and `gate_run` as sha256 hashes where a budget
overrun is two byte counts, which is more specification change than one key, not less. Rejected, and
it is the alternative that reads best from a distance.

**A fifth check result, `advisory`.** The precedent is real: `not-implemented` is already a state
that is neither pass nor fail, and section 5 says what a phase carrying one means. It puts the
property on the check rather than on the thing that is advisory, adds a state every reader of a
result must handle for one check out of fourteen, and would need the status derivation rewritten.
Rejected.

**Change section 5 instead, and let a budget overrun block.** The code is already right and one
paragraph changes. It gives up the reason section 5 gave for itself — "blocking against a number
nobody has experience with yet would be the wrong way round" — and one budget has ever been declared
in this repository, so there is still no experience with the number. Rejected in #235, and it is the
honest option if the field later proves worse than the block.

**A severity on every finding.** `severity: <blocking|advisory>` rather than a boolean absent by
default. More expressive, and it would carry the next case without another amendment. It asks every
existing finding to acquire a field, makes the default explicit where absence already says it, and
opens an argument about where a line sits that this process has not needed. Rejected as a bigger
answer to a smaller question.

**Let the budget finding be approved at P0 every time.** No specification change at all: the
mechanism exists, a person approves it, the phase is `approved` rather than green. It makes a routine
measurement into a governance statement, and A90's objection is nearby — an approval everybody makes
every time is a record nobody reads. Rejected, and it is what the current code effectively demands.

**Say in section 5 that G-Schema does not call `budget`, and move it to its own gate.** The third
shape #235 named. There is no gate whose failure does not bind the sequence, so it would need one,
which is a new gate and section 7's budget. Rejected as the largest of the four for the same result.

**Write the clause generally, as "a check may declare a finding advisory".** Shorter, and it would
not need amending for the next case. It is also how a bounded exception becomes a default, with
nothing but a sentence ever limiting it. Rejected: the paragraph bounding the field is the reason
this amendment is safe to make, and a general clause would delete the bound while sounding tidier.

<!-- xeno:section:impact -->
## Impact

**`docs/process-definition.md`, three places in section 5.** One key in the `gate.yaml` findings
block, two paragraphs after the status derivation, one sentence in the Context economy subsection.
Nothing else in the document changes.

**What becomes buildable.** `Advisory bool` on `model.Finding` with `omitempty`; `result` choosing
`fail` only on a finding that is not advisory; `budget` marking its own; tests that a recorded
context over its budget leaves G-Schema `pass` and the phase green while the finding is present,
has an id, and can be decided. Each was a second-standing-rule violation an hour ago and is now a
step.

**What a reader of a verdict learns that they could not.** That a finding in `gate.yaml` may be one
the phase was not failed for. That is a real cost: "finding" now means two things, and every reader
of a verdict has a distinction to hold. It is the price section 5's sentence was always going to
cost somebody, and it is paid by the reader rather than by the writer.

**Nothing in the trail moves.** The specification is in no `artifacts_hash` and no `rules_hash`, so
every verdict stands and `gate verify` reports the same count before and after. `omitempty` means
that when the field exists, no artifact already written changes either.

**What does not improve.** The budget check still reads nothing at P0, because no P0
`context.lock.yaml` records a `files` list (#267). An advisory finding that is never produced is
still never produced, and the two issues are deliberately separate: this one is about what the
finding does, that one about whether there is anything to find.

**`docs/clause-readers.md` is unchanged and is now one row short.** Section 5's budget clause has no
entry there, which was defensible while the clause had no mechanism and is a gap once it does. The
row belongs to the intent that implements the check, because the column it would carry is the
reader.

**No Go code, no gate, no template, no rule, no field.** Nothing in `cmd/`, `internal/` or
`.xeno/plugin/`, so the suite and `gate verify` assert absence of accident and nothing more.

**Adopters are reached by a document and not by a release.** `docs/process-definition.md` is not
shipped by `xeno init --vendor`. A project tracking the specification sees a key no runner writes
yet, which is the ordinary state of a document that leads its implementation and is the second such
key this session added.
