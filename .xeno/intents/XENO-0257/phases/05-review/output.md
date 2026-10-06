---
intent: github.com/triplem/xeno#257
phase: 05-review
created: "2026-10-06T07:59:18Z"
schema_version: "1.0"
runner_version: dev+77a35b1.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f1fca9129f2490a220dbdd0b2473bec9ca789a9c2519f3b352e524dd93c2664f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Two, each naming what it departs from, and one note about method. Against P2 design, which said what to collect and not what to avoid: the README gains a paragraph instructing that no timing be measured, because four comments on #117 read agent latency as process cost before anybody noticed, and a fixture whose purpose is a comparable measurement should not let that recur. A correction inside P3: ten README prose lines came out at 89 and 90 columns against the 88 the conventions ask, and were reflowed by replacing whole paragraphs rather than breaking at the overflow. And the method note: the twelve template and bundle files were generated from the shipped set by a script and then checked by a second script that parses both trees and compares ids, versions, phases and headings — the headers and the reasons are written, the structure is derived, because transcribing seventeen section entries across six files is where a wrong id goes unnoticed.'
      result: deviation
      rule: deviations-are-traceable
    - note: 'No interface changes. Thirteen files under examples/ that nothing loads: template.Load reads .xeno/config/templates and .xeno/plugin/templates and never this directory, so the fixture is inert by construction and the verdict count is unchanged by it. No code, no command, no field, no artifact shape, and no shipped template altered — every id, version and phase in the candidate matches the shipped set, which a script asserts. The nearest thing to a dependant is a project that chooses to copy a directory into its own config, and the README and each header say how, that the strings bundle must come with it, and what the one section a gate reads is.'
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: None added and no code changed; go.mod is untouched. Twelve of the thirteen files are YAML read by nothing and the thirteenth is Markdown. The one thing that would have needed something new is a named template variant so a project could select a short path rather than overriding per id, and that would be a new field and a specification change — which the fixture exists to avoid needing, since Load already resolves per id and the required flag is the whole lever.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No normative document is touched: section 9 fixes what the sections are
and the `required` flag is the template layer's, which the plan leaves to the project. Nothing is
invented — no field, gate or rule, and no section is added or removed; only flags move, and every id
stays defined. The branch carries one intent, the commit references #257, and the issue carries
`wp20`.

The acceptance criteria. Eleven met, one met by the commit this phase precedes.

The non-goals held. Nothing adopted, here or anywhere — `.xeno/config/templates` does not exist. No
shipped template touched, so no adopter's process moves and the plugin digest is unaffected. No
recommendation that the shortcut be taken. No second language bundle, which the README says. No
section deleted. No figures from running it, because running it is the next step.

What a reviewer should check is one count and one claim. The count: six required out of seventeen,
which a script verifies against the shipped set along with the ids, versions, phases and headings.
The claim: that `release-notes` is the only section any rule names, which is what licenses eleven
drops and which the one apparent counter-example nearly broke — `Scope` in
`internal/rules/rules.go` is a rule's own scope field and not the `scope` section, and taking it at
face value would have made the candidate wrong about what it may touch.

What is left is in P4's gaps and is mostly the point: nobody has run it. The candidate is a question
expressed as data, and the measurement that would answer it needs an intent and a deliberate
comparison with the nine on #117.

One paragraph of the README exists because of this session's own error rather than a finding, and it
is marked as such in P3's deviations: it says not to measure a timing, because four comments on #117
reported agent latency as process cost before anybody noticed.

The evidence was declared from a file whose writer had exited and whose hash was read twice and
compared — the method XENO-0255's mistake taught, where a hash bound to a log still being written
cost that intent two phases.

<!-- xeno:section:release-notes -->
## Release notes

`examples/templates/` holds a candidate reduced section set: six template overrides keeping **6 of
the 17 required sections**. It is enabled nowhere and is not a recommendation. It exists to be tried
and measured, which is what the plan asks of a shortcut's shape.

