#!/bin/sh
# SPDX-License-Identifier: Apache-2.0
#
# Prints the release notes for everything since a tag, grouped by Conventional Commit
# type. Usage:
#
#     scripts/changelog.sh [previous-tag]
#
# The tag is an argument rather than something this script works out, because the
# release creates the new tag before it builds, so by the time the notes are written
# "the last tag" is the one being released and the range would be empty.
set -eu

# $# rather than ${1:-...}: an empty argument means "no previous tag, take the whole
# history", which is the first release, and that has to be distinguishable from no
# argument at all. With the default expansion the first release would look for a tag,
# find the one just created for it, and report an empty range.
if [ $# -ge 1 ]; then
	last=$1
else
	last=$(git describe --tags --abbrev=0 --match 'v[0-9]*' 2>/dev/null || true)
fi
range=${last:+$last..HEAD}
range=${range:-HEAD}

section() {
	# $1 heading, $2.. grep patterns over the subject
	heading=$1
	shift
	body=$(git log --no-merges --format='%s|%h' "$range" | while IFS='|' read -r subject sha; do
		for pat in "$@"; do
			case "$subject" in
			$pat)
				# Strip the type and scope, keep the sentence and the short sha.
				echo "- ${subject#*: } (\`$sha\`)"
				break
				;;
			esac
		done
	done)
	[ -n "$body" ] || return 0
	printf '### %s\n\n%s\n\n' "$heading" "$body"
}

section 'Breaking changes' 'feat!:*' 'fix!:*' 'perf!:*' 'refactor!:*' 'feat(*)!:*' 'fix(*)!:*' 'perf(*)!:*' 'refactor(*)!:*'
section 'Features' 'feat:*' 'feat(*):*'
section 'Fixes' 'fix:*' 'fix(*):*' 'perf:*' 'perf(*):*'

if [ -n "$last" ]; then
	printf 'Full history since %s.\n' "$last"
fi
