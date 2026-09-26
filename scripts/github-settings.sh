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
# Two things are reported and not set. The visibility, because the flip publishes every
# commit and the whole trail at once and is #47, a decision a person made. And the two
# enforcement requirements #84 leaves open, because they interact with how a one person
# project merges.
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
# project.yaml. What is applied is what the project has decided: a pull request with one
# approval, and no force push or deletion. Force pushing the default branch would break
# every content hash a verdict rests on, which is why it is written down rather
# than left to habit.
#
# Requiring the verify check and removing the administrator bypass are the two
# requirements xeno enforcement check reports as unmet, and #84 holds that decision
# with the calls written out. They are not commented out here, so that nothing in this
# file reads as almost done.
echo
echo "== branch protection =="
branch=$(gh api "/repos/$repo" -q .default_branch)
gh api -X PUT "/repos/$repo/branches/$branch/protection" --input - >/dev/null <<JSON
{
  "required_status_checks": null,
  "enforce_admins": false,
  "required_pull_request_reviews": { "required_approving_review_count": 1 },
  "restrictions": null,
  "allow_force_pushes": false,
  "allow_deletions": false
}
JSON
echo "  branch:         $branch"
# One call, one expression per line, so that no line of it runs past the width the
# project wraps at.
gh api "/repos/$repo/branches/$branch/protection" -q '
	"  approvals:      \(.required_pull_request_reviews
	                       .required_approving_review_count)",
	"  force pushes:   \(.allow_force_pushes.enabled)",
	"  deletions:      \(.allow_deletions.enabled)",
	"  admins bypass:  \(.enforce_admins.enabled | not)   see #84",
	"  checks:         \(.required_status_checks.contexts // ["none"]
	                       | join(", "))   see #84"
'
