---
name: xeno-lens-privacy
description: The privacy lens of a Xeno intent — whose data the change touches, where it comes to rest, how long it stays and who can read it. Use while a phase it applies to is under way, when the change touches personal data, identifiers, logs, telemetry or anything leaving the system.
phases: [01-requirements, 02-design, 03-implementation, 05-review]
---

# Privacy lens

A lens, not a phase. It applies inside a phase that is already running and adds to what
that phase writes. It overlaps the security lens and is not the same question: security
asks who should not be able to read the data, privacy asks whether anybody should be
holding it at all, and for how long.

The phases are the four where data has a shape to judge. At `00-intake` there is a
problem and no data yet, and at `04-verification` the question the lens would ask —
whether the checks exercise the retention and the deletion path — is the question it
already asked of the design and reads again at review.

## What it looks for, by phase

- `01-requirements` — what personal data enters, named rather than implied. Who the
  subject is, which field identifies them, and whether a criterion says what happens
  when they ask for it back. A criterion that only describes the happy path leaves
  deletion to nobody.
- `02-design` — where the data comes to rest and for how long. Every store, cache, queue
  and log the decisions add is a copy, and a copy with no retention named outlives the
  system it was built for. Also what crosses a boundary the subject was not told about:
  a third party, another region, a central log.
- `03-implementation` — what the diff now carries. An identifier in a log line, a whole
  record in an error message, a payload in a trace, a fixture made from real data. This
  is the phase where privacy is lost by accident rather than by decision, and the secret
  filter catches credentials, not people.
- `05-review` — whether the residual risk says what is retained and what is not covered,
  and whether the release notes say anything that a subject reading them would be
  surprised by.

## What it may write

Into the phase that is running, through that phase's own commands:

    xeno question record    --intent KEY --phase NN          # the entry on stdin
    xeno assumption record  --intent KEY --phase NN --text TEXT \
        --origin <template-default|repo-convention|rules|user-input> --confidence <high|medium|low>

Whether data may be held, for how long, and under which basis is a person's answer and
not a model's, so it is an open question with its options, their consequences and a
recommendation — never an assumption dressed as agreement. What does belong in the
assumption register is the reading the lens worked from: that a field is not personal
data, that a log is not retained, that a store is inside the boundary. Each of those is
checkable by somebody, which is what makes it an assumption rather than a question.

At `05-review` the finding goes into the review checklist as the structured entry
section 12 gives a lens, which has a command of its own:

    xeno review lens --intent KEY --result <met|deviation|not-applicable> --note TEXT

The entry carries `source: lens` and no rule id, and that is what keeps it outside the
set G-Policy counts: the gate checks the entries rendered from review rules, and this
one answers none. The command therefore takes no rule and refuses one passed anyway.
The note is required whatever the result, because with no rule id nothing else in the
entry says which lens wrote it or what it found — name the privacy lens in it. The
reasoning too long for a note still belongs in the `review-checklist` section's prose.

## What it never writes

An artifact of its own, and a gate verdict. "Compliant" is not a verdict a lens reaches:
a model based judgement is not a deterministic check, and the one thing worse than no
privacy review is a green one nobody can trace to a person. Where the lens wants a
mechanical check — a pattern that must not appear in a log, a retention value that must
be set — that is a review rule or a secret pattern, proposed through `xeno learning
record` and decided by a person.
