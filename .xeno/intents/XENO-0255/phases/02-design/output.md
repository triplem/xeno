---
intent: github.com/triplem/xeno#247
phase: 02-design
created: "2026-10-05T19:35:42Z"
schema_version: "1.0"
runner_version: dev+8e3b29b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1f21452b46f2ad871d8632cf53fce0ce2a999587b00d90f1500968c70bbfb0f9
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

One intent for three issues, because the three are one finding. Each is a clause whose reader can
only be a person, and the work is recording that in the three places this repository has for it;
three intents would write the same P0 three times and cost about 4,200 lines of record for what
one costs. #239 is the precedent for one intent closing more than one issue, and the footer takes
a keyword each.

The reason's row says "no field" before it says "nothing reads it", because the first is why the
second cannot be fixed. A reader meeting "nothing reads it" alone would reasonably ask for a gate;
the row has to say the clause cannot be written down, which is the stronger and less obvious fact.

The mapping row is not touched and the example is mentioned in prose. The reader column means what
would fail if the clause were violated, an unenabled rule fails nothing, and #254 corrected a row
one line away for naming a reader that does not read. Putting the example in the sentence under the
table puts it where somebody deciding whether to adopt it is reading and out of where somebody
counting coverage is reading.

The `CLAUDE.md` paragraph goes under the heading that already covers how to work rather than as a
new section. The file's structure is three standing rules, then conventions, then where things are
written down; asking one question at a time is a convention about the exchange, so it sits with the
conventions and adds no heading.

The paragraph carries XENO-0243 as the instance and the dependency as the reason. "Ask one at a
time" without the reason reads as etiquette; the reason is that Q-2's and Q-3's options assumed an
answer to Q-1 that nobody had given, so a person answering all three answers two of them under an
assumption they have not made. That is checkable in the trail and is what makes the convention a
rule rather than a preference.

The example rule is `scope: org` and `binding: false`, matching
`documentation-follows-the-change.yaml` rather than the shipped set's `scope: builtin`. An example
is adopted by copying it to a level, and its scope has to be the level it is adopted at, which
that file's header already explains.

The rule is worded over the mapping rather than over the criteria. "Every acceptance criterion is
answered in the mapping" asks a person to compare two sections of one intent; "the mapping is
complete" asks them to judge a word. Section 9's sixth clause bounds it the same way — completeness,
not selection and not existence — so the wording follows the clause it stands in for.

The rule's header says what it cannot do. It is answered once per intent, not once per criterion,
which is a weaker instrument than the gate section 7 asks for, and saying so in the file is what
keeps a project from adopting it and believing the clause is mechanically covered.

<!-- xeno:section:alternatives -->
## Alternatives

Three separate intents was the orthodox reading of "a branch carries one intent" and is rejected on
the measurement this session produced. The record is a fixed ~1,400 lines per intent whatever the
change, so three intents for three one-file changes is about 4,200 lines to say one thing three
times. #239's precedent and the fact that the three issues share a single finding are what make one
intent honest rather than merely cheaper.

Waiting for the specification commits was the alternative the maintainer weighed and declined. It
is the better end state for all three — a field, a clause, a numbering convention, and real readers
— and it leaves three issues open indefinitely on a commit only a person can write. Declining it
does not foreclose it: each issue stays open for the mechanism if the maintainer wants it, and what
closes is the question of what reads the clause today.

Putting the sequence rule in the process definition was considered and is not the agent's to do.
#248 names `CLAUDE.md` as the other candidate home and gives the better argument for it anyway: the
rule is about the exchange rather than about a question's shape, and `CLAUDE.md` is read before a
question is asked where section 8 is read when one is being judged.

A `checked` rule for the mapping was considered. A predicate like `section-implies-section` exists
and would assert that `test-mapping` carries text, which `release-notes-are-filled` already does
for its own section and which says nothing about completeness. Section 9's clause is about coverage
and no predicate expresses it, so a checked rule would be a reader that cannot fail — A90's finding
exactly.

Shipping the mapping rule in `given/builtin/` was considered and rejected on A72, which is the
precedent and the measurement: a review rule there reaches every project everywhere, including
every project whose acceptance criteria are prose, and it would be answered not-applicable with a
note at every review for ever.

Adopting the example in this repository's own `.xeno/config/rules/` was considered. It would make
this project read the mapping clause, which is attractive — and it is a decision about this
project's reviews, taken by whoever maintains them, not part of writing the example. It would also
add a rule to the effective set mid-intent, which changes `rules_hash` and is the one thing
criterion 8 exists to prove did not happen.

Adding a row to the audit for the sequence convention was considered and rejected: the audit is a
pass over the two normative documents, and a convention in `CLAUDE.md` is not a clause in either.
A row for it would be the audit describing something outside its own subject.

<!-- xeno:section:impact -->
## Impact

Three files changed, one added. `docs/clause-readers.md` gains a row and a sentence and its count
moves 38 to 39; `CLAUDE.md` gains a paragraph; `examples/rules/` gains a fourth example. Nothing
else, and no code.

Nothing executable changes and nothing can. `rules.Load` walks `.xeno/plugin/rules/` and the
project config, so a file under `examples/` is inert by construction, and the 405 verdicts are the
proof rather than the claim: an example that reached the effective set would change `rules_hash`
and every verdict with it.

What changes for a session is one paragraph in the file it reads before everything. A question
asked one at a time is the only behaviour in this intent that anything will act on, and it will act
on it immediately and for every session, which is more consequence than the other two changes have
between them.

What changes for a reader of the audit is that section 8 is now fully described: three
requirements, three rows, and the third saying the clause cannot be written down. That is the
first clause in the table whose reader column says a field is missing rather than naming a gate or
saying nothing, which is a new kind of entry and worth watching — it is also the most honest
description of the state the table has held.

What a project gains is an option it did not have. The mapping half of G-Test can be read by a
person in the P5 checklist by copying one file, and the file says what that buys and what it does
not. Nothing is adopted by writing it.

What stays open is what the three issues asked for. A field for the reason, a clause for sequence,
a criterion that is identifiable — all three need a normative commit, all three remain available,
and the issues are closed on the question of what reads the clause rather than on the question of
whether a mechanism should exist. That distinction has to survive in the commit message or the
issues will read as settled when they are not.

The honest cost is that two of the three changes are read by nobody who is not already looking.
The audit row and the example rule sit in documents nothing reads, which is A90's own subject; the
`CLAUDE.md` paragraph is the exception and is the only part of this intent with a mechanism behind
it, that mechanism being that the file is sent with every request.
