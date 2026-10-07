---
intent: github.com/triplem/xeno#231
phase: 04-verification
created: "2026-10-07T09:47:09Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6f6d371592ec84b7f94c9ef537209cb49c657e986ebffc9260a74b5b90ef9474
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - kind: test-report
      result: pass
      produced_by: go test ./...
      sha256: 76d54596ae6f1b8c7f1ce47a80d09a4320a804030e105189c3faf3789a40db6b
      path: evidence/go-test.txt
      job: test
    - kind: build-log
      result: pass
      produced_by: go build, gofmt -l ., go vet ./..., xeno gate verify, git status over the touched paths
      sha256: 03f946e209eb3ceb41ae276dd411f0588d371d17e02e4e74e02ba762ffe370ae
      path: evidence/checks.txt
      job: checks
    - kind: test-report
      result: pass
      produced_by: the reference test, then one line of the block changed, then restored
      sha256: 9948d6c57e4dd54d298cd963be890b8dda66af3c8987778df6d5dd6845ceaccf
      path: evidence/reference-test.txt
      job: reference
    - kind: other
      result: pass
      produced_by: the old README's paragraphs against the new set, the command names, and every relative link
      sha256: c143b370d8a55ddc492a1297f307577c0051ef5fee04bb0a786ae11f1f336d3c
      path: evidence/nothing-lost.txt
      job: nothing-lost
    - kind: other
      result: pass
      produced_by: every file under docs/ tested against the index, and every index entry against the tree
      sha256: 21b679a6b485f8d8399efca511356b1160241e662db77dbafc618f20ea4ffbab
      path: evidence/index-covers-docs.txt
      job: index
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Twelve criteria, by number, all passing. Seven are read by a person, because the deliverable is
a document and a check that a heading exists says nothing about whether the page reads. Five
are checks.

| # | What it asserts | What proves it |
|---|---|---|
| 1 | the README opens with what Xeno is for | read: the first paragraph, which is section 1's statement, and the second, which is the ISO/IEC 42001 one |
| 2 | the problem is stated before the tool | read: "produced separately, at different times" is the first sentence and "Xeno makes that binding provable" the third |
| 3 | the name is a sentence, pointing at Appendix C | read: the third paragraph, four sentences, the last naming the appendix |
| 4 | it links the index and names the normative document at the link | read: the "Reading further" section, which names the process definition and says "Where it and the code disagree, it wins" |
| 5 | `docs/README.md` has one line per file in `docs/`, grouped, with the normative two marked | `evidence/index-covers-docs.txt`: every file under `docs/` appears in the index, and the two normative ones carry the word in their own entry |
| 6 | the README names a handful of commands and points at `--help` and the reference | read: four command lines, then `xeno --help` and `docs/commands.md` |
| 7 | `docs/commands.md`'s block is byte-identical to `usage` | `evidence/reference-test.txt`: the test passes |
| 8 | a test fails when they diverge | `evidence/reference-test.txt`: one line of the block was changed, the test failed naming the page, and it passed again when restored |
| 9 | the trail page carries everything that moved | `evidence/nothing-lost.txt`: the version fields, the two shapes, the coverage table and the mutations are all in it, and the development-build paragraph appears exactly once across the four files |
| 10 | nothing is lost | `evidence/nothing-lost.txt`: 14 of the old README's 19 paragraphs are carried verbatim and the other five are accounted for one by one — the replaced opening, the command list now a superset, one paragraph that was two, one renamed heading and one rewritten path |
| 11 | no normative document changes | `evidence/checks.txt`: the changed and added files are `README.md`, `cmd/xeno/main_test.go` and three new files under `docs/` |
| 12 | the suite passes and every link resolves | `evidence/go-test.txt`, `evidence/checks.txt`, and `evidence/nothing-lost.txt`, which resolves all 18 relative links against the tree |

<!-- xeno:section:results -->
## Results

**The reference test holds, checked both ways.** It passes, and with one line of
`docs/commands.md`'s block changed — `xeno version` to `xeno versions` — it fails naming the
page and printing both texts, then passes again when the line is restored.
`evidence/reference-test.txt` carries all three runs. An equality assertion that was never
seen to fail proves that two files exist.

