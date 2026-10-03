---
intent: github.com/triplem/xeno#172
phase: 03-implementation
created: "2026-10-03T10:37:37Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3ace547907c89324f3f52248d6c7f55f55f9f96c2f6b83a9ae9c910e5b1e1ffd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

One function, three tests, two files. 55 lines added, nothing removed.

**`internal/gates`, `links(c)`.** Reads the profile from P0, walks its `links`, and reports one
finding per declared document that is not in the tree. A link with no `docs` declares nothing to
find and is skipped. The cause names the component and the path; where the component is empty the
finding says "a link", because a finding has to name something. The next step says to correct the
path or take the link out, and adds why: a declared link is a claim about a file.

The comment carries the argument rather than the mechanics — that this is the profile's one
unambiguous error, that the finding names the profile because the profile is the claim and the base
is the consequence, that the runner is right to skip the link in silence because it records what
the phase was given, and that the check runs at every phase for the same reason the budget's does.

**`internal/gates`, `schema`.** One line: `links(c)` appended after `budget(c)`, so both profile
checks arrive through the gate section 5 names.

**Tests, `internal/gates/budget_test.go`.** The fixture from #171's budget work answers this with
one more field in the profile, so the three tests are short. A link whose document is missing
produces one finding naming the component, the path and the profile, with a next step. A link whose
document exists produces nothing, and neither does a profile with no links or a repository with no
profile. And the finding arrives through `schema` rather than through a function nobody calls.

**What is not here.** The byte count. Section 5 writes the lock's `files` as a path and a hash, so
recording a size is a field the specification does not have, and the second standing rule makes
that a change to the document before any code. The review carries the sentence a person would add.

<!-- xeno:section:deviations -->
## Deviations from the design

**Nothing departed from the design.** The check went where 02-design put it, in the shape the
budget check established, and the three tests are the three cases 01-requirements named. This is
the first intent of the twenty-eight whose implementation phase has no deviation to report, and the
reason is worth naming: the piece is one condition, its specification was another intent's
criterion rather than a document to interpret, and the shape it had to match already existed in the
tree.

**One thing worth recording that is not a departure.** The empty-component case — a link with
`docs` and no `component` — is not in the criterion and is not a profile a reader would write on
purpose. It produces a finding that says "the link for a link names …", which reads badly. The
alternative was to skip such a link, which would mean a declared document going unchecked because
the other half of the declaration was missing. The ugly sentence is the better of the two and it is
in the gaps.
