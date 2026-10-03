---
intent: github.com/triplem/xeno#210
phase: 05-review
created: "2026-10-03T19:44:06Z"
schema_version: "1.0"
runner_version: dev+9fcc639.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 16f0a09936c4e5b1abceae9571ca4068f4a6a3a36822b697184499ece431f2a2
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
      One, in the implementation phase's deviations section, and it is a correction rather than a
      change of plan: the design asserted that no Go file was over 88 columns, and the tree has
      1 772 such lines in 58 of 64 files. The design phase is sealed, so the measurement and the
      correction are in the implementation phase with the figures.
  - rule: interface-change-needs-a-migration-note
    result: not-applicable
    note: >-
      No interface changes. One paragraph of CLAUDE.md and one row of ASSUMPTIONS.md; no code, no
      command, no field, no gate, no rule and no artifact shape. The only thing that changes is
      what a future session believes about line width.
  - rule: new-dependency-needs-a-rationale
    result: met
    note: >-
      None added, and the alternative that would have needed one is named and deferred: a width
      linter for Go. gofmt does not wrap, so enforcing a Go width means a second tool, and A91
      records that as a conversation rather than a step.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No normative document is touched: 88 appears in neither the
process definition nor the plan, so the first rule is not engaged and no specification
commit precedes this. No field, gate, tool or rule is added. The branch carries one intent,
the commit references #210, and the issue carries `wp0`.

The acceptance criteria. Three rules stated separately; the table and code-block exemption
written down; the 72 carrying a pointer to `CONTRIBUTING.md` rather than a second copy of
its reason; A91 recording why the Go rule is dropped; the paragraph replaced rather than
edited into and read back as prose; nothing in the tree rewrapped.

The non-goals held. No checker, no reflow, no change to `CONTRIBUTING.md`.

The file's own last line. "Keep it short": the paragraph goes from two sentences to four
and the section gains no new heading. The reason for the 72 is one clause, not a paragraph,
because the explanation lives elsewhere.

What the design got wrong is in the implementation's deviations with the numbers, not
quietly dropped: it asserted the Go tree was clean and the tree has 1 772 lines over 88.

<!-- xeno:section:release-notes -->
## Release notes

`CLAUDE.md`'s width convention becomes three conventions, because the one it had was right
about prose and wrong about code and commit messages. Markdown prose wraps at 88 with
tables and code blocks exempt; Go source wraps wherever `gofmt` leaves it; commit messages
and pull request descriptions wrap at 72, with the reason staying in `CONTRIBUTING.md`.

Nothing in the tree is rewrapped and nothing becomes newly broken that was not already
broken — the Go half of the old sentence had 1 772 violations in 58 of 64 files and no
reader, which is what A91 records.

No checker. All three remain conventions held by whoever is reading, and they join the list
in `CLAUSE-READERS.md` of rules whose reader is a person.

<!-- xeno:section:residual-risk -->
## Residual risk

Three stated rules, none enforced. A later session can drift from any of them and nothing
will say so, which is the same class of risk the clause audit catalogued and A91 accepts
knowingly rather than by omission.

The asymmetry suggests where the risk actually sits. Prose compliance was near-perfect
without a checker because a person wraps prose as they write it. Compliance for anything a
tool is supposed to handle was zero. So the unenforced Markdown rule will probably hold,
and the Go statement is now a description of what happens anyway, which is the safest kind
of convention to leave unchecked.

The 90 live prose lines over 88 stay, 62 of them in the two documents the agent cannot
edit. A reader who measures the tree against the new wording will find them, which is why
the verification phase names the figure rather than leaving it to be discovered.

If a width check is written later, A91 is the record of what was deferred and why, and the
decision it needs first is which of the 1 772 Go lines would be rewrapped.
