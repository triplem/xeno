---
intent: github.com/triplem/xeno#167
phase: 00-intake
created: "2026-10-01T16:46:10Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15693cf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c78b7950bdda2a62b83b35f36dab47ec87f6fd6989f1fcefb08c74b19a54c755
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

Thirteen gates read this repository and one kind of statement cannot be made at all: the one a
project's own tool makes. Section 14 fixes the mechanism — a command declared in `project.yaml`
with its path and hash, JSON in and out, the exit status deciding, running only when the hash
matches, every finding marked `provenance: external` — and nothing in the tree reads the
declaration, runs the command or produces such a check.

**The invariants for it exist and have never been met by anything.** `gates.ExternalProvenance`,
`carryForward` declining to attach a decision to an external finding, and `Invariants` refusing
one that carries it anyway all came from #66, and the comment where the verdict is produced says
why it is checked there: "so that a second path into it, an external gate above all, meets the
same rule as the first". There has never been a second path. A rule tested against nothing is in
the same position the rule engine was in last week.

**Three things the documents require and the tree cannot express.**

*The declaration.* `model.Project` carries the runner version, evidence, language, agent and
index. It does not carry `external_gates`, although Appendix A lists it with an `id`, a `path`, a
`sha256` "checked before every run", and the `phases` it applies to.

*The refusal.* "It runs only when the hash matches" is a sentence about what must not happen, and
WP4's done-when makes it a criterion in its own words: a modified external gate "refuses to run
rather than running unnoticed". Nothing today could refuse, because nothing runs.

*The contract.* "Receiving and returning JSON" is as far as either document goes. What the command
is sent, what it may return, and what a malformed answer means are unwritten, so they have to be
decided here and written down — and the decision is load-bearing in a way the other two are not,
because it is the only part a project's tool has to agree with.

**And one thing neither document bounds.** Nothing says how long a foreign command may take. A
gate that hangs is not a deterministic gate, and the first external gate somebody writes will have
a bug in it, so a limit has to exist and be stated rather than inherited from whatever the
pipeline's own timeout happens to be.

**The marking is the point and it is the whole reason this is awkward.** Without external gates the
chain of trust is closed: vendored plugin, hash, signed release, nothing executed. This piece puts
foreign code with repository access inside the gate path, and what makes that acceptable is not a
sandbox — there is none — but that the trail says which statement came from Xeno and which did
not. That is also why execution cannot live in `internal/gates`, whose comment was rewritten once
already for one `git log` and whose subject is reading.

<!-- xeno:section:scope -->
## Scope

**In scope.** The declaration as a type on `model.Project` and in the scaffold's `project.yaml`
comments. A new package that runs a declared command: the hash check before every run, the JSON
contract in both directions, the exit status deciding, a timeout, and a check per declared gate
carrying `provenance: external`. The runner passing those checks into the verdict so that they
meet `Invariants` and `Status` with every other check, and their findings routed through the same
`carryForward`. Tests for the refusal, the contract, the failure modes and the decision that does
not survive.

**Out of scope, and each for its own reason.**

A sandbox. Section 14 says the marking is the point, not the containment. A half-sandbox — a
working directory, a trimmed environment — would read as protection and provide none, and the
honest form is that a project enabling external gates has decided to run its own code in its own
pipeline.

Project-defined predicate types. Section 14 puts them outside v1 and names this piece as the
answer: either Xeno checks something with its own means or it is foreign code and marked as such.

A catalogue, a registry, a discovery mechanism. A project declares a path and a hash. Nothing
searches for gates, nothing installs them, and `xeno init` does not write an example declaration
into a project's configuration, because an example declaration is a command somebody has to
delete.

Anything about what an external gate should check. The one in the tests exists to be run, not to
be imitated.

**Two boundaries worth naming.** The command runs during `gate verify` as well as during
`gate run`, because it is a gate and `verify` recomputes verdicts — so a project that enables one
has foreign code in its CI verification step, which is a consequence of the design rather than of
this implementation. And an external gate's finding ids are derived from its own wording, so a
project that changes its tool's messages changes every id it produced, which is exactly why
section 4 makes the decision not survive.

<!-- xeno:section:context-rationale -->
## Why this context

Section 14's external gates paragraph is the specification, read with the paragraph that follows
it on why the marking is the point, because the second is what decides where execution lives.
Appendix A's `external_gates` block is read for the four fields and for the comment on the hash,
"checked before every run", which is a frequency and not a cache. Appendix A's table is read for
the default, "none run", which is what a project that declares nothing gets.

Section 5 is read for `provenance` on a check, and section 4 for the exception that external
findings make to decisions surviving a re-run, which is the behaviour already implemented and
about to be exercised for the first time.

Section 7 is read for where a gate's result comes from and for `G-Complete`, which section 14 says
accepts an external check like any other — and does, because it reads the checks it is given.

`internal/gates/gates.go` is read for `ExternalProvenance`, `carryForward`, `Invariants` and
`Status`, which between them already define everything an external check has to satisfy. The
reading established that this piece adds no rule: it adds a producer for rules that exist.

`internal/runner/runner.go` is read for `compute`, which is the one place a verdict is assembled,
and for the comment there about a second path into the invariants.

`internal/git/git.go` is read as the precedent for a subprocess in this project: fixed arguments,
failure reported rather than panicked, and the reason written in the package comment. An external
gate is the same shape with the opposite trust, so the differences have to be deliberate.

`internal/secrets` is read for how a project's own file is loaded beside a shipped one, since the
declaration is the first thing a project writes that makes the runner execute something.

`.xeno/config/project.yaml` and `internal/scaffold/files/project.yaml` are read for how a block is
documented where a project will meet it, because a declaration that runs code deserves the comment
rather than the schema.

Nothing outside the repository is needed, and the one thing that would have helped — an external
gate somebody wrote — does not exist, which is why the contract has to be written down rather than
inferred from a user.
