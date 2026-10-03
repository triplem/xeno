---
intent: github.com/triplem/xeno#210
phase: 01-requirements
created: "2026-10-03T19:40:09Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a8ba60aa6696b726be7e629abe0809b5a83b6b6cd5b4a59e4b96ae1f8c22fedd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

`CLAUDE.md`'s Conventions paragraph states three rules where it stated one: Markdown prose
at 88 with tables and code blocks named as exempt, Go source with no width rule beyond
`gofmt`, and 72 for commit messages and pull request descriptions.

The reason for the 72 is named but not restated: `CONTRIBUTING.md` carries it, and the
paragraph points there. The project's own convention is that a passage changed twice is
replaced rather than edited into, and the same logic applies across files — the reason
lives in one place.

`ASSUMPTIONS.md` records why the Go width rule is dropped rather than enforced, so that a
later reader does not read the absence as an oversight and restore it.

The paragraph is replaced as a paragraph, not edited into, and read back as prose. That is
the project's rule for a second change to a passage, and this is the second.

Nothing in the tree is rewrapped to the new rule, and the new rule makes nothing that
exists wrong: every Markdown file here is already at 88 with tables exempt, which is what
the rule now says.

<!-- xeno:section:non-goals -->
## Non goals

No width checker, for Markdown or for Go. The decision is deferred, not taken, and the
reason is the clause audit's: a check that cannot fail is worse than none, and one over
Markdown prose would have to understand tables, code blocks, long links and headings
before it could be trusted.

No reflow of the existing tree. Nothing becomes wrong under the new wording, so a reflow
would be churn that hides the one paragraph that changed.

No change to `CONTRIBUTING.md`. Its passage on 72 is correct, is already the reason the
new wording points at, and touching it would make this change two changes.

No normative document. The 88 is in neither the process definition nor the plan.

<!-- xeno:section:constraints -->
## Constraints

`CLAUDE.md` is sent with every request of every session and says so about itself, with
"Keep it short" as the last line. The replacement is four sentences where there were two,
and that is the whole budget: a convention that explains itself at length stops being read.

The three standing rules are untouched by this, because the file's own rules section is not
the Conventions paragraph. The first rule covers the specification documents, and this is
neither.

The wording was approved by the maintainer before the issue existed, so the work is to
carry it through the process rather than to decide it. A second person still releases
anything the gates find.

Prose at 88, tables exempt — which is the rule being written, applied to the artifacts
that write it.
