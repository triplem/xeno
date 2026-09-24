# What the build pulls in

Two lists. What ends up in a released binary, and what the pipeline runs to produce it.
The second is the longer one, which is the point of writing it down.

## In the binary

One module, vendored under `vendor/`, so a build fetches nothing:

| Module | Version | Licence |
|---|---|---|
| `gopkg.in/yaml.v3` | v3.0.1 | Apache-2.0, MIT |

Its own `go.mod` declares no `go` directive, which is why `vendor/modules.txt` records
it as `## explicit` with no version annotation. That line describes the dependency and
not this module, so it does not move when the `go` directive here does.

Each release carries a CycloneDX bill of materials generated from the shipped binary
rather than from `go.mod`, so it records what the artifact contains and not what the
module declares. It also records its own generator with that generator's hashes.

## In the pipeline

Every entry is pinned to something that cannot move: a commit sha for an action, an
exact version for a tool. A tag can be repointed by whoever owns it, and
`cycjimmy/semantic-release-action@v4` is a *branch*, which moves by design.

| What | Pinned to | Fetched from |
|---|---|---|
| `actions/checkout` | `11d5960a326750d5838078e36cf38b85af677262`, v4.4.0 | github.com |
| `actions/setup-go` | `40f1582b2485089dde7abd97c1529aa768e1baff`, v5.6.0 | github.com |
| Go toolchain | 1.27, from the `go` directive in `go.mod` via `setup-go`, which resolves it to the newest 1.27.x | golang.org |
| `cycjimmy/semantic-release-action` | `16ca923e6ccbb50770c415a0ccd43709a8c5f7a4`, v4.2.2 | github.com |
| semantic-release | 24.2.9 | npm |
| `@semantic-release/changelog` | 7.0.0 | npm |
| `@semantic-release/git` | 11.0.1 | npm |
| `CycloneDX/gh-gomod-generate-sbom` | `efc74245d6802c8cefd925620515442756c70d8f`, v2.0.0 | github.com |
| `cyclonedx-gomod` | v1.12.0 | github.com |

**The versions are the ones that produced v0.4.0**, read from that run's log and from the
bill of materials it published, rather than whichever were newest on the day this was
written. What is wanted is the set that demonstrably works.

Three of the four actions are behind their current major, and two carry a Node 20
deprecation warning. Upgrading is its own change: a run that both pins and upgrades
cannot say which of the two broke it.

## What this does not yet answer

The plan targets a self managed GitLab whose runners may have no route to the public
internet:

> Every step has to work from what the instance holds, which rules out fetching
> toolchains from the public internet at build time and makes the dependency mirror part
> of the bootstrap rather than an afterthought.

Every row of the second table above is such a fetch. For each of them the instance needs
an answer — mirrored, pre-installed on the runner image, or dropped — and none of those
can be decided without the instance. Neither can the module proxy question for a
repository that vendors everything and has no `go.sum`.

That half stays open in issue #6, and A18 and A21 carry it from the other side.
