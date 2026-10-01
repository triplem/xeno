---
intent: github.com/triplem/xeno#162
phase: 00-intake
created: "2026-10-01T15:49:33Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9f9c5ff207244f7fd2dc88adb5ca186a639c591f3ab078ad0945c5a1b08ed728
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
---

# Intake

<!-- xeno:section:problem -->
## Problem

#160 left a registry and no entries, so `kind: checked` is currently worse than useless: a
checked rule turns every phase it applies to red with a finding saying no implementation exists
for its type. That was the right verdict to ship — a rule in force that nothing evaluates is the
silently green claim section 16 catalogues — and it means a project cannot use half of what
section 9 defines.

**Five types are named and none exists.** `section-implies-section`, which is the one the
shipped set needs, and four that read the commit history of the change under review:
`commit-message`, `commit-trailer`, `commit-signature` and `approver-not-author`. Section 9
names them in a table with what each reads, so none of them is a design question; what they need
is an implementation and the two inputs they read from.

**One of the two shipped patterns is missing.** `conventional-commits` is in
`internal/gates/patterns.go` and answers `xeno check commit-message` today. The second — the
same form carrying an issue reference in the subject — is not, and section 9 says what it is for:
a project needs it when the reference has to survive a squash merge, because the squashed
message is built from the merge request title.

**Nothing in this repository reads a commit.** `Base` and `Head` sit on the runner carrying a
comment that says no gate reads them until WP4, and they do not reach the gate context at all.
Four of the five types read a subject, a trailer, a signature status or an author, which is what
`git log` has, and no package here has ever asked it for anything.

**That collides with a sentence in `internal/gates`.** Its package comment says gates "read, they
never run anything and never ask a model". Reading a commit range means running `git log`, and the
sentence as written forbids it. What it was written to mean is that a gate runs no build, no test
and no scanner — G-Build and G-Test read declared results for exactly that reason — but the
sentence is wrong as it stands, and a piece that quietly violated the package's own comment would
leave the next reader to work out which half to believe.

**The range cannot be guessed, and a rule that needs one has to say so.** Section 12 and WP4 both
state it: the range is passed in, never inferred, because a guessed range means different verdicts
locally and in CI from the same repository state. So four of the five types have a second failure
mode that has nothing to do with the commits — the gate was given no range — and WP4's done-when
lists it as its own criterion, which it would not do if a missing range could be treated as an
empty one.

**And one of them reads a verdict rather than an artifact.** `approver-not-author` compares the
`by` of a decision in `gate.yaml` against the authors of the range. It is the only predicate whose
input is something a previous gate run wrote, so it evaluates on the run after the decision — which
the decision's own commit triggers — and it is the only one where a green verdict depends on who
ran the command rather than on what the tree contains.

<!-- xeno:section:scope -->
## Scope

**In scope.** The five predicate types, registered in `internal/gates` so G-Policy evaluates
them where it reported them unimplemented. The second shipped pattern, beside the one that
exists and reachable the same two ways. A reader for the commit range, since a subject, a
trailer, a signature status and an author are what `git log` has and nothing here asks it.
`Base` and `Head` carried from the runner into the gate context, which is what their comment has
been waiting for. A rule that needs a range and did not get one, red, as its own case. The
package comment in `internal/gates` corrected to say what it means. Tests for every type, every
failure mode, and both patterns.

**Out of scope, and each for its own reason.**

The shipped set and the examples. Four generic rules under `given/builtin/`, Conventional
Commits, the `Xeno-Intent:` trailer and the signature rule under `examples/rules/`: all of them
are rule files that name these types, and writing them in the same intent would mean the first
rules anybody reads were written by whoever wrote the types, against no independent reading of
whether the parameters make sense. It is the next piece and this one is what makes it possible.

The two hook templates. `examples/hooks/` is content for the same piece, and the
`prepare-commit-msg` one has nothing to do with a predicate.

External gates. Independent of all of this, as they have been since #158.

Project-defined predicate types. Section 9 puts them outside v1 and says why: an external gate
covers the same ground and keeps the boundary honest. The registry's keys stay the names from
the specification, and nothing reads a type name out of a project's file into the map.

A range the runner works out for itself. Forbidden, and the reason is in the problem.

**Two boundaries worth naming.** The commit reader runs `git`, which no package here has done,
so this piece decides how this repository treats a subprocess in a gate path: read-only
arguments, no network, failure reported as a finding rather than a panic, and the package comment
saying so. And `commit-signature` has to decide what counts as a valid signature, which git
answers in five letters rather than a boolean, so the choice is recorded rather than buried in a
comparison.

<!-- xeno:section:context-rationale -->
## Why this context

Section 9's predicate paragraphs are the specification: the table of five types and what each
reads, the rule that a type is parameterised and never carries an expression, the four
consequences of shipping patterns rather than letting projects write expressions, the two
sentences about `approver-not-author` that say when it evaluates and that it compares two
self-asserted strings, and the paragraph on the two shipped patterns and what the second is for.
Read entire, because the four consequences are the argument this piece has to not undo.

Section 12 is read for the range being passed in rather than inferred, and WP4 of the plan for
the same rule stated as an acceptance criterion, which is what makes a missing range its own
case rather than an empty one.

`internal/gates/patterns.go` is read as the thing to extend rather than replace: the map is
already the single answer for a gate and for `xeno check commit-message`, which is the property
section 9 wants, and the second pattern goes in beside the first. `CheckMessage` is read for how
a subject is taken from a message, since four of the five types read a subject the same way.

`internal/gates/gates.go` is read for the registry #160 left, for `Ctx` and what it currently
carries, and for its package comment, which this piece has to correct.

`internal/runner/runner.go` is read for `Base` and `Head` and the one place `gates.Ctx` is
constructed, which is the whole plumbing change.

`internal/model/model.go` is read for `DecisionOnFinding`, whose `by` is one half of
`approver-not-author`, and for `Gate` and `Check`, since the decision is reached through the
previous run's `gate.yaml`.

`internal/template/template.go` is read for `Parse`, which turns a rendered artifact back into
sections, because `section-implies-section` asks whether a section is non-empty and the artifact
on disk is the rendered form.

`internal/host/gitlab` is read for one thing only: how this repository already words a failure
from something outside its own process, since the commit reader is the second such boundary and
the first one inside a gate.

Nothing outside the repository is needed, and nothing may be: a gate is network free, and `git
log` reads the clone that is already there.
