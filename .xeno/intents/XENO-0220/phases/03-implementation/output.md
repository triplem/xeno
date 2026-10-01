---
intent: github.com/triplem/xeno#165
phase: 03-implementation
created: "2026-10-01T16:34:06Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4f94129.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: aa2023e14ab6ca395f7a8b8de1884ffe94fcb9012a6de14f9b87315f53a4d092
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

Four rule files, four examples, two hooks, one predicate, one vendoring change, three register
rows, and the guard that keeps a new rule set from rewriting the trail behind it.

**`.xeno/plugin/rules/given/builtin/`, four rules.** `deviations-are-traceable` (review, P3),
`interface-change-needs-a-migration-note` (review, P2 and P3, with a comment saying section 9
writes it as checked and why it cannot be here), `new-dependency-needs-a-rationale` (review, P2
and P3), `release-notes-are-filled` (checked, P5, `section-non-empty` over `release-notes`). Each
carries a comment saying it is shipped, pinned and not edited here, and that a project that
disagrees overrides it, because precedence replaces a rule and never deletes one. None is
binding, so a project can always override.

**`section-non-empty`.** `sectionNonEmpty` beside `sectionImpliesSection`, reading one `section`
parameter. Two helpers came out of writing it: `unknownSection`, which reports a rule naming a
section the phase's template does not have and which the implication type now shares, and
`renderedSections`, which parses the artifact back into sections. The implication type lost nine
lines to those two.

**The guard, and the reason it exists.** Writing the set turned all twenty-four sealed P5 phases
of this repository red on `gate verify`: none of them carries the checklist entries the three new
review rules require, and `gate verify` recomputes a sealed verdict and compares it. So `policy`
now reads the phase's own `rules_hash` and judges nothing where it is absent, the placeholder, or
another hash. A74 records it, with the measurement rather than the reasoning, because the
reasoning alone would have sounded like caution.

**`internal/runner/init.go`.** `vendorPlugin` is two calls now: `vendorTree` for the templates,
unchanged in behaviour, and `vendorRules` for `rules/given/builtin`, which is one level deeper and
absent in any plugin released before this piece — so a missing directory is not an error. Both go
through the same `create`, so a second `init` still changes nothing.

**`examples/rules/`, four files.** `conventional-commits`, `xeno-intent-trailer`,
`signed-commits` and `documentation-follows-the-change`. Each is a complete rule with `scope: org`,
so copying it into `given/org/` is the whole adoption, and copying it into `given/project/`
produces a `scope` finding naming both the claim and the path. Each carries in its own text what it
costs and what it does not check: the trailer example says the value is never read and why, the
signature example says what counts as valid and that trust lives in the keyring of whoever runs the
gate, and the documentation example carries A72's reasoning for why it is an example at all.

**`examples/hooks/`, two templates.** `prepare-commit-msg` seeds a `Refs #<issue>` footer from a
branch named `<issue>-<slug>`, leaves a message with any source alone, and says in its own text
that it is comfort, that it is bypassable, and that it validates nothing. `pre-receive` walks the
pushed range, skips deletions and non-branch refs, judges every non-merge subject by calling
`xeno check commit-message`, and says in its own text that it is enforcement, that a host with push
rules should use those instead, that it needs the binary on the server, that a copy of the pattern
has to be kept in step where there is none, and that it judges subjects rather than artifacts. Both
say installation is `core.hooksPath` and that `xeno init` does not touch git configuration.

**Tests, `internal/rules/shipped_test.go`.** The set is read where it lives: every file reaches the
effective set, no statement names a language, a tool or a project from a list of twelve, every rule
is at `builtin` and claims it, none is binding, every statement is at least eight words and ends in
a full stop, every rule names a phase, and every checked rule names a type the registry has. Then
the examples: copied into `given/org/` they all resolve and all claim `scope: org`, and one copied
into `given/project/` produces exactly one finding naming both the claim and the path.

**Tests, `internal/gates`.** `section-non-empty` filled, empty, whitespace-only and absent; an
unknown section and a missing parameter; the shipped release notes rule read from where it ships
and evaluated green and red; the registry's six names with the count asserted so a seventh cannot
arrive without a reason; a phase recording another set not judged, and the same artifact recording
this set judged; and the placeholder and the absent field both leaving a phase unjudged.

**Tests, `internal/runner`.** `init --vendor` carries the rule tree, what arrives resolves with no
finding, the file count matches the rule count, and without `--vendor` nothing is copied.

**The fixtures now record the set they build.** Every fixture in `internal/gates` writes its rule
tree before its artifact and puts `rules_hash` in the frontmatter, because after A74 a fixture that
does not say which set applied is a fixture the gate has nothing to say about.

<!-- xeno:section:deviations -->
## Deviations from the design

**A74 is a whole mechanism the design did not have.** 02-design decided four rules, one
predicate and a vendoring change, and said nothing about what a new rule set does to phases
already sealed. Writing the set answered it: twenty-four divergences on `gate verify`, every
sealed P5 in the repository, because the review rules require entries those artifacts cannot
have. The guard, its reasoning, its test and its register row are all this phase's work, and the
design would have been better for asking the question — a rule set is the first thing in this
project that changes what an old artifact is judged by.

**The guard leaves a hole and the register says so rather than closing it.** A rule added after an
artifact was written leaves that artifact unjudged until something renders it again, because the
hash is written at render time. So a project could, in principle, write its P5 and then add a rule
and have the rule not apply. The honest mitigations are that the next section write re-records the
hash and that the two values are both visible in the artifact, and neither is a check. Recorded in
A74 and in the gaps rather than solved, because solving it means deciding whether a rule set change
makes a phase stale, which is G-Freshness's vocabulary and a specification question.

**The implication type was refactored while adding the second one.** `unknownSection` and
`renderedSections` came out of writing `section-non-empty`, and `sectionImpliesSection` now uses
both. That is a change to code #162 shipped last week, inside a piece whose subject is content,
and it is the right trade — two copies of a template lookup would drift — but it means this diff
touches a predicate nobody asked me to touch.

**The shipped set does not include the rule section 9 writes out in full.** Section 9's own rule
example is `public-interface-change-requires-migration-note`, a checked rule over an `interfaces`
section implying a `migration-notes` section. The shipped templates have neither section, so the
set ships the same statement as a review rule under a shorter id. A reader comparing the
specification with the set will find its one worked example missing, and the rule file says why in
a comment. The finding about the templates is in the gaps.

**Twelve forbidden words is a crude test.** `TestTheShippedSetNamesNothingSpecific` greps the
statements for a list of languages, tools and names. It would not catch a statement that is
specific without using one of those words, and it would fail a legitimate statement that happened
to need one. It is a tripwire rather than a check, and it is in the suite because the property it
guards — nothing specific reaches every project everywhere — has no mechanical form.

**The hooks are untested.** Both parse under `sh -n` and neither is exercised. `pre-receive` reads
a ref update protocol no test here produces, and `prepare-commit-msg` would need a real commit in
flight. They are examples carrying their own limits, and what a reader does with them is read them
first — but a shell example that was never run is a shell example that may not run.
