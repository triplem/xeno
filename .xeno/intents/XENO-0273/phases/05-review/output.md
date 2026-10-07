---
intent: github.com/triplem/xeno#153
phase: 05-review
created: "2026-10-07T14:30:07Z"
schema_version: "1.0"
runner_version: dev+6b48c17.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e15e2eabf0fc5e9dd9a0236e672dfbefdd3951b3f08a9d88a67a4f67d6a2f8ec
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'P3 recorded no deviation: every decision P2 took was carried out as written. It names one thing that departs from nothing — the fenced block is tagged yaml, which P2 did not specify, and the test''s pattern requires the tag, so dropping it fails with carries no yaml block rather than with a mismatch. That is a note about a choice made in implementation rather than a departure from a decision, and P3 says so rather than dressing it as either.'
      result: met
      rule: deviations-are-traceable
    - note: 'No interface changes. No artifact gains or loses a field, no gate changes, no command changes its behaviour or its refusals, and the symbol index format itself is untouched: example-symbols.yaml has no diff and is not in git status. The only Go change is one test function. A project that already produces an index for Xeno needs to do nothing; what it gains is being able to read the format without cloning this repository. xeno gate verify reports 517 verdicts verified.'
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: None added. go.mod and go.sum are untouched and nothing compiles differently. The test's import block gains regexp, which is standard library and already imported by the equivalent test in cmd/xeno.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** No normative document is touched: section 5 and Appendix A are
cited by the page and quoted nowhere. Nothing is invented — the page's own test was that no
sentence may state a field, a default or a value, and `evidence/prose-adds-nothing.txt` is that
test carried out. The change belongs to WP16 by subject, which is #153's own argument for why
it is not WP15's, and to intent XENO-0273 for issue #153.

**What the issue asked, and what it got.** #153's complaint is that the format is discoverable
only by cloning this repository and opening a file under `testdata`. It is now a page in
`docs/`, linked from the index, carrying the same annotated file with a test that fails if the
two part. The issue's own warning — that a second copy of a format description is wrong within
a release while the first copy has a test — is honoured rather than worked around, which is
why there is no field table.

**What a reviewer should look at first.** The three paragraphs, because they are the only thing
here that can be wrong. Each was checked against the line that carries it and nothing holds
them afterwards. The one I would re-read is the sentence sending a reader to section 5 for the
format and to Appendix A for `index.path` and `index.max_age_hours`: it is the kind of claim
that is easy to write from memory and I had it wrong in the intake's first draft, where the
context rationale said "section 5's `index` block", which does not exist.

**The second thing: #153 is closed with half of its own analysis declined.** The issue's "what
it needs to say" section reads like a specification for a prose reference — the eight fields,
which are required, what the keys do — and that page was not written. The reason is the issue's
last paragraph arguing against exactly it, so the issue disagrees with itself and the
maintainer resolved it. Somebody reading the closed issue will find a list of requirements and a
page that does not look like it, and this paragraph is the only place that says why.

**The question no rule asks: is a published YAML comment block documentation?** It is annotated
for a reader who does not write Go, which is what #151 set out to do, and it is still a file
format shown rather than described. A reader wanting to know whether `container` is required
counts on reading a comment that says `container` is omitted where nothing encloses, which is
an answer and not a schema. The judgement is that the annotation is good enough to publish and
that a second description would be worse than a thin one; somebody could reasonably want both,
and what makes that affordable is WP16 generating one rather than a person maintaining it.

**And: two guards now stand where the plan wants a generator.** `docs/commands.md` against
`usage`, this page against the example, both committed copies with a test. WP16 says the
reference is generated from the code. Each intent's alternatives says the test is the stopgap
and the generator is the fix, and that is the only thing stopping the pattern being read as the
design.

**What is deliberately not here.** No prose reference page, no edit to the example, no move out
of `testdata`, no site, no link checker, and no new facts. P1's non-goals carry a reason for
each.

<!-- xeno:section:release-notes -->
## Release notes

The symbol index format is published. [docs/symbol-index.md](symbol-index.md) carries the
annotated worked example in full — the three required provenance fields with their reason, the
five symbol fields and no more, `container` omitted where nothing encloses, and `kind` as a
free string Xeno compares against nothing — so a project that wants Xeno to read an index can
see the shape it has to produce without cloning this repository.

It is not a second copy. The block is
`internal/index/testdata/example-symbols.yaml`, the file `internal/index`'s own test reads, and
a second test fails if the page and the file stop being identical. So what is published is what
the reader accepts, and there is one place the format is written down.

The page says three things the block cannot: that the format is fixed in section 5 of the
process definition under "A symbol index, not a graph", with `index.path` and
`index.max_age_hours` in Appendix A's `project.yaml` rather than in section 5; that a project
produces the index and Xeno ships no indexer, so tree-sitter, ctags, a build system or a
language server all serve; and that nothing depends on the index being there, because absent,
unreadable, malformed and stale are four causes with one outcome and no gate reads it.

`docs/README.md` lists it under "Running it", beside the command reference.

<!-- xeno:section:residual-risk -->
## Residual risk

**The page's three paragraphs are the only thing here that can be wrong, and nothing holds
them.** Each was checked against the line that carries it, once. If section 5's wording moves,
or `index.go`'s degradation rules change, the page is wrong and the suite is green. The block
cannot drift; the prose around it has exactly the property this intent existed to remove, at
three paragraphs instead of a field table.

**The format is published and is still a worked example.** #153 recorded that as the unmet half
of #151's second criterion and it stays unmet by choice: a reader who wants types and
requiredness in a table reads YAML comments. The judgement is that one description with a test
beats two without, and somebody who wants both has to wait for WP16 to generate the second.

**#153 closes with half of its own analysis declined.** Its "what it needs to say" section
reads like a specification for the page that was not written, and its last paragraph is the
argument against writing it. A reader of the closed issue finds requirements and a page that
does not match them, and only P5 says why.

**Two committed copies with tests now stand where WP16 wants a generator.** This page and
`docs/commands.md`. Each intent's alternatives says the test is the stopgap, and nothing stops
a later reader treating the pattern as the answer and leaving the generator unbuilt.

**A document's fence tag is now load bearing for a package's test.** `internal/index`'s test
requires ` ```yaml `. Tidying the Markdown breaks it, which is an odd direction of dependence
and was chosen because the failure message is better than a mismatch.

**Four pages were added under `docs/` today and nothing checks the index lists them.** The gap
XENO-0270 recorded, one page larger.

**Nobody searching the web arrives anywhere.** WP16 has published nothing, so "documented
somewhere a project would look" means the repository's own `docs/` on GitHub. That is less than
#153's title asks for and more than a file under `testdata`, and the remaining distance is
WP16's.
