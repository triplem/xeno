#!/bin/sh
# SPDX-License-Identifier: Apache-2.0
#
# Prepends the notes for a release to CHANGELOG.md. Usage:
#
#     scripts/changelog-update.sh <version> [previous-tag]
#
# The newest release stands at the top, so the file reads in the order a reader wants it
# and an entry once written is never touched again. Insertion is at a marker rather than
# at a line count, so the prose above it can be reworded without breaking this script.
#
# The file is derived and the header says so: the commit history is the source, this is
# a rendering of it.
set -eu

version=${1:?usage: changelog-update.sh <version> [previous-tag]}
prev=${2:-}
file=CHANGELOG.md
marker='<!-- releases below, newest first -->'

if [ ! -f "$file" ]; then
	cat > "$file" <<HEAD
# Changelog

Derived from the commit history by the release pipeline. The commits are the source;
editing this file by hand changes the rendering and not what it renders.

$marker
HEAD
fi

grep -qF "$marker" "$file" || { echo "$file has no marker line: $marker" >&2; exit 1; }

notes=$(sh "$(dirname "$0")/changelog.sh" "$prev")
[ -n "$notes" ] || notes='No categorised changes.'

awk -v marker="$marker" -v entry="## $version, $(date -u +%Y-%m-%d)

$notes" '
	$0 == marker { print; print ""; print entry; next }
	{ print }
' "$file" > "$file.tmp"
mv "$file.tmp" "$file"
