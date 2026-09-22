# Contributing

## Every change is an intent

A change to Xeno starts from an issue and is recorded under `.xeno/intents/<KEY>/`.
While the process can honestly carry only P0 (see `ASSUMPTIONS.md`, A6), an intent
runs its intake through Xeno and the rest is recorded by hand. The commit carries the
intent in a trailer:

    Xeno-Intent: XENO-2

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
