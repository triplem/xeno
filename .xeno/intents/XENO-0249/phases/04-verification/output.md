---
intent: github.com/triplem/xeno#201
phase: 04-verification
created: "2026-10-05T12:55:28Z"
schema_version: "1.0"
runner_version: dev+325b27a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 06ddf9ce80ab73d53969a1ecae8b2a48c09498950bab6b7859b3abfe6dd6e77f
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

Seven tests in `internal/plugin/hash_test.go` against the eleven criteria, plus three rehearsed
runs of the script and one reading of the workflow.

| criterion | what answers it |
|---|---|
| 1, `HashTree` and its contract | `TestHashIsTheTreeAtTheVendoredPath` for the computation, `TestATreeThatIsNotThereHashesToEmpty` for the empty string |
| 2, `Hash` unchanged at its call sites | `TestHashIsTheTreeAtTheVendoredPath`, and `go build ./...` with no call site edited |
| 3, location independence | `TestATreeHashesTheSameAtTwoPaths`, and `TestATreeThatDiffersHashesDifferently` for both halves of the converse |
| 4, `Differences` | `TestDifferencesNamesWhatTheTreesDisagreeAbout`, `TestDifferencesIsEmptyWhereTheTreesAgree`, `TestDifferencesSaysWhichTreeIsNotThere` |
| 5, the script's two paths | rehearsed, not unit tested: no flag prints the digest with a note on standard error; `--require-shipped` with no tree exits 1; with a matching tree prints the digest; with a stale file names `STALE.txt`; with one byte changed names `contents differ: .claude-plugin/plugin.json` |
| 6, the workflow | read, not run. See the gaps section |
| 7, the value changed and nothing compares across it | `grep` for `ExpectedDigest` over the tree, which finds the gate, the variable and tests that compute rather than pin; and the digest printed by the script, `1c9bd14a…`, against what the old code produced |
| 8, the register row | A96, the last row |
| 9, the suite and the verdicts | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l`, `./xeno gate verify` |
| 10, the workflow is read | this phase saying so |
| 11, one commit | the commit itself |

The script has no test of its own because it is `//go:build ignore`, which is what the other
scripts here are and why none of them has one. The logic that could be wrong is in the package
and is tested there; what the script adds is the flag, the exit code and the wording, and those
were exercised by running it four times rather than asserted.

`TestATreeThatDiffersHashesDifferently` covers the stale-file case as well as the
contents-differ case, because the stale copy left by an earlier run in the same checkout is the
specific failure #201 names and the one a contents comparison alone would miss.

The fixture carries `.claude-plugin/plugin.json` so that a walk skipping dot-prefixed names
would fail the test. That is not hypothetical: `embed` needs `all:` for exactly this tree, which
`embedded.go` records, so a hasher with the same blind spot is a plausible mistake.

<!-- xeno:section:results -->
## Results

Ten criteria met, one pending until the commit.

1. Met. `HashTree(dir)` hashes the tree at `dir` with paths relative to it, sorted, and
   returns the empty string where `dir` is absent.

2. Met. `Hash(root)` is one line over `HashTree` and no call site changed: G-Supply at
   `internal/gates/gates.go:280`, `xeno init --vendor` and the script all build untouched.

3. Met. The same tree at `vendored` and at `somewhere/else/entirely` hashes identically; a
   one-line content change and an extra file each change the hash.

4. Met. `Differences` names `contents differ:` and `only in <tree>:` in both directions, is
   empty for identical trees, and says "is not there" or "neither … is a tree" rather than
   listing every file as missing.

5. Met, by rehearsal rather than by a unit test, which the test mapping states. Four runs: no
   flag prints the digest and warns; the flag with no tree exits 1 naming the path the release
   copies to; with a matching tree prints `1c9bd14a…`; with a stale file and with a one-byte
   change, exits 1 and names `STALE.txt` and `.claude-plugin/plugin.json` respectively.

6. Met as a change and unverified as a behaviour. The flag is on the line that produces
   `DIGEST`, between the copy and the build, with the reason beside it. Nothing here runs
   `release.yml`, which the gaps section states rather than this one claiming a run.

7. Met. The digest is now `1c9bd14a9a68ba4e847dcdb08d180d84e3c2e044cd801e58ca85f89d035acf17`
   over the same bytes that hashed to something else before. Every `ExpectedDigest` reference
   in the tree is the gate, the variable itself, or a test that computes the value; none pins
   a literal, so nothing compares across the change.

8. Met. A96 is the last row and carries the reason, the blast radius, and what the change
   unblocks.

9. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints
   nothing, `go test ./...` is `ok` across all eighteen packages that have tests, and
   `./xeno gate verify` exits 0 over 372 verdicts: the 367 that existed are intact and the
   five are this intent's own judged phases.

10. Met. This phase says the workflow was read and not run, and names it in gaps.

11. Pending. The commit comes after this phase is judged, because this phase's artifacts are
    part of what it commits.

<!-- xeno:section:gaps -->
## Gaps

The workflow cannot be rehearsed here and is the half of this change that matters most.
Nothing in this repository runs `release.yml`; the flag was placed by reading, and the script's
behaviour under it was exercised directly. So the thing verified is "the script refuses the
wrong tree" and the thing assumed is "the release calls the script that way". The assumption is
one line, visible in the diff, two lines below the copy it guards — and it is still an
assumption, and the first release after this merges is what tests it.

The guard is opt-in, so a refactor can drop it silently. That is the cost P2 accepted for a
script that stays usable in a development tree, and the mitigation is adjacency rather than a
mechanism. A workflow that dropped `--require-shipped` would go back to the original defect
with nothing reporting it, which is the same shape as the defect itself: correct code, and
nothing checking that it is called.

The embedding is still unchecked. Between the script's comparison and the binary there is a
`go build` reading `//go:embed all:embedded`, and nothing confirms that what was embedded is
what was on disk. Narrower than what this closes — the directive reads the path the copy wrote
— and it is the published-binary check that would close it, which needs a release artifact to
fetch and is #198's.

Nothing recomputes a published binary's digest from its tag. This intent makes it possible and
does not do it, which the scope said; it is the other thing the issue named as unconfirmed.

`internal/plugin/embedded/plugin` is not gitignored, and this phase rehearsed the release,
which means it created exactly the untracked tree that a `git add -A` would commit and that
would make a development build claim to be a release. Found here, left as a finding, named in
P3's deviations. One line of `.gitignore` closes it and it belongs to whoever owns the
release's ergonomics rather than to the digest.
