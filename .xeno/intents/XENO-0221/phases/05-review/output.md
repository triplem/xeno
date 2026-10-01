---
intent: github.com/triplem/xeno#167
phase: 05-review
created: "2026-10-01T16:59:14Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15693cf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0a729959e671222d1411ef2cf88a85cfa2fb2f54373ef416702ae3d456e8b9ee
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
  - rule: deviations-are-traceable
    result: met
    note: >-
      The six deviations of 03-implementation each name what they depart from: the design's
      timeout decision, the alternative of a test that waits a minute, the learning #160 left,
      the applicability criteria, the fixture, and the four pieces that had a demonstration.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      A76 is a contract with code this project does not own, published in the register and in the
      configuration's own comments. There is nothing to migrate from, because no external gate has
      ever existed; gates.Ctx gaining a field is internal and additive.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      No dependency was added. The one vendored dependency is unchanged and internal/external uses
      the standard library only.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The structured answers to the three shipped review rules are in the frontmatter, under
`review_checklist`. What follows is the reasoning and the questions this change raises beyond
them.

**`deviations-are-traceable` — met.** The six deviations of 03-implementation each name what they
depart from: the unbounded wait against the design's own timeout decision, the test seam against
the alternative of a test that waits a minute, the thrown-away gate against the learning #160 left,
the fixed phase in `run1` against the applicability criteria, the fixture against what I assumed it
did, and the missing demonstration against the four pieces that had one.

**`interface-change-needs-a-migration-note` — met.** Two things outside this intent depend on what
changed. `gates.Ctx` gains a field, which is internal and additive. A76 is the one that matters:
it is a contract with code this project does not own, and it is published in the register and in
the configuration's own comments, which is where a project's tool will read it. There is nothing
to migrate from, because no external gate has ever existed.

**`new-dependency-needs-a-rationale` — not-applicable.** No dependency was added. The one vendored
dependency is unchanged, and the external package uses the standard library only.

**Beyond the three rules.**

*Is the mark enough?* Section 14 says it is the point, and this piece provides it: the check
carries `provenance: external`, the finding carries an id derived from foreign wording, and the
decision does not survive. What it does not provide is containment, and the gaps say so in the
first sentence rather than the last. A reader who takes "external gates are supported" for
"external gates are sandboxed" has been told otherwise in three places.

*Did anything skip the verdict path?* No. The producer is injected into `gates.Run` so that an
external check meets `carryForward`, `Invariants` and `Status` with every other check, which is
what #66's comment asked for. A second assembly point was the alternative and it was rejected.

*Was anything invented?* One contract, because section 14 stops at "receiving and returning JSON",
and one timeout, because nothing bounds a foreign command. Both are register rows with their
reasoning, and both are stated where a project meets them. Nothing else: no sandbox, no retry, no
discovery, no predicate type.

*What is the weakest part?* That no external gate has ever run outside a test, and that the tests'
fake commands were written by the person who wrote the mechanism. Four of the five WP4 pieces found
something on a throwaway copy of this repository that no fixture had shown; this one had no such
demonstration, and the thing it found instead — a timeout that bounded nothing — was found by a
test's duration rather than by its assertion.

<!-- xeno:section:release-notes -->
## Release notes

**A project's own tool can make a statement.** An external gate is declared in `project.yaml`
with an `id`, a `path`, a `sha256` and the `phases` it applies to. It runs only when the hash
matches, it is given a JSON object on standard input and answers one on standard output, its exit
status decides the check, and every finding it produces is marked `provenance: external` in
`gate.yaml`.

**A modified gate refuses to run.** The hash is read from disk and compared before anything is
started, on every run. A mismatch is a red check naming the declared hash and the one the file
has — not a skipped check, because a modified tool must not be able to make a verdict quieter.

**A decision on an external finding does not survive the next run.** Its cause is produced by
foreign code and is stable only by that code's promise, so a release is taken again rather than
kept. Where a project wants a lasting release, the honest form is a rule of its own.

**The contract.** On standard input: `intent`, `phase`, `phase_dir`, `artifacts_hash`, `root`. On
standard output: `findings`, each with a required `cause` and an optional `file` and `next`.
Unknown fields in either direction are ignored, so a tool may answer a superset. A command has
sixty seconds. A command that answers something else, exits non-zero with nothing to say, cannot
be started, or does not finish is a red check naming what happened.

**A project that declares none is unaffected**, which is every project today and the default
Appendix A names. The chain of trust stays closed: vendored plugin, hash, signed release, nothing
executed.

**There is no sandbox.** The command runs with whatever the pipeline gives it. What the trail
provides is the mark on every statement foreign code made, which is what section 14 says the
point is.

**WP4 is complete.** The rule format and its four levels, the two kinds, precedence and the
collisions it refuses, the six predicate types, G-Rules and G-Policy, the shipped set with its
examples and hook templates, and external gates.

Closes #167. Refs #1.

<!-- xeno:section:residual-risk -->
## Residual risk

**No external gate has run outside a test, and the tests' gates are mine.** Fourteen fake
commands, written by the person who wrote the mechanism, which is the same closed loop the last
three reviews have named — except that here there is no shipped content to open it. The first
real gate is the first independent reading of A76, and the thing most likely to be wrong is a
field somebody expected and did not get.

**Foreign code runs in the verification step as well as in the run.** `gate verify` recomputes
verdicts, so a project with an external gate executes its own tool once per phase in CI's
verification too. Nobody has measured that on a trail this size, and the measurement that exists
says `gate verify` is already 512–560 ms over 204 verdicts before any external gate runs at all.

**The timeout is one number and the standard error goes nowhere.** Sixty seconds for every gate
with no per-gate override, and a tool that explains itself on standard error says it into
nothing. Both will read as bugs to whoever writes the first real gate, and both are cheap to
change once somebody has.

**A tool can flood a verdict.** Each finding's text is clipped; the number of findings is not. Ten
thousand findings put ten thousand ids in `gate.yaml`, and the only thing between a project and
that is its own tool.

**The mark is a statement about provenance and not about trust.** `provenance: external` says
which code made a statement. It does not say that code was reviewed, pinned to a version, built
from source anybody read, or incapable of writing to the repository while the gate ran. The hash
pins the file's content and nothing pins what the file does.

**A74's hole is wider with external gates in it.** A phase is judged only against the rule set its
artifact recorded, and an external gate is not in that hash at all — the declaration lives in
`project.yaml`, which no artifact hashes. So adding, removing or re-pointing an external gate
changes what a phase is judged by with nothing in the artifact recording that it changed. That is
a question for the specification and it is the sharpest one this work package has produced.
