#!/bin/sh
# SPDX-License-Identifier: Apache-2.0
#
# Writes a CycloneDX bill of materials for one build target. Usage:
#
#     scripts/sbom.sh <goos> <goarch> <output>
#
# The generator is cyclonedx-gomod, pinned below and fetched if it is not on PATH. It
# needs a route to the module proxy, which is the one thing this script assumes and the
# instance's runners may not have; see A18 and A24.
#
# Build constraints decide module selection, so a bill of materials is written per
# target rather than once for the release. The main component takes its version from the
# tag, which is why the release tags before it builds.
#
# -noserial: the serial number is random, and leaving it out keeps two runs over the
# same commit byte identical.
set -eu

CYCLONEDX_GOMOD_VERSION=v1.9.0

goos=${1:?usage: sbom.sh <goos> <goarch> <output>}
goarch=${2:?usage: sbom.sh <goos> <goarch> <output>}
out=${3:?usage: sbom.sh <goos> <goarch> <output>}

tool=$(command -v cyclonedx-gomod || true)
if [ -z "$tool" ]; then
	# Installed from a directory without a go.mod, so that this module's vendored
	# build is never involved in resolving the tool's own dependencies.
	( cd "$(mktemp -d)" && go install "github.com/CycloneDX/cyclonedx-gomod/cmd/cyclonedx-gomod@$CYCLONEDX_GOMOD_VERSION" )
	tool="$(go env GOPATH)/bin/cyclonedx-gomod"
fi

GOOS="$goos" GOARCH="$goarch" CGO_ENABLED=0 "$tool" app \
	-json -licenses -noserial \
	-main ./cmd/xeno \
	-output "$out" \
	.
