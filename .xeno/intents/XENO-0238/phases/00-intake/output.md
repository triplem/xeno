---
intent: github.com/triplem/xeno#210
phase: 00-intake
created: "2026-10-03T19:39:43Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 75a0eb9aaf32ba95ccb8827e2e8fe6aea550aa78c3be4dce6d3ddc8e846a67b6
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

`CLAUDE.md` says "Lines wrap at 88 characters, in code and in prose." One sentence,
covering three things that have three different right answers, and enforcing none of
them.

Go source has no width rule that any tool applies. `gofmt` does not wrap, so a line is as
long as its identifiers make it, and `internal/plugin/plugin.go` sat at 95 columns from
the intent that wrote it until the one that noticed — by accident, while reading for
something else.

Markdown tables cannot obey it. The table in `CLAUSE-READERS.md` is wider than 88 and
wrapping a row breaks the table, so the rule is already broken deliberately in the file
that documents broken rules.

Commit messages and pull request descriptions are 72, not 88. `CONTRIBUTING.md` says so
and gives the reason: GitHub reflows them to that width, so anything wider arrives as
something nobody wrote. Two commits in this session were written at 88 and rewrapped
after the fact.

A convention in `CLAUDE.md` is read before every change and believed, which is what makes
a wrong one expensive rather than untidy.

<!-- xeno:section:scope -->
## Scope

In scope is the Conventions paragraph of `CLAUDE.md`: state the three rules separately,
88 for Markdown prose with tables and code blocks exempt, nothing beyond `gofmt` for Go,
and 72 for commit messages and descriptions with the reason left where
`CONTRIBUTING.md` already carries it. One row in `ASSUMPTIONS.md` for dropping the Go
width rule rather than enforcing it.

Out of scope is a checker. The clause audit's own finding is that a reader which cannot
fail is worse than none, and a width check over Markdown prose would have to understand
tables, code blocks, long links and headings. That is its own decision, and this intent
neither makes it nor forecloses it.

Out of scope is rewrapping anything that exists. The rule changes; the tree is not
reflowed to it, because a reflow of every file would bury the one paragraph that is the
change.

No normative document is touched. The 88 appears in neither the process definition nor
the plan, so no specification commit precedes this.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the Conventions paragraph, `CONTRIBUTING.md`'s passage on description width,
and the evidence that the rule is unenforced: `gofmt -l` is silent on a 95-column file,
and no rule file or workflow step mentions a width.

`CONTRIBUTING.md` is read rather than quoted, because the new wording claims it explains
the 72 and that claim is checkable: it does, under "Write the description at 72
characters", with the reflow as the reason. A convention that points at another document
for its reason is only as good as the pointer.
