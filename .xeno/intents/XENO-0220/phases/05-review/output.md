---
intent: github.com/triplem/xeno#165
phase: 05-review
created: "2026-10-01T16:37:33Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4f94129.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a2e14b1b722a4c2d8feee1698fa3c9a696d307660e895b09ce80f039662806be
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
      The six deviations of 03-implementation each name what they depart from: the design, A74,
      #162's predicate, section 9's rule example, the criterion the forbidden-word test stands in
      for, and the acceptance's reading of the hooks.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      G-Policy's behaviour changes for every project and A74 says what the new behaviour is. A
      dependant does nothing: an artifact written under another rule set stops being judged
      against this one, which is the point rather than a cost.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      No dependency was added. The project still carries one, go.yaml.in/yaml/v3, vendored, and
      the hook templates call the binary so that they add nothing of their own.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

This is the first review checklist in this repository with entries in it, because this is the
piece that shipped the rules. The structured answers are in the frontmatter, under
`review_checklist`, which is what G-Policy counts; what follows is the reasoning behind each one
and the questions the change raises beyond them.

**`deviations-are-traceable` — met.** The six deviations in 03-implementation each name what they
depart from: the design that did not have A74, the hole A74 leaves, the predicate #162 shipped
that this piece refactored, the rule section 9 writes out and this set does not ship, the
criterion the forbidden-word test stands in for, and the hooks the acceptance asked to be read
rather than run. Answering this rule against my own phase is also the first evidence that it is
answerable, and it was: the rule asks for a reference and the references were there.

**`interface-change-needs-a-migration-note` — met.** This change alters two things something
outside the intent depends on: the predicate registry gains a name, which is additive and needs
no migration, and G-Policy's behaviour changes for every project, which does. A74 says what the
new behaviour is, the release notes below say what a reader does about it, and the answer is that
a dependant does nothing — an artifact written under another rule set simply stops being judged
against this one, which is the point rather than a cost.

**`new-dependency-needs-a-rationale` — not-applicable.** No dependency was added. The project
still carries one, `go.yaml.in/yaml/v3`, vendored, and the hook templates call the binary rather
than anything else precisely so that they add nothing.

**Beyond the three rules.** Four questions this change raises that no rule asks:

*Is the set refusable?* The bar for `given/builtin/` is not whether a rule is good but whether a
project could reasonably refuse it. Three statements ask a person to confirm something they
probably did; the fourth asks that a section carry text. None of them is a practice, a language or
a style. Answered yes, and it is why the set is thin.

*Does the guard protect the trail or hide a rule?* Both, and the second is in the gaps. It
protects twenty-four sealed verdicts from being rewritten by a rule adopted today; it also means a
rule added after an artifact was written does not reach it until something renders it again. The
measurement is in the results and the hole is in the register.

*Was anything written that the documents forbid?* No field, no gate, no rule outside the set's
own layer. One predicate type, admitted by section 9's own scoping of the section predicates to
what the shipped set needs, recorded as A73. The plan's open call is answered as A72 rather than
deferred.

*Is the wording answerable by somebody who did not write it?* Two statements were rewritten for
it, which the verification records. The honest limit is that I answered them, so what has been
tested is that they can be answered and not that two people answer them the same way.

<!-- xeno:section:release-notes -->
## Release notes

**A rule set ships.** Four rules under `given/builtin/`, which every project receives and no
project can delete — only override, because precedence replaces a rule rather than removing it.
Three are answered by a person at review: every deviation names what it departs from, an interface
change says how a dependant moves, a new dependency says what it is for and what was weighed. The
fourth is checked: the release notes section of the review artifact carries text.

**Three examples and a fourth, adopted by copying.** `examples/rules/` carries Conventional
Commits, the `Xeno-Intent:` trailer, the signature rule, and documentation following the change.
There is no `enabled` field: a rule that is not in the tree does not apply, so adoption is `cp`
into `given/org/` or `given/project/` with the scope corrected to the level. Copy one to the wrong
level and the gate says so, naming both the claim and the path.

**Two git hook templates.** `examples/hooks/prepare-commit-msg` seeds a reference from the branch
name and says in its own text that it is comfort and bypassable.
`examples/hooks/pre-receive` rejects a push whose subjects do not match a shipped pattern, by
calling `xeno check commit-message` rather than carrying a copy of it, and says what it needs on
the server and what it does not judge. Installation is `core.hooksPath`; `xeno init` still does
not touch a developer's git configuration.

**`xeno init --vendor` now carries the rule tree** beside the templates. A project that vendored
the plugin before this release gets the rules on its next `init`; a plugin without a rule tree is
not an error.

**`rules_hash` starts meaning something.** Every artifact written from now on records a hash over
four rules rather than over an empty rendering.

**A phase is judged only against the rule set it recorded.** Adopting a rule today does not
re-judge phases sealed before it, which is what kept this release from turning every existing P5
red. An artifact that records another set, the placeholder, or nothing is left alone (A74).

**`section-non-empty`** is a new predicate type: one section of the phase's artifact has to carry
text. It is what the release notes rule asks for (A73).

**The documentation consistency rules ship as an example rather than in the set**, which is the
call WP4 left to this work package. A72 says why: the mechanical half has nothing to point at
until the documentation site exists, and a review rule in the shipped set would be answered
`not-applicable` for ever by every project whose documentation lives elsewhere.

Closes #165. Refs #1.

<!-- xeno:section:residual-risk -->
## Residual risk

**The guard has a hole, stated and not closed.** A rule added after an artifact was written does
not reach that artifact until something renders it again, because `rules_hash` is written at
render time. The next section write re-records it and both values are in the file, and neither of
those is a check. Closing it means deciding whether a rule set change makes a phase stale, which
is G-Freshness's vocabulary and a specification question — the first one this package has raised
that is about the trail rather than about rules.

**The specification's one worked rule example cannot be shipped as written.** Section 9 writes
`public-interface-change-requires-migration-note` as a checked rule over an `interfaces` section
implying a `migration-notes` section, and the shipped templates have neither. The set ships the
same statement as a review rule. A reader comparing the document with the package will notice,
and the rule file says why. The fix is two sections in two templates, which is WP3's and a change
to what every phase of every project renders.

**Nothing stops a project editing the shipped set in place.** G-Supply is `not-implemented`, so
the sentence in section 9 — the set "cannot be edited in place in a project" — is true of intent
and not of mechanism. Each file says it is shipped and pinned; that is a comment.

**The set is thin, and that is the deal rather than a defect.** Nothing in these four rules would
have caught any finding this repository has recorded in twenty intents. Three of them ask a person
to confirm something they probably did anyway. Whether a minimal generic set earns its place is a
question the first adopter answers and this suite cannot.

**The review rules have been answered once, by me.** This intent's own checklist is the first, and
two statements were rewritten in the course of answering them, which is the evidence that matters.
What has not been tested is whether two people answer them the same way, and that is the part the
closed loop still hides.

**The hooks were read and never run.** Both parse; neither has been executed, because one reads a
ref update protocol nothing here produces and the other needs a commit in flight. An example that
was never run may not run, and a reader who copies it finds that out first.

**A fifth and sixth rule will be harder than these four.** The easy generic statements are now
taken, so the next addition to the set is either specific — which the layer forbids — or weaker
than what is there. The honest reading of "deliberately minimal" is that four may be the whole set
for v1, and the examples are where the rest belongs.
