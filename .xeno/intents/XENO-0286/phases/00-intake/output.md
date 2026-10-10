---
intent: github.com/triplem/xeno#355
phase: 00-intake
created: "2026-10-10T10:45:43Z"
schema_version: "1.0"
runner_version: dev+90c7227.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 57b0925079ffdfab468947b6826bfc7ee272173b5ccd7b165fa1f73eb5abdc4b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: The base is mirrored into ghcr.io rather than pulled from Docker Hub with a credential
      rationale: 'The two ways out of #355 were a Docker Hub token beside the existing ghcr.io login, and a mirror. The credential is the cheaper change and raises the quota rather than removing the dependency: a free account''s quota is still a quota, and it costs a second registry credential for a pull the release does not otherwise need. The mirror leaves the release depending on one registry, the one it already authenticates to and already pushes to, and Docker Hub is reached when somebody bumps the pin rather than on every release. That is how this repository treats the rest of its supply chain, pinned to something that cannot move and fetched from somewhere accounted for. What it costs is a step that copies a digest-pinned image between registries and a change to what docs/supply-chain.md:80 describes.'
      decided_by: triplem
    - id: D-2
      chosen: The mirror is named by the release, not by the Dockerfile
      rationale: 'ARG BASE keeps docker.io/library/debian:13-slim@sha256:a29215f6… exactly as it is, and the release builds with --build-arg BASE=<mirror>@<the digest read out of the Dockerfile>. The copy preserves the manifest index unmodified, so the digest is the same on both sides and the pin stays the one pin. That is what keeps the regex manager added in #351 working, since a manager matching on docker.io/library/debian is what proposes a newer digest, and what keeps a hand build possible for somebody with no read access to this registry, which the Dockerfile''s own "building this file by hand" paragraph promises. What it costs is that the file says docker.io while the released image came from ghcr.io, so the fetch path is visible in release.yml and in the supply-chain row and not in the Dockerfile.'
      decided_by: triplem
    - id: D-3
      chosen: The release fills the mirror on a miss, so the mirror is a cache rather than a thing with an owner
      rationale: 'The image step asks whether the mirror already holds the pinned digest and copies it from Docker Hub when it does not, then builds from the mirror. So "who refreshes it" stops being a question instead of getting an answer, and Docker Hub is reached once per base bump rather than once per release. What that leaves is the one release following a bump, which still depends on Docker Hub: a bump landing on a spent quota fails that release, the same failure as 2026-10-09, once per bump rather than every time, and the next run refills the cache. The alternative was a mirror workflow triggered by the pin changing, which removes even that; it was rejected because both workflows fire on the same push and a release cut from that push races its own mirror job, which is a failure that depends on timing and would not reproduce.'
      decided_by: triplem
    - id: D-4
      chosen: v0.60.2 gets its image through the fix rather than by hand
      rationale: 'A hand-pushed image would resolve the claim the v0.60.2 binaries make today, at the price of one image in the registry that no run log accounts for, pushed with a credential the pipeline does not use, and it would leave the recovery path untested at the moment it is most clearly needed. Publishing through the dispatch this intent builds buys one way images come to exist. What it costs is the wait: v0.60.2 names an image that does not exist until both changes are merged and the trigger is run, so the recovery is an act after this intent and not a commit inside it.'
      decided_by: triplem
---

# Intake

<!-- xeno:section:problem -->
## Problem

Approved by @triplem on 2026-10-10T10:20:24Z: The mirror and the recovery trigger as decided on #355 and #356, as one intent closing both, with the `docs/supply-chain.md` row as its own commit first. Approval given by the maintainer in session on 2026-10-10 and written here by the agent on instruction, so that the intake's problem section records it. The decision is the maintainer's; the transcription is not.

