---
intent: github.com/triplem/xeno#239
phase: 05-review
created: "2026-10-05T09:02:14Z"
schema_version: "1.0"
runner_version: dev+3f5ad34.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7443f34739d0db836995f4d4f41d199b4af4003f6ba1789cb5b3fd206a5999f2
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
      Three, each naming what it departs from. The exemption of the plan's line 2009 names
      acceptance criterion 2, which declared it before the work began. The five rewrapped
      paragraphs name the sweep that caused them. The restated reference counts name #239's
      table and say why a count needs its date. P4 adds a fourth against criterion 6 itself.
  - rule: interface-change-needs-a-migration-note
    result: not-applicable
    note: >-
      No interface changes. Four renames, 27 corrected references, one register row and two
      lines of prose; no command, field, gate, rule or artifact shape. Nothing outside this
      intent depends on a document's path, because nothing opens one by path. The nearest
      thing to a dependant is a reader holding an old link, and the release notes tell them
      where the file went and that git log --follow still reaches it.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      None added, and none was weighed. The one tool this move would benefit from is a link
      checker, which is WP16's and is named in P4's gaps as the gap that keeps criterion 2 a
      grep somebody has to remember to run.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. The first is engaged and held: `docs/implementation-plan.md`
names `ASSUMPTIONS.md` at line 2009 and is not touched, so the specification keeps its one
stale reference until a person changes it, and A94 says so rather than leaving a reader to
find it. No field, gate, tool or rule is added; nothing executable changes at all. The
branch carries one intent, the commit references both issues, and both carry `wp16`.

The acceptance criteria. Seven met as written, one met in substance with its wording
corrected in P4: criterion 6's verdict count could not survive an intent that writes its
own verdicts, which is a defect in the criterion and not in the work.

The non-goals held. No link checker. The five root documents stay. No sealed artifact is
touched — nothing under `.xeno/intents/` is in the diff but this intent's own directory.
Neither normative document is edited. No content rewritten beyond the five paragraphs the
sweep pushed over 88 columns, and those were replaced as paragraphs and read back as prose.

#222's two open questions are both answered rather than inherited. The name is
`docs/m0-gate-job.md`, a third option neither the issue nor this intent's design began
with, and the file says at the top that it is a record. Both are decisions a reader can
disagree with, which is why they are in P2 with their reasons.

What the design got wrong is in P3's deviations and P4's results with the figures, not
here: the reference counts were restated from this tree because #239's were a day and five
intents old, and criterion 6 was unsatisfiable as written.

One pre-existing defect is left deliberately and named in P3: `docs/m0-gate-job.md` line 8
breaks a sentence mid-phrase. It is WP16's with the link checker.

<!-- xeno:section:release-notes -->
## Release notes

Four documents move out of the repository root and under `docs/`, where this project's
prose lives:

| was | is |
|---|---|
| `ASSUMPTIONS.md` | `docs/assumptions.md` |
| `SUPPLY-CHAIN.md` | `docs/supply-chain.md` |
| `CLAUSE-READERS.md` | `docs/clause-readers.md` |
| `M0.md` | `docs/m0-gate-job.md` |

The root keeps `README.md`, `CONTRIBUTING.md`, `SECURITY.md`, `LICENSE` and `NOTICE`: the
entry point, and the four the host surfaces from there.

Every reference to the four outside the sealed trail is corrected — 27 of them across
`README.md`, `CONTRIBUTING.md`, `CLAUDE.md`, five workflow files, one Go comment and seven
rows of the register. Nothing executable reads these documents by path, so no build, test or
workflow behaviour changes.

Two references are knowingly left. `docs/implementation-plan.md` names `ASSUMPTIONS.md` in
one sentence and is normative, so correcting it is a person's commit rather than this one.
And 274 mentions of the old paths sit inside sealed artifacts, where rewriting one would
stale every verdict over it. A94 records both.

`docs/m0-gate-job.md` is renamed from `M0.md` rather than merely moved, because every file
under `docs/` is a descriptive hyphenated name and the guide is about standing up the M0
gate job. It now says at the top that it records how that was done once, which is what a
walkthrough in a documentation directory needs to say.

For anyone holding a link to an old path: the file is under `docs/` with a lowercase name,
and `git log --follow` reaches its history across the rename.

<!-- xeno:section:residual-risk -->
## Residual risk

The plan's line 2009 stays stale until a person corrects it. Low cost and certain: it is
inline code in a sentence about the core's other choices, nothing resolves it, and the
worst case is a reader who goes looking for `ASSUMPTIONS.md` in the root and does not find
it. It needs one specification commit, and A94 is where whoever makes the next one will
find it described.

A link checker added later will report the 274 sealed references unless it is told to skip
`.xeno/intents/`. This is the risk this intent most plausibly creates for another: the
exclusion is obvious once stated and invisible until the checker's first run, when it will
produce 274 findings nobody is permitted to act on. A94 states it; WP16 has to read it.

Someone cloning onto a case-insensitive filesystem may see the `M0.md` to
`docs/m0-gate-job.md` rename collapse. The index is correct, which is what `git ls-files`
verified, and the name changes as well as the directory, so the collision a pure case
change would risk does not arise. Low.

The register's closure banner still says the record is closed and that A77 is the last row,
with A94 now the seventeenth row added after it. This intent did not repair that, because
repairing it is a second change and would ride in a move commit; it is recorded as P2's
learning instead, and it is a finding about the register rather than about this move. It is
the likeliest thing in this diff for a reviewer to object to, and the objection is fair.

Five paragraphs were rewrapped. Each was read back as prose rather than as a diff, which is
what this project's conventions ask, but rewrapping is where a word gets lost and no tool
would catch it. `README.md`'s is the one worth a reviewer's eye: it dropped a second
mention of a long path in favour of "that guide".

What is not a risk: the 351 pre-existing verdicts, which `gate verify` confirms at exit 0,
and the build, which cannot depend on a document's path because nothing opens one.
