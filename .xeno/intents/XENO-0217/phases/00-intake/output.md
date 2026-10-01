---
intent: github.com/triplem/xeno#158
phase: 00-intake
created: "2026-10-01T14:58:53Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a9a037652b76487b9306843a208685d14b37d5b28ab11b8e25bbc478b695433a
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

WP4 is the substance of M1 and nothing of it exists. `G-Rules` and `G-Policy` stand as
`not-implemented` in the gate table, there is no `internal/rules`, and no `given/builtin/`
tree. Section 9 of the process definition fixes a format, two axes, a precedence order and a
set of configuration errors, and none of it is expressed anywhere in this repository.

**The rule set is the one part of the process with no representation at all.** Every other
object the documents define has a type: artifacts, gates, findings, evidence, assumptions,
the cost record, the symbol index. A rule has none, so there is nothing to validate, nothing
to resolve and nothing for two gates to judge. The gate table says G-Rules checks "rule
collisions resolved, `binding` respected, `scope` matches the path" from P0, and that row has
been a promise since the table was written.

**`rules_hash: by-hand` is in every artifact in this tree.** Seventy-one intents carry it,
written by the harness because no code computes one. It sits inside `artifacts_hash`, so it
is part of every verdict this repository has recorded, and what it currently asserts is that
a person says a rule set was in force. Appendix B fixes every other hash to the byte and
delegates this one to the package that writes it, which is this package.

**Three things the documents require and the tree cannot express.**

*The file.* `id`, `version`, `scope`, `binding`, `kind`, `applies_to`, `statement`, `check`,
and `abstract` where a learned rule above the project level requires it. Section 9 also makes
two combinations configuration errors rather than oddities: `checked` without a `check`, and
`review` with one.

*The reach.* The path decides the level and `scope` repeats it, checked against the path, so
that a moved file cannot silently change which projects a rule applies to. `binding` is
allowed only under `given/`, is rejected under `given/project/` where it would have nothing
left to bind, and `learned/builtin/` does not exist at all because nothing learns into a
pinned, hashed package.

*The resolution.* More specific beats less specific, `given` beats the matching `learned`, and
`binding` above both. Where two binding rules on different levels collide the gate is red and
stays red, which section 9 states as a matter between two organisations and not something a
gate may resolve. That last one is the only precedence rule that is also a refusal, and it is
the one a resolver written without reading the section would get wrong.

**What makes this piece first rather than the predicates or the checklist.** A predicate
evaluates against a rule that was resolved; a checklist renders from an effective set. Both
need a set to exist, and neither can be judged before one does. The order is not a preference.

<!-- xeno:section:scope -->
## Scope

**In scope.** `internal/rules`: the rule file as a type, the tree walked across both axes,
every configuration error section 9 names, precedence resolution into an effective set, and a
byte-exact `rules_hash` over that set with its definition recorded as an assumption, as A62
did for `secrets_hash`. `G-Rules` implemented in `internal/gates`, from P0, replacing its
`not-implemented` row. The runner writing the hash it computes in place of the harness's
`by-hand`. Tests for each finding and for the resolution.

**Out of scope, and each for its own reason.**

The predicate types. A `check` is read as a declared type with parameters and evaluated by
nothing here. `section-implies-section` reads an artifact and the four commit types read a
range, and both are a different question from whether a set resolves — a rule set that does
not resolve has no business running a predicate. Their own issue under wp4.

G-Policy and the P5 checklist. The checklist is a rendered section, so it reaches into the
template set and the renderer as well as the gates, and section 9 ties what G-Policy counts to
what a lens entry carries. It is the next piece and it is deliberately not this one.

The shipped set under `given/builtin/` and the examples. Four generic rules, the two patterns,
`examples/rules/` and the two hook templates are content, and content written before the
reader exists is content written against a guess. This piece ends with a reader that finds no
rules in this repository and says so by staying green.

External gates. Declared in `project.yaml` with path and hash, hash match or refuse, JSON in
and out. They share the work package and nothing else: no rule file, no precedence, no
`rules_hash`. They can go before or after this and the order does not matter.

Whether the documentation consistency rules ship enabled or as examples. The plan explicitly
leaves that call to this work package, and it is a call about the shipped set, which is the
piece after next. Naming it here would be deciding it in the intent that cannot test it.

**A boundary worth stating.** This piece changes what `rules_hash` means in every artifact
written after it: today the field is a person's word, afterwards it is a hash over a resolved
set. Artifacts already sealed keep what they carry, because the field sits inside
`artifacts_hash` and rewriting one would change every verdict in that intent. So the trail
will carry two meanings of the same field, divided by this commit, and that division has to be
written down rather than left for a reader to notice.

<!-- xeno:section:context-rationale -->
## Why this context

Section 9 of the process definition is the whole specification for this piece and is read
entire: the two axes and the table of levels, the rule format with every field, why
`version` claims nothing, where `binding` and `abstract` are allowed, the two kinds and the
two configuration errors, the precedence list and its reasoning, and the sentence that makes
a binding collision a refusal rather than a resolution. It is read as the authority and not
as background, because every finding this piece reports is named there.

Section 7 is read for G-Rules's row — what it checks and that it applies from P0 — and for
why G-Rules and G-Policy sit together at the end of the table. That last point is read as a
constraint on the next piece rather than this one: the tree should be resolved once, so the
resolver this piece writes is the one G-Policy will call, which is an argument for returning
an effective set rather than a verdict.

Appendix B is read for what it fixes and what it delegates. It fixes every hash to the byte
except the ones whose inputs are assembled by code, and A62 is the precedent for how that
delegation is answered: a canonical rendering, sorted, with the rendering written down in the
register rather than left in the function.

`internal/gates/gates.go` is read for the gate table, the `spec` shape with its `from` column,
`notImplemented`, and how an existing gate reports a finding with file, cause and next step.
`internal/secrets/secrets.go` is read as the closest structural analogue in the tree: a set
assembled from a shipped layer plus a project layer, hashed canonically, absent without error.
`internal/template/template.go` is read for how the plugin layer and the project layer are
resolved against each other, since the rule tree has the same shape with four levels instead
of two.

`internal/model/model.go` is read for how a type that appears in an artifact is declared, and
for `HashFields` and `WriterlessHash`, which is where `rules_hash` is currently listed as a
field no writer produces. That list is the tree's own record of this gap.

`.xeno/config/project.yaml` is read for what a project declares today, because the rule tree
needs no configuration at all — the paths are fixed by section 9 — and confirming that is what
keeps a configuration field from being invented.

Nothing outside the repository is needed. The format is Xeno's own and fixed by its own
specification, so there is no host to ask, and the one thing that could have come from outside,
a second project's rule tree, does not exist yet.
