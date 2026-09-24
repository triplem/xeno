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
| `cycjimmy/semantic-release-action` | `b12c8f6015dc215fe37bc154d4ad456dd3833c90`, v6.0.0 | github.com |
| semantic-release | 24.2.9 | npm |
| `@semantic-release/changelog` | 7.0.0 | npm |
| `@semantic-release/git` | 11.0.1 | npm |
| `CycloneDX/gh-gomod-generate-sbom` | `efc74245d6802c8cefd925620515442756c70d8f`, v2.0.0 | github.com |
| `cyclonedx-gomod` | v1.12.0 | github.com |
| `aquasecurity/trivy-action` | `ed142fd0673e97e23eac54620cfb913e5ce36c25`, v0.36.0 | github.com |
| `trivy` | v0.74.0 | github.com |
| trivy's vulnerability database | not pinned, and cannot be | ghcr.io |

**Where a version in this table comes from.** It is read off a run that produced a
release, from that run's log and from the bill of materials it published, rather than
taken as whichever is newest on the day somebody looks. What is wanted is the set that
demonstrably works, and a release is the only thing that demonstrates it.

The evidence travels with each release rather than living here. The bill of materials
records its own generator together with that generator's hashes, and the run log names
the version of semantic-release that ran. This table is therefore a convenience and the
release is the record; where the two disagree, the release is right.

**The vulnerability database is the one row that cannot be pinned.** A scan answers
what is known today, so a database fixed at a version would answer what was known when
somebody fixed it, which is the opposite of the question. It is fetched on every run.

## What this does not yet answer

The plan targets a self managed GitLab whose runners may have no route to the public
internet:

> Every step has to work from what the instance holds, which rules out fetching
> toolchains from the public internet at build time and makes the dependency mirror part
> of the bootstrap rather than an afterthought.

Every row of the second table above is such a fetch, and one of them, the vulnerability
database, is fetched on every run rather than pinned. Trivy supports an air gapped form
of it, which is the shape that question takes here. For each of them the instance needs
an answer — mirrored, pre-installed on the runner image, or dropped — and none of those
can be decided without the instance. Neither can the module proxy question for a
repository that vendors everything and has no `go.sum`.

That half stays open in issue #6, and A18 and A21 carry it from the other side.
