---
intent: github.com/triplem/xeno#239
phase: 01-requirements
created: "2026-10-05T08:51:54Z"
schema_version: "1.0"
runner_version: dev+3f5ad34
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6eb8203f35bcd313189576cae8ed45d3e7d6bf88ae8d84fc0ddea3c5b07f259d
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

1. `docs/assumptions.md`, `docs/supply-chain.md`, `docs/clause-readers.md` and
   `docs/m0-gate-job.md` exist, and no file of any of those four names remains in the
   repository root. The moves are recorded as renames, so `git log --follow` reaches the
   history of each.

2. Every reference to the four outside the trail resolves, with the single exception named
   in intake: `docs/implementation-plan.md` line 2009. That means `README.md`,
   `CONTRIBUTING.md`, `CLAUDE.md`, the five workflow files under `.github/workflows/`,
   `internal/runner/exchange.go`, and the cross-references the four documents make to each
   other.

3. `M0.md`'s own layout block names its own new place, and the three other paths that block
   prints are correct for the tree after the move.

4. The moved M0 document says at the top that it records how something was done once.

5. `ASSUMPTIONS.md` carries one new row for the references the move leaves stale: the
   sealed ones in the trail and the one line of the plan.

6. `./xeno gate verify` exits 0 over 351 verdicts, the count on this tree before the move.
   A different count means an artifact was touched, which this intent must not do.

7. `go build ./cmd/xeno`, `go test ./...` and `go vet ./...` pass, and
   `gofmt -l . | grep -v '^vendor/'` prints nothing.

8. One commit, closing both #239 and #222.

<!-- xeno:section:non-goals -->
## Non goals

Not a link checker. WP16 owns that, and the reason to do this move first is that a checker
over a tree whose documents are split between two places by no rule has to be taught the
exception before it can be trusted.

Not a move of the five documents that stay. `README.md` is the entry point;
`CONTRIBUTING.md`, `SECURITY.md`, `LICENSE` and `NOTICE` are surfaced by the host from the
root, and the convention a reader already has is worth more than the consistency.

Not a rewrite of any sealed artifact. The old paths appear throughout the trail inside
`artifacts_hash`, and correcting one would stale every verdict over it. They are recorded
as stale, not chased.

Not an edit to either normative document, including the one line of the plan that names
`ASSUMPTIONS.md`. That correction is a person's commit, made before the code that follows
from it, and this intent neither makes it nor waits for it.

Not a reflow or a rewrite of the four documents' contents. `M0.md` gains one line and has
its layout block corrected; otherwise the text that moves is the text that arrives, so the
diff is a rename plus the references and nothing a reviewer has to read twice.

Not a renumbering of anything. A-numbers are cited without the filename in most places,
which is why eight references outside the trail rather than eighty, and no assumption
changes its identifier because the file it lives in changes its path.

<!-- xeno:section:constraints -->
## Constraints

The first standing rule binds the plan. One line of it names `ASSUMPTIONS.md`, and no
amount of this intent's work may touch it, so the acceptance criteria exempt it explicitly
rather than quietly failing criterion 2.

`artifacts_hash` binds the trail. 351 verdicts stand on this tree and every one of them
must still verify afterwards, which means the move may not touch a single file under
`.xeno/intents/`. The sealed references to the old paths are therefore not a thing to fix
but a thing to state, and the verdict count is the check that nothing was fixed by
accident.

This intent's own P0 scope names the old paths, and so joins the sealed references the
moment the move lands. That is not avoidable: the scope records what the phase was given,
the phase was given the pre-move tree, and the scope lies inside P0's `artifacts_hash`.

Nothing in the repository opens any of the four by path. Every reference in Go and in the
workflows is a comment, so no build, run or workflow step can break on the rename, and the
test suite is evidence about the sweep rather than about the move.

Renames go through `git mv`, so the history of each file stays reachable with
`git log --follow`. `M0.md` to `docs/m0-gate-job.md` changes both directory and case, which
is the one rename a case-insensitive clone could collapse; it is done as a single `git mv`
and the result is read back from the index rather than from the working tree.

One commit. Both issues close on it, `Closes #239, closes #222`, because a host reads only
the first reference after one keyword.
