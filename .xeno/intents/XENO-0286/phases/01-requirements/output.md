---
intent: github.com/triplem/xeno#355
phase: 01-requirements
created: "2026-10-10T12:12:01Z"
schema_version: "1.0"
runner_version: dev+90c7227.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cd40d0813e2c9424330be0f81c028ff8b8dc3061ce8c262c3d7192fe1661d336
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. **A release pulls its base from `ghcr.io` and not from Docker Hub.** The image step
   builds with `--build-arg BASE=<mirror>@<digest>`, and the only registry it
   authenticates to is the one it already pushes to. Shown on the run log of a release:
   the `load metadata` line names `ghcr.io`, and no request reaches
   `registry-1.docker.io` on a release whose base has not moved.

2. **The copy preserves the digest.** `docker buildx imagetools create` applied to a
   manifest index, with one source and no annotations, yields the same `sha256:` on both
   sides. This is the claim the whole arrangement rests on — it is what makes the mirror
   the same pin rather than a second one — and it is measured against a registry rather
   than taken from documentation. If it does not hold, the design is wrong and not the
   implementation.

3. **The digest is written once.** `ARG BASE` in the `Dockerfile` is the only place the
   base image's digest appears. The release reads it out of that file at run time; it does
   not carry a copy. `internal/model/supply_chain_test.go`'s `digestPin` reads digests out
   of every workflow, so a copy in `release.yml` would arrive as a second pin the page is
   asked to carry, and that test passing is how this criterion is checked rather than
   reviewed.

4. **`ARG BASE` names `docker.io` still, with its tag and its digest.** Unchanged, byte
   for byte: the regex manager #351 added matches on that reference and proposes the next
   digest for it, and the `Dockerfile`'s own "building this file by hand" paragraph
   promises a build to somebody with no read access to this registry. Checked by `git
   diff main -- Dockerfile` being empty, and by running the `customManager` regex against
   the file and reading what it resolves to.

5. **Docker Hub is reached on a miss and on nothing else.** The step asks whether the
   mirror already holds the pinned digest and copies from Docker Hub only when it does
   not. So a release whose base has not moved makes no request to Docker Hub at all, and
   the first release after a bump makes one.

6. **The image step has a second way in, and it publishes an image for an existing tag.**
   `workflow_dispatch` with a version input reaches the image step for a version already
   released, and publishes `ghcr.io/triplem/xeno:<that version>`. Demonstrated by running
   it for v0.60.2 after the merge, which is the recovery #356 decided and the first use of
   the path.

7. **A dispatch publishes the bytes that release already published.** The image carries
   the binary from that release's own assets, verified against that release's
   `SHA256SUMS` before it is copied in, rather than a second build of the same source.
   That is the property the image step's comment says it exists to preserve, and on a
   dispatch it is the only way to have it: the tree at `main` is no longer the tree the
   tag was cut from.

8. **A dispatch for a version with no release fails, and publishes nothing.** #356 names
   the risk the second way in creates — "dispatched with the wrong version it would push
   an image tagged for a release it was not built from" — so the download and the
   checksum comparison are what stand in the way, and they come before the push.

9. **A dispatch cuts no tag and publishes no release.** semantic-release does not run on
   a dispatch, or its running changes nothing: no tag, no release, no assets, no change to
   `main`. The binary channel's four steps — `build`, the module set, the bill of
   materials, `checksums and upload` — stay gated on a version cut in the same run, and
   only the image step admits the other trigger.

10. **The image step does not depend on a step that a dispatch skips.** `IMAGE` is
    written into `GITHUB_ENV` by the build step today, and a dispatch skips the build
    step. After this change the image step resolves the registry for itself, so the two
    triggers cannot disagree about where the image goes.

11. **The supply-chain page says where the base is fetched from.** The row keeps the
    digest it is pinned to and names `ghcr.io` as what the release fetches it from, with
    `docker.io` as its origin and as the cache's source. It lands in its own commit,
    before the code that follows from it.

12. **Every pin the change introduces is on the page, in both directions.**
    `internal/model/supply_chain_test.go` passes. `docker buildx` is the same docker the
    page's last row already says is pinned by somebody else; if the change makes that row
    say less than it should, the row is what moves.

13. **The whole gate suite passes on the tree as it stands.** `go build`, `go test ./...`,
    `gofmt -l .` empty outside `vendor/`, `go vet ./...`, and `./xeno gate verify` at exit
    0, with the change applied rather than reverted out.

<!-- xeno:section:non-goals -->
## Non goals

