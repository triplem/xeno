# Contributing

## Every change is an intent

A change to Xeno starts from an issue and is recorded under `.xeno/intents/<KEY>/`.
While the process can honestly carry only P0 (see `ASSUMPTIONS.md`, A6), an intent
runs its intake through Xeno and the rest is recorded by hand. The commit carries the
intent in a trailer:

    Xeno-Intent: XENO-2

## Commit messages

The subject follows Conventional Commits with the issue reference at its end:

    fix(gates): never carry a decision forward for an external finding (#3)

The reference belongs at the end and not after the type. GitHub reads a subject for
closing keywords, `fix` is one of them, and it closes an issue on `fix: #3` even with
the colon in between. That closed issue #2 of this repository on the third of its six
commits, while the work ran on for three more, and nothing reopens an issue afterwards.

A scope separates them again, so `fix(gates): #3` leaves the issue open. That is the
awkward part rather than the reassuring one: it would mean a commit closes its issue or
not depending on whether you wrote a scope. At the end of the subject the reference
survives a squash just as well and closes nothing, whatever else the line contains.

Because nothing closes by accident any more, closing is deliberate. The commit that
finishes the work carries the keyword in its body, and so does the pull request
description:

    Closes #3

Both, because a squashed message is built from the title and the description, so a
trailer in a single commit body does not survive the squash.

## Developer Certificate of Origin

Every commit is signed off under the Developer Certificate of Origin 1.1
(https://developercertificate.org). Signing off states that you have the right to
submit the change under this project's licence:

    git commit -s

which adds a line of the form

    Signed-off-by: Name <email>

using your git identity. A commit without it is not merged.

## Contributions from outside

The repository is private and changes come from its maintainers. How a contribution
from outside is handled, with no intent and no artifacts behind it, is decided before
the repository is made public; the question is recorded in the implementation plan
under the documentation package.

## Before sending a change

    go test ./...
    go build -o xeno ./cmd/xeno && ./xeno gate verify
