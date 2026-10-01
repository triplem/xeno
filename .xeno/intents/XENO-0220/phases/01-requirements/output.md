---
intent: github.com/triplem/xeno#165
phase: 01-requirements
created: "2026-10-01T16:22:50Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+46d4cb2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: caa2de485d0d465a24ecbc6832fe7811fb55574dc554ab4920b091c10edd1ab2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**The set is four rules and every one of them passes the gate that reads it.** Each file resolves
under `given/builtin/`, declares `scope: builtin`, carries a statement a person can answer
against, and produces no finding from G-Rules. The set is checked by running the gate over this
repository rather than by inspection.

**None of them names a supplier, a customer or a project.** The statements are about artifacts and
about the process, not about Go, not about this repository's layout, and not about anything a
reader would have to work here to understand.

**`release-notes-are-filled` is evaluated and goes red over an empty section.** A P5 artifact whose
`release-notes` section carries nothing is red, naming the artifact and the section; one with text
is green. This is the only `checked` rule in the set and the one that proves the machinery of #162
against a real rule file.

**`section-non-empty` is in the registry and named nowhere else.** One new predicate type, reading
one section of the phase's artifact, with a configuration error where the section is not in the
phase's template, matching what `section-implies-section` already does.

**The three review rules produce checklist entries, and this repository answers them.** At this
intent's own P5, G-Policy requires an entry per review rule of the effective set, and the P5
artifact carries three answers with notes where they are not `met`. That is the first time the
process applies a rule to itself, and it is a criterion rather than a side effect: if the rules
are unanswerable, writing the answers is where it shows.

**`xeno init` carries the rule tree.** An `init --vendor` into an empty repository produces
`.xeno/plugin/rules/given/builtin/` with the four files, and the rules resolve there with no
finding. Running it twice changes nothing the second time, which is WP9's own criterion and must
keep holding.

**The examples are adopted by copying and by nothing else.** Three files under `examples/rules/`,
each a complete rule with `scope: org`, each naming a shipped predicate type and pattern. Nothing
in the runner reads `examples/`, no `enabled` field exists anywhere, and a test asserts the
examples are well formed when copied into a tree rather than where they live.

**The two hook templates carry their limits in their own text.** `prepare-commit-msg` says it is
comfort and bypassable; `pre-receive` says it is enforcement, that it is the only kind there is on
a host without push rules, and that a copy of a pattern on a server without the binary has to be
kept in step. Installation is `core.hooksPath` and both say so. No test asserts prose, so this
one is read rather than checked.

**The documentation consistency call is answered in the register**, with what it decides and what
the alternative was, because the plan leaves it to this work package and an unanswered call would
be the fourth piece deferring it.

**Every finding these rules can produce is read as a stranger would read it.** The three review
statements, the checked rule's red message, the configuration errors for a missing section, and
the examples' own text. Where a wording only makes sense to somebody who built this, it is
rewritten. This is a judgement, recorded as one, and it is the criterion two earlier reviews
deferred here.

**Nothing else moves.** `rules_hash` changes for every artifact written after the set exists, which
is correct and is the first time that field has carried anything but the hash of an empty
rendering. No sealed artifact is rewritten, `./xeno gate verify` stays at exit 0, and the whole
suite stays green.

<!-- xeno:section:non-goals -->
## Non goals

**No new section in any template.** Three review rules would be checkable if the design and
implementation templates carried an `interfaces` section and a `migration-notes` section. Adding
them changes what every phase of every project renders, which is WP3's, and it is written down as
a finding rather than done here.

**No sixth predicate beyond `section-non-empty`.** The registry stays the specification's names
plus the one the set needs, and a project naming its own type still gets #160's finding.

**No rule about this repository.** Nothing about Go, `gofmt`, the 88 column rule, SPDX headers or
the layout of `internal/`. Those are this project's conventions and they belong in its own
`given/project/`, which this piece does not create either: the set is what every project
everywhere inherits.

**No rule that needs a host.** `approver-not-author` is in the examples rather than the set,
because a project without a review culture would be red from its first phase on something no
document asked it to adopt.

**No `enabled`, no switch, no default-off.** The examples are adopted by copying.

**No documentation site, no generated page.** Whatever the documentation consistency call decides,
the pages are WP16's and this piece writes none.

**No change to G-Supply.** The set lives with the plugin and nothing yet checks that a project has
not edited it in place. That gap is inherited and stated, not closed.

**No backfill.** Artifacts already sealed keep the `rules_hash` they carry, which for everything
written this week is the hash of an empty rendering. That value was correct when it was written and
rewriting it would change sealed verdicts.

<!-- xeno:section:constraints -->
## Constraints

**`given/builtin/` reaches every project everywhere**, which is the whole constraint on what may be
written there. A rule that is wrong for one project is wrong in the package, and there is no
mechanism for a project to remove one: precedence overrides, it does not delete. So the bar is not
"is this a good rule" but "is this a rule no project could reasonably refuse".

**A `checked` rule may only name sections that exist in the templates of the phases it applies
to.** This is what decides three of the four kinds, and it is a constraint rather than a
preference: a rule naming a section nobody has is a configuration error that shows only when
somebody adopts it.

**The statements are read by a person and have to say what to do.** A review rule's whole output is
its statement in a checklist, so a statement that describes a property rather than asking a
question produces an entry nobody can answer.

**Section 9's four consequences of shipped patterns hold for shipped rules too.** The failure
message, comparability across projects, cost and behaviour, provenance. A rule written loosely here
is a finding fifty projects cannot compare.

**This repository has to answer its own review rules**, in this intent's P5, under the rules this
intent adds. If that is unpleasant, the rules are wrong, and it is the cheapest review available.

**`xeno init` must stay idempotent.** WP9's criterion is that running it twice changes nothing the
second time, and the vendoring change must not break it.

**One dependency, 88 columns, SPDX on Go files, `gofmt`, `go vet`, the suite, and
`./xeno gate verify` at exit 0.** The hook templates are shell and carry no SPDX line, matching
`scripts/`.

**One intent, one branch, one issue.** `165-the-shipped-rule-set`, #165, labelled wp4, branched off
a `main` that carries all three earlier pieces — no stacking this time.
