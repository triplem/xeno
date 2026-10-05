---
intent: github.com/triplem/xeno#239
phase: 04-verification
created: "2026-10-05T09:00:51Z"
schema_version: "1.0"
runner_version: dev+3f5ad34.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c6f18d4c582d7cb0c7fd508ccd86a41b496a6ffc44a785397693bbbd79edc05b
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

No test is added, and the reason is a fact established in P0 rather than a concession:
nothing in this repository opens any of the four documents by path. Every reference in Go
and in CI is a comment. So there is no behaviour to assert and a test asserting a path
would be a test of this commit's text.

What stands in for it is the existing suite plus three counts:

| criterion | what answers it |
|---|---|
| 1, the four moved | `git status --porcelain` shows four `R` entries; `git ls-files docs/` shows the new paths from the index rather than the working tree |
| 2, references resolve | one grep over the tree excluding `.xeno/`, `vendor/` and `.git/` for all four old names |
| 3, the layout block | read back by eye against `ls docs/` and the root listing |
| 4, the record note | read back by eye |
| 5, A94 | the rename-aware diff of the register |
| 6, verdicts | `./xeno gate verify`, with `--intent XENO-0246` to separate this intent's own |
| 7, the suite | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l` |
| 8, one commit | the commit itself |

The suite is evidence about the sweep, not about the rename: it would pass identically had
every reference been left stale, which is why criterion 2's grep rather than `go test` is
the check that matters here.

<!-- xeno:section:results -->
## Results

Eight criteria, seven met as written and one met in substance after its wording was found
wrong.

1. Met. Four renames, `R` in `git status`, and `git ls-files` confirms
   `docs/m0-gate-job.md` from the index, which is the check that matters for the one
   rename that changes case as well as directory.

2. Met, with the declared exemption. One grep for all four old names over the tree outside
   `.xeno/`, `vendor/` and `.git/` returns exactly one line,
   `docs/implementation-plan.md:2009`, which criterion 2 exempted before the work began.

3. Met. The layout block lists `docs/assumptions.md` and `docs/m0-gate-job.md` in the
   `docs/` group with their old names as provenance, and the three other `docs/` paths it
   already printed are correct for this tree.

4. Met. Two lines under the title.

5. Met. A94, one row, naming both kinds of reference it covers.

6. Met in substance, and the criterion was wrong. It asked for exit 0 over 351 verdicts,
   the count before the move, which this intent cannot produce: a verdict is written per
   judged phase, so finishing six phases adds six. `./xeno gate verify` reports 355 at exit
   0 with four phases judged, and `--intent XENO-0246` reports 4 of them, so the 351 that
   existed before are intact and the growth is this intent's own. The criterion should have
   said that the pre-existing 351 still verify and the total grows only by this intent's
   phases. Stated here rather than corrected in P1, because P1 is sealed.

7. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/`
   prints nothing, and `go test ./...` is `ok` across all eighteen packages that have
   tests, with two reporting no test files as before.

8. Pending at the time of writing: the commit is made after this phase is judged, because
   the artifacts of this phase are part of what it commits.

<!-- xeno:section:gaps -->
## Gaps

Nothing checks that a reference to a document resolves. That is the whole of what this
intent relies on a person for, and it is WP16's to close: a link checker over the tree
would have caught the four old paths in `README.md` without being told they were there,
and would catch the next one. Until it exists, criterion 2 is a grep somebody has to
remember to run, which is the class of defect `docs/clause-readers.md` was written to
enumerate and A90 explains.

The exempt line is unchecked by construction. `docs/implementation-plan.md` line 2009 will
stay stale until a person corrects it, and nothing will report it in the meantime — a link
checker would have to be taught to ignore it or it would fail CI on a line the agent may
not touch, which is a question for WP16 rather than an answer this intent has.

The sealed references are unchecked and must stay so. 274 mentions of the old paths sit in
artifacts whose verdicts depend on their bytes, and any checker this project adds has to
exclude `.xeno/intents/` or it reports 274 findings nobody may act on. That exclusion is
not written down anywhere a checker's author would find it, which is this intent's one
contribution to the next one's difficulty; A94 is where it is now written.

This intent's own P0 scope names the four old root paths and is now sealed, so the trail's
count of stale references includes four this intent added. That is not avoidable and P1's
constraints said so in advance: the scope records the tree the phase was given.

Case sensitivity is verified on one filesystem. `git ls-files` reads the index, so the
rename is correct in what is committed, but a reviewer cloning onto a case-insensitive
filesystem is outside what this phase can check.
