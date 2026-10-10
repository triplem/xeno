---
intent: github.com/triplem/xeno#355
phase: 05-review
created: "2026-10-10T13:16:01Z"
schema_version: "1.0"
runner_version: dev+b9beef5.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ceae7fc2a6836fad9c9ed7e75b121b3991e3e70557103c585dee84153889144b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'One deviation, written in P3 with its cause. The design meant to widen the docker row of docs/supply-chain.md to say it copies as well as building and pushing; the exempt map of internal/model/supply_chain_test.go is keyed on that row''s exact first cell, so moving the words would have unexempted the row and made the documents commit a code change as well. The sentence went into the prose beside the fetch path instead. Everything else is the design: the four workflow edits in the order the impact section named them, and nothing in the Dockerfile. What was weighed and not taken is recorded as a gap in P4 rather than slipped in: a job that builds the image and runs xeno version in it would catch the mode, the base resolution and the COPY path at once, and it is work against the verify workflow and not against the release.'
      result: met
      rule: deviations-are-traceable
    - note: No interface change. A trigger is added to one workflow of this repository, a step is added to it, one step's condition is widened and another's is narrowed. No command, flag, artifact field, template or gate moves; the Dockerfile is byte for byte what it was, so a hand build takes the arguments it took before; and the generated CI wrapper is untouched, so nothing in an adopter's repository changes. The published image is at the same reference with the same tag scheme and is built from the same binary.
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: 'No new dependency. go.mod and vendor/ are untouched, no action is added, and nothing is installed. The copy uses docker buildx imagetools, a plugin of the docker CLI the hosted runner image already ships, which docs/supply-chain.md already records as pinned by somebody else -- which is why crane, regctl and skopeo were each rejected in P2: a copy that happens once per base bump does not earn a row, a version to bump and a download on every release. gh is the host CLI the checksums step already uses. The base is the same digest from a second address, not a second base.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
    - note: 'Supply chain lens: the release now builds from an address this repository controls, which is the point and is also a change in who can move the bytes. On docker.io a digest is immutable and the registry is a third party nobody here can write to; on ghcr.io the same digest is immutable and the registry is one this repository pushes to, so a credential that leaked could put different content at a different digest and a release would still fetch the pinned one. The pin is what carries the guarantee in both cases and it has not moved. What is genuinely new is that the cache is filled by a run rather than by a person: the first release after a bump copies whatever docker.io then serves for that digest, unreviewed, which is the same trust a direct pull placed in the digest and now places in it one step earlier. Named rather than resolved, because the alternative is a person re-mirroring by hand and the decision on #355 retired that question deliberately.'
      result: deviation
      source: lens
---

# Review

<!-- xeno:section:release-notes -->
## Release notes

**A release no longer depends on a quota it shares with strangers.** The image's base was
pulled from Docker Hub without authenticating, so every release spent an allowance shared
with every other unauthenticated puller on the runner's address. On 2026-10-09 that
allowance ran out: the build failed at `load metadata` with `429 Too Many Requests`, and
v0.60.2 was published with its seven assets and no image. The credential the step holds is
for `ghcr.io`, the push target, so the push was authenticated and the pull it depended on
was not.

The base is now fetched from `ghcr.io`, the one registry the release already authenticates
to and already pushes to. The image step asks whether that registry already holds the
pinned digest and copies it from Docker Hub when it does not, so Docker Hub is reached once
per bump of the pin rather than once per release, and the mirror is a cache rather than a
thing with an owner.

**The pin has not moved and the `Dockerfile` has not changed.** `ARG BASE` still names
`docker.io/library/debian:13-slim@sha256:a29215f6…`. The release reads that line and builds
against the same digest at its mirrored address, so a reader of the file and a reader of
`docs/supply-chain.md` are told the same bytes, the dependency bot still proposes newer
digests for the right image, and anybody with no read access to this registry can still
build the file by hand from the reference it carries. What it costs is that the file says
`docker.io` while a released image came from `ghcr.io`; the page's row and the workflow say
so.

