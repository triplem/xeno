---
intent: github.com/triplem/xeno#201
phase: 03-implementation
created: "2026-10-05T12:54:46Z"
schema_version: "1.0"
runner_version: dev+325b27a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5d5810587b460276bb638ecc34590b2e1526c4cd9b7d6242062dda7114583ff1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`internal/plugin/plugin.go`. `Hash(root)` is now one line, `HashTree(filepath.Join(root, Dir))`,
and keeps its signature, its contract and its callers. `HashTree(dir)` is the computation with
each line's path relative to `dir`; `treeLines(dir)` is the walk, the per-file `Normalise` and
the sort, shared by the hash and the comparison so that neither can drift from the other.
`Differences(a, b)` returns sorted strings naming a path only in one tree and a path whose
contents differ. `ShippedDir` is a new constant beside `Dir`, because the two paths this is
about belong declared together.

`Hash`'s doc comment is corrected where it said the path is relative to the repository root,
and gains the paragraph saying the value changed and why nothing compares across the change.
The sentence was load bearing: Appendix B's reason for defining a hash to the byte is that a
hash described rather than defined is how two implementations end up one byte apart, and this
one is now described correctly.

`scripts/plugin-digest.go` takes `--require-shipped`. With it, the digest is printed only if
the tree at `ShippedDir` exists and hashes equal to the vendored one; without it the digest is
printed and standard error says nothing was verified. A refusal prints both digests and then
the differing paths.

`.github/workflows/release.yml` passes the flag, on the line that was already there, with the
comment two lines below the `cp` it guards.

`internal/plugin/hash_test.go`, new, seven tests in the existing external test package: the
same tree at two paths hashes the same; a tree differing in contents and a tree with an extra
file both hash differently; `Hash` is the tree at the vendored path; an absent tree hashes to
empty, for both entry points; `Differences` names contents-differ and only-in for both
directions; it is empty where the trees agree; and it says plainly which tree is not there
rather than reporting every file as missing.

The fixture tree carries a dot-prefixed directory on purpose. The manifest lives under
`.claude-plugin/`, and a walk that skipped dot names would hash a different tree than the one
shipped, which is the kind of thing `embed`'s own `all:` prefix exists to prevent.

The release sequence was rehearsed locally rather than reasoned about: copy the tree, run with
the flag, see the digest; add a stale file, see the refusal name `STALE.txt`; change one byte
of the manifest, see `contents differ: .claude-plugin/plugin.json`. The copy was then removed,
because `internal/plugin/embedded/plugin` is a release artifact and a committed one would make
a development build claim to be a release.

<!-- xeno:section:deviations -->
## Deviations from the design

One addition beyond the design, which named `HashTree` and `Differences` and did not name
`treeLines`. The walk is shared between them rather than written twice, because `Differences`
needs the same per-file `Normalise` the hash applies: a comparison over raw bytes would
disagree with the hash on a file differing only in line endings, which is the objection P2
raised against `diff -r` and would have been just as true of a hand-rolled comparison here.
Extracting the walk is what makes the two provably the same read.

One deviation from P3's own method rather than from the plan, and it is an improvement worth
recording. The three intents before this one wrote their code between P2's verdict and
`phase start --phase 03`, so P3's elapsed figure measured two sections being written and not
the work; the figures comments on #117 reported it three times as a defect in the measurement.
This phase was opened first, at 12:49:35Z, and the work done inside it. The elapsed time is
therefore the work for the first time in four intents, which is a change to how the phase is
run and not to what it produces.

No deviation on the digest's value, which P1 declared would change and P2 explained. It did:
the vendored tree now hashes to `1c9bd14a` where before the prefix on every line made it
something else. Nothing compared the two, which was checked rather than assumed — no test
pinned a literal and this repository records no digest.

One thing found and left, because it is adjacent and not this issue's. `internal/plugin/
embedded/plugin` is not in `.gitignore`, so a developer who rehearses the release, as this
phase did, is one `git add -A` away from committing a tree that would make
`plugin.Shipped()` return true for a development build — the exact claim `embedded.go` says a
development build must not make. One line would close it. It is a finding about the release's
ergonomics rather than about the digest, and absorbing it here would put an unrelated guard in
this commit.
