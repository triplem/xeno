---
intent: github.com/triplem/xeno#258
phase: 02-design
created: "2026-10-06T17:31:01Z"
schema_version: "1.0"
runner_version: dev+d3983d3
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ab4398bbb32877b3807dd47220b00fdb56e761287e85a7b7a48adf0cf701828a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: 'One clause, one gate: G-Schema reads the numbered list on a P1 declaring requirements@1.1.0 or later, G-Test reads the mapping''s completeness on a P4 declaring verification@1.1.0 or later, both keyed on the declared version with a numeric comparison. Reason on model.Option is required by QuestionAsked, the writer only. The mapping accepts a criterion''s number as a standalone token anywhere in the section.'
      rationale: 'The gate is chosen by the section that states the requirement: section 5 is about an artifact''s shape, which G-Schema reads, and section 7 asks G-Test for the mapping. Putting both in G-Test would make a green G-Test a statement about an artifact two phases away and leave a P1 unjudged until it was sealed. The reason stays in the writer because one sealed question in this trail already fails the recommendation requirement, so a gate reading either half would re-judge the trail — #229''s measurement, unchanged. The loose citation token can produce a false green and never a false red, which is the right way round when a false red stops a phase. Nothing stricter than a numbered list is required, because consecutiveness and a lower bound are not in section 5 and the second standing rule makes each an invention. The cost accepted and measured: bumping the two templates removes the strings_hash reader from 130 sealed artifacts, red to green on a corrupted value, and cannot be avoided without guessing — though tampering is still caught by artifacts_hash and the successor''s freshness, which exits 1.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**One clause, one gate, chosen by the section that states the requirement.** Section 5 says the
`acceptance-criteria` section is a numbered list, which is a statement about an artifact's shape, and
G-Schema is what reads shape. Section 7 asks G-Test for "mapping of acceptance criteria complete",
so the completeness lands there. The alternative was both in G-Test, which would make a green G-Test
mean something about an artifact two phases away.

**Both checks key on the version the artifact declares, compared numerically.** `Ref()` is
`id@version`; the comparison parses the three parts and compares them as integers, because a string
compare puts `1.10.0` before `1.9.0` and the check would stop applying at the tenth minor version
with nothing to say so.

**The mapping check accepts a number as a standalone token anywhere in the section.** Section 5 says
the mapping "names each criterion by its number" and does not say how. A table cell, a list item and
a sentence all name it. A stricter reading would fail a format the clause permits, which is A44's
objection; the cost is that a stray number in prose satisfies the check, so it can produce a false
green and never a false red. Given that a false red here stops a phase, that is the right way round.

**Nothing is required of the numbering beyond being numbered.** Not consecutiveness, not starting at
one, not a bound. Each is an invention under the second standing rule and each would fail a
legitimate artifact for a rule nobody wrote.

**`Reason` is required by the writer and not by the gate.** `QuestionAsked` gains it;
`QuestionShape`, which the gate calls, does not. One sealed question in this trail already fails the
recommendation requirement, so a gate reading either half would re-judge the trail — the measurement
is in `QuestionShape`'s comment and #229 settled it. The new row in `docs/clause-readers.md`
therefore says "`QuestionAsked`, the writer only", in the same words the row above it uses.

**The mapping check does nothing where the P1 artifact is missing or older.** A P4 at 1.1.0 behind a
P1 at 1.0.0 is legitimate — it is every intent already under way when the bump lands — and reading
it as a failure would make the bump re-judge exactly what the anchor exists to protect.

**`testReport`'s comment is replaced rather than edited.** It has two statements that are now wrong,
the section number and the figure, and a paragraph with two corrections in it is the case the
convention is about. What replaces it says what the check does and what it costs.

**The cost of the bump is written into the code, beside `goneBundle`.** Not only into this intent's
artifacts, because the trail is not where somebody wondering why a sealed artifact stopped being
checked will look. A reader has the current tree and nothing else.

**Only the two templates the clauses name are bumped.** `intake`, `design`, `implementation` and
`review` stay at 1.0.0: a bump costs each of them the `strings_hash` reader on every artifact that
declares the old version, and no clause asks anything of their sections.

<!-- xeno:section:alternatives -->
## Alternatives

**Put both checks in G-Test.** One gate, one place to read, and the completeness check already needs
the P1 artifact so it has both files open. It makes a green G-Test a statement about an artifact two
phases back, and it leaves a P1 that declares 1.1.0 and numbers nothing unjudged until P4 — three
phases and possibly days later, when the artifact is sealed. Rejected: the phase that can still fix
it is the phase that should fail.

