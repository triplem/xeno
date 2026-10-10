---
intent: github.com/triplem/xeno#355
phase: 02-design
created: "2026-10-10T12:15:52Z"
schema_version: "1.0"
runner_version: dev+90c7227.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a3044769b7b891b8cba111c1407e2b1d851e0e4546870058ac818bf45c649d2a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:alternatives -->
## Alternatives

**`docker pull`, `docker tag`, `docker push` instead of `imagetools`.** The three commands
every runner has, and they cannot do this job. A pull resolves the index to the one
manifest matching the runner's platform, and the push writes that manifest back, so the
mirror would hold a single-platform image with a digest of its own. The pin would then be
two pins that have to be kept equal by hand, which is the thing the decision on #355 was
taken to avoid. Rejected because it loses the property the design rests on, not because it
is more work.

**A copying tool pinned by this repository — `crane`, `regctl`, `skopeo`.** Each copies a
manifest index verbatim and each would be a row on `docs/supply-chain.md`, a version to
bump and a download on every release. `docker buildx` is a plugin of the docker CLI the
runner image already ships, which the page's last row already says is pinned by somebody
else. A new tool to pin for a copy that happens once per base bump is not a trade this
pipeline needs, and "adding a second dependency is a decision, not a step" is the
repository's rule for the binary and reads the same here.

**Mirroring by tag and building from the tag.** `--build-arg BASE=$IMAGE/base:13-slim`
reads more clearly and is wrong: that tag is written by this workflow, so a pin the
workflow can repoint is not a pin. The tag exists on the mirror because
`imagetools create` needs a destination to write, and the build names the digest.

**Widening all five publishing conditions uniformly.** The simpler-looking edit is to add
`|| github.event_name == 'workflow_dispatch'` to each of the five steps that carry the
`new_release_published` condition. It was rejected because it inverts where the care has
to go: a dispatch would then reach `build`, the module set, the bill of materials and
`checksums and upload`, and each would need something else to stop it publishing a second
set of binaries over a released version's assets. Gating the `release` step itself on
`github.event_name == 'push'` leaves the other four conditions untouched and false by
construction, because a step that did not run produces no output. One line instead of
four, and the four that matter are the ones not edited.

**Reading the base's digest from a literal in the workflow.** `BASE_DIGEST: sha256:a292…`
as a job-level `env` is the shortest path to a mirror reference and makes two copies of one
pin. `internal/model/supply_chain_test.go` would then ask the page to carry it as a second
row, which is that test working rather than failing, and the two copies would be correct
on the day they were written. The release reads `ARG BASE` out of the `Dockerfile` at run
time instead, which is also what keeps criterion 4 cheap: the file nobody edits is the
file everybody reads.

**Rebuilding the binary on a dispatch.** `go build` at the dispatched version would give
one release two binaries whose only difference is the tree they came from, and the newer
one would not be covered by the `SHA256SUMS` that release published. That is the asymmetry
#356 was written about, so reproducing it inside the cure would be perverse. The dispatch
downloads that release's own asset and checks it against that release's own `SHA256SUMS`.

**A mirror workflow triggered by the pin changing.** Rejected on #355 and restated here
because the design makes the reason concrete: both workflows fire on the same push to
`main`, so a release cut from that push and the mirror job filling the cache for it run at
the same time. Whether the release wins is a question about scheduling, which means the
failure appears on some bumps and not others and reproduces on none. What is accepted
instead is one release per bump still reaching Docker Hub.

**A `workflow_dispatch` with no input, publishing the latest release's image.** It removes
the chance of typing the wrong version, and it removes the only use the trigger has: the
release that is missing its image is not always the latest, and by the time somebody
dispatches it may not be. The input is required, and criterion 8 is what makes a wrong
one harmless.

<!-- xeno:section:impact -->
## Impact

Two files change. `docs/supply-chain.md` first and in its own commit, then
`.github/workflows/release.yml`. Nothing else, and the `Dockerfile` least of all.

**`docs/supply-chain.md`.** The base image's row keeps the digest it is pinned to and its
third column becomes `ghcr.io`, with `docker.io` named as the origin and as what the cache
is filled from. The paragraph "What a published digest promises, and what it does not"
gains the fetch path, because a row is where a reader looks and a paragraph is where the
reason lives. The row for `docker` says it builds and pushes the image; it now also copies
one, and that is the same unpinned tool rather than a new one. The test over the page reads
the first two columns and not the third, so none of this is a test change — which is why it
can be and has to be the documents commit the standing rule asks for.

