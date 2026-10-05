---
intent: github.com/triplem/xeno#243
phase: 01-requirements
created: "2026-10-05T12:35:22Z"
schema_version: "1.0"
runner_version: dev+f2f65d5
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1e80bc73e032e2cec539e030d3e6a0faa17d2605a27d64f785af749d9a149d82
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. `docs/assumptions.md`'s opening no longer says the record is closed, that A77 is the last
   row, or that nothing is added here. No sentence in the file claims the register is closed.

2. The opening says the register is open, and for what: assumptions and decisions about this
   repository's construction.

3. The opening says what did close at M0, in the plan's own terms — the record kept as a file
   in the branch rather than an artifact under `.xeno/` — and that this is section 4's
   sentence, so the inference #174 made is visible as an inference rather than repeated.

4. The opening keeps the rule that a learning never goes here, routing through the phase's
   `learning.yaml` and a merge request against the rule set, as section 10 describes.

5. The opening says what distinguishes a row here from a decision in an intent's design
   phase: a decision sealed inside one intent belongs to that intent, and a row here is for
   what outlives the intent that found it.

6. The paragraph framing the whole file as the pre-M0 record is corrected, since it is the
   same claim in the fourth paragraph and would otherwise survive the banner's removal.

7. The paragraph saying every row stays with its state column is unchanged, and no existing
   row is edited, renumbered or deleted.

8. A95 records this correction, with the misreading as its reason, and is the last row.

9. Every paragraph changed is replaced whole and read back as prose rather than edited into,
   and Markdown prose stays within 88 columns.

10. `CLAUDE.md` is checked against the new opening and changed only if it disagrees. It names
    this file as where assumptions and decisions are written down, which the correction makes
    true rather than false, so the expected outcome is no change.

11. `./xeno gate verify` exits 0 with the 363 verdicts that exist now intact, growing only by
    this intent's own judged phases. `go build`, `go test ./...`, `go vet ./...` pass and
    `gofmt -l` outside `vendor/` prints nothing, none of which this change can affect.

12. One commit, `Closes #243`, and the issue carries `wp5`.

<!-- xeno:section:non-goals -->
## Non goals

Not a judgement on the seventeen rows. Whether A78 through A94 were each worth a row is
seventeen questions and this intent asks none of them; it settles where such a row belongs.

Not a deletion, renumbering or rewrite of any row. The file's third paragraph says every row
stays with its state column and gives the reason, which this intent leaves standing.

Not a change to either normative document. The finding is that #174 misread section 4, not
that section 4 is wrong, so the plan keeps its sentence and no specification commit precedes
this.

Not a gate, a checker or a rule. Nothing will compare this repository against the register's
opening, and A90's finding is why that is stated rather than fixed: a reader which cannot fail
is worse than none, and a rule about where a decision belongs is one a person reads.

Not a rule about when a construction choice earns a row. The new opening says what the
register is for and what belongs to an intent instead; how significant a choice must be to
earn a row is left to the person writing it, as it was for the first seventy-seven.

Not a reopening of A77's closure of the hand held record. That record did close at M0 and the
new opening says so; what changes is which file the sentence was about.

Not a migration. The rows written while the banner said otherwise are not marked, annotated or
distinguished, because the correction says they were right all along, and labelling them would
imply the opposite.

<!-- xeno:section:constraints -->
## Constraints

The first standing rule binds the plan. Section 4's sentence stays exactly as it is, and this
intent's whole claim is about what it means, so the correction has to quote it rather than
paraphrase it; a paraphrase is how the banner went wrong.

The conventions bind the editing method. Each paragraph being changed has been changed before,
so it is replaced whole and read back as a paragraph rather than edited into, and prose stays
within 88 columns. A sentence patched in place is exactly where the words around it get left
behind.

`CLAUDE.md` is sent with every request of every session, which is why it is in scope even
though it is expected not to change. A register whose opening says one thing and whose pointer
in `CLAUDE.md` says another is how a banner survived twelve contradictions, and the cheap
guard is to read both at once.

A95 is the next number and the sequence is not renumbered. The numbers reach A94 and the gap
that would open by skipping one would read as a withdrawn row, which the file has a
convention for and this is not.

The register is not covered by `artifacts_hash`. Nothing seals it, no verdict depends on its
bytes, and the 363 verdicts are therefore indifferent to this change — which is the same
property that let it drift, and the reason the correction is prose rather than a mechanism.

One commit, `Closes #243`, on a branch carrying one intent, with the issue labelled `wp5`.
