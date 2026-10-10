# What the build pulls in

Two lists. What ends up in a released binary, and what the pipeline runs to produce it.
The second is the longer one, which is the point of writing it down.

## In the binary

One module, vendored under `vendor/`, so a build fetches nothing:

| Module | Version | Licence |
|---|---|---|
| `go.yaml.in/yaml/v3` | v3.0.5 | MIT, Apache-2.0 |

`vendor/modules.txt` records it as `## explicit; go 1.16`. That annotation is the `go`
directive of the dependency, not of this module, so it does not move when the directive
in `go.mod` here does.

It is load bearing rather than cosmetic. Go compiles each module under its own declared
language version, so this package's files are built with 1.16 semantics and ours with
1.27, in one build and one binary. For this package that means no generics and the loop
variable behaviour from before 1.22, which is what its code was written against. A
dependency declaring a version below ours is never a constraint on us; one declaring a
version above ours would be, and could not be built at all.

What it does not say is whether the module is patched. That follows from its version and
from somebody maintaining it, and the YAML organisation took this package over after
go-yaml was marked unmaintained in April 2025, which is why it is the one vendored here.

`go.sum` carries its checksums. A vendored build does not consult it, so it is a record
rather than a gate, and `go mod verify` is what reads it.

Each release carries a CycloneDX bill of materials generated from the shipped binary
rather than from `go.mod`, so it records what the artifact contains and not what the
module declares. It also records its own generator with that generator's hashes.

One document covers all five binaries, and that is a checked condition rather than an
observation. Build constraints can make a target select different modules, and the
document describes modules, so before it is generated the release compares the module
set of every target it builds and stops where two of them disagree. Package sets are not
compared: they differ on every platform by way of the standard library and say nothing
about the document. The day a platform dependent dependency enters the tree, the release
fails and the repair is one document per target.

## In the pipeline

Every entry is pinned to something that cannot move: a commit sha for an action, an
exact version for a tool. A tag can be repointed by whoever owns it, and some publishers
keep moving major refs besides: `cycjimmy/semantic-release-action` carries `v1` through
`v6` as *branches*, which move by design and look exactly like tags in a `uses:` line.

| What | Pinned to | Fetched from |
|---|---|---|
| `actions/checkout` | `3d3c42e5aac5ba805825da76410c181273ba90b1`, v7.0.1 | github.com |
| `actions/setup-go` | `b7ad1dad31e06c5925ef5d2fc7ad053ef454303e`, v7.0.0 | github.com |
| Go toolchain | 1.27, from the `go` directive in `go.mod` via `setup-go`, which resolves it to the newest 1.27.x | golang.org |
| `actions/setup-node` | `820762786026740c76f36085b0efc47a31fe5020`, v7.0.0 | github.com |
| Node toolchain | 24, from `node-version` in `audit.yml` and `release.yml` via `setup-node`, which resolves it to the newest 24.x | github.com, nodejs.org |
| `actions/setup-python` | `5fda3b95a4ea91299a34e894583c3862153e4b97`, v7.0.0 | github.com |
| Python toolchain | 3.14, from `python-version` in `docs.yml` via `setup-python`, which resolves it to the newest 3.14.x | github.com |
| `zensical` | 0.0.69, through `pip install`; its fourteen dependencies resolve unpinned at install | pypi.org |
| `actions/configure-pages` | `45bfe0192ca1faeb007ade9deae92b16b8254a0d`, v6.0.0 | github.com |
| `actions/upload-pages-artifact` | `fc324d3547104276b827a68afc52ff2a11cc49c9`, v5.0.0 | github.com |
| `actions/deploy-pages` | `368f82528645a54fb793d4d04e342629a3f51346`, v5.0.1 | github.com |
| `cycjimmy/semantic-release-action` | `b12c8f6015dc215fe37bc154d4ad456dd3833c90`, v6.0.0 | github.com |
| semantic-release | 25.0.9 | npm |
| `CycloneDX/gh-gomod-generate-sbom` | `efc74245d6802c8cefd925620515442756c70d8f`, v2.0.0 | github.com |
| `cyclonedx-gomod` | v1.12.0 | github.com |
| `actions/upload-artifact` | `043fb46d1a93c77aae656e7c1c64a875d1fc6a0a`, v7.0.1 | github.com |
| `renovatebot/github-action` | `230ce922b08968d0a4f6f70f295601daff09ef1d`, v46.3.7 | github.com |
| `aquasecurity/trivy-action` | `ed142fd0673e97e23eac54620cfb913e5ce36c25`, v0.36.0 | github.com |
| `trivy` | v0.74.0 | github.com |
| trivy's vulnerability database | not pinned, and cannot be | ghcr.io |
| `golangci/golangci-lint-action` | `ba0d7d2ec06a0ea1cb5fa41b2e4a3ab91d21278a`, v9.3.0 | github.com |
| `golangci-lint` | v2.14.0 | github.com |
| `gosec` | v2.29.1-0.20261009120814-7b1b5cebe007, the pseudo-version of a commit on its main branch, as `github.com/securego/gosec/v2` through `go install`, checksummed by the module sum database; a commit until 2.30.0 because release 2.29.0 cannot read Go 1.27.2's export data (#343) | proxy.golang.org |
| `govulncheck` | v1.8.0, as `golang.org/x/vuln` through `go install`, checksummed by the module sum database | proxy.golang.org |
| govulncheck's vulnerability database | not pinned, and cannot be | vuln.go.dev |
| gitleaks' rules | v8.30.1, translated into `.xeno/plugin/secrets.yaml`, not fetched | — |
| `gitleaks` | 8.30.1, by release tarball and sha256 `551f6fc8…` | github.com |
| the image's base, `debian` 13-slim | `sha256:a29215f6…`, the manifest index, by digest | ghcr.io, mirrored from docker.io |
| `ca-certificates` and `git`, installed into the image | not pinned, and cannot be | deb.debian.org |
| `docker`, which builds and pushes the image | the hosted runner's image, which this repository does not pin | — |

