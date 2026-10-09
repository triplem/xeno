#!/bin/sh
# SPDX-License-Identifier: Apache-2.0
#
# Creates the one label Xeno asks a project's tracker for. Section 12 says an issue
# becomes an intent only where it carries the label `xeno-approved` and a comment whose
# first line is `/xeno approved`; the label is a person's act on the host, which `xeno
# init` names among the settings it does not make, and this script is that act written
# down so that it can be run rather than remembered (#332).
#
#     .xeno/plugin/bin/xeno-labels.sh github owner/repo
#     .xeno/plugin/bin/xeno-labels.sh gitlab group/project
#
# Needs gh or glab, authenticated, with the right to manage labels. Running it twice
# changes nothing the second time: a label that is there is reported and left alone.
# That rests on the read before the write asking for the whole answer: a list read with
# its default page stops at thirty labels, and the second run on a repository past that
# tried to create the label again and exited 1 (#346). gh reads by name, which is
# exact after grep because --search also matches descriptions; glab has no search and
# reads a page of a hundred, the most the host returns, so past a hundred labels the
# GitLab branch meets the same fault. The description carries the other half of the
# clause, so that a reader of the label list sees why setting the label alone starts
# nothing.

set -eu

host=${1:?usage: xeno-labels.sh github|gitlab OWNER/REPO}
project=${2:?usage: xeno-labels.sh github|gitlab OWNER/REPO}

label=xeno-approved
description="An issue xeno intent start may start; needs a comment beginning /xeno approved as well"
colour=1d76db

case "$host" in
github)
	if gh label list --search "$label" --repo "$project" --json name --jq '.[].name' | grep -qx "$label"; then
		echo "$label exists on $project; nothing changed"
	else
		gh label create "$label" --repo "$project" --description "$description" --color "$colour"
		echo "$label created on $project"
	fi
	;;
gitlab)
	if glab label list --per-page 100 --repo "$project" --output json 2>/dev/null | grep -q "\"name\":\"$label\""; then
		echo "$label exists on $project; nothing changed"
	else
		glab label create --repo "$project" --name "$label" --description "$description" --color "#$colour"
		echo "$label created on $project"
	fi
	;;
*)
	echo "unknown host $host: github or gitlab" >&2
	exit 1
	;;
esac
