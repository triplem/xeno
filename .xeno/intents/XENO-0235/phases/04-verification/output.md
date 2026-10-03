---
intent: github.com/triplem/xeno#199
phase: 04-verification
created: "2026-10-03T18:24:51Z"
schema_version: "1.0"
runner_version: dev+a897f2a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3f648b87612c8857a89f076fbdb1a4f6534428bd7fc05bb21165be93dfa04a69
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Each acceptance criterion of P1 against what holds it.

**A released binary in an empty repository vendors the plugin** — measured, not tested: a binary
built with the release's flags, run with `git init` and nothing else, reported `plugin from this
release` and wrote 34 files. No Go test can hold this, because the embedded tree only exists in a
build the release made.

**The vendored tree's digest is the one the binary carries** — two things, because the property has
two halves. `TestTheVendoredTreeHasTheDigestTheSourceHas` holds the half a test can: the tree `init`
writes hashes the same as the tree it copied from. The other half — that the binary's compiled-in
value is that digest — is the release's ordering, held by the workflow and measured once by running
the release's steps and then `phase finish` against what they produced: `G-Supply pass`.

**The adopter's `plugin_version` is the release's number** — measured: `0.33.0` in the first
artifact, from a manifest stamped `0.33.0`.

**The digest is taken after the stamp** — held by the workflow's step order, with the reason written
between the steps. No test; the failure is in a build nobody here runs, which is why the comment is
where somebody editing the steps will read it.

**Every file is vendored, at any depth** — `TestVendorCopiesEveryFileOfThePlugin`, comparing two
sorted lists so a failure names what is missing, and `TestTheVendoredTreeHasTheDigestTheSourceHas`,
which would also fail. Two tests for one property on purpose: the list names the gap and the digest
catches a byte.

**The entry point is executable** — `TestTheVendoredEntryPointIsExecutable`.

**A development build with no source refuses, naming `--plugin-from`** —
`TestWithoutAPluginToVendorInitRefuses`, asserting both the refusal and the two phrases a reader
needs.

**The embedded copy wins over the directory** — read in `pluginSource`, and measured in the opposite
direction: the release binary ignored the repository it was run in and said `this release`.

**The repository's manifest is `0.0.0-dev` and the client accepts it** — `claude plugin validate`,
run against the tree.

**The test compares digests, not file lists** — both exist, and the digest one is the one that would
catch a changed byte.

**Nothing in this trail changes** — `gate verify` 285 at exit 0, and the three existing vendor tests
pass untouched, which is what says the walk lost nothing the list carried.

<!-- xeno:section:results -->
## Results

**`go test ./...`** — eighteen packages ok, five cases new, three existing vendor tests unchanged and
passing.

**`gofmt -l .` outside `vendor/`** and **`go vet ./...`** — nothing.

**`./xeno gate verify`** — `verified 285 verdicts`, exit 0, exit code captured.

**`claude plugin validate .xeno/plugin`** — `✔ Validation passed`, with the manifest at `0.0.0-dev`.
That was the one thing promised before starting and it is the reason the placeholder is a version
rather than an empty string.

**The release's own steps, run in a copy**, with `VERSION=0.33.0`:

    stamped: 0.33.0
    digest:  6b99de9c21e5aade8b8bf45d2805609ffa426988a926fcc366d9e26c2c81bf25
    xeno 0.33.0

**And that binary in an empty repository** — `git init` and nothing else, no plugin on disk, no
`--plugin-from`:

    plugin from this release
    files vendored:   34
    manifest version: 0.33.0
    entry point:      -rwxr-xr-x
    G-Supply          pass
    plugin_version:   0.33.0

Every line of that was impossible before this change. The first said nothing, because there was no
embedded tree; the second was 32; the fourth did not exist; the fifth would have failed, because the
tree was two files short of the one the digest covered.

**Before and after, for the two omissions.** `init --vendor` wrote 32 files and now writes 34.
`secrets.yaml`, which `internal/secrets` reads, and `bin/xeno-env.sh`, which the plugin's own
`hooks.json` names, were both absent and both arrive.

**The digest equality in a development build** — the vendored tree and this repository's plugin hash
to the same value, `c116e0cc…`, which is the half of the anchor property a test can hold.

<!-- xeno:section:gaps -->
## Gaps

**The release path is still untested by a release.** Everything above used a binary built locally
with the release's flags and steps run by hand in a copy. The first real exercise is the next
release, and two mistakes would be quiet: a stamp that did not happen leaves `0.0.0-dev` in an
adopter's artifacts, and a digest taken before the stamp fails for every adopter rather than for the
release. The ordering comment is where somebody editing the steps will read it, and that is the whole
of the protection.

**Nothing confirms a published binary carries the digest of the tag's tree.** Open from the intent
before this one, unchanged, and now more consequential because the binary also carries the tree.

**The embedded tree and `.xeno/plugin/` can diverge in a release's checkout without anybody
noticing**, because the copy is a `cp -R` in a shell step. If it failed partially the digest would
still be taken over `.xeno/plugin/` and would not describe the embedded copy. A check comparing the
two after copying would close it and is three lines nobody has written.

**A development build cannot test the embedded path at all.** `Shipped()` returns false, so the
branch that matters to every adopter is exercised only by building the release's way by hand. That is
what was done, once, and it is not in CI.

**`mcp.json` is still absent**, so an adopter's plugin declares no server. Section 13's `--vendor`
list names it; the walk will copy it the day it exists, which is the point of a walk, and nothing
here creates one.

**The repository's own plugin is not a released plugin**, so a released binary run here fails
G-Supply. Correct, and it means the only way to exercise a release against this trail is to vendor
the release's plugin over this repository's own, which would be a change nobody wants to commit.

**Two copies of the plugin now exist in a release's checkout**, briefly. The embedded one is written
into a tracked directory, so a release that failed midway could leave it behind in a dirty tree.
Nothing commits from that checkout, so it is contained, and it is worth knowing that
`internal/plugin/embedded/plugin/` can exist locally and must not be committed.
