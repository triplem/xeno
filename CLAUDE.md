# Xeno

A network free runner for a six phase process: hashing, gates, the phase sequence,
evidence attachment. The process is defined in `.xeno/docs/process-definition.md`, the
build order in `.xeno/docs/implementation-plan.md`. Both are normative.

## Build and test

    go build -o xeno ./cmd/xeno
    go test ./...
    gofmt -l . | grep -v '^vendor/'      # must print nothing
    go vet ./...
    ./xeno gate verify                   # must exit 0

## Conventions

Lines wrap at 88 characters, in code and in prose. Go source carries
`SPDX-License-Identifier: Apache-2.0` and no per file copyright line. One dependency,
`go.yaml.in/yaml/v3`, vendored; adding a second is a decision, not a step.

Where a paragraph is being changed for the second time, replace it rather than edit
into it, and read it back as a paragraph rather than as a diff. Editing into a sentence
leaves the words around it behind, and a small diff is exactly when that is not noticed.

Headings name their section in words, in issues as much as in files. A leading number
indexes a list the reader cannot see, and the numbers this project does have belong to
the process definition's sections and the plan's steps, so a borrowed one reads as a
reference to them.

A comment or a passage of prose says what the construction is and why it is that way.
Where the reason is genuinely the thing it replaced, that belongs in the commit message:
the reader has the current tree and nothing else, and after a squash the history does
not carry the intermediate states either. This holds for the documents as much as for
the code; prose in a file is as readable by a stranger and ages the same way.

Commit subjects are Conventional Commits without an issue reference. The reference goes
in the footer, `Refs #123`, and `Closes #123` on the commit that finishes the work and
in the pull request description. `CONTRIBUTING.md` says why, and names the host
setting it depends on.

An intent key is `XENO-` and the issue number padded to four digits, `XENO-0049`, and
the padding is for sorting alone. Intents sealed before this keep their names: the key
sits inside `artifacts_hash` and inside the merge commits that name them, so renaming
one changes every verdict in it.

## Three standing rules

**The documents are not editable by the agent.** Where the code and the specification
disagree, the specification wins until a person changes it, and a change to it is its
own commit made before the code that follows from it.

**No invented fields, gates, tools or rules.** Everything the artifacts carry is
enumerated in the process definition. An addition is a spec change first, by the rule
above. The MCP tool surface and the gate list are budgets, not lists.

**Every change belongs to a work package and to an intent.** A branch carries one
intent, the commit references its issue, and the issue carries the label of its package.
Where something needed belongs to no package, that is a finding about the plan and gets
written down rather than absorbed.

## Where things are written down

Assumptions and decisions taken while building the core: `ASSUMPTIONS.md`. Learnings do
not go here or in this file. Section 10 routes them from a phase's `learning.yaml`
through a merge request against the rule set, so that they take effect after review and
not on being noticed.

This file is sent with every request of every session. Keep it short.
