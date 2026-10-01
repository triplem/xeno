---
intent: github.com/triplem/xeno#165
phase: 02-design
created: "2026-10-01T16:24:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+46d4cb2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fe47bd10d628135679c86db23e83da99ca958c09ca4df7e9f6ad878c97ace407
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

**The four rules, with their kinds following the templates rather than the statements.**

*`deviations-are-traceable`*, review, P3. "Every deviation this phase records names what it
departs from: a decision, an assumption, or an acceptance criterion of this intent." The
implementation template has a `deviations` section, so a reader has something to check against,
and what makes it unanswerable by a gate is the word "names": a deterministic check could count
entries and not whether any of them points at anything.

*`interface-change-needs-a-migration-note`*, review, P2 and P3. "Where this change alters an
interface something outside this intent depends on, the artifact says how a dependant moves to
the new one." Section 9 writes this as its own `checked` example against an `interfaces` section
and a `migration-notes` section, and the shipped templates have neither, so here it is a review
rule. The phrase "something outside this intent depends on" is what keeps it from firing on every
internal rename.

*`new-dependency-needs-a-rationale`*, review, P2 and P3. "Where this change adds a dependency, the
artifact says what it is for and what was weighed against adding it." Deliberately not "whether
the dependency is justified", which is a judgement, and not a count of `go.mod` lines, which is a
language.

*`release-notes-are-filled`*, checked, P5, `section-non-empty` over `release-notes`. The one
mechanical statement of the four, and the one that proves #162's registry against a real rule
file rather than a fixture.

**`section-non-empty` is `sectionImpliesSection` without an antecedent.** One parameter, `section`,
the same configuration error where the phase's template does not have it, the same reading of the
rendered artifact. It is the second entry in the registry that section 9 does not name in its
table, and the sentence that admits it is the same one: this work package delivers "the section
predicates the shipped set needs".

**Three review rules and one checked rule is the honest ratio, and it is a finding about the
templates rather than about the predicates.** The three would be checkable if the design and
implementation templates carried an `interfaces` section and a `migration-notes` section. That is
WP3's to change and it is written down rather than absorbed.

**Documentation consistency ships as an example, not in the set.** The plan leaves the call here.
The mechanical part it describes — a `checked` rule over derivable structure, with G-Freshness
providing the hash binding — has nothing to point at until WP16 generates the pages it would
check, and the hash binding that package specifies is for "a few high value pairs only". A
documentation rule in `given/builtin/` would also reach every project everywhere, including every
project whose documentation lives somewhere other than the repository, and that is precisely the
case the examples mechanism exists for. So `examples/rules/documentation-follows-the-change.yaml`,
a review rule, adopted by copying. Recorded with its reasoning, because the plan asked for a call
and not for a deferral.

**`xeno init` vendors the rule tree the same way it vendors the templates.** One more loop over one
more directory, with the same `create` so that a second run changes nothing. The alternative — a
project copying rules in by hand — would make the shipped set something a project assembles rather
than something it receives.

**The examples are complete rules, not fragments.** Each is a file somebody can copy into
`given/org/` and have resolve on the first run, which means each carries `scope: org`. A reader
who copies it into `given/project/` gets a `scope` finding naming both the claim and the path,
which is the gate teaching the layout better than a comment in the file could.

**The hooks call the binary rather than carrying a pattern.** `pre-receive` runs
`xeno check commit-message` per commit, so the expression lives in one place, which is the reason
that command exists. Both templates carry their limits in their own text: the
`prepare-commit-msg` says it is comfort and bypassable, the `pre-receive` says it is enforcement
and the only kind there is on a host with no push rules, and that where a server has no binary a
copied pattern has to be kept in step.

<!-- xeno:section:alternatives -->
## Alternatives

**Add the two sections and ship three checked rules.** An `interfaces` section and a
`migration-notes` section in the design and implementation templates would make the migration
note rule the `checked` rule section 9 writes it as, and would give the dependency rationale
something to read. Rejected as the wrong package: the template set is WP3's, the sections appear
in every phase of every project, and both bundles would need headings. It is the strongest
alternative here and the finding says so, because what this piece ships instead is three rules a
person answers where one of them was specified as mechanical.

**Ship nothing checked at all.** All four as review rules, no new predicate, no registry change.
Rejected because the plan names a filled release notes section and because a first set that
demonstrated none of the last three pieces' machinery would leave `kind: checked` untested by
anything but fixtures.

**Make `release-notes-are-filled` a G-Schema matter instead.** A required section that is empty
could be a structural finding, which would cover every template's required sections at once
rather than one section by one rule. Rejected because it changes what every artifact in the trail
is judged by — several phases here carry a required section with a thin paragraph — and because
the gate list is a budget: G-Schema checks fields, and a rule is how a project says it wants more.
Worth proposing as a spec change and not worth doing quietly.

**Put documentation consistency in the set as a review rule.** It would reach every project, cost
nothing mechanically, and the plan calls it a candidate for the set. Rejected because a review
rule in `given/builtin/` is an entry every project answers at every P5 for ever, and a project
whose documentation lives in a wiki would answer `not-applicable` with a note every time. The
examples mechanism is for exactly that, and section 9 says a rule not in the tree does not apply.

**Write the rules as this project's own `given/project/` set instead.** Easier to make them sharp,
and they would be exercised harder. Rejected: the task is the shipped set, and a `given/project/`
set here would be the fifth piece dressed as the fourth. The project layer is where this
repository's conventions go and nothing in the plan asks for it yet.

**Carry the Conventional Commits pattern in the `pre-receive` hook.** Avoids needing the binary on
the server, which is the case the plan says a copy is unavoidable for. Rejected as the default:
three copies of a pattern produce a push that rejects what the gate accepts, which costs more
trust than the check is worth. The hook calls the binary, and its own text says what to do where
there is none.

<!-- xeno:section:impact -->
## Impact

**A rule applies, for the first time, in this repository.** Every phase after this set lands
resolves four rules, evaluates one of them, and requires three answers at P5. The gates stop
being green over nothing.

**`rules_hash` starts carrying something.** Every artifact written after this piece records a
hash over four rules rather than over an empty rendering, so A66's field finally says which set
was in force. The trail therefore has three eras, not two: the placeholder, the empty set, and
the shipped set — divided by #158 and by this commit.

**This intent's own P5 has to answer three review rules.** G-Policy will require an entry per
review rule, so the last phase of this intent is the first artifact in the project to carry a
`review_checklist`. If a statement cannot be answered honestly here, it will not be answerable
anywhere.

**Every project that runs `xeno init` receives the set.** Four rules it did not ask for and cannot
delete, only override, which is what `given/builtin/` means. That is the weight behind the
constraint that no project could reasonably refuse them, and it is why three of the four are
statements about artifacts rather than about engineering practice.

**A phase can now go red for something a person has to answer rather than fix.** A missing
checklist entry is red until somebody writes `met`, `deviation` or `not-applicable` with a note.
That is the intended mechanism and it is also the first time this process can block on a judgement
rather than on a structure.

**The examples make the layout learnable by failing.** A reader who copies an example into the
wrong level gets a `scope` finding naming the claim and the path. Nothing documents the two axes
better than that, and nothing else in this repository teaches them at all until WP16.

**WP4 has one piece left.** External gates, which depend on none of this. After them the package
is done, and the plan's `G-Rules` and `G-Policy` rows are both implemented, which was the
substance of M1.
