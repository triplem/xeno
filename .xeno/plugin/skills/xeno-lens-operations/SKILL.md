---
name: xeno-lens-operations
description: The operations lens of a Xeno intent — what the change needs to run, what it needs to be watched by, and what the person woken up by it will have. Use while a phase it applies to is under way, when the change adds configuration, a migration, a dependency on something running, or a new failure mode.
phases: [02-design, 03-implementation, 04-verification, 05-review]
---

# Operations lens

A lens, not a phase. It applies inside a phase that is already running and adds to what
that phase writes. Its reader is nobody in this intent: it is whoever runs the result at
three in the morning with the artifacts and the tree and nothing else.

It starts at `02-design`, because operability is a consequence of decisions rather than
a sentence in the requirements — an availability target written as a criterion before
anything is designed is a number with nothing behind it. It is the one lens that applies
at `04-verification`, because what a check proved and under which conditions is an
operational claim as much as a test result.

## What it looks for, by phase

- `02-design` — what has to exist for this to run, and what says it is running. Every
  decision that adds a component adds something to configure, something to restart and
  something that can be down; a decision that adds none is worth noting as cheap. Ask
  what the failure modes are and which of them is visible from outside.
- `03-implementation` — the run time surface the diff created. A new setting with no
  default, a migration with no way back, a timeout that is now a constant, a dependency
  that must be reachable at start. Whether the change can be deployed while the previous
  version is still running is the question most often found afterwards.
- `04-verification` — the conditions the results hold under. A check that ran against a
  fixture proves the shape, not the load; a figure with no environment named is a figure
  nobody can reproduce. The gaps section is where this belongs when the answer is that
  nothing proved it.
- `05-review` — what the release notes owe an operator: what to set, what to watch, what
  changed about rollback, and what the residual risk is in terms of what will be seen
  rather than what is wrong internally.

## What it may write

Into the phase that is running, through that phase's own commands:

    xeno question record    --intent KEY --phase NN          # the entry on stdin
    xeno assumption record  --intent KEY --phase NN --text TEXT \
        --origin <template-default|repo-convention|rules|user-input> --confidence <high|medium|low>

Most of what this lens finds is an assumption rather than a question, which is the way
round that distinguishes it from the other three: the lens usually knows what it assumed
about the environment — a value that is set, a service that is reachable, a job that
runs — and the register is where a reader finds out which of those nobody confirmed. A
question is for the choice a person owns: what downtime is acceptable, who is paged,
whether a migration may be irreversible.

At `05-review` the finding belongs in the `review-checklist` section's prose, named as
the operations lens's own. The structured entry section 5 allows — `source: lens`, no
rule id, outside what G-Policy counts — has no command behind it yet. The one command
that writes that list refuses `--source`, because every entry it writes answers a rule
of the effective set.

## What it never writes

An artifact of its own, and a gate verdict. A readiness judgement is the obvious thing
to want as a gate and the clearest case for why a lens is not one: it is a model reading
an artifact, and a deployment blocked or released by that would rest on a judgement with
no evidence behind it. Where the lens wants something mechanical — a declared build log,
a required setting, a smoke check as evidence — that is an evidence kind or a review
rule, proposed through `xeno learning record` and decided by a person.