**Put the numbered-list check in the writer, as `section set` refusing an unnumbered
`acceptance-criteria`.** It is where #229 put the recommendation check and it catches the fault at
the keystroke. The section is written once and read by a gate many times, and a writer check would
not reach an artifact written by hand — which section 5 explicitly allows, with G-Schema named as the
backstop. Rejected, and it is the better half of a both-and that this intent did not have room for.

**Require the criteria to be consecutive from one.** It is what every artifact in this trail does and
it would make the mapping check exact rather than generous. Section 5 says "a numbered list" and
nothing more, so the second standing rule makes it an invention, and an artifact numbering 1, 2, 2a
would go red for a rule nobody wrote. Rejected.

**Require a citation format for the mapping — a table whose first column is the number.** It is what
every P4 artifact in this trail that cites numbers does, and it would let the check be exact. It
would also fail prose that names criterion three in a sentence, which section 5 permits, and A44's
objection is precisely to a check that fires where nothing is wrong. Rejected in favour of the loose
token, with the false-green cost named.

**Compare the declared version against the current template's rather than against a literal
1.1.0.** Then the check applies whenever the artifact is at the current version, with no constant to
maintain. It would stop applying the day either template is bumped again, silently, which is the
same class of fault as the string compare. Rejected: the clause names a version and the code should
too.

**Do not bump the templates; make the checks apply to everything and approve the trail.** It is the
route #258 weighed and declined at 57 verdicts, and the template bump exists to avoid it. Rejected
as a decision already taken, and named because it is the only route that costs no `strings_hash`
reader.

**Bump the templates and also fix `goneBundle` so the reader survives.** The attractive one. It
cannot be done honestly: when a recorded hash differs from the current bundle's, there is no way to
tell a wrong hash from a right one about a bundle that has been replaced, and the two need opposite
answers. Making it pass when it matches adds nothing, since a match is already not a finding.
Rejected as unachievable rather than unwanted, and recorded so the next reader does not spend the
afternoon on it.

**Put `Reason` on `Question` after all.** It would need no conditional requirement and no empty field
on three options out of four. Section 8 now says it belongs to the option, so this is settled by the
commit that preceded this intent. Rejected on the first standing rule.

<!-- xeno:section:impact -->
## Impact

**`internal/model/model.go`.** `Option` gains `Reason string` with `omitempty`. Nothing else.

**`internal/gates/gates.go`.** `QuestionAsked` gains the reason requirement. `schema` gains the
numbered-criteria check for P1. `testReport` gains the mapping check for P4, and its comment is
replaced. A version comparison helper, and a comment beside `goneBundle`'s consequence recording what
a bump costs.

**`.xeno/plugin/templates/requirements/template.yaml` and `.../verification/template.yaml`.**
`version: 1.0.0` becomes `1.1.0`. One line each.

**What changes for the 130 sealed P1 and P4 artifacts.** They declare 1.0.0, so neither new check
applies and no verdict moves — and they lose the `strings_hash` recomputation, because `goneBundle`
becomes true for each of them. Measured: a corrupted `strings_hash` goes from red to green. Editing
one is still caught, by `artifacts_hash` and the successor's G-Freshness, which exits 1.

**What changes for the next intent.** Its P1 must number its criteria or G-Schema fails it, and its
P4 mapping must name each of those numbers or G-Test fails it. That is the convention arriving, and
it arrives for this intent first.

**What changes for an intent already under way when this lands.** Its P1 is sealed at 1.0.0, so its
P4 at 1.1.0 finds an older predecessor and the mapping check does nothing. No intent is caught
mid-flight.

**What changes for an adopter.** Both: the templates ship in the plugin and the checks ship in the
binary, so a project that vendors the next release gets the convention and the gates together. A
project with its own `requirements` template under `.xeno/config/templates/` keeps whatever version
it declares, and the checks apply from 1.1.0 of the id — so an override at 1.0.0 opts out, which is
the per-id override section 5 already describes.

**The plugin digest moves.** Two template files change, so `plugin.sha256` in a new
`context.lock.yaml` differs from the one sealed locks record. G-Supply is `not-implemented`, so
nothing compares them today; when it is implemented it will have to read the lock's own recorded
value rather than the current tree's, which is what section 5 says it does.

**Two rows of `docs/clause-readers.md` move from no reader to a reader**, and the two paragraphs
XENO-0260 wrote under the table change from "waits on the commit" to naming what reads each. That is
the third replacement of that passage, which XENO-0260 predicted.
