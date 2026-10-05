---
intent: github.com/triplem/xeno#247
phase: 01-requirements
created: "2026-10-05T19:34:52Z"
schema_version: "1.0"
runner_version: dev+8e3b29b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 34716316feb1b66fb7dd09d911c77220ac15a9e9de58cf26642862cf29f6067c
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

1. `docs/clause-readers.md` gains a row for section 8's reason: the clause as the section words
   it, and a reader column saying there is no field for it and nothing reads it. The
   tool-requirement count moves 38 to 39 and the table counts 39.

2. The mapping row still reads `nothing`, and the paragraph under the table gains a sentence
   saying an example rule exists for a project that wants a person to read it. Naming the example
   in the reader column is the error #254 corrected one row away and this must not repeat it.

3. `CLAUDE.md` gains one paragraph: a decision put to a person is put one at a time, with the
   reason, and with XENO-0243 as the instance. It goes in a section that already exists rather
   than adding a heading, and the file stays short enough that the addition earns its place.

4. `examples/rules/` gains one `review` rule, that every acceptance criterion P1 declared is
   answered in P4's mapping, with `scope: org`, `binding: false`, and a header saying why it is
   an example and how to adopt it — the shape the four beside it have.

5. The example rule names what it cannot do: it is answered once per intent rather than once per
   criterion, which is a different instrument from the gate section 7 asks for and is the honest
   ceiling until a criterion is identifiable.

6. No normative document is touched, and `git diff --stat` shows it: three files changed and one
   added, none of them under `docs/` but the audit.

7. No code changes, and nothing enables the example rule. `rules.Load` reads
   `.xeno/plugin/rules/` and `.xeno/config/rules/`, not `examples/`, so the effective set is
   unchanged — asserted by `gate verify` rather than by reading.

8. `./xeno gate verify` exits 0 with the 405 verdicts that exist now intact, plus this intent's
   own judged phases. The example rule must not reach the effective set, and an unchanged
   `rules_hash` across the trail is what proves it.

9. `go build`, `go test ./...`, `go vet ./...` pass and `gofmt -l` outside `vendor/` prints
   nothing.

10. Markdown prose stays within 88 columns in both documents; tables are exempt.

11. One commit, `Closes #247, closes #248, closes #250`, each issue carrying its own label.

<!-- xeno:section:non-goals -->
## Non goals

Not a field for the recommendation's reason. Adding one to `model.Option` is an addition to what
section 8 enumerates, so a specification change first and a person's commit by the first standing
rule. #247 stays open on that if the maintainer wants the mechanism; what closes is the question
of what reads the clause today.

Not a clause about sequence. Writing "questions are put one at a time" into a normative document
is the same kind of change. The convention goes in `CLAUDE.md`, which the issue itself names as a
candidate home, and section 8 keeps its wording.

Not a numbering convention for acceptance criteria, which is what a mapping gate would need and
is section 9's.

Not a gate for any of the three. A90's finding is why, and the audit row is where that is
recorded rather than argued again.

Not enabling the example rule. Not in `given/builtin/`, which would reach every adopter and be
answered at every review for ever — A72's reasoning — and not in this repository's own
`.xeno/config/rules/` either, because adopting it is a decision about this project's reviews and
not part of writing it.

Not a change to the mapping row's reader column. It says `nothing` and goes on saying it, because
an example nobody enabled fails nothing.

Not a second pass over the audit. #254 corrected one row and said plainly that thirty-six are
unexamined; this adds one and leaves that true.

<!-- xeno:section:constraints -->
## Constraints

`CLAUDE.md` asks not to be added to. Its last line is "This file is sent with every request of
every session. Keep it short", and it is 78 lines. The addition is one paragraph and has to be
worth a session reading it before every question; anything longer would be the file arguing with
itself.

The reader column has a definition and it has just been got wrong once. #254 corrected a row that
named a real gate which does not read the clause; naming an unenabled example rule here would be
the same error in the same table, so the mapping row stays `nothing` and the example is mentioned
in prose where a reader is not counting coverage.

The example rule must not reach the effective set. `rules.Load` walks the plugin tree and the
project config, so a file under `examples/` is inert by construction — but `rules_hash` is
recorded in every artifact and `gate verify` recomputes it, so an unchanged verdict count over the
whole trail is the check rather than the reasoning.

A72 binds where the rule goes. A review rule in `given/builtin/` reaches every project everywhere,
including every project whose acceptance criteria are not a numbered list, which would answer
not-applicable at every review for ever. The example's header has to carry that reason, as the
four beside it carry theirs.

Three issues, one commit, and the footer takes a keyword each: `Closes #247, closes #250` reads as
one reference and a mention on GitHub, so each needs its own keyword. `CONTRIBUTING.md` says so.

The three issues carry three labels — `wp5`, `wp5`, `wp7` — and the branch carries one intent. One
intent closing three issues follows #239, which closed two.
