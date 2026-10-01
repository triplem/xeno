---
intent: github.com/triplem/xeno#160
phase: 02-design
created: "2026-10-01T15:33:07Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e32006ffc4ee12de0147fd7cc798bf337f7324399be4527fbbd148a3c78946d3
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**Completeness is counted over the review rules of the effective set, and the count runs from
the rules to the entries.** Section 12 forbids a lens from adding to or subtracting from the
counted set, which only holds if the set is the rules; and section 9's pairing of "every entry
carries a result" with "an omission nobody sees" only holds if a missing entry is a finding. A
gate that walked the entries and checked each had a rule would pass an empty checklist.

**Every review rule in the effective set is answered at P5, whatever its `applies_to`.** The
checklist exists once, in the P5 artifact, and section 9 says it is rendered from the effective
rule set rather than from a phase's slice of it. So for a review rule `applies_to` names the
phases whose work the rule is about, and the answer is given when the change is reviewed. The
alternative reading — only rules naming `05-review` are answered — would leave a review rule
about design answered nowhere, which is a rule that does nothing. Recorded as an assumption,
because the section can be read both ways and this piece has to pick one.

**A checked rule is reported, not evaluated, and the seam is a registry.** `internal/gates`
holds a map from predicate type to implementation; it is empty in this piece. A checked rule
that applies to the phase and whose type is not in the map produces a finding naming the rule
and the type. The next piece fills the map and the finding disappears for the types it fills.
Nothing else about a checked rule is said here, and in particular its absence from the map is
reported rather than treated as a pass.

**The registry lives with the gate rather than with the resolver.** A predicate reads an
artifact or a commit range, which is what a gate has and `internal/rules` deliberately does
not: that package walks files and sorts, and giving it artifacts to read would make the
resolver depend on the phase it is resolving for.

**The entry is a frontmatter list, `review_checklist`, beside the three that exist.** Hyphenated
section id, snake_case key, as `open-questions` and `open_questions` already do. A3 is the
assumption that put structured lists in the frontmatter and it covers this one; what it does not
cover is the key's spelling, which section 9 writes under the section's own name, so the
convention is followed and recorded rather than read out of the example.

**An entry carries `rule`, `result`, `note` and `source`.** `source` exists because section 12
requires a lens entry to be distinguishable, and the absence of `rule` is what the gate actually
keys on: an entry with no rule is not counted towards any rule, whatever its `source` says. So a
lens that forgot the field cannot accidentally answer a rule, and a lens that forged a rule id is
indistinguishable from an agent that wrote one — which is the honest limit, since both are
self-asserted strings in the same file.

**An entry naming a rule outside the effective set is a finding.** It asserts an answer to a
rule that does not apply, which is either a rule that was removed or an id that was mistyped,
and both are worth a verdict rather than silence. This is the one finding not written in section
9, and it is recorded in the register beside the completeness reading.

**The gate is quiet rather than absent at P0 to P4.** It resolves the tree, reports an
unimplemented checked type where one applies, and passes. The row says `pass`, because nothing
applied and that is a fact about the phase, not about the gate.

<!-- xeno:section:alternatives -->
## Alternatives

**Count the entries rather than the rules.** Walk the checklist, check each entry has a result
and names a known rule, and stop. It is what section 9's "every entry carries a result, and
nothing more" says if that clause is read alone, and it is one loop instead of two. Rejected
because it passes an empty checklist, and because section 12's clause about a lens would then
have nothing to protect: a set defined by the entries present is a set a lens can add to by
definition.

**Treat an unimplemented predicate type as a pass until the predicates land.** The gate would be
green today in every repository and the next piece would turn it red where a type was missing,
which keeps this piece from making a checked rule unusable. Rejected because that green is
exactly the claim section 16 catalogues: a rule in force that nothing evaluated. The cost is
real — a project adopting a checked rule before the next piece gets a red gate — and it is the
right cost, because the alternative is a project believing a check ran.

**Report the unimplemented type as the gate's result rather than as a finding.** `G-Policy` could
read `not-implemented` where the set contains a type it cannot evaluate, which is the vocabulary
the four unbuilt gates already use. Rejected: `not-implemented` is a statement about the runner,
and this is a statement about one rule in one project's tree. A gate that flipped to
`not-implemented` because of a project's rule file would also hide the checklist half, which is
implemented and has findings to report.

**Filter review rules by `applies_to` naming `05-review`.** It gives `applies_to` an effect for
review rules and makes the set smaller and more explicit. Rejected because a review rule about
design would then be answered nowhere, and because section 9 says the checklist is rendered from
the effective rule set. It is the reading this piece is most likely to be argued with over, which
is why it is in the register rather than only here.

**Put the predicate registry in `internal/rules`.** Everything about rules in one package, and
the resolver could report an unimplemented type itself. Rejected because a predicate reads an
artifact or a commit range, and the resolver reads neither; the dependency would run the wrong
way and make the package that must be callable from two gates depend on what a gate holds.

**Derive the `review-checklist` section from the entries, for this list only.** Section 9 says
the section is rendered from the rule set, and doing it here would make the P5 artifact's
checklist impossible to write inconsistently with its frontmatter. Rejected for the reason in the
non-goals: three other lists are in the same position, A3 claims derivation for all of them, and
one of four derived is a worse state to leave than none. The issue it needs is named in the
residual risk.

<!-- xeno:section:impact -->
## Impact

**A review rule stops being a suggestion.** After this a project can write a review rule and the
gate will require an answer to it at P5, with a note where the answer is not `met`. That is the
whole mechanism section 9 describes for the part of a rule set a deterministic gate cannot
evaluate, and it is the first thing in this repository that makes a person's judgement a
recorded step rather than a habit.

**A checked rule becomes unusable until the predicates land, visibly.** Adopting one now turns
the phases it applies to red with a finding naming the type. That is a deliberate consequence
of refusing the silent pass, and it is the strongest argument for the predicate piece being
next.

**Three quiet rows instead of four.** G-Supply, G-Secret and G-Test remain `not-implemented`.
A phase here will report eleven gates running of fourteen.

**`output.md` gains a fourth structured list**, so the frontmatter order grows by one key and
`internal/model`'s `Output` by one field. Nothing that reads the existing three changes, and an
artifact without the key is an artifact with no entries, which is what every artifact in the
trail is.

**The resolver gets its second caller, which is what its shape was argued for.** #158 returned a
set rather than a verdict on the strength of this gate existing; this piece is where that claim
is paid or not. It is paid: G-Policy calls `Load` and `Effective` and adds no resolution of its
own.

**Nothing in the trail moves.** No set changes, so no `rules_hash` changes; no artifact is
rewritten; `gate verify` stays at exit 0. The phases of this intent are the first in the
repository judged by eleven gates.

**The next piece inherits a seam and a bill.** The registry is one map and the finding it
produces is the list of types it has to fill: `section-implies-section` plus the four that read
a commit range. What it must not do is widen the map's key — a type is a name from the
specification, and a project naming its own would be the project-defined predicate v1 does not
have.
