---
intent: github.com/triplem/xeno#277
phase: 01-requirements
created: "2026-10-07T13:28:20Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5973737622e49269831018c24f084e3145350abeea617e0ff79f83d3c1acba79
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

1. **`honest` no longer reads the `tool` field.** The `byHand` variable is gone and the
   expression keeps the two terms that are facts about the tree.

2. **No gate reads `model`, `tool` or `tool_version` to decide anything.** `tool` was the only
   one that did, and a grep over the gates establishes that none remains.

3. **The comment above `honest` says what was removed and why.** The reason the term was there
   is sound and only its premise was wrong, so the comment records both rather than leaving a
   reader to wonder why the obvious case is missing.

4. **`hashShape`'s doc comment describes the behaviour that exists.** It no longer says the
   placeholder is honest where the artifact says it was manual, and it says what the two
   remaining cases have in common.

5. **The test case that asserted the old behaviour asserts the new one**, with a comment saying
   why the field is not evidence, and the five other cases of `TestHashFieldShape` are
   unchanged and still pass.

6. **Exactly two phases change verdict**, XENO-1's and XENO-2's intakes, on exactly four
   findings, all `context_hash says by-hand where a writer exists`. No other phase in the trail
   moves.

7. **Both phases read `approved` rather than red**, each finding released by the maintainer with
   a reason naming the artifact's date, that the lock beside it is hashable today, and that
   section 11 forbids rewriting a sealed artifact.

8. **The releases are approvals and not overrides**, so `xeno intent status` lists no obligation
   owed for them.

9. **`xeno gate verify` exits zero**, which it did not between the code change and the
   approvals, and that is the thing CI runs.

10. **No artifact's content is edited.** The four files keep `context_hash: by-hand`; what
    changed is the verdict over them and the decision recorded in their `gate.yaml`.

11. **No normative document changes.**

12. **`docs/assumptions.md` carries one row** with the decision, the measurement, the four
    approvals and the date.

13. **`go test ./...`, `gofmt -l .`, `go vet ./...` and `go build` all pass.**

<!-- xeno:section:non-goals -->
## Non goals

**No change to the process definition.** Section 12's paragraph on the triple says nothing
corroborates it, which stays true and is the clause this change acts on. A sentence recording
that a gate used to act on the triple would be history in a normative document, and the
register is where history goes.

**The four artifacts are not corrected.** Writing the real `context_hash` into them would
satisfy the check and rewrite two sealed phases, which section 11 forbids in as many words.
That is why the findings are released, and the reason recorded with each release says so.

**No `gate override`.** An override carries an obligation and `xeno intent status` would list
four `obligation close` calls owed on artifacts nobody can correct. "Later, whoever gets to it"
has no later here, so the release is an acceptance.

**Nothing about `model` or `tool_version`.** Neither is read by any gate. `tool` was the only
one of the three that decided anything, which is what made this a change rather than a clause.

**`goneBundle` stays.** It is a fact about the tree: the shipped set either answers the
recorded ref or it does not. It is also what keeps the two pre-M0 artifacts' `strings_hash`
honest without their declaring anything, and taking it out would turn the trail red for a
reason nobody has argued for.

**The `tool` field is not removed or made optional.** Section 12 requires it and a project with
a provider register needs it. What changes is that no verdict depends on one.

**No new gate, finding kind or result.** The four findings are the existing
`context_hash says by-hand where a writer exists`, produced by the existing code path.

**No check that an artifact's `tool` matches anything.** #207 refused that in the document and
the reason holds: the only other copy is `project.yaml`, which is where the runner read it, so
a comparison would hide the gap rather than close it.

**Nothing is done about the two pre-M0 intakes beyond releasing the findings.** They stay in
the trail with an approved verdict and a reason, which is what the process does with a finding
that is true and cannot be fixed.

<!-- xeno:section:constraints -->
## Constraints

**What is sealed is never rewritten.** Section 11. It is why the four findings are released
rather than fixed, and it is also the constraint this change bumps against hardest: two
`gate.yaml` files are rewritten, which is the process's own mechanism for a verdict that
changes rather than an exception to the rule. The artifacts the verdicts are about are
untouched.

**A release is a statement by a second person.** The runner refuses to offer `gate approve` as
a command for exactly this reason, and the suggestion machinery's comment says so. The four
reasons were drafted and put to the maintainer, who chose the shape and approved the wording on
2026-10-07; the commands were then run with `--by` naming them.

**A negative result is evidence only when the thing checked was there to be found.** The claim
that nothing else in the trail moves is an absence, so it is established by running
`gate verify` over the whole trail before and after and by naming which other terms cover the
other fields, not by inspection.

**The first standing rule.** `docs/process-definition.md` is not the agent's to edit and is not
edited. `docs/assumptions.md` is the open register.

**A paragraph changed a second time is replaced rather than edited into.** `hashShape`'s doc
comment is in that state, so it is rewritten whole and read back as a paragraph.

**No invented fields, gates, tools or rules.** The change removes a term. Nothing is added to
an artifact, to the gate list or to the finding set.

**One dependency.** Untouched.

**Every change belongs to a work package and to an intent.** WP1 by subject, since the artifact
schema and the core gates are where `hashes` lives, and intent XENO-0272 for issue #277. The
key was passed explicitly, because three intents are open on other branches and the sequence
counts from this one.
