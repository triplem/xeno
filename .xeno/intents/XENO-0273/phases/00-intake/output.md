---
intent: github.com/triplem/xeno#153
phase: 00-intake
created: "2026-10-07T14:09:51Z"
schema_version: "1.0"
runner_version: dev+c1f6405
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8a48246c302d57d6df4178b48dd5f9694606b523d9b06ac9ddc21364470b2eaf
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

#151 documented the symbol index format in `internal/index/testdata/example-symbols.yaml`:
annotated for a reader who does not write Go, and read by `index_test` so it cannot drift from
what the reader accepts. It is published nowhere. So a project that wants to produce an index
discovers the format by cloning this repository and opening a file under `testdata`, which
#153 calls a worked example doing the job of an interface description — the half of #151's
second criterion that was recorded as unmet rather than claimed.

## What the example already says

Read rather than assumed, because the page's shape depends on it. Every item on #153's list is
in that file, in its comments: the three provenance fields and that all three are required and
why; the five symbol fields and no more, with the reason call relationships and type resolution
are absent; that `container` is omitted where nothing encloses; that `kind` is a free string
Xeno compares against nothing, with the reason a closed set would impose Xeno's idea of a
symbol on the tool you chose; where the file goes, `index.path`, conventionally
`.xeno/local/index/symbols.yaml` and gitignored because derived data does not belong in the
trail; that an index older than `index.max_age_hours`, default 24, is treated as absent; and
that Xeno ships no indexer, so tree-sitter, ctags, a build system or a language server all
serve.

So what is missing is not the description. It is that the description sits where only somebody
already inside this repository will find it.

## Why the obvious page is the wrong page

#153 names the hazard in its last paragraph: "a second copy of a format description is wrong
within a release, and this one has a test holding the first copy honest". A prose reference —
a field table with types and requiredness — would be that second copy, and nothing could hold a
table against a YAML file. WP16 says the reference is generated from the code for exactly this
reason.

#231 landed the pattern earlier today. `docs/commands.md` carries the `usage` constant in a
fenced block and `TestThePublishedReferenceIsTheUsage` fails if the two part, so the page is
published and machine checked and there is one place a command name is written down. The same
shape fits here, and the thing being carried is already prose-annotated rather than bare.

## What the maintainer chose

Put the three options on 2026-10-07 with the reading above: publish the example held by a test,
write a prose reference page, or point at the file in the repository and add nothing. The first.

<!-- xeno:section:scope -->
## Scope

In scope is `docs/symbol-index.md`, new. It carries
`internal/index/testdata/example-symbols.yaml` verbatim in a fenced block, with a short page of
prose around it saying what the format is for, that the block is the file a test in this
repository reads, that Xeno ships no indexer, and where the configuration keys live. The prose
says nothing the example does not: a page that explained a field for the first time would be
the second copy #153 warns about.

In scope is a test holding the two together, in `internal/index/index_test.go` beside the one
that reads the example. It compares the page's first fenced block against the file on disk and
fails if they part. That is the shape `docs/commands.md` and
`TestThePublishedReferenceIsTheUsage` established for the same hazard earlier today.

In scope is one entry in `docs/README.md`, under the heading about running Xeno where
`commands.md` already sits, because a published format description a project needs before it
can produce an index belongs where a reader looks for how to work with the tool.

## Out of scope

Out of scope is a prose reference page. It was put to the maintainer as the second option, with
the argument that #153's own "what it needs to say" section reads like a specification for one,
and it was declined for the reason the issue itself gives: a field table cannot be held against
a YAML file, so it would be the second copy, and the first copy has a test.

Out of scope is changing `example-symbols.yaml`. It is read by `index_test` and it already says
everything the page needs. Editing it to read better as a published page would be editing a
fixture for a reader it was not written for, and its annotation was the deliverable of #151.

Out of scope is anything about the index reader's behaviour, the configuration schema, or
`index.max_age_hours`. All three are described in the example's comments and fixed in section 5;
nothing about them changes.

Out of scope is the documentation site. WP16 is unstarted and its generator is under question in
#226. This is a file a site can publish, and it is the third such file after `commands.md` and
`the-trail-in-this-repository.md`.

Out of scope is a link checker, and the gap is the same one XENO-0270 recorded: nothing will
notice when a file added under `docs/` is missing from the index. One more page makes that
slightly worse and does not change the argument, which belongs to its own change.

Out of scope is moving `example-symbols.yaml` out of `testdata`. It would make the published
thing the primary copy and the fixture the reference, which is tempting and is a change to how
`internal/index` is tested. #153 asks for the format to be published, not relocated.

<!-- xeno:section:context-rationale -->
## Why this context

Nine files, 316515 bytes.

`internal/index/testdata/example-symbols.yaml` is the thing being published, and it was read
whole before anything was written: the decision between publishing it and writing a prose page
rests on whether its comments already say what #153 lists, and they do.

`internal/index/index.go` is read for what the reader actually does with a stale or malformed
index — `DefaultMaxAge` of 24 hours, and the comment that absent, unreadable, malformed and
stale are four causes with one outcome. It is what lets the page's prose be checked against
behaviour rather than against the example's own claims about it.

`internal/index/index_test.go` holds the test that reads the example, which is where the new
test belongs and what it has to sit beside without duplicating.

`docs/commands.md` and `cmd/xeno/main_test.go` are the pattern. The first is a page carrying a
constant with prose around it; the second holds it with a comparison over the first fenced
block, and its comment says why the block and not the file. Both are read so that this follows
them rather than inventing a second shape for the same problem.

`docs/README.md` is the index the entry goes into, read for its grouping.

`docs/process-definition.md` carries the format itself, in section 5 under "A symbol index, not
a graph": what Xeno requires of an index, that it ships no indexer and why, and that the index
is used while a phase runs and never by a gate, so it may be missing, stale or wrong without
the trail suffering. The two configuration keys are not there but in Appendix A's full
`project.yaml`, which is a distinction worth having read rather than assumed: a page citing
"section 5" for `index.max_age_hours` would send a reader to the wrong place.

`docs/implementation-plan.md` carries WP15, which fixes the format, and WP16, which owns
publishing it and says the reference is generated from the code. #153's own argument for being
WP16's rather than WP15's is in it.

`CLAUDE.md` carries the standing rules: the first, which is why neither normative document is
touched, and the one about a negative result, which applies to the claim that the page says
nothing the example does not.

The links block declares `internal/index` against the process definition, because the format is
fixed in section 5 and read by that package, and a page about it that cited only one would be
describing half of what binds it.

Not in scope: `.xeno/plugin/`, because no skill or template mentions the index format, and the
trail, because nothing in it moves.
