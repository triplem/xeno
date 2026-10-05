---
intent: github.com/triplem/xeno#201
phase: 02-design
created: "2026-10-05T12:48:57Z"
schema_version: "1.0"
runner_version: dev+f2f65d5.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b377d6a2765e66ff3555636654789a1798ffd07ec543c6f1cab6724197eff685
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

`HashTree(dir)` is the primitive and `Hash(root)` becomes a two-line wrapper over it. The
walk, the line format, the sort and the normalisation move unchanged; only the base passed to
`filepath.Rel` differs, and it becomes `dir` instead of `root`. One definition, two entry
points, which is what keeps the release and the gate from disagreeing — the property
`plugin-digest.go`'s own comment already claims and could not deliver.

The digest's basis changes rather than a second scheme being added. A location-independent
hash beside a location-dependent one would be two definitions of the same thing, and the next
reader would have to work out which the gate uses. A86's arrangement is unaffected because
`Hash`'s signature and contract are unchanged; what changes is the number it returns, and
A96 records that.

`Differences(a, b)` is a separate function from the hash, returning sorted strings. It is not
part of the comparison — the hashes decide equality — and exists only so a refusal can be
acted on. Keeping it separate means the equality check stays the gate's definition exactly,
and the diagnosis cannot change the verdict.

The comparison lives in the script and not in the package. `HashTree` and `Differences` are
the package's; deciding that a release may not proceed is the release's. A package function
that refused would put a policy about workflows inside the code the gate calls.

`--require-shipped` is opt-in, because a development tree carries no embedded copy and
`embedded.go` says why. Without the flag the script prints the digest and says on standard
error that nothing was verified, so a person running it by hand is told what they did not get
rather than left to assume.

The flag is passed in the workflow between the copy and the build, two lines below the `cp`.
A guard adjacent to the thing it guards is one a refactor moves with; a guard in a separate
step further down is one a refactor leaves behind.

The embedded tree's path is a constant beside `Dir` rather than a literal in the script.
`Dir` names the vendored tree and the new constant names the copy, so the two paths the bug
is about are declared together in the package that defines the hash over them.

The register row is A96 and it goes in the register rather than only in this intent's P2, by
the rule #243 established: why a digest from before this release does not match is a fact
every later reader of a release needs, so it outlives the intent that changed it.

<!-- xeno:section:alternatives -->
## Alternatives

`diff -r .xeno/plugin internal/plugin/embedded/plugin` was the issue's first form: three
lines, no code. Rejected, and the issue itself prefers the other. `diff` compares bytes, and
the hash compares normalised bytes, so a file differing only in line endings would pass the
gate and fail the `diff`, or the reverse on a different checkout — a guard that disagrees with
the thing it guards. It would also leave the hash location-dependent, so the stronger check
this enables, recomputing a published binary's digest, would still be out of reach.

Keeping the hash as it is and comparing the two trees file by file was considered. It is the
`diff` objection in Go: it works, and it is a second definition of what makes two trees equal,
which is the thing this intent exists to remove.

Making the comparison unconditional was weighed against the flag. It is stronger in the
release and wrong everywhere else: a development tree has no embedded copy by design, so the
script would fail in the tree where a person is most likely to run it, and a script that fails
normally is one whose failure stops meaning anything. The flag is the cost of that, and the
cost is real: a refactor can drop it. Placing it two lines below the `cp` is the mitigation,
and the verification phase names it as the gap rather than claiming it is closed.

Hashing the embedded tree through `embed.FS` rather than from disk was considered, so that the
check would read what the binary actually carries rather than what is about to be embedded.
Rejected: the embedding happens at compile time, so at the moment the script runs there is
nothing embedded yet, and a check after the build would need a built binary to introspect
itself. That is the published-binary check, which is #198's and needs a release to fetch.

Replacing `cp -R` with `go run` over the tree, or with `git archive`, was considered. Any of
them narrows the ways a copy can go wrong without checking that it did not, and the issue's
own framing is that `cp -R` already fails loudly on the ordinary error; what is unguarded is
the copy that succeeds.

Leaving the digest basis alone and accepting that the two trees cannot be compared was the
null option, and the reason to reject it is the blast radius rather than the likelihood: the
failure is improbable, silent at release, and lands on every phase of every adopter of that
release with nothing they can do.

<!-- xeno:section:impact -->
## Impact

Four files. `internal/plugin/plugin.go` gains `HashTree`, `Differences` and a constant, and
`Hash` becomes a wrapper. `internal/plugin/plugin_test.go` gains the location-independence
tests. `scripts/plugin-digest.go` gains the flag and the refusal.
`.github/workflows/release.yml` gains the flag on one line. `docs/assumptions.md` gains A96.

No call site changes. `Hash(root)` keeps its signature and its empty-string contract, so
G-Supply at `internal/gates/gates.go:280`, `xeno init --vendor` and the script's default path
are untouched. Nothing in the repository records a digest, so nothing needs rewriting.

The digest's value changes for identical bytes, by the dropped `.xeno/plugin/` prefix on
every line. Nothing compares across the change: a released binary carries the digest taken at
its own release and the code that recomputes it, so each release is internally consistent, and
no test pins a literal — which was checked rather than assumed. A96 is there because this is
nonetheless the thing most likely to be mistaken for a bug by whoever meets it first.

What the next release gains is that the anchor cannot be produced from the wrong tree. The
digest is printed only after the embedded copy has been hashed with the same code and found
equal, so the failure mode the issue describes becomes a release that refuses to build rather
than a release that ships and breaks every adopter.

What this also unlocks is the check nothing has: a location-independent hash can be recomputed
over a tree extracted from a published artifact, which is what confirming a binary against its
tag needs. This intent does not do it and the scope says so, but it stops being blocked.

The honest limit is that the guard is opt-in and the workflow is unrehearsed. A refactor that
moves the copy and drops the flag restores the original defect silently, and nothing in this
repository can run `release.yml` to notice. The mitigation is adjacency and a comment; the gap
is real and the verification phase states it rather than counting criterion 6 as covered by a
test.

The second limit is that this checks the copy and not the embedding. Between the script's
comparison and the binary there is still a `go build` reading `//go:embed all:embedded`, and
nothing confirms that what was embedded is what was on disk. That gap is narrower than the one
being closed — the embed directive reads the same path the copy wrote — and it is the
published-binary check that would close it.
