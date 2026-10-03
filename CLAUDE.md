# Xeno

A network free runner for a six phase process: hashing, gates, the phase sequence,
evidence attachment. The process is defined in `docs/process-definition.md`, the build
order in `docs/implementation-plan.md`. Both are normative.

## Build and test

    go build -o xeno ./cmd/xeno
    go test ./...
    gofmt -l . | grep -v '^vendor/'      # must print nothing
    go vet ./...
    ./xeno gate verify                   # must exit 0

## Conventions

Markdown prose wraps at 88 characters; tables and code blocks do not. Go source has no
width rule beyond `gofmt`, carries `SPDX-License-Identifier: Apache-2.0`, and no per
file copyright line. Commit messages and pull request descriptions wrap at 72, because
GitHub reflows them to that width, which `CONTRIBUTING.md` explains. One dependency,
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
in the pull request description. Several issues take a keyword each, `Closes #1, closes
#2`, because a host reads only the first reference after one. `CONTRIBUTING.md` says
why, and names the host settings it depends on.

An intent key is `XENO-` and the next number of a sequence of its own, padded to four
digits, beginning at `XENO-0200`. The issue it belongs to is the `intent` field of
`intent.yaml`, which carries the host and the repository with it; the key does not
repeat it. `xeno intent status` without an intent lists them in the order they were
created, which is what the key used to be asked to express and could not.

Keys assigned before this were the issue number, and they stay: the key sits inside
`artifacts_hash` and inside the merge commits that name them, so renaming one changes
every verdict in it. Those numbers reach `XENO-0121`, so the sequence starts beyond
anything they can collide with, and the gap is how a reader tells which scheme a key
follows.

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
