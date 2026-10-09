#!/bin/sh
# SPDX-License-Identifier: Apache-2.0
#
# Sets the repository settings this project depends on. Usage:
#
#     scripts/github-settings.sh [owner/repo]
#
# Three groups, each with its reason at its step: the merge settings the commit
# convention rests on, private vulnerability reporting, which SECURITY.md names as the
# channel, and the branch protection of the default branch.
#
# A setting is not a commit, so no gate covers it and nobody notices it drifting. This
# script exists so that every one of them is written down here and can be re-applied
# rather than remembered. Running it twice changes nothing the second time.
#
# One thing is reported and not set: the visibility, because the flip publishes every
# commit and the whole trail at once and is #47, a decision a person made.
#
# The GitLab equivalent of the merge group is the squash commit message template, which
# has to contain %{description}; it is set in the project's merge request settings and
# has no API call shaped like this one.
#
# Needs gh, authenticated, with admin rights on the repository.
set -eu

repo=${1:-$(gh repo view --json nameWithOwner -q .nameWithOwner)}

echo "repository: $repo"

echo
echo "== visibility, reported and not set =="
gh api "/repos/$repo" -q '"  visibility: \(.visibility)"'

echo
echo "== merge settings =="
echo "before:"
gh api "/repos/$repo" -q '"  squash title:   \(.squash_merge_commit_title)\n  squash message: \(.squash_merge_commit_message)\n  methods:        squash=\(.allow_squash_merge) merge=\(.allow_merge_commit) rebase=\(.allow_rebase_merge)\n  delete branch:  \(.delete_branch_on_merge)"'

# Squash is also made the only method on offer. A merge commit and a rebase both
# discard the pull request description, and the Closes footer with it, so leaving them
# available would make the convention depend on which button somebody presses.
#
# The head branch is deleted on merge, which is not tidiness. GitHub retargets a pull
# request whose base branch merges only when that branch is deleted, and without it a
# stacked pull request goes on pointing at a branch that has already left for main and
# merges into it: #182 and #185 did exactly that, read as merged, and put nothing on
# main. Deleting the branch is what makes the retarget happen, and it also means a
# branch that is still there is a branch with something in it (#186).
gh api -X PATCH "/repos/$repo" \
	-f squash_merge_commit_title=PR_TITLE \
	-f squash_merge_commit_message=PR_BODY \
	-F allow_squash_merge=true \
	-F allow_merge_commit=false \
	-F allow_rebase_merge=false \
	-F delete_branch_on_merge=true \
	>/dev/null

echo
echo "after:"
gh api "/repos/$repo" -q '"  squash title:   \(.squash_merge_commit_title)\n  squash message: \(.squash_merge_commit_message)\n  methods:        squash=\(.allow_squash_merge) merge=\(.allow_merge_commit) rebase=\(.allow_rebase_merge)\n  delete branch:  \(.delete_branch_on_merge)"'
echo
echo "The squashed message is now built from the pull request title and description,"
echo "so a Closes footer in the description survives the merge, and squash is the only"
echo "method on offer, so no other path can discard it."

# Private vulnerability reporting, which SECURITY.md names as the channel instead of an
# address. It exists on a public repository and is off by default, so a file naming it
# would be pointing at something switched off. Reversible, and off again by DELETE.
echo
echo "== private vulnerability reporting =="
gh api -X PUT "/repos/$repo/private-vulnerability-reporting" >/dev/null
gh api "/repos/$repo/private-vulnerability-reporting" -q '"  enabled: \(.enabled)"'

# The protection of the default branch, per section 7 and the enforcement block of
# project.yaml, and it is what #84 decided.
#
# Every check that can run on a pull request is required, and administrators cannot
# bypass any of them: together those are the only arrangement in which a red gate stops
# a merge, because a requirement an administrator can step over is a report rather than
# a barrier.
#
# All five rather than `verify` alone, which is what this script set until #186.
# `verify` recomputes every verdict, and the other four each answer a question the trail
# cannot, so a merge past any of them red is a merge past something nobody looked at.
# Six commits reached main in one morning while `audit` was red, each from a pull
# request that was honestly green, because `audit` was advisory and did not run on a
# pull request at all.
#
# The names are the job ids of the workflows, which is what GitHub reports a check
# under. Three of them said `scan` until #186 and could not be required individually,
# since a required check is matched by that name and three identical ones name no
# workflow. Renaming a check means this list and the workflow's job id have to move
# together, and a required context that never reports blocks every merge, so the rename
# lands first.
#
# No approval is required, and that is deliberate rather than an omission. A single
# maintainer cannot approve their own pull request, so a required approval leaves
# every merge to a bypass, which is what made the bypass load bearing before this. The
# machine gate binds everybody instead, including whoever owns the repository.
#
# `strict` is true: a branch has to contain the current head of the default branch
# before it merges. It was false, which is how a branch that had diverged from main
# merged while every check on it was green — the checks had passed, on a tree that was
# not the one the merge would produce. The cost is a rebase whenever main moves, which
# is small where one pull request is open at a time and is the arrangement #186 asks
# for anyway.
#
# Force push and deletion stay refused. A force push to the default branch would break
# every content hash a verdict rests on.
echo
echo "== branch protection =="
branch=$(gh api "/repos/$repo" -q .default_branch)
gh api -X PUT "/repos/$repo/branches/$branch/protection" --input - >/dev/null <<JSON
{
  "required_status_checks": {
    "strict": true,
    "contexts": ["verify", "audit", "gitleaks", "lint", "gosec", "govulncheck", "trivy"]
  },
  "enforce_admins": true,
  "required_pull_request_reviews": null,
  "restrictions": null,
  "allow_force_pushes": false,
  "allow_deletions": false
}
JSON
echo "  branch:         $branch"
# One call, one expression per line, so that no line of it runs past the width the
# project wraps at.
gh api "/repos/$repo/branches/$branch/protection" -q '
	"  checks:         \(.required_status_checks.contexts // ["none"] | join(", "))",
	"  up to date:     \(.required_status_checks.strict // false)",
	"  approvals:      \(.required_pull_request_reviews
	                       .required_approving_review_count // 0)",
	"  admins bypass:  \(.enforce_admins.enabled | not)",
	"  force pushes:   \(.allow_force_pushes.enabled)",
	"  deletions:      \(.allow_deletions.enabled)"
'
