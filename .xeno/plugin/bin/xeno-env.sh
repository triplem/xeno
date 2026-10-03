#!/bin/sh
# SPDX-License-Identifier: Apache-2.0
#
# The thin entry point section 7 describes: it normalises the environment and then
# starts the runner. "The hook calls a thin entry point that normalises the environment
# before starting the runner. The runner only ever sees XENO_* and does not know which
# harness it runs under."
#
#     .xeno/plugin/bin/xeno-env.sh phase start --intent KEY --phase 00
#
# It sets what it can establish and nothing it cannot. A variable already set is left
# alone, because the caller knew something this script does not.
#
# XENO_PLUGIN_ROOT is deliberately not set here, and that is the one thing this script
# does not do. Section 7 ranks a --plugin-root argument and XENO_PLUGIN_ROOT above the
# vendored tree, and internal/gates reads internal/rules and internal/template, both of
# which resolve from the vendored directory — so an override would make rules_hash and a
# rendered artifact depend on the environment. Measured on this repository, a valid rule
# tree that differs leaves `xeno gate verify` at exit 0 over 273 verdicts while G-Policy
# silently stops judging every phase that recorded the previous hash, because A74 judges
# a phase only against the set its own artifact names. Enabling that before G-Supply
# exists would remove the only thing that would catch it. #183 carries the finding and
# the decision is section 7's.
set -eu

# The git root, so the script works from anywhere inside a working tree.
root=$(git rev-parse --show-toplevel 2>/dev/null || pwd)

# Always .xeno/local/, which is what section 7 says it is. Exported so that whatever the
# runner starts sees the same answer rather than each process deciding.
: "${XENO_PLUGIN_DATA:=$root/.xeno/local}"
export XENO_PLUGIN_DATA

# Recorded only, and read by the runner into section 5's `tool` field. The client's own
# variables are read here and nowhere else: this is the one place that is allowed to
# know which harness it is, and the runner is not.
if [ -z "${XENO_HARNESS:-}" ]; then
	if [ -n "${CLAUDE_PLUGIN_ROOT:-}" ] || [ -n "${CLAUDECODE:-}" ]; then
		XENO_HARNESS=claude-code
	elif [ -n "${CODEX_HOME:-}" ]; then
		XENO_HARNESS=codex
	else
		XENO_HARNESS=unknown
	fi
fi
export XENO_HARNESS

# The harness's own version, for the one field of section 5 the runner cannot know.
# Absent where the client does not say, because a plausible value in a field nobody
# produced is worse than an absent one (A35, A81).
if [ -z "${XENO_HARNESS_VERSION:-}" ] && [ "$XENO_HARNESS" = claude-code ]; then
	XENO_HARNESS_VERSION=$(claude --version 2>/dev/null | awk '{print $1}' || true)
fi
[ -n "${XENO_HARNESS_VERSION:-}" ] && export XENO_HARNESS_VERSION

exec xeno "$@"
