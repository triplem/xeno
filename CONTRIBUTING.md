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

**A description is a commit message.** It becomes the body of the squashed commit, so
anything in it that a machine acts on takes effect. `Closes #3` is used that way
deliberately. The markers that switch CI off are the same mechanism pointed the other
way: one quoted as an example in a description once meant the merge ran nothing at all.
The verify workflow refuses a description carrying any of them and holds the
authoritative list, so nothing here has to write one out in order to warn about it. To
mention one, describe it.

**Write the description at 72 characters.** It becomes the body of the squashed commit,
and GitHub rewraps it to that width on the way, so anything wider is reflowed into
something nobody wrote. It is a commit message, not a note about one.

### The setting this depends on

A footer survives a squash only where the squashed message is built from the pull
request description. On this repository that is *Settings, General, Pull Requests,
Squash merging*, set to **Pull request title and description**. The same setting exists
elsewhere under another name; on GitLab it is *Settings, Merge requests, Squash commit
message template*, which has to include `%{description}`.

**Without that setting the reference is lost at the merge.** It is not a nicety; it is
what makes the convention work at all.

Squash is also the only method the repository offers. A merge commit and a rebase both
discard the description, and the footer with it, so leaving them available would make
the convention depend on which button somebody presses.

On GitHub all three are applied by `scripts/github-settings.sh`, which prints them
before and after so that a drift is visible rather than assumed. A setting is not a
commit and no gate covers it, so the script is how it is written down at all. Another
host's equivalent may have no API call shaped like it and be set by hand.

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
