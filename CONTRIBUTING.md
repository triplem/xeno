# Contributing

## Every change is an intent

A change to Xeno starts from an issue and is recorded under `.xeno/intents/<KEY>/`.
While the process can honestly carry only P0 (see `ASSUMPTIONS.md`, A6), an intent
runs its intake through Xeno and the rest is recorded by hand. The commit carries the
intent in a trailer:

    Xeno-Intent: XENO-2

## Commit messages

The subject is plain Conventional Commits and carries no issue reference:

    feat(docs): a description

The issue goes in the footer, and so does the closing keyword on the commit that
finishes the work:

    Refs #3

    Closes #3

A pull request description carries the closing line too; the template fills it in.

### The setting this depends on

A footer survives a squash only where the squashed message is built from the pull
request description. On this repository that is *Settings, General, Pull Requests,
Squash merging*, set to **Pull request title and description**. On GitLab it is
*Settings, Merge requests, Squash commit message template*, which has to include
`%{description}`.

**Without that setting the reference is lost at the merge.** It is not a nicety; it is
what makes the convention work at all.

On GitHub it is applied by `scripts/github-settings.sh`, which prints the setting before
and after so that a drift is visible rather than assumed. A setting is not a commit and
no gate covers it, so the script is how it is written down at all. GitLab's template has
no API call shaped like it and is set in the project's merge request settings by hand.

### Why the reference is not in the subject

It used to be, and it closed issues by accident. GitHub reads a subject for closing
keywords, `fix` is one of them, and `fix: #3` closes the issue even with the colon
between. That closed issue #2 of this repository on the third of its six commits, while
the work ran on for three more, and nothing reopens an issue afterwards. A scope
separates them again, so `fix(gates): #3` does not close — which is worse, because then
it depends on whether you wrote a scope.

In the footer nothing closes unless you write `Closes`.

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
