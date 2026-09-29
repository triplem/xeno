---
intent: github.com/triplem/xeno#120
phase: 00-intake
created: "2026-09-29T05:57:30Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ad0a764.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 8fcfd55b4b10b31e02abf6f4a6fab73b535416c78f29f67911caea87667cddc0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

Section 16 states the reason this work exists, and it is #120's own argument in the
specification's words:

> The digest is produced in two steps, and the split is the point. The agent writes the
summary > text; the runner filters it against the effective secret filter, hashes it and
writes the file. > The filtering is therefore deterministic and outside the model's
reach, so a model cannot route > around it, not even by accident.

The runner now writes the digest, since #120's first half. It does not filter.
`secrets_hash` is absent from every artifact it writes, and the summary it copies is the
summary it was handed, so the one place the specification puts beyond a model's reach is
the one place nothing stands.

The filter has no file. Section 4 defines it as `.xeno/plugin/secrets.yaml`, shipped,
plus `.xeno/config/secrets.yaml`, a project's additions, with `patterns` carrying an id
and a regex and `paths_never_digested` carrying globs. Neither file exists anywhere in
this tree, so there is nothing to filter against, nothing to hash, and `secrets_hash`
has carried `by-hand` in sixty intents.

`by-hand` was the honest value while no writer existed, which Appendix B says outright.
It stops being honest the moment one does.

So the state is this: the runner writes a file whose name says it was produced, the
specification says that file is filtered, and a summary carrying a secret reaches the
repository unchanged. The review of XENO-0201 named this the thing to close first, and
it named the reason — the file now looks processed, which is worse than a hand written
one that visibly was not.

<!-- xeno:section:scope -->
## Scope

`.xeno/plugin/secrets.yaml` is shipped, with the patterns and the never digested paths
section 4 gives, kept small on purpose.

A package assembles the effective filter from the shipped file and a project's
additions, reduces it to one value for `secrets_hash`, and redacts a text against it.
Appendix B leaves that reduction to "the package that first writes them", so the
computation is defined here and recorded as an assumption rather than invented silently.

`phase finish` filters the summary before it writes the digest, and writes
`secrets_hash` into it. `section set` writes `secrets_hash` into `output.md`.

A35 is amended again: two fields remain without a writer, `tool_version` and
`rules_hash`.

Not `rules_hash`. WP4.

Not `tool_version`. Section 12's agent block names the tool and the model and not the
version, so the field has no home short of the harness, and WP11 decides where it comes
from.

Not G-Secret. The gate scans artifacts for secrets and stays `not-implemented`; this is
the digest writer's half, which section 4 says reads the same file. What the gate does
when it finds one is its own decision.

Not the honesty rule for `by-hand`. Tightening it turns every sealed artifact red, which
was measured rather than assumed, and the design records the number.

Not `paths_never_digested` applied to the summary. The runner does not source the text
it is given, so the globs enter the hash and govern nothing yet.

<!-- xeno:section:context-rationale -->
## Why this context

**The filtering belongs in the runner because section 16 says why, and the reason is
this issue.** Deterministic, outside the model's reach, unroutable even by accident.
Every other part of #120 asks whether an agent can be prevented from doing something,
and section 5 answers that it cannot; this part is the one place the specification says
the prevention is real, because the runner holds the text between the agent and the
file.

**The shipped filter is small on purpose.** A pattern set is a security decision and a
large one written quickly is worse than a small one written carefully: a regex that
matches too much redacts prose and teaches whoever reads the digest to distrust the
redaction, and one that matches too little is a claim nobody checked. Section 4 names
two patterns as its example and two paths, and the file ships close to that, with the
project file as the documented way to add more.

**The computation of `secrets_hash` is a decision this intent takes, because Appendix B
refuses to take it.** It says the two set hashes are "defined by the package that first
writes them", since how several files reduce to one value belongs with the code that
assembles the set. So the choice is recorded as an assumption, and it hashes the
effective set in a canonical form rather than the bytes of the files: a comment edited
in the shipped filter would otherwise change every digest's hash without changing what
any digest passed through.

**A project adds and never removes, which decides what the union is.** Section 4 says a
filter a project can switch off is not a filter. Matching is a disjunction, so keeping
both a shipped pattern and a project one with the same id can only widen what is caught,
and that is the safe direction.

**`by-hand` stays honest for `secrets_hash`, and the number is the argument.** Appendix
B says `by-hand` is honest where a field has no writer yet and wrong once one exists, so
tightening the rule is what this change ought to do. Tightening it was measured: every
one of the eighty-seven verdicts in this repository diverges, because sixty intents were
sealed with `by-hand` when that was the only honest value. A verdict is a statement
about a moment, and `gate verify` recomputes now, so the two cannot be reconciled
without rewriting sealed history. The loophole is left open, recorded, and narrower than
it sounds: new artifacts carry a real hash because the writer writes one.

**What this does not settle about #120.** `tool_version` still has no source, so an
agent still cannot produce a complete artifact with the tool alone, and the enforcement
question stays where the scoping left it.
