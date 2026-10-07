---
intent: github.com/triplem/xeno#231
phase: 00-intake
created: "2026-10-07T09:36:54Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 441bd135d84cb648244b895bdd0ebea3a0641a8d0af7d03bef5a61eaa623d606
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

`README.md` opens with "The Xeno runner: hashing, gates, the phase sequence, evidence
attachment, the digest and its secret filter, the cost record, the branch rules port and the
symbol index reader", then twenty-four command lines, then nine paragraphs about field values
in this repository's own artifacts, then a coverage table of twelve work packages. It is a
note from one developer to another about a package they are both already inside.

Three things a first reader needs are absent. What the tool is for: the answer is the first
page of the process definition, 2284 lines into a document somebody has to be told to open.
Where the name comes from: Appendix C of the same document has it, and the implementation plan
goes out of its way to decline the question — "the namespace and repository name are chosen and
deliberately not recorded here" — so the one fact about the project that nothing a newcomer
reads explains is its own name. And where to read further: there is no index, and `docs/` now
holds eight files with no statement of which two are normative.

## The command list is two copies of one fact

Twenty-four lines in the README against the `usage` constant in `cmd/xeno/main.go`, both
maintained by hand. The list has drifted: `xeno question record`, `xeno decision record`,
`xeno evidence declare`, `xeno intent verify` and `xeno review answer` exist and are not in
it, and `xeno scope set` is not either.

The repository already refuses this shape elsewhere. `TestEveryCommandInTheUsageResolves`
reads the command names out of `usage` with a regexp rather than listing them again, and says
why: "a third copy would be a third place to be wrong". The README is the second copy that
test was written to avoid, and nothing holds it.

## What the maintainer asked for

On the issue, 2026-10-07: point to an index file in `docs/`, highlighting the process
definition; name the most used commands briefly, then link `xeno --help` and the reference;
and move the rest, naming the gate path's absence of a network call as the example, because it
is deep information that belongs in the documents.

That last part is the half the issue itself got wrong in the reader's favour. #231 says "the
existing material stays ... further down" and lists the no-network-call property, the one
command that does call out, and the two shapes of the trail as things a reader needs. They are
things a reader needs; they are not things a first reader needs, and the no-network-call
property is already in section 12 of the process definition under "Two surfaces, two
promises", which is where a reader who has got that far will be.

## There is nothing to link to yet, which is the honest half

WP16 is unstarted, there is no `mkdocs.yml`, and the generator is under question (#226). So a
reference and an index mean files under `docs/` today. #153 argues against exactly that for a
format description — "not a fifth file beside them" — and it is the same argument. What
settles it here is that the maintainer asked for the index and the reference by name, and that
a page WP16 later publishes is the page WP16 would have had to write anyway.

<!-- xeno:section:scope -->
## Scope

In scope is `README.md`, rewritten rather than edited. It opens with what Xeno is for, in the
terms section 1 uses: requirements, design, code and evidence are produced separately, at
different times, partly by people and partly by models, and the only thing holding them
together is their binding to an intent. It says where the name comes from, in a sentence, and
points at Appendix C for the rest. It links an index. It names the handful of commands a
reader starts with and points at `xeno --help` and the reference for the others. It keeps the
two short facts a first reader can act on, that no telemetry leaves the repository and that
every command naming a next step can be silenced with `--no-next`, and nothing else.

In scope is `docs/README.md`, an index. One line per document with what it is for, the process
definition and the implementation plan marked as the two normative ones, and the others
grouped by what a reader would come to them for. It is what the maintainer asked for by name
and what WP16 publishes as a landing page rather than writing again.

In scope is `docs/commands.md`, the reference the README points at. Its command block is the
`usage` constant and not a copy of it: a test asserts the two are identical, so the page
cannot drift the way the README's list did. That is the pattern this repository already uses
for the same hazard, in `TestEveryCommandInTheUsageResolves`, which reads the names out of
`usage` rather than listing them because "a third copy would be a third place to be wrong".

In scope is `docs/the-trail-in-this-repository.md`, which takes what moves out of the README:
the nine paragraphs about what `runner_version`, `plugin_version` and `tool_version` say in
this trail and why, the two shapes of the intents directory, the coverage table against the
implementation plan, and the two deliberate mutations. The material is unchanged; what changes
is who finds it.

In scope is the test that holds `docs/commands.md` against `usage`, in `cmd/xeno/main_test.go`
beside the one that reads the command names out of the same constant.

## Out of scope

Out of scope is any change to the two normative documents. The purpose sentences are taken
from section 1 and the name from Appendix C by quoting and pointing, not by moving text, and
the implementation plan's refusal to record the name stays as it is: it is a reasonable thing
for a plan to say and the README is the right place for the answer.

Out of scope is generating the reference from the flag declarations, which is what WP16
describes. `usage` is a hand written constant and the test can only hold a page against it;
making the constant itself derived is WP16's and is a larger change than this issue.

Out of scope is `mkdocs.yml` and anything else about the documentation site. WP16 is unstarted
and its generator is under question in #226, so the pages here are files a site can later
publish rather than a site.

Out of scope are the command lists in the skills. Checked rather than assumed: `CONTRIBUTING.md`
carries none, and each of the seven skills carries four to seven lines naming the commands of
its own phase, which is seven short lists and not a seventh copy of this one. The review skill
mentions `xeno gate verify` once in prose, about what the pipeline does. Each is correct
because it names only what its phase runs, and nothing a test could hold would make it more so.
Left alone deliberately.

Out of scope is shortening `usage` itself. It is the reference's source and a reader of
`--help` wants all of it.

<!-- xeno:section:context-rationale -->
## Why this context

Seven files, 350879 bytes.

`README.md` is the file being rewritten, and is read whole first: the issue says the existing
material stays further down, the maintainer's comment says most of it moves, and deciding
between those needs the file rather than a diff of it.

`docs/process-definition.md` carries the two passages the new opening rests on: section 1's
purpose, which is where "produced separately, at different times" and the ISO/IEC 42001
sentence come from, and Appendix C, which is the only place the name is explained. It also
carries section 12's "Two surfaces, two promises", which is why the no-network-call paragraph
does not need a new home.

`docs/implementation-plan.md` is read for WP16, which says the reference is generated from the
declarations and is what this intent deliberately does not do, and for the sentence declining
to record the name, which is what makes the README the only place the answer can live.

`cmd/xeno/main.go` holds `usage`, which is the source the reference page is held against and
the thing the README's list has drifted from. Read rather than recalled: the drift was measured
at six commands in `usage` and absent from the README, none the other way, 27 against 21.

`cmd/xeno/main_test.go` holds `TestEveryCommandInTheUsageResolves`, which reads the command
names out of `usage` with a regexp and says why a third copy would be a third place to be
wrong. It is the pattern the new test follows and sits beside.

`CONTRIBUTING.md` is read to confirm it carries no command list of its own, which it does not,
so nothing there has to move or be held.

`CLAUDE.md` carries the standing rules: the first, which is why the normative documents are
pointed at rather than moved, and the one about replacing a paragraph rather than editing into
it, which for a file being restructured means writing it rather than patching it.

The links block declares `cmd/xeno` against the implementation plan, because the question of
who owns the command reference is WP16's and the answer here is deliberately smaller than what
WP16 describes.

Not in scope: `internal/`, because nothing in the runner changes, and `docs/assumptions.md`,
because nothing here is an assumption taken while building — the one judgement, how much of the
README moves, was put to the maintainer and answered on the issue.
