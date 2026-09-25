---
intent: github.com/triplem/xeno#34
phase: 00-intake
created: "2026-09-25T20:39:07Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+81b8426
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 8810d992b4bdfffd654cfc9783068689e6b20111d184cbcb84ab6d6800822e92
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

Issues here carried headings like "## 1, and why the schedule matters here", with the
next one starting at 7. The number was an index into a list the reader never saw, left
behind when the text around it was rearranged. It reads as a reference, so a reader
looks for what it refers to and finds nothing.

The numbers that do exist in this project belong to the process definition's sections
and to the plan's steps. A heading beginning with a bare number borrows that form
without meaning it.

<!-- xeno:section:scope -->
## Scope

The convention, written into `CLAUDE.md` where the other prose rules are. The three
headings in the open issues were repaired when #34 was filed; a scan of every open
issue confirms none is left. Three remain in comments on closed issues and stay there:
a comment is a record of what was said at the time.

While in that file, one stale fact goes with it: the conventions name
`gopkg.in/yaml.v3` as the vendored dependency, and it has been `go.yaml.in/yaml/v3`
since the fork was taken over by the YAML organisation.

<!-- xeno:section:context-rationale -->
## Why this context

A rule in `CLAUDE.md` is the only one of these that acts before the text is written,
because that file is sent with every request of every session. Repairing the headings
fixed the instances; nothing yet stopped the next one.

It is the third rule in that file about prose that was assembled rather than written,
after the splice rule from #42. The same shape: something true of an earlier draft
survives into the current one, and it is invisible in a diff because it is a fragment
of what was already there.
