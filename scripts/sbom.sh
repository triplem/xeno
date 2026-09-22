#!/bin/sh
# SPDX-License-Identifier: Apache-2.0
#
# Writes a CycloneDX 1.5 software bill of materials for a built binary, read from the
# build information the Go toolchain embeds in it. Usage:
#
#     scripts/sbom.sh <binary> <version> > sbom.cdx.json
#
# The binary is the source rather than go.mod, because what a release ships is what the
# binary contains: a module the build did not reach never appears, and a replace never
# goes unnoticed. It also needs no network and no go.sum, which matters for an instance
# whose runners have no route out (WP0, "Runners and egress"). Recorded as A24.
#
# No serial number is written. It is optional in CycloneDX, and leaving it out keeps two
# runs over the same binary byte identical, which is worth more here than a fresh uuid.
set -eu

bin=${1:?usage: sbom.sh <binary> <version>}
version=${2:?usage: sbom.sh <binary> <version>}
name=$(basename "$bin")
stamp=${SOURCE_DATE_EPOCH:+$(date -u -d "@$SOURCE_DATE_EPOCH" +%Y-%m-%dT%H:%M:%SZ)}
stamp=${stamp:-$(date -u +%Y-%m-%dT%H:%M:%SZ)}

components=$(go version -m "$bin" | awk '
	$1 == "dep" {
		# dep <module> <version> [<hash>]
		printf "%s\t%s\n", $2, $3
	}' | sort -u | awk -F'\t' '
	{
		if (NR > 1) printf ",\n"
		printf "    {\n"
		printf "      \"type\": \"library\",\n"
		printf "      \"name\": \"%s\",\n", $1
		printf "      \"version\": \"%s\",\n", $2
		printf "      \"purl\": \"pkg:golang/%s@%s\"\n", $1, $2
		printf "    }"
	}')

cat <<JSON
{
  "bomFormat": "CycloneDX",
  "specVersion": "1.5",
  "version": 1,
  "metadata": {
    "timestamp": "$stamp",
    "tools": [
      { "name": "scripts/sbom.sh", "vendor": "conet Deutschland GmbH" }
    ],
    "component": {
      "type": "application",
      "name": "$name",
      "version": "$version",
      "purl": "pkg:golang/github.com/triplem/xeno@$version",
      "licenses": [ { "license": { "id": "Apache-2.0" } } ]
    }
  },
  "components": [
$components
  ]
}
JSON