**Publishing v0.60.2's image.** The dispatch is built here and run after the merge. It is
an act rather than a commit, and until it is run v0.60.2 goes on naming an image that is
not there. The decision on #356 is that this is the cost of having one way images come to
exist, rather than a hand-pushed image no run log accounts for.

**Authenticating to Docker Hub.** Rejected on #355: it raises the quota rather than
removing the dependency, and costs a second registry credential for a pull the release
would no longer need.

**Removing Docker Hub from the release's reach entirely.** The one release following a
base bump still fetches from Docker Hub, because the mirror is filled on a miss and a
bump is a miss. A separate mirror workflow triggered by the pin changing would remove even
that and was rejected: both workflows fire on the same push, so a release cut from that
push races its own mirror job, and that is a failure which depends on timing and would
not reproduce.

**A second architecture for the image.** `linux/amd64` only, as today. #282 holds that
question and nothing here touches it; a mirror of a manifest index is not a build of a
second platform.

**Pinning what goes into the image.** `ca-certificates` and `git` resolve against
Debian's suite on the day, and the page's own paragraph says why pinning them would turn
each move of that suite into a failed release. Mirroring the base changes where the base
comes from and not what `apt-get` reaches.

**A dispatch that can rebuild the binaries.** The dispatch reaches the image step and not
the binary channel. Letting it rebuild would give one version two sets of binaries whose
only difference is the tree they were built from, which is the asymmetry #356 was written
about rather than a cure for it.

**Anything in the `Dockerfile`.** `ARG BASE` stays as it is, which is criterion 4 and the
reason the regex manager and a hand build both keep working. A file that names
`docker.io` while the released image came from `ghcr.io` is the cost that buys both, and
it is paid in `release.yml` and on the supply-chain page.

**The re-run that proves nothing.** #356 reports that the re-run of `37991808351` is
green because semantic-release correctly found nothing to do, which is indistinguishable
from a recovery. Making a green re-run distinguishable from a recovery is a separate
defect about what a run log says, and this intent gives the recovery a trigger of its own
rather than changing what a re-run reports.

<!-- xeno:section:constraints -->
## Constraints

**One credential, which is the point.** `secrets.GITHUB_TOKEN` with `packages: write` is
what the job already has, and the mirror has to be readable and writable by it or the
change has bought a second credential after all. That constrains where the mirror lives:
under this repository's own owner, as a package the release itself creates on the first
miss, so that the token's scope reaches it without anybody granting anything.

**The mirror's visibility is not this workflow's to set.** A container package `ghcr.io`
creates on a push is private, and only the release reads it. Nothing here makes it public,
and nothing here depends on it being public — that is what criterion 4 is for: a hand
build goes to `docker.io`, which the `Dockerfile` still names.

**`docker buildx` is the hosted runner's, like `docker`.** The copy uses a plugin of the
docker CLI the runner image ships, which this repository does not pin and cannot. The
page's last row already says that of `docker`; what it may have to say is that the same
unpinned tool now copies as well as builds.

**The gates above the publishing steps run on a dispatch too.** `format`, `vet`, `test`
and `verify` execute against `main` as it stands, not against the tree the dispatched tag
was cut from. That is a cost rather than a defect — a dispatch that publishes an image for
an old tag from a `main` that does not build is a dispatch that should stop — and it is
named here so that it is not discovered as a surprise.

**The image's bytes come from the release, so the tree is not consulted for them.** On a
dispatch the tree at `main` is no longer the tree the tag was cut from, which is exactly
why criterion 7 asks for the published asset and its `SHA256SUMS` rather than a build. The
`Dockerfile` only copies, so this costs nothing: the same file serves both triggers.

**`concurrency: group: release` is shared.** A dispatch and a push to `main` queue behind
each other rather than running together, which is what is wanted: both can push to the
same registry path, and `cancel-in-progress: false` means neither interrupts the other.

**`persist-credentials: false` on the checkout.** Anything that talks to the host needs
its token named in the step, as `checksums and upload` already does with `GH_TOKEN`. The
download on a dispatch is subject to the same rule.

**The release step's output drives four conditions and must not drive the fifth.**
`steps.release.outputs.new_release_published` is read by `build`, the module set, the bill
of materials and `checksums and upload`, and those stay as they are. Only the image step's
condition widens, and it has to widen in a way that cannot make a dispatch publish
binaries or a push skip them.

**No second dependency.** `go.yaml.in/yaml/v3` is the one, and nothing here is a Go
change. The workflow may use what the runner image ships and the host's own CLI, and
adding a tool to pin — `crane`, `regctl`, `skopeo` — would be a decision rather than a
step.
