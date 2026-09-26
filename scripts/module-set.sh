#!/bin/sh
# One bill of materials covers every target only while every target selects the same
# modules. Build constraints can change that selection, and a CycloneDX gomod document
# describes modules, so this compares module sets and not package sets: package sets
# differ on every platform by way of the standard library and say nothing about the
# document.
#
# Usage: scripts/module-set.sh <os/arch>...
# Exits non-zero, naming the targets that disagree, where one document would be a claim
# about targets it does not describe.
set -eu

first=""
firstset=""
diverged=""

for target in "$@"; do
	os=${target%/*}
	arch=${target#*/}
	set=$(GOOS="$os" GOARCH="$arch" go list -deps -f '{{if .Module}}{{.Module}}{{end}}' ./cmd/xeno | sort -u)
	if [ -z "$first" ]; then
		first=$target
		firstset=$set
		continue
	fi
	if [ "$set" != "$firstset" ]; then
		diverged="$diverged $target"
		printf '%s and %s select different modules:\n' "$first" "$target" >&2
		printf '%s\n' "$firstset" > "${TMPDIR:-/tmp}/module-set-first.$$"
		printf '%s\n' "$set" > "${TMPDIR:-/tmp}/module-set-other.$$"
		diff "${TMPDIR:-/tmp}/module-set-first.$$" "${TMPDIR:-/tmp}/module-set-other.$$" >&2 || true
		rm -f "${TMPDIR:-/tmp}/module-set-first.$$" "${TMPDIR:-/tmp}/module-set-other.$$"
	fi
done

if [ -n "$diverged" ]; then
	printf '\none bill of materials cannot describe these targets; generate one per target\n' >&2
	exit 1
fi

printf 'every target selects the same modules:\n%s\n' "$firstset"
