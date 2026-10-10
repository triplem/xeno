---
intent: github.com/triplem/xeno#355
phase: 03-implementation
created: "2026-10-10T13:07:15Z"
schema_version: "1.0"
runner_version: dev+b9beef5.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a18fbb9eef30fc310d01c4800521cb39248878e647d3df0934e82abb4f3f56bd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

Two commits, in the order the approval on #355 names.

**`docs(supply-chain)`, first and on its own.** The base image's row keeps
`sha256:a29215f6…`, the manifest index, by digest, and its third column becomes `ghcr.io,
mirrored from docker.io`. Two paragraphs follow the one about what a published digest
promises: where the base is fetched from and where it comes from, with the 2026-10-09
failure as the reason; and that the pin does not move with the fetch path, because `ARG
BASE` is unchanged and the copy preserves the index, so a reader of either place is told
the same bytes. The `docker` row's text is not touched — `exempt` in
`internal/model/supply_chain_test.go` is keyed on that row's exact first cell, so moving
the words would have made the documents commit a test change — and the sentence that the
same unpinned CLI now copies as well as builds and pushes went into the prose instead.

**`.github/workflows/release.yml`.** Four edits.

*The trigger.* `workflow_dispatch` with a required `version` input, described as the
version of an existing release whose image is missing, without the leading `v`. The
comment above it says what it is for, in the terms #356 reports: a re-run restarts the job
at the gate rather than at the thing the gate guards, and is green because
semantic-release correctly does nothing.

*`where the image goes`.* A new step, unconditional, after the two setup actions and
before the gates: `echo "IMAGE=ghcr.io/${GITHUB_REPOSITORY,,}" >> "$GITHUB_ENV"`. The
build step's eleven lines that computed the same value are gone and it reads `$IMAGE`
for the `ldflags`. This is criterion 10. The value was produced by a step a dispatch
skips, so the image step would have resolved its registry to the empty string and pushed
to a name rather than failing. It is a step and not a job level `env` because `${VAR,,}`
is a shell expansion and GitHub's expressions have no lowercase.

*The `release` step is gated on `github.event_name == 'push'`.* One line. The four
conditions below it are untouched and false by construction on a dispatch, because a step
that did not run has no outputs. It also stops a dispatch cutting a version out of
whatever has landed on `main` since the last release.

*The image step.* Its condition admits the dispatch, `VERSION` becomes
`${{ steps.release.outputs.new_release_version || inputs.version }}`, and `GH_TOKEN`
joins the two registry variables in `env`. The script gains three blocks before the
build:

On a dispatch, the binary. `gh release download` for `xeno-$VERSION-linux-amd64` into
`dist/` and for `SHA256SUMS` into the working directory, the one line for that binary cut
out with `grep`, `sha256sum -c` against it, `chmod 0755`, and both checksum files removed
so that `.dockerignore`'s `!dist` leaves the context holding only what the `COPY` names.
The `chmod` is the defect P2 found by running the download: the asset arrives 0644, `go
build` produces 0755, and `COPY` preserves the mode.

The base, read rather than restated. `sed -n 's/^ARG BASE=\(.*\)$/\1/p' Dockerfile`, then
`${base##*@}` for the digest and `${base%@*}` with `${...##*:}` for the tag. `test -n
"$base"` is what stops a rewritten `ARG` line arriving as an empty reference. Both are
echoed, so the run log says which base and which mirror.

The mirror, asked and proved by one command. `docker buildx imagetools inspect
"$mirror@$base_digest"` decides whether a copy is needed; on a miss,
`docker buildx imagetools create --tag "$mirror:$base_tag" "$base"` copies from the
origin and the same inspect runs again, so a copy that re-serialised the index fails
there with the reference it could not resolve rather than three lines later as a build
that cannot find its base. `$mirror` is `$IMAGE/base`, derived from the one place the
registry is named.

Then `docker build` gains `--build-arg "BASE=$mirror@$base_digest"` beside the `VERSION`
it already passed. The push, and the `docker image inspect` that prints the published
digest, are unchanged.

**What was run against the real thing rather than reasoned about.** The dispatch block
was executed for v0.60.2 outside the repository, with the real token and the real
assets: the download succeeds, `sha256sum -c` prints `xeno-0.60.2-linux-amd64: OK`, the
`chmod` leaves `-rwxr-xr-x`, the context holds the binary and nothing else, and the binary
reports `xeno 0.60.2`. Criterion 8's two failure modes were run too: `v9.9.9` exits 1 with
"release not found", and an asset pattern naming a version the release does not carry
exits 1 with "no assets match the file pattern". Both stop under `set -eu` before the
`docker login`, so nothing is pushed.

**What the `Dockerfile` carries after this.** Nothing new. `git diff main -- Dockerfile`
is empty, which is criterion 4, and it is what the regex manager of #351 and a hand build
both depend on.

<!-- xeno:section:deviations -->
## Deviations from the design

One, and it is a line the design did not know it would have to leave alone.

The design said the row for `docker` "says it builds and pushes the image; it now also
copies one". It does not say so on the row. `exempt` in
`internal/model/supply_chain_test.go` is a map keyed on a row's exact first cell —
`"`docker`, which builds and pushes the image"` — and `TestEveryRowOfTheSupplyChainPageIsAPinTheTreeCarries`
looks the row up by that string. Moving the words would have made the row unexempt, which
would have failed the test, which would have made the documents commit a code change as
well. So the sentence went into the prose paragraph beside the fetch path instead, which
is where the reason lives anyway, and the documents commit stayed what the standing rule
asks for: its own commit, before the code, with nothing else in it.

Nothing else departs from the design. The four workflow edits are the four the impact
section named, in that order, and the two assumptions it recorded are still assumptions:
neither the copy's digest preservation nor the mirror's readability by the same token on a
later run can be exercised without a runner, which is why one of them is a check the
pipeline makes rather than a premise, and the other has a failure mode that costs a copy
per release rather than a release.
