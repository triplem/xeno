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
gh api "/repos/$repo" -q '"  squash title:   \(.squash_merge_commit_title)\n  squash message: \(.squash_merge_commit_message)\n  methods:        squash=\(.allow_squash_merge) merge=\(.allow_merge_commit) rebase=\(.allow_rebase_merge)"'

# Squash is also made the only method on offer. A merge commit and a rebase both
# discard the pull request description, and the Closes footer with it, so leaving them
# available would make the convention depend on which button somebody presses.
gh api -X PATCH "/repos/$repo" \
	-f squash_merge_commit_title=PR_TITLE \
	-f squash_merge_commit_message=PR_BODY \
	-F allow_squash_merge=true \
	-F allow_merge_commit=false \
	-F allow_rebase_merge=false \
	>/dev/null

echo
echo "after:"
gh api "/repos/$repo" -q '"  squash title:   \(.squash_merge_commit_title)\n  squash message: \(.squash_merge_commit_message)\n  methods:        squash=\(.allow_squash_merge) merge=\(.allow_merge_commit) rebase=\(.allow_rebase_merge)"'
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
# `verify` is required and administrators cannot bypass it, which together are the only
# arrangement in which a red gate stops a merge: the gate recomputes every verdict, and
# a requirement an administrator can step over is a report rather than a barrier.
#
# No approval is required, and that is deliberate rather than an omission. A single
# maintainer cannot approve their own pull request, so a required approval leaves
# every merge to a bypass, which is what made the bypass load bearing before this. The
# machine gate binds everybody instead, including whoever owns the repository.
#
# `strict` is false: a branch does not have to be rebased onto the current head before
# it merges. `verify` still has to pass on the branch, and requiring more would have
# made today's stacked branches unmergeable without a rebase apiece.
#
# Force push and deletion stay refused. A force push to the default branch would break
# every content hash a verdict rests on.
echo
echo "== branch protection =="
branch=$(gh api "/repos/$repo" -q .default_branch)
gh api -X PUT "/repos/$repo/branches/$branch/protection" --input - >/dev/null <<JSON
{
  "required_status_checks": { "strict": false, "contexts": ["verify"] },
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