**Nothing is lost, counted rather than claimed.** 14 of the old README's 19 paragraphs appear
verbatim in the new set. The other five are accounted for one by one in
`evidence/nothing-lost.txt`: the opening, which is the thing the issue is about and whose one
fact — that the register's last two sections are what is built and what is not — is in the new
README; the command list, which `usage` now carries as a superset; the `learning record`
paragraph, which was two paragraphs with no blank line between them and is now two; the
heading the move renamed; and the two-shapes paragraph, whose paths the move rewrote.

**The command list is strictly larger than it was.** 27 commands in `usage` against the old
README's 21, with `decision record`, `evidence declare`, `intent verify`, `question record`,
`review answer` and `scope set` added and nothing dropped. That is the drift measured at P0,
closed by the page carrying the constant instead of a copy.

**The index covers `docs/` and names nothing absent.** All ten other `.md` files under `docs/`
appear, both normative documents carry the word in their own entry, and every link resolves.
Checked by listing the directory and testing each name against the page, not by reading the
page.

**Eighteen relative links, all resolving.** Across the four files, each resolved against the
tree by path. Three were wrong on the first attempt and are the subject of P3's third
deviation.

**Two claims in the new prose were wrong and were corrected before this phase.** The index
said starting a phase is the one command that may use the network; section 12 says "everything
else in the CLI may use the network: starting a phase reads the issue, finishing one writes a
comment", so the entry now says that. And `docs/README.md` described
`orchestrator-evaluation.md` and `v2-delta.md` as retrospective comparisons; both are v2
architecture records and say so in their own first sections. Both were caught by opening the
files rather than by trusting the names.

**The checks.** `go test ./...` passes across 20 packages, `internal/secrets` the slow one at
138.0 s, no failures. `go build`, `gofmt -l .` outside `vendor/` and `go vet ./...` are clean.
`xeno gate verify` recomputes and matches 499 verdicts. The files touched are `README.md`,
`cmd/xeno/main_test.go` and three new files under `docs/`, and no normative document is among
them.

<!-- xeno:section:gaps -->
## Gaps

**Whether the new README reads well is not checked and cannot be.** Seven of the twelve
criteria are a person reading a document. What is checked is that the facts are there, that
nothing was lost, and that the links resolve; whether a stranger deciding if Xeno applies to
them gets there in the first screen is a judgement, and the only reader who has made it is the
one who wrote it.

**Three of the four files have no test at all.** `docs/commands.md`'s block is held against
`usage`. `README.md`, `docs/README.md` and `docs/the-trail-in-this-repository.md` are held by
nothing: the links were resolved once, here, and a file added under `docs/` tomorrow will not
appear in the index and nothing will say so. P1 ruled a link checker out of scope with its
reason, and the index's coverage is the gap that argues for one.

**The per command notes are the second copy the intent set out to remove, one level down.**
They describe behaviour in prose beside a block that a test holds, and nothing holds them. If
`--for` stops accepting a bare number, the note in `docs/commands.md` is wrong and the suite
is green. The old README had the same property; what changed is that the list beside the notes
is now safe and the notes are not, which is a smaller version of the same defect rather than
its absence.

**The coverage table moved unexamined.** Twelve rows naming tests against the plan's
acceptance criteria, carried verbatim as P1 required, and nobody checked that the tests it
names still exist or still assert what the row claims. That was deliberate — a table that
moved and changed in one commit is unreadable both ways — and it means the page now carries a
table whose last verification is however old it was in the README.

**WP16 inherits three pages and a guard.** The reference's test is a guard; the plan's WP16
says the reference is generated from the declarations, which would remove the need for it. A
later reader could take the test as the design rather than as the stopgap it is. The
alternatives section says so and nothing enforces that reading.

**`docs/README.md` is not `index.md`.** Chosen for GitHub, which is where a reader following a
link arrives today. A generator will want the other name, so this is a rename WP16 has to make
and a decision this intent took on a reader's behalf without a generator to test it against.

**Two wrong claims were written and corrected inside this intent.** The network sentence and
the descriptions of the two v2 drafts. Both were caught, and the thing that caught them was
opening the files, which is not a method anything enforces. A third would have shipped.
