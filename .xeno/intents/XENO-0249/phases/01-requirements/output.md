---
intent: github.com/triplem/xeno#201
phase: 01-requirements
created: "2026-10-05T12:48:15Z"
schema_version: "1.0"
runner_version: dev+f2f65d5.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f2099f96549ae6cf3d8b7441128f6fbfe8a4f53d621d9afa9baae057504a3f2b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. `plugin.HashTree(dir)` hashes the tree rooted at `dir`, each line a sum and the file's
   path relative to `dir`, sorted in byte order, and returns the empty string where `dir` is
   not there — which is `Hash`'s existing contract for an absent plugin.

2. `plugin.Hash(root)` is `HashTree` applied to `root/.xeno/plugin` and keeps its signature,
   so G-Supply, `xeno init` and the script are unchanged at their call sites.

3. The same tree at two paths hashes the same, and a tree differing by one byte does not.
   This is the property the comparison rests on and the one a test must assert directly.

4. `plugin.Differences(a, b)` names what differs between two trees: a path present in one
   and not the other, and a path whose contents differ. Sorted, and empty when the trees
   agree.

5. `scripts/plugin-digest.go --require-shipped` refuses, with exit 1, when the embedded tree
   is absent or hashes differently from the vendored one, and prints the differing paths
   rather than two hex strings alone. Without the flag it prints the digest as before and
   says on standard error that nothing was verified.

6. `.github/workflows/release.yml` passes `--require-shipped`, between the copy and the
   build, so the digest the release embeds cannot come from a tree the binaries do not carry.

7. The digest's value changes, by exactly the dropped `.xeno/plugin/` prefix, and nothing
   compares an old value against a new one: no test pins a literal, and a binary carries the
   digest taken at its own release together with the code that recomputes it.

8. A register row records the basis change, because the next person to wonder why a digest
   from before this release does not match will not find the answer in the diff.

9. `go build`, `go test ./...`, `go vet ./...` pass, `gofmt -l` outside `vendor/` prints
   nothing, and `./xeno gate verify` exits 0 with the 367 verdicts that exist now intact.

10. The workflow's change is checked by reading rather than by running, and the verification
    phase says so: a release cannot be rehearsed here, and claiming otherwise would be the
    defect this intent is about.

11. One commit, `Closes #201`, and the issue carries `wp7`.

<!-- xeno:section:non-goals -->
## Non goals

Not a change to G-Supply. The gate is right and this intent does not touch it. What changes
is that the digest it compares against can no longer be taken over a tree the binaries do not
carry.

Not a check that a published binary matches its tag. The issue names it as the other thing
nothing confirms, and location-independent hashing is what it would need; doing it needs a
released artifact to fetch and belongs with #198's verification phase.

Not a replacement for `cp -R`. A copy that cannot produce a different tree is the other way
to close this and is rejected in the design: the check is cheap, general, and catches a
refactor that a cleverer copy would not.

Not a digest stored in the repository. Section 13 is explicit that nothing in the repository
states the expected value, and a tree hash written beside the tree would prove only that both
were written by the same hand.

Not a new definition of the line format. The format stays a sum, two spaces and a path; only
the base the path is relative to changes, so Appendix B's shape is untouched.

Not a migration of any digest. No value anywhere needs rewriting: the only digests that exist
are inside released binaries, each with the code that recomputes it, and nothing in this
repository records one.

Not an unconditional comparison. A development tree carries no embedded copy by design, which
`embedded.go` explains, so a script that always required one would fail in the tree where it
is most often run by hand.

<!-- xeno:section:constraints -->
## Constraints

The release cannot be rehearsed. Nothing in this repository runs `release.yml`, so the
workflow half of the change is verified by reading it and by the script's own behaviour under
both flags, and the verification phase has to say that rather than imply a run. Claiming a
release was tested would be the same shape of defect as the one being fixed.

The hash basis change is one-way and silent. A digest computed before this commit and one
computed after differ for identical bytes, and nothing will report the difference, because
nothing compares across versions. That is safe and it is also exactly the kind of thing that
reads as a bug to whoever meets it first, which is why criterion 8 wants a register row rather
than leaving it in a commit message.

`HashTree`'s empty-string contract has to survive. `Hash` returns `""` for an absent plugin
and G-Supply reads that as a tree it cannot judge; a refactor that returned a hash of nothing
instead would make an absent plugin compare unequal rather than unjudged, which is a verdict
where there was none.

Normalisation stays where it is. `hashing.Normalise` is applied per file today and must go on
being, because the digest has to agree across a Windows checkout and a Linux runner; the
comparison this intent adds inherits that and must not compare raw bytes instead.

The script is `//go:build ignore` and run with `go run`. It can import `internal/plugin`
because it is inside the module, and it has no test of its own for the same reason the other
scripts have none; what is testable is the package function it calls, which is where the
logic goes.

One intent, one branch, `Closes #201`, and the issue carries `wp7`.
