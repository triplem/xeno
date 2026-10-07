# The documents

Two of these are normative. Where the code and a normative document disagree, the
document wins until a person changes it, and a change to one is its own commit made
before the code that follows from it.

## What the process is

- **[process-definition.md](process-definition.md) — normative.** The six phases, the
  artifacts each produces, the gates, the assumption register, the rule model, learning,
  cost, and the environment. Section 1 says what the process is for; section 6 fixes who
  does what with which command and in which order. Section 12 keeps the two surfaces
  apart: `xeno gate ...` never touches the network, which is the property the verdicts
  rest on, while everything else in the CLI may. Appendix C explains the name.
- **[implementation-plan.md](implementation-plan.md) — normative.** The build order:
  twenty work packages with what each is done when, the sequence they are taken in, and
  the verification points answered rather than assumed. Section 9 holds the open
  questions and their answers.

## Running it

- [commands.md](commands.md) — every command with its flags, which is the text
  `xeno --help` prints.
- [symbol-index.md](symbol-index.md) — the format of the symbol index a project produces
  for Xeno to read, as the annotated worked example the reader is tested against. Xeno
  ships no indexer, and a phase runs without one.

## What is being built, and what it has cost

- [the-trail-in-this-repository.md](the-trail-in-this-repository.md) — Xeno governs its
  own construction, so `.xeno/` here is the longest worked example there is. What the
  version fields say in this trail and why, the two shapes of the intents directory, and
  the coverage against the plan's acceptance criteria.
- [m0-gate-job.md](m0-gate-job.md) — how M0 was reached once, in the tree and with the
  tooling of the day. A record and not a procedure to follow now; it is also the guide
  the intents up to `XENO-0106` were run under, one intake per work package and no
  further.

## Why the tree is the way it is

- [assumptions.md](assumptions.md) — the open register of assumptions and decisions
  taken while building, one row each, with what a person has said about it. A row
  belongs here when the fact outlives the intent that found it.
- [clause-readers.md](clause-readers.md) — one pass over both normative documents,
  listing what each clause requires and what would fail if it were violated, including
  the clauses that have nothing.
- [supply-chain.md](supply-chain.md) — what ends up in a released binary, and what the
  pipeline runs to produce it. The second list is the longer one, which is the point of
  writing it down.

## Where v2 goes

Both are drafts and neither is ratified. They are here so that v1's choices can be read
against what they are deliberately not.

- [orchestrator-evaluation.md](orchestrator-evaluation.md) — an architecture decision
  record: which agent platform v2 delegates a phase to, against which criteria, what was
  rejected, and when the decision is revisited.
- [v2-delta.md](v2-delta.md) — what making Xeno multi-agent changes in the process
  definition, section by section, so the next revision can be written from a list rather
  than from memory.

Beside these, in the repository root: [CONTRIBUTING.md](../CONTRIBUTING.md) for the
commit and pull request conventions and the DCO, [SECURITY.md](../SECURITY.md), and
[CLAUDE.md](../CLAUDE.md), which every session in this repository is sent with.
