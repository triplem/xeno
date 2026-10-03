# Contributing

## Every change is an intent

A change to Xeno starts from an issue and is recorded under `.xeno/intents/<KEY>/`.
While the process can honestly carry only P0 (see `ASSUMPTIONS.md`, A6), an intent
runs its intake through Xeno and the rest is recorded by hand. The commit carries the
intent in a trailer:

    Xeno-Intent: XENO-2

**The trail only grows.** What is sealed is never rewritten, which section 7 states as a
rule holding throughout, so a change adds to `.xeno/intents/` and never removes from it.
An edit to a sealed artifact is caught by `gate verify`, which recomputes the hash and
reports the divergence. A deletion was caught by nothing until #193, and the verify
workflow now refuses a pull request that removes or renames any path under that
directory: a verdict that is gone is not distinguishable from one that never existed,
and the remainder verifies clean.

Adding a directory and dropping it again inside one branch is not that, and is not
refused. The comparison is against the base of the pull request, so what never reached
main is nobody's business.

## Commit messages

The subject is plain Conventional Commits and carries no issue reference:

    feat(docs): a description

**The pull request title is that subject.** The squashed message is built from the title
and the description, so the subject written on a branch is discarded at the merge and
the title is the only place the rule above has any effect. The verify workflow refuses a
title that is not a Conventional Commit, with the shipped `conventional-commits` pattern
and the same `xeno check commit-message` anybody can run on one by hand.

The host appends the pull request's number to it, so `feat(docs): a description` reaches
main as `feat(docs): a description (#3)`. That form is what the second shipped pattern,
`conventional-commits-with-issue`, exists for, and it is why the title is written
without a reference rather than with one.

**A prose title releases nothing, silently.** semantic-release reads the subjects on
main to decide whether there is a version to cut; a subject it cannot parse is not a
small release, it is no release. Fifteen merges of this repository carried prose titles
and no release ran for five days, which is what #190 was.

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

In the footer nothing closes unless you write `Closes`, and each issue needs its own
keyword. `Closes #50, #11, #12` closed #50 and left the other two open, because GitHub
reads only the first reference after a keyword; `Closes #50, closes #11, closes #12` is
the form that works. The first version looks like a list and is read as one item, which
is why it went unnoticed until somebody asked why two issues were still open.

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

    gofmt -l . | grep -v '^vendor/'      # must print nothing
    go vet ./...

## A red check is fixed before the next merge

Every check the pipeline runs gates the merge. Not one of them: all of them. `verify` is
the one that judges the trail, and `audit`, `gitleaks`, `semgrep` and `trivy` each
answer a question the trail cannot, so a merge past any of them red is a merge past
something nobody looked at.

Two of them run on a schedule as well as on a change, because an advisory arrives
without anything in the repository moving. A scheduled run going red is therefore the
normal way to learn that something needs doing, and it is the case this section exists
for: **the work it names comes before the next merge, not after the next feature.**

This was written because it failed. Six commits reached `main` in one morning while the
`audit` job was red, each from a pull request that was honestly green: the job did not
run on a pull request at all, so it could only report after the merge, and a check that
cannot run before a merge cannot gate one. Six green reports, one red gate, and nothing
wrong with anybody's diligence (#186).

**What a red check is not.** It is not a thing to raise the threshold past. The npm
audit baseline fails on drift rather than on presence, so a count that rises is a change
somebody has to read; raising the number is sometimes the right answer and is never the
quick one. The commit that raises it says what was read and why nothing else was
available (#187).