**Why now.** #117 has nine six-phase intents measured. The record is a fixed cost — 17 sections,
32 to 33 files, 1,350 to 1,499 lines — against changes spanning 13 to 440 lines. The record varies by
11% and what it describes by a factor of thirty-four, so proportionality is not a ratio a shortcut
could improve by a percentage; the only lever is which of the seventeen a small change may omit.

**The measurement that makes that answerable.** Of the seventeen, **one is read by a gate and
sixteen are read by the next-step suggestion.** No Go file outside the tests names a section by id;
across the shipped rules and the examples only `release-notes` appears, through `section-non-empty`;
and `sectionNonEmpty`'s own comment states the general case and cites A73. G-Policy is not a
counter-example — its checklist requirement reads the `review_checklist` frontmatter, not the
`review-checklist` section.

**What the candidate keeps, and on what grounds** — the grounds differ in kind and the files say
which:

| section | why |
|---|---|
| `release-notes` | **a gate reads it**: `release-notes-are-filled` through `section-non-empty` |
| `acceptance-criteria`, `test-mapping` | **by decision, not a reader**: #258 settled that section 9 will make a criterion identifiable and a gate will read the mapping |
| `problem`, `changes`, `results` | why the change exists, what was done, whether it worked |

The eleven drops each carry a reason in the file that drops them, so disagreement lands on a
sentence rather than on the set.

**Nothing is deleted.** Every section the shipped template defines is still defined, with the same
`id`, `version` and `phase` — asserted by a script that parses both trees. Only the `required` flag
moves, which is the whole lever: `Missing` reports required sections that carry nothing and nothing
else reads the flag. So `xeno section set` still accepts all seventeen, a rule naming one cannot find
it unknown, and an author who wants a dropped section simply writes it.

**Trying it** is `cp -R examples/templates/. .xeno/config/templates/`, or one directory alone, since
a project template beats the shipped one per id. The strings bundle must come with each directory: a
missing bundle is an error and never a fall back to another language.

Nothing is adopted here. `.xeno/config/templates` does not exist after this change, deliberately —
switching this repository mid-stream would make the tenth intent incomparable with the nine the
fixture exists to be compared against.

**Nobody has run it**, which the README says: whether six sections produce an artifact a person can
review in six months is the question, and only running it answers that.

<!-- xeno:section:residual-risk -->
## Residual risk

Nobody has run it, and that is the risk as much as the point. A candidate that parses and is wrong
about what a reader needs is harder to dislodge than no candidate, because it looks finished. The
README says plainly that trying it is how the question gets answered; nothing stops somebody
adopting it on the strength of the table instead.

It can age with nothing to notice. If a shipped template gains a section the candidate will not have
it, and if `release-notes-are-filled` stops being the only rule naming a section the header asserting
that becomes false. Nothing reads these files — which is what makes them safe and also what makes
them unverifiable — so the defence is the date in the README. That is `docs/clause-readers.md`'s
defence, and #254 is the record of it failing there.

The eleven reasons are assertions and one of them is the weakest link. "A change with no plan beyond
its criteria cannot depart from one" is the argument for dropping `deviations`, and this session
produced three intents whose deviations section carried the most useful thing in the record —
including two where a criterion was falsified by its own test. If the reduced set is ever adopted for
a change that turns out to be larger than it looked, `deviations` being optional is where the loss
would be, and the section stays defined precisely so it can be written anyway.

The measurement this enables is not protected from the mistake that prompted it. The README says not
to measure a timing and nothing enforces it; the next person can still read `context.lock.yaml` and
report agent latency as process cost, as four comments on #117 did before anybody noticed. One
paragraph is the whole guard.

Adopting it would break comparability if done casually. The nine recorded intents were measured
against the seventeen, so a project that adopts the candidate and keeps comparing to #117's figures
is comparing two different processes. The README says the comparison is valid because only the
required flags change, which is true for one intent run deliberately and false for a drift into using
it by default.

What is not a risk: the verdicts, confirmed intact at exit 0; this repository's own process, since
`.xeno/config/templates` does not exist; any adopter of the plugin, since no shipped template moved;
and anything executable, since `template.Load` cannot reach `examples/`.
