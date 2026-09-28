---
intent: github.com/triplem/xeno#107
phase: 01-requirements
created: "2026-09-28T18:50:39Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+907c220.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 5340b675d034ba9dbe4d446a1fcf4d316286f96602259e91450423c2a249ae59
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**AC1.** `Status` returns an error, and no status, for a check whose `result` is `fail`
and whose `findings` are empty. The error names the gate and says what is wrong with the
verdict rather than with the phase.

**AC2.** The three callers of `Status` carry the refusal out. `phase finish` and `gate
run` refuse rather than write a verdict, a decision that would rewrite a stored status
refuses rather than rewriting it, and `intent close` refuses rather than sealing an
abandonment.

**AC3.** A check with `result: fail` and at least one finding is unaffected: the phase
is red, which is what it was.

**AC4.** The three results that carry no finding legitimately are unaffected. `pass`
with no finding is green, which is every passing gate in the tree. `pending` with no
finding is provisional. `not-implemented` with no finding is neither, and a phase of
nothing but those is green.

**AC5.** The existing refusal is unchanged: two findings sharing an id still refuse,
with the error it already carries.

**AC6.** Everything green stays green. `go test ./...` passes and `gate verify` matches
every verdict the repository carries, which is the check that no real gate reaches the
new refusal.

**AC7.** A hand edited `gate.yaml` carrying `fail` with no finding is refused when a
decision sends it back through `Status`, which is the path the intake found and the
reason the rule is not in `Invariants`.

<!-- xeno:section:non-goals -->
## Non goals

`Invariants`. It keeps the rules it has. The design says why the new one is not added
there, and that reasoning is the deliverable rather than a second implementation.

The external gate contract. What a foreign gate must send, and how it is invoked, is
section 14's and WP4's. This change decides only what happens to a verdict that arrives
malformed.

`result()`. It writes `fail` only where there are findings and stays as it is. A check
on the producer cannot cover the producers that do not go through it, which is the whole
argument.

A finding for the malformed check. Refusing is not reporting: no id, no cause, no next
step, nothing new in what a verdict carries.

Any other disagreement between `result` and `findings`. A `pass` carrying findings is a
different shape with a different argument, and it is not in this issue. Named here so
that its absence is deliberate.

A gate that reports `pass` while something did fail. Undetectable locally, by this or by
anything else.

<!-- xeno:section:constraints -->
## Constraints

Nothing under `docs/` changes. Section 14 already makes external gates the extension
point and section 5 already gives them their provenance; nothing here adds to either.

No invented field, gate or finding. The rule reads two values a check already carries
and returns the error `Status` is already able to return.

`Status` keeps its signature. It returns a status and an error today, and one caller in
three already handles a refusal from it, so the callers need no new shape.

The error is the runner's own wording, not a gate's. It describes the verdict, so it
belongs to the code that derives one, and it must not read as a finding about the phase:
a person meeting it has a malformed check to fix, not an artifact.

The refusal must reach a person through every caller. A refusal swallowed by one of the
three would leave a verdict silently unwritten, which is worse than the green it
replaces.

One intent, one work package. The issue carries `wp4`, the branch carries this intent,
the commits reference #107.
