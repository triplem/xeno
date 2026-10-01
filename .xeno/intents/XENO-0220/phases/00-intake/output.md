---
intent: github.com/triplem/xeno#165
phase: 00-intake
created: "2026-10-01T16:21:45Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+46d4cb2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 623ad5f80509828133015f2559fd8231e45bd9d6f460ca7e515d8d8d0c6220ce
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

Three pieces of WP4 are built and nothing has ever applied a rule. #158 resolves a tree, #160
counts the answers to a review rule, #162 evaluates five predicates, and `given/builtin/` is
empty, so in this repository and in every other one the machinery runs over nothing. The gates
are green because there is no rule set, which is the weakest kind of green there is.

**The set the plan names is four rules and none of them exists.** Traceability of deviations,
migration notes for interface changes, a rationale for new dependencies, a filled release notes
section at P5, "deliberately minimal and generic". Also three examples nobody should inherit
unasked — Conventional Commits, the `Xeno-Intent:` trailer, the signature rule — and two hook
templates with their limits written on them.

**A set that ships and does not arrive is not shipped.** `vendorPlugin` copies
`.xeno/plugin/templates/` and nothing else, so `xeno init` into a project would leave the rules
behind. That has been true since WP9 and nobody noticed, because until this week there were no
rules to leave behind. A project that ran `init` today and expected the shipped set would get a
repository with a template set, an empty rule tree, and a green G-Rules telling it nothing.

**Three of the four rules cannot be `checked` and it is not the predicates' fault.** The six
shipped templates have no `interfaces` section, no `migration-notes` section and nothing about
dependencies. `section-implies-section` reads sections of the phase's artifact, so a rule about
an interface change has nothing to read, and the same goes for the dependency rationale and for
the traceability of deviations. They are `review` rules, which is what section 9 calls "the
honest admission that the rest needs a person" — and it means the first set anybody sees is
three quarters prose.

**The fourth needs a predicate nobody wrote.** "A filled release notes section at P5" asks
whether one section is non-empty. `section-implies-section` needs an antecedent, and nothing
else exists. Nothing else checks it either: `Missing` is read only by the next-step suggestion,
so a phase whose required section is empty passes every gate. Section 9 says this work package
delivers "the section predicates the shipped set needs", which is plural and scoped to the set,
so the absence is a gap in this package rather than a wall.

**And the set has to be read as a stranger would.** Two reviews in a row have ended by saying
their findings had been read only by their author, both naming this piece as where that stops.
Four rules, three examples and two hooks are the first text in this project written to be acted
on by somebody who did not build it, and the wording of what a rule says when it goes red is
what they will act on.

<!-- xeno:section:scope -->
## Scope

**In scope.** The four rules under `.xeno/plugin/rules/given/builtin/`, generic and minimal.
`section-non-empty` as the one addition to the predicate registry, because the set needs it and
section 9 scopes the section predicates to what the set needs. `xeno init` vendoring the rule tree
beside the templates. The three examples under `examples/rules/` and the two hook templates under
`examples/hooks/`, each carrying its limits in its own text. The call the plan leaves here, on
whether documentation consistency ships enabled or as an example, answered in the register. And
this repository answering its own review rules in its own P5, which is the first time the process
applies a rule to itself.

**Out of scope, and each for its own reason.**

External gates. The last piece of WP4, dependent on none of this.

A section for interfaces, migration notes or dependencies in the shipped templates. Adding one
would make three review rules checkable and it is a change to the template set, which is WP3's,
and to what every phase of every project renders. It is a finding about the shipped templates and
goes in writing rather than into this intent.

Any rule specific to a supplier, a customer or a project. Section 9 puts those in that party's
own layer and says the package is where they do not belong.

An `enabled` field, a default-off state, or anything for G-Rules to read as a switch. A rule that
is not in the tree does not apply, which section 9 states twice and which is what makes the
examples mechanism work.

The documentation site. Whatever the call above decides, the pages are WP16's.

**One boundary worth naming.** The shipped set lives with the plugin and is therefore covered by
the plugin digest G-Supply will check — and G-Supply is `not-implemented`, so today nothing
detects a project editing the shipped rules in place. That is a gap this piece inherits rather
than creates, and the honest version of the sentence "cannot be edited in place" is "is not meant
to be, and nothing yet notices".

<!-- xeno:section:context-rationale -->
## Why this context

Section 9's shipped-rules paragraphs are the specification for what goes in the set and what does
not: minimal and generic, the four subjects, the examples mechanism with its four consequences,
and the sentence that a rule not in the tree does not apply. The precedence table is read again
for what `given/builtin/` means — every project everywhere — because that is the reach of
everything written here and the reason the set has to be dull.

WP4 of the plan is read for the two sentences this piece has to answer rather than inherit:
documentation consistency as "a candidate for the shipped set rather than a topic of its own",
and the open call on whether those rules ship enabled or as examples. Its paragraph on
`examples/hooks/` is read for what each hook is and what it cannot do, since the limits go in the
files themselves.

The six shipped templates are read section by section, because a `checked` rule can only name
sections that exist and that reading is what moved three of the four rules to `review`.
`internal/template/template.go` is read for `Missing`, to confirm that nothing enforces a filled
required section, which is what makes the fourth rule meaningful.

`internal/gates/gates.go` is read for the registry as #162 left it and for
`sectionImpliesSection`, since `section-non-empty` is the same reader with no antecedent.
`internal/runner/init.go` is read for `vendorPlugin`, which is the piece's one change outside the
rule tree.

`internal/rules` is read as the judge of everything written here: the four rules have to pass the
gate that reads them, and writing a rule file by hand against that package is the first
independent reading its findings have had.

`CONTRIBUTING.md` is read for what this project asks of a commit, because two of the three
examples — Conventional Commits and the `Xeno-Intent:` trailer — describe conventions this
repository either follows or deliberately does not, and the examples should not contradict the
project that ships them.

Nothing outside the repository is needed. The one thing that would have helped — another
project's rule tree to compare against — does not exist, which is exactly why the set has to be
written to be read by a stranger rather than to be convenient here.
