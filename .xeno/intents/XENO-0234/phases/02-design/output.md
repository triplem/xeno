---
intent: github.com/triplem/xeno#183
phase: 02-design
created: "2026-10-03T17:31:49Z"
schema_version: "1.0"
runner_version: dev+9140756.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fb9fb1fca4ca5bb218239f24876d9fb541242b14b323efba3044d0ce400dc90f
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

**`plugin.ExpectedDigest` is a package variable the release sets through ldflags**, empty by default.
It sits beside `Hash`, which computes the value it is compared with, so the two halves of the
arrangement are in one file and a reader of either finds the other.

**`supply(c Ctx)` is three branches and no fourth.** No anchor is `notImplemented(c)`; no plugin is a
finding; a difference is a finding; otherwise `result(nil)`. Nothing reads a version, a manifest or
the lock.

**The no-anchor branch returns `notImplemented` rather than a result of its own.** It is the same
state for the same reason — a check this runner did not perform — and inventing a fourth result value
would be a schema change for a case section 5 already covers.

**The finding names both digests.** The gate holds the expected and the found value at the moment it
fails, and a reader who has to recompute one of them to understand the finding is being asked to do
the gate's work.

**The finding's file is `plugin.Dir`**, the directory, because that is what was judged. Other gates
name a file and this one has no single file to name; naming the tree is more use than naming the
first file in it.

**The release computes the digest with `go run scripts/plugin-digest.go`.** The same code the gate
compares with, so the two cannot disagree about the definition — and that is the one way this
arrangement fails silently, because a digest computed differently is a red verdict on every phase of
every project that installs the release. The `//go:build ignore` pattern is the one
`scripts/go-symbols.go` already uses.

**The script fails loudly where there is no plugin** rather than printing an empty digest, which
would compile in an empty anchor and turn the gate off for the whole release.

**`internal/plugin`'s tests become an external test package.** The gate imports the plugin package
and those tests import the gate, so an internal test is a cycle. External is what the language
provides for it, and the alternative — the gate duplicating the hash computation to avoid the
import — is the one thing the release's single-implementation argument forbids.

**The two status paragraphs in `ASSUMPTIONS.md` are corrected rather than appended to.** One listed
G-Supply as written `not-implemented` and the other as waiting on a harness, and both were true
until this change. A reader who finds a stale list stops trusting the others.

<!-- xeno:section:alternatives -->
## Alternatives

**Comparing against the lock's `plugin.sha256`.** Refused, and it is the obvious implementation now
that the field exists: the lock is written by the same runner into the same repository, so it is "an
expected hash stored beside the thing it describes", which section 13 says proves only that both were
written by the same hand. It would also compare a phase's record of what it was given against the
tree, which is G-Freshness's question about a different file.

**Comparing `plugin_version` against the runner's version.** Refused on section 13's own sentence
about the downgrade: the digest settles it "without any separate version check". A version comparison
is also weaker — a plugin can declare any version — and it is what #177 just removed from the field.

**Reading the expected digest from a file the release writes into the repository.** Refused for the
same reason as the lock, and it would reintroduce the problem #121 created: a release that writes a
file into main cannot, because a push with `GITHUB_TOKEN` reports no required check.

**A `pass` where the binary carries no anchor.** Refused: it is the sentence `not-implemented` exists
for, "a gate that is skipped silently makes a green verdict mean less than it appears to". It would
also make this repository's own trail claim a check that never ran.

**A `fail` where the binary carries no anchor.** Refused: it is not a failure of the repository, it
is a property of the build, and every phase of every development build going red would make the gate
something to switch off rather than something to read.

**A fourth result value for "ran but had nothing to compare".** Refused as a schema change for a case
section 5 covers, and the gate list is a budget.

**Duplicating the hash computation inside `internal/gates` to avoid the import.** Refused: two
implementations of one definition is precisely what makes a release's digest disagree with the
checker's, which is the failure the single implementation exists to prevent.

**Making the gate skip where the tree is a development checkout.** Refused — it would be a gate
branching on something about its environment, and the whole arrangement depends on the anchor being
the only thing that varies.

<!-- xeno:section:impact -->
## Impact

**`internal/plugin/plugin.go`.** `ExpectedDigest`, a variable with the section 13 quotation as its
reason and the no-anchor case stated.

**`internal/gates/gates.go`.** The table entry becomes `supply`. The function, with the three
branches and a doc comment carrying what the gate must not read and why. An import of
`internal/plugin`.

**`internal/plugin/plugin_test.go`.** `package plugin_test`, with the cycle as the stated reason.

**`scripts/plugin-digest.go`, new.** `//go:build ignore`, imports `internal/plugin`, prints the
digest, exits non-zero where there is no plugin.

**`.github/workflows/release.yml`.** `DIGEST=$(go run scripts/plugin-digest.go)`, echoed into the log
so a release says what it anchored to, and a second `-X` on the ldflags line.

**Tests.** Seven in `internal/gates`: the no-anchor case, the match, a one-byte change asserting both
digests in the finding, an earlier plugin, no plugin at all, the anchor not being readable from the
repository, and the gate's position in `Applicable`. Two helpers — one that sets and restores the
linker variable, because there is no `t.Setenv` for one, and one that writes a plugin tree and
returns its digest, so a test anchors to what it wrote rather than to a literal that would pin this
repository's tree.

**`ASSUMPTIONS.md`.** A86. Both status paragraphs corrected: G-Supply leaves the `not-implemented`
list, and the "Not built" paragraph stops saying it waits on a harness. G-Rules and G-Policy were
also stale in the first list and are corrected with it.

**`README.md`.** Why a reader of this trail meets `not-implemented` and what a released binary
reports instead.

**No document change, no change to any other gate, no change to the derived status.**