> **github.com/triplem/xeno#355** — The release pulls its base image from Docker Hub unauthenticated, and a spent quota cost v0.60.2 its image
>
> The release builds the image from a base it pulls from Docker Hub without
> authenticating, so every release depends on a quota shared with every other
> unauthenticated puller on the runner's address. On 2026-10-09 that quota ran
> out and v0.60.2 was published without its image.
>
> ## What happened
>
> Run `37991808351`, at 2026-10-09T21:17:00Z, on `90c7227`:
>
>     #3 [internal] load metadata for docker.io/library/debian:13-slim@sha256:a29215f6…
>     #3 ERROR: failed to copy: httpReadSeeker: failed open: unexpected status
>     code https://registry-1.docker.io/v2/library/debian/manifests/sha256:a29215f6…:
>     429 Too Many Requests - Server message: toomanyrequests: You have reached
>     your unauthenticated pull rate limit.
>
> `docs/supply-chain.md:80` records the base as pinned by digest, which it is.
> The pin is not the problem; the fetch is.
>
> ## Why it is unauthenticated
>
> `release.yml:261` logs in, and it logs in to one registry:
>
>     printf '%s' "$REGISTRY_TOKEN" | docker login ghcr.io -u "$REGISTRY_USER" --password-stdin
>     docker build --build-arg "VERSION=$VERSION" -t "$IMAGE:$VERSION" .
>
> That credential is for `ghcr.io`, the push target. The `docker build` on the
> next line resolves `docker.io/library/debian` from the `ARG BASE` in the
> `Dockerfile`, and nothing has authenticated to `docker.io`. So the push is
> authenticated and the pull it depends on is not.
>
> ## What it cost
>
> v0.60.2 exists, is tagged, and carries its seven assets — five binaries, the
> bill of materials and `SHA256SUMS`, the same seven v0.60.1 carries. It has no
> container image. Section 13 gives the release two channels and this release
> published one of them.
>
> `release.yml:227` anticipates the shape of this: "the two are published
> independently and a registry that refuses a push is then a release missing its
> image rather than a release missing everything." What refused was the pull
> rather than the push, one step earlier, and the consequence is the one that
> comment describes.
>
> ## What this is not
>
> Not the tag added in #351. The request in the error is
> `/v2/library/debian/manifests/sha256:a29215f6…`, resolved by digest, so
> buildkit asked for exactly what the previous digest-only reference asked for,
> and the tag is absent from the request path. The six release runs before this
> one made the same request and succeeded.
>
> It is also not settled by a re-run. The re-run of `37991808351` on 2026-10-10
> is green and proves nothing: semantic-release found no new version to cut, so
> steps 10 to 14 including `image` were skipped and the pull never happened.
> That is a separate defect, reported separately.
>
> ## Two ways out, and the trade between them
>
> **Authenticate the pull.** A Docker Hub account's token in a secret, and a
> `docker login docker.io` beside the existing one. Cheapest change, and it
> raises the quota rather than removing the dependency. Costs a second registry
> credential, and a free account's quota is still a quota.
>
> **Mirror the base into `ghcr.io` and pull from there.** The release then
> depends on one registry, the one it already authenticates to and already
> pushes to, and Docker Hub is touched when somebody deliberately re-mirrors
> rather than on every release. Costs a step that copies a digest-pinned image
> between registries, and a decision about where the mirrored tag lives and who
> refreshes it. It also changes what `docs/supply-chain.md:80` describes, since
> the base would then be fetched from `ghcr.io` with Docker Hub as its origin.
>
> The second is the one consistent with how this repository treats the rest of
> its supply chain — pinned to something that cannot move, fetched from
> somewhere accounted for — but it is the larger change and the row on the page
> would have to say the new thing. Recommending it rather than taking it.
>
> ## Which package
>
> Not obvious. wp10 is the CI wrapper this project generates for adopters, and
> this is the repository's own release workflow, which no package in
> `docs/implementation-plan.md` plainly owns. Left unlabelled rather than
> labelled wrongly; if that is itself a gap in the plan, it is one worth
> knowing about.
>

The issue as it stood at 2026-10-10T10:45:43Z, read by `xeno phase start` and quoted rather than summarised. What this phase concludes about it belongs below.

<!-- xeno:section:scope -->
## Scope

This intent closes two issues with one branch, because neither half is worth releasing
alone. #355 is why the image step failed: the release builds from a base it pulls from
Docker Hub unauthenticated, and on 2026-10-09 that shared quota ran out. #356 is why the
failure is permanent: every publishing step is gated on semantic-release cutting a
version in the same run, the workflow has no other trigger, and a re-run therefore
re-runs the gate rather than the thing the gate guards. v0.60.2 is the case both were
written from — tagged, seven assets, no image.

Two changes follow, in the order the approval names.

**The base is mirrored into `ghcr.io`.** `ARG BASE` in the `Dockerfile` keeps
`docker.io/library/debian:13-slim@sha256:a29215f6…` unchanged, and the release builds
with `--build-arg BASE=<mirror>@<the digest read out of the Dockerfile>`. The image step
asks whether the mirror already holds that digest and copies it from Docker Hub when it
does not, so the mirror is a cache rather than a thing with an owner. The release then
reaches Docker Hub once per base bump instead of once per release, and reaches one
registry — the one it already authenticates to and already pushes to — on every other
release.