**Where a version in this table comes from.** It is read off a run that produced a
release, from that run's log and from the bill of materials it published, rather than
taken as whichever is newest on the day somebody looks. What is wanted is the set that
demonstrably works, and a release is the only thing that demonstrates it.

The evidence travels with each release rather than living here. The bill of materials
records its own generator together with that generator's hashes, and the run log names
the version of semantic-release that ran. This table is therefore a convenience and the
release is the record; where the two disagree, the release is right.

**What a published digest promises, and what it does not.** The image's base is pinned
by the digest of a manifest index rather than by `13-slim`, and the image a release
publishes is referred to by the digest the run prints, so neither the bytes the build
starts from nor the bytes a pipeline pulls can be changed under anybody afterwards. What
that leaves open is the two packages inside. `apt-get install ca-certificates git`
resolves against Debian's suite on the day the release is cut, and the suite moves;
pinning `git=1:2.47.3-1` would turn each of those moves into a failed release, which is
why the row sits in the vulnerability database's class and carries no version. So a
published image is fixed and two images a month apart are not the same image, and the
`git` in either of them is read off the build log of the run that produced it. The bill
of materials cannot answer it, because it describes the binary and not the image around
it.

**Where the base is fetched from, and where it comes from.** Those are two registries
and the row names both. The release builds from `ghcr.io`, which is the registry it
already authenticates to and already pushes the image to, and what it finds there is a
copy of the digest above, mirrored from `docker.io`. The copy is made by the release
itself, on the first run that asks for a digest the mirror does not hold, so Docker Hub
is reached once per bump of the pin and not once per release. That is the whole of the
reason: an unauthenticated pull from Docker Hub spends a quota shared with every other
puller on the runner's address, and on 2026-10-09 it ran out and cost v0.60.2 its
image (#355).

The pin does not move with the fetch path. `ARG BASE` in the `Dockerfile` names
`docker.io/library/debian:13-slim@sha256:a29215f6…` unchanged, the copy preserves the
manifest index, and the release builds against that same digest at its mirrored address.
So a reader of the `Dockerfile` and a reader of this row are told the same bytes, and
somebody with no read access to this registry can still build the file by hand from the
reference it carries. The `docker` of the last row copies as well as building and
pushing, with the same plugin of the same unpinned CLI.

**Zensical's dependencies are the one install here that resolves unpinned.** The
generator itself is pinned to a version, and `pip` resolves the fourteen packages behind
it to whatever each range allows on the day. A lockfile would pin them, and the day a
transitive package breaks the build is the day it earns its place; until then the row
says what it is, and the build is strict enough to say when it breaks. The site it
produces is a derivative of the repository and nothing a verdict rests on, which is why
this is tolerable here and would not be on the gate path.

**The `docker` CLI is pinned by the runner rather than by this repository.** The image
step runs `docker login`, `docker build` and `docker push` instead of
`docker/login-action`, `docker/setup-buildx-action` and `docker/build-push-action`. That
is a trade taken on purpose — three action shas off this table against one tool whose
version comes from the hosted runner image — and the honest form of it is a row saying
the pin is somebody else's. The cost is the one the staleness paragraph below describes:
the day that runner image moves the version, nothing here changes and nothing here
notices.

**The vulnerability database is the one row that cannot be pinned.** A scan answers
what is known today, so a database fixed at a version would answer what was known when
somebody fixed it, which is the opposite of the question. It is fetched on every run,
and it is built upstream on a 24 hour cycle, which the daily scan is aligned with.

It is also the one row where being cut off from the network is not the interesting
question. `--db-repository` takes a list of OCI repositories and the database is an OCI
artifact, so a runner without egress reads it from a mirror in a registry it can reach.
**Currency is then the mirror's sync cadence and not a property of the air gap**:
mirrored daily it is as fresh as fetching it directly, mirrored weekly it is up to seven
days blind. That is a number somebody chooses, and it belongs wherever the scan result
is read.

**What is watched, and what is not.** Trivy reads the binary a release ships. The
`audit` workflow reads the tree that produces it, and reading that tree means performing
somebody else's install: the release installs semantic-release through
`cycjimmy/semantic-release-action`, which runs `npm ci` against its own committed
lockfile and installs the pinned version on top of the result. So the audit checks the
action out at the sha pinned above and runs those steps, taking both identifiers out of
`release.yml` rather than keeping copies. It used to install the pinned package into a
bare `npm init -y` instead, and that is a different tree: measured on 2026-10-06 it
carried twelve advisories and no critical, against twenty-eight and two for the tree the
release installs. The difference is the lockfile — a fresh resolution takes the newest
each range allows, and the action's lockfile holds older versions, so the pin is what
keeps the vulnerable ones and the unpinned resolution is what hid them (#260). The
counts are recorded in `.github/npm-audit-baseline.json` and the job fails when one
rises, not when one is non-zero: none of them is this project's to fix, and the sixteen
that only the corrected tree shows can move only when the action's sha does.

Two things stay unwatched and are written here rather than assumed covered. **Staleness
is not watched at all** — nothing says whether a pinned action or tool is still the one
to be on, and a count that does not move is not evidence that it is. And the commit shas
of the actions themselves have no advisory feed here; Dependabot would give one, and was
rejected in #11 for a reason that has not changed: its pull requests are changes without
an intent.

**The linters' policy goes the other way, and the contrast is the point.**
`.golangci.yml` and `.gosec.json` say which linters run and which findings are accepted,
the tools that read them are pinned, and nothing fetches either at run time, so a run
reports what the previous run reported unless somebody changed one of them in a commit.
The patterns of the secret filter follow the same rule for the same reason: they are
translated from gitleaks' rules at a named version into `.xeno/plugin/secrets.yaml`, and
`secrets_hash` means what it says only because nothing fetches them. A vulnerability
database is facts about the world that other people discover, so it goes stale by time
passing. A rule set is patterns somebody chose to enforce, and one that changes
underneath a project can fail a build that nothing in the repository touched. Facts want
currency; policy wants a commit.

govulncheck's database is trivy's case and not the linter's: fetched on every run, from
`vuln.go.dev`, and matched against call paths rather than versions. Its report carries
the database's own date, `db_last_modified`, so it answers for its own coverage where
trivy's report has to be accompanied by the cache's metadata.

Two things follow from the database being data rather than code. The case for letting it
through a gap is a different case from the one for a toolchain, since nothing in it is
executed. And a scan is evidence in this process, which does not have to be produced on
a runner without egress: it can be produced where there is a route out and bound in by
`uri` and `sha256` like any other evidence item.

## What this does not yet answer

A hosted runner has a route out, so none of this blocks the repository today. It blocks
any project adopting Xeno on its own infrastructure, and it would block this one the day
it moves to a self hosted runner, which is why WP0 keeps the requirement:

> Every step has to work from what that machine already holds, which rules out fetching
> toolchains from the public internet at build time and makes the dependency mirror part
> of the bootstrap rather than an afterthought.

**Every row of the second table above is such a fetch**, bar the one the runner image
already carries, and each of them needs an answer for a machine without a route out:
mirrored, pre-installed on the runner image, or dropped. None of those can be decided
without such a machine in front of somebody. Two of the rows belong to the image rather
than to the build: its base, which this repository pulls when it builds, and the image
itself, which an adopter's pipeline pulls before every job. The second is the one the
generated wrapper already answers, on the image line whose comment says to point it at a
registry the pipeline can reach; the first is a row like any other.

The Go module question is narrower than it looks. A vendored build reads
`vendor/modules.txt` and not `go.sum`, so nothing in the release path fetches a module
at all; only `go mod tidy` and `go mod vendor` reach the proxy, and neither runs in CI.
What is fetched is the toolchain, which is a row of the table like any other.

The vulnerability database is the row that is fetched on every run rather than pinned,
and the paragraph above the table says what the question becomes for it: a mirror and
its cadence, not an air gap.

## One gap that is not about egress at all

**A Trivy report records when the scan ran and not when its database was built.** So a
report cannot be judged for its coverage from itself, whether or not the machine has a
route out. `UpdatedAt` sits in the cache's `metadata.json` and nowhere in the output.

Where the report becomes a declared `kind: scan` item, that value belongs beside
`produced_by` and `result`, or the trail records that a scan ran without recording what
it could have known.

The egress half stays open in issue #6. This one belongs with the evidence
declaration and is recorded in #11.
