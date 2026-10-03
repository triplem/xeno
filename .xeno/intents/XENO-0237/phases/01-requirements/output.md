---
intent: github.com/triplem/xeno#202
phase: 01-requirements
created: "2026-10-03T19:24:26Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bb7dbda3783f9c05f328efc51b287be157611f2a85f07c52879bca14ad06413c
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

Every normative clause in both documents has a named reader or is listed as having none.
A reader is named as the thing that would fail if the clause were violated — a gate, a
refusal in the runner, or a step in a workflow — and is looked up in the code, not
recalled.

The list is a file in the repository, beside `ASSUMPTIONS.md`, so that it is read by
whoever reads the conventions rather than living in an issue comment.

Clauses with no reader are separated from clauses whose reader is a person, and a reader
that cannot fail is called out as one. That is the `plugin_version` case: the field was
in G-Schema's required set and filled by a constant, so it was enforced and meaningless
at the same time.

The counts of all four kinds are stated, including the two kinds that are not
enumerated, so that a reader can tell the list is a pass over everything and not a pass
over what was interesting.

<!-- xeno:section:non-goals -->
## Non goals

Not a gate. Nothing in this intent reads the documents at run time; a clause-coverage
check would be a new gate and section 7 calls the gate list a budget.

Not a fix. The findings stay as findings, and the ones worth closing become issues after
this merges.

Not a judgement about whether a clause deserves a reader. "Xeno never runs a test" needs
none until somebody adds a convenience, and deciding that for eleven properties is a
second pass by a person.

Not a pass over the conventions files. `CLAUDE.md` and `CONTRIBUTING.md` are out by the
intake's scope.

<!-- xeno:section:constraints -->
## Constraints

The documents are not editable here, by the first standing rule, so where a clause is
ambiguous the list records the ambiguity rather than resolving it.

The list is prose and a table, which means it ages: a reader named by symbol is wrong
the day the symbol is renamed, and nothing will say so. It carries the date of its pass
and names the commit the readers were looked up in, which is the most a document can do
about that.

One dependency, so no tooling is added to produce or check the list. The extraction was
a search and a reading; both are repeatable from the file's own description of how it
was made.

Prose wraps at 88 characters, tables do not.