**The image has a way back.** `workflow_dispatch` with a version input, and the
publishing steps' condition widened to admit it. Dispatched for an existing tag the
image step downloads that release's own published binary rather than building a second
one, so the image carries the bytes `SHA256SUMS` already covers, which is the property
the step's comment says it exists to preserve.

The `docs/supply-chain.md` row at line 80 changes with the first of those: the base stays
pinned by the same digest and is fetched from `ghcr.io`, with `docker.io` as its origin
and as the cache's source. That is a documents change and so its own commit, made before
the code that follows from it.

What this does not do.

**It does not publish v0.60.2's image.** That is the dispatch being run once after both
changes are merged, which is an act and not a commit, and the decision recorded on #356
is that it happens through the fix rather than by hand. So the intent ends with v0.60.2
still naming an image that does not exist, and the recovery is the first use of the path
this intent builds.

**It does not authenticate to Docker Hub.** A second registry credential was the other
way out of #355 and was rejected: it raises the quota rather than removing the
dependency, and a free account's quota is still a quota.

**It does not add a mirror workflow triggered by the pin changing.** That would remove
the one release following a base bump from Docker Hub's reach as well, and it was
rejected because both workflows fire on the same push: a release cut from that push races
its own mirror job, which is a failure that depends on timing and would not reproduce.
What is accepted instead is that a bump landing on a spent quota fails that one release,
the same failure as 2026-10-09, once per bump rather than every time.

**It does not touch the five steps' other gates or the audit's reading of this file.**
`audit.yml:79` reads `semantic_version:` and the action's sha out of `release.yml`, and
neither moves.

<!-- xeno:section:context-rationale -->
## Why this context

Six files, 72,508 bytes. The budget is those two figures and not a round number above
them, because the scope is enumerated file by file rather than reached for with
`.github/workflows/**`, which would pull in `audit.yml`, `trivy.yml` and six others for
nothing.

**What changes.** `.github/workflows/release.yml` and `docs/supply-chain.md`. The
workflow gains the mirror and the dispatch; the page's base image row gains `ghcr.io` as
where the base is fetched from and keeps the digest it is pinned to.

**What the change is read off rather than argued from.** The `Dockerfile` holds `ARG
BASE`, which is the one place the digest is written and the value the release has to read
out in order to name the mirror, and its "building this file by hand" paragraph is the
promise the decision on #355 is keeping by leaving the `ARG` alone. `CONTRIBUTING.md`
holds the footer rule, which this branch needs for the shape no previous branch here has
needed: two issues closed by one intent take a keyword each, `Closes #355, closes #356`,
because a host reads only the first reference after one.

**What says whether the change is right.** `internal/model/supply_chain_test.go` holds
the page against the tree in both directions, and it is the test the two edits could
break in opposite ways. Its `digestPin` reads digests out of every workflow as well as
out of the `Dockerfile`, so a mirror reference written into `release.yml` with a digest
in it would become a second pin the page is asked to carry — which is the reason the
release reads the digest out of the `Dockerfile` at run time rather than repeating it,
and reading the regex is how that is known before the edit and not after. Its
`versionPin` does the same for anything shaped like `1.2.3` in a new step, and its
`tableRow` comment says the third column is not read, which is what makes the row's
change a documents change and not a test change. `internal/model/version_test.go` holds
`release.yml` and the `Dockerfile` together on the name of the binary and on the two
`ldflags` the release stamps, and the dispatch path is the first thing that would put a
binary in `dist/` by some other means than the build loop.

What is deliberately out, and why each is not needed.

`internal/model/model.go` and `internal/runner/wrapper.go` carry `RunnerImage`, which is
why #356 is more than a missing artefact: the v0.60.2 binaries generate a reference to an
image that is not there. That argument is quoted in the problem section and nothing here
changes either file, so 34,000 bytes for one function is not economy. What would catch a
drift in it is `version_test.go`, which is in scope.

`.github/workflows/audit.yml` reads two pins out of `release.yml` by `sed`.
Neither — `semantic_version:` and the action's sha — is moved by this change, and both
are in the half of the file above the first publishing step.

`.xeno/intents/**`, because the trail is the record and not input.