**`.github/workflows/release.yml`.** Four edits, and the order of them is the design.

*The trigger.* `workflow_dispatch` with a required `version` input, described as the
version of a release whose image is missing, without the leading `v`.

*Where the image goes becomes its own step.* `IMAGE` is computed from
`ghcr.io/${GITHUB_REPOSITORY,,}` into `GITHUB_ENV` by an unconditional step after the
checkout, and the `build` step reads it instead of computing it. This is criterion 10, and
it is the one edit that is not about either issue: today the value is produced by a step a
dispatch skips, so the image step would have resolved its registry to the empty string.
Doing it this way keeps one copy of the value rather than giving each trigger its own,
which is the property the existing comment claims for it — "the registry the binaries claim
and the registry the image is in cannot disagree."

*The `release` step is gated on `push`.* One line, and it is what lets the four
binary-channel conditions stay exactly as they are: a step that did not run has no
outputs, so `new_release_published` is empty on a dispatch and the four are false by
construction. It also stops the other thing a dispatch could do by accident, which is cut
a version from whatever has landed on `main` since the last release.

*The image step.* Its condition admits the dispatch, its `VERSION` falls back to the
input, and its script gains three things in this order: on a dispatch, the release's own
`linux-amd64` asset downloaded and checked against that release's own `SHA256SUMS`; the
base read out of `ARG BASE`; and the mirror asked for the pinned digest and filled from
Docker Hub when it does not hold it. Then the build, with `--build-arg BASE=` naming the
mirror by that same digest.

**The mirror is asked and proved by one command.** `docker buildx imagetools inspect
"$IMAGE/base@$digest"` is both the hit check and the proof that the copy preserved the
digest: it is run before the copy to decide whether one is needed, and again after the copy
so that a copy which re-serialised the index fails there, with the reference it could not
resolve, rather than three lines later as a build that cannot find its base.

That matters because the property is somebody else's and is not measured here. That
`imagetools create` with one digest-pinned source and no annotations writes the index
through unmodified is taken from its documentation; the local docker has no buildx plugin
and no daemon, so nothing in this intent could run it. The design is arranged so that it
does not rest on the claim: criterion 2 is enforced by the pipeline on every miss rather
than measured once and assumed after, and a design that rests on a property of another
tool should fail on the day the property stops holding.

**A defect in the dispatch path, found by running it.** The download was simulated against
v0.60.2 before this design was written, and the asset arrives as mode 0644. A release asset
carries no permissions, and `gh release download` writes one with the default umask; the
`go build` in the build loop produces 0755. `COPY` preserves the source mode, so a dispatch
that only downloaded would publish an image whose `/usr/local/bin/xeno` is not executable,
and every gate job in it would fail on permission rather than on anything a reader would
connect to a release workflow. So the dispatch path chmods what it downloaded, and the
reason is here rather than in a comment on the line.

**The second thing taken rather than measured.** A container package that
`secrets.GITHUB_TOKEN` creates at `ghcr.io/triplem/xeno/base` from this workflow is linked
to this repository and inherits its permissions, so a later run may read it back with the
same token and no grant from anybody. If that is false the first release after the merge
fills the mirror and the next cannot see it, which arrives as a miss on every release: the
failure mode is a copy each time rather than a broken release, and the repair is one access
setting, which is the maintainer's and not this workflow's. It is written here because the
whole case for the mirror is that it needs one credential, and that is the sentence which
would be wrong.

**What is reached on a release after this.** `ghcr.io` for the base and for the push, and
Docker Hub on the first release after a base bump and on no other. `deb.debian.org` as
before, through `apt-get` inside the build, which mirroring the base does not move.

**What a dispatch costs.** It runs the gates at the top of the job against `main` as it
stands, not against the tree the dispatched tag was cut from. A dispatch for an old tag
therefore stops if `main` does not build, which is the right way round, and it is a cost
rather than a defect. It shares `concurrency: group: release` with pushes, so a dispatch
and a release queue rather than racing for the same registry path.

**What nothing else notices.** `audit.yml` reads `semantic_version:` and the action's sha
out of `release.yml` by `sed`, both above the first publishing step and neither moved.
`internal/model/version_test.go` asserts the two `ldflags` the build stamps and the path
the `Dockerfile` copies; the build keeps both and the `Dockerfile` is untouched.
`.dockerignore` is `*` then `!dist`, so the dispatch keeps `SHA256SUMS` out of `dist/` and
leaves only the binary the `COPY` names.
