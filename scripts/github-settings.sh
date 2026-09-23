#!/bin/sh
# SPDX-License-Identifier: Apache-2.0
#
# Sets the repository settings the commit convention depends on. Usage:
#
#     scripts/github-settings.sh [owner/repo]
#
# The convention puts the issue reference in the footer rather than in the subject, and
# a footer survives a squash only where the squashed message is built from the pull
# request description. Without these two settings the reference is lost at the merge,
# which is worse than any of the placements considered before it.
#
# A setting is not a commit, so it is not covered by any gate and nobody notices it
# drifting. This script exists so that it is at least written down in the repository and
# can be re-applied rather than remembered. The GitLab equivalent is the squash commit
# message template, which has to contain %{description}; it is set in the project's merge
# request settings and has no API call shaped like this one.
#
# Needs gh, authenticated, with admin rights on the repository.
set -eu

repo=${1:-$(gh repo view --json nameWithOwner -q .nameWithOwner)}

echo "repository: $repo"
echo
echo "before:"
gh api "/repos/$repo" -q '"  squash title:   \(.squash_merge_commit_title)\n  squash message: \(.squash_merge_commit_message)"'

gh api -X PATCH "/repos/$repo" \
	-f squash_merge_commit_title=PR_TITLE \
	-f squash_merge_commit_message=PR_BODY \
	>/dev/null

echo
echo "after:"
gh api "/repos/$repo" -q '"  squash title:   \(.squash_merge_commit_title)\n  squash message: \(.squash_merge_commit_message)"'
echo
echo "The squashed message is now built from the pull request title and description,"
echo "so a Closes footer in the description survives the merge."
