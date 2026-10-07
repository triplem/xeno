---
intent: github.com/triplem/xeno#153
phase: 04-verification
created: "2026-10-07T14:27:24Z"
schema_version: "1.0"
runner_version: dev+6b48c17.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 48f3d510201dec69493b140981a9c6c2c8d553b882e9ccf3cde3145d828a0cbd
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
      sha256: 7ec9468beebbf53d30cee63e0d94c546a3cfb73b5db8dc016abee9c126cc2638
      path: evidence/go-test.txt
      job: test
    - kind: build-log
      result: pass
      produced_by: go build, gofmt, go vet, xeno gate verify, the files touched, and every relative link resolved
      sha256: 87d07f359e0d1d5cb7326e358823e439909d971e74d3e535bd6117c6ff096677
      path: evidence/checks.txt
      job: checks
    - kind: test-report
      result: pass
      produced_by: the page test, then the block changed, then the yaml tag dropped, then restored
      sha256: 6c2935159102edbbb2d180d77a394ddf984cc5e4c244b108accc68abcf2f18f3
      path: evidence/page-test.txt
      job: page
    - kind: other
      result: pass
      produced_by: each claim in the page's prose against the line that carries it, and the field-name check
      sha256: 8e455df11b26247a30ae82323cfc87fb7d4e6db62f2514fc97ec7d3ea2779b16
      path: evidence/prose-adds-nothing.txt
      job: prose
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Ten criteria, by number, all passing. Six are checks and four are read in the files.

| # | What it asserts | What proves it |
|---|---|---|
| 1 | the page's first fenced block is byte-identical to the example | `evidence/page-test.txt`: the test passes |
| 2 | a test fails when the two part, checked by making them part | `evidence/page-test.txt`: one line of the block changed, the test failed naming both files; and the fence's `yaml` tag dropped, which failed with "carries no yaml block" rather than with a mismatch |
| 3 | the page's prose says nothing the example does not | `evidence/prose-adds-nothing.txt`: each of its claims against the line that carries it in the example, section 5, Appendix A or `internal/index/index.go`, and no field name appears as a backticked token in the prose |
| 4 | it says where the format is fixed, and that the two keys are in Appendix A | `evidence/prose-adds-nothing.txt`: line 659 is inside section 5 and lines 2293–2294 inside Appendix A |
| 5 | it says a project produces the index and Xeno ships no indexer | read: the first paragraph, against section 5 line 662 and the example's line 5 |
| 6 | it says absent, stale, unreadable or malformed is not an error | read: the second paragraph, against `index.go` line 59 and section 5 line 670 |
| 7 | `docs/README.md` carries an entry under "Running it" | read: beside `commands.md`, and the link resolves |
| 8 | `example-symbols.yaml` is unchanged | `evidence/checks.txt`: no diff against it, and it is not in `git status` |
| 9 | no normative document changes | `evidence/checks.txt`: the touched files are `docs/README.md`, `docs/symbol-index.md` and `internal/index/index_test.go` |
| 10 | the suite passes and every link resolves | `evidence/go-test.txt`, `evidence/checks.txt`, which resolves all fourteen links against the tree |

<!-- xeno:section:results -->
## Results

**The test holds the pair, checked three ways.** It passes. With one line of the page's block
changed — `tool: go-symbols` to `go-symbolz` — it fails naming both files and printing each.
With the fence's `yaml` tag removed it fails with "carries no yaml block, so there is nothing
to hold against the example", which is the right message for that mistake rather than a
mismatch. Restored, it passes. `evidence/page-test.txt` carries all four runs.

**The prose was checked claim by claim, not asserted.** `evidence/prose-adds-nothing.txt` puts
each sentence beside the line that carries it: section 5 line 670 and `index.go` line 13 for a
verdict never depending on the index, `index.go` line 59 for the four causes with one outcome,
section 5 line 662 and the example's line 5 for Xeno shipping no indexer, section 5 line 665
and the example's line 6 for fixing the shape of the answer, and `index_test.go` line 160 for
the test that reads the file.

**And where the format is fixed was checked rather than remembered.** Line 659 is inside
`## 5. Artifact schema` and lines 2293–2294 inside `## Appendix A. The full project.yaml`, so
the page's sentence sending a reader to section 5 for the format and to Appendix A for
`index.path` and `index.max_age_hours` is right in both halves. A page saying "section 5" for
the configuration keys would have sent a reader to the wrong place.

**No field name appears in the prose.** Checked by looking for each of the nine field names as
a backticked token above the fenced block, which is how the page would name one: none. A looser
grep for the same words as bare text returns one line, "`internal/index` reads that file",
where the match is the English word; that is recorded, because the looser grep is the one
somebody would run first and its hit looks like a failure.

**The example is untouched.** No diff against it, and it is not in `git status`. The three
files the change touches are `docs/symbol-index.md`, `docs/README.md` and
`internal/index/index_test.go`.

**The checks.** `go test ./...` passes across 20 packages. `go build`, `gofmt -l .` outside
`vendor/` and `go vet ./...` are clean. `xeno gate verify` matches 517 verdicts. All fourteen
relative links in the two touched pages resolve against the tree.

<!-- xeno:section:gaps -->
## Gaps

**The page's three paragraphs are held by nothing.** The block cannot be wrong; the prose can.
Each claim was checked once, here, against the line that carries it, and if section 5's wording
changes or `index.go`'s degradation rules change, the page goes stale silently. That is the
whole surface this change adds and it is three paragraphs wide on purpose.

**The format is published and still not an interface description.** #153 called the worked
example "a worked example doing the job of an interface description" and recorded that half of
#151's second criterion as unmet. Publishing it closes the discoverability half. A reader who
wants a field table with types still reads YAML comments, which the maintainer chose
deliberately over a second copy, and the issue's own argument is why.

**Nothing notices a fourth page missing from the index.** `docs/` now holds four pages added
today and nothing checks that `docs/README.md` lists them. XENO-0270 recorded that gap and
ruled a link checker out of scope with its reason; this makes it marginally worse.

**Two guards now stand where WP16 wants a generator.** `docs/commands.md` against `usage`, and
this page against the example. Both are committed copies with a test, which is the cheap half
of the guarantee WP16 means to get by generating. A later reader could take the pattern as the
design, and the only thing against that is a paragraph in each intent's alternatives.

**The test couples the page's rendering to an assertion.** It requires the fence to be tagged
`yaml`. That is deliberate and it means a renderer change, or somebody tidying the Markdown,
breaks a test in `internal/index` — a coupling from a package's test to a document's
formatting, which is unusual in this tree and is the price of anchoring on the tag.

**Whether the page is findable is untested.** It is linked from `docs/README.md`, which is
linked from the README. Nobody searching the web for a symbol index format arrives anywhere,
because WP16 has not published anything. The issue's complaint was about a project that wants
to produce an index, and such a project now finds the format in the repository without cloning
it, which is less than being published and more than a file under `testdata`.
