---
intent: github.com/triplem/xeno#188
phase: 03-implementation
created: "2026-10-03T21:28:37Z"
schema_version: "1.0"
runner_version: dev+efc44a9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9d4ce1497faf8247d67a0f700420d69bdc21f56890c1b886db349d20c543a2a1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
open_questions:
    - key: Q-2
      text: The figure marks nine of the ten rows the default listing shows, because nine of the last ten intents asked nothing. Is the listing the right place for a mark that is on almost every row of it?
      options:
        - text: Leave it on both forms. The listing is loud because the fact is, and it stops being loud as intents start recording the exchange.
          consequence: 'a reader sees the pattern across intents, which is the fact #188 measured and the one no single intent''s page can show; until the habit changes the column is nearly uniform, which is what a measurement of a habit looks like'
          recommended: true
        - text: Keep the line in the one-intent form and take the mark out of the listing.
          consequence: 'the listing stays as #136 left it and the figure is there for whoever looks at one intent before its merge; the pattern across intents then has no reader, which is the half of the issue that says a figure somebody looks at is what was missing'
        - text: 'Invert it: mark the rows that did record an exchange instead.'
          consequence: the mark is rare and so reads as information rather than as wallpaper, but it reports the compliant case and leaves the omission unmarked, which is the thing nothing currently sees
        - text: Something else, in the words of whoever decides
          free: true
---

# Implementation

<!-- xeno:section:changes -->
## Changes

**`internal/gates/gates.go`.** `questionShape` and `decisionShape` become
`QuestionShape` and `DecisionShape`, unchanged in behaviour, each with a sentence
saying why it is exported: the command that writes one of these entries refuses on
these checks before the write, and a second opinion about the shape of a question
would be a second definition of one. G-Questions itself is not touched.

**`internal/runner/exchange.go`, new.** `RecordQuestion` takes the entry as YAML and
`RecordDecision` takes the built record; both append one entry to the frontmatter
of the phase named and return what they wrote. Around them: `exchange`, which is
what the intent has raised and settled over every phase; `nextExchangeID`, which
continues the intent's own numbering rather than counting entries; `shapeRefusal`,
which turns what the gate would have reported into a refusal; `artifact`, which
reads an `output.md` three ways and refuses where there is none; and `amendFront`,
which replaces one field and writes the body back byte for byte.

**The question is decoded with unknown fields refused.**
`yaml.Decoder.KnownFields(true)`, so a mistyped `consequence` is a refusal rather
than a dropped half of an entry. The standing rule about invented fields applied
where silence is the alternative. It earned itself immediately: the first question
recorded through the command in this repository was refused for a stray key, which
is in this phase's own trail.

**`internal/runner/runner.go`.** `IntentSummary` gains `Questions` and `Decisions`,
filled by `Summarise` on the walk it already does, and the `AskedNothing` predicate with
the reasoning for both of its exclusions beside it. `summarise` becomes `Summarise`,
because the one-intent form of the status command prints the figure the listing
marks and a second walk computing it would be a second definition of it.

**`cmd/xeno/main.go`.** Two entries in `commands`, both needing a key and a phase. Three
flags: `--chosen`, `--proposed-by` and `--withdraw`; `--reason`, `--by` and `--resolves`
are reused, and `--resolves`'s help text loses the word assumption, since it now names
a question for two commands. `cmdQuestionRecord` reads the entry the way `section set`
reads a section; `cmdDecisionRecord` builds the record from flags and prints what it
recorded, with `resolving` or `withdrawing` where the entry names a question. Three
lines in `usage`. The row mark in `cmdIntentList` and the line after the table in
`cmdIntentStatus`.

**`internal/runner/exchange_test.go`, new.** Fifteen tests: the key the runner
assigns and the sequence continuing across phases; a question with no options
refused with the artifact unchanged; the honest `no_options` accepted; a given key
refused; an unknown field refused; the five refusals a decision has, in a table;
a decision resolving a question nobody raised refused; the loop end to end, raised
in P0 and decided in P1 with G-Questions passing at P5; a withdrawal resolving as
well; the body and the other fields untouched; a phase with no artifact refused;
and the figure's three answers, including that an abandoned intent is left alone
and that one entry of either kind is enough to answer it.

**`cmd/xeno/main_test.go`.** One test driving both commands as a session does,
the question from a file and the decision from flags, asserting the assigned ids
in what is printed, the refusal without `--by`, both entries in the frontmatter,
and the section written before them still in the body. With it `inputFile`, because
the existing `writeTemp` names its file after its content and cannot hold a document.

<!-- xeno:section:deviations -->
## Deviations from the design

**A refusal the design did not name: `--resolves` has to name a question the intent
raised.** Added because the alternative is silent and late. A typo in the key
leaves the question unresolved, the decision looks like its answer, and nothing says
otherwise until G-Questions reports it at P5 with an apparent resolution sitting in
the trail. The check costs the walk `exchange` already does.

**`summarise` was renamed rather than wrapped.** The design said the predicate is read
by both forms of the command; it did not say how the one-intent form would reach it. An
exported wrapper beside an unexported function of the same name would be two names
for one thing, which is the objection this tree records about helpers that exist twice.

**`--resolves`'s help text changed.** It said "the open question this assumption
answers" and now says "the open question this answers", because three commands
pass it and only one of them is about an assumption. A behaviour change to nothing,
recorded because it is a line in the diff that no acceptance criterion asked for.

**The printed figure is not asserted by a test.** The predicate behind it is, in all
three of its answers. What a cmd level test would add is that the mark reaches the
row, and the repository's own listing shows that on 97 intents, which is evidence
of a kind a fixture cannot be: it is the measurement #188 filed. P4 records the
output. The judgement is that a test which built a six-phase intent in a temporary
tree to assert a suffix would assert the suffix and not the figure.

**Nothing else departs from the design.** The channels, the assigned ids, the refusals
from the gate's own checks, `amendFront` leaving the body alone, no gate change and
no document change are as P2 wrote them.
