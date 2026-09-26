---
intent: github.com/triplem/xeno#68
phase: 00-intake
created: "2026-09-26T21:27:00Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+a420490
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 0ff55ee8836013584bc062cff8285b6592639ca0676368f72aa8db10561693f9
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

`xeno enforcement check` compared three of the five settings `project.yaml` can state.
`approvals.not_by_author` and `merge_method` were read from the file and compared
against nothing, so a project stating either got a green report that did not cover it,
which is the quiet half of the problem the report exists to solve. That is A41.

<!-- xeno:section:scope -->
## Scope

`approvals.not_by_author`, compared. `merge_method`, reported as unchecked rather than
compared, because the specification says so.

Section 13's table gives `merge_method` the default behaviour "not checked", and the
plan says it is meaningful once a commit predicate is active, which no gate in this
runner has. So the issue that asked for both was half wrong, and the half it got wrong
is settled by the documents rather than by a preference.

<!-- xeno:section:context-rationale -->
## Why this context

**What satisfies the author clause is not a setting.** GitHub refuses an approval from a
pull request's own author, so wherever an approval is required, the approval that exists
is somebody else's. The host's own rule meets the requirement, and
`require_last_push_approval` is the stronger form, which the report names in a note
where it is off rather than demanding.

**An absence does not satisfy it.** With no approval required the state is unmet,
because a requirement about who gives an approval cannot be met by there being none.
This repository is that case and does not hit it, because it declares no author clause
alongside its zero approvals, which is the consistent declaration rather than a lucky
one.

**Reporting a field as unchecked is the honest third answer.** Leaving `merge_method`
silent would let a project believe it is covered; comparing it would make a check the
documents say is not made, which is a specification change first and a check for a
consumer that does not exist. `unknown` is neither met nor unmet, so a project declaring
it learns the truth without its run failing over it.

**The issue was written before the table was read.** #68 asked for both fields and cited
the quiet half correctly; it did not cite section 13, which answers one of the two. The
finding is recorded in that issue and in A41 rather than quietly narrowed.