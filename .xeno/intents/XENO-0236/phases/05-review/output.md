---
intent: github.com/triplem/xeno#183
phase: 05-review
created: "2026-10-03T19:09:31Z"
schema_version: "1.0"
runner_version: dev+0462eb7.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 099a08eced6b81e47f672b2ba0cd1697792042e68d24e53a3bd29201232dfb21
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
      Two: the comment rewrap, which touched more of internal/plugin/plugin.go than the design
      anticipated because the file had been at 95 columns since an earlier intent of mine; and the
      shape of this intent, where the specification change is the work and the code change is
      comments, which is what the first standing rule produces when a document is what is wrong.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      A documented mechanism is gone and nothing migrates, because nothing implemented it: no code
      to delete, no behaviour to change, no artifact that ever recorded a plugin root. The only
      part somebody could have depended on is XENO_PLUGIN_ROOT leaving the variable list, and
      nothing read it.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Two, each naming what it departs from: the comment rewrap to 88
columns, which touched more of `internal/plugin/plugin.go` than the design anticipated because the
file had been at 95 throughout since an earlier intent of mine; and the shape of this intent, where
the specification change is the work and the code change is comments, which is what the first
standing rule produces when the thing corrected is a document.

**`interface-change-needs-a-migration-note` — met, and it is a removal rather than a change.** A
documented mechanism is gone. Nothing migrates, because nothing implemented it: no code to delete, no
behaviour to change, no artifact that ever recorded a plugin root. What a reader loses is the ability
to find a mechanism they could not have used, and what they gain is four paragraphs saying why.
`XENO_PLUGIN_ROOT` leaving the variable list is the only part somebody could have depended on, and
nothing read it.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added.

**Beyond the three rules.**

*Was the removal argued or asserted?* Argued, from two measurements, and both are in the document
rather than only in this trail. That was deliberate: a clause removed with the conclusion alone is a
clause somebody reinstates by disagreeing with an absence, and one removed with the evidence beside
it has to be argued with.

*Did the first standing rule get followed in the unusual direction?* Yes, and it is worth naming. The
rule orders the specification change before the code, and here the code was already correct — it read
the vendored directory all along. So the document moved to catch up with the code, and what followed
was comments that had argued for not implementing a clause the document no longer has. The ordering
rule holds for a removal and produces almost no second commit.

*What did this intent retract?* One of my own recommendations, made twice: the split by purpose, with
the gate path pinned to the vendored tree and everything else honouring the order. It is in the
alternatives with what killed it, because the reasoning that failed is the reasoning somebody would
repeat.

*Is anything left pointing the wrong way?* One thing, and it is in P4's gaps: P3 justified the rewrap
by a convention the maintainer retired two exchanges later. The phase is sealed and keeps the reason
it was written with, which is what judging a phase against what was in force means — and a reader of
that deviation will find a convention the repository no longer has. That is the process working, not
a defect, and it is the clearest example this trail has of why a sealed phase is not corrected.

*Could this change a verdict behind it?* No. `gate verify` is 285 at exit 0, the suite is unchanged,
and no gate, rule or artifact moves.

<!-- xeno:section:release-notes -->
## Release notes

**The plugin is the vendored one.** `.xeno/plugin/`, found from the git root, and nothing else.

Section 7 described a resolution order above it — a `--plugin-root` argument, then
`XENO_PLUGIN_ROOT`, then the vendored tree, then a client's own variable — and no longer does.
`XENO_PLUGIN_ROOT` is not among the variables the normalised environment carries. Three remain:
`XENO_PLUGIN_DATA`, `XENO_HARNESS`, `XENO_HARNESS_VERSION`.

**It was removed rather than implemented, and the section says why.** The gate path reads the rule
set and the templates from the plugin, so a root taken from the environment makes `rules_hash`,
`strings_hash` and a rendered artifact depend on it — and a phase is judged only against the rule set
its own artifact records, so a changed set does not disagree with the trail, it stops judging it.
Measured: a valid rule tree that differs leaves `xeno gate verify` at exit 0 over 273 verdicts while
G-Policy reports nothing about the 75 phases that recorded the previous hash.

**The client's variable at the bottom could not work at all.** A project with no vendored plugin
resolves no template and has no `plugin_version`, so no phase of it renders and G-Schema reports it.
Falling back to a plugin the client installed let such a project fail differently rather than work.

**Nothing changes about how the plugin is found**, because nothing implemented the order. No code was
deleted and no behaviour moved; the document caught up with the code.

**One sentence is kept for the day an override is wanted:** a runner that cannot verify a plugin does
not accept one from outside the repository.

<!-- xeno:section:residual-risk -->
## Residual risk

**A sealed phase of this intent cites a convention that no longer exists.** P3 justified the comment
rewrap by the 88-column rule in `CLAUDE.md`, and the maintainer decided two exchanges later that Go
source has no width rule beyond `gofmt`. The phase keeps its reason, which is correct and is the rule
this process rests on, and a reader will find a justification pointing at nothing. The clearest
example in this trail of why a sealed phase is not corrected, and worth knowing before somebody tries
to correct it.

**`XENO_PLUGIN_DATA` is in the state the removed entry was in.** Listed in section 7, exported by the
entry point, read by nothing, because the section pins it to one value. A reader cannot tell a
variable with no reader from one whose reader is a constant, and A84 records it. The list is one
entry shorter and not yet honest.

**The removal rests on two measurements of this repository.** A project whose gate path did not read
the plugin would have different numbers. The document carries the measurements so that the argument
is visible as particular rather than general, and a later reader could still take the conclusion
without them.

**The condition kept in section 7 has no reader and cannot have one.** It describes what a future
override would have to do. Deliberate, and it is the category #202 warns against counting as a gap —
this intent added one of those on purpose, which is worth being explicit about in an issue whose
subject is clauses without readers.

**Reinstating the clause is easy and reinstating the mechanism is not.** That asymmetry is the
protection and it is not a check: somebody could add the paragraph back without the resolution, which
would put the document back in the state this change took it out of.

**What is not a risk.** Any verdict behind this intent, any behaviour, and any project's plugin
resolution — which was the vendored directory before this change and is the vendored directory
after it.
