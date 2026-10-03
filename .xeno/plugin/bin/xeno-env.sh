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
# There is no XENO_PLUGIN_ROOT to set. Section 7 described a resolution order above the
# vendored plugin and no longer does: the gate path reads the rule set and the templates
# from the plugin, so a root taken from the environment would make rules_hash and a
# rendered artifact depend on it, and a phase is judged only against the rule set its
# own artifact records — so a changed set stops judging the trail rather than
# disagreeing with it. The plugin is `.xeno/plugin/` found from the git root, and this
# script does not touch it (#183).

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
