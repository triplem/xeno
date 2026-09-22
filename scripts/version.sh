#!/bin/sh
# SPDX-License-Identifier: Apache-2.0
#
# Derives the next version from Conventional Commits since the last tag and prints it,
# or prints nothing where no commit calls for a release. The version number lives here
# and in the tag, never in a source file, which is what WP0 asks of this package.
#
# Below 1.0.0 a breaking change raises the minor rather than the major, which is the
# usual reading of semver while a project is unstable. Recorded as A22.
set -eu

last=$(git describe --tags --abbrev=0 --match 'v[0-9]*' 2>/dev/null || true)
if [ -n "$last" ]; then
	range="$last..HEAD"
	v=${last#v}
else
	# No tag yet: everything in the history counts, and the first release is 0.1.0.
	range="HEAD"
	v="0.0.0"
fi

major=$(echo "$v" | cut -d. -f1)
minor=$(echo "$v" | cut -d. -f2)
patch=$(echo "$v" | cut -d. -f3)

bump=none
while IFS= read -r line; do
	case "$line" in
	# A "!" before the colon, or a BREAKING CHANGE footer, is the loudest signal.
	feat!:* | fix!:* | perf!:* | refactor!:* | feat\(*\)!:* | fix\(*\)!:* | perf\(*\)!:* | refactor\(*\)!:*)
		bump=breaking ;;
	"BREAKING CHANGE:"* | "BREAKING-CHANGE:"*)
		bump=breaking ;;
	feat:* | feat\(*\):*)
		[ "$bump" = breaking ] || bump=feat ;;
	fix:* | fix\(*\):* | perf:* | perf\(*\):*)
		[ "$bump" = none ] && bump=fix ;;
	esac
done <<EOT
$(git log --format='%s%n%b' "$range")
EOT

case "$bump" in
none) exit 0 ;;
breaking)
	if [ "$major" -eq 0 ]; then
		minor=$((minor + 1)); patch=0
	else
		major=$((major + 1)); minor=0; patch=0
	fi ;;
feat) minor=$((minor + 1)); patch=0 ;;
fix) patch=$((patch + 1)) ;;
esac

echo "${major}.${minor}.${patch}"
