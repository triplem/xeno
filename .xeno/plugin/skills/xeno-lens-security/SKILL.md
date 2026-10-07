---
name: xeno-lens-security
description: The security lens of a Xeno intent — who the attacker is, what the change exposes, and what the trail owes a reader who has to judge the risk. Use while a phase it applies to is under way, when the change touches authentication, authorisation, input from outside, secrets, or a dependency.
phases: [01-requirements, 02-design, 03-implementation, 05-review]
---

# Security lens

A lens, not a phase. It applies inside a phase that is already running, adds to what
that phase writes, and produces nothing of its own. Whether it applies at all is a
judgement about the change, which is why it is loaded by its description rather than
typed.

It does not apply at `00-intake`. The intake fixes the problem and the scope, and a
security judgement there has nothing to attach to yet: there is no interface, no store
and no dependency to be wrong about. `01-requirements` is the first phase where there
is.

## What it looks for, by phase

- `01-requirements` — the criterion that names the attacker. Most acceptance criteria
  are written for somebody who wants the feature to work; the one worth adding is the
  state of the tree under somebody who does not. Also what the intake excluded that an
  attacker would reach anyway, which is a finding about the scope rather than a
  criterion.
- `02-design` — the trust boundaries the decisions move. Which component now believes
  something it did not verify, where input crosses in from outside, what holds a secret
  that did not hold one before, and what a compromise of each new component reaches. A
  decision recorded with no boundary named is the one to ask about.
- `03-implementation` — what the diff opened. Input parsed before it is checked, an
  error path that reports more than its caller may know, a permission widened to make a
  test pass, a dependency added for one function. The deviations section already holds
  the honest ones; read it before reading the code.
- `05-review` — whether the residual risk says what is still reachable, in words an
  operator can act on. A residual risk written as "hardening deferred" names nothing.

## What it may write

Into the phase that is running, through that phase's own commands:

    xeno question record    --intent KEY --phase NN          # the entry on stdin
    xeno assumption record  --intent KEY --phase NN --text TEXT \
        --origin <template-default|repo-convention|rules|user-input> --confidence <high|medium|low>

An open question is the right form where the answer is a person's: what the threat model
is, whether a risk is accepted, who the data belongs to. One question at a time, each
option with its consequence, and the recommendation named — section 8's rule holds for a
lens exactly as it holds for a phase.

An assumption is the right form where the lens went on with something it could not
check. `origin` says where the assumption came from and not that a lens wrote it: a
convention of this repository is `repo-convention`, a rule in force is `rules`. The lens
is not an origin.

At `05-review` the finding belongs in the `review-checklist` section's prose, named as
the security lens's own. The structured entry section 5 allows — `source: lens`, no rule
id, outside what G-Policy counts — has no command behind it yet: `xeno review answer`
refuses `--source`, because every entry it writes answers a rule of the effective set.
Nothing is lost from the gate by writing the prose instead: G-Policy counts rule
entries, and a lens never had one.

## What it never writes

An artifact of its own, and a gate verdict. A model based judgement is not a
deterministic check, and a lens that produced either would be read as one. Where the
lens finds something that ought to be checked mechanically, that is a review rule or a
declared scan, proposed through `xeno learning record` and decided by a person, not a
verdict this lens reaches.

It writes no phase it does not apply to either. A security finding about a sealed phase
is recorded in the phase that is running, naming the one it is about; what is sealed is
never rewritten.
