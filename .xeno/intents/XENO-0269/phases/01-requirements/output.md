---
intent: github.com/triplem/xeno#207
phase: 01-requirements
created: "2026-10-07T09:05:30Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ba89364d4ddd823d535bfe7d98eb4eb3728b8dfe0078ad56faf4c548f519e161
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

Numbered, and P4's mapping cites the numbers.

1. **Section 12 says the triple is a declaration and not a measurement**, in the paragraph
   after the one that calls it the raw material for a provider register, so a reader who has
   just been told what the three fields are for learns in the same place what they are worth.

2. **It says where each of the three comes from**, and each is what the code does:
   `tool_version` from `--tool-version` or `XENO_HARNESS_VERSION`, `model` from
   `project.yaml`'s default, `tool` from `project.yaml` unless `XENO_HARNESS` says otherwise.

3. **It says what G-Schema does with them**: presence and shape, and nothing more. That is
   `sessionFields` and the absence of these fields from anything else in the gate.

4. **It says why the runner cannot corroborate one**, naming section 7's rule rather than
   asserting an impossibility.

5. **It says why a check against the second copy is refused**, so the thing #207 warns against
   is written down as refused rather than merely not built.

6. **It names the gateway's record and what it is worth**, and why it is out of reach: the
   runner makes no model request in v1, and the gate path never touches the network. The
   plan's own sentence about a routing record rather than an attestation is what it rests on.

7. **It ends by saying how the register is read.** That is the sentence #207 asked for in as
   many words: the register "says what an agent claimed rather than what ran", and section 12
   now says so itself.

8. **No normative word in the paragraph asks a tool for anything.** What it states is an
   absence, so nothing can violate it by writing the wrong thing, which is what makes it an
   explanation in `docs/clause-readers.md`'s four kinds rather than a tool requirement.

9. **`docs/clause-readers.md` moves its explanation count from 129 to 130** and carries a
   paragraph saying which clause moved it and why it has no row, in the shape #267's three
   clauses were given in the same document.

10. **`docs/assumptions.md` carries one row** with the decision, who took it and when, the
    option not chosen with its cost, and `plugin_version` named as the contrast.

11. **Nothing else changes.** No gate, no field, no artifact, no Go file. `go test ./...`,
    `gofmt -l .`, `go vet ./...` and `xeno gate verify` pass, and no verdict in the trail
    moves.

12. **The specification change is its own commit and comes first**, with the exception to the
    first standing rule named in its message, and with the correction to the drafted wording
    named there too, because the draft was approved with an error in it.

<!-- xeno:section:non-goals -->
## Non goals

**No check against `project.yaml`.** The second copy of `tool` and `model` is the same
declaration in another file, so a green result would mean the two copies agree. #207 names this
as the thing that should not happen, in its own last sentence, and the reason is that it would
make the gap invisible instead of closing it. Refused, and the paragraph says it is refused
rather than leaving the next reader to invent it.

**No gateway comparison.** The one corroboration that would be worth something. It needs a
gateway, which section 12 puts in v2, and it could only ever be a figure outside the gates,
because `xeno gate ...` never touches the network and the runner makes no model request in v1.
Put to the maintainer with that cost and not chosen; this intent does not assign it to a
package either, because assigning work is the plan's and not an intent's.

**No new gate result and no finding for an unverified field.** A result outside the four that
A4 and A42 fix would be a specification change of its own, and a finding that fired on every
artifact forever is #235's shape, answered there.

**No field removed and none added.** Section 12 requires the three and a project with a
register needs them. What was wrong is what a reader took them for.

**No change to `plugin_version`.** It is named in the register row as the contrast, because it
is the field that had exactly this shape and now has an anchor the checked side cannot
influence. Naming it is what makes the row legible; changing it is a different subject.

**No claim that the triple is wrong.** Nothing here says a recorded value is false. The claim
is narrower and that is deliberate: nobody has found a wrong value, a harness that lies about
its version is not a threat model this project has, and #207 says so.

**No change to section 7.** The rule against branching on the harness is the reason
corroboration is unavailable, and it is cited rather than revisited. A paragraph that reopened
it would be arguing with the clause it depends on.

<!-- xeno:section:constraints -->
## Constraints

**The first standing rule, and the exception.** `docs/process-definition.md` is not editable by
the agent. The three options were put to the maintainer with their consequences and costs on
2026-10-07, the choice was theirs, the wording was then drafted and put to them as a separate
question, and "write it as drafted" is what was acted on. The exception is named in the
specification commit's message.

**A negative result is evidence only when the thing checked was there to be found.** This whole
paragraph is a negative — that nothing corroborates the triple — so the three writers and
G-Schema's involvement were read in the files rather than recalled, which is what caught the
error in the approved draft.

**What is sealed is never rewritten.** No verdict and no artifact of another intent is touched,
and the paragraph cannot move a figure anywhere in the trail: the documents are read by context
scopes and compared against what each lock recorded, not against today's tree.

**Prose wraps at 88 characters**, tables and code blocks do not, and a paragraph changed a
second time is replaced rather than edited into. The paragraph was replaced twice — once to
correct `model`, once to drop an em dash the document uses exactly once in 132 kilobytes — and
read back whole each time rather than as a diff.

**Headings name their section in words.** The paragraph is a bold lead-in in the register
section 12 already uses, not a numbered subsection.

**No invented fields, gates, tools or rules.** Nothing is added to an artifact, to the gate list
or to the MCP surface.

**One dependency.** Untouched; no Go file is in the diff.

**The specification commit comes first**, before the commit carrying the count and the row.

**Every change belongs to a work package and to an intent.** WP11, intent XENO-0269 for issue
#207. The key was passed explicitly, because the sequence counts from what this branch can see
and two other intents are open and unmerged.
