---
name: xeno-lens-architecture
description: The architecture conformance lens of a Xeno intent — whether the change holds the structure the project already decided, and where it is quietly inventing a second one. Use while a phase it applies to is under way, when the change adds a package, a dependency between two, a second way of doing something, or a pattern the tree has not used before.
phases: [02-design, 03-implementation, 05-review]
---

# Architecture conformance lens

A lens, not a phase. It applies inside a phase that is already running and adds to what
that phase writes. Conformance is the whole of it: the lens does not ask what the right
architecture would be, it asks whether this change holds the one the project has, and
says so where it does not.

Three phases, which is the narrowest set of the four. `02-design` is where a departure
is still cheap, `03-implementation` is where one happens without being decided, and
`05-review` is where what was left standing has to be written down. Nothing in the
requirements binds structure yet, and `04-verification` judges coverage rather than
shape.

## What it looks for, by phase

- `02-design` — the precedent. A decision that argues from first principles for
  something the tree already settled is the commonest waste of this phase, and the
  lens's first job is to name the place it was settled. Then the direction of every
  dependency the design adds: which package may know about which, and whether the new
  one points the way the existing ones do.
- `03-implementation` — the second way of doing something. A helper beside a helper that
  already existed, a type declared twice so two readers can each have one, a call that
  reaches past the layer that was supposed to own it, a package importing one it is
  supposed to be imported by. These arrive as convenience and are read later as a
  decision nobody took.
- `05-review` — what is now true of the structure that was not true before, and whether
  a departure left standing is recorded as a deviation naming what it departed from. A
  departure that names nothing becomes the precedent the next intent reads.

## What it may write

Into the phase that is running, through that phase's own commands:

    xeno question record    --intent KEY --phase NN          # the entry on stdin
    xeno assumption record  --intent KEY --phase NN --text TEXT \
        --origin <template-default|repo-convention|rules|user-input> --confidence <high|medium|low>

A departure the lens can name is not a question: it is a finding for the phase's own
sections — a deviation at `03-implementation`, an alternative at `02-design` — and the
lens writes it there rather than asking whether it is allowed. The open question is for
the case where the project's structure is genuinely undecided and this change would
decide it, which is a person's to settle because the answer binds every intent after
this one.

Where the lens read the structure from the tree rather than from a document, that
reading is an assumption with `origin: repo-convention`. It is the one this lens owes
most often: the convention it enforced may be three files that happen to agree.

At `05-review` the finding goes into the review checklist as the structured entry
section 12 gives a lens, which has a command of its own:

    xeno review lens --intent KEY --result <met|deviation|not-applicable> --note TEXT

The entry carries `source: lens` and no rule id, and that is what keeps it outside the
set G-Policy counts: the gate checks the entries rendered from review rules, and this
one answers none. The command therefore takes no rule and refuses one passed anyway.
The note is required whatever the result, because with no rule id nothing else in the
entry says which lens wrote it or what it found — name the architecture lens in it. The
reasoning too long for a note still belongs in the `review-checklist` section's prose.

## What it never writes

An artifact of its own, and a gate verdict. "Conformant" is the verdict this lens would
be most readily believed for and the one it may least give: conformance to a structure
nobody wrote down is a model's reading of a tree, and a gate saying so would turn that
reading into the structure. Where the departure is mechanically checkable — an import
that may not exist, a layer that may not be reached — that is a review rule or a linter
whose output is declared evidence, proposed through `xeno learning record` and decided
by a person.
