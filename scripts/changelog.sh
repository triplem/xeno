#!/bin/sh
# SPDX-License-Identifier: Apache-2.0
#
# Prints the release notes for everything since the last tag, grouped by Conventional
# Commit type. The notes go into the release rather than into a file in the repository:
# a pipeline that commits a CHANGELOG.md writes into the repository it is verifying and
# starts the next pipeline doing it. Recorded as A23.
set -eu

last=$(git describe --tags --abbrev=0 --match 'v[0-9]*' 2>/dev/null || true)
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
