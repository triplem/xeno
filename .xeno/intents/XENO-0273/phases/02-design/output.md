---
intent: github.com/triplem/xeno#153
phase: 02-design
created: "2026-10-07T14:23:56Z"
schema_version: "1.0"
runner_version: dev+6b48c17.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d358c4cd6b87c132b7edd0c45313b9ae720520b775490a485f36299c7febd7f4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: Publish internal/index/testdata/example-symbols.yaml as docs/symbol-index.md, verbatim in a fenced block with three paragraphs of prose above it, held against the file by a test in internal/index/index_test.go, with one entry in docs/README.md.
      rationale: '#153 says the format is documented only in a file under testdata and warns in its last paragraph that a second copy of a format description is wrong within a release, since the first copy has a test holding it honest. Reading the example established that every item on the issue''s own list of what a page needs to say is already in its comments, so what is missing is where the description sits and not the description. #231 landed the pattern for the same hazard earlier the same day: docs/commands.md carries the usage constant and a test fails if the two part. The prose reference page was put as the second option with the argument that #153''s what-it-needs-to-say section reads like a specification for one, and declined because a field table cannot be held against a YAML file. Pointing at the file and adding no page was put as the third and declined because a project producing an index would still open a file under testdata, which is the complaint. Generating the page remains WP16''s end state; the test is the cheap half of the same guarantee.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The page carries the file, not a description of it.** The example's comments are already
written for a reader who does not write Go, which is what #151 did, so publishing them is
publishing the description. A page that said the same things in its own words would be the
second copy #153 names as the hazard, and nothing could hold a paragraph against YAML.

**The prose above the block says only what the block cannot.** Three things: that this is the
format Xeno reads and where it is fixed, that the block is the file a test in this repository
reads so it cannot drift, and where to put the result. Everything else — the fields, which are
required, what `kind` may hold, the staleness rule — is in the comments a reader is about to
read. The test for every sentence was whether removing it would leave a reader unable to use
the block, and the field descriptions all failed it.

**The prose goes above the block and not below.** `docs/commands.md` puts its prose above too,
and the reason is the same: a reader who has scrolled past a hundred lines of YAML has already
decided what the page is. The notes in `commands.md` sit below because they are about
particular commands rather than about the page.

**The test compares the first fenced block against the file on disk, not against an embedded
copy.** `index_test` already reads `testdata/example-symbols.yaml`, so the file is the source
either way; comparing the page against the file keeps one authority and one direction of
dependence. Embedding the example in the test would make the test a third copy.

**It lives in `internal/index/index_test.go` and not in `cmd/xeno`.** `TestThePublishedReferenceIsTheUsage`
is in `cmd/xeno` because the thing it holds, `usage`, is declared there. The thing held here is
a file in `internal/index/testdata`, read by that package's test, so the assertion belongs
beside the reader that defines what the format means.

**The index entry goes under "Running it".** `docs/README.md` groups by the errand a reader
arrives with, and a project that wants Xeno to use a symbol index is working out how to run it.
The alternative was a group of its own for formats, which would be a heading over one entry.

**The page is named `symbol-index.md`.** Not `index.md`, which in a `docs/` directory is what a
static site generator takes for the landing page, and not `symbols.yaml.md`. The specification
calls it a symbol index.

**No link from the example back to the page.** The dependence runs one way on purpose: the
example is the authority and does not know it is published. A comment pointing at the page
would have to be kept true by hand, and the test already names the page in its failure message,
which is where somebody who breaks the pair is standing.

<!-- xeno:section:alternatives -->
## Alternatives

**A prose reference page: a field table with types and requiredness, the configuration keys,
the degradation rules.** What #153's "what it needs to say" section reads like a specification
for, and what a reader arriving from a search engine would expect. It is the second copy the
issue's last paragraph warns about, and the asymmetry is the point: the example has a test
holding it honest and a table would have nothing. Declined by the maintainer with that
argument put.

**Point at the file in the repository and add no page.** One copy, no new test, and the
discoverability half of #153 closed by a link. It leaves the other half exactly as it was: a
project producing an index still opens a file under `testdata` to learn the format, which is
the complaint. Declined.

**Generate the page from the example at build time.** What WP16 describes for the command
reference, applied here. It removes the possibility of drift rather than asserting its absence,
and it needs a generator in the build and a check that the committed output matches — which is
a change to how this repository builds, for one page, before WP16 has decided what generates
anything. The test is the cheap half of the same guarantee and the alternatives section is
where that is said so WP16 does not inherit it as the design.

**Move `example-symbols.yaml` under `docs/` and have the test read it there.** Tempting: the
published copy becomes primary and `testdata` stops holding a document. It makes
`internal/index`'s test read across the tree into `docs/`, which is a dependency from a package
to documentation, and it would move a file #151 put where it is on purpose — annotated for a
reader and read by a test, in the directory the test owns.

**Put the format in section 5 of the process definition, in full.** It would be normative,
which a format a project has to produce arguably should be. Section 5 already fixes what Xeno
requires of an index in prose and deliberately does not give a schema; adding one is a
specification change with its own argument, and it would make the example a copy of a normative
block rather than the other way round.

**Nothing, and close #153 as not worth a page.** The index is optional, no gate reads it, and a
phase runs without one. It was not put as an option because the issue is about a format a
project must match exactly to get any benefit, and "optional" is a reason for the index, not
for the description.

<!-- xeno:section:impact -->
## Impact

`docs/symbol-index.md`: new. Three short paragraphs and the example verbatim in a fenced block.

`internal/index/index_test.go`: one test, beside the one that reads the example.

`docs/README.md`: one entry under "Running it".

Nothing else. No normative document, no change to the example, no change to the reader.

**For a project that wants to produce an index.** The format is on the web, in one page, with
the annotations #151 wrote and without a clone. That is the whole of what #153 asked for.

**For a reader who finds the page later.** It cannot be wrong about the fields, because the
fields are the file. It can be wrong about the three things its prose says, which is the
surface this change adds and the reason the prose is three paragraphs rather than ten.

**For the trail.** Nothing. No artifact, gate or verdict is touched, and `xeno gate verify`
recomputes as before.

**For WP16.** A third page it would otherwise write, and a second guard it may read as a
design. The alternatives say plainly that generating the page is the fix and the test is the
stopgap.

**What gets slightly worse.** A fourth file under `docs/` that nothing checks for presence in
the index. XENO-0270 recorded that gap; this adds to it without changing the argument.