**A release that loses its image can now get one.** Every publishing step was gated on
semantic-release cutting a version in the same run, and once it has cut one it will not cut
it again. The workflow had no other trigger, so a re-run restarted the job at the gate
rather than at the thing the gate guards and was green because semantic-release correctly
did nothing. `workflow_dispatch` with a version input now reaches the image step for a tag
that already exists.

It publishes that release's own bytes. The binary comes from the release's assets and is
checked against that release's own `SHA256SUMS` before it is copied in, rather than built
again from a tree that has moved on, so the image carries what `SHA256SUMS` already covers.
A version with no release stops before anything is pushed, and the dispatch cuts no tag,
publishes no release and touches no asset: only the image step admits it, because the
`release` step runs on a push alone and the binary channel's four conditions are therefore
false on a dispatch by construction.

**v0.60.2's image is still missing.** It is recovered by running this dispatch once for
`0.60.2`, which is deliberately the first use of the path rather than a hand-pushed image no
run log accounts for.

**`docs/supply-chain.md` records the change.** The base image's row keeps the digest it is
pinned to and names `ghcr.io` as what the release fetches it from, with `docker.io` as its
origin and as what the cache is filled from.

<!-- xeno:section:residual-risk -->
## Residual risk

**Three criteria are not met yet and cannot be before the merge.** That a release fetches
from `ghcr.io`, that the copy preserves the digest, and that the dispatch publishes an image
are properties of runs. The first release after this lands answers the first two on its
first miss; the dispatch for `0.60.2` answers the third and is the recovery #356 decided.
Until both have run, the change is checked as far as a tree can check it and no further.

**Nothing re-checks that the release fetches from the mirror.** An edit that dropped
`--build-arg BASE=` would leave a release pulling from Docker Hub again with every test in
this repository still passing: `internal/model/supply_chain_test.go` holds the page to the
pins and says nothing about which registry a build fetches from, and nothing else reads the
image step. The failure would return as the next spent quota — months of silence and then
the same 429. The same holds for the dispatch's `chmod`: dropped, it would publish an image
whose runner is not executable, and the failure would appear as a permission error in
somebody else's repository. P4 records both, with what would close them.

**The one release following a base bump still depends on Docker Hub.** A bump is a miss, so
that release copies from `docker.io` and a bump landing on a spent quota fails it — the
2026-10-09 failure, once per bump rather than every time, with the next run refilling the
cache. This is the accepted cost of the decision on #355 and not an oversight: a mirror
workflow on the pin changing would remove it and would race the release cut from the same
push.

**The mirror's readability is assumed.** A container package `GITHUB_TOKEN` creates is taken
to be linked to this repository and readable by the same token on a later run. If that is
false the mirror is filled and never hit, which shows as `docker.io` in every release log
rather than in one per bump, and the repair is one access setting, which is the maintainer's.
The failure mode is a copy per release and not a broken release, which is why it was
accepted as an assumption.

**A dispatch runs the gates against `main`, not against the dispatched tag's tree.** A
dispatch for an old version stops if `main` does not build. That is the right way round and
it means the recovery of an old release depends on the current tree being healthy.

**The context budget is exceeded, and the finding is honest.** G-Schema carries F-786b8c
from P4 onward: the recorded context is 80,490 bytes against the 73,000 the intake set. The
cause is the change itself — `release.yml` grew by 121 lines and the page by twenty — and
the budget sits inside the intake's `artifacts_hash`, so it cannot be moved now without
invalidating that verdict. Nothing about the scope was wrong; the figure was measured before
the work that enlarged the files it measured. The learning recorded against it is the
general form.

**What this change does not fix.** A green re-run of a release workflow is still
indistinguishable from a recovery: #356 reports that the re-run of `37991808351` is green
because semantic-release correctly found nothing to do. This intent gives the recovery its
own trigger instead of changing what a re-run reports, which leaves that as a separate
defect about run logs.
