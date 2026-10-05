---
intent: github.com/triplem/xeno#201
phase: 05-review
created: "2026-10-05T12:56:23Z"
schema_version: "1.0"
runner_version: dev+325b27a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7b0ddbf8f29c5ac79b8e99c1bebadd756c057afabbe7ed7409cd9728d0231275
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Two, both named against what they depart from. One addition beyond P2''s design: treeLines, extracted so Differences inherits the hash''s per-file Normalise rather than reimplementing the comparison — which is P2''s own objection to diff -r, and would have been as true of a hand-rolled comparison. One departure from P3''s method rather than from the plan: the phase was opened before the work, where the previous three intents opened it after, so its elapsed figure measures the work; P3''s learning records the habit. A third thing was found and deliberately left, the ungitignored embedded tree, which is a finding about the release''s ergonomics and not about the digest.'
      result: deviation
      rule: deviations-are-traceable
    - note: 'An interface outside this intent does change, and nothing migrates because nothing can depend on it. plugin.Hash keeps its signature and contract, so every call site is untouched, but its VALUE changes for identical bytes: the lines now carry paths relative to the plugin directory. The dependants would be anything holding a digest computed by the old code, and there are none — a released binary carries the digest taken at its own release together with the code that recomputes it, no test pins a literal, and section 13 forbids recording one in the repository. So the migration note is that there is nothing to migrate, and A96 is where a reader who meets a moved number finds out why.'
      result: deviation
      rule: interface-change-needs-a-migration-note
    - note: None added. go.mod is untouched; the change uses os, filepath, sort, strings and the existing internal/hashing, all already imported by this file, and flag in a script that is build-ignored. The one dependency weighed and rejected was on diff(1), which P2's alternatives rejected because it compares raw bytes where the hash compares normalised ones, so the guard would have disagreed with the thing it guards.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No normative document is touched: section 13 describes the anchor and
says nothing about the line format, so the basis is this package's to choose and no
specification commit precedes this. No field, gate, tool or rule is added — `HashTree`,
`Differences` and `ShippedDir` are new surface in an existing package, and G-Supply is
unchanged. The branch carries one intent, the commit references #201, and the issue carries
`wp7`.

The acceptance criteria. Ten met, one met by the commit this phase precedes. Criterion 6 is met
as a change and unverified as a behaviour, which P4's results and gaps both say rather than
counting the workflow as tested.

The non-goals held. G-Supply untouched. No digest written into the repository. No new line
format — a sum, two spaces, a path, with only the base changed. No replacement for `cp -R`. No
published-binary check. No unconditional comparison, for the reason `embedded.go` gives.

What a reviewer should check is the one assumption the tests cannot reach: that
`release.yml` calls the script with `--require-shipped`. It is one line in the diff, two lines
below the `cp`. Everything else is covered — the location independence by test, the refusals by
rehearsal, including the stale-file case the issue names.

What changed in how this intent was run, rather than in what it produced: P3 was opened before
the work instead of after it, so its elapsed figure measures the work for the first time in
four intents. P3's learning records why the previous three did not.

Two findings are left rather than absorbed, both named in P4's gaps: the embedding between the
script and the binary is still unchecked, and `internal/plugin/embedded/plugin` is not
gitignored, which this phase's own rehearsal demonstrated by creating it.

<!-- xeno:section:release-notes -->
## Release notes

The release now checks that the plugin tree the binaries carry is the tree the digest was taken
over. Until this it compared them with nothing.

`.github/workflows/release.yml` stamps the plugin's version, copies `.xeno/plugin` to
`internal/plugin/embedded/plugin` for the binaries to carry, and takes the digest over
`.xeno/plugin`. The two trees agreed because `cp -R` had run a moment earlier. `set -eu`
catches a copy that fails; nothing caught a copy that succeeds while producing a different tree
— a stale file from an earlier run in the same checkout, a path the hasher and `cp` treat
differently, or a later refactor that changes one path and not the other.

What that would have cost falls on adopters. `xeno init --vendor` gives them the embedded tree,
G-Supply hashes what they vendored and compares it against the digest the binary carries, and if
the trees differed at release time then every phase of every project on that release fails
G-Supply, with a finding naming two hex strings and nothing they can act on. The release that
made the mistake passes its own checks.

So `scripts/plugin-digest.go --require-shipped` now prints a digest only if the embedded tree
exists and hashes equal to the vendored one, and the workflow passes the flag beside the copy it
guards. A refusal names the differing paths — "only in …: STALE.txt", "contents differ:
.claude-plugin/plugin.json" — rather than two digests and nothing to do.

**The digest's value changes for identical bytes.** `plugin.Hash` keyed each line on the file's
path relative to the repository root, so the same tree at two paths hashed differently by
construction and the only definition that could have compared them was unusable. Lines now
carry the path relative to the plugin directory, so a tree hashes to what it is rather than to
where it sits. Nothing compares a digest across the change: a released binary carries the digest
taken at its own release together with the code that recomputes it, no test pinned a literal,
and section 13 forbids recording one in the repository. A96 records it, because a number that
moved for no reason visible in a diff reads as a bug.

For adopters: nothing to do, and no action is available or needed. `plugin.Hash` keeps its
signature and contract, G-Supply is unchanged, and no digest anywhere needs rewriting.

What this also unblocks, and does not do: recomputing a published binary's digest from its tag
needs a hash that is a function of the tree alone, and now there is one. That check is still
nothing's job (#198).

<!-- xeno:section:residual-risk -->
## Residual risk

The guard is opt-in and the workflow is unrehearsed, which is one risk with two faces. Nothing
here runs `release.yml`, so what is verified is that the script refuses the wrong tree and what
is assumed is that the release calls it with the flag. A refactor that moves the copy and drops
`--require-shipped` restores the original defect silently — correct code, nothing checking that
it is called, which is the shape of the defect being fixed. Placing the flag two lines below the
`cp` with the reason beside it is the whole mitigation. The first release after this merges is
what tests it, and a reviewer reading one line of the diff is worth more here than any test in
this PR.

The digest moved and nothing will say so at the moment it matters. The reasoning that nothing
compares across the change is sound and was checked rather than assumed, but it rests on an
absence — no pinned literal, no recorded digest — and an absence is only as good as the search
for it. If some artifact outside this repository holds a digest from an earlier release and
compares it against a newly built binary, it will see a mismatch and the finding will name two
hex strings, which is the experience this intent exists to prevent. A96 is where that reader
lands; it is the best available answer and it is a document, not a mechanism.

The embedding is still unchecked, which is narrower than what closed and is not closed. Between
the comparison and the binary sits a `go build` reading `//go:embed all:embedded`, and nothing
confirms what was embedded is what was on disk. The directive reads the path the copy wrote, so
the window is small; the thing that would shut it is the published-binary check.

`internal/plugin/embedded/plugin` is not gitignored and this intent's own rehearsal created it.
A developer doing the same and running `git add -A` commits a tree that makes
`plugin.Shipped()` return true for a development build — precisely the claim `embedded.go` says
a development build must never make, and it would then vendor a plugin whose digest no runner
expects. One line closes it; it is left as a finding because it belongs to the release's
ergonomics rather than to the digest, and absorbing it would put an unrelated guard in this
commit.

What is not a risk: the 367 pre-existing verdicts, which `gate verify` confirms at exit 0; every
call site of `plugin.Hash`, which is unchanged in signature and contract; and a development
build, which carries no anchor and so has G-Supply report `not-implemented` exactly as before.
